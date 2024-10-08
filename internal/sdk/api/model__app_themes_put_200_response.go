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

// checks if the AppThemesPut200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppThemesPut200Response{}

// AppThemesPut200Response struct for AppThemesPut200Response
type AppThemesPut200Response struct {
	// API request succeeded
	Success bool `json:"success"`
	Data AppThemesPut200ResponseData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _AppThemesPut200Response AppThemesPut200Response

// NewAppThemesPut200Response instantiates a new AppThemesPut200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppThemesPut200Response(success bool, data AppThemesPut200ResponseData) *AppThemesPut200Response {
	this := AppThemesPut200Response{}
	this.Success = success
	this.Data = data
	return &this
}

// NewAppThemesPut200ResponseWithDefaults instantiates a new AppThemesPut200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppThemesPut200ResponseWithDefaults() *AppThemesPut200Response {
	this := AppThemesPut200Response{}
	return &this
}

// GetSuccess returns the Success field value
func (o *AppThemesPut200Response) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *AppThemesPut200Response) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *AppThemesPut200Response) SetSuccess(v bool) {
	o.Success = v
}

// GetData returns the Data field value
func (o *AppThemesPut200Response) GetData() AppThemesPut200ResponseData {
	if o == nil {
		var ret AppThemesPut200ResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppThemesPut200Response) GetDataOk() (*AppThemesPut200ResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppThemesPut200Response) SetData(v AppThemesPut200ResponseData) {
	o.Data = v
}

func (o AppThemesPut200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppThemesPut200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppThemesPut200Response) UnmarshalJSON(data []byte) (err error) {
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

	varAppThemesPut200Response := _AppThemesPut200Response{}

	err = json.Unmarshal(data, &varAppThemesPut200Response)

	if err != nil {
		return err
	}

	*o = AppThemesPut200Response(varAppThemesPut200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppThemesPut200Response struct {
	value *AppThemesPut200Response
	isSet bool
}

func (v NullableAppThemesPut200Response) Get() *AppThemesPut200Response {
	return v.value
}

func (v *NullableAppThemesPut200Response) Set(val *AppThemesPut200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAppThemesPut200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAppThemesPut200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppThemesPut200Response(val *AppThemesPut200Response) *NullableAppThemesPut200Response {
	return &NullableAppThemesPut200Response{value: val, isSet: true}
}

func (v NullableAppThemesPut200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppThemesPut200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


