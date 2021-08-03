// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package cmp

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hpe-hcss/vmaas-cmp-go-sdk/pkg/client"
	"github.com/hpe-hcss/vmaas-cmp-go-sdk/pkg/models"
	"github.com/hpe-hcss/vmaas-terraform-resources/internal/utils"
)

const (
	instanceCloneRetryDelay   = time.Second * 60
	instanceCloneRetryTimeout = time.Second * 30
	instanceCloneRetryCount   = 20
)

// instanceClone implements functions related to cmp instanceClones
type instanceClone struct {
	// expose Instance API service to instanceClones related operations
	iClient *client.InstancesAPIService
}

func (i *instanceClone) getIClient() *client.InstancesAPIService {
	return i.iClient
}

func newInstanceClone(iClient *client.InstancesAPIService) *instanceClone {
	return &instanceClone{
		iClient: iClient,
	}
}

// Create instanceClone
func (i *instanceClone) Create(ctx context.Context, d *utils.Data, meta interface{}) error {
	log.Printf("[INFO] Cloning instance")

	req := &models.CreateInstanceBody{
		CloneName: d.GetString("name"),
		ZoneID:    d.GetJSONNumber("cloud_id"),
		Instance: &models.CreateInstanceBodyInstance{
			InstanceType: &models.CreateInstanceBodyInstanceInstanceType{
				Code: d.GetString("instance_type_code"),
			},
			Plan: &models.CreateInstanceBodyInstancePlan{
				ID: d.GetJSONNumber("plan_id"),
			},
			Site: &models.CreateInstanceBodyInstanceSite{
				ID: d.GetInt("group_id"),
			},
			HostName:          d.GetString("hostname"),
			EnvironmentPrefix: d.GetString("env_prefix"),
		},
		Environment:       d.GetString("environment_code"),
		Evars:             instanceGetEvars(d.GetMap("evars")),
		Labels:            d.GetStringList("labels"),
		NetworkInterfaces: instanceGetNetwork(d.GetListMap("network")),
		Tags:              instanceGetTags(d.GetMap("tags")),
		LayoutSize:        d.GetInt("scale"),
		PowerScheduleType: utils.JSONNumber(d.GetInt("power_schedule_id")),
	}

	c := d.GetListMap("config")
	if len(c) > 0 {
		req.Config = instanceGetConfig(c[0])

		// Get template id instance type is vmware
		if strings.ToLower(req.Instance.InstanceType.Code) == vmware {
			templateID := c[0]["template_id"]
			if templateID == nil {
				return errors.New("error, template id is required for vmware instance type")
			}
			req.Config.Template = templateID.(int)
		}
	}
	// Pre check
	if err := d.Error(); err != nil {
		return err
	}

	// Get source instance
	sourceID := d.GetInt("source_instance_id")
	err := copyInstanceAttribsToClone(ctx, i, meta, req, d.GetListMap("volume"), sourceID)
	if err != nil {
		return err
	}

	// clone the instance
	log.Printf("[INFO] Cloning the instance with %d", sourceID)
	err = cloneInstance(ctx, i, meta, req, sourceID)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Check history")
	err = checkInstanceCloneHistory(ctx, i, meta, sourceID)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Get all instances")
	getInstanceRetry := &utils.CustomRetry{
		Delay:        instanceCloneRetryDelay,
		RetryTimeout: instanceCloneRetryTimeout,
		RetryCount:   instanceCloneRetryCount,
		Cond: func(resp interface{}, err error) (bool, error) {
			if err != nil {
				return false, nil
			}
			instancesList := resp.(models.Instances)

			return len(instancesList.Instances) == 1, nil
		},
	}
	// get cloned instance ID
	instancesResp, err := getInstanceRetry.Retry(ctx, meta, func(ctx context.Context) (interface{}, error) {
		return i.iClient.GetAllInstances(ctx, map[string]string{
			nameKey: req.CloneName,
		})
	})
	if err != nil {
		return err
	}

	instancesList := instancesResp.(models.Instances)
	if len(instancesList.Instances) != 1 {
		return errors.New("get cloned instance is failed")
	}

	if err := instanceWaitUntilCreated(ctx, i, meta, instancesList.Instances[0].ID); err != nil {
		return err
	}

	if snapshot := d.GetListMap("snapshot"); len(snapshot) == 1 {
		err := createInstanceSnapshot(ctx, i, meta, instancesList.Instances[0].ID, models.SnapshotBody{
			Snapshot: &models.SnapshotBodySnapshot{
				Name:        snapshot[0]["name"].(string),
				Description: snapshot[0]["description"].(string),
			},
		})
		if err != nil {
			return err
		}
	}

	d.SetID(instancesList.Instances[0].ID)

	// post check
	return d.Error()
}

// Update instance including poweroff, powerOn, restart, suspend
// changing volumes and instance properties such as labels
// groups and tags
func (i *instanceClone) Update(ctx context.Context, d *utils.Data, meta interface{}) error {
	return updateInstance(ctx, i, d, meta)
}

