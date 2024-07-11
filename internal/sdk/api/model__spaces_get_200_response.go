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

// checks if the SpacesGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpacesGet200Response{}

// SpacesGet200Response struct for SpacesGet200Response
type SpacesGet200Response struct {
	// API request succeeded
	Success bool `json:"success"`
	// An array of requested items
	Data []SpacesGet200ResponseDataInner `json:"data"`
	// Total number of items in the response
	TotalCount float32 `json:"total_count"`
	// A token to retrieve the next page of items in the collection
	NextToken NullableString `json:"next_token"`
	// Whether there are more items in the collection
	HasMore bool `json:"has_more"`
	AdditionalProperties map[string]interface{}
}

type _SpacesGet200Response SpacesGet200Response

// NewSpacesGet200Response instantiates a new SpacesGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpacesGet200Response(success bool, data []SpacesGet200ResponseDataInner, totalCount float32, nextToken NullableString, hasMore bool) *SpacesGet200Response {
	this := SpacesGet200Response{}
	this.Success = success
	this.Data = data
	this.TotalCount = totalCount
	this.NextToken = nextToken
	this.HasMore = hasMore
	return &this
}

// NewSpacesGet200ResponseWithDefaults instantiates a new SpacesGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpacesGet200ResponseWithDefaults() *SpacesGet200Response {
	this := SpacesGet200Response{}
	return &this
}

// GetSuccess returns the Success field value
func (o *SpacesGet200Response) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *SpacesGet200Response) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *SpacesGet200Response) SetSuccess(v bool) {
	o.Success = v
}

// GetData returns the Data field value
func (o *SpacesGet200Response) GetData() []SpacesGet200ResponseDataInner {
	if o == nil {
		var ret []SpacesGet200ResponseDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SpacesGet200Response) GetDataOk() ([]SpacesGet200ResponseDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *SpacesGet200Response) SetData(v []SpacesGet200ResponseDataInner) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value
func (o *SpacesGet200Response) GetTotalCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *SpacesGet200Response) GetTotalCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *SpacesGet200Response) SetTotalCount(v float32) {
	o.TotalCount = v
}

// GetNextToken returns the NextToken field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SpacesGet200Response) GetNextToken() string {
	if o == nil || o.NextToken.Get() == nil {
		var ret string
		return ret
	}

	return *o.NextToken.Get()
}

// GetNextTokenOk returns a tuple with the NextToken field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpacesGet200Response) GetNextTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextToken.Get(), o.NextToken.IsSet()
}

// SetNextToken sets field value
func (o *SpacesGet200Response) SetNextToken(v string) {
	o.NextToken.Set(&v)
}

// GetHasMore returns the HasMore field value
func (o *SpacesGet200Response) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *SpacesGet200Response) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *SpacesGet200Response) SetHasMore(v bool) {
	o.HasMore = v
}

func (o SpacesGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpacesGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	toSerialize["data"] = o.Data
	toSerialize["total_count"] = o.TotalCount
	toSerialize["next_token"] = o.NextToken.Get()
	toSerialize["has_more"] = o.HasMore

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SpacesGet200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"success",
		"data",
		"total_count",
		"next_token",
		"has_more",
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

	varSpacesGet200Response := _SpacesGet200Response{}

	err = json.Unmarshal(data, &varSpacesGet200Response)

	if err != nil {
		return err
	}

	*o = SpacesGet200Response(varSpacesGet200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		delete(additionalProperties, "data")
		delete(additionalProperties, "total_count")
		delete(additionalProperties, "next_token")
		delete(additionalProperties, "has_more")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSpacesGet200Response struct {
	value *SpacesGet200Response
	isSet bool
}

func (v NullableSpacesGet200Response) Get() *SpacesGet200Response {
	return v.value
}

func (v *NullableSpacesGet200Response) Set(val *SpacesGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSpacesGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSpacesGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpacesGet200Response(val *SpacesGet200Response) *NullableSpacesGet200Response {
	return &NullableSpacesGet200Response{value: val, isSet: true}
}

func (v NullableSpacesGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpacesGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


