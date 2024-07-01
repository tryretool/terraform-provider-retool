# SourceControlConfigGet200ResponseDataAnyOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**SourceControlConfigGet200ResponseDataAnyOf1Config**](SourceControlConfigGet200ResponseDataAnyOf1Config.md) |  | 
**Provider** | **string** |  | 
**Org** | **string** | The user or organization to which the repository belongs to. | 
**Repo** | **string** | The name of the repository you created to use with Retool. | 
**DefaultBranch** | **string** | The default branch, e.g., main. | 
**RepoVersion** | Pointer to **string** | Repositories using Toolscript are 2.0.0. Repositories using legacy YAML are 1.0.0. | [optional] 

## Methods

### NewSourceControlConfigGet200ResponseDataAnyOf1

`func NewSourceControlConfigGet200ResponseDataAnyOf1(config SourceControlConfigGet200ResponseDataAnyOf1Config, provider string, org string, repo string, defaultBranch string, ) *SourceControlConfigGet200ResponseDataAnyOf1`

NewSourceControlConfigGet200ResponseDataAnyOf1 instantiates a new SourceControlConfigGet200ResponseDataAnyOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlConfigGet200ResponseDataAnyOf1WithDefaults

`func NewSourceControlConfigGet200ResponseDataAnyOf1WithDefaults() *SourceControlConfigGet200ResponseDataAnyOf1`

NewSourceControlConfigGet200ResponseDataAnyOf1WithDefaults instantiates a new SourceControlConfigGet200ResponseDataAnyOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *SourceControlConfigGet200ResponseDataAnyOf1) GetConfig() SourceControlConfigGet200ResponseDataAnyOf1Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SourceControlConfigGet200ResponseDataAnyOf1) GetConfigOk() (*SourceControlConfigGet200ResponseDataAnyOf1Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SourceControlConfigGet200ResponseDataAnyOf1) SetConfig(v SourceControlConfigGet200ResponseDataAnyOf1Config)`

SetConfig sets Config field to given value.


### GetProvider

`func (o *SourceControlConfigGet200ResponseDataAnyOf1) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SourceControlConfigGet200ResponseDataAnyOf1) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SourceControlConfigGet200ResponseDataAnyOf1) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetOrg

`func (o *SourceControlConfigGet200ResponseDataAnyOf1) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *SourceControlConfigGet200ResponseDataAnyOf1) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *SourceControlConfigGet200ResponseDataAnyOf1) SetOrg(v string)`

SetOrg sets Org field to given value.


### GetRepo

`func (o *SourceControlConfigGet200ResponseDataAnyOf1) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *SourceControlConfigGet200ResponseDataAnyOf1) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *SourceControlConfigGet200ResponseDataAnyOf1) SetRepo(v string)`

SetRepo sets Repo field to given value.


### GetDefaultBranch

`func (o *SourceControlConfigGet200ResponseDataAnyOf1) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *SourceControlConfigGet200ResponseDataAnyOf1) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *SourceControlConfigGet200ResponseDataAnyOf1) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.


### GetRepoVersion

`func (o *SourceControlConfigGet200ResponseDataAnyOf1) GetRepoVersion() string`

GetRepoVersion returns the RepoVersion field if non-nil, zero value otherwise.

### GetRepoVersionOk

`func (o *SourceControlConfigGet200ResponseDataAnyOf1) GetRepoVersionOk() (*string, bool)`

GetRepoVersionOk returns a tuple with the RepoVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoVersion

`func (o *SourceControlConfigGet200ResponseDataAnyOf1) SetRepoVersion(v string)`

SetRepoVersion sets RepoVersion field to given value.

### HasRepoVersion

`func (o *SourceControlConfigGet200ResponseDataAnyOf1) HasRepoVersion() bool`

HasRepoVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


