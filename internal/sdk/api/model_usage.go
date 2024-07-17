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

// checks if the Usage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Usage{}

// Usage struct for Usage
type Usage struct {
	// The number of page saves for all the organizations within the date range
	CountCurrentPageSaves float32 `json:"count_current_page_saves"`
	// The number of page views for all the organizations within the date range
	CountCurrentPageViews float32 `json:"count_current_page_views"`
	// The total number of unique users who have performed any tracked activity in all the organizations within the date range
	CountCurrentUsers float32 `json:"count_current_users"`
	// The number of unique users who have been active over the trailing 30 days from the end date provided
	CountT30Users float32 `json:"count_T30_users"`
	// The percentage growth in page saves compared to the previous cycle of the same length as the provided date range
	PercentGrowthPageSaves float32 `json:"percent_growth_page_saves"`
	// The percentage growth in page views compared to the previous cycle of the same length as the provided date range
	PercentGrowthPageViews float32 `json:"percent_growth_page_views"`
	// The percentage growth in unique users compared to the previous cycle of the same length as the provided date range
	PercentGrowthUsers float32 `json:"percent_growth_users"`
	// The percentage growth in T30 users (trailing 30 days users) compared to the previous cycle of the same length as the provided date range
	PercentGrowthT30Users float32 `json:"percent_growth_T30_users"`
	DailyT30Usage []UsageGet200ResponseDataDailyT30UsageInner `json:"daily_T30_usage"`
}

type _Usage Usage

// NewUsage instantiates a new Usage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsage(countCurrentPageSaves float32, countCurrentPageViews float32, countCurrentUsers float32, countT30Users float32, percentGrowthPageSaves float32, percentGrowthPageViews float32, percentGrowthUsers float32, percentGrowthT30Users float32, dailyT30Usage []UsageGet200ResponseDataDailyT30UsageInner) *Usage {
	this := Usage{}
	this.CountCurrentPageSaves = countCurrentPageSaves
	this.CountCurrentPageViews = countCurrentPageViews
	this.CountCurrentUsers = countCurrentUsers
	this.CountT30Users = countT30Users
	this.PercentGrowthPageSaves = percentGrowthPageSaves
	this.PercentGrowthPageViews = percentGrowthPageViews
	this.PercentGrowthUsers = percentGrowthUsers
	this.PercentGrowthT30Users = percentGrowthT30Users
	this.DailyT30Usage = dailyT30Usage
	return &this
}

// NewUsageWithDefaults instantiates a new Usage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageWithDefaults() *Usage {
	this := Usage{}
	return &this
}

// GetCountCurrentPageSaves returns the CountCurrentPageSaves field value
func (o *Usage) GetCountCurrentPageSaves() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CountCurrentPageSaves
}

// GetCountCurrentPageSavesOk returns a tuple with the CountCurrentPageSaves field value
// and a boolean to check if the value has been set.
func (o *Usage) GetCountCurrentPageSavesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountCurrentPageSaves, true
}

// SetCountCurrentPageSaves sets field value
func (o *Usage) SetCountCurrentPageSaves(v float32) {
	o.CountCurrentPageSaves = v
}

// GetCountCurrentPageViews returns the CountCurrentPageViews field value
func (o *Usage) GetCountCurrentPageViews() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CountCurrentPageViews
}

// GetCountCurrentPageViewsOk returns a tuple with the CountCurrentPageViews field value
// and a boolean to check if the value has been set.
func (o *Usage) GetCountCurrentPageViewsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountCurrentPageViews, true
}

// SetCountCurrentPageViews sets field value
func (o *Usage) SetCountCurrentPageViews(v float32) {
	o.CountCurrentPageViews = v
}

// GetCountCurrentUsers returns the CountCurrentUsers field value
func (o *Usage) GetCountCurrentUsers() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CountCurrentUsers
}

// GetCountCurrentUsersOk returns a tuple with the CountCurrentUsers field value
// and a boolean to check if the value has been set.
func (o *Usage) GetCountCurrentUsersOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountCurrentUsers, true
}

// SetCountCurrentUsers sets field value
func (o *Usage) SetCountCurrentUsers(v float32) {
	o.CountCurrentUsers = v
}

// GetCountT30Users returns the CountT30Users field value
func (o *Usage) GetCountT30Users() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CountT30Users
}

// GetCountT30UsersOk returns a tuple with the CountT30Users field value
// and a boolean to check if the value has been set.
func (o *Usage) GetCountT30UsersOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountT30Users, true
}

// SetCountT30Users sets field value
func (o *Usage) SetCountT30Users(v float32) {
	o.CountT30Users = v
}

