# GroupsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the group. | 
**Members** | Pointer to [**[]GroupsGroupIdPutRequestMembersInner**](GroupsGroupIdPutRequestMembersInner.md) |  | [optional] 
**UniversalAppAccess** | Pointer to **string** | The universal app access level for the group. This denotes the access level that this group has for all apps. | [optional] 
**UniversalResourceAccess** | Pointer to **string** | The universal resource access level for the group. This denotes the access level that this group has for all resources. | [optional] 
**UniversalWorkflowAccess** | Pointer to **string** | The universal workflow access level for the group. This denotes the access level that this group has for all workflows. | [optional] 
**UserInvites** | Pointer to [**[]GroupsGroupIdGet200ResponseDataUserInvitesInner**](GroupsGroupIdGet200ResponseDataUserInvitesInner.md) | A list of user invites that will be added to the group | [optional] 
**UserListAccess** | Pointer to **bool** | Whether the group has access to the user list | [optional] 
**AuditLogAccess** | Pointer to **bool** | Whether the group has access to the audit log | [optional] 
**UnpublishedReleaseAccess** | Pointer to **bool** | Whether the group has access to unpublished releases | [optional] 
**UsageAnalyticsAccess** | Pointer to **bool** | Whether the group has access to usage analytics | [optional] 
**AccountDetailsAccess** | Pointer to **bool** | Whether the group has access to account details | [optional] 
**LandingPageAppId** | Pointer to **NullableString** | The app ID of the landing page | [optional] 

## Methods

### NewGroupsPostRequest

`func NewGroupsPostRequest(name string, ) *GroupsPostRequest`

NewGroupsPostRequest instantiates a new GroupsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupsPostRequestWithDefaults

`func NewGroupsPostRequestWithDefaults() *GroupsPostRequest`

NewGroupsPostRequestWithDefaults instantiates a new GroupsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GroupsPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupsPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupsPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetMembers

`func (o *GroupsPostRequest) GetMembers() []GroupsGroupIdPutRequestMembersInner`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GroupsPostRequest) GetMembersOk() (*[]GroupsGroupIdPutRequestMembersInner, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GroupsPostRequest) SetMembers(v []GroupsGroupIdPutRequestMembersInner)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *GroupsPostRequest) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetUniversalAppAccess

`func (o *GroupsPostRequest) GetUniversalAppAccess() string`

GetUniversalAppAccess returns the UniversalAppAccess field if non-nil, zero value otherwise.

### GetUniversalAppAccessOk

`func (o *GroupsPostRequest) GetUniversalAppAccessOk() (*string, bool)`

GetUniversalAppAccessOk returns a tuple with the UniversalAppAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversalAppAccess

`func (o *GroupsPostRequest) SetUniversalAppAccess(v string)`

SetUniversalAppAccess sets UniversalAppAccess field to given value.

### HasUniversalAppAccess

`func (o *GroupsPostRequest) HasUniversalAppAccess() bool`

HasUniversalAppAccess returns a boolean if a field has been set.

### GetUniversalResourceAccess

`func (o *GroupsPostRequest) GetUniversalResourceAccess() string`

GetUniversalResourceAccess returns the UniversalResourceAccess field if non-nil, zero value otherwise.

### GetUniversalResourceAccessOk

`func (o *GroupsPostRequest) GetUniversalResourceAccessOk() (*string, bool)`

GetUniversalResourceAccessOk returns a tuple with the UniversalResourceAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversalResourceAccess

`func (o *GroupsPostRequest) SetUniversalResourceAccess(v string)`

SetUniversalResourceAccess sets UniversalResourceAccess field to given value.

### HasUniversalResourceAccess

`func (o *GroupsPostRequest) HasUniversalResourceAccess() bool`

HasUniversalResourceAccess returns a boolean if a field has been set.

### GetUniversalWorkflowAccess

`func (o *GroupsPostRequest) GetUniversalWorkflowAccess() string`

GetUniversalWorkflowAccess returns the UniversalWorkflowAccess field if non-nil, zero value otherwise.

### GetUniversalWorkflowAccessOk

`func (o *GroupsPostRequest) GetUniversalWorkflowAccessOk() (*string, bool)`

GetUniversalWorkflowAccessOk returns a tuple with the UniversalWorkflowAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversalWorkflowAccess

`func (o *GroupsPostRequest) SetUniversalWorkflowAccess(v string)`

SetUniversalWorkflowAccess sets UniversalWorkflowAccess field to given value.

### HasUniversalWorkflowAccess

`func (o *GroupsPostRequest) HasUniversalWorkflowAccess() bool`

HasUniversalWorkflowAccess returns a boolean if a field has been set.

### GetUserInvites

`func (o *GroupsPostRequest) GetUserInvites() []GroupsGroupIdGet200ResponseDataUserInvitesInner`

GetUserInvites returns the UserInvites field if non-nil, zero value otherwise.

### GetUserInvitesOk

`func (o *GroupsPostRequest) GetUserInvitesOk() (*[]GroupsGroupIdGet200ResponseDataUserInvitesInner, bool)`

GetUserInvitesOk returns a tuple with the UserInvites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInvites

`func (o *GroupsPostRequest) SetUserInvites(v []GroupsGroupIdGet200ResponseDataUserInvitesInner)`

