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

// AccountV2StatusEnumProcessing struct for AccountV2StatusEnumProcessing
type AccountV2StatusEnumProcessing struct {
	PROCESSING string `json:"PROCESSING"`
}

// NewAccountV2StatusEnumProcessing instantiates a new AccountV2StatusEnumProcessing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountV2StatusEnumProcessing(pROCESSING string) *AccountV2StatusEnumProcessing {
	this := AccountV2StatusEnumProcessing{}
	this.PROCESSING = pROCESSING
	return &this
}

// NewAccountV2StatusEnumProcessingWithDefaults instantiates a new AccountV2StatusEnumProcessing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountV2StatusEnumProcessingWithDefaults() *AccountV2StatusEnumProcessing {
	this := AccountV2StatusEnumProcessing{}
	return &this
}

// GetPROCESSING returns the PROCESSING field value
func (o *AccountV2StatusEnumProcessing) GetPROCESSING() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PROCESSING
}

// GetPROCESSINGOk returns a tuple with the PROCESSING field value
// and a boolean to check if the value has been set.
func (o *AccountV2StatusEnumProcessing) GetPROCESSINGOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PROCESSING, true
}

// SetPROCESSING sets field value
func (o *AccountV2StatusEnumProcessing) SetPROCESSING(v string) {
	o.PROCESSING = v
}

func (o AccountV2StatusEnumProcessing) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["PROCESSING"] = o.PROCESSING
	}
	return json.Marshal(toSerialize)
}

type NullableAccountV2StatusEnumProcessing struct {
	value *AccountV2StatusEnumProcessing
	isSet bool
}

func (v NullableAccountV2StatusEnumProcessing) Get() *AccountV2StatusEnumProcessing {
	return v.value
}

func (v *NullableAccountV2StatusEnumProcessing) Set(val *AccountV2StatusEnumProcessing) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountV2StatusEnumProcessing) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountV2StatusEnumProcessing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountV2StatusEnumProcessing(val *AccountV2StatusEnumProcessing) *NullableAccountV2StatusEnumProcessing {
	return &NullableAccountV2StatusEnumProcessing{value: val, isSet: true}
}

func (v NullableAccountV2StatusEnumProcessing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountV2StatusEnumProcessing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


