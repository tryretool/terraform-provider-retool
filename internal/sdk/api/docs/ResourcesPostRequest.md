# ResourcesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of resource. | 
**DisplayName** | **string** |  | 
**Options** | [**ResourcesPostRequestOptions**](ResourcesPostRequestOptions.md) |  | 

## Methods

### NewResourcesPostRequest

`func NewResourcesPostRequest(type_ string, displayName string, options ResourcesPostRequestOptions, ) *ResourcesPostRequest`

NewResourcesPostRequest instantiates a new ResourcesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcesPostRequestWithDefaults

`func NewResourcesPostRequestWithDefaults() *ResourcesPostRequest`

NewResourcesPostRequestWithDefaults instantiates a new ResourcesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ResourcesPostRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourcesPostRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourcesPostRequest) SetType(v string)`

SetType sets Type field to given value.


### GetDisplayName

`func (o *ResourcesPostRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ResourcesPostRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ResourcesPostRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetOptions

`func (o *ResourcesPostRequest) GetOptions() ResourcesPostRequestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ResourcesPostRequest) GetOptionsOk() (*ResourcesPostRequestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ResourcesPostRequest) SetOptions(v ResourcesPostRequestOptions)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


