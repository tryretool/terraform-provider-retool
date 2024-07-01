# PostgresOptionsDatabaseOptionsAnyOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionString** | Pointer to **string** |  | [optional] 
**DisableConvertingQueriesToPreparedStatements** | Pointer to **bool** | This allows you to use Javascript to dynamically generate SQL but also turns off SQL injection protection. | [optional] 
**ShowWriteGuiOnly** | Pointer to **bool** | This allows you to enable writing via only the restrictive GUI query editor. | [optional] 

## Methods

### NewPostgresOptionsDatabaseOptionsAnyOf1

`func NewPostgresOptionsDatabaseOptionsAnyOf1() *PostgresOptionsDatabaseOptionsAnyOf1`

NewPostgresOptionsDatabaseOptionsAnyOf1 instantiates a new PostgresOptionsDatabaseOptionsAnyOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresOptionsDatabaseOptionsAnyOf1WithDefaults

`func NewPostgresOptionsDatabaseOptionsAnyOf1WithDefaults() *PostgresOptionsDatabaseOptionsAnyOf1`

NewPostgresOptionsDatabaseOptionsAnyOf1WithDefaults instantiates a new PostgresOptionsDatabaseOptionsAnyOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionString

`func (o *PostgresOptionsDatabaseOptionsAnyOf1) GetConnectionString() string`

GetConnectionString returns the ConnectionString field if non-nil, zero value otherwise.

### GetConnectionStringOk

`func (o *PostgresOptionsDatabaseOptionsAnyOf1) GetConnectionStringOk() (*string, bool)`

GetConnectionStringOk returns a tuple with the ConnectionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionString

`func (o *PostgresOptionsDatabaseOptionsAnyOf1) SetConnectionString(v string)`

SetConnectionString sets ConnectionString field to given value.

### HasConnectionString

`func (o *PostgresOptionsDatabaseOptionsAnyOf1) HasConnectionString() bool`

HasConnectionString returns a boolean if a field has been set.

### GetDisableConvertingQueriesToPreparedStatements

`func (o *PostgresOptionsDatabaseOptionsAnyOf1) GetDisableConvertingQueriesToPreparedStatements() bool`

GetDisableConvertingQueriesToPreparedStatements returns the DisableConvertingQueriesToPreparedStatements field if non-nil, zero value otherwise.

### GetDisableConvertingQueriesToPreparedStatementsOk

`func (o *PostgresOptionsDatabaseOptionsAnyOf1) GetDisableConvertingQueriesToPreparedStatementsOk() (*bool, bool)`

GetDisableConvertingQueriesToPreparedStatementsOk returns a tuple with the DisableConvertingQueriesToPreparedStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableConvertingQueriesToPreparedStatements

`func (o *PostgresOptionsDatabaseOptionsAnyOf1) SetDisableConvertingQueriesToPreparedStatements(v bool)`

SetDisableConvertingQueriesToPreparedStatements sets DisableConvertingQueriesToPreparedStatements field to given value.

### HasDisableConvertingQueriesToPreparedStatements

`func (o *PostgresOptionsDatabaseOptionsAnyOf1) HasDisableConvertingQueriesToPreparedStatements() bool`

HasDisableConvertingQueriesToPreparedStatements returns a boolean if a field has been set.

### GetShowWriteGuiOnly

`func (o *PostgresOptionsDatabaseOptionsAnyOf1) GetShowWriteGuiOnly() bool`

GetShowWriteGuiOnly returns the ShowWriteGuiOnly field if non-nil, zero value otherwise.

### GetShowWriteGuiOnlyOk

`func (o *PostgresOptionsDatabaseOptionsAnyOf1) GetShowWriteGuiOnlyOk() (*bool, bool)`

GetShowWriteGuiOnlyOk returns a tuple with the ShowWriteGuiOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowWriteGuiOnly

`func (o *PostgresOptionsDatabaseOptionsAnyOf1) SetShowWriteGuiOnly(v bool)`

SetShowWriteGuiOnly sets ShowWriteGuiOnly field to given value.

### HasShowWriteGuiOnly

`func (o *PostgresOptionsDatabaseOptionsAnyOf1) HasShowWriteGuiOnly() bool`

HasShowWriteGuiOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


