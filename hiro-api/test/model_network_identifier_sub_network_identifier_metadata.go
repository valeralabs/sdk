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

// NetworkIdentifierSubNetworkIdentifierMetadata Meta data from subnetwork identifier
type NetworkIdentifierSubNetworkIdentifierMetadata struct {
	// producer
	Producer string `json:"producer"`
}

// NewNetworkIdentifierSubNetworkIdentifierMetadata instantiates a new NetworkIdentifierSubNetworkIdentifierMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkIdentifierSubNetworkIdentifierMetadata(producer string) *NetworkIdentifierSubNetworkIdentifierMetadata {
	this := NetworkIdentifierSubNetworkIdentifierMetadata{}
	this.Producer = producer
	return &this
}

// NewNetworkIdentifierSubNetworkIdentifierMetadataWithDefaults instantiates a new NetworkIdentifierSubNetworkIdentifierMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkIdentifierSubNetworkIdentifierMetadataWithDefaults() *NetworkIdentifierSubNetworkIdentifierMetadata {
	this := NetworkIdentifierSubNetworkIdentifierMetadata{}
	return &this
}

// GetProducer returns the Producer field value
func (o *NetworkIdentifierSubNetworkIdentifierMetadata) GetProducer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Producer
}

// GetProducerOk returns a tuple with the Producer field value
// and a boolean to check if the value has been set.
func (o *NetworkIdentifierSubNetworkIdentifierMetadata) GetProducerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Producer, true
}

// SetProducer sets field value
func (o *NetworkIdentifierSubNetworkIdentifierMetadata) SetProducer(v string) {
	o.Producer = v
}

func (o NetworkIdentifierSubNetworkIdentifierMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["producer"] = o.Producer
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkIdentifierSubNetworkIdentifierMetadata struct {
	value *NetworkIdentifierSubNetworkIdentifierMetadata
	isSet bool
}

func (v NullableNetworkIdentifierSubNetworkIdentifierMetadata) Get() *NetworkIdentifierSubNetworkIdentifierMetadata {
	return v.value
}

func (v *NullableNetworkIdentifierSubNetworkIdentifierMetadata) Set(val *NetworkIdentifierSubNetworkIdentifierMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkIdentifierSubNetworkIdentifierMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkIdentifierSubNetworkIdentifierMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkIdentifierSubNetworkIdentifierMetadata(val *NetworkIdentifierSubNetworkIdentifierMetadata) *NullableNetworkIdentifierSubNetworkIdentifierMetadata {
	return &NullableNetworkIdentifierSubNetworkIdentifierMetadata{value: val, isSet: true}
}

func (v NullableNetworkIdentifierSubNetworkIdentifierMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkIdentifierSubNetworkIdentifierMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


