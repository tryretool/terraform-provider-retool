# SourceControlTestDeployPostRequestDeployParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitSha** | **string** | The full commit SHA to dry deploy | 
**IsFullSync** | Pointer to **bool** |  | [optional] 

## Methods

### NewSourceControlTestDeployPostRequestDeployParams

`func NewSourceControlTestDeployPostRequestDeployParams(commitSha string, ) *SourceControlTestDeployPostRequestDeployParams`

NewSourceControlTestDeployPostRequestDeployParams instantiates a new SourceControlTestDeployPostRequestDeployParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlTestDeployPostRequestDeployParamsWithDefaults

`func NewSourceControlTestDeployPostRequestDeployParamsWithDefaults() *SourceControlTestDeployPostRequestDeployParams`

NewSourceControlTestDeployPostRequestDeployParamsWithDefaults instantiates a new SourceControlTestDeployPostRequestDeployParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitSha

`func (o *SourceControlTestDeployPostRequestDeployParams) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *SourceControlTestDeployPostRequestDeployParams) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *SourceControlTestDeployPostRequestDeployParams) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.


### GetIsFullSync

`func (o *SourceControlTestDeployPostRequestDeployParams) GetIsFullSync() bool`

GetIsFullSync returns the IsFullSync field if non-nil, zero value otherwise.

### GetIsFullSyncOk

`func (o *SourceControlTestDeployPostRequestDeployParams) GetIsFullSyncOk() (*bool, bool)`

GetIsFullSyncOk returns a tuple with the IsFullSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFullSync

`func (o *SourceControlTestDeployPostRequestDeployParams) SetIsFullSync(v bool)`

SetIsFullSync sets IsFullSync field to given value.

### HasIsFullSync

`func (o *SourceControlTestDeployPostRequestDeployParams) HasIsFullSync() bool`

HasIsFullSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


