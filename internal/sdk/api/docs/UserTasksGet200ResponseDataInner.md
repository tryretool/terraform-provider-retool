# UserTasksGet200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the user task. | 
**WorkflowId** | **string** | The ID of the workflow this task belongs to. | 
**WorkflowName** | **string** | The name of the workflow. | 
**WorkflowRunId** | **string** | The ID of the workflow run. | 
**TaskName** | **string** | The name of the task. | 
**Status** | **string** | The current status of the task. | 
**Context** | **map[string]interface{}** | Additional context for the task. | 
**CreatedAt** | **string** | The creation date time of the task. | 
**CompletedAt** | **NullableString** | The completion date time of the task. | 
**TaskUrl** | **[]string** | URL links to apps associated with the task. | 
**Assignees** | [**[]UserTasksGet200ResponseDataInnerAssigneesInner**](UserTasksGet200ResponseDataInnerAssigneesInner.md) | The groups or users assigned to this task. | 
**WorkflowRelease** | **NullableString** | The workflow release version, if any. | 
**ExpiresAt** | **NullableString** | The expiration date time of the task, if set. | 
**Output** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUserTasksGet200ResponseDataInner

`func NewUserTasksGet200ResponseDataInner(id string, workflowId string, workflowName string, workflowRunId string, taskName string, status string, context map[string]interface{}, createdAt string, completedAt NullableString, taskUrl []string, assignees []UserTasksGet200ResponseDataInnerAssigneesInner, workflowRelease NullableString, expiresAt NullableString, ) *UserTasksGet200ResponseDataInner`

NewUserTasksGet200ResponseDataInner instantiates a new UserTasksGet200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserTasksGet200ResponseDataInnerWithDefaults

`func NewUserTasksGet200ResponseDataInnerWithDefaults() *UserTasksGet200ResponseDataInner`

NewUserTasksGet200ResponseDataInnerWithDefaults instantiates a new UserTasksGet200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserTasksGet200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserTasksGet200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserTasksGet200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetWorkflowId

`func (o *UserTasksGet200ResponseDataInner) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *UserTasksGet200ResponseDataInner) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *UserTasksGet200ResponseDataInner) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetWorkflowName

`func (o *UserTasksGet200ResponseDataInner) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *UserTasksGet200ResponseDataInner) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *UserTasksGet200ResponseDataInner) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.


### GetWorkflowRunId

`func (o *UserTasksGet200ResponseDataInner) GetWorkflowRunId() string`

GetWorkflowRunId returns the WorkflowRunId field if non-nil, zero value otherwise.

### GetWorkflowRunIdOk

`func (o *UserTasksGet200ResponseDataInner) GetWorkflowRunIdOk() (*string, bool)`

GetWorkflowRunIdOk returns a tuple with the WorkflowRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowRunId

`func (o *UserTasksGet200ResponseDataInner) SetWorkflowRunId(v string)`

SetWorkflowRunId sets WorkflowRunId field to given value.


### GetTaskName

`func (o *UserTasksGet200ResponseDataInner) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *UserTasksGet200ResponseDataInner) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *UserTasksGet200ResponseDataInner) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.


### GetStatus

`func (o *UserTasksGet200ResponseDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserTasksGet200ResponseDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserTasksGet200ResponseDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetContext

`func (o *UserTasksGet200ResponseDataInner) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *UserTasksGet200ResponseDataInner) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *UserTasksGet200ResponseDataInner) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.


### GetCreatedAt

`func (o *UserTasksGet200ResponseDataInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserTasksGet200ResponseDataInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserTasksGet200ResponseDataInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCompletedAt

`func (o *UserTasksGet200ResponseDataInner) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *UserTasksGet200ResponseDataInner) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *UserTasksGet200ResponseDataInner) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.


### SetCompletedAtNil

`func (o *UserTasksGet200ResponseDataInner) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *UserTasksGet200ResponseDataInner) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetTaskUrl

`func (o *UserTasksGet200ResponseDataInner) GetTaskUrl() []string`

GetTaskUrl returns the TaskUrl field if non-nil, zero value otherwise.

### GetTaskUrlOk

`func (o *UserTasksGet200ResponseDataInner) GetTaskUrlOk() (*[]string, bool)`

GetTaskUrlOk returns a tuple with the TaskUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskUrl

`func (o *UserTasksGet200ResponseDataInner) SetTaskUrl(v []string)`

SetTaskUrl sets TaskUrl field to given value.


### GetAssignees

`func (o *UserTasksGet200ResponseDataInner) GetAssignees() []UserTasksGet200ResponseDataInnerAssigneesInner`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *UserTasksGet200ResponseDataInner) GetAssigneesOk() (*[]UserTasksGet200ResponseDataInnerAssigneesInner, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *UserTasksGet200ResponseDataInner) SetAssignees(v []UserTasksGet200ResponseDataInnerAssigneesInner)`

SetAssignees sets Assignees field to given value.


### GetWorkflowRelease

`func (o *UserTasksGet200ResponseDataInner) GetWorkflowRelease() string`

GetWorkflowRelease returns the WorkflowRelease field if non-nil, zero value otherwise.

### GetWorkflowReleaseOk

`func (o *UserTasksGet200ResponseDataInner) GetWorkflowReleaseOk() (*string, bool)`

GetWorkflowReleaseOk returns a tuple with the WorkflowRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowRelease

`func (o *UserTasksGet200ResponseDataInner) SetWorkflowRelease(v string)`

SetWorkflowRelease sets WorkflowRelease field to given value.


### SetWorkflowReleaseNil

`func (o *UserTasksGet200ResponseDataInner) SetWorkflowReleaseNil(b bool)`

 SetWorkflowReleaseNil sets the value for WorkflowRelease to be an explicit nil

### UnsetWorkflowRelease
`func (o *UserTasksGet200ResponseDataInner) UnsetWorkflowRelease()`

UnsetWorkflowRelease ensures that no value is present for WorkflowRelease, not even an explicit nil
### GetExpiresAt

`func (o *UserTasksGet200ResponseDataInner) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *UserTasksGet200ResponseDataInner) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *UserTasksGet200ResponseDataInner) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *UserTasksGet200ResponseDataInner) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *UserTasksGet200ResponseDataInner) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetOutput

`func (o *UserTasksGet200ResponseDataInner) GetOutput() map[string]interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *UserTasksGet200ResponseDataInner) GetOutputOk() (*map[string]interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *UserTasksGet200ResponseDataInner) SetOutput(v map[string]interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *UserTasksGet200ResponseDataInner) HasOutput() bool`

HasOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