// GetPercentGrowthPageSaves returns the PercentGrowthPageSaves field value
func (o *Usage) GetPercentGrowthPageSaves() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PercentGrowthPageSaves
}

// GetPercentGrowthPageSavesOk returns a tuple with the PercentGrowthPageSaves field value
// and a boolean to check if the value has been set.
func (o *Usage) GetPercentGrowthPageSavesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PercentGrowthPageSaves, true
}

// SetPercentGrowthPageSaves sets field value
func (o *Usage) SetPercentGrowthPageSaves(v float32) {
	o.PercentGrowthPageSaves = v
}

// GetPercentGrowthPageViews returns the PercentGrowthPageViews field value
func (o *Usage) GetPercentGrowthPageViews() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PercentGrowthPageViews
}

// GetPercentGrowthPageViewsOk returns a tuple with the PercentGrowthPageViews field value
// and a boolean to check if the value has been set.
func (o *Usage) GetPercentGrowthPageViewsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PercentGrowthPageViews, true
}

// SetPercentGrowthPageViews sets field value
func (o *Usage) SetPercentGrowthPageViews(v float32) {
	o.PercentGrowthPageViews = v
}

// GetPercentGrowthUsers returns the PercentGrowthUsers field value
func (o *Usage) GetPercentGrowthUsers() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PercentGrowthUsers
}

// GetPercentGrowthUsersOk returns a tuple with the PercentGrowthUsers field value
// and a boolean to check if the value has been set.
func (o *Usage) GetPercentGrowthUsersOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PercentGrowthUsers, true
}

// SetPercentGrowthUsers sets field value
func (o *Usage) SetPercentGrowthUsers(v float32) {
	o.PercentGrowthUsers = v
}

// GetPercentGrowthT30Users returns the PercentGrowthT30Users field value
func (o *Usage) GetPercentGrowthT30Users() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PercentGrowthT30Users
}

// GetPercentGrowthT30UsersOk returns a tuple with the PercentGrowthT30Users field value
// and a boolean to check if the value has been set.
func (o *Usage) GetPercentGrowthT30UsersOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PercentGrowthT30Users, true
}

// SetPercentGrowthT30Users sets field value
func (o *Usage) SetPercentGrowthT30Users(v float32) {
	o.PercentGrowthT30Users = v
}

// GetDailyT30Usage returns the DailyT30Usage field value
func (o *Usage) GetDailyT30Usage() []UsageGet200ResponseDataDailyT30UsageInner {
	if o == nil {
		var ret []UsageGet200ResponseDataDailyT30UsageInner
		return ret
	}

	return o.DailyT30Usage
}

// GetDailyT30UsageOk returns a tuple with the DailyT30Usage field value
// and a boolean to check if the value has been set.
func (o *Usage) GetDailyT30UsageOk() ([]UsageGet200ResponseDataDailyT30UsageInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.DailyT30Usage, true
}

// SetDailyT30Usage sets field value
func (o *Usage) SetDailyT30Usage(v []UsageGet200ResponseDataDailyT30UsageInner) {
	o.DailyT30Usage = v
}

func (o Usage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Usage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count_current_page_saves"] = o.CountCurrentPageSaves
	toSerialize["count_current_page_views"] = o.CountCurrentPageViews
	toSerialize["count_current_users"] = o.CountCurrentUsers
	toSerialize["count_T30_users"] = o.CountT30Users
	toSerialize["percent_growth_page_saves"] = o.PercentGrowthPageSaves
	toSerialize["percent_growth_page_views"] = o.PercentGrowthPageViews
	toSerialize["percent_growth_users"] = o.PercentGrowthUsers
	toSerialize["percent_growth_T30_users"] = o.PercentGrowthT30Users
	toSerialize["daily_T30_usage"] = o.DailyT30Usage
	return toSerialize, nil
}

func (o *Usage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count_current_page_saves",
		"count_current_page_views",
		"count_current_users",
		"count_T30_users",
		"percent_growth_page_saves",
		"percent_growth_page_views",
		"percent_growth_users",
		"percent_growth_T30_users",
		"daily_T30_usage",
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

	varUsage := _Usage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsage)

	if err != nil {
		return err
	}

	*o = Usage(varUsage)

	return err
}

type NullableUsage struct {
	value *Usage
	isSet bool
}

func (v NullableUsage) Get() *Usage {
	return v.value
}

func (v *NullableUsage) Set(val *Usage) {
	v.value = val
	v.isSet = true
}

func (v NullableUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsage(val *Usage) *NullableUsage {
	return &NullableUsage{value: val, isSet: true}
}

func (v NullableUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


