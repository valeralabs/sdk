package api

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Defines a 200 (List of blocks) response for GetBlockList.
type GetBlockListResponse struct {
	// The number of blocks to return.
	Limit int `json:"limit"`
	// The number to blocks to skip (starting at `0`).
	Offset  int   `json:"offset"`
	Results []any `json:"results"`
	// The number of blocks available.
	Total int `json:"total"`
}

// Defines parameters for GetBlockList
type GetBlockListParams struct {
	// Max number of blocks to fetch.
	// Optional. Use `-1` to use the default value (20).
	// Max value is 200.
	Limit int
	// Index of first block to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
}

// Defines a 200 (List of contracts implement given trait) response for GetContractsByTrait.
type GetContractsByTraitResponse struct {
	// The number of contracts to return.
	Limit int `json:"limit"`
	// The number to contracts to skip (starting at `0`).
	Offset  int   `json:"offset"`
	Results []any `json:"results"`
}

// Defines parameters for GetContractsByTrait
type GetContractsByTraitParams struct {
	// Max number of contracts fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Index of first contract event to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// JSON abi of the trait.
	TraitABI string
}

// Defines a 200 (Transactions) response for GetUnanchoredTxs.
type GetUnanchoredTxsResponse struct {
	Results []any `json:"results"`
	// The number of unanchored transactions available.
	Total int `json:"total"`
}

// Defines a 200 (Success) response for GetNamespacePrice.
type GetNamespacePriceResponse struct {
	Amount string `json:"amount"`
	Units  string `json:"units"`
}

// Defines parameters for GetNamespacePrice
type GetNamespacePriceParams struct {
	// The namespace to fetch price for.
	TLD string
}

// Defines a 200 (List of microblocks) response for GetMicroblockList.
type GetMicroblockListResponse struct {
	// The number of microblocks to return.
	Limit int `json:"limit"`
	// The number to microblocks to skip (starting at `0`).
	Offset  int   `json:"offset"`
	Results []any `json:"results"`
	// The number of microblocks available.
	Total int `json:"total"`
}

// Defines parameters for GetMicroblockList
type GetMicroblockListParams struct {
	// Max number of microblocks to fetch.
	// Optional. Use `-1` to use the default value (20).
	// Max value is 200.
	Limit int
	// Index of the first microblock to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
}

// Defines a 200 (Success) response for GetSTXSupply.
type GetSTXSupplyResponse struct {
	// The block height at which this information was queried.
	BlockHeight int `json:"block_height"`
	// String quoted decimal number of the total possible number of STX.
	TotalSTX string `json:"total_stx"`
	// String quoted decimal number of the percentage of STX that have unlocked.
	UnlockedPercent string `json:"unlocked_percent"`
	// String quoted decimal number of the STX that have been mined or unlocked.
	UnlockedSTX string `json:"unlocked_stx"`
}

// Defines parameters for GetSTXSupply
type GetSTXSupplyParams struct {
	// The block height at which to query supply details from, if not provided then the latest block height is used.
	Height int
}

// Defines a 200 (List of Transactions) response for GetTransactionsByBlockHash.
type GetTransactionsByBlockHashResponse struct {
	// The number of transactions to return.
	Limit int `json:"limit"`
	// The number to transactions to skip (starting at `0`).
	Offset  int   `json:"offset"`
	Results []any `json:"results"`
	// The number of transactions available.
	Total int `json:"total"`
}

// Defines parameters for GetTransactionsByBlockHash
type GetTransactionsByBlockHashParams struct {
	// Max number of transactions to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Hash of block.
	BlockHash string
	// Index of first transaction to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
}

// Defines a 200 (Block) response for GetBlockByBurnBlockHeight.
type GetBlockByBurnBlockHeightResponse struct {
	// An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) indicating when this block was mined.
	BurnBlockTimeIso string `json:"burn_block_time_iso"`
	// Execution cost write length.
	ExecutionCostWriteLength int `json:"execution_cost_write_length"`
	// List of microblocks that were streamed/produced by this anchor block's miner. This list only includes microblocks that were accepted in the following anchor block. Microblocks that were orphaned are not included in this list.
	MicroblocksStreamed []string `json:"microblocks_streamed"`
	// The hash of the last streamed block that precedes this block to which this block is to be appended. Not every anchored block will have a parent microblock stream. An anchored block that does not have a parent microblock stream has the parent microblock hash set to an empty string, and the parent microblock sequence number set to -1.
	ParentMicroblockSequence int `json:"parent_microblock_sequence"`
	// Height of the anchor chain block.
	BurnBlockHeight int `json:"burn_block_height"`
	// Unix timestamp (in seconds) indicating when this block was mined.
	BurnBlockTime int `json:"burn_block_time"`
	// Execution cost write count.
	ExecutionCostWriteCount int `json:"execution_cost_write_count"`
	// The hash of the last streamed block that precedes this block to which this block is to be appended. Not every anchored block will have a parent microblock stream. An anchored block that does not have a parent microblock stream has the parent microblock hash set to an empty string, and the parent microblock sequence number set to -1.
	ParentMicroblockHash string `json:"parent_microblock_hash"`
	// Set to `true` if block corresponds to the canonical chain tip.
	Canonical bool `json:"canonical"`
	// Execution cost read count.
	ExecutionCostReadCount int `json:"execution_cost_read_count"`
	// Height of the block.
	Height int `json:"height"`
	// Anchor chain transaction ID.
	MinerTxid string `json:"miner_txid"`
	// List of transactions included in the block.
	Txs []string `json:"txs"`
	// Execution cost read length.
	ExecutionCostReadLength int `json:"execution_cost_read_length"`
	// Execution cost runtime.
	ExecutionCostRuntime int `json:"execution_cost_runtime"`
	// List of microblocks that were accepted in this anchor block. Not every anchored block will have a accepted all (or any) of the previously streamed microblocks. Microblocks that were orphaned are not included in this list.
	MicroblocksAccepted []string `json:"microblocks_accepted"`
	// Hash of the parent block.
	ParentBlockHash string `json:"parent_block_hash"`
	// Hash of the anchor chain block.
	BurnBlockHash string `json:"burn_block_hash"`
	// Hash representing the block.
	Hash string `json:"hash"`
}

// Defines a 200 (List of burnchain reward recipients and amounts) response for GetBurnchainRewardList.
type GetBurnchainRewardListResponse struct {
	// The number of burnchain rewards to return.
	Limit int `json:"limit"`
	// The number to burnchain rewards to skip (starting at `0`).
	Offset  int   `json:"offset"`
	Results []any `json:"results"`
}

// Defines parameters for GetBurnchainRewardList
type GetBurnchainRewardListParams struct {
	// Index of first rewards to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Max number of rewards to fetch.
	// Optional. Use `-1` to use the default value (96).
	// Max value is 250.
	Limit int
}

// Defines a 200 (List of events) response for GetContractEventsByID.
type GetContractEventsByIDResponse struct{}

// Defines parameters for GetContractEventsByID
type GetContractEventsByIDParams struct {
	// Max number of contract events to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Index of first contract event to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Contract identifier formatted as `<contract_address>.<contract_name>`.
	ContractID string
}

