# SourceControlManifestsManifestNamePutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Manifest** | Pointer to [**SourceControlManifestsManifestNamePutRequestManifest**](SourceControlManifestsManifestNamePutRequestManifest.md) |  | [optional] 
**CommitMessage** | Pointer to **string** | Message to use for the commit that updates the specified manifest. If a message is not provided, a default will be used. | [optional] 

## Methods

### NewSourceControlManifestsManifestNamePutRequest

`func NewSourceControlManifestsManifestNamePutRequest() *SourceControlManifestsManifestNamePutRequest`

NewSourceControlManifestsManifestNamePutRequest instantiates a new SourceControlManifestsManifestNamePutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlManifestsManifestNamePutRequestWithDefaults

`func NewSourceControlManifestsManifestNamePutRequestWithDefaults() *SourceControlManifestsManifestNamePutRequest`

NewSourceControlManifestsManifestNamePutRequestWithDefaults instantiates a new SourceControlManifestsManifestNamePutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManifest

`func (o *SourceControlManifestsManifestNamePutRequest) GetManifest() SourceControlManifestsManifestNamePutRequestManifest`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *SourceControlManifestsManifestNamePutRequest) GetManifestOk() (*SourceControlManifestsManifestNamePutRequestManifest, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *SourceControlManifestsManifestNamePutRequest) SetManifest(v SourceControlManifestsManifestNamePutRequestManifest)`

SetManifest sets Manifest field to given value.

### HasManifest

`func (o *SourceControlManifestsManifestNamePutRequest) HasManifest() bool`

HasManifest returns a boolean if a field has been set.

### GetCommitMessage

`func (o *SourceControlManifestsManifestNamePutRequest) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *SourceControlManifestsManifestNamePutRequest) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *SourceControlManifestsManifestNamePutRequest) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.

### HasCommitMessage

`func (o *SourceControlManifestsManifestNamePutRequest) HasCommitMessage() bool`

HasCommitMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


