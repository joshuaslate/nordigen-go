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

// AccountDetail struct for AccountDetail
type AccountDetail struct {
	ResourceId *string `json:"resourceId,omitempty"`
	Iban *string `json:"iban,omitempty"`
	Currency *string `json:"currency,omitempty"`
	OwnerName *string `json:"ownerName,omitempty"`
	Name *string `json:"name,omitempty"`
	Product *string `json:"product,omitempty"`
	CashAccountType *string `json:"cashAccountType,omitempty"`
}

// NewAccountDetail instantiates a new AccountDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountDetail() *AccountDetail {
	this := AccountDetail{}
	return &this
}

// NewAccountDetailWithDefaults instantiates a new AccountDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountDetailWithDefaults() *AccountDetail {
	this := AccountDetail{}
	return &this
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *AccountDetail) GetResourceId() string {
	if o == nil || o.ResourceId == nil {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetail) GetResourceIdOk() (*string, bool) {
	if o == nil || o.ResourceId == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *AccountDetail) HasResourceId() bool {
	if o != nil && o.ResourceId != nil {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *AccountDetail) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetIban returns the Iban field value if set, zero value otherwise.
func (o *AccountDetail) GetIban() string {
	if o == nil || o.Iban == nil {
		var ret string
		return ret
	}
	return *o.Iban
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetail) GetIbanOk() (*string, bool) {
	if o == nil || o.Iban == nil {
		return nil, false
	}
	return o.Iban, true
}

// HasIban returns a boolean if a field has been set.
func (o *AccountDetail) HasIban() bool {
	if o != nil && o.Iban != nil {
		return true
	}

	return false
}

// SetIban gets a reference to the given string and assigns it to the Iban field.
func (o *AccountDetail) SetIban(v string) {
	o.Iban = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *AccountDetail) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetail) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *AccountDetail) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *AccountDetail) SetCurrency(v string) {
	o.Currency = &v
}

// GetOwnerName returns the OwnerName field value if set, zero value otherwise.
func (o *AccountDetail) GetOwnerName() string {
	if o == nil || o.OwnerName == nil {
		var ret string
		return ret
	}
	return *o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetail) GetOwnerNameOk() (*string, bool) {
	if o == nil || o.OwnerName == nil {
		return nil, false
	}
	return o.OwnerName, true
}

// HasOwnerName returns a boolean if a field has been set.
func (o *AccountDetail) HasOwnerName() bool {
	if o != nil && o.OwnerName != nil {
		return true
	}

	return false
}

// SetOwnerName gets a reference to the given string and assigns it to the OwnerName field.
func (o *AccountDetail) SetOwnerName(v string) {
	o.OwnerName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccountDetail) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetail) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccountDetail) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccountDetail) SetName(v string) {
	o.Name = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *AccountDetail) GetProduct() string {
	if o == nil || o.Product == nil {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetail) GetProductOk() (*string, bool) {
	if o == nil || o.Product == nil {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *AccountDetail) HasProduct() bool {
	if o != nil && o.Product != nil {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *AccountDetail) SetProduct(v string) {
	o.Product = &v
}

// GetCashAccountType returns the CashAccountType field value if set, zero value otherwise.
func (o *AccountDetail) GetCashAccountType() string {
	if o == nil || o.CashAccountType == nil {
		var ret string
		return ret
	}
	return *o.CashAccountType
}

// GetCashAccountTypeOk returns a tuple with the CashAccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetail) GetCashAccountTypeOk() (*string, bool) {
	if o == nil || o.CashAccountType == nil {
		return nil, false
	}
	return o.CashAccountType, true
}

// HasCashAccountType returns a boolean if a field has been set.
func (o *AccountDetail) HasCashAccountType() bool {
	if o != nil && o.CashAccountType != nil {
		return true
	}

	return false
}

// SetCashAccountType gets a reference to the given string and assigns it to the CashAccountType field.
func (o *AccountDetail) SetCashAccountType(v string) {
	o.CashAccountType = &v
}

func (o AccountDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourceId != nil {
		toSerialize["resourceId"] = o.ResourceId
	}
	if o.Iban != nil {
		toSerialize["iban"] = o.Iban
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.OwnerName != nil {
		toSerialize["ownerName"] = o.OwnerName
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Product != nil {
		toSerialize["product"] = o.Product
	}
	if o.CashAccountType != nil {
		toSerialize["cashAccountType"] = o.CashAccountType
	}
	return json.Marshal(toSerialize)
}

type NullableAccountDetail struct {
	value *AccountDetail
	isSet bool
}

func (v NullableAccountDetail) Get() *AccountDetail {
	return v.value
}

func (v *NullableAccountDetail) Set(val *AccountDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountDetail(val *AccountDetail) *NullableAccountDetail {
	return &NullableAccountDetail{value: val, isSet: true}
}

func (v NullableAccountDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


