/*
Retool API

Go to Settings > API to get started. Once you generate an API token, use bearer token authentication to make requests.

API version: 2.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the CustomComponentLibraryRevisionFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomComponentLibraryRevisionFile{}

// CustomComponentLibraryRevisionFile Custom Component Library Revision
type CustomComponentLibraryRevisionFile struct {
	Filepath string `json:"filepath"`
	FileContents string `json:"file_contents"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type _CustomComponentLibraryRevisionFile CustomComponentLibraryRevisionFile

// NewCustomComponentLibraryRevisionFile instantiates a new CustomComponentLibraryRevisionFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomComponentLibraryRevisionFile(filepath string, fileContents string, createdAt time.Time, updatedAt time.Time) *CustomComponentLibraryRevisionFile {
	this := CustomComponentLibraryRevisionFile{}
	this.Filepath = filepath
	this.FileContents = fileContents
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewCustomComponentLibraryRevisionFileWithDefaults instantiates a new CustomComponentLibraryRevisionFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomComponentLibraryRevisionFileWithDefaults() *CustomComponentLibraryRevisionFile {
	this := CustomComponentLibraryRevisionFile{}
	return &this
}

// GetFilepath returns the Filepath field value
func (o *CustomComponentLibraryRevisionFile) GetFilepath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filepath
}

// GetFilepathOk returns a tuple with the Filepath field value
// and a boolean to check if the value has been set.
func (o *CustomComponentLibraryRevisionFile) GetFilepathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filepath, true
}

// SetFilepath sets field value
func (o *CustomComponentLibraryRevisionFile) SetFilepath(v string) {
	o.Filepath = v
}

// GetFileContents returns the FileContents field value
func (o *CustomComponentLibraryRevisionFile) GetFileContents() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileContents
}

// GetFileContentsOk returns a tuple with the FileContents field value
// and a boolean to check if the value has been set.
func (o *CustomComponentLibraryRevisionFile) GetFileContentsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileContents, true
}

// SetFileContents sets field value
func (o *CustomComponentLibraryRevisionFile) SetFileContents(v string) {
	o.FileContents = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CustomComponentLibraryRevisionFile) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CustomComponentLibraryRevisionFile) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CustomComponentLibraryRevisionFile) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *CustomComponentLibraryRevisionFile) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CustomComponentLibraryRevisionFile) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *CustomComponentLibraryRevisionFile) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o CustomComponentLibraryRevisionFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomComponentLibraryRevisionFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filepath"] = o.Filepath
	toSerialize["file_contents"] = o.FileContents
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *CustomComponentLibraryRevisionFile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filepath",
		"file_contents",
		"created_at",
		"updated_at",
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

	varCustomComponentLibraryRevisionFile := _CustomComponentLibraryRevisionFile{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomComponentLibraryRevisionFile)

	if err != nil {
		return err
	}

	*o = CustomComponentLibraryRevisionFile(varCustomComponentLibraryRevisionFile)

	return err
}

type NullableCustomComponentLibraryRevisionFile struct {
	value *CustomComponentLibraryRevisionFile
	isSet bool
}

func (v NullableCustomComponentLibraryRevisionFile) Get() *CustomComponentLibraryRevisionFile {
	return v.value
}

func (v *NullableCustomComponentLibraryRevisionFile) Set(val *CustomComponentLibraryRevisionFile) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomComponentLibraryRevisionFile) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomComponentLibraryRevisionFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomComponentLibraryRevisionFile(val *CustomComponentLibraryRevisionFile) *NullableCustomComponentLibraryRevisionFile {
	return &NullableCustomComponentLibraryRevisionFile{value: val, isSet: true}
}

func (v NullableCustomComponentLibraryRevisionFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomComponentLibraryRevisionFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


