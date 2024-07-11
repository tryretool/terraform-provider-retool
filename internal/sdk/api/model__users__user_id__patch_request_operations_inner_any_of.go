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

// checks if the UsersUserIdPatchRequestOperationsInnerAnyOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsersUserIdPatchRequestOperationsInnerAnyOf{}

// UsersUserIdPatchRequestOperationsInnerAnyOf struct for UsersUserIdPatchRequestOperationsInnerAnyOf
type UsersUserIdPatchRequestOperationsInnerAnyOf struct {
	Op string `json:"op"`
	Path string `json:"path"`
	// A JSON value
	Value interface{} `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UsersUserIdPatchRequestOperationsInnerAnyOf UsersUserIdPatchRequestOperationsInnerAnyOf

// NewUsersUserIdPatchRequestOperationsInnerAnyOf instantiates a new UsersUserIdPatchRequestOperationsInnerAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersUserIdPatchRequestOperationsInnerAnyOf(op string, path string) *UsersUserIdPatchRequestOperationsInnerAnyOf {
	this := UsersUserIdPatchRequestOperationsInnerAnyOf{}
	this.Op = op
	this.Path = path
	return &this
}

// NewUsersUserIdPatchRequestOperationsInnerAnyOfWithDefaults instantiates a new UsersUserIdPatchRequestOperationsInnerAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersUserIdPatchRequestOperationsInnerAnyOfWithDefaults() *UsersUserIdPatchRequestOperationsInnerAnyOf {
	this := UsersUserIdPatchRequestOperationsInnerAnyOf{}
	return &this
}

// GetOp returns the Op field value
func (o *UsersUserIdPatchRequestOperationsInnerAnyOf) GetOp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *UsersUserIdPatchRequestOperationsInnerAnyOf) GetOpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *UsersUserIdPatchRequestOperationsInnerAnyOf) SetOp(v string) {
	o.Op = v
}

// GetPath returns the Path field value
func (o *UsersUserIdPatchRequestOperationsInnerAnyOf) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *UsersUserIdPatchRequestOperationsInnerAnyOf) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *UsersUserIdPatchRequestOperationsInnerAnyOf) SetPath(v string) {
	o.Path = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsersUserIdPatchRequestOperationsInnerAnyOf) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsersUserIdPatchRequestOperationsInnerAnyOf) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *UsersUserIdPatchRequestOperationsInnerAnyOf) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *UsersUserIdPatchRequestOperationsInnerAnyOf) SetValue(v interface{}) {
	o.Value = v
}

func (o UsersUserIdPatchRequestOperationsInnerAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsersUserIdPatchRequestOperationsInnerAnyOf) ToMap() (map[string]interface{}, error) {
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

func (o *UsersUserIdPatchRequestOperationsInnerAnyOf) UnmarshalJSON(data []byte) (err error) {
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

	varUsersUserIdPatchRequestOperationsInnerAnyOf := _UsersUserIdPatchRequestOperationsInnerAnyOf{}

	err = json.Unmarshal(data, &varUsersUserIdPatchRequestOperationsInnerAnyOf)

	if err != nil {
		return err
	}

	*o = UsersUserIdPatchRequestOperationsInnerAnyOf(varUsersUserIdPatchRequestOperationsInnerAnyOf)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "op")
		delete(additionalProperties, "path")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUsersUserIdPatchRequestOperationsInnerAnyOf struct {
	value *UsersUserIdPatchRequestOperationsInnerAnyOf
	isSet bool
}

func (v NullableUsersUserIdPatchRequestOperationsInnerAnyOf) Get() *UsersUserIdPatchRequestOperationsInnerAnyOf {
	return v.value
}

func (v *NullableUsersUserIdPatchRequestOperationsInnerAnyOf) Set(val *UsersUserIdPatchRequestOperationsInnerAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersUserIdPatchRequestOperationsInnerAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersUserIdPatchRequestOperationsInnerAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersUserIdPatchRequestOperationsInnerAnyOf(val *UsersUserIdPatchRequestOperationsInnerAnyOf) *NullableUsersUserIdPatchRequestOperationsInnerAnyOf {
	return &NullableUsersUserIdPatchRequestOperationsInnerAnyOf{value: val, isSet: true}
}

func (v NullableUsersUserIdPatchRequestOperationsInnerAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersUserIdPatchRequestOperationsInnerAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


