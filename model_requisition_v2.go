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

// RequisitionV2 Get requisition by id.
type RequisitionV2 struct {
	Id *string `json:"id,omitempty"`
	// The date & time at which the requisition was created.
	Created NullableTime `json:"created,omitempty"`
	// redirect URL to your application after end-user authorization with ASPSP
	Redirect NullableString `json:"redirect"`
	Status NullableStatus1c5Enum `json:"status,omitempty"`
	// an Institution ID for this Requisition
	InstitutionId string `json:"institution_id"`
	// EUA associated with this requisition
	Agreement *string `json:"agreement,omitempty"`
	// additional ID to identify the end user
	Reference *string `json:"reference,omitempty"`
	// array of account IDs retrieved within a scope of this requisition
	Accounts []string `json:"accounts,omitempty"`
	// A two-letter country code (ISO 639-1)
	UserLanguage *string `json:"user_language,omitempty"`
	// link to initiate authorization with Institution
	Link *string `json:"link,omitempty"`
	// optional SSN field to verify ownership of the account
	Ssn *string `json:"ssn,omitempty"`
	// option to enable account selection view for the end user
	AccountSelection *bool `json:"account_selection,omitempty"`
	// enable redirect back to the client after account list received
	RedirectImmediate *bool `json:"redirect_immediate,omitempty"`
}

// NewRequisitionV2 instantiates a new RequisitionV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequisitionV2(redirect NullableString, institutionId string) *RequisitionV2 {
	this := RequisitionV2{}
	this.Redirect = redirect
	this.InstitutionId = institutionId
	var accountSelection bool = false
	this.AccountSelection = &accountSelection
	var redirectImmediate bool = false
	this.RedirectImmediate = &redirectImmediate
	return &this
}

// NewRequisitionV2WithDefaults instantiates a new RequisitionV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequisitionV2WithDefaults() *RequisitionV2 {
	this := RequisitionV2{}
	var accountSelection bool = false
	this.AccountSelection = &accountSelection
	var redirectImmediate bool = false
	this.RedirectImmediate = &redirectImmediate
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RequisitionV2) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequisitionV2) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RequisitionV2) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RequisitionV2) SetId(v string) {
	o.Id = &v
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequisitionV2) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequisitionV2) GetCreatedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *RequisitionV2) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *RequisitionV2) SetCreated(v time.Time) {
	o.Created.Set(&v)
}
// SetCreatedNil sets the value for Created to be an explicit nil
func (o *RequisitionV2) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *RequisitionV2) UnsetCreated() {
	o.Created.Unset()
}

// GetRedirect returns the Redirect field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RequisitionV2) GetRedirect() string {
	if o == nil || o.Redirect.Get() == nil {
		var ret string
		return ret
	}

	return *o.Redirect.Get()
}

// GetRedirectOk returns a tuple with the Redirect field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequisitionV2) GetRedirectOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Redirect.Get(), o.Redirect.IsSet()
}

// SetRedirect sets field value
func (o *RequisitionV2) SetRedirect(v string) {
	o.Redirect.Set(&v)
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequisitionV2) GetStatus() Status1c5Enum {
	if o == nil || o.Status.Get() == nil {
		var ret Status1c5Enum
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequisitionV2) GetStatusOk() (*Status1c5Enum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *RequisitionV2) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableStatus1c5Enum and assigns it to the Status field.
func (o *RequisitionV2) SetStatus(v Status1c5Enum) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *RequisitionV2) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *RequisitionV2) UnsetStatus() {
	o.Status.Unset()
}

// GetInstitutionId returns the InstitutionId field value
func (o *RequisitionV2) GetInstitutionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstitutionId
}

// GetInstitutionIdOk returns a tuple with the InstitutionId field value
// and a boolean to check if the value has been set.
func (o *RequisitionV2) GetInstitutionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InstitutionId, true
}

// SetInstitutionId sets field value
func (o *RequisitionV2) SetInstitutionId(v string) {
	o.InstitutionId = v
}

// GetAgreement returns the Agreement field value if set, zero value otherwise.
func (o *RequisitionV2) GetAgreement() string {
	if o == nil || o.Agreement == nil {
		var ret string
		return ret
	}
	return *o.Agreement
}

// GetAgreementOk returns a tuple with the Agreement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequisitionV2) GetAgreementOk() (*string, bool) {
	if o == nil || o.Agreement == nil {
		return nil, false
	}
	return o.Agreement, true
}

// HasAgreement returns a boolean if a field has been set.
func (o *RequisitionV2) HasAgreement() bool {
	if o != nil && o.Agreement != nil {
		return true
	}

	return false
}

