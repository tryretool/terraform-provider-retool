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

// checks if the SourceControlSettingsPutRequestSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceControlSettingsPutRequestSettings{}

// SourceControlSettingsPutRequestSettings struct for SourceControlSettingsPutRequestSettings
type SourceControlSettingsPutRequestSettings struct {
	// When enabled, Retool automatically suggests a branch name on branch creation. Defaults to true.
	AutoBranchNamingEnabled *bool `json:"auto_branch_naming_enabled,omitempty"`
	// When enabled, Retool will use the template specified to create pull requests. Defaults to false.
	CustomPullRequestTemplateEnabled *bool `json:"custom_pull_request_template_enabled,omitempty"`
	// Pull requests created from Retool will use the template specified.
	CustomPullRequestTemplate *string `json:"custom_pull_request_template,omitempty"`
	// When set to true, creates a read-only instance of Retool, where app editing is disabled. Defaults to false.
	VersionControlLocked *bool `json:"version_control_locked,omitempty"`
	// When set to true, creates a uuid mapping for protected elements to be used in the source control repo. Defaults to false.
	ForceUuidMapping *bool `json:"force_uuid_mapping,omitempty"`
}

// NewSourceControlSettingsPutRequestSettings instantiates a new SourceControlSettingsPutRequestSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceControlSettingsPutRequestSettings() *SourceControlSettingsPutRequestSettings {
	this := SourceControlSettingsPutRequestSettings{}
	return &this
}

// NewSourceControlSettingsPutRequestSettingsWithDefaults instantiates a new SourceControlSettingsPutRequestSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceControlSettingsPutRequestSettingsWithDefaults() *SourceControlSettingsPutRequestSettings {
	this := SourceControlSettingsPutRequestSettings{}
	return &this
}

// GetAutoBranchNamingEnabled returns the AutoBranchNamingEnabled field value if set, zero value otherwise.
func (o *SourceControlSettingsPutRequestSettings) GetAutoBranchNamingEnabled() bool {
	if o == nil || IsNil(o.AutoBranchNamingEnabled) {
		var ret bool
		return ret
	}
	return *o.AutoBranchNamingEnabled
}

// GetAutoBranchNamingEnabledOk returns a tuple with the AutoBranchNamingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceControlSettingsPutRequestSettings) GetAutoBranchNamingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoBranchNamingEnabled) {
		return nil, false
	}
	return o.AutoBranchNamingEnabled, true
}

// HasAutoBranchNamingEnabled returns a boolean if a field has been set.
func (o *SourceControlSettingsPutRequestSettings) HasAutoBranchNamingEnabled() bool {
	if o != nil && !IsNil(o.AutoBranchNamingEnabled) {
		return true
	}

	return false
}

// SetAutoBranchNamingEnabled gets a reference to the given bool and assigns it to the AutoBranchNamingEnabled field.
func (o *SourceControlSettingsPutRequestSettings) SetAutoBranchNamingEnabled(v bool) {
	o.AutoBranchNamingEnabled = &v
}

// GetCustomPullRequestTemplateEnabled returns the CustomPullRequestTemplateEnabled field value if set, zero value otherwise.
func (o *SourceControlSettingsPutRequestSettings) GetCustomPullRequestTemplateEnabled() bool {
	if o == nil || IsNil(o.CustomPullRequestTemplateEnabled) {
		var ret bool
		return ret
	}
	return *o.CustomPullRequestTemplateEnabled
}

// GetCustomPullRequestTemplateEnabledOk returns a tuple with the CustomPullRequestTemplateEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceControlSettingsPutRequestSettings) GetCustomPullRequestTemplateEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CustomPullRequestTemplateEnabled) {
		return nil, false
	}
	return o.CustomPullRequestTemplateEnabled, true
}

// HasCustomPullRequestTemplateEnabled returns a boolean if a field has been set.
func (o *SourceControlSettingsPutRequestSettings) HasCustomPullRequestTemplateEnabled() bool {
	if o != nil && !IsNil(o.CustomPullRequestTemplateEnabled) {
		return true
	}

	return false
}

// SetCustomPullRequestTemplateEnabled gets a reference to the given bool and assigns it to the CustomPullRequestTemplateEnabled field.
func (o *SourceControlSettingsPutRequestSettings) SetCustomPullRequestTemplateEnabled(v bool) {
	o.CustomPullRequestTemplateEnabled = &v
}

// GetCustomPullRequestTemplate returns the CustomPullRequestTemplate field value if set, zero value otherwise.
func (o *SourceControlSettingsPutRequestSettings) GetCustomPullRequestTemplate() string {
	if o == nil || IsNil(o.CustomPullRequestTemplate) {
		var ret string
		return ret
	}
	return *o.CustomPullRequestTemplate
}

