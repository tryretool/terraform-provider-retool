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

// checks if the UserInvitesUserInviteIdUserAttributesPost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserInvitesUserInviteIdUserAttributesPost200Response{}

// UserInvitesUserInviteIdUserAttributesPost200Response struct for UserInvitesUserInviteIdUserAttributesPost200Response
type UserInvitesUserInviteIdUserAttributesPost200Response struct {
	// API request succeeded
	Success bool `json:"success"`
	Data UserInvitesUserInviteIdUserAttributesPost200ResponseData `json:"data"`
}

type _UserInvitesUserInviteIdUserAttributesPost200Response UserInvitesUserInviteIdUserAttributesPost200Response

// NewUserInvitesUserInviteIdUserAttributesPost200Response instantiates a new UserInvitesUserInviteIdUserAttributesPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInvitesUserInviteIdUserAttributesPost200Response(success bool, data UserInvitesUserInviteIdUserAttributesPost200ResponseData) *UserInvitesUserInviteIdUserAttributesPost200Response {
	this := UserInvitesUserInviteIdUserAttributesPost200Response{}
	this.Success = success
	this.Data = data
	return &this
}

// NewUserInvitesUserInviteIdUserAttributesPost200ResponseWithDefaults instantiates a new UserInvitesUserInviteIdUserAttributesPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInvitesUserInviteIdUserAttributesPost200ResponseWithDefaults() *UserInvitesUserInviteIdUserAttributesPost200Response {
	this := UserInvitesUserInviteIdUserAttributesPost200Response{}
	return &this
}

// GetSuccess returns the Success field value
func (o *UserInvitesUserInviteIdUserAttributesPost200Response) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *UserInvitesUserInviteIdUserAttributesPost200Response) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *UserInvitesUserInviteIdUserAttributesPost200Response) SetSuccess(v bool) {
	o.Success = v
}

// GetData returns the Data field value
func (o *UserInvitesUserInviteIdUserAttributesPost200Response) GetData() UserInvitesUserInviteIdUserAttributesPost200ResponseData {
	if o == nil {
		var ret UserInvitesUserInviteIdUserAttributesPost200ResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UserInvitesUserInviteIdUserAttributesPost200Response) GetDataOk() (*UserInvitesUserInviteIdUserAttributesPost200ResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UserInvitesUserInviteIdUserAttributesPost200Response) SetData(v UserInvitesUserInviteIdUserAttributesPost200ResponseData) {
	o.Data = v
}

func (o UserInvitesUserInviteIdUserAttributesPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserInvitesUserInviteIdUserAttributesPost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *UserInvitesUserInviteIdUserAttributesPost200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"success",
		"data",
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

	varUserInvitesUserInviteIdUserAttributesPost200Response := _UserInvitesUserInviteIdUserAttributesPost200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserInvitesUserInviteIdUserAttributesPost200Response)

	if err != nil {
		return err
	}

	*o = UserInvitesUserInviteIdUserAttributesPost200Response(varUserInvitesUserInviteIdUserAttributesPost200Response)

	return err
}

type NullableUserInvitesUserInviteIdUserAttributesPost200Response struct {
	value *UserInvitesUserInviteIdUserAttributesPost200Response
	isSet bool
}

func (v NullableUserInvitesUserInviteIdUserAttributesPost200Response) Get() *UserInvitesUserInviteIdUserAttributesPost200Response {
	return v.value
}

func (v *NullableUserInvitesUserInviteIdUserAttributesPost200Response) Set(val *UserInvitesUserInviteIdUserAttributesPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInvitesUserInviteIdUserAttributesPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInvitesUserInviteIdUserAttributesPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInvitesUserInviteIdUserAttributesPost200Response(val *UserInvitesUserInviteIdUserAttributesPost200Response) *NullableUserInvitesUserInviteIdUserAttributesPost200Response {
	return &NullableUserInvitesUserInviteIdUserAttributesPost200Response{value: val, isSet: true}
}

func (v NullableUserInvitesUserInviteIdUserAttributesPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInvitesUserInviteIdUserAttributesPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


