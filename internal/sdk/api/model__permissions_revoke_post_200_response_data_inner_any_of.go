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

// checks if the PermissionsRevokePost200ResponseDataInnerAnyOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionsRevokePost200ResponseDataInnerAnyOf{}

// PermissionsRevokePost200ResponseDataInnerAnyOf struct for PermissionsRevokePost200ResponseDataInnerAnyOf
type PermissionsRevokePost200ResponseDataInnerAnyOf struct {
	Type string `json:"type"`
	// The id of the folder
	Id string `json:"id"`
	// The access level of the folder
	AccessLevel string `json:"access_level"`
	AdditionalProperties map[string]interface{}
}

type _PermissionsRevokePost200ResponseDataInnerAnyOf PermissionsRevokePost200ResponseDataInnerAnyOf

// NewPermissionsRevokePost200ResponseDataInnerAnyOf instantiates a new PermissionsRevokePost200ResponseDataInnerAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionsRevokePost200ResponseDataInnerAnyOf(type_ string, id string, accessLevel string) *PermissionsRevokePost200ResponseDataInnerAnyOf {
	this := PermissionsRevokePost200ResponseDataInnerAnyOf{}
	this.Type = type_
	this.Id = id
	this.AccessLevel = accessLevel
	return &this
}

// NewPermissionsRevokePost200ResponseDataInnerAnyOfWithDefaults instantiates a new PermissionsRevokePost200ResponseDataInnerAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionsRevokePost200ResponseDataInnerAnyOfWithDefaults() *PermissionsRevokePost200ResponseDataInnerAnyOf {
	this := PermissionsRevokePost200ResponseDataInnerAnyOf{}
	return &this
}

// GetType returns the Type field value
func (o *PermissionsRevokePost200ResponseDataInnerAnyOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PermissionsRevokePost200ResponseDataInnerAnyOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PermissionsRevokePost200ResponseDataInnerAnyOf) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *PermissionsRevokePost200ResponseDataInnerAnyOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PermissionsRevokePost200ResponseDataInnerAnyOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PermissionsRevokePost200ResponseDataInnerAnyOf) SetId(v string) {
	o.Id = v
}

// GetAccessLevel returns the AccessLevel field value
func (o *PermissionsRevokePost200ResponseDataInnerAnyOf) GetAccessLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessLevel
}

// GetAccessLevelOk returns a tuple with the AccessLevel field value
// and a boolean to check if the value has been set.
func (o *PermissionsRevokePost200ResponseDataInnerAnyOf) GetAccessLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessLevel, true
}

// SetAccessLevel sets field value
func (o *PermissionsRevokePost200ResponseDataInnerAnyOf) SetAccessLevel(v string) {
	o.AccessLevel = v
}

func (o PermissionsRevokePost200ResponseDataInnerAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionsRevokePost200ResponseDataInnerAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["access_level"] = o.AccessLevel

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PermissionsRevokePost200ResponseDataInnerAnyOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"access_level",
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

	varPermissionsRevokePost200ResponseDataInnerAnyOf := _PermissionsRevokePost200ResponseDataInnerAnyOf{}

	err = json.Unmarshal(data, &varPermissionsRevokePost200ResponseDataInnerAnyOf)

	if err != nil {
		return err
	}

	*o = PermissionsRevokePost200ResponseDataInnerAnyOf(varPermissionsRevokePost200ResponseDataInnerAnyOf)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "access_level")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePermissionsRevokePost200ResponseDataInnerAnyOf struct {
	value *PermissionsRevokePost200ResponseDataInnerAnyOf
	isSet bool
}

func (v NullablePermissionsRevokePost200ResponseDataInnerAnyOf) Get() *PermissionsRevokePost200ResponseDataInnerAnyOf {
	return v.value
}

func (v *NullablePermissionsRevokePost200ResponseDataInnerAnyOf) Set(val *PermissionsRevokePost200ResponseDataInnerAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionsRevokePost200ResponseDataInnerAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionsRevokePost200ResponseDataInnerAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionsRevokePost200ResponseDataInnerAnyOf(val *PermissionsRevokePost200ResponseDataInnerAnyOf) *NullablePermissionsRevokePost200ResponseDataInnerAnyOf {
	return &NullablePermissionsRevokePost200ResponseDataInnerAnyOf{value: val, isSet: true}
}

func (v NullablePermissionsRevokePost200ResponseDataInnerAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionsRevokePost200ResponseDataInnerAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


