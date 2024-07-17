# Usage

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

### NewUsage

`func NewUsage(countCurrentPageSaves float32, countCurrentPageViews float32, countCurrentUsers float32, countT30Users float32, percentGrowthPageSaves float32, percentGrowthPageViews float32, percentGrowthUsers float32, percentGrowthT30Users float32, dailyT30Usage []UsageGet200ResponseDataDailyT30UsageInner, ) *Usage`

NewUsage instantiates a new Usage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageWithDefaults

`func NewUsageWithDefaults() *Usage`

NewUsageWithDefaults instantiates a new Usage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountCurrentPageSaves

`func (o *Usage) GetCountCurrentPageSaves() float32`

GetCountCurrentPageSaves returns the CountCurrentPageSaves field if non-nil, zero value otherwise.

### GetCountCurrentPageSavesOk

`func (o *Usage) GetCountCurrentPageSavesOk() (*float32, bool)`

GetCountCurrentPageSavesOk returns a tuple with the CountCurrentPageSaves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountCurrentPageSaves

`func (o *Usage) SetCountCurrentPageSaves(v float32)`

SetCountCurrentPageSaves sets CountCurrentPageSaves field to given value.


### GetCountCurrentPageViews

`func (o *Usage) GetCountCurrentPageViews() float32`

GetCountCurrentPageViews returns the CountCurrentPageViews field if non-nil, zero value otherwise.

### GetCountCurrentPageViewsOk

`func (o *Usage) GetCountCurrentPageViewsOk() (*float32, bool)`

GetCountCurrentPageViewsOk returns a tuple with the CountCurrentPageViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountCurrentPageViews

`func (o *Usage) SetCountCurrentPageViews(v float32)`

SetCountCurrentPageViews sets CountCurrentPageViews field to given value.


### GetCountCurrentUsers

`func (o *Usage) GetCountCurrentUsers() float32`

GetCountCurrentUsers returns the CountCurrentUsers field if non-nil, zero value otherwise.

### GetCountCurrentUsersOk

`func (o *Usage) GetCountCurrentUsersOk() (*float32, bool)`

GetCountCurrentUsersOk returns a tuple with the CountCurrentUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountCurrentUsers

`func (o *Usage) SetCountCurrentUsers(v float32)`

SetCountCurrentUsers sets CountCurrentUsers field to given value.


### GetCountT30Users

`func (o *Usage) GetCountT30Users() float32`

GetCountT30Users returns the CountT30Users field if non-nil, zero value otherwise.

### GetCountT30UsersOk

`func (o *Usage) GetCountT30UsersOk() (*float32, bool)`

GetCountT30UsersOk returns a tuple with the CountT30Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountT30Users

`func (o *Usage) SetCountT30Users(v float32)`

SetCountT30Users sets CountT30Users field to given value.


### GetPercentGrowthPageSaves

`func (o *Usage) GetPercentGrowthPageSaves() float32`

GetPercentGrowthPageSaves returns the PercentGrowthPageSaves field if non-nil, zero value otherwise.

### GetPercentGrowthPageSavesOk

`func (o *Usage) GetPercentGrowthPageSavesOk() (*float32, bool)`

GetPercentGrowthPageSavesOk returns a tuple with the PercentGrowthPageSaves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentGrowthPageSaves

`func (o *Usage) SetPercentGrowthPageSaves(v float32)`

SetPercentGrowthPageSaves sets PercentGrowthPageSaves field to given value.


### GetPercentGrowthPageViews

`func (o *Usage) GetPercentGrowthPageViews() float32`

GetPercentGrowthPageViews returns the PercentGrowthPageViews field if non-nil, zero value otherwise.

### GetPercentGrowthPageViewsOk

`func (o *Usage) GetPercentGrowthPageViewsOk() (*float32, bool)`

GetPercentGrowthPageViewsOk returns a tuple with the PercentGrowthPageViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentGrowthPageViews

`func (o *Usage) SetPercentGrowthPageViews(v float32)`

SetPercentGrowthPageViews sets PercentGrowthPageViews field to given value.


### GetPercentGrowthUsers

`func (o *Usage) GetPercentGrowthUsers() float32`

GetPercentGrowthUsers returns the PercentGrowthUsers field if non-nil, zero value otherwise.

### GetPercentGrowthUsersOk

`func (o *Usage) GetPercentGrowthUsersOk() (*float32, bool)`

GetPercentGrowthUsersOk returns a tuple with the PercentGrowthUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentGrowthUsers

`func (o *Usage) SetPercentGrowthUsers(v float32)`

SetPercentGrowthUsers sets PercentGrowthUsers field to given value.


### GetPercentGrowthT30Users

`func (o *Usage) GetPercentGrowthT30Users() float32`

GetPercentGrowthT30Users returns the PercentGrowthT30Users field if non-nil, zero value otherwise.

### GetPercentGrowthT30UsersOk

`func (o *Usage) GetPercentGrowthT30UsersOk() (*float32, bool)`

GetPercentGrowthT30UsersOk returns a tuple with the PercentGrowthT30Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentGrowthT30Users

`func (o *Usage) SetPercentGrowthT30Users(v float32)`

SetPercentGrowthT30Users sets PercentGrowthT30Users field to given value.


### GetDailyT30Usage

`func (o *Usage) GetDailyT30Usage() []UsageGet200ResponseDataDailyT30UsageInner`

GetDailyT30Usage returns the DailyT30Usage field if non-nil, zero value otherwise.

### GetDailyT30UsageOk

`func (o *Usage) GetDailyT30UsageOk() (*[]UsageGet200ResponseDataDailyT30UsageInner, bool)`

GetDailyT30UsageOk returns a tuple with the DailyT30Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyT30Usage

`func (o *Usage) SetDailyT30Usage(v []UsageGet200ResponseDataDailyT30UsageInner)`

SetDailyT30Usage sets DailyT30Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


