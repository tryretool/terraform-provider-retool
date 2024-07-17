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

// checks if the FoldersFolderIdPatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FoldersFolderIdPatchRequest{}

// FoldersFolderIdPatchRequest struct for FoldersFolderIdPatchRequest
type FoldersFolderIdPatchRequest struct {
	// A list of operations to apply to the folder. See https://tools.ietf.org/html/rfc6902 for details.
	Operations []FoldersFolderIdPatchRequestOperationsInner `json:"operations"`
}

type _FoldersFolderIdPatchRequest FoldersFolderIdPatchRequest

// NewFoldersFolderIdPatchRequest instantiates a new FoldersFolderIdPatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFoldersFolderIdPatchRequest(operations []FoldersFolderIdPatchRequestOperationsInner) *FoldersFolderIdPatchRequest {
	this := FoldersFolderIdPatchRequest{}
	this.Operations = operations
	return &this
}

// NewFoldersFolderIdPatchRequestWithDefaults instantiates a new FoldersFolderIdPatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFoldersFolderIdPatchRequestWithDefaults() *FoldersFolderIdPatchRequest {
	this := FoldersFolderIdPatchRequest{}
	return &this
}

// GetOperations returns the Operations field value
func (o *FoldersFolderIdPatchRequest) GetOperations() []FoldersFolderIdPatchRequestOperationsInner {
	if o == nil {
		var ret []FoldersFolderIdPatchRequestOperationsInner
		return ret
	}

	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value
// and a boolean to check if the value has been set.
func (o *FoldersFolderIdPatchRequest) GetOperationsOk() ([]FoldersFolderIdPatchRequestOperationsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operations, true
}

// SetOperations sets field value
func (o *FoldersFolderIdPatchRequest) SetOperations(v []FoldersFolderIdPatchRequestOperationsInner) {
	o.Operations = v
}

func (o FoldersFolderIdPatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FoldersFolderIdPatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operations"] = o.Operations
	return toSerialize, nil
}

func (o *FoldersFolderIdPatchRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"operations",
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

	varFoldersFolderIdPatchRequest := _FoldersFolderIdPatchRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFoldersFolderIdPatchRequest)

	if err != nil {
		return err
	}

	*o = FoldersFolderIdPatchRequest(varFoldersFolderIdPatchRequest)

	return err
}

type NullableFoldersFolderIdPatchRequest struct {
	value *FoldersFolderIdPatchRequest
	isSet bool
}

func (v NullableFoldersFolderIdPatchRequest) Get() *FoldersFolderIdPatchRequest {
	return v.value
}

func (v *NullableFoldersFolderIdPatchRequest) Set(val *FoldersFolderIdPatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFoldersFolderIdPatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFoldersFolderIdPatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFoldersFolderIdPatchRequest(val *FoldersFolderIdPatchRequest) *NullableFoldersFolderIdPatchRequest {
	return &NullableFoldersFolderIdPatchRequest{value: val, isSet: true}
}

func (v NullableFoldersFolderIdPatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFoldersFolderIdPatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

