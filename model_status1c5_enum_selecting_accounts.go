/*
Nordigen Account Information Services API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0 (v2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nordigen

import (
	"encoding/json"
)

// Status1c5EnumSelectingAccounts struct for Status1c5EnumSelectingAccounts
type Status1c5EnumSelectingAccounts struct {
	Short string `json:"short"`
	Long string `json:"long"`
	Description string `json:"description"`
}

// NewStatus1c5EnumSelectingAccounts instantiates a new Status1c5EnumSelectingAccounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatus1c5EnumSelectingAccounts(short string, long string, description string) *Status1c5EnumSelectingAccounts {
	this := Status1c5EnumSelectingAccounts{}
	this.Short = short
	this.Long = long
	this.Description = description
	return &this
}

// NewStatus1c5EnumSelectingAccountsWithDefaults instantiates a new Status1c5EnumSelectingAccounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatus1c5EnumSelectingAccountsWithDefaults() *Status1c5EnumSelectingAccounts {
	this := Status1c5EnumSelectingAccounts{}
	return &this
}

// GetShort returns the Short field value
func (o *Status1c5EnumSelectingAccounts) GetShort() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Short
}

// GetShortOk returns a tuple with the Short field value
// and a boolean to check if the value has been set.
func (o *Status1c5EnumSelectingAccounts) GetShortOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Short, true
}

// SetShort sets field value
func (o *Status1c5EnumSelectingAccounts) SetShort(v string) {
	o.Short = v
}

// GetLong returns the Long field value
func (o *Status1c5EnumSelectingAccounts) GetLong() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Long
}

// GetLongOk returns a tuple with the Long field value
// and a boolean to check if the value has been set.
func (o *Status1c5EnumSelectingAccounts) GetLongOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Long, true
}

// SetLong sets field value
func (o *Status1c5EnumSelectingAccounts) SetLong(v string) {
	o.Long = v
}

// GetDescription returns the Description field value
func (o *Status1c5EnumSelectingAccounts) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Status1c5EnumSelectingAccounts) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Status1c5EnumSelectingAccounts) SetDescription(v string) {
	o.Description = v
}

func (o Status1c5EnumSelectingAccounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["short"] = o.Short
	}
	if true {
		toSerialize["long"] = o.Long
	}
	if true {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableStatus1c5EnumSelectingAccounts struct {
	value *Status1c5EnumSelectingAccounts
	isSet bool
}

func (v NullableStatus1c5EnumSelectingAccounts) Get() *Status1c5EnumSelectingAccounts {
	return v.value
}

func (v *NullableStatus1c5EnumSelectingAccounts) Set(val *Status1c5EnumSelectingAccounts) {
	v.value = val
	v.isSet = true
}

func (v NullableStatus1c5EnumSelectingAccounts) IsSet() bool {
	return v.isSet
}

func (v *NullableStatus1c5EnumSelectingAccounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatus1c5EnumSelectingAccounts(val *Status1c5EnumSelectingAccounts) *NullableStatus1c5EnumSelectingAccounts {
	return &NullableStatus1c5EnumSelectingAccounts{value: val, isSet: true}
}

func (v NullableStatus1c5EnumSelectingAccounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatus1c5EnumSelectingAccounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


