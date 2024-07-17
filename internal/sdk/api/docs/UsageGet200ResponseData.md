# UsageGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountCurrentPageSaves** | **float32** | The number of page saves for all the organizations within the date range | 
**CountCurrentPageViews** | **float32** | The number of page views for all the organizations within the date range | 
**CountCurrentUsers** | **float32** | The total number of unique users who have performed any tracked activity in all the organizations within the date range | 
**CountT30Users** | **float32** | The number of unique users who have been active over the trailing 30 days from the end date provided | 
**PercentGrowthPageSaves** | **float32** | The percentage growth in page saves compared to the previous cycle of the same length as the provided date range | 
**PercentGrowthPageViews** | **float32** | The percentage growth in page views compared to the previous cycle of the same length as the provided date range | 
**PercentGrowthUsers** | **float32** | The percentage growth in unique users compared to the previous cycle of the same length as the provided date range | 
**PercentGrowthT30Users** | **float32** | The percentage growth in T30 users (trailing 30 days users) compared to the previous cycle of the same length as the provided date range | 
**DailyT30Usage** | [**[]UsageGet200ResponseDataDailyT30UsageInner**](UsageGet200ResponseDataDailyT30UsageInner.md) |  | 

## Methods

### NewUsageGet200ResponseData

`func NewUsageGet200ResponseData(countCurrentPageSaves float32, countCurrentPageViews float32, countCurrentUsers float32, countT30Users float32, percentGrowthPageSaves float32, percentGrowthPageViews float32, percentGrowthUsers float32, percentGrowthT30Users float32, dailyT30Usage []UsageGet200ResponseDataDailyT30UsageInner, ) *UsageGet200ResponseData`

NewUsageGet200ResponseData instantiates a new UsageGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageGet200ResponseDataWithDefaults

`func NewUsageGet200ResponseDataWithDefaults() *UsageGet200ResponseData`

NewUsageGet200ResponseDataWithDefaults instantiates a new UsageGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountCurrentPageSaves

`func (o *UsageGet200ResponseData) GetCountCurrentPageSaves() float32`

GetCountCurrentPageSaves returns the CountCurrentPageSaves field if non-nil, zero value otherwise.

### GetCountCurrentPageSavesOk

`func (o *UsageGet200ResponseData) GetCountCurrentPageSavesOk() (*float32, bool)`

GetCountCurrentPageSavesOk returns a tuple with the CountCurrentPageSaves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountCurrentPageSaves

`func (o *UsageGet200ResponseData) SetCountCurrentPageSaves(v float32)`

SetCountCurrentPageSaves sets CountCurrentPageSaves field to given value.


### GetCountCurrentPageViews

`func (o *UsageGet200ResponseData) GetCountCurrentPageViews() float32`

GetCountCurrentPageViews returns the CountCurrentPageViews field if non-nil, zero value otherwise.

### GetCountCurrentPageViewsOk

`func (o *UsageGet200ResponseData) GetCountCurrentPageViewsOk() (*float32, bool)`

GetCountCurrentPageViewsOk returns a tuple with the CountCurrentPageViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountCurrentPageViews

`func (o *UsageGet200ResponseData) SetCountCurrentPageViews(v float32)`

SetCountCurrentPageViews sets CountCurrentPageViews field to given value.


### GetCountCurrentUsers

`func (o *UsageGet200ResponseData) GetCountCurrentUsers() float32`

GetCountCurrentUsers returns the CountCurrentUsers field if non-nil, zero value otherwise.

### GetCountCurrentUsersOk

`func (o *UsageGet200ResponseData) GetCountCurrentUsersOk() (*float32, bool)`

GetCountCurrentUsersOk returns a tuple with the CountCurrentUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountCurrentUsers

`func (o *UsageGet200ResponseData) SetCountCurrentUsers(v float32)`

SetCountCurrentUsers sets CountCurrentUsers field to given value.


### GetCountT30Users