SetUserInvites sets UserInvites field to given value.

### HasUserInvites

`func (o *GroupsPostRequest) HasUserInvites() bool`

HasUserInvites returns a boolean if a field has been set.

### GetUserListAccess

`func (o *GroupsPostRequest) GetUserListAccess() bool`

GetUserListAccess returns the UserListAccess field if non-nil, zero value otherwise.

### GetUserListAccessOk

`func (o *GroupsPostRequest) GetUserListAccessOk() (*bool, bool)`

GetUserListAccessOk returns a tuple with the UserListAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserListAccess

`func (o *GroupsPostRequest) SetUserListAccess(v bool)`

SetUserListAccess sets UserListAccess field to given value.

### HasUserListAccess

`func (o *GroupsPostRequest) HasUserListAccess() bool`

HasUserListAccess returns a boolean if a field has been set.

### GetAuditLogAccess

`func (o *GroupsPostRequest) GetAuditLogAccess() bool`

GetAuditLogAccess returns the AuditLogAccess field if non-nil, zero value otherwise.

### GetAuditLogAccessOk

`func (o *GroupsPostRequest) GetAuditLogAccessOk() (*bool, bool)`

GetAuditLogAccessOk returns a tuple with the AuditLogAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogAccess

`func (o *GroupsPostRequest) SetAuditLogAccess(v bool)`

SetAuditLogAccess sets AuditLogAccess field to given value.

### HasAuditLogAccess

`func (o *GroupsPostRequest) HasAuditLogAccess() bool`

HasAuditLogAccess returns a boolean if a field has been set.

### GetUnpublishedReleaseAccess

`func (o *GroupsPostRequest) GetUnpublishedReleaseAccess() bool`

GetUnpublishedReleaseAccess returns the UnpublishedReleaseAccess field if non-nil, zero value otherwise.

### GetUnpublishedReleaseAccessOk

`func (o *GroupsPostRequest) GetUnpublishedReleaseAccessOk() (*bool, bool)`

GetUnpublishedReleaseAccessOk returns a tuple with the UnpublishedReleaseAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpublishedReleaseAccess

`func (o *GroupsPostRequest) SetUnpublishedReleaseAccess(v bool)`

SetUnpublishedReleaseAccess sets UnpublishedReleaseAccess field to given value.

### HasUnpublishedReleaseAccess

`func (o *GroupsPostRequest) HasUnpublishedReleaseAccess() bool`

HasUnpublishedReleaseAccess returns a boolean if a field has been set.

### GetUsageAnalyticsAccess

`func (o *GroupsPostRequest) GetUsageAnalyticsAccess() bool`

GetUsageAnalyticsAccess returns the UsageAnalyticsAccess field if non-nil, zero value otherwise.

### GetUsageAnalyticsAccessOk

`func (o *GroupsPostRequest) GetUsageAnalyticsAccessOk() (*bool, bool)`

GetUsageAnalyticsAccessOk returns a tuple with the UsageAnalyticsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageAnalyticsAccess

`func (o *GroupsPostRequest) SetUsageAnalyticsAccess(v bool)`

SetUsageAnalyticsAccess sets UsageAnalyticsAccess field to given value.

### HasUsageAnalyticsAccess

`func (o *GroupsPostRequest) HasUsageAnalyticsAccess() bool`

HasUsageAnalyticsAccess returns a boolean if a field has been set.

### GetAccountDetailsAccess

`func (o *GroupsPostRequest) GetAccountDetailsAccess() bool`

GetAccountDetailsAccess returns the AccountDetailsAccess field if non-nil, zero value otherwise.

### GetAccountDetailsAccessOk

`func (o *GroupsPostRequest) GetAccountDetailsAccessOk() (*bool, bool)`

GetAccountDetailsAccessOk returns a tuple with the AccountDetailsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDetailsAccess

`func (o *GroupsPostRequest) SetAccountDetailsAccess(v bool)`

SetAccountDetailsAccess sets AccountDetailsAccess field to given value.

### HasAccountDetailsAccess

`func (o *GroupsPostRequest) HasAccountDetailsAccess() bool`

HasAccountDetailsAccess returns a boolean if a field has been set.

### GetLandingPageAppId

`func (o *GroupsPostRequest) GetLandingPageAppId() string`

GetLandingPageAppId returns the LandingPageAppId field if non-nil, zero value otherwise.

### GetLandingPageAppIdOk

`func (o *GroupsPostRequest) GetLandingPageAppIdOk() (*string, bool)`

GetLandingPageAppIdOk returns a tuple with the LandingPageAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingPageAppId

`func (o *GroupsPostRequest) SetLandingPageAppId(v string)`

SetLandingPageAppId sets LandingPageAppId field to given value.

### HasLandingPageAppId

`func (o *GroupsPostRequest) HasLandingPageAppId() bool`

HasLandingPageAppId returns a boolean if a field has been set.

### SetLandingPageAppIdNil

`func (o *GroupsPostRequest) SetLandingPageAppIdNil(b bool)`

 SetLandingPageAppIdNil sets the value for LandingPageAppId to be an explicit nil

### UnsetLandingPageAppId
`func (o *GroupsPostRequest) UnsetLandingPageAppId()`

UnsetLandingPageAppId ensures that no value is present for LandingPageAppId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


