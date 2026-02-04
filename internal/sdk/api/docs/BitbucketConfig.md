# BitbucketConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Your Bitbucket username. | 
**AppPassword** | **string** | Your Bitbucket app password. | 
**Url** | Pointer to **string** | The domain used to access your self-hosted Bitbucket instance. Defaults to https://bitbucket.org/. | [optional] 
**EnterpriseApiUrl** | Pointer to **string** | The REST API route for your self-hosted Bitbucket instance. Defaults to https://api.bitbucket.org/2.0.  | [optional] 
**Type** | **string** |  | 
**Token** | **string** | Your Bitbucket API token. | 

## Methods

### NewBitbucketConfig

`func NewBitbucketConfig(username string, appPassword string, type_ string, token string, ) *BitbucketConfig`

NewBitbucketConfig instantiates a new BitbucketConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBitbucketConfigWithDefaults

`func NewBitbucketConfigWithDefaults() *BitbucketConfig`

NewBitbucketConfigWithDefaults instantiates a new BitbucketConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *BitbucketConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *BitbucketConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *BitbucketConfig) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetAppPassword

`func (o *BitbucketConfig) GetAppPassword() string`

GetAppPassword returns the AppPassword field if non-nil, zero value otherwise.

### GetAppPasswordOk

`func (o *BitbucketConfig) GetAppPasswordOk() (*string, bool)`

GetAppPasswordOk returns a tuple with the AppPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPassword

`func (o *BitbucketConfig) SetAppPassword(v string)`

SetAppPassword sets AppPassword field to given value.


### GetUrl

`func (o *BitbucketConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BitbucketConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BitbucketConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *BitbucketConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetEnterpriseApiUrl

`func (o *BitbucketConfig) GetEnterpriseApiUrl() string`

GetEnterpriseApiUrl returns the EnterpriseApiUrl field if non-nil, zero value otherwise.

### GetEnterpriseApiUrlOk

`func (o *BitbucketConfig) GetEnterpriseApiUrlOk() (*string, bool)`

GetEnterpriseApiUrlOk returns a tuple with the EnterpriseApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseApiUrl

`func (o *BitbucketConfig) SetEnterpriseApiUrl(v string)`

SetEnterpriseApiUrl sets EnterpriseApiUrl field to given value.

### HasEnterpriseApiUrl

`func (o *BitbucketConfig) HasEnterpriseApiUrl() bool`

HasEnterpriseApiUrl returns a boolean if a field has been set.

### GetType

`func (o *BitbucketConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BitbucketConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BitbucketConfig) SetType(v string)`

SetType sets Type field to given value.


### GetToken

`func (o *BitbucketConfig) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *BitbucketConfig) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *BitbucketConfig) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


