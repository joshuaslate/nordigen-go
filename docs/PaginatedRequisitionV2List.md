# PaginatedRequisitionV2List

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | Pointer to [**[]RequisitionV2**](RequisitionV2.md) |  | [optional] 

## Methods

### NewPaginatedRequisitionV2List

`func NewPaginatedRequisitionV2List() *PaginatedRequisitionV2List`

NewPaginatedRequisitionV2List instantiates a new PaginatedRequisitionV2List object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedRequisitionV2ListWithDefaults

`func NewPaginatedRequisitionV2ListWithDefaults() *PaginatedRequisitionV2List`

NewPaginatedRequisitionV2ListWithDefaults instantiates a new PaginatedRequisitionV2List object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedRequisitionV2List) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedRequisitionV2List) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedRequisitionV2List) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedRequisitionV2List) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNext

`func (o *PaginatedRequisitionV2List) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedRequisitionV2List) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedRequisitionV2List) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedRequisitionV2List) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedRequisitionV2List) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedRequisitionV2List) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedRequisitionV2List) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedRequisitionV2List) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedRequisitionV2List) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedRequisitionV2List) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedRequisitionV2List) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedRequisitionV2List) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedRequisitionV2List) GetResults() []RequisitionV2`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedRequisitionV2List) GetResultsOk() (*[]RequisitionV2, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedRequisitionV2List) SetResults(v []RequisitionV2)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedRequisitionV2List) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


