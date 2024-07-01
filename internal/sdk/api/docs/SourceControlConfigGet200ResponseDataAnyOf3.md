# SourceControlConfigGet200ResponseDataAnyOf3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**SourceControlConfigGet200ResponseDataAnyOf3Config**](SourceControlConfigGet200ResponseDataAnyOf3Config.md) |  | 
**Provider** | **string** |  | 
**Org** | **string** | The user or organization to which the repository belongs to. | 
**Repo** | **string** | The name of the repository you created to use with Retool. | 
**DefaultBranch** | **string** | The default branch, e.g., main. | 
**RepoVersion** | Pointer to **string** | Repositories using Toolscript are 2.0.0. Repositories using legacy YAML are 1.0.0. | [optional] 

## Methods

### NewSourceControlConfigGet200ResponseDataAnyOf3

`func NewSourceControlConfigGet200ResponseDataAnyOf3(config SourceControlConfigGet200ResponseDataAnyOf3Config, provider string, org string, repo string, defaultBranch string, ) *SourceControlConfigGet200ResponseDataAnyOf3`

NewSourceControlConfigGet200ResponseDataAnyOf3 instantiates a new SourceControlConfigGet200ResponseDataAnyOf3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlConfigGet200ResponseDataAnyOf3WithDefaults

`func NewSourceControlConfigGet200ResponseDataAnyOf3WithDefaults() *SourceControlConfigGet200ResponseDataAnyOf3`

NewSourceControlConfigGet200ResponseDataAnyOf3WithDefaults instantiates a new SourceControlConfigGet200ResponseDataAnyOf3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *SourceControlConfigGet200ResponseDataAnyOf3) GetConfig() SourceControlConfigGet200ResponseDataAnyOf3Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SourceControlConfigGet200ResponseDataAnyOf3) GetConfigOk() (*SourceControlConfigGet200ResponseDataAnyOf3Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SourceControlConfigGet200ResponseDataAnyOf3) SetConfig(v SourceControlConfigGet200ResponseDataAnyOf3Config)`

SetConfig sets Config field to given value.


### GetProvider

`func (o *SourceControlConfigGet200ResponseDataAnyOf3) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SourceControlConfigGet200ResponseDataAnyOf3) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SourceControlConfigGet200ResponseDataAnyOf3) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetOrg

`func (o *SourceControlConfigGet200ResponseDataAnyOf3) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *SourceControlConfigGet200ResponseDataAnyOf3) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *SourceControlConfigGet200ResponseDataAnyOf3) SetOrg(v string)`

SetOrg sets Org field to given value.


### GetRepo

`func (o *SourceControlConfigGet200ResponseDataAnyOf3) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *SourceControlConfigGet200ResponseDataAnyOf3) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *SourceControlConfigGet200ResponseDataAnyOf3) SetRepo(v string)`

SetRepo sets Repo field to given value.


### GetDefaultBranch

`func (o *SourceControlConfigGet200ResponseDataAnyOf3) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *SourceControlConfigGet200ResponseDataAnyOf3) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *SourceControlConfigGet200ResponseDataAnyOf3) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.


### GetRepoVersion

`func (o *SourceControlConfigGet200ResponseDataAnyOf3) GetRepoVersion() string`

GetRepoVersion returns the RepoVersion field if non-nil, zero value otherwise.

### GetRepoVersionOk

`func (o *SourceControlConfigGet200ResponseDataAnyOf3) GetRepoVersionOk() (*string, bool)`

GetRepoVersionOk returns a tuple with the RepoVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoVersion

`func (o *SourceControlConfigGet200ResponseDataAnyOf3) SetRepoVersion(v string)`

SetRepoVersion sets RepoVersion field to given value.

### HasRepoVersion

`func (o *SourceControlConfigGet200ResponseDataAnyOf3) HasRepoVersion() bool`

HasRepoVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


