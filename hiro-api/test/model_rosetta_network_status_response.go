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

// RosettaNetworkStatusResponse NetworkStatusResponse contains basic information about the node's view of a blockchain network. It is assumed that any BlockIdentifier.Index less than or equal to CurrentBlockIdentifier.Index can be queried. If a Rosetta implementation prunes historical state, it should populate the optional oldest_block_identifier field with the oldest block available to query. If this is not populated, it is assumed that the genesis_block_identifier is the oldest queryable block. If a Rosetta implementation performs some pre-sync before it is possible to query blocks, sync_status should be populated so that clients can still monitor healthiness. Without this field, it may appear that the implementation is stuck syncing and needs to be terminated.
type RosettaNetworkStatusResponse struct {
	CurrentBlockIdentifier Object `json:"current_block_identifier"`
	// The timestamp of the block in milliseconds since the Unix Epoch. The timestamp is stored in milliseconds because some blockchains produce blocks more often than once a second.
	CurrentBlockTimestamp int32 `json:"current_block_timestamp"`
	GenesisBlockIdentifier RosettaGenesisBlockIdentifier `json:"genesis_block_identifier"`
	OldestBlockIdentifier *RosettaOldestBlockIdentifier `json:"oldest_block_identifier,omitempty"`
	SyncStatus *RosettaSyncStatus `json:"sync_status,omitempty"`
	// Peers information
	Peers []RosettaPeers `json:"peers"`
}

// NewRosettaNetworkStatusResponse instantiates a new RosettaNetworkStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaNetworkStatusResponse(currentBlockIdentifier Object, currentBlockTimestamp int32, genesisBlockIdentifier RosettaGenesisBlockIdentifier, peers []RosettaPeers) *RosettaNetworkStatusResponse {
	this := RosettaNetworkStatusResponse{}
	this.CurrentBlockIdentifier = currentBlockIdentifier
	this.CurrentBlockTimestamp = currentBlockTimestamp
	this.GenesisBlockIdentifier = genesisBlockIdentifier
	this.Peers = peers
	return &this
}

// NewRosettaNetworkStatusResponseWithDefaults instantiates a new RosettaNetworkStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaNetworkStatusResponseWithDefaults() *RosettaNetworkStatusResponse {
	this := RosettaNetworkStatusResponse{}
	return &this
}

// GetCurrentBlockIdentifier returns the CurrentBlockIdentifier field value
func (o *RosettaNetworkStatusResponse) GetCurrentBlockIdentifier() Object {
	if o == nil {
		var ret Object
		return ret
	}

	return o.CurrentBlockIdentifier
}

// GetCurrentBlockIdentifierOk returns a tuple with the CurrentBlockIdentifier field value
// and a boolean to check if the value has been set.
func (o *RosettaNetworkStatusResponse) GetCurrentBlockIdentifierOk() (*Object, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentBlockIdentifier, true
}

// SetCurrentBlockIdentifier sets field value
func (o *RosettaNetworkStatusResponse) SetCurrentBlockIdentifier(v Object) {
	o.CurrentBlockIdentifier = v
}

// GetCurrentBlockTimestamp returns the CurrentBlockTimestamp field value
func (o *RosettaNetworkStatusResponse) GetCurrentBlockTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CurrentBlockTimestamp
}

// GetCurrentBlockTimestampOk returns a tuple with the CurrentBlockTimestamp field value
// and a boolean to check if the value has been set.
func (o *RosettaNetworkStatusResponse) GetCurrentBlockTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentBlockTimestamp, true
}

// SetCurrentBlockTimestamp sets field value
func (o *RosettaNetworkStatusResponse) SetCurrentBlockTimestamp(v int32) {
	o.CurrentBlockTimestamp = v
}

// GetGenesisBlockIdentifier returns the GenesisBlockIdentifier field value
func (o *RosettaNetworkStatusResponse) GetGenesisBlockIdentifier() RosettaGenesisBlockIdentifier {
	if o == nil {
		var ret RosettaGenesisBlockIdentifier
		return ret
	}

	return o.GenesisBlockIdentifier
}

// GetGenesisBlockIdentifierOk returns a tuple with the GenesisBlockIdentifier field value
// and a boolean to check if the value has been set.
func (o *RosettaNetworkStatusResponse) GetGenesisBlockIdentifierOk() (*RosettaGenesisBlockIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GenesisBlockIdentifier, true
}

// SetGenesisBlockIdentifier sets field value
func (o *RosettaNetworkStatusResponse) SetGenesisBlockIdentifier(v RosettaGenesisBlockIdentifier) {
	o.GenesisBlockIdentifier = v
}

