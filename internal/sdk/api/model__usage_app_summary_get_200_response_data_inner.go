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

// checks if the UsageAppSummaryGet200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageAppSummaryGet200ResponseDataInner{}

// UsageAppSummaryGet200ResponseDataInner struct for UsageAppSummaryGet200ResponseDataInner
type UsageAppSummaryGet200ResponseDataInner struct {
	// The name of the app
	AppName string `json:"app_name"`
	// The organization host
	Host string `json:"host"`
	// The number of views of the app in the time range specified
	CountAppViews float32 `json:"count_app_views"`
	// The number of saves of the app in the time range specified
	CountAppSaves float32 `json:"count_app_saves"`
	// The number of viewers of the app in the time range specified
	CountViewers float32 `json:"count_viewers"`
	// The number of editors of the app in the time range specified
	CountEditors float32 `json:"count_editors"`
}

type _UsageAppSummaryGet200ResponseDataInner UsageAppSummaryGet200ResponseDataInner

// NewUsageAppSummaryGet200ResponseDataInner instantiates a new UsageAppSummaryGet200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageAppSummaryGet200ResponseDataInner(appName string, host string, countAppViews float32, countAppSaves float32, countViewers float32, countEditors float32) *UsageAppSummaryGet200ResponseDataInner {
	this := UsageAppSummaryGet200ResponseDataInner{}
	this.AppName = appName
	this.Host = host
	this.CountAppViews = countAppViews
	this.CountAppSaves = countAppSaves
	this.CountViewers = countViewers
	this.CountEditors = countEditors
	return &this
}

// NewUsageAppSummaryGet200ResponseDataInnerWithDefaults instantiates a new UsageAppSummaryGet200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageAppSummaryGet200ResponseDataInnerWithDefaults() *UsageAppSummaryGet200ResponseDataInner {
	this := UsageAppSummaryGet200ResponseDataInner{}
	return &this
}

// GetAppName returns the AppName field value
func (o *UsageAppSummaryGet200ResponseDataInner) GetAppName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value
// and a boolean to check if the value has been set.
func (o *UsageAppSummaryGet200ResponseDataInner) GetAppNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppName, true
}

// SetAppName sets field value
func (o *UsageAppSummaryGet200ResponseDataInner) SetAppName(v string) {
	o.AppName = v
}

// GetHost returns the Host field value
func (o *UsageAppSummaryGet200ResponseDataInner) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *UsageAppSummaryGet200ResponseDataInner) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *UsageAppSummaryGet200ResponseDataInner) SetHost(v string) {
	o.Host = v
}

// GetCountAppViews returns the CountAppViews field value
func (o *UsageAppSummaryGet200ResponseDataInner) GetCountAppViews() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CountAppViews
}

// GetCountAppViewsOk returns a tuple with the CountAppViews field value
// and a boolean to check if the value has been set.
func (o *UsageAppSummaryGet200ResponseDataInner) GetCountAppViewsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountAppViews, true
}

// SetCountAppViews sets field value
func (o *UsageAppSummaryGet200ResponseDataInner) SetCountAppViews(v float32) {
	o.CountAppViews = v
}

// GetCountAppSaves returns the CountAppSaves field value
func (o *UsageAppSummaryGet200ResponseDataInner) GetCountAppSaves() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CountAppSaves
}

// GetCountAppSavesOk returns a tuple with the CountAppSaves field value
// and a boolean to check if the value has been set.
func (o *UsageAppSummaryGet200ResponseDataInner) GetCountAppSavesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountAppSaves, true
}

// SetCountAppSaves sets field value
func (o *UsageAppSummaryGet200ResponseDataInner) SetCountAppSaves(v float32) {
	o.CountAppSaves = v
}

// GetCountViewers returns the CountViewers field value
func (o *UsageAppSummaryGet200ResponseDataInner) GetCountViewers() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CountViewers
}

// GetCountViewersOk returns a tuple with the CountViewers field value
// and a boolean to check if the value has been set.
func (o *UsageAppSummaryGet200ResponseDataInner) GetCountViewersOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountViewers, true
}

// SetCountViewers sets field value
func (o *UsageAppSummaryGet200ResponseDataInner) SetCountViewers(v float32) {
	o.CountViewers = v
}

// GetCountEditors returns the CountEditors field value
func (o *UsageAppSummaryGet200ResponseDataInner) GetCountEditors() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CountEditors
}

// GetCountEditorsOk returns a tuple with the CountEditors field value
// and a boolean to check if the value has been set.
func (o *UsageAppSummaryGet200ResponseDataInner) GetCountEditorsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountEditors, true
}

// SetCountEditors sets field value
func (o *UsageAppSummaryGet200ResponseDataInner) SetCountEditors(v float32) {
	o.CountEditors = v
}

func (o UsageAppSummaryGet200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageAppSummaryGet200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["app_name"] = o.AppName
	toSerialize["host"] = o.Host
	toSerialize["count_app_views"] = o.CountAppViews
	toSerialize["count_app_saves"] = o.CountAppSaves
	toSerialize["count_viewers"] = o.CountViewers
	toSerialize["count_editors"] = o.CountEditors
	return toSerialize, nil
}

func (o *UsageAppSummaryGet200ResponseDataInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"app_name",
		"host",
		"count_app_views",
		"count_app_saves",
		"count_viewers",
		"count_editors",
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

	varUsageAppSummaryGet200ResponseDataInner := _UsageAppSummaryGet200ResponseDataInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsageAppSummaryGet200ResponseDataInner)

	if err != nil {
		return err
	}

	*o = UsageAppSummaryGet200ResponseDataInner(varUsageAppSummaryGet200ResponseDataInner)

	return err
}

type NullableUsageAppSummaryGet200ResponseDataInner struct {
	value *UsageAppSummaryGet200ResponseDataInner
	isSet bool
}

func (v NullableUsageAppSummaryGet200ResponseDataInner) Get() *UsageAppSummaryGet200ResponseDataInner {
	return v.value
}

func (v *NullableUsageAppSummaryGet200ResponseDataInner) Set(val *UsageAppSummaryGet200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageAppSummaryGet200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageAppSummaryGet200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageAppSummaryGet200ResponseDataInner(val *UsageAppSummaryGet200ResponseDataInner) *NullableUsageAppSummaryGet200ResponseDataInner {
	return &NullableUsageAppSummaryGet200ResponseDataInner{value: val, isSet: true}
}

func (v NullableUsageAppSummaryGet200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageAppSummaryGet200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

