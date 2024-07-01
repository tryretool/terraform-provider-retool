# PostgresOptionsDatabaseOptionsAnyOf

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

## Methods

### NewPostgresOptionsDatabaseOptionsAnyOf

`func NewPostgresOptionsDatabaseOptionsAnyOf(host string, port string, ) *PostgresOptionsDatabaseOptionsAnyOf`

NewPostgresOptionsDatabaseOptionsAnyOf instantiates a new PostgresOptionsDatabaseOptionsAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresOptionsDatabaseOptionsAnyOfWithDefaults

`func NewPostgresOptionsDatabaseOptionsAnyOfWithDefaults() *PostgresOptionsDatabaseOptionsAnyOf`

NewPostgresOptionsDatabaseOptionsAnyOfWithDefaults instantiates a new PostgresOptionsDatabaseOptionsAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *PostgresOptionsDatabaseOptionsAnyOf) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *PostgresOptionsDatabaseOptionsAnyOf) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *PostgresOptionsDatabaseOptionsAnyOf) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *PostgresOptionsDatabaseOptionsAnyOf) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PostgresOptionsDatabaseOptionsAnyOf) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PostgresOptionsDatabaseOptionsAnyOf) SetPort(v string)`

SetPort sets Port field to given value.


### GetName

`func (o *PostgresOptionsDatabaseOptionsAnyOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostgresOptionsDatabaseOptionsAnyOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostgresOptionsDatabaseOptionsAnyOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostgresOptionsDatabaseOptionsAnyOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUsername

`func (o *PostgresOptionsDatabaseOptionsAnyOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PostgresOptionsDatabaseOptionsAnyOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PostgresOptionsDatabaseOptionsAnyOf) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PostgresOptionsDatabaseOptionsAnyOf) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *PostgresOptionsDatabaseOptionsAnyOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PostgresOptionsDatabaseOptionsAnyOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PostgresOptionsDatabaseOptionsAnyOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PostgresOptionsDatabaseOptionsAnyOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSslSettings

`func (o *PostgresOptionsDatabaseOptionsAnyOf) GetSslSettings() PostgresOptionsDatabaseOptionsAnyOfSslSettings`

GetSslSettings returns the SslSettings field if non-nil, zero value otherwise.

### GetSslSettingsOk

`func (o *PostgresOptionsDatabaseOptionsAnyOf) GetSslSettingsOk() (*PostgresOptionsDatabaseOptionsAnyOfSslSettings, bool)`

GetSslSettingsOk returns a tuple with the SslSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslSettings

`func (o *PostgresOptionsDatabaseOptionsAnyOf) SetSslSettings(v PostgresOptionsDatabaseOptionsAnyOfSslSettings)`

SetSslSettings sets SslSettings field to given value.

### HasSslSettings

`func (o *PostgresOptionsDatabaseOptionsAnyOf) HasSslSettings() bool`

HasSslSettings returns a boolean if a field has been set.

### GetSshTunnelOptions

`func (o *PostgresOptionsDatabaseOptionsAnyOf) GetSshTunnelOptions() PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions`

GetSshTunnelOptions returns the SshTunnelOptions field if non-nil, zero value otherwise.

### GetSshTunnelOptionsOk

`func (o *PostgresOptionsDatabaseOptionsAnyOf) GetSshTunnelOptionsOk() (*PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions, bool)`

GetSshTunnelOptionsOk returns a tuple with the SshTunnelOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshTunnelOptions

`func (o *PostgresOptionsDatabaseOptionsAnyOf) SetSshTunnelOptions(v PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions)`

SetSshTunnelOptions sets SshTunnelOptions field to given value.

### HasSshTunnelOptions

`func (o *PostgresOptionsDatabaseOptionsAnyOf) HasSshTunnelOptions() bool`

HasSshTunnelOptions returns a boolean if a field has been set.

### GetDisableConvertingQueriesToPreparedStatements

`func (o *PostgresOptionsDatabaseOptionsAnyOf) GetDisableConvertingQueriesToPreparedStatements() bool`

GetDisableConvertingQueriesToPreparedStatements returns the DisableConvertingQueriesToPreparedStatements field if non-nil, zero value otherwise.

### GetDisableConvertingQueriesToPreparedStatementsOk

`func (o *PostgresOptionsDatabaseOptionsAnyOf) GetDisableConvertingQueriesToPreparedStatementsOk() (*bool, bool)`

GetDisableConvertingQueriesToPreparedStatementsOk returns a tuple with the DisableConvertingQueriesToPreparedStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableConvertingQueriesToPreparedStatements

`func (o *PostgresOptionsDatabaseOptionsAnyOf) SetDisableConvertingQueriesToPreparedStatements(v bool)`

SetDisableConvertingQueriesToPreparedStatements sets DisableConvertingQueriesToPreparedStatements field to given value.

### HasDisableConvertingQueriesToPreparedStatements

`func (o *PostgresOptionsDatabaseOptionsAnyOf) HasDisableConvertingQueriesToPreparedStatements() bool`

HasDisableConvertingQueriesToPreparedStatements returns a boolean if a field has been set.

### GetShowWriteGuiOnly

`func (o *PostgresOptionsDatabaseOptionsAnyOf) GetShowWriteGuiOnly() bool`

GetShowWriteGuiOnly returns the ShowWriteGuiOnly field if non-nil, zero value otherwise.

### GetShowWriteGuiOnlyOk

`func (o *PostgresOptionsDatabaseOptionsAnyOf) GetShowWriteGuiOnlyOk() (*bool, bool)`

GetShowWriteGuiOnlyOk returns a tuple with the ShowWriteGuiOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowWriteGuiOnly

`func (o *PostgresOptionsDatabaseOptionsAnyOf) SetShowWriteGuiOnly(v bool)`

SetShowWriteGuiOnly sets ShowWriteGuiOnly field to given value.

### HasShowWriteGuiOnly

`func (o *PostgresOptionsDatabaseOptionsAnyOf) HasShowWriteGuiOnly() bool`

HasShowWriteGuiOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


