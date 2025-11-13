# SourceControlManifestsManifestNamePut200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitSha** | **string** | The SHA1 hash for the git commit that the created branch points to | 
**Url** | **string** | A URL to create a pull/merge request to merge the created branch into the default branch | 
**BranchName** | **string** | The name for the created branch | 

## Methods

### NewSourceControlManifestsManifestNamePut200ResponseData

`func NewSourceControlManifestsManifestNamePut200ResponseData(commitSha string, url string, branchName string, ) *SourceControlManifestsManifestNamePut200ResponseData`

NewSourceControlManifestsManifestNamePut200ResponseData instantiates a new SourceControlManifestsManifestNamePut200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlManifestsManifestNamePut200ResponseDataWithDefaults

`func NewSourceControlManifestsManifestNamePut200ResponseDataWithDefaults() *SourceControlManifestsManifestNamePut200ResponseData`

NewSourceControlManifestsManifestNamePut200ResponseDataWithDefaults instantiates a new SourceControlManifestsManifestNamePut200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitSha

`func (o *SourceControlManifestsManifestNamePut200ResponseData) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *SourceControlManifestsManifestNamePut200ResponseData) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *SourceControlManifestsManifestNamePut200ResponseData) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.


### GetUrl

`func (o *SourceControlManifestsManifestNamePut200ResponseData) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SourceControlManifestsManifestNamePut200ResponseData) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SourceControlManifestsManifestNamePut200ResponseData) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetBranchName

`func (o *SourceControlManifestsManifestNamePut200ResponseData) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *SourceControlManifestsManifestNamePut200ResponseData) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *SourceControlManifestsManifestNamePut200ResponseData) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


