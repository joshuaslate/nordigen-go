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

// AccountV2StatusEnumExpired struct for AccountV2StatusEnumExpired
type AccountV2StatusEnumExpired struct {
	EXPIRED string `json:"EXPIRED"`
}

// NewAccountV2StatusEnumExpired instantiates a new AccountV2StatusEnumExpired object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountV2StatusEnumExpired(eXPIRED string) *AccountV2StatusEnumExpired {
	this := AccountV2StatusEnumExpired{}
	this.EXPIRED = eXPIRED
	return &this
}

// NewAccountV2StatusEnumExpiredWithDefaults instantiates a new AccountV2StatusEnumExpired object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountV2StatusEnumExpiredWithDefaults() *AccountV2StatusEnumExpired {
	this := AccountV2StatusEnumExpired{}
	return &this
}

// GetEXPIRED returns the EXPIRED field value
func (o *AccountV2StatusEnumExpired) GetEXPIRED() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EXPIRED
}

// GetEXPIREDOk returns a tuple with the EXPIRED field value
// and a boolean to check if the value has been set.
func (o *AccountV2StatusEnumExpired) GetEXPIREDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EXPIRED, true
}

// SetEXPIRED sets field value
func (o *AccountV2StatusEnumExpired) SetEXPIRED(v string) {
	o.EXPIRED = v
}

func (o AccountV2StatusEnumExpired) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["EXPIRED"] = o.EXPIRED
	}
	return json.Marshal(toSerialize)
}

type NullableAccountV2StatusEnumExpired struct {
	value *AccountV2StatusEnumExpired
	isSet bool
}

func (v NullableAccountV2StatusEnumExpired) Get() *AccountV2StatusEnumExpired {
	return v.value
}

func (v *NullableAccountV2StatusEnumExpired) Set(val *AccountV2StatusEnumExpired) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountV2StatusEnumExpired) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountV2StatusEnumExpired) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountV2StatusEnumExpired(val *AccountV2StatusEnumExpired) *NullableAccountV2StatusEnumExpired {
	return &NullableAccountV2StatusEnumExpired{value: val, isSet: true}
}

func (v NullableAccountV2StatusEnumExpired) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountV2StatusEnumExpired) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


