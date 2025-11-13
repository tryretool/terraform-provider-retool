# UserAttributesIdPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operations** | [**[]ReplaceOperation**](ReplaceOperation.md) | A list of operations to apply to the user attribute. See the [JSON PATCH specification](https://tools.ietf.org/html/rfc6902) for more details. | 
**UpdateExisting** | Pointer to **bool** | Whether to update existing users with the deleted attribute | [optional] 

## Methods

### NewUserAttributesIdPatchRequest

`func NewUserAttributesIdPatchRequest(operations []ReplaceOperation, ) *UserAttributesIdPatchRequest`

NewUserAttributesIdPatchRequest instantiates a new UserAttributesIdPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAttributesIdPatchRequestWithDefaults

`func NewUserAttributesIdPatchRequestWithDefaults() *UserAttributesIdPatchRequest`

NewUserAttributesIdPatchRequestWithDefaults instantiates a new UserAttributesIdPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperations

`func (o *UserAttributesIdPatchRequest) GetOperations() []ReplaceOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *UserAttributesIdPatchRequest) GetOperationsOk() (*[]ReplaceOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *UserAttributesIdPatchRequest) SetOperations(v []ReplaceOperation)`

SetOperations sets Operations field to given value.


### GetUpdateExisting

`func (o *UserAttributesIdPatchRequest) GetUpdateExisting() bool`

GetUpdateExisting returns the UpdateExisting field if non-nil, zero value otherwise.

### GetUpdateExistingOk

`func (o *UserAttributesIdPatchRequest) GetUpdateExistingOk() (*bool, bool)`

GetUpdateExistingOk returns a tuple with the UpdateExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateExisting

`func (o *UserAttributesIdPatchRequest) SetUpdateExisting(v bool)`

SetUpdateExisting sets UpdateExisting field to given value.

### HasUpdateExisting

`func (o *UserAttributesIdPatchRequest) HasUpdateExisting() bool`

HasUpdateExisting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


