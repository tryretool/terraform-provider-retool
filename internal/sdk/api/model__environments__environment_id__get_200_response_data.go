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

// checks if the EnvironmentsEnvironmentIdGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentsEnvironmentIdGet200ResponseData{}

// EnvironmentsEnvironmentIdGet200ResponseData The requested environment
type EnvironmentsEnvironmentIdGet200ResponseData struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description NullableString `json:"description"`
	Color string `json:"color"`
	Default bool `json:"default"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type _EnvironmentsEnvironmentIdGet200ResponseData EnvironmentsEnvironmentIdGet200ResponseData

// NewEnvironmentsEnvironmentIdGet200ResponseData instantiates a new EnvironmentsEnvironmentIdGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentsEnvironmentIdGet200ResponseData(id string, name string, description NullableString, color string, default_ bool, createdAt string, updatedAt string) *EnvironmentsEnvironmentIdGet200ResponseData {
	this := EnvironmentsEnvironmentIdGet200ResponseData{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.Color = color
	this.Default = default_
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewEnvironmentsEnvironmentIdGet200ResponseDataWithDefaults instantiates a new EnvironmentsEnvironmentIdGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentsEnvironmentIdGet200ResponseDataWithDefaults() *EnvironmentsEnvironmentIdGet200ResponseData {
	this := EnvironmentsEnvironmentIdGet200ResponseData{}
	return &this
}

// GetId returns the Id field value
func (o *EnvironmentsEnvironmentIdGet200ResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EnvironmentsEnvironmentIdGet200ResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EnvironmentsEnvironmentIdGet200ResponseData) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *EnvironmentsEnvironmentIdGet200ResponseData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EnvironmentsEnvironmentIdGet200ResponseData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EnvironmentsEnvironmentIdGet200ResponseData) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EnvironmentsEnvironmentIdGet200ResponseData) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentsEnvironmentIdGet200ResponseData) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *EnvironmentsEnvironmentIdGet200ResponseData) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetColor returns the Color field value
func (o *EnvironmentsEnvironmentIdGet200ResponseData) GetColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
func (o *EnvironmentsEnvironmentIdGet200ResponseData) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value
func (o *EnvironmentsEnvironmentIdGet200ResponseData) SetColor(v string) {
	o.Color = v
}

// GetDefault returns the Default field value
func (o *EnvironmentsEnvironmentIdGet200ResponseData) GetDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *EnvironmentsEnvironmentIdGet200ResponseData) GetDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value
func (o *EnvironmentsEnvironmentIdGet200ResponseData) SetDefault(v bool) {
	o.Default = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *EnvironmentsEnvironmentIdGet200ResponseData) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EnvironmentsEnvironmentIdGet200ResponseData) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *EnvironmentsEnvironmentIdGet200ResponseData) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *EnvironmentsEnvironmentIdGet200ResponseData) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *EnvironmentsEnvironmentIdGet200ResponseData) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *EnvironmentsEnvironmentIdGet200ResponseData) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o EnvironmentsEnvironmentIdGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentsEnvironmentIdGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description.Get()
	toSerialize["color"] = o.Color
	toSerialize["default"] = o.Default
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *EnvironmentsEnvironmentIdGet200ResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"description",
		"color",
		"default",
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

	varEnvironmentsEnvironmentIdGet200ResponseData := _EnvironmentsEnvironmentIdGet200ResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEnvironmentsEnvironmentIdGet200ResponseData)

	if err != nil {
		return err
	}

	*o = EnvironmentsEnvironmentIdGet200ResponseData(varEnvironmentsEnvironmentIdGet200ResponseData)

	return err
}

type NullableEnvironmentsEnvironmentIdGet200ResponseData struct {
	value *EnvironmentsEnvironmentIdGet200ResponseData
	isSet bool
}

func (v NullableEnvironmentsEnvironmentIdGet200ResponseData) Get() *EnvironmentsEnvironmentIdGet200ResponseData {
	return v.value
}

func (v *NullableEnvironmentsEnvironmentIdGet200ResponseData) Set(val *EnvironmentsEnvironmentIdGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentsEnvironmentIdGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentsEnvironmentIdGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentsEnvironmentIdGet200ResponseData(val *EnvironmentsEnvironmentIdGet200ResponseData) *NullableEnvironmentsEnvironmentIdGet200ResponseData {
	return &NullableEnvironmentsEnvironmentIdGet200ResponseData{value: val, isSet: true}
}

func (v NullableEnvironmentsEnvironmentIdGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentsEnvironmentIdGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


