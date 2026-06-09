# UsersGet200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the user. | 
**LegacyId** | **float32** | The legacy ID of the user. | 
**Email** | **string** | The email of the user. | 
**Active** | **bool** | Whether the user is active or not. | 
**CreatedAt** | **time.Time** | The timestamp when the user was created. | 
**LastActive** | **NullableTime** | The timestamp when the user was last active. | 
**FirstName** | **NullableString** | The first name of the user. | 
**LastName** | **NullableString** | The last name of the user. | 
**Metadata** | **map[string]interface{}** | Custom metadata associated with the user. | [default to {}]
**IsAdmin** | **bool** | Whether the user is an admin or not. | 
**UserType** | **string** | The user type. | 
**SeatType** | Pointer to **NullableString** | The user&#39;s seat type. Only meaningful on plans that have the Named Seats feature; on other plans the value is informational. Setting this field via API requires the Named Seats feature. | [optional] 
**TwoFactorAuthEnabled** | **bool** | Whether two-factor authentication is enabled for this user. | 

## Methods

### NewUsersGet200ResponseDataInner

`func NewUsersGet200ResponseDataInner(id string, legacyId float32, email string, active bool, createdAt time.Time, lastActive NullableTime, firstName NullableString, lastName NullableString, metadata map[string]interface{}, isAdmin bool, userType string, twoFactorAuthEnabled bool, ) *UsersGet200ResponseDataInner`

NewUsersGet200ResponseDataInner instantiates a new UsersGet200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersGet200ResponseDataInnerWithDefaults

`func NewUsersGet200ResponseDataInnerWithDefaults() *UsersGet200ResponseDataInner`

NewUsersGet200ResponseDataInnerWithDefaults instantiates a new UsersGet200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UsersGet200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsersGet200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsersGet200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetLegacyId

`func (o *UsersGet200ResponseDataInner) GetLegacyId() float32`

GetLegacyId returns the LegacyId field if non-nil, zero value otherwise.

### GetLegacyIdOk

`func (o *UsersGet200ResponseDataInner) GetLegacyIdOk() (*float32, bool)`

GetLegacyIdOk returns a tuple with the LegacyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyId

`func (o *UsersGet200ResponseDataInner) SetLegacyId(v float32)`

SetLegacyId sets LegacyId field to given value.


### GetEmail

`func (o *UsersGet200ResponseDataInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UsersGet200ResponseDataInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UsersGet200ResponseDataInner) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetActive

`func (o *UsersGet200ResponseDataInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UsersGet200ResponseDataInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UsersGet200ResponseDataInner) SetActive(v bool)`

SetActive sets Active field to given value.


### GetCreatedAt

`func (o *UsersGet200ResponseDataInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UsersGet200ResponseDataInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UsersGet200ResponseDataInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetLastActive

`func (o *UsersGet200ResponseDataInner) GetLastActive() time.Time`

GetLastActive returns the LastActive field if non-nil, zero value otherwise.

### GetLastActiveOk

`func (o *UsersGet200ResponseDataInner) GetLastActiveOk() (*time.Time, bool)`

GetLastActiveOk returns a tuple with the LastActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActive

`func (o *UsersGet200ResponseDataInner) SetLastActive(v time.Time)`

SetLastActive sets LastActive field to given value.


### SetLastActiveNil

`func (o *UsersGet200ResponseDataInner) SetLastActiveNil(b bool)`

 SetLastActiveNil sets the value for LastActive to be an explicit nil

### UnsetLastActive
`func (o *UsersGet200ResponseDataInner) UnsetLastActive()`

UnsetLastActive ensures that no value is present for LastActive, not even an explicit nil
### GetFirstName

`func (o *UsersGet200ResponseDataInner) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UsersGet200ResponseDataInner) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UsersGet200ResponseDataInner) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### SetFirstNameNil

`func (o *UsersGet200ResponseDataInner) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *UsersGet200ResponseDataInner) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *UsersGet200ResponseDataInner) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UsersGet200ResponseDataInner) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UsersGet200ResponseDataInner) SetLastName(v string)`

SetLastName sets LastName field to given value.


### SetLastNameNil

`func (o *UsersGet200ResponseDataInner) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *UsersGet200ResponseDataInner) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetMetadata

`func (o *UsersGet200ResponseDataInner) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UsersGet200ResponseDataInner) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UsersGet200ResponseDataInner) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetIsAdmin

`func (o *UsersGet200ResponseDataInner) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *UsersGet200ResponseDataInner) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *UsersGet200ResponseDataInner) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.


### GetUserType

`func (o *UsersGet200ResponseDataInner) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *UsersGet200ResponseDataInner) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *UsersGet200ResponseDataInner) SetUserType(v string)`

SetUserType sets UserType field to given value.


### GetSeatType

`func (o *UsersGet200ResponseDataInner) GetSeatType() string`

GetSeatType returns the SeatType field if non-nil, zero value otherwise.

### GetSeatTypeOk

`func (o *UsersGet200ResponseDataInner) GetSeatTypeOk() (*string, bool)`

GetSeatTypeOk returns a tuple with the SeatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatType

`func (o *UsersGet200ResponseDataInner) SetSeatType(v string)`

SetSeatType sets SeatType field to given value.

### HasSeatType

`func (o *UsersGet200ResponseDataInner) HasSeatType() bool`

HasSeatType returns a boolean if a field has been set.

### SetSeatTypeNil

`func (o *UsersGet200ResponseDataInner) SetSeatTypeNil(b bool)`

 SetSeatTypeNil sets the value for SeatType to be an explicit nil

### UnsetSeatType
`func (o *UsersGet200ResponseDataInner) UnsetSeatType()`

UnsetSeatType ensures that no value is present for SeatType, not even an explicit nil
### GetTwoFactorAuthEnabled

`func (o *UsersGet200ResponseDataInner) GetTwoFactorAuthEnabled() bool`

GetTwoFactorAuthEnabled returns the TwoFactorAuthEnabled field if non-nil, zero value otherwise.

### GetTwoFactorAuthEnabledOk

`func (o *UsersGet200ResponseDataInner) GetTwoFactorAuthEnabledOk() (*bool, bool)`

GetTwoFactorAuthEnabledOk returns a tuple with the TwoFactorAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuthEnabled

`func (o *UsersGet200ResponseDataInner) SetTwoFactorAuthEnabled(v bool)`

SetTwoFactorAuthEnabled sets TwoFactorAuthEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


