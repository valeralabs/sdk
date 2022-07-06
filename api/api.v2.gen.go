package api

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type Server string

const (
	ValeraMN Server = "mn.stx.valera.sh"
	ValeraTN Server = "tn.stx.valera.sh"
	HiroMN   Server = "stacks-node-api.mainnet.stacks.co"
	HiroTN   Server = "stacks-node-api.testnet.stacks.co"
)

func check(err error, save *error) {
	if err != nil {
		*save = err
	}
}

func makeGetReq(url string, returnedErr *error) ([]byte, int) {
	*returnedErr = nil
	resp, err := http.Get(url)
	check(err, returnedErr)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err, returnedErr)

	return body, resp.StatusCode
}

func makePostReq(url string, sendBody string, returnedErr *error) ([]byte, int) {
	*returnedErr = nil
	resp, err := http.Post(url, "application/json", strings.NewReader(sendBody))
	check(err, returnedErr)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err, returnedErr)

	return body, resp.StatusCode
}

// Defines a 200 (Success) response for GetCoreAPIInfo.
type GetCoreAPIInfoResponse struct {
	// The burn chain (i.e., bitcoin) consensus hash at the time that stacks_tip was mined.
	StacksTipConsensusHash string
	// The latest microblock hash if any microblocks were processed. if no microblock has been processed for the current block, a 000.., hex array is returned.
	UnanchoredTip string
	// The block height at which the testnet network will be reset. not applicable for mainnet.
	ExitAtBlockHeight int
	// Is similar to peer_version and will be used to differentiate between different testnets. this value will be different between mainnet and testnet. once launched, this value will not change.
	NetworkID int
	// Same as network_id, but for bitcoin.
	ParentNetworkID int
	// Same as burn_consensus, but evaluated at stable_burn_block_height.
	StablePoxConsensus string
	// The best known block hash for the Stack chain (not including any pending microblocks).
	StacksTip string
	// The latest Stacks chain height. Stacks forks can occur independent of the Bitcoin chain, that height doesn't increase 1-to-1 with the Bitcoin height.
	StacksTipHeight int
	// Latest bitcoin chain height.
	BurnBlockHeight int
	// Identifies the version number for the networking communication, this should not change while a node is running, and will only change if there's an upgrade.
	PeerVersion int
	// Is a hash used to identify the burnchain view for a node. it incorporates bitcoin chain information and PoX information. nodes that disagree on this value will appear to each other as forks. this value will change after every block.
	PoxConsensus string
	// Is a version descriptor.
	ServerVersion string
	// Leftover from stacks 1.0, basically always burn_block_height - 1.
	StableBurnBlockHeight int
}
type GetContractDataMapEntryBody struct{}

// Defines a 200 (Success) response for GetContractDataMapEntry.
type GetContractDataMapEntryResponse struct {
	// Hex-encoded string of clarity value. It is always an optional tuple.
	Data string
	// Hex-encoded string of the MARF proof for the data.
	Proof string
}

// Defines a 400 (Failed loading data map) response for GetContractDataMapEntry.
type GetContractDataMapEntry400Error struct{}

// Defines parameters for GetContractDataMapEntry
type GetContractDataMapEntryParams struct {
	// Returns object without the proof field when set to 0.
	// Optional. Use `-1` to use the default value.
	Proof int
	// Stacks address.
	ContractAddress string
	// Map name.
	MapName string
	// Contract name.
	ContractName string
	// The Stacks chain tip to query from.
	// Optional. Pass an empty string to use the default value.
	Tip string
}

// Defines a 404 (Cannot find block with given height) response for GetBlockByHeight.
type GetBlockByHeight404Error struct{}

// Defines a 200 (Block) response for GetBlockByHeight.
type GetBlockByHeightResponse struct {
	// Hash of the anchor chain block.
	BurnBlockHash string
	// Execution cost read length.
	ExecutionCostReadLength int
	// Height of the block.
	Height int
	// Height of the anchor chain block.
	BurnBlockHeight int
	// Execution cost read count.
	ExecutionCostReadCount int
	// List of microblocks that were accepted in this anchor block. Not every anchored block will have a accepted all (or any) of the previously streamed microblocks. Microblocks that were orphaned are not included in this list.
	MicroblocksAccepted []string
	// Hash of the parent block.
	ParentBlockHash string
	// Anchor chain transaction ID.
	MinerTxid string
	// List of transactions included in the block.
	Txs []string
	// Unix timestamp (in seconds) indicating when this block was mined.
	BurnBlockTime int
	// Set to `true` if block corresponds to the canonical chain tip.
	Canonical bool
	// Hash representing the block.
	Hash string
	// List of microblocks that were streamed/produced by this anchor block's miner. This list only includes microblocks that were accepted in the following anchor block. Microblocks that were orphaned are not included in this list.
	MicroblocksStreamed []string
	// The hash of the last streamed block that precedes this block to which this block is to be appended. Not every anchored block will have a parent microblock stream. An anchored block that does not have a parent microblock stream has the parent microblock hash set to an empty string, and the parent microblock sequence number set to -1.
	ParentMicroblockHash string
	// The hash of the last streamed block that precedes this block to which this block is to be appended. Not every anchored block will have a parent microblock stream. An anchored block that does not have a parent microblock stream has the parent microblock hash set to an empty string, and the parent microblock sequence number set to -1.
	ParentMicroblockSequence int
	// An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) indicating when this block was mined.
	BurnBlockTimeIso string
	// Execution cost runtime.
	ExecutionCostRuntime int
	// Execution cost write count.
	ExecutionCostWriteCount int
	// Execution cost write length.
	ExecutionCostWriteLength int
}

