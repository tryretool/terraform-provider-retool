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

// checks if the RestAPIOptionsAuthenticationOptionsAnyOf4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestAPIOptionsAuthenticationOptionsAnyOf4{}

// RestAPIOptionsAuthenticationOptionsAnyOf4 struct for RestAPIOptionsAuthenticationOptionsAnyOf4
type RestAPIOptionsAuthenticationOptionsAnyOf4 struct {
	AuthenticationType string `json:"authentication_type"`
	Oauth1SignatureMethod RestAPIOptionsAuthenticationOptionsAnyOf4Oauth1SignatureMethod `json:"oauth1_signature_method"`
	Oauth1ConsumerKey string `json:"oauth1_consumer_key"`
	Oauth1ConsumerSecret string `json:"oauth1_consumer_secret"`
	Oauth1TokenKey string `json:"oauth1_token_key"`
	Oauth1TokenSecret string `json:"oauth1_token_secret"`
	Oauth1RealmParameter string `json:"oauth1_realm_parameter"`
}

type _RestAPIOptionsAuthenticationOptionsAnyOf4 RestAPIOptionsAuthenticationOptionsAnyOf4

// NewRestAPIOptionsAuthenticationOptionsAnyOf4 instantiates a new RestAPIOptionsAuthenticationOptionsAnyOf4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestAPIOptionsAuthenticationOptionsAnyOf4(authenticationType string, oauth1SignatureMethod RestAPIOptionsAuthenticationOptionsAnyOf4Oauth1SignatureMethod, oauth1ConsumerKey string, oauth1ConsumerSecret string, oauth1TokenKey string, oauth1TokenSecret string, oauth1RealmParameter string) *RestAPIOptionsAuthenticationOptionsAnyOf4 {
	this := RestAPIOptionsAuthenticationOptionsAnyOf4{}
	this.AuthenticationType = authenticationType
	this.Oauth1SignatureMethod = oauth1SignatureMethod
	this.Oauth1ConsumerKey = oauth1ConsumerKey
	this.Oauth1ConsumerSecret = oauth1ConsumerSecret
	this.Oauth1TokenKey = oauth1TokenKey
	this.Oauth1TokenSecret = oauth1TokenSecret
	this.Oauth1RealmParameter = oauth1RealmParameter
	return &this
}

// NewRestAPIOptionsAuthenticationOptionsAnyOf4WithDefaults instantiates a new RestAPIOptionsAuthenticationOptionsAnyOf4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestAPIOptionsAuthenticationOptionsAnyOf4WithDefaults() *RestAPIOptionsAuthenticationOptionsAnyOf4 {
	this := RestAPIOptionsAuthenticationOptionsAnyOf4{}
	return &this
}

// GetAuthenticationType returns the AuthenticationType field value
func (o *RestAPIOptionsAuthenticationOptionsAnyOf4) GetAuthenticationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthenticationType
}

// GetAuthenticationTypeOk returns a tuple with the AuthenticationType field value
// and a boolean to check if the value has been set.
func (o *RestAPIOptionsAuthenticationOptionsAnyOf4) GetAuthenticationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationType, true
}

// SetAuthenticationType sets field value
func (o *RestAPIOptionsAuthenticationOptionsAnyOf4) SetAuthenticationType(v string) {
	o.AuthenticationType = v
}

// GetOauth1SignatureMethod returns the Oauth1SignatureMethod field value
func (o *RestAPIOptionsAuthenticationOptionsAnyOf4) GetOauth1SignatureMethod() RestAPIOptionsAuthenticationOptionsAnyOf4Oauth1SignatureMethod {
	if o == nil {
		var ret RestAPIOptionsAuthenticationOptionsAnyOf4Oauth1SignatureMethod
		return ret
	}

	return o.Oauth1SignatureMethod
}

// GetOauth1SignatureMethodOk returns a tuple with the Oauth1SignatureMethod field value
// and a boolean to check if the value has been set.
func (o *RestAPIOptionsAuthenticationOptionsAnyOf4) GetOauth1SignatureMethodOk() (*RestAPIOptionsAuthenticationOptionsAnyOf4Oauth1SignatureMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Oauth1SignatureMethod, true
}

// SetOauth1SignatureMethod sets field value
func (o *RestAPIOptionsAuthenticationOptionsAnyOf4) SetOauth1SignatureMethod(v RestAPIOptionsAuthenticationOptionsAnyOf4Oauth1SignatureMethod) {
	o.Oauth1SignatureMethod = v
}

// GetOauth1ConsumerKey returns the Oauth1ConsumerKey field value
func (o *RestAPIOptionsAuthenticationOptionsAnyOf4) GetOauth1ConsumerKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Oauth1ConsumerKey
}

// GetOauth1ConsumerKeyOk returns a tuple with the Oauth1ConsumerKey field value
// and a boolean to check if the value has been set.
func (o *RestAPIOptionsAuthenticationOptionsAnyOf4) GetOauth1ConsumerKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Oauth1ConsumerKey, true
}

// SetOauth1ConsumerKey sets field value
func (o *RestAPIOptionsAuthenticationOptionsAnyOf4) SetOauth1ConsumerKey(v string) {
	o.Oauth1ConsumerKey = v
}

// GetOauth1ConsumerSecret returns the Oauth1ConsumerSecret field value
func (o *RestAPIOptionsAuthenticationOptionsAnyOf4) GetOauth1ConsumerSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Oauth1ConsumerSecret
}

