package group_test

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/tryretool/terraform-provider-retool/internal/acctest"
	"github.com/tryretool/terraform-provider-retool/internal/provider/utils"
)

const testGroupConfig = `
	resource "retool_group" "test_group" {
		name = "tf-acc-test-group"
		universal_app_access = "use"
		universal_resource_access = "own"
		universal_workflow_access = "none"
		universal_query_library_access = "edit"
		user_list_access = false
		audit_log_access = true
		unpublished_release_access = false
		usage_analytics_access = true
		account_details_access = true
	}
		` // not testing landing_page_app_id because it has to match an existing app id, which is hard to achieve when hitting random Retool instance

const testUpdatedGroupConfig = `
	resource "retool_group" "test_group" {
		name = "tf-acc-test-group-modified"
		universal_app_access = "own"
		universal_resource_access = "none"
		universal_workflow_access = "edit"
		universal_query_library_access = "none"
		user_list_access = true
		audit_log_access = false
		unpublished_release_access = true
		usage_analytics_access = false
		account_details_access = false
	}
		`

const testDefaultValuesConfig = `
	resource "retool_group" "test_group_with_defaults" {
		name = "tf-acc-test-group-with-defaults"
	}
		`

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func TestAccGroup(t *testing.T) {
	acctest.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			// Read and Create
			{
				Config: testGroupConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_group.test_group", "name", "tf-acc-test-group"),
					resource.TestCheckResourceAttr("retool_group.test_group", "universal_app_access", "use"),
					resource.TestCheckResourceAttr("retool_group.test_group", "universal_resource_access", "own"),
					resource.TestCheckResourceAttr("retool_group.test_group", "universal_workflow_access", "none"),
					resource.TestCheckResourceAttr("retool_group.test_group", "universal_query_library_access", "edit"),
					resource.TestCheckResourceAttr("retool_group.test_group", "user_list_access", "false"),
					resource.TestCheckResourceAttr("retool_group.test_group", "audit_log_access", "true"),
					resource.TestCheckResourceAttr("retool_group.test_group", "unpublished_release_access", "false"),
					resource.TestCheckResourceAttr("retool_group.test_group", "usage_analytics_access", "true"),
					resource.TestCheckResourceAttr("retool_group.test_group", "account_details_access", "true"),
					resource.TestCheckResourceAttrSet("retool_group.test_group", "id"),
					resource.TestCheckResourceAttrSet("retool_group.test_group", "legacy_id"),
					resource.TestCheckNoResourceAttr("retool_group.test_group", "landing_page_app_id"),
				),
			},
			// Import state
			{
				ResourceName:      "retool_group.test_group",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read
			{
				Config: testUpdatedGroupConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_group.test_group", "name", "tf-acc-test-group-modified"),
					resource.TestCheckResourceAttr("retool_group.test_group", "universal_app_access", "own"),
					resource.TestCheckResourceAttr("retool_group.test_group", "universal_resource_access", "none"),
					resource.TestCheckResourceAttr("retool_group.test_group", "universal_workflow_access", "edit"),
					resource.TestCheckResourceAttr("retool_group.test_group", "universal_query_library_access", "none"),
					resource.TestCheckResourceAttr("retool_group.test_group", "user_list_access", "true"),
					resource.TestCheckResourceAttr("retool_group.test_group", "audit_log_access", "false"),
					resource.TestCheckResourceAttr("retool_group.test_group", "unpublished_release_access", "true"),
					resource.TestCheckResourceAttr("retool_group.test_group", "usage_analytics_access", "false"),
					resource.TestCheckResourceAttr("retool_group.test_group", "account_details_access", "false"),
					resource.TestCheckResourceAttrSet("retool_group.test_group", "id"),
					resource.TestCheckResourceAttrSet("retool_group.test_group", "legacy_id"),
					resource.TestCheckNoResourceAttr("retool_group.test_group", "landing_page_app_id"),
				),
			},
			// Check default values
			{
				Config: testDefaultValuesConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_group.test_group_with_defaults", "name", "tf-acc-test-group-with-defaults"),
					resource.TestCheckResourceAttr("retool_group.test_group_with_defaults", "universal_app_access", "none"),
					resource.TestCheckResourceAttr("retool_group.test_group_with_defaults", "universal_resource_access", "none"),
					resource.TestCheckResourceAttr("retool_group.test_group_with_defaults", "universal_workflow_access", "none"),
					resource.TestCheckResourceAttr("retool_group.test_group_with_defaults", "universal_query_library_access", "none"),
					resource.TestCheckResourceAttr("retool_group.test_group_with_defaults", "user_list_access", "false"),
					resource.TestCheckResourceAttr("retool_group.test_group_with_defaults", "audit_log_access", "false"),
					resource.TestCheckResourceAttr("retool_group.test_group_with_defaults", "unpublished_release_access", "false"),
					resource.TestCheckResourceAttr("retool_group.test_group_with_defaults", "usage_analytics_access", "false"),
					resource.TestCheckResourceAttr("retool_group.test_group_with_defaults", "account_details_access", "false"),
					resource.TestCheckResourceAttrSet("retool_group.test_group_with_defaults", "id"),
					resource.TestCheckResourceAttrSet("retool_group.test_group_with_defaults", "legacy_id"),
					resource.TestCheckNoResourceAttr("retool_group.test_group_with_defaults", "landing_page_app_id"),
				),
			},
		},
	})
}

func sweepGroups(region string) error {
	log.Printf("Sweeping groups in region %s", region)
	client, err := acctest.SweeperClient()
	if err != nil {
		return err
	}

	groups, _, err := client.GroupsAPI.GroupsGet(context.Background()).Execute()
	if err != nil {
		return fmt.Errorf("Error reading groups: %s", err.Error())
	}

	for _, group := range groups.Data {
		if strings.HasPrefix(group.Name, "tf-acc") {
			log.Printf("Deleting group %s", group.Name)
			_, err := client.GroupsAPI.GroupsGroupIdDelete(context.Background(), utils.Float32PtrToIntString(group.Id.Get())).Execute()
			if err != nil {
				return fmt.Errorf("Error deleting group %s: %s", group.Name, err.Error())
			}
		}
	}
	return nil
}

func init() {
	resource.AddTestSweepers("retool_group", &resource.Sweeper{
		Name: "retool_group",
		F:    sweepGroups,
	})
}
