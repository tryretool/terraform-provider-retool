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

// checks if the SpacesCopyElementsPost201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpacesCopyElementsPost201Response{}

// SpacesCopyElementsPost201Response Elements copied successfully
type SpacesCopyElementsPost201Response struct {
	// API request succeeded
	Success bool `json:"success"`
	Data SpacesCopyElementsPost201ResponseData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _SpacesCopyElementsPost201Response SpacesCopyElementsPost201Response

// NewSpacesCopyElementsPost201Response instantiates a new SpacesCopyElementsPost201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpacesCopyElementsPost201Response(success bool, data SpacesCopyElementsPost201ResponseData) *SpacesCopyElementsPost201Response {
	this := SpacesCopyElementsPost201Response{}
	this.Success = success
	this.Data = data
	return &this
}

// NewSpacesCopyElementsPost201ResponseWithDefaults instantiates a new SpacesCopyElementsPost201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpacesCopyElementsPost201ResponseWithDefaults() *SpacesCopyElementsPost201Response {
	this := SpacesCopyElementsPost201Response{}
	return &this
}

// GetSuccess returns the Success field value
func (o *SpacesCopyElementsPost201Response) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *SpacesCopyElementsPost201Response) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *SpacesCopyElementsPost201Response) SetSuccess(v bool) {
	o.Success = v
}

// GetData returns the Data field value
func (o *SpacesCopyElementsPost201Response) GetData() SpacesCopyElementsPost201ResponseData {
	if o == nil {
		var ret SpacesCopyElementsPost201ResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SpacesCopyElementsPost201Response) GetDataOk() (*SpacesCopyElementsPost201ResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SpacesCopyElementsPost201Response) SetData(v SpacesCopyElementsPost201ResponseData) {
	o.Data = v
}

func (o SpacesCopyElementsPost201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpacesCopyElementsPost201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SpacesCopyElementsPost201Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"success",
		"data",
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

	varSpacesCopyElementsPost201Response := _SpacesCopyElementsPost201Response{}

	err = json.Unmarshal(data, &varSpacesCopyElementsPost201Response)

	if err != nil {
		return err
	}

	*o = SpacesCopyElementsPost201Response(varSpacesCopyElementsPost201Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSpacesCopyElementsPost201Response struct {
	value *SpacesCopyElementsPost201Response
	isSet bool
}

func (v NullableSpacesCopyElementsPost201Response) Get() *SpacesCopyElementsPost201Response {
	return v.value
}

func (v *NullableSpacesCopyElementsPost201Response) Set(val *SpacesCopyElementsPost201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSpacesCopyElementsPost201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSpacesCopyElementsPost201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpacesCopyElementsPost201Response(val *SpacesCopyElementsPost201Response) *NullableSpacesCopyElementsPost201Response {
	return &NullableSpacesCopyElementsPost201Response{value: val, isSet: true}
}

func (v NullableSpacesCopyElementsPost201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpacesCopyElementsPost201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


