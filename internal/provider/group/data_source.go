// Package group provides implementation of the Group resource and Groups data source.
package group

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

// NewDataSource creates a new data source for groups.
func NewDataSource() datasource.DataSource {
	return &groupsDataSource{}
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &groupsDataSource{}
	_ datasource.DataSourceWithConfigure = &groupsDataSource{}
)

type groupsDataSource struct {
	client *api.APIClient
}

type groupsDataSourceModel struct {
	Groups []groupModel `tfsdk:"groups"`
}

type groupMemberModel struct {
	ID           types.String `tfsdk:"id"`
	Email        types.String `tfsdk:"email"`
	IsGroupAdmin types.Bool   `tfsdk:"is_group_admin"`
}

type groupUserInvitesModel struct {
	ID           types.String `tfsdk:"id"`
	LegacyID     types.String `tfsdk:"legacy_id"`
	InvitedByID  types.String `tfsdk:"invited_by_id"`
	InvitedEmail types.String `tfsdk:"invited_email"`
	ExpiresAt    types.String `tfsdk:"expires_at"`
	ClaimedBy    types.String `tfsdk:"claimed_by"`
	ClaimedAt    types.String `tfsdk:"claimed_at"`
	UserType     types.String `tfsdk:"user_type"`
	CreatedAt    types.String `tfsdk:"created_at"`
	InviteLink   types.String `tfsdk:"invite_link"`
}

type groupModel struct {
	ID                          types.String            `tfsdk:"id"`
	LegacyID                    types.String            `tfsdk:"legacy_id"`
	Name                        types.String            `tfsdk:"name"`
	Members                     []groupMemberModel      `tfsdk:"members"`
	UniversalAppAccess          types.String            `tfsdk:"universal_app_access"`
	UniversalResourceAccess     types.String            `tfsdk:"universal_resource_access"`
	UniversalWorkflowAccess     types.String            `tfsdk:"universal_workflow_access"`
	UniversalQueryLibraryAccess types.String            `tfsdk:"universal_query_library_access"`
	UserInvites                 []groupUserInvitesModel `tfsdk:"user_invites"`
	UserListAccess              types.Bool              `tfsdk:"user_list_access"`
	AuditLogAccess              types.Bool              `tfsdk:"audit_log_access"`
	UnpublishedReleaseAccess    types.Bool              `tfsdk:"unpublished_release_access"`
	UsageAnalyticsAccess        types.Bool              `tfsdk:"usage_analytics_access"`
	ThemeAccess                 types.Bool              `tfsdk:"theme_access"`
	AccountDetailsAccess        types.Bool              `tfsdk:"account_details_access"`
	LandingPageAppID            types.String            `tfsdk:"landing_page_app_id"`
	CreatedAt                   types.String            `tfsdk:"created_at"`
	UpdatedAt                   types.String            `tfsdk:"updated_at"`
}

// Configure adds the provider configured client to the data source.
func (d *groupsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *groupsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_groups"
}

func (d *groupsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Groups data source allows you to retrieve a list of groups in Retool.",
		Attributes: map[string]schema.Attribute{
			"groups": schema.ListNestedAttribute{
				Computed:    true,
				Description: "A list of groups",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:    true,
							Description: "The id of the group.",
						},
						"legacy_id": schema.StringAttribute{
							Computed:    true,
							Description: "The legacy id of the group.",
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Description: "The name of the group.",
						},
						"members": schema.ListNestedAttribute{
							Computed:    true,
							Description: "The members of the group.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Computed:    true,
										Description: "The id of the member.",
									},
									"email": schema.StringAttribute{
										Computed:    true,
										Description: "The email of the member.",
									},
									"is_group_admin": schema.BoolAttribute{
										Computed:    true,
										Description: "Whether the member is a group admin.",
									},
								},
							},
						},
						"universal_app_access": schema.StringAttribute{
							Computed:    true,
							Description: "The universal app access level of the group.",
						},
						"universal_resource_access": schema.StringAttribute{
							Computed:    true,
							Description: "The universal resource access level of the group.",
						},
						"universal_workflow_access": schema.StringAttribute{
							Computed:    true,
							Description: "The universal workflow access level of the group.",
						},
						"universal_query_library_access": schema.StringAttribute{
							Computed:    true,
							Description: "The universal query library access level of the group.",
						},
						"user_invites": schema.ListNestedAttribute{
							Computed:    true,
							Description: "The user invites for the group.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Computed:    true,
										Description: "The id of the user invite.",
									},
									"legacy_id": schema.StringAttribute{
										Computed:    true,
										Description: "The legacy id of the user invite.",
									},
									"invited_by_id": schema.StringAttribute{
										Computed:    true,
										Description: "The id of the user who invited the user.",
									},
									"invited_email": schema.StringAttribute{
										Computed:    true,
										Description: "The email of the invited user.",
									},
									"expires_at": schema.StringAttribute{
										Computed:    true,
										Description: "The expiration date of the invite.",
									},
									"claimed_by": schema.StringAttribute{
										Computed:    true,
										Description: "The id of the user who claimed the invite.",
									},
									"claimed_at": schema.StringAttribute{
										Computed:    true,
										Description: "The date the invite was claimed.",
									},
									"user_type": schema.StringAttribute{
										Computed:    true,
										Description: "The type of user invited.",
									},
									"created_at": schema.StringAttribute{
										Computed:    true,
										Description: "The date the invite was created.",
									},
									"invite_link": schema.StringAttribute{
										Computed:    true,
										Description: "The invite link.",
									},
								},
							},
						},
						"user_list_access": schema.BoolAttribute{
							Computed:    true,
							Description: "Whether the group has user list access.",
						},
						"audit_log_access": schema.BoolAttribute{
							Computed:    true,
							Description: "Whether the group has audit log access.",
						},
						"unpublished_release_access": schema.BoolAttribute{
							Computed:    true,
							Description: "Whether the group has unpublished release access.",
						},
					"usage_analytics_access": schema.BoolAttribute{
						Computed:    true,
						Description: "Whether the group has usage analytics access.",
					},
					"theme_access": schema.BoolAttribute{
						Computed:    true,
						Description: "Whether the group has access to edit themes.",
					},
					"account_details_access": schema.BoolAttribute{
						Computed:    true,
						Description: "Whether the group has access to account details.",
					},
					"landing_page_app_id": schema.StringAttribute{
						Computed:    true,
						Description: "The id of the landing page app for the group.",
					},
						"created_at": schema.StringAttribute{
							Computed:    true,
							Description: "The date the group was created.",
						},
						"updated_at": schema.StringAttribute{
							Computed:    true,
							Description: "The date the group was last updated.",
						},
					},
				},
			},
		},
	}
}

