# AppThemesPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**LegacyId** | **float32** |  | 
**Name** | **string** | The name of the app theme. | 
**Theme** | **map[string]interface{}** | The theme object. | 

## Methods

### NewAppThemesPutRequest

`func NewAppThemesPutRequest(id float32, legacyId float32, name string, theme map[string]interface{}, ) *AppThemesPutRequest`

NewAppThemesPutRequest instantiates a new AppThemesPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppThemesPutRequestWithDefaults

`func NewAppThemesPutRequestWithDefaults() *AppThemesPutRequest`

NewAppThemesPutRequestWithDefaults instantiates a new AppThemesPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppThemesPutRequest) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppThemesPutRequest) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppThemesPutRequest) SetId(v float32)`

SetId sets Id field to given value.


### GetLegacyId

`func (o *AppThemesPutRequest) GetLegacyId() float32`

GetLegacyId returns the LegacyId field if non-nil, zero value otherwise.

### GetLegacyIdOk

`func (o *AppThemesPutRequest) GetLegacyIdOk() (*float32, bool)`

GetLegacyIdOk returns a tuple with the LegacyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyId

`func (o *AppThemesPutRequest) SetLegacyId(v float32)`

SetLegacyId sets LegacyId field to given value.


### GetName

`func (o *AppThemesPutRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppThemesPutRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppThemesPutRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTheme

`func (o *AppThemesPutRequest) GetTheme() map[string]interface{}`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *AppThemesPutRequest) GetThemeOk() (*map[string]interface{}, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *AppThemesPutRequest) SetTheme(v map[string]interface{})`

SetTheme sets Theme field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


