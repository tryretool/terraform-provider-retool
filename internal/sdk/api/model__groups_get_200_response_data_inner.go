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

// checks if the GroupsGet200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupsGet200ResponseDataInner{}

// GroupsGet200ResponseDataInner This object represents a permission group. A permission group is a list of users by which to restrict access. You add users to one or more groups and then configure the group's level of access. This applies the same permissions to all group members. See https://docs.retool.com/org-users/concepts/permission-groups or more information.
type GroupsGet200ResponseDataInner struct {
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
	// Level of access that the group has to the Query Library.
	UniversalQueryLibraryAccess string `json:"universal_query_library_access"`
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

type _GroupsGet200ResponseDataInner GroupsGet200ResponseDataInner

// NewGroupsGet200ResponseDataInner instantiates a new GroupsGet200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupsGet200ResponseDataInner(id NullableFloat32, legacyId NullableFloat32, name string, members []GroupsGroupIdGet200ResponseDataMembersInner, universalAppAccess string, universalResourceAccess string, universalWorkflowAccess string, universalQueryLibraryAccess string, userInvites []GroupsGroupIdGet200ResponseDataUserInvitesInner, userListAccess bool, auditLogAccess bool, unpublishedReleaseAccess bool, usageAnalyticsAccess bool, accountDetailsAccess bool, landingPageAppId NullableString, createdAt string, updatedAt string) *GroupsGet200ResponseDataInner {
	this := GroupsGet200ResponseDataInner{}
	this.Id = id
	this.LegacyId = legacyId
	this.Name = name
	this.Members = members
	this.UniversalAppAccess = universalAppAccess
	this.UniversalResourceAccess = universalResourceAccess
	this.UniversalWorkflowAccess = universalWorkflowAccess
	this.UniversalQueryLibraryAccess = universalQueryLibraryAccess
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

// NewGroupsGet200ResponseDataInnerWithDefaults instantiates a new GroupsGet200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupsGet200ResponseDataInnerWithDefaults() *GroupsGet200ResponseDataInner {
	this := GroupsGet200ResponseDataInner{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *GroupsGet200ResponseDataInner) GetId() float32 {
	if o == nil || o.Id.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupsGet200ResponseDataInner) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *GroupsGet200ResponseDataInner) SetId(v float32) {
	o.Id.Set(&v)
}

// GetLegacyId returns the LegacyId field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *GroupsGet200ResponseDataInner) GetLegacyId() float32 {
	if o == nil || o.LegacyId.Get() == nil {
		var ret float32
		return ret
	}

	return *o.LegacyId.Get()
}

// GetLegacyIdOk returns a tuple with the LegacyId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupsGet200ResponseDataInner) GetLegacyIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LegacyId.Get(), o.LegacyId.IsSet()
}

// SetLegacyId sets field value
func (o *GroupsGet200ResponseDataInner) SetLegacyId(v float32) {
	o.LegacyId.Set(&v)
}

// GetName returns the Name field value
func (o *GroupsGet200ResponseDataInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GroupsGet200ResponseDataInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GroupsGet200ResponseDataInner) SetName(v string) {
	o.Name = v
}

// GetMembers returns the Members field value
func (o *GroupsGet200ResponseDataInner) GetMembers() []GroupsGroupIdGet200ResponseDataMembersInner {
	if o == nil {
		var ret []GroupsGroupIdGet200ResponseDataMembersInner
		return ret
	}

	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *GroupsGet200ResponseDataInner) GetMembersOk() ([]GroupsGroupIdGet200ResponseDataMembersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Members, true
}

// SetMembers sets field value
func (o *GroupsGet200ResponseDataInner) SetMembers(v []GroupsGroupIdGet200ResponseDataMembersInner) {
	o.Members = v
}

// GetUniversalAppAccess returns the UniversalAppAccess field value
func (o *GroupsGet200ResponseDataInner) GetUniversalAppAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UniversalAppAccess
}

// GetUniversalAppAccessOk returns a tuple with the UniversalAppAccess field value
// and a boolean to check if the value has been set.
func (o *GroupsGet200ResponseDataInner) GetUniversalAppAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UniversalAppAccess, true
}

// SetUniversalAppAccess sets field value
func (o *GroupsGet200ResponseDataInner) SetUniversalAppAccess(v string) {
	o.UniversalAppAccess = v
}

