# ResourceConfiguration

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

### NewResourceConfiguration

`func NewResourceConfiguration(id string, resource ResourcesGet200ResponseDataInner, environment ResourceConfigurationsGet200ResponseDataInnerEnvironment, options ResourcesPostRequestOptions, createdAt string, updatedAt string, ) *ResourceConfiguration`

NewResourceConfiguration instantiates a new ResourceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceConfigurationWithDefaults

`func NewResourceConfigurationWithDefaults() *ResourceConfiguration`

NewResourceConfigurationWithDefaults instantiates a new ResourceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceConfiguration) SetId(v string)`

SetId sets Id field to given value.


### GetResource

`func (o *ResourceConfiguration) GetResource() ResourcesGet200ResponseDataInner`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ResourceConfiguration) GetResourceOk() (*ResourcesGet200ResponseDataInner, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ResourceConfiguration) SetResource(v ResourcesGet200ResponseDataInner)`

SetResource sets Resource field to given value.


### GetEnvironment

`func (o *ResourceConfiguration) GetEnvironment() ResourceConfigurationsGet200ResponseDataInnerEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ResourceConfiguration) GetEnvironmentOk() (*ResourceConfigurationsGet200ResponseDataInnerEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ResourceConfiguration) SetEnvironment(v ResourceConfigurationsGet200ResponseDataInnerEnvironment)`

SetEnvironment sets Environment field to given value.


### GetOptions

`func (o *ResourceConfiguration) GetOptions() ResourcesPostRequestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ResourceConfiguration) GetOptionsOk() (*ResourcesPostRequestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ResourceConfiguration) SetOptions(v ResourcesPostRequestOptions)`

SetOptions sets Options field to given value.


### GetCreatedAt

`func (o *ResourceConfiguration) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResourceConfiguration) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResourceConfiguration) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ResourceConfiguration) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResourceConfiguration) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResourceConfiguration) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


