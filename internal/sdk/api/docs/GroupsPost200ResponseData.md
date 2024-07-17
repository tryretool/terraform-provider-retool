# GroupsPost200ResponseData

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
**UserInvites** | [**[]GroupsGroupIdGet200ResponseDataUserInvitesInner**](GroupsGroupIdGet200ResponseDataUserInvitesInner.md) | A list of user invites that will be added to the group | 
**UserListAccess** | **bool** | Whether the group has access to the user list | 
**AuditLogAccess** | **bool** | Whether the group has access to the audit log | 
**UnpublishedReleaseAccess** | **bool** | Whether the group has access to unpublished releases | 
**UsageAnalyticsAccess** | **bool** | Whether the group has access to usage analytics | 
**AccountDetailsAccess** | **bool** | Whether the group has access to account details | 
**LandingPageAppId** | **NullableString** | The app ID of the landing page | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 

## Methods

### NewGroupsPost200ResponseData

`func NewGroupsPost200ResponseData(id NullableFloat32, legacyId NullableFloat32, name string, members []GroupsGroupIdGet200ResponseDataMembersInner, universalAppAccess string, universalResourceAccess string, universalWorkflowAccess string, userInvites []GroupsGroupIdGet200ResponseDataUserInvitesInner, userListAccess bool, auditLogAccess bool, unpublishedReleaseAccess bool, usageAnalyticsAccess bool, accountDetailsAccess bool, landingPageAppId NullableString, createdAt string, updatedAt string, ) *GroupsPost200ResponseData`

NewGroupsPost200ResponseData instantiates a new GroupsPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupsPost200ResponseDataWithDefaults

`func NewGroupsPost200ResponseDataWithDefaults() *GroupsPost200ResponseData`

