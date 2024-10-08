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

// checks if the PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions{}

// PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions struct for PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions
type PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions struct {
	Enabled bool `json:"enabled"`
	Host string `json:"host"`
	Port string `json:"port"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	PrivateKeyName *string `json:"private_key_name,omitempty"`
}

type _PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions

// NewPostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions instantiates a new PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions(enabled bool, host string, port string) *PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions {
	this := PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions{}
	this.Enabled = enabled
	this.Host = host
	this.Port = port
	return &this
}

// NewPostgresOptionsDatabaseOptionsAnyOfSshTunnelOptionsWithDefaults instantiates a new PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostgresOptionsDatabaseOptionsAnyOfSshTunnelOptionsWithDefaults() *PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions {
	this := PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) SetEnabled(v bool) {
	o.Enabled = v
}

// GetHost returns the Host field value
func (o *PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value
func (o *PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) GetPort() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) GetPortOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) SetPort(v string) {
	o.Port = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) SetPassword(v string) {
	o.Password = &v
}

// GetPrivateKeyName returns the PrivateKeyName field value if set, zero value otherwise.
func (o *PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) GetPrivateKeyName() string {
	if o == nil || IsNil(o.PrivateKeyName) {
		var ret string
		return ret
	}
	return *o.PrivateKeyName
}

// GetPrivateKeyNameOk returns a tuple with the PrivateKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) GetPrivateKeyNameOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateKeyName) {
		return nil, false
	}
	return o.PrivateKeyName, true
}

// HasPrivateKeyName returns a boolean if a field has been set.
func (o *PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) HasPrivateKeyName() bool {
	if o != nil && !IsNil(o.PrivateKeyName) {
		return true
	}

	return false
}

// SetPrivateKeyName gets a reference to the given string and assigns it to the PrivateKeyName field.
func (o *PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) SetPrivateKeyName(v string) {
	o.PrivateKeyName = &v
}

func (o PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	toSerialize["host"] = o.Host
	toSerialize["port"] = o.Port
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.PrivateKeyName) {
		toSerialize["private_key_name"] = o.PrivateKeyName
	}
	return toSerialize, nil
}

func (o *PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
		"host",
		"port",
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

	varPostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions := _PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions)

	if err != nil {
		return err
	}

	*o = PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions(varPostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions)

	return err
}

type NullablePostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions struct {
	value *PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions
	isSet bool
}

func (v NullablePostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) Get() *PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions {
	return v.value
}

func (v *NullablePostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) Set(val *PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) {
	v.value = val
	v.isSet = true
}

func (v NullablePostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) IsSet() bool {
	return v.isSet
}

func (v *NullablePostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions(val *PostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) *NullablePostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions {
	return &NullablePostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions{value: val, isSet: true}
}

func (v NullablePostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostgresOptionsDatabaseOptionsAnyOfSshTunnelOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


