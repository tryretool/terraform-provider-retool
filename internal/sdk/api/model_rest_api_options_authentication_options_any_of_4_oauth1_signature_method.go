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

// RestAPIOptionsAuthenticationOptionsAnyOf4Oauth1SignatureMethod struct for RestAPIOptionsAuthenticationOptionsAnyOf4Oauth1SignatureMethod
type RestAPIOptionsAuthenticationOptionsAnyOf4Oauth1SignatureMethod struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RestAPIOptionsAuthenticationOptionsAnyOf4Oauth1SignatureMethod) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(RestAPIOptionsAuthenticationOptionsAnyOf4Oauth1SignatureMethod)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RestAPIOptionsAuthenticationOptionsAnyOf4Oauth1SignatureMethod) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRestAPIOptionsAuthenticationOptionsAnyOf4Oauth1SignatureMethod struct {
	value *RestAPIOptionsAuthenticationOptionsAnyOf4Oauth1SignatureMethod
	isSet bool
}

func (v NullableRestAPIOptionsAuthenticationOptionsAnyOf4Oauth1SignatureMethod) Get() *RestAPIOptionsAuthenticationOptionsAnyOf4Oauth1SignatureMethod {
	return v.value
}

func (v *NullableRestAPIOptionsAuthenticationOptionsAnyOf4Oauth1SignatureMethod) Set(val *RestAPIOptionsAuthenticationOptionsAnyOf4Oauth1SignatureMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableRestAPIOptionsAuthenticationOptionsAnyOf4Oauth1SignatureMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableRestAPIOptionsAuthenticationOptionsAnyOf4Oauth1SignatureMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestAPIOptionsAuthenticationOptionsAnyOf4Oauth1SignatureMethod(val *RestAPIOptionsAuthenticationOptionsAnyOf4Oauth1SignatureMethod) *NullableRestAPIOptionsAuthenticationOptionsAnyOf4Oauth1SignatureMethod {
	return &NullableRestAPIOptionsAuthenticationOptionsAnyOf4Oauth1SignatureMethod{value: val, isSet: true}
}

func (v NullableRestAPIOptionsAuthenticationOptionsAnyOf4Oauth1SignatureMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestAPIOptionsAuthenticationOptionsAnyOf4Oauth1SignatureMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


