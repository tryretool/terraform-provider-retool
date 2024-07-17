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

// checks if the ResourceConfigurationsConfigurationIdGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceConfigurationsConfigurationIdGet200Response{}

// ResourceConfigurationsConfigurationIdGet200Response Resource Configuration
type ResourceConfigurationsConfigurationIdGet200Response struct {
	// The uuid for the resource configuration.
	Id string `json:"id"`
	Resource ResourcesGet200ResponseDataInner `json:"resource"`
	Environment ResourceConfigurationsGet200ResponseDataInnerEnvironment `json:"environment"`
	Options ResourcesPostRequestOptions `json:"options"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type _ResourceConfigurationsConfigurationIdGet200Response ResourceConfigurationsConfigurationIdGet200Response

// NewResourceConfigurationsConfigurationIdGet200Response instantiates a new ResourceConfigurationsConfigurationIdGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceConfigurationsConfigurationIdGet200Response(id string, resource ResourcesGet200ResponseDataInner, environment ResourceConfigurationsGet200ResponseDataInnerEnvironment, options ResourcesPostRequestOptions, createdAt string, updatedAt string) *ResourceConfigurationsConfigurationIdGet200Response {
	this := ResourceConfigurationsConfigurationIdGet200Response{}
	this.Id = id
	this.Resource = resource
	this.Environment = environment
	this.Options = options
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewResourceConfigurationsConfigurationIdGet200ResponseWithDefaults instantiates a new ResourceConfigurationsConfigurationIdGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceConfigurationsConfigurationIdGet200ResponseWithDefaults() *ResourceConfigurationsConfigurationIdGet200Response {
	this := ResourceConfigurationsConfigurationIdGet200Response{}
	return &this
}

// GetId returns the Id field value
func (o *ResourceConfigurationsConfigurationIdGet200Response) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResourceConfigurationsConfigurationIdGet200Response) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResourceConfigurationsConfigurationIdGet200Response) SetId(v string) {
	o.Id = v
}

// GetResource returns the Resource field value
func (o *ResourceConfigurationsConfigurationIdGet200Response) GetResource() ResourcesGet200ResponseDataInner {
	if o == nil {
		var ret ResourcesGet200ResponseDataInner
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *ResourceConfigurationsConfigurationIdGet200Response) GetResourceOk() (*ResourcesGet200ResponseDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *ResourceConfigurationsConfigurationIdGet200Response) SetResource(v ResourcesGet200ResponseDataInner) {
	o.Resource = v
}

// GetEnvironment returns the Environment field value
func (o *ResourceConfigurationsConfigurationIdGet200Response) GetEnvironment() ResourceConfigurationsGet200ResponseDataInnerEnvironment {
	if o == nil {
		var ret ResourceConfigurationsGet200ResponseDataInnerEnvironment
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *ResourceConfigurationsConfigurationIdGet200Response) GetEnvironmentOk() (*ResourceConfigurationsGet200ResponseDataInnerEnvironment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *ResourceConfigurationsConfigurationIdGet200Response) SetEnvironment(v ResourceConfigurationsGet200ResponseDataInnerEnvironment) {
	o.Environment = v
}

// GetOptions returns the Options field value
func (o *ResourceConfigurationsConfigurationIdGet200Response) GetOptions() ResourcesPostRequestOptions {
	if o == nil {
		var ret ResourcesPostRequestOptions
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *ResourceConfigurationsConfigurationIdGet200Response) GetOptionsOk() (*ResourcesPostRequestOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *ResourceConfigurationsConfigurationIdGet200Response) SetOptions(v ResourcesPostRequestOptions) {
	o.Options = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ResourceConfigurationsConfigurationIdGet200Response) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ResourceConfigurationsConfigurationIdGet200Response) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ResourceConfigurationsConfigurationIdGet200Response) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ResourceConfigurationsConfigurationIdGet200Response) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ResourceConfigurationsConfigurationIdGet200Response) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ResourceConfigurationsConfigurationIdGet200Response) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o ResourceConfigurationsConfigurationIdGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceConfigurationsConfigurationIdGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["resource"] = o.Resource
	toSerialize["environment"] = o.Environment
	toSerialize["options"] = o.Options
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *ResourceConfigurationsConfigurationIdGet200Response) UnmarshalJSON(data []byte) (err error) {
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

	varResourceConfigurationsConfigurationIdGet200Response := _ResourceConfigurationsConfigurationIdGet200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResourceConfigurationsConfigurationIdGet200Response)

	if err != nil {
		return err
	}

	*o = ResourceConfigurationsConfigurationIdGet200Response(varResourceConfigurationsConfigurationIdGet200Response)

	return err
}

type NullableResourceConfigurationsConfigurationIdGet200Response struct {
	value *ResourceConfigurationsConfigurationIdGet200Response
	isSet bool
}

func (v NullableResourceConfigurationsConfigurationIdGet200Response) Get() *ResourceConfigurationsConfigurationIdGet200Response {
	return v.value
}

func (v *NullableResourceConfigurationsConfigurationIdGet200Response) Set(val *ResourceConfigurationsConfigurationIdGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceConfigurationsConfigurationIdGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceConfigurationsConfigurationIdGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceConfigurationsConfigurationIdGet200Response(val *ResourceConfigurationsConfigurationIdGet200Response) *NullableResourceConfigurationsConfigurationIdGet200Response {
	return &NullableResourceConfigurationsConfigurationIdGet200Response{value: val, isSet: true}
}

func (v NullableResourceConfigurationsConfigurationIdGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceConfigurationsConfigurationIdGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


