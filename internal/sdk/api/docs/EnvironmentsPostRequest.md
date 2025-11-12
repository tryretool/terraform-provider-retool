# EnvironmentsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the environment. | 
**Description** | Pointer to **string** | A description for the environment. | [optional] 
**Color** | **string** | The (hexadecimal) color of the environment. | 

## Methods

### NewEnvironmentsPostRequest

`func NewEnvironmentsPostRequest(name string, color string, ) *EnvironmentsPostRequest`

NewEnvironmentsPostRequest instantiates a new EnvironmentsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentsPostRequestWithDefaults

`func NewEnvironmentsPostRequestWithDefaults() *EnvironmentsPostRequest`

NewEnvironmentsPostRequestWithDefaults instantiates a new EnvironmentsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnvironmentsPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentsPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentsPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EnvironmentsPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnvironmentsPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnvironmentsPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnvironmentsPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetColor

`func (o *EnvironmentsPostRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *EnvironmentsPostRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *EnvironmentsPostRequest) SetColor(v string)`

SetColor sets Color field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


