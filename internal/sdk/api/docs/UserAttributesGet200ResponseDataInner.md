# UserAttributesGet200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the user attribute | 
**Name** | **string** | The name of the user attribute | 
**Label** | **string** | The label of the user attribute | 
**DataType** | **string** | The data type of the user attribute. One of types &#x60;string&#x60;, &#x60;number&#x60;, or &#x60;json&#x60; | 
**DefaultValue** | **NullableString** | The default value of the user attribute, assigned to all users without the attribute set | 
**IntercomAttributeName** | **NullableString** | The name of the Intercom user attribute that this attribute should be mapped to | 

## Methods

### NewUserAttributesGet200ResponseDataInner

`func NewUserAttributesGet200ResponseDataInner(id string, name string, label string, dataType string, defaultValue NullableString, intercomAttributeName NullableString, ) *UserAttributesGet200ResponseDataInner`

NewUserAttributesGet200ResponseDataInner instantiates a new UserAttributesGet200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAttributesGet200ResponseDataInnerWithDefaults

`func NewUserAttributesGet200ResponseDataInnerWithDefaults() *UserAttributesGet200ResponseDataInner`

NewUserAttributesGet200ResponseDataInnerWithDefaults instantiates a new UserAttributesGet200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserAttributesGet200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserAttributesGet200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserAttributesGet200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UserAttributesGet200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserAttributesGet200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserAttributesGet200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *UserAttributesGet200ResponseDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UserAttributesGet200ResponseDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UserAttributesGet200ResponseDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDataType

`func (o *UserAttributesGet200ResponseDataInner) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *UserAttributesGet200ResponseDataInner) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *UserAttributesGet200ResponseDataInner) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetDefaultValue

`func (o *UserAttributesGet200ResponseDataInner) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *UserAttributesGet200ResponseDataInner) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *UserAttributesGet200ResponseDataInner) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.


### SetDefaultValueNil

`func (o *UserAttributesGet200ResponseDataInner) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *UserAttributesGet200ResponseDataInner) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetIntercomAttributeName

`func (o *UserAttributesGet200ResponseDataInner) GetIntercomAttributeName() string`

GetIntercomAttributeName returns the IntercomAttributeName field if non-nil, zero value otherwise.

### GetIntercomAttributeNameOk

`func (o *UserAttributesGet200ResponseDataInner) GetIntercomAttributeNameOk() (*string, bool)`

GetIntercomAttributeNameOk returns a tuple with the IntercomAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntercomAttributeName

`func (o *UserAttributesGet200ResponseDataInner) SetIntercomAttributeName(v string)`

SetIntercomAttributeName sets IntercomAttributeName field to given value.


### SetIntercomAttributeNameNil

`func (o *UserAttributesGet200ResponseDataInner) SetIntercomAttributeNameNil(b bool)`

 SetIntercomAttributeNameNil sets the value for IntercomAttributeName to be an explicit nil

### UnsetIntercomAttributeName
`func (o *UserAttributesGet200ResponseDataInner) UnsetIntercomAttributeName()`

UnsetIntercomAttributeName ensures that no value is present for IntercomAttributeName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


