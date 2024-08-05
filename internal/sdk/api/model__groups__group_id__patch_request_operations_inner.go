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

// GroupsGroupIdPatchRequestOperationsInner struct for GroupsGroupIdPatchRequestOperationsInner
type GroupsGroupIdPatchRequestOperationsInner struct {
	AddOperation *AddOperation
	RemoveOperation *RemoveOperation
	ReplaceOperation *ReplaceOperation
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *GroupsGroupIdPatchRequestOperationsInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AddOperation
	err = json.Unmarshal(data, &dst.AddOperation);
	if err == nil {
		jsonAddOperation, _ := json.Marshal(dst.AddOperation)
		if string(jsonAddOperation) == "{}" { // empty struct
			dst.AddOperation = nil
		} else {
			return nil // data stored in dst.AddOperation, return on the first match
		}
	} else {
		dst.AddOperation = nil
	}

	// try to unmarshal JSON data into RemoveOperation
	err = json.Unmarshal(data, &dst.RemoveOperation);
	if err == nil {
		jsonRemoveOperation, _ := json.Marshal(dst.RemoveOperation)
		if string(jsonRemoveOperation) == "{}" { // empty struct
			dst.RemoveOperation = nil
		} else {
			return nil // data stored in dst.RemoveOperation, return on the first match
		}
	} else {
		dst.RemoveOperation = nil
	}

	// try to unmarshal JSON data into ReplaceOperation
	err = json.Unmarshal(data, &dst.ReplaceOperation);
	if err == nil {
		jsonReplaceOperation, _ := json.Marshal(dst.ReplaceOperation)
		if string(jsonReplaceOperation) == "{}" { // empty struct
			dst.ReplaceOperation = nil
		} else {
			return nil // data stored in dst.ReplaceOperation, return on the first match
		}
	} else {
		dst.ReplaceOperation = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(GroupsGroupIdPatchRequestOperationsInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *GroupsGroupIdPatchRequestOperationsInner) MarshalJSON() ([]byte, error) {
	if src.AddOperation != nil {
		return json.Marshal(&src.AddOperation)
	}

	if src.RemoveOperation != nil {
		return json.Marshal(&src.RemoveOperation)
	}

	if src.ReplaceOperation != nil {
		return json.Marshal(&src.ReplaceOperation)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableGroupsGroupIdPatchRequestOperationsInner struct {
	value *GroupsGroupIdPatchRequestOperationsInner
	isSet bool
}

func (v NullableGroupsGroupIdPatchRequestOperationsInner) Get() *GroupsGroupIdPatchRequestOperationsInner {
	return v.value
}

func (v *NullableGroupsGroupIdPatchRequestOperationsInner) Set(val *GroupsGroupIdPatchRequestOperationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupsGroupIdPatchRequestOperationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupsGroupIdPatchRequestOperationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupsGroupIdPatchRequestOperationsInner(val *GroupsGroupIdPatchRequestOperationsInner) *NullableGroupsGroupIdPatchRequestOperationsInner {
	return &NullableGroupsGroupIdPatchRequestOperationsInner{value: val, isSet: true}
}

func (v NullableGroupsGroupIdPatchRequestOperationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupsGroupIdPatchRequestOperationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

