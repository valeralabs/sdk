package api

// GetSTXSupplyCirculatingPlainParams defines parameters for GetSTXSupplyCirculatingPlain
type GetSTXSupplyCirculatingPlainParams struct{}

// GetTXListDetailsParams defines parameters for GetTXListDetails
type GetTXListDetailsParams struct{}

// GetTransactionByIDParams defines parameters for GetTransactionByID
type GetTransactionByIDParams struct{}

// GetAllNamesParams defines parameters for GetAllNames
type GetAllNamesParams struct {
	// Names are returned in pages of size 100, so specify the page number.
	Page int
}

// GetContractDataMapEntryParams defines parameters for GetContractDataMapEntry
type GetContractDataMapEntryParams struct {
	// The Stacks chain tip to query from.
	// Optional. Pass an empty string to use the default value.
	Tip string
	// Map name.
	MapName string
	// Returns object without the proof field when set to 0.
	// Optional. Use `-1` to use the default value.
	Proof int
	// Contract name.
	ContractName string
	// Stacks address.
	ContractAddress string
}

// GetNamePriceParams defines parameters for GetNamePrice
type GetNamePriceParams struct {
	// The name to query price information for.
	Name string
}

// GetAccountBalanceParams defines parameters for GetAccountBalance
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

// GetMicroblockListParams defines parameters for GetMicroblockList
type GetMicroblockListParams struct {
	// Index of the first microblock to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Max number of microblocks to fetch.
	// Optional. Use `-1` to use the default value (20).
	// Max value is 200.
	Limit int
}

// PostCoreNodeTransactionsParams defines parameters for PostCoreNodeTransactions
type PostCoreNodeTransactionsParams struct{}

// GetContractsByTraitParams defines parameters for GetContractsByTrait
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

// GetContractEventsByIDParams defines parameters for GetContractEventsByID
type GetContractEventsByIDParams struct {
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Index of first contract event to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Contract identifier formatted as `<contract_address>.<contract_name>`.
	ContractID string
	// Max number of contract events to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
}

// RunFaucetSTXParams defines parameters for RunFaucetSTX
type RunFaucetSTXParams struct {
	// Use required number of tokens for stacking.
	Stacking bool
	// STX testnet address.
	Address string
}

// GetFTMetadataListParams defines parameters for GetFTMetadataList
type GetFTMetadataListParams struct {
	// Index of first tokens to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Max number of tokens to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
}

// GetNFTHoldingsParams defines parameters for GetNFTHoldings
type GetNFTHoldingsParams struct {
	// Whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times.
	// Optional. Use `false` as default.
	TXMetadata bool
	// Max number of tokens to fetch.
	// Optional. Use `-1` to use the default value (50).
	Limit int
	// Whether or not to include tokens from unconfirmed transactions.
	// Optional. Use `false` as default.
	Unanchored bool
	// Token owner's STX address or Smart Contract ID.
	Principal string
	// Index of first tokens to fetch.
	// Optional. Use `-1` to use the default value (0).
	Offset int
	// Identifiers of the token asset classes to filter for.
}

// GetNFTMetadataListParams defines parameters for GetNFTMetadataList
type GetNFTMetadataListParams struct {
	// Index of first tokens to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Max number of tokens to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
}

// GetBlockByBurnBlockHashParams defines parameters for GetBlockByBurnBlockHash
type GetBlockByBurnBlockHashParams struct{}

// GetBurnchainRewardsTotalByAddressParams defines parameters for GetBurnchainRewardsTotalByAddress
type GetBurnchainRewardsTotalByAddressParams struct {
	// Reward recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format.
	Address string
}
type GetTransactionListType int64

const (
	Coinbase GetTransactionListType = iota
	TokenTransfer
	SmartContract
	ContractCall
	PoisonMicroblock
)

// GetTransactionListParams defines parameters for GetTransactionList
type GetTransactionListParams struct {
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Index of first transaction to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Max number of transactions to fetch.
	// Optional. Use `-1` to use the default value (96).
	// Max value is 200.
	Limit int
	// Filter by transaction type.
	// [Coinbase TokenTransfer SmartContract ContractCall PoisonMicroblock]
	Type GetTransactionListType
}

// GetRawTransactionByIDParams defines parameters for GetRawTransactionByID
type GetRawTransactionByIDParams struct{}

