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

// checks if the AddOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddOperation{}

// AddOperation struct for AddOperation
type AddOperation struct {
	Op string `json:"op"`
	Path string `json:"path"`
	// A JSON value
	Value interface{} `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AddOperation AddOperation

// NewAddOperation instantiates a new AddOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddOperation(op string, path string) *AddOperation {
	this := AddOperation{}
	this.Op = op
	this.Path = path
	return &this
}

// NewAddOperationWithDefaults instantiates a new AddOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddOperationWithDefaults() *AddOperation {
	this := AddOperation{}
	return &this
}

// GetOp returns the Op field value
func (o *AddOperation) GetOp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *AddOperation) GetOpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *AddOperation) SetOp(v string) {
	o.Op = v
}

// GetPath returns the Path field value
func (o *AddOperation) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *AddOperation) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *AddOperation) SetPath(v string) {
	o.Path = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddOperation) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddOperation) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AddOperation) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *AddOperation) SetValue(v interface{}) {
	o.Value = v
}

func (o AddOperation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["op"] = o.Op
	toSerialize["path"] = o.Path
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AddOperation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"op",
		"path",
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

	varAddOperation := _AddOperation{}

	err = json.Unmarshal(data, &varAddOperation)

	if err != nil {
		return err
	}

	*o = AddOperation(varAddOperation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "op")
		delete(additionalProperties, "path")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAddOperation struct {
	value *AddOperation
	isSet bool
}

func (v NullableAddOperation) Get() *AddOperation {
	return v.value
}

func (v *NullableAddOperation) Set(val *AddOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableAddOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableAddOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddOperation(val *AddOperation) *NullableAddOperation {
	return &NullableAddOperation{value: val, isSet: true}
}

func (v NullableAddOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


