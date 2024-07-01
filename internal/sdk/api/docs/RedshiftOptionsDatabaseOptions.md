# RedshiftOptionsDatabaseOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** |  | 
**Port** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**SslSettings** | Pointer to [**RedshiftOptionsDatabaseOptionsSslSettings**](RedshiftOptionsDatabaseOptionsSslSettings.md) |  | [optional] 
**DisableConvertingQueriesToPreparedStatements** | Pointer to **bool** | This allows you to use Javascript to dynamically generate SQL but also turns off SQL injection protection. | [optional] 
**ShowWriteGuiOnly** | Pointer to **bool** | This allows you to enable writing via only the restrictive GUI query editor. | [optional] 

## Methods

### NewRedshiftOptionsDatabaseOptions

`func NewRedshiftOptionsDatabaseOptions(host string, port string, ) *RedshiftOptionsDatabaseOptions`

NewRedshiftOptionsDatabaseOptions instantiates a new RedshiftOptionsDatabaseOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedshiftOptionsDatabaseOptionsWithDefaults

`func NewRedshiftOptionsDatabaseOptionsWithDefaults() *RedshiftOptionsDatabaseOptions`

NewRedshiftOptionsDatabaseOptionsWithDefaults instantiates a new RedshiftOptionsDatabaseOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *RedshiftOptionsDatabaseOptions) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RedshiftOptionsDatabaseOptions) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RedshiftOptionsDatabaseOptions) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *RedshiftOptionsDatabaseOptions) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RedshiftOptionsDatabaseOptions) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RedshiftOptionsDatabaseOptions) SetPort(v string)`

SetPort sets Port field to given value.


### GetName

`func (o *RedshiftOptionsDatabaseOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RedshiftOptionsDatabaseOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RedshiftOptionsDatabaseOptions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RedshiftOptionsDatabaseOptions) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUsername

`func (o *RedshiftOptionsDatabaseOptions) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RedshiftOptionsDatabaseOptions) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RedshiftOptionsDatabaseOptions) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *RedshiftOptionsDatabaseOptions) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *RedshiftOptionsDatabaseOptions) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RedshiftOptionsDatabaseOptions) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RedshiftOptionsDatabaseOptions) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RedshiftOptionsDatabaseOptions) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSslSettings

`func (o *RedshiftOptionsDatabaseOptions) GetSslSettings() RedshiftOptionsDatabaseOptionsSslSettings`

GetSslSettings returns the SslSettings field if non-nil, zero value otherwise.

### GetSslSettingsOk

`func (o *RedshiftOptionsDatabaseOptions) GetSslSettingsOk() (*RedshiftOptionsDatabaseOptionsSslSettings, bool)`

GetSslSettingsOk returns a tuple with the SslSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslSettings

`func (o *RedshiftOptionsDatabaseOptions) SetSslSettings(v RedshiftOptionsDatabaseOptionsSslSettings)`

SetSslSettings sets SslSettings field to given value.

### HasSslSettings

`func (o *RedshiftOptionsDatabaseOptions) HasSslSettings() bool`

HasSslSettings returns a boolean if a field has been set.

### GetDisableConvertingQueriesToPreparedStatements

`func (o *RedshiftOptionsDatabaseOptions) GetDisableConvertingQueriesToPreparedStatements() bool`

GetDisableConvertingQueriesToPreparedStatements returns the DisableConvertingQueriesToPreparedStatements field if non-nil, zero value otherwise.

### GetDisableConvertingQueriesToPreparedStatementsOk

`func (o *RedshiftOptionsDatabaseOptions) GetDisableConvertingQueriesToPreparedStatementsOk() (*bool, bool)`

GetDisableConvertingQueriesToPreparedStatementsOk returns a tuple with the DisableConvertingQueriesToPreparedStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableConvertingQueriesToPreparedStatements

`func (o *RedshiftOptionsDatabaseOptions) SetDisableConvertingQueriesToPreparedStatements(v bool)`

SetDisableConvertingQueriesToPreparedStatements sets DisableConvertingQueriesToPreparedStatements field to given value.

### HasDisableConvertingQueriesToPreparedStatements

`func (o *RedshiftOptionsDatabaseOptions) HasDisableConvertingQueriesToPreparedStatements() bool`

HasDisableConvertingQueriesToPreparedStatements returns a boolean if a field has been set.

### GetShowWriteGuiOnly

`func (o *RedshiftOptionsDatabaseOptions) GetShowWriteGuiOnly() bool`

GetShowWriteGuiOnly returns the ShowWriteGuiOnly field if non-nil, zero value otherwise.

### GetShowWriteGuiOnlyOk

`func (o *RedshiftOptionsDatabaseOptions) GetShowWriteGuiOnlyOk() (*bool, bool)`

GetShowWriteGuiOnlyOk returns a tuple with the ShowWriteGuiOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowWriteGuiOnly

`func (o *RedshiftOptionsDatabaseOptions) SetShowWriteGuiOnly(v bool)`

SetShowWriteGuiOnly sets ShowWriteGuiOnly field to given value.

### HasShowWriteGuiOnly

`func (o *RedshiftOptionsDatabaseOptions) HasShowWriteGuiOnly() bool`

HasShowWriteGuiOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


