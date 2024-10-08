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

// SourceControlTestConnectionGet200ResponseData Test connection response
type SourceControlTestConnectionGet200ResponseData struct {
	SourceControlTestConnectionGet200ResponseDataAnyOf *SourceControlTestConnectionGet200ResponseDataAnyOf
	SourceControlTestConnectionGet200ResponseDataAnyOf1 *SourceControlTestConnectionGet200ResponseDataAnyOf1
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SourceControlTestConnectionGet200ResponseData) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SourceControlTestConnectionGet200ResponseDataAnyOf
	err = json.Unmarshal(data, &dst.SourceControlTestConnectionGet200ResponseDataAnyOf);
	if err == nil {
		jsonSourceControlTestConnectionGet200ResponseDataAnyOf, _ := json.Marshal(dst.SourceControlTestConnectionGet200ResponseDataAnyOf)
		if string(jsonSourceControlTestConnectionGet200ResponseDataAnyOf) == "{}" { // empty struct
			dst.SourceControlTestConnectionGet200ResponseDataAnyOf = nil
		} else {
			return nil // data stored in dst.SourceControlTestConnectionGet200ResponseDataAnyOf, return on the first match
		}
	} else {
		dst.SourceControlTestConnectionGet200ResponseDataAnyOf = nil
	}

	// try to unmarshal JSON data into SourceControlTestConnectionGet200ResponseDataAnyOf1
	err = json.Unmarshal(data, &dst.SourceControlTestConnectionGet200ResponseDataAnyOf1);
	if err == nil {
		jsonSourceControlTestConnectionGet200ResponseDataAnyOf1, _ := json.Marshal(dst.SourceControlTestConnectionGet200ResponseDataAnyOf1)
		if string(jsonSourceControlTestConnectionGet200ResponseDataAnyOf1) == "{}" { // empty struct
			dst.SourceControlTestConnectionGet200ResponseDataAnyOf1 = nil
		} else {
			return nil // data stored in dst.SourceControlTestConnectionGet200ResponseDataAnyOf1, return on the first match
		}
	} else {
		dst.SourceControlTestConnectionGet200ResponseDataAnyOf1 = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SourceControlTestConnectionGet200ResponseData)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SourceControlTestConnectionGet200ResponseData) MarshalJSON() ([]byte, error) {
	if src.SourceControlTestConnectionGet200ResponseDataAnyOf != nil {
		return json.Marshal(&src.SourceControlTestConnectionGet200ResponseDataAnyOf)
	}

	if src.SourceControlTestConnectionGet200ResponseDataAnyOf1 != nil {
		return json.Marshal(&src.SourceControlTestConnectionGet200ResponseDataAnyOf1)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSourceControlTestConnectionGet200ResponseData struct {
	value *SourceControlTestConnectionGet200ResponseData
	isSet bool
}

func (v NullableSourceControlTestConnectionGet200ResponseData) Get() *SourceControlTestConnectionGet200ResponseData {
	return v.value
}

func (v *NullableSourceControlTestConnectionGet200ResponseData) Set(val *SourceControlTestConnectionGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceControlTestConnectionGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceControlTestConnectionGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceControlTestConnectionGet200ResponseData(val *SourceControlTestConnectionGet200ResponseData) *NullableSourceControlTestConnectionGet200ResponseData {
	return &NullableSourceControlTestConnectionGet200ResponseData{value: val, isSet: true}
}

func (v NullableSourceControlTestConnectionGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceControlTestConnectionGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


