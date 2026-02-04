# BitbucketConfigAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Your Bitbucket username. | 
**AppPassword** | **string** | Your Bitbucket app password. | 
**Url** | Pointer to **string** | The domain used to access your self-hosted Bitbucket instance. Defaults to https://bitbucket.org/. | [optional] 
**EnterpriseApiUrl** | Pointer to **string** | The REST API route for your self-hosted Bitbucket instance. Defaults to https://api.bitbucket.org/2.0.  | [optional] 

## Methods

### NewBitbucketConfigAnyOf

`func NewBitbucketConfigAnyOf(username string, appPassword string, ) *BitbucketConfigAnyOf`

NewBitbucketConfigAnyOf instantiates a new BitbucketConfigAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBitbucketConfigAnyOfWithDefaults

`func NewBitbucketConfigAnyOfWithDefaults() *BitbucketConfigAnyOf`

NewBitbucketConfigAnyOfWithDefaults instantiates a new BitbucketConfigAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *BitbucketConfigAnyOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *BitbucketConfigAnyOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *BitbucketConfigAnyOf) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetAppPassword

`func (o *BitbucketConfigAnyOf) GetAppPassword() string`

GetAppPassword returns the AppPassword field if non-nil, zero value otherwise.

### GetAppPasswordOk

`func (o *BitbucketConfigAnyOf) GetAppPasswordOk() (*string, bool)`

GetAppPasswordOk returns a tuple with the AppPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPassword

`func (o *BitbucketConfigAnyOf) SetAppPassword(v string)`

SetAppPassword sets AppPassword field to given value.


### GetUrl

`func (o *BitbucketConfigAnyOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BitbucketConfigAnyOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BitbucketConfigAnyOf) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *BitbucketConfigAnyOf) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetEnterpriseApiUrl

`func (o *BitbucketConfigAnyOf) GetEnterpriseApiUrl() string`

GetEnterpriseApiUrl returns the EnterpriseApiUrl field if non-nil, zero value otherwise.

### GetEnterpriseApiUrlOk

`func (o *BitbucketConfigAnyOf) GetEnterpriseApiUrlOk() (*string, bool)`

GetEnterpriseApiUrlOk returns a tuple with the EnterpriseApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseApiUrl

`func (o *BitbucketConfigAnyOf) SetEnterpriseApiUrl(v string)`

SetEnterpriseApiUrl sets EnterpriseApiUrl field to given value.

### HasEnterpriseApiUrl

`func (o *BitbucketConfigAnyOf) HasEnterpriseApiUrl() bool`

HasEnterpriseApiUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


