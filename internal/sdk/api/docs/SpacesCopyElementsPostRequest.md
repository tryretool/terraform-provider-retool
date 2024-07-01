# SpacesCopyElementsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceIds** | [**[]SpacesCopyElementsPostRequestResourceIdsInner**](SpacesCopyElementsPostRequestResourceIdsInner.md) | List of resource uuids to copy to the new space. | 
**QueryLibraryQueryIds** | **[]string** | List of query library query uuids to copy to the new space. | 
**AppIds** | **[]string** | List of app or module uuids to copy to the new space. | 
**WorkflowIds** | **[]string** | List of workflow ids to copy to the new space. | 
**DestinationSpaceId** | **string** | The id of the space to copy the elements to. | 

## Methods

### NewSpacesCopyElementsPostRequest

`func NewSpacesCopyElementsPostRequest(resourceIds []SpacesCopyElementsPostRequestResourceIdsInner, queryLibraryQueryIds []string, appIds []string, workflowIds []string, destinationSpaceId string, ) *SpacesCopyElementsPostRequest`

NewSpacesCopyElementsPostRequest instantiates a new SpacesCopyElementsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpacesCopyElementsPostRequestWithDefaults

`func NewSpacesCopyElementsPostRequestWithDefaults() *SpacesCopyElementsPostRequest`

NewSpacesCopyElementsPostRequestWithDefaults instantiates a new SpacesCopyElementsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceIds

`func (o *SpacesCopyElementsPostRequest) GetResourceIds() []SpacesCopyElementsPostRequestResourceIdsInner`

GetResourceIds returns the ResourceIds field if non-nil, zero value otherwise.

### GetResourceIdsOk

`func (o *SpacesCopyElementsPostRequest) GetResourceIdsOk() (*[]SpacesCopyElementsPostRequestResourceIdsInner, bool)`

GetResourceIdsOk returns a tuple with the ResourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceIds

`func (o *SpacesCopyElementsPostRequest) SetResourceIds(v []SpacesCopyElementsPostRequestResourceIdsInner)`

SetResourceIds sets ResourceIds field to given value.


### GetQueryLibraryQueryIds

`func (o *SpacesCopyElementsPostRequest) GetQueryLibraryQueryIds() []string`

GetQueryLibraryQueryIds returns the QueryLibraryQueryIds field if non-nil, zero value otherwise.

### GetQueryLibraryQueryIdsOk

`func (o *SpacesCopyElementsPostRequest) GetQueryLibraryQueryIdsOk() (*[]string, bool)`

GetQueryLibraryQueryIdsOk returns a tuple with the QueryLibraryQueryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryLibraryQueryIds

`func (o *SpacesCopyElementsPostRequest) SetQueryLibraryQueryIds(v []string)`

SetQueryLibraryQueryIds sets QueryLibraryQueryIds field to given value.


### GetAppIds

`func (o *SpacesCopyElementsPostRequest) GetAppIds() []string`

GetAppIds returns the AppIds field if non-nil, zero value otherwise.

### GetAppIdsOk

`func (o *SpacesCopyElementsPostRequest) GetAppIdsOk() (*[]string, bool)`

GetAppIdsOk returns a tuple with the AppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIds

`func (o *SpacesCopyElementsPostRequest) SetAppIds(v []string)`

SetAppIds sets AppIds field to given value.


### GetWorkflowIds

`func (o *SpacesCopyElementsPostRequest) GetWorkflowIds() []string`

GetWorkflowIds returns the WorkflowIds field if non-nil, zero value otherwise.

### GetWorkflowIdsOk

`func (o *SpacesCopyElementsPostRequest) GetWorkflowIdsOk() (*[]string, bool)`

GetWorkflowIdsOk returns a tuple with the WorkflowIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowIds

`func (o *SpacesCopyElementsPostRequest) SetWorkflowIds(v []string)`

SetWorkflowIds sets WorkflowIds field to given value.


### GetDestinationSpaceId

`func (o *SpacesCopyElementsPostRequest) GetDestinationSpaceId() string`

GetDestinationSpaceId returns the DestinationSpaceId field if non-nil, zero value otherwise.

### GetDestinationSpaceIdOk

`func (o *SpacesCopyElementsPostRequest) GetDestinationSpaceIdOk() (*string, bool)`

GetDestinationSpaceIdOk returns a tuple with the DestinationSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationSpaceId

`func (o *SpacesCopyElementsPostRequest) SetDestinationSpaceId(v string)`

SetDestinationSpaceId sets DestinationSpaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


