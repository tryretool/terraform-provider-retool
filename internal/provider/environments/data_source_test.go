package environments_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/tryretool/terraform-provider-retool/internal/acctest"
)

func checkEnvironmentsListIsNotEmpty(state *terraform.State) error {
	environments := state.RootModule().Resources["data.retool_environments.test_environments"]
	if environments == nil {
		return fmt.Errorf("environments not found in state")
	}

	if len(environments.Primary.Attributes) == 0 {
		return fmt.Errorf("empty resource")
	}
	numEnvironments, err := strconv.Atoi(environments.Primary.Attributes["environments.#"])
	if err != nil || numEnvironments == 0 {
		return fmt.Errorf("no environments found")
	}

	return nil
}

const testEnvironmentDataSourceConfig = `
	data "retool_environments" "test_environments" {}
`

func TestAccEnvironmentsDataSource(t *testing.T) {
	acctest.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: testEnvironmentDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					checkEnvironmentsListIsNotEmpty,
					resource.TestCheckResourceAttrSet("data.retool_environments.test_environments", "environments.0.id"),
					resource.TestCheckResourceAttrSet("data.retool_environments.test_environments", "environments.0.name"),
					resource.TestCheckResourceAttrSet("data.retool_environments.test_environments", "environments.0.description"),
					resource.TestCheckResourceAttrSet("data.retool_environments.test_environments", "environments.0.default"),
					resource.TestCheckResourceAttrSet("data.retool_environments.test_environments", "environments.0.created_at"),
					resource.TestCheckResourceAttrSet("data.retool_environments.test_environments", "environments.0.updated_at"),
				),
			},
		},
	})
}
