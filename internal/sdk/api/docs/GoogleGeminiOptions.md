# GoogleGeminiOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeySource** | **string** | Always &#39;byok&#39; (bring your own key) — this provider requires you to supply credentials. Defaults to &#39;byok&#39; and may be omitted. | [default to "byok"]
**BaseUrl** | **string** |  | [default to "https://generativelanguage.googleapis.com"]
**ApiKey** | **string** |  | 
**WebGroundingEnabled** | Pointer to **bool** |  | [optional] 
**CustomHeaders** | Pointer to **[][]string** |  | [optional] 

## Methods

### NewGoogleGeminiOptions

`func NewGoogleGeminiOptions(keySource string, baseUrl string, apiKey string, ) *GoogleGeminiOptions`

NewGoogleGeminiOptions instantiates a new GoogleGeminiOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleGeminiOptionsWithDefaults

`func NewGoogleGeminiOptionsWithDefaults() *GoogleGeminiOptions`

NewGoogleGeminiOptionsWithDefaults instantiates a new GoogleGeminiOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeySource

`func (o *GoogleGeminiOptions) GetKeySource() string`

GetKeySource returns the KeySource field if non-nil, zero value otherwise.

### GetKeySourceOk

`func (o *GoogleGeminiOptions) GetKeySourceOk() (*string, bool)`

GetKeySourceOk returns a tuple with the KeySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySource

`func (o *GoogleGeminiOptions) SetKeySource(v string)`

SetKeySource sets KeySource field to given value.


### GetBaseUrl

`func (o *GoogleGeminiOptions) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *GoogleGeminiOptions) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *GoogleGeminiOptions) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetApiKey

`func (o *GoogleGeminiOptions) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *GoogleGeminiOptions) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *GoogleGeminiOptions) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetWebGroundingEnabled

`func (o *GoogleGeminiOptions) GetWebGroundingEnabled() bool`

GetWebGroundingEnabled returns the WebGroundingEnabled field if non-nil, zero value otherwise.

### GetWebGroundingEnabledOk

`func (o *GoogleGeminiOptions) GetWebGroundingEnabledOk() (*bool, bool)`

GetWebGroundingEnabledOk returns a tuple with the WebGroundingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebGroundingEnabled

`func (o *GoogleGeminiOptions) SetWebGroundingEnabled(v bool)`

SetWebGroundingEnabled sets WebGroundingEnabled field to given value.

### HasWebGroundingEnabled

`func (o *GoogleGeminiOptions) HasWebGroundingEnabled() bool`

HasWebGroundingEnabled returns a boolean if a field has been set.

### GetCustomHeaders

`func (o *GoogleGeminiOptions) GetCustomHeaders() [][]string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *GoogleGeminiOptions) GetCustomHeadersOk() (*[][]string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *GoogleGeminiOptions) SetCustomHeaders(v [][]string)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *GoogleGeminiOptions) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


