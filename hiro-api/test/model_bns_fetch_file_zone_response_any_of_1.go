/*
Stacks Blockchain API

Welcome to the API reference overview for the <a href=\"https://docs.hiro.so/get-started/stacks-blockchain-api\">Stacks Blockchain API</a>.  <a href=\"https://hirosystems.github.io/stacks-blockchain-api/collection.json\" download=\"stacks-api-collection.json\">Download Postman collection</a> 

API version: STACKS_API_VERSION
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// BnsFetchFileZoneResponseAnyOf1 struct for BnsFetchFileZoneResponseAnyOf1
type BnsFetchFileZoneResponseAnyOf1 struct {
	Error *string `json:"error,omitempty"`
}

// NewBnsFetchFileZoneResponseAnyOf1 instantiates a new BnsFetchFileZoneResponseAnyOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBnsFetchFileZoneResponseAnyOf1() *BnsFetchFileZoneResponseAnyOf1 {
	this := BnsFetchFileZoneResponseAnyOf1{}
	return &this
}

// NewBnsFetchFileZoneResponseAnyOf1WithDefaults instantiates a new BnsFetchFileZoneResponseAnyOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBnsFetchFileZoneResponseAnyOf1WithDefaults() *BnsFetchFileZoneResponseAnyOf1 {
	this := BnsFetchFileZoneResponseAnyOf1{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *BnsFetchFileZoneResponseAnyOf1) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BnsFetchFileZoneResponseAnyOf1) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *BnsFetchFileZoneResponseAnyOf1) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *BnsFetchFileZoneResponseAnyOf1) SetError(v string) {
	o.Error = &v
}

func (o BnsFetchFileZoneResponseAnyOf1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableBnsFetchFileZoneResponseAnyOf1 struct {
	value *BnsFetchFileZoneResponseAnyOf1
	isSet bool
}

func (v NullableBnsFetchFileZoneResponseAnyOf1) Get() *BnsFetchFileZoneResponseAnyOf1 {
	return v.value
}

func (v *NullableBnsFetchFileZoneResponseAnyOf1) Set(val *BnsFetchFileZoneResponseAnyOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableBnsFetchFileZoneResponseAnyOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableBnsFetchFileZoneResponseAnyOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBnsFetchFileZoneResponseAnyOf1(val *BnsFetchFileZoneResponseAnyOf1) *NullableBnsFetchFileZoneResponseAnyOf1 {
	return &NullableBnsFetchFileZoneResponseAnyOf1{value: val, isSet: true}
}

func (v NullableBnsFetchFileZoneResponseAnyOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBnsFetchFileZoneResponseAnyOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


