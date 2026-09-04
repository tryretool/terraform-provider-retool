package retoolresource

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func baseModel() retoolResourceModel {
	return retoolResourceModel{
		Type:        types.StringValue("restapi"),
		DisplayName: types.StringValue("My API"),
		Options:     types.StringValue(`{"base_url":"https://api.example.com"}`),
		FolderID:    types.StringNull(),
	}
}

func TestBuildResourceCreateBody_OmitsFolderIDWhenUnset(t *testing.T) {
	t.Parallel()

	body, err := buildResourceCreateBody(baseModel())
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if _, ok := body["folder_id"]; ok {
		t.Errorf("folder_id should be omitted when unset, got %v", body["folder_id"])
	}
	if body["type"] != "restapi" {
		t.Errorf("expected type restapi, got %v", body["type"])
	}
	if body["display_name"] != "My API" {
		t.Errorf("expected display_name My API, got %v", body["display_name"])
	}
}

func TestBuildResourceCreateBody_OmitsFolderIDWhenUnknown(t *testing.T) {
	t.Parallel()

	plan := baseModel()
	plan.FolderID = types.StringUnknown()

	body, err := buildResourceCreateBody(plan)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if _, ok := body["folder_id"]; ok {
		t.Errorf("folder_id should be omitted when unknown, got %v", body["folder_id"])
	}
}

func TestBuildResourceCreateBody_IncludesFolderIDWhenSet(t *testing.T) {
	t.Parallel()

	plan := baseModel()
	plan.FolderID = types.StringValue("folder_123")

	body, err := buildResourceCreateBody(plan)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if body["folder_id"] != "folder_123" {
		t.Errorf("expected folder_id folder_123, got %v", body["folder_id"])
	}
}

func TestBuildResourceCreateBody_InvalidOptions(t *testing.T) {
	t.Parallel()

	plan := baseModel()
	plan.Options = types.StringValue("not json")

	if _, err := buildResourceCreateBody(plan); err == nil {
		t.Error("expected error for invalid options JSON, got nil")
	}
}
