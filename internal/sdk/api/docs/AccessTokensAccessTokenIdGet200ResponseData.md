# AccessTokensAccessTokenIdGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The uuid of the access token. | 
**Label** | **string** | The name of the access token. | 
**Description** | **NullableString** | The description of the access token. | 
**Last4** | **string** | Last 4 characters of the access token secret | 
**OwnerLegacyId** | **float32** | The legacy id of the user that created the token | 
**Scopes** | **[]string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewAccessTokensAccessTokenIdGet200ResponseData

`func NewAccessTokensAccessTokenIdGet200ResponseData(id string, label string, description NullableString, last4 string, ownerLegacyId float32, scopes []string, createdAt time.Time, updatedAt time.Time, ) *AccessTokensAccessTokenIdGet200ResponseData`

NewAccessTokensAccessTokenIdGet200ResponseData instantiates a new AccessTokensAccessTokenIdGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokensAccessTokenIdGet200ResponseDataWithDefaults

`func NewAccessTokensAccessTokenIdGet200ResponseDataWithDefaults() *AccessTokensAccessTokenIdGet200ResponseData`

NewAccessTokensAccessTokenIdGet200ResponseDataWithDefaults instantiates a new AccessTokensAccessTokenIdGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccessTokensAccessTokenIdGet200ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessTokensAccessTokenIdGet200ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessTokensAccessTokenIdGet200ResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *AccessTokensAccessTokenIdGet200ResponseData) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AccessTokensAccessTokenIdGet200ResponseData) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AccessTokensAccessTokenIdGet200ResponseData) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *AccessTokensAccessTokenIdGet200ResponseData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessTokensAccessTokenIdGet200ResponseData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessTokensAccessTokenIdGet200ResponseData) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *AccessTokensAccessTokenIdGet200ResponseData) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AccessTokensAccessTokenIdGet200ResponseData) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLast4

`func (o *AccessTokensAccessTokenIdGet200ResponseData) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *AccessTokensAccessTokenIdGet200ResponseData) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *AccessTokensAccessTokenIdGet200ResponseData) SetLast4(v string)`

SetLast4 sets Last4 field to given value.


### GetOwnerLegacyId

`func (o *AccessTokensAccessTokenIdGet200ResponseData) GetOwnerLegacyId() float32`

GetOwnerLegacyId returns the OwnerLegacyId field if non-nil, zero value otherwise.

### GetOwnerLegacyIdOk

`func (o *AccessTokensAccessTokenIdGet200ResponseData) GetOwnerLegacyIdOk() (*float32, bool)`

GetOwnerLegacyIdOk returns a tuple with the OwnerLegacyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerLegacyId

`func (o *AccessTokensAccessTokenIdGet200ResponseData) SetOwnerLegacyId(v float32)`

SetOwnerLegacyId sets OwnerLegacyId field to given value.


### GetScopes

`func (o *AccessTokensAccessTokenIdGet200ResponseData) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AccessTokensAccessTokenIdGet200ResponseData) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AccessTokensAccessTokenIdGet200ResponseData) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetCreatedAt

`func (o *AccessTokensAccessTokenIdGet200ResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccessTokensAccessTokenIdGet200ResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccessTokensAccessTokenIdGet200ResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AccessTokensAccessTokenIdGet200ResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AccessTokensAccessTokenIdGet200ResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AccessTokensAccessTokenIdGet200ResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


