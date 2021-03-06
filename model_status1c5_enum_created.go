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

// Status1c5EnumCreated the model 'Status1c5EnumCreated'
type Status1c5EnumCreated string

// List of Status1c5EnumCreated
const (
	STATUS1C5ENUMCREATED_CR Status1c5EnumCreated = "CR"
	STATUS1C5ENUMCREATED_CREATED Status1c5EnumCreated = "CREATED"
)

// All allowed values of Status1c5EnumCreated enum
var AllowedStatus1c5EnumCreatedEnumValues = []Status1c5EnumCreated{
	"CR",
	"CREATED",
}

func (v *Status1c5EnumCreated) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Status1c5EnumCreated(value)
	for _, existing := range AllowedStatus1c5EnumCreatedEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Status1c5EnumCreated", value)
}

// NewStatus1c5EnumCreatedFromValue returns a pointer to a valid Status1c5EnumCreated
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStatus1c5EnumCreatedFromValue(v string) (*Status1c5EnumCreated, error) {
	ev := Status1c5EnumCreated(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Status1c5EnumCreated: valid values are %v", v, AllowedStatus1c5EnumCreatedEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Status1c5EnumCreated) IsValid() bool {
	for _, existing := range AllowedStatus1c5EnumCreatedEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Status1c5EnumCreated value
func (v Status1c5EnumCreated) Ptr() *Status1c5EnumCreated {
	return &v
}

type NullableStatus1c5EnumCreated struct {
	value *Status1c5EnumCreated
	isSet bool
}

func (v NullableStatus1c5EnumCreated) Get() *Status1c5EnumCreated {
	return v.value
}

func (v *NullableStatus1c5EnumCreated) Set(val *Status1c5EnumCreated) {
	v.value = val
	v.isSet = true
}

func (v NullableStatus1c5EnumCreated) IsSet() bool {
	return v.isSet
}

func (v *NullableStatus1c5EnumCreated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatus1c5EnumCreated(val *Status1c5EnumCreated) *NullableStatus1c5EnumCreated {
	return &NullableStatus1c5EnumCreated{value: val, isSet: true}
}

func (v NullableStatus1c5EnumCreated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatus1c5EnumCreated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

