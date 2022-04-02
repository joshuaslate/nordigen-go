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

// SpectacularJWTRefresh Refresh Access token.
type SpectacularJWTRefresh struct {
	// Your access token
	Access *string `json:"access,omitempty"`
	// Access token expires in seconds
	AccessExpires *int32 `json:"access_expires,omitempty"`
}

// NewSpectacularJWTRefresh instantiates a new SpectacularJWTRefresh object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpectacularJWTRefresh() *SpectacularJWTRefresh {
	this := SpectacularJWTRefresh{}
	return &this
}

// NewSpectacularJWTRefreshWithDefaults instantiates a new SpectacularJWTRefresh object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpectacularJWTRefreshWithDefaults() *SpectacularJWTRefresh {
	this := SpectacularJWTRefresh{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *SpectacularJWTRefresh) GetAccess() string {
	if o == nil || o.Access == nil {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpectacularJWTRefresh) GetAccessOk() (*string, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *SpectacularJWTRefresh) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *SpectacularJWTRefresh) SetAccess(v string) {
	o.Access = &v
}

// GetAccessExpires returns the AccessExpires field value if set, zero value otherwise.
func (o *SpectacularJWTRefresh) GetAccessExpires() int32 {
	if o == nil || o.AccessExpires == nil {
		var ret int32
		return ret
	}
	return *o.AccessExpires
}

// GetAccessExpiresOk returns a tuple with the AccessExpires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpectacularJWTRefresh) GetAccessExpiresOk() (*int32, bool) {
	if o == nil || o.AccessExpires == nil {
		return nil, false
	}
	return o.AccessExpires, true
}

// HasAccessExpires returns a boolean if a field has been set.
func (o *SpectacularJWTRefresh) HasAccessExpires() bool {
	if o != nil && o.AccessExpires != nil {
		return true
	}

	return false
}

// SetAccessExpires gets a reference to the given int32 and assigns it to the AccessExpires field.
func (o *SpectacularJWTRefresh) SetAccessExpires(v int32) {
	o.AccessExpires = &v
}

func (o SpectacularJWTRefresh) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	if o.AccessExpires != nil {
		toSerialize["access_expires"] = o.AccessExpires
	}
	return json.Marshal(toSerialize)
}

type NullableSpectacularJWTRefresh struct {
	value *SpectacularJWTRefresh
	isSet bool
}

func (v NullableSpectacularJWTRefresh) Get() *SpectacularJWTRefresh {
	return v.value
}

func (v *NullableSpectacularJWTRefresh) Set(val *SpectacularJWTRefresh) {
	v.value = val
	v.isSet = true
}

func (v NullableSpectacularJWTRefresh) IsSet() bool {
	return v.isSet
}

func (v *NullableSpectacularJWTRefresh) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpectacularJWTRefresh(val *SpectacularJWTRefresh) *NullableSpectacularJWTRefresh {
	return &NullableSpectacularJWTRefresh{value: val, isSet: true}
}

func (v NullableSpectacularJWTRefresh) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpectacularJWTRefresh) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


