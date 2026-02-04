# UsageOrganizationsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**[]UsageOrganizationsGet200ResponseDataInner**](UsageOrganizationsGet200ResponseDataInner.md) | An array of requested items | 

## Methods

### NewUsageOrganizationsGet200Response

`func NewUsageOrganizationsGet200Response(success bool, data []UsageOrganizationsGet200ResponseDataInner, ) *UsageOrganizationsGet200Response`

NewUsageOrganizationsGet200Response instantiates a new UsageOrganizationsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageOrganizationsGet200ResponseWithDefaults

`func NewUsageOrganizationsGet200ResponseWithDefaults() *UsageOrganizationsGet200Response`

NewUsageOrganizationsGet200ResponseWithDefaults instantiates a new UsageOrganizationsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *UsageOrganizationsGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UsageOrganizationsGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UsageOrganizationsGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *UsageOrganizationsGet200Response) GetData() []UsageOrganizationsGet200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UsageOrganizationsGet200Response) GetDataOk() (*[]UsageOrganizationsGet200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UsageOrganizationsGet200Response) SetData(v []UsageOrganizationsGet200ResponseDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


