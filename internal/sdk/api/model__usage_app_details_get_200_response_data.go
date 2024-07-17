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

// checks if the UsageAppDetailsGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageAppDetailsGet200ResponseData{}

// UsageAppDetailsGet200ResponseData struct for UsageAppDetailsGet200ResponseData
type UsageAppDetailsGet200ResponseData struct {
	OverallSummary []UsageAppDetailsGet200ResponseDataOverallSummaryInner `json:"overall_summary"`
	WeeklySummary []UsageAppDetailsGet200ResponseDataWeeklySummaryInner `json:"weekly_summary"`
	ViewerSummary []UsageAppDetailsGet200ResponseDataViewerSummaryInner `json:"viewer_summary"`
	EditorSummary []UsageAppDetailsGet200ResponseDataViewerSummaryInner `json:"editor_summary"`
}

type _UsageAppDetailsGet200ResponseData UsageAppDetailsGet200ResponseData

// NewUsageAppDetailsGet200ResponseData instantiates a new UsageAppDetailsGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageAppDetailsGet200ResponseData(overallSummary []UsageAppDetailsGet200ResponseDataOverallSummaryInner, weeklySummary []UsageAppDetailsGet200ResponseDataWeeklySummaryInner, viewerSummary []UsageAppDetailsGet200ResponseDataViewerSummaryInner, editorSummary []UsageAppDetailsGet200ResponseDataViewerSummaryInner) *UsageAppDetailsGet200ResponseData {
	this := UsageAppDetailsGet200ResponseData{}
	this.OverallSummary = overallSummary
	this.WeeklySummary = weeklySummary
	this.ViewerSummary = viewerSummary
	this.EditorSummary = editorSummary
	return &this
}

// NewUsageAppDetailsGet200ResponseDataWithDefaults instantiates a new UsageAppDetailsGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageAppDetailsGet200ResponseDataWithDefaults() *UsageAppDetailsGet200ResponseData {
	this := UsageAppDetailsGet200ResponseData{}
	return &this
}

// GetOverallSummary returns the OverallSummary field value
func (o *UsageAppDetailsGet200ResponseData) GetOverallSummary() []UsageAppDetailsGet200ResponseDataOverallSummaryInner {
	if o == nil {
		var ret []UsageAppDetailsGet200ResponseDataOverallSummaryInner
		return ret
	}

	return o.OverallSummary
}

// GetOverallSummaryOk returns a tuple with the OverallSummary field value
// and a boolean to check if the value has been set.
func (o *UsageAppDetailsGet200ResponseData) GetOverallSummaryOk() ([]UsageAppDetailsGet200ResponseDataOverallSummaryInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.OverallSummary, true
}

// SetOverallSummary sets field value
func (o *UsageAppDetailsGet200ResponseData) SetOverallSummary(v []UsageAppDetailsGet200ResponseDataOverallSummaryInner) {
	o.OverallSummary = v
}

// GetWeeklySummary returns the WeeklySummary field value
func (o *UsageAppDetailsGet200ResponseData) GetWeeklySummary() []UsageAppDetailsGet200ResponseDataWeeklySummaryInner {
	if o == nil {
		var ret []UsageAppDetailsGet200ResponseDataWeeklySummaryInner
		return ret
	}

	return o.WeeklySummary
}

// GetWeeklySummaryOk returns a tuple with the WeeklySummary field value
// and a boolean to check if the value has been set.
func (o *UsageAppDetailsGet200ResponseData) GetWeeklySummaryOk() ([]UsageAppDetailsGet200ResponseDataWeeklySummaryInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.WeeklySummary, true
}

// SetWeeklySummary sets field value
func (o *UsageAppDetailsGet200ResponseData) SetWeeklySummary(v []UsageAppDetailsGet200ResponseDataWeeklySummaryInner) {
	o.WeeklySummary = v
}

// GetViewerSummary returns the ViewerSummary field value
func (o *UsageAppDetailsGet200ResponseData) GetViewerSummary() []UsageAppDetailsGet200ResponseDataViewerSummaryInner {
	if o == nil {
		var ret []UsageAppDetailsGet200ResponseDataViewerSummaryInner
		return ret
	}

	return o.ViewerSummary
}

// GetViewerSummaryOk returns a tuple with the ViewerSummary field value
// and a boolean to check if the value has been set.
func (o *UsageAppDetailsGet200ResponseData) GetViewerSummaryOk() ([]UsageAppDetailsGet200ResponseDataViewerSummaryInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.ViewerSummary, true
}

// SetViewerSummary sets field value
func (o *UsageAppDetailsGet200ResponseData) SetViewerSummary(v []UsageAppDetailsGet200ResponseDataViewerSummaryInner) {
	o.ViewerSummary = v
}

// GetEditorSummary returns the EditorSummary field value
func (o *UsageAppDetailsGet200ResponseData) GetEditorSummary() []UsageAppDetailsGet200ResponseDataViewerSummaryInner {
	if o == nil {
		var ret []UsageAppDetailsGet200ResponseDataViewerSummaryInner
		return ret
	}

	return o.EditorSummary
}

// GetEditorSummaryOk returns a tuple with the EditorSummary field value
// and a boolean to check if the value has been set.
func (o *UsageAppDetailsGet200ResponseData) GetEditorSummaryOk() ([]UsageAppDetailsGet200ResponseDataViewerSummaryInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.EditorSummary, true
}

// SetEditorSummary sets field value
func (o *UsageAppDetailsGet200ResponseData) SetEditorSummary(v []UsageAppDetailsGet200ResponseDataViewerSummaryInner) {
	o.EditorSummary = v
}

func (o UsageAppDetailsGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageAppDetailsGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["overall_summary"] = o.OverallSummary
	toSerialize["weekly_summary"] = o.WeeklySummary
	toSerialize["viewer_summary"] = o.ViewerSummary
	toSerialize["editor_summary"] = o.EditorSummary
	return toSerialize, nil
}

func (o *UsageAppDetailsGet200ResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"overall_summary",
		"weekly_summary",
		"viewer_summary",
		"editor_summary",
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

	varUsageAppDetailsGet200ResponseData := _UsageAppDetailsGet200ResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsageAppDetailsGet200ResponseData)

	if err != nil {
		return err
	}

	*o = UsageAppDetailsGet200ResponseData(varUsageAppDetailsGet200ResponseData)

	return err
}

type NullableUsageAppDetailsGet200ResponseData struct {
	value *UsageAppDetailsGet200ResponseData
	isSet bool
}

func (v NullableUsageAppDetailsGet200ResponseData) Get() *UsageAppDetailsGet200ResponseData {
	return v.value
}

func (v *NullableUsageAppDetailsGet200ResponseData) Set(val *UsageAppDetailsGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageAppDetailsGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageAppDetailsGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageAppDetailsGet200ResponseData(val *UsageAppDetailsGet200ResponseData) *NullableUsageAppDetailsGet200ResponseData {
	return &NullableUsageAppDetailsGet200ResponseData{value: val, isSet: true}
}

func (v NullableUsageAppDetailsGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageAppDetailsGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


