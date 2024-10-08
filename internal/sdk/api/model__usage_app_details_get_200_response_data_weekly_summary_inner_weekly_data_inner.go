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

// checks if the UsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner{}

// UsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner struct for UsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner
type UsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner struct {
	Week string `json:"week"`
	CountAppSaves float32 `json:"count_app_saves"`
	CountAppViews float32 `json:"count_app_views"`
	AdditionalProperties map[string]interface{}
}

type _UsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner UsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner

// NewUsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner instantiates a new UsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner(week string, countAppSaves float32, countAppViews float32) *UsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner {
	this := UsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner{}
	this.Week = week
	this.CountAppSaves = countAppSaves
	this.CountAppViews = countAppViews
	return &this
}

// NewUsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInnerWithDefaults instantiates a new UsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInnerWithDefaults() *UsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner {
	this := UsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner{}
	return &this
}

// GetWeek returns the Week field value
func (o *UsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner) GetWeek() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Week
}

// GetWeekOk returns a tuple with the Week field value
// and a boolean to check if the value has been set.
func (o *UsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner) GetWeekOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Week, true
}

// SetWeek sets field value
func (o *UsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner) SetWeek(v string) {
	o.Week = v
}

// GetCountAppSaves returns the CountAppSaves field value
func (o *UsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner) GetCountAppSaves() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CountAppSaves
}

// GetCountAppSavesOk returns a tuple with the CountAppSaves field value
// and a boolean to check if the value has been set.
func (o *UsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner) GetCountAppSavesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountAppSaves, true
}

// SetCountAppSaves sets field value
func (o *UsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner) SetCountAppSaves(v float32) {
	o.CountAppSaves = v
}

// GetCountAppViews returns the CountAppViews field value
func (o *UsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner) GetCountAppViews() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CountAppViews
}

// GetCountAppViewsOk returns a tuple with the CountAppViews field value
// and a boolean to check if the value has been set.
func (o *UsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner) GetCountAppViewsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountAppViews, true
}

// SetCountAppViews sets field value
func (o *UsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner) SetCountAppViews(v float32) {
	o.CountAppViews = v
}

func (o UsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["week"] = o.Week
	toSerialize["count_app_saves"] = o.CountAppSaves
	toSerialize["count_app_views"] = o.CountAppViews

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"week",
		"count_app_saves",
		"count_app_views",
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

	varUsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner := _UsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner{}

	err = json.Unmarshal(data, &varUsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner)

	if err != nil {
		return err
	}

	*o = UsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner(varUsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "week")
		delete(additionalProperties, "count_app_saves")
		delete(additionalProperties, "count_app_views")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner struct {
	value *UsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner
	isSet bool
}

func (v NullableUsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner) Get() *UsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner {
	return v.value
}

func (v *NullableUsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner) Set(val *UsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner(val *UsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner) *NullableUsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner {
	return &NullableUsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner{value: val, isSet: true}
}

func (v NullableUsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageAppDetailsGet200ResponseDataWeeklySummaryInnerWeeklyDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


