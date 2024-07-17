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

// checks if the GroupsGroupIdPut200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupsGroupIdPut200ResponseData{}

// GroupsGroupIdPut200ResponseData The updated group
type GroupsGroupIdPut200ResponseData struct {
	// The ID of the group. Currently this is the same legacy_id but will change in the future.
	Id NullableFloat32 `json:"id"`
	// The legacy ID of the group.
	LegacyId NullableFloat32 `json:"legacy_id"`
	// The name of the group.
	Name string `json:"name"`
	// The members of the group.
	Members []GroupsGroupIdGet200ResponseDataMembersInner `json:"members"`
	// The universal app access level for the group. This denotes the access level that this group has for all apps.
	UniversalAppAccess string `json:"universal_app_access"`
	// The universal resource access level for the group. This denotes the access level that this group has for all resources.
	UniversalResourceAccess string `json:"universal_resource_access"`
	// The universal workflow access level for the group. This denotes the access level that this group has for all workflows.
	UniversalWorkflowAccess string `json:"universal_workflow_access"`
	// A list of user invites that will be added to the group
	UserInvites []GroupsGroupIdGet200ResponseDataUserInvitesInner `json:"user_invites"`
	// Whether the group has access to the user list
	UserListAccess bool `json:"user_list_access"`
	// Whether the group has access to the audit log
	AuditLogAccess bool `json:"audit_log_access"`
	// Whether the group has access to unpublished releases
	UnpublishedReleaseAccess bool `json:"unpublished_release_access"`
	// Whether the group has access to usage analytics
	UsageAnalyticsAccess bool `json:"usage_analytics_access"`
	// Whether the group has access to account details
	AccountDetailsAccess bool `json:"account_details_access"`
	// The app ID of the landing page
	LandingPageAppId NullableString `json:"landing_page_app_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _GroupsGroupIdPut200ResponseData GroupsGroupIdPut200ResponseData

// NewGroupsGroupIdPut200ResponseData instantiates a new GroupsGroupIdPut200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupsGroupIdPut200ResponseData(id NullableFloat32, legacyId NullableFloat32, name string, members []GroupsGroupIdGet200ResponseDataMembersInner, universalAppAccess string, universalResourceAccess string, universalWorkflowAccess string, userInvites []GroupsGroupIdGet200ResponseDataUserInvitesInner, userListAccess bool, auditLogAccess bool, unpublishedReleaseAccess bool, usageAnalyticsAccess bool, accountDetailsAccess bool, landingPageAppId NullableString, createdAt string, updatedAt string) *GroupsGroupIdPut200ResponseData {
	this := GroupsGroupIdPut200ResponseData{}
	this.Id = id
	this.LegacyId = legacyId
	this.Name = name
	this.Members = members
	this.UniversalAppAccess = universalAppAccess
	this.UniversalResourceAccess = universalResourceAccess
	this.UniversalWorkflowAccess = universalWorkflowAccess
	this.UserInvites = userInvites
	this.UserListAccess = userListAccess
	this.AuditLogAccess = auditLogAccess
	this.UnpublishedReleaseAccess = unpublishedReleaseAccess
	this.UsageAnalyticsAccess = usageAnalyticsAccess
	this.AccountDetailsAccess = accountDetailsAccess
	this.LandingPageAppId = landingPageAppId
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewGroupsGroupIdPut200ResponseDataWithDefaults instantiates a new GroupsGroupIdPut200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupsGroupIdPut200ResponseDataWithDefaults() *GroupsGroupIdPut200ResponseData {
	this := GroupsGroupIdPut200ResponseData{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *GroupsGroupIdPut200ResponseData) GetId() float32 {
	if o == nil || o.Id.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupsGroupIdPut200ResponseData) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *GroupsGroupIdPut200ResponseData) SetId(v float32) {
	o.Id.Set(&v)
}

// GetLegacyId returns the LegacyId field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *GroupsGroupIdPut200ResponseData) GetLegacyId() float32 {
	if o == nil || o.LegacyId.Get() == nil {
		var ret float32
		return ret
	}

	return *o.LegacyId.Get()
}

// GetLegacyIdOk returns a tuple with the LegacyId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupsGroupIdPut200ResponseData) GetLegacyIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LegacyId.Get(), o.LegacyId.IsSet()
}

// SetLegacyId sets field value
func (o *GroupsGroupIdPut200ResponseData) SetLegacyId(v float32) {
	o.LegacyId.Set(&v)
}

// GetName returns the Name field value
func (o *GroupsGroupIdPut200ResponseData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdPut200ResponseData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GroupsGroupIdPut200ResponseData) SetName(v string) {
	o.Name = v
}

// GetMembers returns the Members field value
func (o *GroupsGroupIdPut200ResponseData) GetMembers() []GroupsGroupIdGet200ResponseDataMembersInner {
	if o == nil {
		var ret []GroupsGroupIdGet200ResponseDataMembersInner
		return ret
	}

	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdPut200ResponseData) GetMembersOk() ([]GroupsGroupIdGet200ResponseDataMembersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Members, true
}

// SetMembers sets field value
func (o *GroupsGroupIdPut200ResponseData) SetMembers(v []GroupsGroupIdGet200ResponseDataMembersInner) {
	o.Members = v
}

// GetUniversalAppAccess returns the UniversalAppAccess field value
func (o *GroupsGroupIdPut200ResponseData) GetUniversalAppAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UniversalAppAccess
}

// GetUniversalAppAccessOk returns a tuple with the UniversalAppAccess field value
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdPut200ResponseData) GetUniversalAppAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UniversalAppAccess, true
}

// SetUniversalAppAccess sets field value
func (o *GroupsGroupIdPut200ResponseData) SetUniversalAppAccess(v string) {
	o.UniversalAppAccess = v
}

// GetUniversalResourceAccess returns the UniversalResourceAccess field value
func (o *GroupsGroupIdPut200ResponseData) GetUniversalResourceAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UniversalResourceAccess
}

// GetUniversalResourceAccessOk returns a tuple with the UniversalResourceAccess field value
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdPut200ResponseData) GetUniversalResourceAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UniversalResourceAccess, true
}

// SetUniversalResourceAccess sets field value
func (o *GroupsGroupIdPut200ResponseData) SetUniversalResourceAccess(v string) {
	o.UniversalResourceAccess = v
}

// GetUniversalWorkflowAccess returns the UniversalWorkflowAccess field value
func (o *GroupsGroupIdPut200ResponseData) GetUniversalWorkflowAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UniversalWorkflowAccess
}

// GetUniversalWorkflowAccessOk returns a tuple with the UniversalWorkflowAccess field value
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdPut200ResponseData) GetUniversalWorkflowAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UniversalWorkflowAccess, true
}

// SetUniversalWorkflowAccess sets field value
func (o *GroupsGroupIdPut200ResponseData) SetUniversalWorkflowAccess(v string) {
	o.UniversalWorkflowAccess = v
}

// GetUserInvites returns the UserInvites field value
func (o *GroupsGroupIdPut200ResponseData) GetUserInvites() []GroupsGroupIdGet200ResponseDataUserInvitesInner {
	if o == nil {
		var ret []GroupsGroupIdGet200ResponseDataUserInvitesInner
		return ret
	}

	return o.UserInvites
}

// GetUserInvitesOk returns a tuple with the UserInvites field value
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdPut200ResponseData) GetUserInvitesOk() ([]GroupsGroupIdGet200ResponseDataUserInvitesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserInvites, true
}

// SetUserInvites sets field value
func (o *GroupsGroupIdPut200ResponseData) SetUserInvites(v []GroupsGroupIdGet200ResponseDataUserInvitesInner) {
	o.UserInvites = v
}

// GetUserListAccess returns the UserListAccess field value
func (o *GroupsGroupIdPut200ResponseData) GetUserListAccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UserListAccess
}

// GetUserListAccessOk returns a tuple with the UserListAccess field value
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdPut200ResponseData) GetUserListAccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserListAccess, true
}

// SetUserListAccess sets field value
func (o *GroupsGroupIdPut200ResponseData) SetUserListAccess(v bool) {
	o.UserListAccess = v
}

// GetAuditLogAccess returns the AuditLogAccess field value
func (o *GroupsGroupIdPut200ResponseData) GetAuditLogAccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AuditLogAccess
}

// GetAuditLogAccessOk returns a tuple with the AuditLogAccess field value
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdPut200ResponseData) GetAuditLogAccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuditLogAccess, true
}

// SetAuditLogAccess sets field value
func (o *GroupsGroupIdPut200ResponseData) SetAuditLogAccess(v bool) {
	o.AuditLogAccess = v
}

// GetUnpublishedReleaseAccess returns the UnpublishedReleaseAccess field value
func (o *GroupsGroupIdPut200ResponseData) GetUnpublishedReleaseAccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UnpublishedReleaseAccess
}

// GetUnpublishedReleaseAccessOk returns a tuple with the UnpublishedReleaseAccess field value
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdPut200ResponseData) GetUnpublishedReleaseAccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnpublishedReleaseAccess, true
}

// SetUnpublishedReleaseAccess sets field value
func (o *GroupsGroupIdPut200ResponseData) SetUnpublishedReleaseAccess(v bool) {
	o.UnpublishedReleaseAccess = v
}

// GetUsageAnalyticsAccess returns the UsageAnalyticsAccess field value
func (o *GroupsGroupIdPut200ResponseData) GetUsageAnalyticsAccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UsageAnalyticsAccess
}

// GetUsageAnalyticsAccessOk returns a tuple with the UsageAnalyticsAccess field value
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdPut200ResponseData) GetUsageAnalyticsAccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsageAnalyticsAccess, true
}

// SetUsageAnalyticsAccess sets field value
func (o *GroupsGroupIdPut200ResponseData) SetUsageAnalyticsAccess(v bool) {
	o.UsageAnalyticsAccess = v
}

// GetAccountDetailsAccess returns the AccountDetailsAccess field value
func (o *GroupsGroupIdPut200ResponseData) GetAccountDetailsAccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AccountDetailsAccess
}

// GetAccountDetailsAccessOk returns a tuple with the AccountDetailsAccess field value
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdPut200ResponseData) GetAccountDetailsAccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountDetailsAccess, true
}

// SetAccountDetailsAccess sets field value
func (o *GroupsGroupIdPut200ResponseData) SetAccountDetailsAccess(v bool) {
	o.AccountDetailsAccess = v
}

// GetLandingPageAppId returns the LandingPageAppId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GroupsGroupIdPut200ResponseData) GetLandingPageAppId() string {
	if o == nil || o.LandingPageAppId.Get() == nil {
		var ret string
		return ret
	}

	return *o.LandingPageAppId.Get()
}

// GetLandingPageAppIdOk returns a tuple with the LandingPageAppId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupsGroupIdPut200ResponseData) GetLandingPageAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LandingPageAppId.Get(), o.LandingPageAppId.IsSet()
}

// SetLandingPageAppId sets field value
func (o *GroupsGroupIdPut200ResponseData) SetLandingPageAppId(v string) {
	o.LandingPageAppId.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *GroupsGroupIdPut200ResponseData) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdPut200ResponseData) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *GroupsGroupIdPut200ResponseData) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *GroupsGroupIdPut200ResponseData) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdPut200ResponseData) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *GroupsGroupIdPut200ResponseData) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o GroupsGroupIdPut200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupsGroupIdPut200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	toSerialize["legacy_id"] = o.LegacyId.Get()
	toSerialize["name"] = o.Name
	toSerialize["members"] = o.Members
	toSerialize["universal_app_access"] = o.UniversalAppAccess
	toSerialize["universal_resource_access"] = o.UniversalResourceAccess
	toSerialize["universal_workflow_access"] = o.UniversalWorkflowAccess
	toSerialize["user_invites"] = o.UserInvites
	toSerialize["user_list_access"] = o.UserListAccess
	toSerialize["audit_log_access"] = o.AuditLogAccess
	toSerialize["unpublished_release_access"] = o.UnpublishedReleaseAccess
	toSerialize["usage_analytics_access"] = o.UsageAnalyticsAccess
	toSerialize["account_details_access"] = o.AccountDetailsAccess
	toSerialize["landing_page_app_id"] = o.LandingPageAppId.Get()
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupsGroupIdPut200ResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"legacy_id",
		"name",
		"members",
		"universal_app_access",
		"universal_resource_access",
		"universal_workflow_access",
		"user_invites",
		"user_list_access",
		"audit_log_access",
		"unpublished_release_access",
		"usage_analytics_access",
		"account_details_access",
		"landing_page_app_id",
		"created_at",
		"updated_at",
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

	varGroupsGroupIdPut200ResponseData := _GroupsGroupIdPut200ResponseData{}

	err = json.Unmarshal(data, &varGroupsGroupIdPut200ResponseData)

	if err != nil {
		return err
	}

	*o = GroupsGroupIdPut200ResponseData(varGroupsGroupIdPut200ResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "legacy_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "members")
		delete(additionalProperties, "universal_app_access")
		delete(additionalProperties, "universal_resource_access")
		delete(additionalProperties, "universal_workflow_access")
		delete(additionalProperties, "user_invites")
		delete(additionalProperties, "user_list_access")
		delete(additionalProperties, "audit_log_access")
		delete(additionalProperties, "unpublished_release_access")
		delete(additionalProperties, "usage_analytics_access")
		delete(additionalProperties, "account_details_access")
		delete(additionalProperties, "landing_page_app_id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupsGroupIdPut200ResponseData struct {
	value *GroupsGroupIdPut200ResponseData
	isSet bool
}

func (v NullableGroupsGroupIdPut200ResponseData) Get() *GroupsGroupIdPut200ResponseData {
	return v.value
}

func (v *NullableGroupsGroupIdPut200ResponseData) Set(val *GroupsGroupIdPut200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupsGroupIdPut200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupsGroupIdPut200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupsGroupIdPut200ResponseData(val *GroupsGroupIdPut200ResponseData) *NullableGroupsGroupIdPut200ResponseData {
	return &NullableGroupsGroupIdPut200ResponseData{value: val, isSet: true}
}

func (v NullableGroupsGroupIdPut200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupsGroupIdPut200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


