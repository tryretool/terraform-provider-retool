# SourceControlBranchesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchName** | **string** | The Retool branch to create. | 
**ElementId** | Pointer to **string** | Optional ID of a single Retool element to add to the branch. | [optional] 
**ElementType** | Pointer to **string** | Required when element_id is provided. The classic Retool element type to add to the branch. Supported values are PAGE, PLAYGROUND_QUERY, WORKFLOW, PROCESS, and AGENT. | [optional] 
**Shared** | Pointer to **bool** | Whether the branch should be shared with other users. Defaults to false. | [optional] 

## Methods

### NewSourceControlBranchesPostRequest

`func NewSourceControlBranchesPostRequest(branchName string, ) *SourceControlBranchesPostRequest`

NewSourceControlBranchesPostRequest instantiates a new SourceControlBranchesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlBranchesPostRequestWithDefaults

`func NewSourceControlBranchesPostRequestWithDefaults() *SourceControlBranchesPostRequest`

NewSourceControlBranchesPostRequestWithDefaults instantiates a new SourceControlBranchesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranchName

`func (o *SourceControlBranchesPostRequest) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *SourceControlBranchesPostRequest) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *SourceControlBranchesPostRequest) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.


### GetElementId

`func (o *SourceControlBranchesPostRequest) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *SourceControlBranchesPostRequest) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *SourceControlBranchesPostRequest) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *SourceControlBranchesPostRequest) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetElementType

`func (o *SourceControlBranchesPostRequest) GetElementType() string`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *SourceControlBranchesPostRequest) GetElementTypeOk() (*string, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *SourceControlBranchesPostRequest) SetElementType(v string)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *SourceControlBranchesPostRequest) HasElementType() bool`

HasElementType returns a boolean if a field has been set.

### GetShared

`func (o *SourceControlBranchesPostRequest) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *SourceControlBranchesPostRequest) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *SourceControlBranchesPostRequest) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *SourceControlBranchesPostRequest) HasShared() bool`

HasShared returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


