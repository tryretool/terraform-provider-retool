/*
Retool API

Go to Settings > API to get started. Once you generate an API token, use bearer token authentication to make requests.

API version: 2.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// checks if the SourceControlConfigGet200ResponseDataAnyOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceControlConfigGet200ResponseDataAnyOf{}

// SourceControlConfigGet200ResponseDataAnyOf struct for SourceControlConfigGet200ResponseDataAnyOf
type SourceControlConfigGet200ResponseDataAnyOf struct {
	Config SourceControlConfigGet200ResponseDataAnyOfConfig `json:"config"`
	Provider string `json:"provider"`
	// The user or organization to which the repository belongs to.
	Org string `json:"org"`
	// The name of the repository you created to use with Retool.
	Repo string `json:"repo"`
	// The default branch, e.g., main.
	DefaultBranch string `json:"default_branch"`
	// Repositories using Toolscript are 2.0.0. Repositories using legacy YAML are 1.0.0.
	RepoVersion *string `json:"repo_version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SourceControlConfigGet200ResponseDataAnyOf SourceControlConfigGet200ResponseDataAnyOf

// NewSourceControlConfigGet200ResponseDataAnyOf instantiates a new SourceControlConfigGet200ResponseDataAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceControlConfigGet200ResponseDataAnyOf(config SourceControlConfigGet200ResponseDataAnyOfConfig, provider string, org string, repo string, defaultBranch string) *SourceControlConfigGet200ResponseDataAnyOf {
	this := SourceControlConfigGet200ResponseDataAnyOf{}
	this.Config = config
	this.Provider = provider
	this.Org = org
	this.Repo = repo
	this.DefaultBranch = defaultBranch
	return &this
}

// NewSourceControlConfigGet200ResponseDataAnyOfWithDefaults instantiates a new SourceControlConfigGet200ResponseDataAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceControlConfigGet200ResponseDataAnyOfWithDefaults() *SourceControlConfigGet200ResponseDataAnyOf {
	this := SourceControlConfigGet200ResponseDataAnyOf{}
	return &this
}

// GetConfig returns the Config field value
func (o *SourceControlConfigGet200ResponseDataAnyOf) GetConfig() SourceControlConfigGet200ResponseDataAnyOfConfig {
	if o == nil {
		var ret SourceControlConfigGet200ResponseDataAnyOfConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *SourceControlConfigGet200ResponseDataAnyOf) GetConfigOk() (*SourceControlConfigGet200ResponseDataAnyOfConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *SourceControlConfigGet200ResponseDataAnyOf) SetConfig(v SourceControlConfigGet200ResponseDataAnyOfConfig) {
	o.Config = v
}

// GetProvider returns the Provider field value
func (o *SourceControlConfigGet200ResponseDataAnyOf) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *SourceControlConfigGet200ResponseDataAnyOf) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *SourceControlConfigGet200ResponseDataAnyOf) SetProvider(v string) {
	o.Provider = v
}

// GetOrg returns the Org field value
func (o *SourceControlConfigGet200ResponseDataAnyOf) GetOrg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Org
}

// GetOrgOk returns a tuple with the Org field value
// and a boolean to check if the value has been set.
func (o *SourceControlConfigGet200ResponseDataAnyOf) GetOrgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Org, true
}

// SetOrg sets field value
func (o *SourceControlConfigGet200ResponseDataAnyOf) SetOrg(v string) {
	o.Org = v
}

// GetRepo returns the Repo field value
func (o *SourceControlConfigGet200ResponseDataAnyOf) GetRepo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Repo
}

// GetRepoOk returns a tuple with the Repo field value
// and a boolean to check if the value has been set.
func (o *SourceControlConfigGet200ResponseDataAnyOf) GetRepoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repo, true
}

// SetRepo sets field value
func (o *SourceControlConfigGet200ResponseDataAnyOf) SetRepo(v string) {
	o.Repo = v
}

// GetDefaultBranch returns the DefaultBranch field value
func (o *SourceControlConfigGet200ResponseDataAnyOf) GetDefaultBranch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultBranch
}

// GetDefaultBranchOk returns a tuple with the DefaultBranch field value
// and a boolean to check if the value has been set.
func (o *SourceControlConfigGet200ResponseDataAnyOf) GetDefaultBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultBranch, true
}

// SetDefaultBranch sets field value
func (o *SourceControlConfigGet200ResponseDataAnyOf) SetDefaultBranch(v string) {
	o.DefaultBranch = v
}

// GetRepoVersion returns the RepoVersion field value if set, zero value otherwise.
func (o *SourceControlConfigGet200ResponseDataAnyOf) GetRepoVersion() string {
	if o == nil || IsNil(o.RepoVersion) {
		var ret string
		return ret
	}
	return *o.RepoVersion
}

// GetRepoVersionOk returns a tuple with the RepoVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceControlConfigGet200ResponseDataAnyOf) GetRepoVersionOk() (*string, bool) {
	if o == nil || IsNil(o.RepoVersion) {
		return nil, false
	}
	return o.RepoVersion, true
}

// HasRepoVersion returns a boolean if a field has been set.
func (o *SourceControlConfigGet200ResponseDataAnyOf) HasRepoVersion() bool {
	if o != nil && !IsNil(o.RepoVersion) {
		return true
	}

	return false
}

// SetRepoVersion gets a reference to the given string and assigns it to the RepoVersion field.
func (o *SourceControlConfigGet200ResponseDataAnyOf) SetRepoVersion(v string) {
	o.RepoVersion = &v
}

func (o SourceControlConfigGet200ResponseDataAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceControlConfigGet200ResponseDataAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["config"] = o.Config
	toSerialize["provider"] = o.Provider
	toSerialize["org"] = o.Org
	toSerialize["repo"] = o.Repo
	toSerialize["default_branch"] = o.DefaultBranch
	if !IsNil(o.RepoVersion) {
		toSerialize["repo_version"] = o.RepoVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SourceControlConfigGet200ResponseDataAnyOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"config",
		"provider",
		"org",
		"repo",
		"default_branch",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSourceControlConfigGet200ResponseDataAnyOf := _SourceControlConfigGet200ResponseDataAnyOf{}

	err = json.Unmarshal(data, &varSourceControlConfigGet200ResponseDataAnyOf)

	if err != nil {
		return err
	}

	*o = SourceControlConfigGet200ResponseDataAnyOf(varSourceControlConfigGet200ResponseDataAnyOf)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "config")
		delete(additionalProperties, "provider")
		delete(additionalProperties, "org")
		delete(additionalProperties, "repo")
		delete(additionalProperties, "default_branch")
		delete(additionalProperties, "repo_version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSourceControlConfigGet200ResponseDataAnyOf struct {
	value *SourceControlConfigGet200ResponseDataAnyOf
	isSet bool
}

func (v NullableSourceControlConfigGet200ResponseDataAnyOf) Get() *SourceControlConfigGet200ResponseDataAnyOf {
	return v.value
}

func (v *NullableSourceControlConfigGet200ResponseDataAnyOf) Set(val *SourceControlConfigGet200ResponseDataAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceControlConfigGet200ResponseDataAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceControlConfigGet200ResponseDataAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceControlConfigGet200ResponseDataAnyOf(val *SourceControlConfigGet200ResponseDataAnyOf) *NullableSourceControlConfigGet200ResponseDataAnyOf {
	return &NullableSourceControlConfigGet200ResponseDataAnyOf{value: val, isSet: true}
}

func (v NullableSourceControlConfigGet200ResponseDataAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceControlConfigGet200ResponseDataAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


