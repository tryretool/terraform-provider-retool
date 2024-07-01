# UserInvitesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email of the invitee | 
**DefaultGroupIds** | Pointer to **[]float32** | Group IDs to invite the user to | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUserInvitesPostRequest

`func NewUserInvitesPostRequest(email string, ) *UserInvitesPostRequest`

NewUserInvitesPostRequest instantiates a new UserInvitesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInvitesPostRequestWithDefaults

`func NewUserInvitesPostRequestWithDefaults() *UserInvitesPostRequest`

NewUserInvitesPostRequestWithDefaults instantiates a new UserInvitesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserInvitesPostRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserInvitesPostRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserInvitesPostRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetDefaultGroupIds

`func (o *UserInvitesPostRequest) GetDefaultGroupIds() []float32`

GetDefaultGroupIds returns the DefaultGroupIds field if non-nil, zero value otherwise.

### GetDefaultGroupIdsOk

`func (o *UserInvitesPostRequest) GetDefaultGroupIdsOk() (*[]float32, bool)`

GetDefaultGroupIdsOk returns a tuple with the DefaultGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGroupIds

`func (o *UserInvitesPostRequest) SetDefaultGroupIds(v []float32)`

SetDefaultGroupIds sets DefaultGroupIds field to given value.

### HasDefaultGroupIds

`func (o *UserInvitesPostRequest) HasDefaultGroupIds() bool`

HasDefaultGroupIds returns a boolean if a field has been set.

### GetMetadata

`func (o *UserInvitesPostRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UserInvitesPostRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UserInvitesPostRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UserInvitesPostRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


