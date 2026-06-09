# AppsToolscriptValidatePost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Valid** | **bool** |  | 
**Diagnostics** | [**[]AppsToolscriptValidatePost200ResponseDataDiagnosticsInner**](AppsToolscriptValidatePost200ResponseDataDiagnosticsInner.md) |  | 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewAppsToolscriptValidatePost200ResponseData

`func NewAppsToolscriptValidatePost200ResponseData(valid bool, diagnostics []AppsToolscriptValidatePost200ResponseDataDiagnosticsInner, ) *AppsToolscriptValidatePost200ResponseData`

NewAppsToolscriptValidatePost200ResponseData instantiates a new AppsToolscriptValidatePost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppsToolscriptValidatePost200ResponseDataWithDefaults

`func NewAppsToolscriptValidatePost200ResponseDataWithDefaults() *AppsToolscriptValidatePost200ResponseData`

NewAppsToolscriptValidatePost200ResponseDataWithDefaults instantiates a new AppsToolscriptValidatePost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValid

`func (o *AppsToolscriptValidatePost200ResponseData) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *AppsToolscriptValidatePost200ResponseData) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *AppsToolscriptValidatePost200ResponseData) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetDiagnostics

`func (o *AppsToolscriptValidatePost200ResponseData) GetDiagnostics() []AppsToolscriptValidatePost200ResponseDataDiagnosticsInner`

GetDiagnostics returns the Diagnostics field if non-nil, zero value otherwise.

### GetDiagnosticsOk

`func (o *AppsToolscriptValidatePost200ResponseData) GetDiagnosticsOk() (*[]AppsToolscriptValidatePost200ResponseDataDiagnosticsInner, bool)`

GetDiagnosticsOk returns a tuple with the Diagnostics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnostics

`func (o *AppsToolscriptValidatePost200ResponseData) SetDiagnostics(v []AppsToolscriptValidatePost200ResponseDataDiagnosticsInner)`

SetDiagnostics sets Diagnostics field to given value.


### GetError

`func (o *AppsToolscriptValidatePost200ResponseData) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AppsToolscriptValidatePost200ResponseData) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AppsToolscriptValidatePost200ResponseData) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AppsToolscriptValidatePost200ResponseData) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


