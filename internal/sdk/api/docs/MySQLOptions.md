# MySQLOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseOptions** | [**MySQLOptionsDatabaseOptions**](MySQLOptionsDatabaseOptions.md) |  | 

## Methods

### NewMySQLOptions

`func NewMySQLOptions(databaseOptions MySQLOptionsDatabaseOptions, ) *MySQLOptions`

NewMySQLOptions instantiates a new MySQLOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMySQLOptionsWithDefaults

`func NewMySQLOptionsWithDefaults() *MySQLOptions`

NewMySQLOptionsWithDefaults instantiates a new MySQLOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseOptions

`func (o *MySQLOptions) GetDatabaseOptions() MySQLOptionsDatabaseOptions`

GetDatabaseOptions returns the DatabaseOptions field if non-nil, zero value otherwise.

### GetDatabaseOptionsOk

`func (o *MySQLOptions) GetDatabaseOptionsOk() (*MySQLOptionsDatabaseOptions, bool)`

GetDatabaseOptionsOk returns a tuple with the DatabaseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseOptions

`func (o *MySQLOptions) SetDatabaseOptions(v MySQLOptionsDatabaseOptions)`

SetDatabaseOptions sets DatabaseOptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


