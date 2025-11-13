# ResourceConfigurationsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | A UUID that uniquely identifies a resource. This is also referenced within the Retool app. For example, when you edit a resource, the ID can be found in the url. | 
**EnvironmentId** | **string** | A UUID that uniquely identifies an environment. | 
**Options** | [**ResourcesPostRequestOptions**](ResourcesPostRequestOptions.md) |  | 

## Methods

### NewResourceConfigurationsPostRequest

`func NewResourceConfigurationsPostRequest(resourceId string, environmentId string, options ResourcesPostRequestOptions, ) *ResourceConfigurationsPostRequest`

NewResourceConfigurationsPostRequest instantiates a new ResourceConfigurationsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceConfigurationsPostRequestWithDefaults

`func NewResourceConfigurationsPostRequestWithDefaults() *ResourceConfigurationsPostRequest`

NewResourceConfigurationsPostRequestWithDefaults instantiates a new ResourceConfigurationsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *ResourceConfigurationsPostRequest) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ResourceConfigurationsPostRequest) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ResourceConfigurationsPostRequest) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetEnvironmentId

`func (o *ResourceConfigurationsPostRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ResourceConfigurationsPostRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ResourceConfigurationsPostRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetOptions

`func (o *ResourceConfigurationsPostRequest) GetOptions() ResourcesPostRequestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ResourceConfigurationsPostRequest) GetOptionsOk() (*ResourcesPostRequestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ResourceConfigurationsPostRequest) SetOptions(v ResourcesPostRequestOptions)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