// GetUniversalResourceAccess returns the UniversalResourceAccess field value
func (o *GroupsGet200ResponseDataInner) GetUniversalResourceAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UniversalResourceAccess
}

// GetUniversalResourceAccessOk returns a tuple with the UniversalResourceAccess field value
// and a boolean to check if the value has been set.
func (o *GroupsGet200ResponseDataInner) GetUniversalResourceAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UniversalResourceAccess, true
}

// SetUniversalResourceAccess sets field value
func (o *GroupsGet200ResponseDataInner) SetUniversalResourceAccess(v string) {
	o.UniversalResourceAccess = v
}

// GetUniversalWorkflowAccess returns the UniversalWorkflowAccess field value
func (o *GroupsGet200ResponseDataInner) GetUniversalWorkflowAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UniversalWorkflowAccess
}

// GetUniversalWorkflowAccessOk returns a tuple with the UniversalWorkflowAccess field value
// and a boolean to check if the value has been set.
func (o *GroupsGet200ResponseDataInner) GetUniversalWorkflowAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UniversalWorkflowAccess, true
}

// SetUniversalWorkflowAccess sets field value
func (o *GroupsGet200ResponseDataInner) SetUniversalWorkflowAccess(v string) {
	o.UniversalWorkflowAccess = v
}

// GetUniversalQueryLibraryAccess returns the UniversalQueryLibraryAccess field value
func (o *GroupsGet200ResponseDataInner) GetUniversalQueryLibraryAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UniversalQueryLibraryAccess
}

// GetUniversalQueryLibraryAccessOk returns a tuple with the UniversalQueryLibraryAccess field value
// and a boolean to check if the value has been set.
func (o *GroupsGet200ResponseDataInner) GetUniversalQueryLibraryAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UniversalQueryLibraryAccess, true
}

// SetUniversalQueryLibraryAccess sets field value
func (o *GroupsGet200ResponseDataInner) SetUniversalQueryLibraryAccess(v string) {
	o.UniversalQueryLibraryAccess = v
}

// GetUserInvites returns the UserInvites field value
func (o *GroupsGet200ResponseDataInner) GetUserInvites() []GroupsGroupIdGet200ResponseDataUserInvitesInner {
	if o == nil {
		var ret []GroupsGroupIdGet200ResponseDataUserInvitesInner
		return ret
	}

	return o.UserInvites
}

// GetUserInvitesOk returns a tuple with the UserInvites field value
// and a boolean to check if the value has been set.
func (o *GroupsGet200ResponseDataInner) GetUserInvitesOk() ([]GroupsGroupIdGet200ResponseDataUserInvitesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserInvites, true
}

// SetUserInvites sets field value
func (o *GroupsGet200ResponseDataInner) SetUserInvites(v []GroupsGroupIdGet200ResponseDataUserInvitesInner) {
	o.UserInvites = v
}

// GetUserListAccess returns the UserListAccess field value
func (o *GroupsGet200ResponseDataInner) GetUserListAccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UserListAccess
}

// GetUserListAccessOk returns a tuple with the UserListAccess field value
// and a boolean to check if the value has been set.
func (o *GroupsGet200ResponseDataInner) GetUserListAccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserListAccess, true
}

// SetUserListAccess sets field value
func (o *GroupsGet200ResponseDataInner) SetUserListAccess(v bool) {
	o.UserListAccess = v
}

// GetAuditLogAccess returns the AuditLogAccess field value
func (o *GroupsGet200ResponseDataInner) GetAuditLogAccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AuditLogAccess
}

// GetAuditLogAccessOk returns a tuple with the AuditLogAccess field value
// and a boolean to check if the value has been set.
func (o *GroupsGet200ResponseDataInner) GetAuditLogAccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuditLogAccess, true
}

// SetAuditLogAccess sets field value
func (o *GroupsGet200ResponseDataInner) SetAuditLogAccess(v bool) {
	o.AuditLogAccess = v
}

// GetUnpublishedReleaseAccess returns the UnpublishedReleaseAccess field value
func (o *GroupsGet200ResponseDataInner) GetUnpublishedReleaseAccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UnpublishedReleaseAccess
}

// GetUnpublishedReleaseAccessOk returns a tuple with the UnpublishedReleaseAccess field value
// and a boolean to check if the value has been set.
func (o *GroupsGet200ResponseDataInner) GetUnpublishedReleaseAccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnpublishedReleaseAccess, true
}

