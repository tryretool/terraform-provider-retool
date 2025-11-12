# BitbucketConfigAnyOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Token** | **string** | Your Bitbucket API token. | 
**Url** | Pointer to **string** | The domain used to access your self-hosted Bitbucket instance. Defaults to https://bitbucket.org/. | [optional] 
**EnterpriseApiUrl** | Pointer to **string** | The REST API route for your self-hosted Bitbucket instance. Defaults to https://api.bitbucket.org/2.0.  | [optional] 

## Methods

### NewBitbucketConfigAnyOf1

`func NewBitbucketConfigAnyOf1(token string, ) *BitbucketConfigAnyOf1`

NewBitbucketConfigAnyOf1 instantiates a new BitbucketConfigAnyOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBitbucketConfigAnyOf1WithDefaults

`func NewBitbucketConfigAnyOf1WithDefaults() *BitbucketConfigAnyOf1`

NewBitbucketConfigAnyOf1WithDefaults instantiates a new BitbucketConfigAnyOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BitbucketConfigAnyOf1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BitbucketConfigAnyOf1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BitbucketConfigAnyOf1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BitbucketConfigAnyOf1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetToken

`func (o *BitbucketConfigAnyOf1) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *BitbucketConfigAnyOf1) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *BitbucketConfigAnyOf1) SetToken(v string)`

SetToken sets Token field to given value.


### GetUrl

`func (o *BitbucketConfigAnyOf1) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BitbucketConfigAnyOf1) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BitbucketConfigAnyOf1) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *BitbucketConfigAnyOf1) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetEnterpriseApiUrl

`func (o *BitbucketConfigAnyOf1) GetEnterpriseApiUrl() string`

GetEnterpriseApiUrl returns the EnterpriseApiUrl field if non-nil, zero value otherwise.

### GetEnterpriseApiUrlOk

`func (o *BitbucketConfigAnyOf1) GetEnterpriseApiUrlOk() (*string, bool)`

GetEnterpriseApiUrlOk returns a tuple with the EnterpriseApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseApiUrl

`func (o *BitbucketConfigAnyOf1) SetEnterpriseApiUrl(v string)`

SetEnterpriseApiUrl sets EnterpriseApiUrl field to given value.

### HasEnterpriseApiUrl

`func (o *BitbucketConfigAnyOf1) HasEnterpriseApiUrl() bool`

HasEnterpriseApiUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


