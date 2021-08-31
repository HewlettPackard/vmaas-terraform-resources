// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package acceptancetest

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/HewlettPackard/hpegl-vmaas-terraform-resources/pkg/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	libUtils "github.com/HewlettPackard/hpegl-provider-lib/pkg/utils"
)

var (
	testAccProviders map[string]*schema.Provider
	testAccProvider  *schema.Provider
)

func init() {
	testAccProvider = utils.ProviderFunc()()
	testAccProviders = map[string]*schema.Provider{
		"hpegl": testAccProvider,
	}
}

func testAccPreCheck(t *testing.T) {
	// validate all required envs are present, if not then throws error
	requiredenvs := []string{"CMP_SUBJECT"}
	for _, r := range requiredenvs {
		if os.Getenv(r) == "" {
			panic(r + " env is required, but not found")
		}
	}

	t.Helper()
}

func TestProvider(t *testing.T) {
	if err := utils.ProviderFunc()().InternalValidate(); err != nil {
		t.Fatalf("%s\n", err)
	}
	testAccPreCheck(t)
}

func TestMain(m *testing.M) {
	// nolint
	_, b, _, _ := runtime.Caller(0)
	// Root folder of this project
	d := filepath.Join(filepath.Dir(b), "../..")
	libUtils.ReadAccConfig(d)
	m.Run()
	os.Exit(0)
}
