# AccountV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of this Account, used to refer to this account in other API calls. | [optional] [readonly] 
**Created** | Pointer to **time.Time** | The date &amp; time at which the account object was created. | [optional] [readonly] 
**LastAccessed** | Pointer to **NullableTime** | The date &amp; time at which the account object was last accessed. | [optional] [readonly] 
**Iban** | Pointer to **string** | The Account IBAN | [optional] [readonly] 
**InstitutionId** | Pointer to **string** | The ASPSP associated with this account. | [optional] [readonly] 
**Status** | Pointer to [**AccountV2StatusEnum**](AccountV2StatusEnum.md) |  | [optional] 

## Methods

### NewAccountV2

`func NewAccountV2() *AccountV2`

NewAccountV2 instantiates a new AccountV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountV2WithDefaults

`func NewAccountV2WithDefaults() *AccountV2`

NewAccountV2WithDefaults instantiates a new AccountV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreated

`func (o *AccountV2) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AccountV2) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AccountV2) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AccountV2) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastAccessed

`func (o *AccountV2) GetLastAccessed() time.Time`

GetLastAccessed returns the LastAccessed field if non-nil, zero value otherwise.

### GetLastAccessedOk

`func (o *AccountV2) GetLastAccessedOk() (*time.Time, bool)`

GetLastAccessedOk returns a tuple with the LastAccessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessed

`func (o *AccountV2) SetLastAccessed(v time.Time)`

SetLastAccessed sets LastAccessed field to given value.

### HasLastAccessed

`func (o *AccountV2) HasLastAccessed() bool`

HasLastAccessed returns a boolean if a field has been set.

### SetLastAccessedNil

`func (o *AccountV2) SetLastAccessedNil(b bool)`

 SetLastAccessedNil sets the value for LastAccessed to be an explicit nil

### UnsetLastAccessed
`func (o *AccountV2) UnsetLastAccessed()`

UnsetLastAccessed ensures that no value is present for LastAccessed, not even an explicit nil
### GetIban

`func (o *AccountV2) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *AccountV2) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *AccountV2) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *AccountV2) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetInstitutionId

`func (o *AccountV2) GetInstitutionId() string`

GetInstitutionId returns the InstitutionId field if non-nil, zero value otherwise.

### GetInstitutionIdOk

`func (o *AccountV2) GetInstitutionIdOk() (*string, bool)`

GetInstitutionIdOk returns a tuple with the InstitutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionId

`func (o *AccountV2) SetInstitutionId(v string)`

SetInstitutionId sets InstitutionId field to given value.

### HasInstitutionId

`func (o *AccountV2) HasInstitutionId() bool`

HasInstitutionId returns a boolean if a field has been set.

### GetStatus

`func (o *AccountV2) GetStatus() AccountV2StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountV2) GetStatusOk() (*AccountV2StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountV2) SetStatus(v AccountV2StatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountV2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


