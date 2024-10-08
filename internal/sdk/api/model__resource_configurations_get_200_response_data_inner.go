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

// checks if the ResourceConfigurationsGet200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceConfigurationsGet200ResponseDataInner{}

// ResourceConfigurationsGet200ResponseDataInner A list of resource configurations.
type ResourceConfigurationsGet200ResponseDataInner struct {
	// The uuid for the resource configuration.
	Id string `json:"id"`
	Resource ResourcesGet200ResponseDataInner `json:"resource"`
	Environment ResourceConfigurationsGet200ResponseDataInnerEnvironment `json:"environment"`
	Options ResourcesPostRequestOptions `json:"options"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type _ResourceConfigurationsGet200ResponseDataInner ResourceConfigurationsGet200ResponseDataInner

// NewResourceConfigurationsGet200ResponseDataInner instantiates a new ResourceConfigurationsGet200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceConfigurationsGet200ResponseDataInner(id string, resource ResourcesGet200ResponseDataInner, environment ResourceConfigurationsGet200ResponseDataInnerEnvironment, options ResourcesPostRequestOptions, createdAt string, updatedAt string) *ResourceConfigurationsGet200ResponseDataInner {
	this := ResourceConfigurationsGet200ResponseDataInner{}
	this.Id = id
	this.Resource = resource
	this.Environment = environment
	this.Options = options
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewResourceConfigurationsGet200ResponseDataInnerWithDefaults instantiates a new ResourceConfigurationsGet200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceConfigurationsGet200ResponseDataInnerWithDefaults() *ResourceConfigurationsGet200ResponseDataInner {
	this := ResourceConfigurationsGet200ResponseDataInner{}
	return &this
}

// GetId returns the Id field value
func (o *ResourceConfigurationsGet200ResponseDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResourceConfigurationsGet200ResponseDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResourceConfigurationsGet200ResponseDataInner) SetId(v string) {
	o.Id = v
}

// GetResource returns the Resource field value
func (o *ResourceConfigurationsGet200ResponseDataInner) GetResource() ResourcesGet200ResponseDataInner {
	if o == nil {
		var ret ResourcesGet200ResponseDataInner
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *ResourceConfigurationsGet200ResponseDataInner) GetResourceOk() (*ResourcesGet200ResponseDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *ResourceConfigurationsGet200ResponseDataInner) SetResource(v ResourcesGet200ResponseDataInner) {
	o.Resource = v
}

// GetEnvironment returns the Environment field value
func (o *ResourceConfigurationsGet200ResponseDataInner) GetEnvironment() ResourceConfigurationsGet200ResponseDataInnerEnvironment {
	if o == nil {
		var ret ResourceConfigurationsGet200ResponseDataInnerEnvironment
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *ResourceConfigurationsGet200ResponseDataInner) GetEnvironmentOk() (*ResourceConfigurationsGet200ResponseDataInnerEnvironment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *ResourceConfigurationsGet200ResponseDataInner) SetEnvironment(v ResourceConfigurationsGet200ResponseDataInnerEnvironment) {
	o.Environment = v
}

// GetOptions returns the Options field value
func (o *ResourceConfigurationsGet200ResponseDataInner) GetOptions() ResourcesPostRequestOptions {
	if o == nil {
		var ret ResourcesPostRequestOptions
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *ResourceConfigurationsGet200ResponseDataInner) GetOptionsOk() (*ResourcesPostRequestOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *ResourceConfigurationsGet200ResponseDataInner) SetOptions(v ResourcesPostRequestOptions) {
	o.Options = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ResourceConfigurationsGet200ResponseDataInner) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ResourceConfigurationsGet200ResponseDataInner) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ResourceConfigurationsGet200ResponseDataInner) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ResourceConfigurationsGet200ResponseDataInner) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ResourceConfigurationsGet200ResponseDataInner) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ResourceConfigurationsGet200ResponseDataInner) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o ResourceConfigurationsGet200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceConfigurationsGet200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["resource"] = o.Resource
	toSerialize["environment"] = o.Environment
	toSerialize["options"] = o.Options
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *ResourceConfigurationsGet200ResponseDataInner) UnmarshalJSON(data []byte) (err error) {
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

	varResourceConfigurationsGet200ResponseDataInner := _ResourceConfigurationsGet200ResponseDataInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResourceConfigurationsGet200ResponseDataInner)

	if err != nil {
		return err
	}

	*o = ResourceConfigurationsGet200ResponseDataInner(varResourceConfigurationsGet200ResponseDataInner)

	return err
}

type NullableResourceConfigurationsGet200ResponseDataInner struct {
	value *ResourceConfigurationsGet200ResponseDataInner
	isSet bool
}

func (v NullableResourceConfigurationsGet200ResponseDataInner) Get() *ResourceConfigurationsGet200ResponseDataInner {
	return v.value
}

func (v *NullableResourceConfigurationsGet200ResponseDataInner) Set(val *ResourceConfigurationsGet200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceConfigurationsGet200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceConfigurationsGet200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceConfigurationsGet200ResponseDataInner(val *ResourceConfigurationsGet200ResponseDataInner) *NullableResourceConfigurationsGet200ResponseDataInner {
	return &NullableResourceConfigurationsGet200ResponseDataInner{value: val, isSet: true}
}

func (v NullableResourceConfigurationsGet200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceConfigurationsGet200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


