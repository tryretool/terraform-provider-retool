package folder

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tryretool/terraform-provider-retool/internal/provider/utils"
	"github.com/tryretool/terraform-provider-retool/internal/sdk/api"
)

func NewDataSource() datasource.DataSource {
	return &foldersDataSource{}
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &foldersDataSource{}
	_ datasource.DataSourceWithConfigure = &foldersDataSource{}
)

type foldersDataSource struct {
	client            *api.APIClient
	rootFolderIdCache *map[string]string
}

type foldersDataSourceModel struct {
	Folders []folderModel `tfsdk:"folders"`
}

type folderModel struct {
	Id             types.String `tfsdk:"id"`
	LegacyId       types.String `tfsdk:"legacy_id"`
	Name           types.String `tfsdk:"name"`
	ParentFolderId types.String `tfsdk:"parent_folder_id"`
	IsSystemFolder types.Bool   `tfsdk:"is_system_folder"`
	FolderType     types.String `tfsdk:"folder_type"`
}

// Configure adds the provider configured client to the data source.
func (d *foldersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	providerData, ok := req.ProviderData.(*utils.ProviderData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *utils.ProviderData, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = providerData.Client
	d.rootFolderIdCache = providerData.RootFolderIdCache
}

func (d *foldersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_folders"
}

func (d *foldersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"folders": schema.ListNestedAttribute{
				Computed:    true,
				Description: "A list of folders",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:    true,
							Description: "The id of the folder. Currently this is the same as legacy_id but will be different in the future.",
						},
						"legacy_id": schema.StringAttribute{
							Computed:    true,
							Description: "The legacy id of the folder.",
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Description: "The name of the folder.",
						},
						"parent_folder_id": schema.StringAttribute{
							Computed:    true,
							Description: "The id of the parent folder.",
						},
						"is_system_folder": schema.BoolAttribute{
							Computed:    true,
							Description: "Whether the folder is a system folder.",
						},
						"folder_type": schema.StringAttribute{
							Computed:    true,
							Description: "The type of the folder: (app|file|resource|workflow).",
						},
					},
				},
			},
		},
	}
}

func (d *foldersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state foldersDataSourceModel

	folders, _, err := d.client.FoldersAPI.FoldersGet(ctx).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Folders via Retool API",
			err.Error(),
		)
		return
	}
	tflog.Info(ctx, "Read Folders via Retool API", map[string]any{"num_folders": folders.TotalCount})

	// Fun fact: even though the response has next_token and has_more fields, the API does not support pagination and always returns all folders.
	for _, folder := range folders.Data {
		state.Folders = append(state.Folders, folderModel{
			Id:             types.StringValue(folder.Id),
			LegacyId:       types.StringValue(folder.LegacyId),
			Name:           types.StringValue(folder.Name),
			ParentFolderId: types.StringPointerValue(maybeReplaceRootFolderIdWithConstant(ctx, folder.FolderType, folder.ParentFolderId.Get(), d.client, d.rootFolderIdCache, &resp.Diagnostics)),
			IsSystemFolder: types.BoolValue(folder.IsSystemFolder),
			FolderType:     types.StringValue(folder.FolderType),
		})
		if resp.Diagnostics.HasError() {
			return
		}
	}

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
