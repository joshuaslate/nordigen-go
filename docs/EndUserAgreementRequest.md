# EndUserAgreementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxHistoricalDays** | Pointer to **int32** | Maximum number of days of transaction data to retrieve. | [optional] [default to 90]
**AccessValidForDays** | Pointer to **int32** | Number of days from acceptance that the access can be used. | [optional] [default to 90]
**AccessScope** | Pointer to **[]string** | Array containing one or several values of [&#39;balances&#39;, &#39;details&#39;, &#39;transactions&#39;] | [optional] [default to ["balances","details","transactions"]]
**InstitutionId** | **string** | an Institution ID for this EUA | 

## Methods

### NewEndUserAgreementRequest

`func NewEndUserAgreementRequest(institutionId string, ) *EndUserAgreementRequest`

NewEndUserAgreementRequest instantiates a new EndUserAgreementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndUserAgreementRequestWithDefaults

`func NewEndUserAgreementRequestWithDefaults() *EndUserAgreementRequest`

NewEndUserAgreementRequestWithDefaults instantiates a new EndUserAgreementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxHistoricalDays

`func (o *EndUserAgreementRequest) GetMaxHistoricalDays() int32`

GetMaxHistoricalDays returns the MaxHistoricalDays field if non-nil, zero value otherwise.

### GetMaxHistoricalDaysOk

`func (o *EndUserAgreementRequest) GetMaxHistoricalDaysOk() (*int32, bool)`

GetMaxHistoricalDaysOk returns a tuple with the MaxHistoricalDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHistoricalDays

`func (o *EndUserAgreementRequest) SetMaxHistoricalDays(v int32)`

SetMaxHistoricalDays sets MaxHistoricalDays field to given value.

### HasMaxHistoricalDays

`func (o *EndUserAgreementRequest) HasMaxHistoricalDays() bool`

HasMaxHistoricalDays returns a boolean if a field has been set.

### GetAccessValidForDays

`func (o *EndUserAgreementRequest) GetAccessValidForDays() int32`

GetAccessValidForDays returns the AccessValidForDays field if non-nil, zero value otherwise.

### GetAccessValidForDaysOk

`func (o *EndUserAgreementRequest) GetAccessValidForDaysOk() (*int32, bool)`

GetAccessValidForDaysOk returns a tuple with the AccessValidForDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessValidForDays

`func (o *EndUserAgreementRequest) SetAccessValidForDays(v int32)`

SetAccessValidForDays sets AccessValidForDays field to given value.

### HasAccessValidForDays

`func (o *EndUserAgreementRequest) HasAccessValidForDays() bool`

HasAccessValidForDays returns a boolean if a field has been set.

### GetAccessScope

`func (o *EndUserAgreementRequest) GetAccessScope() []string`

GetAccessScope returns the AccessScope field if non-nil, zero value otherwise.

### GetAccessScopeOk

`func (o *EndUserAgreementRequest) GetAccessScopeOk() (*[]string, bool)`

GetAccessScopeOk returns a tuple with the AccessScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessScope

`func (o *EndUserAgreementRequest) SetAccessScope(v []string)`

SetAccessScope sets AccessScope field to given value.

### HasAccessScope

`func (o *EndUserAgreementRequest) HasAccessScope() bool`

HasAccessScope returns a boolean if a field has been set.

### GetInstitutionId

`func (o *EndUserAgreementRequest) GetInstitutionId() string`

GetInstitutionId returns the InstitutionId field if non-nil, zero value otherwise.

### GetInstitutionIdOk

`func (o *EndUserAgreementRequest) GetInstitutionIdOk() (*string, bool)`

GetInstitutionIdOk returns a tuple with the InstitutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionId

`func (o *EndUserAgreementRequest) SetInstitutionId(v string)`

SetInstitutionId sets InstitutionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


