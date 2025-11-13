# UserTasksTaskIdSubmitPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Email of user completing the task. | 
**Output** | **map[string]interface{}** | Assignee-defined output required for a task to be submitted. | 

## Methods

### NewUserTasksTaskIdSubmitPatchRequest

`func NewUserTasksTaskIdSubmitPatchRequest(email string, output map[string]interface{}, ) *UserTasksTaskIdSubmitPatchRequest`

NewUserTasksTaskIdSubmitPatchRequest instantiates a new UserTasksTaskIdSubmitPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserTasksTaskIdSubmitPatchRequestWithDefaults

`func NewUserTasksTaskIdSubmitPatchRequestWithDefaults() *UserTasksTaskIdSubmitPatchRequest`

NewUserTasksTaskIdSubmitPatchRequestWithDefaults instantiates a new UserTasksTaskIdSubmitPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserTasksTaskIdSubmitPatchRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserTasksTaskIdSubmitPatchRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserTasksTaskIdSubmitPatchRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetOutput

`func (o *UserTasksTaskIdSubmitPatchRequest) GetOutput() map[string]interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *UserTasksTaskIdSubmitPatchRequest) GetOutputOk() (*map[string]interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *UserTasksTaskIdSubmitPatchRequest) SetOutput(v map[string]interface{})`

SetOutput sets Output field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


