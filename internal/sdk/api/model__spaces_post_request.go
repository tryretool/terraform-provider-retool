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

// checks if the SpacesPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpacesPostRequest{}

// SpacesPostRequest struct for SpacesPostRequest
type SpacesPostRequest struct {
	Name string `json:"name"`
	// The domain of the space. On Retool Cloud, specify subdomain of the space instead.
	Domain string `json:"domain"`
	Options *SpacesPostRequestOptions `json:"options,omitempty"`
}

type _SpacesPostRequest SpacesPostRequest

// NewSpacesPostRequest instantiates a new SpacesPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpacesPostRequest(name string, domain string) *SpacesPostRequest {
	this := SpacesPostRequest{}
	this.Name = name
	this.Domain = domain
	return &this
}

// NewSpacesPostRequestWithDefaults instantiates a new SpacesPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpacesPostRequestWithDefaults() *SpacesPostRequest {
	this := SpacesPostRequest{}
	return &this
}

// GetName returns the Name field value
func (o *SpacesPostRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SpacesPostRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SpacesPostRequest) SetName(v string) {
	o.Name = v
}

// GetDomain returns the Domain field value
func (o *SpacesPostRequest) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *SpacesPostRequest) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *SpacesPostRequest) SetDomain(v string) {
	o.Domain = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *SpacesPostRequest) GetOptions() SpacesPostRequestOptions {
	if o == nil || IsNil(o.Options) {
		var ret SpacesPostRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpacesPostRequest) GetOptionsOk() (*SpacesPostRequestOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *SpacesPostRequest) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given SpacesPostRequestOptions and assigns it to the Options field.
func (o *SpacesPostRequest) SetOptions(v SpacesPostRequestOptions) {
	o.Options = &v
}

func (o SpacesPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpacesPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["domain"] = o.Domain
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

func (o *SpacesPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"domain",
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

	varSpacesPostRequest := _SpacesPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSpacesPostRequest)

	if err != nil {
		return err
	}

	*o = SpacesPostRequest(varSpacesPostRequest)

	return err
}

type NullableSpacesPostRequest struct {
	value *SpacesPostRequest
	isSet bool
}

func (v NullableSpacesPostRequest) Get() *SpacesPostRequest {
	return v.value
}

func (v *NullableSpacesPostRequest) Set(val *SpacesPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSpacesPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSpacesPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpacesPostRequest(val *SpacesPostRequest) *NullableSpacesPostRequest {
	return &NullableSpacesPostRequest{value: val, isSet: true}
}

func (v NullableSpacesPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpacesPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


