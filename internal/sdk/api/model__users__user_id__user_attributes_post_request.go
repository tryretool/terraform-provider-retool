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

// checks if the UsersUserIdUserAttributesPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsersUserIdUserAttributesPostRequest{}

// UsersUserIdUserAttributesPostRequest The body of a request to add a user attribute to a user or update an existing user attribute.
type UsersUserIdUserAttributesPostRequest struct {
	// The name of the user attribute (must match an existing attribute exactly)
	Name string `json:"name"`
	// The value of the user attribute
	Value NullableString `json:"value"`
}

type _UsersUserIdUserAttributesPostRequest UsersUserIdUserAttributesPostRequest

// NewUsersUserIdUserAttributesPostRequest instantiates a new UsersUserIdUserAttributesPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersUserIdUserAttributesPostRequest(name string, value NullableString) *UsersUserIdUserAttributesPostRequest {
	this := UsersUserIdUserAttributesPostRequest{}
	this.Name = name
	this.Value = value
	return &this
}

// NewUsersUserIdUserAttributesPostRequestWithDefaults instantiates a new UsersUserIdUserAttributesPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersUserIdUserAttributesPostRequestWithDefaults() *UsersUserIdUserAttributesPostRequest {
	this := UsersUserIdUserAttributesPostRequest{}
	return &this
}

// GetName returns the Name field value
func (o *UsersUserIdUserAttributesPostRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UsersUserIdUserAttributesPostRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UsersUserIdUserAttributesPostRequest) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UsersUserIdUserAttributesPostRequest) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}

	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsersUserIdUserAttributesPostRequest) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// SetValue sets field value
func (o *UsersUserIdUserAttributesPostRequest) SetValue(v string) {
	o.Value.Set(&v)
}

func (o UsersUserIdUserAttributesPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsersUserIdUserAttributesPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value.Get()
	return toSerialize, nil
}

func (o *UsersUserIdUserAttributesPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"value",
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

	varUsersUserIdUserAttributesPostRequest := _UsersUserIdUserAttributesPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsersUserIdUserAttributesPostRequest)

	if err != nil {
		return err
	}

	*o = UsersUserIdUserAttributesPostRequest(varUsersUserIdUserAttributesPostRequest)

	return err
}

type NullableUsersUserIdUserAttributesPostRequest struct {
	value *UsersUserIdUserAttributesPostRequest
	isSet bool
}

func (v NullableUsersUserIdUserAttributesPostRequest) Get() *UsersUserIdUserAttributesPostRequest {
	return v.value
}

func (v *NullableUsersUserIdUserAttributesPostRequest) Set(val *UsersUserIdUserAttributesPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersUserIdUserAttributesPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersUserIdUserAttributesPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersUserIdUserAttributesPostRequest(val *UsersUserIdUserAttributesPostRequest) *NullableUsersUserIdUserAttributesPostRequest {
	return &NullableUsersUserIdUserAttributesPostRequest{value: val, isSet: true}
}

func (v NullableUsersUserIdUserAttributesPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersUserIdUserAttributesPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


