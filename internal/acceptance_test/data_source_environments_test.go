// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package acceptancetest

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceEnvironment(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		IsUnitTest: false,
		PreCheck:   func() { testAccPreCheck(t) },
		Providers:  testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceEnvironmentConfig(),
				Check: resource.ComposeTestCheckFunc(
					validateDataSourceID("data.hpegl_vmaas_environment." +
						viper.GetString("vmaas.data_source_environment.environmentLocalName")),
				),
			},
		},
	})
}

func testAccDataSourceEnvironmentConfig() string {
	return providerStanza + fmt.Sprintf(`
data "hpegl_vmaas_environment" "%s" {
	name = "%s"
  }
`, viper.GetString("vmaas.data_source_environment.environmentLocalName"),
		viper.GetString("vmaas.data_source_environment.environmentName"))
}
