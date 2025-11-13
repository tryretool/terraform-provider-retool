# SourceControlManifestsManifestNameDeletePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitMessage** | Pointer to **string** | Message to use for the commit that updates the specified manifest. If a message is not provided, a default will be used. | [optional] 

## Methods

### NewSourceControlManifestsManifestNameDeletePostRequest

`func NewSourceControlManifestsManifestNameDeletePostRequest() *SourceControlManifestsManifestNameDeletePostRequest`

NewSourceControlManifestsManifestNameDeletePostRequest instantiates a new SourceControlManifestsManifestNameDeletePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlManifestsManifestNameDeletePostRequestWithDefaults

`func NewSourceControlManifestsManifestNameDeletePostRequestWithDefaults() *SourceControlManifestsManifestNameDeletePostRequest`

NewSourceControlManifestsManifestNameDeletePostRequestWithDefaults instantiates a new SourceControlManifestsManifestNameDeletePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitMessage

`func (o *SourceControlManifestsManifestNameDeletePostRequest) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *SourceControlManifestsManifestNameDeletePostRequest) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *SourceControlManifestsManifestNameDeletePostRequest) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.

### HasCommitMessage

`func (o *SourceControlManifestsManifestNameDeletePostRequest) HasCommitMessage() bool`

HasCommitMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


