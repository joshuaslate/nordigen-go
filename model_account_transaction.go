/*
Nordigen Account Information Services API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0 (v2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nordigen

import (
	"encoding/json"
	"time"
)

// AccountTransaction struct for AccountTransaction
type AccountTransaction struct {
	TransactionId *string `json:"transactionId,omitempty"`
	DebtorName *string `json:"debtorName,omitempty"`
	DebtorAccount *AccountTransactionDebtorAccount `json:"debtorAccount,omitempty"`
	TransactionAmount *AccountTransactionTransactionAmount `json:"transactionAmount,omitempty"`
	BankTransactionCode *string `json:"bankTransactionCode,omitempty"`
	BookingDate *time.Time `json:"bookingDate,omitempty"`
	ValueDate *time.Time `json:"valueDate,omitempty"`
	RemittanceInformationUnstructured *string `json:"remittanceInformationUnstructured,omitempty"`
}

// NewAccountTransaction instantiates a new AccountTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountTransaction() *AccountTransaction {
	this := AccountTransaction{}
	return &this
}

// NewAccountTransactionWithDefaults instantiates a new AccountTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountTransactionWithDefaults() *AccountTransaction {
	this := AccountTransaction{}
	return &this
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *AccountTransaction) GetTransactionId() string {
	if o == nil || o.TransactionId == nil {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTransaction) GetTransactionIdOk() (*string, bool) {
	if o == nil || o.TransactionId == nil {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *AccountTransaction) HasTransactionId() bool {
	if o != nil && o.TransactionId != nil {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *AccountTransaction) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetDebtorName returns the DebtorName field value if set, zero value otherwise.
func (o *AccountTransaction) GetDebtorName() string {
	if o == nil || o.DebtorName == nil {
		var ret string
		return ret
	}
	return *o.DebtorName
}

// GetDebtorNameOk returns a tuple with the DebtorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTransaction) GetDebtorNameOk() (*string, bool) {
	if o == nil || o.DebtorName == nil {
		return nil, false
	}
	return o.DebtorName, true
}

// HasDebtorName returns a boolean if a field has been set.
func (o *AccountTransaction) HasDebtorName() bool {
	if o != nil && o.DebtorName != nil {
		return true
	}

	return false
}

// SetDebtorName gets a reference to the given string and assigns it to the DebtorName field.
func (o *AccountTransaction) SetDebtorName(v string) {
	o.DebtorName = &v
}

// GetDebtorAccount returns the DebtorAccount field value if set, zero value otherwise.
func (o *AccountTransaction) GetDebtorAccount() AccountTransactionDebtorAccount {
	if o == nil || o.DebtorAccount == nil {
		var ret AccountTransactionDebtorAccount
		return ret
	}
	return *o.DebtorAccount
}

// GetDebtorAccountOk returns a tuple with the DebtorAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTransaction) GetDebtorAccountOk() (*AccountTransactionDebtorAccount, bool) {
	if o == nil || o.DebtorAccount == nil {
		return nil, false
	}
	return o.DebtorAccount, true
}

// HasDebtorAccount returns a boolean if a field has been set.
func (o *AccountTransaction) HasDebtorAccount() bool {
	if o != nil && o.DebtorAccount != nil {
		return true
	}

	return false
}

// SetDebtorAccount gets a reference to the given AccountTransactionDebtorAccount and assigns it to the DebtorAccount field.
func (o *AccountTransaction) SetDebtorAccount(v AccountTransactionDebtorAccount) {
	o.DebtorAccount = &v
}

// GetTransactionAmount returns the TransactionAmount field value if set, zero value otherwise.
func (o *AccountTransaction) GetTransactionAmount() AccountTransactionTransactionAmount {
	if o == nil || o.TransactionAmount == nil {
		var ret AccountTransactionTransactionAmount
		return ret
	}
	return *o.TransactionAmount
}

// GetTransactionAmountOk returns a tuple with the TransactionAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTransaction) GetTransactionAmountOk() (*AccountTransactionTransactionAmount, bool) {
	if o == nil || o.TransactionAmount == nil {
		return nil, false
	}
	return o.TransactionAmount, true
}

// HasTransactionAmount returns a boolean if a field has been set.
func (o *AccountTransaction) HasTransactionAmount() bool {
	if o != nil && o.TransactionAmount != nil {
		return true
	}

	return false
}

// SetTransactionAmount gets a reference to the given AccountTransactionTransactionAmount and assigns it to the TransactionAmount field.
func (o *AccountTransaction) SetTransactionAmount(v AccountTransactionTransactionAmount) {
	o.TransactionAmount = &v
}

// GetBankTransactionCode returns the BankTransactionCode field value if set, zero value otherwise.
func (o *AccountTransaction) GetBankTransactionCode() string {
	if o == nil || o.BankTransactionCode == nil {
		var ret string
		return ret
	}
	return *o.BankTransactionCode
}

// GetBankTransactionCodeOk returns a tuple with the BankTransactionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTransaction) GetBankTransactionCodeOk() (*string, bool) {
	if o == nil || o.BankTransactionCode == nil {
		return nil, false
	}
	return o.BankTransactionCode, true
}

// HasBankTransactionCode returns a boolean if a field has been set.
func (o *AccountTransaction) HasBankTransactionCode() bool {
	if o != nil && o.BankTransactionCode != nil {
		return true
	}

	return false
}

// SetBankTransactionCode gets a reference to the given string and assigns it to the BankTransactionCode field.
func (o *AccountTransaction) SetBankTransactionCode(v string) {
	o.BankTransactionCode = &v
}

// GetBookingDate returns the BookingDate field value if set, zero value otherwise.
func (o *AccountTransaction) GetBookingDate() time.Time {
	if o == nil || o.BookingDate == nil {
		var ret time.Time
		return ret
	}
	return *o.BookingDate
}

// GetBookingDateOk returns a tuple with the BookingDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTransaction) GetBookingDateOk() (*time.Time, bool) {
	if o == nil || o.BookingDate == nil {
		return nil, false
	}
	return o.BookingDate, true
}

// HasBookingDate returns a boolean if a field has been set.
func (o *AccountTransaction) HasBookingDate() bool {
	if o != nil && o.BookingDate != nil {
		return true
	}

	return false
}

// SetBookingDate gets a reference to the given time.Time and assigns it to the BookingDate field.
func (o *AccountTransaction) SetBookingDate(v time.Time) {
	o.BookingDate = &v
}

// GetValueDate returns the ValueDate field value if set, zero value otherwise.
func (o *AccountTransaction) GetValueDate() time.Time {
	if o == nil || o.ValueDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ValueDate
}

// GetValueDateOk returns a tuple with the ValueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTransaction) GetValueDateOk() (*time.Time, bool) {
	if o == nil || o.ValueDate == nil {
		return nil, false
	}
	return o.ValueDate, true
}

// HasValueDate returns a boolean if a field has been set.
func (o *AccountTransaction) HasValueDate() bool {
	if o != nil && o.ValueDate != nil {
		return true
	}

	return false
}

// SetValueDate gets a reference to the given time.Time and assigns it to the ValueDate field.
func (o *AccountTransaction) SetValueDate(v time.Time) {
	o.ValueDate = &v
}

// GetRemittanceInformationUnstructured returns the RemittanceInformationUnstructured field value if set, zero value otherwise.
func (o *AccountTransaction) GetRemittanceInformationUnstructured() string {
	if o == nil || o.RemittanceInformationUnstructured == nil {
		var ret string
		return ret
	}
	return *o.RemittanceInformationUnstructured
}

// GetRemittanceInformationUnstructuredOk returns a tuple with the RemittanceInformationUnstructured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTransaction) GetRemittanceInformationUnstructuredOk() (*string, bool) {
	if o == nil || o.RemittanceInformationUnstructured == nil {
		return nil, false
	}
	return o.RemittanceInformationUnstructured, true
}

// HasRemittanceInformationUnstructured returns a boolean if a field has been set.
func (o *AccountTransaction) HasRemittanceInformationUnstructured() bool {
	if o != nil && o.RemittanceInformationUnstructured != nil {
		return true
	}

	return false
}

// SetRemittanceInformationUnstructured gets a reference to the given string and assigns it to the RemittanceInformationUnstructured field.
func (o *AccountTransaction) SetRemittanceInformationUnstructured(v string) {
	o.RemittanceInformationUnstructured = &v
}

func (o AccountTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TransactionId != nil {
		toSerialize["transactionId"] = o.TransactionId
	}
	if o.DebtorName != nil {
		toSerialize["debtorName"] = o.DebtorName
	}
	if o.DebtorAccount != nil {
		toSerialize["debtorAccount"] = o.DebtorAccount
	}
	if o.TransactionAmount != nil {
		toSerialize["transactionAmount"] = o.TransactionAmount
	}
	if o.BankTransactionCode != nil {
		toSerialize["bankTransactionCode"] = o.BankTransactionCode
	}
	if o.BookingDate != nil {
		toSerialize["bookingDate"] = o.BookingDate
	}
	if o.ValueDate != nil {
		toSerialize["valueDate"] = o.ValueDate
	}
	if o.RemittanceInformationUnstructured != nil {
		toSerialize["remittanceInformationUnstructured"] = o.RemittanceInformationUnstructured
	}
	return json.Marshal(toSerialize)
}

type NullableAccountTransaction struct {
	value *AccountTransaction
	isSet bool
}

func (v NullableAccountTransaction) Get() *AccountTransaction {
	return v.value
}

func (v *NullableAccountTransaction) Set(val *AccountTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountTransaction(val *AccountTransaction) *NullableAccountTransaction {
	return &NullableAccountTransaction{value: val, isSet: true}
}

func (v NullableAccountTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

