# AccessTokensPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | The name of the access token. | 
**Description** | Pointer to **NullableString** | A short, user-facing description of the access token. | [optional] 
**Scopes** | **[]string** | Scopes to include in the access token. All valid scopes can be found in the [documentation](https://docs.retool.com/reference/api/v2). Any invalid scopes are ignored. | 

## Methods

### NewAccessTokensPostRequest

`func NewAccessTokensPostRequest(label string, scopes []string, ) *AccessTokensPostRequest`

NewAccessTokensPostRequest instantiates a new AccessTokensPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokensPostRequestWithDefaults

`func NewAccessTokensPostRequestWithDefaults() *AccessTokensPostRequest`

NewAccessTokensPostRequestWithDefaults instantiates a new AccessTokensPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *AccessTokensPostRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AccessTokensPostRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AccessTokensPostRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *AccessTokensPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessTokensPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessTokensPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccessTokensPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AccessTokensPostRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AccessTokensPostRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetScopes

`func (o *AccessTokensPostRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AccessTokensPostRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AccessTokensPostRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


