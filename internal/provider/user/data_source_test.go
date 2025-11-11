package user_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/tryretool/terraform-provider-retool/internal/acctest"
)

func checkUsersListIsNotEmpty(state *terraform.State) error {
	users := state.RootModule().Resources["data.retool_users.test_users"]
	if users == nil {
		return fmt.Errorf("users not found in state")
	}

	if len(users.Primary.Attributes) == 0 {
		return fmt.Errorf("empty resource")
	}
	numUsers, err := strconv.Atoi(users.Primary.Attributes["users.#"])
	if err != nil || numUsers == 0 {
		return fmt.Errorf("no users found")
	}

	return nil
}

const testUsersDataSourceConfig = `
	data "retool_users" "test_users" {}
`

func TestAccUsersDataSource(t *testing.T) {
	acctest.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: testUsersDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					checkUsersListIsNotEmpty,
					resource.TestCheckResourceAttrSet("data.retool_users.test_users", "users.0.id"),
					resource.TestCheckResourceAttrSet("data.retool_users.test_users", "users.0.email"),
					resource.TestCheckResourceAttrSet("data.retool_users.test_users", "users.0.legacy_id"),
					resource.TestCheckResourceAttrSet("data.retool_users.test_users", "users.0.created_at"),
					resource.TestCheckResourceAttrSet("data.retool_users.test_users", "users.0.active"),
					resource.TestCheckResourceAttrSet("data.retool_users.test_users", "users.0.is_admin"),
					resource.TestCheckResourceAttrSet("data.retool_users.test_users", "users.0.two_factor_auth_enabled"),
				),
			},
		},
	})
}

