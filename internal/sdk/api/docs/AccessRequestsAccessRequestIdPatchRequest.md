# AccessRequestsAccessRequestIdPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operations** | [**[]UsersUserIdPatchRequestOperationsInnerAnyOf2**](UsersUserIdPatchRequestOperationsInnerAnyOf2.md) | The body of a request to update an access request&#39;s status. | 

## Methods

### NewAccessRequestsAccessRequestIdPatchRequest

`func NewAccessRequestsAccessRequestIdPatchRequest(operations []UsersUserIdPatchRequestOperationsInnerAnyOf2, ) *AccessRequestsAccessRequestIdPatchRequest`

NewAccessRequestsAccessRequestIdPatchRequest instantiates a new AccessRequestsAccessRequestIdPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessRequestsAccessRequestIdPatchRequestWithDefaults

`func NewAccessRequestsAccessRequestIdPatchRequestWithDefaults() *AccessRequestsAccessRequestIdPatchRequest`

NewAccessRequestsAccessRequestIdPatchRequestWithDefaults instantiates a new AccessRequestsAccessRequestIdPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperations

`func (o *AccessRequestsAccessRequestIdPatchRequest) GetOperations() []UsersUserIdPatchRequestOperationsInnerAnyOf2`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *AccessRequestsAccessRequestIdPatchRequest) GetOperationsOk() (*[]UsersUserIdPatchRequestOperationsInnerAnyOf2, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *AccessRequestsAccessRequestIdPatchRequest) SetOperations(v []UsersUserIdPatchRequestOperationsInnerAnyOf2)`

SetOperations sets Operations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


