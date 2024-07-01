# PermissionsGrantPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | [**PermissionsListObjectsPostRequestSubject**](PermissionsListObjectsPostRequestSubject.md) |  | 
**Object** | [**PermissionsGrantPostRequestObject**](PermissionsGrantPostRequestObject.md) |  | 
**AccessLevel** | **string** | The access level that the group should have for the object | 

## Methods

### NewPermissionsGrantPostRequest

`func NewPermissionsGrantPostRequest(subject PermissionsListObjectsPostRequestSubject, object PermissionsGrantPostRequestObject, accessLevel string, ) *PermissionsGrantPostRequest`

NewPermissionsGrantPostRequest instantiates a new PermissionsGrantPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionsGrantPostRequestWithDefaults

`func NewPermissionsGrantPostRequestWithDefaults() *PermissionsGrantPostRequest`

NewPermissionsGrantPostRequestWithDefaults instantiates a new PermissionsGrantPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *PermissionsGrantPostRequest) GetSubject() PermissionsListObjectsPostRequestSubject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *PermissionsGrantPostRequest) GetSubjectOk() (*PermissionsListObjectsPostRequestSubject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *PermissionsGrantPostRequest) SetSubject(v PermissionsListObjectsPostRequestSubject)`

SetSubject sets Subject field to given value.


### GetObject

`func (o *PermissionsGrantPostRequest) GetObject() PermissionsGrantPostRequestObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PermissionsGrantPostRequest) GetObjectOk() (*PermissionsGrantPostRequestObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PermissionsGrantPostRequest) SetObject(v PermissionsGrantPostRequestObject)`

SetObject sets Object field to given value.


### GetAccessLevel

`func (o *PermissionsGrantPostRequest) GetAccessLevel() string`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *PermissionsGrantPostRequest) GetAccessLevelOk() (*string, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *PermissionsGrantPostRequest) SetAccessLevel(v string)`

SetAccessLevel sets AccessLevel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


