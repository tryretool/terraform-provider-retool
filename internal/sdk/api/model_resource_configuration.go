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

// checks if the ResourceConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceConfiguration{}

// ResourceConfiguration Resource Configuration
type ResourceConfiguration struct {
	// The uuid for the resource configuration.
	Id string `json:"id"`
	Resource ResourcesGet200ResponseDataInner `json:"resource"`
	Environment ResourceConfigurationsGet200ResponseDataInnerEnvironment `json:"environment"`
	Options ResourceConfigurationsGet200ResponseDataInnerOptions `json:"options"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type _ResourceConfiguration ResourceConfiguration

// NewResourceConfiguration instantiates a new ResourceConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceConfiguration(id string, resource ResourcesGet200ResponseDataInner, environment ResourceConfigurationsGet200ResponseDataInnerEnvironment, options ResourceConfigurationsGet200ResponseDataInnerOptions, createdAt string, updatedAt string) *ResourceConfiguration {
	this := ResourceConfiguration{}
	this.Id = id
	this.Resource = resource
	this.Environment = environment
	this.Options = options
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewResourceConfigurationWithDefaults instantiates a new ResourceConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceConfigurationWithDefaults() *ResourceConfiguration {
	this := ResourceConfiguration{}
	return &this
}

// GetId returns the Id field value
func (o *ResourceConfiguration) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResourceConfiguration) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResourceConfiguration) SetId(v string) {
	o.Id = v
}

// GetResource returns the Resource field value
func (o *ResourceConfiguration) GetResource() ResourcesGet200ResponseDataInner {
	if o == nil {
		var ret ResourcesGet200ResponseDataInner
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *ResourceConfiguration) GetResourceOk() (*ResourcesGet200ResponseDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *ResourceConfiguration) SetResource(v ResourcesGet200ResponseDataInner) {
	o.Resource = v
}

// GetEnvironment returns the Environment field value
func (o *ResourceConfiguration) GetEnvironment() ResourceConfigurationsGet200ResponseDataInnerEnvironment {
	if o == nil {
		var ret ResourceConfigurationsGet200ResponseDataInnerEnvironment
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *ResourceConfiguration) GetEnvironmentOk() (*ResourceConfigurationsGet200ResponseDataInnerEnvironment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *ResourceConfiguration) SetEnvironment(v ResourceConfigurationsGet200ResponseDataInnerEnvironment) {
	o.Environment = v
}

// GetOptions returns the Options field value
func (o *ResourceConfiguration) GetOptions() ResourceConfigurationsGet200ResponseDataInnerOptions {
	if o == nil {
		var ret ResourceConfigurationsGet200ResponseDataInnerOptions
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *ResourceConfiguration) GetOptionsOk() (*ResourceConfigurationsGet200ResponseDataInnerOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *ResourceConfiguration) SetOptions(v ResourceConfigurationsGet200ResponseDataInnerOptions) {
	o.Options = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ResourceConfiguration) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ResourceConfiguration) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ResourceConfiguration) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ResourceConfiguration) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ResourceConfiguration) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ResourceConfiguration) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o ResourceConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["resource"] = o.Resource
	toSerialize["environment"] = o.Environment
	toSerialize["options"] = o.Options
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *ResourceConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"resource",
		"environment",
		"options",
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

	varResourceConfiguration := _ResourceConfiguration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResourceConfiguration)

	if err != nil {
		return err
	}

	*o = ResourceConfiguration(varResourceConfiguration)

	return err
}

type NullableResourceConfiguration struct {
	value *ResourceConfiguration
	isSet bool
}

func (v NullableResourceConfiguration) Get() *ResourceConfiguration {
	return v.value
}

func (v *NullableResourceConfiguration) Set(val *ResourceConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceConfiguration(val *ResourceConfiguration) *NullableResourceConfiguration {
	return &NullableResourceConfiguration{value: val, isSet: true}
}

func (v NullableResourceConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


