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

// AccountV2StatusEnumSuspended struct for AccountV2StatusEnumSuspended
type AccountV2StatusEnumSuspended struct {
	SUSPENDED string `json:"SUSPENDED"`
}

// NewAccountV2StatusEnumSuspended instantiates a new AccountV2StatusEnumSuspended object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountV2StatusEnumSuspended(sUSPENDED string) *AccountV2StatusEnumSuspended {
	this := AccountV2StatusEnumSuspended{}
	this.SUSPENDED = sUSPENDED
	return &this
}

// NewAccountV2StatusEnumSuspendedWithDefaults instantiates a new AccountV2StatusEnumSuspended object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountV2StatusEnumSuspendedWithDefaults() *AccountV2StatusEnumSuspended {
	this := AccountV2StatusEnumSuspended{}
	return &this
}

// GetSUSPENDED returns the SUSPENDED field value
func (o *AccountV2StatusEnumSuspended) GetSUSPENDED() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUSPENDED
}

// GetSUSPENDEDOk returns a tuple with the SUSPENDED field value
// and a boolean to check if the value has been set.
func (o *AccountV2StatusEnumSuspended) GetSUSPENDEDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SUSPENDED, true
}

// SetSUSPENDED sets field value
func (o *AccountV2StatusEnumSuspended) SetSUSPENDED(v string) {
	o.SUSPENDED = v
}

func (o AccountV2StatusEnumSuspended) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["SUSPENDED"] = o.SUSPENDED
	}
	return json.Marshal(toSerialize)
}

type NullableAccountV2StatusEnumSuspended struct {
	value *AccountV2StatusEnumSuspended
	isSet bool
}

func (v NullableAccountV2StatusEnumSuspended) Get() *AccountV2StatusEnumSuspended {
	return v.value
}

func (v *NullableAccountV2StatusEnumSuspended) Set(val *AccountV2StatusEnumSuspended) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountV2StatusEnumSuspended) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountV2StatusEnumSuspended) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountV2StatusEnumSuspended(val *AccountV2StatusEnumSuspended) *NullableAccountV2StatusEnumSuspended {
	return &NullableAccountV2StatusEnumSuspended{value: val, isSet: true}
}

func (v NullableAccountV2StatusEnumSuspended) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountV2StatusEnumSuspended) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


