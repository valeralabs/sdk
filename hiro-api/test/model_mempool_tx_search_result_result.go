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

// MempoolTxSearchResultResult This object carries the search result
type MempoolTxSearchResultResult struct {
	// The id used to search this query.
	EntityId string `json:"entity_id"`
	EntityType string `json:"entity_type"`
	TxData MempoolTxSearchResultResultTxData `json:"tx_data"`
	Metadata *Object `json:"metadata,omitempty"`
}

// NewMempoolTxSearchResultResult instantiates a new MempoolTxSearchResultResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMempoolTxSearchResultResult(entityId string, entityType string, txData MempoolTxSearchResultResultTxData) *MempoolTxSearchResultResult {
	this := MempoolTxSearchResultResult{}
	this.EntityId = entityId
	this.EntityType = entityType
	this.TxData = txData
	return &this
}

// NewMempoolTxSearchResultResultWithDefaults instantiates a new MempoolTxSearchResultResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMempoolTxSearchResultResultWithDefaults() *MempoolTxSearchResultResult {
	this := MempoolTxSearchResultResult{}
	return &this
}

// GetEntityId returns the EntityId field value
func (o *MempoolTxSearchResultResult) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *MempoolTxSearchResultResult) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *MempoolTxSearchResultResult) SetEntityId(v string) {
	o.EntityId = v
}

// GetEntityType returns the EntityType field value
func (o *MempoolTxSearchResultResult) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *MempoolTxSearchResultResult) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *MempoolTxSearchResultResult) SetEntityType(v string) {
	o.EntityType = v
}

// GetTxData returns the TxData field value
func (o *MempoolTxSearchResultResult) GetTxData() MempoolTxSearchResultResultTxData {
	if o == nil {
		var ret MempoolTxSearchResultResultTxData
		return ret
	}

	return o.TxData
}

// GetTxDataOk returns a tuple with the TxData field value
// and a boolean to check if the value has been set.
func (o *MempoolTxSearchResultResult) GetTxDataOk() (*MempoolTxSearchResultResultTxData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxData, true
}

// SetTxData sets field value
func (o *MempoolTxSearchResultResult) SetTxData(v MempoolTxSearchResultResultTxData) {
	o.TxData = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *MempoolTxSearchResultResult) GetMetadata() Object {
	if o == nil || o.Metadata == nil {
		var ret Object
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MempoolTxSearchResultResult) GetMetadataOk() (*Object, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *MempoolTxSearchResultResult) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given Object and assigns it to the Metadata field.
func (o *MempoolTxSearchResultResult) SetMetadata(v Object) {
	o.Metadata = &v
}

func (o MempoolTxSearchResultResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entity_id"] = o.EntityId
	}
	if true {
		toSerialize["entity_type"] = o.EntityType
	}
	if true {
		toSerialize["tx_data"] = o.TxData
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableMempoolTxSearchResultResult struct {
	value *MempoolTxSearchResultResult
	isSet bool
}

func (v NullableMempoolTxSearchResultResult) Get() *MempoolTxSearchResultResult {
	return v.value
}

func (v *NullableMempoolTxSearchResultResult) Set(val *MempoolTxSearchResultResult) {
	v.value = val
	v.isSet = true
}

func (v NullableMempoolTxSearchResultResult) IsSet() bool {
	return v.isSet
}

func (v *NullableMempoolTxSearchResultResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMempoolTxSearchResultResult(val *MempoolTxSearchResultResult) *NullableMempoolTxSearchResultResult {
	return &NullableMempoolTxSearchResultResult{value: val, isSet: true}
}

func (v NullableMempoolTxSearchResultResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMempoolTxSearchResultResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


