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

// checks if the AccessRequestsAccessRequestIdGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRequestsAccessRequestIdGet200Response{}

// AccessRequestsAccessRequestIdGet200Response A single access request.
type AccessRequestsAccessRequestIdGet200Response struct {
	// API request succeeded
	Success bool `json:"success"`
	Data AccessRequestsGet200ResponseDataInner `json:"data"`
}

type _AccessRequestsAccessRequestIdGet200Response AccessRequestsAccessRequestIdGet200Response

// NewAccessRequestsAccessRequestIdGet200Response instantiates a new AccessRequestsAccessRequestIdGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRequestsAccessRequestIdGet200Response(success bool, data AccessRequestsGet200ResponseDataInner) *AccessRequestsAccessRequestIdGet200Response {
	this := AccessRequestsAccessRequestIdGet200Response{}
	this.Success = success
	this.Data = data
	return &this
}

// NewAccessRequestsAccessRequestIdGet200ResponseWithDefaults instantiates a new AccessRequestsAccessRequestIdGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRequestsAccessRequestIdGet200ResponseWithDefaults() *AccessRequestsAccessRequestIdGet200Response {
	this := AccessRequestsAccessRequestIdGet200Response{}
	return &this
}

// GetSuccess returns the Success field value
func (o *AccessRequestsAccessRequestIdGet200Response) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *AccessRequestsAccessRequestIdGet200Response) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *AccessRequestsAccessRequestIdGet200Response) SetSuccess(v bool) {
	o.Success = v
}

// GetData returns the Data field value
func (o *AccessRequestsAccessRequestIdGet200Response) GetData() AccessRequestsGet200ResponseDataInner {
	if o == nil {
		var ret AccessRequestsGet200ResponseDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AccessRequestsAccessRequestIdGet200Response) GetDataOk() (*AccessRequestsGet200ResponseDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AccessRequestsAccessRequestIdGet200Response) SetData(v AccessRequestsGet200ResponseDataInner) {
	o.Data = v
}

func (o AccessRequestsAccessRequestIdGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRequestsAccessRequestIdGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AccessRequestsAccessRequestIdGet200Response) UnmarshalJSON(data []byte) (err error) {
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

	varAccessRequestsAccessRequestIdGet200Response := _AccessRequestsAccessRequestIdGet200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccessRequestsAccessRequestIdGet200Response)

	if err != nil {
		return err
	}

	*o = AccessRequestsAccessRequestIdGet200Response(varAccessRequestsAccessRequestIdGet200Response)

	return err
}

type NullableAccessRequestsAccessRequestIdGet200Response struct {
	value *AccessRequestsAccessRequestIdGet200Response
	isSet bool
}

func (v NullableAccessRequestsAccessRequestIdGet200Response) Get() *AccessRequestsAccessRequestIdGet200Response {
	return v.value
}

func (v *NullableAccessRequestsAccessRequestIdGet200Response) Set(val *AccessRequestsAccessRequestIdGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequestsAccessRequestIdGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequestsAccessRequestIdGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequestsAccessRequestIdGet200Response(val *AccessRequestsAccessRequestIdGet200Response) *NullableAccessRequestsAccessRequestIdGet200Response {
	return &NullableAccessRequestsAccessRequestIdGet200Response{value: val, isSet: true}
}

func (v NullableAccessRequestsAccessRequestIdGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequestsAccessRequestIdGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


