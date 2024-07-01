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

// checks if the ResourcesGet200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourcesGet200ResponseDataInner{}

// ResourcesGet200ResponseDataInner Resource
type ResourcesGet200ResponseDataInner struct {
	Id PermissionsListObjectsPost200ResponseDataInnerOneOf2Id `json:"id"`
	// The type of resource.
	Type string `json:"type"`
	DisplayName string `json:"display_name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type _ResourcesGet200ResponseDataInner ResourcesGet200ResponseDataInner

// NewResourcesGet200ResponseDataInner instantiates a new ResourcesGet200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourcesGet200ResponseDataInner(id PermissionsListObjectsPost200ResponseDataInnerOneOf2Id, type_ string, displayName string, createdAt string, updatedAt string) *ResourcesGet200ResponseDataInner {
	this := ResourcesGet200ResponseDataInner{}
	this.Id = id
	this.Type = type_
	this.DisplayName = displayName
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewResourcesGet200ResponseDataInnerWithDefaults instantiates a new ResourcesGet200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourcesGet200ResponseDataInnerWithDefaults() *ResourcesGet200ResponseDataInner {
	this := ResourcesGet200ResponseDataInner{}
	return &this
}

// GetId returns the Id field value
func (o *ResourcesGet200ResponseDataInner) GetId() PermissionsListObjectsPost200ResponseDataInnerOneOf2Id {
	if o == nil {
		var ret PermissionsListObjectsPost200ResponseDataInnerOneOf2Id
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResourcesGet200ResponseDataInner) GetIdOk() (*PermissionsListObjectsPost200ResponseDataInnerOneOf2Id, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResourcesGet200ResponseDataInner) SetId(v PermissionsListObjectsPost200ResponseDataInnerOneOf2Id) {
	o.Id = v
}

// GetType returns the Type field value
func (o *ResourcesGet200ResponseDataInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ResourcesGet200ResponseDataInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ResourcesGet200ResponseDataInner) SetType(v string) {
	o.Type = v
}

// GetDisplayName returns the DisplayName field value
func (o *ResourcesGet200ResponseDataInner) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *ResourcesGet200ResponseDataInner) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *ResourcesGet200ResponseDataInner) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ResourcesGet200ResponseDataInner) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ResourcesGet200ResponseDataInner) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ResourcesGet200ResponseDataInner) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ResourcesGet200ResponseDataInner) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ResourcesGet200ResponseDataInner) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ResourcesGet200ResponseDataInner) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o ResourcesGet200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourcesGet200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["display_name"] = o.DisplayName
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *ResourcesGet200ResponseDataInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"display_name",
		"created_at",
		"updated_at",
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

	varResourcesGet200ResponseDataInner := _ResourcesGet200ResponseDataInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResourcesGet200ResponseDataInner)

	if err != nil {
		return err
	}

	*o = ResourcesGet200ResponseDataInner(varResourcesGet200ResponseDataInner)

	return err
}

type NullableResourcesGet200ResponseDataInner struct {
	value *ResourcesGet200ResponseDataInner
	isSet bool
}

func (v NullableResourcesGet200ResponseDataInner) Get() *ResourcesGet200ResponseDataInner {
	return v.value
}

func (v *NullableResourcesGet200ResponseDataInner) Set(val *ResourcesGet200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcesGet200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcesGet200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcesGet200ResponseDataInner(val *ResourcesGet200ResponseDataInner) *NullableResourcesGet200ResponseDataInner {
	return &NullableResourcesGet200ResponseDataInner{value: val, isSet: true}
}

func (v NullableResourcesGet200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcesGet200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


