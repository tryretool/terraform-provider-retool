# SourceControlBranchesPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchId** | **string** | The created Retool branch ID. | 
**BranchName** | **string** | The created Retool branch name. | 
**Shared** | **bool** | Whether the branch is shared with other users. | 
**Element** | [**NullableSourceControlBranchesPost200ResponseDataElement**](SourceControlBranchesPost200ResponseDataElement.md) |  | 

## Methods

### NewSourceControlBranchesPost200ResponseData

`func NewSourceControlBranchesPost200ResponseData(branchId string, branchName string, shared bool, element NullableSourceControlBranchesPost200ResponseDataElement, ) *SourceControlBranchesPost200ResponseData`

NewSourceControlBranchesPost200ResponseData instantiates a new SourceControlBranchesPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlBranchesPost200ResponseDataWithDefaults

`func NewSourceControlBranchesPost200ResponseDataWithDefaults() *SourceControlBranchesPost200ResponseData`

NewSourceControlBranchesPost200ResponseDataWithDefaults instantiates a new SourceControlBranchesPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranchId

`func (o *SourceControlBranchesPost200ResponseData) GetBranchId() string`

GetBranchId returns the BranchId field if non-nil, zero value otherwise.

### GetBranchIdOk

`func (o *SourceControlBranchesPost200ResponseData) GetBranchIdOk() (*string, bool)`

GetBranchIdOk returns a tuple with the BranchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchId

`func (o *SourceControlBranchesPost200ResponseData) SetBranchId(v string)`

SetBranchId sets BranchId field to given value.


### GetBranchName

`func (o *SourceControlBranchesPost200ResponseData) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *SourceControlBranchesPost200ResponseData) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *SourceControlBranchesPost200ResponseData) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.


### GetShared

`func (o *SourceControlBranchesPost200ResponseData) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *SourceControlBranchesPost200ResponseData) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *SourceControlBranchesPost200ResponseData) SetShared(v bool)`

SetShared sets Shared field to given value.


### GetElement

`func (o *SourceControlBranchesPost200ResponseData) GetElement() SourceControlBranchesPost200ResponseDataElement`

GetElement returns the Element field if non-nil, zero value otherwise.

### GetElementOk

`func (o *SourceControlBranchesPost200ResponseData) GetElementOk() (*SourceControlBranchesPost200ResponseDataElement, bool)`

GetElementOk returns a tuple with the Element field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElement

`func (o *SourceControlBranchesPost200ResponseData) SetElement(v SourceControlBranchesPost200ResponseDataElement)`

SetElement sets Element field to given value.


### SetElementNil

`func (o *SourceControlBranchesPost200ResponseData) SetElementNil(b bool)`

 SetElementNil sets the value for Element to be an explicit nil

### UnsetElement
`func (o *SourceControlBranchesPost200ResponseData) UnsetElement()`

UnsetElement ensures that no value is present for Element, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


