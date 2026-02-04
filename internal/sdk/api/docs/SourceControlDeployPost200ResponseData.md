# SourceControlDeployPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The deployment ID | 
**Status** | **string** | The deployment status | 
**CommitSha** | **NullableString** | The commit SHA that was deployed. Will be null while the deployment is in PENDING_START status. | 
**TriggeredBy** | **NullableString** | The user ID who triggered the deployment, or null if triggered automatically | 
**CreatedAt** | **string** | ISO 8601 timestamp when the deployment was created | 
**CompletedAt** | **NullableString** | ISO 8601 timestamp when the deployment completed, or null if still in progress | 

## Methods

### NewSourceControlDeployPost200ResponseData

`func NewSourceControlDeployPost200ResponseData(id string, status string, commitSha NullableString, triggeredBy NullableString, createdAt string, completedAt NullableString, ) *SourceControlDeployPost200ResponseData`

NewSourceControlDeployPost200ResponseData instantiates a new SourceControlDeployPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlDeployPost200ResponseDataWithDefaults

`func NewSourceControlDeployPost200ResponseDataWithDefaults() *SourceControlDeployPost200ResponseData`

NewSourceControlDeployPost200ResponseDataWithDefaults instantiates a new SourceControlDeployPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SourceControlDeployPost200ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SourceControlDeployPost200ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SourceControlDeployPost200ResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *SourceControlDeployPost200ResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SourceControlDeployPost200ResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SourceControlDeployPost200ResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCommitSha

`func (o *SourceControlDeployPost200ResponseData) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *SourceControlDeployPost200ResponseData) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *SourceControlDeployPost200ResponseData) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.


### SetCommitShaNil

`func (o *SourceControlDeployPost200ResponseData) SetCommitShaNil(b bool)`

 SetCommitShaNil sets the value for CommitSha to be an explicit nil

### UnsetCommitSha
`func (o *SourceControlDeployPost200ResponseData) UnsetCommitSha()`

UnsetCommitSha ensures that no value is present for CommitSha, not even an explicit nil
### GetTriggeredBy

`func (o *SourceControlDeployPost200ResponseData) GetTriggeredBy() string`

GetTriggeredBy returns the TriggeredBy field if non-nil, zero value otherwise.

### GetTriggeredByOk

`func (o *SourceControlDeployPost200ResponseData) GetTriggeredByOk() (*string, bool)`

GetTriggeredByOk returns a tuple with the TriggeredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredBy

`func (o *SourceControlDeployPost200ResponseData) SetTriggeredBy(v string)`

SetTriggeredBy sets TriggeredBy field to given value.


### SetTriggeredByNil

`func (o *SourceControlDeployPost200ResponseData) SetTriggeredByNil(b bool)`

 SetTriggeredByNil sets the value for TriggeredBy to be an explicit nil

### UnsetTriggeredBy
`func (o *SourceControlDeployPost200ResponseData) UnsetTriggeredBy()`

UnsetTriggeredBy ensures that no value is present for TriggeredBy, not even an explicit nil
### GetCreatedAt

`func (o *SourceControlDeployPost200ResponseData) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SourceControlDeployPost200ResponseData) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SourceControlDeployPost200ResponseData) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCompletedAt

`func (o *SourceControlDeployPost200ResponseData) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *SourceControlDeployPost200ResponseData) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *SourceControlDeployPost200ResponseData) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.


### SetCompletedAtNil

`func (o *SourceControlDeployPost200ResponseData) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *SourceControlDeployPost200ResponseData) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


