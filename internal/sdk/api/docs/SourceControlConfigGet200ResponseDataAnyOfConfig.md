# SourceControlConfigGet200ResponseDataAnyOfConfig

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

### NewSourceControlConfigGet200ResponseDataAnyOfConfig

`func NewSourceControlConfigGet200ResponseDataAnyOfConfig(type_ string, appId string, installationId string, privateKey string, personalAccessToken string, ) *SourceControlConfigGet200ResponseDataAnyOfConfig`

NewSourceControlConfigGet200ResponseDataAnyOfConfig instantiates a new SourceControlConfigGet200ResponseDataAnyOfConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlConfigGet200ResponseDataAnyOfConfigWithDefaults

`func NewSourceControlConfigGet200ResponseDataAnyOfConfigWithDefaults() *SourceControlConfigGet200ResponseDataAnyOfConfig`

NewSourceControlConfigGet200ResponseDataAnyOfConfigWithDefaults instantiates a new SourceControlConfigGet200ResponseDataAnyOfConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfig) SetType(v string)`

SetType sets Type field to given value.


### GetAppId

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfig) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfig) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfig) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetInstallationId

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfig) GetInstallationId() string`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfig) GetInstallationIdOk() (*string, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfig) SetInstallationId(v string)`

SetInstallationId sets InstallationId field to given value.


### GetPrivateKey

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfig) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfig) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfig) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetUrl

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetEnterpriseApiUrl

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfig) GetEnterpriseApiUrl() string`

GetEnterpriseApiUrl returns the EnterpriseApiUrl field if non-nil, zero value otherwise.

### GetEnterpriseApiUrlOk

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfig) GetEnterpriseApiUrlOk() (*string, bool)`

GetEnterpriseApiUrlOk returns a tuple with the EnterpriseApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseApiUrl

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfig) SetEnterpriseApiUrl(v string)`

SetEnterpriseApiUrl sets EnterpriseApiUrl field to given value.

### HasEnterpriseApiUrl

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfig) HasEnterpriseApiUrl() bool`

HasEnterpriseApiUrl returns a boolean if a field has been set.

### GetPersonalAccessToken

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfig) GetPersonalAccessToken() string`

GetPersonalAccessToken returns the PersonalAccessToken field if non-nil, zero value otherwise.

### GetPersonalAccessTokenOk

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfig) GetPersonalAccessTokenOk() (*string, bool)`

GetPersonalAccessTokenOk returns a tuple with the PersonalAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalAccessToken

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfig) SetPersonalAccessToken(v string)`

SetPersonalAccessToken sets PersonalAccessToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


