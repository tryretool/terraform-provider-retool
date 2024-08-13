package folder

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/tryretool/terraform-provider-retool/internal/provider/utils"
	"github.com/tryretool/terraform-provider-retool/internal/sdk/api"
)

// Fake "id" that we use to represent the root folder.
const RootFolderID string = "ROOT"

func getRootFolderID(ctx context.Context, folderType string, client *api.APIClient, cache *map[string]string) (string, error) {
	if cache != nil {
		if rootFolderID, ok := (*cache)[folderType]; ok {
			return rootFolderID, nil
		}
	}
	tflog.Info(ctx, "Getting root folder ID", map[string]any{"folderType": folderType})
	response, httpResponse, err := client.FoldersAPI.FoldersGet(ctx).Execute()
	if err != nil {
		tflog.Error(ctx, "Error getting root folder ID", utils.AddHTTPStatusCode(map[string]any{"error": err.Error()}, httpResponse))
		return "", err
	}

	for _, folder := range response.Data {
		if folder.FolderType == folderType && folder.ParentFolderId.Get() == nil {
			if cache != nil {
				(*cache)[folderType] = folder.Id
			}
			return folder.Id, nil
		}
	}
	tflog.Error(ctx, "Root folder not found", map[string]any{"folderType": folderType})

	return "", fmt.Errorf("root folder not found for type %s", folderType)
}

func maybeReplaceRootFolderIDWithConstant(ctx context.Context, folderType string, folderID *string, client *api.APIClient, cache *map[string]string, diags *diag.Diagnostics) *string {
	if folderID != nil {
		rootFolderID, err := getRootFolderID(ctx, folderType, client, cache)
		if err != nil {
			diags.AddError(
				"Error reading folder",
				"Could not find root folder for type "+folderType+": "+err.Error(),
			)
			return nil
		}
		if *folderID == rootFolderID {
			tempStr := RootFolderID
			return &tempStr
		}
	}
	return folderID
}
