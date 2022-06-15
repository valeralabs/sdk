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

// TransactionEventNonFungibleAssetAllOf struct for TransactionEventNonFungibleAssetAllOf
type TransactionEventNonFungibleAssetAllOf struct {
	EventType string `json:"event_type"`
	TxId string `json:"tx_id"`
	Asset TransactionEventNonFungibleAssetAllOfAsset `json:"asset"`
}

// NewTransactionEventNonFungibleAssetAllOf instantiates a new TransactionEventNonFungibleAssetAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionEventNonFungibleAssetAllOf(eventType string, txId string, asset TransactionEventNonFungibleAssetAllOfAsset) *TransactionEventNonFungibleAssetAllOf {
	this := TransactionEventNonFungibleAssetAllOf{}
	this.EventType = eventType
	this.TxId = txId
	this.Asset = asset
	return &this
}

// NewTransactionEventNonFungibleAssetAllOfWithDefaults instantiates a new TransactionEventNonFungibleAssetAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionEventNonFungibleAssetAllOfWithDefaults() *TransactionEventNonFungibleAssetAllOf {
	this := TransactionEventNonFungibleAssetAllOf{}
	return &this
}

// GetEventType returns the EventType field value
func (o *TransactionEventNonFungibleAssetAllOf) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *TransactionEventNonFungibleAssetAllOf) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *TransactionEventNonFungibleAssetAllOf) SetEventType(v string) {
	o.EventType = v
}

// GetTxId returns the TxId field value
func (o *TransactionEventNonFungibleAssetAllOf) GetTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *TransactionEventNonFungibleAssetAllOf) GetTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *TransactionEventNonFungibleAssetAllOf) SetTxId(v string) {
	o.TxId = v
}

// GetAsset returns the Asset field value
func (o *TransactionEventNonFungibleAssetAllOf) GetAsset() TransactionEventNonFungibleAssetAllOfAsset {
	if o == nil {
		var ret TransactionEventNonFungibleAssetAllOfAsset
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *TransactionEventNonFungibleAssetAllOf) GetAssetOk() (*TransactionEventNonFungibleAssetAllOfAsset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *TransactionEventNonFungibleAssetAllOf) SetAsset(v TransactionEventNonFungibleAssetAllOfAsset) {
	o.Asset = v
}

func (o TransactionEventNonFungibleAssetAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["event_type"] = o.EventType
	}
	if true {
		toSerialize["tx_id"] = o.TxId
	}
	if true {
		toSerialize["asset"] = o.Asset
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionEventNonFungibleAssetAllOf struct {
	value *TransactionEventNonFungibleAssetAllOf
	isSet bool
}

func (v NullableTransactionEventNonFungibleAssetAllOf) Get() *TransactionEventNonFungibleAssetAllOf {
	return v.value
}

func (v *NullableTransactionEventNonFungibleAssetAllOf) Set(val *TransactionEventNonFungibleAssetAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionEventNonFungibleAssetAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionEventNonFungibleAssetAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionEventNonFungibleAssetAllOf(val *TransactionEventNonFungibleAssetAllOf) *NullableTransactionEventNonFungibleAssetAllOf {
	return &NullableTransactionEventNonFungibleAssetAllOf{value: val, isSet: true}
}

func (v NullableTransactionEventNonFungibleAssetAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionEventNonFungibleAssetAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


