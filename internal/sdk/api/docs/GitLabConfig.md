# GitLabConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **float32** | The numerical project ID for your GitLab project. Find this ID listed below the project&#39;s name on the project&#39;s homepage. | 
**Url** | **string** | Your base GitLab URL. On GitLab Cloud, this is always https://gitlab.com. On GitLab self-managed, this is the URL where your instance is hosted. | 
**ProjectAccessToken** | **string** | The GitLab project access token to authenticate to the GitLab API. | 

## Methods

### NewGitLabConfig

`func NewGitLabConfig(projectId float32, url string, projectAccessToken string, ) *GitLabConfig`

NewGitLabConfig instantiates a new GitLabConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitLabConfigWithDefaults

`func NewGitLabConfigWithDefaults() *GitLabConfig`

NewGitLabConfigWithDefaults instantiates a new GitLabConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *GitLabConfig) GetProjectId() float32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GitLabConfig) GetProjectIdOk() (*float32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GitLabConfig) SetProjectId(v float32)`

SetProjectId sets ProjectId field to given value.


### GetUrl

`func (o *GitLabConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GitLabConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GitLabConfig) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetProjectAccessToken

`func (o *GitLabConfig) GetProjectAccessToken() string`

GetProjectAccessToken returns the ProjectAccessToken field if non-nil, zero value otherwise.

### GetProjectAccessTokenOk

`func (o *GitLabConfig) GetProjectAccessTokenOk() (*string, bool)`

GetProjectAccessTokenOk returns a tuple with the ProjectAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectAccessToken

`func (o *GitLabConfig) SetProjectAccessToken(v string)`

SetProjectAccessToken sets ProjectAccessToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


