/*
Retool API

Go to Settings > API to get started. Once you generate an API token, use bearer token authentication to make requests.

API version: 2.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SSOConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SSOConfig{}

// SSOConfig struct for SSOConfig
type SSOConfig struct {
	ConfigType string `json:"config_type"`
	GoogleClientId *string `json:"google_client_id,omitempty"`
	GoogleClientSecret *string `json:"google_client_secret,omitempty"`
	DisableEmailPasswordLogin bool `json:"disable_email_password_login"`
	OidcClientId *string `json:"oidc_client_id,omitempty"`
	OidcClientSecret *string `json:"oidc_client_secret,omitempty"`
	OidcScopes *string `json:"oidc_scopes,omitempty"`
	OidcAuthUrl *string `json:"oidc_auth_url,omitempty"`
	OidcTokenUrl *string `json:"oidc_token_url,omitempty"`
	OidcUserinfoUrl *string `json:"oidc_userinfo_url,omitempty"`
	OidcAudience *string `json:"oidc_audience,omitempty"`
	JwtEmailKey *string `json:"jwt_email_key,omitempty"`
	JwtRolesKey *string `json:"jwt_roles_key,omitempty"`
	JwtFirstNameKey *string `json:"jwt_first_name_key,omitempty"`
	JwtLastNameKey *string `json:"jwt_last_name_key,omitempty"`
	RolesMapping *string `json:"roles_mapping,omitempty"`
	JitEnabled *bool `json:"jit_enabled,omitempty"`
	RestrictedDomain *string `json:"restricted_domain,omitempty"`
	TriggerLoginAutomatically *bool `json:"trigger_login_automatically,omitempty"`
	IdpMetadataXml *string `json:"idp_metadata_xml,omitempty"`
	SamlFirstNameAttribute *string `json:"saml_first_name_attribute,omitempty"`
	SamlLastNameAttribute *string `json:"saml_last_name_attribute,omitempty"`
	SamlGroupsAttribute *string `json:"saml_groups_attribute,omitempty"`
	SamlSyncGroupClaims *bool `json:"saml_sync_group_claims,omitempty"`
	LdapSyncGroupClaims *bool `json:"ldap_sync_group_claims,omitempty"`
	LdapRoleMapping *string `json:"ldap_role_mapping,omitempty"`
	LdapServerUrl *string `json:"ldap_server_url,omitempty"`
	LdapBaseDomainComponents *string `json:"ldap_base_domain_components,omitempty"`
	LdapServerName *string `json:"ldap_server_name,omitempty"`
	LdapServerKey *string `json:"ldap_server_key,omitempty"`
	LdapServerCertificate *string `json:"ldap_server_certificate,omitempty"`
}

type _SSOConfig SSOConfig

// NewSSOConfig instantiates a new SSOConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSSOConfig(configType string, disableEmailPasswordLogin bool) *SSOConfig {
	this := SSOConfig{}
	this.ConfigType = configType
	this.DisableEmailPasswordLogin = disableEmailPasswordLogin
	return &this
}

// NewSSOConfigWithDefaults instantiates a new SSOConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSSOConfigWithDefaults() *SSOConfig {
	this := SSOConfig{}
	return &this
}

// GetConfigType returns the ConfigType field value
func (o *SSOConfig) GetConfigType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigType
}

// GetConfigTypeOk returns a tuple with the ConfigType field value
// and a boolean to check if the value has been set.
func (o *SSOConfig) GetConfigTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigType, true
}

// SetConfigType sets field value
func (o *SSOConfig) SetConfigType(v string) {
	o.ConfigType = v
}

// GetGoogleClientId returns the GoogleClientId field value if set, zero value otherwise.
func (o *SSOConfig) GetGoogleClientId() string {
	if o == nil || IsNil(o.GoogleClientId) {
		var ret string
		return ret
	}
	return *o.GoogleClientId
}

// GetGoogleClientIdOk returns a tuple with the GoogleClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOConfig) GetGoogleClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.GoogleClientId) {
		return nil, false
	}
	return o.GoogleClientId, true
}

// HasGoogleClientId returns a boolean if a field has been set.
func (o *SSOConfig) HasGoogleClientId() bool {
	if o != nil && !IsNil(o.GoogleClientId) {
		return true
	}

	return false
}

// SetGoogleClientId gets a reference to the given string and assigns it to the GoogleClientId field.
func (o *SSOConfig) SetGoogleClientId(v string) {
	o.GoogleClientId = &v
}

// GetGoogleClientSecret returns the GoogleClientSecret field value if set, zero value otherwise.
func (o *SSOConfig) GetGoogleClientSecret() string {
	if o == nil || IsNil(o.GoogleClientSecret) {
		var ret string
		return ret
	}
	return *o.GoogleClientSecret
}

// GetGoogleClientSecretOk returns a tuple with the GoogleClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOConfig) GetGoogleClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.GoogleClientSecret) {
		return nil, false
	}
	return o.GoogleClientSecret, true
}

// HasGoogleClientSecret returns a boolean if a field has been set.
func (o *SSOConfig) HasGoogleClientSecret() bool {
	if o != nil && !IsNil(o.GoogleClientSecret) {
		return true
	}

	return false
}

// SetGoogleClientSecret gets a reference to the given string and assigns it to the GoogleClientSecret field.
func (o *SSOConfig) SetGoogleClientSecret(v string) {
	o.GoogleClientSecret = &v
}

// GetDisableEmailPasswordLogin returns the DisableEmailPasswordLogin field value
func (o *SSOConfig) GetDisableEmailPasswordLogin() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableEmailPasswordLogin
}

// GetDisableEmailPasswordLoginOk returns a tuple with the DisableEmailPasswordLogin field value
// and a boolean to check if the value has been set.
func (o *SSOConfig) GetDisableEmailPasswordLoginOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableEmailPasswordLogin, true
}

// SetDisableEmailPasswordLogin sets field value
func (o *SSOConfig) SetDisableEmailPasswordLogin(v bool) {
	o.DisableEmailPasswordLogin = v
}

// GetOidcClientId returns the OidcClientId field value if set, zero value otherwise.
func (o *SSOConfig) GetOidcClientId() string {
	if o == nil || IsNil(o.OidcClientId) {
		var ret string
		return ret
	}
	return *o.OidcClientId
}

// GetOidcClientIdOk returns a tuple with the OidcClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOConfig) GetOidcClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.OidcClientId) {
		return nil, false
	}
	return o.OidcClientId, true
}

// HasOidcClientId returns a boolean if a field has been set.
func (o *SSOConfig) HasOidcClientId() bool {
	if o != nil && !IsNil(o.OidcClientId) {
		return true
	}

	return false
}

// SetOidcClientId gets a reference to the given string and assigns it to the OidcClientId field.
func (o *SSOConfig) SetOidcClientId(v string) {
	o.OidcClientId = &v
}

// GetOidcClientSecret returns the OidcClientSecret field value if set, zero value otherwise.
func (o *SSOConfig) GetOidcClientSecret() string {
	if o == nil || IsNil(o.OidcClientSecret) {
		var ret string
		return ret
	}
	return *o.OidcClientSecret
}

// GetOidcClientSecretOk returns a tuple with the OidcClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOConfig) GetOidcClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.OidcClientSecret) {
		return nil, false
	}
	return o.OidcClientSecret, true
}

// HasOidcClientSecret returns a boolean if a field has been set.
func (o *SSOConfig) HasOidcClientSecret() bool {
	if o != nil && !IsNil(o.OidcClientSecret) {
		return true
	}

	return false
}

// SetOidcClientSecret gets a reference to the given string and assigns it to the OidcClientSecret field.
func (o *SSOConfig) SetOidcClientSecret(v string) {
	o.OidcClientSecret = &v
}

// GetOidcScopes returns the OidcScopes field value if set, zero value otherwise.
func (o *SSOConfig) GetOidcScopes() string {
	if o == nil || IsNil(o.OidcScopes) {
		var ret string
		return ret
	}
	return *o.OidcScopes
}

// GetOidcScopesOk returns a tuple with the OidcScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOConfig) GetOidcScopesOk() (*string, bool) {
	if o == nil || IsNil(o.OidcScopes) {
		return nil, false
	}
	return o.OidcScopes, true
}

// HasOidcScopes returns a boolean if a field has been set.
func (o *SSOConfig) HasOidcScopes() bool {
	if o != nil && !IsNil(o.OidcScopes) {
		return true
	}

	return false
}

// SetOidcScopes gets a reference to the given string and assigns it to the OidcScopes field.
func (o *SSOConfig) SetOidcScopes(v string) {
	o.OidcScopes = &v
}

// GetOidcAuthUrl returns the OidcAuthUrl field value if set, zero value otherwise.
func (o *SSOConfig) GetOidcAuthUrl() string {
	if o == nil || IsNil(o.OidcAuthUrl) {
		var ret string
		return ret
	}
	return *o.OidcAuthUrl
}

// GetOidcAuthUrlOk returns a tuple with the OidcAuthUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOConfig) GetOidcAuthUrlOk() (*string, bool) {
	if o == nil || IsNil(o.OidcAuthUrl) {
		return nil, false
	}
	return o.OidcAuthUrl, true
}

// HasOidcAuthUrl returns a boolean if a field has been set.
func (o *SSOConfig) HasOidcAuthUrl() bool {
	if o != nil && !IsNil(o.OidcAuthUrl) {
		return true
	}

	return false
}

// SetOidcAuthUrl gets a reference to the given string and assigns it to the OidcAuthUrl field.
func (o *SSOConfig) SetOidcAuthUrl(v string) {
	o.OidcAuthUrl = &v
}

// GetOidcTokenUrl returns the OidcTokenUrl field value if set, zero value otherwise.
func (o *SSOConfig) GetOidcTokenUrl() string {
	if o == nil || IsNil(o.OidcTokenUrl) {
		var ret string
		return ret
	}
	return *o.OidcTokenUrl
}

// GetOidcTokenUrlOk returns a tuple with the OidcTokenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOConfig) GetOidcTokenUrlOk() (*string, bool) {
	if o == nil || IsNil(o.OidcTokenUrl) {
		return nil, false
	}
	return o.OidcTokenUrl, true
}

// HasOidcTokenUrl returns a boolean if a field has been set.
func (o *SSOConfig) HasOidcTokenUrl() bool {
	if o != nil && !IsNil(o.OidcTokenUrl) {
		return true
	}

	return false
}

// SetOidcTokenUrl gets a reference to the given string and assigns it to the OidcTokenUrl field.
func (o *SSOConfig) SetOidcTokenUrl(v string) {
	o.OidcTokenUrl = &v
}

// GetOidcUserinfoUrl returns the OidcUserinfoUrl field value if set, zero value otherwise.
func (o *SSOConfig) GetOidcUserinfoUrl() string {
	if o == nil || IsNil(o.OidcUserinfoUrl) {
		var ret string
		return ret
	}
	return *o.OidcUserinfoUrl
}

// GetOidcUserinfoUrlOk returns a tuple with the OidcUserinfoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOConfig) GetOidcUserinfoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.OidcUserinfoUrl) {
		return nil, false
	}
	return o.OidcUserinfoUrl, true
}

// HasOidcUserinfoUrl returns a boolean if a field has been set.
func (o *SSOConfig) HasOidcUserinfoUrl() bool {
	if o != nil && !IsNil(o.OidcUserinfoUrl) {
		return true
	}

	return false
}

// SetOidcUserinfoUrl gets a reference to the given string and assigns it to the OidcUserinfoUrl field.
func (o *SSOConfig) SetOidcUserinfoUrl(v string) {
	o.OidcUserinfoUrl = &v
}

// GetOidcAudience returns the OidcAudience field value if set, zero value otherwise.
func (o *SSOConfig) GetOidcAudience() string {
	if o == nil || IsNil(o.OidcAudience) {
		var ret string
		return ret
	}
	return *o.OidcAudience
}

// GetOidcAudienceOk returns a tuple with the OidcAudience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOConfig) GetOidcAudienceOk() (*string, bool) {
	if o == nil || IsNil(o.OidcAudience) {
		return nil, false
	}
	return o.OidcAudience, true
}

// HasOidcAudience returns a boolean if a field has been set.
func (o *SSOConfig) HasOidcAudience() bool {
	if o != nil && !IsNil(o.OidcAudience) {
		return true
	}

	return false
}

// SetOidcAudience gets a reference to the given string and assigns it to the OidcAudience field.
func (o *SSOConfig) SetOidcAudience(v string) {
	o.OidcAudience = &v
}

// GetJwtEmailKey returns the JwtEmailKey field value if set, zero value otherwise.
func (o *SSOConfig) GetJwtEmailKey() string {
	if o == nil || IsNil(o.JwtEmailKey) {
		var ret string
		return ret
	}
	return *o.JwtEmailKey
}

// GetJwtEmailKeyOk returns a tuple with the JwtEmailKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOConfig) GetJwtEmailKeyOk() (*string, bool) {
	if o == nil || IsNil(o.JwtEmailKey) {
		return nil, false
	}
	return o.JwtEmailKey, true
}

// HasJwtEmailKey returns a boolean if a field has been set.
func (o *SSOConfig) HasJwtEmailKey() bool {
	if o != nil && !IsNil(o.JwtEmailKey) {
		return true
	}

	return false
}

// SetJwtEmailKey gets a reference to the given string and assigns it to the JwtEmailKey field.
func (o *SSOConfig) SetJwtEmailKey(v string) {
	o.JwtEmailKey = &v
}

// GetJwtRolesKey returns the JwtRolesKey field value if set, zero value otherwise.
func (o *SSOConfig) GetJwtRolesKey() string {
	if o == nil || IsNil(o.JwtRolesKey) {
		var ret string
		return ret
	}
	return *o.JwtRolesKey
}

// GetJwtRolesKeyOk returns a tuple with the JwtRolesKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOConfig) GetJwtRolesKeyOk() (*string, bool) {
	if o == nil || IsNil(o.JwtRolesKey) {
		return nil, false
	}
	return o.JwtRolesKey, true
}

// HasJwtRolesKey returns a boolean if a field has been set.
func (o *SSOConfig) HasJwtRolesKey() bool {
	if o != nil && !IsNil(o.JwtRolesKey) {
		return true
	}

	return false
}

// SetJwtRolesKey gets a reference to the given string and assigns it to the JwtRolesKey field.
func (o *SSOConfig) SetJwtRolesKey(v string) {
	o.JwtRolesKey = &v
}

// GetJwtFirstNameKey returns the JwtFirstNameKey field value if set, zero value otherwise.
func (o *SSOConfig) GetJwtFirstNameKey() string {
	if o == nil || IsNil(o.JwtFirstNameKey) {
		var ret string
		return ret
	}
	return *o.JwtFirstNameKey
}

// GetJwtFirstNameKeyOk returns a tuple with the JwtFirstNameKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOConfig) GetJwtFirstNameKeyOk() (*string, bool) {
	if o == nil || IsNil(o.JwtFirstNameKey) {
		return nil, false
	}
	return o.JwtFirstNameKey, true
}

// HasJwtFirstNameKey returns a boolean if a field has been set.
func (o *SSOConfig) HasJwtFirstNameKey() bool {
	if o != nil && !IsNil(o.JwtFirstNameKey) {
		return true
	}

	return false
}

// SetJwtFirstNameKey gets a reference to the given string and assigns it to the JwtFirstNameKey field.
func (o *SSOConfig) SetJwtFirstNameKey(v string) {
	o.JwtFirstNameKey = &v
}

// GetJwtLastNameKey returns the JwtLastNameKey field value if set, zero value otherwise.
func (o *SSOConfig) GetJwtLastNameKey() string {
	if o == nil || IsNil(o.JwtLastNameKey) {
		var ret string
		return ret
	}
	return *o.JwtLastNameKey
}

// GetJwtLastNameKeyOk returns a tuple with the JwtLastNameKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOConfig) GetJwtLastNameKeyOk() (*string, bool) {
	if o == nil || IsNil(o.JwtLastNameKey) {
		return nil, false
	}
	return o.JwtLastNameKey, true
}

// HasJwtLastNameKey returns a boolean if a field has been set.
func (o *SSOConfig) HasJwtLastNameKey() bool {
	if o != nil && !IsNil(o.JwtLastNameKey) {
		return true
	}

	return false
}

// SetJwtLastNameKey gets a reference to the given string and assigns it to the JwtLastNameKey field.
func (o *SSOConfig) SetJwtLastNameKey(v string) {
	o.JwtLastNameKey = &v
}

// GetRolesMapping returns the RolesMapping field value if set, zero value otherwise.
func (o *SSOConfig) GetRolesMapping() string {
	if o == nil || IsNil(o.RolesMapping) {
		var ret string
		return ret
	}
	return *o.RolesMapping
}

// GetRolesMappingOk returns a tuple with the RolesMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOConfig) GetRolesMappingOk() (*string, bool) {
	if o == nil || IsNil(o.RolesMapping) {
		return nil, false
	}
	return o.RolesMapping, true
}

// HasRolesMapping returns a boolean if a field has been set.
func (o *SSOConfig) HasRolesMapping() bool {
	if o != nil && !IsNil(o.RolesMapping) {
		return true
	}

	return false
}

// SetRolesMapping gets a reference to the given string and assigns it to the RolesMapping field.
func (o *SSOConfig) SetRolesMapping(v string) {
	o.RolesMapping = &v
}

// GetJitEnabled returns the JitEnabled field value if set, zero value otherwise.
func (o *SSOConfig) GetJitEnabled() bool {
	if o == nil || IsNil(o.JitEnabled) {
		var ret bool
		return ret
	}
	return *o.JitEnabled
}

// GetJitEnabledOk returns a tuple with the JitEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOConfig) GetJitEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.JitEnabled) {
		return nil, false
	}
	return o.JitEnabled, true
}

// HasJitEnabled returns a boolean if a field has been set.
func (o *SSOConfig) HasJitEnabled() bool {
	if o != nil && !IsNil(o.JitEnabled) {
		return true
	}

	return false
}

// SetJitEnabled gets a reference to the given bool and assigns it to the JitEnabled field.
func (o *SSOConfig) SetJitEnabled(v bool) {
	o.JitEnabled = &v
}

// GetRestrictedDomain returns the RestrictedDomain field value if set, zero value otherwise.
func (o *SSOConfig) GetRestrictedDomain() string {
	if o == nil || IsNil(o.RestrictedDomain) {
		var ret string
		return ret
	}
	return *o.RestrictedDomain
}

// GetRestrictedDomainOk returns a tuple with the RestrictedDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOConfig) GetRestrictedDomainOk() (*string, bool) {
	if o == nil || IsNil(o.RestrictedDomain) {
		return nil, false
	}
	return o.RestrictedDomain, true
}

// HasRestrictedDomain returns a boolean if a field has been set.
func (o *SSOConfig) HasRestrictedDomain() bool {
	if o != nil && !IsNil(o.RestrictedDomain) {
		return true
	}

	return false
}

// SetRestrictedDomain gets a reference to the given string and assigns it to the RestrictedDomain field.
func (o *SSOConfig) SetRestrictedDomain(v string) {
	o.RestrictedDomain = &v
}

// GetTriggerLoginAutomatically returns the TriggerLoginAutomatically field value if set, zero value otherwise.
func (o *SSOConfig) GetTriggerLoginAutomatically() bool {
	if o == nil || IsNil(o.TriggerLoginAutomatically) {
		var ret bool
		return ret
	}
	return *o.TriggerLoginAutomatically
}

// GetTriggerLoginAutomaticallyOk returns a tuple with the TriggerLoginAutomatically field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOConfig) GetTriggerLoginAutomaticallyOk() (*bool, bool) {
	if o == nil || IsNil(o.TriggerLoginAutomatically) {
		return nil, false
	}
	return o.TriggerLoginAutomatically, true
}

// HasTriggerLoginAutomatically returns a boolean if a field has been set.
func (o *SSOConfig) HasTriggerLoginAutomatically() bool {
	if o != nil && !IsNil(o.TriggerLoginAutomatically) {
		return true
	}

	return false
}

// SetTriggerLoginAutomatically gets a reference to the given bool and assigns it to the TriggerLoginAutomatically field.
func (o *SSOConfig) SetTriggerLoginAutomatically(v bool) {
	o.TriggerLoginAutomatically = &v
}

// GetIdpMetadataXml returns the IdpMetadataXml field value if set, zero value otherwise.
func (o *SSOConfig) GetIdpMetadataXml() string {
	if o == nil || IsNil(o.IdpMetadataXml) {
		var ret string
		return ret
	}
	return *o.IdpMetadataXml
}

// GetIdpMetadataXmlOk returns a tuple with the IdpMetadataXml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOConfig) GetIdpMetadataXmlOk() (*string, bool) {
	if o == nil || IsNil(o.IdpMetadataXml) {
		return nil, false
	}
	return o.IdpMetadataXml, true
}

// HasIdpMetadataXml returns a boolean if a field has been set.
func (o *SSOConfig) HasIdpMetadataXml() bool {
	if o != nil && !IsNil(o.IdpMetadataXml) {
		return true
	}

	return false
}

// SetIdpMetadataXml gets a reference to the given string and assigns it to the IdpMetadataXml field.
func (o *SSOConfig) SetIdpMetadataXml(v string) {
	o.IdpMetadataXml = &v
}

// GetSamlFirstNameAttribute returns the SamlFirstNameAttribute field value if set, zero value otherwise.
func (o *SSOConfig) GetSamlFirstNameAttribute() string {
	if o == nil || IsNil(o.SamlFirstNameAttribute) {
		var ret string
		return ret
	}
	return *o.SamlFirstNameAttribute
}

// GetSamlFirstNameAttributeOk returns a tuple with the SamlFirstNameAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOConfig) GetSamlFirstNameAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.SamlFirstNameAttribute) {
		return nil, false
	}
	return o.SamlFirstNameAttribute, true
}

// HasSamlFirstNameAttribute returns a boolean if a field has been set.
func (o *SSOConfig) HasSamlFirstNameAttribute() bool {
	if o != nil && !IsNil(o.SamlFirstNameAttribute) {
		return true
	}

	return false
}

// SetSamlFirstNameAttribute gets a reference to the given string and assigns it to the SamlFirstNameAttribute field.
func (o *SSOConfig) SetSamlFirstNameAttribute(v string) {
	o.SamlFirstNameAttribute = &v
}

// GetSamlLastNameAttribute returns the SamlLastNameAttribute field value if set, zero value otherwise.
func (o *SSOConfig) GetSamlLastNameAttribute() string {
	if o == nil || IsNil(o.SamlLastNameAttribute) {
		var ret string
		return ret
	}
	return *o.SamlLastNameAttribute
}

// GetSamlLastNameAttributeOk returns a tuple with the SamlLastNameAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOConfig) GetSamlLastNameAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.SamlLastNameAttribute) {
		return nil, false
	}
	return o.SamlLastNameAttribute, true
}

// HasSamlLastNameAttribute returns a boolean if a field has been set.
func (o *SSOConfig) HasSamlLastNameAttribute() bool {
	if o != nil && !IsNil(o.SamlLastNameAttribute) {
		return true
	}

	return false
}

// SetSamlLastNameAttribute gets a reference to the given string and assigns it to the SamlLastNameAttribute field.
func (o *SSOConfig) SetSamlLastNameAttribute(v string) {
	o.SamlLastNameAttribute = &v
}

// GetSamlGroupsAttribute returns the SamlGroupsAttribute field value if set, zero value otherwise.
func (o *SSOConfig) GetSamlGroupsAttribute() string {
	if o == nil || IsNil(o.SamlGroupsAttribute) {
		var ret string
		return ret
	}
	return *o.SamlGroupsAttribute
}

// GetSamlGroupsAttributeOk returns a tuple with the SamlGroupsAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOConfig) GetSamlGroupsAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.SamlGroupsAttribute) {
		return nil, false
	}
	return o.SamlGroupsAttribute, true
}

// HasSamlGroupsAttribute returns a boolean if a field has been set.
func (o *SSOConfig) HasSamlGroupsAttribute() bool {
	if o != nil && !IsNil(o.SamlGroupsAttribute) {
		return true
	}

	return false
}

// SetSamlGroupsAttribute gets a reference to the given string and assigns it to the SamlGroupsAttribute field.
func (o *SSOConfig) SetSamlGroupsAttribute(v string) {
	o.SamlGroupsAttribute = &v
}

// GetSamlSyncGroupClaims returns the SamlSyncGroupClaims field value if set, zero value otherwise.
func (o *SSOConfig) GetSamlSyncGroupClaims() bool {
	if o == nil || IsNil(o.SamlSyncGroupClaims) {
		var ret bool
		return ret
	}
	return *o.SamlSyncGroupClaims
}

// GetSamlSyncGroupClaimsOk returns a tuple with the SamlSyncGroupClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOConfig) GetSamlSyncGroupClaimsOk() (*bool, bool) {
	if o == nil || IsNil(o.SamlSyncGroupClaims) {
		return nil, false
	}
	return o.SamlSyncGroupClaims, true
}

// HasSamlSyncGroupClaims returns a boolean if a field has been set.
func (o *SSOConfig) HasSamlSyncGroupClaims() bool {
	if o != nil && !IsNil(o.SamlSyncGroupClaims) {
		return true
	}

	return false
}

// SetSamlSyncGroupClaims gets a reference to the given bool and assigns it to the SamlSyncGroupClaims field.
func (o *SSOConfig) SetSamlSyncGroupClaims(v bool) {
	o.SamlSyncGroupClaims = &v
}

// GetLdapSyncGroupClaims returns the LdapSyncGroupClaims field value if set, zero value otherwise.
func (o *SSOConfig) GetLdapSyncGroupClaims() bool {
	if o == nil || IsNil(o.LdapSyncGroupClaims) {
		var ret bool
		return ret
	}
	return *o.LdapSyncGroupClaims
}

// GetLdapSyncGroupClaimsOk returns a tuple with the LdapSyncGroupClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOConfig) GetLdapSyncGroupClaimsOk() (*bool, bool) {
	if o == nil || IsNil(o.LdapSyncGroupClaims) {
		return nil, false
	}
	return o.LdapSyncGroupClaims, true
}

// HasLdapSyncGroupClaims returns a boolean if a field has been set.
func (o *SSOConfig) HasLdapSyncGroupClaims() bool {
	if o != nil && !IsNil(o.LdapSyncGroupClaims) {
		return true
	}

	return false
}

// SetLdapSyncGroupClaims gets a reference to the given bool and assigns it to the LdapSyncGroupClaims field.
func (o *SSOConfig) SetLdapSyncGroupClaims(v bool) {
	o.LdapSyncGroupClaims = &v
}

// GetLdapRoleMapping returns the LdapRoleMapping field value if set, zero value otherwise.
func (o *SSOConfig) GetLdapRoleMapping() string {
	if o == nil || IsNil(o.LdapRoleMapping) {
		var ret string
		return ret
	}
	return *o.LdapRoleMapping
}

// GetLdapRoleMappingOk returns a tuple with the LdapRoleMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOConfig) GetLdapRoleMappingOk() (*string, bool) {
	if o == nil || IsNil(o.LdapRoleMapping) {
		return nil, false
	}
	return o.LdapRoleMapping, true
}

// HasLdapRoleMapping returns a boolean if a field has been set.
func (o *SSOConfig) HasLdapRoleMapping() bool {
	if o != nil && !IsNil(o.LdapRoleMapping) {
		return true
	}

	return false
}

// SetLdapRoleMapping gets a reference to the given string and assigns it to the LdapRoleMapping field.
func (o *SSOConfig) SetLdapRoleMapping(v string) {
	o.LdapRoleMapping = &v
}

// GetLdapServerUrl returns the LdapServerUrl field value if set, zero value otherwise.
func (o *SSOConfig) GetLdapServerUrl() string {
	if o == nil || IsNil(o.LdapServerUrl) {
		var ret string
		return ret
	}
	return *o.LdapServerUrl
}

// GetLdapServerUrlOk returns a tuple with the LdapServerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOConfig) GetLdapServerUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LdapServerUrl) {
		return nil, false
	}
	return o.LdapServerUrl, true
}

// HasLdapServerUrl returns a boolean if a field has been set.
func (o *SSOConfig) HasLdapServerUrl() bool {
	if o != nil && !IsNil(o.LdapServerUrl) {
		return true
	}

	return false
}

// SetLdapServerUrl gets a reference to the given string and assigns it to the LdapServerUrl field.
func (o *SSOConfig) SetLdapServerUrl(v string) {
	o.LdapServerUrl = &v
}

// GetLdapBaseDomainComponents returns the LdapBaseDomainComponents field value if set, zero value otherwise.
func (o *SSOConfig) GetLdapBaseDomainComponents() string {
	if o == nil || IsNil(o.LdapBaseDomainComponents) {
		var ret string
		return ret
	}
	return *o.LdapBaseDomainComponents
}

// GetLdapBaseDomainComponentsOk returns a tuple with the LdapBaseDomainComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOConfig) GetLdapBaseDomainComponentsOk() (*string, bool) {
	if o == nil || IsNil(o.LdapBaseDomainComponents) {
		return nil, false
	}
	return o.LdapBaseDomainComponents, true
}

// HasLdapBaseDomainComponents returns a boolean if a field has been set.
func (o *SSOConfig) HasLdapBaseDomainComponents() bool {
	if o != nil && !IsNil(o.LdapBaseDomainComponents) {
		return true
	}

	return false
}

// SetLdapBaseDomainComponents gets a reference to the given string and assigns it to the LdapBaseDomainComponents field.
func (o *SSOConfig) SetLdapBaseDomainComponents(v string) {
	o.LdapBaseDomainComponents = &v
}

// GetLdapServerName returns the LdapServerName field value if set, zero value otherwise.
func (o *SSOConfig) GetLdapServerName() string {
	if o == nil || IsNil(o.LdapServerName) {
		var ret string
		return ret
	}
	return *o.LdapServerName
}

// GetLdapServerNameOk returns a tuple with the LdapServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOConfig) GetLdapServerNameOk() (*string, bool) {
	if o == nil || IsNil(o.LdapServerName) {
		return nil, false
	}
	return o.LdapServerName, true
}

// HasLdapServerName returns a boolean if a field has been set.
func (o *SSOConfig) HasLdapServerName() bool {
	if o != nil && !IsNil(o.LdapServerName) {
		return true
	}

	return false
}

// SetLdapServerName gets a reference to the given string and assigns it to the LdapServerName field.
func (o *SSOConfig) SetLdapServerName(v string) {
	o.LdapServerName = &v
}

// GetLdapServerKey returns the LdapServerKey field value if set, zero value otherwise.
func (o *SSOConfig) GetLdapServerKey() string {
	if o == nil || IsNil(o.LdapServerKey) {
		var ret string
		return ret
	}
	return *o.LdapServerKey
}

// GetLdapServerKeyOk returns a tuple with the LdapServerKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOConfig) GetLdapServerKeyOk() (*string, bool) {
	if o == nil || IsNil(o.LdapServerKey) {
		return nil, false
	}
	return o.LdapServerKey, true
}

// HasLdapServerKey returns a boolean if a field has been set.
func (o *SSOConfig) HasLdapServerKey() bool {
	if o != nil && !IsNil(o.LdapServerKey) {
		return true
	}

	return false
}

// SetLdapServerKey gets a reference to the given string and assigns it to the LdapServerKey field.
func (o *SSOConfig) SetLdapServerKey(v string) {
	o.LdapServerKey = &v
}

// GetLdapServerCertificate returns the LdapServerCertificate field value if set, zero value otherwise.
func (o *SSOConfig) GetLdapServerCertificate() string {
	if o == nil || IsNil(o.LdapServerCertificate) {
		var ret string
		return ret
	}
	return *o.LdapServerCertificate
}

// GetLdapServerCertificateOk returns a tuple with the LdapServerCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOConfig) GetLdapServerCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.LdapServerCertificate) {
		return nil, false
	}
	return o.LdapServerCertificate, true
}

// HasLdapServerCertificate returns a boolean if a field has been set.
func (o *SSOConfig) HasLdapServerCertificate() bool {
	if o != nil && !IsNil(o.LdapServerCertificate) {
		return true
	}

	return false
}

// SetLdapServerCertificate gets a reference to the given string and assigns it to the LdapServerCertificate field.
func (o *SSOConfig) SetLdapServerCertificate(v string) {
	o.LdapServerCertificate = &v
}

func (o SSOConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SSOConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["config_type"] = o.ConfigType
	if !IsNil(o.GoogleClientId) {
		toSerialize["google_client_id"] = o.GoogleClientId
	}
	if !IsNil(o.GoogleClientSecret) {
		toSerialize["google_client_secret"] = o.GoogleClientSecret
	}
	toSerialize["disable_email_password_login"] = o.DisableEmailPasswordLogin
	if !IsNil(o.OidcClientId) {
		toSerialize["oidc_client_id"] = o.OidcClientId
	}
	if !IsNil(o.OidcClientSecret) {
		toSerialize["oidc_client_secret"] = o.OidcClientSecret
	}
	if !IsNil(o.OidcScopes) {
		toSerialize["oidc_scopes"] = o.OidcScopes
	}
	if !IsNil(o.OidcAuthUrl) {
		toSerialize["oidc_auth_url"] = o.OidcAuthUrl
	}
	if !IsNil(o.OidcTokenUrl) {
		toSerialize["oidc_token_url"] = o.OidcTokenUrl
	}
	if !IsNil(o.OidcUserinfoUrl) {
		toSerialize["oidc_userinfo_url"] = o.OidcUserinfoUrl
	}
	if !IsNil(o.OidcAudience) {
		toSerialize["oidc_audience"] = o.OidcAudience
	}
	if !IsNil(o.JwtEmailKey) {
		toSerialize["jwt_email_key"] = o.JwtEmailKey
	}
	if !IsNil(o.JwtRolesKey) {
		toSerialize["jwt_roles_key"] = o.JwtRolesKey
	}
	if !IsNil(o.JwtFirstNameKey) {
		toSerialize["jwt_first_name_key"] = o.JwtFirstNameKey
	}
	if !IsNil(o.JwtLastNameKey) {
		toSerialize["jwt_last_name_key"] = o.JwtLastNameKey
	}
	if !IsNil(o.RolesMapping) {
		toSerialize["roles_mapping"] = o.RolesMapping
	}
	if !IsNil(o.JitEnabled) {
		toSerialize["jit_enabled"] = o.JitEnabled
	}
	if !IsNil(o.RestrictedDomain) {
		toSerialize["restricted_domain"] = o.RestrictedDomain
	}
	if !IsNil(o.TriggerLoginAutomatically) {
		toSerialize["trigger_login_automatically"] = o.TriggerLoginAutomatically
	}
	if !IsNil(o.IdpMetadataXml) {
		toSerialize["idp_metadata_xml"] = o.IdpMetadataXml
	}
	if !IsNil(o.SamlFirstNameAttribute) {
		toSerialize["saml_first_name_attribute"] = o.SamlFirstNameAttribute
	}
	if !IsNil(o.SamlLastNameAttribute) {
		toSerialize["saml_last_name_attribute"] = o.SamlLastNameAttribute
	}
	if !IsNil(o.SamlGroupsAttribute) {
		toSerialize["saml_groups_attribute"] = o.SamlGroupsAttribute
	}
	if !IsNil(o.SamlSyncGroupClaims) {
		toSerialize["saml_sync_group_claims"] = o.SamlSyncGroupClaims
	}
	if !IsNil(o.LdapSyncGroupClaims) {
		toSerialize["ldap_sync_group_claims"] = o.LdapSyncGroupClaims
	}
	if !IsNil(o.LdapRoleMapping) {
		toSerialize["ldap_role_mapping"] = o.LdapRoleMapping
	}
	if !IsNil(o.LdapServerUrl) {
		toSerialize["ldap_server_url"] = o.LdapServerUrl
	}
	if !IsNil(o.LdapBaseDomainComponents) {
		toSerialize["ldap_base_domain_components"] = o.LdapBaseDomainComponents
	}
	if !IsNil(o.LdapServerName) {
		toSerialize["ldap_server_name"] = o.LdapServerName
	}
	if !IsNil(o.LdapServerKey) {
		toSerialize["ldap_server_key"] = o.LdapServerKey
	}
	if !IsNil(o.LdapServerCertificate) {
		toSerialize["ldap_server_certificate"] = o.LdapServerCertificate
	}
	return toSerialize, nil
}

func (o *SSOConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"config_type",
		"disable_email_password_login",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSSOConfig := _SSOConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSSOConfig)

	if err != nil {
		return err
	}

	*o = SSOConfig(varSSOConfig)

	return err
}

type NullableSSOConfig struct {
	value *SSOConfig
	isSet bool
}

func (v NullableSSOConfig) Get() *SSOConfig {
	return v.value
}

func (v *NullableSSOConfig) Set(val *SSOConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSSOConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSSOConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSSOConfig(val *SSOConfig) *NullableSSOConfig {
	return &NullableSSOConfig{value: val, isSet: true}
}

func (v NullableSSOConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSSOConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


