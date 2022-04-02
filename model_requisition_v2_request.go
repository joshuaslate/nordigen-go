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

// RequisitionV2Request Get requisition by id.
type RequisitionV2Request struct {
	// redirect URL to your application after end-user authorization with ASPSP
	Redirect NullableString `json:"redirect"`
	// an Institution ID for this Requisition
	InstitutionId string `json:"institution_id"`
	// EUA associated with this requisition
	Agreement *string `json:"agreement,omitempty"`
	// additional ID to identify the end user
	Reference *string `json:"reference,omitempty"`
	// A two-letter country code (ISO 639-1)
	UserLanguage *string `json:"user_language,omitempty"`
	// optional SSN field to verify ownership of the account
	Ssn *string `json:"ssn,omitempty"`
	// option to enable account selection view for the end user
	AccountSelection *bool `json:"account_selection,omitempty"`
	// enable redirect back to the client after account list received
	RedirectImmediate *bool `json:"redirect_immediate,omitempty"`
}

// NewRequisitionV2Request instantiates a new RequisitionV2Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequisitionV2Request(redirect NullableString, institutionId string) *RequisitionV2Request {
	this := RequisitionV2Request{}
	this.Redirect = redirect
	this.InstitutionId = institutionId
	var accountSelection bool = false
	this.AccountSelection = &accountSelection
	var redirectImmediate bool = false
	this.RedirectImmediate = &redirectImmediate
	return &this
}

// NewRequisitionV2RequestWithDefaults instantiates a new RequisitionV2Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequisitionV2RequestWithDefaults() *RequisitionV2Request {
	this := RequisitionV2Request{}
	var accountSelection bool = false
	this.AccountSelection = &accountSelection
	var redirectImmediate bool = false
	this.RedirectImmediate = &redirectImmediate
	return &this
}

// GetRedirect returns the Redirect field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RequisitionV2Request) GetRedirect() string {
	if o == nil || o.Redirect.Get() == nil {
		var ret string
		return ret
	}

	return *o.Redirect.Get()
}

// GetRedirectOk returns a tuple with the Redirect field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequisitionV2Request) GetRedirectOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Redirect.Get(), o.Redirect.IsSet()
}

// SetRedirect sets field value
func (o *RequisitionV2Request) SetRedirect(v string) {
	o.Redirect.Set(&v)
}

// GetInstitutionId returns the InstitutionId field value
func (o *RequisitionV2Request) GetInstitutionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstitutionId
}

// GetInstitutionIdOk returns a tuple with the InstitutionId field value
// and a boolean to check if the value has been set.
func (o *RequisitionV2Request) GetInstitutionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InstitutionId, true
}

// SetInstitutionId sets field value
func (o *RequisitionV2Request) SetInstitutionId(v string) {
	o.InstitutionId = v
}

// GetAgreement returns the Agreement field value if set, zero value otherwise.
func (o *RequisitionV2Request) GetAgreement() string {
	if o == nil || o.Agreement == nil {
		var ret string
		return ret
	}
	return *o.Agreement
}

// GetAgreementOk returns a tuple with the Agreement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequisitionV2Request) GetAgreementOk() (*string, bool) {
	if o == nil || o.Agreement == nil {
		return nil, false
	}
	return o.Agreement, true
}

// HasAgreement returns a boolean if a field has been set.
func (o *RequisitionV2Request) HasAgreement() bool {
	if o != nil && o.Agreement != nil {
		return true
	}

	return false
}

// SetAgreement gets a reference to the given string and assigns it to the Agreement field.
func (o *RequisitionV2Request) SetAgreement(v string) {
	o.Agreement = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *RequisitionV2Request) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequisitionV2Request) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *RequisitionV2Request) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *RequisitionV2Request) SetReference(v string) {
	o.Reference = &v
}

// GetUserLanguage returns the UserLanguage field value if set, zero value otherwise.
func (o *RequisitionV2Request) GetUserLanguage() string {
	if o == nil || o.UserLanguage == nil {
		var ret string
		return ret
	}
	return *o.UserLanguage
}

// GetUserLanguageOk returns a tuple with the UserLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequisitionV2Request) GetUserLanguageOk() (*string, bool) {
	if o == nil || o.UserLanguage == nil {
		return nil, false
	}
	return o.UserLanguage, true
}

// HasUserLanguage returns a boolean if a field has been set.
func (o *RequisitionV2Request) HasUserLanguage() bool {
	if o != nil && o.UserLanguage != nil {
		return true
	}

	return false
}

