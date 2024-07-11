/*
Retool API

Go to Settings > API to get started. Once you generate an API token, use bearer token authentication to make requests.

API version: 2.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the CustomComponentLibrariesLibraryIdGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomComponentLibrariesLibraryIdGet200ResponseData{}

// CustomComponentLibrariesLibraryIdGet200ResponseData Custom Component Library
type CustomComponentLibrariesLibraryIdGet200ResponseData struct {
	Id string `json:"id"`
	// How the library will be referred to in Toolscript, and other code based usages.
	Name string `json:"name"`
	Description NullableString `json:"description"`
	// How the library will be referred to in the Retool UI. Should be human readable.
	Label string `json:"label"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _CustomComponentLibrariesLibraryIdGet200ResponseData CustomComponentLibrariesLibraryIdGet200ResponseData

// NewCustomComponentLibrariesLibraryIdGet200ResponseData instantiates a new CustomComponentLibrariesLibraryIdGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomComponentLibrariesLibraryIdGet200ResponseData(id string, name string, description NullableString, label string, createdAt time.Time, updatedAt time.Time) *CustomComponentLibrariesLibraryIdGet200ResponseData {
	this := CustomComponentLibrariesLibraryIdGet200ResponseData{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.Label = label
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewCustomComponentLibrariesLibraryIdGet200ResponseDataWithDefaults instantiates a new CustomComponentLibrariesLibraryIdGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomComponentLibrariesLibraryIdGet200ResponseDataWithDefaults() *CustomComponentLibrariesLibraryIdGet200ResponseData {
	this := CustomComponentLibrariesLibraryIdGet200ResponseData{}
	return &this
}

// GetId returns the Id field value
func (o *CustomComponentLibrariesLibraryIdGet200ResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomComponentLibrariesLibraryIdGet200ResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomComponentLibrariesLibraryIdGet200ResponseData) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *CustomComponentLibrariesLibraryIdGet200ResponseData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomComponentLibrariesLibraryIdGet200ResponseData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomComponentLibrariesLibraryIdGet200ResponseData) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CustomComponentLibrariesLibraryIdGet200ResponseData) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomComponentLibrariesLibraryIdGet200ResponseData) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *CustomComponentLibrariesLibraryIdGet200ResponseData) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetLabel returns the Label field value
func (o *CustomComponentLibrariesLibraryIdGet200ResponseData) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *CustomComponentLibrariesLibraryIdGet200ResponseData) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *CustomComponentLibrariesLibraryIdGet200ResponseData) SetLabel(v string) {
	o.Label = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CustomComponentLibrariesLibraryIdGet200ResponseData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CustomComponentLibrariesLibraryIdGet200ResponseData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CustomComponentLibrariesLibraryIdGet200ResponseData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *CustomComponentLibrariesLibraryIdGet200ResponseData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CustomComponentLibrariesLibraryIdGet200ResponseData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *CustomComponentLibrariesLibraryIdGet200ResponseData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o CustomComponentLibrariesLibraryIdGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomComponentLibrariesLibraryIdGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description.Get()
	toSerialize["label"] = o.Label
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomComponentLibrariesLibraryIdGet200ResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"description",
		"label",
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

	varCustomComponentLibrariesLibraryIdGet200ResponseData := _CustomComponentLibrariesLibraryIdGet200ResponseData{}

	err = json.Unmarshal(data, &varCustomComponentLibrariesLibraryIdGet200ResponseData)

	if err != nil {
		return err
	}

	*o = CustomComponentLibrariesLibraryIdGet200ResponseData(varCustomComponentLibrariesLibraryIdGet200ResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "label")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomComponentLibrariesLibraryIdGet200ResponseData struct {
	value *CustomComponentLibrariesLibraryIdGet200ResponseData
	isSet bool
}

func (v NullableCustomComponentLibrariesLibraryIdGet200ResponseData) Get() *CustomComponentLibrariesLibraryIdGet200ResponseData {
	return v.value
}

func (v *NullableCustomComponentLibrariesLibraryIdGet200ResponseData) Set(val *CustomComponentLibrariesLibraryIdGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomComponentLibrariesLibraryIdGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomComponentLibrariesLibraryIdGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomComponentLibrariesLibraryIdGet200ResponseData(val *CustomComponentLibrariesLibraryIdGet200ResponseData) *NullableCustomComponentLibrariesLibraryIdGet200ResponseData {
	return &NullableCustomComponentLibrariesLibraryIdGet200ResponseData{value: val, isSet: true}
}

func (v NullableCustomComponentLibrariesLibraryIdGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomComponentLibrariesLibraryIdGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


