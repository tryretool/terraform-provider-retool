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

// SourceControlConfigurationAnyOfConfig struct for SourceControlConfigurationAnyOfConfig
type SourceControlConfigurationAnyOfConfig struct {
	SourceControlConfigGet200ResponseDataOneOfConfigOneOf *SourceControlConfigGet200ResponseDataOneOfConfigOneOf
	SourceControlConfigGet200ResponseDataOneOfConfigOneOf1 *SourceControlConfigGet200ResponseDataOneOfConfigOneOf1
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SourceControlConfigurationAnyOfConfig) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SourceControlConfigGet200ResponseDataOneOfConfigOneOf
	err = json.Unmarshal(data, &dst.SourceControlConfigGet200ResponseDataOneOfConfigOneOf);
	if err == nil {
		jsonSourceControlConfigGet200ResponseDataOneOfConfigOneOf, _ := json.Marshal(dst.SourceControlConfigGet200ResponseDataOneOfConfigOneOf)
		if string(jsonSourceControlConfigGet200ResponseDataOneOfConfigOneOf) == "{}" { // empty struct
			dst.SourceControlConfigGet200ResponseDataOneOfConfigOneOf = nil
		} else {
			return nil // data stored in dst.SourceControlConfigGet200ResponseDataOneOfConfigOneOf, return on the first match
		}
	} else {
		dst.SourceControlConfigGet200ResponseDataOneOfConfigOneOf = nil
	}

	// try to unmarshal JSON data into SourceControlConfigGet200ResponseDataOneOfConfigOneOf1
	err = json.Unmarshal(data, &dst.SourceControlConfigGet200ResponseDataOneOfConfigOneOf1);
	if err == nil {
		jsonSourceControlConfigGet200ResponseDataOneOfConfigOneOf1, _ := json.Marshal(dst.SourceControlConfigGet200ResponseDataOneOfConfigOneOf1)
		if string(jsonSourceControlConfigGet200ResponseDataOneOfConfigOneOf1) == "{}" { // empty struct
			dst.SourceControlConfigGet200ResponseDataOneOfConfigOneOf1 = nil
		} else {
			return nil // data stored in dst.SourceControlConfigGet200ResponseDataOneOfConfigOneOf1, return on the first match
		}
	} else {
		dst.SourceControlConfigGet200ResponseDataOneOfConfigOneOf1 = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SourceControlConfigurationAnyOfConfig)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SourceControlConfigurationAnyOfConfig) MarshalJSON() ([]byte, error) {
	if src.SourceControlConfigGet200ResponseDataOneOfConfigOneOf != nil {
		return json.Marshal(&src.SourceControlConfigGet200ResponseDataOneOfConfigOneOf)
	}

	if src.SourceControlConfigGet200ResponseDataOneOfConfigOneOf1 != nil {
		return json.Marshal(&src.SourceControlConfigGet200ResponseDataOneOfConfigOneOf1)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSourceControlConfigurationAnyOfConfig struct {
	value *SourceControlConfigurationAnyOfConfig
	isSet bool
}

func (v NullableSourceControlConfigurationAnyOfConfig) Get() *SourceControlConfigurationAnyOfConfig {
	return v.value
}

func (v *NullableSourceControlConfigurationAnyOfConfig) Set(val *SourceControlConfigurationAnyOfConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceControlConfigurationAnyOfConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceControlConfigurationAnyOfConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceControlConfigurationAnyOfConfig(val *SourceControlConfigurationAnyOfConfig) *NullableSourceControlConfigurationAnyOfConfig {
	return &NullableSourceControlConfigurationAnyOfConfig{value: val, isSet: true}
}

func (v NullableSourceControlConfigurationAnyOfConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceControlConfigurationAnyOfConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


