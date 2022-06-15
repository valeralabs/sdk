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

// PostConditionNonFungibleAllOf struct for PostConditionNonFungibleAllOf
type PostConditionNonFungibleAllOf struct {
	// A non-fungible condition code encodes a statement being made about a non-fungible token, with respect to whether or not the particular non-fungible token is owned by the account.
	ConditionCode string `json:"condition_code"`
	Type string `json:"type"`
	AssetValue PostConditionNonFungibleAllOfAssetValue `json:"asset_value"`
	Asset PostConditionFungibleAllOf1Asset `json:"asset"`
}

// NewPostConditionNonFungibleAllOf instantiates a new PostConditionNonFungibleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostConditionNonFungibleAllOf(conditionCode string, type_ string, assetValue PostConditionNonFungibleAllOfAssetValue, asset PostConditionFungibleAllOf1Asset) *PostConditionNonFungibleAllOf {
	this := PostConditionNonFungibleAllOf{}
	this.ConditionCode = conditionCode
	this.Type = type_
	this.AssetValue = assetValue
	this.Asset = asset
	return &this
}

// NewPostConditionNonFungibleAllOfWithDefaults instantiates a new PostConditionNonFungibleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostConditionNonFungibleAllOfWithDefaults() *PostConditionNonFungibleAllOf {
	this := PostConditionNonFungibleAllOf{}
	return &this
}

// GetConditionCode returns the ConditionCode field value
func (o *PostConditionNonFungibleAllOf) GetConditionCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConditionCode
}

// GetConditionCodeOk returns a tuple with the ConditionCode field value
// and a boolean to check if the value has been set.
func (o *PostConditionNonFungibleAllOf) GetConditionCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConditionCode, true
}

// SetConditionCode sets field value
func (o *PostConditionNonFungibleAllOf) SetConditionCode(v string) {
	o.ConditionCode = v
}

// GetType returns the Type field value
func (o *PostConditionNonFungibleAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PostConditionNonFungibleAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PostConditionNonFungibleAllOf) SetType(v string) {
	o.Type = v
}

// GetAssetValue returns the AssetValue field value
func (o *PostConditionNonFungibleAllOf) GetAssetValue() PostConditionNonFungibleAllOfAssetValue {
	if o == nil {
		var ret PostConditionNonFungibleAllOfAssetValue
		return ret
	}

	return o.AssetValue
}

// GetAssetValueOk returns a tuple with the AssetValue field value
// and a boolean to check if the value has been set.
func (o *PostConditionNonFungibleAllOf) GetAssetValueOk() (*PostConditionNonFungibleAllOfAssetValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetValue, true
}

// SetAssetValue sets field value
func (o *PostConditionNonFungibleAllOf) SetAssetValue(v PostConditionNonFungibleAllOfAssetValue) {
	o.AssetValue = v
}

// GetAsset returns the Asset field value
func (o *PostConditionNonFungibleAllOf) GetAsset() PostConditionFungibleAllOf1Asset {
	if o == nil {
		var ret PostConditionFungibleAllOf1Asset
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *PostConditionNonFungibleAllOf) GetAssetOk() (*PostConditionFungibleAllOf1Asset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *PostConditionNonFungibleAllOf) SetAsset(v PostConditionFungibleAllOf1Asset) {
	o.Asset = v
}

func (o PostConditionNonFungibleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["condition_code"] = o.ConditionCode
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["asset_value"] = o.AssetValue
	}
	if true {
		toSerialize["asset"] = o.Asset
	}
	return json.Marshal(toSerialize)
}

type NullablePostConditionNonFungibleAllOf struct {
	value *PostConditionNonFungibleAllOf
	isSet bool
}

func (v NullablePostConditionNonFungibleAllOf) Get() *PostConditionNonFungibleAllOf {
	return v.value
}

func (v *NullablePostConditionNonFungibleAllOf) Set(val *PostConditionNonFungibleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePostConditionNonFungibleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePostConditionNonFungibleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostConditionNonFungibleAllOf(val *PostConditionNonFungibleAllOf) *NullablePostConditionNonFungibleAllOf {
	return &NullablePostConditionNonFungibleAllOf{value: val, isSet: true}
}

func (v NullablePostConditionNonFungibleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostConditionNonFungibleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


