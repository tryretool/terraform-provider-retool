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

// checks if the GroupsGroupIdGet200ResponseDataUserInvitesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupsGroupIdGet200ResponseDataUserInvitesInner{}

// GroupsGroupIdGet200ResponseDataUserInvitesInner User Invite
type GroupsGroupIdGet200ResponseDataUserInvitesInner struct {
	Id float32 `json:"id"`
	LegacyId float32 `json:"legacy_id"`
	InvitedBy string `json:"invited_by"`
	InvitedEmail string `json:"invited_email"`
	ExpiresAt string `json:"expires_at"`
	ClaimedBy NullableString `json:"claimed_by"`
	ClaimedAt NullableString `json:"claimed_at"`
	UserType NullableString `json:"user_type"`
	Metadata map[string]interface{} `json:"metadata"`
	CreatedAt string `json:"created_at"`
	InviteLink *string `json:"invite_link,omitempty"`
}

type _GroupsGroupIdGet200ResponseDataUserInvitesInner GroupsGroupIdGet200ResponseDataUserInvitesInner

// NewGroupsGroupIdGet200ResponseDataUserInvitesInner instantiates a new GroupsGroupIdGet200ResponseDataUserInvitesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupsGroupIdGet200ResponseDataUserInvitesInner(id float32, legacyId float32, invitedBy string, invitedEmail string, expiresAt string, claimedBy NullableString, claimedAt NullableString, userType NullableString, metadata map[string]interface{}, createdAt string) *GroupsGroupIdGet200ResponseDataUserInvitesInner {
	this := GroupsGroupIdGet200ResponseDataUserInvitesInner{}
	this.Id = id
	this.LegacyId = legacyId
	this.InvitedBy = invitedBy
	this.InvitedEmail = invitedEmail
	this.ExpiresAt = expiresAt
	this.ClaimedBy = claimedBy
	this.ClaimedAt = claimedAt
	this.UserType = userType
	this.Metadata = metadata
	this.CreatedAt = createdAt
	return &this
}

// NewGroupsGroupIdGet200ResponseDataUserInvitesInnerWithDefaults instantiates a new GroupsGroupIdGet200ResponseDataUserInvitesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupsGroupIdGet200ResponseDataUserInvitesInnerWithDefaults() *GroupsGroupIdGet200ResponseDataUserInvitesInner {
	this := GroupsGroupIdGet200ResponseDataUserInvitesInner{}
	return &this
}

// GetId returns the Id field value
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) SetId(v float32) {
	o.Id = v
}

// GetLegacyId returns the LegacyId field value
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetLegacyId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.LegacyId
}

// GetLegacyIdOk returns a tuple with the LegacyId field value
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetLegacyIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LegacyId, true
}

// SetLegacyId sets field value
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) SetLegacyId(v float32) {
	o.LegacyId = v
}

// GetInvitedBy returns the InvitedBy field value
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetInvitedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvitedBy
}

// GetInvitedByOk returns a tuple with the InvitedBy field value
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetInvitedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvitedBy, true
}

// SetInvitedBy sets field value
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) SetInvitedBy(v string) {
	o.InvitedBy = v
}

// GetInvitedEmail returns the InvitedEmail field value
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetInvitedEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvitedEmail
}

// GetInvitedEmailOk returns a tuple with the InvitedEmail field value
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetInvitedEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvitedEmail, true
}

// SetInvitedEmail sets field value
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) SetInvitedEmail(v string) {
	o.InvitedEmail = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetExpiresAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetExpiresAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) SetExpiresAt(v string) {
	o.ExpiresAt = v
}

// GetClaimedBy returns the ClaimedBy field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetClaimedBy() string {
	if o == nil || o.ClaimedBy.Get() == nil {
		var ret string
		return ret
	}

	return *o.ClaimedBy.Get()
}

// GetClaimedByOk returns a tuple with the ClaimedBy field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetClaimedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClaimedBy.Get(), o.ClaimedBy.IsSet()
}

// SetClaimedBy sets field value
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) SetClaimedBy(v string) {
	o.ClaimedBy.Set(&v)
}

