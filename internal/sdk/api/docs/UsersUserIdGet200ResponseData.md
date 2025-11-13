# UsersUserIdGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the user | 
**LegacyId** | **float32** | The legacy id of the user | 
**Email** | **string** | The email of the user | 
**Active** | **bool** | Whether the user is active or not | 
**CreatedAt** | **time.Time** |  | 
**LastActive** | **NullableTime** |  | 
**FirstName** | **NullableString** | The first name of the user | 
**LastName** | **NullableString** | The last name of the user | 
**Metadata** | **map[string]interface{}** |  | [default to {}]
**IsAdmin** | **bool** | Whether the user is an admin or not | 
**UserType** | **string** | The user type | 
**TwoFactorAuthEnabled** | **bool** | Whether two factor authentication is enabled for this user | 

## Methods

### NewUsersUserIdGet200ResponseData

`func NewUsersUserIdGet200ResponseData(id string, legacyId float32, email string, active bool, createdAt time.Time, lastActive NullableTime, firstName NullableString, lastName NullableString, metadata map[string]interface{}, isAdmin bool, userType string, twoFactorAuthEnabled bool, ) *UsersUserIdGet200ResponseData`

NewUsersUserIdGet200ResponseData instantiates a new UsersUserIdGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersUserIdGet200ResponseDataWithDefaults

`func NewUsersUserIdGet200ResponseDataWithDefaults() *UsersUserIdGet200ResponseData`

NewUsersUserIdGet200ResponseDataWithDefaults instantiates a new UsersUserIdGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UsersUserIdGet200ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsersUserIdGet200ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsersUserIdGet200ResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetLegacyId

`func (o *UsersUserIdGet200ResponseData) GetLegacyId() float32`

GetLegacyId returns the LegacyId field if non-nil, zero value otherwise.

### GetLegacyIdOk

`func (o *UsersUserIdGet200ResponseData) GetLegacyIdOk() (*float32, bool)`

GetLegacyIdOk returns a tuple with the LegacyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyId

`func (o *UsersUserIdGet200ResponseData) SetLegacyId(v float32)`

SetLegacyId sets LegacyId field to given value.


### GetEmail

`func (o *UsersUserIdGet200ResponseData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UsersUserIdGet200ResponseData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UsersUserIdGet200ResponseData) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetActive

`func (o *UsersUserIdGet200ResponseData) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UsersUserIdGet200ResponseData) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UsersUserIdGet200ResponseData) SetActive(v bool)`

SetActive sets Active field to given value.


### GetCreatedAt

`func (o *UsersUserIdGet200ResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UsersUserIdGet200ResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UsersUserIdGet200ResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetLastActive

`func (o *UsersUserIdGet200ResponseData) GetLastActive() time.Time`

GetLastActive returns the LastActive field if non-nil, zero value otherwise.

### GetLastActiveOk

`func (o *UsersUserIdGet200ResponseData) GetLastActiveOk() (*time.Time, bool)`

GetLastActiveOk returns a tuple with the LastActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActive

`func (o *UsersUserIdGet200ResponseData) SetLastActive(v time.Time)`

SetLastActive sets LastActive field to given value.


### SetLastActiveNil

`func (o *UsersUserIdGet200ResponseData) SetLastActiveNil(b bool)`

 SetLastActiveNil sets the value for LastActive to be an explicit nil

### UnsetLastActive
`func (o *UsersUserIdGet200ResponseData) UnsetLastActive()`

UnsetLastActive ensures that no value is present for LastActive, not even an explicit nil
### GetFirstName

`func (o *UsersUserIdGet200ResponseData) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UsersUserIdGet200ResponseData) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UsersUserIdGet200ResponseData) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### SetFirstNameNil

`func (o *UsersUserIdGet200ResponseData) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *UsersUserIdGet200ResponseData) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *UsersUserIdGet200ResponseData) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UsersUserIdGet200ResponseData) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UsersUserIdGet200ResponseData) SetLastName(v string)`

SetLastName sets LastName field to given value.


### SetLastNameNil

`func (o *UsersUserIdGet200ResponseData) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *UsersUserIdGet200ResponseData) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetMetadata

`func (o *UsersUserIdGet200ResponseData) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UsersUserIdGet200ResponseData) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UsersUserIdGet200ResponseData) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetIsAdmin

`func (o *UsersUserIdGet200ResponseData) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *UsersUserIdGet200ResponseData) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *UsersUserIdGet200ResponseData) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.


### GetUserType

`func (o *UsersUserIdGet200ResponseData) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *UsersUserIdGet200ResponseData) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *UsersUserIdGet200ResponseData) SetUserType(v string)`

SetUserType sets UserType field to given value.


### GetTwoFactorAuthEnabled

`func (o *UsersUserIdGet200ResponseData) GetTwoFactorAuthEnabled() bool`

GetTwoFactorAuthEnabled returns the TwoFactorAuthEnabled field if non-nil, zero value otherwise.

### GetTwoFactorAuthEnabledOk

`func (o *UsersUserIdGet200ResponseData) GetTwoFactorAuthEnabledOk() (*bool, bool)`

GetTwoFactorAuthEnabledOk returns a tuple with the TwoFactorAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuthEnabled

`func (o *UsersUserIdGet200ResponseData) SetTwoFactorAuthEnabled(v bool)`

SetTwoFactorAuthEnabled sets TwoFactorAuthEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