NewGroupsPost200ResponseDataWithDefaults instantiates a new GroupsPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupsPost200ResponseData) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupsPost200ResponseData) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupsPost200ResponseData) SetId(v float32)`

SetId sets Id field to given value.


### SetIdNil

`func (o *GroupsPost200ResponseData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GroupsPost200ResponseData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLegacyId

`func (o *GroupsPost200ResponseData) GetLegacyId() float32`

GetLegacyId returns the LegacyId field if non-nil, zero value otherwise.

### GetLegacyIdOk

`func (o *GroupsPost200ResponseData) GetLegacyIdOk() (*float32, bool)`

GetLegacyIdOk returns a tuple with the LegacyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyId

`func (o *GroupsPost200ResponseData) SetLegacyId(v float32)`

SetLegacyId sets LegacyId field to given value.


### SetLegacyIdNil

`func (o *GroupsPost200ResponseData) SetLegacyIdNil(b bool)`

 SetLegacyIdNil sets the value for LegacyId to be an explicit nil

### UnsetLegacyId
`func (o *GroupsPost200ResponseData) UnsetLegacyId()`

UnsetLegacyId ensures that no value is present for LegacyId, not even an explicit nil
### GetName

`func (o *GroupsPost200ResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupsPost200ResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupsPost200ResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetMembers

`func (o *GroupsPost200ResponseData) GetMembers() []GroupsGroupIdGet200ResponseDataMembersInner`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GroupsPost200ResponseData) GetMembersOk() (*[]GroupsGroupIdGet200ResponseDataMembersInner, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GroupsPost200ResponseData) SetMembers(v []GroupsGroupIdGet200ResponseDataMembersInner)`

SetMembers sets Members field to given value.


### GetUniversalAppAccess

`func (o *GroupsPost200ResponseData) GetUniversalAppAccess() string`

GetUniversalAppAccess returns the UniversalAppAccess field if non-nil, zero value otherwise.

### GetUniversalAppAccessOk

`func (o *GroupsPost200ResponseData) GetUniversalAppAccessOk() (*string, bool)`

GetUniversalAppAccessOk returns a tuple with the UniversalAppAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversalAppAccess

`func (o *GroupsPost200ResponseData) SetUniversalAppAccess(v string)`

SetUniversalAppAccess sets UniversalAppAccess field to given value.


### GetUniversalResourceAccess

`func (o *GroupsPost200ResponseData) GetUniversalResourceAccess() string`

GetUniversalResourceAccess returns the UniversalResourceAccess field if non-nil, zero value otherwise.

### GetUniversalResourceAccessOk

`func (o *GroupsPost200ResponseData) GetUniversalResourceAccessOk() (*string, bool)`

GetUniversalResourceAccessOk returns a tuple with the UniversalResourceAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversalResourceAccess

`func (o *GroupsPost200ResponseData) SetUniversalResourceAccess(v string)`

SetUniversalResourceAccess sets UniversalResourceAccess field to given value.


### GetUniversalWorkflowAccess

`func (o *GroupsPost200ResponseData) GetUniversalWorkflowAccess() string`

GetUniversalWorkflowAccess returns the UniversalWorkflowAccess field if non-nil, zero value otherwise.

### GetUniversalWorkflowAccessOk

`func (o *GroupsPost200ResponseData) GetUniversalWorkflowAccessOk() (*string, bool)`

GetUniversalWorkflowAccessOk returns a tuple with the UniversalWorkflowAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversalWorkflowAccess

`func (o *GroupsPost200ResponseData) SetUniversalWorkflowAccess(v string)`

SetUniversalWorkflowAccess sets UniversalWorkflowAccess field to given value.


### GetUserInvites

`func (o *GroupsPost200ResponseData) GetUserInvites() []GroupsGroupIdGet200ResponseDataUserInvitesInner`

GetUserInvites returns the UserInvites field if non-nil, zero value otherwise.

### GetUserInvitesOk

`func (o *GroupsPost200ResponseData) GetUserInvitesOk() (*[]GroupsGroupIdGet200ResponseDataUserInvitesInner, bool)`

GetUserInvitesOk returns a tuple with the UserInvites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInvites

`func (o *GroupsPost200ResponseData) SetUserInvites(v []GroupsGroupIdGet200ResponseDataUserInvitesInner)`

SetUserInvites sets UserInvites field to given value.


### GetUserListAccess

`func (o *GroupsPost200ResponseData) GetUserListAccess() bool`

GetUserListAccess returns the UserListAccess field if non-nil, zero value otherwise.

### GetUserListAccessOk

`func (o *GroupsPost200ResponseData) GetUserListAccessOk() (*bool, bool)`

GetUserListAccessOk returns a tuple with the UserListAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserListAccess

`func (o *GroupsPost200ResponseData) SetUserListAccess(v bool)`

SetUserListAccess sets UserListAccess field to given value.


### GetAuditLogAccess

`func (o *GroupsPost200ResponseData) GetAuditLogAccess() bool`

GetAuditLogAccess returns the AuditLogAccess field if non-nil, zero value otherwise.

### GetAuditLogAccessOk

`func (o *GroupsPost200ResponseData) GetAuditLogAccessOk() (*bool, bool)`

GetAuditLogAccessOk returns a tuple with the AuditLogAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogAccess

`func (o *GroupsPost200ResponseData) SetAuditLogAccess(v bool)`

SetAuditLogAccess sets AuditLogAccess field to given value.


### GetUnpublishedReleaseAccess

`func (o *GroupsPost200ResponseData) GetUnpublishedReleaseAccess() bool`

GetUnpublishedReleaseAccess returns the UnpublishedReleaseAccess field if non-nil, zero value otherwise.

### GetUnpublishedReleaseAccessOk

`func (o *GroupsPost200ResponseData) GetUnpublishedReleaseAccessOk() (*bool, bool)`

GetUnpublishedReleaseAccessOk returns a tuple with the UnpublishedReleaseAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpublishedReleaseAccess

`func (o *GroupsPost200ResponseData) SetUnpublishedReleaseAccess(v bool)`

SetUnpublishedReleaseAccess sets UnpublishedReleaseAccess field to given value.


### GetUsageAnalyticsAccess

`func (o *GroupsPost200ResponseData) GetUsageAnalyticsAccess() bool`

GetUsageAnalyticsAccess returns the UsageAnalyticsAccess field if non-nil, zero value otherwise.

### GetUsageAnalyticsAccessOk

`func (o *GroupsPost200ResponseData) GetUsageAnalyticsAccessOk() (*bool, bool)`

GetUsageAnalyticsAccessOk returns a tuple with the UsageAnalyticsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageAnalyticsAccess

`func (o *GroupsPost200ResponseData) SetUsageAnalyticsAccess(v bool)`

SetUsageAnalyticsAccess sets UsageAnalyticsAccess field to given value.


### GetAccountDetailsAccess

`func (o *GroupsPost200ResponseData) GetAccountDetailsAccess() bool`

GetAccountDetailsAccess returns the AccountDetailsAccess field if non-nil, zero value otherwise.

### GetAccountDetailsAccessOk

`func (o *GroupsPost200ResponseData) GetAccountDetailsAccessOk() (*bool, bool)`

GetAccountDetailsAccessOk returns a tuple with the AccountDetailsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDetailsAccess

`func (o *GroupsPost200ResponseData) SetAccountDetailsAccess(v bool)`

SetAccountDetailsAccess sets AccountDetailsAccess field to given value.


### GetLandingPageAppId

`func (o *GroupsPost200ResponseData) GetLandingPageAppId() string`

GetLandingPageAppId returns the LandingPageAppId field if non-nil, zero value otherwise.

### GetLandingPageAppIdOk

`func (o *GroupsPost200ResponseData) GetLandingPageAppIdOk() (*string, bool)`

GetLandingPageAppIdOk returns a tuple with the LandingPageAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingPageAppId

`func (o *GroupsPost200ResponseData) SetLandingPageAppId(v string)`

SetLandingPageAppId sets LandingPageAppId field to given value.


### SetLandingPageAppIdNil

`func (o *GroupsPost200ResponseData) SetLandingPageAppIdNil(b bool)`

 SetLandingPageAppIdNil sets the value for LandingPageAppId to be an explicit nil

### UnsetLandingPageAppId
`func (o *GroupsPost200ResponseData) UnsetLandingPageAppId()`

UnsetLandingPageAppId ensures that no value is present for LandingPageAppId, not even an explicit nil
### GetCreatedAt

`func (o *GroupsPost200ResponseData) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GroupsPost200ResponseData) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GroupsPost200ResponseData) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *GroupsPost200ResponseData) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GroupsPost200ResponseData) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GroupsPost200ResponseData) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


