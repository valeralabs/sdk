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

// BurnchainRewardsTotal Total burnchain rewards made to a recipient
type BurnchainRewardsTotal struct {
	// The recipient address that received the burnchain rewards, in the format native to the burnchain (e.g. B58 encoded for Bitcoin)
	RewardRecipient string `json:"reward_recipient"`
	// The total amount of burnchain tokens rewarded to the recipient, in the smallest unit (e.g. satoshis for Bitcoin)
	RewardAmount string `json:"reward_amount"`
}

// NewBurnchainRewardsTotal instantiates a new BurnchainRewardsTotal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBurnchainRewardsTotal(rewardRecipient string, rewardAmount string) *BurnchainRewardsTotal {
	this := BurnchainRewardsTotal{}
	this.RewardRecipient = rewardRecipient
	this.RewardAmount = rewardAmount
	return &this
}

// NewBurnchainRewardsTotalWithDefaults instantiates a new BurnchainRewardsTotal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBurnchainRewardsTotalWithDefaults() *BurnchainRewardsTotal {
	this := BurnchainRewardsTotal{}
	return &this
}

// GetRewardRecipient returns the RewardRecipient field value
func (o *BurnchainRewardsTotal) GetRewardRecipient() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RewardRecipient
}

// GetRewardRecipientOk returns a tuple with the RewardRecipient field value
// and a boolean to check if the value has been set.
func (o *BurnchainRewardsTotal) GetRewardRecipientOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RewardRecipient, true
}

// SetRewardRecipient sets field value
func (o *BurnchainRewardsTotal) SetRewardRecipient(v string) {
	o.RewardRecipient = v
}

// GetRewardAmount returns the RewardAmount field value
func (o *BurnchainRewardsTotal) GetRewardAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RewardAmount
}

// GetRewardAmountOk returns a tuple with the RewardAmount field value
// and a boolean to check if the value has been set.
func (o *BurnchainRewardsTotal) GetRewardAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RewardAmount, true
}

// SetRewardAmount sets field value
func (o *BurnchainRewardsTotal) SetRewardAmount(v string) {
	o.RewardAmount = v
}

func (o BurnchainRewardsTotal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["reward_recipient"] = o.RewardRecipient
	}
	if true {
		toSerialize["reward_amount"] = o.RewardAmount
	}
	return json.Marshal(toSerialize)
}

type NullableBurnchainRewardsTotal struct {
	value *BurnchainRewardsTotal
	isSet bool
}

func (v NullableBurnchainRewardsTotal) Get() *BurnchainRewardsTotal {
	return v.value
}

func (v *NullableBurnchainRewardsTotal) Set(val *BurnchainRewardsTotal) {
	v.value = val
	v.isSet = true
}

func (v NullableBurnchainRewardsTotal) IsSet() bool {
	return v.isSet
}

func (v *NullableBurnchainRewardsTotal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBurnchainRewardsTotal(val *BurnchainRewardsTotal) *NullableBurnchainRewardsTotal {
	return &NullableBurnchainRewardsTotal{value: val, isSet: true}
}

func (v NullableBurnchainRewardsTotal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBurnchainRewardsTotal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


