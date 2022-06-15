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

// OtherTransactionIdentifier The transaction_identifier uniquely identifies a transaction in a particular network and block or in the mempool.
type OtherTransactionIdentifier struct {
	// Any transactions that are attributable only to a block (ex: a block event) should use the hash of the block as the identifier.
	Hash string `json:"hash"`
}

// NewOtherTransactionIdentifier instantiates a new OtherTransactionIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOtherTransactionIdentifier(hash string) *OtherTransactionIdentifier {
	this := OtherTransactionIdentifier{}
	this.Hash = hash
	return &this
}

// NewOtherTransactionIdentifierWithDefaults instantiates a new OtherTransactionIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOtherTransactionIdentifierWithDefaults() *OtherTransactionIdentifier {
	this := OtherTransactionIdentifier{}
	return &this
}

// GetHash returns the Hash field value
func (o *OtherTransactionIdentifier) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *OtherTransactionIdentifier) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *OtherTransactionIdentifier) SetHash(v string) {
	o.Hash = v
}

func (o OtherTransactionIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hash"] = o.Hash
	}
	return json.Marshal(toSerialize)
}

type NullableOtherTransactionIdentifier struct {
	value *OtherTransactionIdentifier
	isSet bool
}

func (v NullableOtherTransactionIdentifier) Get() *OtherTransactionIdentifier {
	return v.value
}

func (v *NullableOtherTransactionIdentifier) Set(val *OtherTransactionIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableOtherTransactionIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableOtherTransactionIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOtherTransactionIdentifier(val *OtherTransactionIdentifier) *NullableOtherTransactionIdentifier {
	return &NullableOtherTransactionIdentifier{value: val, isSet: true}
}

func (v NullableOtherTransactionIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOtherTransactionIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


