# PermissionsListObjectsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | [**PermissionsListObjectsPostRequestSubject**](PermissionsListObjectsPostRequestSubject.md) |  | 
**ObjectType** | **string** |  | 
**IncludeInheritedAccess** | Pointer to **bool** |  | [optional] 

## Methods

### NewPermissionsListObjectsPostRequest

`func NewPermissionsListObjectsPostRequest(subject PermissionsListObjectsPostRequestSubject, objectType string, ) *PermissionsListObjectsPostRequest`

NewPermissionsListObjectsPostRequest instantiates a new PermissionsListObjectsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionsListObjectsPostRequestWithDefaults

`func NewPermissionsListObjectsPostRequestWithDefaults() *PermissionsListObjectsPostRequest`

NewPermissionsListObjectsPostRequestWithDefaults instantiates a new PermissionsListObjectsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *PermissionsListObjectsPostRequest) GetSubject() PermissionsListObjectsPostRequestSubject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *PermissionsListObjectsPostRequest) GetSubjectOk() (*PermissionsListObjectsPostRequestSubject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *PermissionsListObjectsPostRequest) SetSubject(v PermissionsListObjectsPostRequestSubject)`

SetSubject sets Subject field to given value.


### GetObjectType

`func (o *PermissionsListObjectsPostRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PermissionsListObjectsPostRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PermissionsListObjectsPostRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIncludeInheritedAccess

`func (o *PermissionsListObjectsPostRequest) GetIncludeInheritedAccess() bool`

GetIncludeInheritedAccess returns the IncludeInheritedAccess field if non-nil, zero value otherwise.

### GetIncludeInheritedAccessOk

`func (o *PermissionsListObjectsPostRequest) GetIncludeInheritedAccessOk() (*bool, bool)`

GetIncludeInheritedAccessOk returns a tuple with the IncludeInheritedAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInheritedAccess

`func (o *PermissionsListObjectsPostRequest) SetIncludeInheritedAccess(v bool)`

SetIncludeInheritedAccess sets IncludeInheritedAccess field to given value.

### HasIncludeInheritedAccess

`func (o *PermissionsListObjectsPostRequest) HasIncludeInheritedAccess() bool`

HasIncludeInheritedAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


