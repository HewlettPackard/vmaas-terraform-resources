// (C) Copyright 2022 Hewlett Packard Enterprise Development LP

package resources

import (
	"context"

	diffvalidation "github.com/HewlettPackard/hpegl-vmaas-terraform-resources/internal/resources/diffValidation"
	"github.com/HewlettPackard/hpegl-vmaas-terraform-resources/internal/resources/schemas"
	"github.com/HewlettPackard/hpegl-vmaas-terraform-resources/internal/resources/validations"
	"github.com/HewlettPackard/hpegl-vmaas-terraform-resources/internal/utils"
	"github.com/HewlettPackard/hpegl-vmaas-terraform-resources/pkg/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func LoadBalancerVirtualServers() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"lb_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Parent lb ID, lb_id can be obtained by using LB datasource/resource.",
				ForceNew:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of Network loadbalancer virtual server name",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of Network loadbalancer virtual server",
			},
			"vip_address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Vip_address of Network loadbalancer virtual server",
			},
			"vip_port": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Vip_port of network loadbalancer virtual server",
			},
			"pool": {
				Type:     schema.TypeInt,
				Required: true,
				Description: "It is recommended that you attach a pool to the Virtual Server to have a correct LB functionality." +
					"ID of the loadBalncer Pool. Use " + DSLBPool + "datasource to obtain the id here",
			},
			"type": {
				Type: schema.TypeString,
				ValidateDiagFunc: validations.StringInSlice([]string{
					"http",
					"tcp",
					"udp",
				}, false),
				Required:     true,
				InputDefault: "http",
				Description:  "Vip protocol of Network loadbalancer virtual server",
			},
			"tcp_application_profile":  schemas.TCPAppProfileSchema(),
			"udp_application_profile":  schemas.UDPAppProfileSchema(),
			"http_application_profile": schemas.HTTPAppProfileSchema(),
			"persistence": {
				Type: schema.TypeString,
				ValidateDiagFunc: validations.StringInSlice([]string{
					"SOURCE_IP",
					"COOKIE",
				}, false),
				Optional:    true,
				Description: "Persistence type for Network loadbalancer virtual server",
			},
			"cookie_persistence_profile":   schemas.CookiePersProfileSchema(),
			"sourceip_persistence_profile": schemas.SourceipPersProfileSchema(),
			"ssl_server_cert": {
				Type:     schema.TypeInt,
				Required: true,
				Description: "SSLServerCert is needed only for https based load balancer" +
					"ID of the ssl_server_cert. Use " + DSLBVirtualServerSslCert + "datasource to obtain the id  here",
			},
			"ssl_server_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "virtual server Configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ssl_server_profile": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "ID of the ssl_server_profile. Use " + DSLBProfile + "datasource to obtain the id  here",
						},
					},
				},
			},
			"ssl_client_cert": {
				Type:     schema.TypeInt,
				Required: true,
				Description: "SSLClientCert is needed only for https based load balancer." +
					"ID of the ssl_client_cert. Use " + DSLBVirtualServerSslCert + "datasource to obtain the id  here",
			},
			"ssl_client_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "virtual server Configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ssl_client_profile": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "ID of the ssl_client_profile. Use " + DSLBProfile + "datasource to obtain the id  here",
						},
					},
				},
			},
		},
		ReadContext:   loadbalancerVirtualServerReadContext,
		UpdateContext: loadbalancerVirtualServerUpdateContext,
		CreateContext: loadbalancerVirtualServerCreateContext,
		DeleteContext: loadbalancerVirtualServerDeleteContext,
		CustomizeDiff: virtualServerCustomDiff,
		Description: `loadbalancer Virtual Server resource facilitates creating,
		and deleting NSX-T  Network Load Balancers.`,
	}
}

func loadbalancerVirtualServerUpdateContext(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, err := client.GetClientFromMetaMap(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	data := utils.NewData(rd)
	if err := c.CmpClient.LoadBalancerVirtualServer.Update(ctx, data, meta); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func loadbalancerVirtualServerReadContext(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, err := client.GetClientFromMetaMap(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	data := utils.NewData(rd)
	if err := c.CmpClient.LoadBalancerVirtualServer.Read(ctx, data, meta); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func loadbalancerVirtualServerCreateContext(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, err := client.GetClientFromMetaMap(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	data := utils.NewData(rd)
	if err := c.CmpClient.LoadBalancerVirtualServer.Create(ctx, data, meta); err != nil {
		return diag.FromErr(err)
	}

	return loadbalancerVirtualServerReadContext(ctx, rd, meta)
}

func loadbalancerVirtualServerDeleteContext(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, err := client.GetClientFromMetaMap(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	data := utils.NewData(rd)
	if err := c.CmpClient.LoadBalancerVirtualServer.Delete(ctx, data, meta); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func virtualServerCustomDiff(ctx context.Context, diff *schema.ResourceDiff, meta interface{}) error {
	return diffvalidation.NewLoadBalancerVirtualServerValidate(diff).DiffValidate()
}
