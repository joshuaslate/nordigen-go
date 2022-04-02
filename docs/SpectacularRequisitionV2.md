# SpectacularRequisitionV2

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
**Accounts** | Pointer to **[]interface{}** | array of account IDs retrieved within a scope of this requisition | [optional] [readonly] [default to []]
**UserLanguage** | Pointer to **string** | A two-letter country code (ISO 639-1) | [optional] 
**Link** | Pointer to **string** | link to initiate authorization with Institution | [optional] [readonly] [default to "https://ob.nordigen.com/psd2/start/3fa85f64-5717-4562-b3fc-2c963f66afa6/{$INSTITUTION_ID}"]
**Ssn** | Pointer to **string** | optional SSN field to verify ownership of the account | [optional] 
**AccountSelection** | Pointer to **bool** | option to enable account selection view for the end user | [optional] [default to false]
**RedirectImmediate** | Pointer to **bool** | enable redirect back to the client after account list received | [optional] [default to false]

## Methods

### NewSpectacularRequisitionV2

`func NewSpectacularRequisitionV2(redirect NullableString, institutionId string, ) *SpectacularRequisitionV2`

NewSpectacularRequisitionV2 instantiates a new SpectacularRequisitionV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpectacularRequisitionV2WithDefaults

`func NewSpectacularRequisitionV2WithDefaults() *SpectacularRequisitionV2`

NewSpectacularRequisitionV2WithDefaults instantiates a new SpectacularRequisitionV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SpectacularRequisitionV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpectacularRequisitionV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpectacularRequisitionV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SpectacularRequisitionV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreated

`func (o *SpectacularRequisitionV2) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SpectacularRequisitionV2) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SpectacularRequisitionV2) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SpectacularRequisitionV2) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *SpectacularRequisitionV2) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *SpectacularRequisitionV2) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetRedirect

`func (o *SpectacularRequisitionV2) GetRedirect() string`

GetRedirect returns the Redirect field if non-nil, zero value otherwise.

### GetRedirectOk

`func (o *SpectacularRequisitionV2) GetRedirectOk() (*string, bool)`

GetRedirectOk returns a tuple with the Redirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirect

`func (o *SpectacularRequisitionV2) SetRedirect(v string)`

SetRedirect sets Redirect field to given value.


### SetRedirectNil

`func (o *SpectacularRequisitionV2) SetRedirectNil(b bool)`

 SetRedirectNil sets the value for Redirect to be an explicit nil

### UnsetRedirect
`func (o *SpectacularRequisitionV2) UnsetRedirect()`

UnsetRedirect ensures that no value is present for Redirect, not even an explicit nil
### GetStatus

`func (o *SpectacularRequisitionV2) GetStatus() Status1c5Enum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SpectacularRequisitionV2) GetStatusOk() (*Status1c5Enum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SpectacularRequisitionV2) SetStatus(v Status1c5Enum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SpectacularRequisitionV2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *SpectacularRequisitionV2) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *SpectacularRequisitionV2) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetInstitutionId

`func (o *SpectacularRequisitionV2) GetInstitutionId() string`

GetInstitutionId returns the InstitutionId field if non-nil, zero value otherwise.

### GetInstitutionIdOk

`func (o *SpectacularRequisitionV2) GetInstitutionIdOk() (*string, bool)`

GetInstitutionIdOk returns a tuple with the InstitutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionId

`func (o *SpectacularRequisitionV2) SetInstitutionId(v string)`

SetInstitutionId sets InstitutionId field to given value.


### GetAgreement

`func (o *SpectacularRequisitionV2) GetAgreement() string`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *SpectacularRequisitionV2) GetAgreementOk() (*string, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *SpectacularRequisitionV2) SetAgreement(v string)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *SpectacularRequisitionV2) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetReference

`func (o *SpectacularRequisitionV2) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *SpectacularRequisitionV2) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *SpectacularRequisitionV2) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *SpectacularRequisitionV2) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetAccounts

`func (o *SpectacularRequisitionV2) GetAccounts() []interface{}`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *SpectacularRequisitionV2) GetAccountsOk() (*[]interface{}, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *SpectacularRequisitionV2) SetAccounts(v []interface{})`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *SpectacularRequisitionV2) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetUserLanguage

`func (o *SpectacularRequisitionV2) GetUserLanguage() string`

GetUserLanguage returns the UserLanguage field if non-nil, zero value otherwise.

### GetUserLanguageOk

`func (o *SpectacularRequisitionV2) GetUserLanguageOk() (*string, bool)`

GetUserLanguageOk returns a tuple with the UserLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLanguage

`func (o *SpectacularRequisitionV2) SetUserLanguage(v string)`

SetUserLanguage sets UserLanguage field to given value.

### HasUserLanguage

`func (o *SpectacularRequisitionV2) HasUserLanguage() bool`

HasUserLanguage returns a boolean if a field has been set.

### GetLink

`func (o *SpectacularRequisitionV2) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *SpectacularRequisitionV2) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *SpectacularRequisitionV2) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *SpectacularRequisitionV2) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetSsn

`func (o *SpectacularRequisitionV2) GetSsn() string`

GetSsn returns the Ssn field if non-nil, zero value otherwise.

### GetSsnOk

`func (o *SpectacularRequisitionV2) GetSsnOk() (*string, bool)`

GetSsnOk returns a tuple with the Ssn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn

`func (o *SpectacularRequisitionV2) SetSsn(v string)`

SetSsn sets Ssn field to given value.

### HasSsn

`func (o *SpectacularRequisitionV2) HasSsn() bool`

HasSsn returns a boolean if a field has been set.

### GetAccountSelection

`func (o *SpectacularRequisitionV2) GetAccountSelection() bool`

GetAccountSelection returns the AccountSelection field if non-nil, zero value otherwise.

### GetAccountSelectionOk

`func (o *SpectacularRequisitionV2) GetAccountSelectionOk() (*bool, bool)`

GetAccountSelectionOk returns a tuple with the AccountSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSelection

`func (o *SpectacularRequisitionV2) SetAccountSelection(v bool)`

SetAccountSelection sets AccountSelection field to given value.

### HasAccountSelection

`func (o *SpectacularRequisitionV2) HasAccountSelection() bool`

HasAccountSelection returns a boolean if a field has been set.

### GetRedirectImmediate

`func (o *SpectacularRequisitionV2) GetRedirectImmediate() bool`

GetRedirectImmediate returns the RedirectImmediate field if non-nil, zero value otherwise.

### GetRedirectImmediateOk

`func (o *SpectacularRequisitionV2) GetRedirectImmediateOk() (*bool, bool)`

GetRedirectImmediateOk returns a tuple with the RedirectImmediate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectImmediate

`func (o *SpectacularRequisitionV2) SetRedirectImmediate(v bool)`

SetRedirectImmediate sets RedirectImmediate field to given value.

### HasRedirectImmediate

`func (o *SpectacularRequisitionV2) HasRedirectImmediate() bool`

HasRedirectImmediate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


