# ResourceConfigurationsGet200ResponseDataInnerOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationOptions** | [**GRPCOptionsAuthenticationOptions**](GRPCOptionsAuthenticationOptions.md) |  | 
**DatabaseOptions** | [**SnowflakeOptionsDatabaseOptions**](SnowflakeOptionsDatabaseOptions.md) |  | 
**AccountIdentifier** | **string** |  | 
**UserRole** | Pointer to **string** |  | [optional] 
**BaseUrl** | **string** |  | 
**UrlParameters** | Pointer to **[][]string** |  | [optional] 
**Headers** | Pointer to **[][]string** |  | [optional] 
**ExtraBodyValues** | Pointer to **[][]string** | Extra body values are not passed for GET or HEAD requests. | [optional] 
**CookiesToForward** | Pointer to **[]string** | You can use the pattern COOKIE_your_cookie_name in the Headers section in order to implement the double-cookie submit pattern. | [optional] 
**ForwardAllCookies** | Pointer to **bool** | This is useful if you have dynamic cookie names. | [optional] 
**MaxReceiveMessageLength** | Pointer to **string** |  | [optional] 
**MaxSendMessageLength** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **[][]string** |  | [optional] 
**Options** | Pointer to **interface{}** |  | [optional] 
**ProtoFileUrl** | Pointer to **string** |  | [optional] 
**ProtoFileUrlHeaders** | Pointer to **[][]string** |  | [optional] 
**ProtoSource** | Pointer to **string** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 
**ServiceNames** | Pointer to **[]string** |  | [optional] 
**VerifySessionActionEnabled** | Pointer to **bool** |  | [optional] 
**KeySource** | **string** | Always &#39;byok&#39; (bring your own key) — this provider requires you to supply credentials. Defaults to &#39;byok&#39; and may be omitted. | [default to "byok"]
**ApiKey** | **string** |  | 
**CustomHeaders** | Pointer to **[][]string** |  | [optional] 
**Region** | **string** |  | 
**AccessKeyId** | Pointer to **string** |  | [optional] 
**SecretAccessKey** | Pointer to **string** |  | [optional] 
**SessionToken** | Pointer to **string** |  | [optional] 
**AssumeRole** | Pointer to **string** |  | [optional] 
**AuthWithDefaultCredentialProviderChain** | Pointer to **bool** | Retool will source AWS credentials from the credential provider chain. Use this option to authenticate with credentials provided in environment variables or the underlying instance role. | [optional] 
**DeploymentNames** | **[]string** |  | 
**ApiVersion** | Pointer to **string** |  | [optional] 
**WebGroundingEnabled** | Pointer to **bool** |  | [optional] 
**ProjectId** | **string** |  | 
**Location** | **string** |  | 
**ServiceAccountKey** | **string** |  | 
**CompatibleSchema** | **string** |  | 
**ModelList** | **[]string** |  | 

## Methods

### NewResourceConfigurationsGet200ResponseDataInnerOptions

`func NewResourceConfigurationsGet200ResponseDataInnerOptions(authenticationOptions GRPCOptionsAuthenticationOptions, databaseOptions SnowflakeOptionsDatabaseOptions, accountIdentifier string, baseUrl string, keySource string, apiKey string, region string, deploymentNames []string, projectId string, location string, serviceAccountKey string, compatibleSchema string, modelList []string, ) *ResourceConfigurationsGet200ResponseDataInnerOptions`

NewResourceConfigurationsGet200ResponseDataInnerOptions instantiates a new ResourceConfigurationsGet200ResponseDataInnerOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceConfigurationsGet200ResponseDataInnerOptionsWithDefaults

`func NewResourceConfigurationsGet200ResponseDataInnerOptionsWithDefaults() *ResourceConfigurationsGet200ResponseDataInnerOptions`

NewResourceConfigurationsGet200ResponseDataInnerOptionsWithDefaults instantiates a new ResourceConfigurationsGet200ResponseDataInnerOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationOptions

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetAuthenticationOptions() GRPCOptionsAuthenticationOptions`

GetAuthenticationOptions returns the AuthenticationOptions field if non-nil, zero value otherwise.

### GetAuthenticationOptionsOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetAuthenticationOptionsOk() (*GRPCOptionsAuthenticationOptions, bool)`

GetAuthenticationOptionsOk returns a tuple with the AuthenticationOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOptions

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetAuthenticationOptions(v GRPCOptionsAuthenticationOptions)`

SetAuthenticationOptions sets AuthenticationOptions field to given value.


### GetDatabaseOptions

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetDatabaseOptions() SnowflakeOptionsDatabaseOptions`

