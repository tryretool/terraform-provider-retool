# GroupsGroupIdPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the group. | [optional] 
**Members** | Pointer to [**[]GroupsGroupIdPutRequestMembersInner**](GroupsGroupIdPutRequestMembersInner.md) | Users to add to the group. | [optional] 
**UniversalAppAccess** | Pointer to **string** | The universal app access level for the group. This denotes the access level that this group has for all apps. | [optional] 
**UniversalResourceAccess** | Pointer to **string** | The universal resource access level for the group. This denotes the access level that this group has for all resources. | [optional] 
**UniversalWorkflowAccess** | Pointer to **string** | The universal workflow access level for the group. This denotes the access level that this group has for all workflows. | [optional] 
**UniversalQueryLibraryAccess** | Pointer to **string** | Level of access that the group has to the Query Library. | [optional] 
**UserInvites** | Pointer to [**[]GroupsGroupIdGet200ResponseDataUserInvitesInner**](GroupsGroupIdGet200ResponseDataUserInvitesInner.md) | A list of user invites that will be added to the group | [optional] 
**UserListAccess** | Pointer to **bool** | Whether the group has access to the user list | [optional] 
**AuditLogAccess** | Pointer to **bool** | Whether the group has access to the audit log | [optional] 
**UnpublishedReleaseAccess** | Pointer to **bool** | Whether the group has access to unpublished releases | [optional] 
**UsageAnalyticsAccess** | Pointer to **bool** | Whether the group has access to usage analytics | [optional] 
**ThemeAccess** | Pointer to **bool** | Whether the group has access to edit themes | [optional] 
**AccountDetailsAccess** | Pointer to **bool** | Whether the group has access to account details | [optional] 
**LandingPageAppId** | Pointer to **NullableString** | The app ID of the landing page | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewGroupsGroupIdPutRequest

`func NewGroupsGroupIdPutRequest() *GroupsGroupIdPutRequest`

NewGroupsGroupIdPutRequest instantiates a new GroupsGroupIdPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupsGroupIdPutRequestWithDefaults

`func NewGroupsGroupIdPutRequestWithDefaults() *GroupsGroupIdPutRequest`

NewGroupsGroupIdPutRequestWithDefaults instantiates a new GroupsGroupIdPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GroupsGroupIdPutRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupsGroupIdPutRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupsGroupIdPutRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupsGroupIdPutRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMembers

`func (o *GroupsGroupIdPutRequest) GetMembers() []GroupsGroupIdPutRequestMembersInner`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GroupsGroupIdPutRequest) GetMembersOk() (*[]GroupsGroupIdPutRequestMembersInner, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GroupsGroupIdPutRequest) SetMembers(v []GroupsGroupIdPutRequestMembersInner)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *GroupsGroupIdPutRequest) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetUniversalAppAccess

`func (o *GroupsGroupIdPutRequest) GetUniversalAppAccess() string`

GetUniversalAppAccess returns the UniversalAppAccess field if non-nil, zero value otherwise.

### GetUniversalAppAccessOk

`func (o *GroupsGroupIdPutRequest) GetUniversalAppAccessOk() (*string, bool)`

GetUniversalAppAccessOk returns a tuple with the UniversalAppAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversalAppAccess

`func (o *GroupsGroupIdPutRequest) SetUniversalAppAccess(v string)`

SetUniversalAppAccess sets UniversalAppAccess field to given value.

### HasUniversalAppAccess

`func (o *GroupsGroupIdPutRequest) HasUniversalAppAccess() bool`

HasUniversalAppAccess returns a boolean if a field has been set.

### GetUniversalResourceAccess

`func (o *GroupsGroupIdPutRequest) GetUniversalResourceAccess() string`

GetUniversalResourceAccess returns the UniversalResourceAccess field if non-nil, zero value otherwise.

### GetUniversalResourceAccessOk

`func (o *GroupsGroupIdPutRequest) GetUniversalResourceAccessOk() (*string, bool)`

GetUniversalResourceAccessOk returns a tuple with the UniversalResourceAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversalResourceAccess

`func (o *GroupsGroupIdPutRequest) SetUniversalResourceAccess(v string)`

SetUniversalResourceAccess sets UniversalResourceAccess field to given value.

### HasUniversalResourceAccess

`func (o *GroupsGroupIdPutRequest) HasUniversalResourceAccess() bool`

HasUniversalResourceAccess returns a boolean if a field has been set.

### GetUniversalWorkflowAccess

`func (o *GroupsGroupIdPutRequest) GetUniversalWorkflowAccess() string`

