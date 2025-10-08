package group_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/tryretool/terraform-provider-retool/internal/acctest"
)

func checkGroupsListIsNotEmpty(state *terraform.State) error {
	groups := state.RootModule().Resources["data.retool_groups.test_groups"]
	if groups == nil {
		return fmt.Errorf("group not found in state")
	}

	if len(groups.Primary.Attributes) == 0 {
		return fmt.Errorf("empty resource")
	}
	numGroups, err := strconv.Atoi(groups.Primary.Attributes["groups.#"])
	if err != nil || numGroups == 0 {
		return fmt.Errorf("no groups found")
	}

	return nil
}

const testGroupsDataSourceConfig = `
	data "retool_groups" "test_groups" {}
`

func TestAccGroupsDataSource(t *testing.T) {
	// By default, retool instances has 4 groups: Admin, Editor, Viewer, and All Users.
	// In this test we just validate that we managed to fetch 4 groups.
	acctest.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: testGroupsDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					checkGroupsListIsNotEmpty,
					resource.TestCheckResourceAttrSet("data.retool_groups.test_groups", "groups.0.id"),
					resource.TestCheckResourceAttrSet("data.retool_groups.test_groups", "groups.0.name"),
					resource.TestCheckResourceAttrSet("data.retool_groups.test_groups", "groups.1.id"),
					resource.TestCheckResourceAttrSet("data.retool_groups.test_groups", "groups.1.name"),
					resource.TestCheckResourceAttrSet("data.retool_groups.test_groups", "groups.2.id"),
					resource.TestCheckResourceAttrSet("data.retool_groups.test_groups", "groups.2.name"),
					resource.TestCheckResourceAttrSet("data.retool_groups.test_groups", "groups.3.id"),
					resource.TestCheckResourceAttrSet("data.retool_groups.test_groups", "groups.3.name"),
				),
			},
		}})
}