GetDatabaseOptions returns the DatabaseOptions field if non-nil, zero value otherwise.

### GetDatabaseOptionsOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetDatabaseOptionsOk() (*SnowflakeOptionsDatabaseOptions, bool)`

GetDatabaseOptionsOk returns a tuple with the DatabaseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseOptions

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetDatabaseOptions(v SnowflakeOptionsDatabaseOptions)`

SetDatabaseOptions sets DatabaseOptions field to given value.


### GetAccountIdentifier

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetAccountIdentifier() string`

GetAccountIdentifier returns the AccountIdentifier field if non-nil, zero value otherwise.

### GetAccountIdentifierOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetAccountIdentifierOk() (*string, bool)`

GetAccountIdentifierOk returns a tuple with the AccountIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIdentifier

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetAccountIdentifier(v string)`

SetAccountIdentifier sets AccountIdentifier field to given value.


### GetUserRole

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetUserRole() string`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetUserRoleOk() (*string, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetUserRole(v string)`

SetUserRole sets UserRole field to given value.

### HasUserRole

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) HasUserRole() bool`

HasUserRole returns a boolean if a field has been set.

### GetBaseUrl

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetUrlParameters

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetUrlParameters() [][]string`

GetUrlParameters returns the UrlParameters field if non-nil, zero value otherwise.

### GetUrlParametersOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetUrlParametersOk() (*[][]string, bool)`

GetUrlParametersOk returns a tuple with the UrlParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlParameters

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetUrlParameters(v [][]string)`

SetUrlParameters sets UrlParameters field to given value.

### HasUrlParameters

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) HasUrlParameters() bool`

HasUrlParameters returns a boolean if a field has been set.

### GetHeaders

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetHeaders() [][]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetHeadersOk() (*[][]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetHeaders(v [][]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetExtraBodyValues

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetExtraBodyValues() [][]string`

GetExtraBodyValues returns the ExtraBodyValues field if non-nil, zero value otherwise.

### GetExtraBodyValuesOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetExtraBodyValuesOk() (*[][]string, bool)`

GetExtraBodyValuesOk returns a tuple with the ExtraBodyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraBodyValues

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetExtraBodyValues(v [][]string)`

SetExtraBodyValues sets ExtraBodyValues field to given value.

### HasExtraBodyValues

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) HasExtraBodyValues() bool`

HasExtraBodyValues returns a boolean if a field has been set.

### GetCookiesToForward

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetCookiesToForward() []string`

GetCookiesToForward returns the CookiesToForward field if non-nil, zero value otherwise.

### GetCookiesToForwardOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetCookiesToForwardOk() (*[]string, bool)`

GetCookiesToForwardOk returns a tuple with the CookiesToForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookiesToForward

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetCookiesToForward(v []string)`

SetCookiesToForward sets CookiesToForward field to given value.

### HasCookiesToForward

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) HasCookiesToForward() bool`

HasCookiesToForward returns a boolean if a field has been set.

### GetForwardAllCookies

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetForwardAllCookies() bool`

GetForwardAllCookies returns the ForwardAllCookies field if non-nil, zero value otherwise.

### GetForwardAllCookiesOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetForwardAllCookiesOk() (*bool, bool)`

GetForwardAllCookiesOk returns a tuple with the ForwardAllCookies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardAllCookies

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetForwardAllCookies(v bool)`

SetForwardAllCookies sets ForwardAllCookies field to given value.

### HasForwardAllCookies

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) HasForwardAllCookies() bool`

HasForwardAllCookies returns a boolean if a field has been set.

### GetMaxReceiveMessageLength

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetMaxReceiveMessageLength() string`

GetMaxReceiveMessageLength returns the MaxReceiveMessageLength field if non-nil, zero value otherwise.

### GetMaxReceiveMessageLengthOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetMaxReceiveMessageLengthOk() (*string, bool)`

GetMaxReceiveMessageLengthOk returns a tuple with the MaxReceiveMessageLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReceiveMessageLength

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetMaxReceiveMessageLength(v string)`

SetMaxReceiveMessageLength sets MaxReceiveMessageLength field to given value.

### HasMaxReceiveMessageLength

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) HasMaxReceiveMessageLength() bool`

HasMaxReceiveMessageLength returns a boolean if a field has been set.

### GetMaxSendMessageLength

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetMaxSendMessageLength() string`

GetMaxSendMessageLength returns the MaxSendMessageLength field if non-nil, zero value otherwise.

### GetMaxSendMessageLengthOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetMaxSendMessageLengthOk() (*string, bool)`

GetMaxSendMessageLengthOk returns a tuple with the MaxSendMessageLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSendMessageLength

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetMaxSendMessageLength(v string)`

SetMaxSendMessageLength sets MaxSendMessageLength field to given value.

### HasMaxSendMessageLength

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) HasMaxSendMessageLength() bool`

HasMaxSendMessageLength returns a boolean if a field has been set.

### GetMetadata

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetMetadata() [][]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetMetadataOk() (*[][]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetMetadata(v [][]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOptions

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetOptions() interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetOptionsOk() (*interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetOptions(v interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetProtoFileUrl

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetProtoFileUrl() string`

GetProtoFileUrl returns the ProtoFileUrl field if non-nil, zero value otherwise.

### GetProtoFileUrlOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetProtoFileUrlOk() (*string, bool)`

GetProtoFileUrlOk returns a tuple with the ProtoFileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtoFileUrl

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetProtoFileUrl(v string)`

SetProtoFileUrl sets ProtoFileUrl field to given value.

### HasProtoFileUrl

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) HasProtoFileUrl() bool`

HasProtoFileUrl returns a boolean if a field has been set.

### GetProtoFileUrlHeaders

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetProtoFileUrlHeaders() [][]string`

GetProtoFileUrlHeaders returns the ProtoFileUrlHeaders field if non-nil, zero value otherwise.

### GetProtoFileUrlHeadersOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetProtoFileUrlHeadersOk() (*[][]string, bool)`

GetProtoFileUrlHeadersOk returns a tuple with the ProtoFileUrlHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtoFileUrlHeaders

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetProtoFileUrlHeaders(v [][]string)`

SetProtoFileUrlHeaders sets ProtoFileUrlHeaders field to given value.

### HasProtoFileUrlHeaders

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) HasProtoFileUrlHeaders() bool`

HasProtoFileUrlHeaders returns a boolean if a field has been set.

### SetProtoFileUrlHeadersNil

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetProtoFileUrlHeadersNil(b bool)`

 SetProtoFileUrlHeadersNil sets the value for ProtoFileUrlHeaders to be an explicit nil

### UnsetProtoFileUrlHeaders
`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) UnsetProtoFileUrlHeaders()`

UnsetProtoFileUrlHeaders ensures that no value is present for ProtoFileUrlHeaders, not even an explicit nil
### GetProtoSource

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetProtoSource() string`

GetProtoSource returns the ProtoSource field if non-nil, zero value otherwise.

### GetProtoSourceOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetProtoSourceOk() (*string, bool)`

GetProtoSourceOk returns a tuple with the ProtoSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtoSource

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetProtoSource(v string)`

SetProtoSource sets ProtoSource field to given value.

### HasProtoSource

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) HasProtoSource() bool`

HasProtoSource returns a boolean if a field has been set.

### GetServiceName

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetServiceNames

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetServiceNames() []string`

GetServiceNames returns the ServiceNames field if non-nil, zero value otherwise.

### GetServiceNamesOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetServiceNamesOk() (*[]string, bool)`

GetServiceNamesOk returns a tuple with the ServiceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNames

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetServiceNames(v []string)`

SetServiceNames sets ServiceNames field to given value.

### HasServiceNames

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) HasServiceNames() bool`

HasServiceNames returns a boolean if a field has been set.

### GetVerifySessionActionEnabled

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetVerifySessionActionEnabled() bool`

GetVerifySessionActionEnabled returns the VerifySessionActionEnabled field if non-nil, zero value otherwise.

### GetVerifySessionActionEnabledOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetVerifySessionActionEnabledOk() (*bool, bool)`

GetVerifySessionActionEnabledOk returns a tuple with the VerifySessionActionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySessionActionEnabled

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetVerifySessionActionEnabled(v bool)`

SetVerifySessionActionEnabled sets VerifySessionActionEnabled field to given value.

### HasVerifySessionActionEnabled

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) HasVerifySessionActionEnabled() bool`

HasVerifySessionActionEnabled returns a boolean if a field has been set.

### GetKeySource

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetKeySource() string`

GetKeySource returns the KeySource field if non-nil, zero value otherwise.

### GetKeySourceOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetKeySourceOk() (*string, bool)`

GetKeySourceOk returns a tuple with the KeySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySource

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetKeySource(v string)`

SetKeySource sets KeySource field to given value.


### GetApiKey

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetCustomHeaders

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetCustomHeaders() [][]string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetCustomHeadersOk() (*[][]string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetCustomHeaders(v [][]string)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.

### GetRegion

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetAccessKeyId

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetSecretAccessKey

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.

### GetSessionToken

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetSessionToken() string`

GetSessionToken returns the SessionToken field if non-nil, zero value otherwise.

### GetSessionTokenOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetSessionTokenOk() (*string, bool)`

GetSessionTokenOk returns a tuple with the SessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionToken

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetSessionToken(v string)`

SetSessionToken sets SessionToken field to given value.

### HasSessionToken

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) HasSessionToken() bool`

HasSessionToken returns a boolean if a field has been set.

### GetAssumeRole

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetAssumeRole() string`

GetAssumeRole returns the AssumeRole field if non-nil, zero value otherwise.

### GetAssumeRoleOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetAssumeRoleOk() (*string, bool)`

GetAssumeRoleOk returns a tuple with the AssumeRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssumeRole

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetAssumeRole(v string)`

SetAssumeRole sets AssumeRole field to given value.

### HasAssumeRole

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) HasAssumeRole() bool`

HasAssumeRole returns a boolean if a field has been set.

### GetAuthWithDefaultCredentialProviderChain

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetAuthWithDefaultCredentialProviderChain() bool`

GetAuthWithDefaultCredentialProviderChain returns the AuthWithDefaultCredentialProviderChain field if non-nil, zero value otherwise.

### GetAuthWithDefaultCredentialProviderChainOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetAuthWithDefaultCredentialProviderChainOk() (*bool, bool)`

GetAuthWithDefaultCredentialProviderChainOk returns a tuple with the AuthWithDefaultCredentialProviderChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthWithDefaultCredentialProviderChain

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetAuthWithDefaultCredentialProviderChain(v bool)`

SetAuthWithDefaultCredentialProviderChain sets AuthWithDefaultCredentialProviderChain field to given value.

### HasAuthWithDefaultCredentialProviderChain

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) HasAuthWithDefaultCredentialProviderChain() bool`

HasAuthWithDefaultCredentialProviderChain returns a boolean if a field has been set.

### GetDeploymentNames

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetDeploymentNames() []string`

GetDeploymentNames returns the DeploymentNames field if non-nil, zero value otherwise.

### GetDeploymentNamesOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetDeploymentNamesOk() (*[]string, bool)`

GetDeploymentNamesOk returns a tuple with the DeploymentNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentNames

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetDeploymentNames(v []string)`

SetDeploymentNames sets DeploymentNames field to given value.


### GetApiVersion

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetWebGroundingEnabled

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetWebGroundingEnabled() bool`

GetWebGroundingEnabled returns the WebGroundingEnabled field if non-nil, zero value otherwise.

### GetWebGroundingEnabledOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetWebGroundingEnabledOk() (*bool, bool)`

GetWebGroundingEnabledOk returns a tuple with the WebGroundingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebGroundingEnabled

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetWebGroundingEnabled(v bool)`

SetWebGroundingEnabled sets WebGroundingEnabled field to given value.

### HasWebGroundingEnabled

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) HasWebGroundingEnabled() bool`

HasWebGroundingEnabled returns a boolean if a field has been set.

### GetProjectId

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetLocation

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetServiceAccountKey

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetServiceAccountKey() string`

GetServiceAccountKey returns the ServiceAccountKey field if non-nil, zero value otherwise.

### GetServiceAccountKeyOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetServiceAccountKeyOk() (*string, bool)`

GetServiceAccountKeyOk returns a tuple with the ServiceAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountKey

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetServiceAccountKey(v string)`

SetServiceAccountKey sets ServiceAccountKey field to given value.


### GetCompatibleSchema

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetCompatibleSchema() string`

GetCompatibleSchema returns the CompatibleSchema field if non-nil, zero value otherwise.

### GetCompatibleSchemaOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetCompatibleSchemaOk() (*string, bool)`

GetCompatibleSchemaOk returns a tuple with the CompatibleSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibleSchema

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetCompatibleSchema(v string)`

SetCompatibleSchema sets CompatibleSchema field to given value.


### GetModelList

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetModelList() []string`

GetModelList returns the ModelList field if non-nil, zero value otherwise.

### GetModelListOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetModelListOk() (*[]string, bool)`

GetModelListOk returns a tuple with the ModelList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelList

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetModelList(v []string)`

SetModelList sets ModelList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


