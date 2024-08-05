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

// SourceControlConfigGet200ResponseData - Source Control Provider Config
type SourceControlConfigGet200ResponseData struct {
	SourceControlConfigGet200ResponseDataOneOf *SourceControlConfigGet200ResponseDataOneOf
	SourceControlConfigGet200ResponseDataOneOf1 *SourceControlConfigGet200ResponseDataOneOf1
	SourceControlConfigGet200ResponseDataOneOf2 *SourceControlConfigGet200ResponseDataOneOf2
	SourceControlConfigGet200ResponseDataOneOf3 *SourceControlConfigGet200ResponseDataOneOf3
	SourceControlConfigGet200ResponseDataOneOf4 *SourceControlConfigGet200ResponseDataOneOf4
}

// SourceControlConfigGet200ResponseDataOneOfAsSourceControlConfigGet200ResponseData is a convenience function that returns SourceControlConfigGet200ResponseDataOneOf wrapped in SourceControlConfigGet200ResponseData
func SourceControlConfigGet200ResponseDataOneOfAsSourceControlConfigGet200ResponseData(v *SourceControlConfigGet200ResponseDataOneOf) SourceControlConfigGet200ResponseData {
	return SourceControlConfigGet200ResponseData{
		SourceControlConfigGet200ResponseDataOneOf: v,
	}
}

// SourceControlConfigGet200ResponseDataOneOf1AsSourceControlConfigGet200ResponseData is a convenience function that returns SourceControlConfigGet200ResponseDataOneOf1 wrapped in SourceControlConfigGet200ResponseData
func SourceControlConfigGet200ResponseDataOneOf1AsSourceControlConfigGet200ResponseData(v *SourceControlConfigGet200ResponseDataOneOf1) SourceControlConfigGet200ResponseData {
	return SourceControlConfigGet200ResponseData{
		SourceControlConfigGet200ResponseDataOneOf1: v,
	}
}

// SourceControlConfigGet200ResponseDataOneOf2AsSourceControlConfigGet200ResponseData is a convenience function that returns SourceControlConfigGet200ResponseDataOneOf2 wrapped in SourceControlConfigGet200ResponseData
func SourceControlConfigGet200ResponseDataOneOf2AsSourceControlConfigGet200ResponseData(v *SourceControlConfigGet200ResponseDataOneOf2) SourceControlConfigGet200ResponseData {
	return SourceControlConfigGet200ResponseData{
		SourceControlConfigGet200ResponseDataOneOf2: v,
	}
}

// SourceControlConfigGet200ResponseDataOneOf3AsSourceControlConfigGet200ResponseData is a convenience function that returns SourceControlConfigGet200ResponseDataOneOf3 wrapped in SourceControlConfigGet200ResponseData
func SourceControlConfigGet200ResponseDataOneOf3AsSourceControlConfigGet200ResponseData(v *SourceControlConfigGet200ResponseDataOneOf3) SourceControlConfigGet200ResponseData {
	return SourceControlConfigGet200ResponseData{
		SourceControlConfigGet200ResponseDataOneOf3: v,
	}
}

