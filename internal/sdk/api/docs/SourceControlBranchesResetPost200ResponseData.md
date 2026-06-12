# SourceControlBranchesResetPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchName** | **string** | The Retool app branch that was reset. | 
**RemoteBranchName** | **string** | The remote Git branch used as the reset source. | 
**Status** | **string** | The reset result. | 
**UpdatedApps** | [**[]SourceControlBranchesResetPost200ResponseDataUpdatedAppsInner**](SourceControlBranchesResetPost200ResponseDataUpdatedAppsInner.md) | Apps that were reset from the remote branch. | 

## Methods

### NewSourceControlBranchesResetPost200ResponseData

`func NewSourceControlBranchesResetPost200ResponseData(branchName string, remoteBranchName string, status string, updatedApps []SourceControlBranchesResetPost200ResponseDataUpdatedAppsInner, ) *SourceControlBranchesResetPost200ResponseData`

NewSourceControlBranchesResetPost200ResponseData instantiates a new SourceControlBranchesResetPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlBranchesResetPost200ResponseDataWithDefaults

`func NewSourceControlBranchesResetPost200ResponseDataWithDefaults() *SourceControlBranchesResetPost200ResponseData`

NewSourceControlBranchesResetPost200ResponseDataWithDefaults instantiates a new SourceControlBranchesResetPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranchName

`func (o *SourceControlBranchesResetPost200ResponseData) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *SourceControlBranchesResetPost200ResponseData) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *SourceControlBranchesResetPost200ResponseData) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.


### GetRemoteBranchName

`func (o *SourceControlBranchesResetPost200ResponseData) GetRemoteBranchName() string`

GetRemoteBranchName returns the RemoteBranchName field if non-nil, zero value otherwise.

### GetRemoteBranchNameOk

`func (o *SourceControlBranchesResetPost200ResponseData) GetRemoteBranchNameOk() (*string, bool)`

GetRemoteBranchNameOk returns a tuple with the RemoteBranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteBranchName

`func (o *SourceControlBranchesResetPost200ResponseData) SetRemoteBranchName(v string)`

SetRemoteBranchName sets RemoteBranchName field to given value.


### GetStatus

`func (o *SourceControlBranchesResetPost200ResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SourceControlBranchesResetPost200ResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SourceControlBranchesResetPost200ResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedApps

`func (o *SourceControlBranchesResetPost200ResponseData) GetUpdatedApps() []SourceControlBranchesResetPost200ResponseDataUpdatedAppsInner`

GetUpdatedApps returns the UpdatedApps field if non-nil, zero value otherwise.

### GetUpdatedAppsOk

`func (o *SourceControlBranchesResetPost200ResponseData) GetUpdatedAppsOk() (*[]SourceControlBranchesResetPost200ResponseDataUpdatedAppsInner, bool)`

GetUpdatedAppsOk returns a tuple with the UpdatedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedApps

`func (o *SourceControlBranchesResetPost200ResponseData) SetUpdatedApps(v []SourceControlBranchesResetPost200ResponseDataUpdatedAppsInner)`

SetUpdatedApps sets UpdatedApps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


