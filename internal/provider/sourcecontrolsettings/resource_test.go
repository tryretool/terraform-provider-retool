package sourcecontrolsettings_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/tryretool/terraform-provider-retool/internal/acctest"
	"github.com/tryretool/terraform-provider-retool/internal/sdk/api"
)

const testSCMSettings = `
	resource "retool_source_control_settings" "scm_settings" {
		auto_branch_naming_enabled = false
		custom_pull_request_template_enabled = true
		custom_pull_request_template = "custom-pull-request-template"
		version_control_locked = true
		force_uuid_mapping = true
	}
`

const testSCMSettingsUpdated = `
	resource "retool_source_control_settings" "scm_settings" {
		auto_branch_naming_enabled = true
		custom_pull_request_template_enabled = false
		custom_pull_request_template = "custom-pull-request-template-updated"
		version_control_locked = false
		force_uuid_mapping = false
	}
`

const testSCMSettingsDefaults = `
	resource "retool_source_control_settings" "scm_settings" {
	}
`

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func TestAccSourceControlSettings(t *testing.T) {
	acctest.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			// Read and Create.
			{
				Config: testSCMSettings,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_source_control_settings.scm_settings", "auto_branch_naming_enabled", "false"),
					resource.TestCheckResourceAttr("retool_source_control_settings.scm_settings", "custom_pull_request_template_enabled", "true"),
					resource.TestCheckResourceAttr("retool_source_control_settings.scm_settings", "custom_pull_request_template", "custom-pull-request-template"),
					resource.TestCheckResourceAttr("retool_source_control_settings.scm_settings", "version_control_locked", "true"),
					resource.TestCheckResourceAttr("retool_source_control_settings.scm_settings", "force_uuid_mapping", "true"),
				),
			},
			// Import state.
			{
				ResourceName:                         "retool_source_control_settings.scm_settings",
				ImportState:                          true,
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "version_control_locked",
			},
			// Update and Read.
			{
				Config: testSCMSettingsUpdated,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_source_control_settings.scm_settings", "auto_branch_naming_enabled", "true"),
					resource.TestCheckResourceAttr("retool_source_control_settings.scm_settings", "custom_pull_request_template_enabled", "false"),
					resource.TestCheckResourceAttr("retool_source_control_settings.scm_settings", "custom_pull_request_template", "custom-pull-request-template-updated"),
					resource.TestCheckResourceAttr("retool_source_control_settings.scm_settings", "version_control_locked", "false"),
					resource.TestCheckResourceAttr("retool_source_control_settings.scm_settings", "force_uuid_mapping", "false"),
				),
			},
			// Use default values.
			{
				Config: testSCMSettingsDefaults,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_source_control_settings.scm_settings", "auto_branch_naming_enabled", "true"),
					resource.TestCheckResourceAttr("retool_source_control_settings.scm_settings", "custom_pull_request_template_enabled", "false"),
					resource.TestCheckResourceAttr("retool_source_control_settings.scm_settings", "custom_pull_request_template", ""),
					resource.TestCheckResourceAttr("retool_source_control_settings.scm_settings", "version_control_locked", "false"),
					resource.TestCheckResourceAttr("retool_source_control_settings.scm_settings", "force_uuid_mapping", "false"),
				),
			},
		},
	})
}

func sweepSourceControlSettings(region string) error {
	log.Printf("Sweeping Source Control settings in region %s", region)
	client, err := acctest.SweeperClient()
	if err != nil {
		return err
	}

	falseValue := false
	trueValue := true
	emptyString := ""
	apiRequest := api.SourceControlSettingsPutRequest{
		Settings: api.SourceControlSettingsPutRequestSettings{
			AutoBranchNamingEnabled:          &trueValue,
			CustomPullRequestTemplateEnabled: &falseValue,
			CustomPullRequestTemplate:        &emptyString,
			VersionControlLocked:             &falseValue,
		},
	}
	_, _, err = client.SourceControlAPI.SourceControlSettingsPut(context.Background()).SourceControlSettingsPutRequest(apiRequest).Execute()
	if err != nil {
		return fmt.Errorf("Error resetting Source Control settings: %s", err.Error())
	}
	return nil
}

func init() {
	resource.AddTestSweepers("retool_source_control_settings", &resource.Sweeper{
		Name: "retool_source_control_settings",
		F:    sweepSourceControlSettings,
	})
}
