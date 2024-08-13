package space_test

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/tryretool/terraform-provider-retool/internal/acctest"
)

const testSpaceConfig = `
	resource "retool_space" "test_space" {
		name = "tf-acc-test-space"
		domain = "tfspace.localhost"
		create_options = {
			copy_sso_settings = true
			copy_branding_and_themes_settings = true
			create_admin_user = false
			users_to_copy_as_admins = ["admin@example.com"]
		}
	}
		`

const testUpdatedSpaceConfig = `
	resource "retool_space" "test_space" {
		name = "tf-acc-test-space-updated"
		domain = "tfspaceupd.localhost"
		create_options = {
			copy_sso_settings = true
			copy_branding_and_themes_settings = true
			create_admin_user = false
			users_to_copy_as_admins = ["admin@example.com"]
		}
	}
		`

const testDefaultValuesConfig = `
	resource "retool_space" "test_space_with_defaults" {
		name = "tf-acc-test-space-with-defaults"
		domain = "tfspacedefault.localhost"
	}
		`

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func TestAccSpace(t *testing.T) {
	acctest.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			// Read and Create.
			{
				Config: testSpaceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_space.test_space", "name", "tf-acc-test-space"),
					resource.TestCheckResourceAttr("retool_space.test_space", "domain", "tfspace.localhost"),
					resource.TestCheckResourceAttr("retool_space.test_space", "create_options.copy_sso_settings", "true"),
					resource.TestCheckResourceAttr("retool_space.test_space", "create_options.copy_branding_and_themes_settings", "true"),
					resource.TestCheckResourceAttr("retool_space.test_space", "create_options.create_admin_user", "false"),
					resource.TestCheckResourceAttr("retool_space.test_space", "create_options.users_to_copy_as_admins.#", "1"),
					resource.TestCheckResourceAttr("retool_space.test_space", "create_options.users_to_copy_as_admins.0", "admin@example.com"),
				),
			},
			// Import state.
			{
				ResourceName: "retool_space.test_space",
				ImportState:  true,
			},
			// Update and Read.
			{
				Config: testUpdatedSpaceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_space.test_space", "name", "tf-acc-test-space-updated"),
					resource.TestCheckResourceAttr("retool_space.test_space", "domain", "tfspaceupd.localhost"),
					resource.TestCheckResourceAttr("retool_space.test_space", "create_options.copy_sso_settings", "true"),
					resource.TestCheckResourceAttr("retool_space.test_space", "create_options.copy_branding_and_themes_settings", "true"),
					resource.TestCheckResourceAttr("retool_space.test_space", "create_options.create_admin_user", "false"),
					resource.TestCheckResourceAttr("retool_space.test_space", "create_options.users_to_copy_as_admins.#", "1"),
					resource.TestCheckResourceAttr("retool_space.test_space", "create_options.users_to_copy_as_admins.0", "admin@example.com"),
				),
			},
			// Check default values.
			{
				Config: testDefaultValuesConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_space.test_space_with_defaults", "name", "tf-acc-test-space-with-defaults"),
					resource.TestCheckResourceAttr("retool_space.test_space_with_defaults", "domain", "tfspacedefault.localhost"),
					resource.TestCheckResourceAttr("retool_space.test_space_with_defaults", "create_options.copy_sso_settings", "false"),
					resource.TestCheckResourceAttr("retool_space.test_space_with_defaults", "create_options.copy_branding_and_themes_settings", "false"),
					resource.TestCheckResourceAttr("retool_space.test_space_with_defaults", "create_options.create_admin_user", "true"),
					resource.TestCheckResourceAttr("retool_space.test_space_with_defaults", "create_options.users_to_copy_as_admins.#", "0"),
				),
			},
		},
	})
}

func sweepSpaces(region string) error {
	log.Printf("Sweeping spaces in region %s", region)
	client, err := acctest.SweeperClient()
	if err != nil {
		return err
	}

	spaces, _, err := client.SpacesAPI.SpacesGet(context.Background()).Execute()
	if err != nil {
		return fmt.Errorf("Error reading spaces: %s", err.Error())
	}

	for _, space := range spaces.Data {
		if strings.HasPrefix(space.Name, "tf-acc") {
			log.Printf("Deleting space %s", space.Name)
			_, err := client.SpacesAPI.SpacesSpaceIdDelete(context.Background(), space.Id).Execute()
			if err != nil {
				return fmt.Errorf("Error deleting group %s: %s", space.Name, err.Error())
			}
		}
	}
	return nil
}

func init() {
	resource.AddTestSweepers("retool_space", &resource.Sweeper{
		Name: "retool_space",
		F:    sweepSpaces,
	})
}
