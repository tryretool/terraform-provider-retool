# SourceControlConfigPutRequestConfigOneOfConfig

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

### NewSourceControlConfigPutRequestConfigOneOfConfig

`func NewSourceControlConfigPutRequestConfigOneOfConfig(type_ string, appId string, installationId string, privateKey string, personalAccessToken string, ) *SourceControlConfigPutRequestConfigOneOfConfig`

NewSourceControlConfigPutRequestConfigOneOfConfig instantiates a new SourceControlConfigPutRequestConfigOneOfConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlConfigPutRequestConfigOneOfConfigWithDefaults

`func NewSourceControlConfigPutRequestConfigOneOfConfigWithDefaults() *SourceControlConfigPutRequestConfigOneOfConfig`

NewSourceControlConfigPutRequestConfigOneOfConfigWithDefaults instantiates a new SourceControlConfigPutRequestConfigOneOfConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SourceControlConfigPutRequestConfigOneOfConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SourceControlConfigPutRequestConfigOneOfConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SourceControlConfigPutRequestConfigOneOfConfig) SetType(v string)`

SetType sets Type field to given value.


### GetAppId

`func (o *SourceControlConfigPutRequestConfigOneOfConfig) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *SourceControlConfigPutRequestConfigOneOfConfig) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *SourceControlConfigPutRequestConfigOneOfConfig) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetInstallationId

`func (o *SourceControlConfigPutRequestConfigOneOfConfig) GetInstallationId() string`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *SourceControlConfigPutRequestConfigOneOfConfig) GetInstallationIdOk() (*string, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *SourceControlConfigPutRequestConfigOneOfConfig) SetInstallationId(v string)`

SetInstallationId sets InstallationId field to given value.


### GetPrivateKey

`func (o *SourceControlConfigPutRequestConfigOneOfConfig) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *SourceControlConfigPutRequestConfigOneOfConfig) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *SourceControlConfigPutRequestConfigOneOfConfig) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetUrl

`func (o *SourceControlConfigPutRequestConfigOneOfConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SourceControlConfigPutRequestConfigOneOfConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SourceControlConfigPutRequestConfigOneOfConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SourceControlConfigPutRequestConfigOneOfConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetEnterpriseApiUrl

`func (o *SourceControlConfigPutRequestConfigOneOfConfig) GetEnterpriseApiUrl() string`

GetEnterpriseApiUrl returns the EnterpriseApiUrl field if non-nil, zero value otherwise.

### GetEnterpriseApiUrlOk

`func (o *SourceControlConfigPutRequestConfigOneOfConfig) GetEnterpriseApiUrlOk() (*string, bool)`

GetEnterpriseApiUrlOk returns a tuple with the EnterpriseApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseApiUrl

`func (o *SourceControlConfigPutRequestConfigOneOfConfig) SetEnterpriseApiUrl(v string)`

SetEnterpriseApiUrl sets EnterpriseApiUrl field to given value.

### HasEnterpriseApiUrl

`func (o *SourceControlConfigPutRequestConfigOneOfConfig) HasEnterpriseApiUrl() bool`

HasEnterpriseApiUrl returns a boolean if a field has been set.

### GetPersonalAccessToken

`func (o *SourceControlConfigPutRequestConfigOneOfConfig) GetPersonalAccessToken() string`

GetPersonalAccessToken returns the PersonalAccessToken field if non-nil, zero value otherwise.

### GetPersonalAccessTokenOk

`func (o *SourceControlConfigPutRequestConfigOneOfConfig) GetPersonalAccessTokenOk() (*string, bool)`

GetPersonalAccessTokenOk returns a tuple with the PersonalAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalAccessToken

`func (o *SourceControlConfigPutRequestConfigOneOfConfig) SetPersonalAccessToken(v string)`

SetPersonalAccessToken sets PersonalAccessToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


