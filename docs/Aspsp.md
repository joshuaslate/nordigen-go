# Aspsp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Bic** | Pointer to **string** |  | [optional] 
**TransactionTotalDays** | Pointer to **string** |  | [optional] [default to "90"]
**Countries** | **[]string** |  | 
**Logo** | **string** |  | 

## Methods

### NewAspsp

`func NewAspsp(id string, name string, countries []string, logo string, ) *Aspsp`

NewAspsp instantiates a new Aspsp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAspspWithDefaults

`func NewAspspWithDefaults() *Aspsp`

NewAspspWithDefaults instantiates a new Aspsp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Aspsp) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Aspsp) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Aspsp) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Aspsp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Aspsp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Aspsp) SetName(v string)`

SetName sets Name field to given value.


### GetBic

`func (o *Aspsp) GetBic() string`

GetBic returns the Bic field if non-nil, zero value otherwise.

### GetBicOk

`func (o *Aspsp) GetBicOk() (*string, bool)`

GetBicOk returns a tuple with the Bic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBic

`func (o *Aspsp) SetBic(v string)`

SetBic sets Bic field to given value.

### HasBic

`func (o *Aspsp) HasBic() bool`

HasBic returns a boolean if a field has been set.

### GetTransactionTotalDays

`func (o *Aspsp) GetTransactionTotalDays() string`

GetTransactionTotalDays returns the TransactionTotalDays field if non-nil, zero value otherwise.

### GetTransactionTotalDaysOk

`func (o *Aspsp) GetTransactionTotalDaysOk() (*string, bool)`

GetTransactionTotalDaysOk returns a tuple with the TransactionTotalDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTotalDays

`func (o *Aspsp) SetTransactionTotalDays(v string)`

SetTransactionTotalDays sets TransactionTotalDays field to given value.

### HasTransactionTotalDays

`func (o *Aspsp) HasTransactionTotalDays() bool`

HasTransactionTotalDays returns a boolean if a field has been set.

### GetCountries

`func (o *Aspsp) GetCountries() []string`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *Aspsp) GetCountriesOk() (*[]string, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *Aspsp) SetCountries(v []string)`

SetCountries sets Countries field to given value.


### GetLogo

`func (o *Aspsp) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *Aspsp) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *Aspsp) SetLogo(v string)`

SetLogo sets Logo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


