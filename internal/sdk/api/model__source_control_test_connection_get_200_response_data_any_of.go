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

// checks if the SourceControlTestConnectionGet200ResponseDataAnyOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceControlTestConnectionGet200ResponseDataAnyOf{}

// SourceControlTestConnectionGet200ResponseDataAnyOf struct for SourceControlTestConnectionGet200ResponseDataAnyOf
type SourceControlTestConnectionGet200ResponseDataAnyOf struct {
	// Test connection failed
	Success bool `json:"success"`
	// Error message
	Message string `json:"message"`
}

type _SourceControlTestConnectionGet200ResponseDataAnyOf SourceControlTestConnectionGet200ResponseDataAnyOf

// NewSourceControlTestConnectionGet200ResponseDataAnyOf instantiates a new SourceControlTestConnectionGet200ResponseDataAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceControlTestConnectionGet200ResponseDataAnyOf(success bool, message string) *SourceControlTestConnectionGet200ResponseDataAnyOf {
	this := SourceControlTestConnectionGet200ResponseDataAnyOf{}
	this.Success = success
	this.Message = message
	return &this
}

// NewSourceControlTestConnectionGet200ResponseDataAnyOfWithDefaults instantiates a new SourceControlTestConnectionGet200ResponseDataAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceControlTestConnectionGet200ResponseDataAnyOfWithDefaults() *SourceControlTestConnectionGet200ResponseDataAnyOf {
	this := SourceControlTestConnectionGet200ResponseDataAnyOf{}
	return &this
}

// GetSuccess returns the Success field value
func (o *SourceControlTestConnectionGet200ResponseDataAnyOf) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *SourceControlTestConnectionGet200ResponseDataAnyOf) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *SourceControlTestConnectionGet200ResponseDataAnyOf) SetSuccess(v bool) {
	o.Success = v
}

// GetMessage returns the Message field value
func (o *SourceControlTestConnectionGet200ResponseDataAnyOf) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SourceControlTestConnectionGet200ResponseDataAnyOf) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SourceControlTestConnectionGet200ResponseDataAnyOf) SetMessage(v string) {
	o.Message = v
}

func (o SourceControlTestConnectionGet200ResponseDataAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceControlTestConnectionGet200ResponseDataAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *SourceControlTestConnectionGet200ResponseDataAnyOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"success",
		"message",
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

	varSourceControlTestConnectionGet200ResponseDataAnyOf := _SourceControlTestConnectionGet200ResponseDataAnyOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSourceControlTestConnectionGet200ResponseDataAnyOf)

	if err != nil {
		return err
	}

	*o = SourceControlTestConnectionGet200ResponseDataAnyOf(varSourceControlTestConnectionGet200ResponseDataAnyOf)

	return err
}

type NullableSourceControlTestConnectionGet200ResponseDataAnyOf struct {
	value *SourceControlTestConnectionGet200ResponseDataAnyOf
	isSet bool
}

func (v NullableSourceControlTestConnectionGet200ResponseDataAnyOf) Get() *SourceControlTestConnectionGet200ResponseDataAnyOf {
	return v.value
}

func (v *NullableSourceControlTestConnectionGet200ResponseDataAnyOf) Set(val *SourceControlTestConnectionGet200ResponseDataAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceControlTestConnectionGet200ResponseDataAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceControlTestConnectionGet200ResponseDataAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceControlTestConnectionGet200ResponseDataAnyOf(val *SourceControlTestConnectionGet200ResponseDataAnyOf) *NullableSourceControlTestConnectionGet200ResponseDataAnyOf {
	return &NullableSourceControlTestConnectionGet200ResponseDataAnyOf{value: val, isSet: true}
}

func (v NullableSourceControlTestConnectionGet200ResponseDataAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceControlTestConnectionGet200ResponseDataAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


