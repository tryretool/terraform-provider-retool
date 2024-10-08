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

// checks if the ConfigurationVariable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigurationVariable{}

// ConfigurationVariable struct for ConfigurationVariable
type ConfigurationVariable struct {
	// The ID of the configuration variable
	Id string `json:"id"`
	// The name of the configuration variable
	Name string `json:"name"`
	// The description of the configuration variable
	Description *string `json:"description,omitempty"`
	// Whether the configuration variable is a secret
	Secret bool `json:"secret"`
	Values []ConfigurationVariablesGet200ResponseDataInnerValuesInner `json:"values"`
	AdditionalProperties map[string]interface{}
}

type _ConfigurationVariable ConfigurationVariable

// NewConfigurationVariable instantiates a new ConfigurationVariable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationVariable(id string, name string, secret bool, values []ConfigurationVariablesGet200ResponseDataInnerValuesInner) *ConfigurationVariable {
	this := ConfigurationVariable{}
	this.Id = id
	this.Name = name
	this.Secret = secret
	this.Values = values
	return &this
}

// NewConfigurationVariableWithDefaults instantiates a new ConfigurationVariable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationVariableWithDefaults() *ConfigurationVariable {
	this := ConfigurationVariable{}
	return &this
}

// GetId returns the Id field value
func (o *ConfigurationVariable) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConfigurationVariable) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConfigurationVariable) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ConfigurationVariable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConfigurationVariable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConfigurationVariable) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConfigurationVariable) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationVariable) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConfigurationVariable) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConfigurationVariable) SetDescription(v string) {
	o.Description = &v
}

// GetSecret returns the Secret field value
func (o *ConfigurationVariable) GetSecret() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *ConfigurationVariable) GetSecretOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *ConfigurationVariable) SetSecret(v bool) {
	o.Secret = v
}

// GetValues returns the Values field value
func (o *ConfigurationVariable) GetValues() []ConfigurationVariablesGet200ResponseDataInnerValuesInner {
	if o == nil {
		var ret []ConfigurationVariablesGet200ResponseDataInnerValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *ConfigurationVariable) GetValuesOk() ([]ConfigurationVariablesGet200ResponseDataInnerValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *ConfigurationVariable) SetValues(v []ConfigurationVariablesGet200ResponseDataInnerValuesInner) {
	o.Values = v
}

func (o ConfigurationVariable) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigurationVariable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
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

func (o *ConfigurationVariable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varConfigurationVariable := _ConfigurationVariable{}

	err = json.Unmarshal(data, &varConfigurationVariable)

	if err != nil {
		return err
	}

	*o = ConfigurationVariable(varConfigurationVariable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "secret")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConfigurationVariable struct {
	value *ConfigurationVariable
	isSet bool
}

func (v NullableConfigurationVariable) Get() *ConfigurationVariable {
	return v.value
}

func (v *NullableConfigurationVariable) Set(val *ConfigurationVariable) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationVariable) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationVariable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationVariable(val *ConfigurationVariable) *NullableConfigurationVariable {
	return &NullableConfigurationVariable{value: val, isSet: true}
}

func (v NullableConfigurationVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationVariable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


