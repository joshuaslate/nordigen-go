/*
Nordigen Account Information Services API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0 (v2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nordigen

import (
	"encoding/json"
)

// JWTRefreshRequest Refresh access token.
type JWTRefreshRequest struct {
	Refresh string `json:"refresh"`
}

// NewJWTRefreshRequest instantiates a new JWTRefreshRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJWTRefreshRequest(refresh string) *JWTRefreshRequest {
	this := JWTRefreshRequest{}
	this.Refresh = refresh
	return &this
}

// NewJWTRefreshRequestWithDefaults instantiates a new JWTRefreshRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJWTRefreshRequestWithDefaults() *JWTRefreshRequest {
	this := JWTRefreshRequest{}
	return &this
}

// GetRefresh returns the Refresh field value
func (o *JWTRefreshRequest) GetRefresh() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Refresh
}

// GetRefreshOk returns a tuple with the Refresh field value
// and a boolean to check if the value has been set.
func (o *JWTRefreshRequest) GetRefreshOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Refresh, true
}

// SetRefresh sets field value
func (o *JWTRefreshRequest) SetRefresh(v string) {
	o.Refresh = v
}

func (o JWTRefreshRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["refresh"] = o.Refresh
	}
	return json.Marshal(toSerialize)
}

type NullableJWTRefreshRequest struct {
	value *JWTRefreshRequest
	isSet bool
}

func (v NullableJWTRefreshRequest) Get() *JWTRefreshRequest {
	return v.value
}

func (v *NullableJWTRefreshRequest) Set(val *JWTRefreshRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableJWTRefreshRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableJWTRefreshRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJWTRefreshRequest(val *JWTRefreshRequest) *NullableJWTRefreshRequest {
	return &NullableJWTRefreshRequest{value: val, isSet: true}
}

func (v NullableJWTRefreshRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJWTRefreshRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


