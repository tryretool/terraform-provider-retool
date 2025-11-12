# GroupsGet200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableFloat32** | The ID of the group. Currently this is the same legacy_id but will change in the future. | 
**LegacyId** | **NullableFloat32** | The legacy ID of the group. | 
**Name** | **string** | The name of the group. | 
**Members** | [**[]GroupsGroupIdGet200ResponseDataMembersInner**](GroupsGroupIdGet200ResponseDataMembersInner.md) | The members of the group. | 
**UniversalAppAccess** | **string** | The universal app access level for the group. This denotes the access level that this group has for all apps. | 
**UniversalResourceAccess** | **string** | The universal resource access level for the group. This denotes the access level that this group has for all resources. | 
**UniversalWorkflowAccess** | **string** | The universal workflow access level for the group. This denotes the access level that this group has for all workflows. | 
**UniversalQueryLibraryAccess** | **string** | Level of access that the group has to the Query Library. | 
**UserInvites** | [**[]GroupsGroupIdGet200ResponseDataUserInvitesInner**](GroupsGroupIdGet200ResponseDataUserInvitesInner.md) | A list of user invites that will be added to the group | 
**UserListAccess** | **bool** | Whether the group has access to the user list | 
**AuditLogAccess** | **bool** | Whether the group has access to the audit log | 
**UnpublishedReleaseAccess** | **bool** | Whether the group has access to unpublished releases | 
**UsageAnalyticsAccess** | **bool** | Whether the group has access to usage analytics | 
**ThemeAccess** | **bool** | Whether the group has access to edit themes | 
**AccountDetailsAccess** | **bool** | Whether the group has access to account details | 
**LandingPageAppId** | **NullableString** | The app ID of the landing page | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 

## Methods

### NewGroupsGet200ResponseDataInner

`func NewGroupsGet200ResponseDataInner(id NullableFloat32, legacyId NullableFloat32, name string, members []GroupsGroupIdGet200ResponseDataMembersInner, universalAppAccess string, universalResourceAccess string, universalWorkflowAccess string, universalQueryLibraryAccess string, userInvites []GroupsGroupIdGet200ResponseDataUserInvitesInner, userListAccess bool, auditLogAccess bool, unpublishedReleaseAccess bool, usageAnalyticsAccess bool, themeAccess bool, accountDetailsAccess bool, landingPageAppId NullableString, createdAt string, updatedAt string, ) *GroupsGet200ResponseDataInner`

NewGroupsGet200ResponseDataInner instantiates a new GroupsGet200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupsGet200ResponseDataInnerWithDefaults

`func NewGroupsGet200ResponseDataInnerWithDefaults() *GroupsGet200ResponseDataInner`

NewGroupsGet200ResponseDataInnerWithDefaults instantiates a new GroupsGet200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupsGet200ResponseDataInner) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupsGet200ResponseDataInner) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupsGet200ResponseDataInner) SetId(v float32)`

SetId sets Id field to given value.


### SetIdNil

`func (o *GroupsGet200ResponseDataInner) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GroupsGet200ResponseDataInner) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLegacyId

`func (o *GroupsGet200ResponseDataInner) GetLegacyId() float32`

GetLegacyId returns the LegacyId field if non-nil, zero value otherwise.

### GetLegacyIdOk

`func (o *GroupsGet200ResponseDataInner) GetLegacyIdOk() (*float32, bool)`

GetLegacyIdOk returns a tuple with the LegacyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyId

`func (o *GroupsGet200ResponseDataInner) SetLegacyId(v float32)`

SetLegacyId sets LegacyId field to given value.


### SetLegacyIdNil

`func (o *GroupsGet200ResponseDataInner) SetLegacyIdNil(b bool)`

 SetLegacyIdNil sets the value for LegacyId to be an explicit nil

### UnsetLegacyId
`func (o *GroupsGet200ResponseDataInner) UnsetLegacyId()`

UnsetLegacyId ensures that no value is present for LegacyId, not even an explicit nil
### GetName

`func (o *GroupsGet200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupsGet200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupsGet200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.


### GetMembers

`func (o *GroupsGet200ResponseDataInner) GetMembers() []GroupsGroupIdGet200ResponseDataMembersInner`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GroupsGet200ResponseDataInner) GetMembersOk() (*[]GroupsGroupIdGet200ResponseDataMembersInner, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GroupsGet200ResponseDataInner) SetMembers(v []GroupsGroupIdGet200ResponseDataMembersInner)`