// Defines a 200 (Fungible tokens metadata for contract id) response for GetContractFTMetadata.
type GetContractFTMetadataResponse struct {
	// The number of decimal places in a token.
	Decimals int
	// Describes the asset to which this token represents.
	Description string
	// The original image URI specified by the contract. A URI pointing to a resource with mime type image/* representing the asset to which this token represents. Consider making any images at a width between 320 and 1080 pixels and aspect ratio between 1.91:1 and 4:5 inclusive.
	ImageCanonicalUri string
	// Identifies the asset to which this token represents.
	Name string
	// A shorter representation of a token. This is sometimes referred to as a "ticker". Examples: "STX", "COOL", etc. Typically, a token could be referred to as $SYMBOL when referencing it in writing.
	Symbol string
	// A URI pointing to a resource with mime type image/* representing the asset to which this token represents. The API may provide a URI to a cached resource, dependending on configuration. Otherwise, this can be the same value as the canonical image URI.
	ImageUri string
	// Principle that deployed the contract.
	SenderAddress string
	// An optional string that is a valid URI which resolves to this token's metadata. Can be empty.
	TokenUri string
	// Tx id that deployed the contract.
	TXID string
}

// Defines parameters for GetContractFTMetadata
type GetContractFTMetadataParams struct {
	// Token's contract id.
	ContractId string
}

