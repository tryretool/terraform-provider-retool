# PostgresOptionsDatabaseOptions

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
**ConnectionString** | Pointer to **string** |  | [optional] 

## Methods

### NewPostgresOptionsDatabaseOptions

`func NewPostgresOptionsDatabaseOptions(host string, port string, ) *PostgresOptionsDatabaseOptions`

NewPostgresOptionsDatabaseOptions instantiates a new PostgresOptionsDatabaseOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresOptionsDatabaseOptionsWithDefaults

`func NewPostgresOptionsDatabaseOptionsWithDefaults() *PostgresOptionsDatabaseOptions`

NewPostgresOptionsDatabaseOptionsWithDefaults instantiates a new PostgresOptionsDatabaseOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *PostgresOptionsDatabaseOptions) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *PostgresOptionsDatabaseOptions) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *PostgresOptionsDatabaseOptions) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *PostgresOptionsDatabaseOptions) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PostgresOptionsDatabaseOptions) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PostgresOptionsDatabaseOptions) SetPort(v string)`

SetPort sets Port field to given value.


### GetName

`func (o *PostgresOptionsDatabaseOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostgresOptionsDatabaseOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostgresOptionsDatabaseOptions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostgresOptionsDatabaseOptions) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUsername

`func (o *PostgresOptionsDatabaseOptions) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PostgresOptionsDatabaseOptions) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PostgresOptionsDatabaseOptions) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PostgresOptionsDatabaseOptions) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *PostgresOptionsDatabaseOptions) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PostgresOptionsDatabaseOptions) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PostgresOptionsDatabaseOptions) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PostgresOptionsDatabaseOptions) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSslSettings

`func (o *PostgresOptionsDatabaseOptions) GetSslSettings() PostgresOptionsDatabaseOptionsAnyOfSslSettings`

GetSslSettings returns the SslSettings field if non-nil, zero value otherwise.

### GetSslSettingsOk

`func (o *PostgresOptionsDatabaseOptions) GetSslSettingsOk() (*PostgresOptionsDatabaseOptionsAnyOfSslSettings, bool)`

GetSslSettingsOk returns a tuple with the SslSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslSettings

`func (o *PostgresOptionsDatabaseOptions) SetSslSettings(v PostgresOptionsDatabaseOptionsAnyOfSslSettings)`

SetSslSettings sets SslSettings field to given value.

### HasSslSettings

`func (o *PostgresOptionsDatabaseOptions) HasSslSettings() bool`

HasSslSettings returns a boolean if a field has been set.

### GetSshTunnelOptions

`func (o *PostgresOptionsDatabaseOptions) GetSshTunnelOptions() PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions`

GetSshTunnelOptions returns the SshTunnelOptions field if non-nil, zero value otherwise.

### GetSshTunnelOptionsOk

`func (o *PostgresOptionsDatabaseOptions) GetSshTunnelOptionsOk() (*PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions, bool)`

GetSshTunnelOptionsOk returns a tuple with the SshTunnelOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshTunnelOptions

`func (o *PostgresOptionsDatabaseOptions) SetSshTunnelOptions(v PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions)`

SetSshTunnelOptions sets SshTunnelOptions field to given value.

### HasSshTunnelOptions

`func (o *PostgresOptionsDatabaseOptions) HasSshTunnelOptions() bool`

HasSshTunnelOptions returns a boolean if a field has been set.

### GetDisableConvertingQueriesToPreparedStatements

`func (o *PostgresOptionsDatabaseOptions) GetDisableConvertingQueriesToPreparedStatements() bool`

GetDisableConvertingQueriesToPreparedStatements returns the DisableConvertingQueriesToPreparedStatements field if non-nil, zero value otherwise.

### GetDisableConvertingQueriesToPreparedStatementsOk

`func (o *PostgresOptionsDatabaseOptions) GetDisableConvertingQueriesToPreparedStatementsOk() (*bool, bool)`

GetDisableConvertingQueriesToPreparedStatementsOk returns a tuple with the DisableConvertingQueriesToPreparedStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableConvertingQueriesToPreparedStatements

`func (o *PostgresOptionsDatabaseOptions) SetDisableConvertingQueriesToPreparedStatements(v bool)`

SetDisableConvertingQueriesToPreparedStatements sets DisableConvertingQueriesToPreparedStatements field to given value.

### HasDisableConvertingQueriesToPreparedStatements

`func (o *PostgresOptionsDatabaseOptions) HasDisableConvertingQueriesToPreparedStatements() bool`

HasDisableConvertingQueriesToPreparedStatements returns a boolean if a field has been set.

### GetShowWriteGuiOnly

`func (o *PostgresOptionsDatabaseOptions) GetShowWriteGuiOnly() bool`

GetShowWriteGuiOnly returns the ShowWriteGuiOnly field if non-nil, zero value otherwise.

### GetShowWriteGuiOnlyOk

`func (o *PostgresOptionsDatabaseOptions) GetShowWriteGuiOnlyOk() (*bool, bool)`

GetShowWriteGuiOnlyOk returns a tuple with the ShowWriteGuiOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowWriteGuiOnly

`func (o *PostgresOptionsDatabaseOptions) SetShowWriteGuiOnly(v bool)`

SetShowWriteGuiOnly sets ShowWriteGuiOnly field to given value.

### HasShowWriteGuiOnly

`func (o *PostgresOptionsDatabaseOptions) HasShowWriteGuiOnly() bool`

HasShowWriteGuiOnly returns a boolean if a field has been set.

### GetConnectionString

`func (o *PostgresOptionsDatabaseOptions) GetConnectionString() string`

GetConnectionString returns the ConnectionString field if non-nil, zero value otherwise.

### GetConnectionStringOk

`func (o *PostgresOptionsDatabaseOptions) GetConnectionStringOk() (*string, bool)`

GetConnectionStringOk returns a tuple with the ConnectionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionString

`func (o *PostgresOptionsDatabaseOptions) SetConnectionString(v string)`

SetConnectionString sets ConnectionString field to given value.

### HasConnectionString

`func (o *PostgresOptionsDatabaseOptions) HasConnectionString() bool`

HasConnectionString returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