GetUniversalWorkflowAccess returns the UniversalWorkflowAccess field if non-nil, zero value otherwise.

### GetUniversalWorkflowAccessOk

`func (o *GroupsGroupIdPutRequest) GetUniversalWorkflowAccessOk() (*string, bool)`

GetUniversalWorkflowAccessOk returns a tuple with the UniversalWorkflowAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversalWorkflowAccess

`func (o *GroupsGroupIdPutRequest) SetUniversalWorkflowAccess(v string)`

SetUniversalWorkflowAccess sets UniversalWorkflowAccess field to given value.

### HasUniversalWorkflowAccess

`func (o *GroupsGroupIdPutRequest) HasUniversalWorkflowAccess() bool`

HasUniversalWorkflowAccess returns a boolean if a field has been set.

### GetUniversalQueryLibraryAccess

`func (o *GroupsGroupIdPutRequest) GetUniversalQueryLibraryAccess() string`

GetUniversalQueryLibraryAccess returns the UniversalQueryLibraryAccess field if non-nil, zero value otherwise.

### GetUniversalQueryLibraryAccessOk

`func (o *GroupsGroupIdPutRequest) GetUniversalQueryLibraryAccessOk() (*string, bool)`

GetUniversalQueryLibraryAccessOk returns a tuple with the UniversalQueryLibraryAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversalQueryLibraryAccess

`func (o *GroupsGroupIdPutRequest) SetUniversalQueryLibraryAccess(v string)`

SetUniversalQueryLibraryAccess sets UniversalQueryLibraryAccess field to given value.

### HasUniversalQueryLibraryAccess

`func (o *GroupsGroupIdPutRequest) HasUniversalQueryLibraryAccess() bool`

HasUniversalQueryLibraryAccess returns a boolean if a field has been set.

### GetUserInvites

`func (o *GroupsGroupIdPutRequest) GetUserInvites() []GroupsGroupIdGet200ResponseDataUserInvitesInner`

GetUserInvites returns the UserInvites field if non-nil, zero value otherwise.

### GetUserInvitesOk

`func (o *GroupsGroupIdPutRequest) GetUserInvitesOk() (*[]GroupsGroupIdGet200ResponseDataUserInvitesInner, bool)`

GetUserInvitesOk returns a tuple with the UserInvites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInvites

`func (o *GroupsGroupIdPutRequest) SetUserInvites(v []GroupsGroupIdGet200ResponseDataUserInvitesInner)`

SetUserInvites sets UserInvites field to given value.

### HasUserInvites

`func (o *GroupsGroupIdPutRequest) HasUserInvites() bool`

HasUserInvites returns a boolean if a field has been set.

### GetUserListAccess

`func (o *GroupsGroupIdPutRequest) GetUserListAccess() bool`

GetUserListAccess returns the UserListAccess field if non-nil, zero value otherwise.

### GetUserListAccessOk

`func (o *GroupsGroupIdPutRequest) GetUserListAccessOk() (*bool, bool)`

GetUserListAccessOk returns a tuple with the UserListAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserListAccess

`func (o *GroupsGroupIdPutRequest) SetUserListAccess(v bool)`

SetUserListAccess sets UserListAccess field to given value.

### HasUserListAccess

`func (o *GroupsGroupIdPutRequest) HasUserListAccess() bool`

HasUserListAccess returns a boolean if a field has been set.

### GetAuditLogAccess

`func (o *GroupsGroupIdPutRequest) GetAuditLogAccess() bool`

GetAuditLogAccess returns the AuditLogAccess field if non-nil, zero value otherwise.

### GetAuditLogAccessOk

`func (o *GroupsGroupIdPutRequest) GetAuditLogAccessOk() (*bool, bool)`

GetAuditLogAccessOk returns a tuple with the AuditLogAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogAccess

`func (o *GroupsGroupIdPutRequest) SetAuditLogAccess(v bool)`

SetAuditLogAccess sets AuditLogAccess field to given value.

### HasAuditLogAccess

`func (o *GroupsGroupIdPutRequest) HasAuditLogAccess() bool`

HasAuditLogAccess returns a boolean if a field has been set.

### GetUnpublishedReleaseAccess

`func (o *GroupsGroupIdPutRequest) GetUnpublishedReleaseAccess() bool`

GetUnpublishedReleaseAccess returns the UnpublishedReleaseAccess field if non-nil, zero value otherwise.

### GetUnpublishedReleaseAccessOk

`func (o *GroupsGroupIdPutRequest) GetUnpublishedReleaseAccessOk() (*bool, bool)`

