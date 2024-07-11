/*
Retool API

Go to Settings > API to get started. Once you generate an API token, use bearer token authentication to make requests.

API version: 2.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GroupsGroupIdPutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupsGroupIdPutRequest{}

// GroupsGroupIdPutRequest struct for GroupsGroupIdPutRequest
type GroupsGroupIdPutRequest struct {
	// The name of the group.
	Name *string `json:"name,omitempty"`
	Members []GroupsGroupIdPutRequestMembersInner `json:"members,omitempty"`
	// The universal app access level for the group. This denotes the access level that this group has for all apps.
	UniversalAppAccess *string `json:"universal_app_access,omitempty"`
	// The universal resource access level for the group. This denotes the access level that this group has for all resources.
	UniversalResourceAccess *string `json:"universal_resource_access,omitempty"`
	// The universal workflow access level for the group. This denotes the access level that this group has for all workflows.
	UniversalWorkflowAccess *string `json:"universal_workflow_access,omitempty"`
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
	AdditionalProperties map[string]interface{}
}

type _GroupsGroupIdPutRequest GroupsGroupIdPutRequest

// NewGroupsGroupIdPutRequest instantiates a new GroupsGroupIdPutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupsGroupIdPutRequest() *GroupsGroupIdPutRequest {
	this := GroupsGroupIdPutRequest{}
	return &this
}

// NewGroupsGroupIdPutRequestWithDefaults instantiates a new GroupsGroupIdPutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupsGroupIdPutRequestWithDefaults() *GroupsGroupIdPutRequest {
	this := GroupsGroupIdPutRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GroupsGroupIdPutRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdPutRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GroupsGroupIdPutRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GroupsGroupIdPutRequest) SetName(v string) {
	o.Name = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *GroupsGroupIdPutRequest) GetMembers() []GroupsGroupIdPutRequestMembersInner {
	if o == nil || IsNil(o.Members) {
		var ret []GroupsGroupIdPutRequestMembersInner
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdPutRequest) GetMembersOk() ([]GroupsGroupIdPutRequestMembersInner, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *GroupsGroupIdPutRequest) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []GroupsGroupIdPutRequestMembersInner and assigns it to the Members field.
func (o *GroupsGroupIdPutRequest) SetMembers(v []GroupsGroupIdPutRequestMembersInner) {
	o.Members = v
}

// GetUniversalAppAccess returns the UniversalAppAccess field value if set, zero value otherwise.
func (o *GroupsGroupIdPutRequest) GetUniversalAppAccess() string {
	if o == nil || IsNil(o.UniversalAppAccess) {
		var ret string
		return ret
	}
	return *o.UniversalAppAccess
}

// GetUniversalAppAccessOk returns a tuple with the UniversalAppAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdPutRequest) GetUniversalAppAccessOk() (*string, bool) {
	if o == nil || IsNil(o.UniversalAppAccess) {
		return nil, false
	}
	return o.UniversalAppAccess, true
}

// HasUniversalAppAccess returns a boolean if a field has been set.
func (o *GroupsGroupIdPutRequest) HasUniversalAppAccess() bool {
	if o != nil && !IsNil(o.UniversalAppAccess) {
		return true
	}

	return false
}

// SetUniversalAppAccess gets a reference to the given string and assigns it to the UniversalAppAccess field.
func (o *GroupsGroupIdPutRequest) SetUniversalAppAccess(v string) {
	o.UniversalAppAccess = &v
}

// GetUniversalResourceAccess returns the UniversalResourceAccess field value if set, zero value otherwise.
func (o *GroupsGroupIdPutRequest) GetUniversalResourceAccess() string {
	if o == nil || IsNil(o.UniversalResourceAccess) {
		var ret string
		return ret
	}
	return *o.UniversalResourceAccess
}

// GetUniversalResourceAccessOk returns a tuple with the UniversalResourceAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdPutRequest) GetUniversalResourceAccessOk() (*string, bool) {
	if o == nil || IsNil(o.UniversalResourceAccess) {
		return nil, false
	}
	return o.UniversalResourceAccess, true
}

// HasUniversalResourceAccess returns a boolean if a field has been set.
func (o *GroupsGroupIdPutRequest) HasUniversalResourceAccess() bool {
	if o != nil && !IsNil(o.UniversalResourceAccess) {
		return true
	}

	return false
}

// SetUniversalResourceAccess gets a reference to the given string and assigns it to the UniversalResourceAccess field.
func (o *GroupsGroupIdPutRequest) SetUniversalResourceAccess(v string) {
	o.UniversalResourceAccess = &v
}

// GetUniversalWorkflowAccess returns the UniversalWorkflowAccess field value if set, zero value otherwise.
func (o *GroupsGroupIdPutRequest) GetUniversalWorkflowAccess() string {
	if o == nil || IsNil(o.UniversalWorkflowAccess) {
		var ret string
		return ret
	}
	return *o.UniversalWorkflowAccess
}

// GetUniversalWorkflowAccessOk returns a tuple with the UniversalWorkflowAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdPutRequest) GetUniversalWorkflowAccessOk() (*string, bool) {
	if o == nil || IsNil(o.UniversalWorkflowAccess) {
		return nil, false
	}
	return o.UniversalWorkflowAccess, true
}

// HasUniversalWorkflowAccess returns a boolean if a field has been set.
func (o *GroupsGroupIdPutRequest) HasUniversalWorkflowAccess() bool {
	if o != nil && !IsNil(o.UniversalWorkflowAccess) {
		return true
	}

	return false
}

// SetUniversalWorkflowAccess gets a reference to the given string and assigns it to the UniversalWorkflowAccess field.
func (o *GroupsGroupIdPutRequest) SetUniversalWorkflowAccess(v string) {
	o.UniversalWorkflowAccess = &v
}

// GetUserInvites returns the UserInvites field value if set, zero value otherwise.
func (o *GroupsGroupIdPutRequest) GetUserInvites() []GroupsGroupIdGet200ResponseDataUserInvitesInner {
	if o == nil || IsNil(o.UserInvites) {
		var ret []GroupsGroupIdGet200ResponseDataUserInvitesInner
		return ret
	}
	return o.UserInvites
}

// GetUserInvitesOk returns a tuple with the UserInvites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdPutRequest) GetUserInvitesOk() ([]GroupsGroupIdGet200ResponseDataUserInvitesInner, bool) {
	if o == nil || IsNil(o.UserInvites) {
		return nil, false
	}
	return o.UserInvites, true
}

// HasUserInvites returns a boolean if a field has been set.
func (o *GroupsGroupIdPutRequest) HasUserInvites() bool {
	if o != nil && !IsNil(o.UserInvites) {
		return true
	}

	return false
}

// SetUserInvites gets a reference to the given []GroupsGroupIdGet200ResponseDataUserInvitesInner and assigns it to the UserInvites field.
func (o *GroupsGroupIdPutRequest) SetUserInvites(v []GroupsGroupIdGet200ResponseDataUserInvitesInner) {
	o.UserInvites = v
}

// GetUserListAccess returns the UserListAccess field value if set, zero value otherwise.
func (o *GroupsGroupIdPutRequest) GetUserListAccess() bool {
	if o == nil || IsNil(o.UserListAccess) {
		var ret bool
		return ret
	}
	return *o.UserListAccess
}

// GetUserListAccessOk returns a tuple with the UserListAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdPutRequest) GetUserListAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.UserListAccess) {
		return nil, false
	}
	return o.UserListAccess, true
}

// HasUserListAccess returns a boolean if a field has been set.
func (o *GroupsGroupIdPutRequest) HasUserListAccess() bool {
	if o != nil && !IsNil(o.UserListAccess) {
		return true
	}

	return false
}

// SetUserListAccess gets a reference to the given bool and assigns it to the UserListAccess field.
func (o *GroupsGroupIdPutRequest) SetUserListAccess(v bool) {
	o.UserListAccess = &v
}

// GetAuditLogAccess returns the AuditLogAccess field value if set, zero value otherwise.
func (o *GroupsGroupIdPutRequest) GetAuditLogAccess() bool {
	if o == nil || IsNil(o.AuditLogAccess) {
		var ret bool
		return ret
	}
	return *o.AuditLogAccess
}

// GetAuditLogAccessOk returns a tuple with the AuditLogAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdPutRequest) GetAuditLogAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.AuditLogAccess) {
		return nil, false
	}
	return o.AuditLogAccess, true
}

// HasAuditLogAccess returns a boolean if a field has been set.
func (o *GroupsGroupIdPutRequest) HasAuditLogAccess() bool {
	if o != nil && !IsNil(o.AuditLogAccess) {
		return true
	}

	return false
}

// SetAuditLogAccess gets a reference to the given bool and assigns it to the AuditLogAccess field.
func (o *GroupsGroupIdPutRequest) SetAuditLogAccess(v bool) {
	o.AuditLogAccess = &v
}

// GetUnpublishedReleaseAccess returns the UnpublishedReleaseAccess field value if set, zero value otherwise.
func (o *GroupsGroupIdPutRequest) GetUnpublishedReleaseAccess() bool {
	if o == nil || IsNil(o.UnpublishedReleaseAccess) {
		var ret bool
		return ret
	}
	return *o.UnpublishedReleaseAccess
}

// GetUnpublishedReleaseAccessOk returns a tuple with the UnpublishedReleaseAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdPutRequest) GetUnpublishedReleaseAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.UnpublishedReleaseAccess) {
		return nil, false
	}
	return o.UnpublishedReleaseAccess, true
}

// HasUnpublishedReleaseAccess returns a boolean if a field has been set.
func (o *GroupsGroupIdPutRequest) HasUnpublishedReleaseAccess() bool {
	if o != nil && !IsNil(o.UnpublishedReleaseAccess) {
		return true
	}

	return false
}

// SetUnpublishedReleaseAccess gets a reference to the given bool and assigns it to the UnpublishedReleaseAccess field.
func (o *GroupsGroupIdPutRequest) SetUnpublishedReleaseAccess(v bool) {
	o.UnpublishedReleaseAccess = &v
}

// GetUsageAnalyticsAccess returns the UsageAnalyticsAccess field value if set, zero value otherwise.
func (o *GroupsGroupIdPutRequest) GetUsageAnalyticsAccess() bool {
	if o == nil || IsNil(o.UsageAnalyticsAccess) {
		var ret bool
		return ret
	}
	return *o.UsageAnalyticsAccess
}

// GetUsageAnalyticsAccessOk returns a tuple with the UsageAnalyticsAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdPutRequest) GetUsageAnalyticsAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.UsageAnalyticsAccess) {
		return nil, false
	}
	return o.UsageAnalyticsAccess, true
}

// HasUsageAnalyticsAccess returns a boolean if a field has been set.
func (o *GroupsGroupIdPutRequest) HasUsageAnalyticsAccess() bool {
	if o != nil && !IsNil(o.UsageAnalyticsAccess) {
		return true
	}

	return false
}

// SetUsageAnalyticsAccess gets a reference to the given bool and assigns it to the UsageAnalyticsAccess field.
func (o *GroupsGroupIdPutRequest) SetUsageAnalyticsAccess(v bool) {
	o.UsageAnalyticsAccess = &v
}

// GetAccountDetailsAccess returns the AccountDetailsAccess field value if set, zero value otherwise.
func (o *GroupsGroupIdPutRequest) GetAccountDetailsAccess() bool {
	if o == nil || IsNil(o.AccountDetailsAccess) {
		var ret bool
		return ret
	}
	return *o.AccountDetailsAccess
}

// GetAccountDetailsAccessOk returns a tuple with the AccountDetailsAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGroupIdPutRequest) GetAccountDetailsAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.AccountDetailsAccess) {
		return nil, false
	}
	return o.AccountDetailsAccess, true
}

// HasAccountDetailsAccess returns a boolean if a field has been set.
func (o *GroupsGroupIdPutRequest) HasAccountDetailsAccess() bool {
	if o != nil && !IsNil(o.AccountDetailsAccess) {
		return true
	}

	return false
}

// SetAccountDetailsAccess gets a reference to the given bool and assigns it to the AccountDetailsAccess field.
func (o *GroupsGroupIdPutRequest) SetAccountDetailsAccess(v bool) {
	o.AccountDetailsAccess = &v
}

// GetLandingPageAppId returns the LandingPageAppId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupsGroupIdPutRequest) GetLandingPageAppId() string {
	if o == nil || IsNil(o.LandingPageAppId.Get()) {
		var ret string
		return ret
	}
	return *o.LandingPageAppId.Get()
}

// GetLandingPageAppIdOk returns a tuple with the LandingPageAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupsGroupIdPutRequest) GetLandingPageAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LandingPageAppId.Get(), o.LandingPageAppId.IsSet()
}

// HasLandingPageAppId returns a boolean if a field has been set.
func (o *GroupsGroupIdPutRequest) HasLandingPageAppId() bool {
	if o != nil && o.LandingPageAppId.IsSet() {
		return true
	}

	return false
}

// SetLandingPageAppId gets a reference to the given NullableString and assigns it to the LandingPageAppId field.
func (o *GroupsGroupIdPutRequest) SetLandingPageAppId(v string) {
	o.LandingPageAppId.Set(&v)
}
// SetLandingPageAppIdNil sets the value for LandingPageAppId to be an explicit nil
func (o *GroupsGroupIdPutRequest) SetLandingPageAppIdNil() {
	o.LandingPageAppId.Set(nil)
}

// UnsetLandingPageAppId ensures that no value is present for LandingPageAppId, not even an explicit nil
func (o *GroupsGroupIdPutRequest) UnsetLandingPageAppId() {
	o.LandingPageAppId.Unset()
}

func (o GroupsGroupIdPutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupsGroupIdPutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupsGroupIdPutRequest) UnmarshalJSON(data []byte) (err error) {
	varGroupsGroupIdPutRequest := _GroupsGroupIdPutRequest{}

	err = json.Unmarshal(data, &varGroupsGroupIdPutRequest)

	if err != nil {
		return err
	}

	*o = GroupsGroupIdPutRequest(varGroupsGroupIdPutRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
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
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupsGroupIdPutRequest struct {
	value *GroupsGroupIdPutRequest
	isSet bool
}

func (v NullableGroupsGroupIdPutRequest) Get() *GroupsGroupIdPutRequest {
	return v.value
}

func (v *NullableGroupsGroupIdPutRequest) Set(val *GroupsGroupIdPutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupsGroupIdPutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupsGroupIdPutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupsGroupIdPutRequest(val *GroupsGroupIdPutRequest) *NullableGroupsGroupIdPutRequest {
	return &NullableGroupsGroupIdPutRequest{value: val, isSet: true}
}

func (v NullableGroupsGroupIdPutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupsGroupIdPutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