// SearchByIDParams defines parameters for SearchByID
type SearchByIDParams struct {
	// This includes the detailed data for purticular hash in the response.
	// Optional. Use `<nil>` as default.
	IncludeMetadata bool
	// The hex hash string for a block or transaction, account address, or contract address.
	ID string
}

// GetBurnchainRewardListByAddressParams defines parameters for GetBurnchainRewardListByAddress
type GetBurnchainRewardListByAddressParams struct {
	// Index of first rewards to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Max number of rewards to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Reward recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format.
	Address string
}

// FetchFeeRateParams defines parameters for FetchFeeRate
type FetchFeeRateParams struct {
	// A serialized transaction.
	Transaction string
}

// GetHistoricalZoneFileParams defines parameters for GetHistoricalZoneFile
type GetHistoricalZoneFileParams struct {
	// Zone file hash.
	ZoneFileHash string
	// Fully-qualified name.
	Name string
}

// GetFeeTransferParams defines parameters for GetFeeTransfer
type GetFeeTransferParams struct{}

// GetAccountAssetsParams defines parameters for GetAccountAssets
type GetAccountAssetsParams struct {
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
	// Returned data representing the state at that point in time, rather than the current block.
	// Optional. Pass an empty string to use the default value.
	UntilBlock string
	// Max number of account assets to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Index of first account assets to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
}

// GetAccountNoncesParams defines parameters for GetAccountNonces
type GetAccountNoncesParams struct {
	// Optionally get the nonce at a given block hash.
	// Optional. Pass an empty string to use the default value.
	BlockHash string
	// Stacks address (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0`).
	Principal string
	// Optionally get the nonce at a given block height.
	BlockHeight int
}
type GetFilteredEventsType int64

const (
	SmartContractLog GetFilteredEventsType = iota
	STXLock
	STXAsset
	FungibleTokenAsset
	NonFungibleTokenAsset
)

// GetFilteredEventsParams defines parameters for GetFilteredEvents
type GetFilteredEventsParams struct {
	// Stacks address or a Contract identifier.
	// Optional. Pass an empty string to use the default value.
	Address string
	// Filter the events on event type.
	// [SmartContractLog STXLock STXAsset FungibleTokenAsset NonFungibleTokenAsset]
	Type GetFilteredEventsType
	// Hash of transaction.
	// Optional. Pass an empty string to use the default value.
	TXID string
	// Number of items to skip.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Number of items to return.
	// Optional. Use `-1` to use the default value.
	Limit int
}

// GetNamesOwnedByAddressParams defines parameters for GetNamesOwnedByAddress
type GetNamesOwnedByAddressParams struct {
	// The address to lookup.
	Address string
	// The layer-1 blockchain for the address.
	Blockchain string
}

// GetUnanchoredTxsParams defines parameters for GetUnanchoredTxs
type GetUnanchoredTxsParams struct{}

// GetNFTHistoryParams defines parameters for GetNFTHistory
type GetNFTHistoryParams struct {
	// Whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times.
	// Optional. Use `false` as default.
	TXMetadata bool
	// Max number of events to fetch.
	// Optional. Use `-1` to use the default value (50).
	Limit int
	// Token asset class identifier.
	AssetIdentifier string
	// Index of first event to fetch.
	// Optional. Use `-1` to use the default value (0).
	Offset int
	// Whether or not to include events from unconfirmed transactions.
	// Optional. Use `false` as default.
	Unanchored bool
	// Hex representation of the token's unique value.
	Value string
}

// GetTotalSTXSupplyLegacyFormatParams defines parameters for GetTotalSTXSupplyLegacyFormat
type GetTotalSTXSupplyLegacyFormatParams struct {
	// The block height at which to query supply details from, if not provided then the latest block height is used.
	Height int
}

// GetAccountTransactionsParams defines parameters for GetAccountTransactions
type GetAccountTransactionsParams struct {
	// Returned data representing the state up until that point in time, rather than the current block.
	// Optional. Pass an empty string to use the default value.
	UntilBlock string
	// Index of first account transaction to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
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
}

// GetSTXSupplyParams defines parameters for GetSTXSupply
type GetSTXSupplyParams struct {
	// The block height at which to query supply details from, if not provided then the latest block height is used.
	Height int
}

// GetTransactionsByBlockHashParams defines parameters for GetTransactionsByBlockHash
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

