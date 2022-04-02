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

// AccountTransactionTransactionAmount struct for AccountTransactionTransactionAmount
type AccountTransactionTransactionAmount struct {
	Amount *string `json:"amount,omitempty"`
	Currency *string `json:"currency,omitempty"`
}

// NewAccountTransactionTransactionAmount instantiates a new AccountTransactionTransactionAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountTransactionTransactionAmount() *AccountTransactionTransactionAmount {
	this := AccountTransactionTransactionAmount{}
	return &this
}

// NewAccountTransactionTransactionAmountWithDefaults instantiates a new AccountTransactionTransactionAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountTransactionTransactionAmountWithDefaults() *AccountTransactionTransactionAmount {
	this := AccountTransactionTransactionAmount{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AccountTransactionTransactionAmount) GetAmount() string {
	if o == nil || o.Amount == nil {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTransactionTransactionAmount) GetAmountOk() (*string, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AccountTransactionTransactionAmount) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *AccountTransactionTransactionAmount) SetAmount(v string) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *AccountTransactionTransactionAmount) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTransactionTransactionAmount) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *AccountTransactionTransactionAmount) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *AccountTransactionTransactionAmount) SetCurrency(v string) {
	o.Currency = &v
}

func (o AccountTransactionTransactionAmount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	return json.Marshal(toSerialize)
}

type NullableAccountTransactionTransactionAmount struct {
	value *AccountTransactionTransactionAmount
	isSet bool
}

func (v NullableAccountTransactionTransactionAmount) Get() *AccountTransactionTransactionAmount {
	return v.value
}

func (v *NullableAccountTransactionTransactionAmount) Set(val *AccountTransactionTransactionAmount) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountTransactionTransactionAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountTransactionTransactionAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountTransactionTransactionAmount(val *AccountTransactionTransactionAmount) *NullableAccountTransactionTransactionAmount {
	return &NullableAccountTransactionTransactionAmount{value: val, isSet: true}
}

func (v NullableAccountTransactionTransactionAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountTransactionTransactionAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


