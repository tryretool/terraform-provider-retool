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

// checks if the SourceControlConfigGet200ResponseDataOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceControlConfigGet200ResponseDataOneOf{}

// SourceControlConfigGet200ResponseDataOneOf struct for SourceControlConfigGet200ResponseDataOneOf
type SourceControlConfigGet200ResponseDataOneOf struct {
	Config SourceControlConfigGet200ResponseDataOneOfConfig `json:"config"`
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

type _SourceControlConfigGet200ResponseDataOneOf SourceControlConfigGet200ResponseDataOneOf

// NewSourceControlConfigGet200ResponseDataOneOf instantiates a new SourceControlConfigGet200ResponseDataOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceControlConfigGet200ResponseDataOneOf(config SourceControlConfigGet200ResponseDataOneOfConfig, provider string, org string, repo string, defaultBranch string) *SourceControlConfigGet200ResponseDataOneOf {
	this := SourceControlConfigGet200ResponseDataOneOf{}
	this.Config = config
	this.Provider = provider
	this.Org = org
	this.Repo = repo
	this.DefaultBranch = defaultBranch
	return &this
}

// NewSourceControlConfigGet200ResponseDataOneOfWithDefaults instantiates a new SourceControlConfigGet200ResponseDataOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceControlConfigGet200ResponseDataOneOfWithDefaults() *SourceControlConfigGet200ResponseDataOneOf {
	this := SourceControlConfigGet200ResponseDataOneOf{}
	return &this
}

// GetConfig returns the Config field value
func (o *SourceControlConfigGet200ResponseDataOneOf) GetConfig() SourceControlConfigGet200ResponseDataOneOfConfig {
	if o == nil {
		var ret SourceControlConfigGet200ResponseDataOneOfConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *SourceControlConfigGet200ResponseDataOneOf) GetConfigOk() (*SourceControlConfigGet200ResponseDataOneOfConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *SourceControlConfigGet200ResponseDataOneOf) SetConfig(v SourceControlConfigGet200ResponseDataOneOfConfig) {
	o.Config = v
}

// GetProvider returns the Provider field value
func (o *SourceControlConfigGet200ResponseDataOneOf) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *SourceControlConfigGet200ResponseDataOneOf) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *SourceControlConfigGet200ResponseDataOneOf) SetProvider(v string) {
	o.Provider = v
}

// GetOrg returns the Org field value
func (o *SourceControlConfigGet200ResponseDataOneOf) GetOrg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Org
}

// GetOrgOk returns a tuple with the Org field value
// and a boolean to check if the value has been set.
func (o *SourceControlConfigGet200ResponseDataOneOf) GetOrgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Org, true
}

// SetOrg sets field value
func (o *SourceControlConfigGet200ResponseDataOneOf) SetOrg(v string) {
	o.Org = v
}

// GetRepo returns the Repo field value
func (o *SourceControlConfigGet200ResponseDataOneOf) GetRepo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Repo
}

// GetRepoOk returns a tuple with the Repo field value
// and a boolean to check if the value has been set.
func (o *SourceControlConfigGet200ResponseDataOneOf) GetRepoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repo, true
}

// SetRepo sets field value
func (o *SourceControlConfigGet200ResponseDataOneOf) SetRepo(v string) {
	o.Repo = v
}

// GetDefaultBranch returns the DefaultBranch field value
func (o *SourceControlConfigGet200ResponseDataOneOf) GetDefaultBranch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultBranch
}

// GetDefaultBranchOk returns a tuple with the DefaultBranch field value
// and a boolean to check if the value has been set.
func (o *SourceControlConfigGet200ResponseDataOneOf) GetDefaultBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultBranch, true
}

// SetDefaultBranch sets field value
func (o *SourceControlConfigGet200ResponseDataOneOf) SetDefaultBranch(v string) {
	o.DefaultBranch = v
}

// GetRepoVersion returns the RepoVersion field value if set, zero value otherwise.
func (o *SourceControlConfigGet200ResponseDataOneOf) GetRepoVersion() string {
	if o == nil || IsNil(o.RepoVersion) {
		var ret string
		return ret
	}
	return *o.RepoVersion
}

// GetRepoVersionOk returns a tuple with the RepoVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceControlConfigGet200ResponseDataOneOf) GetRepoVersionOk() (*string, bool) {
	if o == nil || IsNil(o.RepoVersion) {
		return nil, false
	}
	return o.RepoVersion, true
}

// HasRepoVersion returns a boolean if a field has been set.
func (o *SourceControlConfigGet200ResponseDataOneOf) HasRepoVersion() bool {
	if o != nil && !IsNil(o.RepoVersion) {
		return true
	}

	return false
}

// SetRepoVersion gets a reference to the given string and assigns it to the RepoVersion field.
func (o *SourceControlConfigGet200ResponseDataOneOf) SetRepoVersion(v string) {
	o.RepoVersion = &v
}

func (o SourceControlConfigGet200ResponseDataOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceControlConfigGet200ResponseDataOneOf) ToMap() (map[string]interface{}, error) {
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

func (o *SourceControlConfigGet200ResponseDataOneOf) UnmarshalJSON(data []byte) (err error) {
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

	varSourceControlConfigGet200ResponseDataOneOf := _SourceControlConfigGet200ResponseDataOneOf{}

	err = json.Unmarshal(data, &varSourceControlConfigGet200ResponseDataOneOf)

	if err != nil {
		return err
	}

	*o = SourceControlConfigGet200ResponseDataOneOf(varSourceControlConfigGet200ResponseDataOneOf)

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

type NullableSourceControlConfigGet200ResponseDataOneOf struct {
	value *SourceControlConfigGet200ResponseDataOneOf
	isSet bool
}

func (v NullableSourceControlConfigGet200ResponseDataOneOf) Get() *SourceControlConfigGet200ResponseDataOneOf {
	return v.value
}

func (v *NullableSourceControlConfigGet200ResponseDataOneOf) Set(val *SourceControlConfigGet200ResponseDataOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceControlConfigGet200ResponseDataOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceControlConfigGet200ResponseDataOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceControlConfigGet200ResponseDataOneOf(val *SourceControlConfigGet200ResponseDataOneOf) *NullableSourceControlConfigGet200ResponseDataOneOf {
	return &NullableSourceControlConfigGet200ResponseDataOneOf{value: val, isSet: true}
}

func (v NullableSourceControlConfigGet200ResponseDataOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceControlConfigGet200ResponseDataOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


