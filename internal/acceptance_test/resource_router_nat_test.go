// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package acceptancetest

import (
	"context"
	"testing"

	api_client "github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/client"
	"github.com/HewlettPackard/hpegl-vmaas-terraform-resources/pkg/atf"
)

func TestVmaasRouterNatPlan(t *testing.T) {
	acc := &atf.Acc{
		PreCheck:     testAccPreCheck,
		Providers:    testAccProviders,
		ResourceName: "hpegl_vmaas_router_nat_rule",
	}
	acc.RunResourcePlanTest(t)
}

func TestAccResourceRouterNatCreate(t *testing.T) {
	acc := &atf.Acc{
		ResourceName: "hpegl_vmaas_router_nat_rule",
		PreCheck:     testAccPreCheck,
		Providers:    testAccProviders,
		GetApi: func(attr map[string]string) (interface{}, error) {
			cl, cfg := getAPIClient()
			iClient := api_client.RouterAPIService{
				Client: cl,
				Cfg:    cfg,
			}
			id := toInt(attr["id"])
			routerID := toInt(attr["router_id"])

			return iClient.GetSpecificRouterNat(context.Background(), routerID, id)
		},
	}

	acc.RunTests(t)
}
