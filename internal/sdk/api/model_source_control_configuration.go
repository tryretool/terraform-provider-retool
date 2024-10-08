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

// SourceControlConfiguration This object represents the Source Control provider configuration for the organization or space. See https://docs.retool.com/source-control for more information.
type SourceControlConfiguration struct {
	SourceControlConfigGet200ResponseDataOneOf1 *SourceControlConfigGet200ResponseDataOneOf1
	SourceControlConfigGet200ResponseDataOneOf2 *SourceControlConfigGet200ResponseDataOneOf2
	SourceControlConfigGet200ResponseDataOneOf3 *SourceControlConfigGet200ResponseDataOneOf3
	SourceControlConfigGet200ResponseDataOneOf4 *SourceControlConfigGet200ResponseDataOneOf4
	SourceControlConfigurationAnyOf *SourceControlConfigurationAnyOf
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SourceControlConfiguration) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SourceControlConfigGet200ResponseDataOneOf1
	err = json.Unmarshal(data, &dst.SourceControlConfigGet200ResponseDataOneOf1);
	if err == nil {
		jsonSourceControlConfigGet200ResponseDataOneOf1, _ := json.Marshal(dst.SourceControlConfigGet200ResponseDataOneOf1)
		if string(jsonSourceControlConfigGet200ResponseDataOneOf1) == "{}" { // empty struct
			dst.SourceControlConfigGet200ResponseDataOneOf1 = nil
		} else {
			return nil // data stored in dst.SourceControlConfigGet200ResponseDataOneOf1, return on the first match
		}
	} else {
		dst.SourceControlConfigGet200ResponseDataOneOf1 = nil
	}

	// try to unmarshal JSON data into SourceControlConfigGet200ResponseDataOneOf2
	err = json.Unmarshal(data, &dst.SourceControlConfigGet200ResponseDataOneOf2);
	if err == nil {
		jsonSourceControlConfigGet200ResponseDataOneOf2, _ := json.Marshal(dst.SourceControlConfigGet200ResponseDataOneOf2)
		if string(jsonSourceControlConfigGet200ResponseDataOneOf2) == "{}" { // empty struct
			dst.SourceControlConfigGet200ResponseDataOneOf2 = nil
		} else {
			return nil // data stored in dst.SourceControlConfigGet200ResponseDataOneOf2, return on the first match
		}
	} else {
		dst.SourceControlConfigGet200ResponseDataOneOf2 = nil
	}

	// try to unmarshal JSON data into SourceControlConfigGet200ResponseDataOneOf3
	err = json.Unmarshal(data, &dst.SourceControlConfigGet200ResponseDataOneOf3);
	if err == nil {
		jsonSourceControlConfigGet200ResponseDataOneOf3, _ := json.Marshal(dst.SourceControlConfigGet200ResponseDataOneOf3)
		if string(jsonSourceControlConfigGet200ResponseDataOneOf3) == "{}" { // empty struct
			dst.SourceControlConfigGet200ResponseDataOneOf3 = nil
		} else {
			return nil // data stored in dst.SourceControlConfigGet200ResponseDataOneOf3, return on the first match
		}
	} else {
		dst.SourceControlConfigGet200ResponseDataOneOf3 = nil
	}

	// try to unmarshal JSON data into SourceControlConfigGet200ResponseDataOneOf4
	err = json.Unmarshal(data, &dst.SourceControlConfigGet200ResponseDataOneOf4);
	if err == nil {
		jsonSourceControlConfigGet200ResponseDataOneOf4, _ := json.Marshal(dst.SourceControlConfigGet200ResponseDataOneOf4)
		if string(jsonSourceControlConfigGet200ResponseDataOneOf4) == "{}" { // empty struct
			dst.SourceControlConfigGet200ResponseDataOneOf4 = nil
		} else {
			return nil // data stored in dst.SourceControlConfigGet200ResponseDataOneOf4, return on the first match
		}
	} else {
		dst.SourceControlConfigGet200ResponseDataOneOf4 = nil
	}

	// try to unmarshal JSON data into SourceControlConfigurationAnyOf
	err = json.Unmarshal(data, &dst.SourceControlConfigurationAnyOf);
	if err == nil {
		jsonSourceControlConfigurationAnyOf, _ := json.Marshal(dst.SourceControlConfigurationAnyOf)
		if string(jsonSourceControlConfigurationAnyOf) == "{}" { // empty struct
			dst.SourceControlConfigurationAnyOf = nil
		} else {
			return nil // data stored in dst.SourceControlConfigurationAnyOf, return on the first match
		}
	} else {
		dst.SourceControlConfigurationAnyOf = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SourceControlConfiguration)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SourceControlConfiguration) MarshalJSON() ([]byte, error) {
	if src.SourceControlConfigGet200ResponseDataOneOf1 != nil {
		return json.Marshal(&src.SourceControlConfigGet200ResponseDataOneOf1)
	}

	if src.SourceControlConfigGet200ResponseDataOneOf2 != nil {
		return json.Marshal(&src.SourceControlConfigGet200ResponseDataOneOf2)
	}

	if src.SourceControlConfigGet200ResponseDataOneOf3 != nil {
		return json.Marshal(&src.SourceControlConfigGet200ResponseDataOneOf3)
	}

	if src.SourceControlConfigGet200ResponseDataOneOf4 != nil {
		return json.Marshal(&src.SourceControlConfigGet200ResponseDataOneOf4)
	}

	if src.SourceControlConfigurationAnyOf != nil {
		return json.Marshal(&src.SourceControlConfigurationAnyOf)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSourceControlConfiguration struct {
	value *SourceControlConfiguration
	isSet bool
}

func (v NullableSourceControlConfiguration) Get() *SourceControlConfiguration {
	return v.value
}

func (v *NullableSourceControlConfiguration) Set(val *SourceControlConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceControlConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceControlConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceControlConfiguration(val *SourceControlConfiguration) *NullableSourceControlConfiguration {
	return &NullableSourceControlConfiguration{value: val, isSet: true}
}

func (v NullableSourceControlConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceControlConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


