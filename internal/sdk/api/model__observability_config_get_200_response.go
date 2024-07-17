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

// checks if the ObservabilityConfigGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObservabilityConfigGet200Response{}

// ObservabilityConfigGet200Response struct for ObservabilityConfigGet200Response
type ObservabilityConfigGet200Response struct {
	// API request succeeded
	Success bool `json:"success"`
	Data []ObservabilityConfigGet200ResponseDataInner `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _ObservabilityConfigGet200Response ObservabilityConfigGet200Response

// NewObservabilityConfigGet200Response instantiates a new ObservabilityConfigGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObservabilityConfigGet200Response(success bool, data []ObservabilityConfigGet200ResponseDataInner) *ObservabilityConfigGet200Response {
	this := ObservabilityConfigGet200Response{}
	this.Success = success
	this.Data = data
	return &this
}

// NewObservabilityConfigGet200ResponseWithDefaults instantiates a new ObservabilityConfigGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObservabilityConfigGet200ResponseWithDefaults() *ObservabilityConfigGet200Response {
	this := ObservabilityConfigGet200Response{}
	return &this
}

// GetSuccess returns the Success field value
func (o *ObservabilityConfigGet200Response) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *ObservabilityConfigGet200Response) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *ObservabilityConfigGet200Response) SetSuccess(v bool) {
	o.Success = v
}

// GetData returns the Data field value
func (o *ObservabilityConfigGet200Response) GetData() []ObservabilityConfigGet200ResponseDataInner {
	if o == nil {
		var ret []ObservabilityConfigGet200ResponseDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ObservabilityConfigGet200Response) GetDataOk() ([]ObservabilityConfigGet200ResponseDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ObservabilityConfigGet200Response) SetData(v []ObservabilityConfigGet200ResponseDataInner) {
	o.Data = v
}

func (o ObservabilityConfigGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObservabilityConfigGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ObservabilityConfigGet200Response) UnmarshalJSON(data []byte) (err error) {
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

	varObservabilityConfigGet200Response := _ObservabilityConfigGet200Response{}

	err = json.Unmarshal(data, &varObservabilityConfigGet200Response)

	if err != nil {
		return err
	}

	*o = ObservabilityConfigGet200Response(varObservabilityConfigGet200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableObservabilityConfigGet200Response struct {
	value *ObservabilityConfigGet200Response
	isSet bool
}

func (v NullableObservabilityConfigGet200Response) Get() *ObservabilityConfigGet200Response {
	return v.value
}

func (v *NullableObservabilityConfigGet200Response) Set(val *ObservabilityConfigGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableObservabilityConfigGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableObservabilityConfigGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObservabilityConfigGet200Response(val *ObservabilityConfigGet200Response) *NullableObservabilityConfigGet200Response {
	return &NullableObservabilityConfigGet200Response{value: val, isSet: true}
}

func (v NullableObservabilityConfigGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObservabilityConfigGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

