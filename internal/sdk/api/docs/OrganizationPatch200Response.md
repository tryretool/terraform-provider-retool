# OrganizationPatch200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**OrganizationPatch200ResponseData**](OrganizationPatch200ResponseData.md) |  | 

## Methods

### NewOrganizationPatch200Response

`func NewOrganizationPatch200Response(success bool, data OrganizationPatch200ResponseData, ) *OrganizationPatch200Response`

NewOrganizationPatch200Response instantiates a new OrganizationPatch200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationPatch200ResponseWithDefaults

`func NewOrganizationPatch200ResponseWithDefaults() *OrganizationPatch200Response`

NewOrganizationPatch200ResponseWithDefaults instantiates a new OrganizationPatch200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *OrganizationPatch200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *OrganizationPatch200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *OrganizationPatch200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *OrganizationPatch200Response) GetData() OrganizationPatch200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OrganizationPatch200Response) GetDataOk() (*OrganizationPatch200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OrganizationPatch200Response) SetData(v OrganizationPatch200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


