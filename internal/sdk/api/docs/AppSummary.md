# AppSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | **string** | The name of the app | 
**Host** | **string** | The organization host | 
**CountAppViews** | **float32** | The number of views of the app in the time range specified | 
**CountAppSaves** | **float32** | The number of saves of the app in the time range specified | 
**CountViewers** | **float32** | The number of viewers of the app in the time range specified | 
**CountEditors** | **float32** | The number of editors of the app in the time range specified | 

## Methods

### NewAppSummary

`func NewAppSummary(appName string, host string, countAppViews float32, countAppSaves float32, countViewers float32, countEditors float32, ) *AppSummary`

NewAppSummary instantiates a new AppSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSummaryWithDefaults

`func NewAppSummaryWithDefaults() *AppSummary`

NewAppSummaryWithDefaults instantiates a new AppSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *AppSummary) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *AppSummary) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *AppSummary) SetAppName(v string)`

SetAppName sets AppName field to given value.


### GetHost

`func (o *AppSummary) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AppSummary) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AppSummary) SetHost(v string)`

SetHost sets Host field to given value.


### GetCountAppViews

`func (o *AppSummary) GetCountAppViews() float32`

GetCountAppViews returns the CountAppViews field if non-nil, zero value otherwise.

### GetCountAppViewsOk

`func (o *AppSummary) GetCountAppViewsOk() (*float32, bool)`

GetCountAppViewsOk returns a tuple with the CountAppViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountAppViews

`func (o *AppSummary) SetCountAppViews(v float32)`

SetCountAppViews sets CountAppViews field to given value.


### GetCountAppSaves

`func (o *AppSummary) GetCountAppSaves() float32`

GetCountAppSaves returns the CountAppSaves field if non-nil, zero value otherwise.

### GetCountAppSavesOk

`func (o *AppSummary) GetCountAppSavesOk() (*float32, bool)`

GetCountAppSavesOk returns a tuple with the CountAppSaves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountAppSaves

`func (o *AppSummary) SetCountAppSaves(v float32)`

SetCountAppSaves sets CountAppSaves field to given value.


### GetCountViewers

`func (o *AppSummary) GetCountViewers() float32`

GetCountViewers returns the CountViewers field if non-nil, zero value otherwise.

### GetCountViewersOk

`func (o *AppSummary) GetCountViewersOk() (*float32, bool)`

GetCountViewersOk returns a tuple with the CountViewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountViewers

`func (o *AppSummary) SetCountViewers(v float32)`

SetCountViewers sets CountViewers field to given value.


### GetCountEditors

`func (o *AppSummary) GetCountEditors() float32`

GetCountEditors returns the CountEditors field if non-nil, zero value otherwise.

### GetCountEditorsOk

`func (o *AppSummary) GetCountEditorsOk() (*float32, bool)`

GetCountEditorsOk returns a tuple with the CountEditors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountEditors

`func (o *AppSummary) SetCountEditors(v float32)`

SetCountEditors sets CountEditors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


