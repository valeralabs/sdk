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

// TokenTransferTransactionMetadataTokenTransfer struct for TokenTransferTransactionMetadataTokenTransfer
type TokenTransferTransactionMetadataTokenTransfer struct {
	RecipientAddress string `json:"recipient_address"`
	// Transfer amount as Integer string (64-bit unsigned integer)
	Amount string `json:"amount"`
	// Hex encoded arbitrary message, up to 34 bytes length (should try decoding to an ASCII string)
	Memo string `json:"memo"`
}

// NewTokenTransferTransactionMetadataTokenTransfer instantiates a new TokenTransferTransactionMetadataTokenTransfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenTransferTransactionMetadataTokenTransfer(recipientAddress string, amount string, memo string) *TokenTransferTransactionMetadataTokenTransfer {
	this := TokenTransferTransactionMetadataTokenTransfer{}
	this.RecipientAddress = recipientAddress
	this.Amount = amount
	this.Memo = memo
	return &this
}

// NewTokenTransferTransactionMetadataTokenTransferWithDefaults instantiates a new TokenTransferTransactionMetadataTokenTransfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenTransferTransactionMetadataTokenTransferWithDefaults() *TokenTransferTransactionMetadataTokenTransfer {
	this := TokenTransferTransactionMetadataTokenTransfer{}
	return &this
}

// GetRecipientAddress returns the RecipientAddress field value
func (o *TokenTransferTransactionMetadataTokenTransfer) GetRecipientAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecipientAddress
}

// GetRecipientAddressOk returns a tuple with the RecipientAddress field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransactionMetadataTokenTransfer) GetRecipientAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecipientAddress, true
}

// SetRecipientAddress sets field value
func (o *TokenTransferTransactionMetadataTokenTransfer) SetRecipientAddress(v string) {
	o.RecipientAddress = v
}

// GetAmount returns the Amount field value
func (o *TokenTransferTransactionMetadataTokenTransfer) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransactionMetadataTokenTransfer) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TokenTransferTransactionMetadataTokenTransfer) SetAmount(v string) {
	o.Amount = v
}

// GetMemo returns the Memo field value
func (o *TokenTransferTransactionMetadataTokenTransfer) GetMemo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Memo
}

// GetMemoOk returns a tuple with the Memo field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransactionMetadataTokenTransfer) GetMemoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Memo, true
}

// SetMemo sets field value
func (o *TokenTransferTransactionMetadataTokenTransfer) SetMemo(v string) {
	o.Memo = v
}

func (o TokenTransferTransactionMetadataTokenTransfer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["recipient_address"] = o.RecipientAddress
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["memo"] = o.Memo
	}
	return json.Marshal(toSerialize)
}

type NullableTokenTransferTransactionMetadataTokenTransfer struct {
	value *TokenTransferTransactionMetadataTokenTransfer
	isSet bool
}

func (v NullableTokenTransferTransactionMetadataTokenTransfer) Get() *TokenTransferTransactionMetadataTokenTransfer {
	return v.value
}

func (v *NullableTokenTransferTransactionMetadataTokenTransfer) Set(val *TokenTransferTransactionMetadataTokenTransfer) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenTransferTransactionMetadataTokenTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenTransferTransactionMetadataTokenTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenTransferTransactionMetadataTokenTransfer(val *TokenTransferTransactionMetadataTokenTransfer) *NullableTokenTransferTransactionMetadataTokenTransfer {
	return &NullableTokenTransferTransactionMetadataTokenTransfer{value: val, isSet: true}
}

func (v NullableTokenTransferTransactionMetadataTokenTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenTransferTransactionMetadataTokenTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


