# UsersUserIdUserAttributesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the user attribute (must match an existing attribute exactly) | 
**Value** | **NullableString** | The value of the user attribute | 

## Methods

### NewUsersUserIdUserAttributesPostRequest

`func NewUsersUserIdUserAttributesPostRequest(name string, value NullableString, ) *UsersUserIdUserAttributesPostRequest`

NewUsersUserIdUserAttributesPostRequest instantiates a new UsersUserIdUserAttributesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersUserIdUserAttributesPostRequestWithDefaults

`func NewUsersUserIdUserAttributesPostRequestWithDefaults() *UsersUserIdUserAttributesPostRequest`

NewUsersUserIdUserAttributesPostRequestWithDefaults instantiates a new UsersUserIdUserAttributesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UsersUserIdUserAttributesPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UsersUserIdUserAttributesPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UsersUserIdUserAttributesPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *UsersUserIdUserAttributesPostRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UsersUserIdUserAttributesPostRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UsersUserIdUserAttributesPostRequest) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *UsersUserIdUserAttributesPostRequest) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *UsersUserIdUserAttributesPostRequest) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


