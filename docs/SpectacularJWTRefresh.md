# SpectacularJWTRefresh

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to **string** | Your access token | [optional] [readonly] 
**AccessExpires** | Pointer to **int32** | Access token expires in seconds | [optional] [readonly] [default to 86400]

## Methods

### NewSpectacularJWTRefresh

`func NewSpectacularJWTRefresh() *SpectacularJWTRefresh`

NewSpectacularJWTRefresh instantiates a new SpectacularJWTRefresh object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpectacularJWTRefreshWithDefaults

`func NewSpectacularJWTRefreshWithDefaults() *SpectacularJWTRefresh`

NewSpectacularJWTRefreshWithDefaults instantiates a new SpectacularJWTRefresh object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *SpectacularJWTRefresh) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *SpectacularJWTRefresh) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *SpectacularJWTRefresh) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *SpectacularJWTRefresh) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetAccessExpires

`func (o *SpectacularJWTRefresh) GetAccessExpires() int32`

GetAccessExpires returns the AccessExpires field if non-nil, zero value otherwise.

### GetAccessExpiresOk

`func (o *SpectacularJWTRefresh) GetAccessExpiresOk() (*int32, bool)`

GetAccessExpiresOk returns a tuple with the AccessExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpires

`func (o *SpectacularJWTRefresh) SetAccessExpires(v int32)`

SetAccessExpires sets AccessExpires field to given value.

### HasAccessExpires

`func (o *SpectacularJWTRefresh) HasAccessExpires() bool`

HasAccessExpires returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