// GetSingleTransactionWithTransfersParams defines parameters for GetSingleTransactionWithTransfers
type GetSingleTransactionWithTransfersParams struct {
	// Transaction id.
	TXID string
	// Stacks address or a contract identifier.
	Principal string
}

// GetContractNFTMetadataParams defines parameters for GetContractNFTMetadata
type GetContractNFTMetadataParams struct {
	// Token's contract id.
	ContractId string
}

// GetBlockByBurnBlockHeightParams defines parameters for GetBlockByBurnBlockHeight
type GetBlockByBurnBlockHeightParams struct{}

// GetNFTMintsParams defines parameters for GetNFTMints
type GetNFTMintsParams struct {
	// Max number of events to fetch.
	// Optional. Use `-1` to use the default value (50).
	Limit int
	// Whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times.
	// Optional. Use `false` as default.
	TXMetadata bool
	// Index of first event to fetch.
	// Optional. Use `-1` to use the default value (0).
	Offset int
	// Whether or not to include events from unconfirmed transactions.
	// Optional. Use `false` as default.
	Unanchored bool
	// Token asset class identifier.
	AssetIdentifier string
}

// GetContractFTMetadataParams defines parameters for GetContractFTMetadata
type GetContractFTMetadataParams struct {
	// Token's contract id.
	ContractId string
}

// GetAccountSTXBalanceParams defines parameters for GetAccountSTXBalance
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

// GetAccountInboundParams defines parameters for GetAccountInbound
type GetAccountInboundParams struct {
	// Returned data representing the state up until that point in time, rather than the current block.
	// Optional. Pass an empty string to use the default value.
	UntilBlock string
	// Number of items to return.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
	// Filter for transfers only at this given block height.
	Height int
	// Number of items to skip.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
}

// GetAllNamespacesParams defines parameters for GetAllNamespaces
type GetAllNamespacesParams struct{}

// GetAccountTransactionsWithTransfersParams defines parameters for GetAccountTransactionsWithTransfers
type GetAccountTransactionsWithTransfersParams struct {
	// Returned data representing the state up until that point in time, rather than the current block.
	// Optional. Pass an empty string to use the default value.
	UntilBlock string
	// Filter for transactions only at this given block height.
	Height int
	// Index of first account transaction to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
	// Max number of account transactions to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
}

// GetAccountInfoParams defines parameters for GetAccountInfo
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

// GetMempoolTransactionListParams defines parameters for GetMempoolTransactionList
type GetMempoolTransactionListParams struct {
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Filter to only return transactions with this address as the sender or recipient (recipient only applicable for STX transfer tx types).
	// Optional. Pass an empty string to use the default value.
	Address string
	// Filter to only return transactions with this sender address.
	// Optional. Pass an empty string to use the default value.
	SenderAddress string
	// Filter to only return transactions with this recipient address (only applicable for STX transfer tx types).
	// Optional. Pass an empty string to use the default value.
	RecipientAddress string
	// Index of first mempool transaction to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Max number of mempool transactions to fetch.
	// Optional. Use `-1` to use the default value (96).
	// Max value is 200.
	Limit int
}

// FetchZoneFileParams defines parameters for FetchZoneFile
type FetchZoneFileParams struct {
	// Fully-qualified name.
	Name string
}

// GetContractInterfaceParams defines parameters for GetContractInterface
type GetContractInterfaceParams struct{}

// GetPoxInfoParams defines parameters for GetPoxInfo
type GetPoxInfoParams struct{}

// GetBurnchainRewardSlotHoldersParams defines parameters for GetBurnchainRewardSlotHolders
type GetBurnchainRewardSlotHoldersParams struct {
	// Index of the first items to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Max number of items to fetch.
	// Optional. Use `-1` to use the default value (96).
	// Max value is 250.
	Limit int
}

// GetTransactionsByBlockHeightParams defines parameters for GetTransactionsByBlockHeight
type GetTransactionsByBlockHeightParams struct {
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Max number of transactions to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Height of block.
	Height int
	// Index of first transaction to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
}

// GetNetworkBlockTimesParams defines parameters for GetNetworkBlockTimes
type GetNetworkBlockTimesParams struct{}