SetMembers sets Members field to given value.


### GetUniversalAppAccess

`func (o *GroupsGet200ResponseDataInner) GetUniversalAppAccess() string`

GetUniversalAppAccess returns the UniversalAppAccess field if non-nil, zero value otherwise.

### GetUniversalAppAccessOk

`func (o *GroupsGet200ResponseDataInner) GetUniversalAppAccessOk() (*string, bool)`

GetUniversalAppAccessOk returns a tuple with the UniversalAppAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversalAppAccess

`func (o *GroupsGet200ResponseDataInner) SetUniversalAppAccess(v string)`

SetUniversalAppAccess sets UniversalAppAccess field to given value.


### GetUniversalResourceAccess

`func (o *GroupsGet200ResponseDataInner) GetUniversalResourceAccess() string`

GetUniversalResourceAccess returns the UniversalResourceAccess field if non-nil, zero value otherwise.

### GetUniversalResourceAccessOk

`func (o *GroupsGet200ResponseDataInner) GetUniversalResourceAccessOk() (*string, bool)`

GetUniversalResourceAccessOk returns a tuple with the UniversalResourceAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversalResourceAccess

`func (o *GroupsGet200ResponseDataInner) SetUniversalResourceAccess(v string)`

SetUniversalResourceAccess sets UniversalResourceAccess field to given value.


### GetUniversalWorkflowAccess

`func (o *GroupsGet200ResponseDataInner) GetUniversalWorkflowAccess() string`

GetUniversalWorkflowAccess returns the UniversalWorkflowAccess field if non-nil, zero value otherwise.

### GetUniversalWorkflowAccessOk

`func (o *GroupsGet200ResponseDataInner) GetUniversalWorkflowAccessOk() (*string, bool)`

GetUniversalWorkflowAccessOk returns a tuple with the UniversalWorkflowAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversalWorkflowAccess

`func (o *GroupsGet200ResponseDataInner) SetUniversalWorkflowAccess(v string)`

SetUniversalWorkflowAccess sets UniversalWorkflowAccess field to given value.


### GetUniversalQueryLibraryAccess

`func (o *GroupsGet200ResponseDataInner) GetUniversalQueryLibraryAccess() string`

GetUniversalQueryLibraryAccess returns the UniversalQueryLibraryAccess field if non-nil, zero value otherwise.

### GetUniversalQueryLibraryAccessOk

`func (o *GroupsGet200ResponseDataInner) GetUniversalQueryLibraryAccessOk() (*string, bool)`

GetUniversalQueryLibraryAccessOk returns a tuple with the UniversalQueryLibraryAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversalQueryLibraryAccess

`func (o *GroupsGet200ResponseDataInner) SetUniversalQueryLibraryAccess(v string)`

SetUniversalQueryLibraryAccess sets UniversalQueryLibraryAccess field to given value.


### GetUserInvites

`func (o *GroupsGet200ResponseDataInner) GetUserInvites() []GroupsGroupIdGet200ResponseDataUserInvitesInner`

GetUserInvites returns the UserInvites field if non-nil, zero value otherwise.

### GetUserInvitesOk

`func (o *GroupsGet200ResponseDataInner) GetUserInvitesOk() (*[]GroupsGroupIdGet200ResponseDataUserInvitesInner, bool)`

GetUserInvitesOk returns a tuple with the UserInvites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInvites

`func (o *GroupsGet200ResponseDataInner) SetUserInvites(v []GroupsGroupIdGet200ResponseDataUserInvitesInner)`

SetUserInvites sets UserInvites field to given value.


### GetUserListAccess

`func (o *GroupsGet200ResponseDataInner) GetUserListAccess() bool`

GetUserListAccess returns the UserListAccess field if non-nil, zero value otherwise.

### GetUserListAccessOk

`func (o *GroupsGet200ResponseDataInner) GetUserListAccessOk() (*bool, bool)`

GetUserListAccessOk returns a tuple with the UserListAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserListAccess

`func (o *GroupsGet200ResponseDataInner) SetUserListAccess(v bool)`

SetUserListAccess sets UserListAccess field to given value.


### GetAuditLogAccess

`func (o *GroupsGet200ResponseDataInner) GetAuditLogAccess() bool`

GetAuditLogAccess returns the AuditLogAccess field if non-nil, zero value otherwise.

### GetAuditLogAccessOk

