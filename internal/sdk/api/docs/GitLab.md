# GitLab

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**GitLabConfig**](GitLabConfig.md) |  | 
**Provider** | **string** |  | 
**Org** | **string** | The user or organization to which the repository belongs to. | 
**Repo** | **string** | The name of the repository you created to use with Retool. | 
**DefaultBranch** | **string** | The default branch, e.g., main. | 
**RepoVersion** | Pointer to **string** | Repositories using Toolscript are 2.0.0. Repositories using legacy YAML are 1.0.0. | [optional] 

## Methods

### NewGitLab

`func NewGitLab(config GitLabConfig, provider string, org string, repo string, defaultBranch string, ) *GitLab`

NewGitLab instantiates a new GitLab object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitLabWithDefaults

`func NewGitLabWithDefaults() *GitLab`

NewGitLabWithDefaults instantiates a new GitLab object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *GitLab) GetConfig() GitLabConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GitLab) GetConfigOk() (*GitLabConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GitLab) SetConfig(v GitLabConfig)`

SetConfig sets Config field to given value.


### GetProvider

`func (o *GitLab) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *GitLab) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *GitLab) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetOrg

`func (o *GitLab) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *GitLab) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *GitLab) SetOrg(v string)`

SetOrg sets Org field to given value.


### GetRepo

`func (o *GitLab) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *GitLab) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *GitLab) SetRepo(v string)`

SetRepo sets Repo field to given value.


### GetDefaultBranch

`func (o *GitLab) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *GitLab) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *GitLab) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.


### GetRepoVersion

`func (o *GitLab) GetRepoVersion() string`

GetRepoVersion returns the RepoVersion field if non-nil, zero value otherwise.

### GetRepoVersionOk

`func (o *GitLab) GetRepoVersionOk() (*string, bool)`

GetRepoVersionOk returns a tuple with the RepoVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoVersion

`func (o *GitLab) SetRepoVersion(v string)`

SetRepoVersion sets RepoVersion field to given value.

### HasRepoVersion

`func (o *GitLab) HasRepoVersion() bool`

HasRepoVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


