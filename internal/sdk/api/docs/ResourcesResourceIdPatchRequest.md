# ResourcesResourceIdPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operations** | [**[]ReplaceOperation**](ReplaceOperation.md) | A list of operations to apply to the resource. See the [JSON PATCH specification](https://tools.ietf.org/html/rfc6902) for more details. | 

## Methods

### NewResourcesResourceIdPatchRequest

`func NewResourcesResourceIdPatchRequest(operations []ReplaceOperation, ) *ResourcesResourceIdPatchRequest`

NewResourcesResourceIdPatchRequest instantiates a new ResourcesResourceIdPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcesResourceIdPatchRequestWithDefaults

`func NewResourcesResourceIdPatchRequestWithDefaults() *ResourcesResourceIdPatchRequest`

NewResourcesResourceIdPatchRequestWithDefaults instantiates a new ResourcesResourceIdPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperations

`func (o *ResourcesResourceIdPatchRequest) GetOperations() []ReplaceOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *ResourcesResourceIdPatchRequest) GetOperationsOk() (*[]ReplaceOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *ResourcesResourceIdPatchRequest) SetOperations(v []ReplaceOperation)`

SetOperations sets Operations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


