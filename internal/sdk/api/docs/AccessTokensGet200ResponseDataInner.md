# AccessTokensGet200ResponseDataInner

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

### NewAccessTokensGet200ResponseDataInner

`func NewAccessTokensGet200ResponseDataInner(id string, label string, description NullableString, last4 string, ownerLegacyId float32, scopes []string, createdAt time.Time, updatedAt time.Time, ) *AccessTokensGet200ResponseDataInner`

NewAccessTokensGet200ResponseDataInner instantiates a new AccessTokensGet200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokensGet200ResponseDataInnerWithDefaults

`func NewAccessTokensGet200ResponseDataInnerWithDefaults() *AccessTokensGet200ResponseDataInner`

NewAccessTokensGet200ResponseDataInnerWithDefaults instantiates a new AccessTokensGet200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccessTokensGet200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessTokensGet200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessTokensGet200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *AccessTokensGet200ResponseDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AccessTokensGet200ResponseDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AccessTokensGet200ResponseDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *AccessTokensGet200ResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessTokensGet200ResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessTokensGet200ResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *AccessTokensGet200ResponseDataInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AccessTokensGet200ResponseDataInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLast4

`func (o *AccessTokensGet200ResponseDataInner) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *AccessTokensGet200ResponseDataInner) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *AccessTokensGet200ResponseDataInner) SetLast4(v string)`

SetLast4 sets Last4 field to given value.


### GetOwnerLegacyId

`func (o *AccessTokensGet200ResponseDataInner) GetOwnerLegacyId() float32`

GetOwnerLegacyId returns the OwnerLegacyId field if non-nil, zero value otherwise.

### GetOwnerLegacyIdOk

`func (o *AccessTokensGet200ResponseDataInner) GetOwnerLegacyIdOk() (*float32, bool)`

GetOwnerLegacyIdOk returns a tuple with the OwnerLegacyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerLegacyId

`func (o *AccessTokensGet200ResponseDataInner) SetOwnerLegacyId(v float32)`

SetOwnerLegacyId sets OwnerLegacyId field to given value.


### GetScopes

`func (o *AccessTokensGet200ResponseDataInner) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AccessTokensGet200ResponseDataInner) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AccessTokensGet200ResponseDataInner) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetCreatedAt

`func (o *AccessTokensGet200ResponseDataInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccessTokensGet200ResponseDataInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccessTokensGet200ResponseDataInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AccessTokensGet200ResponseDataInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AccessTokensGet200ResponseDataInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AccessTokensGet200ResponseDataInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


