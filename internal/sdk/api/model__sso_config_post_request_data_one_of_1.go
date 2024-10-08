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

// checks if the SsoConfigPostRequestDataOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SsoConfigPostRequestDataOneOf1{}

// SsoConfigPostRequestDataOneOf1 struct for SsoConfigPostRequestDataOneOf1
type SsoConfigPostRequestDataOneOf1 struct {
	ConfigType string `json:"config_type"`
	OidcClientId string `json:"oidc_client_id"`
	OidcClientSecret string `json:"oidc_client_secret"`
	OidcScopes string `json:"oidc_scopes"`
	OidcAuthUrl string `json:"oidc_auth_url"`
	OidcTokenUrl string `json:"oidc_token_url"`
	OidcUserinfoUrl *string `json:"oidc_userinfo_url,omitempty"`
	OidcAudience *string `json:"oidc_audience,omitempty"`
	JwtEmailKey string `json:"jwt_email_key"`
	JwtRolesKey *string `json:"jwt_roles_key,omitempty"`
	JwtFirstNameKey string `json:"jwt_first_name_key"`
	JwtLastNameKey string `json:"jwt_last_name_key"`
	RolesMapping *string `json:"roles_mapping,omitempty"`
	JitEnabled bool `json:"jit_enabled"`
	RestrictedDomain *string `json:"restricted_domain,omitempty"`
	TriggerLoginAutomatically bool `json:"trigger_login_automatically"`
	DisableEmailPasswordLogin bool `json:"disable_email_password_login"`
}

type _SsoConfigPostRequestDataOneOf1 SsoConfigPostRequestDataOneOf1

// NewSsoConfigPostRequestDataOneOf1 instantiates a new SsoConfigPostRequestDataOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsoConfigPostRequestDataOneOf1(configType string, oidcClientId string, oidcClientSecret string, oidcScopes string, oidcAuthUrl string, oidcTokenUrl string, jwtEmailKey string, jwtFirstNameKey string, jwtLastNameKey string, jitEnabled bool, triggerLoginAutomatically bool, disableEmailPasswordLogin bool) *SsoConfigPostRequestDataOneOf1 {
	this := SsoConfigPostRequestDataOneOf1{}
	this.ConfigType = configType
	this.OidcClientId = oidcClientId
	this.OidcClientSecret = oidcClientSecret
	this.OidcScopes = oidcScopes
	this.OidcAuthUrl = oidcAuthUrl
	this.OidcTokenUrl = oidcTokenUrl
	this.JwtEmailKey = jwtEmailKey
	this.JwtFirstNameKey = jwtFirstNameKey
	this.JwtLastNameKey = jwtLastNameKey
	this.JitEnabled = jitEnabled
	this.TriggerLoginAutomatically = triggerLoginAutomatically
	this.DisableEmailPasswordLogin = disableEmailPasswordLogin
	return &this
}

// NewSsoConfigPostRequestDataOneOf1WithDefaults instantiates a new SsoConfigPostRequestDataOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoConfigPostRequestDataOneOf1WithDefaults() *SsoConfigPostRequestDataOneOf1 {
	this := SsoConfigPostRequestDataOneOf1{}
	return &this
}

// GetConfigType returns the ConfigType field value
func (o *SsoConfigPostRequestDataOneOf1) GetConfigType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigType
}

// GetConfigTypeOk returns a tuple with the ConfigType field value
// and a boolean to check if the value has been set.
func (o *SsoConfigPostRequestDataOneOf1) GetConfigTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigType, true
}

// SetConfigType sets field value
func (o *SsoConfigPostRequestDataOneOf1) SetConfigType(v string) {
	o.ConfigType = v
}

// GetOidcClientId returns the OidcClientId field value
func (o *SsoConfigPostRequestDataOneOf1) GetOidcClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OidcClientId
}

// GetOidcClientIdOk returns a tuple with the OidcClientId field value
// and a boolean to check if the value has been set.
func (o *SsoConfigPostRequestDataOneOf1) GetOidcClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OidcClientId, true
}

// SetOidcClientId sets field value
func (o *SsoConfigPostRequestDataOneOf1) SetOidcClientId(v string) {
	o.OidcClientId = v
}

// GetOidcClientSecret returns the OidcClientSecret field value
func (o *SsoConfigPostRequestDataOneOf1) GetOidcClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OidcClientSecret
}

// GetOidcClientSecretOk returns a tuple with the OidcClientSecret field value
// and a boolean to check if the value has been set.
func (o *SsoConfigPostRequestDataOneOf1) GetOidcClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OidcClientSecret, true
}

// SetOidcClientSecret sets field value
func (o *SsoConfigPostRequestDataOneOf1) SetOidcClientSecret(v string) {
	o.OidcClientSecret = v
}

