/*
Retool API

Go to Settings > API to get started. Once you generate an API token, use bearer token authentication to make requests.

API version: 2.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the SpacesPostRequestOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpacesPostRequestOptions{}

// SpacesPostRequestOptions struct for SpacesPostRequestOptions
type SpacesPostRequestOptions struct {
	// Copy SSO settings from the admin space.
	CopySsoSettings *bool `json:"copy_sso_settings,omitempty"`
	// Copy Branding and Themes settings from the admin space.
	CopyBrandingAndThemesSettings *bool `json:"copy_branding_and_themes_settings,omitempty"`
	// List of emails of users from the admin space that need to be added to the new space as admins.
	UsersToCopyAsAdmins []string `json:"users_to_copy_as_admins,omitempty"`
	// Create an admin user in the new space for the creator instead of just sending out an invite.
	CreateAdminUser *bool `json:"create_admin_user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SpacesPostRequestOptions SpacesPostRequestOptions

// NewSpacesPostRequestOptions instantiates a new SpacesPostRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpacesPostRequestOptions() *SpacesPostRequestOptions {
	this := SpacesPostRequestOptions{}
	return &this
}

// NewSpacesPostRequestOptionsWithDefaults instantiates a new SpacesPostRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpacesPostRequestOptionsWithDefaults() *SpacesPostRequestOptions {
	this := SpacesPostRequestOptions{}
	return &this
}

// GetCopySsoSettings returns the CopySsoSettings field value if set, zero value otherwise.
func (o *SpacesPostRequestOptions) GetCopySsoSettings() bool {
	if o == nil || IsNil(o.CopySsoSettings) {
		var ret bool
		return ret
	}
	return *o.CopySsoSettings
}

// GetCopySsoSettingsOk returns a tuple with the CopySsoSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpacesPostRequestOptions) GetCopySsoSettingsOk() (*bool, bool) {
	if o == nil || IsNil(o.CopySsoSettings) {
		return nil, false
	}
	return o.CopySsoSettings, true
}

// HasCopySsoSettings returns a boolean if a field has been set.
func (o *SpacesPostRequestOptions) HasCopySsoSettings() bool {
	if o != nil && !IsNil(o.CopySsoSettings) {
		return true
	}

	return false
}

// SetCopySsoSettings gets a reference to the given bool and assigns it to the CopySsoSettings field.
func (o *SpacesPostRequestOptions) SetCopySsoSettings(v bool) {
	o.CopySsoSettings = &v
}

// GetCopyBrandingAndThemesSettings returns the CopyBrandingAndThemesSettings field value if set, zero value otherwise.
func (o *SpacesPostRequestOptions) GetCopyBrandingAndThemesSettings() bool {
	if o == nil || IsNil(o.CopyBrandingAndThemesSettings) {
		var ret bool
		return ret
	}
	return *o.CopyBrandingAndThemesSettings
}

// GetCopyBrandingAndThemesSettingsOk returns a tuple with the CopyBrandingAndThemesSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpacesPostRequestOptions) GetCopyBrandingAndThemesSettingsOk() (*bool, bool) {
	if o == nil || IsNil(o.CopyBrandingAndThemesSettings) {
		return nil, false
	}
	return o.CopyBrandingAndThemesSettings, true
}

// HasCopyBrandingAndThemesSettings returns a boolean if a field has been set.
func (o *SpacesPostRequestOptions) HasCopyBrandingAndThemesSettings() bool {
	if o != nil && !IsNil(o.CopyBrandingAndThemesSettings) {
		return true
	}

	return false
}

// SetCopyBrandingAndThemesSettings gets a reference to the given bool and assigns it to the CopyBrandingAndThemesSettings field.
func (o *SpacesPostRequestOptions) SetCopyBrandingAndThemesSettings(v bool) {
	o.CopyBrandingAndThemesSettings = &v
}

// GetUsersToCopyAsAdmins returns the UsersToCopyAsAdmins field value if set, zero value otherwise.
func (o *SpacesPostRequestOptions) GetUsersToCopyAsAdmins() []string {
	if o == nil || IsNil(o.UsersToCopyAsAdmins) {
		var ret []string
		return ret
	}
	return o.UsersToCopyAsAdmins
}

// GetUsersToCopyAsAdminsOk returns a tuple with the UsersToCopyAsAdmins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpacesPostRequestOptions) GetUsersToCopyAsAdminsOk() ([]string, bool) {
	if o == nil || IsNil(o.UsersToCopyAsAdmins) {
		return nil, false
	}
	return o.UsersToCopyAsAdmins, true
}

// HasUsersToCopyAsAdmins returns a boolean if a field has been set.
func (o *SpacesPostRequestOptions) HasUsersToCopyAsAdmins() bool {
	if o != nil && !IsNil(o.UsersToCopyAsAdmins) {
		return true
	}

	return false
}

// SetUsersToCopyAsAdmins gets a reference to the given []string and assigns it to the UsersToCopyAsAdmins field.
func (o *SpacesPostRequestOptions) SetUsersToCopyAsAdmins(v []string) {
	o.UsersToCopyAsAdmins = v
}

// GetCreateAdminUser returns the CreateAdminUser field value if set, zero value otherwise.
func (o *SpacesPostRequestOptions) GetCreateAdminUser() bool {
	if o == nil || IsNil(o.CreateAdminUser) {
		var ret bool
		return ret
	}
	return *o.CreateAdminUser
}

// GetCreateAdminUserOk returns a tuple with the CreateAdminUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpacesPostRequestOptions) GetCreateAdminUserOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateAdminUser) {
		return nil, false
	}
	return o.CreateAdminUser, true
}

// HasCreateAdminUser returns a boolean if a field has been set.
func (o *SpacesPostRequestOptions) HasCreateAdminUser() bool {
	if o != nil && !IsNil(o.CreateAdminUser) {
		return true
	}

	return false
}

// SetCreateAdminUser gets a reference to the given bool and assigns it to the CreateAdminUser field.
func (o *SpacesPostRequestOptions) SetCreateAdminUser(v bool) {
	o.CreateAdminUser = &v
}

func (o SpacesPostRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpacesPostRequestOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CopySsoSettings) {
		toSerialize["copy_sso_settings"] = o.CopySsoSettings
	}
	if !IsNil(o.CopyBrandingAndThemesSettings) {
		toSerialize["copy_branding_and_themes_settings"] = o.CopyBrandingAndThemesSettings
	}
	if !IsNil(o.UsersToCopyAsAdmins) {
		toSerialize["users_to_copy_as_admins"] = o.UsersToCopyAsAdmins
	}
	if !IsNil(o.CreateAdminUser) {
		toSerialize["create_admin_user"] = o.CreateAdminUser
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SpacesPostRequestOptions) UnmarshalJSON(data []byte) (err error) {
	varSpacesPostRequestOptions := _SpacesPostRequestOptions{}

	err = json.Unmarshal(data, &varSpacesPostRequestOptions)

	if err != nil {
		return err
	}

	*o = SpacesPostRequestOptions(varSpacesPostRequestOptions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "copy_sso_settings")
		delete(additionalProperties, "copy_branding_and_themes_settings")
		delete(additionalProperties, "users_to_copy_as_admins")
		delete(additionalProperties, "create_admin_user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSpacesPostRequestOptions struct {
	value *SpacesPostRequestOptions
	isSet bool
}

func (v NullableSpacesPostRequestOptions) Get() *SpacesPostRequestOptions {
	return v.value
}

func (v *NullableSpacesPostRequestOptions) Set(val *SpacesPostRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSpacesPostRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSpacesPostRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpacesPostRequestOptions(val *SpacesPostRequestOptions) *NullableSpacesPostRequestOptions {
	return &NullableSpacesPostRequestOptions{value: val, isSet: true}
}

func (v NullableSpacesPostRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpacesPostRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


