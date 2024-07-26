# GroupsGroupIdPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operations** | [**[]GroupsGroupIdPatchRequestOperationsInner**](GroupsGroupIdPatchRequestOperationsInner.md) | A list of operations to apply to the group. See the [JSON PATCH specification](https://tools.ietf.org/html/rfc6902) for more details. | 

## Methods

### NewGroupsGroupIdPatchRequest

`func NewGroupsGroupIdPatchRequest(operations []GroupsGroupIdPatchRequestOperationsInner, ) *GroupsGroupIdPatchRequest`

NewGroupsGroupIdPatchRequest instantiates a new GroupsGroupIdPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupsGroupIdPatchRequestWithDefaults

`func NewGroupsGroupIdPatchRequestWithDefaults() *GroupsGroupIdPatchRequest`

NewGroupsGroupIdPatchRequestWithDefaults instantiates a new GroupsGroupIdPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperations

`func (o *GroupsGroupIdPatchRequest) GetOperations() []GroupsGroupIdPatchRequestOperationsInner`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *GroupsGroupIdPatchRequest) GetOperationsOk() (*[]GroupsGroupIdPatchRequestOperationsInner, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *GroupsGroupIdPatchRequest) SetOperations(v []GroupsGroupIdPatchRequestOperationsInner)`

SetOperations sets Operations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


