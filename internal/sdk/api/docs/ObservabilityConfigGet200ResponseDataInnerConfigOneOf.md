# ObservabilityConfigGet200ResponseDataInnerConfigOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** | This is an enumerated field that can only take the value \&quot;Sentry\&quot;. It identifies the provider of the observability configuration. | 
**Dsn** | **string** | This is a string field that represents the Data Source Name (DSN). It is the unique identifier that helps Sentry know where to send errors and exceptions. | 
**CaCerts** | Pointer to **[]string** | This is an optional array of Certificate Authority (CA) certificates. These are passed as &#x60;caCerts&#x60; to Sentry&#39;s NodeClient | [optional] 

## Methods

### NewObservabilityConfigGet200ResponseDataInnerConfigOneOf

`func NewObservabilityConfigGet200ResponseDataInnerConfigOneOf(provider string, dsn string, ) *ObservabilityConfigGet200ResponseDataInnerConfigOneOf`

NewObservabilityConfigGet200ResponseDataInnerConfigOneOf instantiates a new ObservabilityConfigGet200ResponseDataInnerConfigOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObservabilityConfigGet200ResponseDataInnerConfigOneOfWithDefaults

`func NewObservabilityConfigGet200ResponseDataInnerConfigOneOfWithDefaults() *ObservabilityConfigGet200ResponseDataInnerConfigOneOf`

NewObservabilityConfigGet200ResponseDataInnerConfigOneOfWithDefaults instantiates a new ObservabilityConfigGet200ResponseDataInnerConfigOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *ObservabilityConfigGet200ResponseDataInnerConfigOneOf) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ObservabilityConfigGet200ResponseDataInnerConfigOneOf) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ObservabilityConfigGet200ResponseDataInnerConfigOneOf) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetDsn

`func (o *ObservabilityConfigGet200ResponseDataInnerConfigOneOf) GetDsn() string`

GetDsn returns the Dsn field if non-nil, zero value otherwise.

### GetDsnOk

`func (o *ObservabilityConfigGet200ResponseDataInnerConfigOneOf) GetDsnOk() (*string, bool)`

GetDsnOk returns a tuple with the Dsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsn

`func (o *ObservabilityConfigGet200ResponseDataInnerConfigOneOf) SetDsn(v string)`

SetDsn sets Dsn field to given value.


### GetCaCerts

`func (o *ObservabilityConfigGet200ResponseDataInnerConfigOneOf) GetCaCerts() []string`

GetCaCerts returns the CaCerts field if non-nil, zero value otherwise.

### GetCaCertsOk

`func (o *ObservabilityConfigGet200ResponseDataInnerConfigOneOf) GetCaCertsOk() (*[]string, bool)`

GetCaCertsOk returns a tuple with the CaCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCerts

`func (o *ObservabilityConfigGet200ResponseDataInnerConfigOneOf) SetCaCerts(v []string)`

SetCaCerts sets CaCerts field to given value.

### HasCaCerts

`func (o *ObservabilityConfigGet200ResponseDataInnerConfigOneOf) HasCaCerts() bool`

HasCaCerts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


