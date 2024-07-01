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

// checks if the UserInvitesPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserInvitesPostRequest{}

// UserInvitesPostRequest Body of request to create a user invite
type UserInvitesPostRequest struct {
	// The email of the invitee
	Email string `json:"email"`
	// Group IDs to invite the user to
	DefaultGroupIds []float32 `json:"defaultGroupIds,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type _UserInvitesPostRequest UserInvitesPostRequest

// NewUserInvitesPostRequest instantiates a new UserInvitesPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInvitesPostRequest(email string) *UserInvitesPostRequest {
	this := UserInvitesPostRequest{}
	this.Email = email
	return &this
}

// NewUserInvitesPostRequestWithDefaults instantiates a new UserInvitesPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInvitesPostRequestWithDefaults() *UserInvitesPostRequest {
	this := UserInvitesPostRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *UserInvitesPostRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserInvitesPostRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserInvitesPostRequest) SetEmail(v string) {
	o.Email = v
}

// GetDefaultGroupIds returns the DefaultGroupIds field value if set, zero value otherwise.
func (o *UserInvitesPostRequest) GetDefaultGroupIds() []float32 {
	if o == nil || IsNil(o.DefaultGroupIds) {
		var ret []float32
		return ret
	}
	return o.DefaultGroupIds
}

// GetDefaultGroupIdsOk returns a tuple with the DefaultGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvitesPostRequest) GetDefaultGroupIdsOk() ([]float32, bool) {
	if o == nil || IsNil(o.DefaultGroupIds) {
		return nil, false
	}
	return o.DefaultGroupIds, true
}

// HasDefaultGroupIds returns a boolean if a field has been set.
func (o *UserInvitesPostRequest) HasDefaultGroupIds() bool {
	if o != nil && !IsNil(o.DefaultGroupIds) {
		return true
	}

	return false
}

// SetDefaultGroupIds gets a reference to the given []float32 and assigns it to the DefaultGroupIds field.
func (o *UserInvitesPostRequest) SetDefaultGroupIds(v []float32) {
	o.DefaultGroupIds = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UserInvitesPostRequest) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvitesPostRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UserInvitesPostRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *UserInvitesPostRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o UserInvitesPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserInvitesPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	if !IsNil(o.DefaultGroupIds) {
		toSerialize["defaultGroupIds"] = o.DefaultGroupIds
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *UserInvitesPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
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

	varUserInvitesPostRequest := _UserInvitesPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserInvitesPostRequest)

	if err != nil {
		return err
	}

	*o = UserInvitesPostRequest(varUserInvitesPostRequest)

	return err
}

type NullableUserInvitesPostRequest struct {
	value *UserInvitesPostRequest
	isSet bool
}

func (v NullableUserInvitesPostRequest) Get() *UserInvitesPostRequest {
	return v.value
}

func (v *NullableUserInvitesPostRequest) Set(val *UserInvitesPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInvitesPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInvitesPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInvitesPostRequest(val *UserInvitesPostRequest) *NullableUserInvitesPostRequest {
	return &NullableUserInvitesPostRequest{value: val, isSet: true}
}

func (v NullableUserInvitesPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInvitesPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


