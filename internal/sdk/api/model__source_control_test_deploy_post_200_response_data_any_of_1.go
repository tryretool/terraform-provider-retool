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

// checks if the SourceControlTestDeployPost200ResponseDataAnyOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceControlTestDeployPost200ResponseDataAnyOf1{}

// SourceControlTestDeployPost200ResponseDataAnyOf1 struct for SourceControlTestDeployPost200ResponseDataAnyOf1
type SourceControlTestDeployPost200ResponseDataAnyOf1 struct {
	// Deployment succeeded
	Success bool `json:"success"`
}

type _SourceControlTestDeployPost200ResponseDataAnyOf1 SourceControlTestDeployPost200ResponseDataAnyOf1

// NewSourceControlTestDeployPost200ResponseDataAnyOf1 instantiates a new SourceControlTestDeployPost200ResponseDataAnyOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceControlTestDeployPost200ResponseDataAnyOf1(success bool) *SourceControlTestDeployPost200ResponseDataAnyOf1 {
	this := SourceControlTestDeployPost200ResponseDataAnyOf1{}
	this.Success = success
	return &this
}

// NewSourceControlTestDeployPost200ResponseDataAnyOf1WithDefaults instantiates a new SourceControlTestDeployPost200ResponseDataAnyOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceControlTestDeployPost200ResponseDataAnyOf1WithDefaults() *SourceControlTestDeployPost200ResponseDataAnyOf1 {
	this := SourceControlTestDeployPost200ResponseDataAnyOf1{}
	return &this
}

// GetSuccess returns the Success field value
func (o *SourceControlTestDeployPost200ResponseDataAnyOf1) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *SourceControlTestDeployPost200ResponseDataAnyOf1) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *SourceControlTestDeployPost200ResponseDataAnyOf1) SetSuccess(v bool) {
	o.Success = v
}

func (o SourceControlTestDeployPost200ResponseDataAnyOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceControlTestDeployPost200ResponseDataAnyOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	return toSerialize, nil
}

func (o *SourceControlTestDeployPost200ResponseDataAnyOf1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"success",
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

	varSourceControlTestDeployPost200ResponseDataAnyOf1 := _SourceControlTestDeployPost200ResponseDataAnyOf1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSourceControlTestDeployPost200ResponseDataAnyOf1)

	if err != nil {
		return err
	}

	*o = SourceControlTestDeployPost200ResponseDataAnyOf1(varSourceControlTestDeployPost200ResponseDataAnyOf1)

	return err
}

type NullableSourceControlTestDeployPost200ResponseDataAnyOf1 struct {
	value *SourceControlTestDeployPost200ResponseDataAnyOf1
	isSet bool
}

func (v NullableSourceControlTestDeployPost200ResponseDataAnyOf1) Get() *SourceControlTestDeployPost200ResponseDataAnyOf1 {
	return v.value
}

func (v *NullableSourceControlTestDeployPost200ResponseDataAnyOf1) Set(val *SourceControlTestDeployPost200ResponseDataAnyOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceControlTestDeployPost200ResponseDataAnyOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceControlTestDeployPost200ResponseDataAnyOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceControlTestDeployPost200ResponseDataAnyOf1(val *SourceControlTestDeployPost200ResponseDataAnyOf1) *NullableSourceControlTestDeployPost200ResponseDataAnyOf1 {
	return &NullableSourceControlTestDeployPost200ResponseDataAnyOf1{value: val, isSet: true}
}

func (v NullableSourceControlTestDeployPost200ResponseDataAnyOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceControlTestDeployPost200ResponseDataAnyOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


