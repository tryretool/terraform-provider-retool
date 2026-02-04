# UsageAppSummaryGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**[]UsageAppSummaryGet200ResponseDataInner**](UsageAppSummaryGet200ResponseDataInner.md) | An array of requested items | 

## Methods

### NewUsageAppSummaryGet200Response

`func NewUsageAppSummaryGet200Response(success bool, data []UsageAppSummaryGet200ResponseDataInner, ) *UsageAppSummaryGet200Response`

NewUsageAppSummaryGet200Response instantiates a new UsageAppSummaryGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageAppSummaryGet200ResponseWithDefaults

`func NewUsageAppSummaryGet200ResponseWithDefaults() *UsageAppSummaryGet200Response`

NewUsageAppSummaryGet200ResponseWithDefaults instantiates a new UsageAppSummaryGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *UsageAppSummaryGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UsageAppSummaryGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UsageAppSummaryGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *UsageAppSummaryGet200Response) GetData() []UsageAppSummaryGet200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UsageAppSummaryGet200Response) GetDataOk() (*[]UsageAppSummaryGet200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UsageAppSummaryGet200Response) SetData(v []UsageAppSummaryGet200ResponseDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