// GetOauth1ConsumerSecretOk returns a tuple with the Oauth1ConsumerSecret field value
// and a boolean to check if the value has been set.
func (o *RestAPIOptionsAuthenticationOptionsAnyOf4) GetOauth1ConsumerSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Oauth1ConsumerSecret, true
}

// SetOauth1ConsumerSecret sets field value
func (o *RestAPIOptionsAuthenticationOptionsAnyOf4) SetOauth1ConsumerSecret(v string) {
	o.Oauth1ConsumerSecret = v
}

// GetOauth1TokenKey returns the Oauth1TokenKey field value
func (o *RestAPIOptionsAuthenticationOptionsAnyOf4) GetOauth1TokenKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Oauth1TokenKey
}

// GetOauth1TokenKeyOk returns a tuple with the Oauth1TokenKey field value
// and a boolean to check if the value has been set.
func (o *RestAPIOptionsAuthenticationOptionsAnyOf4) GetOauth1TokenKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Oauth1TokenKey, true
}

// SetOauth1TokenKey sets field value
func (o *RestAPIOptionsAuthenticationOptionsAnyOf4) SetOauth1TokenKey(v string) {
	o.Oauth1TokenKey = v
}

// GetOauth1TokenSecret returns the Oauth1TokenSecret field value
func (o *RestAPIOptionsAuthenticationOptionsAnyOf4) GetOauth1TokenSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Oauth1TokenSecret
}

// GetOauth1TokenSecretOk returns a tuple with the Oauth1TokenSecret field value
// and a boolean to check if the value has been set.
func (o *RestAPIOptionsAuthenticationOptionsAnyOf4) GetOauth1TokenSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Oauth1TokenSecret, true
}

// SetOauth1TokenSecret sets field value
func (o *RestAPIOptionsAuthenticationOptionsAnyOf4) SetOauth1TokenSecret(v string) {
	o.Oauth1TokenSecret = v
}

// GetOauth1RealmParameter returns the Oauth1RealmParameter field value
func (o *RestAPIOptionsAuthenticationOptionsAnyOf4) GetOauth1RealmParameter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Oauth1RealmParameter
}

// GetOauth1RealmParameterOk returns a tuple with the Oauth1RealmParameter field value
// and a boolean to check if the value has been set.
func (o *RestAPIOptionsAuthenticationOptionsAnyOf4) GetOauth1RealmParameterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Oauth1RealmParameter, true
}

// SetOauth1RealmParameter sets field value
func (o *RestAPIOptionsAuthenticationOptionsAnyOf4) SetOauth1RealmParameter(v string) {
	o.Oauth1RealmParameter = v
}

func (o RestAPIOptionsAuthenticationOptionsAnyOf4) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestAPIOptionsAuthenticationOptionsAnyOf4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authentication_type"] = o.AuthenticationType
	toSerialize["oauth1_signature_method"] = o.Oauth1SignatureMethod
	toSerialize["oauth1_consumer_key"] = o.Oauth1ConsumerKey
	toSerialize["oauth1_consumer_secret"] = o.Oauth1ConsumerSecret
	toSerialize["oauth1_token_key"] = o.Oauth1TokenKey
	toSerialize["oauth1_token_secret"] = o.Oauth1TokenSecret
	toSerialize["oauth1_realm_parameter"] = o.Oauth1RealmParameter
	return toSerialize, nil
}

func (o *RestAPIOptionsAuthenticationOptionsAnyOf4) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authentication_type",
		"oauth1_signature_method",
		"oauth1_consumer_key",
		"oauth1_consumer_secret",
		"oauth1_token_key",
		"oauth1_token_secret",
		"oauth1_realm_parameter",
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

	varRestAPIOptionsAuthenticationOptionsAnyOf4 := _RestAPIOptionsAuthenticationOptionsAnyOf4{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRestAPIOptionsAuthenticationOptionsAnyOf4)

	if err != nil {
		return err
	}

	*o = RestAPIOptionsAuthenticationOptionsAnyOf4(varRestAPIOptionsAuthenticationOptionsAnyOf4)

	return err
}

type NullableRestAPIOptionsAuthenticationOptionsAnyOf4 struct {
	value *RestAPIOptionsAuthenticationOptionsAnyOf4
	isSet bool
}

func (v NullableRestAPIOptionsAuthenticationOptionsAnyOf4) Get() *RestAPIOptionsAuthenticationOptionsAnyOf4 {
	return v.value
}

func (v *NullableRestAPIOptionsAuthenticationOptionsAnyOf4) Set(val *RestAPIOptionsAuthenticationOptionsAnyOf4) {
	v.value = val
	v.isSet = true
}

func (v NullableRestAPIOptionsAuthenticationOptionsAnyOf4) IsSet() bool {
	return v.isSet
}

func (v *NullableRestAPIOptionsAuthenticationOptionsAnyOf4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestAPIOptionsAuthenticationOptionsAnyOf4(val *RestAPIOptionsAuthenticationOptionsAnyOf4) *NullableRestAPIOptionsAuthenticationOptionsAnyOf4 {
	return &NullableRestAPIOptionsAuthenticationOptionsAnyOf4{value: val, isSet: true}
}

func (v NullableRestAPIOptionsAuthenticationOptionsAnyOf4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestAPIOptionsAuthenticationOptionsAnyOf4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


