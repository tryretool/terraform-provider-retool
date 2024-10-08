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
	"bytes"
	"fmt"
)

// checks if the UsageOrganizationsGet200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageOrganizationsGet200ResponseDataInner{}

// UsageOrganizationsGet200ResponseDataInner An organization in Retool
type UsageOrganizationsGet200ResponseDataInner struct {
	// The id of the organization
	OrgId string `json:"org_id"`
	// The host of the organization
	Host string `json:"host"`
	LastActive time.Time `json:"last_active"`
}

type _UsageOrganizationsGet200ResponseDataInner UsageOrganizationsGet200ResponseDataInner

// NewUsageOrganizationsGet200ResponseDataInner instantiates a new UsageOrganizationsGet200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageOrganizationsGet200ResponseDataInner(orgId string, host string, lastActive time.Time) *UsageOrganizationsGet200ResponseDataInner {
	this := UsageOrganizationsGet200ResponseDataInner{}
	this.OrgId = orgId
	this.Host = host
	this.LastActive = lastActive
	return &this
}

// NewUsageOrganizationsGet200ResponseDataInnerWithDefaults instantiates a new UsageOrganizationsGet200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageOrganizationsGet200ResponseDataInnerWithDefaults() *UsageOrganizationsGet200ResponseDataInner {
	this := UsageOrganizationsGet200ResponseDataInner{}
	return &this
}

// GetOrgId returns the OrgId field value
func (o *UsageOrganizationsGet200ResponseDataInner) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *UsageOrganizationsGet200ResponseDataInner) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *UsageOrganizationsGet200ResponseDataInner) SetOrgId(v string) {
	o.OrgId = v
}

// GetHost returns the Host field value
func (o *UsageOrganizationsGet200ResponseDataInner) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *UsageOrganizationsGet200ResponseDataInner) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *UsageOrganizationsGet200ResponseDataInner) SetHost(v string) {
	o.Host = v
}

// GetLastActive returns the LastActive field value
func (o *UsageOrganizationsGet200ResponseDataInner) GetLastActive() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastActive
}

// GetLastActiveOk returns a tuple with the LastActive field value
// and a boolean to check if the value has been set.
func (o *UsageOrganizationsGet200ResponseDataInner) GetLastActiveOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastActive, true
}

// SetLastActive sets field value
func (o *UsageOrganizationsGet200ResponseDataInner) SetLastActive(v time.Time) {
	o.LastActive = v
}

func (o UsageOrganizationsGet200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageOrganizationsGet200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["org_id"] = o.OrgId
	toSerialize["host"] = o.Host
	toSerialize["last_active"] = o.LastActive
	return toSerialize, nil
}

func (o *UsageOrganizationsGet200ResponseDataInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"org_id",
		"host",
		"last_active",
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

	varUsageOrganizationsGet200ResponseDataInner := _UsageOrganizationsGet200ResponseDataInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsageOrganizationsGet200ResponseDataInner)

	if err != nil {
		return err
	}

	*o = UsageOrganizationsGet200ResponseDataInner(varUsageOrganizationsGet200ResponseDataInner)

	return err
}

type NullableUsageOrganizationsGet200ResponseDataInner struct {
	value *UsageOrganizationsGet200ResponseDataInner
	isSet bool
}

func (v NullableUsageOrganizationsGet200ResponseDataInner) Get() *UsageOrganizationsGet200ResponseDataInner {
	return v.value
}

func (v *NullableUsageOrganizationsGet200ResponseDataInner) Set(val *UsageOrganizationsGet200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageOrganizationsGet200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageOrganizationsGet200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageOrganizationsGet200ResponseDataInner(val *UsageOrganizationsGet200ResponseDataInner) *NullableUsageOrganizationsGet200ResponseDataInner {
	return &NullableUsageOrganizationsGet200ResponseDataInner{value: val, isSet: true}
}

func (v NullableUsageOrganizationsGet200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageOrganizationsGet200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


