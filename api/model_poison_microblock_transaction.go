/*
Stacks Blockchain API

Welcome to the API reference overview for the <a href=\"https://docs.hiro.so/get-started/stacks-blockchain-api\">Stacks Blockchain API</a>.  <a href=\"https://hirosystems.github.io/stacks-blockchain-api/collection.json\" download=\"stacks-api-collection.json\">Download Postman collection</a> 

API version: v4.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// PoisonMicroblockTransaction Describes representation of a Type 3 Stacks 2.0 transaction: Poison Microblock
type PoisonMicroblockTransaction struct {
	// Transaction ID
	TxId string `json:"tx_id"`
	// Used for ordering the transactions originating from and paying from an account. The nonce ensures that a transaction is processed at most once. The nonce counts the number of times an account's owner(s) have authorized a transaction. The first transaction from an account will have a nonce value equal to 0, the second will have a nonce value equal to 1, and so on.
	Nonce int32 `json:"nonce"`
	// Transaction fee as Integer string (64-bit unsigned integer).
	FeeRate string `json:"fee_rate"`
	// Address of the transaction initiator
	SenderAddress string `json:"sender_address"`
	SponsorNonce *int32 `json:"sponsor_nonce,omitempty"`
	// Denotes whether the originating account is the same as the paying account
	Sponsored bool `json:"sponsored"`
	SponsorAddress *string `json:"sponsor_address,omitempty"`
	// 
	PostConditionMode string `json:"post_condition_mode"`
	PostConditions []PostCondition `json:"post_conditions"`
	// `on_chain_only`: the transaction MUST be included in an anchored block, `off_chain_only`: the transaction MUST be included in a microblock, `any`: the leader can choose where to include the transaction.
	AnchorMode string `json:"anchor_mode"`
	// Hash of the blocked this transactions was associated with
	BlockHash string `json:"block_hash"`
	// Height of the block this transactions was associated with
	BlockHeight int32 `json:"block_height"`
	// Unix timestamp (in seconds) indicating when this block was mined
	BurnBlockTime int32 `json:"burn_block_time"`
	// An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) timestamp indicating when this block was mined.
	BurnBlockTimeIso string `json:"burn_block_time_iso"`
	// Unix timestamp (in seconds) indicating when this parent block was mined
	ParentBurnBlockTime int32 `json:"parent_burn_block_time"`
	// An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) timestamp indicating when this parent block was mined.
	ParentBurnBlockTimeIso string `json:"parent_burn_block_time_iso"`
	// Set to `true` if block corresponds to the canonical chain tip
	Canonical bool `json:"canonical"`
	// Index of the transaction, indicating the order. Starts at `0` and increases with each transaction
	TxIndex int32 `json:"tx_index"`
	// Status of the transaction
	TxStatus string `json:"tx_status"`
	TxResult AbstractTransactionAllOfTxResult `json:"tx_result"`
	// Number of transaction events
	EventCount int32 `json:"event_count"`
	// Hash of the previous block.
	ParentBlockHash string `json:"parent_block_hash"`
	// True if the transaction is included in a microblock that has not been confirmed by an anchor block.
	IsUnanchored bool `json:"is_unanchored"`
	// The microblock hash that this transaction was streamed in. If the transaction was batched in an anchor block (not included within a microblock) then this value will be an empty string.
	MicroblockHash string `json:"microblock_hash"`
	// The microblock sequence number that this transaction was streamed in. If the transaction was batched in an anchor block (not included within a microblock) then this value will be 2147483647 (0x7fffffff, the max int32 value), this value preserves logical transaction ordering on (block_height, microblock_sequence, tx_index).
	MicroblockSequence int32 `json:"microblock_sequence"`
	// Set to `true` if microblock is anchored in the canonical chain tip, `false` if the transaction was orphaned in a micro-fork.
	MicroblockCanonical bool `json:"microblock_canonical"`
	// Execution cost read count.
	ExecutionCostReadCount int32 `json:"execution_cost_read_count"`
	// Execution cost read length.
	ExecutionCostReadLength int32 `json:"execution_cost_read_length"`
	// Execution cost runtime.
	ExecutionCostRuntime int32 `json:"execution_cost_runtime"`
	// Execution cost write count.
	ExecutionCostWriteCount int32 `json:"execution_cost_write_count"`
	// Execution cost write length.
	ExecutionCostWriteLength int32 `json:"execution_cost_write_length"`
	// List of transaction events
	Events []TransactionEvent `json:"events"`
	TxType string `json:"tx_type"`
	PoisonMicroblock PoisonMicroblockTransactionMetadataPoisonMicroblock `json:"poison_microblock"`
}

// NewPoisonMicroblockTransaction instantiates a new PoisonMicroblockTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoisonMicroblockTransaction(txId string, nonce int32, feeRate string, senderAddress string, sponsored bool, postConditionMode string, postConditions []PostCondition, anchorMode string, blockHash string, blockHeight int32, burnBlockTime int32, burnBlockTimeIso string, parentBurnBlockTime int32, parentBurnBlockTimeIso string, canonical bool, txIndex int32, txStatus string, txResult AbstractTransactionAllOfTxResult, eventCount int32, parentBlockHash string, isUnanchored bool, microblockHash string, microblockSequence int32, microblockCanonical bool, executionCostReadCount int32, executionCostReadLength int32, executionCostRuntime int32, executionCostWriteCount int32, executionCostWriteLength int32, events []TransactionEvent, txType string, poisonMicroblock PoisonMicroblockTransactionMetadataPoisonMicroblock) *PoisonMicroblockTransaction {
	this := PoisonMicroblockTransaction{}
	this.TxId = txId
	this.Nonce = nonce
	this.FeeRate = feeRate
	this.SenderAddress = senderAddress
	this.Sponsored = sponsored
	this.PostConditionMode = postConditionMode
	this.PostConditions = postConditions
	this.AnchorMode = anchorMode
	this.BlockHash = blockHash
	this.BlockHeight = blockHeight
	this.BurnBlockTime = burnBlockTime
	this.BurnBlockTimeIso = burnBlockTimeIso
	this.ParentBurnBlockTime = parentBurnBlockTime
	this.ParentBurnBlockTimeIso = parentBurnBlockTimeIso
	this.Canonical = canonical
	this.TxIndex = txIndex
	this.TxStatus = txStatus
	this.TxResult = txResult
	this.EventCount = eventCount
	this.ParentBlockHash = parentBlockHash
	this.IsUnanchored = isUnanchored
	this.MicroblockHash = microblockHash
	this.MicroblockSequence = microblockSequence
	this.MicroblockCanonical = microblockCanonical
	this.ExecutionCostReadCount = executionCostReadCount
	this.ExecutionCostReadLength = executionCostReadLength
	this.ExecutionCostRuntime = executionCostRuntime
	this.ExecutionCostWriteCount = executionCostWriteCount
	this.ExecutionCostWriteLength = executionCostWriteLength
	this.Events = events
	this.TxType = txType
	this.PoisonMicroblock = poisonMicroblock
	return &this
}

// NewPoisonMicroblockTransactionWithDefaults instantiates a new PoisonMicroblockTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoisonMicroblockTransactionWithDefaults() *PoisonMicroblockTransaction {
	this := PoisonMicroblockTransaction{}
	return &this
}

// GetTxId returns the TxId field value
func (o *PoisonMicroblockTransaction) GetTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *PoisonMicroblockTransaction) SetTxId(v string) {
	o.TxId = v
}

// GetNonce returns the Nonce field value
func (o *PoisonMicroblockTransaction) GetNonce() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetNonceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *PoisonMicroblockTransaction) SetNonce(v int32) {
	o.Nonce = v
}

// GetFeeRate returns the FeeRate field value
func (o *PoisonMicroblockTransaction) GetFeeRate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeeRate
}

// GetFeeRateOk returns a tuple with the FeeRate field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetFeeRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeRate, true
}

// SetFeeRate sets field value
func (o *PoisonMicroblockTransaction) SetFeeRate(v string) {
	o.FeeRate = v
}

// GetSenderAddress returns the SenderAddress field value
func (o *PoisonMicroblockTransaction) GetSenderAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SenderAddress
}

// GetSenderAddressOk returns a tuple with the SenderAddress field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetSenderAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderAddress, true
}

// SetSenderAddress sets field value
func (o *PoisonMicroblockTransaction) SetSenderAddress(v string) {
	o.SenderAddress = v
}

// GetSponsorNonce returns the SponsorNonce field value if set, zero value otherwise.
func (o *PoisonMicroblockTransaction) GetSponsorNonce() int32 {
	if o == nil || o.SponsorNonce == nil {
		var ret int32
		return ret
	}
	return *o.SponsorNonce
}

// GetSponsorNonceOk returns a tuple with the SponsorNonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetSponsorNonceOk() (*int32, bool) {
	if o == nil || o.SponsorNonce == nil {
		return nil, false
	}
	return o.SponsorNonce, true
}

// HasSponsorNonce returns a boolean if a field has been set.
func (o *PoisonMicroblockTransaction) HasSponsorNonce() bool {
	if o != nil && o.SponsorNonce != nil {
		return true
	}

	return false
}

// SetSponsorNonce gets a reference to the given int32 and assigns it to the SponsorNonce field.
func (o *PoisonMicroblockTransaction) SetSponsorNonce(v int32) {
	o.SponsorNonce = &v
}

// GetSponsored returns the Sponsored field value
func (o *PoisonMicroblockTransaction) GetSponsored() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Sponsored
}

// GetSponsoredOk returns a tuple with the Sponsored field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetSponsoredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sponsored, true
}

// SetSponsored sets field value
func (o *PoisonMicroblockTransaction) SetSponsored(v bool) {
	o.Sponsored = v
}

// GetSponsorAddress returns the SponsorAddress field value if set, zero value otherwise.
func (o *PoisonMicroblockTransaction) GetSponsorAddress() string {
	if o == nil || o.SponsorAddress == nil {
		var ret string
		return ret
	}
	return *o.SponsorAddress
}

// GetSponsorAddressOk returns a tuple with the SponsorAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetSponsorAddressOk() (*string, bool) {
	if o == nil || o.SponsorAddress == nil {
		return nil, false
	}
	return o.SponsorAddress, true
}

// HasSponsorAddress returns a boolean if a field has been set.
func (o *PoisonMicroblockTransaction) HasSponsorAddress() bool {
	if o != nil && o.SponsorAddress != nil {
		return true
	}

	return false
}

// SetSponsorAddress gets a reference to the given string and assigns it to the SponsorAddress field.
func (o *PoisonMicroblockTransaction) SetSponsorAddress(v string) {
	o.SponsorAddress = &v
}

// GetPostConditionMode returns the PostConditionMode field value
func (o *PoisonMicroblockTransaction) GetPostConditionMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostConditionMode
}

// GetPostConditionModeOk returns a tuple with the PostConditionMode field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetPostConditionModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostConditionMode, true
}

// SetPostConditionMode sets field value
func (o *PoisonMicroblockTransaction) SetPostConditionMode(v string) {
	o.PostConditionMode = v
}

// GetPostConditions returns the PostConditions field value
func (o *PoisonMicroblockTransaction) GetPostConditions() []PostCondition {
	if o == nil {
		var ret []PostCondition
		return ret
	}

	return o.PostConditions
}

// GetPostConditionsOk returns a tuple with the PostConditions field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetPostConditionsOk() ([]PostCondition, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostConditions, true
}

// SetPostConditions sets field value
func (o *PoisonMicroblockTransaction) SetPostConditions(v []PostCondition) {
	o.PostConditions = v
}

// GetAnchorMode returns the AnchorMode field value
func (o *PoisonMicroblockTransaction) GetAnchorMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AnchorMode
}

// GetAnchorModeOk returns a tuple with the AnchorMode field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetAnchorModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnchorMode, true
}

// SetAnchorMode sets field value
func (o *PoisonMicroblockTransaction) SetAnchorMode(v string) {
	o.AnchorMode = v
}

// GetBlockHash returns the BlockHash field value
func (o *PoisonMicroblockTransaction) GetBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHash, true
}

// SetBlockHash sets field value
func (o *PoisonMicroblockTransaction) SetBlockHash(v string) {
	o.BlockHash = v
}

// GetBlockHeight returns the BlockHeight field value
func (o *PoisonMicroblockTransaction) GetBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BlockHeight
}

// GetBlockHeightOk returns a tuple with the BlockHeight field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHeight, true
}

// SetBlockHeight sets field value
func (o *PoisonMicroblockTransaction) SetBlockHeight(v int32) {
	o.BlockHeight = v
}

// GetBurnBlockTime returns the BurnBlockTime field value
func (o *PoisonMicroblockTransaction) GetBurnBlockTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BurnBlockTime
}

// GetBurnBlockTimeOk returns a tuple with the BurnBlockTime field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetBurnBlockTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockTime, true
}

// SetBurnBlockTime sets field value
func (o *PoisonMicroblockTransaction) SetBurnBlockTime(v int32) {
	o.BurnBlockTime = v
}

// GetBurnBlockTimeIso returns the BurnBlockTimeIso field value
func (o *PoisonMicroblockTransaction) GetBurnBlockTimeIso() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BurnBlockTimeIso
}

// GetBurnBlockTimeIsoOk returns a tuple with the BurnBlockTimeIso field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetBurnBlockTimeIsoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockTimeIso, true
}

// SetBurnBlockTimeIso sets field value
func (o *PoisonMicroblockTransaction) SetBurnBlockTimeIso(v string) {
	o.BurnBlockTimeIso = v
}

// GetParentBurnBlockTime returns the ParentBurnBlockTime field value
func (o *PoisonMicroblockTransaction) GetParentBurnBlockTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ParentBurnBlockTime
}

// GetParentBurnBlockTimeOk returns a tuple with the ParentBurnBlockTime field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetParentBurnBlockTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentBurnBlockTime, true
}

// SetParentBurnBlockTime sets field value
func (o *PoisonMicroblockTransaction) SetParentBurnBlockTime(v int32) {
	o.ParentBurnBlockTime = v
}

// GetParentBurnBlockTimeIso returns the ParentBurnBlockTimeIso field value
func (o *PoisonMicroblockTransaction) GetParentBurnBlockTimeIso() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentBurnBlockTimeIso
}

// GetParentBurnBlockTimeIsoOk returns a tuple with the ParentBurnBlockTimeIso field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetParentBurnBlockTimeIsoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentBurnBlockTimeIso, true
}

// SetParentBurnBlockTimeIso sets field value
func (o *PoisonMicroblockTransaction) SetParentBurnBlockTimeIso(v string) {
	o.ParentBurnBlockTimeIso = v
}

// GetCanonical returns the Canonical field value
func (o *PoisonMicroblockTransaction) GetCanonical() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Canonical
}

// GetCanonicalOk returns a tuple with the Canonical field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetCanonicalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Canonical, true
}

// SetCanonical sets field value
func (o *PoisonMicroblockTransaction) SetCanonical(v bool) {
	o.Canonical = v
}

// GetTxIndex returns the TxIndex field value
func (o *PoisonMicroblockTransaction) GetTxIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TxIndex
}

// GetTxIndexOk returns a tuple with the TxIndex field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetTxIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxIndex, true
}

// SetTxIndex sets field value
func (o *PoisonMicroblockTransaction) SetTxIndex(v int32) {
	o.TxIndex = v
}

// GetTxStatus returns the TxStatus field value
func (o *PoisonMicroblockTransaction) GetTxStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxStatus
}

// GetTxStatusOk returns a tuple with the TxStatus field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetTxStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxStatus, true
}

// SetTxStatus sets field value
func (o *PoisonMicroblockTransaction) SetTxStatus(v string) {
	o.TxStatus = v
}

// GetTxResult returns the TxResult field value
func (o *PoisonMicroblockTransaction) GetTxResult() AbstractTransactionAllOfTxResult {
	if o == nil {
		var ret AbstractTransactionAllOfTxResult
		return ret
	}

	return o.TxResult
}

// GetTxResultOk returns a tuple with the TxResult field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetTxResultOk() (*AbstractTransactionAllOfTxResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxResult, true
}

// SetTxResult sets field value
func (o *PoisonMicroblockTransaction) SetTxResult(v AbstractTransactionAllOfTxResult) {
	o.TxResult = v
}

// GetEventCount returns the EventCount field value
func (o *PoisonMicroblockTransaction) GetEventCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventCount
}

// GetEventCountOk returns a tuple with the EventCount field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetEventCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventCount, true
}

// SetEventCount sets field value
func (o *PoisonMicroblockTransaction) SetEventCount(v int32) {
	o.EventCount = v
}

// GetParentBlockHash returns the ParentBlockHash field value
func (o *PoisonMicroblockTransaction) GetParentBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentBlockHash
}

// GetParentBlockHashOk returns a tuple with the ParentBlockHash field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetParentBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentBlockHash, true
}

// SetParentBlockHash sets field value
func (o *PoisonMicroblockTransaction) SetParentBlockHash(v string) {
	o.ParentBlockHash = v
}

// GetIsUnanchored returns the IsUnanchored field value
func (o *PoisonMicroblockTransaction) GetIsUnanchored() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsUnanchored
}

// GetIsUnanchoredOk returns a tuple with the IsUnanchored field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetIsUnanchoredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsUnanchored, true
}

// SetIsUnanchored sets field value
func (o *PoisonMicroblockTransaction) SetIsUnanchored(v bool) {
	o.IsUnanchored = v
}

// GetMicroblockHash returns the MicroblockHash field value
func (o *PoisonMicroblockTransaction) GetMicroblockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicroblockHash
}

// GetMicroblockHashOk returns a tuple with the MicroblockHash field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetMicroblockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicroblockHash, true
}

// SetMicroblockHash sets field value
func (o *PoisonMicroblockTransaction) SetMicroblockHash(v string) {
	o.MicroblockHash = v
}

// GetMicroblockSequence returns the MicroblockSequence field value
func (o *PoisonMicroblockTransaction) GetMicroblockSequence() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MicroblockSequence
}

// GetMicroblockSequenceOk returns a tuple with the MicroblockSequence field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetMicroblockSequenceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicroblockSequence, true
}

// SetMicroblockSequence sets field value
func (o *PoisonMicroblockTransaction) SetMicroblockSequence(v int32) {
	o.MicroblockSequence = v
}

// GetMicroblockCanonical returns the MicroblockCanonical field value
func (o *PoisonMicroblockTransaction) GetMicroblockCanonical() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MicroblockCanonical
}

// GetMicroblockCanonicalOk returns a tuple with the MicroblockCanonical field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetMicroblockCanonicalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicroblockCanonical, true
}

// SetMicroblockCanonical sets field value
func (o *PoisonMicroblockTransaction) SetMicroblockCanonical(v bool) {
	o.MicroblockCanonical = v
}

// GetExecutionCostReadCount returns the ExecutionCostReadCount field value
func (o *PoisonMicroblockTransaction) GetExecutionCostReadCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExecutionCostReadCount
}

// GetExecutionCostReadCountOk returns a tuple with the ExecutionCostReadCount field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetExecutionCostReadCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionCostReadCount, true
}

// SetExecutionCostReadCount sets field value
func (o *PoisonMicroblockTransaction) SetExecutionCostReadCount(v int32) {
	o.ExecutionCostReadCount = v
}

// GetExecutionCostReadLength returns the ExecutionCostReadLength field value
func (o *PoisonMicroblockTransaction) GetExecutionCostReadLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExecutionCostReadLength
}

// GetExecutionCostReadLengthOk returns a tuple with the ExecutionCostReadLength field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetExecutionCostReadLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionCostReadLength, true
}

// SetExecutionCostReadLength sets field value
func (o *PoisonMicroblockTransaction) SetExecutionCostReadLength(v int32) {
	o.ExecutionCostReadLength = v
}

// GetExecutionCostRuntime returns the ExecutionCostRuntime field value
func (o *PoisonMicroblockTransaction) GetExecutionCostRuntime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExecutionCostRuntime
}

// GetExecutionCostRuntimeOk returns a tuple with the ExecutionCostRuntime field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetExecutionCostRuntimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionCostRuntime, true
}

// SetExecutionCostRuntime sets field value
func (o *PoisonMicroblockTransaction) SetExecutionCostRuntime(v int32) {
	o.ExecutionCostRuntime = v
}

// GetExecutionCostWriteCount returns the ExecutionCostWriteCount field value
func (o *PoisonMicroblockTransaction) GetExecutionCostWriteCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExecutionCostWriteCount
}

// GetExecutionCostWriteCountOk returns a tuple with the ExecutionCostWriteCount field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetExecutionCostWriteCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionCostWriteCount, true
}

// SetExecutionCostWriteCount sets field value
func (o *PoisonMicroblockTransaction) SetExecutionCostWriteCount(v int32) {
	o.ExecutionCostWriteCount = v
}

// GetExecutionCostWriteLength returns the ExecutionCostWriteLength field value
func (o *PoisonMicroblockTransaction) GetExecutionCostWriteLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExecutionCostWriteLength
}

// GetExecutionCostWriteLengthOk returns a tuple with the ExecutionCostWriteLength field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetExecutionCostWriteLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionCostWriteLength, true
}

// SetExecutionCostWriteLength sets field value
func (o *PoisonMicroblockTransaction) SetExecutionCostWriteLength(v int32) {
	o.ExecutionCostWriteLength = v
}

// GetEvents returns the Events field value
func (o *PoisonMicroblockTransaction) GetEvents() []TransactionEvent {
	if o == nil {
		var ret []TransactionEvent
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetEventsOk() ([]TransactionEvent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *PoisonMicroblockTransaction) SetEvents(v []TransactionEvent) {
	o.Events = v
}

// GetTxType returns the TxType field value
func (o *PoisonMicroblockTransaction) GetTxType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxType
}

// GetTxTypeOk returns a tuple with the TxType field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetTxTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxType, true
}

// SetTxType sets field value
func (o *PoisonMicroblockTransaction) SetTxType(v string) {
	o.TxType = v
}

// GetPoisonMicroblock returns the PoisonMicroblock field value
func (o *PoisonMicroblockTransaction) GetPoisonMicroblock() PoisonMicroblockTransactionMetadataPoisonMicroblock {
	if o == nil {
		var ret PoisonMicroblockTransactionMetadataPoisonMicroblock
		return ret
	}

	return o.PoisonMicroblock
}

// GetPoisonMicroblockOk returns a tuple with the PoisonMicroblock field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction) GetPoisonMicroblockOk() (*PoisonMicroblockTransactionMetadataPoisonMicroblock, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoisonMicroblock, true
}

// SetPoisonMicroblock sets field value
func (o *PoisonMicroblockTransaction) SetPoisonMicroblock(v PoisonMicroblockTransactionMetadataPoisonMicroblock) {
	o.PoisonMicroblock = v
}

func (o PoisonMicroblockTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tx_id"] = o.TxId
	}
	if true {
		toSerialize["nonce"] = o.Nonce
	}
	if true {
		toSerialize["fee_rate"] = o.FeeRate
	}
	if true {
		toSerialize["sender_address"] = o.SenderAddress
	}
	if o.SponsorNonce != nil {
		toSerialize["sponsor_nonce"] = o.SponsorNonce
	}
	if true {
		toSerialize["sponsored"] = o.Sponsored
	}
	if o.SponsorAddress != nil {
		toSerialize["sponsor_address"] = o.SponsorAddress
	}
	if true {
		toSerialize["post_condition_mode"] = o.PostConditionMode
	}
	if true {
		toSerialize["post_conditions"] = o.PostConditions
	}
	if true {
		toSerialize["anchor_mode"] = o.AnchorMode
	}
	if true {
		toSerialize["block_hash"] = o.BlockHash
	}
	if true {
		toSerialize["block_height"] = o.BlockHeight
	}
	if true {
		toSerialize["burn_block_time"] = o.BurnBlockTime
	}
	if true {
		toSerialize["burn_block_time_iso"] = o.BurnBlockTimeIso
	}
	if true {
		toSerialize["parent_burn_block_time"] = o.ParentBurnBlockTime
	}
	if true {
		toSerialize["parent_burn_block_time_iso"] = o.ParentBurnBlockTimeIso
	}
	if true {
		toSerialize["canonical"] = o.Canonical
	}
	if true {
		toSerialize["tx_index"] = o.TxIndex
	}
	if true {
		toSerialize["tx_status"] = o.TxStatus
	}
	if true {
		toSerialize["tx_result"] = o.TxResult
	}
	if true {
		toSerialize["event_count"] = o.EventCount
	}
	if true {
		toSerialize["parent_block_hash"] = o.ParentBlockHash
	}
	if true {
		toSerialize["is_unanchored"] = o.IsUnanchored
	}
	if true {
		toSerialize["microblock_hash"] = o.MicroblockHash
	}
	if true {
		toSerialize["microblock_sequence"] = o.MicroblockSequence
	}
	if true {
		toSerialize["microblock_canonical"] = o.MicroblockCanonical
	}
	if true {
		toSerialize["execution_cost_read_count"] = o.ExecutionCostReadCount
	}
	if true {
		toSerialize["execution_cost_read_length"] = o.ExecutionCostReadLength
	}
	if true {
		toSerialize["execution_cost_runtime"] = o.ExecutionCostRuntime
	}
	if true {
		toSerialize["execution_cost_write_count"] = o.ExecutionCostWriteCount
	}
	if true {
		toSerialize["execution_cost_write_length"] = o.ExecutionCostWriteLength
	}
	if true {
		toSerialize["events"] = o.Events
	}
	if true {
		toSerialize["tx_type"] = o.TxType
	}
	if true {
		toSerialize["poison_microblock"] = o.PoisonMicroblock
	}
	return json.Marshal(toSerialize)
}

type NullablePoisonMicroblockTransaction struct {
	value *PoisonMicroblockTransaction
	isSet bool
}

func (v NullablePoisonMicroblockTransaction) Get() *PoisonMicroblockTransaction {
	return v.value
}

func (v *NullablePoisonMicroblockTransaction) Set(val *PoisonMicroblockTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullablePoisonMicroblockTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullablePoisonMicroblockTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoisonMicroblockTransaction(val *PoisonMicroblockTransaction) *NullablePoisonMicroblockTransaction {
	return &NullablePoisonMicroblockTransaction{value: val, isSet: true}
}

func (v NullablePoisonMicroblockTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoisonMicroblockTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


