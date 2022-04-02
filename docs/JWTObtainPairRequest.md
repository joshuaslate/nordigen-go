# JWTObtainPairRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretId** | **string** | Secret id from /user-secrets/ | 
**SecretKey** | **string** | Secret key from /user-secrets/ | 

## Methods

### NewJWTObtainPairRequest

`func NewJWTObtainPairRequest(secretId string, secretKey string, ) *JWTObtainPairRequest`

NewJWTObtainPairRequest instantiates a new JWTObtainPairRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJWTObtainPairRequestWithDefaults

`func NewJWTObtainPairRequestWithDefaults() *JWTObtainPairRequest`

NewJWTObtainPairRequestWithDefaults instantiates a new JWTObtainPairRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecretId

`func (o *JWTObtainPairRequest) GetSecretId() string`

GetSecretId returns the SecretId field if non-nil, zero value otherwise.

### GetSecretIdOk

`func (o *JWTObtainPairRequest) GetSecretIdOk() (*string, bool)`

GetSecretIdOk returns a tuple with the SecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretId

`func (o *JWTObtainPairRequest) SetSecretId(v string)`

SetSecretId sets SecretId field to given value.


### GetSecretKey

`func (o *JWTObtainPairRequest) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *JWTObtainPairRequest) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *JWTObtainPairRequest) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


