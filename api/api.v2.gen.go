package api

// Defines parameters for GetTransactionsByBlockHeight
type GetTransactionsByBlockHeightParams struct {
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Index of first transaction to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Max number of transactions to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Height of block.
	Height int
}

// Defines parameters for GetTXListDetails
type GetTXListDetailsParams struct{}

// Defines parameters for PostFeeTransaction
type PostFeeTransactionParams struct {
	//
	EstimatedLen int
	//
	TransactionPayload string
}

// Defines parameters for GetCoreAPIInfo
type GetCoreAPIInfoParams struct{}

// Defines parameters for GetAccountTransactions
type GetAccountTransactionsParams struct {
	// Returned data representing the state up until that point in time, rather than the current block.
	// Optional. Pass an empty string to use the default value.
	UntilBlock string
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
	// Max number of account transactions to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Filter for transactions only at this given block height.
	Height int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Index of first account transaction to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
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

// Defines parameters for GetNetworkBlockTimes
type GetNetworkBlockTimesParams struct{}

// Defines parameters for GetNFTMints
type GetNFTMintsParams struct {
	// Whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times.
	// Optional. Use `false` as default.
	TXMetadata bool
	// Token asset class identifier.
	AssetIdentifier string
	// Index of first event to fetch.
	// Optional. Use `-1` to use the default value (0).
	Offset int
	// Max number of events to fetch.
	// Optional. Use `-1` to use the default value (50).
	Limit int
	// Whether or not to include events from unconfirmed transactions.
	// Optional. Use `false` as default.
	Unanchored bool
}

// Defines parameters for FetchSubdomainsListForName
type FetchSubdomainsListForNameParams struct {
	// Fully-qualified name.
	Name string
}

// Defines parameters for GetContractDataMapEntry
type GetContractDataMapEntryParams struct {
	// The Stacks chain tip to query from.
	// Optional. Pass an empty string to use the default value.
	Tip string
	// Map name.
	MapName string
	// Contract name.
	ContractName string
	// Stacks address.
	ContractAddress string
	// Returns object without the proof field when set to 0.
	// Optional. Use `-1` to use the default value.
	Proof int
}

// Defines parameters for GetAccountNFT
type GetAccountNFTParams struct {
	// Returned data representing the state up until that point in time, rather than the current block.
	// Optional. Pass an empty string to use the default value.
	UntilBlock string
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Number of items to return.
	// Optional. Use `-1` to use the default value.
	Limit int
}

// Defines parameters for GetSingleTransactionWithTransfers
type GetSingleTransactionWithTransfersParams struct {
	// Transaction id.
	TXID string
	// Stacks address or a contract identifier.
	Principal string
}

// Defines parameters for GetBlockByBurnBlockHash
type GetBlockByBurnBlockHashParams struct{}

// Defines parameters for RunFaucetSTX
type RunFaucetSTXParams struct {
	// Use required number of tokens for stacking.
	Stacking bool
	// STX testnet address.
	Address string
}

// Defines parameters for GetMicroblockList
type GetMicroblockListParams struct {
	// Index of the first microblock to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Max number of microblocks to fetch.
	// Optional. Use `-1` to use the default value (20).
	// Max value is 200.
	Limit int
}

// Defines parameters for GetContractNFTMetadata
type GetContractNFTMetadataParams struct {
	// Token's contract id.
	ContractId string
}

// Defines parameters for CallReadOnlyFunction
type CallReadOnlyFunctionParams struct {
	// The simulated tx-sender.
	Sender string
	// The Stacks chain tip to query from.
	// Optional. Pass an empty string to use the default value.
	Tip string
	// Function name.
	FunctionName string
	// Stacks address.
	ContractAddress string
	// Contract name.
	ContractName string
	// An array of hex serialized Clarity values.
	Arguments CallReadOnlyFunctionArguments
}

// Defines parameters for GetAccountBalance
type GetAccountBalanceParams struct {
	// Returned data representing the state up until that point in time, rather than the current block.
	// Optional. Pass an empty string to use the default value.
	UntilBlock string
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
}

// Defines parameters for GetContractsByTrait
type GetContractsByTraitParams struct {
	// Index of first contract event to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// JSON abi of the trait.
	TraitABI string
	// Max number of contracts fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
}

// Defines parameters for FetchFeeRate
type FetchFeeRateParams struct {
	// A serialized transaction.
	Transaction string
}

// Defines parameters for GetSTXSupplyCirculatingPlain
type GetSTXSupplyCirculatingPlainParams struct{}

// Defines parameters for GetNamePrice
type GetNamePriceParams struct {
	// The name to query price information for.
	Name string
}

// Defines parameters for GetNamespacePrice
type GetNamespacePriceParams struct {
	// The namespace to fetch price for.
	TLD string
}

// Defines parameters for GetAccountTransactionsWithTransfers
type GetAccountTransactionsWithTransfersParams struct {
	// Returned data representing the state up until that point in time, rather than the current block.
	// Optional. Pass an empty string to use the default value.
	UntilBlock string
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
	// Index of first account transaction to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Max number of account transactions to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Filter for transactions only at this given block height.
	Height int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
}

// Defines parameters for GetSTXSupplyTotalSupplyPlain
type GetSTXSupplyTotalSupplyPlainParams struct{}

// Defines parameters for GetAddressMempoolTransactions
type GetAddressMempoolTransactionsParams struct {
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Max number of transactions to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Index of first transaction to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Transactions for the address.
	Address string
}

// Defines parameters for GetBlockByBurnBlockHeight
type GetBlockByBurnBlockHeightParams struct{}

// Defines parameters for GetBurnchainRewardListByAddress
type GetBurnchainRewardListByAddressParams struct {
	// Index of first rewards to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Reward recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format.
	Address string
	// Max number of rewards to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
}

// Defines parameters for GetTransactionsByBlockHash
type GetTransactionsByBlockHashParams struct {
	// Index of first transaction to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Hash of block.
	BlockHash string
	// Max number of transactions to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
}

// Defines parameters for GetContractSource
type GetContractSourceParams struct{}

// Defines parameters for PostCoreNodeTransactions
type PostCoreNodeTransactionsParams struct{}

// Defines parameters for RunFaucetBTC
type RunFaucetBTCParams struct {
	// BTC testnet address.
	Address string
}

// Defines parameters for GetMempoolTransactionList
type GetMempoolTransactionListParams struct {
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Filter to only return transactions with this sender address.
	// Optional. Pass an empty string to use the default value.
	SenderAddress string
	// Filter to only return transactions with this address as the sender or recipient (recipient only applicable for STX transfer tx types).
	// Optional. Pass an empty string to use the default value.
	Address string
	// Filter to only return transactions with this recipient address (only applicable for STX transfer tx types).
	// Optional. Pass an empty string to use the default value.
	RecipientAddress string
	// Max number of mempool transactions to fetch.
	// Optional. Use `-1` to use the default value (96).
	// Max value is 200.
	Limit int
	// Index of first mempool transaction to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
}

// Defines parameters for GetPoxInfo
type GetPoxInfoParams struct{}

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

// Defines parameters for GetUnanchoredTxs
type GetUnanchoredTxsParams struct{}

// Defines parameters for GetMicroblockByHash
type GetMicroblockByHashParams struct{}

// Defines parameters for SearchByID
type SearchByIDParams struct {
	// This includes the detailed data for purticular hash in the response.
	// Optional. Use `<nil>` as default.
	IncludeMetadata bool
	// The hex hash string for a block or transaction, account address, or contract address.
	ID string
}

// Defines parameters for GetAccountInfo
type GetAccountInfoParams struct {
	// The Stacks chain tip to query from.
	// Optional. Pass an empty string to use the default value.
	Tip string
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
	// Returns object without the proof field if set to 0.
	// Optional. Use `-1` to use the default value.
	Proof int
}

// Defines parameters for GetAccountSTXBalance
type GetAccountSTXBalanceParams struct {
	// Returned data representing the state up until that point in time, rather than the current block.
	// Optional. Pass an empty string to use the default value.
	UntilBlock string
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
}

// Defines parameters for GetBlockByHeight
type GetBlockByHeightParams struct{}

// Defines parameters for GetTotalSTXSupplyLegacyFormat
type GetTotalSTXSupplyLegacyFormatParams struct {
	// The block height at which to query supply details from, if not provided then the latest block height is used.
	Height int
}

// Defines parameters for GetAllNamespaces
type GetAllNamespacesParams struct{}

// Defines parameters for GetAccountInbound
type GetAccountInboundParams struct {
	// Returned data representing the state up until that point in time, rather than the current block.
	// Optional. Pass an empty string to use the default value.
	UntilBlock string
	// Number of items to skip.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Filter for transfers only at this given block height.
	Height int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
	// Number of items to return.
	// Optional. Use `-1` to use the default value.
	Limit int
}

// Defines parameters for GetStatus
type GetStatusParams struct{}

// Defines parameters for GetContractFTMetadata
type GetContractFTMetadataParams struct {
	// Token's contract id.
	ContractId string
}

// Defines parameters for FetchZoneFile
type FetchZoneFileParams struct {
	// Fully-qualified name.
	Name string
}

// Defines parameters for GetHistoricalZoneFile
type GetHistoricalZoneFileParams struct {
	// Zone file hash.
	ZoneFileHash string
	// Fully-qualified name.
	Name string
}

// Defines parameters for GetBurnchainRewardsTotalByAddress
type GetBurnchainRewardsTotalByAddressParams struct {
	// Reward recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format.
	Address string
}

// Defines parameters for GetNFTHoldings
type GetNFTHoldingsParams struct {
	// Whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times.
	// Optional. Use `false` as default.
	TXMetadata bool
	// Token owner's STX address or Smart Contract ID.
	Principal string
	// Index of first tokens to fetch.
	// Optional. Use `-1` to use the default value (0).
	Offset int
	// Max number of tokens to fetch.
	// Optional. Use `-1` to use the default value (50).
	Limit int
	// Identifiers of the token asset classes to filter for.
	// Whether or not to include tokens from unconfirmed transactions.
	// Optional. Use `false` as default.
	Unanchored bool
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

// Defines parameters for GetNFTHistory
type GetNFTHistoryParams struct {
	// Whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times.
	// Optional. Use `false` as default.
	TXMetadata bool
	// Max number of events to fetch.
	// Optional. Use `-1` to use the default value (50).
	Limit int
	// Index of first event to fetch.
	// Optional. Use `-1` to use the default value (0).
	Offset int
	// Whether or not to include events from unconfirmed transactions.
	// Optional. Use `false` as default.
	Unanchored bool
	// Hex representation of the token's unique value.
	Value string
	// Token asset class identifier.
	AssetIdentifier string
}

// Defines parameters for GetTransactionByID
type GetTransactionByIDParams struct{}

// Defines parameters for GetNamesOwnedByAddress
type GetNamesOwnedByAddressParams struct {
	// The address to lookup.
	Address string
	// The layer-1 blockchain for the address.
	Blockchain string
}

// Defines parameters for GetAccountAssets
type GetAccountAssetsParams struct {
	// Returned data representing the state at that point in time, rather than the current block.
	// Optional. Pass an empty string to use the default value.
	UntilBlock string
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
	// Index of first account assets to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Max number of account assets to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
}

// Defines parameters for GetBlockByHash
type GetBlockByHashParams struct{}

// Defines parameters for GetContractEventsByID
type GetContractEventsByIDParams struct {
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Contract identifier formatted as `<contract_address>.<contract_name>`.
	ContractID string
	// Index of first contract event to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Max number of contract events to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
}

// Defines parameters for GetFTMetadataList
type GetFTMetadataListParams struct {
	// Index of first tokens to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Max number of tokens to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
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
	// Stacks address or a Contract identifier.
	// Optional. Pass an empty string to use the default value.
	Address string
	// Hash of transaction.
	// Optional. Pass an empty string to use the default value.
	TXID string
	// Number of items to skip.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Filter the events on event type.
	// [SmartContractLog STXLock STXAsset FungibleTokenAsset NonFungibleTokenAsset]
	Type GetFilteredEventsType
}

// Defines parameters for GetDroppedMempoolTransactionList
type GetDroppedMempoolTransactionListParams struct {
	// Index of first mempool transaction to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Max number of mempool transactions to fetch.
	// Optional. Use `-1` to use the default value (96).
	// Max value is 200.
	Limit int
}

// Defines parameters for GetRawTransactionByID
type GetRawTransactionByIDParams struct{}

// Defines parameters for GetAllNames
type GetAllNamesParams struct {
	// Names are returned in pages of size 100, so specify the page number.
	Page int
}

// Defines parameters for GetNamespaceNames
type GetNamespaceNamesParams struct {
	// Names are returned in pages of size 100, so specify the page number.
	Page int
	// The namespace to fetch names from.
	TLD string
}

// Defines parameters for GetContractInterface
type GetContractInterfaceParams struct{}

// Defines parameters for GetNetworkBlockTimeByNetwork
type GetNetworkBlockTimeByNetworkParams struct {
	// Which network to retrieve the target block time of.
	Network string
}

// Defines parameters for GetSTXSupply
type GetSTXSupplyParams struct {
	// The block height at which to query supply details from, if not provided then the latest block height is used.
	Height int
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
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Max number of transactions to fetch.
	// Optional. Use `-1` to use the default value (96).
	// Max value is 200.
	Limit int
	// Index of first transaction to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Filter by transaction type.
	// [Coinbase TokenTransfer SmartContract ContractCall PoisonMicroblock]
	Type GetTransactionListType
}

// Defines parameters for GetNameInfo
type GetNameInfoParams struct {
	// Fully-qualified name.
	Name string
}

// Defines parameters for GetFeeTransfer
type GetFeeTransferParams struct{}

// Defines parameters for GetBurnchainRewardSlotHoldersByAddress
type GetBurnchainRewardSlotHoldersByAddressParams struct {
	// Index of the first items to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Reward slot holder recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format.
	Address string
	// Max number of items to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
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

// Defines parameters for GetContractByID
type GetContractByIDParams struct{}
type CallReadOnlyFunctionArguments struct {
	// The simulated tx-sender
	Sender string
	// An array of hex serialized Clarity values
	Arguments []string
}
