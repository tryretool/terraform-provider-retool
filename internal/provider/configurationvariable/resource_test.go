package configurationvariable_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/tryretool/terraform-provider-retool/internal/acctest"
)

const testConfigrationVariableResourceBasic = `
resource "retool_configuration_variable" "test_configuration_variable" {
  name        = "tf-acc-test-configuration-variable"
  description = "Terraform acceptance test configuration variable"
  secret      = false
  values = [
	{
	  environment_id = "env_123"
	  value          = "value1"
	},
	{
	  environment_id = "env_456"
	  value          = "value2"
	}
  ]
}
`

const testConfigrationVariableResourceUpdated = `
resource "retool_configuration_variable" "test_configuration_variable" {
  name        = "tf-acc-test-configuration-variable-modified"
  description = "Terraform acceptance test configuration variable modified"
  secret      = true
  values = [
	{
	  environment_id = "env_123"
	  value          = "new_value1"
	},
	{
	  environment_id = "env_456"
	  value          = "new_value2"
	}
  ]
}
`

func testMain(m *testing.M) {
	resource.TestMain(m)
}

func TestAccConfigurationVariable(t *testing.T) {
	acctest.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			// Read and Create.
			{
				Config: testConfigrationVariableResourceBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_configuration_variable.test_configuration_variable", "name", "tf-acc-test-configuration-variable"),
					resource.TestCheckResourceAttr("retool_configuration_variable.test_configuration_variable", "description", "Terraform acceptance test configuration variable"),
					resource.TestCheckResourceAttr("retool_configuration_variable.test_configuration_variable", "secret", "false"),
					resource.TestCheckResourceAttrSet("retool_configuration_variable.test_configuration_variable", "id"),
					resource.TestCheckResourceAttr("retool_configuration_variable.test_configuration_variable", "values.#", "2"),
					resource.TestCheckResourceAttrSet("retool_configuration_variable.test_configuration_variable", "values.0.environment_id"),
					resource.TestCheckResourceAttr("retool_configuration_variable.test_configuration_variable", "values.0.value", "value1"),
					resource.TestCheckResourceAttrSet("retool_configuration_variable.test_configuration_variable", "values.1.environment_id"),
					resource.TestCheckResourceAttr("retool_configuration_variable.test_configuration_variable", "values.1.value", "value2"),
				),
			},
			// Import state.
			{
				ResourceName:      "retool_configuration_variable.test_configuration_variable",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read.
			{
				Config: testConfigrationVariableResourceUpdated,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_configuration_variable.test_configuration_variable", "name", "tf-acc-test-configuration-variable-modified"),
					resource.TestCheckResourceAttr("retool_configuration_variable.test_configuration_variable", "description", "Terraform acceptance test configuration variable modified"),
					resource.TestCheckResourceAttr("retool_configuration_variable.test_configuration_variable", "secret", "true"),
					resource.TestCheckResourceAttrSet("retool_configuration_variable.test_configuration_variable", "id"),
					resource.TestCheckResourceAttr("retool_configuration_variable.test_configuration_variable", "values.#", "2"),
					resource.TestCheckResourceAttrSet("retool_configuration_variable.test_configuration_variable", "values.0.environment_id"),
					resource.TestCheckResourceAttr("retool_configuration_variable.test_configuration_variable", "values.0.value", "new_value1"),
					resource.TestCheckResourceAttrSet("retool_configuration_variable.test_configuration_variable", "values.1.environment_id"),
					resource.TestCheckResourceAttr("retool_configuration_variable.test_configuration_variable", "values.1.value", "new_value2"),
				),
			},
		},
	})
}

func sweepConfigurationVariables(region string) error {
	client, err := acctest.SweeperClient()
	if err != nil {
		return fmt.Errorf("could not create Retool API client: %w", err)
	}

	configuration_variables, _, err := client.ConfigurationVariablesAPI.ConfigurationVariablesGet(context.Background()).Execute()
	if err != nil {
		return fmt.Errorf("error reading configuration variables: %s", err.Error())
	}

	for _, configuration_variable := range configuration_variables.Data {
		if len(configuration_variable.Name) >= 16 && configuration_variable.Name[:16] == "tf-acc-test-" {
			_, err := client.ConfigurationVariablesAPI.ConfigurationVariablesIdDelete(context.Background(), configuration_variable.Id).Execute()
			if err != nil {
				return fmt.Errorf("error deleting configuration variable %s: %s", configuration_variable.Id, err.Error())
			}
			log.Printf("Deleted configuration variable %s\n", configuration_variable.Id)
		}
	}
	return nil
}

func init() {
	resource.AddTestSweepers("retool_configuration_variable", &resource.Sweeper{
		Name: "retool_configuration_variable",
		F:    sweepConfigurationVariables,
	})
}