// Defines a 200 (Non-Fungible Token mints) response for GetNFTMints.
type GetNFTMintsResponse struct {
	// The number of mint events to return.
	Limit int `json:"limit"`
	// The number to mint events to skip (starting at `0`).
	Offset  int   `json:"offset"`
	Results []any `json:"results"`
	// The number of mint events available.
	Total int `json:"total"`
}

// Defines parameters for GetNFTMints
type GetNFTMintsParams struct {
	// Index of first event to fetch.
	// Optional. Use `-1` to use the default value (0).
	Offset int
	// Whether or not to include events from unconfirmed transactions.
	// Optional. Use `false` as default.
	Unanchored bool
	// Token asset class identifier.
	AssetIdentifier string
	// Whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times.
	// Optional. Use `false` as default.
	TXMetadata bool
	// Max number of events to fetch.
	// Optional. Use `-1` to use the default value (50).
	Limit int
}

// Defines a 200 (List of Transactions) response for GetTransactionsByBlockHeight.
type GetTransactionsByBlockHeightResponse struct {
	// The number of transactions available.
	Total int `json:"total"`
	// The number of transactions to return.
	Limit int `json:"limit"`
	// The number to transactions to skip (starting at `0`).
	Offset  int   `json:"offset"`
	Results []any `json:"results"`
}

// Defines parameters for GetTransactionsByBlockHeight
type GetTransactionsByBlockHeightParams struct {
	// Max number of transactions to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Index of first transaction to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Height of block.
	Height int
}

// Defines a 200 (Transaction) response for GetTransactionByID.
type GetTransactionByIDResponse struct{}

// Defines a 200 (Success) response for GetContractSource.
type GetContractSourceResponse struct {
	PublishHeight int    `json:"publish_height"`
	Source        string `json:"source"`
	Proof         string `json:"proof"`
}

// Defines a 200 (Success) response for GetAccountTransactionsWithTransfers.
type GetAccountTransactionsWithTransfersResponse struct {
	Total   int   `json:"total"`
	Limit   int   `json:"limit"`
	Offset  int   `json:"offset"`
	Results []any `json:"results"`
}

// Defines parameters for GetAccountTransactionsWithTransfers
type GetAccountTransactionsWithTransfersParams struct {
	// Index of first account transaction to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Filter for transactions only at this given block height.
	Height int
	// Max number of account transactions to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Returned data representing the state up until that point in time, rather than the current block.
	// Optional. Pass an empty string to use the default value.
	UntilBlock string
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
}

// Defines a 200 (Block) response for GetBlockByHash.
type GetBlockByHashResponse struct {
	// Hash of the anchor chain block.
	BurnBlockHash string `json:"burn_block_hash"`
	// Height of the block.
	Height int `json:"height"`
	// List of microblocks that were streamed/produced by this anchor block's miner. This list only includes microblocks that were accepted in the following anchor block. Microblocks that were orphaned are not included in this list.
	MicroblocksStreamed []string `json:"microblocks_streamed"`
	// Height of the anchor chain block.
	BurnBlockHeight int `json:"burn_block_height"`
	// An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) indicating when this block was mined.
	BurnBlockTimeIso string `json:"burn_block_time_iso"`
	// Set to `true` if block corresponds to the canonical chain tip.
	Canonical bool `json:"canonical"`
	// Execution cost read length.
	ExecutionCostReadLength int `json:"execution_cost_read_length"`
	// Execution cost runtime.
	ExecutionCostRuntime int `json:"execution_cost_runtime"`
	// Hash representing the block.
	Hash string `json:"hash"`
	// Hash of the parent block.
	ParentBlockHash string `json:"parent_block_hash"`
	// Unix timestamp (in seconds) indicating when this block was mined.
	BurnBlockTime int `json:"burn_block_time"`
	// Anchor chain transaction ID.
	MinerTxid string `json:"miner_txid"`
	// The hash of the last streamed block that precedes this block to which this block is to be appended. Not every anchored block will have a parent microblock stream. An anchored block that does not have a parent microblock stream has the parent microblock hash set to an empty string, and the parent microblock sequence number set to -1.
	ParentMicroblockHash string `json:"parent_microblock_hash"`
	// The hash of the last streamed block that precedes this block to which this block is to be appended. Not every anchored block will have a parent microblock stream. An anchored block that does not have a parent microblock stream has the parent microblock hash set to an empty string, and the parent microblock sequence number set to -1.
	ParentMicroblockSequence int `json:"parent_microblock_sequence"`
	// List of transactions included in the block.
	Txs []string `json:"txs"`
	// Execution cost read count.
	ExecutionCostReadCount int `json:"execution_cost_read_count"`
	// Execution cost write count.
	ExecutionCostWriteCount int `json:"execution_cost_write_count"`
	// Execution cost write length.
	ExecutionCostWriteLength int `json:"execution_cost_write_length"`
	// List of microblocks that were accepted in this anchor block. Not every anchored block will have a accepted all (or any) of the previously streamed microblocks. Microblocks that were orphaned are not included in this list.
	MicroblocksAccepted []string `json:"microblocks_accepted"`
}

// Defines a 200 (List of burnchain reward recipients and amounts) response for GetBurnchainRewardSlotHolders.
type GetBurnchainRewardSlotHoldersResponse struct {
	// The number of items to return.
	Limit int `json:"limit"`
	// The number of items to skip (starting at `0`).
	Offset  int   `json:"offset"`
	Results []any `json:"results"`
	// Total number of available items.
	Total int `json:"total"`
}

// Defines parameters for GetBurnchainRewardSlotHolders
type GetBurnchainRewardSlotHoldersParams struct {
	// Index of the first items to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Max number of items to fetch.
	// Optional. Use `-1` to use the default value (96).
	// Max value is 250.
	Limit int
}

// Defines a 200 (List of burnchain reward recipients and amounts) response for GetBurnchainRewardListByAddress.
type GetBurnchainRewardListByAddressResponse struct {
	// The number of burnchain rewards to return.
	Limit int `json:"limit"`
	// The number to burnchain rewards to skip (starting at `0`).
	Offset  int   `json:"offset"`
	Results []any `json:"results"`
}

// Defines parameters for GetBurnchainRewardListByAddress
type GetBurnchainRewardListByAddressParams struct {
	// Max number of rewards to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Index of first rewards to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Reward recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format.
	Address string
}

// Defines a 200 (Success) response for GetNetworkBlockTimeByNetwork.
type GetNetworkBlockTimeByNetworkResponse struct{}

// Defines parameters for GetNetworkBlockTimeByNetwork
type GetNetworkBlockTimeByNetworkParams struct {
	// Which network to retrieve the target block time of.
	Network string
}

// Defines a 200 (Success) response for GetAllNamespaces.
type GetAllNamespacesResponse struct {
	Namespaces []string `json:"namespaces"`
}

// Defines a 200 (Success) response for GetAccountBalance.
type GetAccountBalanceResponse struct {
	FungibleTokens    any `json:"fungible_tokens"`
	NonFungibleTokens any `json:"non_fungible_tokens"`
	STX               any `json:"stx"`
	// Token Offering Locked.
	TokenOfferingLocked any `json:"token_offering_locked"`
}

// Defines parameters for GetAccountBalance
type GetAccountBalanceParams struct {
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
	// Returned data representing the state up until that point in time, rather than the current block.
	// Optional. Pass an empty string to use the default value.
	UntilBlock string
}

