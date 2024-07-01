# RestAPIOptionsAuthenticationOptionsAnyOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationType** | **string** |  | 
**AmazonAwsRegion** | **string** |  | 
**AmazonServiceName** | **string** |  | 
**AmazonAccessKeyId** | **string** |  | 
**AmazonSecretAccessKey** | **string** |  | 
**AuthWithDefaultCredentialProviderChain** | Pointer to **bool** |  | [optional] 
**EnableAwsv4AuthenticationViaHeaders** | Pointer to **bool** |  | [optional] 

## Methods

### NewRestAPIOptionsAuthenticationOptionsAnyOf1

`func NewRestAPIOptionsAuthenticationOptionsAnyOf1(authenticationType string, amazonAwsRegion string, amazonServiceName string, amazonAccessKeyId string, amazonSecretAccessKey string, ) *RestAPIOptionsAuthenticationOptionsAnyOf1`

NewRestAPIOptionsAuthenticationOptionsAnyOf1 instantiates a new RestAPIOptionsAuthenticationOptionsAnyOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestAPIOptionsAuthenticationOptionsAnyOf1WithDefaults

`func NewRestAPIOptionsAuthenticationOptionsAnyOf1WithDefaults() *RestAPIOptionsAuthenticationOptionsAnyOf1`

NewRestAPIOptionsAuthenticationOptionsAnyOf1WithDefaults instantiates a new RestAPIOptionsAuthenticationOptionsAnyOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationType

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf1) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf1) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf1) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.


### GetAmazonAwsRegion

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf1) GetAmazonAwsRegion() string`

GetAmazonAwsRegion returns the AmazonAwsRegion field if non-nil, zero value otherwise.

### GetAmazonAwsRegionOk

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf1) GetAmazonAwsRegionOk() (*string, bool)`

GetAmazonAwsRegionOk returns a tuple with the AmazonAwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonAwsRegion

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf1) SetAmazonAwsRegion(v string)`

SetAmazonAwsRegion sets AmazonAwsRegion field to given value.


### GetAmazonServiceName

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf1) GetAmazonServiceName() string`

GetAmazonServiceName returns the AmazonServiceName field if non-nil, zero value otherwise.

### GetAmazonServiceNameOk

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf1) GetAmazonServiceNameOk() (*string, bool)`

GetAmazonServiceNameOk returns a tuple with the AmazonServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonServiceName

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf1) SetAmazonServiceName(v string)`

SetAmazonServiceName sets AmazonServiceName field to given value.


### GetAmazonAccessKeyId

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf1) GetAmazonAccessKeyId() string`

GetAmazonAccessKeyId returns the AmazonAccessKeyId field if non-nil, zero value otherwise.

### GetAmazonAccessKeyIdOk

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf1) GetAmazonAccessKeyIdOk() (*string, bool)`

GetAmazonAccessKeyIdOk returns a tuple with the AmazonAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonAccessKeyId

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf1) SetAmazonAccessKeyId(v string)`

SetAmazonAccessKeyId sets AmazonAccessKeyId field to given value.


### GetAmazonSecretAccessKey

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf1) GetAmazonSecretAccessKey() string`

GetAmazonSecretAccessKey returns the AmazonSecretAccessKey field if non-nil, zero value otherwise.

### GetAmazonSecretAccessKeyOk

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf1) GetAmazonSecretAccessKeyOk() (*string, bool)`

GetAmazonSecretAccessKeyOk returns a tuple with the AmazonSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonSecretAccessKey

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf1) SetAmazonSecretAccessKey(v string)`

SetAmazonSecretAccessKey sets AmazonSecretAccessKey field to given value.


### GetAuthWithDefaultCredentialProviderChain

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf1) GetAuthWithDefaultCredentialProviderChain() bool`

GetAuthWithDefaultCredentialProviderChain returns the AuthWithDefaultCredentialProviderChain field if non-nil, zero value otherwise.

### GetAuthWithDefaultCredentialProviderChainOk

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf1) GetAuthWithDefaultCredentialProviderChainOk() (*bool, bool)`

GetAuthWithDefaultCredentialProviderChainOk returns a tuple with the AuthWithDefaultCredentialProviderChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthWithDefaultCredentialProviderChain

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf1) SetAuthWithDefaultCredentialProviderChain(v bool)`

SetAuthWithDefaultCredentialProviderChain sets AuthWithDefaultCredentialProviderChain field to given value.

### HasAuthWithDefaultCredentialProviderChain

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf1) HasAuthWithDefaultCredentialProviderChain() bool`

HasAuthWithDefaultCredentialProviderChain returns a boolean if a field has been set.

### GetEnableAwsv4AuthenticationViaHeaders

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf1) GetEnableAwsv4AuthenticationViaHeaders() bool`

GetEnableAwsv4AuthenticationViaHeaders returns the EnableAwsv4AuthenticationViaHeaders field if non-nil, zero value otherwise.

### GetEnableAwsv4AuthenticationViaHeadersOk

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf1) GetEnableAwsv4AuthenticationViaHeadersOk() (*bool, bool)`

GetEnableAwsv4AuthenticationViaHeadersOk returns a tuple with the EnableAwsv4AuthenticationViaHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAwsv4AuthenticationViaHeaders

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf1) SetEnableAwsv4AuthenticationViaHeaders(v bool)`

SetEnableAwsv4AuthenticationViaHeaders sets EnableAwsv4AuthenticationViaHeaders field to given value.

### HasEnableAwsv4AuthenticationViaHeaders

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf1) HasEnableAwsv4AuthenticationViaHeaders() bool`

HasEnableAwsv4AuthenticationViaHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


