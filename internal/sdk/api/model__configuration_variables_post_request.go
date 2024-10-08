/*
Retool API

Go to Settings > API to get started. Once you generate an API token, use bearer token authentication to make requests.

API version: 2.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// checks if the ConfigurationVariablesPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigurationVariablesPostRequest{}

// ConfigurationVariablesPostRequest struct for ConfigurationVariablesPostRequest
type ConfigurationVariablesPostRequest struct {
	// The name of the configuration variable
	Name string `json:"name"`
	// The description of the configuration variable
	Description *string `json:"description,omitempty"`
	// Whether the configuration variable is a secret
	Secret bool `json:"secret"`
	Values []ConfigurationVariablesGet200ResponseDataInnerValuesInner `json:"values"`
	AdditionalProperties map[string]interface{}
}

type _ConfigurationVariablesPostRequest ConfigurationVariablesPostRequest

// NewConfigurationVariablesPostRequest instantiates a new ConfigurationVariablesPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationVariablesPostRequest(name string, secret bool, values []ConfigurationVariablesGet200ResponseDataInnerValuesInner) *ConfigurationVariablesPostRequest {
	this := ConfigurationVariablesPostRequest{}
	this.Name = name
	this.Secret = secret
	this.Values = values
	return &this
}

// NewConfigurationVariablesPostRequestWithDefaults instantiates a new ConfigurationVariablesPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationVariablesPostRequestWithDefaults() *ConfigurationVariablesPostRequest {
	this := ConfigurationVariablesPostRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ConfigurationVariablesPostRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConfigurationVariablesPostRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConfigurationVariablesPostRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConfigurationVariablesPostRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationVariablesPostRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConfigurationVariablesPostRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConfigurationVariablesPostRequest) SetDescription(v string) {
	o.Description = &v
}

// GetSecret returns the Secret field value
func (o *ConfigurationVariablesPostRequest) GetSecret() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *ConfigurationVariablesPostRequest) GetSecretOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *ConfigurationVariablesPostRequest) SetSecret(v bool) {
	o.Secret = v
}

// GetValues returns the Values field value
func (o *ConfigurationVariablesPostRequest) GetValues() []ConfigurationVariablesGet200ResponseDataInnerValuesInner {
	if o == nil {
		var ret []ConfigurationVariablesGet200ResponseDataInnerValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *ConfigurationVariablesPostRequest) GetValuesOk() ([]ConfigurationVariablesGet200ResponseDataInnerValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *ConfigurationVariablesPostRequest) SetValues(v []ConfigurationVariablesGet200ResponseDataInnerValuesInner) {
	o.Values = v
}

func (o ConfigurationVariablesPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigurationVariablesPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["secret"] = o.Secret
	toSerialize["values"] = o.Values

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConfigurationVariablesPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"secret",
		"values",
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

	varConfigurationVariablesPostRequest := _ConfigurationVariablesPostRequest{}

	err = json.Unmarshal(data, &varConfigurationVariablesPostRequest)

	if err != nil {
		return err
	}

	*o = ConfigurationVariablesPostRequest(varConfigurationVariablesPostRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "secret")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConfigurationVariablesPostRequest struct {
	value *ConfigurationVariablesPostRequest
	isSet bool
}

func (v NullableConfigurationVariablesPostRequest) Get() *ConfigurationVariablesPostRequest {
	return v.value
}

func (v *NullableConfigurationVariablesPostRequest) Set(val *ConfigurationVariablesPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationVariablesPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationVariablesPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationVariablesPostRequest(val *ConfigurationVariablesPostRequest) *NullableConfigurationVariablesPostRequest {
	return &NullableConfigurationVariablesPostRequest{value: val, isSet: true}
}

func (v NullableConfigurationVariablesPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationVariablesPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


