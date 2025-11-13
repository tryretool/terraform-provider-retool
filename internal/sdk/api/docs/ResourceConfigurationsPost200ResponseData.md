# ResourceConfigurationsPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The uuid for the resource configuration. | 
**Resource** | [**ResourcesGet200ResponseDataInner**](ResourcesGet200ResponseDataInner.md) |  | 
**Environment** | [**ResourceConfigurationsGet200ResponseDataInnerEnvironment**](ResourceConfigurationsGet200ResponseDataInnerEnvironment.md) |  | 
**Options** | [**ResourcesPostRequestOptions**](ResourcesPostRequestOptions.md) |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 

## Methods

### NewResourceConfigurationsPost200ResponseData

`func NewResourceConfigurationsPost200ResponseData(id string, resource ResourcesGet200ResponseDataInner, environment ResourceConfigurationsGet200ResponseDataInnerEnvironment, options ResourcesPostRequestOptions, createdAt string, updatedAt string, ) *ResourceConfigurationsPost200ResponseData`

NewResourceConfigurationsPost200ResponseData instantiates a new ResourceConfigurationsPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceConfigurationsPost200ResponseDataWithDefaults

`func NewResourceConfigurationsPost200ResponseDataWithDefaults() *ResourceConfigurationsPost200ResponseData`

NewResourceConfigurationsPost200ResponseDataWithDefaults instantiates a new ResourceConfigurationsPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceConfigurationsPost200ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceConfigurationsPost200ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceConfigurationsPost200ResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetResource

`func (o *ResourceConfigurationsPost200ResponseData) GetResource() ResourcesGet200ResponseDataInner`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ResourceConfigurationsPost200ResponseData) GetResourceOk() (*ResourcesGet200ResponseDataInner, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ResourceConfigurationsPost200ResponseData) SetResource(v ResourcesGet200ResponseDataInner)`

SetResource sets Resource field to given value.


### GetEnvironment

`func (o *ResourceConfigurationsPost200ResponseData) GetEnvironment() ResourceConfigurationsGet200ResponseDataInnerEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ResourceConfigurationsPost200ResponseData) GetEnvironmentOk() (*ResourceConfigurationsGet200ResponseDataInnerEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ResourceConfigurationsPost200ResponseData) SetEnvironment(v ResourceConfigurationsGet200ResponseDataInnerEnvironment)`

SetEnvironment sets Environment field to given value.


### GetOptions

`func (o *ResourceConfigurationsPost200ResponseData) GetOptions() ResourcesPostRequestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ResourceConfigurationsPost200ResponseData) GetOptionsOk() (*ResourcesPostRequestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ResourceConfigurationsPost200ResponseData) SetOptions(v ResourcesPostRequestOptions)`

SetOptions sets Options field to given value.


### GetCreatedAt

`func (o *ResourceConfigurationsPost200ResponseData) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResourceConfigurationsPost200ResponseData) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResourceConfigurationsPost200ResponseData) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ResourceConfigurationsPost200ResponseData) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResourceConfigurationsPost200ResponseData) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResourceConfigurationsPost200ResponseData) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


