# AzureReposConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | Your base Azure URL. For Azure Cloud, this is always http://dev.azure.com. For Azure self-managed, this is the URL where your instance is hosted. | 
**Project** | **string** | Your new or existing Azure DevOps project. | 
**User** | **string** | The Azure Repos username. | 
**PersonalAccessToken** | **string** | The Azure project access tokens to authenticate to the Azure API. | 
**UseBasicAuth** | **bool** | Set this to true if you are using self-hosted Azure Repos. | 

## Methods

### NewAzureReposConfig

`func NewAzureReposConfig(url string, project string, user string, personalAccessToken string, useBasicAuth bool, ) *AzureReposConfig`

NewAzureReposConfig instantiates a new AzureReposConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureReposConfigWithDefaults

`func NewAzureReposConfigWithDefaults() *AzureReposConfig`

NewAzureReposConfigWithDefaults instantiates a new AzureReposConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *AzureReposConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AzureReposConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AzureReposConfig) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetProject

`func (o *AzureReposConfig) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AzureReposConfig) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AzureReposConfig) SetProject(v string)`

SetProject sets Project field to given value.


### GetUser

`func (o *AzureReposConfig) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AzureReposConfig) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AzureReposConfig) SetUser(v string)`

SetUser sets User field to given value.


### GetPersonalAccessToken

`func (o *AzureReposConfig) GetPersonalAccessToken() string`

GetPersonalAccessToken returns the PersonalAccessToken field if non-nil, zero value otherwise.

### GetPersonalAccessTokenOk

`func (o *AzureReposConfig) GetPersonalAccessTokenOk() (*string, bool)`

GetPersonalAccessTokenOk returns a tuple with the PersonalAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalAccessToken

`func (o *AzureReposConfig) SetPersonalAccessToken(v string)`

SetPersonalAccessToken sets PersonalAccessToken field to given value.


### GetUseBasicAuth

`func (o *AzureReposConfig) GetUseBasicAuth() bool`

GetUseBasicAuth returns the UseBasicAuth field if non-nil, zero value otherwise.

### GetUseBasicAuthOk

`func (o *AzureReposConfig) GetUseBasicAuthOk() (*bool, bool)`

GetUseBasicAuthOk returns a tuple with the UseBasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBasicAuth

`func (o *AzureReposConfig) SetUseBasicAuth(v bool)`

SetUseBasicAuth sets UseBasicAuth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