// GetClaimedAt returns the ClaimedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetClaimedAt() string {
	if o == nil || o.ClaimedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.ClaimedAt.Get()
}

// GetClaimedAtOk returns a tuple with the ClaimedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetClaimedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClaimedAt.Get(), o.ClaimedAt.IsSet()
}

// SetClaimedAt sets field value
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) SetClaimedAt(v string) {
	o.ClaimedAt.Set(&v)
}

// GetUserType returns the UserType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetUserType() string {
	if o == nil || o.UserType.Get() == nil {
		var ret string
		return ret
	}

	return *o.UserType.Get()
}

// GetUserTypeOk returns a tuple with the UserType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetUserTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserType.Get(), o.UserType.IsSet()
}

// SetUserType sets field value
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) SetUserType(v string) {
	o.UserType.Set(&v)
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetInviteLink returns the InviteLink field value if set, zero value otherwise.
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetInviteLink() string {
	if o == nil || IsNil(o.InviteLink) {
		var ret string
		return ret
	}
	return *o.InviteLink
}

// GetInviteLinkOk returns a tuple with the InviteLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetInviteLinkOk() (*string, bool) {
	if o == nil || IsNil(o.InviteLink) {
		return nil, false
	}
	return o.InviteLink, true
}

// HasInviteLink returns a boolean if a field has been set.
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) HasInviteLink() bool {
	if o != nil && !IsNil(o.InviteLink) {
		return true
	}

	return false
}

// SetInviteLink gets a reference to the given string and assigns it to the InviteLink field.
func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) SetInviteLink(v string) {
	o.InviteLink = &v
}

func (o GroupsGroupIdGet200ResponseDataUserInvitesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupsGroupIdGet200ResponseDataUserInvitesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["legacy_id"] = o.LegacyId
	toSerialize["invited_by"] = o.InvitedBy
	toSerialize["invited_email"] = o.InvitedEmail
	toSerialize["expires_at"] = o.ExpiresAt
	toSerialize["claimed_by"] = o.ClaimedBy.Get()
	toSerialize["claimed_at"] = o.ClaimedAt.Get()
	toSerialize["user_type"] = o.UserType.Get()
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.InviteLink) {
		toSerialize["invite_link"] = o.InviteLink
	}
	return toSerialize, nil
}

func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"legacy_id",
		"invited_by",
		"invited_email",
		"expires_at",
		"claimed_by",
		"claimed_at",
		"user_type",
		"metadata",
		"created_at",
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

	varGroupsGroupIdGet200ResponseDataUserInvitesInner := _GroupsGroupIdGet200ResponseDataUserInvitesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGroupsGroupIdGet200ResponseDataUserInvitesInner)

	if err != nil {
		return err
	}

	*o = GroupsGroupIdGet200ResponseDataUserInvitesInner(varGroupsGroupIdGet200ResponseDataUserInvitesInner)

	return err
}

type NullableGroupsGroupIdGet200ResponseDataUserInvitesInner struct {
	value *GroupsGroupIdGet200ResponseDataUserInvitesInner
	isSet bool
}

func (v NullableGroupsGroupIdGet200ResponseDataUserInvitesInner) Get() *GroupsGroupIdGet200ResponseDataUserInvitesInner {
	return v.value
}

func (v *NullableGroupsGroupIdGet200ResponseDataUserInvitesInner) Set(val *GroupsGroupIdGet200ResponseDataUserInvitesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupsGroupIdGet200ResponseDataUserInvitesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupsGroupIdGet200ResponseDataUserInvitesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupsGroupIdGet200ResponseDataUserInvitesInner(val *GroupsGroupIdGet200ResponseDataUserInvitesInner) *NullableGroupsGroupIdGet200ResponseDataUserInvitesInner {
	return &NullableGroupsGroupIdGet200ResponseDataUserInvitesInner{value: val, isSet: true}
}

func (v NullableGroupsGroupIdGet200ResponseDataUserInvitesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupsGroupIdGet200ResponseDataUserInvitesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


