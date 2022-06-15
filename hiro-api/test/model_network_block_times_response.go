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

// NetworkBlockTimesResponse GET request that returns network target block times
type NetworkBlockTimesResponse struct {
	Mainnet TargetBlockTime `json:"mainnet"`
	Testnet Object `json:"testnet"`
}

// NewNetworkBlockTimesResponse instantiates a new NetworkBlockTimesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkBlockTimesResponse(mainnet TargetBlockTime, testnet Object) *NetworkBlockTimesResponse {
	this := NetworkBlockTimesResponse{}
	this.Mainnet = mainnet
	this.Testnet = testnet
	return &this
}

// NewNetworkBlockTimesResponseWithDefaults instantiates a new NetworkBlockTimesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkBlockTimesResponseWithDefaults() *NetworkBlockTimesResponse {
	this := NetworkBlockTimesResponse{}
	return &this
}

// GetMainnet returns the Mainnet field value
func (o *NetworkBlockTimesResponse) GetMainnet() TargetBlockTime {
	if o == nil {
		var ret TargetBlockTime
		return ret
	}

	return o.Mainnet
}

// GetMainnetOk returns a tuple with the Mainnet field value
// and a boolean to check if the value has been set.
func (o *NetworkBlockTimesResponse) GetMainnetOk() (*TargetBlockTime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mainnet, true
}

// SetMainnet sets field value
func (o *NetworkBlockTimesResponse) SetMainnet(v TargetBlockTime) {
	o.Mainnet = v
}

// GetTestnet returns the Testnet field value
func (o *NetworkBlockTimesResponse) GetTestnet() Object {
	if o == nil {
		var ret Object
		return ret
	}

	return o.Testnet
}

// GetTestnetOk returns a tuple with the Testnet field value
// and a boolean to check if the value has been set.
func (o *NetworkBlockTimesResponse) GetTestnetOk() (*Object, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Testnet, true
}

// SetTestnet sets field value
func (o *NetworkBlockTimesResponse) SetTestnet(v Object) {
	o.Testnet = v
}

func (o NetworkBlockTimesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mainnet"] = o.Mainnet
	}
	if true {
		toSerialize["testnet"] = o.Testnet
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkBlockTimesResponse struct {
	value *NetworkBlockTimesResponse
	isSet bool
}

func (v NullableNetworkBlockTimesResponse) Get() *NetworkBlockTimesResponse {
	return v.value
}

func (v *NullableNetworkBlockTimesResponse) Set(val *NetworkBlockTimesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkBlockTimesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkBlockTimesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkBlockTimesResponse(val *NetworkBlockTimesResponse) *NullableNetworkBlockTimesResponse {
	return &NullableNetworkBlockTimesResponse{value: val, isSet: true}
}

func (v NullableNetworkBlockTimesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkBlockTimesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


