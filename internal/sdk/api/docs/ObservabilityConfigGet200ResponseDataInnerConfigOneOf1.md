# ObservabilityConfigGet200ResponseDataInnerConfigOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** | This is an enumerated field that can only take the value \&quot;Datadog\&quot;. It identifies the provider of the observability configuration. | 
**ApiKey** | **string** | This is a string field that represents the API key for Datadog. This key is used to authenticate with the Datadog API. | 

## Methods

### NewObservabilityConfigGet200ResponseDataInnerConfigOneOf1

`func NewObservabilityConfigGet200ResponseDataInnerConfigOneOf1(provider string, apiKey string, ) *ObservabilityConfigGet200ResponseDataInnerConfigOneOf1`

NewObservabilityConfigGet200ResponseDataInnerConfigOneOf1 instantiates a new ObservabilityConfigGet200ResponseDataInnerConfigOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObservabilityConfigGet200ResponseDataInnerConfigOneOf1WithDefaults

`func NewObservabilityConfigGet200ResponseDataInnerConfigOneOf1WithDefaults() *ObservabilityConfigGet200ResponseDataInnerConfigOneOf1`

NewObservabilityConfigGet200ResponseDataInnerConfigOneOf1WithDefaults instantiates a new ObservabilityConfigGet200ResponseDataInnerConfigOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *ObservabilityConfigGet200ResponseDataInnerConfigOneOf1) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ObservabilityConfigGet200ResponseDataInnerConfigOneOf1) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ObservabilityConfigGet200ResponseDataInnerConfigOneOf1) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetApiKey

`func (o *ObservabilityConfigGet200ResponseDataInnerConfigOneOf1) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ObservabilityConfigGet200ResponseDataInnerConfigOneOf1) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ObservabilityConfigGet200ResponseDataInnerConfigOneOf1) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


