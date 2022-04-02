# RequisitionV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Created** | Pointer to **NullableTime** | The date &amp; time at which the requisition was created. | [optional] [readonly] 
**Redirect** | **NullableString** | redirect URL to your application after end-user authorization with ASPSP | 
**Status** | Pointer to [**NullableStatus1c5Enum**](Status1c5Enum.md) |  | [optional] 
**InstitutionId** | **string** | an Institution ID for this Requisition | 
**Agreement** | Pointer to **string** | EUA associated with this requisition | [optional] 
**Reference** | Pointer to **string** | additional ID to identify the end user | [optional] 
**Accounts** | Pointer to **[]string** | array of account IDs retrieved within a scope of this requisition | [optional] [readonly] 
**UserLanguage** | Pointer to **string** | A two-letter country code (ISO 639-1) | [optional] 
**Link** | Pointer to **string** | link to initiate authorization with Institution | [optional] [readonly] [default to "https://ob.nordigen.com/psd2/start/3fa85f64-5717-4562-b3fc-2c963f66afa6/{$INSTITUTION_ID}"]
**Ssn** | Pointer to **string** | optional SSN field to verify ownership of the account | [optional] 
**AccountSelection** | Pointer to **bool** | option to enable account selection view for the end user | [optional] [default to false]
**RedirectImmediate** | Pointer to **bool** | enable redirect back to the client after account list received | [optional] [default to false]

## Methods

### NewRequisitionV2

`func NewRequisitionV2(redirect NullableString, institutionId string, ) *RequisitionV2`

NewRequisitionV2 instantiates a new RequisitionV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequisitionV2WithDefaults

`func NewRequisitionV2WithDefaults() *RequisitionV2`

NewRequisitionV2WithDefaults instantiates a new RequisitionV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequisitionV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequisitionV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequisitionV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RequisitionV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreated

`func (o *RequisitionV2) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RequisitionV2) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RequisitionV2) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RequisitionV2) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *RequisitionV2) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *RequisitionV2) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetRedirect

`func (o *RequisitionV2) GetRedirect() string`

GetRedirect returns the Redirect field if non-nil, zero value otherwise.

### GetRedirectOk

`func (o *RequisitionV2) GetRedirectOk() (*string, bool)`

GetRedirectOk returns a tuple with the Redirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirect

`func (o *RequisitionV2) SetRedirect(v string)`

SetRedirect sets Redirect field to given value.


### SetRedirectNil

`func (o *RequisitionV2) SetRedirectNil(b bool)`

 SetRedirectNil sets the value for Redirect to be an explicit nil

### UnsetRedirect
`func (o *RequisitionV2) UnsetRedirect()`

UnsetRedirect ensures that no value is present for Redirect, not even an explicit nil
### GetStatus

`func (o *RequisitionV2) GetStatus() Status1c5Enum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequisitionV2) GetStatusOk() (*Status1c5Enum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequisitionV2) SetStatus(v Status1c5Enum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RequisitionV2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *RequisitionV2) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *RequisitionV2) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetInstitutionId

`func (o *RequisitionV2) GetInstitutionId() string`

GetInstitutionId returns the InstitutionId field if non-nil, zero value otherwise.

### GetInstitutionIdOk

`func (o *RequisitionV2) GetInstitutionIdOk() (*string, bool)`

GetInstitutionIdOk returns a tuple with the InstitutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionId

`func (o *RequisitionV2) SetInstitutionId(v string)`

SetInstitutionId sets InstitutionId field to given value.


### GetAgreement

`func (o *RequisitionV2) GetAgreement() string`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *RequisitionV2) GetAgreementOk() (*string, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *RequisitionV2) SetAgreement(v string)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *RequisitionV2) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetReference

`func (o *RequisitionV2) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *RequisitionV2) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *RequisitionV2) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *RequisitionV2) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetAccounts

`func (o *RequisitionV2) GetAccounts() []string`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *RequisitionV2) GetAccountsOk() (*[]string, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *RequisitionV2) SetAccounts(v []string)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *RequisitionV2) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetUserLanguage

`func (o *RequisitionV2) GetUserLanguage() string`

GetUserLanguage returns the UserLanguage field if non-nil, zero value otherwise.

### GetUserLanguageOk

`func (o *RequisitionV2) GetUserLanguageOk() (*string, bool)`

GetUserLanguageOk returns a tuple with the UserLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLanguage

`func (o *RequisitionV2) SetUserLanguage(v string)`

SetUserLanguage sets UserLanguage field to given value.

### HasUserLanguage

`func (o *RequisitionV2) HasUserLanguage() bool`

HasUserLanguage returns a boolean if a field has been set.

### GetLink

`func (o *RequisitionV2) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *RequisitionV2) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *RequisitionV2) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *RequisitionV2) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetSsn

`func (o *RequisitionV2) GetSsn() string`

GetSsn returns the Ssn field if non-nil, zero value otherwise.

### GetSsnOk

`func (o *RequisitionV2) GetSsnOk() (*string, bool)`

GetSsnOk returns a tuple with the Ssn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn

`func (o *RequisitionV2) SetSsn(v string)`

SetSsn sets Ssn field to given value.

### HasSsn

`func (o *RequisitionV2) HasSsn() bool`

HasSsn returns a boolean if a field has been set.

### GetAccountSelection

`func (o *RequisitionV2) GetAccountSelection() bool`

GetAccountSelection returns the AccountSelection field if non-nil, zero value otherwise.

### GetAccountSelectionOk

`func (o *RequisitionV2) GetAccountSelectionOk() (*bool, bool)`

GetAccountSelectionOk returns a tuple with the AccountSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSelection

`func (o *RequisitionV2) SetAccountSelection(v bool)`

SetAccountSelection sets AccountSelection field to given value.

### HasAccountSelection

`func (o *RequisitionV2) HasAccountSelection() bool`

HasAccountSelection returns a boolean if a field has been set.

### GetRedirectImmediate

`func (o *RequisitionV2) GetRedirectImmediate() bool`

GetRedirectImmediate returns the RedirectImmediate field if non-nil, zero value otherwise.

### GetRedirectImmediateOk

`func (o *RequisitionV2) GetRedirectImmediateOk() (*bool, bool)`

GetRedirectImmediateOk returns a tuple with the RedirectImmediate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectImmediate

`func (o *RequisitionV2) SetRedirectImmediate(v bool)`

SetRedirectImmediate sets RedirectImmediate field to given value.

### HasRedirectImmediate

`func (o *RequisitionV2) HasRedirectImmediate() bool`

HasRedirectImmediate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


