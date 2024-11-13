# SourceControlSettingsPut200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoBranchNamingEnabled** | **bool** | When enabled, Retool automatically suggests a branch name on branch creation. Defaults to true. | 
**CustomPullRequestTemplateEnabled** | **bool** | When enabled, Retool will use the template specified to create pull requests. Defaults to false. | 
**CustomPullRequestTemplate** | **string** | Pull requests created from Retool will use the template specified. | 
**VersionControlLocked** | **bool** | When set to true, creates a read-only instance of Retool, where app editing is disabled. Defaults to false. | 
**ForceUuidMapping** | **bool** | When set to true, creates a uuid mapping for protected elements to be used in the source control repo. Defaults to false. | 

## Methods

### NewSourceControlSettingsPut200ResponseData

`func NewSourceControlSettingsPut200ResponseData(autoBranchNamingEnabled bool, customPullRequestTemplateEnabled bool, customPullRequestTemplate string, versionControlLocked bool, forceUuidMapping bool, ) *SourceControlSettingsPut200ResponseData`

NewSourceControlSettingsPut200ResponseData instantiates a new SourceControlSettingsPut200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlSettingsPut200ResponseDataWithDefaults

`func NewSourceControlSettingsPut200ResponseDataWithDefaults() *SourceControlSettingsPut200ResponseData`

NewSourceControlSettingsPut200ResponseDataWithDefaults instantiates a new SourceControlSettingsPut200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoBranchNamingEnabled

`func (o *SourceControlSettingsPut200ResponseData) GetAutoBranchNamingEnabled() bool`

GetAutoBranchNamingEnabled returns the AutoBranchNamingEnabled field if non-nil, zero value otherwise.

### GetAutoBranchNamingEnabledOk

`func (o *SourceControlSettingsPut200ResponseData) GetAutoBranchNamingEnabledOk() (*bool, bool)`

GetAutoBranchNamingEnabledOk returns a tuple with the AutoBranchNamingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBranchNamingEnabled

`func (o *SourceControlSettingsPut200ResponseData) SetAutoBranchNamingEnabled(v bool)`

SetAutoBranchNamingEnabled sets AutoBranchNamingEnabled field to given value.


### GetCustomPullRequestTemplateEnabled

`func (o *SourceControlSettingsPut200ResponseData) GetCustomPullRequestTemplateEnabled() bool`

GetCustomPullRequestTemplateEnabled returns the CustomPullRequestTemplateEnabled field if non-nil, zero value otherwise.

### GetCustomPullRequestTemplateEnabledOk

`func (o *SourceControlSettingsPut200ResponseData) GetCustomPullRequestTemplateEnabledOk() (*bool, bool)`

GetCustomPullRequestTemplateEnabledOk returns a tuple with the CustomPullRequestTemplateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPullRequestTemplateEnabled

`func (o *SourceControlSettingsPut200ResponseData) SetCustomPullRequestTemplateEnabled(v bool)`

SetCustomPullRequestTemplateEnabled sets CustomPullRequestTemplateEnabled field to given value.


### GetCustomPullRequestTemplate

`func (o *SourceControlSettingsPut200ResponseData) GetCustomPullRequestTemplate() string`

GetCustomPullRequestTemplate returns the CustomPullRequestTemplate field if non-nil, zero value otherwise.

### GetCustomPullRequestTemplateOk

`func (o *SourceControlSettingsPut200ResponseData) GetCustomPullRequestTemplateOk() (*string, bool)`

GetCustomPullRequestTemplateOk returns a tuple with the CustomPullRequestTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPullRequestTemplate

`func (o *SourceControlSettingsPut200ResponseData) SetCustomPullRequestTemplate(v string)`

SetCustomPullRequestTemplate sets CustomPullRequestTemplate field to given value.


### GetVersionControlLocked

`func (o *SourceControlSettingsPut200ResponseData) GetVersionControlLocked() bool`

GetVersionControlLocked returns the VersionControlLocked field if non-nil, zero value otherwise.

### GetVersionControlLockedOk

`func (o *SourceControlSettingsPut200ResponseData) GetVersionControlLockedOk() (*bool, bool)`

GetVersionControlLockedOk returns a tuple with the VersionControlLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionControlLocked

`func (o *SourceControlSettingsPut200ResponseData) SetVersionControlLocked(v bool)`

SetVersionControlLocked sets VersionControlLocked field to given value.


### GetForceUuidMapping

`func (o *SourceControlSettingsPut200ResponseData) GetForceUuidMapping() bool`

GetForceUuidMapping returns the ForceUuidMapping field if non-nil, zero value otherwise.

### GetForceUuidMappingOk

`func (o *SourceControlSettingsPut200ResponseData) GetForceUuidMappingOk() (*bool, bool)`

GetForceUuidMappingOk returns a tuple with the ForceUuidMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceUuidMapping

`func (o *SourceControlSettingsPut200ResponseData) SetForceUuidMapping(v bool)`

SetForceUuidMapping sets ForceUuidMapping field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