// Defines a 200 (List of non fungible tokens metadata) response for GetNFTMetadataList.
type GetNFTMetadataListResponse struct {
	// The number of tokens metadata to return.
	Limit int `json:"limit"`
	// The number to tokens metadata to skip (starting at `0`).
	Offset  int   `json:"offset"`
	Results []any `json:"results"`
	// The number of tokens metadata available.
	Total int `json:"total"`
}

// Defines parameters for GetNFTMetadataList
type GetNFTMetadataListParams struct {
	// Index of first tokens to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Max number of tokens to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
}

// Defines a 200 (Success) response for GetAccountInfo.
type GetAccountInfoResponse struct {
	BalanceProof string `json:"balance_proof"`
	Locked       string `json:"locked"`
	Nonce        int    `json:"nonce"`
	NonceProof   string `json:"nonce_proof"`
	UnlockHeight int    `json:"unlock_height"`
	Balance      string `json:"balance"`
}

// Defines parameters for GetAccountInfo
type GetAccountInfoParams struct {
	// Returns object without the proof field if set to 0.
	// Optional. Use `-1` to use the default value.
	Proof int
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
	// The Stacks chain tip to query from.
	// Optional. Pass an empty string to use the default value.
	Tip string
}

// Defines a 200 (Hex encoded serialized transaction) response for GetRawTransactionByID.
type GetRawTransactionByIDResponse struct {
	// A hex encoded serialized transaction.
	RawTX string `json:"raw_tx"`
}
type PostFeeTransactionBody struct{}

/*
Get an estimated fee for the supplied transaction.  This
estimates the execution cost of the transaction, the current
fee rate of the network, and returns estimates for fee
amounts.
* transaction_payload is a hex-encoded serialization of
  the TransactionPayload for the transaction.
* estimated_len is an optional argument that provides the
  endpoint with an estimation of the final length (in bytes)
  of the transaction, including any post-conditions and
  signatures
If the node cannot provide an estimate for the transaction
(e.g., if the node has never seen a contract-call for the
given contract and function) or if estimation is not
configured on this node, a 400 response is returned.
The 400 response will be a JSON error containing a reason
field which can be one of the following:
* DatabaseError - this Stacks node has had an internal
  database error while trying to estimate the costs of the
  supplied transaction.
* NoEstimateAvailable - this Stacks node has not seen this
  kind of contract-call before, and it cannot provide an
  estimate yet.
* CostEstimationDisabled - this Stacks node does not perform
  fee or cost estimation, and it cannot respond on this
  endpoint.
The 200 response contains the following data:
* estimated_cost - the estimated multi-dimensional cost of
  executing the Clarity VM on the provided transaction.
* estimated_cost_scalar - a unitless integer that the Stacks
  node uses to compare how much of the block limit is consumed
  by different transactions. This value incorporates the
  estimated length of the transaction and the estimated
  execution cost of the transaction. The range of this integer
  may vary between different Stacks nodes. In order to compute
  an estimate of total fee amount for the transaction, this
  value is multiplied by the same Stacks node's estimated fee
  rate.
* cost_scalar_change_by_byte - a float value that indicates how
  much the estimated_cost_scalar value would increase for every
  additional byte in the final transaction.
* estimations - an array of estimated fee rates and total fees to
  pay in microSTX for the transaction. This array provides a range of
  estimates (default: 3) that may be used. Each element of the array
  contains the following fields:
    * fee_rate - the estimated value for the current fee
      rates in the network
    * fee - the estimated value for the total fee in
      microSTX that the given transaction should pay. These
      values are the result of computing:
      fee_rate x estimated_cost_scalar.
      If the estimated fees are less than the minimum relay
      fee (1 ustx x estimated_len), then that minimum relay
      fee will be returned here instead.
Note: If the final transaction's byte size is larger than
supplied to estimated_len, then applications should increase
this fee amount by:
  fee_rate x cost_scalar_change_by_byte x (final_size - estimated_size)
*/
func postFeeTransaction(server Server, body PostFeeTransactionBody) (PostFeeTransactionResponse, error) {
	url := fmt.Sprintf("%s%s", server, "/v2/fees/transaction")
	var returnedErr error
	sendBody, _ := json.Marshal(body)
	resp, status := makePostReq(url, string(sendBody), &returnedErr)

	if returnedErr != nil {
		return PostFeeTransactionResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj PostFeeTransactionResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return PostFeeTransactionResponse{}, err
		}
		return PostFeeTransactionResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Estimated fees for the transaction) response for PostFeeTransaction.
type PostFeeTransactionResponse struct {
	EstimatedCostScalar    int   `json:"estimated_cost_scalar"`
	Estimations            []any `json:"estimations"`
	CostScalarChangeByByte int   `json:"cost_scalar_change_by_byte"`
	EstimatedCost          any   `json:"estimated_cost"`
}

// Defines a 200 (Success) response for GetCoreAPIInfo.
type GetCoreAPIInfoResponse struct {
	// The best known block hash for the Stack chain (not including any pending microblocks).
	StacksTip string `json:"stacks_tip"`
	// The burn chain (i.e., bitcoin) consensus hash at the time that stacks_tip was mined.
	StacksTipConsensusHash string `json:"stacks_tip_consensus_hash"`
	// Is similar to peer_version and will be used to differentiate between different testnets. this value will be different between mainnet and testnet. once launched, this value will not change.
	NetworkID int `json:"network_id"`
	// Same as network_id, but for bitcoin.
	ParentNetworkID int `json:"parent_network_id"`
	// Is a hash used to identify the burnchain view for a node. it incorporates bitcoin chain information and PoX information. nodes that disagree on this value will appear to each other as forks. this value will change after every block.
	PoxConsensus string `json:"pox_consensus"`
	// Leftover from stacks 1.0, basically always burn_block_height - 1.
	StableBurnBlockHeight int `json:"stable_burn_block_height"`
	// Same as burn_consensus, but evaluated at stable_burn_block_height.
	StablePoxConsensus string `json:"stable_pox_consensus"`
	// The latest Stacks chain height. Stacks forks can occur independent of the Bitcoin chain, that height doesn't increase 1-to-1 with the Bitcoin height.
	StacksTipHeight int `json:"stacks_tip_height"`
	// The latest microblock hash if any microblocks were processed. if no microblock has been processed for the current block, a 000.., hex array is returned.
	UnanchoredTip string `json:"unanchored_tip"`
	// Latest bitcoin chain height.
	BurnBlockHeight int `json:"burn_block_height"`
	// The block height at which the testnet network will be reset. not applicable for mainnet.
	ExitAtBlockHeight int `json:"exit_at_block_height"`
	// Identifies the version number for the networking communication, this should not change while a node is running, and will only change if there's an upgrade.
	PeerVersion int `json:"peer_version"`
	// Is a version descriptor.
	ServerVersion string `json:"server_version"`
}

// Defines a 200 (Success) response for GetAccountTransactions.
type GetAccountTransactionsResponse struct {
	Limit   int   `json:"limit"`
	Offset  int   `json:"offset"`
	Results []any `json:"results"`
	Total   int   `json:"total"`
}

// Defines parameters for GetAccountTransactions
type GetAccountTransactionsParams struct {
	// Index of first account transaction to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Filter for transactions only at this given block height.
	Height int
	// Returned data representing the state up until that point in time, rather than the current block.
	// Optional. Pass an empty string to use the default value.
	UntilBlock string
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
	// Max number of account transactions to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
}

