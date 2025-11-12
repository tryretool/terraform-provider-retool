# SourceControlManifestsGet200ResponseDataManifestsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Apps** | Pointer to [**[]SourceControlManifestsGet200ResponseDataManifestsInnerAppsInner**](SourceControlManifestsGet200ResponseDataManifestsInnerAppsInner.md) |  | [optional] 
**Workflows** | Pointer to [**[]SourceControlManifestsGet200ResponseDataManifestsInnerAppsInner**](SourceControlManifestsGet200ResponseDataManifestsInnerAppsInner.md) |  | [optional] 

## Methods

### NewSourceControlManifestsGet200ResponseDataManifestsInner

`func NewSourceControlManifestsGet200ResponseDataManifestsInner(name string, ) *SourceControlManifestsGet200ResponseDataManifestsInner`

NewSourceControlManifestsGet200ResponseDataManifestsInner instantiates a new SourceControlManifestsGet200ResponseDataManifestsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlManifestsGet200ResponseDataManifestsInnerWithDefaults

`func NewSourceControlManifestsGet200ResponseDataManifestsInnerWithDefaults() *SourceControlManifestsGet200ResponseDataManifestsInner`

NewSourceControlManifestsGet200ResponseDataManifestsInnerWithDefaults instantiates a new SourceControlManifestsGet200ResponseDataManifestsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SourceControlManifestsGet200ResponseDataManifestsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourceControlManifestsGet200ResponseDataManifestsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourceControlManifestsGet200ResponseDataManifestsInner) SetName(v string)`

SetName sets Name field to given value.


### GetApps

`func (o *SourceControlManifestsGet200ResponseDataManifestsInner) GetApps() []SourceControlManifestsGet200ResponseDataManifestsInnerAppsInner`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *SourceControlManifestsGet200ResponseDataManifestsInner) GetAppsOk() (*[]SourceControlManifestsGet200ResponseDataManifestsInnerAppsInner, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *SourceControlManifestsGet200ResponseDataManifestsInner) SetApps(v []SourceControlManifestsGet200ResponseDataManifestsInnerAppsInner)`

SetApps sets Apps field to given value.

### HasApps

`func (o *SourceControlManifestsGet200ResponseDataManifestsInner) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetWorkflows

`func (o *SourceControlManifestsGet200ResponseDataManifestsInner) GetWorkflows() []SourceControlManifestsGet200ResponseDataManifestsInnerAppsInner`

GetWorkflows returns the Workflows field if non-nil, zero value otherwise.

### GetWorkflowsOk

`func (o *SourceControlManifestsGet200ResponseDataManifestsInner) GetWorkflowsOk() (*[]SourceControlManifestsGet200ResponseDataManifestsInnerAppsInner, bool)`

GetWorkflowsOk returns a tuple with the Workflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflows

`func (o *SourceControlManifestsGet200ResponseDataManifestsInner) SetWorkflows(v []SourceControlManifestsGet200ResponseDataManifestsInnerAppsInner)`

SetWorkflows sets Workflows field to given value.

### HasWorkflows

`func (o *SourceControlManifestsGet200ResponseDataManifestsInner) HasWorkflows() bool`

HasWorkflows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


