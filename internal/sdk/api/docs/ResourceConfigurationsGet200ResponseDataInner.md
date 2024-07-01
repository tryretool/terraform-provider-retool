# ResourceConfigurationsGet200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The uuid for the resource configuration. | 
**Resource** | [**ResourcesGet200ResponseDataInner**](ResourcesGet200ResponseDataInner.md) |  | 
**Environment** | [**ResourceConfigurationsGet200ResponseDataInnerEnvironment**](ResourceConfigurationsGet200ResponseDataInnerEnvironment.md) |  | 
**Options** | [**ResourceConfigurationsGet200ResponseDataInnerOptions**](ResourceConfigurationsGet200ResponseDataInnerOptions.md) |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 

## Methods

### NewResourceConfigurationsGet200ResponseDataInner

`func NewResourceConfigurationsGet200ResponseDataInner(id string, resource ResourcesGet200ResponseDataInner, environment ResourceConfigurationsGet200ResponseDataInnerEnvironment, options ResourceConfigurationsGet200ResponseDataInnerOptions, createdAt string, updatedAt string, ) *ResourceConfigurationsGet200ResponseDataInner`

NewResourceConfigurationsGet200ResponseDataInner instantiates a new ResourceConfigurationsGet200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceConfigurationsGet200ResponseDataInnerWithDefaults

`func NewResourceConfigurationsGet200ResponseDataInnerWithDefaults() *ResourceConfigurationsGet200ResponseDataInner`

NewResourceConfigurationsGet200ResponseDataInnerWithDefaults instantiates a new ResourceConfigurationsGet200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceConfigurationsGet200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceConfigurationsGet200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceConfigurationsGet200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetResource

`func (o *ResourceConfigurationsGet200ResponseDataInner) GetResource() ResourcesGet200ResponseDataInner`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ResourceConfigurationsGet200ResponseDataInner) GetResourceOk() (*ResourcesGet200ResponseDataInner, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ResourceConfigurationsGet200ResponseDataInner) SetResource(v ResourcesGet200ResponseDataInner)`

SetResource sets Resource field to given value.


### GetEnvironment

`func (o *ResourceConfigurationsGet200ResponseDataInner) GetEnvironment() ResourceConfigurationsGet200ResponseDataInnerEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ResourceConfigurationsGet200ResponseDataInner) GetEnvironmentOk() (*ResourceConfigurationsGet200ResponseDataInnerEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ResourceConfigurationsGet200ResponseDataInner) SetEnvironment(v ResourceConfigurationsGet200ResponseDataInnerEnvironment)`

SetEnvironment sets Environment field to given value.


### GetOptions

`func (o *ResourceConfigurationsGet200ResponseDataInner) GetOptions() ResourceConfigurationsGet200ResponseDataInnerOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ResourceConfigurationsGet200ResponseDataInner) GetOptionsOk() (*ResourceConfigurationsGet200ResponseDataInnerOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ResourceConfigurationsGet200ResponseDataInner) SetOptions(v ResourceConfigurationsGet200ResponseDataInnerOptions)`

SetOptions sets Options field to given value.


### GetCreatedAt

`func (o *ResourceConfigurationsGet200ResponseDataInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResourceConfigurationsGet200ResponseDataInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResourceConfigurationsGet200ResponseDataInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ResourceConfigurationsGet200ResponseDataInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResourceConfigurationsGet200ResponseDataInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResourceConfigurationsGet200ResponseDataInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


