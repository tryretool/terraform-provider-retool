# SourceControlConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**SourceControlConfigGet200ResponseDataAnyOf4Config**](SourceControlConfigGet200ResponseDataAnyOf4Config.md) |  | 
**Provider** | **string** |  | 
**Org** | **string** | The user or organization to which the repository belongs to. | 
**Repo** | **string** | The name of the repository you created to use with Retool. | 
**DefaultBranch** | **string** | The default branch, e.g., main. | 
**RepoVersion** | Pointer to **string** | Repositories using Toolscript are 2.0.0. Repositories using legacy YAML are 1.0.0. | [optional] 

## Methods

### NewSourceControlConfiguration

`func NewSourceControlConfiguration(config SourceControlConfigGet200ResponseDataAnyOf4Config, provider string, org string, repo string, defaultBranch string, ) *SourceControlConfiguration`

NewSourceControlConfiguration instantiates a new SourceControlConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlConfigurationWithDefaults

`func NewSourceControlConfigurationWithDefaults() *SourceControlConfiguration`

NewSourceControlConfigurationWithDefaults instantiates a new SourceControlConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *SourceControlConfiguration) GetConfig() SourceControlConfigGet200ResponseDataAnyOf4Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SourceControlConfiguration) GetConfigOk() (*SourceControlConfigGet200ResponseDataAnyOf4Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SourceControlConfiguration) SetConfig(v SourceControlConfigGet200ResponseDataAnyOf4Config)`

SetConfig sets Config field to given value.


### GetProvider

`func (o *SourceControlConfiguration) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SourceControlConfiguration) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SourceControlConfiguration) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetOrg

`func (o *SourceControlConfiguration) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *SourceControlConfiguration) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *SourceControlConfiguration) SetOrg(v string)`

SetOrg sets Org field to given value.


### GetRepo

`func (o *SourceControlConfiguration) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *SourceControlConfiguration) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *SourceControlConfiguration) SetRepo(v string)`

SetRepo sets Repo field to given value.


### GetDefaultBranch

`func (o *SourceControlConfiguration) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *SourceControlConfiguration) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *SourceControlConfiguration) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.


### GetRepoVersion

`func (o *SourceControlConfiguration) GetRepoVersion() string`

GetRepoVersion returns the RepoVersion field if non-nil, zero value otherwise.

### GetRepoVersionOk

`func (o *SourceControlConfiguration) GetRepoVersionOk() (*string, bool)`

GetRepoVersionOk returns a tuple with the RepoVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoVersion

`func (o *SourceControlConfiguration) SetRepoVersion(v string)`

SetRepoVersion sets RepoVersion field to given value.

### HasRepoVersion

`func (o *SourceControlConfiguration) HasRepoVersion() bool`

HasRepoVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