// GetOidcScopes returns the OidcScopes field value
func (o *SsoConfigPostRequestDataOneOf1) GetOidcScopes() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OidcScopes
}

// GetOidcScopesOk returns a tuple with the OidcScopes field value
// and a boolean to check if the value has been set.
func (o *SsoConfigPostRequestDataOneOf1) GetOidcScopesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OidcScopes, true
}

// SetOidcScopes sets field value
func (o *SsoConfigPostRequestDataOneOf1) SetOidcScopes(v string) {
	o.OidcScopes = v
}

// GetOidcAuthUrl returns the OidcAuthUrl field value
func (o *SsoConfigPostRequestDataOneOf1) GetOidcAuthUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OidcAuthUrl
}

// GetOidcAuthUrlOk returns a tuple with the OidcAuthUrl field value
// and a boolean to check if the value has been set.
func (o *SsoConfigPostRequestDataOneOf1) GetOidcAuthUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OidcAuthUrl, true
}

// SetOidcAuthUrl sets field value
func (o *SsoConfigPostRequestDataOneOf1) SetOidcAuthUrl(v string) {
	o.OidcAuthUrl = v
}

// GetOidcTokenUrl returns the OidcTokenUrl field value
func (o *SsoConfigPostRequestDataOneOf1) GetOidcTokenUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OidcTokenUrl
}

// GetOidcTokenUrlOk returns a tuple with the OidcTokenUrl field value
// and a boolean to check if the value has been set.
func (o *SsoConfigPostRequestDataOneOf1) GetOidcTokenUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OidcTokenUrl, true
}

// SetOidcTokenUrl sets field value
func (o *SsoConfigPostRequestDataOneOf1) SetOidcTokenUrl(v string) {
	o.OidcTokenUrl = v
}

// GetOidcUserinfoUrl returns the OidcUserinfoUrl field value if set, zero value otherwise.
func (o *SsoConfigPostRequestDataOneOf1) GetOidcUserinfoUrl() string {
	if o == nil || IsNil(o.OidcUserinfoUrl) {
		var ret string
		return ret
	}
	return *o.OidcUserinfoUrl
}

// GetOidcUserinfoUrlOk returns a tuple with the OidcUserinfoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoConfigPostRequestDataOneOf1) GetOidcUserinfoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.OidcUserinfoUrl) {
		return nil, false
	}
	return o.OidcUserinfoUrl, true
}

// HasOidcUserinfoUrl returns a boolean if a field has been set.
func (o *SsoConfigPostRequestDataOneOf1) HasOidcUserinfoUrl() bool {
	if o != nil && !IsNil(o.OidcUserinfoUrl) {
		return true
	}

	return false
}

// SetOidcUserinfoUrl gets a reference to the given string and assigns it to the OidcUserinfoUrl field.
func (o *SsoConfigPostRequestDataOneOf1) SetOidcUserinfoUrl(v string) {
	o.OidcUserinfoUrl = &v
}

// GetOidcAudience returns the OidcAudience field value if set, zero value otherwise.
func (o *SsoConfigPostRequestDataOneOf1) GetOidcAudience() string {
	if o == nil || IsNil(o.OidcAudience) {
		var ret string
		return ret
	}
	return *o.OidcAudience
}

// GetOidcAudienceOk returns a tuple with the OidcAudience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoConfigPostRequestDataOneOf1) GetOidcAudienceOk() (*string, bool) {
	if o == nil || IsNil(o.OidcAudience) {
		return nil, false
	}
	return o.OidcAudience, true
}

// HasOidcAudience returns a boolean if a field has been set.
func (o *SsoConfigPostRequestDataOneOf1) HasOidcAudience() bool {
	if o != nil && !IsNil(o.OidcAudience) {
		return true
	}

	return false
}

// SetOidcAudience gets a reference to the given string and assigns it to the OidcAudience field.
func (o *SsoConfigPostRequestDataOneOf1) SetOidcAudience(v string) {
	o.OidcAudience = &v
}

// GetJwtEmailKey returns the JwtEmailKey field value
func (o *SsoConfigPostRequestDataOneOf1) GetJwtEmailKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JwtEmailKey
}

// GetJwtEmailKeyOk returns a tuple with the JwtEmailKey field value
// and a boolean to check if the value has been set.
func (o *SsoConfigPostRequestDataOneOf1) GetJwtEmailKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JwtEmailKey, true
}

// SetJwtEmailKey sets field value
func (o *SsoConfigPostRequestDataOneOf1) SetJwtEmailKey(v string) {
	o.JwtEmailKey = v
}

