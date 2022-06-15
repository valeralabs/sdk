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

// AddressNonces The latest nonce values used by an account by inspecting the mempool, microblock transactions, and anchored transactions
type AddressNonces struct {
	// The latest nonce found within mempool transactions sent by this address. Will be null if there are no current mempool transactions for this address.
	LastMempoolTxNonce NullableInt32 `json:"last_mempool_tx_nonce"`
	// The latest nonce found within transactions sent by this address, including unanchored microblock transactions. Will be null if there are no current transactions for this address.
	LastExecutedTxNonce NullableInt32 `json:"last_executed_tx_nonce"`
	// The likely nonce required for creating the next transaction, based on the last nonces seen by the API. This can be incorrect if the API's mempool or transactions aren't fully synchronized, even by a small amount, or if a previous transaction is still propagating through the Stacks blockchain network when this endpoint is called.
	PossibleNextNonce int32 `json:"possible_next_nonce"`
	// Nonces that appear to be missing and likely causing a mempool transaction to be stuck.
	DetectedMissingNonces []int32 `json:"detected_missing_nonces"`
}

// NewAddressNonces instantiates a new AddressNonces object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressNonces(lastMempoolTxNonce NullableInt32, lastExecutedTxNonce NullableInt32, possibleNextNonce int32, detectedMissingNonces []int32) *AddressNonces {
	this := AddressNonces{}
	this.LastMempoolTxNonce = lastMempoolTxNonce
	this.LastExecutedTxNonce = lastExecutedTxNonce
	this.PossibleNextNonce = possibleNextNonce
	this.DetectedMissingNonces = detectedMissingNonces
	return &this
}

// NewAddressNoncesWithDefaults instantiates a new AddressNonces object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressNoncesWithDefaults() *AddressNonces {
	this := AddressNonces{}
	return &this
}

// GetLastMempoolTxNonce returns the LastMempoolTxNonce field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *AddressNonces) GetLastMempoolTxNonce() int32 {
	if o == nil || o.LastMempoolTxNonce.Get() == nil {
		var ret int32
		return ret
	}

	return *o.LastMempoolTxNonce.Get()
}

// GetLastMempoolTxNonceOk returns a tuple with the LastMempoolTxNonce field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressNonces) GetLastMempoolTxNonceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastMempoolTxNonce.Get(), o.LastMempoolTxNonce.IsSet()
}

// SetLastMempoolTxNonce sets field value
func (o *AddressNonces) SetLastMempoolTxNonce(v int32) {
	o.LastMempoolTxNonce.Set(&v)
}

// GetLastExecutedTxNonce returns the LastExecutedTxNonce field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *AddressNonces) GetLastExecutedTxNonce() int32 {
	if o == nil || o.LastExecutedTxNonce.Get() == nil {
		var ret int32
		return ret
	}

	return *o.LastExecutedTxNonce.Get()
}

// GetLastExecutedTxNonceOk returns a tuple with the LastExecutedTxNonce field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressNonces) GetLastExecutedTxNonceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastExecutedTxNonce.Get(), o.LastExecutedTxNonce.IsSet()
}

// SetLastExecutedTxNonce sets field value
func (o *AddressNonces) SetLastExecutedTxNonce(v int32) {
	o.LastExecutedTxNonce.Set(&v)
}

// GetPossibleNextNonce returns the PossibleNextNonce field value
func (o *AddressNonces) GetPossibleNextNonce() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PossibleNextNonce
}

// GetPossibleNextNonceOk returns a tuple with the PossibleNextNonce field value
// and a boolean to check if the value has been set.
func (o *AddressNonces) GetPossibleNextNonceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PossibleNextNonce, true
}

// SetPossibleNextNonce sets field value
func (o *AddressNonces) SetPossibleNextNonce(v int32) {
	o.PossibleNextNonce = v
}

// GetDetectedMissingNonces returns the DetectedMissingNonces field value
func (o *AddressNonces) GetDetectedMissingNonces() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.DetectedMissingNonces
}

// GetDetectedMissingNoncesOk returns a tuple with the DetectedMissingNonces field value
// and a boolean to check if the value has been set.
func (o *AddressNonces) GetDetectedMissingNoncesOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DetectedMissingNonces, true
}

// SetDetectedMissingNonces sets field value
func (o *AddressNonces) SetDetectedMissingNonces(v []int32) {
	o.DetectedMissingNonces = v
}

func (o AddressNonces) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["last_mempool_tx_nonce"] = o.LastMempoolTxNonce.Get()
	}
	if true {
		toSerialize["last_executed_tx_nonce"] = o.LastExecutedTxNonce.Get()
	}
	if true {
		toSerialize["possible_next_nonce"] = o.PossibleNextNonce
	}
	if true {
		toSerialize["detected_missing_nonces"] = o.DetectedMissingNonces
	}
	return json.Marshal(toSerialize)
}

type NullableAddressNonces struct {
	value *AddressNonces
	isSet bool
}

func (v NullableAddressNonces) Get() *AddressNonces {
	return v.value
}

func (v *NullableAddressNonces) Set(val *AddressNonces) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressNonces) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressNonces) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressNonces(val *AddressNonces) *NullableAddressNonces {
	return &NullableAddressNonces{value: val, isSet: true}
}

func (v NullableAddressNonces) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressNonces) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


