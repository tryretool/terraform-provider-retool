# WorkflowRunIdGet200Response

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

### NewWorkflowRunIdGet200Response

`func NewWorkflowRunIdGet200Response(status string, workflowId string, id string, triggerType string, triggerId string, createdAt string, ) *WorkflowRunIdGet200Response`

NewWorkflowRunIdGet200Response instantiates a new WorkflowRunIdGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowRunIdGet200ResponseWithDefaults

`func NewWorkflowRunIdGet200ResponseWithDefaults() *WorkflowRunIdGet200Response`

NewWorkflowRunIdGet200ResponseWithDefaults instantiates a new WorkflowRunIdGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *WorkflowRunIdGet200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowRunIdGet200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowRunIdGet200Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetWorkflowId

`func (o *WorkflowRunIdGet200Response) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowRunIdGet200Response) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowRunIdGet200Response) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetId

`func (o *WorkflowRunIdGet200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowRunIdGet200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowRunIdGet200Response) SetId(v string)`

SetId sets Id field to given value.


### GetTriggerType

`func (o *WorkflowRunIdGet200Response) GetTriggerType() string`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *WorkflowRunIdGet200Response) GetTriggerTypeOk() (*string, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *WorkflowRunIdGet200Response) SetTriggerType(v string)`

SetTriggerType sets TriggerType field to given value.


### GetTriggerId

`func (o *WorkflowRunIdGet200Response) GetTriggerId() string`

GetTriggerId returns the TriggerId field if non-nil, zero value otherwise.

### GetTriggerIdOk

`func (o *WorkflowRunIdGet200Response) GetTriggerIdOk() (*string, bool)`

GetTriggerIdOk returns a tuple with the TriggerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerId

`func (o *WorkflowRunIdGet200Response) SetTriggerId(v string)`

SetTriggerId sets TriggerId field to given value.


### GetCreatedAt

`func (o *WorkflowRunIdGet200Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkflowRunIdGet200Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkflowRunIdGet200Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUserTasks

`func (o *WorkflowRunIdGet200Response) GetUserTasks() []UserTasksGet200ResponseDataInner`

GetUserTasks returns the UserTasks field if non-nil, zero value otherwise.

### GetUserTasksOk

`func (o *WorkflowRunIdGet200Response) GetUserTasksOk() (*[]UserTasksGet200ResponseDataInner, bool)`

GetUserTasksOk returns a tuple with the UserTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTasks

`func (o *WorkflowRunIdGet200Response) SetUserTasks(v []UserTasksGet200ResponseDataInner)`

SetUserTasks sets UserTasks field to given value.

### HasUserTasks

`func (o *WorkflowRunIdGet200Response) HasUserTasks() bool`

HasUserTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


