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

// checks if the FoldersPost409Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FoldersPost409Response{}

// FoldersPost409Response API error response
type FoldersPost409Response struct {
	// API request failed
	Success bool `json:"success"`
	// Error message
	Message string `json:"message"`
	AdditionalProperties map[string]interface{}
}

type _FoldersPost409Response FoldersPost409Response

// NewFoldersPost409Response instantiates a new FoldersPost409Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFoldersPost409Response(success bool, message string) *FoldersPost409Response {
	this := FoldersPost409Response{}
	this.Success = success
	this.Message = message
	return &this
}

// NewFoldersPost409ResponseWithDefaults instantiates a new FoldersPost409Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFoldersPost409ResponseWithDefaults() *FoldersPost409Response {
	this := FoldersPost409Response{}
	return &this
}

// GetSuccess returns the Success field value
func (o *FoldersPost409Response) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *FoldersPost409Response) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *FoldersPost409Response) SetSuccess(v bool) {
	o.Success = v
}

// GetMessage returns the Message field value
func (o *FoldersPost409Response) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *FoldersPost409Response) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *FoldersPost409Response) SetMessage(v string) {
	o.Message = v
}

func (o FoldersPost409Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FoldersPost409Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	toSerialize["message"] = o.Message

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FoldersPost409Response) UnmarshalJSON(data []byte) (err error) {
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

	varFoldersPost409Response := _FoldersPost409Response{}

	err = json.Unmarshal(data, &varFoldersPost409Response)

	if err != nil {
		return err
	}

	*o = FoldersPost409Response(varFoldersPost409Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		delete(additionalProperties, "message")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFoldersPost409Response struct {
	value *FoldersPost409Response
	isSet bool
}

func (v NullableFoldersPost409Response) Get() *FoldersPost409Response {
	return v.value
}

func (v *NullableFoldersPost409Response) Set(val *FoldersPost409Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFoldersPost409Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFoldersPost409Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFoldersPost409Response(val *FoldersPost409Response) *NullableFoldersPost409Response {
	return &NullableFoldersPost409Response{value: val, isSet: true}
}

func (v NullableFoldersPost409Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFoldersPost409Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


