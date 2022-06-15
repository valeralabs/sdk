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

// RosettaNetworkOptionsResponse NetworkOptionsResponse contains information about the versioning of the node and the allowed operation statuses, operation types, and errors.
type RosettaNetworkOptionsResponse struct {
	Version RosettaNetworkOptionsResponseVersion `json:"version"`
	Allow RosettaNetworkOptionsResponseAllow `json:"allow"`
}

// NewRosettaNetworkOptionsResponse instantiates a new RosettaNetworkOptionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaNetworkOptionsResponse(version RosettaNetworkOptionsResponseVersion, allow RosettaNetworkOptionsResponseAllow) *RosettaNetworkOptionsResponse {
	this := RosettaNetworkOptionsResponse{}
	this.Version = version
	this.Allow = allow
	return &this
}

// NewRosettaNetworkOptionsResponseWithDefaults instantiates a new RosettaNetworkOptionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaNetworkOptionsResponseWithDefaults() *RosettaNetworkOptionsResponse {
	this := RosettaNetworkOptionsResponse{}
	return &this
}

// GetVersion returns the Version field value
func (o *RosettaNetworkOptionsResponse) GetVersion() RosettaNetworkOptionsResponseVersion {
	if o == nil {
		var ret RosettaNetworkOptionsResponseVersion
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *RosettaNetworkOptionsResponse) GetVersionOk() (*RosettaNetworkOptionsResponseVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *RosettaNetworkOptionsResponse) SetVersion(v RosettaNetworkOptionsResponseVersion) {
	o.Version = v
}

// GetAllow returns the Allow field value
func (o *RosettaNetworkOptionsResponse) GetAllow() RosettaNetworkOptionsResponseAllow {
	if o == nil {
		var ret RosettaNetworkOptionsResponseAllow
		return ret
	}

	return o.Allow
}

// GetAllowOk returns a tuple with the Allow field value
// and a boolean to check if the value has been set.
func (o *RosettaNetworkOptionsResponse) GetAllowOk() (*RosettaNetworkOptionsResponseAllow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Allow, true
}

// SetAllow sets field value
func (o *RosettaNetworkOptionsResponse) SetAllow(v RosettaNetworkOptionsResponseAllow) {
	o.Allow = v
}

func (o RosettaNetworkOptionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["allow"] = o.Allow
	}
	return json.Marshal(toSerialize)
}

type NullableRosettaNetworkOptionsResponse struct {
	value *RosettaNetworkOptionsResponse
	isSet bool
}

func (v NullableRosettaNetworkOptionsResponse) Get() *RosettaNetworkOptionsResponse {
	return v.value
}

func (v *NullableRosettaNetworkOptionsResponse) Set(val *RosettaNetworkOptionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaNetworkOptionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaNetworkOptionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaNetworkOptionsResponse(val *RosettaNetworkOptionsResponse) *NullableRosettaNetworkOptionsResponse {
	return &NullableRosettaNetworkOptionsResponse{value: val, isSet: true}
}

func (v NullableRosettaNetworkOptionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaNetworkOptionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


