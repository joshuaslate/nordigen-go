# EndUserAgreement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of this End User Agreement, used to refer to this end user agreement in other API calls. | [optional] [readonly] 
**Created** | Pointer to **time.Time** | The date &amp; time at which the end user agreement was created. | [optional] [readonly] 
**MaxHistoricalDays** | Pointer to **int32** | Maximum number of days of transaction data to retrieve. | [optional] [default to 90]
**AccessValidForDays** | Pointer to **int32** | Number of days from acceptance that the access can be used. | [optional] [default to 90]
**AccessScope** | Pointer to **[]string** | Array containing one or several values of [&#39;balances&#39;, &#39;details&#39;, &#39;transactions&#39;] | [optional] [default to ["balances","details","transactions"]]
**Accepted** | Pointer to **NullableTime** | The date &amp; time at which the end user accepted the agreement. | [optional] [readonly] 
**InstitutionId** | **string** | an Institution ID for this EUA | 

## Methods

### NewEndUserAgreement

`func NewEndUserAgreement(institutionId string, ) *EndUserAgreement`

NewEndUserAgreement instantiates a new EndUserAgreement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndUserAgreementWithDefaults

`func NewEndUserAgreementWithDefaults() *EndUserAgreement`

NewEndUserAgreementWithDefaults instantiates a new EndUserAgreement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EndUserAgreement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EndUserAgreement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EndUserAgreement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EndUserAgreement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreated

`func (o *EndUserAgreement) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EndUserAgreement) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *EndUserAgreement) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *EndUserAgreement) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetMaxHistoricalDays

`func (o *EndUserAgreement) GetMaxHistoricalDays() int32`

GetMaxHistoricalDays returns the MaxHistoricalDays field if non-nil, zero value otherwise.

### GetMaxHistoricalDaysOk

`func (o *EndUserAgreement) GetMaxHistoricalDaysOk() (*int32, bool)`

GetMaxHistoricalDaysOk returns a tuple with the MaxHistoricalDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHistoricalDays

`func (o *EndUserAgreement) SetMaxHistoricalDays(v int32)`

SetMaxHistoricalDays sets MaxHistoricalDays field to given value.

### HasMaxHistoricalDays

`func (o *EndUserAgreement) HasMaxHistoricalDays() bool`

HasMaxHistoricalDays returns a boolean if a field has been set.

### GetAccessValidForDays

`func (o *EndUserAgreement) GetAccessValidForDays() int32`

GetAccessValidForDays returns the AccessValidForDays field if non-nil, zero value otherwise.

### GetAccessValidForDaysOk

`func (o *EndUserAgreement) GetAccessValidForDaysOk() (*int32, bool)`

GetAccessValidForDaysOk returns a tuple with the AccessValidForDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessValidForDays

`func (o *EndUserAgreement) SetAccessValidForDays(v int32)`

SetAccessValidForDays sets AccessValidForDays field to given value.

### HasAccessValidForDays

`func (o *EndUserAgreement) HasAccessValidForDays() bool`

HasAccessValidForDays returns a boolean if a field has been set.

### GetAccessScope

`func (o *EndUserAgreement) GetAccessScope() []string`

GetAccessScope returns the AccessScope field if non-nil, zero value otherwise.

### GetAccessScopeOk

`func (o *EndUserAgreement) GetAccessScopeOk() (*[]string, bool)`

GetAccessScopeOk returns a tuple with the AccessScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessScope

`func (o *EndUserAgreement) SetAccessScope(v []string)`

SetAccessScope sets AccessScope field to given value.

### HasAccessScope

`func (o *EndUserAgreement) HasAccessScope() bool`

HasAccessScope returns a boolean if a field has been set.

### GetAccepted

`func (o *EndUserAgreement) GetAccepted() time.Time`

GetAccepted returns the Accepted field if non-nil, zero value otherwise.

### GetAcceptedOk

`func (o *EndUserAgreement) GetAcceptedOk() (*time.Time, bool)`

GetAcceptedOk returns a tuple with the Accepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepted

`func (o *EndUserAgreement) SetAccepted(v time.Time)`

SetAccepted sets Accepted field to given value.

### HasAccepted

`func (o *EndUserAgreement) HasAccepted() bool`

HasAccepted returns a boolean if a field has been set.

### SetAcceptedNil

`func (o *EndUserAgreement) SetAcceptedNil(b bool)`

 SetAcceptedNil sets the value for Accepted to be an explicit nil

### UnsetAccepted
`func (o *EndUserAgreement) UnsetAccepted()`

UnsetAccepted ensures that no value is present for Accepted, not even an explicit nil
### GetInstitutionId

`func (o *EndUserAgreement) GetInstitutionId() string`

GetInstitutionId returns the InstitutionId field if non-nil, zero value otherwise.

### GetInstitutionIdOk

`func (o *EndUserAgreement) GetInstitutionIdOk() (*string, bool)`

GetInstitutionIdOk returns a tuple with the InstitutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionId

`func (o *EndUserAgreement) SetInstitutionId(v string)`

SetInstitutionId sets InstitutionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