// Delete instance and set ID as ""
func (i *instanceClone) Delete(ctx context.Context, d *utils.Data, meta interface{}) error {
	return deleteInstance(ctx, i, d, meta)
}

// Read instance and set state values accordingly
func (i *instanceClone) Read(ctx context.Context, d *utils.Data, meta interface{}) error {
	id := d.GetID()

	log.Printf("[INFO] Get instance with ID %d", id)

	// Precheck
	if err := d.Error(); err != nil {
		return err
	}

	resp, err := utils.Retry(ctx, meta, func(ctx context.Context) (interface{}, error) {
		return i.iClient.GetASpecificInstance(ctx, id)
	})
	if err != nil {
		return err
	}
	instance := resp.(models.GetInstanceResponse)

	volumes := d.GetListMap("volume")
	// Assign proper ID for the volume, since response may contains more
	// volumes than schema, check the name and assign ip
	for i := range volumes {
		for _, vModel := range instance.Instance.Volumes {
			if vModel.Name == volumes[i]["name"].(string) {
				volumes[i]["id"] = vModel.ID
			}
		}
	}

	d.Set("volume", volumes)

	// Write IPs in to state file
	instanceSetIP(d, instance)
	instanceSetHostname(d, instance)

	d.Set("layout_id", instance.Instance.Layout.ID)
	d.SetString("status", instance.Instance.Status)
	d.SetID(instance.Instance.ID)

	// post check
	return d.Error()
}

func checkInstanceCloneHistory(ctx context.Context, i *instanceClone, meta interface{}, instanceID int) error {
	errCount := 0
	historyRetry := utils.CustomRetry{
		Delay:        instanceCloneRetryDelay,
		RetryTimeout: instanceCloneRetryTimeout,
		RetryCount:   instanceCloneRetryCount,
		Cond: func(response interface{}, ResponseErr error) (bool, error) {
			if ResponseErr != nil {
				errCount++
				if errCount == 3 {
					return false, ResponseErr
				}

				return false, nil
			}
			errCount = 0

			instanceHistory := response.(models.GetInstanceHistory)
			if len(instanceHistory.Processes) > 0 && instanceHistory.Processes[0].ProcessType.Code == "cloning" {
				if instanceHistory.Processes[0].Status == "failed" {
					return false, fmt.Errorf("failed to clone instance")
				} else if instanceHistory.Processes[0].Status == "success" {
					return true, nil
				}
			}

			return false, nil
		},
	}
	_, err := historyRetry.Retry(ctx, meta, func(ctx context.Context) (interface{}, error) {
		return i.iClient.GetInstanceHistory(ctx, instanceID)
	})

	return err
}

func cloneInstance(ctx context.Context, i *instanceClone, meta interface{}, req *models.CreateInstanceBody, sourceID int) error {
	cloneRetry := &utils.CustomRetry{
		Cond: func(response interface{}, ResponseErr error) (bool, error) {
			if ResponseErr != nil {
				return false, nil
			}
			if !response.(models.SuccessOrErrorMessage).Success {
				return false, fmt.Errorf("failed to clone instance")
			}

			return true, nil
		},
	}
	_, err := cloneRetry.Retry(ctx, meta, func(ctx context.Context) (interface{}, error) {
		return i.iClient.CloneAnInstance(ctx, sourceID, req)
	})
	return err
}

func copyInstanceAttribsToClone(
	ctx context.Context,
	i *instanceClone,
	meta interface{},
	req *models.CreateInstanceBody,
	volumes []map[string]interface{},
	sourceID int,
) error {
	sourceInstanceResp, err := utils.Retry(ctx, meta, func(ctx context.Context) (interface{}, error) {
		return i.iClient.GetASpecificInstance(ctx, sourceID)
	})
	if err != nil {
		return err
	}
	sourceInstance := sourceInstanceResp.(models.GetInstanceResponse)

	if utils.IsEmpty(req.ZoneID) {
		req.ZoneID = utils.JSONNumber(sourceInstance.Instance.Cloud.ID)
	}
	if utils.IsEmpty(req.Instance.Plan.ID) {
		req.Instance.Plan.ID = utils.JSONNumber(sourceInstance.Instance.Plan.ID)
	}
	if req.Instance.InstanceType.Code == "" {
		req.Instance.InstanceType.Code = sourceInstance.Instance.InstanceType.Code
	}
	if req.Instance.Site.ID == 0 {
		req.Instance.Site.ID = sourceInstance.Instance.Group.ID
	}
	if req.LayoutSize == 0 {
		req.LayoutSize = sourceInstance.Instance.Config.Layoutsize
	}
	if req.Labels == nil {
		req.Labels = sourceInstance.Instance.Labels
	}
	if req.Tags == nil {
		req.Tags = sourceInstance.Instance.Tags
	}

	req.Volumes = instanceCloneCompareVolume(volumes, sourceInstance.Instance.Volumes)
	req.Instance.Layout = &models.CreateInstanceBodyInstanceLayout{
		ID: utils.JSONNumber(sourceInstance.Instance.Layout.ID),
	}

	return nil
}
