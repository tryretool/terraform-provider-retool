# GroupsGroupIdGet200ResponseDataUserInvitesInner

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
**InviteLink** | Pointer to **string** |  | [optional] 

## Methods

### NewGroupsGroupIdGet200ResponseDataUserInvitesInner

`func NewGroupsGroupIdGet200ResponseDataUserInvitesInner(id float32, legacyId float32, invitedBy string, invitedEmail string, expiresAt string, claimedBy NullableString, claimedAt NullableString, userType NullableString, metadata map[string]interface{}, createdAt string, ) *GroupsGroupIdGet200ResponseDataUserInvitesInner`

NewGroupsGroupIdGet200ResponseDataUserInvitesInner instantiates a new GroupsGroupIdGet200ResponseDataUserInvitesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupsGroupIdGet200ResponseDataUserInvitesInnerWithDefaults

`func NewGroupsGroupIdGet200ResponseDataUserInvitesInnerWithDefaults() *GroupsGroupIdGet200ResponseDataUserInvitesInner`

NewGroupsGroupIdGet200ResponseDataUserInvitesInnerWithDefaults instantiates a new GroupsGroupIdGet200ResponseDataUserInvitesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) SetId(v float32)`

SetId sets Id field to given value.


### GetLegacyId

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetLegacyId() float32`

GetLegacyId returns the LegacyId field if non-nil, zero value otherwise.

### GetLegacyIdOk

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetLegacyIdOk() (*float32, bool)`

GetLegacyIdOk returns a tuple with the LegacyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyId

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) SetLegacyId(v float32)`

SetLegacyId sets LegacyId field to given value.


### GetInvitedBy

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetInvitedBy() string`

GetInvitedBy returns the InvitedBy field if non-nil, zero value otherwise.

### GetInvitedByOk

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetInvitedByOk() (*string, bool)`

GetInvitedByOk returns a tuple with the InvitedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedBy

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) SetInvitedBy(v string)`

SetInvitedBy sets InvitedBy field to given value.


### GetInvitedEmail

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetInvitedEmail() string`

GetInvitedEmail returns the InvitedEmail field if non-nil, zero value otherwise.

### GetInvitedEmailOk

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetInvitedEmailOk() (*string, bool)`

GetInvitedEmailOk returns a tuple with the InvitedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedEmail

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) SetInvitedEmail(v string)`

SetInvitedEmail sets InvitedEmail field to given value.


### GetExpiresAt

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### GetClaimedBy

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetClaimedBy() string`

GetClaimedBy returns the ClaimedBy field if non-nil, zero value otherwise.

### GetClaimedByOk

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetClaimedByOk() (*string, bool)`

GetClaimedByOk returns a tuple with the ClaimedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedBy

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) SetClaimedBy(v string)`

SetClaimedBy sets ClaimedBy field to given value.


### SetClaimedByNil

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) SetClaimedByNil(b bool)`

 SetClaimedByNil sets the value for ClaimedBy to be an explicit nil

### UnsetClaimedBy
`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) UnsetClaimedBy()`

UnsetClaimedBy ensures that no value is present for ClaimedBy, not even an explicit nil
### GetClaimedAt

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetClaimedAt() string`

GetClaimedAt returns the ClaimedAt field if non-nil, zero value otherwise.

### GetClaimedAtOk

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetClaimedAtOk() (*string, bool)`

GetClaimedAtOk returns a tuple with the ClaimedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedAt

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) SetClaimedAt(v string)`

SetClaimedAt sets ClaimedAt field to given value.


### SetClaimedAtNil

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) SetClaimedAtNil(b bool)`

 SetClaimedAtNil sets the value for ClaimedAt to be an explicit nil

### UnsetClaimedAt
`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) UnsetClaimedAt()`

UnsetClaimedAt ensures that no value is present for ClaimedAt, not even an explicit nil
### GetUserType

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) SetUserType(v string)`

SetUserType sets UserType field to given value.


### SetUserTypeNil

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) SetUserTypeNil(b bool)`

 SetUserTypeNil sets the value for UserType to be an explicit nil

### UnsetUserType
`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) UnsetUserType()`

UnsetUserType ensures that no value is present for UserType, not even an explicit nil
### GetMetadata

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCreatedAt

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetInviteLink

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetInviteLink() string`

GetInviteLink returns the InviteLink field if non-nil, zero value otherwise.

### GetInviteLinkOk

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) GetInviteLinkOk() (*string, bool)`

GetInviteLinkOk returns a tuple with the InviteLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteLink

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) SetInviteLink(v string)`

SetInviteLink sets InviteLink field to given value.

### HasInviteLink

`func (o *GroupsGroupIdGet200ResponseDataUserInvitesInner) HasInviteLink() bool`

HasInviteLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


