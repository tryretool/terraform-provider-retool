# ReplaceOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | **string** |  | 
**Path** | **string** |  | 
**Value** | Pointer to **interface{}** | A JSON value | [optional] 

## Methods

### NewReplaceOperation

`func NewReplaceOperation(op string, path string, ) *ReplaceOperation`

NewReplaceOperation instantiates a new ReplaceOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplaceOperationWithDefaults

`func NewReplaceOperationWithDefaults() *ReplaceOperation`

NewReplaceOperationWithDefaults instantiates a new ReplaceOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *ReplaceOperation) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *ReplaceOperation) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *ReplaceOperation) SetOp(v string)`

SetOp sets Op field to given value.


### GetPath

`func (o *ReplaceOperation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ReplaceOperation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ReplaceOperation) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *ReplaceOperation) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ReplaceOperation) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ReplaceOperation) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *ReplaceOperation) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ReplaceOperation) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ReplaceOperation) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