// SetUserLanguage gets a reference to the given string and assigns it to the UserLanguage field.
func (o *RequisitionV2Request) SetUserLanguage(v string) {
	o.UserLanguage = &v
}

// GetSsn returns the Ssn field value if set, zero value otherwise.
func (o *RequisitionV2Request) GetSsn() string {
	if o == nil || o.Ssn == nil {
		var ret string
		return ret
	}
	return *o.Ssn
}

// GetSsnOk returns a tuple with the Ssn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequisitionV2Request) GetSsnOk() (*string, bool) {
	if o == nil || o.Ssn == nil {
		return nil, false
	}
	return o.Ssn, true
}

// HasSsn returns a boolean if a field has been set.
func (o *RequisitionV2Request) HasSsn() bool {
	if o != nil && o.Ssn != nil {
		return true
	}

	return false
}

// SetSsn gets a reference to the given string and assigns it to the Ssn field.
func (o *RequisitionV2Request) SetSsn(v string) {
	o.Ssn = &v
}

// GetAccountSelection returns the AccountSelection field value if set, zero value otherwise.
func (o *RequisitionV2Request) GetAccountSelection() bool {
	if o == nil || o.AccountSelection == nil {
		var ret bool
		return ret
	}
	return *o.AccountSelection
}

// GetAccountSelectionOk returns a tuple with the AccountSelection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequisitionV2Request) GetAccountSelectionOk() (*bool, bool) {
	if o == nil || o.AccountSelection == nil {
		return nil, false
	}
	return o.AccountSelection, true
}

// HasAccountSelection returns a boolean if a field has been set.
func (o *RequisitionV2Request) HasAccountSelection() bool {
	if o != nil && o.AccountSelection != nil {
		return true
	}

	return false
}

// SetAccountSelection gets a reference to the given bool and assigns it to the AccountSelection field.
func (o *RequisitionV2Request) SetAccountSelection(v bool) {
	o.AccountSelection = &v
}

// GetRedirectImmediate returns the RedirectImmediate field value if set, zero value otherwise.
func (o *RequisitionV2Request) GetRedirectImmediate() bool {
	if o == nil || o.RedirectImmediate == nil {
		var ret bool
		return ret
	}
	return *o.RedirectImmediate
}

// GetRedirectImmediateOk returns a tuple with the RedirectImmediate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequisitionV2Request) GetRedirectImmediateOk() (*bool, bool) {
	if o == nil || o.RedirectImmediate == nil {
		return nil, false
	}
	return o.RedirectImmediate, true
}

// HasRedirectImmediate returns a boolean if a field has been set.
func (o *RequisitionV2Request) HasRedirectImmediate() bool {
	if o != nil && o.RedirectImmediate != nil {
		return true
	}

	return false
}

// SetRedirectImmediate gets a reference to the given bool and assigns it to the RedirectImmediate field.
func (o *RequisitionV2Request) SetRedirectImmediate(v bool) {
	o.RedirectImmediate = &v
}

func (o RequisitionV2Request) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["redirect"] = o.Redirect.Get()
	}
	if true {
		toSerialize["institution_id"] = o.InstitutionId
	}
	if o.Agreement != nil {
		toSerialize["agreement"] = o.Agreement
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.UserLanguage != nil {
		toSerialize["user_language"] = o.UserLanguage
	}
	if o.Ssn != nil {
		toSerialize["ssn"] = o.Ssn
	}
	if o.AccountSelection != nil {
		toSerialize["account_selection"] = o.AccountSelection
	}
	if o.RedirectImmediate != nil {
		toSerialize["redirect_immediate"] = o.RedirectImmediate
	}
	return json.Marshal(toSerialize)
}

type NullableRequisitionV2Request struct {
	value *RequisitionV2Request
	isSet bool
}

func (v NullableRequisitionV2Request) Get() *RequisitionV2Request {
	return v.value
}

func (v *NullableRequisitionV2Request) Set(val *RequisitionV2Request) {
	v.value = val
	v.isSet = true
}

func (v NullableRequisitionV2Request) IsSet() bool {
	return v.isSet
}

func (v *NullableRequisitionV2Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequisitionV2Request(val *RequisitionV2Request) *NullableRequisitionV2Request {
	return &NullableRequisitionV2Request{value: val, isSet: true}
}

func (v NullableRequisitionV2Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequisitionV2Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

