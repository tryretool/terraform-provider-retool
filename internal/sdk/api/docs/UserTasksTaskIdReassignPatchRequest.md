# UserTasksTaskIdReassignPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Email of user reassigning the task. | 
**Assignees** | [**[]UserTasksTaskIdReassignPatchRequestAssigneesInner**](UserTasksTaskIdReassignPatchRequestAssigneesInner.md) | An array of assignees the user task should be reassigned to. Can include both groups and users. | 

## Methods

### NewUserTasksTaskIdReassignPatchRequest

`func NewUserTasksTaskIdReassignPatchRequest(email string, assignees []UserTasksTaskIdReassignPatchRequestAssigneesInner, ) *UserTasksTaskIdReassignPatchRequest`

NewUserTasksTaskIdReassignPatchRequest instantiates a new UserTasksTaskIdReassignPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserTasksTaskIdReassignPatchRequestWithDefaults

`func NewUserTasksTaskIdReassignPatchRequestWithDefaults() *UserTasksTaskIdReassignPatchRequest`

NewUserTasksTaskIdReassignPatchRequestWithDefaults instantiates a new UserTasksTaskIdReassignPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserTasksTaskIdReassignPatchRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserTasksTaskIdReassignPatchRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserTasksTaskIdReassignPatchRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetAssignees

`func (o *UserTasksTaskIdReassignPatchRequest) GetAssignees() []UserTasksTaskIdReassignPatchRequestAssigneesInner`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *UserTasksTaskIdReassignPatchRequest) GetAssigneesOk() (*[]UserTasksTaskIdReassignPatchRequestAssigneesInner, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *UserTasksTaskIdReassignPatchRequest) SetAssignees(v []UserTasksTaskIdReassignPatchRequestAssigneesInner)`

SetAssignees sets Assignees field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


