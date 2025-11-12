# WorkflowRunWithUserTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The status of the workflow run. | 
**WorkflowId** | **string** | The ID of the workflow. | 
**Id** | **string** | The ID of the workflow run. | 
**TriggerType** | **string** | The type of trigger that started the workflow. | 
**TriggerId** | **string** | The ID of the trigger. | 
**CreatedAt** | **string** | The creation date time of the workflow run. | 
**UserTasks** | Pointer to [**[]UserTasksGet200ResponseDataInner**](UserTasksGet200ResponseDataInner.md) | List of user tasks of the workfow run. | [optional] 

## Methods

### NewWorkflowRunWithUserTask

`func NewWorkflowRunWithUserTask(status string, workflowId string, id string, triggerType string, triggerId string, createdAt string, ) *WorkflowRunWithUserTask`

NewWorkflowRunWithUserTask instantiates a new WorkflowRunWithUserTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowRunWithUserTaskWithDefaults

`func NewWorkflowRunWithUserTaskWithDefaults() *WorkflowRunWithUserTask`

NewWorkflowRunWithUserTaskWithDefaults instantiates a new WorkflowRunWithUserTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *WorkflowRunWithUserTask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowRunWithUserTask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowRunWithUserTask) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetWorkflowId

`func (o *WorkflowRunWithUserTask) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowRunWithUserTask) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowRunWithUserTask) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetId

`func (o *WorkflowRunWithUserTask) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowRunWithUserTask) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowRunWithUserTask) SetId(v string)`

SetId sets Id field to given value.


### GetTriggerType

`func (o *WorkflowRunWithUserTask) GetTriggerType() string`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *WorkflowRunWithUserTask) GetTriggerTypeOk() (*string, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *WorkflowRunWithUserTask) SetTriggerType(v string)`

SetTriggerType sets TriggerType field to given value.


### GetTriggerId

`func (o *WorkflowRunWithUserTask) GetTriggerId() string`

GetTriggerId returns the TriggerId field if non-nil, zero value otherwise.

### GetTriggerIdOk

`func (o *WorkflowRunWithUserTask) GetTriggerIdOk() (*string, bool)`

GetTriggerIdOk returns a tuple with the TriggerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerId

`func (o *WorkflowRunWithUserTask) SetTriggerId(v string)`

SetTriggerId sets TriggerId field to given value.


### GetCreatedAt

`func (o *WorkflowRunWithUserTask) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkflowRunWithUserTask) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkflowRunWithUserTask) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUserTasks

`func (o *WorkflowRunWithUserTask) GetUserTasks() []UserTasksGet200ResponseDataInner`

GetUserTasks returns the UserTasks field if non-nil, zero value otherwise.

### GetUserTasksOk

`func (o *WorkflowRunWithUserTask) GetUserTasksOk() (*[]UserTasksGet200ResponseDataInner, bool)`

GetUserTasksOk returns a tuple with the UserTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTasks

`func (o *WorkflowRunWithUserTask) SetUserTasks(v []UserTasksGet200ResponseDataInner)`

SetUserTasks sets UserTasks field to given value.

### HasUserTasks

`func (o *WorkflowRunWithUserTask) HasUserTasks() bool`

HasUserTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


