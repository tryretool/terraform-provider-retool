# SpacesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Domain** | **string** | The domain of the space. On Retool Cloud, specify subdomain of the space instead. | 
**Options** | Pointer to [**SpacesPostRequestOptions**](SpacesPostRequestOptions.md) |  | [optional] 

## Methods

### NewSpacesPostRequest

`func NewSpacesPostRequest(name string, domain string, ) *SpacesPostRequest`

NewSpacesPostRequest instantiates a new SpacesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpacesPostRequestWithDefaults

`func NewSpacesPostRequestWithDefaults() *SpacesPostRequest`

NewSpacesPostRequestWithDefaults instantiates a new SpacesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SpacesPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpacesPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpacesPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDomain

`func (o *SpacesPostRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *SpacesPostRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *SpacesPostRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetOptions

`func (o *SpacesPostRequest) GetOptions() SpacesPostRequestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SpacesPostRequest) GetOptionsOk() (*SpacesPostRequestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SpacesPostRequest) SetOptions(v SpacesPostRequestOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SpacesPostRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


