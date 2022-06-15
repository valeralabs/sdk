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

// MempoolTxSearchResultResultTxData Returns basic search result information about the requested id
type MempoolTxSearchResultResultTxData struct {
	TxType string `json:"tx_type"`
}

// NewMempoolTxSearchResultResultTxData instantiates a new MempoolTxSearchResultResultTxData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMempoolTxSearchResultResultTxData(txType string) *MempoolTxSearchResultResultTxData {
	this := MempoolTxSearchResultResultTxData{}
	this.TxType = txType
	return &this
}

// NewMempoolTxSearchResultResultTxDataWithDefaults instantiates a new MempoolTxSearchResultResultTxData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMempoolTxSearchResultResultTxDataWithDefaults() *MempoolTxSearchResultResultTxData {
	this := MempoolTxSearchResultResultTxData{}
	return &this
}

// GetTxType returns the TxType field value
func (o *MempoolTxSearchResultResultTxData) GetTxType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxType
}

// GetTxTypeOk returns a tuple with the TxType field value
// and a boolean to check if the value has been set.
func (o *MempoolTxSearchResultResultTxData) GetTxTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxType, true
}

// SetTxType sets field value
func (o *MempoolTxSearchResultResultTxData) SetTxType(v string) {
	o.TxType = v
}

func (o MempoolTxSearchResultResultTxData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tx_type"] = o.TxType
	}
	return json.Marshal(toSerialize)
}

type NullableMempoolTxSearchResultResultTxData struct {
	value *MempoolTxSearchResultResultTxData
	isSet bool
}

func (v NullableMempoolTxSearchResultResultTxData) Get() *MempoolTxSearchResultResultTxData {
	return v.value
}

func (v *NullableMempoolTxSearchResultResultTxData) Set(val *MempoolTxSearchResultResultTxData) {
	v.value = val
	v.isSet = true
}

func (v NullableMempoolTxSearchResultResultTxData) IsSet() bool {
	return v.isSet
}

func (v *NullableMempoolTxSearchResultResultTxData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMempoolTxSearchResultResultTxData(val *MempoolTxSearchResultResultTxData) *NullableMempoolTxSearchResultResultTxData {
	return &NullableMempoolTxSearchResultResultTxData{value: val, isSet: true}
}

func (v NullableMempoolTxSearchResultResultTxData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMempoolTxSearchResultResultTxData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


