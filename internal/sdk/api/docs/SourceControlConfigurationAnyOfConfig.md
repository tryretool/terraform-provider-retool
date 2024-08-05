# SourceControlConfigurationAnyOfConfig

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

### NewSourceControlConfigurationAnyOfConfig

`func NewSourceControlConfigurationAnyOfConfig(type_ string, appId string, installationId string, privateKey string, personalAccessToken string, ) *SourceControlConfigurationAnyOfConfig`

NewSourceControlConfigurationAnyOfConfig instantiates a new SourceControlConfigurationAnyOfConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlConfigurationAnyOfConfigWithDefaults

`func NewSourceControlConfigurationAnyOfConfigWithDefaults() *SourceControlConfigurationAnyOfConfig`

NewSourceControlConfigurationAnyOfConfigWithDefaults instantiates a new SourceControlConfigurationAnyOfConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SourceControlConfigurationAnyOfConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SourceControlConfigurationAnyOfConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SourceControlConfigurationAnyOfConfig) SetType(v string)`

SetType sets Type field to given value.


### GetAppId

`func (o *SourceControlConfigurationAnyOfConfig) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *SourceControlConfigurationAnyOfConfig) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *SourceControlConfigurationAnyOfConfig) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetInstallationId

`func (o *SourceControlConfigurationAnyOfConfig) GetInstallationId() string`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *SourceControlConfigurationAnyOfConfig) GetInstallationIdOk() (*string, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *SourceControlConfigurationAnyOfConfig) SetInstallationId(v string)`

SetInstallationId sets InstallationId field to given value.


### GetPrivateKey

`func (o *SourceControlConfigurationAnyOfConfig) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *SourceControlConfigurationAnyOfConfig) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *SourceControlConfigurationAnyOfConfig) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetUrl

`func (o *SourceControlConfigurationAnyOfConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SourceControlConfigurationAnyOfConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SourceControlConfigurationAnyOfConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SourceControlConfigurationAnyOfConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetEnterpriseApiUrl

`func (o *SourceControlConfigurationAnyOfConfig) GetEnterpriseApiUrl() string`

GetEnterpriseApiUrl returns the EnterpriseApiUrl field if non-nil, zero value otherwise.

### GetEnterpriseApiUrlOk

`func (o *SourceControlConfigurationAnyOfConfig) GetEnterpriseApiUrlOk() (*string, bool)`

GetEnterpriseApiUrlOk returns a tuple with the EnterpriseApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseApiUrl

`func (o *SourceControlConfigurationAnyOfConfig) SetEnterpriseApiUrl(v string)`

SetEnterpriseApiUrl sets EnterpriseApiUrl field to given value.

### HasEnterpriseApiUrl

`func (o *SourceControlConfigurationAnyOfConfig) HasEnterpriseApiUrl() bool`

HasEnterpriseApiUrl returns a boolean if a field has been set.

### GetPersonalAccessToken

`func (o *SourceControlConfigurationAnyOfConfig) GetPersonalAccessToken() string`

GetPersonalAccessToken returns the PersonalAccessToken field if non-nil, zero value otherwise.

### GetPersonalAccessTokenOk

`func (o *SourceControlConfigurationAnyOfConfig) GetPersonalAccessTokenOk() (*string, bool)`

GetPersonalAccessTokenOk returns a tuple with the PersonalAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalAccessToken

`func (o *SourceControlConfigurationAnyOfConfig) SetPersonalAccessToken(v string)`

SetPersonalAccessToken sets PersonalAccessToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


