package folder_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/tryretool/terraform-provider-retool/internal/acctest"
)

const testFolderConfig = `
	resource "retool_folder" "test_folder" {
		name = "tf-acc-test-folder"
		folder_type = "app"
	}

	resource "retool_folder" "child_folder" {
		name = "tf-acc-test-child-folder"
		folder_type = "app"
		parent_folder_id = retool_folder.test_folder.id
	}
`

func TestAccFolder(t *testing.T) {
	acctest.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			// Read and Create
			{
				Config: testFolderConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_folder.test_folder", "name", "tf-acc-test-folder"),
					resource.TestCheckResourceAttr("retool_folder.test_folder", "folder_type", "app"),
					resource.TestCheckResourceAttr("retool_folder.test_folder", "is_system_folder", "false"),
					resource.TestCheckResourceAttrSet("retool_folder.test_folder", "id"),
					resource.TestCheckResourceAttrSet("retool_folder.test_folder", "legacy_id"),
					resource.TestCheckResourceAttrSet("retool_folder.test_folder", "parent_folder_id"),

					resource.TestCheckResourceAttr("retool_folder.child_folder", "name", "tf-acc-test-child-folder"),
					resource.TestCheckResourceAttr("retool_folder.child_folder", "folder_type", "app"),
					resource.TestCheckResourceAttr("retool_folder.child_folder", "is_system_folder", "false"),
					resource.TestCheckResourceAttrPair("retool_folder.child_folder", "parent_folder_id", "retool_folder.test_folder", "id"),
					resource.TestCheckResourceAttrSet("retool_folder.child_folder", "id"),
					resource.TestCheckResourceAttrSet("retool_folder.child_folder", "legacy_id"),
				),
			},
			// Import state
			{
				ResourceName:      "retool_folder.test_folder",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
