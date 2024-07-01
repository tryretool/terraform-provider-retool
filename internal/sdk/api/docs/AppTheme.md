# AppTheme

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**LegacyId** | **float32** |  | 
**Name** | **string** | The name of the app theme. | 
**Theme** | **map[string]interface{}** | The theme object. | 

## Methods

### NewAppTheme

`func NewAppTheme(id float32, legacyId float32, name string, theme map[string]interface{}, ) *AppTheme`

NewAppTheme instantiates a new AppTheme object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppThemeWithDefaults

`func NewAppThemeWithDefaults() *AppTheme`

NewAppThemeWithDefaults instantiates a new AppTheme object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppTheme) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppTheme) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppTheme) SetId(v float32)`

SetId sets Id field to given value.


### GetLegacyId

`func (o *AppTheme) GetLegacyId() float32`

GetLegacyId returns the LegacyId field if non-nil, zero value otherwise.

### GetLegacyIdOk

`func (o *AppTheme) GetLegacyIdOk() (*float32, bool)`

GetLegacyIdOk returns a tuple with the LegacyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyId

`func (o *AppTheme) SetLegacyId(v float32)`

SetLegacyId sets LegacyId field to given value.


### GetName

`func (o *AppTheme) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppTheme) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppTheme) SetName(v string)`

SetName sets Name field to given value.


### GetTheme

`func (o *AppTheme) GetTheme() map[string]interface{}`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *AppTheme) GetThemeOk() (*map[string]interface{}, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *AppTheme) SetTheme(v map[string]interface{})`

SetTheme sets Theme field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


