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

// checks if the UsageUserSummaryGet200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageUserSummaryGet200ResponseDataInner{}

// UsageUserSummaryGet200ResponseDataInner struct for UsageUserSummaryGet200ResponseDataInner
type UsageUserSummaryGet200ResponseDataInner struct {
	// The id of the organization to which this user belongs
	OrgId string `json:"org_id"`
	// The id of the user
	UserId string `json:"user_id"`
	// The email of the user
	Email string `json:"email"`
	// The host of the organization to which this user belongs
	Host string `json:"host"`
	// The number of times the user viewed an app in the time range specified
	CountAppViews float32 `json:"count_app_views"`
	// The number of times the user edited an app in the time range specified
	CountAppSaves float32 `json:"count_app_saves"`
	// The number of unique apps edited in the time range specified
	CountUniqueApps float32 `json:"count_unique_apps"`
	LastActive time.Time `json:"last_active"`
}

type _UsageUserSummaryGet200ResponseDataInner UsageUserSummaryGet200ResponseDataInner

// NewUsageUserSummaryGet200ResponseDataInner instantiates a new UsageUserSummaryGet200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageUserSummaryGet200ResponseDataInner(orgId string, userId string, email string, host string, countAppViews float32, countAppSaves float32, countUniqueApps float32, lastActive time.Time) *UsageUserSummaryGet200ResponseDataInner {
	this := UsageUserSummaryGet200ResponseDataInner{}
	this.OrgId = orgId
	this.UserId = userId
	this.Email = email
	this.Host = host
	this.CountAppViews = countAppViews
	this.CountAppSaves = countAppSaves
	this.CountUniqueApps = countUniqueApps
	this.LastActive = lastActive
	return &this
}

// NewUsageUserSummaryGet200ResponseDataInnerWithDefaults instantiates a new UsageUserSummaryGet200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageUserSummaryGet200ResponseDataInnerWithDefaults() *UsageUserSummaryGet200ResponseDataInner {
	this := UsageUserSummaryGet200ResponseDataInner{}
	return &this
}

// GetOrgId returns the OrgId field value
func (o *UsageUserSummaryGet200ResponseDataInner) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *UsageUserSummaryGet200ResponseDataInner) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *UsageUserSummaryGet200ResponseDataInner) SetOrgId(v string) {
	o.OrgId = v
}

// GetUserId returns the UserId field value
func (o *UsageUserSummaryGet200ResponseDataInner) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UsageUserSummaryGet200ResponseDataInner) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UsageUserSummaryGet200ResponseDataInner) SetUserId(v string) {
	o.UserId = v
}

// GetEmail returns the Email field value
func (o *UsageUserSummaryGet200ResponseDataInner) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UsageUserSummaryGet200ResponseDataInner) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UsageUserSummaryGet200ResponseDataInner) SetEmail(v string) {
	o.Email = v
}

// GetHost returns the Host field value
func (o *UsageUserSummaryGet200ResponseDataInner) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *UsageUserSummaryGet200ResponseDataInner) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *UsageUserSummaryGet200ResponseDataInner) SetHost(v string) {
	o.Host = v
}

// GetCountAppViews returns the CountAppViews field value
func (o *UsageUserSummaryGet200ResponseDataInner) GetCountAppViews() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CountAppViews
}

// GetCountAppViewsOk returns a tuple with the CountAppViews field value
// and a boolean to check if the value has been set.
func (o *UsageUserSummaryGet200ResponseDataInner) GetCountAppViewsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountAppViews, true
}

// SetCountAppViews sets field value
func (o *UsageUserSummaryGet200ResponseDataInner) SetCountAppViews(v float32) {
	o.CountAppViews = v
}

// GetCountAppSaves returns the CountAppSaves field value
func (o *UsageUserSummaryGet200ResponseDataInner) GetCountAppSaves() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CountAppSaves
}

// GetCountAppSavesOk returns a tuple with the CountAppSaves field value
// and a boolean to check if the value has been set.
func (o *UsageUserSummaryGet200ResponseDataInner) GetCountAppSavesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountAppSaves, true
}

// SetCountAppSaves sets field value
func (o *UsageUserSummaryGet200ResponseDataInner) SetCountAppSaves(v float32) {
	o.CountAppSaves = v
}

// GetCountUniqueApps returns the CountUniqueApps field value
func (o *UsageUserSummaryGet200ResponseDataInner) GetCountUniqueApps() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CountUniqueApps
}

// GetCountUniqueAppsOk returns a tuple with the CountUniqueApps field value
// and a boolean to check if the value has been set.
func (o *UsageUserSummaryGet200ResponseDataInner) GetCountUniqueAppsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountUniqueApps, true
}

// SetCountUniqueApps sets field value
func (o *UsageUserSummaryGet200ResponseDataInner) SetCountUniqueApps(v float32) {
	o.CountUniqueApps = v
}

// GetLastActive returns the LastActive field value
func (o *UsageUserSummaryGet200ResponseDataInner) GetLastActive() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastActive
}

// GetLastActiveOk returns a tuple with the LastActive field value
// and a boolean to check if the value has been set.
func (o *UsageUserSummaryGet200ResponseDataInner) GetLastActiveOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastActive, true
}

// SetLastActive sets field value
func (o *UsageUserSummaryGet200ResponseDataInner) SetLastActive(v time.Time) {
	o.LastActive = v
}

func (o UsageUserSummaryGet200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageUserSummaryGet200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["org_id"] = o.OrgId
	toSerialize["user_id"] = o.UserId
	toSerialize["email"] = o.Email
	toSerialize["host"] = o.Host
	toSerialize["count_app_views"] = o.CountAppViews
	toSerialize["count_app_saves"] = o.CountAppSaves
	toSerialize["count_unique_apps"] = o.CountUniqueApps
	toSerialize["last_active"] = o.LastActive
	return toSerialize, nil
}

func (o *UsageUserSummaryGet200ResponseDataInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"org_id",
		"user_id",
		"email",
		"host",
		"count_app_views",
		"count_app_saves",
		"count_unique_apps",
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

	varUsageUserSummaryGet200ResponseDataInner := _UsageUserSummaryGet200ResponseDataInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsageUserSummaryGet200ResponseDataInner)

	if err != nil {
		return err
	}

	*o = UsageUserSummaryGet200ResponseDataInner(varUsageUserSummaryGet200ResponseDataInner)

	return err
}

type NullableUsageUserSummaryGet200ResponseDataInner struct {
	value *UsageUserSummaryGet200ResponseDataInner
	isSet bool
}

func (v NullableUsageUserSummaryGet200ResponseDataInner) Get() *UsageUserSummaryGet200ResponseDataInner {
	return v.value
}

func (v *NullableUsageUserSummaryGet200ResponseDataInner) Set(val *UsageUserSummaryGet200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageUserSummaryGet200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageUserSummaryGet200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageUserSummaryGet200ResponseDataInner(val *UsageUserSummaryGet200ResponseDataInner) *NullableUsageUserSummaryGet200ResponseDataInner {
	return &NullableUsageUserSummaryGet200ResponseDataInner{value: val, isSet: true}
}

func (v NullableUsageUserSummaryGet200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageUserSummaryGet200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