GetUnpublishedReleaseAccessOk returns a tuple with the UnpublishedReleaseAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpublishedReleaseAccess

`func (o *GroupsGroupIdPutRequest) SetUnpublishedReleaseAccess(v bool)`

SetUnpublishedReleaseAccess sets UnpublishedReleaseAccess field to given value.

### HasUnpublishedReleaseAccess

`func (o *GroupsGroupIdPutRequest) HasUnpublishedReleaseAccess() bool`

HasUnpublishedReleaseAccess returns a boolean if a field has been set.

### GetUsageAnalyticsAccess

`func (o *GroupsGroupIdPutRequest) GetUsageAnalyticsAccess() bool`

GetUsageAnalyticsAccess returns the UsageAnalyticsAccess field if non-nil, zero value otherwise.

### GetUsageAnalyticsAccessOk

`func (o *GroupsGroupIdPutRequest) GetUsageAnalyticsAccessOk() (*bool, bool)`

GetUsageAnalyticsAccessOk returns a tuple with the UsageAnalyticsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageAnalyticsAccess

`func (o *GroupsGroupIdPutRequest) SetUsageAnalyticsAccess(v bool)`

SetUsageAnalyticsAccess sets UsageAnalyticsAccess field to given value.

### HasUsageAnalyticsAccess

`func (o *GroupsGroupIdPutRequest) HasUsageAnalyticsAccess() bool`

HasUsageAnalyticsAccess returns a boolean if a field has been set.

### GetThemeAccess

`func (o *GroupsGroupIdPutRequest) GetThemeAccess() bool`

GetThemeAccess returns the ThemeAccess field if non-nil, zero value otherwise.

### GetThemeAccessOk

`func (o *GroupsGroupIdPutRequest) GetThemeAccessOk() (*bool, bool)`

GetThemeAccessOk returns a tuple with the ThemeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemeAccess

`func (o *GroupsGroupIdPutRequest) SetThemeAccess(v bool)`

SetThemeAccess sets ThemeAccess field to given value.

### HasThemeAccess

`func (o *GroupsGroupIdPutRequest) HasThemeAccess() bool`

HasThemeAccess returns a boolean if a field has been set.

### GetAccountDetailsAccess

`func (o *GroupsGroupIdPutRequest) GetAccountDetailsAccess() bool`

GetAccountDetailsAccess returns the AccountDetailsAccess field if non-nil, zero value otherwise.

### GetAccountDetailsAccessOk

`func (o *GroupsGroupIdPutRequest) GetAccountDetailsAccessOk() (*bool, bool)`

GetAccountDetailsAccessOk returns a tuple with the AccountDetailsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDetailsAccess

`func (o *GroupsGroupIdPutRequest) SetAccountDetailsAccess(v bool)`

SetAccountDetailsAccess sets AccountDetailsAccess field to given value.

### HasAccountDetailsAccess

`func (o *GroupsGroupIdPutRequest) HasAccountDetailsAccess() bool`

HasAccountDetailsAccess returns a boolean if a field has been set.

### GetLandingPageAppId

`func (o *GroupsGroupIdPutRequest) GetLandingPageAppId() string`

GetLandingPageAppId returns the LandingPageAppId field if non-nil, zero value otherwise.

### GetLandingPageAppIdOk

`func (o *GroupsGroupIdPutRequest) GetLandingPageAppIdOk() (*string, bool)`

GetLandingPageAppIdOk returns a tuple with the LandingPageAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingPageAppId

`func (o *GroupsGroupIdPutRequest) SetLandingPageAppId(v string)`

SetLandingPageAppId sets LandingPageAppId field to given value.

### HasLandingPageAppId

`func (o *GroupsGroupIdPutRequest) HasLandingPageAppId() bool`

HasLandingPageAppId returns a boolean if a field has been set.

### SetLandingPageAppIdNil

`func (o *GroupsGroupIdPutRequest) SetLandingPageAppIdNil(b bool)`

 SetLandingPageAppIdNil sets the value for LandingPageAppId to be an explicit nil

### UnsetLandingPageAppId
`func (o *GroupsGroupIdPutRequest) UnsetLandingPageAppId()`

UnsetLandingPageAppId ensures that no value is present for LandingPageAppId, not even an explicit nil
### GetCreatedAt

`func (o *GroupsGroupIdPutRequest) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GroupsGroupIdPutRequest) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GroupsGroupIdPutRequest) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GroupsGroupIdPutRequest) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GroupsGroupIdPutRequest) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GroupsGroupIdPutRequest) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GroupsGroupIdPutRequest) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GroupsGroupIdPutRequest) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


