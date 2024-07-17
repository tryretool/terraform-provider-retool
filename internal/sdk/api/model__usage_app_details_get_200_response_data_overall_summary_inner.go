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

// checks if the UsageAppDetailsGet200ResponseDataOverallSummaryInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageAppDetailsGet200ResponseDataOverallSummaryInner{}

// UsageAppDetailsGet200ResponseDataOverallSummaryInner struct for UsageAppDetailsGet200ResponseDataOverallSummaryInner
type UsageAppDetailsGet200ResponseDataOverallSummaryInner struct {
	OrgId string `json:"org_id"`
	AppName string `json:"app_name"`
	CountAppSaves float32 `json:"count_app_saves"`
	CountAppViews float32 `json:"count_app_views"`
	Host string `json:"host"`
}

type _UsageAppDetailsGet200ResponseDataOverallSummaryInner UsageAppDetailsGet200ResponseDataOverallSummaryInner

// NewUsageAppDetailsGet200ResponseDataOverallSummaryInner instantiates a new UsageAppDetailsGet200ResponseDataOverallSummaryInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageAppDetailsGet200ResponseDataOverallSummaryInner(orgId string, appName string, countAppSaves float32, countAppViews float32, host string) *UsageAppDetailsGet200ResponseDataOverallSummaryInner {
	this := UsageAppDetailsGet200ResponseDataOverallSummaryInner{}
	this.OrgId = orgId
	this.AppName = appName
	this.CountAppSaves = countAppSaves
	this.CountAppViews = countAppViews
	this.Host = host
	return &this
}

// NewUsageAppDetailsGet200ResponseDataOverallSummaryInnerWithDefaults instantiates a new UsageAppDetailsGet200ResponseDataOverallSummaryInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageAppDetailsGet200ResponseDataOverallSummaryInnerWithDefaults() *UsageAppDetailsGet200ResponseDataOverallSummaryInner {
	this := UsageAppDetailsGet200ResponseDataOverallSummaryInner{}
	return &this
}

// GetOrgId returns the OrgId field value
func (o *UsageAppDetailsGet200ResponseDataOverallSummaryInner) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *UsageAppDetailsGet200ResponseDataOverallSummaryInner) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *UsageAppDetailsGet200ResponseDataOverallSummaryInner) SetOrgId(v string) {
	o.OrgId = v
}

// GetAppName returns the AppName field value
func (o *UsageAppDetailsGet200ResponseDataOverallSummaryInner) GetAppName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value
// and a boolean to check if the value has been set.
func (o *UsageAppDetailsGet200ResponseDataOverallSummaryInner) GetAppNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppName, true
}

// SetAppName sets field value
func (o *UsageAppDetailsGet200ResponseDataOverallSummaryInner) SetAppName(v string) {
	o.AppName = v
}

// GetCountAppSaves returns the CountAppSaves field value
func (o *UsageAppDetailsGet200ResponseDataOverallSummaryInner) GetCountAppSaves() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CountAppSaves
}

// GetCountAppSavesOk returns a tuple with the CountAppSaves field value
// and a boolean to check if the value has been set.
func (o *UsageAppDetailsGet200ResponseDataOverallSummaryInner) GetCountAppSavesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountAppSaves, true
}

// SetCountAppSaves sets field value
func (o *UsageAppDetailsGet200ResponseDataOverallSummaryInner) SetCountAppSaves(v float32) {
	o.CountAppSaves = v
}

// GetCountAppViews returns the CountAppViews field value
func (o *UsageAppDetailsGet200ResponseDataOverallSummaryInner) GetCountAppViews() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CountAppViews
}

// GetCountAppViewsOk returns a tuple with the CountAppViews field value
// and a boolean to check if the value has been set.
func (o *UsageAppDetailsGet200ResponseDataOverallSummaryInner) GetCountAppViewsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountAppViews, true
}

// SetCountAppViews sets field value
func (o *UsageAppDetailsGet200ResponseDataOverallSummaryInner) SetCountAppViews(v float32) {
	o.CountAppViews = v
}

// GetHost returns the Host field value
func (o *UsageAppDetailsGet200ResponseDataOverallSummaryInner) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *UsageAppDetailsGet200ResponseDataOverallSummaryInner) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *UsageAppDetailsGet200ResponseDataOverallSummaryInner) SetHost(v string) {
	o.Host = v
}

func (o UsageAppDetailsGet200ResponseDataOverallSummaryInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageAppDetailsGet200ResponseDataOverallSummaryInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["org_id"] = o.OrgId
	toSerialize["app_name"] = o.AppName
	toSerialize["count_app_saves"] = o.CountAppSaves
	toSerialize["count_app_views"] = o.CountAppViews
	toSerialize["host"] = o.Host
	return toSerialize, nil
}

func (o *UsageAppDetailsGet200ResponseDataOverallSummaryInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"org_id",
		"app_name",
		"count_app_saves",
		"count_app_views",
		"host",
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

	varUsageAppDetailsGet200ResponseDataOverallSummaryInner := _UsageAppDetailsGet200ResponseDataOverallSummaryInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsageAppDetailsGet200ResponseDataOverallSummaryInner)

	if err != nil {
		return err
	}

	*o = UsageAppDetailsGet200ResponseDataOverallSummaryInner(varUsageAppDetailsGet200ResponseDataOverallSummaryInner)

	return err
}

type NullableUsageAppDetailsGet200ResponseDataOverallSummaryInner struct {
	value *UsageAppDetailsGet200ResponseDataOverallSummaryInner
	isSet bool
}

func (v NullableUsageAppDetailsGet200ResponseDataOverallSummaryInner) Get() *UsageAppDetailsGet200ResponseDataOverallSummaryInner {
	return v.value
}

func (v *NullableUsageAppDetailsGet200ResponseDataOverallSummaryInner) Set(val *UsageAppDetailsGet200ResponseDataOverallSummaryInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageAppDetailsGet200ResponseDataOverallSummaryInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageAppDetailsGet200ResponseDataOverallSummaryInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageAppDetailsGet200ResponseDataOverallSummaryInner(val *UsageAppDetailsGet200ResponseDataOverallSummaryInner) *NullableUsageAppDetailsGet200ResponseDataOverallSummaryInner {
	return &NullableUsageAppDetailsGet200ResponseDataOverallSummaryInner{value: val, isSet: true}
}

func (v NullableUsageAppDetailsGet200ResponseDataOverallSummaryInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageAppDetailsGet200ResponseDataOverallSummaryInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

