# UserTasksTaskIdCancelPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Email of user cancelling the task. | 
**Output** | Pointer to **map[string]interface{}** | Assignee-defined output required for a task to be cancelled. | [optional] 

## Methods

### NewUserTasksTaskIdCancelPatchRequest

`func NewUserTasksTaskIdCancelPatchRequest(email string, ) *UserTasksTaskIdCancelPatchRequest`

NewUserTasksTaskIdCancelPatchRequest instantiates a new UserTasksTaskIdCancelPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserTasksTaskIdCancelPatchRequestWithDefaults

`func NewUserTasksTaskIdCancelPatchRequestWithDefaults() *UserTasksTaskIdCancelPatchRequest`

NewUserTasksTaskIdCancelPatchRequestWithDefaults instantiates a new UserTasksTaskIdCancelPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserTasksTaskIdCancelPatchRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserTasksTaskIdCancelPatchRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserTasksTaskIdCancelPatchRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetOutput

`func (o *UserTasksTaskIdCancelPatchRequest) GetOutput() map[string]interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *UserTasksTaskIdCancelPatchRequest) GetOutputOk() (*map[string]interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *UserTasksTaskIdCancelPatchRequest) SetOutput(v map[string]interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *UserTasksTaskIdCancelPatchRequest) HasOutput() bool`

HasOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


