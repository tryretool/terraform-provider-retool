# AzureRepos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**AzureReposConfig**](AzureReposConfig.md) |  | 
**Provider** | **string** |  | 
**Org** | **string** | The user or organization to which the repository belongs to. | 
**Repo** | **string** | The name of the repository you created to use with Retool. | 
**DefaultBranch** | **string** | The default branch, e.g., main. | 
**RepoVersion** | Pointer to **string** | Repositories using Toolscript are 2.0.0. Repositories using legacy YAML are 1.0.0. | [optional] 

## Methods

### NewAzureRepos

`func NewAzureRepos(config AzureReposConfig, provider string, org string, repo string, defaultBranch string, ) *AzureRepos`

NewAzureRepos instantiates a new AzureRepos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureReposWithDefaults

`func NewAzureReposWithDefaults() *AzureRepos`

NewAzureReposWithDefaults instantiates a new AzureRepos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *AzureRepos) GetConfig() AzureReposConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AzureRepos) GetConfigOk() (*AzureReposConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AzureRepos) SetConfig(v AzureReposConfig)`

SetConfig sets Config field to given value.


### GetProvider

`func (o *AzureRepos) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AzureRepos) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AzureRepos) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetOrg

`func (o *AzureRepos) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *AzureRepos) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *AzureRepos) SetOrg(v string)`

SetOrg sets Org field to given value.


### GetRepo

`func (o *AzureRepos) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *AzureRepos) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *AzureRepos) SetRepo(v string)`

SetRepo sets Repo field to given value.


### GetDefaultBranch

`func (o *AzureRepos) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *AzureRepos) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *AzureRepos) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.


### GetRepoVersion

`func (o *AzureRepos) GetRepoVersion() string`

GetRepoVersion returns the RepoVersion field if non-nil, zero value otherwise.

### GetRepoVersionOk

`func (o *AzureRepos) GetRepoVersionOk() (*string, bool)`

GetRepoVersionOk returns a tuple with the RepoVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoVersion

`func (o *AzureRepos) SetRepoVersion(v string)`

SetRepoVersion sets RepoVersion field to given value.

### HasRepoVersion

`func (o *AzureRepos) HasRepoVersion() bool`

HasRepoVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


