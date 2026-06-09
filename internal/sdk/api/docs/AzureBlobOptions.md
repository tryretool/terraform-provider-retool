# AzureBlobOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationOptions** | [**AzureBlobOptionsAuthenticationOptions**](AzureBlobOptionsAuthenticationOptions.md) |  | 

## Methods

### NewAzureBlobOptions

`func NewAzureBlobOptions(authenticationOptions AzureBlobOptionsAuthenticationOptions, ) *AzureBlobOptions`

NewAzureBlobOptions instantiates a new AzureBlobOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureBlobOptionsWithDefaults

`func NewAzureBlobOptionsWithDefaults() *AzureBlobOptions`

NewAzureBlobOptionsWithDefaults instantiates a new AzureBlobOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationOptions

`func (o *AzureBlobOptions) GetAuthenticationOptions() AzureBlobOptionsAuthenticationOptions`

GetAuthenticationOptions returns the AuthenticationOptions field if non-nil, zero value otherwise.

### GetAuthenticationOptionsOk

`func (o *AzureBlobOptions) GetAuthenticationOptionsOk() (*AzureBlobOptionsAuthenticationOptions, bool)`

GetAuthenticationOptionsOk returns a tuple with the AuthenticationOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOptions

`func (o *AzureBlobOptions) SetAuthenticationOptions(v AzureBlobOptionsAuthenticationOptions)`

SetAuthenticationOptions sets AuthenticationOptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


