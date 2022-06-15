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

// BurnchainRewardListResponse GET request that returns blocks
type BurnchainRewardListResponse struct {
	// The number of burnchain rewards to return
	Limit int32 `json:"limit"`
	// The number to burnchain rewards to skip (starting at `0`)
	Offset int32 `json:"offset"`
	Results []BurnchainReward `json:"results"`
}

// NewBurnchainRewardListResponse instantiates a new BurnchainRewardListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBurnchainRewardListResponse(limit int32, offset int32, results []BurnchainReward) *BurnchainRewardListResponse {
	this := BurnchainRewardListResponse{}
	this.Limit = limit
	this.Offset = offset
	this.Results = results
	return &this
}

// NewBurnchainRewardListResponseWithDefaults instantiates a new BurnchainRewardListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBurnchainRewardListResponseWithDefaults() *BurnchainRewardListResponse {
	this := BurnchainRewardListResponse{}
	var offset int32 = 0
	this.Offset = offset
	return &this
}

// GetLimit returns the Limit field value
func (o *BurnchainRewardListResponse) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *BurnchainRewardListResponse) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *BurnchainRewardListResponse) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *BurnchainRewardListResponse) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *BurnchainRewardListResponse) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *BurnchainRewardListResponse) SetOffset(v int32) {
	o.Offset = v
}

// GetResults returns the Results field value
func (o *BurnchainRewardListResponse) GetResults() []BurnchainReward {
	if o == nil {
		var ret []BurnchainReward
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *BurnchainRewardListResponse) GetResultsOk() ([]BurnchainReward, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *BurnchainRewardListResponse) SetResults(v []BurnchainReward) {
	o.Results = v
}

func (o BurnchainRewardListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["limit"] = o.Limit
	}
	if true {
		toSerialize["offset"] = o.Offset
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableBurnchainRewardListResponse struct {
	value *BurnchainRewardListResponse
	isSet bool
}

func (v NullableBurnchainRewardListResponse) Get() *BurnchainRewardListResponse {
	return v.value
}

func (v *NullableBurnchainRewardListResponse) Set(val *BurnchainRewardListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBurnchainRewardListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBurnchainRewardListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBurnchainRewardListResponse(val *BurnchainRewardListResponse) *NullableBurnchainRewardListResponse {
	return &NullableBurnchainRewardListResponse{value: val, isSet: true}
}

func (v NullableBurnchainRewardListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBurnchainRewardListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


