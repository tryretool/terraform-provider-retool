# AccessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**Status** | **string** |  | 
**LegacyId** | **float32** |  | 
**RequestingEmail** | **string** |  | 
**UpdatedById** | **NullableString** |  | 

## Methods

### NewAccessRequest

`func NewAccessRequest(id float32, status string, legacyId float32, requestingEmail string, updatedById NullableString, ) *AccessRequest`

NewAccessRequest instantiates a new AccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessRequestWithDefaults

`func NewAccessRequestWithDefaults() *AccessRequest`

NewAccessRequestWithDefaults instantiates a new AccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccessRequest) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessRequest) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessRequest) SetId(v float32)`

SetId sets Id field to given value.


### GetStatus

`func (o *AccessRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccessRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccessRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLegacyId

`func (o *AccessRequest) GetLegacyId() float32`

GetLegacyId returns the LegacyId field if non-nil, zero value otherwise.

### GetLegacyIdOk

`func (o *AccessRequest) GetLegacyIdOk() (*float32, bool)`

GetLegacyIdOk returns a tuple with the LegacyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyId

`func (o *AccessRequest) SetLegacyId(v float32)`

SetLegacyId sets LegacyId field to given value.


### GetRequestingEmail

`func (o *AccessRequest) GetRequestingEmail() string`

GetRequestingEmail returns the RequestingEmail field if non-nil, zero value otherwise.

### GetRequestingEmailOk

`func (o *AccessRequest) GetRequestingEmailOk() (*string, bool)`

GetRequestingEmailOk returns a tuple with the RequestingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestingEmail

`func (o *AccessRequest) SetRequestingEmail(v string)`

SetRequestingEmail sets RequestingEmail field to given value.


### GetUpdatedById

`func (o *AccessRequest) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *AccessRequest) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *AccessRequest) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


### SetUpdatedByIdNil

`func (o *AccessRequest) SetUpdatedByIdNil(b bool)`

 SetUpdatedByIdNil sets the value for UpdatedById to be an explicit nil

### UnsetUpdatedById
`func (o *AccessRequest) UnsetUpdatedById()`

UnsetUpdatedById ensures that no value is present for UpdatedById, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


