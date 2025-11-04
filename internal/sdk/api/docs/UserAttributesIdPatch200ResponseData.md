# UserAttributesIdPatch200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the user attribute | 
**Name** | **string** | The name of the user attribute | 
**Label** | **string** | The label of the user attribute | 
**DataType** | **string** | The data type of the user attribute. Valid types are &#x60;string&#x60;, &#x60;number&#x60;, or &#x60;json&#x60; | 
**DefaultValue** | **NullableString** | The default value of the user attribute, assigned to all users without the attribute set | 
**IntercomAttributeName** | **NullableString** | The name of the Intercom user attribute that this attribute should be mapped to | 

## Methods

### NewUserAttributesIdPatch200ResponseData

`func NewUserAttributesIdPatch200ResponseData(id string, name string, label string, dataType string, defaultValue NullableString, intercomAttributeName NullableString, ) *UserAttributesIdPatch200ResponseData`

NewUserAttributesIdPatch200ResponseData instantiates a new UserAttributesIdPatch200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAttributesIdPatch200ResponseDataWithDefaults

`func NewUserAttributesIdPatch200ResponseDataWithDefaults() *UserAttributesIdPatch200ResponseData`

NewUserAttributesIdPatch200ResponseDataWithDefaults instantiates a new UserAttributesIdPatch200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserAttributesIdPatch200ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserAttributesIdPatch200ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserAttributesIdPatch200ResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UserAttributesIdPatch200ResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserAttributesIdPatch200ResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserAttributesIdPatch200ResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *UserAttributesIdPatch200ResponseData) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UserAttributesIdPatch200ResponseData) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UserAttributesIdPatch200ResponseData) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDataType

`func (o *UserAttributesIdPatch200ResponseData) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *UserAttributesIdPatch200ResponseData) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *UserAttributesIdPatch200ResponseData) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetDefaultValue

`func (o *UserAttributesIdPatch200ResponseData) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *UserAttributesIdPatch200ResponseData) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *UserAttributesIdPatch200ResponseData) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.


### SetDefaultValueNil

`func (o *UserAttributesIdPatch200ResponseData) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *UserAttributesIdPatch200ResponseData) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetIntercomAttributeName

`func (o *UserAttributesIdPatch200ResponseData) GetIntercomAttributeName() string`

GetIntercomAttributeName returns the IntercomAttributeName field if non-nil, zero value otherwise.

### GetIntercomAttributeNameOk

`func (o *UserAttributesIdPatch200ResponseData) GetIntercomAttributeNameOk() (*string, bool)`

GetIntercomAttributeNameOk returns a tuple with the IntercomAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntercomAttributeName

`func (o *UserAttributesIdPatch200ResponseData) SetIntercomAttributeName(v string)`

SetIntercomAttributeName sets IntercomAttributeName field to given value.


### SetIntercomAttributeNameNil

`func (o *UserAttributesIdPatch200ResponseData) SetIntercomAttributeNameNil(b bool)`

 SetIntercomAttributeNameNil sets the value for IntercomAttributeName to be an explicit nil

### UnsetIntercomAttributeName
`func (o *UserAttributesIdPatch200ResponseData) UnsetIntercomAttributeName()`

UnsetIntercomAttributeName ensures that no value is present for IntercomAttributeName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


