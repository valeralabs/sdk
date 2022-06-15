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

// TransactionEventStxLockAllOfStxLockEvent struct for TransactionEventStxLockAllOfStxLockEvent
type TransactionEventStxLockAllOfStxLockEvent struct {
	LockedAmount string `json:"locked_amount"`
	UnlockHeight int32 `json:"unlock_height"`
	LockedAddress string `json:"locked_address"`
}

// NewTransactionEventStxLockAllOfStxLockEvent instantiates a new TransactionEventStxLockAllOfStxLockEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionEventStxLockAllOfStxLockEvent(lockedAmount string, unlockHeight int32, lockedAddress string) *TransactionEventStxLockAllOfStxLockEvent {
	this := TransactionEventStxLockAllOfStxLockEvent{}
	this.LockedAmount = lockedAmount
	this.UnlockHeight = unlockHeight
	this.LockedAddress = lockedAddress
	return &this
}

// NewTransactionEventStxLockAllOfStxLockEventWithDefaults instantiates a new TransactionEventStxLockAllOfStxLockEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionEventStxLockAllOfStxLockEventWithDefaults() *TransactionEventStxLockAllOfStxLockEvent {
	this := TransactionEventStxLockAllOfStxLockEvent{}
	return &this
}

// GetLockedAmount returns the LockedAmount field value
func (o *TransactionEventStxLockAllOfStxLockEvent) GetLockedAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LockedAmount
}

// GetLockedAmountOk returns a tuple with the LockedAmount field value
// and a boolean to check if the value has been set.
func (o *TransactionEventStxLockAllOfStxLockEvent) GetLockedAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LockedAmount, true
}

// SetLockedAmount sets field value
func (o *TransactionEventStxLockAllOfStxLockEvent) SetLockedAmount(v string) {
	o.LockedAmount = v
}

// GetUnlockHeight returns the UnlockHeight field value
func (o *TransactionEventStxLockAllOfStxLockEvent) GetUnlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UnlockHeight
}

// GetUnlockHeightOk returns a tuple with the UnlockHeight field value
// and a boolean to check if the value has been set.
func (o *TransactionEventStxLockAllOfStxLockEvent) GetUnlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnlockHeight, true
}

// SetUnlockHeight sets field value
func (o *TransactionEventStxLockAllOfStxLockEvent) SetUnlockHeight(v int32) {
	o.UnlockHeight = v
}

// GetLockedAddress returns the LockedAddress field value
func (o *TransactionEventStxLockAllOfStxLockEvent) GetLockedAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LockedAddress
}

// GetLockedAddressOk returns a tuple with the LockedAddress field value
// and a boolean to check if the value has been set.
func (o *TransactionEventStxLockAllOfStxLockEvent) GetLockedAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LockedAddress, true
}

// SetLockedAddress sets field value
func (o *TransactionEventStxLockAllOfStxLockEvent) SetLockedAddress(v string) {
	o.LockedAddress = v
}

func (o TransactionEventStxLockAllOfStxLockEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["locked_amount"] = o.LockedAmount
	}
	if true {
		toSerialize["unlock_height"] = o.UnlockHeight
	}
	if true {
		toSerialize["locked_address"] = o.LockedAddress
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionEventStxLockAllOfStxLockEvent struct {
	value *TransactionEventStxLockAllOfStxLockEvent
	isSet bool
}

func (v NullableTransactionEventStxLockAllOfStxLockEvent) Get() *TransactionEventStxLockAllOfStxLockEvent {
	return v.value
}

func (v *NullableTransactionEventStxLockAllOfStxLockEvent) Set(val *TransactionEventStxLockAllOfStxLockEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionEventStxLockAllOfStxLockEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionEventStxLockAllOfStxLockEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionEventStxLockAllOfStxLockEvent(val *TransactionEventStxLockAllOfStxLockEvent) *NullableTransactionEventStxLockAllOfStxLockEvent {
	return &NullableTransactionEventStxLockAllOfStxLockEvent{value: val, isSet: true}
}

func (v NullableTransactionEventStxLockAllOfStxLockEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionEventStxLockAllOfStxLockEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


