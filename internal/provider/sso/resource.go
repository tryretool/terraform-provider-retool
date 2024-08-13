// Package sso provides implementation of the SSO resource.
package sso

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/tryretool/terraform-provider-retool/internal/provider/utils"
	"github.com/tryretool/terraform-provider-retool/internal/sdk/api"
)

// Ensure SSOResource implements the tfsdk.Resource interface.
var (
	_ resource.Resource              = &ssoResource{}
	_ resource.ResourceWithConfigure = &ssoResource{}
)

// ssoResource schema structure.
type ssoResource struct {
	client *api.APIClient
}

type googleConfigModel struct {
	ClientID              types.String `tfsdk:"client_id"`
	ClientSecret          types.String `tfsdk:"client_secret"`
	EncryptedClientSecret types.String `tfsdk:"encrypted_client_secret"`
}

func (m googleConfigModel) attributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"client_id":               types.StringType,
		"client_secret":           types.StringType,
		"encrypted_client_secret": types.StringType,
	}
}

type oidcConfigModel struct {
	ClientID                  types.String `tfsdk:"client_id"`
	ClientSecret              types.String `tfsdk:"client_secret"`
	EncryptedClientSecret     types.String `tfsdk:"encrypted_client_secret"`
	Scopes                    types.String `tfsdk:"scopes"`
	AuthURL                   types.String `tfsdk:"auth_url"`
	TokenURL                  types.String `tfsdk:"token_url"`
	UserInfoURL               types.String `tfsdk:"userinfo_url"`
	Audience                  types.String `tfsdk:"audience"`
	JWTEmailKey               types.String `tfsdk:"jwt_email_key"`
	JWTRolesKey               types.String `tfsdk:"jwt_roles_key"`
	JWTFirstNameKey           types.String `tfsdk:"jwt_first_name_key"`
	JWTLastNameKey            types.String `tfsdk:"jwt_last_name_key"`
	RolesMapping              types.Map    `tfsdk:"roles_mapping"`
	TriggerLoginAutomatically types.Bool   `tfsdk:"trigger_login_automatically"`
	RestrictedDomains         types.List   `tfsdk:"restricted_domains"`
	JITEnabled                types.Bool   `tfsdk:"jit_enabled"`
}

func (m oidcConfigModel) attributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"client_id":                   types.StringType,
		"client_secret":               types.StringType,
		"encrypted_client_secret":     types.StringType,
		"scopes":                      types.StringType,
		"auth_url":                    types.StringType,
		"token_url":                   types.StringType,
		"userinfo_url":                types.StringType,
		"audience":                    types.StringType,
		"jwt_email_key":               types.StringType,
		"jwt_roles_key":               types.StringType,
		"jwt_first_name_key":          types.StringType,
		"jwt_last_name_key":           types.StringType,
		"roles_mapping":               types.MapType{ElemType: types.StringType},
		"trigger_login_automatically": types.BoolType,
		"restricted_domains":          types.ListType{ElemType: types.StringType},
		"jit_enabled":                 types.BoolType,
	}
}

type ldapConfigModel struct {
	ServerURL                  types.String `tfsdk:"server_url"`
	BaseDomainComponents       types.String `tfsdk:"base_domain_components"`
	ServerName                 types.String `tfsdk:"server_name"`
	ServerKey                  types.String `tfsdk:"server_key"`
	EncryptedServerKey         types.String `tfsdk:"encrypted_server_key"`
	ServerCertificate          types.String `tfsdk:"server_certificate"`
	EncryptedServerCertificate types.String `tfsdk:"encrypted_server_certificate"`
}

func (m ldapConfigModel) attributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"server_url":                   types.StringType,
		"base_domain_components":       types.StringType,
		"server_name":                  types.StringType,
		"server_key":                   types.StringType,
		"encrypted_server_key":         types.StringType,
		"server_certificate":           types.StringType,
		"encrypted_server_certificate": types.StringType,
	}
}

type samlConfigModel struct {
	IDPMetadataXML            types.String `tfsdk:"idp_metadata_xml"`
	FirstNameAttribute        types.String `tfsdk:"first_name_attribute"`
	LastNameAttribute         types.String `tfsdk:"last_name_attribute"`
	GroupsAttribute           types.String `tfsdk:"groups_attribute"`
	SyncGroupClaims           types.Bool   `tfsdk:"sync_group_claims"`
	RolesMapping              types.Map    `tfsdk:"roles_mapping"`
	LDAPSyncGroupClaims       types.Bool   `tfsdk:"ldap_sync_group_claims"`
	LDAPConfig                types.Object `tfsdk:"ldap_config"`
	TriggerLoginAutomatically types.Bool   `tfsdk:"trigger_login_automatically"`
	RestrictedDomains         types.List   `tfsdk:"restricted_domains"`
	JITEnabled                types.Bool   `tfsdk:"jit_enabled"`
}

func (m samlConfigModel) attributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"idp_metadata_xml":            types.StringType,
		"first_name_attribute":        types.StringType,
		"last_name_attribute":         types.StringType,
		"groups_attribute":            types.StringType,
		"sync_group_claims":           types.BoolType,
		"roles_mapping":               types.MapType{ElemType: types.ListType{ElemType: types.StringType}},
		"ldap_sync_group_claims":      types.BoolType,
		"ldap_config":                 types.ObjectType{AttrTypes: ldapConfigModel{}.attributeTypes()},
		"trigger_login_automatically": types.BoolType,
		"restricted_domains":          types.ListType{ElemType: types.StringType},
		"jit_enabled":                 types.BoolType,
	}
}

