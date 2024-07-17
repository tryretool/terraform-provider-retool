# UsageUserSummaryGet200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **string** | The id of the organization to which this user belongs | 
**UserId** | **string** | The id of the user | 
**Email** | **string** | The email of the user | 
**Host** | **string** | The host of the organization to which this user belongs | 
**CountAppViews** | **float32** | The number of times the user viewed an app in the time range specified | 
**CountAppSaves** | **float32** | The number of times the user edited an app in the time range specified | 
**CountUniqueApps** | **float32** | The number of unique apps edited in the time range specified | 
**LastActive** | **time.Time** |  | 

## Methods

### NewUsageUserSummaryGet200ResponseDataInner

`func NewUsageUserSummaryGet200ResponseDataInner(orgId string, userId string, email string, host string, countAppViews float32, countAppSaves float32, countUniqueApps float32, lastActive time.Time, ) *UsageUserSummaryGet200ResponseDataInner`

NewUsageUserSummaryGet200ResponseDataInner instantiates a new UsageUserSummaryGet200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageUserSummaryGet200ResponseDataInnerWithDefaults

`func NewUsageUserSummaryGet200ResponseDataInnerWithDefaults() *UsageUserSummaryGet200ResponseDataInner`

NewUsageUserSummaryGet200ResponseDataInnerWithDefaults instantiates a new UsageUserSummaryGet200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *UsageUserSummaryGet200ResponseDataInner) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *UsageUserSummaryGet200ResponseDataInner) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *UsageUserSummaryGet200ResponseDataInner) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetUserId

`func (o *UsageUserSummaryGet200ResponseDataInner) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UsageUserSummaryGet200ResponseDataInner) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UsageUserSummaryGet200ResponseDataInner) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetEmail

`func (o *UsageUserSummaryGet200ResponseDataInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UsageUserSummaryGet200ResponseDataInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UsageUserSummaryGet200ResponseDataInner) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetHost

`func (o *UsageUserSummaryGet200ResponseDataInner) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *UsageUserSummaryGet200ResponseDataInner) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *UsageUserSummaryGet200ResponseDataInner) SetHost(v string)`

SetHost sets Host field to given value.


### GetCountAppViews

`func (o *UsageUserSummaryGet200ResponseDataInner) GetCountAppViews() float32`

GetCountAppViews returns the CountAppViews field if non-nil, zero value otherwise.

### GetCountAppViewsOk

`func (o *UsageUserSummaryGet200ResponseDataInner) GetCountAppViewsOk() (*float32, bool)`

GetCountAppViewsOk returns a tuple with the CountAppViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountAppViews

`func (o *UsageUserSummaryGet200ResponseDataInner) SetCountAppViews(v float32)`

SetCountAppViews sets CountAppViews field to given value.


### GetCountAppSaves

`func (o *UsageUserSummaryGet200ResponseDataInner) GetCountAppSaves() float32`

GetCountAppSaves returns the CountAppSaves field if non-nil, zero value otherwise.

### GetCountAppSavesOk

`func (o *UsageUserSummaryGet200ResponseDataInner) GetCountAppSavesOk() (*float32, bool)`

GetCountAppSavesOk returns a tuple with the CountAppSaves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountAppSaves

`func (o *UsageUserSummaryGet200ResponseDataInner) SetCountAppSaves(v float32)`

SetCountAppSaves sets CountAppSaves field to given value.


### GetCountUniqueApps

`func (o *UsageUserSummaryGet200ResponseDataInner) GetCountUniqueApps() float32`

GetCountUniqueApps returns the CountUniqueApps field if non-nil, zero value otherwise.

### GetCountUniqueAppsOk

`func (o *UsageUserSummaryGet200ResponseDataInner) GetCountUniqueAppsOk() (*float32, bool)`

GetCountUniqueAppsOk returns a tuple with the CountUniqueApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountUniqueApps

`func (o *UsageUserSummaryGet200ResponseDataInner) SetCountUniqueApps(v float32)`

SetCountUniqueApps sets CountUniqueApps field to given value.


### GetLastActive

`func (o *UsageUserSummaryGet200ResponseDataInner) GetLastActive() time.Time`

GetLastActive returns the LastActive field if non-nil, zero value otherwise.

### GetLastActiveOk

`func (o *UsageUserSummaryGet200ResponseDataInner) GetLastActiveOk() (*time.Time, bool)`

GetLastActiveOk returns a tuple with the LastActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActive

`func (o *UsageUserSummaryGet200ResponseDataInner) SetLastActive(v time.Time)`

SetLastActive sets LastActive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