`func (o *UsageGet200ResponseData) GetCountT30Users() float32`

GetCountT30Users returns the CountT30Users field if non-nil, zero value otherwise.

### GetCountT30UsersOk

`func (o *UsageGet200ResponseData) GetCountT30UsersOk() (*float32, bool)`

GetCountT30UsersOk returns a tuple with the CountT30Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountT30Users

`func (o *UsageGet200ResponseData) SetCountT30Users(v float32)`

SetCountT30Users sets CountT30Users field to given value.


### GetPercentGrowthPageSaves

`func (o *UsageGet200ResponseData) GetPercentGrowthPageSaves() float32`

GetPercentGrowthPageSaves returns the PercentGrowthPageSaves field if non-nil, zero value otherwise.

### GetPercentGrowthPageSavesOk

`func (o *UsageGet200ResponseData) GetPercentGrowthPageSavesOk() (*float32, bool)`

GetPercentGrowthPageSavesOk returns a tuple with the PercentGrowthPageSaves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentGrowthPageSaves

`func (o *UsageGet200ResponseData) SetPercentGrowthPageSaves(v float32)`

SetPercentGrowthPageSaves sets PercentGrowthPageSaves field to given value.


### GetPercentGrowthPageViews

`func (o *UsageGet200ResponseData) GetPercentGrowthPageViews() float32`

GetPercentGrowthPageViews returns the PercentGrowthPageViews field if non-nil, zero value otherwise.

### GetPercentGrowthPageViewsOk

`func (o *UsageGet200ResponseData) GetPercentGrowthPageViewsOk() (*float32, bool)`

GetPercentGrowthPageViewsOk returns a tuple with the PercentGrowthPageViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentGrowthPageViews

`func (o *UsageGet200ResponseData) SetPercentGrowthPageViews(v float32)`

SetPercentGrowthPageViews sets PercentGrowthPageViews field to given value.


### GetPercentGrowthUsers

`func (o *UsageGet200ResponseData) GetPercentGrowthUsers() float32`

GetPercentGrowthUsers returns the PercentGrowthUsers field if non-nil, zero value otherwise.

### GetPercentGrowthUsersOk

`func (o *UsageGet200ResponseData) GetPercentGrowthUsersOk() (*float32, bool)`

GetPercentGrowthUsersOk returns a tuple with the PercentGrowthUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentGrowthUsers

`func (o *UsageGet200ResponseData) SetPercentGrowthUsers(v float32)`

SetPercentGrowthUsers sets PercentGrowthUsers field to given value.


### GetPercentGrowthT30Users

`func (o *UsageGet200ResponseData) GetPercentGrowthT30Users() float32`

GetPercentGrowthT30Users returns the PercentGrowthT30Users field if non-nil, zero value otherwise.

### GetPercentGrowthT30UsersOk

`func (o *UsageGet200ResponseData) GetPercentGrowthT30UsersOk() (*float32, bool)`

GetPercentGrowthT30UsersOk returns a tuple with the PercentGrowthT30Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentGrowthT30Users

`func (o *UsageGet200ResponseData) SetPercentGrowthT30Users(v float32)`

SetPercentGrowthT30Users sets PercentGrowthT30Users field to given value.


### GetDailyT30Usage

`func (o *UsageGet200ResponseData) GetDailyT30Usage() []UsageGet200ResponseDataDailyT30UsageInner`

GetDailyT30Usage returns the DailyT30Usage field if non-nil, zero value otherwise.

### GetDailyT30UsageOk

`func (o *UsageGet200ResponseData) GetDailyT30UsageOk() (*[]UsageGet200ResponseDataDailyT30UsageInner, bool)`

GetDailyT30UsageOk returns a tuple with the DailyT30Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyT30Usage

`func (o *UsageGet200ResponseData) SetDailyT30Usage(v []UsageGet200ResponseDataDailyT30UsageInner)`

SetDailyT30Usage sets DailyT30Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


