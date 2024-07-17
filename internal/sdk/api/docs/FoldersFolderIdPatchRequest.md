# FoldersFolderIdPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operations** | [**[]FoldersFolderIdPatchRequestOperationsInner**](FoldersFolderIdPatchRequestOperationsInner.md) | A list of operations to apply to the folder. See https://tools.ietf.org/html/rfc6902 for details. | 

## Methods

### NewFoldersFolderIdPatchRequest

`func NewFoldersFolderIdPatchRequest(operations []FoldersFolderIdPatchRequestOperationsInner, ) *FoldersFolderIdPatchRequest`

NewFoldersFolderIdPatchRequest instantiates a new FoldersFolderIdPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFoldersFolderIdPatchRequestWithDefaults

`func NewFoldersFolderIdPatchRequestWithDefaults() *FoldersFolderIdPatchRequest`

NewFoldersFolderIdPatchRequestWithDefaults instantiates a new FoldersFolderIdPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperations

`func (o *FoldersFolderIdPatchRequest) GetOperations() []FoldersFolderIdPatchRequestOperationsInner`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *FoldersFolderIdPatchRequest) GetOperationsOk() (*[]FoldersFolderIdPatchRequestOperationsInner, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *FoldersFolderIdPatchRequest) SetOperations(v []FoldersFolderIdPatchRequestOperationsInner)`

SetOperations sets Operations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