type ssoResourceModel struct {
	Google                    types.Object `tfsdk:"google"`
	OIDC                      types.Object `tfsdk:"oidc"`
	SAML                      types.Object `tfsdk:"saml"`
	DisableEmailPasswordLogin types.Bool   `tfsdk:"disable_email_password_login"`
}

func NewResource() resource.Resource {
	return &ssoResource{}
}

// Configure adds the provider configured client to the resource.
func (r *ssoResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	providerData, ok := req.ProviderData.(*utils.ProviderData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			"Expected *utils.ProviderData, got: %T. Please report this issue to the provider developers.",
		)
	}
	r.client = providerData.Client
}

// Metadata associated with the SSO resource.
func (r *ssoResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sso"
}

// Schema returns the schema for the SSO resource.
func (r *ssoResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"google": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "Google SSO settings",
				Attributes: map[string]schema.Attribute{
					"client_id": schema.StringAttribute{
						Required:    true,
						Description: "Google client ID",
					},
					"client_secret": schema.StringAttribute{
						Required:    true,
						Description: "Google client secret",
						Sensitive:   true,
					},
					"encrypted_client_secret": schema.StringAttribute{
						Computed:    true,
						Description: "Encrypted Google client secret",
						Sensitive:   true,
					},
				},
			},
			"oidc": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "OpenID SSO settings",
				Attributes: map[string]schema.Attribute{
					"client_id": schema.StringAttribute{
						Required:    true,
						Description: "OpenID client ID",
					},
					"client_secret": schema.StringAttribute{
						Required:    true,
						Description: "OpenID client secret",
						Sensitive:   true,
					},
					"encrypted_client_secret": schema.StringAttribute{
						Computed:    true,
						Description: "Encrypted OpenID client secret",
						Sensitive:   true,
					},
					"scopes": schema.StringAttribute{
						Required:    true,
						Description: "OpenID scopes",
					},
					"auth_url": schema.StringAttribute{
						Required:    true,
						Description: "OpenID Auth URL",
					},
					"token_url": schema.StringAttribute{
						Required:    true,
						Description: "OpenID Token URL",
					},
					"userinfo_url": schema.StringAttribute{
						Optional:    true,
						Description: "OpenID User Info URL (Fat token URL)",
					},
					"audience": schema.StringAttribute{
						Optional:    true,
						Description: "SSO Audience",
					},
					"jwt_email_key": schema.StringAttribute{
						Required:    true,
						Description: "Email Key in JWT Token",
					},
					"jwt_roles_key": schema.StringAttribute{
						Optional:    true,
						Description: "Role Key in JWT Token",
					},
					"jwt_first_name_key": schema.StringAttribute{
						Required:    true,
						Description: "First Name Key in JWT Token",
					},
					"jwt_last_name_key": schema.StringAttribute{
						Required:    true,
						Description: "Last Name Key in JWT Token",
					},
					"roles_mapping": schema.MapAttribute{
						Optional:    true,
						Description: "Role Mapping",
						ElementType: types.StringType,
					},
					"trigger_login_automatically": schema.BoolAttribute{
						Optional:    true,
						Computed:    true,
						Default:     booldefault.StaticBool(false),
						Description: "Trigger Login Automatically",
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"restricted_domains": schema.ListAttribute{
						Optional:    true,
						Description: "Restricted Domains",
						ElementType: types.StringType,
					},
					"jit_enabled": schema.BoolAttribute{
						Optional:    true,
						Computed:    true,
						Default:     booldefault.StaticBool(false),
						Description: "Enable Just-In-Time User Provisioning",
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"saml": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "SAML SSO settings",
				Attributes: map[string]schema.Attribute{
					"idp_metadata_xml": schema.StringAttribute{
						Required:    true,
						Description: "IdP Metadata XML",
					},
					"first_name_attribute": schema.StringAttribute{
						Required:    true,
						Description: "First Name Attribute",
					},
					"last_name_attribute": schema.StringAttribute{
						Required:    true,
						Description: "Last Name Attribute",
					},
					"groups_attribute": schema.StringAttribute{
						Optional:    true,
						Description: "Groups Attribute",
					},
					"sync_group_claims": schema.BoolAttribute{
						Optional:    true,
						Computed:    true,
						Default:     booldefault.StaticBool(false),
						Description: "Sync Group Claims",
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"roles_mapping": schema.MapAttribute{
						Optional:    true,
						Description: "Role Mapping",
						ElementType: types.ListType{ElemType: types.StringType},
					},
					"ldap_sync_group_claims": schema.BoolAttribute{
						Optional:    true,
						Computed:    true,
						Default:     booldefault.StaticBool(false),
						Description: "LDAP Sync Group Claims",
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"ldap_config": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "LDAP Configuration",
						Attributes: map[string]schema.Attribute{
							"server_url": schema.StringAttribute{
								Required:    true,
								Description: "Server URL",
							},
							"base_domain_components": schema.StringAttribute{
								Required:    true,
								Description: "Base Domain Components",
							},
							"server_name": schema.StringAttribute{
								Required:    true,
								Description: "Server Name",
							},
							"server_key": schema.StringAttribute{
								Required:    true,
								Description: "Server Key",
								Sensitive:   true,
							},
							"encrypted_server_key": schema.StringAttribute{
								Computed:    true,
								Description: "Encrypted Server Key",
								Sensitive:   true,
							},
							"server_certificate": schema.StringAttribute{
								Required:    true,
								Description: "Server Certificate",
								Sensitive:   true,
							},
							"encrypted_server_certificate": schema.StringAttribute{
								Computed:    true,
								Description: "Encrypted Server Certificate",
								Sensitive:   true,
							},
						},
					},
					"trigger_login_automatically": schema.BoolAttribute{
						Optional:    true,
						Computed:    true,
						Default:     booldefault.StaticBool(false),
						Description: "Trigger Login Automatically",
					},
					"restricted_domains": schema.ListAttribute{
						Optional:    true,
						Description: "Restricted Domains",
						ElementType: types.StringType,
					},
					"jit_enabled": schema.BoolAttribute{
						Optional:    true,
						Computed:    true,
						Default:     booldefault.StaticBool(false),
						Description: "Enable Just-In-Time User Provisioning",
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"disable_email_password_login": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: "Disable email/password login",
				Default:     booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

// Custom validation implementation to prevent common errors in the SSO config.
func (r *ssoResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var model ssoResourceModel
	diags := req.Config.Get(ctx, &model)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Either google, oidc, or saml must be configured.
	if utils.IsEmptyObject(model.Google) && utils.IsEmptyObject(model.OIDC) && utils.IsEmptyObject(model.SAML) {
		resp.Diagnostics.AddError("sso_config_validation", "Either google, oidc, or saml must be configured")
		return
	}

	// You can't configure both SAML and OIDC at the same time.
	if !utils.IsEmptyObject(model.OIDC) && !utils.IsEmptyObject(model.SAML) {
		resp.Diagnostics.AddError("sso_config_validation", "You can't configure both SAML and OIDC at the same time")
	}
}

// Converts a saml/ldap role mapping from a map to a string
// The encoding is "multivalue", so { "key1": ["value1", "value2"], "key2": ["value3"] } becomes "key1->value1,key1->value2,key2->value3".
func getLdapRolesMapping(ctx context.Context, samlConfig samlConfigModel, diags *diag.Diagnostics) *string {
	if utils.IsEmptyMap(samlConfig.RolesMapping) {
		return nil
	}

	elements := make(map[string]types.List, len(samlConfig.RolesMapping.Elements()))
	d := samlConfig.RolesMapping.ElementsAs(ctx, &elements, false)
	diags.Append(d...)
	if diags.HasError() {
		return nil
	}
	// Iterate over elements, convert each value to a list of strings, and add to the tokens array.
	roleMapping := make([]string, 0, len(elements))
	for key, value := range elements {
		for _, v := range utils.GetStringListValue(ctx, value, diags) {
			roleMapping = append(roleMapping, key+"->"+v)
		}
	}
	if diags.HasError() {
		return nil
	}
	result := strings.Join(roleMapping, ",")
	return &result
}

// Converts OIDC role mapping from a map to a string
// { "key1": "value1", "key2: value2"} is converted to "key1->value1,key2->value2".
func getOidcRolesMapping(ctx context.Context, oidcConfig oidcConfigModel, diags *diag.Diagnostics) *string {
	if utils.IsEmptyMap(oidcConfig.RolesMapping) {
		return nil
	}

	elements := make(map[string]types.String, len(oidcConfig.RolesMapping.Elements()))
	d := oidcConfig.RolesMapping.ElementsAs(ctx, &elements, false)
	diags.Append(d...)
	if diags.HasError() {
		return nil
	}
	// Iterate over elements, convert each value to a string and add it to the list of strings as "key->value".
	roleMapping := make([]string, 0, len(elements))
	for key, value := range elements {
		roleMapping = append(roleMapping, key+"->"+value.ValueString())
	}
	result := strings.Join(roleMapping, ",")
	return &result
}

// Converts RestrictedDomains list attribute to a comma-separated string.
func getRestrictedDomains(ctx context.Context, restrictedDomains types.List, diags *diag.Diagnostics) *string {
	if utils.IsEmptyList(restrictedDomains) {
		return nil
	}

	restrictedDomainsList := utils.GetStringListValue(ctx, restrictedDomains, diags)
	result := strings.Join(restrictedDomainsList, ",")
	return &result
}

type encryptedSecrets struct {
	googleClientSecret    *string
	oidcClientSecret      *string
	ldapServerKey         *string
	ldapServerCertificate *string
}

// Helper method that prepares and makes the API POST request to update the SSO config
// Returns a struct containing the encrypted secrets - these need to be stored in the state.
func (r *ssoResource) updateSSOConfig(ctx context.Context, plan ssoResourceModel, globalDiags *diag.Diagnostics) *encryptedSecrets {
	// Generate API request body from plan.
	var apiRequest api.SsoConfigPostRequest

	// Request object is a union of 5 structs - one for each SSO type
	// Since the structs don't share any common interfaces, we need to fill each of them separately, meaning that the code for Google is repeated 3 times, and for OIDC and SAML 2 times. I'm sorry.
	if !utils.IsEmptyObject(plan.Google) {
		tflog.Info(ctx, "Google SSO settings detected")
		if !utils.IsEmptyObject(plan.SAML) {
			tflog.Info(ctx, "Google & SAML SSO settings detected")
			// SSO type is "google & saml"
			// Get google and saml configs.
			var googleConfig googleConfigModel
			diags := plan.Google.As(ctx, &googleConfig, basetypes.ObjectAsOptions{})
			globalDiags.Append(diags...)
			if globalDiags.HasError() {
				return nil
			}
			var samlConfig samlConfigModel
			diags = plan.SAML.As(ctx, &samlConfig, basetypes.ObjectAsOptions{})
			globalDiags.Append(diags...)
			if globalDiags.HasError() {
				return nil
			}
			ssoConfig := api.SsoConfigPostRequestDataOneOf4{
				ConfigType:                "google & saml",
				GoogleClientId:            googleConfig.ClientID.ValueString(),
				GoogleClientSecret:        googleConfig.ClientSecret.ValueString(),
				DisableEmailPasswordLogin: plan.DisableEmailPasswordLogin.ValueBool(),
				IdpMetadataXml:            samlConfig.IDPMetadataXML.ValueString(),
				SamlFirstNameAttribute:    samlConfig.FirstNameAttribute.ValueString(),
				SamlLastNameAttribute:     samlConfig.LastNameAttribute.ValueString(),
				SamlGroupsAttribute:       samlConfig.GroupsAttribute.ValueStringPointer(),
				SamlSyncGroupClaims:       samlConfig.SyncGroupClaims.ValueBool(),
				LdapSyncGroupClaims:       samlConfig.LDAPSyncGroupClaims.ValueBoolPointer(),
				LdapRoleMapping:           getLdapRolesMapping(ctx, samlConfig, globalDiags),
				JitEnabled:                samlConfig.JITEnabled.ValueBool(),
				TriggerLoginAutomatically: samlConfig.TriggerLoginAutomatically.ValueBool(),
				RestrictedDomain:          getRestrictedDomains(ctx, samlConfig.RestrictedDomains, globalDiags),
			}
			// Add ldap config if present.
			if !utils.IsEmptyObject(samlConfig.LDAPConfig) {
				var ldapConfig ldapConfigModel
				diags = samlConfig.LDAPConfig.As(ctx, &ldapConfig, basetypes.ObjectAsOptions{})
				globalDiags.Append(diags...)
				if globalDiags.HasError() {
					return nil
				}
				ssoConfig.LdapServerUrl = ldapConfig.ServerURL.ValueStringPointer()
				ssoConfig.LdapBaseDomainComponents = ldapConfig.BaseDomainComponents.ValueStringPointer()
				ssoConfig.LdapServerName = ldapConfig.ServerName.ValueStringPointer()
				ssoConfig.LdapServerKey = ldapConfig.ServerKey.ValueStringPointer()
				ssoConfig.LdapServerCertificate = ldapConfig.ServerCertificate.ValueStringPointer()
			}
			if globalDiags.HasError() {
				return nil
			}
			apiRequest.Data.SsoConfigPostRequestDataOneOf4 = &ssoConfig
		} else if !utils.IsEmptyObject(plan.OIDC) {
			// SSO type is "google & oidc"
			// Get google and oidc configs.
			var googleConfig googleConfigModel
			diags := plan.Google.As(ctx, &googleConfig, basetypes.ObjectAsOptions{})
			globalDiags.Append(diags...)
			if globalDiags.HasError() {
				return nil
			}
			var oidcConfig oidcConfigModel
			diags = plan.OIDC.As(ctx, &oidcConfig, basetypes.ObjectAsOptions{})
			globalDiags.Append(diags...)
			if globalDiags.HasError() {
				return nil
			}
			ssoConfig := api.SsoConfigPostRequestDataOneOf2{
				ConfigType:                "google & oidc",
				GoogleClientId:            googleConfig.ClientID.ValueString(),
				GoogleClientSecret:        googleConfig.ClientSecret.ValueString(),
				DisableEmailPasswordLogin: plan.DisableEmailPasswordLogin.ValueBool(),
				OidcClientId:              oidcConfig.ClientID.ValueString(),
				OidcClientSecret:          oidcConfig.ClientSecret.ValueString(),
				OidcScopes:                oidcConfig.Scopes.ValueString(),
				OidcAuthUrl:               oidcConfig.AuthURL.ValueString(),
				OidcTokenUrl:              oidcConfig.TokenURL.ValueString(),
				OidcUserinfoUrl:           oidcConfig.UserInfoURL.ValueStringPointer(),
				OidcAudience:              oidcConfig.Audience.ValueStringPointer(),
				JwtEmailKey:               oidcConfig.JWTEmailKey.ValueString(),
				JwtRolesKey:               oidcConfig.JWTRolesKey.ValueStringPointer(),
				JwtFirstNameKey:           oidcConfig.JWTFirstNameKey.ValueString(),
				JwtLastNameKey:            oidcConfig.JWTLastNameKey.ValueString(),
				RolesMapping:              getOidcRolesMapping(ctx, oidcConfig, globalDiags),
				JitEnabled:                oidcConfig.JITEnabled.ValueBool(),
				TriggerLoginAutomatically: oidcConfig.TriggerLoginAutomatically.ValueBool(),
				RestrictedDomain:          getRestrictedDomains(ctx, oidcConfig.RestrictedDomains, globalDiags),
			}
			if globalDiags.HasError() {
				return nil
			}
			apiRequest.Data.SsoConfigPostRequestDataOneOf2 = &ssoConfig
		} else {
			// SSO type is "google".
			var googleConfig googleConfigModel
			diags := plan.Google.As(ctx, &googleConfig, basetypes.ObjectAsOptions{})
			globalDiags.Append(diags...)
			if globalDiags.HasError() {
				return nil
			}
			ssoConfig := api.SsoConfigPostRequestDataOneOf{
				ConfigType:                "google",
				GoogleClientId:            googleConfig.ClientID.ValueString(),
				GoogleClientSecret:        googleConfig.ClientSecret.ValueString(),
				DisableEmailPasswordLogin: plan.DisableEmailPasswordLogin.ValueBool(),
			}
			if globalDiags.HasError() {
				return nil
			}
			apiRequest.Data.SsoConfigPostRequestDataOneOf = &ssoConfig
		}
	} else if !utils.IsEmptyObject(plan.OIDC) {
		// SSO type is "oidc".
		var oidcConfig oidcConfigModel
		diags := plan.OIDC.As(ctx, &oidcConfig, basetypes.ObjectAsOptions{})
		globalDiags.Append(diags...)
		if globalDiags.HasError() {
			return nil
		}
		ssoConfig := api.SsoConfigPostRequestDataOneOf1{
			ConfigType:                "oidc",
			DisableEmailPasswordLogin: plan.DisableEmailPasswordLogin.ValueBool(),
			OidcClientId:              oidcConfig.ClientID.ValueString(),
			OidcClientSecret:          oidcConfig.ClientSecret.ValueString(),
			OidcScopes:                oidcConfig.Scopes.ValueString(),
			OidcAuthUrl:               oidcConfig.AuthURL.ValueString(),
			OidcTokenUrl:              oidcConfig.TokenURL.ValueString(),
			OidcUserinfoUrl:           oidcConfig.UserInfoURL.ValueStringPointer(),
			OidcAudience:              oidcConfig.Audience.ValueStringPointer(),
			JwtEmailKey:               oidcConfig.JWTEmailKey.ValueString(),
			JwtRolesKey:               oidcConfig.JWTRolesKey.ValueStringPointer(),
			JwtFirstNameKey:           oidcConfig.JWTFirstNameKey.ValueString(),
			JwtLastNameKey:            oidcConfig.JWTLastNameKey.ValueString(),
			RolesMapping:              getOidcRolesMapping(ctx, oidcConfig, globalDiags),
			JitEnabled:                oidcConfig.JITEnabled.ValueBool(),
			TriggerLoginAutomatically: oidcConfig.TriggerLoginAutomatically.ValueBool(),
			RestrictedDomain:          getRestrictedDomains(ctx, oidcConfig.RestrictedDomains, globalDiags),
		}
		if globalDiags.HasError() {
			return nil
		}
		apiRequest.Data.SsoConfigPostRequestDataOneOf1 = &ssoConfig
	} else if !utils.IsEmptyObject(plan.SAML) {
		// SSO type is "saml".
		var samlConfig samlConfigModel
		diags := plan.SAML.As(ctx, &samlConfig, basetypes.ObjectAsOptions{})
		globalDiags.Append(diags...)
		if globalDiags.HasError() {
			return nil
		}
		ssoConfig := api.SsoConfigPostRequestDataOneOf3{
			ConfigType:                "saml",
			DisableEmailPasswordLogin: plan.DisableEmailPasswordLogin.ValueBool(),
			IdpMetadataXml:            samlConfig.IDPMetadataXML.ValueString(),
			SamlFirstNameAttribute:    samlConfig.FirstNameAttribute.ValueString(),
			SamlLastNameAttribute:     samlConfig.LastNameAttribute.ValueString(),
			SamlGroupsAttribute:       samlConfig.GroupsAttribute.ValueStringPointer(),
			SamlSyncGroupClaims:       samlConfig.SyncGroupClaims.ValueBool(),
			LdapSyncGroupClaims:       samlConfig.LDAPSyncGroupClaims.ValueBoolPointer(),
			LdapRoleMapping:           getLdapRolesMapping(ctx, samlConfig, globalDiags),
			JitEnabled:                samlConfig.JITEnabled.ValueBool(),
			TriggerLoginAutomatically: samlConfig.TriggerLoginAutomatically.ValueBool(),
			RestrictedDomain:          getRestrictedDomains(ctx, samlConfig.RestrictedDomains, globalDiags),
		}
		// Add ldap config if present.
		if !utils.IsEmptyObject(samlConfig.LDAPConfig) {
			var ldapConfig ldapConfigModel
			diags = samlConfig.LDAPConfig.As(ctx, &ldapConfig, basetypes.ObjectAsOptions{})
			globalDiags.Append(diags...)
			if globalDiags.HasError() {
				return nil
			}
			ssoConfig.LdapServerUrl = ldapConfig.ServerURL.ValueStringPointer()
			ssoConfig.LdapBaseDomainComponents = ldapConfig.BaseDomainComponents.ValueStringPointer()
			ssoConfig.LdapServerName = ldapConfig.ServerName.ValueStringPointer()
			ssoConfig.LdapServerKey = ldapConfig.ServerKey.ValueStringPointer()
			ssoConfig.LdapServerCertificate = ldapConfig.ServerCertificate.ValueStringPointer()
		}
		if globalDiags.HasError() {
			return nil
		}
		apiRequest.Data.SsoConfigPostRequestDataOneOf3 = &ssoConfig
	} else {
		globalDiags.AddError(
			"Error creating SSO config",
			"One of attributes google, oidc, or saml must be set",
		)
		return nil
	}

	tflog.Info(ctx, "Creating SSO config")
	response, httpResponse, err := r.client.SSOAPI.SsoConfigPost(ctx).SsoConfigPostRequest(apiRequest).Execute()

	if err != nil {
		globalDiags.AddError(
			"Error creating SSO config",
			"Could not create SSO config, unexpected error: "+err.Error(),
		)
		tflog.Error(ctx, "Error creating SSO config", utils.AddHttpStatusCode(map[string]interface{}{"error": err.Error()}, httpResponse))
		return nil
	}
	secrets := &encryptedSecrets{
		googleClientSecret:    response.Data.SSOConfig.GoogleClientSecret,
		oidcClientSecret:      response.Data.SSOConfig.OidcClientSecret,
		ldapServerKey:         response.Data.SSOConfig.LdapServerKey,
		ldapServerCertificate: response.Data.SSOConfig.LdapServerCertificate,
	}
	return secrets
}

// Sets encrypted secrets in the state.
func setEncryptedSecrets(ctx context.Context, state *tfsdk.State, secrets *encryptedSecrets, globalDiags *diag.Diagnostics) {
	var obj types.Object
	if diags := state.GetAttribute(ctx, path.Root("google"), &obj); !diags.HasError() && !obj.IsNull() {
		diags = state.SetAttribute(ctx, path.Root("google").AtName("encrypted_client_secret"), types.StringPointerValue(secrets.googleClientSecret))
		globalDiags.Append(diags...)
	}
	if diags := state.GetAttribute(ctx, path.Root("oidc"), &obj); !diags.HasError() && !obj.IsNull() {
		diags = state.SetAttribute(ctx, path.Root("oidc").AtName("encrypted_client_secret"), types.StringPointerValue(secrets.oidcClientSecret))
		globalDiags.Append(diags...)
	}
	if diags := state.GetAttribute(ctx, path.Root("saml").AtName("ldap_config"), &obj); !diags.HasError() && !obj.IsNull() {
		diags = state.SetAttribute(ctx, path.Root("saml").AtName("ldap_config").AtName("encrypted_server_key"), types.StringPointerValue(secrets.ldapServerKey))
		globalDiags.Append(diags...)
		diags = state.SetAttribute(ctx, path.Root("saml").AtName("ldap_config").AtName("encrypted_server_certificate"), types.StringPointerValue(secrets.ldapServerCertificate))
		globalDiags.Append(diags...)
	}
}

// Updates SSO config and sets state.
func (r *ssoResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan.
	tflog.Info(ctx, "Getting the plan model")
	var plan ssoResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Updating SSO config")
	secrets := r.updateSSOConfig(ctx, plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set state to fully populated data.
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Error creating SSO conifg", map[string]interface{}{"error": "Could not set state"})
		return
	}
	// Update encrypted secret strings in the state.
	setEncryptedSecrets(ctx, &resp.State, secrets, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Error creating SSO config", map[string]interface{}{"error": "Could not set encrypted secrets"})
		return
	}

	tflog.Info(ctx, "SSO config created")
}

// Check if encrypted secret matches the state. If it does, we return the secret from the state. If it doesn't, we return null string.
func getSecretValue(ctx context.Context, state *tfsdk.State, newEncryptedSecret types.String, secretPath path.Path, encryptedSecretPath path.Path, globalDiags *diag.Diagnostics) types.String {
	var encryptedSecret types.String
	diags := state.GetAttribute(ctx, encryptedSecretPath, &encryptedSecret)
	globalDiags.Append(diags...)
	if globalDiags.HasError() {
		return types.StringNull()
	}
	if encryptedSecret.Equal(newEncryptedSecret) {
		var secret types.String
		diags = state.GetAttribute(ctx, secretPath, &secret)
		globalDiags.Append(diags...)
		if globalDiags.HasError() {
			return types.StringNull()
		}
		return secret
	}
	return types.StringNull()
}

// Converts a string of the form "key1->value1,key1->value2,key2->value3" to a map of the form { "key1": ["value1", "value2"], "key2": ["value3"] }.
func parseLdapRolesMappingValue(ctx context.Context, rolesMappingValue string) map[string][]string {
	rolesMapping := make(map[string][]string)
	roleMapTuples := strings.Split(rolesMappingValue, ",")
	for _, roleMapping := range roleMapTuples {
		roleMappingParts := strings.Split(roleMapping, "->")
		if len(roleMappingParts) != 2 {
			tflog.Warn(ctx, "Invalid role mapping: "+roleMapping)
			continue
		}
		key := strings.TrimSpace(roleMappingParts[0])
		value := strings.TrimSpace(roleMappingParts[1])
		rolesMapping[key] = append(rolesMapping[key], value)
	}
	return rolesMapping
}

// Check if any of the LDAP config attributes are present.
func isLdapConfigPresent(ssoConfig *api.SSOConfig) bool {
	return ssoConfig.LdapServerUrl != nil ||
		ssoConfig.LdapBaseDomainComponents != nil ||
		ssoConfig.LdapServerName != nil ||
		ssoConfig.LdapServerKey != nil ||
		ssoConfig.LdapServerCertificate != nil
}

// Read SSO config.
func (r *ssoResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ssoResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, httpResponse, err := r.client.SSOAPI.SsoConfigGet(ctx).Execute()
	if err != nil {
		if httpResponse != nil && httpResponse.StatusCode == 404 {
			tflog.Info(ctx, "SSO is not configured")
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Error reading SSO config",
			fmt.Sprintf("Could not read SSO config: %s", err.Error()),
		)
		tflog.Error(ctx, "Error reading SSO config", utils.AddHttpStatusCode(map[string]any{"error": err.Error()}, httpResponse))
		return
	}

	state.DisableEmailPasswordLogin = types.BoolValue(response.Data.SSOConfig.DisableEmailPasswordLogin)
	if response.Data.SSOConfig.ConfigType == "google" || response.Data.SSOConfig.ConfigType == "google & oidc" || response.Data.SSOConfig.ConfigType == "google & saml" {
		googleConfig := googleConfigModel{
			ClientID:              types.StringPointerValue(response.Data.SSOConfig.GoogleClientId),
			EncryptedClientSecret: types.StringPointerValue(response.Data.SSOConfig.GoogleClientSecret),
		}
		googleConfig.ClientSecret = getSecretValue(
			ctx,
			&resp.State,
			googleConfig.EncryptedClientSecret,
			path.Root("google").AtName("client_secret"),
			path.Root("google").AtName("encrypted_client_secret"),
			&resp.Diagnostics)
		if resp.Diagnostics.HasError() {
			return
		}
		googleConfigObject, diags := types.ObjectValueFrom(ctx, googleConfig.attributeTypes(), googleConfig)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		state.Google = googleConfigObject
	}
	if response.Data.SSOConfig.ConfigType == "oidc" || response.Data.SSOConfig.ConfigType == "google & oidc" {
		oidcConfig := oidcConfigModel{
			ClientID:                  types.StringPointerValue(response.Data.SSOConfig.OidcClientId),
			EncryptedClientSecret:     types.StringPointerValue(response.Data.SSOConfig.OidcClientSecret),
			Scopes:                    types.StringPointerValue(response.Data.SSOConfig.OidcScopes),
			AuthURL:                   types.StringPointerValue(response.Data.SSOConfig.OidcAuthUrl),
			TokenURL:                  types.StringPointerValue(response.Data.SSOConfig.OidcTokenUrl),
			UserInfoURL:               types.StringPointerValue(response.Data.SSOConfig.OidcUserinfoUrl),
			Audience:                  types.StringPointerValue(response.Data.SSOConfig.OidcAudience),
			JWTEmailKey:               types.StringPointerValue(response.Data.SSOConfig.JwtEmailKey),
			JWTRolesKey:               types.StringPointerValue(response.Data.SSOConfig.JwtRolesKey),
			JWTFirstNameKey:           types.StringPointerValue(response.Data.SSOConfig.JwtFirstNameKey),
			JWTLastNameKey:            types.StringPointerValue(response.Data.SSOConfig.JwtLastNameKey),
			TriggerLoginAutomatically: types.BoolPointerValue(response.Data.SSOConfig.TriggerLoginAutomatically),
			JITEnabled:                types.BoolPointerValue(response.Data.SSOConfig.JitEnabled),
		}
		oidcConfig.ClientSecret = getSecretValue(
			ctx,
			&resp.State,
			oidcConfig.EncryptedClientSecret,
			path.Root("oidc").AtName("client_secret"),
			path.Root("oidc").AtName("encrypted_client_secret"),
			&resp.Diagnostics)
		if resp.Diagnostics.HasError() {
			return
		}

		if response.Data.SSOConfig.RolesMapping != nil {
			tflog.Info(ctx, "Roles mapping detected"+*response.Data.SSOConfig.RolesMapping)
			roleMapTuples := strings.Split(*response.Data.SSOConfig.RolesMapping, ",")
			rolesMapping := make(map[string]types.String, len(roleMapTuples))
			for _, roleMapping := range roleMapTuples {
				roleMappingParts := strings.Split(roleMapping, "->")
				if len(roleMappingParts) != 2 {
					tflog.Warn(ctx, "Invalid role mapping: "+roleMapping)
					continue
				}
				rolesMapping[strings.TrimSpace(roleMappingParts[0])] = types.StringValue(strings.TrimSpace(roleMappingParts[1]))
			}
			oidcConfig.RolesMapping, diags = types.MapValueFrom(ctx, types.StringType, rolesMapping)
			resp.Diagnostics.Append(diags...)
		} else {
			oidcConfig.RolesMapping = types.MapNull(types.StringType)
		}
		if response.Data.SSOConfig.RestrictedDomain != nil {
			restrictedDomains := strings.Split(*response.Data.SSOConfig.RestrictedDomain, ",")
			oidcConfig.RestrictedDomains, diags = types.ListValueFrom(ctx, types.StringType, restrictedDomains)
			resp.Diagnostics.Append(diags...)
		} else {
			oidcConfig.RestrictedDomains = types.ListNull(types.StringType)
		}

		oidcConfigObject, diags := types.ObjectValueFrom(ctx, oidcConfig.attributeTypes(), oidcConfig)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		state.OIDC = oidcConfigObject
	}
	if response.Data.SSOConfig.ConfigType == "saml" || response.Data.SSOConfig.ConfigType == "google & saml" {
		samlConfig := samlConfigModel{
			IDPMetadataXML:            types.StringPointerValue(response.Data.SSOConfig.IdpMetadataXml),
			FirstNameAttribute:        types.StringPointerValue(response.Data.SSOConfig.SamlFirstNameAttribute),
			LastNameAttribute:         types.StringPointerValue(response.Data.SSOConfig.SamlLastNameAttribute),
			GroupsAttribute:           types.StringPointerValue(response.Data.SSOConfig.SamlGroupsAttribute),
			SyncGroupClaims:           types.BoolPointerValue(response.Data.SSOConfig.SamlSyncGroupClaims),
			LDAPSyncGroupClaims:       types.BoolPointerValue(response.Data.SSOConfig.LdapSyncGroupClaims),
			JITEnabled:                types.BoolPointerValue(response.Data.SSOConfig.JitEnabled),
			TriggerLoginAutomatically: types.BoolPointerValue(response.Data.SSOConfig.TriggerLoginAutomatically),
		}
		if response.Data.SSOConfig.LdapRoleMapping != nil {
			rolesMapping := parseLdapRolesMappingValue(ctx, *response.Data.SSOConfig.LdapRoleMapping)
			rolesMappingTF := make(map[string]types.List, len(rolesMapping))
			for key, value := range rolesMapping {
				rolesMappingTF[key], diags = types.ListValueFrom(ctx, types.StringType, value)
				resp.Diagnostics.Append(diags...)
			}
			samlConfig.RolesMapping, diags = types.MapValueFrom(ctx, types.ListType{ElemType: types.StringType}, rolesMappingTF)
			resp.Diagnostics.Append(diags...)
		} else {
			samlConfig.RolesMapping = types.MapNull(types.ListType{ElemType: types.StringType})
		}
		if response.Data.SSOConfig.RestrictedDomain != nil {
			restrictedDomains := strings.Split(*response.Data.SSOConfig.RestrictedDomain, ",")
			samlConfig.RestrictedDomains, diags = types.ListValueFrom(ctx, types.StringType, restrictedDomains)
			resp.Diagnostics.Append(diags...)
		} else {
			samlConfig.RestrictedDomains = types.ListNull(types.StringType)
		}
		if isLdapConfigPresent(response.Data.SSOConfig) {
			ldapConfig := ldapConfigModel{
				ServerURL:                  types.StringPointerValue(response.Data.SSOConfig.LdapServerUrl),
				BaseDomainComponents:       types.StringPointerValue(response.Data.SSOConfig.LdapBaseDomainComponents),
				ServerName:                 types.StringPointerValue(response.Data.SSOConfig.LdapServerName),
				EncryptedServerKey:         types.StringPointerValue(response.Data.SSOConfig.LdapServerKey),
				EncryptedServerCertificate: types.StringPointerValue(response.Data.SSOConfig.LdapServerCertificate),
			}
			ldapConfig.ServerKey = getSecretValue(
				ctx,
				&resp.State,
				ldapConfig.EncryptedServerKey,
				path.Root("saml").AtName("ldap_config").AtName("server_key"),
				path.Root("saml").AtName("ldap_config").AtName("encrypted_server_key"),
				&resp.Diagnostics)
			if resp.Diagnostics.HasError() {
				return
			}
			ldapConfig.ServerCertificate = getSecretValue(
				ctx,
				&resp.State,
				ldapConfig.EncryptedServerCertificate,
				path.Root("saml").AtName("ldap_config").AtName("server_certificate"),
				path.Root("saml").AtName("ldap_config").AtName("encrypted_server_certificate"),
				&resp.Diagnostics)
			if resp.Diagnostics.HasError() {
				return
			}

			ldapConfigObject, diags := types.ObjectValueFrom(ctx, ldapConfig.attributeTypes(), ldapConfig)
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}
			samlConfig.LDAPConfig = ldapConfigObject
		} else {
			samlConfig.LDAPConfig = types.ObjectNull(ldapConfigModel{}.attributeTypes())
		}
		samlConfigObject, diags := types.ObjectValueFrom(ctx, samlConfig.attributeTypes(), samlConfig)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		state.SAML = samlConfigObject
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update SSO config.
func (r *ssoResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ssoResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	secrets := r.updateSSOConfig(ctx, plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update encrypted secret strings in the state.
	setEncryptedSecrets(ctx, &resp.State, secrets, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete SSO config.
func (r *ssoResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	httpResponse, err := r.client.SSOAPI.SsoConfigDelete(ctx).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting SSO config",
			"Could not delete SSO config: "+err.Error(),
		)
		tflog.Error(ctx, "Error Deleting SSO config", utils.AddHttpStatusCode(map[string]any{"error": err.Error()}, httpResponse))
		return
	}
}
