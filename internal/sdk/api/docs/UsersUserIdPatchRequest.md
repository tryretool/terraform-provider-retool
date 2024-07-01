# UsersUserIdPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operations** | [**[]UsersUserIdPatchRequestOperationsInner**](UsersUserIdPatchRequestOperationsInner.md) | The body of a request to update a user. | 

## Methods

### NewUsersUserIdPatchRequest

`func NewUsersUserIdPatchRequest(operations []UsersUserIdPatchRequestOperationsInner, ) *UsersUserIdPatchRequest`

NewUsersUserIdPatchRequest instantiates a new UsersUserIdPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersUserIdPatchRequestWithDefaults

`func NewUsersUserIdPatchRequestWithDefaults() *UsersUserIdPatchRequest`

NewUsersUserIdPatchRequestWithDefaults instantiates a new UsersUserIdPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperations

`func (o *UsersUserIdPatchRequest) GetOperations() []UsersUserIdPatchRequestOperationsInner`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *UsersUserIdPatchRequest) GetOperationsOk() (*[]UsersUserIdPatchRequestOperationsInner, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *UsersUserIdPatchRequest) SetOperations(v []UsersUserIdPatchRequestOperationsInner)`

SetOperations sets Operations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