// Defines a 200 (Success) response for GetFilteredEvents.
type GetFilteredEventsResponse struct {
	Limit   int
	Offset  int
	Results []any
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
	// Hash of transaction.
	// Optional. Pass an empty string to use the default value.
	TXID string
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

// Defines a 200 (Contract interface) response for GetContractInterface.
type GetContractInterfaceResponse struct {
	// List of fungible tokens in the contract.
	FungibleTokens []any
	// List of defined data-maps.
	Maps []any
	// List of non-fungible tokens in the contract.
	NonFungibleTokens []any
	// List of defined variables.
	Variables []any
	// List of defined methods.
	Functions []any
}

// Defines a 200 (Success) response for GetAccountSTXBalance.
type GetAccountSTXBalanceResponse struct{}

// Defines parameters for GetAccountSTXBalance
type GetAccountSTXBalanceParams struct {
	// Returned data representing the state up until that point in time, rather than the current block.
	// Optional. Pass an empty string to use the default value.
	UntilBlock string
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
}

// Defines a 200 (List of microblocks) response for GetMicroblockList.
type GetMicroblockListResponse struct {
	// The number to microblocks to skip (starting at `0`).
	Offset  int
	Results []any
	// The number of microblocks available.
	Total int
	// The number of microblocks to return.
	Limit int
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

// Defines a 404 (Not found) response for SearchByID.
type SearchByID404Error struct{}

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

// Defines a 200 (Success) response for FetchZoneFile.
type FetchZoneFileResponse struct{}

// Defines a 404 (Error) response for FetchZoneFile.
type FetchZoneFile404Error struct {
	Error string
}

// Defines a 400 (Error) response for FetchZoneFile.
type FetchZoneFile400Error struct {
	Error string
}

// Defines parameters for FetchZoneFile
type FetchZoneFileParams struct {
	// Fully-qualified name.
	Name string
}

// Defines a 200 (Success) response for GetAccountBalance.
type GetAccountBalanceResponse struct {
	STX interface{}
	// Token Offering Locked.
	TokenOfferingLocked interface{}
	FungibleTokens      interface{}
	NonFungibleTokens   interface{}
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

// Defines a 404 (Cannot find block with given ID) response for GetBlockByHash.
type GetBlockByHash404Error struct{}

// Defines a 200 (Block) response for GetBlockByHash.
type GetBlockByHashResponse struct {
	// Height of the block.
	Height int
	// List of microblocks that were accepted in this anchor block. Not every anchored block will have a accepted all (or any) of the previously streamed microblocks. Microblocks that were orphaned are not included in this list.
	MicroblocksAccepted []string
	// Height of the anchor chain block.
	BurnBlockHeight int
	// An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) indicating when this block was mined.
	BurnBlockTimeIso string
	// Set to `true` if block corresponds to the canonical chain tip.
	Canonical bool
	// Execution cost write count.
	ExecutionCostWriteCount int
	// Execution cost write length.
	ExecutionCostWriteLength int
	// Hash representing the block.
	Hash string
	// Anchor chain transaction ID.
	MinerTxid string
	// Hash of the parent block.
	ParentBlockHash string
	// The hash of the last streamed block that precedes this block to which this block is to be appended. Not every anchored block will have a parent microblock stream. An anchored block that does not have a parent microblock stream has the parent microblock hash set to an empty string, and the parent microblock sequence number set to -1.
	ParentMicroblockHash string
	// Execution cost read count.
	ExecutionCostReadCount int
	// Execution cost read length.
	ExecutionCostReadLength int
	// List of microblocks that were streamed/produced by this anchor block's miner. This list only includes microblocks that were accepted in the following anchor block. Microblocks that were orphaned are not included in this list.
	MicroblocksStreamed []string
	// Hash of the anchor chain block.
	BurnBlockHash string
	// Unix timestamp (in seconds) indicating when this block was mined.
	BurnBlockTime int
	// Execution cost runtime.
	ExecutionCostRuntime int
	// The hash of the last streamed block that precedes this block to which this block is to be appended. Not every anchored block will have a parent microblock stream. An anchored block that does not have a parent microblock stream has the parent microblock hash set to an empty string, and the parent microblock sequence number set to -1.
	ParentMicroblockSequence int
	// List of transactions included in the block.
	Txs []string
}

// Defines a 200 (Non-Fungible Token mints) response for GetNFTMints.
type GetNFTMintsResponse struct {
	// The number of mint events to return.
	Limit int
	// The number to mint events to skip (starting at `0`).
	Offset  int
	Results []any
	// The number of mint events available.
	Total int
}

// Defines parameters for GetNFTMints
type GetNFTMintsParams struct {
	// Index of first event to fetch.
	// Optional. Use `-1` to use the default value (0).
	Offset int
	// Whether or not to include events from unconfirmed transactions.
	// Optional. Use `false` as default.
	Unanchored bool
	// Max number of events to fetch.
	// Optional. Use `-1` to use the default value (50).
	Limit int
	// Whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times.
	// Optional. Use `false` as default.
	TXMetadata bool
	// Token asset class identifier.
	AssetIdentifier string
}

// Defines a 404 (Error) response for GetHistoricalZoneFile.
type GetHistoricalZoneFile404Error struct {
	Error string
}

// Defines a 200 (Success) response for GetHistoricalZoneFile.
type GetHistoricalZoneFileResponse struct{}

// Defines a 400 (Error) response for GetHistoricalZoneFile.
type GetHistoricalZoneFile400Error struct {
	Error string
}

// Defines parameters for GetHistoricalZoneFile
type GetHistoricalZoneFileParams struct {
	// Fully-qualified name.
	Name string
	// Zone file hash.
	ZoneFileHash string
}

// Defines a 200 (Success) response for GetNamePrice.
type GetNamePriceResponse struct {
	Units  string
	Amount string
}

// Defines parameters for GetNamePrice
type GetNamePriceParams struct {
	// The name to query price information for.
	Name string
}
type FetchFeeRateBody struct{}

// Defines a 200 (Transaction fee rate) response for FetchFeeRate.
type FetchFeeRateResponse struct {
	FeeRate int
}

// Defines a 200 (Success) response for GetSTXSupply.
type GetSTXSupplyResponse struct {
	// The block height at which this information was queried.
	BlockHeight int
	// String quoted decimal number of the total possible number of STX.
	TotalSTX string
	// String quoted decimal number of the percentage of STX that have unlocked.
	UnlockedPercent string
	// String quoted decimal number of the STX that have been mined or unlocked.
	UnlockedSTX string
}

// Defines parameters for GetSTXSupply
type GetSTXSupplyParams struct {
	// The block height at which to query supply details from, if not provided then the latest block height is used.
	Height int
}

// Defines a 200 (Success) response for GetAccountTransactionsWithTransfers.
type GetAccountTransactionsWithTransfersResponse struct {
	Limit   int
	Offset  int
	Results []any
	Total   int
}

// Defines parameters for GetAccountTransactionsWithTransfers
type GetAccountTransactionsWithTransfersParams struct {
	// Index of first account transaction to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Returned data representing the state up until that point in time, rather than the current block.
	// Optional. Pass an empty string to use the default value.
	UntilBlock string
	// Filter for transactions only at this given block height.
	Height int
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
	// Max number of account transactions to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
}

// Defines a 404 (Cannot find contract of given ID) response for GetContractByID.
type GetContractByID404Error struct{}

// Defines a 200 (Contract found) response for GetContractByID.
type GetContractByIDResponse struct{}

// Defines a 404 (Cannot find transaction for given ID) response for GetTransactionByID.
type GetTransactionByID404Error struct{}

// Defines a 200 (Transaction) response for GetTransactionByID.
type GetTransactionByIDResponse struct{}

// Defines a 404 (Cannot find transaction for given ID) response for GetRawTransactionByID.
type GetRawTransactionByID404Error struct{}

// Defines a 200 (Hex encoded serialized transaction) response for GetRawTransactionByID.
type GetRawTransactionByIDResponse struct {
	// A hex encoded serialized transaction.
	RawTX string
}

// Defines a 200 (Success) response for GetPoxInfo.
type GetPoxInfoResponse struct {
	TotalLiquidSupplyUstx      int
	FirstBurnchainBlockHeight  int
	RegistrationWindowLength   int
	RejectionVotesLeftRequired int
	RewardCycleLength          int
	ContractID                 string
	MinAmountUstx              int
	RejectionFraction          int
	RewardCycleID              int
}

// Defines a 200 (List of Transactions) response for GetAddressMempoolTransactions.
type GetAddressMempoolTransactionsResponse struct {
	Limit   int
	Offset  int
	Results []any
	Total   int
}

// Defines parameters for GetAddressMempoolTransactions
type GetAddressMempoolTransactionsParams struct {
	// Max number of transactions to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Index of first transaction to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Transactions for the address.
	Address string
}

// Defines a 200 (Success) response for GetAccountNonces.
type GetAccountNoncesResponse struct {
	// Nonces that appear to be missing and likely causing a mempool transaction to be stuck.
	DetectedMissingNonces []int
	// The latest nonce found within transactions sent by this address, including unanchored microblock transactions. Will be null if there are no current transactions for this address.
	LastExecutedTXNonce int
	// The latest nonce found within mempool transactions sent by this address. Will be null if there are no current mempool transactions for this address.
	LastMempoolTXNonce int
	// The likely nonce required for creating the next transaction, based on the last nonces seen by the API. This can be incorrect if the API's mempool or transactions aren't fully synchronized, even by a small amount, or if a previous transaction is still propagating through the Stacks blockchain network when this endpoint is called.
	PossibleNextNonce int
}

// Defines parameters for GetAccountNonces
type GetAccountNoncesParams struct {
	// Optionally get the nonce at a given block height.
	BlockHeight int
	// Optionally get the nonce at a given block hash.
	// Optional. Pass an empty string to use the default value.
	BlockHash string
	// Stacks address (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0`).
	Principal string
}

// Defines a 200 (List of fungible tokens metadata) response for GetFTMetadataList.
type GetFTMetadataListResponse struct {
	// The number of tokens metadata to return.
	Limit int
	// The number to tokens metadata to skip (starting at `0`).
	Offset  int
	Results []any
	// The number of tokens metadata available.
	Total int
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

// Defines a 200 (List of dropped mempool transactions) response for GetDroppedMempoolTransactionList.
type GetDroppedMempoolTransactionListResponse struct {
	Limit   int
	Offset  int
	Results []any
	Total   int
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

// Defines a 200 (Success) response for GetAccountTransactions.
type GetAccountTransactionsResponse struct {
	Offset  int
	Results []any
	Total   int
	Limit   int
}

// Defines parameters for GetAccountTransactions
type GetAccountTransactionsParams struct {
	// Index of first account transaction to fetch.
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
	// Filter for transactions only at this given block height.
	Height int
	// Max number of account transactions to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
}

// Defines a 200 (success) response for GetSTXSupplyTotalSupplyPlain.
type GetSTXSupplyTotalSupplyPlainResponse struct{}

// Defines a 400 (Error) response for GetAllNames.
type GetAllNames400Error struct {
	Error string
}

// Defines a 200 (Success) response for GetAllNames.
type GetAllNamesResponse struct{}

// Defines parameters for GetAllNames
type GetAllNamesParams struct {
	// Names are returned in pages of size 100, so specify the page number.
	Page int
}

// Defines a 200 (Success) response for GetAllNamespaces.
type GetAllNamespacesResponse struct {
	Namespaces []string
}
type CallReadOnlyFunctionBody struct{}

// Defines a 200 (Success) response for CallReadOnlyFunction.
type CallReadOnlyFunctionResponse struct {
	Cause  string
	Okay   bool
	Result string
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

// Defines a 404 (Cannot find block with given height) response for GetBlockByBurnBlockHeight.
type GetBlockByBurnBlockHeight404Error struct{}

// Defines a 200 (Block) response for GetBlockByBurnBlockHeight.
type GetBlockByBurnBlockHeightResponse struct {
	// List of microblocks that were streamed/produced by this anchor block's miner. This list only includes microblocks that were accepted in the following anchor block. Microblocks that were orphaned are not included in this list.
	MicroblocksStreamed []string
	// Anchor chain transaction ID.
	MinerTxid string
	// Hash of the anchor chain block.
	BurnBlockHash string
	// Set to `true` if block corresponds to the canonical chain tip.
	Canonical bool
	// Execution cost write length.
	ExecutionCostWriteLength int
	// Height of the anchor chain block.
	BurnBlockHeight int
	// Execution cost read length.
	ExecutionCostReadLength int
	// Execution cost runtime.
	ExecutionCostRuntime int
	// Height of the block.
	Height int
	// Hash of the parent block.
	ParentBlockHash string
	// The hash of the last streamed block that precedes this block to which this block is to be appended. Not every anchored block will have a parent microblock stream. An anchored block that does not have a parent microblock stream has the parent microblock hash set to an empty string, and the parent microblock sequence number set to -1.
	ParentMicroblockSequence int
	// Unix timestamp (in seconds) indicating when this block was mined.
	BurnBlockTime int
	// An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) indicating when this block was mined.
	BurnBlockTimeIso string
	// Execution cost read count.
	ExecutionCostReadCount int
	// The hash of the last streamed block that precedes this block to which this block is to be appended. Not every anchored block will have a parent microblock stream. An anchored block that does not have a parent microblock stream has the parent microblock hash set to an empty string, and the parent microblock sequence number set to -1.
	ParentMicroblockHash string
	// List of transactions included in the block.
	Txs []string
	// Execution cost write count.
	ExecutionCostWriteCount int
	// Hash representing the block.
	Hash string
	// List of microblocks that were accepted in this anchor block. Not every anchored block will have a accepted all (or any) of the previously streamed microblocks. Microblocks that were orphaned are not included in this list.
	MicroblocksAccepted []string
}
type RunFaucetSTXBody struct{}

// Defines a 500 (Faucet call failed) response for RunFaucetSTX.
type RunFaucetSTX500Error struct{}

// Defines a 200 (Success) response for RunFaucetSTX.
type RunFaucetSTXResponse struct {
	// The transaction ID for the faucet call.
	TxId string
	// Raw transaction in hex string representation.
	TxRaw string
	// Indicates if the faucet call was successful.
	Success bool
}

// Defines a 200 (Success) response for GetNetworkBlockTimeByNetwork.
type GetNetworkBlockTimeByNetworkResponse struct{}

// Defines parameters for GetNetworkBlockTimeByNetwork
type GetNetworkBlockTimeByNetworkParams struct {
	// Which network to retrieve the target block time of.
	Network string
}

// Defines a 200 (Success) response for GetNetworkBlockTimes.
type GetNetworkBlockTimesResponse struct {
	Mainnet interface{}
	Testnet interface{}
}

// Defines a 200 (Success) response for GetStatus.
type GetStatusResponse struct {
	// Current chain tip information.
	ChainTip interface{}
	// The server version that is currently running.
	ServerVersion string
	// The current server status.
	Status string
}

// Defines a 200 (Success) response for GetTotalSTXSupplyLegacyFormat.
type GetTotalSTXSupplyLegacyFormatResponse struct {
	// String quoted decimal number of the percentage of STX that have unlocked.
	UnlockedPercent string
	// String quoted decimal number of the STX that have been mined or unlocked.
	UnlockedSupply string
	// Same as `unlockedSupply` but formatted with comma thousands separators.
	UnlockedSupplyFormatted string
	// The block height at which this information was queried.
	BlockHeight string
	// String quoted decimal number of the total possible number of STX.
	TotalStacks string
	// Same as `totalStacks` but formatted with comma thousands separators.
	TotalStacksFormatted string
}

// Defines parameters for GetTotalSTXSupplyLegacyFormat
type GetTotalSTXSupplyLegacyFormatParams struct {
	// The block height at which to query supply details from, if not provided then the latest block height is used.
	Height int
}

// Defines a 404 (Error) response for GetNameInfo.
type GetNameInfo404Error struct {
	Error string
}

// Defines a 400 (Error) response for GetNameInfo.
type GetNameInfo400Error struct {
	Error string
}

// Defines a 200 (Success) response for GetNameInfo.
type GetNameInfoResponse struct {
	Resolver     string
	Status       string
	Zonefile     string
	GracePeriod  int
	LastTxid     string
	ExpireBlock  int
	ZonefileHash string
	Address      string
	Blockchain   string
}

// Defines parameters for GetNameInfo
type GetNameInfoParams struct {
	// Fully-qualified name.
	Name string
}

// Defines a 200 (List of burnchain reward recipients and amounts) response for GetBurnchainRewardListByAddress.
type GetBurnchainRewardListByAddressResponse struct {
	// The number of burnchain rewards to return.
	Limit int
	// The number to burnchain rewards to skip (starting at `0`).
	Offset  int
	Results []any
}

// Defines parameters for GetBurnchainRewardListByAddress
type GetBurnchainRewardListByAddressParams struct {
	// Max number of rewards to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Reward recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format.
	Address string
	// Index of first rewards to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
}
type RunFaucetBTCBody struct{}

// Defines a 500 (Faucet call failed) response for RunFaucetBTC.
type RunFaucetBTC500Error struct{}

// Defines a 200 (Success) response for RunFaucetBTC.
type RunFaucetBTCResponse struct {
	// Indicates if the faucet call was successful.
	Success bool
	// The transaction ID for the faucet call.
	TxId string
	// Raw transaction in hex string representation.
	TxRaw string
}

// Defines a 400 (Error) response for GetNamespaceNames.
type GetNamespaceNames400Error struct {
	Error string
}

// Defines a 404 (Error) response for GetNamespaceNames.
type GetNamespaceNames404Error struct {
	Error string
}

// Defines a 200 (Success) response for GetNamespaceNames.
type GetNamespaceNamesResponse struct{}

// Defines parameters for GetNamespaceNames
type GetNamespaceNamesParams struct {
	// Names are returned in pages of size 100, so specify the page number.
	Page int
	// The namespace to fetch names from.
	TLD string
}

// Defines a 200 (Success) response for GetAccountInfo.
type GetAccountInfoResponse struct {
	Balance      string
	BalanceProof string
	Locked       string
	Nonce        int
	NonceProof   string
	UnlockHeight int
}

// Defines parameters for GetAccountInfo
type GetAccountInfoParams struct {
	// Returns object without the proof field if set to 0.
	// Optional. Use `-1` to use the default value.
	Proof int
	// The Stacks chain tip to query from.
	// Optional. Pass an empty string to use the default value.
	Tip string
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
}

// Defines a 200 (List of burnchain reward recipients and amounts) response for GetBurnchainRewardSlotHoldersByAddress.
type GetBurnchainRewardSlotHoldersByAddressResponse struct {
	// The number of items to return.
	Limit int
	// The number of items to skip (starting at `0`).
	Offset  int
	Results []any
	// Total number of available items.
	Total int
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

// Defines a 200 (List of burnchain reward recipients and amounts) response for GetBurnchainRewardList.
type GetBurnchainRewardListResponse struct {
	// The number of burnchain rewards to return.
	Limit int
	// The number to burnchain rewards to skip (starting at `0`).
	Offset  int
	Results []any
}

// Defines parameters for GetBurnchainRewardList
type GetBurnchainRewardListParams struct {
	// Max number of rewards to fetch.
	// Optional. Use `-1` to use the default value (96).
	// Max value is 250.
	Limit int
	// Index of first rewards to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
}

// Defines a 200 (Non-Fungible Token event history) response for GetNFTHistory.
type GetNFTHistoryResponse struct {
	// The number of events to return.
	Limit int
	// The number to events to skip (starting at `0`).
	Offset  int
	Results []any
	// The number of events available.
	Total int
}

// Defines parameters for GetNFTHistory
type GetNFTHistoryParams struct {
	// Max number of events to fetch.
	// Optional. Use `-1` to use the default value (50).
	Limit int
	// Whether or not to include events from unconfirmed transactions.
	// Optional. Use `false` as default.
	Unanchored bool
	// Token asset class identifier.
	AssetIdentifier string
	// Whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times.
	// Optional. Use `false` as default.
	TXMetadata bool
	// Index of first event to fetch.
	// Optional. Use `-1` to use the default value (0).
	Offset int
	// Hex representation of the token's unique value.
	Value string
}

// Defines a 200 (List of Transactions) response for GetTransactionsByBlockHeight.
type GetTransactionsByBlockHeightResponse struct {
	// The number of transactions available.
	Total int
	// The number of transactions to return.
	Limit int
	// The number to transactions to skip (starting at `0`).
	Offset  int
	Results []any
}

// Defines parameters for GetTransactionsByBlockHeight
type GetTransactionsByBlockHeightParams struct {
	// Max number of transactions to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Index of first transaction to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Height of block.
	Height int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
}

// Defines a 200 (Success) response for GetAccountInbound.
type GetAccountInboundResponse struct {
	Limit   int
	Offset  int
	Results []any
	Total   int
}

// Defines parameters for GetAccountInbound
type GetAccountInboundParams struct {
	// Number of items to skip.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Returned data representing the state up until that point in time, rather than the current block.
	// Optional. Pass an empty string to use the default value.
	UntilBlock string
	// Filter for transfers only at this given block height.
	Height int
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
	// Number of items to return.
	// Optional. Use `-1` to use the default value.
	Limit int
}

// Defines a 404 (Cannot find block with given height) response for GetBlockByBurnBlockHash.
type GetBlockByBurnBlockHash404Error struct{}

// Defines a 200 (Block) response for GetBlockByBurnBlockHash.
type GetBlockByBurnBlockHashResponse struct {
	// Hash of the anchor chain block.
	BurnBlockHash string
	// Execution cost write count.
	ExecutionCostWriteCount int
	// Anchor chain transaction ID.
	MinerTxid string
	// List of transactions included in the block.
	Txs []string
	// The hash of the last streamed block that precedes this block to which this block is to be appended. Not every anchored block will have a parent microblock stream. An anchored block that does not have a parent microblock stream has the parent microblock hash set to an empty string, and the parent microblock sequence number set to -1.
	ParentMicroblockHash string
	// An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) indicating when this block was mined.
	BurnBlockTimeIso string
	// Execution cost read length.
	ExecutionCostReadLength int
	// Hash representing the block.
	Hash string
	// Height of the block.
	Height int
	// List of microblocks that were accepted in this anchor block. Not every anchored block will have a accepted all (or any) of the previously streamed microblocks. Microblocks that were orphaned are not included in this list.
	MicroblocksAccepted []string
	// Hash of the parent block.
	ParentBlockHash string
	// Set to `true` if block corresponds to the canonical chain tip.
	Canonical bool
	// Execution cost read count.
	ExecutionCostReadCount int
	// Execution cost runtime.
	ExecutionCostRuntime int
	// Execution cost write length.
	ExecutionCostWriteLength int
	// The hash of the last streamed block that precedes this block to which this block is to be appended. Not every anchored block will have a parent microblock stream. An anchored block that does not have a parent microblock stream has the parent microblock hash set to an empty string, and the parent microblock sequence number set to -1.
	ParentMicroblockSequence int
	// Height of the anchor chain block.
	BurnBlockHeight int
	// Unix timestamp (in seconds) indicating when this block was mined.
	BurnBlockTime int
	// List of microblocks that were streamed/produced by this anchor block's miner. This list only includes microblocks that were accepted in the following anchor block. Microblocks that were orphaned are not included in this list.
	MicroblocksStreamed []string
}

// Defines a 200 (Success) response for GetContractSource.
type GetContractSourceResponse struct {
	Proof         string
	PublishHeight int
	Source        string
}

// Defines a 200 (List of contracts implement given trait) response for GetContractsByTrait.
type GetContractsByTraitResponse struct {
	// The number of contracts to return.
	Limit int
	// The number to contracts to skip (starting at `0`).
	Offset  int
	Results []any
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

// Defines a 200 (List of transactions) response for GetTransactionList.
type GetTransactionListResponse struct {
	// The number of transactions to return.
	Limit int
	// The number to transactions to skip (starting at `0`).
	Offset  int
	Results []any
	// The number of transactions available.
	Total int
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
	// Max number of transactions to fetch.
	// Optional. Use `-1` to use the default value (96).
	// Max value is 200.
	Limit int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Filter by transaction type.
	// [Coinbase TokenTransfer SmartContract ContractCall PoisonMicroblock]
	Type GetTransactionListType
}
type PostFeeTransactionBody struct{}

// Defines a 200 (Estimated fees for the transaction) response for PostFeeTransaction.
type PostFeeTransactionResponse struct {
	EstimatedCost          interface{}
	EstimatedCostScalar    int
	Estimations            []any
	CostScalarChangeByByte int
}

// Defines a 200 (Success) response for GetAccountAssets.
type GetAccountAssetsResponse struct {
	Limit   int
	Offset  int
	Results []any
	Total   int
}

// Defines parameters for GetAccountAssets
type GetAccountAssetsParams struct {
	// Index of first account assets to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Max number of account assets to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
	// Returned data representing the state at that point in time, rather than the current block.
	// Optional. Pass an empty string to use the default value.
	UntilBlock string
}

// Defines a 200 (List of burnchain reward recipients and amounts) response for GetBurnchainRewardsTotalByAddress.
type GetBurnchainRewardsTotalByAddressResponse struct {
	// The total amount of burnchain tokens rewarded to the recipient, in the smallest unit (e.g. satoshis for Bitcoin).
	RewardAmount string
	// The recipient address that received the burnchain rewards, in the format native to the burnchain (e.g. B58 encoded for Bitcoin).
	RewardRecipient string
}

// Defines parameters for GetBurnchainRewardsTotalByAddress
type GetBurnchainRewardsTotalByAddressParams struct {
	// Reward recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format.
	Address string
}

// Defines a 404 (Cannot find microblock with given hash) response for GetMicroblockByHash.
type GetMicroblockByHash404Error struct{}

// Defines a 200 (Microblock) response for GetMicroblockByHash.
type GetMicroblockByHashResponse struct {
	// The SHA512/256 hash of the previous signed microblock in this stream.
	MicroblockParentHash string
	// The hash of the anchor block that preceded this microblock.
	ParentBlockHash string
	// The hash of the Bitcoin block that preceded this microblock.
	ParentBurnBlockHash string
	// The height of the Bitcoin block that preceded this microblock.
	ParentBurnBlockHeight int
	// The ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) formatted block time of the bitcoin block that preceded this microblock.
	ParentBurnBlockTimeIso string
	// The hash of the anchor block that confirmed this microblock. This wil be empty for unanchored microblocks.
	BlockHash string
	// Set to `true` if the microblock corresponds to the canonical chain tip.
	Canonical bool
	// The SHA512/256 hash of this microblock.
	MicroblockHash string
	// Set to `true` if the microblock was not orphaned in a following anchor block. Defaults to `true` if the following anchor block has not yet been created.
	MicroblockCanonical bool
	// List of transactions included in the microblock.
	Txs []string
	// The anchor block height that confirmed this microblock.
	BlockHeight int
	// A hint to describe how to order a set of microblocks. Starts at 0.
	MicroblockSequence int
	// The height of the anchor block that preceded this microblock.
	ParentBlockHeight int
	// The block timestamp of the Bitcoin block that preceded this microblock.
	ParentBurnBlockTime int
}

// Defines a 200 (List of mempool transactions) response for GetMempoolTransactionList.
type GetMempoolTransactionListResponse struct {
	Results []any
	Total   int
	Limit   int
	Offset  int
}

// Defines parameters for GetMempoolTransactionList
type GetMempoolTransactionListParams struct {
	// Filter to only return transactions with this address as the sender or recipient (recipient only applicable for STX transfer tx types).
	// Optional. Pass an empty string to use the default value.
	Address string
	// Filter to only return transactions with this recipient address (only applicable for STX transfer tx types).
	// Optional. Pass an empty string to use the default value.
	RecipientAddress string
	// Filter to only return transactions with this sender address.
	// Optional. Pass an empty string to use the default value.
	SenderAddress string
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Index of first mempool transaction to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Max number of mempool transactions to fetch.
	// Optional. Use `-1` to use the default value (96).
	// Max value is 200.
	Limit int
}

// Defines a 200 (Returns list of transactions with their details for corresponding requested tx_ids.) response for GetTXListDetails.
type GetTXListDetailsResponse struct{}

// Defines a 404 (Could not find any transaction by ID) response for GetTXListDetails.
type GetTXListDetails404Error struct{}

// Defines a 200 (List of non fungible tokens metadata) response for GetNFTMetadataList.
type GetNFTMetadataListResponse struct {
	// The number of tokens metadata to return.
	Limit int
	// The number to tokens metadata to skip (starting at `0`).
	Offset  int
	Results []any
	// The number of tokens metadata available.
	Total int
}

// Defines parameters for GetNFTMetadataList
type GetNFTMetadataListParams struct {
	// Max number of tokens to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Index of first tokens to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
}

// Defines a 200 (Non fungible tokens metadata for contract id) response for GetContractNFTMetadata.
type GetContractNFTMetadataResponse struct {
	// Describes the asset to which this token represents.
	Description string
	// The original image URI specified by the contract. A URI pointing to a resource with mime type image/* representing the asset to which this token represents. Consider making any images at a width between 320 and 1080 pixels and aspect ratio between 1.91:1 and 4:5 inclusive.
	ImageCanonicalUri string
	// A URI pointing to a resource with mime type image/* representing the asset to which this token represents. The API may provide a URI to a cached resource, dependending on configuration. Otherwise, this can be the same value as the canonical image URI.
	ImageUri string
	// Identifies the asset to which this token represents.
	Name string
	// Principle that deployed the contract.
	SenderAddress string
	// An optional string that is a valid URI which resolves to this token's metadata. Can be empty.
	TokenUri string
	// Tx id that deployed the contract.
	TXID string
}

// Defines parameters for GetContractNFTMetadata
type GetContractNFTMetadataParams struct {
	// Token's contract id.
	ContractId string
}
type PostCoreNodeTransactionsBody struct{}

// Defines a 200 (Transaction id of successful post of a raw tx to the node's mempool) response for PostCoreNodeTransactions.
type PostCoreNodeTransactionsResponse struct{}

// Defines a 400 (Rejections result in a 400 error) response for PostCoreNodeTransactions.
type PostCoreNodeTransactions400Error struct {
	// The error.
	Error string
	// The reason for the error.
	Reason string
	// More details about the reason.
	ReasonData interface{}
	// The relevant transaction id.
	Txid string
}

// Defines a 200 (Success) response for GetAccountNFT.
type GetAccountNFTResponse struct {
	Limit     int
	NFTEvents []any
	Offset    int
	Total     int
}

// Defines parameters for GetAccountNFT
type GetAccountNFTParams struct {
	// Number of items to skip.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Returned data representing the state up until that point in time, rather than the current block.
	// Optional. Pass an empty string to use the default value.
	UntilBlock string
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
	// Number of items to return.
	// Optional. Use `-1` to use the default value.
	Limit int
}

// Defines a 404 (Not found) response for GetSingleTransactionWithTransfers.
type GetSingleTransactionWithTransfers404Error struct{}

// Defines a 200 (Success) response for GetSingleTransactionWithTransfers.
type GetSingleTransactionWithTransfersResponse struct {
	FTTransfers  []any
	NFTTransfers []any
	// Total received by the given address in micro-STX as an integer string.
	STXReceived string
	// Total sent from the given address, including the tx fee, in micro-STX as an integer string.
	STXSent      string
	STXTransfers []any
	// Describes all transaction types on Stacks 2.0 blockchain.
	TX interface{}
}

// Defines parameters for GetSingleTransactionWithTransfers
type GetSingleTransactionWithTransfersParams struct {
	// Transaction id.
	TXID string
	// Stacks address or a contract identifier.
	Principal string
}

// Defines a 200 (Transactions) response for GetUnanchoredTxs.
type GetUnanchoredTxsResponse struct {
	Results []any
	// The number of unanchored transactions available.
	Total int
}

// Defines a 200 (List of Transactions) response for GetTransactionsByBlockHash.
type GetTransactionsByBlockHashResponse struct {
	// The number of transactions to return.
	Limit int
	// The number to transactions to skip (starting at `0`).
	Offset  int
	Results []any
	// The number of transactions available.
	Total int
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

// Defines a 200 (List of burnchain reward recipients and amounts) response for GetBurnchainRewardSlotHolders.
type GetBurnchainRewardSlotHoldersResponse struct {
	Results []any
	// Total number of available items.
	Total int
	// The number of items to return.
	Limit int
	// The number of items to skip (starting at `0`).
	Offset int
}

// Defines parameters for GetBurnchainRewardSlotHolders
type GetBurnchainRewardSlotHoldersParams struct {
	// Max number of items to fetch.
	// Optional. Use `-1` to use the default value (96).
	// Max value is 250.
	Limit int
	// Index of the first items to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
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
	// Contract identifier formatted as `<contract_address>.<contract_name>`.
	ContractID string
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
}

// Defines a 404 (Error) response for GetNamesOwnedByAddress.
type GetNamesOwnedByAddress404Error struct {
	Error string
}

// Defines a 200 (Success) response for GetNamesOwnedByAddress.
type GetNamesOwnedByAddressResponse struct {
	Names []string
}

// Defines parameters for GetNamesOwnedByAddress
type GetNamesOwnedByAddressParams struct {
	// The address to lookup.
	Address string
	// The layer-1 blockchain for the address.
	Blockchain string
}

// Defines a 200 (List of blocks) response for GetBlockList.
type GetBlockListResponse struct {
	// The number of blocks to return.
	Limit int
	// The number to blocks to skip (starting at `0`).
	Offset  int
	Results []any
	// The number of blocks available.
	Total int
}

// Defines parameters for GetBlockList
type GetBlockListParams struct {
	// Index of first block to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Max number of blocks to fetch.
	// Optional. Use `-1` to use the default value (20).
	// Max value is 200.
	Limit int
}

// Defines a 200 (success) response for GetSTXSupplyCirculatingPlain.
type GetSTXSupplyCirculatingPlainResponse struct{}

// Defines a 200 (Success) response for GetNamespacePrice.
type GetNamespacePriceResponse struct {
	Amount string
	Units  string
}

// Defines parameters for GetNamespacePrice
type GetNamespacePriceParams struct {
	// The namespace to fetch price for.
	TLD string
}

// Defines a 200 (List of Non-Fungible Token holdings) response for GetNFTHoldings.
type GetNFTHoldingsResponse struct {
	// The number of Non-Fungible Token holdings to return.
	Limit int
	// The number to Non-Fungible Token holdings to skip (starting at `0`).
	Offset  int
	Results []any
	// The number of Non-Fungible Token holdings available.
	Total int
}

// Defines parameters for GetNFTHoldings
type GetNFTHoldingsParams struct {
	// Token owner's STX address or Smart Contract ID.
	Principal string
	// Identifiers of the token asset classes to filter for.
	// Max number of tokens to fetch.
	// Optional. Use `-1` to use the default value (50).
	Limit int
	// Whether or not to include tokens from unconfirmed transactions.
	// Optional. Use `false` as default.
	Unanchored bool
	// Whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times.
	// Optional. Use `false` as default.
	TXMetadata bool
	// Index of first tokens to fetch.
	// Optional. Use `-1` to use the default value (0).
	Offset int
}

// Defines a 200 (Success) response for FetchSubdomainsListForName.
type FetchSubdomainsListForNameResponse struct{}

// Defines parameters for FetchSubdomainsListForName
type FetchSubdomainsListForNameParams struct {
	// Fully-qualified name.
	Name string
}
type CallReadOnlyFunctionArguments struct {
	// The simulated tx-sender
	Sender string
	// An array of hex serialized Clarity values
	Arguments []string
}
