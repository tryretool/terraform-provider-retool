# MySQLOptionsDatabaseOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** |  | 
**Port** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**SslSettings** | Pointer to [**PostgresOptionsDatabaseOptionsAnyOfSslSettings**](PostgresOptionsDatabaseOptionsAnyOfSslSettings.md) |  | [optional] 
**SshTunnelOptions** | Pointer to [**PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions**](PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions.md) |  | [optional] 
**DisableConvertingQueriesToPreparedStatements** | Pointer to **bool** | This allows you to use Javascript to dynamically generate SQL but also turns off SQL injection protection. | [optional] 
**ShowWriteGuiOnly** | Pointer to **bool** | This allows you to enable writing via only the restrictive GUI query editor. | [optional] 
**UseDynamicDatabaseNames** | Pointer to **bool** | Enable this to allow the Database Name to be overridden by a dynamically generated value. This allows using Retool with a database that has been sharded into several different databases. | [optional] 
**UseDynamicDatabaseHosts** | Pointer to **bool** | Enable this to allow the Database Host to be overridden by a dynamically generated value. This allows using Retool with several different databases. | [optional] 
**ConvertMysqlDatesToJavascript** | Pointer to **bool** | This allows you to turn your MySQL date strings into Javascript Date objects. | [optional] [default to true]

## Methods

### NewMySQLOptionsDatabaseOptions

`func NewMySQLOptionsDatabaseOptions(host string, port string, ) *MySQLOptionsDatabaseOptions`

NewMySQLOptionsDatabaseOptions instantiates a new MySQLOptionsDatabaseOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMySQLOptionsDatabaseOptionsWithDefaults

`func NewMySQLOptionsDatabaseOptionsWithDefaults() *MySQLOptionsDatabaseOptions`

NewMySQLOptionsDatabaseOptionsWithDefaults instantiates a new MySQLOptionsDatabaseOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *MySQLOptionsDatabaseOptions) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *MySQLOptionsDatabaseOptions) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *MySQLOptionsDatabaseOptions) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *MySQLOptionsDatabaseOptions) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *MySQLOptionsDatabaseOptions) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *MySQLOptionsDatabaseOptions) SetPort(v string)`

SetPort sets Port field to given value.


### GetName

`func (o *MySQLOptionsDatabaseOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MySQLOptionsDatabaseOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MySQLOptionsDatabaseOptions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MySQLOptionsDatabaseOptions) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUsername

