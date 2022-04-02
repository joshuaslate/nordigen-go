# PremiumAccountQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateFrom** | Pointer to **string** |  | [optional] 
**DateTo** | Pointer to **string** |  | [optional] 
**Country** | Pointer to [**CountryEnum**](CountryEnum.md) |  | [optional] 

## Methods

### NewPremiumAccountQuery

`func NewPremiumAccountQuery() *PremiumAccountQuery`

NewPremiumAccountQuery instantiates a new PremiumAccountQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPremiumAccountQueryWithDefaults

`func NewPremiumAccountQueryWithDefaults() *PremiumAccountQuery`

NewPremiumAccountQueryWithDefaults instantiates a new PremiumAccountQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateFrom

`func (o *PremiumAccountQuery) GetDateFrom() string`

GetDateFrom returns the DateFrom field if non-nil, zero value otherwise.

### GetDateFromOk

`func (o *PremiumAccountQuery) GetDateFromOk() (*string, bool)`

GetDateFromOk returns a tuple with the DateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFrom

`func (o *PremiumAccountQuery) SetDateFrom(v string)`

SetDateFrom sets DateFrom field to given value.

### HasDateFrom

`func (o *PremiumAccountQuery) HasDateFrom() bool`

HasDateFrom returns a boolean if a field has been set.

### GetDateTo

`func (o *PremiumAccountQuery) GetDateTo() string`

GetDateTo returns the DateTo field if non-nil, zero value otherwise.

### GetDateToOk

`func (o *PremiumAccountQuery) GetDateToOk() (*string, bool)`

GetDateToOk returns a tuple with the DateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTo

`func (o *PremiumAccountQuery) SetDateTo(v string)`

SetDateTo sets DateTo field to given value.

### HasDateTo

`func (o *PremiumAccountQuery) HasDateTo() bool`

HasDateTo returns a boolean if a field has been set.

### GetCountry

`func (o *PremiumAccountQuery) GetCountry() CountryEnum`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PremiumAccountQuery) GetCountryOk() (*CountryEnum, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PremiumAccountQuery) SetCountry(v CountryEnum)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PremiumAccountQuery) HasCountry() bool`

HasCountry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


