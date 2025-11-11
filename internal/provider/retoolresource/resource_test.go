package retoolresource_test

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/tryretool/terraform-provider-retool/internal/acctest"
)

// Note: These tests use a REST API resource as it's the simplest to configure
// Options are write-only, so we can't verify them after creation

const testRetoolResourceConfig = `
	resource "retool_resource" "test_resource" {
		display_name = "Test REST API"
		type         = "restapi"
		options      = jsonencode({
			base_url = "https://api.example.com"
		})
	}
`

const testRetoolResourceUpdatedConfig = `
	resource "retool_resource" "test_resource" {
		display_name = "Updated REST API"
		type         = "restapi"
		options      = jsonencode({
			base_url = "https://api.example.com"
		})
	}
`

func TestAccRetoolResource(t *testing.T) {
	acctest.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testRetoolResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_resource.test_resource", "display_name", "Test REST API"),
					resource.TestCheckResourceAttr("retool_resource.test_resource", "type", "restapi"),
					resource.TestCheckResourceAttrSet("retool_resource.test_resource", "id"),
					resource.TestCheckResourceAttrSet("retool_resource.test_resource", "created_at"),
					resource.TestCheckResourceAttrSet("retool_resource.test_resource", "updated_at"),
					resource.TestCheckResourceAttrSet("retool_resource.test_resource", "protected"),
				),
			},
			// Import state
			{
				ResourceName:            "retool_resource.test_resource",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"options"}, // options cannot be read back
			},
			// Update and Read
			{
				Config: testRetoolResourceUpdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_resource.test_resource", "display_name", "Updated REST API"),
					resource.TestCheckResourceAttr("retool_resource.test_resource", "type", "restapi"),
					resource.TestCheckResourceAttrSet("retool_resource.test_resource", "id"),
				),
			},
		},
		// Note: The test framework will automatically call Delete on all resources after the last step
	})
}

func sweepRetoolResources(region string) error {
	log.Printf("Sweeping Retool resources in region %s", region)
	client, err := acctest.SweeperClient()
	if err != nil {
		return err
	}

	resources, _, err := client.ResourcesAPI.ResourcesGet(context.Background()).Execute()
	if err != nil {
		return fmt.Errorf("Error reading resources: %s", err.Error())
	}

	for _, res := range resources.Data {
		if strings.HasPrefix(res.DisplayName, "Test ") || strings.HasPrefix(res.DisplayName, "Updated ") {
			log.Printf("Deleting resource %s (%s)", res.DisplayName, res.Id)
			_, err := client.ResourcesAPI.ResourcesResourceIdDelete(context.Background(), res.Id).Execute()
			if err != nil {
				return fmt.Errorf("Error deleting resource %s: %s", res.Id, err.Error())
			}
		}
	}
	return nil
}

func init() {
	resource.AddTestSweepers("retool_resource", &resource.Sweeper{
		Name: "retool_resource",
		F:    sweepRetoolResources,
	})
}