`func (o *MySQLOptionsDatabaseOptions) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *MySQLOptionsDatabaseOptions) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *MySQLOptionsDatabaseOptions) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *MySQLOptionsDatabaseOptions) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *MySQLOptionsDatabaseOptions) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MySQLOptionsDatabaseOptions) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MySQLOptionsDatabaseOptions) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *MySQLOptionsDatabaseOptions) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSslSettings

`func (o *MySQLOptionsDatabaseOptions) GetSslSettings() PostgresOptionsDatabaseOptionsAnyOfSslSettings`

GetSslSettings returns the SslSettings field if non-nil, zero value otherwise.

### GetSslSettingsOk

`func (o *MySQLOptionsDatabaseOptions) GetSslSettingsOk() (*PostgresOptionsDatabaseOptionsAnyOfSslSettings, bool)`

GetSslSettingsOk returns a tuple with the SslSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslSettings

`func (o *MySQLOptionsDatabaseOptions) SetSslSettings(v PostgresOptionsDatabaseOptionsAnyOfSslSettings)`

SetSslSettings sets SslSettings field to given value.

### HasSslSettings

`func (o *MySQLOptionsDatabaseOptions) HasSslSettings() bool`

HasSslSettings returns a boolean if a field has been set.

### GetSshTunnelOptions

`func (o *MySQLOptionsDatabaseOptions) GetSshTunnelOptions() PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions`

GetSshTunnelOptions returns the SshTunnelOptions field if non-nil, zero value otherwise.

### GetSshTunnelOptionsOk

`func (o *MySQLOptionsDatabaseOptions) GetSshTunnelOptionsOk() (*PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions, bool)`

GetSshTunnelOptionsOk returns a tuple with the SshTunnelOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshTunnelOptions

`func (o *MySQLOptionsDatabaseOptions) SetSshTunnelOptions(v PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions)`

SetSshTunnelOptions sets SshTunnelOptions field to given value.

### HasSshTunnelOptions

`func (o *MySQLOptionsDatabaseOptions) HasSshTunnelOptions() bool`

HasSshTunnelOptions returns a boolean if a field has been set.

### GetDisableConvertingQueriesToPreparedStatements

`func (o *MySQLOptionsDatabaseOptions) GetDisableConvertingQueriesToPreparedStatements() bool`

GetDisableConvertingQueriesToPreparedStatements returns the DisableConvertingQueriesToPreparedStatements field if non-nil, zero value otherwise.

### GetDisableConvertingQueriesToPreparedStatementsOk

`func (o *MySQLOptionsDatabaseOptions) GetDisableConvertingQueriesToPreparedStatementsOk() (*bool, bool)`

GetDisableConvertingQueriesToPreparedStatementsOk returns a tuple with the DisableConvertingQueriesToPreparedStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableConvertingQueriesToPreparedStatements

`func (o *MySQLOptionsDatabaseOptions) SetDisableConvertingQueriesToPreparedStatements(v bool)`

SetDisableConvertingQueriesToPreparedStatements sets DisableConvertingQueriesToPreparedStatements field to given value.

### HasDisableConvertingQueriesToPreparedStatements

`func (o *MySQLOptionsDatabaseOptions) HasDisableConvertingQueriesToPreparedStatements() bool`

HasDisableConvertingQueriesToPreparedStatements returns a boolean if a field has been set.

### GetShowWriteGuiOnly

`func (o *MySQLOptionsDatabaseOptions) GetShowWriteGuiOnly() bool`

GetShowWriteGuiOnly returns the ShowWriteGuiOnly field if non-nil, zero value otherwise.

### GetShowWriteGuiOnlyOk

`func (o *MySQLOptionsDatabaseOptions) GetShowWriteGuiOnlyOk() (*bool, bool)`

GetShowWriteGuiOnlyOk returns a tuple with the ShowWriteGuiOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowWriteGuiOnly

`func (o *MySQLOptionsDatabaseOptions) SetShowWriteGuiOnly(v bool)`

SetShowWriteGuiOnly sets ShowWriteGuiOnly field to given value.

### HasShowWriteGuiOnly

`func (o *MySQLOptionsDatabaseOptions) HasShowWriteGuiOnly() bool`

HasShowWriteGuiOnly returns a boolean if a field has been set.

### GetUseDynamicDatabaseNames

`func (o *MySQLOptionsDatabaseOptions) GetUseDynamicDatabaseNames() bool`

GetUseDynamicDatabaseNames returns the UseDynamicDatabaseNames field if non-nil, zero value otherwise.

### GetUseDynamicDatabaseNamesOk

`func (o *MySQLOptionsDatabaseOptions) GetUseDynamicDatabaseNamesOk() (*bool, bool)`

GetUseDynamicDatabaseNamesOk returns a tuple with the UseDynamicDatabaseNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDynamicDatabaseNames

`func (o *MySQLOptionsDatabaseOptions) SetUseDynamicDatabaseNames(v bool)`

SetUseDynamicDatabaseNames sets UseDynamicDatabaseNames field to given value.

### HasUseDynamicDatabaseNames

`func (o *MySQLOptionsDatabaseOptions) HasUseDynamicDatabaseNames() bool`

HasUseDynamicDatabaseNames returns a boolean if a field has been set.

### GetUseDynamicDatabaseHosts

`func (o *MySQLOptionsDatabaseOptions) GetUseDynamicDatabaseHosts() bool`

GetUseDynamicDatabaseHosts returns the UseDynamicDatabaseHosts field if non-nil, zero value otherwise.

### GetUseDynamicDatabaseHostsOk

`func (o *MySQLOptionsDatabaseOptions) GetUseDynamicDatabaseHostsOk() (*bool, bool)`

GetUseDynamicDatabaseHostsOk returns a tuple with the UseDynamicDatabaseHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDynamicDatabaseHosts

`func (o *MySQLOptionsDatabaseOptions) SetUseDynamicDatabaseHosts(v bool)`

SetUseDynamicDatabaseHosts sets UseDynamicDatabaseHosts field to given value.

### HasUseDynamicDatabaseHosts

`func (o *MySQLOptionsDatabaseOptions) HasUseDynamicDatabaseHosts() bool`

HasUseDynamicDatabaseHosts returns a boolean if a field has been set.

### GetConvertMysqlDatesToJavascript

`func (o *MySQLOptionsDatabaseOptions) GetConvertMysqlDatesToJavascript() bool`

GetConvertMysqlDatesToJavascript returns the ConvertMysqlDatesToJavascript field if non-nil, zero value otherwise.

### GetConvertMysqlDatesToJavascriptOk

`func (o *MySQLOptionsDatabaseOptions) GetConvertMysqlDatesToJavascriptOk() (*bool, bool)`

GetConvertMysqlDatesToJavascriptOk returns a tuple with the ConvertMysqlDatesToJavascript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertMysqlDatesToJavascript

`func (o *MySQLOptionsDatabaseOptions) SetConvertMysqlDatesToJavascript(v bool)`

SetConvertMysqlDatesToJavascript sets ConvertMysqlDatesToJavascript field to given value.

### HasConvertMysqlDatesToJavascript

`func (o *MySQLOptionsDatabaseOptions) HasConvertMysqlDatesToJavascript() bool`

HasConvertMysqlDatesToJavascript returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


