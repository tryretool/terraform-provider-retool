/*
Retool API

Go to Settings > API to get started. Once you generate an API token, use bearer token authentication to make requests.

API version: 2.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SourceControlConfigGet200ResponseDataAnyOf4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceControlConfigGet200ResponseDataAnyOf4{}

// SourceControlConfigGet200ResponseDataAnyOf4 struct for SourceControlConfigGet200ResponseDataAnyOf4
type SourceControlConfigGet200ResponseDataAnyOf4 struct {
	Config SourceControlConfigGet200ResponseDataAnyOf4Config `json:"config"`
	Provider string `json:"provider"`
	// The user or organization to which the repository belongs to.
	Org string `json:"org"`
	// The name of the repository you created to use with Retool.
	Repo string `json:"repo"`
	// The default branch, e.g., main.
	DefaultBranch string `json:"default_branch"`
	// Repositories using Toolscript are 2.0.0. Repositories using legacy YAML are 1.0.0.
	RepoVersion *string `json:"repo_version,omitempty"`
}

type _SourceControlConfigGet200ResponseDataAnyOf4 SourceControlConfigGet200ResponseDataAnyOf4

// NewSourceControlConfigGet200ResponseDataAnyOf4 instantiates a new SourceControlConfigGet200ResponseDataAnyOf4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceControlConfigGet200ResponseDataAnyOf4(config SourceControlConfigGet200ResponseDataAnyOf4Config, provider string, org string, repo string, defaultBranch string) *SourceControlConfigGet200ResponseDataAnyOf4 {
	this := SourceControlConfigGet200ResponseDataAnyOf4{}
	this.Config = config
	this.Provider = provider
	this.Org = org
	this.Repo = repo
	this.DefaultBranch = defaultBranch
	return &this
}

// NewSourceControlConfigGet200ResponseDataAnyOf4WithDefaults instantiates a new SourceControlConfigGet200ResponseDataAnyOf4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceControlConfigGet200ResponseDataAnyOf4WithDefaults() *SourceControlConfigGet200ResponseDataAnyOf4 {
	this := SourceControlConfigGet200ResponseDataAnyOf4{}
	return &this
}

// GetConfig returns the Config field value
func (o *SourceControlConfigGet200ResponseDataAnyOf4) GetConfig() SourceControlConfigGet200ResponseDataAnyOf4Config {
	if o == nil {
		var ret SourceControlConfigGet200ResponseDataAnyOf4Config
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *SourceControlConfigGet200ResponseDataAnyOf4) GetConfigOk() (*SourceControlConfigGet200ResponseDataAnyOf4Config, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *SourceControlConfigGet200ResponseDataAnyOf4) SetConfig(v SourceControlConfigGet200ResponseDataAnyOf4Config) {
	o.Config = v
}

// GetProvider returns the Provider field value
func (o *SourceControlConfigGet200ResponseDataAnyOf4) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *SourceControlConfigGet200ResponseDataAnyOf4) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *SourceControlConfigGet200ResponseDataAnyOf4) SetProvider(v string) {
	o.Provider = v
}

// GetOrg returns the Org field value
func (o *SourceControlConfigGet200ResponseDataAnyOf4) GetOrg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Org
}

// GetOrgOk returns a tuple with the Org field value
// and a boolean to check if the value has been set.
func (o *SourceControlConfigGet200ResponseDataAnyOf4) GetOrgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Org, true
}

// SetOrg sets field value
func (o *SourceControlConfigGet200ResponseDataAnyOf4) SetOrg(v string) {
	o.Org = v
}

// GetRepo returns the Repo field value
func (o *SourceControlConfigGet200ResponseDataAnyOf4) GetRepo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Repo
}

// GetRepoOk returns a tuple with the Repo field value
// and a boolean to check if the value has been set.
func (o *SourceControlConfigGet200ResponseDataAnyOf4) GetRepoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repo, true
}

// SetRepo sets field value
func (o *SourceControlConfigGet200ResponseDataAnyOf4) SetRepo(v string) {
	o.Repo = v
}

// GetDefaultBranch returns the DefaultBranch field value
func (o *SourceControlConfigGet200ResponseDataAnyOf4) GetDefaultBranch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultBranch
}

// GetDefaultBranchOk returns a tuple with the DefaultBranch field value
// and a boolean to check if the value has been set.
func (o *SourceControlConfigGet200ResponseDataAnyOf4) GetDefaultBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultBranch, true
}

// SetDefaultBranch sets field value
func (o *SourceControlConfigGet200ResponseDataAnyOf4) SetDefaultBranch(v string) {
	o.DefaultBranch = v
}

// GetRepoVersion returns the RepoVersion field value if set, zero value otherwise.
func (o *SourceControlConfigGet200ResponseDataAnyOf4) GetRepoVersion() string {
	if o == nil || IsNil(o.RepoVersion) {
		var ret string
		return ret
	}
	return *o.RepoVersion
}

// GetRepoVersionOk returns a tuple with the RepoVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceControlConfigGet200ResponseDataAnyOf4) GetRepoVersionOk() (*string, bool) {
	if o == nil || IsNil(o.RepoVersion) {
		return nil, false
	}
	return o.RepoVersion, true
}

// HasRepoVersion returns a boolean if a field has been set.
func (o *SourceControlConfigGet200ResponseDataAnyOf4) HasRepoVersion() bool {
	if o != nil && !IsNil(o.RepoVersion) {
		return true
	}

	return false
}

// SetRepoVersion gets a reference to the given string and assigns it to the RepoVersion field.
func (o *SourceControlConfigGet200ResponseDataAnyOf4) SetRepoVersion(v string) {
	o.RepoVersion = &v
}

func (o SourceControlConfigGet200ResponseDataAnyOf4) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceControlConfigGet200ResponseDataAnyOf4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["config"] = o.Config
	toSerialize["provider"] = o.Provider
	toSerialize["org"] = o.Org
	toSerialize["repo"] = o.Repo
	toSerialize["default_branch"] = o.DefaultBranch
	if !IsNil(o.RepoVersion) {
		toSerialize["repo_version"] = o.RepoVersion
	}
	return toSerialize, nil
}

func (o *SourceControlConfigGet200ResponseDataAnyOf4) UnmarshalJSON(data []byte) (err error) {
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

	varSourceControlConfigGet200ResponseDataAnyOf4 := _SourceControlConfigGet200ResponseDataAnyOf4{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSourceControlConfigGet200ResponseDataAnyOf4)

	if err != nil {
		return err
	}

	*o = SourceControlConfigGet200ResponseDataAnyOf4(varSourceControlConfigGet200ResponseDataAnyOf4)

	return err
}

type NullableSourceControlConfigGet200ResponseDataAnyOf4 struct {
	value *SourceControlConfigGet200ResponseDataAnyOf4
	isSet bool
}

func (v NullableSourceControlConfigGet200ResponseDataAnyOf4) Get() *SourceControlConfigGet200ResponseDataAnyOf4 {
	return v.value
}

func (v *NullableSourceControlConfigGet200ResponseDataAnyOf4) Set(val *SourceControlConfigGet200ResponseDataAnyOf4) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceControlConfigGet200ResponseDataAnyOf4) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceControlConfigGet200ResponseDataAnyOf4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceControlConfigGet200ResponseDataAnyOf4(val *SourceControlConfigGet200ResponseDataAnyOf4) *NullableSourceControlConfigGet200ResponseDataAnyOf4 {
	return &NullableSourceControlConfigGet200ResponseDataAnyOf4{value: val, isSet: true}
}

func (v NullableSourceControlConfigGet200ResponseDataAnyOf4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceControlConfigGet200ResponseDataAnyOf4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