// GetOldestBlockIdentifier returns the OldestBlockIdentifier field value if set, zero value otherwise.
func (o *RosettaNetworkStatusResponse) GetOldestBlockIdentifier() RosettaOldestBlockIdentifier {
	if o == nil || o.OldestBlockIdentifier == nil {
		var ret RosettaOldestBlockIdentifier
		return ret
	}
	return *o.OldestBlockIdentifier
}

// GetOldestBlockIdentifierOk returns a tuple with the OldestBlockIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaNetworkStatusResponse) GetOldestBlockIdentifierOk() (*RosettaOldestBlockIdentifier, bool) {
	if o == nil || o.OldestBlockIdentifier == nil {
		return nil, false
	}
	return o.OldestBlockIdentifier, true
}

// HasOldestBlockIdentifier returns a boolean if a field has been set.
func (o *RosettaNetworkStatusResponse) HasOldestBlockIdentifier() bool {
	if o != nil && o.OldestBlockIdentifier != nil {
		return true
	}

	return false
}

// SetOldestBlockIdentifier gets a reference to the given RosettaOldestBlockIdentifier and assigns it to the OldestBlockIdentifier field.
func (o *RosettaNetworkStatusResponse) SetOldestBlockIdentifier(v RosettaOldestBlockIdentifier) {
	o.OldestBlockIdentifier = &v
}

// GetSyncStatus returns the SyncStatus field value if set, zero value otherwise.
func (o *RosettaNetworkStatusResponse) GetSyncStatus() RosettaSyncStatus {
	if o == nil || o.SyncStatus == nil {
		var ret RosettaSyncStatus
		return ret
	}
	return *o.SyncStatus
}

// GetSyncStatusOk returns a tuple with the SyncStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaNetworkStatusResponse) GetSyncStatusOk() (*RosettaSyncStatus, bool) {
	if o == nil || o.SyncStatus == nil {
		return nil, false
	}
	return o.SyncStatus, true
}

// HasSyncStatus returns a boolean if a field has been set.
func (o *RosettaNetworkStatusResponse) HasSyncStatus() bool {
	if o != nil && o.SyncStatus != nil {
		return true
	}

	return false
}

// SetSyncStatus gets a reference to the given RosettaSyncStatus and assigns it to the SyncStatus field.
func (o *RosettaNetworkStatusResponse) SetSyncStatus(v RosettaSyncStatus) {
	o.SyncStatus = &v
}

// GetPeers returns the Peers field value
func (o *RosettaNetworkStatusResponse) GetPeers() []RosettaPeers {
	if o == nil {
		var ret []RosettaPeers
		return ret
	}

	return o.Peers
}

// GetPeersOk returns a tuple with the Peers field value
// and a boolean to check if the value has been set.
func (o *RosettaNetworkStatusResponse) GetPeersOk() ([]RosettaPeers, bool) {
	if o == nil {
		return nil, false
	}
	return o.Peers, true
}

// SetPeers sets field value
func (o *RosettaNetworkStatusResponse) SetPeers(v []RosettaPeers) {
	o.Peers = v
}

func (o RosettaNetworkStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["current_block_identifier"] = o.CurrentBlockIdentifier
	}
	if true {
		toSerialize["current_block_timestamp"] = o.CurrentBlockTimestamp
	}
	if true {
		toSerialize["genesis_block_identifier"] = o.GenesisBlockIdentifier
	}
	if o.OldestBlockIdentifier != nil {
		toSerialize["oldest_block_identifier"] = o.OldestBlockIdentifier
	}
	if o.SyncStatus != nil {
		toSerialize["sync_status"] = o.SyncStatus
	}
	if true {
		toSerialize["peers"] = o.Peers
	}
	return json.Marshal(toSerialize)
}

type NullableRosettaNetworkStatusResponse struct {
	value *RosettaNetworkStatusResponse
	isSet bool
}

func (v NullableRosettaNetworkStatusResponse) Get() *RosettaNetworkStatusResponse {
	return v.value
}

func (v *NullableRosettaNetworkStatusResponse) Set(val *RosettaNetworkStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaNetworkStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaNetworkStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaNetworkStatusResponse(val *RosettaNetworkStatusResponse) *NullableRosettaNetworkStatusResponse {
	return &NullableRosettaNetworkStatusResponse{value: val, isSet: true}
}

func (v NullableRosettaNetworkStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaNetworkStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