// SetAgreement gets a reference to the given string and assigns it to the Agreement field.
func (o *RequisitionV2) SetAgreement(v string) {
	o.Agreement = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *RequisitionV2) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequisitionV2) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *RequisitionV2) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *RequisitionV2) SetReference(v string) {
	o.Reference = &v
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *RequisitionV2) GetAccounts() []string {
	if o == nil || o.Accounts == nil {
		var ret []string
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequisitionV2) GetAccountsOk() ([]string, bool) {
	if o == nil || o.Accounts == nil {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *RequisitionV2) HasAccounts() bool {
	if o != nil && o.Accounts != nil {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []string and assigns it to the Accounts field.
func (o *RequisitionV2) SetAccounts(v []string) {
	o.Accounts = v
}

// GetUserLanguage returns the UserLanguage field value if set, zero value otherwise.
func (o *RequisitionV2) GetUserLanguage() string {
	if o == nil || o.UserLanguage == nil {
		var ret string
		return ret
	}
	return *o.UserLanguage
}

// GetUserLanguageOk returns a tuple with the UserLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequisitionV2) GetUserLanguageOk() (*string, bool) {
	if o == nil || o.UserLanguage == nil {
		return nil, false
	}
	return o.UserLanguage, true
}

// HasUserLanguage returns a boolean if a field has been set.
func (o *RequisitionV2) HasUserLanguage() bool {
	if o != nil && o.UserLanguage != nil {
		return true
	}

	return false
}

// SetUserLanguage gets a reference to the given string and assigns it to the UserLanguage field.
func (o *RequisitionV2) SetUserLanguage(v string) {
	o.UserLanguage = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *RequisitionV2) GetLink() string {
	if o == nil || o.Link == nil {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequisitionV2) GetLinkOk() (*string, bool) {
	if o == nil || o.Link == nil {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *RequisitionV2) HasLink() bool {
	if o != nil && o.Link != nil {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *RequisitionV2) SetLink(v string) {
	o.Link = &v
}

// GetSsn returns the Ssn field value if set, zero value otherwise.
func (o *RequisitionV2) GetSsn() string {
	if o == nil || o.Ssn == nil {
		var ret string
		return ret
	}
	return *o.Ssn
}

// GetSsnOk returns a tuple with the Ssn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequisitionV2) GetSsnOk() (*string, bool) {
	if o == nil || o.Ssn == nil {
		return nil, false
	}
	return o.Ssn, true
}

// HasSsn returns a boolean if a field has been set.
func (o *RequisitionV2) HasSsn() bool {
	if o != nil && o.Ssn != nil {
		return true
	}

	return false
}

// SetSsn gets a reference to the given string and assigns it to the Ssn field.
func (o *RequisitionV2) SetSsn(v string) {
	o.Ssn = &v
}

// GetAccountSelection returns the AccountSelection field value if set, zero value otherwise.
func (o *RequisitionV2) GetAccountSelection() bool {
	if o == nil || o.AccountSelection == nil {
		var ret bool
		return ret
	}
	return *o.AccountSelection
}

// GetAccountSelectionOk returns a tuple with the AccountSelection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequisitionV2) GetAccountSelectionOk() (*bool, bool) {
	if o == nil || o.AccountSelection == nil {
		return nil, false
	}
	return o.AccountSelection, true
}

// HasAccountSelection returns a boolean if a field has been set.
func (o *RequisitionV2) HasAccountSelection() bool {
	if o != nil && o.AccountSelection != nil {
		return true
	}

	return false
}

// SetAccountSelection gets a reference to the given bool and assigns it to the AccountSelection field.
func (o *RequisitionV2) SetAccountSelection(v bool) {
	o.AccountSelection = &v
}

// GetRedirectImmediate returns the RedirectImmediate field value if set, zero value otherwise.
func (o *RequisitionV2) GetRedirectImmediate() bool {
	if o == nil || o.RedirectImmediate == nil {
		var ret bool
		return ret
	}
	return *o.RedirectImmediate
}

// GetRedirectImmediateOk returns a tuple with the RedirectImmediate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequisitionV2) GetRedirectImmediateOk() (*bool, bool) {
	if o == nil || o.RedirectImmediate == nil {
		return nil, false
	}
	return o.RedirectImmediate, true
}

// HasRedirectImmediate returns a boolean if a field has been set.
func (o *RequisitionV2) HasRedirectImmediate() bool {
	if o != nil && o.RedirectImmediate != nil {
		return true
	}

	return false
}

// SetRedirectImmediate gets a reference to the given bool and assigns it to the RedirectImmediate field.
func (o *RequisitionV2) SetRedirectImmediate(v bool) {
	o.RedirectImmediate = &v
}

func (o RequisitionV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
	}
	if true {
		toSerialize["redirect"] = o.Redirect.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
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
	if o.Accounts != nil {
		toSerialize["accounts"] = o.Accounts
	}
	if o.UserLanguage != nil {
		toSerialize["user_language"] = o.UserLanguage
	}
	if o.Link != nil {
		toSerialize["link"] = o.Link
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

type NullableRequisitionV2 struct {
	value *RequisitionV2
	isSet bool
}

func (v NullableRequisitionV2) Get() *RequisitionV2 {
	return v.value
}

func (v *NullableRequisitionV2) Set(val *RequisitionV2) {
	v.value = val
	v.isSet = true
}

func (v NullableRequisitionV2) IsSet() bool {
	return v.isSet
}

func (v *NullableRequisitionV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequisitionV2(val *RequisitionV2) *NullableRequisitionV2 {
	return &NullableRequisitionV2{value: val, isSet: true}
}

func (v NullableRequisitionV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequisitionV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


