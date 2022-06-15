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

// RosettaAmount Amount is some Value of a Currency. It is considered invalid to specify a Value without a Currency.
type RosettaAmount struct {
	// Value of the transaction in atomic units represented as an arbitrary-sized signed integer. For example, 1 BTC would be represented by a value of 100000000.
	Value string `json:"value"`
	Currency Object `json:"currency"`
	// 
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewRosettaAmount instantiates a new RosettaAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaAmount(value string, currency Object) *RosettaAmount {
	this := RosettaAmount{}
	this.Value = value
	this.Currency = currency
	return &this
}

// NewRosettaAmountWithDefaults instantiates a new RosettaAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaAmountWithDefaults() *RosettaAmount {
	this := RosettaAmount{}
	return &this
}

// GetValue returns the Value field value
func (o *RosettaAmount) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *RosettaAmount) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *RosettaAmount) SetValue(v string) {
	o.Value = v
}

// GetCurrency returns the Currency field value
func (o *RosettaAmount) GetCurrency() Object {
	if o == nil {
		var ret Object
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *RosettaAmount) GetCurrencyOk() (*Object, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *RosettaAmount) SetCurrency(v Object) {
	o.Currency = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RosettaAmount) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaAmount) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RosettaAmount) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *RosettaAmount) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o RosettaAmount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableRosettaAmount struct {
	value *RosettaAmount
	isSet bool
}

func (v NullableRosettaAmount) Get() *RosettaAmount {
	return v.value
}

func (v *NullableRosettaAmount) Set(val *RosettaAmount) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaAmount(val *RosettaAmount) *NullableRosettaAmount {
	return &NullableRosettaAmount{value: val, isSet: true}
}

func (v NullableRosettaAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


