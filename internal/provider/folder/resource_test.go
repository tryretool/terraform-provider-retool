package folder_test

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-log/tflog"
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

const testUpdatedFolderConfig = `
	resource "retool_folder" "parent_folder" {
		name = "tf-acc-parent-folder"
		folder_type = "app"
	}

	resource "retool_folder" "test_folder" {
		name = "tf-acc-test-folder-renamed"
		folder_type = "app"
		parent_folder_id = retool_folder.parent_folder.id
	}

	resource "retool_folder" "child_folder" {
		name = "tf-acc-test-child-folder"
		folder_type = "app"
	}
`

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func sweepFolders(region string) error {
	log.Printf("Sweeping folders in region %s", region)
	client, err := acctest.SweeperClient()
	if err != nil {
		return err
	}

	folders, _, err := client.FoldersAPI.FoldersGet(context.Background()).Execute()
	if err != nil {
		return fmt.Errorf("Error reading folders: %s", err.Error())
	}

	for _, folder := range folders.Data {
		if strings.HasPrefix(folder.Name, "tf-acc") {
			log.Printf("Deleting folder %s", folder.Name)
			tflog.Info(context.Background(), "Deleting folder", map[string]interface{}{"folder": folder.Name})
			_, err := client.FoldersAPI.FoldersFolderIdDelete(context.Background(), folder.Id).Recursive(true).Execute()
			if err != nil {
				return fmt.Errorf("Error deleting folder %s: %s", folder.Name, err.Error())
			}
		}
	}
	return nil
}

func init() {
	resource.AddTestSweepers("retool_folder", &resource.Sweeper{
		Name: "retool_folder",
		F:    sweepFolders,
	})
}

func TestAccFolder(t *testing.T) {
	acctest.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			// Read and Create.
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
			// Import state.
			{
				ResourceName:      "retool_folder.test_folder",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read.
			{
				Config: testUpdatedFolderConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_folder.test_folder", "name", "tf-acc-test-folder-renamed"),
					resource.TestCheckResourceAttr("retool_folder.test_folder", "folder_type", "app"),
					resource.TestCheckResourceAttr("retool_folder.test_folder", "is_system_folder", "false"),
					resource.TestCheckResourceAttrSet("retool_folder.test_folder", "id"),
					resource.TestCheckResourceAttrSet("retool_folder.test_folder", "legacy_id"),
					resource.TestCheckResourceAttrPair("retool_folder.test_folder", "parent_folder_id", "retool_folder.parent_folder", "id"),
					resource.TestCheckResourceAttr("retool_folder.child_folder", "parent_folder_id", "ROOT"),
				),
			},
		},
	})
}