// Defines a 200 (List of burnchain reward recipients and amounts) response for GetBurnchainRewardSlotHoldersByAddress.
type GetBurnchainRewardSlotHoldersByAddressResponse struct {
	// The number of items to skip (starting at `0`).
	Offset  int   `json:"offset"`
	Results []any `json:"results"`
	// Total number of available items.
	Total int `json:"total"`
	// The number of items to return.
	Limit int `json:"limit"`
}

// Defines parameters for GetBurnchainRewardSlotHoldersByAddress
type GetBurnchainRewardSlotHoldersByAddressParams struct {
	// Max number of items to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Reward slot holder recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format.
	Address string
	// Index of the first items to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
}

// Defines a 200 (Contract found) response for GetContractByID.
type GetContractByIDResponse struct{}

// Defines a 200 (Success) response for SearchByID.
type SearchByIDResponse struct{}

// Defines parameters for SearchByID
type SearchByIDParams struct {
	// This includes the detailed data for purticular hash in the response.
	// Optional. Use `<nil>` as default.
	IncludeMetadata bool
	// The hex hash string for a block or transaction, account address, or contract address.
	ID string
}

// Defines a 200 (Success) response for GetTotalSTXSupplyLegacyFormat.
type GetTotalSTXSupplyLegacyFormatResponse struct {
	// String quoted decimal number of the STX that have been mined or unlocked.
	UnlockedSupply string `json:"unlockedSupply"`
	// Same as `unlockedSupply` but formatted with comma thousands separators.
	UnlockedSupplyFormatted string `json:"unlockedSupplyFormatted"`
	// The block height at which this information was queried.
	BlockHeight string `json:"blockHeight"`
	// String quoted decimal number of the total possible number of STX.
	TotalStacks string `json:"totalStacks"`
	// Same as `totalStacks` but formatted with comma thousands separators.
	TotalStacksFormatted string `json:"totalStacksFormatted"`
	// String quoted decimal number of the percentage of STX that have unlocked.
	UnlockedPercent string `json:"unlockedPercent"`
}

// Defines parameters for GetTotalSTXSupplyLegacyFormat
type GetTotalSTXSupplyLegacyFormatParams struct {
	// The block height at which to query supply details from, if not provided then the latest block height is used.
	Height int
}

// Defines a 200 (Success) response for FetchZoneFile.
type FetchZoneFileResponse struct{}

// Defines parameters for FetchZoneFile
type FetchZoneFileParams struct {
	// Fully-qualified name.
	Name string
}

// Defines a 200 (List of fungible tokens metadata) response for GetFTMetadataList.
type GetFTMetadataListResponse struct {
	// The number of tokens metadata to return.
	Limit int `json:"limit"`
	// The number to tokens metadata to skip (starting at `0`).
	Offset  int   `json:"offset"`
	Results []any `json:"results"`
	// The number of tokens metadata available.
	Total int `json:"total"`
}

// Defines parameters for GetFTMetadataList
type GetFTMetadataListParams struct {
	// Max number of tokens to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Index of first tokens to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
}

// Defines a 200 (Success) response for GetAllNames.
type GetAllNamesResponse struct{}

// Defines parameters for GetAllNames
type GetAllNamesParams struct {
	// Names are returned in pages of size 100, so specify the page number.
	Page int
}

// Defines a 200 (Success) response for GetFilteredEvents.
type GetFilteredEventsResponse struct {
	Limit   int   `json:"limit"`
	Offset  int   `json:"offset"`
	Results []any `json:"results"`
}
type GetFilteredEventsType int64

const (
	SmartContractLog GetFilteredEventsType = iota
	STXLock
	STXAsset
	FungibleTokenAsset
	NonFungibleTokenAsset
)

// Defines parameters for GetFilteredEvents
type GetFilteredEventsParams struct {
	// Number of items to return.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Number of items to skip.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Stacks address or a Contract identifier.
	// Optional. Pass an empty string to use the default value.
	Address string
	// Filter the events on event type.
	// [SmartContractLog STXLock STXAsset FungibleTokenAsset NonFungibleTokenAsset]
	Type GetFilteredEventsType
}

// Defines a 200 (List of Transactions) response for GetAddressMempoolTransactions.
type GetAddressMempoolTransactionsResponse struct {
	Limit   int   `json:"limit"`
	Offset  int   `json:"offset"`
	Results []any `json:"results"`
	Total   int   `json:"total"`
}

// Defines parameters for GetAddressMempoolTransactions
type GetAddressMempoolTransactionsParams struct {
	// Max number of transactions to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Index of first transaction to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Transactions for the address.
	Address string
}

// Defines a 200 (List of burnchain reward recipients and amounts) response for GetBurnchainRewardsTotalByAddress.
type GetBurnchainRewardsTotalByAddressResponse struct {
	// The total amount of burnchain tokens rewarded to the recipient, in the smallest unit (e.g. satoshis for Bitcoin).
	RewardAmount string `json:"reward_amount"`
	// The recipient address that received the burnchain rewards, in the format native to the burnchain (e.g. B58 encoded for Bitcoin).
	RewardRecipient string `json:"reward_recipient"`
}

// Defines parameters for GetBurnchainRewardsTotalByAddress
type GetBurnchainRewardsTotalByAddressParams struct {
	// Reward recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format.
	Address string
}

// Defines a 200 (Success) response for GetNetworkBlockTimes.
type GetNetworkBlockTimesResponse struct {
	Mainnet any `json:"mainnet"`
	Testnet any `json:"testnet"`
}

// Defines a 200 (Success) response for GetStatus.
type GetStatusResponse struct {
	// Current chain tip information.
	ChainTip any `json:"chain_tip"`
	// The server version that is currently running.
	ServerVersion string `json:"server_version"`
	// The current server status.
	Status string `json:"status"`
}

// Defines a 200 (success) response for GetSTXSupplyTotalSupplyPlain.
type GetSTXSupplyTotalSupplyPlainResponse struct{}

// Defines a 200 (Non fungible tokens metadata for contract id) response for GetContractNFTMetadata.
type GetContractNFTMetadataResponse struct {
	// Tx id that deployed the contract.
	TXID string `json:"tx_id"`
	// Describes the asset to which this token represents.
	Description string `json:"description"`
	// The original image URI specified by the contract. A URI pointing to a resource with mime type image/* representing the asset to which this token represents. Consider making any images at a width between 320 and 1080 pixels and aspect ratio between 1.91:1 and 4:5 inclusive.
	ImageCanonicalUri string `json:"image_canonical_uri"`
	// A URI pointing to a resource with mime type image/* representing the asset to which this token represents. The API may provide a URI to a cached resource, dependending on configuration. Otherwise, this can be the same value as the canonical image URI.
	ImageUri string `json:"image_uri"`
	// Identifies the asset to which this token represents.
	Name string `json:"name"`
	// Principle that deployed the contract.
	SenderAddress string `json:"sender_address"`
	// An optional string that is a valid URI which resolves to this token's metadata. Can be empty.
	TokenUri string `json:"token_uri"`
}

// Defines parameters for GetContractNFTMetadata
type GetContractNFTMetadataParams struct {
	// Token's contract id.
	ContractId string
}

