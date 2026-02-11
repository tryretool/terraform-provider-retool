# UserInvite

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

### NewUserInvite

`func NewUserInvite(id float32, legacyId float32, invitedBy string, invitedEmail string, expiresAt string, claimedBy NullableString, claimedAt NullableString, userType NullableString, metadata map[string]interface{}, createdAt string, ) *UserInvite`

NewUserInvite instantiates a new UserInvite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInviteWithDefaults

`func NewUserInviteWithDefaults() *UserInvite`

NewUserInviteWithDefaults instantiates a new UserInvite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserInvite) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserInvite) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserInvite) SetId(v float32)`

SetId sets Id field to given value.


### GetLegacyId

`func (o *UserInvite) GetLegacyId() float32`

GetLegacyId returns the LegacyId field if non-nil, zero value otherwise.

### GetLegacyIdOk

`func (o *UserInvite) GetLegacyIdOk() (*float32, bool)`

GetLegacyIdOk returns a tuple with the LegacyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyId

`func (o *UserInvite) SetLegacyId(v float32)`

SetLegacyId sets LegacyId field to given value.


### GetInvitedBy

`func (o *UserInvite) GetInvitedBy() string`

GetInvitedBy returns the InvitedBy field if non-nil, zero value otherwise.

### GetInvitedByOk

`func (o *UserInvite) GetInvitedByOk() (*string, bool)`

GetInvitedByOk returns a tuple with the InvitedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedBy

`func (o *UserInvite) SetInvitedBy(v string)`

SetInvitedBy sets InvitedBy field to given value.


### GetInvitedEmail

`func (o *UserInvite) GetInvitedEmail() string`

GetInvitedEmail returns the InvitedEmail field if non-nil, zero value otherwise.

### GetInvitedEmailOk

`func (o *UserInvite) GetInvitedEmailOk() (*string, bool)`

GetInvitedEmailOk returns a tuple with the InvitedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedEmail

`func (o *UserInvite) SetInvitedEmail(v string)`

SetInvitedEmail sets InvitedEmail field to given value.


### GetExpiresAt

`func (o *UserInvite) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *UserInvite) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *UserInvite) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### GetClaimedBy

`func (o *UserInvite) GetClaimedBy() string`

GetClaimedBy returns the ClaimedBy field if non-nil, zero value otherwise.

### GetClaimedByOk

`func (o *UserInvite) GetClaimedByOk() (*string, bool)`

GetClaimedByOk returns a tuple with the ClaimedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedBy

`func (o *UserInvite) SetClaimedBy(v string)`

SetClaimedBy sets ClaimedBy field to given value.


### SetClaimedByNil

`func (o *UserInvite) SetClaimedByNil(b bool)`

 SetClaimedByNil sets the value for ClaimedBy to be an explicit nil

### UnsetClaimedBy
`func (o *UserInvite) UnsetClaimedBy()`

UnsetClaimedBy ensures that no value is present for ClaimedBy, not even an explicit nil
### GetClaimedAt

`func (o *UserInvite) GetClaimedAt() string`

GetClaimedAt returns the ClaimedAt field if non-nil, zero value otherwise.

### GetClaimedAtOk

`func (o *UserInvite) GetClaimedAtOk() (*string, bool)`

GetClaimedAtOk returns a tuple with the ClaimedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedAt

`func (o *UserInvite) SetClaimedAt(v string)`

SetClaimedAt sets ClaimedAt field to given value.


### SetClaimedAtNil

`func (o *UserInvite) SetClaimedAtNil(b bool)`

 SetClaimedAtNil sets the value for ClaimedAt to be an explicit nil

### UnsetClaimedAt
`func (o *UserInvite) UnsetClaimedAt()`

UnsetClaimedAt ensures that no value is present for ClaimedAt, not even an explicit nil
### GetUserType

`func (o *UserInvite) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *UserInvite) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *UserInvite) SetUserType(v string)`

SetUserType sets UserType field to given value.


### SetUserTypeNil

`func (o *UserInvite) SetUserTypeNil(b bool)`

 SetUserTypeNil sets the value for UserType to be an explicit nil

### UnsetUserType
`func (o *UserInvite) UnsetUserType()`

UnsetUserType ensures that no value is present for UserType, not even an explicit nil
### GetMetadata

`func (o *UserInvite) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UserInvite) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UserInvite) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *UserInvite) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *UserInvite) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCreatedAt

`func (o *UserInvite) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserInvite) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserInvite) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetInviteLink

`func (o *UserInvite) GetInviteLink() string`

GetInviteLink returns the InviteLink field if non-nil, zero value otherwise.

### GetInviteLinkOk

`func (o *UserInvite) GetInviteLinkOk() (*string, bool)`

GetInviteLinkOk returns a tuple with the InviteLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteLink

`func (o *UserInvite) SetInviteLink(v string)`

SetInviteLink sets InviteLink field to given value.

### HasInviteLink

`func (o *UserInvite) HasInviteLink() bool`

HasInviteLink returns a boolean if a field has been set.

### GetGroups

`func (o *UserInvite) GetGroups() []GroupsGroupIdGet200ResponseDataUserInvitesInnerGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UserInvite) GetGroupsOk() (*[]GroupsGroupIdGet200ResponseDataUserInvitesInnerGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UserInvite) SetGroups(v []GroupsGroupIdGet200ResponseDataUserInvitesInnerGroupsInner)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UserInvite) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


