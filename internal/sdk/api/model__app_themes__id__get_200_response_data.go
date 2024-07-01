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

// checks if the AppThemesIdGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppThemesIdGet200ResponseData{}

// AppThemesIdGet200ResponseData The app theme.
type AppThemesIdGet200ResponseData struct {
	Id float32 `json:"id"`
	LegacyId float32 `json:"legacy_id"`
	// The name of the app theme.
	Name string `json:"name"`
	// The theme object.
	Theme map[string]interface{} `json:"theme"`
}

type _AppThemesIdGet200ResponseData AppThemesIdGet200ResponseData

// NewAppThemesIdGet200ResponseData instantiates a new AppThemesIdGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppThemesIdGet200ResponseData(id float32, legacyId float32, name string, theme map[string]interface{}) *AppThemesIdGet200ResponseData {
	this := AppThemesIdGet200ResponseData{}
	this.Id = id
	this.LegacyId = legacyId
	this.Name = name
	this.Theme = theme
	return &this
}

// NewAppThemesIdGet200ResponseDataWithDefaults instantiates a new AppThemesIdGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppThemesIdGet200ResponseDataWithDefaults() *AppThemesIdGet200ResponseData {
	this := AppThemesIdGet200ResponseData{}
	return &this
}

// GetId returns the Id field value
func (o *AppThemesIdGet200ResponseData) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppThemesIdGet200ResponseData) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppThemesIdGet200ResponseData) SetId(v float32) {
	o.Id = v
}

// GetLegacyId returns the LegacyId field value
func (o *AppThemesIdGet200ResponseData) GetLegacyId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.LegacyId
}

// GetLegacyIdOk returns a tuple with the LegacyId field value
// and a boolean to check if the value has been set.
func (o *AppThemesIdGet200ResponseData) GetLegacyIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LegacyId, true
}

// SetLegacyId sets field value
func (o *AppThemesIdGet200ResponseData) SetLegacyId(v float32) {
	o.LegacyId = v
}

// GetName returns the Name field value
func (o *AppThemesIdGet200ResponseData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AppThemesIdGet200ResponseData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AppThemesIdGet200ResponseData) SetName(v string) {
	o.Name = v
}

// GetTheme returns the Theme field value
func (o *AppThemesIdGet200ResponseData) GetTheme() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Theme
}

// GetThemeOk returns a tuple with the Theme field value
// and a boolean to check if the value has been set.
func (o *AppThemesIdGet200ResponseData) GetThemeOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Theme, true
}

// SetTheme sets field value
func (o *AppThemesIdGet200ResponseData) SetTheme(v map[string]interface{}) {
	o.Theme = v
}

func (o AppThemesIdGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppThemesIdGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["legacy_id"] = o.LegacyId
	toSerialize["name"] = o.Name
	toSerialize["theme"] = o.Theme
	return toSerialize, nil
}

func (o *AppThemesIdGet200ResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"legacy_id",
		"name",
		"theme",
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

	varAppThemesIdGet200ResponseData := _AppThemesIdGet200ResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppThemesIdGet200ResponseData)

	if err != nil {
		return err
	}

	*o = AppThemesIdGet200ResponseData(varAppThemesIdGet200ResponseData)

	return err
}

type NullableAppThemesIdGet200ResponseData struct {
	value *AppThemesIdGet200ResponseData
	isSet bool
}

func (v NullableAppThemesIdGet200ResponseData) Get() *AppThemesIdGet200ResponseData {
	return v.value
}

func (v *NullableAppThemesIdGet200ResponseData) Set(val *AppThemesIdGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppThemesIdGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppThemesIdGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppThemesIdGet200ResponseData(val *AppThemesIdGet200ResponseData) *NullableAppThemesIdGet200ResponseData {
	return &NullableAppThemesIdGet200ResponseData{value: val, isSet: true}
}

func (v NullableAppThemesIdGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppThemesIdGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


