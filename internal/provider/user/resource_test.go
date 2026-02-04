package user_test

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/tryretool/terraform-provider-retool/internal/acctest"
)

const testUserConfig = `
	resource "retool_user" "test_user" {
		email = "tf-acc-user-v4@example.retool.com"
		first_name = "Test"
		last_name = "User"
		active = true
		metadata = "{\"department\":\"engineering\"}"
	}
`

const testUpdatedUserConfig = `
	resource "retool_user" "test_user" {
		email = "tf-acc-user-v4@example.retool.com"
		first_name = "Updated"
		last_name = "TestUser"
		active = true
		metadata = "{\"department\":\"engineering\"}"
	}
`

const testDefaultValuesConfig = `
	resource "retool_user" "test_user_with_defaults" {
		email = "tf-acc-defaults-v4@example.retool.com"
		first_name = "Default"
		last_name = "User"
	}
`

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func TestAccUser(t *testing.T) {
	acctest.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			// Create and Read.
			{
				Config: testUserConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_user.test_user", "email", "tf-acc-user-v4@example.retool.com"),
					resource.TestCheckResourceAttr("retool_user.test_user", "first_name", "Test"),
					resource.TestCheckResourceAttr("retool_user.test_user", "last_name", "User"),
					resource.TestCheckResourceAttr("retool_user.test_user", "active", "true"),
					resource.TestCheckResourceAttr("retool_user.test_user", "metadata", "{\"department\":\"engineering\"}"),
					resource.TestCheckResourceAttrSet("retool_user.test_user", "id"),
					resource.TestCheckResourceAttrSet("retool_user.test_user", "legacy_id"),
					resource.TestCheckResourceAttrSet("retool_user.test_user", "created_at"),
					resource.TestCheckResourceAttrSet("retool_user.test_user", "is_admin"),
					resource.TestCheckResourceAttrSet("retool_user.test_user", "two_factor_auth_enabled"),
				),
			},
			// Import state.
			{
				ResourceName:      "retool_user.test_user",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read.
			{
				Config: testUpdatedUserConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_user.test_user", "email", "tf-acc-user-v4@example.retool.com"),
					resource.TestCheckResourceAttr("retool_user.test_user", "first_name", "Updated"),
					resource.TestCheckResourceAttr("retool_user.test_user", "last_name", "TestUser"),
					resource.TestCheckResourceAttr("retool_user.test_user", "active", "true"),
					resource.TestCheckResourceAttr("retool_user.test_user", "metadata", "{\"department\":\"engineering\"}"),
					resource.TestCheckResourceAttrSet("retool_user.test_user", "id"),
					resource.TestCheckResourceAttrSet("retool_user.test_user", "legacy_id"),
					resource.TestCheckResourceAttrSet("retool_user.test_user", "created_at"),
					resource.TestCheckResourceAttrSet("retool_user.test_user", "is_admin"),
					resource.TestCheckResourceAttrSet("retool_user.test_user", "two_factor_auth_enabled"),
				),
			},
		},
		// Note: The test framework will automatically call Delete on all resources after the last step,
		// which exercises the Delete function by setting the user to inactive.
	})
}

// TestAccUserWithDefaults tests a user resource with minimal configuration (default values).
func TestAccUserWithDefaults(t *testing.T) {
	acctest.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: testDefaultValuesConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_user.test_user_with_defaults", "email", "tf-acc-defaults-v4@example.retool.com"),
					resource.TestCheckResourceAttr("retool_user.test_user_with_defaults", "first_name", "Default"),
					resource.TestCheckResourceAttr("retool_user.test_user_with_defaults", "last_name", "User"),
					resource.TestCheckResourceAttr("retool_user.test_user_with_defaults", "active", "true"),
					resource.TestCheckResourceAttrSet("retool_user.test_user_with_defaults", "id"),
					resource.TestCheckResourceAttrSet("retool_user.test_user_with_defaults", "legacy_id"),
					resource.TestCheckResourceAttrSet("retool_user.test_user_with_defaults", "created_at"),
					resource.TestCheckResourceAttrSet("retool_user.test_user_with_defaults", "is_admin"),
					resource.TestCheckResourceAttrSet("retool_user.test_user_with_defaults", "two_factor_auth_enabled"),
				),
			},
		},
		// The test framework will call Delete after this step, exercising the Delete function.
	})
}

func sweepUsers(region string) error {
	log.Printf("Sweeping users in region %s", region)
	client, err := acctest.SweeperClient()
	if err != nil {
		return err
	}

	users, _, err := client.UsersAPI.UsersGet(context.Background()).Execute()
	if err != nil {
		return fmt.Errorf("Error reading users: %s", err.Error())
	}

	for _, user := range users.Data {
		if strings.HasPrefix(user.Email, "tf-acc") || strings.HasPrefix(user.Email, "terraform-test") {
			log.Printf("Deleting user %s", user.Email)
			_, err := client.UsersAPI.UsersUserIdDelete(context.Background(), user.Id).Execute()
			if err != nil {
				return fmt.Errorf("Error deleting user %s: %s", user.Email, err.Error())
			}
		}
	}
	return nil
}

func init() {
	resource.AddTestSweepers("retool_user", &resource.Sweeper{
		Name: "retool_user",
		F:    sweepUsers,
	})
}
