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

// checks if the GroupsPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupsPostRequest{}

// GroupsPostRequest struct for GroupsPostRequest
type GroupsPostRequest struct {
	// The name of the group.
	Name string `json:"name"`
	// Users to add to the group. Pass in an empty list to create a group with no members.
	Members []GroupsGroupIdPutRequestMembersInner `json:"members,omitempty"`
	// The universal app access level for the group. This denotes the access level that this group has for all apps.
	UniversalAppAccess *string `json:"universal_app_access,omitempty"`
	// The universal resource access level for the group. This denotes the access level that this group has for all resources.
	UniversalResourceAccess *string `json:"universal_resource_access,omitempty"`
	// The universal workflow access level for the group. This denotes the access level that this group has for all workflows.
	UniversalWorkflowAccess *string `json:"universal_workflow_access,omitempty"`
	// Level of access that the group has to the Query Library.
	UniversalQueryLibraryAccess *string `json:"universal_query_library_access,omitempty"`
	// A list of user invites that will be added to the group
	UserInvites []GroupsGroupIdGet200ResponseDataUserInvitesInner `json:"user_invites,omitempty"`
	// Whether the group has access to the user list
	UserListAccess *bool `json:"user_list_access,omitempty"`
	// Whether the group has access to the audit log
	AuditLogAccess *bool `json:"audit_log_access,omitempty"`
	// Whether the group has access to unpublished releases
	UnpublishedReleaseAccess *bool `json:"unpublished_release_access,omitempty"`
	// Whether the group has access to usage analytics
	UsageAnalyticsAccess *bool `json:"usage_analytics_access,omitempty"`
	// Whether the group has access to account details
	AccountDetailsAccess *bool `json:"account_details_access,omitempty"`
	// The app ID of the landing page
	LandingPageAppId NullableString `json:"landing_page_app_id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupsPostRequest GroupsPostRequest

// NewGroupsPostRequest instantiates a new GroupsPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupsPostRequest(name string) *GroupsPostRequest {
	this := GroupsPostRequest{}
	this.Name = name
	return &this
}

// NewGroupsPostRequestWithDefaults instantiates a new GroupsPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupsPostRequestWithDefaults() *GroupsPostRequest {
	this := GroupsPostRequest{}
	return &this
}

// GetName returns the Name field value
func (o *GroupsPostRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GroupsPostRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GroupsPostRequest) SetName(v string) {
	o.Name = v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *GroupsPostRequest) GetMembers() []GroupsGroupIdPutRequestMembersInner {
	if o == nil || IsNil(o.Members) {
		var ret []GroupsGroupIdPutRequestMembersInner
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsPostRequest) GetMembersOk() ([]GroupsGroupIdPutRequestMembersInner, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *GroupsPostRequest) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []GroupsGroupIdPutRequestMembersInner and assigns it to the Members field.
func (o *GroupsPostRequest) SetMembers(v []GroupsGroupIdPutRequestMembersInner) {
	o.Members = v
}

// GetUniversalAppAccess returns the UniversalAppAccess field value if set, zero value otherwise.
func (o *GroupsPostRequest) GetUniversalAppAccess() string {
	if o == nil || IsNil(o.UniversalAppAccess) {
		var ret string
		return ret
	}
	return *o.UniversalAppAccess
}

// GetUniversalAppAccessOk returns a tuple with the UniversalAppAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsPostRequest) GetUniversalAppAccessOk() (*string, bool) {
	if o == nil || IsNil(o.UniversalAppAccess) {
		return nil, false
	}
	return o.UniversalAppAccess, true
}

// HasUniversalAppAccess returns a boolean if a field has been set.
func (o *GroupsPostRequest) HasUniversalAppAccess() bool {
	if o != nil && !IsNil(o.UniversalAppAccess) {
		return true
	}

	return false
}

// SetUniversalAppAccess gets a reference to the given string and assigns it to the UniversalAppAccess field.
func (o *GroupsPostRequest) SetUniversalAppAccess(v string) {
	o.UniversalAppAccess = &v
}

// GetUniversalResourceAccess returns the UniversalResourceAccess field value if set, zero value otherwise.
func (o *GroupsPostRequest) GetUniversalResourceAccess() string {
	if o == nil || IsNil(o.UniversalResourceAccess) {
		var ret string
		return ret
	}
	return *o.UniversalResourceAccess
}

// GetUniversalResourceAccessOk returns a tuple with the UniversalResourceAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsPostRequest) GetUniversalResourceAccessOk() (*string, bool) {
	if o == nil || IsNil(o.UniversalResourceAccess) {
		return nil, false
	}
	return o.UniversalResourceAccess, true
}

// HasUniversalResourceAccess returns a boolean if a field has been set.
func (o *GroupsPostRequest) HasUniversalResourceAccess() bool {
	if o != nil && !IsNil(o.UniversalResourceAccess) {
		return true
	}

	return false
}

// SetUniversalResourceAccess gets a reference to the given string and assigns it to the UniversalResourceAccess field.
func (o *GroupsPostRequest) SetUniversalResourceAccess(v string) {
	o.UniversalResourceAccess = &v
}

// GetUniversalWorkflowAccess returns the UniversalWorkflowAccess field value if set, zero value otherwise.
func (o *GroupsPostRequest) GetUniversalWorkflowAccess() string {
	if o == nil || IsNil(o.UniversalWorkflowAccess) {
		var ret string
		return ret
	}
	return *o.UniversalWorkflowAccess
}

// GetUniversalWorkflowAccessOk returns a tuple with the UniversalWorkflowAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsPostRequest) GetUniversalWorkflowAccessOk() (*string, bool) {
	if o == nil || IsNil(o.UniversalWorkflowAccess) {
		return nil, false
	}
	return o.UniversalWorkflowAccess, true
}

// HasUniversalWorkflowAccess returns a boolean if a field has been set.
func (o *GroupsPostRequest) HasUniversalWorkflowAccess() bool {
	if o != nil && !IsNil(o.UniversalWorkflowAccess) {
		return true
	}

	return false
}

// SetUniversalWorkflowAccess gets a reference to the given string and assigns it to the UniversalWorkflowAccess field.
func (o *GroupsPostRequest) SetUniversalWorkflowAccess(v string) {
	o.UniversalWorkflowAccess = &v
}

// GetUniversalQueryLibraryAccess returns the UniversalQueryLibraryAccess field value if set, zero value otherwise.
func (o *GroupsPostRequest) GetUniversalQueryLibraryAccess() string {
	if o == nil || IsNil(o.UniversalQueryLibraryAccess) {
		var ret string
		return ret
	}
	return *o.UniversalQueryLibraryAccess
}

// GetUniversalQueryLibraryAccessOk returns a tuple with the UniversalQueryLibraryAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsPostRequest) GetUniversalQueryLibraryAccessOk() (*string, bool) {
	if o == nil || IsNil(o.UniversalQueryLibraryAccess) {
		return nil, false
	}
	return o.UniversalQueryLibraryAccess, true
}

// HasUniversalQueryLibraryAccess returns a boolean if a field has been set.
func (o *GroupsPostRequest) HasUniversalQueryLibraryAccess() bool {
	if o != nil && !IsNil(o.UniversalQueryLibraryAccess) {
		return true
	}

	return false
}

// SetUniversalQueryLibraryAccess gets a reference to the given string and assigns it to the UniversalQueryLibraryAccess field.
func (o *GroupsPostRequest) SetUniversalQueryLibraryAccess(v string) {
	o.UniversalQueryLibraryAccess = &v
}

// GetUserInvites returns the UserInvites field value if set, zero value otherwise.
func (o *GroupsPostRequest) GetUserInvites() []GroupsGroupIdGet200ResponseDataUserInvitesInner {
	if o == nil || IsNil(o.UserInvites) {
		var ret []GroupsGroupIdGet200ResponseDataUserInvitesInner
		return ret
	}
	return o.UserInvites
}

// GetUserInvitesOk returns a tuple with the UserInvites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsPostRequest) GetUserInvitesOk() ([]GroupsGroupIdGet200ResponseDataUserInvitesInner, bool) {
	if o == nil || IsNil(o.UserInvites) {
		return nil, false
	}
	return o.UserInvites, true
}

// HasUserInvites returns a boolean if a field has been set.
func (o *GroupsPostRequest) HasUserInvites() bool {
	if o != nil && !IsNil(o.UserInvites) {
		return true
	}

	return false
}

// SetUserInvites gets a reference to the given []GroupsGroupIdGet200ResponseDataUserInvitesInner and assigns it to the UserInvites field.
func (o *GroupsPostRequest) SetUserInvites(v []GroupsGroupIdGet200ResponseDataUserInvitesInner) {
	o.UserInvites = v
}

// GetUserListAccess returns the UserListAccess field value if set, zero value otherwise.
func (o *GroupsPostRequest) GetUserListAccess() bool {
	if o == nil || IsNil(o.UserListAccess) {
		var ret bool
		return ret
	}
	return *o.UserListAccess
}

// GetUserListAccessOk returns a tuple with the UserListAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsPostRequest) GetUserListAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.UserListAccess) {
		return nil, false
	}
	return o.UserListAccess, true
}

// HasUserListAccess returns a boolean if a field has been set.
func (o *GroupsPostRequest) HasUserListAccess() bool {
	if o != nil && !IsNil(o.UserListAccess) {
		return true
	}

	return false
}

// SetUserListAccess gets a reference to the given bool and assigns it to the UserListAccess field.
func (o *GroupsPostRequest) SetUserListAccess(v bool) {
	o.UserListAccess = &v
}

// GetAuditLogAccess returns the AuditLogAccess field value if set, zero value otherwise.
func (o *GroupsPostRequest) GetAuditLogAccess() bool {
	if o == nil || IsNil(o.AuditLogAccess) {
		var ret bool
		return ret
	}
	return *o.AuditLogAccess
}

// GetAuditLogAccessOk returns a tuple with the AuditLogAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsPostRequest) GetAuditLogAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.AuditLogAccess) {
		return nil, false
	}
	return o.AuditLogAccess, true
}

// HasAuditLogAccess returns a boolean if a field has been set.
func (o *GroupsPostRequest) HasAuditLogAccess() bool {
	if o != nil && !IsNil(o.AuditLogAccess) {
		return true
	}

	return false
}

// SetAuditLogAccess gets a reference to the given bool and assigns it to the AuditLogAccess field.
func (o *GroupsPostRequest) SetAuditLogAccess(v bool) {
	o.AuditLogAccess = &v
}

// GetUnpublishedReleaseAccess returns the UnpublishedReleaseAccess field value if set, zero value otherwise.
func (o *GroupsPostRequest) GetUnpublishedReleaseAccess() bool {
	if o == nil || IsNil(o.UnpublishedReleaseAccess) {
		var ret bool
		return ret
	}
	return *o.UnpublishedReleaseAccess
}

// GetUnpublishedReleaseAccessOk returns a tuple with the UnpublishedReleaseAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsPostRequest) GetUnpublishedReleaseAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.UnpublishedReleaseAccess) {
		return nil, false
	}
	return o.UnpublishedReleaseAccess, true
}

// HasUnpublishedReleaseAccess returns a boolean if a field has been set.
func (o *GroupsPostRequest) HasUnpublishedReleaseAccess() bool {
	if o != nil && !IsNil(o.UnpublishedReleaseAccess) {
		return true
	}

	return false
}

// SetUnpublishedReleaseAccess gets a reference to the given bool and assigns it to the UnpublishedReleaseAccess field.
func (o *GroupsPostRequest) SetUnpublishedReleaseAccess(v bool) {
	o.UnpublishedReleaseAccess = &v
}

// GetUsageAnalyticsAccess returns the UsageAnalyticsAccess field value if set, zero value otherwise.
func (o *GroupsPostRequest) GetUsageAnalyticsAccess() bool {
	if o == nil || IsNil(o.UsageAnalyticsAccess) {
		var ret bool
		return ret
	}
	return *o.UsageAnalyticsAccess
}

// GetUsageAnalyticsAccessOk returns a tuple with the UsageAnalyticsAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsPostRequest) GetUsageAnalyticsAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.UsageAnalyticsAccess) {
		return nil, false
	}
	return o.UsageAnalyticsAccess, true
}

// HasUsageAnalyticsAccess returns a boolean if a field has been set.
func (o *GroupsPostRequest) HasUsageAnalyticsAccess() bool {
	if o != nil && !IsNil(o.UsageAnalyticsAccess) {
		return true
	}

	return false
}

// SetUsageAnalyticsAccess gets a reference to the given bool and assigns it to the UsageAnalyticsAccess field.
func (o *GroupsPostRequest) SetUsageAnalyticsAccess(v bool) {
	o.UsageAnalyticsAccess = &v
}

// GetAccountDetailsAccess returns the AccountDetailsAccess field value if set, zero value otherwise.
func (o *GroupsPostRequest) GetAccountDetailsAccess() bool {
	if o == nil || IsNil(o.AccountDetailsAccess) {
		var ret bool
		return ret
	}
	return *o.AccountDetailsAccess
}

// GetAccountDetailsAccessOk returns a tuple with the AccountDetailsAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsPostRequest) GetAccountDetailsAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.AccountDetailsAccess) {
		return nil, false
	}
	return o.AccountDetailsAccess, true
}

// HasAccountDetailsAccess returns a boolean if a field has been set.
func (o *GroupsPostRequest) HasAccountDetailsAccess() bool {
	if o != nil && !IsNil(o.AccountDetailsAccess) {
		return true
	}

	return false
}

// SetAccountDetailsAccess gets a reference to the given bool and assigns it to the AccountDetailsAccess field.
func (o *GroupsPostRequest) SetAccountDetailsAccess(v bool) {
	o.AccountDetailsAccess = &v
}

// GetLandingPageAppId returns the LandingPageAppId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupsPostRequest) GetLandingPageAppId() string {
	if o == nil || IsNil(o.LandingPageAppId.Get()) {
		var ret string
		return ret
	}
	return *o.LandingPageAppId.Get()
}

// GetLandingPageAppIdOk returns a tuple with the LandingPageAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupsPostRequest) GetLandingPageAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LandingPageAppId.Get(), o.LandingPageAppId.IsSet()
}

// HasLandingPageAppId returns a boolean if a field has been set.
func (o *GroupsPostRequest) HasLandingPageAppId() bool {
	if o != nil && o.LandingPageAppId.IsSet() {
		return true
	}

	return false
}

// SetLandingPageAppId gets a reference to the given NullableString and assigns it to the LandingPageAppId field.
func (o *GroupsPostRequest) SetLandingPageAppId(v string) {
	o.LandingPageAppId.Set(&v)
}
// SetLandingPageAppIdNil sets the value for LandingPageAppId to be an explicit nil
func (o *GroupsPostRequest) SetLandingPageAppIdNil() {
	o.LandingPageAppId.Set(nil)
}

// UnsetLandingPageAppId ensures that no value is present for LandingPageAppId, not even an explicit nil
func (o *GroupsPostRequest) UnsetLandingPageAppId() {
	o.LandingPageAppId.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GroupsPostRequest) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsPostRequest) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GroupsPostRequest) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GroupsPostRequest) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GroupsPostRequest) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsPostRequest) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GroupsPostRequest) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GroupsPostRequest) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o GroupsPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupsPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	if !IsNil(o.UniversalAppAccess) {
		toSerialize["universal_app_access"] = o.UniversalAppAccess
	}
	if !IsNil(o.UniversalResourceAccess) {
		toSerialize["universal_resource_access"] = o.UniversalResourceAccess
	}
	if !IsNil(o.UniversalWorkflowAccess) {
		toSerialize["universal_workflow_access"] = o.UniversalWorkflowAccess
	}
	if !IsNil(o.UniversalQueryLibraryAccess) {
		toSerialize["universal_query_library_access"] = o.UniversalQueryLibraryAccess
	}
	if !IsNil(o.UserInvites) {
		toSerialize["user_invites"] = o.UserInvites
	}
	if !IsNil(o.UserListAccess) {
		toSerialize["user_list_access"] = o.UserListAccess
	}
	if !IsNil(o.AuditLogAccess) {
		toSerialize["audit_log_access"] = o.AuditLogAccess
	}
	if !IsNil(o.UnpublishedReleaseAccess) {
		toSerialize["unpublished_release_access"] = o.UnpublishedReleaseAccess
	}
	if !IsNil(o.UsageAnalyticsAccess) {
		toSerialize["usage_analytics_access"] = o.UsageAnalyticsAccess
	}
	if !IsNil(o.AccountDetailsAccess) {
		toSerialize["account_details_access"] = o.AccountDetailsAccess
	}
	if o.LandingPageAppId.IsSet() {
		toSerialize["landing_page_app_id"] = o.LandingPageAppId.Get()
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupsPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varGroupsPostRequest := _GroupsPostRequest{}

	err = json.Unmarshal(data, &varGroupsPostRequest)

	if err != nil {
		return err
	}

	*o = GroupsPostRequest(varGroupsPostRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
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

type NullableGroupsPostRequest struct {
	value *GroupsPostRequest
	isSet bool
}

func (v NullableGroupsPostRequest) Get() *GroupsPostRequest {
	return v.value
}

func (v *NullableGroupsPostRequest) Set(val *GroupsPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupsPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupsPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupsPostRequest(val *GroupsPostRequest) *NullableGroupsPostRequest {
	return &NullableGroupsPostRequest{value: val, isSet: true}
}

func (v NullableGroupsPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupsPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


