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

// checks if the ResourcesPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourcesPostRequest{}

// ResourcesPostRequest struct for ResourcesPostRequest
type ResourcesPostRequest struct {
	// The type of resource.
	Type string `json:"type"`
	DisplayName string `json:"display_name"`
	Options ResourcesPostRequestOptions `json:"options"`
	AdditionalProperties map[string]interface{}
}

type _ResourcesPostRequest ResourcesPostRequest

// NewResourcesPostRequest instantiates a new ResourcesPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourcesPostRequest(type_ string, displayName string, options ResourcesPostRequestOptions) *ResourcesPostRequest {
	this := ResourcesPostRequest{}
	this.Type = type_
	this.DisplayName = displayName
	this.Options = options
	return &this
}

// NewResourcesPostRequestWithDefaults instantiates a new ResourcesPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourcesPostRequestWithDefaults() *ResourcesPostRequest {
	this := ResourcesPostRequest{}
	return &this
}

// GetType returns the Type field value
func (o *ResourcesPostRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ResourcesPostRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ResourcesPostRequest) SetType(v string) {
	o.Type = v
}

// GetDisplayName returns the DisplayName field value
func (o *ResourcesPostRequest) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *ResourcesPostRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *ResourcesPostRequest) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetOptions returns the Options field value
func (o *ResourcesPostRequest) GetOptions() ResourcesPostRequestOptions {
	if o == nil {
		var ret ResourcesPostRequestOptions
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *ResourcesPostRequest) GetOptionsOk() (*ResourcesPostRequestOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *ResourcesPostRequest) SetOptions(v ResourcesPostRequestOptions) {
	o.Options = v
}

func (o ResourcesPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourcesPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["display_name"] = o.DisplayName
	toSerialize["options"] = o.Options

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourcesPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"display_name",
		"options",
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

	varResourcesPostRequest := _ResourcesPostRequest{}

	err = json.Unmarshal(data, &varResourcesPostRequest)

	if err != nil {
		return err
	}

	*o = ResourcesPostRequest(varResourcesPostRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "display_name")
		delete(additionalProperties, "options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourcesPostRequest struct {
	value *ResourcesPostRequest
	isSet bool
}

func (v NullableResourcesPostRequest) Get() *ResourcesPostRequest {
	return v.value
}

func (v *NullableResourcesPostRequest) Set(val *ResourcesPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcesPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcesPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcesPostRequest(val *ResourcesPostRequest) *NullableResourcesPostRequest {
	return &NullableResourcesPostRequest{value: val, isSet: true}
}

func (v NullableResourcesPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcesPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

