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

// checks if the GroupsGroupIdUserInvitesPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupsGroupIdUserInvitesPostRequest{}

// GroupsGroupIdUserInvitesPostRequest User Invites to add to the group
type GroupsGroupIdUserInvitesPostRequest struct {
	UserInviteIds []float32 `json:"userInviteIds"`
}

type _GroupsGroupIdUserInvitesPostRequest GroupsGroupIdUserInvitesPostRequest

// NewGroupsGroupIdUserInvitesPostRequest instantiates a new GroupsGroupIdUserInvitesPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupsGroupIdUserInvitesPostRequest(userInviteIds []float32) *GroupsGroupIdUserInvitesPostRequest {
	this := GroupsGroupIdUserInvitesPostRequest{}
	this.UserInviteIds = userInviteIds
	return &this
}

// NewGroupsGroupIdUserInvitesPostRequestWithDefaults instantiates a new GroupsGroupIdUserInvitesPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupsGroupIdUserInvitesPostRequestWithDefaults() *GroupsGroupIdUserInvitesPostRequest {
	this := GroupsGroupIdUserInvitesPostRequest{}
	return &this
}

// GetUserInviteIds returns the UserInviteIds field value
func (o *GroupsGroupIdUserInvitesPostRequest) GetUserInviteIds() []float32 {
	if o == nil {
		var ret []float32
		return ret
	}

	return o.UserInviteIds
}

// GetUserInviteIdsOk returns a tuple with the UserInviteIds field value
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdUserInvitesPostRequest) GetUserInviteIdsOk() ([]float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserInviteIds, true
}

// SetUserInviteIds sets field value
func (o *GroupsGroupIdUserInvitesPostRequest) SetUserInviteIds(v []float32) {
	o.UserInviteIds = v
}

func (o GroupsGroupIdUserInvitesPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupsGroupIdUserInvitesPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userInviteIds"] = o.UserInviteIds
	return toSerialize, nil
}

func (o *GroupsGroupIdUserInvitesPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"userInviteIds",
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

	varGroupsGroupIdUserInvitesPostRequest := _GroupsGroupIdUserInvitesPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGroupsGroupIdUserInvitesPostRequest)

	if err != nil {
		return err
	}

	*o = GroupsGroupIdUserInvitesPostRequest(varGroupsGroupIdUserInvitesPostRequest)

	return err
}

type NullableGroupsGroupIdUserInvitesPostRequest struct {
	value *GroupsGroupIdUserInvitesPostRequest
	isSet bool
}

func (v NullableGroupsGroupIdUserInvitesPostRequest) Get() *GroupsGroupIdUserInvitesPostRequest {
	return v.value
}

func (v *NullableGroupsGroupIdUserInvitesPostRequest) Set(val *GroupsGroupIdUserInvitesPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupsGroupIdUserInvitesPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupsGroupIdUserInvitesPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupsGroupIdUserInvitesPostRequest(val *GroupsGroupIdUserInvitesPostRequest) *NullableGroupsGroupIdUserInvitesPostRequest {
	return &NullableGroupsGroupIdUserInvitesPostRequest{value: val, isSet: true}
}

func (v NullableGroupsGroupIdUserInvitesPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupsGroupIdUserInvitesPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


