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

// checks if the FoldersGet200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FoldersGet200ResponseDataInner{}

// FoldersGet200ResponseDataInner Folders
type FoldersGet200ResponseDataInner struct {
	// The id of the folder. Currently this is the same as legacy_id but will be different in the future.
	Id string `json:"id"`
	// The legacy id of the folder.
	LegacyId string `json:"legacy_id"`
	// The name of the folder
	Name string `json:"name"`
	// The id of the parent folder
	ParentFolderId NullableString `json:"parent_folder_id,omitempty"`
	// Whether the folder is a system folder
	IsSystemFolder bool `json:"is_system_folder"`
	// The type of the folder
	FolderType string `json:"folder_type"`
}

type _FoldersGet200ResponseDataInner FoldersGet200ResponseDataInner

// NewFoldersGet200ResponseDataInner instantiates a new FoldersGet200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFoldersGet200ResponseDataInner(id string, legacyId string, name string, isSystemFolder bool, folderType string) *FoldersGet200ResponseDataInner {
	this := FoldersGet200ResponseDataInner{}
	this.Id = id
	this.LegacyId = legacyId
	this.Name = name
	this.IsSystemFolder = isSystemFolder
	this.FolderType = folderType
	return &this
}

// NewFoldersGet200ResponseDataInnerWithDefaults instantiates a new FoldersGet200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFoldersGet200ResponseDataInnerWithDefaults() *FoldersGet200ResponseDataInner {
	this := FoldersGet200ResponseDataInner{}
	return &this
}

// GetId returns the Id field value
func (o *FoldersGet200ResponseDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FoldersGet200ResponseDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FoldersGet200ResponseDataInner) SetId(v string) {
	o.Id = v
}

// GetLegacyId returns the LegacyId field value
func (o *FoldersGet200ResponseDataInner) GetLegacyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LegacyId
}

// GetLegacyIdOk returns a tuple with the LegacyId field value
// and a boolean to check if the value has been set.
func (o *FoldersGet200ResponseDataInner) GetLegacyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LegacyId, true
}

// SetLegacyId sets field value
func (o *FoldersGet200ResponseDataInner) SetLegacyId(v string) {
	o.LegacyId = v
}

// GetName returns the Name field value
func (o *FoldersGet200ResponseDataInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FoldersGet200ResponseDataInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FoldersGet200ResponseDataInner) SetName(v string) {
	o.Name = v
}

// GetParentFolderId returns the ParentFolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FoldersGet200ResponseDataInner) GetParentFolderId() string {
	if o == nil || IsNil(o.ParentFolderId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentFolderId.Get()
}

// GetParentFolderIdOk returns a tuple with the ParentFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FoldersGet200ResponseDataInner) GetParentFolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentFolderId.Get(), o.ParentFolderId.IsSet()
}

// HasParentFolderId returns a boolean if a field has been set.
func (o *FoldersGet200ResponseDataInner) HasParentFolderId() bool {
	if o != nil && o.ParentFolderId.IsSet() {
		return true
	}

	return false
}

// SetParentFolderId gets a reference to the given NullableString and assigns it to the ParentFolderId field.
func (o *FoldersGet200ResponseDataInner) SetParentFolderId(v string) {
	o.ParentFolderId.Set(&v)
}
// SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil
func (o *FoldersGet200ResponseDataInner) SetParentFolderIdNil() {
	o.ParentFolderId.Set(nil)
}

// UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil
func (o *FoldersGet200ResponseDataInner) UnsetParentFolderId() {
	o.ParentFolderId.Unset()
}

// GetIsSystemFolder returns the IsSystemFolder field value
func (o *FoldersGet200ResponseDataInner) GetIsSystemFolder() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSystemFolder
}

// GetIsSystemFolderOk returns a tuple with the IsSystemFolder field value
// and a boolean to check if the value has been set.
func (o *FoldersGet200ResponseDataInner) GetIsSystemFolderOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSystemFolder, true
}

// SetIsSystemFolder sets field value
func (o *FoldersGet200ResponseDataInner) SetIsSystemFolder(v bool) {
	o.IsSystemFolder = v
}

// GetFolderType returns the FolderType field value
func (o *FoldersGet200ResponseDataInner) GetFolderType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FolderType
}

// GetFolderTypeOk returns a tuple with the FolderType field value
// and a boolean to check if the value has been set.
func (o *FoldersGet200ResponseDataInner) GetFolderTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FolderType, true
}

// SetFolderType sets field value
func (o *FoldersGet200ResponseDataInner) SetFolderType(v string) {
	o.FolderType = v
}

func (o FoldersGet200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FoldersGet200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["legacy_id"] = o.LegacyId
	toSerialize["name"] = o.Name
	if o.ParentFolderId.IsSet() {
		toSerialize["parent_folder_id"] = o.ParentFolderId.Get()
	}
	toSerialize["is_system_folder"] = o.IsSystemFolder
	toSerialize["folder_type"] = o.FolderType
	return toSerialize, nil
}

func (o *FoldersGet200ResponseDataInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"legacy_id",
		"name",
		"is_system_folder",
		"folder_type",
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

	varFoldersGet200ResponseDataInner := _FoldersGet200ResponseDataInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFoldersGet200ResponseDataInner)

	if err != nil {
		return err
	}

	*o = FoldersGet200ResponseDataInner(varFoldersGet200ResponseDataInner)

	return err
}

type NullableFoldersGet200ResponseDataInner struct {
	value *FoldersGet200ResponseDataInner
	isSet bool
}

func (v NullableFoldersGet200ResponseDataInner) Get() *FoldersGet200ResponseDataInner {
	return v.value
}

func (v *NullableFoldersGet200ResponseDataInner) Set(val *FoldersGet200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFoldersGet200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFoldersGet200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFoldersGet200ResponseDataInner(val *FoldersGet200ResponseDataInner) *NullableFoldersGet200ResponseDataInner {
	return &NullableFoldersGet200ResponseDataInner{value: val, isSet: true}
}

func (v NullableFoldersGet200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFoldersGet200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


