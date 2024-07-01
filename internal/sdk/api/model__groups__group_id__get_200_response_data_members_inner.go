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

// checks if the GroupsGroupIdGet200ResponseDataMembersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupsGroupIdGet200ResponseDataMembersInner{}

// GroupsGroupIdGet200ResponseDataMembersInner struct for GroupsGroupIdGet200ResponseDataMembersInner
type GroupsGroupIdGet200ResponseDataMembersInner struct {
	Id string `json:"id"`
	// The email of the user
	Email string `json:"email"`
	// Whether the user is a group admin
	IsGroupAdmin bool `json:"is_group_admin"`
}

type _GroupsGroupIdGet200ResponseDataMembersInner GroupsGroupIdGet200ResponseDataMembersInner

// NewGroupsGroupIdGet200ResponseDataMembersInner instantiates a new GroupsGroupIdGet200ResponseDataMembersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupsGroupIdGet200ResponseDataMembersInner(id string, email string, isGroupAdmin bool) *GroupsGroupIdGet200ResponseDataMembersInner {
	this := GroupsGroupIdGet200ResponseDataMembersInner{}
	this.Id = id
	this.Email = email
	this.IsGroupAdmin = isGroupAdmin
	return &this
}

// NewGroupsGroupIdGet200ResponseDataMembersInnerWithDefaults instantiates a new GroupsGroupIdGet200ResponseDataMembersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupsGroupIdGet200ResponseDataMembersInnerWithDefaults() *GroupsGroupIdGet200ResponseDataMembersInner {
	this := GroupsGroupIdGet200ResponseDataMembersInner{}
	var isGroupAdmin bool = false
	this.IsGroupAdmin = isGroupAdmin
	return &this
}

// GetId returns the Id field value
func (o *GroupsGroupIdGet200ResponseDataMembersInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdGet200ResponseDataMembersInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GroupsGroupIdGet200ResponseDataMembersInner) SetId(v string) {
	o.Id = v
}

// GetEmail returns the Email field value
func (o *GroupsGroupIdGet200ResponseDataMembersInner) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdGet200ResponseDataMembersInner) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *GroupsGroupIdGet200ResponseDataMembersInner) SetEmail(v string) {
	o.Email = v
}

// GetIsGroupAdmin returns the IsGroupAdmin field value
func (o *GroupsGroupIdGet200ResponseDataMembersInner) GetIsGroupAdmin() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsGroupAdmin
}

// GetIsGroupAdminOk returns a tuple with the IsGroupAdmin field value
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdGet200ResponseDataMembersInner) GetIsGroupAdminOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsGroupAdmin, true
}

// SetIsGroupAdmin sets field value
func (o *GroupsGroupIdGet200ResponseDataMembersInner) SetIsGroupAdmin(v bool) {
	o.IsGroupAdmin = v
}

func (o GroupsGroupIdGet200ResponseDataMembersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupsGroupIdGet200ResponseDataMembersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["email"] = o.Email
	toSerialize["is_group_admin"] = o.IsGroupAdmin
	return toSerialize, nil
}

func (o *GroupsGroupIdGet200ResponseDataMembersInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"email",
		"is_group_admin",
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

	varGroupsGroupIdGet200ResponseDataMembersInner := _GroupsGroupIdGet200ResponseDataMembersInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGroupsGroupIdGet200ResponseDataMembersInner)

	if err != nil {
		return err
	}

	*o = GroupsGroupIdGet200ResponseDataMembersInner(varGroupsGroupIdGet200ResponseDataMembersInner)

	return err
}

type NullableGroupsGroupIdGet200ResponseDataMembersInner struct {
	value *GroupsGroupIdGet200ResponseDataMembersInner
	isSet bool
}

func (v NullableGroupsGroupIdGet200ResponseDataMembersInner) Get() *GroupsGroupIdGet200ResponseDataMembersInner {
	return v.value
}

func (v *NullableGroupsGroupIdGet200ResponseDataMembersInner) Set(val *GroupsGroupIdGet200ResponseDataMembersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupsGroupIdGet200ResponseDataMembersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupsGroupIdGet200ResponseDataMembersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupsGroupIdGet200ResponseDataMembersInner(val *GroupsGroupIdGet200ResponseDataMembersInner) *NullableGroupsGroupIdGet200ResponseDataMembersInner {
	return &NullableGroupsGroupIdGet200ResponseDataMembersInner{value: val, isSet: true}
}

func (v NullableGroupsGroupIdGet200ResponseDataMembersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupsGroupIdGet200ResponseDataMembersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


