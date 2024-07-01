# SourceControlConfigGet200ResponseDataAnyOf3Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Your Bitbucket username. | 
**Url** | Pointer to **string** | The domain used to access your self-hosted Bitbucket instance. Defaults to https://bitbucket.org/. | [optional] 
**EnterpriseApiUrl** | Pointer to **string** | The REST API route for your self-hosted Bitbucket instance. Defaults to https://api.bitbucket.org/2.0.  | [optional] 
**AppPassword** | **string** | Your Bitbucket app password. | 

## Methods

### NewSourceControlConfigGet200ResponseDataAnyOf3Config

`func NewSourceControlConfigGet200ResponseDataAnyOf3Config(username string, appPassword string, ) *SourceControlConfigGet200ResponseDataAnyOf3Config`

NewSourceControlConfigGet200ResponseDataAnyOf3Config instantiates a new SourceControlConfigGet200ResponseDataAnyOf3Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlConfigGet200ResponseDataAnyOf3ConfigWithDefaults

`func NewSourceControlConfigGet200ResponseDataAnyOf3ConfigWithDefaults() *SourceControlConfigGet200ResponseDataAnyOf3Config`

NewSourceControlConfigGet200ResponseDataAnyOf3ConfigWithDefaults instantiates a new SourceControlConfigGet200ResponseDataAnyOf3Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *SourceControlConfigGet200ResponseDataAnyOf3Config) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SourceControlConfigGet200ResponseDataAnyOf3Config) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SourceControlConfigGet200ResponseDataAnyOf3Config) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetUrl

`func (o *SourceControlConfigGet200ResponseDataAnyOf3Config) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SourceControlConfigGet200ResponseDataAnyOf3Config) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SourceControlConfigGet200ResponseDataAnyOf3Config) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SourceControlConfigGet200ResponseDataAnyOf3Config) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetEnterpriseApiUrl

`func (o *SourceControlConfigGet200ResponseDataAnyOf3Config) GetEnterpriseApiUrl() string`

GetEnterpriseApiUrl returns the EnterpriseApiUrl field if non-nil, zero value otherwise.

### GetEnterpriseApiUrlOk

`func (o *SourceControlConfigGet200ResponseDataAnyOf3Config) GetEnterpriseApiUrlOk() (*string, bool)`

GetEnterpriseApiUrlOk returns a tuple with the EnterpriseApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseApiUrl

`func (o *SourceControlConfigGet200ResponseDataAnyOf3Config) SetEnterpriseApiUrl(v string)`

SetEnterpriseApiUrl sets EnterpriseApiUrl field to given value.

### HasEnterpriseApiUrl

`func (o *SourceControlConfigGet200ResponseDataAnyOf3Config) HasEnterpriseApiUrl() bool`

HasEnterpriseApiUrl returns a boolean if a field has been set.

### GetAppPassword

`func (o *SourceControlConfigGet200ResponseDataAnyOf3Config) GetAppPassword() string`

GetAppPassword returns the AppPassword field if non-nil, zero value otherwise.

### GetAppPasswordOk

`func (o *SourceControlConfigGet200ResponseDataAnyOf3Config) GetAppPasswordOk() (*string, bool)`

GetAppPasswordOk returns a tuple with the AppPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPassword

`func (o *SourceControlConfigGet200ResponseDataAnyOf3Config) SetAppPassword(v string)`

SetAppPassword sets AppPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