// SourceControlConfigGet200ResponseDataOneOf4AsSourceControlConfigGet200ResponseData is a convenience function that returns SourceControlConfigGet200ResponseDataOneOf4 wrapped in SourceControlConfigGet200ResponseData
func SourceControlConfigGet200ResponseDataOneOf4AsSourceControlConfigGet200ResponseData(v *SourceControlConfigGet200ResponseDataOneOf4) SourceControlConfigGet200ResponseData {
	return SourceControlConfigGet200ResponseData{
		SourceControlConfigGet200ResponseDataOneOf4: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SourceControlConfigGet200ResponseData) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SourceControlConfigGet200ResponseDataOneOf
	err = newStrictDecoder(data).Decode(&dst.SourceControlConfigGet200ResponseDataOneOf)
	if err == nil {
		jsonSourceControlConfigGet200ResponseDataOneOf, _ := json.Marshal(dst.SourceControlConfigGet200ResponseDataOneOf)
		if string(jsonSourceControlConfigGet200ResponseDataOneOf) == "{}" { // empty struct
			dst.SourceControlConfigGet200ResponseDataOneOf = nil
		} else {
			match++
		}
	} else {
		dst.SourceControlConfigGet200ResponseDataOneOf = nil
	}

	// try to unmarshal data into SourceControlConfigGet200ResponseDataOneOf1
	err = newStrictDecoder(data).Decode(&dst.SourceControlConfigGet200ResponseDataOneOf1)
	if err == nil {
		jsonSourceControlConfigGet200ResponseDataOneOf1, _ := json.Marshal(dst.SourceControlConfigGet200ResponseDataOneOf1)
		if string(jsonSourceControlConfigGet200ResponseDataOneOf1) == "{}" { // empty struct
			dst.SourceControlConfigGet200ResponseDataOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.SourceControlConfigGet200ResponseDataOneOf1 = nil
	}

	// try to unmarshal data into SourceControlConfigGet200ResponseDataOneOf2
	err = newStrictDecoder(data).Decode(&dst.SourceControlConfigGet200ResponseDataOneOf2)
	if err == nil {
		jsonSourceControlConfigGet200ResponseDataOneOf2, _ := json.Marshal(dst.SourceControlConfigGet200ResponseDataOneOf2)
		if string(jsonSourceControlConfigGet200ResponseDataOneOf2) == "{}" { // empty struct
			dst.SourceControlConfigGet200ResponseDataOneOf2 = nil
		} else {
			match++
		}
	} else {
		dst.SourceControlConfigGet200ResponseDataOneOf2 = nil
	}

	// try to unmarshal data into SourceControlConfigGet200ResponseDataOneOf3
	err = newStrictDecoder(data).Decode(&dst.SourceControlConfigGet200ResponseDataOneOf3)
	if err == nil {
		jsonSourceControlConfigGet200ResponseDataOneOf3, _ := json.Marshal(dst.SourceControlConfigGet200ResponseDataOneOf3)
		if string(jsonSourceControlConfigGet200ResponseDataOneOf3) == "{}" { // empty struct
			dst.SourceControlConfigGet200ResponseDataOneOf3 = nil
		} else {
			match++
		}
	} else {
		dst.SourceControlConfigGet200ResponseDataOneOf3 = nil
	}

	// try to unmarshal data into SourceControlConfigGet200ResponseDataOneOf4
	err = newStrictDecoder(data).Decode(&dst.SourceControlConfigGet200ResponseDataOneOf4)
	if err == nil {
		jsonSourceControlConfigGet200ResponseDataOneOf4, _ := json.Marshal(dst.SourceControlConfigGet200ResponseDataOneOf4)
		if string(jsonSourceControlConfigGet200ResponseDataOneOf4) == "{}" { // empty struct
			dst.SourceControlConfigGet200ResponseDataOneOf4 = nil
		} else {
			match++
		}
	} else {
		dst.SourceControlConfigGet200ResponseDataOneOf4 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SourceControlConfigGet200ResponseDataOneOf = nil
		dst.SourceControlConfigGet200ResponseDataOneOf1 = nil
		dst.SourceControlConfigGet200ResponseDataOneOf2 = nil
		dst.SourceControlConfigGet200ResponseDataOneOf3 = nil
		dst.SourceControlConfigGet200ResponseDataOneOf4 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SourceControlConfigGet200ResponseData)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SourceControlConfigGet200ResponseData)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SourceControlConfigGet200ResponseData) MarshalJSON() ([]byte, error) {
	if src.SourceControlConfigGet200ResponseDataOneOf != nil {
		return json.Marshal(&src.SourceControlConfigGet200ResponseDataOneOf)
	}

	if src.SourceControlConfigGet200ResponseDataOneOf1 != nil {
		return json.Marshal(&src.SourceControlConfigGet200ResponseDataOneOf1)
	}

	if src.SourceControlConfigGet200ResponseDataOneOf2 != nil {
		return json.Marshal(&src.SourceControlConfigGet200ResponseDataOneOf2)
	}

	if src.SourceControlConfigGet200ResponseDataOneOf3 != nil {
		return json.Marshal(&src.SourceControlConfigGet200ResponseDataOneOf3)
	}

	if src.SourceControlConfigGet200ResponseDataOneOf4 != nil {
		return json.Marshal(&src.SourceControlConfigGet200ResponseDataOneOf4)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SourceControlConfigGet200ResponseData) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.SourceControlConfigGet200ResponseDataOneOf != nil {
		return obj.SourceControlConfigGet200ResponseDataOneOf
	}

	if obj.SourceControlConfigGet200ResponseDataOneOf1 != nil {
		return obj.SourceControlConfigGet200ResponseDataOneOf1
	}

	if obj.SourceControlConfigGet200ResponseDataOneOf2 != nil {
		return obj.SourceControlConfigGet200ResponseDataOneOf2
	}

	if obj.SourceControlConfigGet200ResponseDataOneOf3 != nil {
		return obj.SourceControlConfigGet200ResponseDataOneOf3
	}

	if obj.SourceControlConfigGet200ResponseDataOneOf4 != nil {
		return obj.SourceControlConfigGet200ResponseDataOneOf4
	}

	// all schemas are nil
	return nil
}

type NullableSourceControlConfigGet200ResponseData struct {
	value *SourceControlConfigGet200ResponseData
	isSet bool
}

func (v NullableSourceControlConfigGet200ResponseData) Get() *SourceControlConfigGet200ResponseData {
	return v.value
}

func (v *NullableSourceControlConfigGet200ResponseData) Set(val *SourceControlConfigGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceControlConfigGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceControlConfigGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceControlConfigGet200ResponseData(val *SourceControlConfigGet200ResponseData) *NullableSourceControlConfigGet200ResponseData {
	return &NullableSourceControlConfigGet200ResponseData{value: val, isSet: true}
}

func (v NullableSourceControlConfigGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceControlConfigGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