// CallReadOnlyFunctionParams defines parameters for CallReadOnlyFunction
type CallReadOnlyFunctionParams struct {
	// The simulated tx-sender.
	Sender string
	// Function name.
	FunctionName string
	// The Stacks chain tip to query from.
	// Optional. Pass an empty string to use the default value.
	Tip string
	// An array of hex serialized Clarity values.
	Arguments CallReadOnlyFunctionArguments
	// Stacks address.
	ContractAddress string
	// Contract name.
	ContractName string
}

// PostFeeTransactionParams defines parameters for PostFeeTransaction
type PostFeeTransactionParams struct {
	//
	TransactionPayload string
	//
	EstimatedLen int
}

// GetBlockByHeightParams defines parameters for GetBlockByHeight
type GetBlockByHeightParams struct{}

// GetContractByIDParams defines parameters for GetContractByID
type GetContractByIDParams struct{}

// GetDroppedMempoolTransactionListParams defines parameters for GetDroppedMempoolTransactionList
type GetDroppedMempoolTransactionListParams struct {
	// Index of first mempool transaction to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Max number of mempool transactions to fetch.
	// Optional. Use `-1` to use the default value (96).
	// Max value is 200.
	Limit int
}

// GetCoreAPIInfoParams defines parameters for GetCoreAPIInfo
type GetCoreAPIInfoParams struct{}

// GetBlockByHashParams defines parameters for GetBlockByHash
type GetBlockByHashParams struct{}

// GetMicroblockByHashParams defines parameters for GetMicroblockByHash
type GetMicroblockByHashParams struct{}

// GetStatusParams defines parameters for GetStatus
type GetStatusParams struct{}

// GetNamespacePriceParams defines parameters for GetNamespacePrice
type GetNamespacePriceParams struct {
	// The namespace to fetch price for.
	TLD string
}

// GetAccountNFTParams defines parameters for GetAccountNFT
type GetAccountNFTParams struct {
	// Returned data representing the state up until that point in time, rather than the current block.
	// Optional. Pass an empty string to use the default value.
	UntilBlock string
	// Number of items to return.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
	// Number of items to skip.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
}

// RunFaucetBTCParams defines parameters for RunFaucetBTC
type RunFaucetBTCParams struct {
	// BTC testnet address.
	Address string
}

// GetBurnchainRewardListParams defines parameters for GetBurnchainRewardList
type GetBurnchainRewardListParams struct {
	// Index of first rewards to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Max number of rewards to fetch.
	// Optional. Use `-1` to use the default value (96).
	// Max value is 250.
	Limit int
}

// GetNetworkBlockTimeByNetworkParams defines parameters for GetNetworkBlockTimeByNetwork
type GetNetworkBlockTimeByNetworkParams struct {
	// Which network to retrieve the target block time of.
	Network string
}

// GetSTXSupplyTotalSupplyPlainParams defines parameters for GetSTXSupplyTotalSupplyPlain
type GetSTXSupplyTotalSupplyPlainParams struct{}

// GetAddressMempoolTransactionsParams defines parameters for GetAddressMempoolTransactions
type GetAddressMempoolTransactionsParams struct {
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Max number of transactions to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Transactions for the address.
	Address string
	// Index of first transaction to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
}

// GetBlockListParams defines parameters for GetBlockList
type GetBlockListParams struct {
	// Index of first block to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Max number of blocks to fetch.
	// Optional. Use `-1` to use the default value (20).
	// Max value is 200.
	Limit int
}

// FetchSubdomainsListForNameParams defines parameters for FetchSubdomainsListForName
type FetchSubdomainsListForNameParams struct {
	// Fully-qualified name.
	Name string
}

// GetNamespaceNamesParams defines parameters for GetNamespaceNames
type GetNamespaceNamesParams struct {
	// Names are returned in pages of size 100, so specify the page number.
	Page int
	// The namespace to fetch names from.
	TLD string
}

// GetContractSourceParams defines parameters for GetContractSource
type GetContractSourceParams struct{}

// GetNameInfoParams defines parameters for GetNameInfo
type GetNameInfoParams struct {
	// Fully-qualified name.
	Name string
}

// GetBurnchainRewardSlotHoldersByAddressParams defines parameters for GetBurnchainRewardSlotHoldersByAddress
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
type CallReadOnlyFunctionArguments struct {
	// The simulated tx-sender
	Sender string
	// An array of hex serialized Clarity values
	Arguments []string
}