`func (o *GroupsGet200ResponseDataInner) GetAuditLogAccessOk() (*bool, bool)`

GetAuditLogAccessOk returns a tuple with the AuditLogAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogAccess

`func (o *GroupsGet200ResponseDataInner) SetAuditLogAccess(v bool)`

SetAuditLogAccess sets AuditLogAccess field to given value.


### GetUnpublishedReleaseAccess

`func (o *GroupsGet200ResponseDataInner) GetUnpublishedReleaseAccess() bool`

GetUnpublishedReleaseAccess returns the UnpublishedReleaseAccess field if non-nil, zero value otherwise.

### GetUnpublishedReleaseAccessOk

`func (o *GroupsGet200ResponseDataInner) GetUnpublishedReleaseAccessOk() (*bool, bool)`

GetUnpublishedReleaseAccessOk returns a tuple with the UnpublishedReleaseAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpublishedReleaseAccess

`func (o *GroupsGet200ResponseDataInner) SetUnpublishedReleaseAccess(v bool)`

SetUnpublishedReleaseAccess sets UnpublishedReleaseAccess field to given value.


### GetUsageAnalyticsAccess

`func (o *GroupsGet200ResponseDataInner) GetUsageAnalyticsAccess() bool`

GetUsageAnalyticsAccess returns the UsageAnalyticsAccess field if non-nil, zero value otherwise.

### GetUsageAnalyticsAccessOk

`func (o *GroupsGet200ResponseDataInner) GetUsageAnalyticsAccessOk() (*bool, bool)`

GetUsageAnalyticsAccessOk returns a tuple with the UsageAnalyticsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageAnalyticsAccess

`func (o *GroupsGet200ResponseDataInner) SetUsageAnalyticsAccess(v bool)`

SetUsageAnalyticsAccess sets UsageAnalyticsAccess field to given value.


### GetThemeAccess

`func (o *GroupsGet200ResponseDataInner) GetThemeAccess() bool`

GetThemeAccess returns the ThemeAccess field if non-nil, zero value otherwise.

### GetThemeAccessOk

`func (o *GroupsGet200ResponseDataInner) GetThemeAccessOk() (*bool, bool)`

GetThemeAccessOk returns a tuple with the ThemeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemeAccess

`func (o *GroupsGet200ResponseDataInner) SetThemeAccess(v bool)`

SetThemeAccess sets ThemeAccess field to given value.


### GetAccountDetailsAccess

`func (o *GroupsGet200ResponseDataInner) GetAccountDetailsAccess() bool`

GetAccountDetailsAccess returns the AccountDetailsAccess field if non-nil, zero value otherwise.

### GetAccountDetailsAccessOk

`func (o *GroupsGet200ResponseDataInner) GetAccountDetailsAccessOk() (*bool, bool)`

GetAccountDetailsAccessOk returns a tuple with the AccountDetailsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDetailsAccess

`func (o *GroupsGet200ResponseDataInner) SetAccountDetailsAccess(v bool)`

SetAccountDetailsAccess sets AccountDetailsAccess field to given value.


### GetLandingPageAppId

`func (o *GroupsGet200ResponseDataInner) GetLandingPageAppId() string`

GetLandingPageAppId returns the LandingPageAppId field if non-nil, zero value otherwise.

### GetLandingPageAppIdOk

`func (o *GroupsGet200ResponseDataInner) GetLandingPageAppIdOk() (*string, bool)`

GetLandingPageAppIdOk returns a tuple with the LandingPageAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingPageAppId

`func (o *GroupsGet200ResponseDataInner) SetLandingPageAppId(v string)`

SetLandingPageAppId sets LandingPageAppId field to given value.


### SetLandingPageAppIdNil

`func (o *GroupsGet200ResponseDataInner) SetLandingPageAppIdNil(b bool)`

 SetLandingPageAppIdNil sets the value for LandingPageAppId to be an explicit nil

### UnsetLandingPageAppId
`func (o *GroupsGet200ResponseDataInner) UnsetLandingPageAppId()`

UnsetLandingPageAppId ensures that no value is present for LandingPageAppId, not even an explicit nil
### GetCreatedAt

`func (o *GroupsGet200ResponseDataInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GroupsGet200ResponseDataInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GroupsGet200ResponseDataInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *GroupsGet200ResponseDataInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GroupsGet200ResponseDataInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GroupsGet200ResponseDataInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


