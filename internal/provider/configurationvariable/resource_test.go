package configurationvariable_test

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/tryretool/terraform-provider-retool/internal/acctest"
)

const testConfigrationVariableResourceBasic = `
resource "retool_configuration_variable" "test_configuration_variable" {
  name        = "tf-acc-test-configuration-variable"
  description = "Terraform acceptance test configuration variable"
  values = [
	{
	  environment_id = "23ba02f3-12c2-456e-826f-13b616daa5ce"
	  value          = "value1"
	},
	{
	  environment_id = "fc4a3c5b-6328-4990-b368-6879a1f38855"
	  value          = "value2"
	}
  ]
}
`

const testConfigrationVariableResourceUpdated = `
resource "retool_configuration_variable" "test_configuration_variable" {
  name        = "tf-acc-test-configuration-variable-modified"
  description = "Terraform acceptance test configuration variable modified"
  values = [
	{
	  environment_id = "23ba02f3-12c2-456e-826f-13b616daa5ce"
	  value          = "new_value1"
	},
	{
	  environment_id = "fc4a3c5b-6328-4990-b368-6879a1f38855"
	  value          = "new_value2"
	}
  ]
}
`

const testConfigurationVariableWithInvalidEnvironmentID = `
resource "retool_configuration_variable" "test_configuration_variable" {
  name        = "tf-acc-test-configuration-variable-invalid-env-id"
  description = "Terraform acceptance test configuration variable with invalid environment ID"
  values = [
	{
	  environment_id = "invalid-environment-id"
	  value          = "value1"
	}
  ]
}
`

const testConfigurationVariableSecret = `
resource "retool_configuration_variable" "test_secret_configuration_variable" {
  name        = "tf-acc-test-secret-configuration-variable"
  description = "Terraform acceptance test secret configuration variable"
  secret      = true
  values = [
	{
	  environment_id = "23ba02f3-12c2-456e-826f-13b616daa5ce"
	  value          = "secret_value1"
	},
	{
	  environment_id = "fc4a3c5b-6328-4990-b368-6879a1f38855"
	  value          = "secret_value2"
	}
  ]
}
`

func TestMain(m *testing.M) {
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
					resource.TestCheckResourceAttr("retool_configuration_variable.test_configuration_variable", "values.#", "2"),
					resource.TestCheckResourceAttrSet("retool_configuration_variable.test_configuration_variable", "values.0.environment_id"),
					resource.TestCheckResourceAttrSet("retool_configuration_variable.test_configuration_variable", "values.0.value"),
					resource.TestCheckResourceAttr("retool_configuration_variable.test_configuration_variable", "values.0.value", "value1"),
					resource.TestCheckResourceAttrSet("retool_configuration_variable.test_configuration_variable", "values.1.environment_id"),
					resource.TestCheckResourceAttrSet("retool_configuration_variable.test_configuration_variable", "values.1.value"),
					resource.TestCheckResourceAttr("retool_configuration_variable.test_configuration_variable", "values.1.value", "value2"),
				),
			},
			// Import state.
			{
				ResourceName:      "retool_configuration_variable.test_configuration_variable",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update.
			{
				Config: testConfigrationVariableResourceUpdated,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_configuration_variable.test_configuration_variable", "name", "tf-acc-test-configuration-variable-modified"),
					resource.TestCheckResourceAttr("retool_configuration_variable.test_configuration_variable", "description", "Terraform acceptance test configuration variable modified"),
					resource.TestCheckResourceAttr("retool_configuration_variable.test_configuration_variable", "secret", "false"),
					resource.TestCheckResourceAttrSet("retool_configuration_variable.test_configuration_variable", "id"),
					resource.TestCheckResourceAttr("retool_configuration_variable.test_configuration_variable", "values.#", "2"),
					resource.TestCheckResourceAttrSet("retool_configuration_variable.test_configuration_variable", "values.0.environment_id"),
					resource.TestCheckResourceAttrSet("retool_configuration_variable.test_configuration_variable", "values.0.value"),
					resource.TestCheckResourceAttr("retool_configuration_variable.test_configuration_variable", "values.0.value", "new_value1"),
					resource.TestCheckResourceAttrSet("retool_configuration_variable.test_configuration_variable", "values.1.environment_id"),
					resource.TestCheckResourceAttrSet("retool_configuration_variable.test_configuration_variable", "values.1.value"),
					resource.TestCheckResourceAttr("retool_configuration_variable.test_configuration_variable", "values.1.value", "new_value2"),
				),
			},
			// Invalid environment ID.
			{
				Config:      testConfigurationVariableWithInvalidEnvironmentID,
				ExpectError: regexp.MustCompile(`400 Bad Request`),
			},
		},
	})
}

func sweepConfigurationVariables(region string) error {
	log.Printf("Sweeping configuration variables in region %s", region)
	client, err := acctest.SweeperClient()
	if err != nil {
		return fmt.Errorf("could not create Retool API client: %w", err)
	}

	configurationVariables, _, err := client.ConfigurationVariablesAPI.ConfigurationVariablesGet(context.Background()).Execute()
	if err != nil {
		return fmt.Errorf("error reading configuration variables: %s", err.Error())
	}

	for _, configurationVariable := range configurationVariables.Data {
		if len(configurationVariable.Name) >= 16 && configurationVariable.Name[:16] == "tf-acc-test-" {
			log.Printf("Deleting configuration variable %s (%s)\n", configurationVariable.Name, configurationVariable.Id)
			_, err := client.ConfigurationVariablesAPI.ConfigurationVariablesIdDelete(context.Background(), configurationVariable.Id).Execute()
			if err != nil {
				return fmt.Errorf("error deleting configuration variable %s: %s", configurationVariable.Id, err.Error())
			}
			log.Printf("Deleted configuration variable %s\n", configurationVariable.Id)
		}
	}
	return nil
}

func TestAccConfigurationVariableSecret(t *testing.T) {
	acctest.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			// Create a secret configuration variable.
			{
				Config: testConfigurationVariableSecret,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_configuration_variable.test_secret_configuration_variable", "name", "tf-acc-test-secret-configuration-variable"),
					resource.TestCheckResourceAttr("retool_configuration_variable.test_secret_configuration_variable", "description", "Terraform acceptance test secret configuration variable"),
					resource.TestCheckResourceAttr("retool_configuration_variable.test_secret_configuration_variable", "secret", "true"),
					resource.TestCheckResourceAttrSet("retool_configuration_variable.test_secret_configuration_variable", "id"),
					resource.TestCheckResourceAttr("retool_configuration_variable.test_secret_configuration_variable", "values.#", "2"),
				),
			},
			// Import state - secret values are set to null on read, so we ignore them in verification.
			{
				ResourceName:            "retool_configuration_variable.test_secret_configuration_variable",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"values.0.value", "values.1.value"}, // Secret values are set to null on read
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("retool_configuration_variable", &resource.Sweeper{
		Name: "retool_configuration_variable",
		F:    sweepConfigurationVariables,
	})
}
