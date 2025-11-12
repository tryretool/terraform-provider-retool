# RolePermissionsRoleGrantsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleId** | **string** | The id of the role | 
**SubjectType** | **string** | The type of the subject | 
**SubjectId** | **string** | The id of the subject | 
**ObjectType** | Pointer to **string** | The type of the object | [optional] 
**ObjectId** | Pointer to **string** | The id of the object | [optional] 

## Methods

### NewRolePermissionsRoleGrantsPostRequest

`func NewRolePermissionsRoleGrantsPostRequest(roleId string, subjectType string, subjectId string, ) *RolePermissionsRoleGrantsPostRequest`

NewRolePermissionsRoleGrantsPostRequest instantiates a new RolePermissionsRoleGrantsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolePermissionsRoleGrantsPostRequestWithDefaults

`func NewRolePermissionsRoleGrantsPostRequestWithDefaults() *RolePermissionsRoleGrantsPostRequest`

NewRolePermissionsRoleGrantsPostRequestWithDefaults instantiates a new RolePermissionsRoleGrantsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleId

`func (o *RolePermissionsRoleGrantsPostRequest) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *RolePermissionsRoleGrantsPostRequest) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *RolePermissionsRoleGrantsPostRequest) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.


### GetSubjectType

`func (o *RolePermissionsRoleGrantsPostRequest) GetSubjectType() string`

GetSubjectType returns the SubjectType field if non-nil, zero value otherwise.

### GetSubjectTypeOk

`func (o *RolePermissionsRoleGrantsPostRequest) GetSubjectTypeOk() (*string, bool)`

GetSubjectTypeOk returns a tuple with the SubjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectType

`func (o *RolePermissionsRoleGrantsPostRequest) SetSubjectType(v string)`

SetSubjectType sets SubjectType field to given value.


### GetSubjectId

`func (o *RolePermissionsRoleGrantsPostRequest) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *RolePermissionsRoleGrantsPostRequest) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *RolePermissionsRoleGrantsPostRequest) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.


### GetObjectType

`func (o *RolePermissionsRoleGrantsPostRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RolePermissionsRoleGrantsPostRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RolePermissionsRoleGrantsPostRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *RolePermissionsRoleGrantsPostRequest) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectId

`func (o *RolePermissionsRoleGrantsPostRequest) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *RolePermissionsRoleGrantsPostRequest) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *RolePermissionsRoleGrantsPostRequest) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *RolePermissionsRoleGrantsPostRequest) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


