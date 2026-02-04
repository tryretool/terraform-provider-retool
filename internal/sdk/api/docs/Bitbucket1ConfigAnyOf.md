# Bitbucket1ConfigAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Username** | **string** | Your Bitbucket username. | 
**AppPassword** | **string** | Your Bitbucket app password. | 
**Url** | Pointer to **string** | The domain used to access your self-hosted Bitbucket instance. Defaults to https://bitbucket.org/. | [optional] 
**EnterpriseApiUrl** | Pointer to **string** | The REST API route for your self-hosted Bitbucket instance. Defaults to https://api.bitbucket.org/2.0.  | [optional] 

## Methods

### NewBitbucket1ConfigAnyOf

`func NewBitbucket1ConfigAnyOf(username string, appPassword string, ) *Bitbucket1ConfigAnyOf`

NewBitbucket1ConfigAnyOf instantiates a new Bitbucket1ConfigAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBitbucket1ConfigAnyOfWithDefaults

`func NewBitbucket1ConfigAnyOfWithDefaults() *Bitbucket1ConfigAnyOf`

NewBitbucket1ConfigAnyOfWithDefaults instantiates a new Bitbucket1ConfigAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Bitbucket1ConfigAnyOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Bitbucket1ConfigAnyOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Bitbucket1ConfigAnyOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Bitbucket1ConfigAnyOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsername

`func (o *Bitbucket1ConfigAnyOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Bitbucket1ConfigAnyOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Bitbucket1ConfigAnyOf) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetAppPassword

`func (o *Bitbucket1ConfigAnyOf) GetAppPassword() string`

GetAppPassword returns the AppPassword field if non-nil, zero value otherwise.

### GetAppPasswordOk

`func (o *Bitbucket1ConfigAnyOf) GetAppPasswordOk() (*string, bool)`

GetAppPasswordOk returns a tuple with the AppPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPassword

`func (o *Bitbucket1ConfigAnyOf) SetAppPassword(v string)`

SetAppPassword sets AppPassword field to given value.


### GetUrl

`func (o *Bitbucket1ConfigAnyOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Bitbucket1ConfigAnyOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Bitbucket1ConfigAnyOf) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Bitbucket1ConfigAnyOf) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetEnterpriseApiUrl

`func (o *Bitbucket1ConfigAnyOf) GetEnterpriseApiUrl() string`

GetEnterpriseApiUrl returns the EnterpriseApiUrl field if non-nil, zero value otherwise.

### GetEnterpriseApiUrlOk

`func (o *Bitbucket1ConfigAnyOf) GetEnterpriseApiUrlOk() (*string, bool)`

GetEnterpriseApiUrlOk returns a tuple with the EnterpriseApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseApiUrl

`func (o *Bitbucket1ConfigAnyOf) SetEnterpriseApiUrl(v string)`

SetEnterpriseApiUrl sets EnterpriseApiUrl field to given value.

### HasEnterpriseApiUrl

`func (o *Bitbucket1ConfigAnyOf) HasEnterpriseApiUrl() bool`

HasEnterpriseApiUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


