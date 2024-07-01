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

// SourceControlConfigPost200ResponseData Created Source Control Provider Config
type SourceControlConfigPost200ResponseData struct {
	SourceControlConfigGet200ResponseDataAnyOf *SourceControlConfigGet200ResponseDataAnyOf
	SourceControlConfigGet200ResponseDataAnyOf1 *SourceControlConfigGet200ResponseDataAnyOf1
	SourceControlConfigGet200ResponseDataAnyOf2 *SourceControlConfigGet200ResponseDataAnyOf2
	SourceControlConfigGet200ResponseDataAnyOf3 *SourceControlConfigGet200ResponseDataAnyOf3
	SourceControlConfigGet200ResponseDataAnyOf4 *SourceControlConfigGet200ResponseDataAnyOf4
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SourceControlConfigPost200ResponseData) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SourceControlConfigGet200ResponseDataAnyOf
	err = json.Unmarshal(data, &dst.SourceControlConfigGet200ResponseDataAnyOf);
	if err == nil {
		jsonSourceControlConfigGet200ResponseDataAnyOf, _ := json.Marshal(dst.SourceControlConfigGet200ResponseDataAnyOf)
		if string(jsonSourceControlConfigGet200ResponseDataAnyOf) == "{}" { // empty struct
			dst.SourceControlConfigGet200ResponseDataAnyOf = nil
		} else {
			return nil // data stored in dst.SourceControlConfigGet200ResponseDataAnyOf, return on the first match
		}
	} else {
		dst.SourceControlConfigGet200ResponseDataAnyOf = nil
	}

	// try to unmarshal JSON data into SourceControlConfigGet200ResponseDataAnyOf1
	err = json.Unmarshal(data, &dst.SourceControlConfigGet200ResponseDataAnyOf1);
	if err == nil {
		jsonSourceControlConfigGet200ResponseDataAnyOf1, _ := json.Marshal(dst.SourceControlConfigGet200ResponseDataAnyOf1)
		if string(jsonSourceControlConfigGet200ResponseDataAnyOf1) == "{}" { // empty struct
			dst.SourceControlConfigGet200ResponseDataAnyOf1 = nil
		} else {
			return nil // data stored in dst.SourceControlConfigGet200ResponseDataAnyOf1, return on the first match
		}
	} else {
		dst.SourceControlConfigGet200ResponseDataAnyOf1 = nil
	}

	// try to unmarshal JSON data into SourceControlConfigGet200ResponseDataAnyOf2
	err = json.Unmarshal(data, &dst.SourceControlConfigGet200ResponseDataAnyOf2);
	if err == nil {
		jsonSourceControlConfigGet200ResponseDataAnyOf2, _ := json.Marshal(dst.SourceControlConfigGet200ResponseDataAnyOf2)
		if string(jsonSourceControlConfigGet200ResponseDataAnyOf2) == "{}" { // empty struct
			dst.SourceControlConfigGet200ResponseDataAnyOf2 = nil
		} else {
			return nil // data stored in dst.SourceControlConfigGet200ResponseDataAnyOf2, return on the first match
		}
	} else {
		dst.SourceControlConfigGet200ResponseDataAnyOf2 = nil
	}

	// try to unmarshal JSON data into SourceControlConfigGet200ResponseDataAnyOf3
	err = json.Unmarshal(data, &dst.SourceControlConfigGet200ResponseDataAnyOf3);
	if err == nil {
		jsonSourceControlConfigGet200ResponseDataAnyOf3, _ := json.Marshal(dst.SourceControlConfigGet200ResponseDataAnyOf3)
		if string(jsonSourceControlConfigGet200ResponseDataAnyOf3) == "{}" { // empty struct
			dst.SourceControlConfigGet200ResponseDataAnyOf3 = nil
		} else {
			return nil // data stored in dst.SourceControlConfigGet200ResponseDataAnyOf3, return on the first match
		}
	} else {
		dst.SourceControlConfigGet200ResponseDataAnyOf3 = nil
	}

	// try to unmarshal JSON data into SourceControlConfigGet200ResponseDataAnyOf4
	err = json.Unmarshal(data, &dst.SourceControlConfigGet200ResponseDataAnyOf4);
	if err == nil {
		jsonSourceControlConfigGet200ResponseDataAnyOf4, _ := json.Marshal(dst.SourceControlConfigGet200ResponseDataAnyOf4)
		if string(jsonSourceControlConfigGet200ResponseDataAnyOf4) == "{}" { // empty struct
			dst.SourceControlConfigGet200ResponseDataAnyOf4 = nil
		} else {
			return nil // data stored in dst.SourceControlConfigGet200ResponseDataAnyOf4, return on the first match
		}
	} else {
		dst.SourceControlConfigGet200ResponseDataAnyOf4 = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SourceControlConfigPost200ResponseData)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SourceControlConfigPost200ResponseData) MarshalJSON() ([]byte, error) {
	if src.SourceControlConfigGet200ResponseDataAnyOf != nil {
		return json.Marshal(&src.SourceControlConfigGet200ResponseDataAnyOf)
	}

	if src.SourceControlConfigGet200ResponseDataAnyOf1 != nil {
		return json.Marshal(&src.SourceControlConfigGet200ResponseDataAnyOf1)
	}

	if src.SourceControlConfigGet200ResponseDataAnyOf2 != nil {
		return json.Marshal(&src.SourceControlConfigGet200ResponseDataAnyOf2)
	}

	if src.SourceControlConfigGet200ResponseDataAnyOf3 != nil {
		return json.Marshal(&src.SourceControlConfigGet200ResponseDataAnyOf3)
	}

	if src.SourceControlConfigGet200ResponseDataAnyOf4 != nil {
		return json.Marshal(&src.SourceControlConfigGet200ResponseDataAnyOf4)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSourceControlConfigPost200ResponseData struct {
	value *SourceControlConfigPost200ResponseData
	isSet bool
}

func (v NullableSourceControlConfigPost200ResponseData) Get() *SourceControlConfigPost200ResponseData {
	return v.value
}

func (v *NullableSourceControlConfigPost200ResponseData) Set(val *SourceControlConfigPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceControlConfigPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceControlConfigPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceControlConfigPost200ResponseData(val *SourceControlConfigPost200ResponseData) *NullableSourceControlConfigPost200ResponseData {
	return &NullableSourceControlConfigPost200ResponseData{value: val, isSet: true}
}

func (v NullableSourceControlConfigPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceControlConfigPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


