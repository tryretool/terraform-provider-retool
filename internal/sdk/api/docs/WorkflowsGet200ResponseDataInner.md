# WorkflowsGet200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Workflow ID. | 
**Name** | **string** | The name of the Workflow | 
**Description** | **NullableString** | The description of the Workflow | 
**Crontab** | **NullableString** | The cron tab of the Workflow | 
**Timezone** | **NullableString** | The timezone of the Workflow | 
**IsEnabled** | **bool** | Whether the Workflow is enabled | 
**FolderId** | **string** | The folder ID of the Workflow | 
**Protected** | **bool** | Whether the Workflow is protected | 
**CreatedBy** | **NullableFloat32** | The user ID of the creator of the Workflow | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewWorkflowsGet200ResponseDataInner

`func NewWorkflowsGet200ResponseDataInner(id string, name string, description NullableString, crontab NullableString, timezone NullableString, isEnabled bool, folderId string, protected bool, createdBy NullableFloat32, createdAt time.Time, updatedAt time.Time, ) *WorkflowsGet200ResponseDataInner`

NewWorkflowsGet200ResponseDataInner instantiates a new WorkflowsGet200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowsGet200ResponseDataInnerWithDefaults

`func NewWorkflowsGet200ResponseDataInnerWithDefaults() *WorkflowsGet200ResponseDataInner`

NewWorkflowsGet200ResponseDataInnerWithDefaults instantiates a new WorkflowsGet200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowsGet200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowsGet200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowsGet200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WorkflowsGet200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowsGet200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowsGet200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WorkflowsGet200ResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowsGet200ResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowsGet200ResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *WorkflowsGet200ResponseDataInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WorkflowsGet200ResponseDataInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCrontab

`func (o *WorkflowsGet200ResponseDataInner) GetCrontab() string`

GetCrontab returns the Crontab field if non-nil, zero value otherwise.

### GetCrontabOk

`func (o *WorkflowsGet200ResponseDataInner) GetCrontabOk() (*string, bool)`

GetCrontabOk returns a tuple with the Crontab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrontab

`func (o *WorkflowsGet200ResponseDataInner) SetCrontab(v string)`

SetCrontab sets Crontab field to given value.


### SetCrontabNil

`func (o *WorkflowsGet200ResponseDataInner) SetCrontabNil(b bool)`

 SetCrontabNil sets the value for Crontab to be an explicit nil

### UnsetCrontab
`func (o *WorkflowsGet200ResponseDataInner) UnsetCrontab()`

UnsetCrontab ensures that no value is present for Crontab, not even an explicit nil
### GetTimezone

`func (o *WorkflowsGet200ResponseDataInner) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *WorkflowsGet200ResponseDataInner) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *WorkflowsGet200ResponseDataInner) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### SetTimezoneNil

`func (o *WorkflowsGet200ResponseDataInner) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *WorkflowsGet200ResponseDataInner) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetIsEnabled

`func (o *WorkflowsGet200ResponseDataInner) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *WorkflowsGet200ResponseDataInner) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *WorkflowsGet200ResponseDataInner) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetFolderId

`func (o *WorkflowsGet200ResponseDataInner) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *WorkflowsGet200ResponseDataInner) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *WorkflowsGet200ResponseDataInner) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.


### GetProtected

`func (o *WorkflowsGet200ResponseDataInner) GetProtected() bool`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *WorkflowsGet200ResponseDataInner) GetProtectedOk() (*bool, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *WorkflowsGet200ResponseDataInner) SetProtected(v bool)`

SetProtected sets Protected field to given value.


### GetCreatedBy

`func (o *WorkflowsGet200ResponseDataInner) GetCreatedBy() float32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WorkflowsGet200ResponseDataInner) GetCreatedByOk() (*float32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WorkflowsGet200ResponseDataInner) SetCreatedBy(v float32)`

SetCreatedBy sets CreatedBy field to given value.


### SetCreatedByNil

`func (o *WorkflowsGet200ResponseDataInner) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *WorkflowsGet200ResponseDataInner) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedAt

`func (o *WorkflowsGet200ResponseDataInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkflowsGet200ResponseDataInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkflowsGet200ResponseDataInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *WorkflowsGet200ResponseDataInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorkflowsGet200ResponseDataInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorkflowsGet200ResponseDataInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