// GetCustomPullRequestTemplateOk returns a tuple with the CustomPullRequestTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceControlSettingsPutRequestSettings) GetCustomPullRequestTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.CustomPullRequestTemplate) {
		return nil, false
	}
	return o.CustomPullRequestTemplate, true
}

// HasCustomPullRequestTemplate returns a boolean if a field has been set.
func (o *SourceControlSettingsPutRequestSettings) HasCustomPullRequestTemplate() bool {
	if o != nil && !IsNil(o.CustomPullRequestTemplate) {
		return true
	}

	return false
}

// SetCustomPullRequestTemplate gets a reference to the given string and assigns it to the CustomPullRequestTemplate field.
func (o *SourceControlSettingsPutRequestSettings) SetCustomPullRequestTemplate(v string) {
	o.CustomPullRequestTemplate = &v
}

// GetVersionControlLocked returns the VersionControlLocked field value if set, zero value otherwise.
func (o *SourceControlSettingsPutRequestSettings) GetVersionControlLocked() bool {
	if o == nil || IsNil(o.VersionControlLocked) {
		var ret bool
		return ret
	}
	return *o.VersionControlLocked
}

// GetVersionControlLockedOk returns a tuple with the VersionControlLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceControlSettingsPutRequestSettings) GetVersionControlLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.VersionControlLocked) {
		return nil, false
	}
	return o.VersionControlLocked, true
}

// HasVersionControlLocked returns a boolean if a field has been set.
func (o *SourceControlSettingsPutRequestSettings) HasVersionControlLocked() bool {
	if o != nil && !IsNil(o.VersionControlLocked) {
		return true
	}

	return false
}

// SetVersionControlLocked gets a reference to the given bool and assigns it to the VersionControlLocked field.
func (o *SourceControlSettingsPutRequestSettings) SetVersionControlLocked(v bool) {
	o.VersionControlLocked = &v
}

// GetForceUuidMapping returns the ForceUuidMapping field value if set, zero value otherwise.
func (o *SourceControlSettingsPutRequestSettings) GetForceUuidMapping() bool {
	if o == nil || IsNil(o.ForceUuidMapping) {
		var ret bool
		return ret
	}
	return *o.ForceUuidMapping
}

// GetForceUuidMappingOk returns a tuple with the ForceUuidMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceControlSettingsPutRequestSettings) GetForceUuidMappingOk() (*bool, bool) {
	if o == nil || IsNil(o.ForceUuidMapping) {
		return nil, false
	}
	return o.ForceUuidMapping, true
}

// HasForceUuidMapping returns a boolean if a field has been set.
func (o *SourceControlSettingsPutRequestSettings) HasForceUuidMapping() bool {
	if o != nil && !IsNil(o.ForceUuidMapping) {
		return true
	}

	return false
}

// SetForceUuidMapping gets a reference to the given bool and assigns it to the ForceUuidMapping field.
func (o *SourceControlSettingsPutRequestSettings) SetForceUuidMapping(v bool) {
	o.ForceUuidMapping = &v
}

func (o SourceControlSettingsPutRequestSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceControlSettingsPutRequestSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoBranchNamingEnabled) {
		toSerialize["auto_branch_naming_enabled"] = o.AutoBranchNamingEnabled
	}
	if !IsNil(o.CustomPullRequestTemplateEnabled) {
		toSerialize["custom_pull_request_template_enabled"] = o.CustomPullRequestTemplateEnabled
	}
	if !IsNil(o.CustomPullRequestTemplate) {
		toSerialize["custom_pull_request_template"] = o.CustomPullRequestTemplate
	}
	if !IsNil(o.VersionControlLocked) {
		toSerialize["version_control_locked"] = o.VersionControlLocked
	}
	if !IsNil(o.ForceUuidMapping) {
		toSerialize["force_uuid_mapping"] = o.ForceUuidMapping
	}
	return toSerialize, nil
}

type NullableSourceControlSettingsPutRequestSettings struct {
	value *SourceControlSettingsPutRequestSettings
	isSet bool
}

func (v NullableSourceControlSettingsPutRequestSettings) Get() *SourceControlSettingsPutRequestSettings {
	return v.value
}

func (v *NullableSourceControlSettingsPutRequestSettings) Set(val *SourceControlSettingsPutRequestSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceControlSettingsPutRequestSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceControlSettingsPutRequestSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceControlSettingsPutRequestSettings(val *SourceControlSettingsPutRequestSettings) *NullableSourceControlSettingsPutRequestSettings {
	return &NullableSourceControlSettingsPutRequestSettings{value: val, isSet: true}
}

func (v NullableSourceControlSettingsPutRequestSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceControlSettingsPutRequestSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


