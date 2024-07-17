# ObservabilityConfigGet200ResponseDataInnerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** | This is an enumerated field that can only take the value \&quot;Datadog\&quot;. It identifies the provider of the observability configuration. | 
**Dsn** | **string** | This is a string field that represents the Data Source Name (DSN). It is the unique identifier that helps Sentry know where to send errors and exceptions. | 
**ApiKey** | **string** | This is a string field that represents the API key for Datadog. This key is used to authenticate with the Datadog API. | 

## Methods

### NewObservabilityConfigGet200ResponseDataInnerConfig

`func NewObservabilityConfigGet200ResponseDataInnerConfig(provider string, dsn string, apiKey string, ) *ObservabilityConfigGet200ResponseDataInnerConfig`

NewObservabilityConfigGet200ResponseDataInnerConfig instantiates a new ObservabilityConfigGet200ResponseDataInnerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObservabilityConfigGet200ResponseDataInnerConfigWithDefaults

`func NewObservabilityConfigGet200ResponseDataInnerConfigWithDefaults() *ObservabilityConfigGet200ResponseDataInnerConfig`

NewObservabilityConfigGet200ResponseDataInnerConfigWithDefaults instantiates a new ObservabilityConfigGet200ResponseDataInnerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *ObservabilityConfigGet200ResponseDataInnerConfig) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ObservabilityConfigGet200ResponseDataInnerConfig) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ObservabilityConfigGet200ResponseDataInnerConfig) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetDsn

`func (o *ObservabilityConfigGet200ResponseDataInnerConfig) GetDsn() string`

GetDsn returns the Dsn field if non-nil, zero value otherwise.

### GetDsnOk

`func (o *ObservabilityConfigGet200ResponseDataInnerConfig) GetDsnOk() (*string, bool)`

GetDsnOk returns a tuple with the Dsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsn

`func (o *ObservabilityConfigGet200ResponseDataInnerConfig) SetDsn(v string)`

SetDsn sets Dsn field to given value.


### GetApiKey

`func (o *ObservabilityConfigGet200ResponseDataInnerConfig) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ObservabilityConfigGet200ResponseDataInnerConfig) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ObservabilityConfigGet200ResponseDataInnerConfig) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


