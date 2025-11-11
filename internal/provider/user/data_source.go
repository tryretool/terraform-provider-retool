// Package user provides implementation of the User resource and Users data source.
package user

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

// NewDataSource creates a new data source for users.
func NewDataSource() datasource.DataSource {
	return &usersDataSource{}
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &usersDataSource{}
	_ datasource.DataSourceWithConfigure = &usersDataSource{}
)

type usersDataSource struct {
	client *api.APIClient
}

type usersDataSourceModel struct {
	Users []userDataSourceModel `tfsdk:"users"`
}

type userDataSourceModel struct {
	ID                   types.String `tfsdk:"id"`
	LegacyID             types.String `tfsdk:"legacy_id"`
	Email                types.String `tfsdk:"email"`
	FirstName            types.String `tfsdk:"first_name"`
	LastName             types.String `tfsdk:"last_name"`
	Active               types.Bool   `tfsdk:"active"`
	Metadata             types.String `tfsdk:"metadata"`
	UserType             types.String `tfsdk:"user_type"`
	CreatedAt            types.String `tfsdk:"created_at"`
	LastActive           types.String `tfsdk:"last_active"`
	IsAdmin              types.Bool   `tfsdk:"is_admin"`
	TwoFactorAuthEnabled types.Bool   `tfsdk:"two_factor_auth_enabled"`
}

// Configure adds the provider configured client to the data source.
func (d *usersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
}

func (d *usersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_users"
}

func (d *usersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Users data source allows you to retrieve a list of users in Retool.",
		Attributes: map[string]schema.Attribute{
			"users": schema.ListNestedAttribute{
				Computed:    true,
				Description: "A list of users",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:    true,
							Description: "The ID of the user. Currently this is the same as legacy_id but will change in the future.",
						},
						"legacy_id": schema.StringAttribute{
							Computed:    true,
							Description: "The legacy ID of the user.",
						},
						"email": schema.StringAttribute{
							Computed:    true,
							Description: "The email address of the user.",
						},
						"first_name": schema.StringAttribute{
							Computed:    true,
							Description: "The first name of the user.",
						},
						"last_name": schema.StringAttribute{
							Computed:    true,
							Description: "The last name of the user.",
						},
						"active": schema.BoolAttribute{
							Computed:    true,
							Description: "Whether the user is active or not.",
						},
						"metadata": schema.StringAttribute{
							Computed:    true,
							Description: "JSON string containing custom metadata for the user.",
						},
						"user_type": schema.StringAttribute{
							Computed:    true,
							Description: "The user type.",
						},
						"created_at": schema.StringAttribute{
							Computed:    true,
							Description: "The timestamp when the user was created.",
						},
						"last_active": schema.StringAttribute{
							Computed:    true,
							Description: "The timestamp when the user was last active.",
						},
						"is_admin": schema.BoolAttribute{
							Computed:    true,
							Description: "Whether the user is an admin or not.",
						},
						"two_factor_auth_enabled": schema.BoolAttribute{
							Computed:    true,
							Description: "Whether two-factor authentication is enabled for this user.",
						},
					},
				},
			},
		},
	}
}

func (d *usersDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state usersDataSourceModel

	users, httpResponse, err := d.client.UsersAPI.UsersGet(ctx).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Users via Retool API",
			err.Error(),
		)
		tflog.Error(ctx, "Error reading users", utils.AddHTTPStatusCode(map[string]any{"error": err.Error()}, httpResponse))
		return
	}
	tflog.Info(ctx, "Read Users via Retool API", map[string]any{"num_users": len(users.Data)})

	for _, user := range users.Data {
		userModel := userDataSourceModel{
			ID:                   types.StringValue(user.Id),
			LegacyID:             types.StringValue(fmt.Sprintf("%.0f", user.LegacyId)),
			Email:                types.StringValue(user.Email),
			Active:               types.BoolValue(user.Active),
			IsAdmin:              types.BoolValue(user.IsAdmin),
			UserType:             types.StringValue(user.UserType),
			TwoFactorAuthEnabled: types.BoolValue(user.TwoFactorAuthEnabled),
			CreatedAt:            types.StringValue(user.CreatedAt.String()),
		}

		// Handle nullable fields
		if user.FirstName.Get() != nil {
			userModel.FirstName = types.StringValue(*user.FirstName.Get())
		} else {
			userModel.FirstName = types.StringNull()
		}

		if user.LastName.Get() != nil {
			userModel.LastName = types.StringValue(*user.LastName.Get())
		} else {
			userModel.LastName = types.StringNull()
		}

		if user.LastActive.Get() != nil {
			userModel.LastActive = types.StringValue(user.LastActive.Get().String())
		} else {
			userModel.LastActive = types.StringNull()
		}

		// Handle metadata
		if len(user.Metadata) > 0 {
			metadataStr, err := utils.MapToJSONString(user.Metadata)
			if err != nil {
				resp.Diagnostics.AddError(
					"Error serializing metadata",
					fmt.Sprintf("Could not serialize metadata for user %s: %s", user.Email, err.Error()),
				)
				return
			}
			userModel.Metadata = types.StringValue(metadataStr)
		} else {
			userModel.Metadata = types.StringNull()
		}

		state.Users = append(state.Users, userModel)
	}

	// Set state.
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

