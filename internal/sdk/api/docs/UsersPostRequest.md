# UsersPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email of the user | 
**FirstName** | **string** | The first name of the user | 
**LastName** | **string** | The last name of the user | 
**Active** | Pointer to **bool** | Whether the user is enabled or not | [optional] [default to true]
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] [default to {}]
**UserType** | Pointer to **string** | The user type | [optional] 

## Methods

### NewUsersPostRequest

`func NewUsersPostRequest(email string, firstName string, lastName string, ) *UsersPostRequest`

NewUsersPostRequest instantiates a new UsersPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersPostRequestWithDefaults

`func NewUsersPostRequestWithDefaults() *UsersPostRequest`

NewUsersPostRequestWithDefaults instantiates a new UsersPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UsersPostRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UsersPostRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UsersPostRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *UsersPostRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UsersPostRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UsersPostRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *UsersPostRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UsersPostRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UsersPostRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetActive

`func (o *UsersPostRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UsersPostRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UsersPostRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UsersPostRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetMetadata

`func (o *UsersPostRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UsersPostRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UsersPostRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UsersPostRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetUserType

`func (o *UsersPostRequest) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *UsersPostRequest) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *UsersPostRequest) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *UsersPostRequest) HasUserType() bool`

HasUserType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


