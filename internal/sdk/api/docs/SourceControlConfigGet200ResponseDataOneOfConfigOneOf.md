# SourceControlConfigGet200ResponseDataOneOfConfigOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**AppId** | **string** | The GitHub App ID. | 
**InstallationId** | **string** | The GitHub installation ID. This can be found at the end of the installation URL. | 
**PrivateKey** | **string** | The base64-encoded private key. | 
**Url** | Pointer to **string** | The domain used to access your self-hosted GitHub instance. | [optional] 
**EnterpriseApiUrl** | Pointer to **string** | The REST API route for your self-hosted GitHub instance. Defaults to https://[hostname]/api/v3. | [optional] 

## Methods

### NewSourceControlConfigGet200ResponseDataOneOfConfigOneOf

`func NewSourceControlConfigGet200ResponseDataOneOfConfigOneOf(type_ string, appId string, installationId string, privateKey string, ) *SourceControlConfigGet200ResponseDataOneOfConfigOneOf`

NewSourceControlConfigGet200ResponseDataOneOfConfigOneOf instantiates a new SourceControlConfigGet200ResponseDataOneOfConfigOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlConfigGet200ResponseDataOneOfConfigOneOfWithDefaults

`func NewSourceControlConfigGet200ResponseDataOneOfConfigOneOfWithDefaults() *SourceControlConfigGet200ResponseDataOneOfConfigOneOf`

NewSourceControlConfigGet200ResponseDataOneOfConfigOneOfWithDefaults instantiates a new SourceControlConfigGet200ResponseDataOneOfConfigOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SourceControlConfigGet200ResponseDataOneOfConfigOneOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SourceControlConfigGet200ResponseDataOneOfConfigOneOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SourceControlConfigGet200ResponseDataOneOfConfigOneOf) SetType(v string)`

SetType sets Type field to given value.


### GetAppId

`func (o *SourceControlConfigGet200ResponseDataOneOfConfigOneOf) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *SourceControlConfigGet200ResponseDataOneOfConfigOneOf) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *SourceControlConfigGet200ResponseDataOneOfConfigOneOf) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetInstallationId

`func (o *SourceControlConfigGet200ResponseDataOneOfConfigOneOf) GetInstallationId() string`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *SourceControlConfigGet200ResponseDataOneOfConfigOneOf) GetInstallationIdOk() (*string, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *SourceControlConfigGet200ResponseDataOneOfConfigOneOf) SetInstallationId(v string)`

SetInstallationId sets InstallationId field to given value.


### GetPrivateKey

`func (o *SourceControlConfigGet200ResponseDataOneOfConfigOneOf) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *SourceControlConfigGet200ResponseDataOneOfConfigOneOf) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *SourceControlConfigGet200ResponseDataOneOfConfigOneOf) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetUrl

`func (o *SourceControlConfigGet200ResponseDataOneOfConfigOneOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SourceControlConfigGet200ResponseDataOneOfConfigOneOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SourceControlConfigGet200ResponseDataOneOfConfigOneOf) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SourceControlConfigGet200ResponseDataOneOfConfigOneOf) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetEnterpriseApiUrl

`func (o *SourceControlConfigGet200ResponseDataOneOfConfigOneOf) GetEnterpriseApiUrl() string`

GetEnterpriseApiUrl returns the EnterpriseApiUrl field if non-nil, zero value otherwise.

### GetEnterpriseApiUrlOk

`func (o *SourceControlConfigGet200ResponseDataOneOfConfigOneOf) GetEnterpriseApiUrlOk() (*string, bool)`

GetEnterpriseApiUrlOk returns a tuple with the EnterpriseApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseApiUrl

`func (o *SourceControlConfigGet200ResponseDataOneOfConfigOneOf) SetEnterpriseApiUrl(v string)`

SetEnterpriseApiUrl sets EnterpriseApiUrl field to given value.

### HasEnterpriseApiUrl

`func (o *SourceControlConfigGet200ResponseDataOneOfConfigOneOf) HasEnterpriseApiUrl() bool`

HasEnterpriseApiUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


