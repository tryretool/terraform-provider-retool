# SourceControlConfigPost200ResponseData

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

### NewSourceControlConfigPost200ResponseData

`func NewSourceControlConfigPost200ResponseData(config SourceControlConfigGet200ResponseDataAnyOf4Config, provider string, org string, repo string, defaultBranch string, ) *SourceControlConfigPost200ResponseData`

NewSourceControlConfigPost200ResponseData instantiates a new SourceControlConfigPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlConfigPost200ResponseDataWithDefaults

`func NewSourceControlConfigPost200ResponseDataWithDefaults() *SourceControlConfigPost200ResponseData`

NewSourceControlConfigPost200ResponseDataWithDefaults instantiates a new SourceControlConfigPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *SourceControlConfigPost200ResponseData) GetConfig() SourceControlConfigGet200ResponseDataAnyOf4Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SourceControlConfigPost200ResponseData) GetConfigOk() (*SourceControlConfigGet200ResponseDataAnyOf4Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SourceControlConfigPost200ResponseData) SetConfig(v SourceControlConfigGet200ResponseDataAnyOf4Config)`

SetConfig sets Config field to given value.


### GetProvider

`func (o *SourceControlConfigPost200ResponseData) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SourceControlConfigPost200ResponseData) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SourceControlConfigPost200ResponseData) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetOrg

`func (o *SourceControlConfigPost200ResponseData) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *SourceControlConfigPost200ResponseData) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *SourceControlConfigPost200ResponseData) SetOrg(v string)`

SetOrg sets Org field to given value.


### GetRepo

`func (o *SourceControlConfigPost200ResponseData) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *SourceControlConfigPost200ResponseData) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *SourceControlConfigPost200ResponseData) SetRepo(v string)`

SetRepo sets Repo field to given value.


### GetDefaultBranch

`func (o *SourceControlConfigPost200ResponseData) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *SourceControlConfigPost200ResponseData) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *SourceControlConfigPost200ResponseData) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.


### GetRepoVersion

`func (o *SourceControlConfigPost200ResponseData) GetRepoVersion() string`

GetRepoVersion returns the RepoVersion field if non-nil, zero value otherwise.

### GetRepoVersionOk

`func (o *SourceControlConfigPost200ResponseData) GetRepoVersionOk() (*string, bool)`

GetRepoVersionOk returns a tuple with the RepoVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoVersion

`func (o *SourceControlConfigPost200ResponseData) SetRepoVersion(v string)`

SetRepoVersion sets RepoVersion field to given value.

### HasRepoVersion

`func (o *SourceControlConfigPost200ResponseData) HasRepoVersion() bool`

HasRepoVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


