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

// PermissionsListObjectsPostRequestSubject - struct for PermissionsListObjectsPostRequestSubject
type PermissionsListObjectsPostRequestSubject struct {
	PermissionsListObjectsPostRequestSubjectOneOf *PermissionsListObjectsPostRequestSubjectOneOf
	PermissionsListObjectsPostRequestSubjectOneOf1 *PermissionsListObjectsPostRequestSubjectOneOf1
}

// PermissionsListObjectsPostRequestSubjectOneOfAsPermissionsListObjectsPostRequestSubject is a convenience function that returns PermissionsListObjectsPostRequestSubjectOneOf wrapped in PermissionsListObjectsPostRequestSubject
func PermissionsListObjectsPostRequestSubjectOneOfAsPermissionsListObjectsPostRequestSubject(v *PermissionsListObjectsPostRequestSubjectOneOf) PermissionsListObjectsPostRequestSubject {
	return PermissionsListObjectsPostRequestSubject{
		PermissionsListObjectsPostRequestSubjectOneOf: v,
	}
}

// PermissionsListObjectsPostRequestSubjectOneOf1AsPermissionsListObjectsPostRequestSubject is a convenience function that returns PermissionsListObjectsPostRequestSubjectOneOf1 wrapped in PermissionsListObjectsPostRequestSubject
func PermissionsListObjectsPostRequestSubjectOneOf1AsPermissionsListObjectsPostRequestSubject(v *PermissionsListObjectsPostRequestSubjectOneOf1) PermissionsListObjectsPostRequestSubject {
	return PermissionsListObjectsPostRequestSubject{
		PermissionsListObjectsPostRequestSubjectOneOf1: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PermissionsListObjectsPostRequestSubject) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PermissionsListObjectsPostRequestSubjectOneOf
	err = newStrictDecoder(data).Decode(&dst.PermissionsListObjectsPostRequestSubjectOneOf)
	if err == nil {
		jsonPermissionsListObjectsPostRequestSubjectOneOf, _ := json.Marshal(dst.PermissionsListObjectsPostRequestSubjectOneOf)
		if string(jsonPermissionsListObjectsPostRequestSubjectOneOf) == "{}" { // empty struct
			dst.PermissionsListObjectsPostRequestSubjectOneOf = nil
		} else {
			match++
		}
	} else {
		dst.PermissionsListObjectsPostRequestSubjectOneOf = nil
	}

	// try to unmarshal data into PermissionsListObjectsPostRequestSubjectOneOf1
	err = newStrictDecoder(data).Decode(&dst.PermissionsListObjectsPostRequestSubjectOneOf1)
	if err == nil {
		jsonPermissionsListObjectsPostRequestSubjectOneOf1, _ := json.Marshal(dst.PermissionsListObjectsPostRequestSubjectOneOf1)
		if string(jsonPermissionsListObjectsPostRequestSubjectOneOf1) == "{}" { // empty struct
			dst.PermissionsListObjectsPostRequestSubjectOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.PermissionsListObjectsPostRequestSubjectOneOf1 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PermissionsListObjectsPostRequestSubjectOneOf = nil
		dst.PermissionsListObjectsPostRequestSubjectOneOf1 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PermissionsListObjectsPostRequestSubject)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PermissionsListObjectsPostRequestSubject)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PermissionsListObjectsPostRequestSubject) MarshalJSON() ([]byte, error) {
	if src.PermissionsListObjectsPostRequestSubjectOneOf != nil {
		return json.Marshal(&src.PermissionsListObjectsPostRequestSubjectOneOf)
	}

	if src.PermissionsListObjectsPostRequestSubjectOneOf1 != nil {
		return json.Marshal(&src.PermissionsListObjectsPostRequestSubjectOneOf1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PermissionsListObjectsPostRequestSubject) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.PermissionsListObjectsPostRequestSubjectOneOf != nil {
		return obj.PermissionsListObjectsPostRequestSubjectOneOf
	}

	if obj.PermissionsListObjectsPostRequestSubjectOneOf1 != nil {
		return obj.PermissionsListObjectsPostRequestSubjectOneOf1
	}

	// all schemas are nil
	return nil
}

type NullablePermissionsListObjectsPostRequestSubject struct {
	value *PermissionsListObjectsPostRequestSubject
	isSet bool
}

func (v NullablePermissionsListObjectsPostRequestSubject) Get() *PermissionsListObjectsPostRequestSubject {
	return v.value
}

func (v *NullablePermissionsListObjectsPostRequestSubject) Set(val *PermissionsListObjectsPostRequestSubject) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionsListObjectsPostRequestSubject) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionsListObjectsPostRequestSubject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionsListObjectsPostRequestSubject(val *PermissionsListObjectsPostRequestSubject) *NullablePermissionsListObjectsPostRequestSubject {
	return &NullablePermissionsListObjectsPostRequestSubject{value: val, isSet: true}
}

func (v NullablePermissionsListObjectsPostRequestSubject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionsListObjectsPostRequestSubject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


