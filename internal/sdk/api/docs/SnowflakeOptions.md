# SnowflakeOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountIdentifier** | **string** |  | 
**DatabaseOptions** | [**SnowflakeOptionsDatabaseOptions**](SnowflakeOptionsDatabaseOptions.md) |  | 
**UserRole** | Pointer to **string** |  | [optional] 
**AuthenticationOptions** | [**SnowflakeOptionsAuthenticationOptions**](SnowflakeOptionsAuthenticationOptions.md) |  | 

## Methods

### NewSnowflakeOptions

`func NewSnowflakeOptions(accountIdentifier string, databaseOptions SnowflakeOptionsDatabaseOptions, authenticationOptions SnowflakeOptionsAuthenticationOptions, ) *SnowflakeOptions`

NewSnowflakeOptions instantiates a new SnowflakeOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnowflakeOptionsWithDefaults

`func NewSnowflakeOptionsWithDefaults() *SnowflakeOptions`

NewSnowflakeOptionsWithDefaults instantiates a new SnowflakeOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountIdentifier

`func (o *SnowflakeOptions) GetAccountIdentifier() string`

GetAccountIdentifier returns the AccountIdentifier field if non-nil, zero value otherwise.

### GetAccountIdentifierOk

`func (o *SnowflakeOptions) GetAccountIdentifierOk() (*string, bool)`

GetAccountIdentifierOk returns a tuple with the AccountIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIdentifier

`func (o *SnowflakeOptions) SetAccountIdentifier(v string)`

SetAccountIdentifier sets AccountIdentifier field to given value.


### GetDatabaseOptions

`func (o *SnowflakeOptions) GetDatabaseOptions() SnowflakeOptionsDatabaseOptions`

GetDatabaseOptions returns the DatabaseOptions field if non-nil, zero value otherwise.

### GetDatabaseOptionsOk

`func (o *SnowflakeOptions) GetDatabaseOptionsOk() (*SnowflakeOptionsDatabaseOptions, bool)`

GetDatabaseOptionsOk returns a tuple with the DatabaseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseOptions

`func (o *SnowflakeOptions) SetDatabaseOptions(v SnowflakeOptionsDatabaseOptions)`

SetDatabaseOptions sets DatabaseOptions field to given value.


### GetUserRole

`func (o *SnowflakeOptions) GetUserRole() string`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *SnowflakeOptions) GetUserRoleOk() (*string, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *SnowflakeOptions) SetUserRole(v string)`

SetUserRole sets UserRole field to given value.

### HasUserRole

`func (o *SnowflakeOptions) HasUserRole() bool`

HasUserRole returns a boolean if a field has been set.

### GetAuthenticationOptions

`func (o *SnowflakeOptions) GetAuthenticationOptions() SnowflakeOptionsAuthenticationOptions`

GetAuthenticationOptions returns the AuthenticationOptions field if non-nil, zero value otherwise.

### GetAuthenticationOptionsOk

`func (o *SnowflakeOptions) GetAuthenticationOptionsOk() (*SnowflakeOptionsAuthenticationOptions, bool)`

GetAuthenticationOptionsOk returns a tuple with the AuthenticationOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOptions

`func (o *SnowflakeOptions) SetAuthenticationOptions(v SnowflakeOptionsAuthenticationOptions)`

SetAuthenticationOptions sets AuthenticationOptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


