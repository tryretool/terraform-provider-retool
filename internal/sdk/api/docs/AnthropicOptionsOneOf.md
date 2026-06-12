# AnthropicOptionsOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeySource** | **string** | &#39;byok&#39; (bring your own key) — authenticates using the api_key you supply at the base_url you specify. | 
**BaseUrl** | **string** |  | [default to "https://api.anthropic.com"]
**ApiKey** | **string** |  | 
**CustomHeaders** | Pointer to **[][]string** |  | [optional] 

## Methods

### NewAnthropicOptionsOneOf

`func NewAnthropicOptionsOneOf(keySource string, baseUrl string, apiKey string, ) *AnthropicOptionsOneOf`

NewAnthropicOptionsOneOf instantiates a new AnthropicOptionsOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnthropicOptionsOneOfWithDefaults

`func NewAnthropicOptionsOneOfWithDefaults() *AnthropicOptionsOneOf`

NewAnthropicOptionsOneOfWithDefaults instantiates a new AnthropicOptionsOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeySource

`func (o *AnthropicOptionsOneOf) GetKeySource() string`

GetKeySource returns the KeySource field if non-nil, zero value otherwise.

### GetKeySourceOk

`func (o *AnthropicOptionsOneOf) GetKeySourceOk() (*string, bool)`

GetKeySourceOk returns a tuple with the KeySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySource

`func (o *AnthropicOptionsOneOf) SetKeySource(v string)`

SetKeySource sets KeySource field to given value.


### GetBaseUrl

`func (o *AnthropicOptionsOneOf) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *AnthropicOptionsOneOf) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *AnthropicOptionsOneOf) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetApiKey

`func (o *AnthropicOptionsOneOf) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *AnthropicOptionsOneOf) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *AnthropicOptionsOneOf) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetCustomHeaders

`func (o *AnthropicOptionsOneOf) GetCustomHeaders() [][]string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *AnthropicOptionsOneOf) GetCustomHeadersOk() (*[][]string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *AnthropicOptionsOneOf) SetCustomHeaders(v [][]string)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *AnthropicOptionsOneOf) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


