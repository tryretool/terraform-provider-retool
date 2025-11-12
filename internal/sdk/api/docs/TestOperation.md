# TestOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | **string** |  | 
**Path** | **string** |  | 
**Value** | Pointer to **interface{}** | A JSON value | [optional] 

## Methods

### NewTestOperation

`func NewTestOperation(op string, path string, ) *TestOperation`

NewTestOperation instantiates a new TestOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestOperationWithDefaults

`func NewTestOperationWithDefaults() *TestOperation`

NewTestOperationWithDefaults instantiates a new TestOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *TestOperation) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *TestOperation) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *TestOperation) SetOp(v string)`

SetOp sets Op field to given value.


### GetPath

`func (o *TestOperation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *TestOperation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *TestOperation) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *TestOperation) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TestOperation) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TestOperation) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *TestOperation) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *TestOperation) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *TestOperation) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