// GetJwtRolesKey returns the JwtRolesKey field value if set, zero value otherwise.
func (o *SsoConfigPostRequestDataOneOf1) GetJwtRolesKey() string {
	if o == nil || IsNil(o.JwtRolesKey) {
		var ret string
		return ret
	}
	return *o.JwtRolesKey
}

// GetJwtRolesKeyOk returns a tuple with the JwtRolesKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoConfigPostRequestDataOneOf1) GetJwtRolesKeyOk() (*string, bool) {
	if o == nil || IsNil(o.JwtRolesKey) {
		return nil, false
	}
	return o.JwtRolesKey, true
}

// HasJwtRolesKey returns a boolean if a field has been set.
func (o *SsoConfigPostRequestDataOneOf1) HasJwtRolesKey() bool {
	if o != nil && !IsNil(o.JwtRolesKey) {
		return true
	}

	return false
}

// SetJwtRolesKey gets a reference to the given string and assigns it to the JwtRolesKey field.
func (o *SsoConfigPostRequestDataOneOf1) SetJwtRolesKey(v string) {
	o.JwtRolesKey = &v
}

// GetJwtFirstNameKey returns the JwtFirstNameKey field value
func (o *SsoConfigPostRequestDataOneOf1) GetJwtFirstNameKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JwtFirstNameKey
}

// GetJwtFirstNameKeyOk returns a tuple with the JwtFirstNameKey field value
// and a boolean to check if the value has been set.
func (o *SsoConfigPostRequestDataOneOf1) GetJwtFirstNameKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JwtFirstNameKey, true
}

// SetJwtFirstNameKey sets field value
func (o *SsoConfigPostRequestDataOneOf1) SetJwtFirstNameKey(v string) {
	o.JwtFirstNameKey = v
}

// GetJwtLastNameKey returns the JwtLastNameKey field value
func (o *SsoConfigPostRequestDataOneOf1) GetJwtLastNameKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JwtLastNameKey
}

// GetJwtLastNameKeyOk returns a tuple with the JwtLastNameKey field value
// and a boolean to check if the value has been set.
func (o *SsoConfigPostRequestDataOneOf1) GetJwtLastNameKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JwtLastNameKey, true
}

// SetJwtLastNameKey sets field value
func (o *SsoConfigPostRequestDataOneOf1) SetJwtLastNameKey(v string) {
	o.JwtLastNameKey = v
}

// GetRolesMapping returns the RolesMapping field value if set, zero value otherwise.
func (o *SsoConfigPostRequestDataOneOf1) GetRolesMapping() string {
	if o == nil || IsNil(o.RolesMapping) {
		var ret string
		return ret
	}
	return *o.RolesMapping
}

// GetRolesMappingOk returns a tuple with the RolesMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoConfigPostRequestDataOneOf1) GetRolesMappingOk() (*string, bool) {
	if o == nil || IsNil(o.RolesMapping) {
		return nil, false
	}
	return o.RolesMapping, true
}

// HasRolesMapping returns a boolean if a field has been set.
func (o *SsoConfigPostRequestDataOneOf1) HasRolesMapping() bool {
	if o != nil && !IsNil(o.RolesMapping) {
		return true
	}

	return false
}

// SetRolesMapping gets a reference to the given string and assigns it to the RolesMapping field.
func (o *SsoConfigPostRequestDataOneOf1) SetRolesMapping(v string) {
	o.RolesMapping = &v
}

// GetJitEnabled returns the JitEnabled field value
func (o *SsoConfigPostRequestDataOneOf1) GetJitEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.JitEnabled
}

// GetJitEnabledOk returns a tuple with the JitEnabled field value
// and a boolean to check if the value has been set.
func (o *SsoConfigPostRequestDataOneOf1) GetJitEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JitEnabled, true
}

// SetJitEnabled sets field value
func (o *SsoConfigPostRequestDataOneOf1) SetJitEnabled(v bool) {
	o.JitEnabled = v
}

// GetRestrictedDomain returns the RestrictedDomain field value if set, zero value otherwise.
func (o *SsoConfigPostRequestDataOneOf1) GetRestrictedDomain() string {
	if o == nil || IsNil(o.RestrictedDomain) {
		var ret string
		return ret
	}
	return *o.RestrictedDomain
}

// GetRestrictedDomainOk returns a tuple with the RestrictedDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoConfigPostRequestDataOneOf1) GetRestrictedDomainOk() (*string, bool) {
	if o == nil || IsNil(o.RestrictedDomain) {
		return nil, false
	}
	return o.RestrictedDomain, true
}

// HasRestrictedDomain returns a boolean if a field has been set.
func (o *SsoConfigPostRequestDataOneOf1) HasRestrictedDomain() bool {
	if o != nil && !IsNil(o.RestrictedDomain) {
		return true
	}

	return false
}

