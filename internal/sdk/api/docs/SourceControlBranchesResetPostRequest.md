# SourceControlBranchesResetPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchName** | **string** | The Retool app branch to reset. | 
**RemoteBranchName** | Pointer to **string** | The remote Git branch to reset from. Defaults to branch_name when omitted. | [optional] 

## Methods

### NewSourceControlBranchesResetPostRequest

`func NewSourceControlBranchesResetPostRequest(branchName string, ) *SourceControlBranchesResetPostRequest`

NewSourceControlBranchesResetPostRequest instantiates a new SourceControlBranchesResetPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlBranchesResetPostRequestWithDefaults

`func NewSourceControlBranchesResetPostRequestWithDefaults() *SourceControlBranchesResetPostRequest`

NewSourceControlBranchesResetPostRequestWithDefaults instantiates a new SourceControlBranchesResetPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranchName

`func (o *SourceControlBranchesResetPostRequest) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *SourceControlBranchesResetPostRequest) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *SourceControlBranchesResetPostRequest) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.


### GetRemoteBranchName

`func (o *SourceControlBranchesResetPostRequest) GetRemoteBranchName() string`

GetRemoteBranchName returns the RemoteBranchName field if non-nil, zero value otherwise.

### GetRemoteBranchNameOk

`func (o *SourceControlBranchesResetPostRequest) GetRemoteBranchNameOk() (*string, bool)`

GetRemoteBranchNameOk returns a tuple with the RemoteBranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteBranchName

`func (o *SourceControlBranchesResetPostRequest) SetRemoteBranchName(v string)`

SetRemoteBranchName sets RemoteBranchName field to given value.

### HasRemoteBranchName

`func (o *SourceControlBranchesResetPostRequest) HasRemoteBranchName() bool`

HasRemoteBranchName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


