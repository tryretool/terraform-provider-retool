# UserInvitesPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**LegacyId** | **float32** |  | 
**InvitedBy** | **string** |  | 
**InvitedEmail** | **string** |  | 
**ExpiresAt** | **string** |  | 
**ClaimedBy** | **NullableString** |  | 
**ClaimedAt** | **NullableString** |  | 
**UserType** | **NullableString** |  | 
**Metadata** | **map[string]interface{}** |  | [default to {}]
**CreatedAt** | **string** |  | 
**InviteLink** | Pointer to **string** | User invite link expires 7 days after generation. | [optional] 
**Groups** | Pointer to [**[]GroupsGroupIdGet200ResponseDataUserInvitesInnerGroupsInner**](GroupsGroupIdGet200ResponseDataUserInvitesInnerGroupsInner.md) | The groups that the user is invited to | [optional] 

## Methods

### NewUserInvitesPost200ResponseData

`func NewUserInvitesPost200ResponseData(id float32, legacyId float32, invitedBy string, invitedEmail string, expiresAt string, claimedBy NullableString, claimedAt NullableString, userType NullableString, metadata map[string]interface{}, createdAt string, ) *UserInvitesPost200ResponseData`

NewUserInvitesPost200ResponseData instantiates a new UserInvitesPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInvitesPost200ResponseDataWithDefaults

`func NewUserInvitesPost200ResponseDataWithDefaults() *UserInvitesPost200ResponseData`

NewUserInvitesPost200ResponseDataWithDefaults instantiates a new UserInvitesPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserInvitesPost200ResponseData) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserInvitesPost200ResponseData) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserInvitesPost200ResponseData) SetId(v float32)`

SetId sets Id field to given value.


### GetLegacyId

`func (o *UserInvitesPost200ResponseData) GetLegacyId() float32`

GetLegacyId returns the LegacyId field if non-nil, zero value otherwise.

### GetLegacyIdOk

`func (o *UserInvitesPost200ResponseData) GetLegacyIdOk() (*float32, bool)`

GetLegacyIdOk returns a tuple with the LegacyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyId

`func (o *UserInvitesPost200ResponseData) SetLegacyId(v float32)`

SetLegacyId sets LegacyId field to given value.


### GetInvitedBy

`func (o *UserInvitesPost200ResponseData) GetInvitedBy() string`

GetInvitedBy returns the InvitedBy field if non-nil, zero value otherwise.

### GetInvitedByOk

`func (o *UserInvitesPost200ResponseData) GetInvitedByOk() (*string, bool)`

GetInvitedByOk returns a tuple with the InvitedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedBy

`func (o *UserInvitesPost200ResponseData) SetInvitedBy(v string)`

SetInvitedBy sets InvitedBy field to given value.


### GetInvitedEmail

`func (o *UserInvitesPost200ResponseData) GetInvitedEmail() string`

GetInvitedEmail returns the InvitedEmail field if non-nil, zero value otherwise.

### GetInvitedEmailOk

`func (o *UserInvitesPost200ResponseData) GetInvitedEmailOk() (*string, bool)`

GetInvitedEmailOk returns a tuple with the InvitedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedEmail

`func (o *UserInvitesPost200ResponseData) SetInvitedEmail(v string)`

SetInvitedEmail sets InvitedEmail field to given value.


### GetExpiresAt

`func (o *UserInvitesPost200ResponseData) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *UserInvitesPost200ResponseData) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *UserInvitesPost200ResponseData) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### GetClaimedBy

`func (o *UserInvitesPost200ResponseData) GetClaimedBy() string`

GetClaimedBy returns the ClaimedBy field if non-nil, zero value otherwise.

### GetClaimedByOk

`func (o *UserInvitesPost200ResponseData) GetClaimedByOk() (*string, bool)`

GetClaimedByOk returns a tuple with the ClaimedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedBy

`func (o *UserInvitesPost200ResponseData) SetClaimedBy(v string)`

SetClaimedBy sets ClaimedBy field to given value.


### SetClaimedByNil

`func (o *UserInvitesPost200ResponseData) SetClaimedByNil(b bool)`

 SetClaimedByNil sets the value for ClaimedBy to be an explicit nil

### UnsetClaimedBy
`func (o *UserInvitesPost200ResponseData) UnsetClaimedBy()`

UnsetClaimedBy ensures that no value is present for ClaimedBy, not even an explicit nil
### GetClaimedAt

`func (o *UserInvitesPost200ResponseData) GetClaimedAt() string`

GetClaimedAt returns the ClaimedAt field if non-nil, zero value otherwise.

### GetClaimedAtOk

`func (o *UserInvitesPost200ResponseData) GetClaimedAtOk() (*string, bool)`

GetClaimedAtOk returns a tuple with the ClaimedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedAt

`func (o *UserInvitesPost200ResponseData) SetClaimedAt(v string)`

SetClaimedAt sets ClaimedAt field to given value.


### SetClaimedAtNil

`func (o *UserInvitesPost200ResponseData) SetClaimedAtNil(b bool)`

 SetClaimedAtNil sets the value for ClaimedAt to be an explicit nil

### UnsetClaimedAt
`func (o *UserInvitesPost200ResponseData) UnsetClaimedAt()`

UnsetClaimedAt ensures that no value is present for ClaimedAt, not even an explicit nil
### GetUserType

`func (o *UserInvitesPost200ResponseData) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *UserInvitesPost200ResponseData) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *UserInvitesPost200ResponseData) SetUserType(v string)`

SetUserType sets UserType field to given value.


### SetUserTypeNil

`func (o *UserInvitesPost200ResponseData) SetUserTypeNil(b bool)`

 SetUserTypeNil sets the value for UserType to be an explicit nil

### UnsetUserType
`func (o *UserInvitesPost200ResponseData) UnsetUserType()`

UnsetUserType ensures that no value is present for UserType, not even an explicit nil
### GetMetadata

`func (o *UserInvitesPost200ResponseData) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UserInvitesPost200ResponseData) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UserInvitesPost200ResponseData) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *UserInvitesPost200ResponseData) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *UserInvitesPost200ResponseData) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCreatedAt

`func (o *UserInvitesPost200ResponseData) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserInvitesPost200ResponseData) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserInvitesPost200ResponseData) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetInviteLink

`func (o *UserInvitesPost200ResponseData) GetInviteLink() string`

GetInviteLink returns the InviteLink field if non-nil, zero value otherwise.

### GetInviteLinkOk

`func (o *UserInvitesPost200ResponseData) GetInviteLinkOk() (*string, bool)`

GetInviteLinkOk returns a tuple with the InviteLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteLink

`func (o *UserInvitesPost200ResponseData) SetInviteLink(v string)`

SetInviteLink sets InviteLink field to given value.

### HasInviteLink

`func (o *UserInvitesPost200ResponseData) HasInviteLink() bool`

HasInviteLink returns a boolean if a field has been set.

### GetGroups

`func (o *UserInvitesPost200ResponseData) GetGroups() []GroupsGroupIdGet200ResponseDataUserInvitesInnerGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UserInvitesPost200ResponseData) GetGroupsOk() (*[]GroupsGroupIdGet200ResponseDataUserInvitesInnerGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UserInvitesPost200ResponseData) SetGroups(v []GroupsGroupIdGet200ResponseDataUserInvitesInnerGroupsInner)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UserInvitesPost200ResponseData) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


