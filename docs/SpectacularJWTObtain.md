# SpectacularJWTObtain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to **string** | Your access token | [optional] [readonly] 
**AccessExpires** | Pointer to **int32** | Access token expires in seconds | [optional] [readonly] [default to 86400]
**Refresh** | Pointer to **string** | Your refresh token | [optional] [readonly] 
**RefreshExpires** | Pointer to **int32** | Refresh token expires in seconds | [optional] [readonly] [default to 2592000]

## Methods

### NewSpectacularJWTObtain

`func NewSpectacularJWTObtain() *SpectacularJWTObtain`

NewSpectacularJWTObtain instantiates a new SpectacularJWTObtain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpectacularJWTObtainWithDefaults

`func NewSpectacularJWTObtainWithDefaults() *SpectacularJWTObtain`

NewSpectacularJWTObtainWithDefaults instantiates a new SpectacularJWTObtain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *SpectacularJWTObtain) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *SpectacularJWTObtain) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *SpectacularJWTObtain) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *SpectacularJWTObtain) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetAccessExpires

`func (o *SpectacularJWTObtain) GetAccessExpires() int32`

GetAccessExpires returns the AccessExpires field if non-nil, zero value otherwise.

### GetAccessExpiresOk

`func (o *SpectacularJWTObtain) GetAccessExpiresOk() (*int32, bool)`

GetAccessExpiresOk returns a tuple with the AccessExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpires

`func (o *SpectacularJWTObtain) SetAccessExpires(v int32)`

SetAccessExpires sets AccessExpires field to given value.

### HasAccessExpires

`func (o *SpectacularJWTObtain) HasAccessExpires() bool`

HasAccessExpires returns a boolean if a field has been set.

### GetRefresh

`func (o *SpectacularJWTObtain) GetRefresh() string`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *SpectacularJWTObtain) GetRefreshOk() (*string, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *SpectacularJWTObtain) SetRefresh(v string)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *SpectacularJWTObtain) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### GetRefreshExpires

`func (o *SpectacularJWTObtain) GetRefreshExpires() int32`

GetRefreshExpires returns the RefreshExpires field if non-nil, zero value otherwise.

### GetRefreshExpiresOk

`func (o *SpectacularJWTObtain) GetRefreshExpiresOk() (*int32, bool)`

GetRefreshExpiresOk returns a tuple with the RefreshExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshExpires

`func (o *SpectacularJWTObtain) SetRefreshExpires(v int32)`

SetRefreshExpires sets RefreshExpires field to given value.

### HasRefreshExpires

`func (o *SpectacularJWTObtain) HasRefreshExpires() bool`

HasRefreshExpires returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


