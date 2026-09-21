package retoolresource

import (
	"encoding/json"
	"testing"

	"github.com/tryretool/terraform-provider-retool/internal/sdk/api"
)

// resourceResponseJSON builds a GET /resources/{id} body, optionally including the
// data_access_enforced field that Retool 4.34+ returns.
func resourceResponseJSON(withDataAccessEnforced bool) string {
	dataAccessEnforced := ""
	if withDataAccessEnforced {
		dataAccessEnforced = `"data_access_enforced": true,`
	}
	return `{
		"success": true,
		"data": {
			"id": "my-api",
			"type": "restapi",
			"display_name": "My API",
			"folder_id": null,
			"protected": false,
			` + dataAccessEnforced + `
			"created_at": "2019-02-08T11:45:48.899Z",
			"updated_at": "2019-02-24T18:28:18.790Z"
		}
	}`
}

func TestResourceResponse_DecodesDataAccessEnforced(t *testing.T) {
	t.Parallel()

	var response api.ResourcesResourceIdGet200Response
	if err := json.Unmarshal([]byte(resourceResponseJSON(true)), &response); err != nil {
		t.Fatalf("unexpected error decoding response: %s", err)
	}

	if !response.Data.HasDataAccessEnforced() {
		t.Fatal("expected data_access_enforced to be set")
	}
	if !response.Data.GetDataAccessEnforced() {
		t.Error("expected data_access_enforced to be true")
	}
	if response.Data.Id != "my-api" {
		t.Errorf("expected id my-api, got %s", response.Data.Id)
	}
}

// Instances older than 4.34 omit the field entirely; the provider still has to read them.
func TestResourceResponse_DecodesWithoutDataAccessEnforced(t *testing.T) {
	t.Parallel()

	var response api.ResourcesResourceIdGet200Response
	if err := json.Unmarshal([]byte(resourceResponseJSON(false)), &response); err != nil {
		t.Fatalf("unexpected error decoding response: %s", err)
	}

	if response.Data.HasDataAccessEnforced() {
		t.Error("expected data_access_enforced to be unset")
	}
}
