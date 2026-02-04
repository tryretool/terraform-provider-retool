# UsageUserSummaryGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**[]UsageUserSummaryGet200ResponseDataInner**](UsageUserSummaryGet200ResponseDataInner.md) | An array of requested items | 

## Methods

### NewUsageUserSummaryGet200Response

`func NewUsageUserSummaryGet200Response(success bool, data []UsageUserSummaryGet200ResponseDataInner, ) *UsageUserSummaryGet200Response`

NewUsageUserSummaryGet200Response instantiates a new UsageUserSummaryGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageUserSummaryGet200ResponseWithDefaults

`func NewUsageUserSummaryGet200ResponseWithDefaults() *UsageUserSummaryGet200Response`

NewUsageUserSummaryGet200ResponseWithDefaults instantiates a new UsageUserSummaryGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *UsageUserSummaryGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UsageUserSummaryGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UsageUserSummaryGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *UsageUserSummaryGet200Response) GetData() []UsageUserSummaryGet200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UsageUserSummaryGet200Response) GetDataOk() (*[]UsageUserSummaryGet200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UsageUserSummaryGet200Response) SetData(v []UsageUserSummaryGet200ResponseDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


