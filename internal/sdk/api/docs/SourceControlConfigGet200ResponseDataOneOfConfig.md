# SourceControlConfigGet200ResponseDataOneOfConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**AppId** | **string** | The GitHub App ID. | 
**InstallationId** | **string** | The GitHub installation ID. This can be found at the end of the installation URL. | 
**PrivateKey** | **string** | The base64-encoded private key. | 
**Url** | Pointer to **string** | The domain used to access your self-hosted GitHub instance. | [optional] 
**EnterpriseApiUrl** | Pointer to **string** | The REST API route for your self-hosted GitHub instance. Defaults to https://[hostname]/api/v3. | [optional] 
**PersonalAccessToken** | **string** | The GitHub project access token to authenticate to the GitHub API.  | 

## Methods

### NewSourceControlConfigGet200ResponseDataOneOfConfig

`func NewSourceControlConfigGet200ResponseDataOneOfConfig(type_ string, appId string, installationId string, privateKey string, personalAccessToken string, ) *SourceControlConfigGet200ResponseDataOneOfConfig`

NewSourceControlConfigGet200ResponseDataOneOfConfig instantiates a new SourceControlConfigGet200ResponseDataOneOfConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlConfigGet200ResponseDataOneOfConfigWithDefaults

`func NewSourceControlConfigGet200ResponseDataOneOfConfigWithDefaults() *SourceControlConfigGet200ResponseDataOneOfConfig`

NewSourceControlConfigGet200ResponseDataOneOfConfigWithDefaults instantiates a new SourceControlConfigGet200ResponseDataOneOfConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SourceControlConfigGet200ResponseDataOneOfConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SourceControlConfigGet200ResponseDataOneOfConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SourceControlConfigGet200ResponseDataOneOfConfig) SetType(v string)`

SetType sets Type field to given value.


### GetAppId

`func (o *SourceControlConfigGet200ResponseDataOneOfConfig) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *SourceControlConfigGet200ResponseDataOneOfConfig) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *SourceControlConfigGet200ResponseDataOneOfConfig) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetInstallationId

`func (o *SourceControlConfigGet200ResponseDataOneOfConfig) GetInstallationId() string`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *SourceControlConfigGet200ResponseDataOneOfConfig) GetInstallationIdOk() (*string, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *SourceControlConfigGet200ResponseDataOneOfConfig) SetInstallationId(v string)`

SetInstallationId sets InstallationId field to given value.


### GetPrivateKey

`func (o *SourceControlConfigGet200ResponseDataOneOfConfig) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *SourceControlConfigGet200ResponseDataOneOfConfig) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *SourceControlConfigGet200ResponseDataOneOfConfig) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetUrl

`func (o *SourceControlConfigGet200ResponseDataOneOfConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SourceControlConfigGet200ResponseDataOneOfConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SourceControlConfigGet200ResponseDataOneOfConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SourceControlConfigGet200ResponseDataOneOfConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetEnterpriseApiUrl

`func (o *SourceControlConfigGet200ResponseDataOneOfConfig) GetEnterpriseApiUrl() string`

GetEnterpriseApiUrl returns the EnterpriseApiUrl field if non-nil, zero value otherwise.

### GetEnterpriseApiUrlOk

`func (o *SourceControlConfigGet200ResponseDataOneOfConfig) GetEnterpriseApiUrlOk() (*string, bool)`

GetEnterpriseApiUrlOk returns a tuple with the EnterpriseApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseApiUrl

`func (o *SourceControlConfigGet200ResponseDataOneOfConfig) SetEnterpriseApiUrl(v string)`

SetEnterpriseApiUrl sets EnterpriseApiUrl field to given value.

### HasEnterpriseApiUrl

`func (o *SourceControlConfigGet200ResponseDataOneOfConfig) HasEnterpriseApiUrl() bool`

HasEnterpriseApiUrl returns a boolean if a field has been set.

### GetPersonalAccessToken

`func (o *SourceControlConfigGet200ResponseDataOneOfConfig) GetPersonalAccessToken() string`

GetPersonalAccessToken returns the PersonalAccessToken field if non-nil, zero value otherwise.

### GetPersonalAccessTokenOk

`func (o *SourceControlConfigGet200ResponseDataOneOfConfig) GetPersonalAccessTokenOk() (*string, bool)`

GetPersonalAccessTokenOk returns a tuple with the PersonalAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalAccessToken

`func (o *SourceControlConfigGet200ResponseDataOneOfConfig) SetPersonalAccessToken(v string)`

SetPersonalAccessToken sets PersonalAccessToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


