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

// checks if the SsoConfigPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SsoConfigPostRequest{}

// SsoConfigPostRequest Shape of SSO config
type SsoConfigPostRequest struct {
	Data SsoConfigPostRequestData `json:"data"`
}

type _SsoConfigPostRequest SsoConfigPostRequest

// NewSsoConfigPostRequest instantiates a new SsoConfigPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsoConfigPostRequest(data SsoConfigPostRequestData) *SsoConfigPostRequest {
	this := SsoConfigPostRequest{}
	this.Data = data
	return &this
}

// NewSsoConfigPostRequestWithDefaults instantiates a new SsoConfigPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoConfigPostRequestWithDefaults() *SsoConfigPostRequest {
	this := SsoConfigPostRequest{}
	return &this
}

// GetData returns the Data field value
func (o *SsoConfigPostRequest) GetData() SsoConfigPostRequestData {
	if o == nil {
		var ret SsoConfigPostRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SsoConfigPostRequest) GetDataOk() (*SsoConfigPostRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SsoConfigPostRequest) SetData(v SsoConfigPostRequestData) {
	o.Data = v
}

func (o SsoConfigPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SsoConfigPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *SsoConfigPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varSsoConfigPostRequest := _SsoConfigPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSsoConfigPostRequest)

	if err != nil {
		return err
	}

	*o = SsoConfigPostRequest(varSsoConfigPostRequest)

	return err
}

type NullableSsoConfigPostRequest struct {
	value *SsoConfigPostRequest
	isSet bool
}

func (v NullableSsoConfigPostRequest) Get() *SsoConfigPostRequest {
	return v.value
}

func (v *NullableSsoConfigPostRequest) Set(val *SsoConfigPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoConfigPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoConfigPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoConfigPostRequest(val *SsoConfigPostRequest) *NullableSsoConfigPostRequest {
	return &NullableSsoConfigPostRequest{value: val, isSet: true}
}

func (v NullableSsoConfigPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoConfigPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


