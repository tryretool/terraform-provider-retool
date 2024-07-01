# CustomComponentLibrary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** | How the library will be referred to in Toolscript, and other code based usages. | 
**Description** | **NullableString** |  | 
**Label** | **string** | How the library will be referred to in the Retool UI. Should be human readable. | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewCustomComponentLibrary

`func NewCustomComponentLibrary(id string, name string, description NullableString, label string, createdAt time.Time, updatedAt time.Time, ) *CustomComponentLibrary`

NewCustomComponentLibrary instantiates a new CustomComponentLibrary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomComponentLibraryWithDefaults

`func NewCustomComponentLibraryWithDefaults() *CustomComponentLibrary`

NewCustomComponentLibraryWithDefaults instantiates a new CustomComponentLibrary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomComponentLibrary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomComponentLibrary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomComponentLibrary) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CustomComponentLibrary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomComponentLibrary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomComponentLibrary) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CustomComponentLibrary) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomComponentLibrary) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomComponentLibrary) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *CustomComponentLibrary) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CustomComponentLibrary) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLabel

`func (o *CustomComponentLibrary) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CustomComponentLibrary) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CustomComponentLibrary) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetCreatedAt

`func (o *CustomComponentLibrary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomComponentLibrary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomComponentLibrary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CustomComponentLibrary) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CustomComponentLibrary) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CustomComponentLibrary) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