// SetRestrictedDomain gets a reference to the given string and assigns it to the RestrictedDomain field.
func (o *SsoConfigPostRequestDataOneOf1) SetRestrictedDomain(v string) {
	o.RestrictedDomain = &v
}

// GetTriggerLoginAutomatically returns the TriggerLoginAutomatically field value
func (o *SsoConfigPostRequestDataOneOf1) GetTriggerLoginAutomatically() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.TriggerLoginAutomatically
}

// GetTriggerLoginAutomaticallyOk returns a tuple with the TriggerLoginAutomatically field value
// and a boolean to check if the value has been set.
func (o *SsoConfigPostRequestDataOneOf1) GetTriggerLoginAutomaticallyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerLoginAutomatically, true
}

// SetTriggerLoginAutomatically sets field value
func (o *SsoConfigPostRequestDataOneOf1) SetTriggerLoginAutomatically(v bool) {
	o.TriggerLoginAutomatically = v
}

// GetDisableEmailPasswordLogin returns the DisableEmailPasswordLogin field value
func (o *SsoConfigPostRequestDataOneOf1) GetDisableEmailPasswordLogin() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableEmailPasswordLogin
}

// GetDisableEmailPasswordLoginOk returns a tuple with the DisableEmailPasswordLogin field value
// and a boolean to check if the value has been set.
func (o *SsoConfigPostRequestDataOneOf1) GetDisableEmailPasswordLoginOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableEmailPasswordLogin, true
}

// SetDisableEmailPasswordLogin sets field value
func (o *SsoConfigPostRequestDataOneOf1) SetDisableEmailPasswordLogin(v bool) {
	o.DisableEmailPasswordLogin = v
}

func (o SsoConfigPostRequestDataOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SsoConfigPostRequestDataOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["config_type"] = o.ConfigType
	toSerialize["oidc_client_id"] = o.OidcClientId
	toSerialize["oidc_client_secret"] = o.OidcClientSecret
	toSerialize["oidc_scopes"] = o.OidcScopes
	toSerialize["oidc_auth_url"] = o.OidcAuthUrl
	toSerialize["oidc_token_url"] = o.OidcTokenUrl
	if !IsNil(o.OidcUserinfoUrl) {
		toSerialize["oidc_userinfo_url"] = o.OidcUserinfoUrl
	}
	if !IsNil(o.OidcAudience) {
		toSerialize["oidc_audience"] = o.OidcAudience
	}
	toSerialize["jwt_email_key"] = o.JwtEmailKey
	if !IsNil(o.JwtRolesKey) {
		toSerialize["jwt_roles_key"] = o.JwtRolesKey
	}
	toSerialize["jwt_first_name_key"] = o.JwtFirstNameKey
	toSerialize["jwt_last_name_key"] = o.JwtLastNameKey
	if !IsNil(o.RolesMapping) {
		toSerialize["roles_mapping"] = o.RolesMapping
	}
	toSerialize["jit_enabled"] = o.JitEnabled
	if !IsNil(o.RestrictedDomain) {
		toSerialize["restricted_domain"] = o.RestrictedDomain
	}
	toSerialize["trigger_login_automatically"] = o.TriggerLoginAutomatically
	toSerialize["disable_email_password_login"] = o.DisableEmailPasswordLogin
	return toSerialize, nil
}

func (o *SsoConfigPostRequestDataOneOf1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"config_type",
		"oidc_client_id",
		"oidc_client_secret",
		"oidc_scopes",
		"oidc_auth_url",
		"oidc_token_url",
		"jwt_email_key",
		"jwt_first_name_key",
		"jwt_last_name_key",
		"jit_enabled",
		"trigger_login_automatically",
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

	varSsoConfigPostRequestDataOneOf1 := _SsoConfigPostRequestDataOneOf1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSsoConfigPostRequestDataOneOf1)

	if err != nil {
		return err
	}

	*o = SsoConfigPostRequestDataOneOf1(varSsoConfigPostRequestDataOneOf1)

	return err
}

type NullableSsoConfigPostRequestDataOneOf1 struct {
	value *SsoConfigPostRequestDataOneOf1
	isSet bool
}

func (v NullableSsoConfigPostRequestDataOneOf1) Get() *SsoConfigPostRequestDataOneOf1 {
	return v.value
}

func (v *NullableSsoConfigPostRequestDataOneOf1) Set(val *SsoConfigPostRequestDataOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoConfigPostRequestDataOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoConfigPostRequestDataOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoConfigPostRequestDataOneOf1(val *SsoConfigPostRequestDataOneOf1) *NullableSsoConfigPostRequestDataOneOf1 {
	return &NullableSsoConfigPostRequestDataOneOf1{value: val, isSet: true}
}

func (v NullableSsoConfigPostRequestDataOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoConfigPostRequestDataOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


