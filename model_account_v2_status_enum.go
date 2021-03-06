/*
Nordigen Account Information Services API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0 (v2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nordigen

import (
	"encoding/json"
	"fmt"
)

// AccountV2StatusEnum Status of account
type AccountV2StatusEnum string

// List of AccountV2StatusEnum
const (
	ACCOUNTV2STATUSENUM_READY AccountV2StatusEnum = "READY"
	ACCOUNTV2STATUSENUM_EXPIRED AccountV2StatusEnum = "EXPIRED"
	ACCOUNTV2STATUSENUM_DISCOVERED AccountV2StatusEnum = "DISCOVERED"
	ACCOUNTV2STATUSENUM_PROCESSING AccountV2StatusEnum = "PROCESSING"
	ACCOUNTV2STATUSENUM_ERROR AccountV2StatusEnum = "ERROR"
	ACCOUNTV2STATUSENUM_SUSPENDED AccountV2StatusEnum = "SUSPENDED"
)

// All allowed values of AccountV2StatusEnum enum
var AllowedAccountV2StatusEnumEnumValues = []AccountV2StatusEnum{
	"READY",
	"EXPIRED",
	"DISCOVERED",
	"PROCESSING",
	"ERROR",
	"SUSPENDED",
}

func (v *AccountV2StatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccountV2StatusEnum(value)
	for _, existing := range AllowedAccountV2StatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccountV2StatusEnum", value)
}

// NewAccountV2StatusEnumFromValue returns a pointer to a valid AccountV2StatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountV2StatusEnumFromValue(v string) (*AccountV2StatusEnum, error) {
	ev := AccountV2StatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccountV2StatusEnum: valid values are %v", v, AllowedAccountV2StatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountV2StatusEnum) IsValid() bool {
	for _, existing := range AllowedAccountV2StatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccountV2StatusEnum value
func (v AccountV2StatusEnum) Ptr() *AccountV2StatusEnum {
	return &v
}

type NullableAccountV2StatusEnum struct {
	value *AccountV2StatusEnum
	isSet bool
}

func (v NullableAccountV2StatusEnum) Get() *AccountV2StatusEnum {
	return v.value
}

func (v *NullableAccountV2StatusEnum) Set(val *AccountV2StatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountV2StatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountV2StatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountV2StatusEnum(val *AccountV2StatusEnum) *NullableAccountV2StatusEnum {
	return &NullableAccountV2StatusEnum{value: val, isSet: true}
}

func (v NullableAccountV2StatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountV2StatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

