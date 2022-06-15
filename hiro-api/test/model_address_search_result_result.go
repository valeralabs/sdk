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

// AddressSearchResultResult This object carries the search result
type AddressSearchResultResult struct {
	// The id used to search this query.
	EntityId string `json:"entity_id"`
	EntityType string `json:"entity_type"`
	Metadata *AddressSearchResultResultMetadata `json:"metadata,omitempty"`
}

// NewAddressSearchResultResult instantiates a new AddressSearchResultResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressSearchResultResult(entityId string, entityType string) *AddressSearchResultResult {
	this := AddressSearchResultResult{}
	this.EntityId = entityId
	this.EntityType = entityType
	return &this
}

// NewAddressSearchResultResultWithDefaults instantiates a new AddressSearchResultResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressSearchResultResultWithDefaults() *AddressSearchResultResult {
	this := AddressSearchResultResult{}
	return &this
}

// GetEntityId returns the EntityId field value
func (o *AddressSearchResultResult) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *AddressSearchResultResult) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *AddressSearchResultResult) SetEntityId(v string) {
	o.EntityId = v
}

// GetEntityType returns the EntityType field value
func (o *AddressSearchResultResult) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *AddressSearchResultResult) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *AddressSearchResultResult) SetEntityType(v string) {
	o.EntityType = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AddressSearchResultResult) GetMetadata() AddressSearchResultResultMetadata {
	if o == nil || o.Metadata == nil {
		var ret AddressSearchResultResultMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressSearchResultResult) GetMetadataOk() (*AddressSearchResultResultMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AddressSearchResultResult) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given AddressSearchResultResultMetadata and assigns it to the Metadata field.
func (o *AddressSearchResultResult) SetMetadata(v AddressSearchResultResultMetadata) {
	o.Metadata = &v
}

func (o AddressSearchResultResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entity_id"] = o.EntityId
	}
	if true {
		toSerialize["entity_type"] = o.EntityType
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableAddressSearchResultResult struct {
	value *AddressSearchResultResult
	isSet bool
}

func (v NullableAddressSearchResultResult) Get() *AddressSearchResultResult {
	return v.value
}

func (v *NullableAddressSearchResultResult) Set(val *AddressSearchResultResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressSearchResultResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressSearchResultResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressSearchResultResult(val *AddressSearchResultResult) *NullableAddressSearchResultResult {
	return &NullableAddressSearchResultResult{value: val, isSet: true}
}

func (v NullableAddressSearchResultResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressSearchResultResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


