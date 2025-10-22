// Package environments contains the data source implementation for Retool Environments.
package environments

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/tryretool/terraform-provider-retool/internal/provider/utils"
	"github.com/tryretool/terraform-provider-retool/internal/sdk/api"
)

// NewDataSource creates a new data source for environments.
func NewDataSource() datasource.DataSource {
	return &environmentsDataSource{}
}

var (
	_ datasource.DataSource              = &environmentsDataSource{}
	_ datasource.DataSourceWithConfigure = &environmentsDataSource{}
)

type environmentsDataSource struct {
	client *api.APIClient
}

type environmentsDataSourceModel struct {
	Environments []environmentModel `tfsdk:"environments"`
}

type environmentModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Default     types.Bool   `tfsdk:"default"`
	CreatedAt   types.String `tfsdk:"created_at"`
	UpdatedAt   types.String `tfsdk:"updated_at"`
}

// Configure adds the provider configured client to the data source.
func (d *environmentsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	providerData, ok := req.ProviderData.(*utils.ProviderData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			"Expected *utils.ProviderData.",
		)

		return
	}

	d.client = providerData.Client
}

func (d *environmentsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environments"
}

func (d *environmentsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Data source for Retool Environments. Environments are used to manage different settings for your Retool apps in different contexts, such as development, staging, and production.",
		Attributes: map[string]schema.Attribute{
			"environments": schema.ListNestedAttribute{
				Description: "List of Retool Environments.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The unique identifier for the environment.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the environment.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "A brief description of the environment.",
							Computed:    true,
						},
						"default": schema.BoolAttribute{
							Description: "Indicates if this is the default environment.",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Description: "The timestamp when the environment was created.",
							Computed:    true,
						},
						"updated_at": schema.StringAttribute{
							Description: "The timestamp when the environment was last updated.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *environmentsDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state environmentsDataSourceModel

	environments, _, err := d.client.EnvironmentAPI.EnvironmentsGet(ctx).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Retool Environments",
			err.Error(),
		)
		return
	}

	tflog.Info(ctx, "Retrieved environments", map[string]interface{}{"num_environments": environments.TotalCount})

	// Currently we do not support pagination for environments.
	for _, env := range environments.Data {
		state.Environments = append(state.Environments, environmentModel{
			ID:          types.StringValue(env.Id),
			Name:        types.StringValue(env.Name),
			Description: types.StringPointerValue(env.Description.Get()),
			Default:     types.BoolValue(env.Default),
			CreatedAt:   types.StringValue(env.CreatedAt),
			UpdatedAt:   types.StringValue(env.UpdatedAt),
		})
		if resp.Diagnostics.HasError() {
			return
		}
	}

	// Set state.
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