func (d *groupsDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state groupsDataSourceModel

	groups, _, err := d.client.GroupsAPI.GroupsGet(ctx).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Groups via Retool API",
			err.Error(),
		)
		return
	}
	tflog.Info(ctx, "Read Groups via Retool API", map[string]any{"num_groups": groups.TotalCount})

	for _, group := range groups.Data {
		tflog.Debug(ctx, "Processing group", map[string]any{"group_id": group.Id, "group_name": group.Name})

		state.Groups = append(state.Groups, groupModel{
			ID:                          types.StringValue(utils.Float32PtrToIntString(group.Id.Get())),
			LegacyID:                    types.StringValue(utils.Float32PtrToIntString(group.LegacyId.Get())),
			Name:                        types.StringValue(group.Name),
			Members:                     convertGroupMembersFromAPIToModel(group.Members),
			UniversalAppAccess:          types.StringValue(group.UniversalAppAccess),
			UniversalResourceAccess:     types.StringValue(group.UniversalResourceAccess),
			UniversalWorkflowAccess:     types.StringValue(group.UniversalWorkflowAccess),
			UniversalQueryLibraryAccess: types.StringValue(group.UniversalQueryLibraryAccess),
			UserInvites:                 convertGroupUserInvitesFromAPIToModel(group.UserInvites),
			UserListAccess:              types.BoolValue(group.UserListAccess),
			AuditLogAccess:              types.BoolValue(group.AuditLogAccess),
			UnpublishedReleaseAccess:    types.BoolValue(group.UnpublishedReleaseAccess),
			UsageAnalyticsAccess:        types.BoolValue(group.UsageAnalyticsAccess),
			ThemeAccess:                 types.BoolValue(group.ThemeAccess),
			AccountDetailsAccess:        types.BoolValue(group.AccountDetailsAccess),
			LandingPageAppID:            types.StringPointerValue(group.LandingPageAppId.Get()),
			CreatedAt:                   types.StringValue(group.CreatedAt),
			UpdatedAt:                   types.StringValue(group.UpdatedAt),
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

func convertGroupMembersFromAPIToModel(members []api.GroupsGroupIdGet200ResponseDataMembersInner) []groupMemberModel {
	var result []groupMemberModel
	for _, member := range members {
		result = append(result, groupMemberModel{
			ID:           types.StringValue(member.Id),
			Email:        types.StringValue(member.Email),
			IsGroupAdmin: types.BoolValue(member.IsGroupAdmin),
		})
	}
	return result
}

func convertGroupUserInvitesFromAPIToModel(invites []api.GroupsGroupIdGet200ResponseDataUserInvitesInner) []groupUserInvitesModel {
	var result []groupUserInvitesModel
	for _, invite := range invites {
		result = append(result, groupUserInvitesModel{
			ID:           types.StringValue(utils.Float32PtrToIntString(&invite.Id)),
			LegacyID:     types.StringValue(utils.Float32PtrToIntString(&invite.LegacyId)),
			InvitedByID:  types.StringValue(invite.InvitedBy),
			InvitedEmail: types.StringValue(invite.InvitedEmail),
			ExpiresAt:    types.StringValue(invite.ExpiresAt),
			ClaimedBy:    types.StringPointerValue(invite.ClaimedBy.Get()),
			ClaimedAt:    types.StringPointerValue(invite.ClaimedAt.Get()),
			UserType:     types.StringPointerValue(invite.UserType.Get()),
			CreatedAt:    types.StringValue(invite.CreatedAt),
			InviteLink:   types.StringPointerValue(invite.InviteLink),
		})
	}
	return result
}
