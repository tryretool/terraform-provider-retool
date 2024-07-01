# PermissionsRevokePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | [**PermissionsListObjectsPostRequestSubject**](PermissionsListObjectsPostRequestSubject.md) |  | 
**Object** | [**PermissionsGrantPostRequestObject**](PermissionsGrantPostRequestObject.md) |  | 

## Methods

### NewPermissionsRevokePostRequest

`func NewPermissionsRevokePostRequest(subject PermissionsListObjectsPostRequestSubject, object PermissionsGrantPostRequestObject, ) *PermissionsRevokePostRequest`

NewPermissionsRevokePostRequest instantiates a new PermissionsRevokePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionsRevokePostRequestWithDefaults

`func NewPermissionsRevokePostRequestWithDefaults() *PermissionsRevokePostRequest`

NewPermissionsRevokePostRequestWithDefaults instantiates a new PermissionsRevokePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *PermissionsRevokePostRequest) GetSubject() PermissionsListObjectsPostRequestSubject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *PermissionsRevokePostRequest) GetSubjectOk() (*PermissionsListObjectsPostRequestSubject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *PermissionsRevokePostRequest) SetSubject(v PermissionsListObjectsPostRequestSubject)`

SetSubject sets Subject field to given value.


### GetObject

`func (o *PermissionsRevokePostRequest) GetObject() PermissionsGrantPostRequestObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PermissionsRevokePostRequest) GetObjectOk() (*PermissionsGrantPostRequestObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PermissionsRevokePostRequest) SetObject(v PermissionsGrantPostRequestObject)`

SetObject sets Object field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


