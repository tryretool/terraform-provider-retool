# SourceControlConfigGet200ResponseDataOneOf4Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | Your base Azure URL. For Azure Cloud, this is always http://dev.azure.com. For Azure self-managed, this is the URL where your instance is hosted. | 
**Project** | **string** | Your new or existing Azure DevOps project. | 
**User** | **string** | The Azure Repos username. | 
**PersonalAccessToken** | **string** | The Azure project access tokens to authenticate to the Azure API. | 
**UseBasicAuth** | **bool** | Set this to true if you are using self-hosted Azure Repos. | 

## Methods

### NewSourceControlConfigGet200ResponseDataOneOf4Config

`func NewSourceControlConfigGet200ResponseDataOneOf4Config(url string, project string, user string, personalAccessToken string, useBasicAuth bool, ) *SourceControlConfigGet200ResponseDataOneOf4Config`

NewSourceControlConfigGet200ResponseDataOneOf4Config instantiates a new SourceControlConfigGet200ResponseDataOneOf4Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlConfigGet200ResponseDataOneOf4ConfigWithDefaults

`func NewSourceControlConfigGet200ResponseDataOneOf4ConfigWithDefaults() *SourceControlConfigGet200ResponseDataOneOf4Config`

NewSourceControlConfigGet200ResponseDataOneOf4ConfigWithDefaults instantiates a new SourceControlConfigGet200ResponseDataOneOf4Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *SourceControlConfigGet200ResponseDataOneOf4Config) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SourceControlConfigGet200ResponseDataOneOf4Config) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SourceControlConfigGet200ResponseDataOneOf4Config) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetProject

`func (o *SourceControlConfigGet200ResponseDataOneOf4Config) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *SourceControlConfigGet200ResponseDataOneOf4Config) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *SourceControlConfigGet200ResponseDataOneOf4Config) SetProject(v string)`

SetProject sets Project field to given value.


### GetUser

`func (o *SourceControlConfigGet200ResponseDataOneOf4Config) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SourceControlConfigGet200ResponseDataOneOf4Config) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SourceControlConfigGet200ResponseDataOneOf4Config) SetUser(v string)`

SetUser sets User field to given value.


### GetPersonalAccessToken

`func (o *SourceControlConfigGet200ResponseDataOneOf4Config) GetPersonalAccessToken() string`

GetPersonalAccessToken returns the PersonalAccessToken field if non-nil, zero value otherwise.

### GetPersonalAccessTokenOk

`func (o *SourceControlConfigGet200ResponseDataOneOf4Config) GetPersonalAccessTokenOk() (*string, bool)`

GetPersonalAccessTokenOk returns a tuple with the PersonalAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalAccessToken

`func (o *SourceControlConfigGet200ResponseDataOneOf4Config) SetPersonalAccessToken(v string)`

SetPersonalAccessToken sets PersonalAccessToken field to given value.


### GetUseBasicAuth

`func (o *SourceControlConfigGet200ResponseDataOneOf4Config) GetUseBasicAuth() bool`

GetUseBasicAuth returns the UseBasicAuth field if non-nil, zero value otherwise.

### GetUseBasicAuthOk

`func (o *SourceControlConfigGet200ResponseDataOneOf4Config) GetUseBasicAuthOk() (*bool, bool)`

GetUseBasicAuthOk returns a tuple with the UseBasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBasicAuth

`func (o *SourceControlConfigGet200ResponseDataOneOf4Config) SetUseBasicAuth(v bool)`

SetUseBasicAuth sets UseBasicAuth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