// Defines a 200 (List of transactions) response for GetTransactionList.
type GetTransactionListResponse struct {
	// The number of transactions to return.
	Limit int `json:"limit"`
	// The number to transactions to skip (starting at `0`).
	Offset  int   `json:"offset"`
	Results []any `json:"results"`
	// The number of transactions available.
	Total int `json:"total"`
}
type GetTransactionListType int64

const (
	Coinbase GetTransactionListType = iota
	TokenTransfer
	SmartContract
	ContractCall
	PoisonMicroblock
)

// Defines parameters for GetTransactionList
type GetTransactionListParams struct {
	// Index of first transaction to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Max number of transactions to fetch.
	// Optional. Use `-1` to use the default value (96).
	// Max value is 200.
	Limit int
	// Filter by transaction type.
	// [Coinbase TokenTransfer SmartContract ContractCall PoisonMicroblock]
	Type GetTransactionListType
}

// Defines a 200 (Success) response for GetNamespaceNames.
type GetNamespaceNamesResponse struct{}

// Defines parameters for GetNamespaceNames
type GetNamespaceNamesParams struct {
	// The namespace to fetch names from.
	TLD string
	// Names are returned in pages of size 100, so specify the page number.
	Page int
}

// Defines a 200 (Success) response for GetHistoricalZoneFile.
type GetHistoricalZoneFileResponse struct{}

// Defines parameters for GetHistoricalZoneFile
type GetHistoricalZoneFileParams struct {
	// Zone file hash.
	ZoneFileHash string
	// Fully-qualified name.
	Name string
}

// Defines a 200 (Contract interface) response for GetContractInterface.
type GetContractInterfaceResponse struct {
	// List of defined data-maps.
	Maps []any `json:"maps"`
	// List of non-fungible tokens in the contract.
	NonFungibleTokens []any `json:"non_fungible_tokens"`
	// List of defined variables.
	Variables []any `json:"variables"`
	// List of defined methods.
	Functions []any `json:"functions"`
	// List of fungible tokens in the contract.
	FungibleTokens []any `json:"fungible_tokens"`
}
type GetContractDataMapEntryBody struct{}