// SetUnpublishedReleaseAccess sets field value
func (o *GroupsGet200ResponseDataInner) SetUnpublishedReleaseAccess(v bool) {
	o.UnpublishedReleaseAccess = v
}

// GetUsageAnalyticsAccess returns the UsageAnalyticsAccess field value
func (o *GroupsGet200ResponseDataInner) GetUsageAnalyticsAccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UsageAnalyticsAccess
}

// GetUsageAnalyticsAccessOk returns a tuple with the UsageAnalyticsAccess field value
// and a boolean to check if the value has been set.
func (o *GroupsGet200ResponseDataInner) GetUsageAnalyticsAccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsageAnalyticsAccess, true
}

// SetUsageAnalyticsAccess sets field value
func (o *GroupsGet200ResponseDataInner) SetUsageAnalyticsAccess(v bool) {
	o.UsageAnalyticsAccess = v
}

// GetAccountDetailsAccess returns the AccountDetailsAccess field value
func (o *GroupsGet200ResponseDataInner) GetAccountDetailsAccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AccountDetailsAccess
}

// GetAccountDetailsAccessOk returns a tuple with the AccountDetailsAccess field value
// and a boolean to check if the value has been set.
func (o *GroupsGet200ResponseDataInner) GetAccountDetailsAccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountDetailsAccess, true
}

// SetAccountDetailsAccess sets field value
func (o *GroupsGet200ResponseDataInner) SetAccountDetailsAccess(v bool) {
	o.AccountDetailsAccess = v
}

// GetLandingPageAppId returns the LandingPageAppId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GroupsGet200ResponseDataInner) GetLandingPageAppId() string {
	if o == nil || o.LandingPageAppId.Get() == nil {
		var ret string
		return ret
	}

	return *o.LandingPageAppId.Get()
}

// GetLandingPageAppIdOk returns a tuple with the LandingPageAppId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupsGet200ResponseDataInner) GetLandingPageAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LandingPageAppId.Get(), o.LandingPageAppId.IsSet()
}

// SetLandingPageAppId sets field value
func (o *GroupsGet200ResponseDataInner) SetLandingPageAppId(v string) {
	o.LandingPageAppId.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *GroupsGet200ResponseDataInner) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *GroupsGet200ResponseDataInner) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *GroupsGet200ResponseDataInner) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *GroupsGet200ResponseDataInner) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *GroupsGet200ResponseDataInner) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *GroupsGet200ResponseDataInner) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o GroupsGet200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupsGet200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	toSerialize["legacy_id"] = o.LegacyId.Get()
	toSerialize["name"] = o.Name
	toSerialize["members"] = o.Members
	toSerialize["universal_app_access"] = o.UniversalAppAccess
	toSerialize["universal_resource_access"] = o.UniversalResourceAccess
	toSerialize["universal_workflow_access"] = o.UniversalWorkflowAccess
	toSerialize["universal_query_library_access"] = o.UniversalQueryLibraryAccess
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

func (o *GroupsGet200ResponseDataInner) UnmarshalJSON(data []byte) (err error) {
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
		"universal_query_library_access",
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

	varGroupsGet200ResponseDataInner := _GroupsGet200ResponseDataInner{}

	err = json.Unmarshal(data, &varGroupsGet200ResponseDataInner)

	if err != nil {
		return err
	}

	*o = GroupsGet200ResponseDataInner(varGroupsGet200ResponseDataInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "legacy_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "members")
		delete(additionalProperties, "universal_app_access")
		delete(additionalProperties, "universal_resource_access")
		delete(additionalProperties, "universal_workflow_access")
		delete(additionalProperties, "universal_query_library_access")
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

type NullableGroupsGet200ResponseDataInner struct {
	value *GroupsGet200ResponseDataInner
	isSet bool
}

func (v NullableGroupsGet200ResponseDataInner) Get() *GroupsGet200ResponseDataInner {
	return v.value
}

func (v *NullableGroupsGet200ResponseDataInner) Set(val *GroupsGet200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupsGet200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupsGet200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupsGet200ResponseDataInner(val *GroupsGet200ResponseDataInner) *NullableGroupsGet200ResponseDataInner {
	return &NullableGroupsGet200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGroupsGet200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupsGet200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


