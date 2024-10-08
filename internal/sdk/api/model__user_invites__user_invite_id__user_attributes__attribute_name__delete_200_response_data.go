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

// checks if the UserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData{}

// UserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData struct for UserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData
type UserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData struct {
	// The updated user metadata
	Metadata map[string]interface{} `json:"metadata"`
	AdditionalProperties map[string]interface{}
}

type _UserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData UserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData

// NewUserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData instantiates a new UserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData(metadata map[string]interface{}) *UserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData {
	this := UserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData{}
	this.Metadata = metadata
	return &this
}

// NewUserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseDataWithDefaults instantiates a new UserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseDataWithDefaults() *UserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData {
	this := UserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData{}
	return &this
}

// GetMetadata returns the Metadata field value
func (o *UserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *UserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *UserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o UserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metadata"] = o.Metadata

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"metadata",
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

	varUserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData := _UserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData{}

	err = json.Unmarshal(data, &varUserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData)

	if err != nil {
		return err
	}

	*o = UserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData(varUserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData struct {
	value *UserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData
	isSet bool
}

func (v NullableUserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData) Get() *UserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData {
	return v.value
}

func (v *NullableUserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData) Set(val *UserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData(val *UserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData) *NullableUserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData {
	return &NullableUserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData{value: val, isSet: true}
}

func (v NullableUserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInvitesUserInviteIdUserAttributesAttributeNameDelete200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


