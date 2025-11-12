# SourceControlSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoBranchNamingEnabled** | **bool** | When enabled, Retool automatically suggests a branch name on branch creation. Defaults to true. | 
**CustomPullRequestTemplateEnabled** | **bool** | When enabled, Retool will use the template specified to create pull requests. Defaults to false. | 
**CustomPullRequestTemplate** | **string** | Pull requests created from Retool will use the template specified. | 
**VersionControlLocked** | **bool** | When set to true, creates a read-only instance of Retool, where app editing is disabled. Defaults to false. | 
**ForceUuidMapping** | **bool** | When set to true, creates a uuid mapping for protected elements to be used in the source control repo. Defaults to false. | 
**AutoCleanupBranchesEnabled** | **bool** | When set to true, Retool will automatically delete branches after changes are merged if the remote branch no longer exists and there are no uncommitted changes. Defaults to true. | 

## Methods

### NewSourceControlSettings

`func NewSourceControlSettings(autoBranchNamingEnabled bool, customPullRequestTemplateEnabled bool, customPullRequestTemplate string, versionControlLocked bool, forceUuidMapping bool, autoCleanupBranchesEnabled bool, ) *SourceControlSettings`

NewSourceControlSettings instantiates a new SourceControlSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlSettingsWithDefaults

`func NewSourceControlSettingsWithDefaults() *SourceControlSettings`

NewSourceControlSettingsWithDefaults instantiates a new SourceControlSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoBranchNamingEnabled

`func (o *SourceControlSettings) GetAutoBranchNamingEnabled() bool`

GetAutoBranchNamingEnabled returns the AutoBranchNamingEnabled field if non-nil, zero value otherwise.

### GetAutoBranchNamingEnabledOk

`func (o *SourceControlSettings) GetAutoBranchNamingEnabledOk() (*bool, bool)`

GetAutoBranchNamingEnabledOk returns a tuple with the AutoBranchNamingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBranchNamingEnabled

`func (o *SourceControlSettings) SetAutoBranchNamingEnabled(v bool)`

SetAutoBranchNamingEnabled sets AutoBranchNamingEnabled field to given value.


### GetCustomPullRequestTemplateEnabled

`func (o *SourceControlSettings) GetCustomPullRequestTemplateEnabled() bool`

GetCustomPullRequestTemplateEnabled returns the CustomPullRequestTemplateEnabled field if non-nil, zero value otherwise.

### GetCustomPullRequestTemplateEnabledOk

`func (o *SourceControlSettings) GetCustomPullRequestTemplateEnabledOk() (*bool, bool)`

GetCustomPullRequestTemplateEnabledOk returns a tuple with the CustomPullRequestTemplateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPullRequestTemplateEnabled

`func (o *SourceControlSettings) SetCustomPullRequestTemplateEnabled(v bool)`

SetCustomPullRequestTemplateEnabled sets CustomPullRequestTemplateEnabled field to given value.


### GetCustomPullRequestTemplate

`func (o *SourceControlSettings) GetCustomPullRequestTemplate() string`

GetCustomPullRequestTemplate returns the CustomPullRequestTemplate field if non-nil, zero value otherwise.

### GetCustomPullRequestTemplateOk

`func (o *SourceControlSettings) GetCustomPullRequestTemplateOk() (*string, bool)`

GetCustomPullRequestTemplateOk returns a tuple with the CustomPullRequestTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPullRequestTemplate

`func (o *SourceControlSettings) SetCustomPullRequestTemplate(v string)`

SetCustomPullRequestTemplate sets CustomPullRequestTemplate field to given value.


### GetVersionControlLocked

`func (o *SourceControlSettings) GetVersionControlLocked() bool`

GetVersionControlLocked returns the VersionControlLocked field if non-nil, zero value otherwise.

### GetVersionControlLockedOk

`func (o *SourceControlSettings) GetVersionControlLockedOk() (*bool, bool)`

GetVersionControlLockedOk returns a tuple with the VersionControlLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionControlLocked

`func (o *SourceControlSettings) SetVersionControlLocked(v bool)`

SetVersionControlLocked sets VersionControlLocked field to given value.


### GetForceUuidMapping

`func (o *SourceControlSettings) GetForceUuidMapping() bool`

GetForceUuidMapping returns the ForceUuidMapping field if non-nil, zero value otherwise.

### GetForceUuidMappingOk

`func (o *SourceControlSettings) GetForceUuidMappingOk() (*bool, bool)`

GetForceUuidMappingOk returns a tuple with the ForceUuidMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceUuidMapping

`func (o *SourceControlSettings) SetForceUuidMapping(v bool)`

SetForceUuidMapping sets ForceUuidMapping field to given value.


### GetAutoCleanupBranchesEnabled

`func (o *SourceControlSettings) GetAutoCleanupBranchesEnabled() bool`

GetAutoCleanupBranchesEnabled returns the AutoCleanupBranchesEnabled field if non-nil, zero value otherwise.

### GetAutoCleanupBranchesEnabledOk

`func (o *SourceControlSettings) GetAutoCleanupBranchesEnabledOk() (*bool, bool)`

GetAutoCleanupBranchesEnabledOk returns a tuple with the AutoCleanupBranchesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCleanupBranchesEnabled

`func (o *SourceControlSettings) SetAutoCleanupBranchesEnabled(v bool)`

SetAutoCleanupBranchesEnabled sets AutoCleanupBranchesEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


