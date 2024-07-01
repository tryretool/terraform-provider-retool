# SnowflakeOptionsDatabaseOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Schema** | Pointer to **string** |  | [optional] 
**Warehouse** | Pointer to **string** |  | [optional] 
**DisableConvertingQueriesToPreparedStatements** | Pointer to **bool** | This allows you to use Javascript to dynamically generate SQL but also turns off SQL injection protection. | [optional] 
**ShowWriteGuiOnly** | Pointer to **bool** | This allows you to enable writing via only the restrictive GUI query editor. | [optional] 

## Methods

### NewSnowflakeOptionsDatabaseOptions

`func NewSnowflakeOptionsDatabaseOptions(name string, ) *SnowflakeOptionsDatabaseOptions`

NewSnowflakeOptionsDatabaseOptions instantiates a new SnowflakeOptionsDatabaseOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnowflakeOptionsDatabaseOptionsWithDefaults

`func NewSnowflakeOptionsDatabaseOptionsWithDefaults() *SnowflakeOptionsDatabaseOptions`

NewSnowflakeOptionsDatabaseOptionsWithDefaults instantiates a new SnowflakeOptionsDatabaseOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SnowflakeOptionsDatabaseOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SnowflakeOptionsDatabaseOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SnowflakeOptionsDatabaseOptions) SetName(v string)`

SetName sets Name field to given value.


### GetSchema

`func (o *SnowflakeOptionsDatabaseOptions) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *SnowflakeOptionsDatabaseOptions) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *SnowflakeOptionsDatabaseOptions) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *SnowflakeOptionsDatabaseOptions) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetWarehouse

`func (o *SnowflakeOptionsDatabaseOptions) GetWarehouse() string`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *SnowflakeOptionsDatabaseOptions) GetWarehouseOk() (*string, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *SnowflakeOptionsDatabaseOptions) SetWarehouse(v string)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *SnowflakeOptionsDatabaseOptions) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.

### GetDisableConvertingQueriesToPreparedStatements

`func (o *SnowflakeOptionsDatabaseOptions) GetDisableConvertingQueriesToPreparedStatements() bool`

GetDisableConvertingQueriesToPreparedStatements returns the DisableConvertingQueriesToPreparedStatements field if non-nil, zero value otherwise.

### GetDisableConvertingQueriesToPreparedStatementsOk

`func (o *SnowflakeOptionsDatabaseOptions) GetDisableConvertingQueriesToPreparedStatementsOk() (*bool, bool)`

GetDisableConvertingQueriesToPreparedStatementsOk returns a tuple with the DisableConvertingQueriesToPreparedStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableConvertingQueriesToPreparedStatements

`func (o *SnowflakeOptionsDatabaseOptions) SetDisableConvertingQueriesToPreparedStatements(v bool)`

SetDisableConvertingQueriesToPreparedStatements sets DisableConvertingQueriesToPreparedStatements field to given value.

### HasDisableConvertingQueriesToPreparedStatements

`func (o *SnowflakeOptionsDatabaseOptions) HasDisableConvertingQueriesToPreparedStatements() bool`

HasDisableConvertingQueriesToPreparedStatements returns a boolean if a field has been set.

### GetShowWriteGuiOnly

`func (o *SnowflakeOptionsDatabaseOptions) GetShowWriteGuiOnly() bool`

GetShowWriteGuiOnly returns the ShowWriteGuiOnly field if non-nil, zero value otherwise.

### GetShowWriteGuiOnlyOk

`func (o *SnowflakeOptionsDatabaseOptions) GetShowWriteGuiOnlyOk() (*bool, bool)`

GetShowWriteGuiOnlyOk returns a tuple with the ShowWriteGuiOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowWriteGuiOnly

`func (o *SnowflakeOptionsDatabaseOptions) SetShowWriteGuiOnly(v bool)`

SetShowWriteGuiOnly sets ShowWriteGuiOnly field to given value.

### HasShowWriteGuiOnly

`func (o *SnowflakeOptionsDatabaseOptions) HasShowWriteGuiOnly() bool`

HasShowWriteGuiOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