/*
Attempt to fetch data from a contract data map. The contract is identified with Stacks Address contract_address and Contract Name contract_address in the URL path. The map is identified with [Map Name].

The key to lookup in the map is supplied via the POST body. This should be supplied as the hex string serialization of the key (which should be a Clarity value). Note, this is a JSON string atom.

In the response, data is the hex serialization of the map response. Note that map responses are Clarity option types, for non-existent values, this is a serialized none, and for all other responses, it is a serialized (some ...) object.
*/
func getContractDataMapEntry(server Server, params GetContractDataMapEntryParams, body GetContractDataMapEntryBody) (GetContractDataMapEntryResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/v2/map_entry/{contract_address}/{contract_name}/{map_name}", params))
	var returnedErr error
	sendBody, _ := json.Marshal(body)
	resp, status := makePostReq(url, string(sendBody), &returnedErr)

	if returnedErr != nil {
		return GetContractDataMapEntryResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetContractDataMapEntryResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetContractDataMapEntryResponse{}, err
		}
		return GetContractDataMapEntryResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for GetContractDataMapEntry.
type GetContractDataMapEntryResponse struct {
	// Hex-encoded string of clarity value. It is always an optional tuple.
	Data string `json:"data"`
	// Hex-encoded string of the MARF proof for the data.
	Proof string `json:"proof"`
}

// Defines parameters for GetContractDataMapEntry
type GetContractDataMapEntryParams struct {
	// Map name.
	MapName string
	// Stacks address.
	ContractAddress string
	// Returns object without the proof field when set to 0.
	// Optional. Use `-1` to use the default value.
	Proof int
	// Contract name.
	ContractName string
	// The Stacks chain tip to query from.
	// Optional. Pass an empty string to use the default value.
	Tip string
}

// Defines a 200 (success) response for GetSTXSupplyCirculatingPlain.
type GetSTXSupplyCirculatingPlainResponse struct{}

// Defines a 200 (List of Non-Fungible Token holdings) response for GetNFTHoldings.
type GetNFTHoldingsResponse struct {
	// The number of Non-Fungible Token holdings to return.
	Limit int `json:"limit"`
	// The number to Non-Fungible Token holdings to skip (starting at `0`).
	Offset  int   `json:"offset"`
	Results []any `json:"results"`
	// The number of Non-Fungible Token holdings available.
	Total int `json:"total"`
}

// Defines parameters for GetNFTHoldings
type GetNFTHoldingsParams struct {
	// Max number of tokens to fetch.
	// Optional. Use `-1` to use the default value (50).
	Limit int
	// Whether or not to include tokens from unconfirmed transactions.
	// Optional. Use `false` as default.
	Unanchored bool
	// Index of first tokens to fetch.
	// Optional. Use `-1` to use the default value (0).
	Offset int
	// Identifiers of the token asset classes to filter for.
	// Token owner's STX address or Smart Contract ID.
	Principal string
	// Whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times.
	// Optional. Use `false` as default.
	TXMetadata bool
}

// Defines a 200 (Fungible tokens metadata for contract id) response for GetContractFTMetadata.
type GetContractFTMetadataResponse struct {
	// The original image URI specified by the contract. A URI pointing to a resource with mime type image/* representing the asset to which this token represents. Consider making any images at a width between 320 and 1080 pixels and aspect ratio between 1.91:1 and 4:5 inclusive.
	ImageCanonicalUri string `json:"image_canonical_uri"`
	// A shorter representation of a token. This is sometimes referred to as a "ticker". Examples: "STX", "COOL", etc. Typically, a token could be referred to as $SYMBOL when referencing it in writing.
	Symbol string `json:"symbol"`
	// The number of decimal places in a token.
	Decimals int `json:"decimals"`
	// Describes the asset to which this token represents.
	Description string `json:"description"`
	// A URI pointing to a resource with mime type image/* representing the asset to which this token represents. The API may provide a URI to a cached resource, dependending on configuration. Otherwise, this can be the same value as the canonical image URI.
	ImageUri string `json:"image_uri"`
	// Identifies the asset to which this token represents.
	Name string `json:"name"`
	// Principle that deployed the contract.
	SenderAddress string `json:"sender_address"`
	// An optional string that is a valid URI which resolves to this token's metadata. Can be empty.
	TokenUri string `json:"token_uri"`
	// Tx id that deployed the contract.
	TXID string `json:"tx_id"`
}

// Defines parameters for GetContractFTMetadata
type GetContractFTMetadataParams struct {
	// Token's contract id.
	ContractId string
}
type PostCoreNodeTransactionsBody struct{}

// Defines a 200 (Transaction id of successful post of a raw tx to the node's mempool) response for PostCoreNodeTransactions.
type PostCoreNodeTransactionsResponse struct{}

// Broadcasts raw transactions on the network. You can use the [@stacks/transactions](https://github.com/blockstack/stacks.js) project to generate a raw transaction payload.
func postCoreNodeTransactions(server Server, body PostCoreNodeTransactionsBody) (PostCoreNodeTransactionsResponse, error) {
	url := fmt.Sprintf("%s%s", server, "/v2/transactions")
	var returnedErr error
	sendBody, _ := json.Marshal(body)
	resp, status := makePostReq(url, string(sendBody), &returnedErr)

	if returnedErr != nil {
		return PostCoreNodeTransactionsResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj PostCoreNodeTransactionsResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return PostCoreNodeTransactionsResponse{}, err
		}
		return PostCoreNodeTransactionsResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for GetAccountNFT.
type GetAccountNFTResponse struct {
	Offset    int   `json:"offset"`
	Total     int   `json:"total"`
	Limit     int   `json:"limit"`
	NFTEvents []any `json:"nft_events"`
}

// Defines parameters for GetAccountNFT
type GetAccountNFTParams struct {
	// Number of items to skip.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
	// Returned data representing the state up until that point in time, rather than the current block.
	// Optional. Pass an empty string to use the default value.
	UntilBlock string
	// Number of items to return.
	// Optional. Use `-1` to use the default value.
	Limit int
}

// Defines a 200 (Success) response for GetAccountInbound.
type GetAccountInboundResponse struct {
	Offset  int   `json:"offset"`
	Results []any `json:"results"`
	Total   int   `json:"total"`
	Limit   int   `json:"limit"`
}

// Defines parameters for GetAccountInbound
type GetAccountInboundParams struct {
	// Number of items to skip.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
	// Number of items to return.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Returned data representing the state up until that point in time, rather than the current block.
	// Optional. Pass an empty string to use the default value.
	UntilBlock string
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Filter for transfers only at this given block height.
	Height int
}
type RunFaucetBTCBody struct{}

/*
Add 1 BTC token to the specified testnet BTC address.

The endpoint returns the transaction ID, which you can use to view the transaction in a testnet Bitcoin block
explorer. The tokens are delivered once the transaction has been included in a block.

**Note:** This is a testnet only endpoint. This endpoint will not work on the mainnet.
*/
func runFaucetBTC(server Server, body RunFaucetBTCBody) (RunFaucetBTCResponse, error) {
	url := fmt.Sprintf("%s%s", server, "/extended/v1/faucets/btc")
	var returnedErr error
	sendBody, _ := json.Marshal(body)
	resp, status := makePostReq(url, string(sendBody), &returnedErr)

	if returnedErr != nil {
		return RunFaucetBTCResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj RunFaucetBTCResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return RunFaucetBTCResponse{}, err
		}
		return RunFaucetBTCResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for RunFaucetBTC.
type RunFaucetBTCResponse struct {
	// Indicates if the faucet call was successful.
	Success bool `json:"success"`
	// The transaction ID for the faucet call.
	TxId string `json:"txId"`
	// Raw transaction in hex string representation.
	TxRaw string `json:"txRaw"`
}

// Defines a 200 (Returns list of transactions with their details for corresponding requested tx_ids.) response for GetTXListDetails.
type GetTXListDetailsResponse struct{}

// Defines a 200 (Block) response for GetBlockByBurnBlockHash.
type GetBlockByBurnBlockHashResponse struct {
	// List of transactions included in the block.
	Txs []string `json:"txs"`
	// Execution cost runtime.
	ExecutionCostRuntime int `json:"execution_cost_runtime"`
	// List of microblocks that were accepted in this anchor block. Not every anchored block will have a accepted all (or any) of the previously streamed microblocks. Microblocks that were orphaned are not included in this list.
	MicroblocksAccepted []string `json:"microblocks_accepted"`
	// Hash of the parent block.
	ParentBlockHash string `json:"parent_block_hash"`
	// The hash of the last streamed block that precedes this block to which this block is to be appended. Not every anchored block will have a parent microblock stream. An anchored block that does not have a parent microblock stream has the parent microblock hash set to an empty string, and the parent microblock sequence number set to -1.
	ParentMicroblockSequence int `json:"parent_microblock_sequence"`
	// Height of the block.
	Height int `json:"height"`
	// Hash of the anchor chain block.
	BurnBlockHash string `json:"burn_block_hash"`
	// Height of the anchor chain block.
	BurnBlockHeight int `json:"burn_block_height"`
	// Set to `true` if block corresponds to the canonical chain tip.
	Canonical bool `json:"canonical"`
	// Execution cost read count.
	ExecutionCostReadCount int `json:"execution_cost_read_count"`
	// Unix timestamp (in seconds) indicating when this block was mined.
	BurnBlockTime int `json:"burn_block_time"`
	// Execution cost write length.
	ExecutionCostWriteLength int `json:"execution_cost_write_length"`
	// Hash representing the block.
	Hash string `json:"hash"`
	// Anchor chain transaction ID.
	MinerTxid string `json:"miner_txid"`
	// The hash of the last streamed block that precedes this block to which this block is to be appended. Not every anchored block will have a parent microblock stream. An anchored block that does not have a parent microblock stream has the parent microblock hash set to an empty string, and the parent microblock sequence number set to -1.
	ParentMicroblockHash string `json:"parent_microblock_hash"`
	// An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) indicating when this block was mined.
	BurnBlockTimeIso string `json:"burn_block_time_iso"`
	// Execution cost read length.
	ExecutionCostReadLength int `json:"execution_cost_read_length"`
	// Execution cost write count.
	ExecutionCostWriteCount int `json:"execution_cost_write_count"`
	// List of microblocks that were streamed/produced by this anchor block's miner. This list only includes microblocks that were accepted in the following anchor block. Microblocks that were orphaned are not included in this list.
	MicroblocksStreamed []string `json:"microblocks_streamed"`
}

// Defines a 200 (Success) response for FetchSubdomainsListForName.
type FetchSubdomainsListForNameResponse struct{}

// Defines parameters for FetchSubdomainsListForName
type FetchSubdomainsListForNameParams struct {
	// Fully-qualified name.
	Name string
}
type CallReadOnlyFunctionBody struct{}

/*
Call a read-only public function on a given smart contract.

The smart contract and function are specified using the URL path. The arguments and the simulated tx-sender are supplied via the POST body in the following JSON format:
*/
func callReadOnlyFunction(server Server, params CallReadOnlyFunctionParams, body CallReadOnlyFunctionBody) (CallReadOnlyFunctionResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/v2/contracts/call-read/{contract_address}/{contract_name}/{function_name}", params))
	var returnedErr error
	sendBody, _ := json.Marshal(body)
	resp, status := makePostReq(url, string(sendBody), &returnedErr)

	if returnedErr != nil {
		return CallReadOnlyFunctionResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj CallReadOnlyFunctionResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return CallReadOnlyFunctionResponse{}, err
		}
		return CallReadOnlyFunctionResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for CallReadOnlyFunction.
type CallReadOnlyFunctionResponse struct {
	Result string `json:"result"`
	Cause  string `json:"cause"`
	Okay   bool   `json:"okay"`
}

// Defines parameters for CallReadOnlyFunction
type CallReadOnlyFunctionParams struct {
	// Function name.
	FunctionName string
	// Stacks address.
	ContractAddress string
	// Contract name.
	ContractName string
	// The Stacks chain tip to query from.
	// Optional. Pass an empty string to use the default value.
	Tip string
}

// Defines a 200 (Success) response for GetFeeTransfer.
type GetFeeTransferResponse struct{}

// Defines a 200 (Success) response for GetAccountNonces.
type GetAccountNoncesResponse struct {
	// Nonces that appear to be missing and likely causing a mempool transaction to be stuck.
	DetectedMissingNonces []int `json:"detected_missing_nonces"`
	// The latest nonce found within transactions sent by this address, including unanchored microblock transactions. Will be null if there are no current transactions for this address.
	LastExecutedTXNonce int `json:"last_executed_tx_nonce"`
	// The latest nonce found within mempool transactions sent by this address. Will be null if there are no current mempool transactions for this address.
	LastMempoolTXNonce int `json:"last_mempool_tx_nonce"`
	// The likely nonce required for creating the next transaction, based on the last nonces seen by the API. This can be incorrect if the API's mempool or transactions aren't fully synchronized, even by a small amount, or if a previous transaction is still propagating through the Stacks blockchain network when this endpoint is called.
	PossibleNextNonce int `json:"possible_next_nonce"`
}

// Defines parameters for GetAccountNonces
type GetAccountNoncesParams struct {
	// Optionally get the nonce at a given block hash.
	// Optional. Pass an empty string to use the default value.
	BlockHash string
	// Stacks address (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0`).
	Principal string
	// Optionally get the nonce at a given block height.
	BlockHeight int
}

// Defines a 200 (Success) response for GetAccountSTXBalance.
type GetAccountSTXBalanceResponse struct{}

// Defines parameters for GetAccountSTXBalance
type GetAccountSTXBalanceParams struct {
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Returned data representing the state up until that point in time, rather than the current block.
	// Optional. Pass an empty string to use the default value.
	UntilBlock string
}

// Defines a 200 (Success) response for GetSingleTransactionWithTransfers.
type GetSingleTransactionWithTransfersResponse struct {
	// Total received by the given address in micro-STX as an integer string.
	STXReceived string `json:"stx_received"`
	// Total sent from the given address, including the tx fee, in micro-STX as an integer string.
	STXSent      string `json:"stx_sent"`
	STXTransfers []any  `json:"stx_transfers"`
	// Describes all transaction types on Stacks 2.0 blockchain.
	TX           any   `json:"tx"`
	FTTransfers  []any `json:"ft_transfers"`
	NFTTransfers []any `json:"nft_transfers"`
}

// Defines parameters for GetSingleTransactionWithTransfers
type GetSingleTransactionWithTransfersParams struct {
	// Stacks address or a contract identifier.
	Principal string
	// Transaction id.
	TXID string
}

// Defines a 200 (Non-Fungible Token event history) response for GetNFTHistory.
type GetNFTHistoryResponse struct {
	// The number of events to return.
	Limit int `json:"limit"`
	// The number to events to skip (starting at `0`).
	Offset  int   `json:"offset"`
	Results []any `json:"results"`
	// The number of events available.
	Total int `json:"total"`
}

// Defines parameters for GetNFTHistory
type GetNFTHistoryParams struct {
	// Max number of events to fetch.
	// Optional. Use `-1` to use the default value (50).
	Limit int
	// Token asset class identifier.
	AssetIdentifier string
	// Whether or not to include events from unconfirmed transactions.
	// Optional. Use `false` as default.
	Unanchored bool
	// Hex representation of the token's unique value.
	Value string
	// Index of first event to fetch.
	// Optional. Use `-1` to use the default value (0).
	Offset int
	// Whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times.
	// Optional. Use `false` as default.
	TXMetadata bool
}

// Defines a 200 (Success) response for GetPoxInfo.
type GetPoxInfoResponse struct {
	ContractID                 string `json:"contract_id"`
	FirstBurnchainBlockHeight  int    `json:"first_burnchain_block_height"`
	MinAmountUstx              int    `json:"min_amount_ustx"`
	RegistrationWindowLength   int    `json:"registration_window_length"`
	RejectionVotesLeftRequired int    `json:"rejection_votes_left_required"`
	RewardCycleID              int    `json:"reward_cycle_id"`
	RejectionFraction          int    `json:"rejection_fraction"`
	RewardCycleLength          int    `json:"reward_cycle_length"`
	TotalLiquidSupplyUstx      int    `json:"total_liquid_supply_ustx"`
}

// Defines a 200 (Success) response for GetNamePrice.
type GetNamePriceResponse struct {
	Amount string `json:"amount"`
	Units  string `json:"units"`
}

// Defines parameters for GetNamePrice
type GetNamePriceParams struct {
	// The name to query price information for.
	Name string
}

// Defines a 200 (Success) response for GetAccountAssets.
type GetAccountAssetsResponse struct {
	Offset  int   `json:"offset"`
	Results []any `json:"results"`
	Total   int   `json:"total"`
	Limit   int   `json:"limit"`
}

// Defines parameters for GetAccountAssets
type GetAccountAssetsParams struct {
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
	// Index of first account assets to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Returned data representing the state at that point in time, rather than the current block.
	// Optional. Pass an empty string to use the default value.
	UntilBlock string
	// Max number of account assets to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
}
type RunFaucetSTXBody struct{}

/*
Add 500 STX tokens to the specified testnet address. Testnet STX addresses begin with ST. If the stacking
parameter is set to true, the faucet will add the required number of tokens for individual stacking to the
specified testnet address.

The endpoint returns the transaction ID, which you can use to view the transaction in the
[Stacks Explorer](https://explorer.stacks.co/?chain=testnet). The tokens are delivered once the transaction has
been included in an anchor block.

A common reason for failed faucet transactions is that the faucet has run out of tokens. If you are experiencing
failed faucet transactions to a testnet address, you can get help in [Discord](https://stacks.chat).

**Note:** This is a testnet only endpoint. This endpoint will not work on the mainnet.
*/
func runFaucetSTX(server Server, body RunFaucetSTXBody) (RunFaucetSTXResponse, error) {
	url := fmt.Sprintf("%s%s", server, "/extended/v1/faucets/stx")
	var returnedErr error
	sendBody, _ := json.Marshal(body)
	resp, status := makePostReq(url, string(sendBody), &returnedErr)

	if returnedErr != nil {
		return RunFaucetSTXResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj RunFaucetSTXResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return RunFaucetSTXResponse{}, err
		}
		return RunFaucetSTXResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for RunFaucetSTX.
type RunFaucetSTXResponse struct {
	// Indicates if the faucet call was successful.
	Success bool `json:"success"`
	// The transaction ID for the faucet call.
	TxId string `json:"txId"`
	// Raw transaction in hex string representation.
	TxRaw string `json:"txRaw"`
}

// Defines a 200 (List of dropped mempool transactions) response for GetDroppedMempoolTransactionList.
type GetDroppedMempoolTransactionListResponse struct {
	Results []any `json:"results"`
	Total   int   `json:"total"`
	Limit   int   `json:"limit"`
	Offset  int   `json:"offset"`
}

// Defines parameters for GetDroppedMempoolTransactionList
type GetDroppedMempoolTransactionListParams struct {
	// Max number of mempool transactions to fetch.
	// Optional. Use `-1` to use the default value (96).
	// Max value is 200.
	Limit int
	// Index of first mempool transaction to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
}

// Defines a 200 (Success) response for GetNamesOwnedByAddress.
type GetNamesOwnedByAddressResponse struct {
	Names []string `json:"names"`
}

// Defines parameters for GetNamesOwnedByAddress
type GetNamesOwnedByAddressParams struct {
	// The layer-1 blockchain for the address.
	Blockchain string
	// The address to lookup.
	Address string
}

// Defines a 200 (Block) response for GetBlockByHeight.
type GetBlockByHeightResponse struct {
	// Height of the anchor chain block.
	BurnBlockHeight int `json:"burn_block_height"`
	// Execution cost write length.
	ExecutionCostWriteLength int `json:"execution_cost_write_length"`
	// The hash of the last streamed block that precedes this block to which this block is to be appended. Not every anchored block will have a parent microblock stream. An anchored block that does not have a parent microblock stream has the parent microblock hash set to an empty string, and the parent microblock sequence number set to -1.
	ParentMicroblockSequence int `json:"parent_microblock_sequence"`
	// Unix timestamp (in seconds) indicating when this block was mined.
	BurnBlockTime int `json:"burn_block_time"`
	// An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) indicating when this block was mined.
	BurnBlockTimeIso string `json:"burn_block_time_iso"`
	// Execution cost runtime.
	ExecutionCostRuntime int `json:"execution_cost_runtime"`
	// Execution cost write count.
	ExecutionCostWriteCount int `json:"execution_cost_write_count"`
	// Height of the block.
	Height int `json:"height"`
	// The hash of the last streamed block that precedes this block to which this block is to be appended. Not every anchored block will have a parent microblock stream. An anchored block that does not have a parent microblock stream has the parent microblock hash set to an empty string, and the parent microblock sequence number set to -1.
	ParentMicroblockHash string `json:"parent_microblock_hash"`
	// Hash of the anchor chain block.
	BurnBlockHash string `json:"burn_block_hash"`
	// Set to `true` if block corresponds to the canonical chain tip.
	Canonical bool `json:"canonical"`
	// Execution cost read length.
	ExecutionCostReadLength int `json:"execution_cost_read_length"`
	// Hash representing the block.
	Hash string `json:"hash"`
	// List of microblocks that were streamed/produced by this anchor block's miner. This list only includes microblocks that were accepted in the following anchor block. Microblocks that were orphaned are not included in this list.
	MicroblocksStreamed []string `json:"microblocks_streamed"`
	// Anchor chain transaction ID.
	MinerTxid string `json:"miner_txid"`
	// Execution cost read count.
	ExecutionCostReadCount int `json:"execution_cost_read_count"`
	// List of microblocks that were accepted in this anchor block. Not every anchored block will have a accepted all (or any) of the previously streamed microblocks. Microblocks that were orphaned are not included in this list.
	MicroblocksAccepted []string `json:"microblocks_accepted"`
	// Hash of the parent block.
	ParentBlockHash string `json:"parent_block_hash"`
	// List of transactions included in the block.
	Txs []string `json:"txs"`
}
type FetchFeeRateBody struct{}

/*
**NOTE:** This endpoint is deprecated in favor of [Get approximate fees for the given transaction](#operation/post_fee_transaction).

Retrieves estimated fee rate.
*/
func fetchFeeRate(server Server, body FetchFeeRateBody) (FetchFeeRateResponse, error) {
	url := fmt.Sprintf("%s%s", server, "/extended/v1/fee_rate")
	var returnedErr error
	sendBody, _ := json.Marshal(body)
	resp, status := makePostReq(url, string(sendBody), &returnedErr)

	if returnedErr != nil {
		return FetchFeeRateResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj FetchFeeRateResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return FetchFeeRateResponse{}, err
		}
		return FetchFeeRateResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Transaction fee rate) response for FetchFeeRate.
type FetchFeeRateResponse struct {
	FeeRate int `json:"fee_rate"`
}

// Defines a 200 (Microblock) response for GetMicroblockByHash.
type GetMicroblockByHashResponse struct {
	// A hint to describe how to order a set of microblocks. Starts at 0.
	MicroblockSequence int `json:"microblock_sequence"`
	// The hash of the Bitcoin block that preceded this microblock.
	ParentBurnBlockHash string `json:"parent_burn_block_hash"`
	// The height of the Bitcoin block that preceded this microblock.
	ParentBurnBlockHeight int `json:"parent_burn_block_height"`
	// The block timestamp of the Bitcoin block that preceded this microblock.
	ParentBurnBlockTime int `json:"parent_burn_block_time"`
	// The ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) formatted block time of the bitcoin block that preceded this microblock.
	ParentBurnBlockTimeIso string `json:"parent_burn_block_time_iso"`
	// List of transactions included in the microblock.
	Txs []string `json:"txs"`
	// The SHA512/256 hash of the previous signed microblock in this stream.
	MicroblockParentHash string `json:"microblock_parent_hash"`
	// The anchor block height that confirmed this microblock.
	BlockHeight int `json:"block_height"`
	// Set to `true` if the microblock was not orphaned in a following anchor block. Defaults to `true` if the following anchor block has not yet been created.
	MicroblockCanonical bool `json:"microblock_canonical"`
	// The hash of the anchor block that confirmed this microblock. This wil be empty for unanchored microblocks.
	BlockHash string `json:"block_hash"`
	// The height of the anchor block that preceded this microblock.
	ParentBlockHeight int `json:"parent_block_height"`
	// The hash of the anchor block that preceded this microblock.
	ParentBlockHash string `json:"parent_block_hash"`
	// The SHA512/256 hash of this microblock.
	MicroblockHash string `json:"microblock_hash"`
	// Set to `true` if the microblock corresponds to the canonical chain tip.
	Canonical bool `json:"canonical"`
}

// Defines a 200 (List of mempool transactions) response for GetMempoolTransactionList.
type GetMempoolTransactionListResponse struct {
	Total   int   `json:"total"`
	Limit   int   `json:"limit"`
	Offset  int   `json:"offset"`
	Results []any `json:"results"`
}

// Defines parameters for GetMempoolTransactionList
type GetMempoolTransactionListParams struct {
	// Filter to only return transactions with this sender address.
	// Optional. Pass an empty string to use the default value.
	SenderAddress string
	// Filter to only return transactions with this recipient address (only applicable for STX transfer tx types).
	// Optional. Pass an empty string to use the default value.
	RecipientAddress string
	// Filter to only return transactions with this address as the sender or recipient (recipient only applicable for STX transfer tx types).
	// Optional. Pass an empty string to use the default value.
	Address string
	// Max number of mempool transactions to fetch.
	// Optional. Use `-1` to use the default value (96).
	// Max value is 200.
	Limit int
	// Index of first mempool transaction to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
}

// Defines a 200 (Success) response for GetNameInfo.
type GetNameInfoResponse struct {
	Resolver     string `json:"resolver"`
	Status       string `json:"status"`
	Address      string `json:"address"`
	Blockchain   string `json:"blockchain"`
	GracePeriod  int    `json:"grace_period"`
	LastTxid     string `json:"last_txid"`
	ExpireBlock  int    `json:"expire_block"`
	Zonefile     string `json:"zonefile"`
	ZonefileHash string `json:"zonefile_hash"`
}

// Defines parameters for GetNameInfo
type GetNameInfoParams struct {
	// Fully-qualified name.
	Name string
}
type CallReadOnlyFunctionArguments struct {
	// The simulated tx-sender
	Sender string
	// An array of hex serialized Clarity values
	Arguments []string
}
