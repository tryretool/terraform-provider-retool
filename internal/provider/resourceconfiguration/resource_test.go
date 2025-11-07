package resourceconfiguration_test

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/tryretool/terraform-provider-retool/internal/acctest"
)

// Note: These tests require an existing resource and environment to configure.
// We'll create a test resource first, then create a configuration for it.

const testResourceConfigurationConfig = `
	# First create a test resource
	resource "retool_resource" "test_resource" {
		display_name = "Test REST API for Config"
		type         = "restapi"
		options      = jsonencode({
			base_url = "https://api.example.com"
		})
	}

	# Get the environments
	data "retool_environments" "all" {}

	# Create a resource configuration for the first non-default environment
	resource "retool_resource_configuration" "test_config" {
		resource_id    = retool_resource.test_resource.id
		environment_id = [for env in data.retool_environments.all.environments : env.id if !env.default][0]
		options        = jsonencode({
			base_url = "https://api.production.example.com"
		})
	}
`

const testResourceConfigurationUpdatedConfig = `
	# First create a test resource
	resource "retool_resource" "test_resource" {
		display_name = "Test REST API for Config"
		type         = "restapi"
		options      = jsonencode({
			base_url = "https://api.example.com"
		})
	}

	# Get the environments
	data "retool_environments" "all" {}

	# Create a resource configuration for the first non-default environment with updated options
	resource "retool_resource_configuration" "test_config" {
		resource_id    = retool_resource.test_resource.id
		environment_id = [for env in data.retool_environments.all.environments : env.id if !env.default][0]
		options        = jsonencode({
			base_url = "https://api.production-v2.example.com"
		})
	}
`

func TestAccResourceConfiguration(t *testing.T) {
	acctest.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			// Create and Read.
			{
				Config: testResourceConfigurationConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("retool_resource_configuration.test_config", "id"),
					resource.TestCheckResourceAttrSet("retool_resource_configuration.test_config", "resource_id"),
					resource.TestCheckResourceAttrSet("retool_resource_configuration.test_config", "environment_id"),
					resource.TestCheckResourceAttrSet("retool_resource_configuration.test_config", "options"),
					resource.TestCheckResourceAttrSet("retool_resource_configuration.test_config", "created_at"),
					resource.TestCheckResourceAttrSet("retool_resource_configuration.test_config", "updated_at"),
				),
				// The API adds default values to options, which may cause a refresh plan
				ExpectNonEmptyPlan: true,
			},
			// Import state.
			{
				ResourceName:      "retool_resource_configuration.test_config",
				ImportState:       true,
				ImportStateVerify: true,
				// The API adds default values to options, so we can't verify exact match
				ImportStateVerifyIgnore: []string{"options"},
			},
			// Update and Read.
			{
				Config: testResourceConfigurationUpdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("retool_resource_configuration.test_config", "id"),
					resource.TestCheckResourceAttrSet("retool_resource_configuration.test_config", "resource_id"),
					resource.TestCheckResourceAttrSet("retool_resource_configuration.test_config", "environment_id"),
				),
				// The API adds default values to options, which may cause a refresh plan
				ExpectNonEmptyPlan: true,
			},
		},
		// Note: The test framework will automatically call Delete on all resources after the last step.
	})
}

func sweepResourceConfigurations(region string) error {
	log.Printf("Sweeping resource configurations in region %s", region)
	client, err := acctest.SweeperClient()
	if err != nil {
		return err
	}

	configurations, _, err := client.ResourceConfigurationsAPI.ResourceConfigurationsGet(context.Background()).Execute()
	if err != nil {
		return fmt.Errorf("Error reading resource configurations: %s", err.Error())
	}

	for _, config := range configurations.Data {
		// Check if the resource associated with this config is a test resource
		if strings.Contains(config.Resource.DisplayName, "Test ") || strings.Contains(config.Resource.DisplayName, "test") {
			log.Printf("Deleting resource configuration %s (resource: %s)", config.Id, config.Resource.DisplayName)
			_, err := client.ResourceConfigurationsAPI.ResourceConfigurationsConfigurationIdDelete(context.Background(), config.Id).Execute()
			if err != nil {
				return fmt.Errorf("Error deleting resource configuration %s: %s", config.Id, err.Error())
			}
		}
	}
	return nil
}

func init() {
	resource.AddTestSweepers("retool_resource_configuration", &resource.Sweeper{
		Name: "retool_resource_configuration",
		F:    sweepResourceConfigurations,
		Dependencies: []string{
			"retool_resource", // Clean up configurations before resources
		},
	})
}
