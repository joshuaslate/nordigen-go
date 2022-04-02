# RequisitionV2Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Redirect** | **NullableString** | redirect URL to your application after end-user authorization with ASPSP | 
**InstitutionId** | **string** | an Institution ID for this Requisition | 
**Agreement** | Pointer to **string** | EUA associated with this requisition | [optional] 
**Reference** | Pointer to **string** | additional ID to identify the end user | [optional] 
**UserLanguage** | Pointer to **string** | A two-letter country code (ISO 639-1) | [optional] 
**Ssn** | Pointer to **string** | optional SSN field to verify ownership of the account | [optional] 
**AccountSelection** | Pointer to **bool** | option to enable account selection view for the end user | [optional] [default to false]
**RedirectImmediate** | Pointer to **bool** | enable redirect back to the client after account list received | [optional] [default to false]

## Methods

### NewRequisitionV2Request

`func NewRequisitionV2Request(redirect NullableString, institutionId string, ) *RequisitionV2Request`

NewRequisitionV2Request instantiates a new RequisitionV2Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequisitionV2RequestWithDefaults

`func NewRequisitionV2RequestWithDefaults() *RequisitionV2Request`

NewRequisitionV2RequestWithDefaults instantiates a new RequisitionV2Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedirect

`func (o *RequisitionV2Request) GetRedirect() string`

GetRedirect returns the Redirect field if non-nil, zero value otherwise.

### GetRedirectOk

`func (o *RequisitionV2Request) GetRedirectOk() (*string, bool)`

GetRedirectOk returns a tuple with the Redirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirect

`func (o *RequisitionV2Request) SetRedirect(v string)`

SetRedirect sets Redirect field to given value.


### SetRedirectNil

`func (o *RequisitionV2Request) SetRedirectNil(b bool)`

 SetRedirectNil sets the value for Redirect to be an explicit nil

### UnsetRedirect
`func (o *RequisitionV2Request) UnsetRedirect()`

UnsetRedirect ensures that no value is present for Redirect, not even an explicit nil
### GetInstitutionId

`func (o *RequisitionV2Request) GetInstitutionId() string`

GetInstitutionId returns the InstitutionId field if non-nil, zero value otherwise.

### GetInstitutionIdOk

`func (o *RequisitionV2Request) GetInstitutionIdOk() (*string, bool)`

GetInstitutionIdOk returns a tuple with the InstitutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionId

`func (o *RequisitionV2Request) SetInstitutionId(v string)`

SetInstitutionId sets InstitutionId field to given value.


### GetAgreement

`func (o *RequisitionV2Request) GetAgreement() string`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *RequisitionV2Request) GetAgreementOk() (*string, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *RequisitionV2Request) SetAgreement(v string)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *RequisitionV2Request) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetReference

`func (o *RequisitionV2Request) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *RequisitionV2Request) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *RequisitionV2Request) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *RequisitionV2Request) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetUserLanguage

`func (o *RequisitionV2Request) GetUserLanguage() string`

GetUserLanguage returns the UserLanguage field if non-nil, zero value otherwise.

### GetUserLanguageOk

`func (o *RequisitionV2Request) GetUserLanguageOk() (*string, bool)`

GetUserLanguageOk returns a tuple with the UserLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLanguage

`func (o *RequisitionV2Request) SetUserLanguage(v string)`

SetUserLanguage sets UserLanguage field to given value.

### HasUserLanguage

`func (o *RequisitionV2Request) HasUserLanguage() bool`

HasUserLanguage returns a boolean if a field has been set.

### GetSsn

`func (o *RequisitionV2Request) GetSsn() string`

GetSsn returns the Ssn field if non-nil, zero value otherwise.

### GetSsnOk

`func (o *RequisitionV2Request) GetSsnOk() (*string, bool)`

GetSsnOk returns a tuple with the Ssn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn

`func (o *RequisitionV2Request) SetSsn(v string)`

SetSsn sets Ssn field to given value.

### HasSsn

`func (o *RequisitionV2Request) HasSsn() bool`

HasSsn returns a boolean if a field has been set.

### GetAccountSelection

`func (o *RequisitionV2Request) GetAccountSelection() bool`

GetAccountSelection returns the AccountSelection field if non-nil, zero value otherwise.

### GetAccountSelectionOk

`func (o *RequisitionV2Request) GetAccountSelectionOk() (*bool, bool)`

GetAccountSelectionOk returns a tuple with the AccountSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSelection

`func (o *RequisitionV2Request) SetAccountSelection(v bool)`

SetAccountSelection sets AccountSelection field to given value.

### HasAccountSelection

`func (o *RequisitionV2Request) HasAccountSelection() bool`

HasAccountSelection returns a boolean if a field has been set.

### GetRedirectImmediate

`func (o *RequisitionV2Request) GetRedirectImmediate() bool`

GetRedirectImmediate returns the RedirectImmediate field if non-nil, zero value otherwise.

### GetRedirectImmediateOk

`func (o *RequisitionV2Request) GetRedirectImmediateOk() (*bool, bool)`

GetRedirectImmediateOk returns a tuple with the RedirectImmediate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectImmediate

`func (o *RequisitionV2Request) SetRedirectImmediate(v bool)`

SetRedirectImmediate sets RedirectImmediate field to given value.

### HasRedirectImmediate

`func (o *RequisitionV2Request) HasRedirectImmediate() bool`

HasRedirectImmediate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


