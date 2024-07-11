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

// checks if the UserInvitesUserInviteIdUserAttributesPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserInvitesUserInviteIdUserAttributesPost200ResponseData{}

// UserInvitesUserInviteIdUserAttributesPost200ResponseData struct for UserInvitesUserInviteIdUserAttributesPost200ResponseData
type UserInvitesUserInviteIdUserAttributesPost200ResponseData struct {
	// The updated user invite metadata, containing the new attribute value
	Metadata map[string]interface{} `json:"metadata"`
	AdditionalProperties map[string]interface{}
}

type _UserInvitesUserInviteIdUserAttributesPost200ResponseData UserInvitesUserInviteIdUserAttributesPost200ResponseData

// NewUserInvitesUserInviteIdUserAttributesPost200ResponseData instantiates a new UserInvitesUserInviteIdUserAttributesPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInvitesUserInviteIdUserAttributesPost200ResponseData(metadata map[string]interface{}) *UserInvitesUserInviteIdUserAttributesPost200ResponseData {
	this := UserInvitesUserInviteIdUserAttributesPost200ResponseData{}
	this.Metadata = metadata
	return &this
}

// NewUserInvitesUserInviteIdUserAttributesPost200ResponseDataWithDefaults instantiates a new UserInvitesUserInviteIdUserAttributesPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInvitesUserInviteIdUserAttributesPost200ResponseDataWithDefaults() *UserInvitesUserInviteIdUserAttributesPost200ResponseData {
	this := UserInvitesUserInviteIdUserAttributesPost200ResponseData{}
	return &this
}

// GetMetadata returns the Metadata field value
func (o *UserInvitesUserInviteIdUserAttributesPost200ResponseData) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *UserInvitesUserInviteIdUserAttributesPost200ResponseData) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *UserInvitesUserInviteIdUserAttributesPost200ResponseData) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o UserInvitesUserInviteIdUserAttributesPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserInvitesUserInviteIdUserAttributesPost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metadata"] = o.Metadata

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserInvitesUserInviteIdUserAttributesPost200ResponseData) UnmarshalJSON(data []byte) (err error) {
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

	varUserInvitesUserInviteIdUserAttributesPost200ResponseData := _UserInvitesUserInviteIdUserAttributesPost200ResponseData{}

	err = json.Unmarshal(data, &varUserInvitesUserInviteIdUserAttributesPost200ResponseData)

	if err != nil {
		return err
	}

	*o = UserInvitesUserInviteIdUserAttributesPost200ResponseData(varUserInvitesUserInviteIdUserAttributesPost200ResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserInvitesUserInviteIdUserAttributesPost200ResponseData struct {
	value *UserInvitesUserInviteIdUserAttributesPost200ResponseData
	isSet bool
}

func (v NullableUserInvitesUserInviteIdUserAttributesPost200ResponseData) Get() *UserInvitesUserInviteIdUserAttributesPost200ResponseData {
	return v.value
}

func (v *NullableUserInvitesUserInviteIdUserAttributesPost200ResponseData) Set(val *UserInvitesUserInviteIdUserAttributesPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInvitesUserInviteIdUserAttributesPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInvitesUserInviteIdUserAttributesPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInvitesUserInviteIdUserAttributesPost200ResponseData(val *UserInvitesUserInviteIdUserAttributesPost200ResponseData) *NullableUserInvitesUserInviteIdUserAttributesPost200ResponseData {
	return &NullableUserInvitesUserInviteIdUserAttributesPost200ResponseData{value: val, isSet: true}
}

func (v NullableUserInvitesUserInviteIdUserAttributesPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInvitesUserInviteIdUserAttributesPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


