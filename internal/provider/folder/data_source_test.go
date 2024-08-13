package folder_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/tryretool/terraform-provider-retool/internal/acctest"
)

func checkFoldersListIsNotEmpty(state *terraform.State) error {
	folders := state.RootModule().Resources["data.retool_folders.test_folders"]
	if folders == nil {
		return fmt.Errorf("folder not found in state")
	}

	if len(folders.Primary.Attributes) == 0 {
		return fmt.Errorf("empty resource")
	}
	numFolders, err := strconv.Atoi(folders.Primary.Attributes["folders.#"])
	if err != nil || numFolders == 0 {
		return fmt.Errorf("no folders found")
	}

	return nil
}

const testFoldersDataSourceConfig = `
	data "retool_folders" "test_folders" {}
`

func TestAccFoldersDataSource(t *testing.T) {
	acctest.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: testFoldersDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					checkFoldersListIsNotEmpty,
					resource.TestCheckResourceAttrSet("data.retool_folders.test_folders", "folders.0.id"),
					resource.TestCheckResourceAttrSet("data.retool_folders.test_folders", "folders.0.name"),
					resource.TestCheckResourceAttrSet("data.retool_folders.test_folders", "folders.0.folder_type"),
					resource.TestCheckResourceAttrSet("data.retool_folders.test_folders", "folders.0.is_system_folder"),
					resource.TestCheckResourceAttrSet("data.retool_folders.test_folders", "folders.0.legacy_id"),
				),
			},
		}})
}
