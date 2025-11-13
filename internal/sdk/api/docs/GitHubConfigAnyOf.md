# GitHubConfigAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**AppId** | **string** | The GitHub App ID. | 
**InstallationId** | **string** | The GitHub installation ID. This can be found at the end of the installation URL. | 
**PrivateKey** | **string** | The base64-encoded private key. | 
**Url** | Pointer to **string** | The domain used to access your self-hosted GitHub instance. | [optional] 
**EnterpriseApiUrl** | Pointer to **string** | The REST API route for your self-hosted GitHub instance. Defaults to https://[hostname]/api/v3. | [optional] 

## Methods

### NewGitHubConfigAnyOf

`func NewGitHubConfigAnyOf(type_ string, appId string, installationId string, privateKey string, ) *GitHubConfigAnyOf`

NewGitHubConfigAnyOf instantiates a new GitHubConfigAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitHubConfigAnyOfWithDefaults

`func NewGitHubConfigAnyOfWithDefaults() *GitHubConfigAnyOf`

NewGitHubConfigAnyOfWithDefaults instantiates a new GitHubConfigAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GitHubConfigAnyOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GitHubConfigAnyOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GitHubConfigAnyOf) SetType(v string)`

SetType sets Type field to given value.


### GetAppId

`func (o *GitHubConfigAnyOf) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *GitHubConfigAnyOf) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *GitHubConfigAnyOf) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetInstallationId

`func (o *GitHubConfigAnyOf) GetInstallationId() string`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *GitHubConfigAnyOf) GetInstallationIdOk() (*string, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *GitHubConfigAnyOf) SetInstallationId(v string)`

SetInstallationId sets InstallationId field to given value.


### GetPrivateKey

`func (o *GitHubConfigAnyOf) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *GitHubConfigAnyOf) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *GitHubConfigAnyOf) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetUrl

`func (o *GitHubConfigAnyOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GitHubConfigAnyOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GitHubConfigAnyOf) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GitHubConfigAnyOf) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetEnterpriseApiUrl

`func (o *GitHubConfigAnyOf) GetEnterpriseApiUrl() string`

GetEnterpriseApiUrl returns the EnterpriseApiUrl field if non-nil, zero value otherwise.

### GetEnterpriseApiUrlOk

`func (o *GitHubConfigAnyOf) GetEnterpriseApiUrlOk() (*string, bool)`

GetEnterpriseApiUrlOk returns a tuple with the EnterpriseApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseApiUrl

`func (o *GitHubConfigAnyOf) SetEnterpriseApiUrl(v string)`

SetEnterpriseApiUrl sets EnterpriseApiUrl field to given value.

### HasEnterpriseApiUrl

`func (o *GitHubConfigAnyOf) HasEnterpriseApiUrl() bool`

HasEnterpriseApiUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


