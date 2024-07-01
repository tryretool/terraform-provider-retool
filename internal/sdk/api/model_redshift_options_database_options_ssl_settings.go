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

// checks if the RedshiftOptionsDatabaseOptionsSslSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RedshiftOptionsDatabaseOptionsSslSettings{}

// RedshiftOptionsDatabaseOptionsSslSettings struct for RedshiftOptionsDatabaseOptionsSslSettings
type RedshiftOptionsDatabaseOptionsSslSettings struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// NewRedshiftOptionsDatabaseOptionsSslSettings instantiates a new RedshiftOptionsDatabaseOptionsSslSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedshiftOptionsDatabaseOptionsSslSettings() *RedshiftOptionsDatabaseOptionsSslSettings {
	this := RedshiftOptionsDatabaseOptionsSslSettings{}
	return &this
}

// NewRedshiftOptionsDatabaseOptionsSslSettingsWithDefaults instantiates a new RedshiftOptionsDatabaseOptionsSslSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedshiftOptionsDatabaseOptionsSslSettingsWithDefaults() *RedshiftOptionsDatabaseOptionsSslSettings {
	this := RedshiftOptionsDatabaseOptionsSslSettings{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *RedshiftOptionsDatabaseOptionsSslSettings) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedshiftOptionsDatabaseOptionsSslSettings) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *RedshiftOptionsDatabaseOptionsSslSettings) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *RedshiftOptionsDatabaseOptionsSslSettings) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o RedshiftOptionsDatabaseOptionsSslSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RedshiftOptionsDatabaseOptionsSslSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableRedshiftOptionsDatabaseOptionsSslSettings struct {
	value *RedshiftOptionsDatabaseOptionsSslSettings
	isSet bool
}

func (v NullableRedshiftOptionsDatabaseOptionsSslSettings) Get() *RedshiftOptionsDatabaseOptionsSslSettings {
	return v.value
}

func (v *NullableRedshiftOptionsDatabaseOptionsSslSettings) Set(val *RedshiftOptionsDatabaseOptionsSslSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableRedshiftOptionsDatabaseOptionsSslSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableRedshiftOptionsDatabaseOptionsSslSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedshiftOptionsDatabaseOptionsSslSettings(val *RedshiftOptionsDatabaseOptionsSslSettings) *NullableRedshiftOptionsDatabaseOptionsSslSettings {
	return &NullableRedshiftOptionsDatabaseOptionsSslSettings{value: val, isSet: true}
}

func (v NullableRedshiftOptionsDatabaseOptionsSslSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedshiftOptionsDatabaseOptionsSslSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


