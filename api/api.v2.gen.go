package api

// GetContractSourceParams defines parameters for GetContractSource
type GetContractSourceParams struct{}

// GetNetworkBlockTimeByNetworkParams defines parameters for GetNetworkBlockTimeByNetwork
type GetNetworkBlockTimeByNetworkParams struct {
	// Which network to retrieve the target block time of
	Network string
}

// GetBlockByBurnBlockHashParams defines parameters for GetBlockByBurnBlockHash
type GetBlockByBurnBlockHashParams struct{}

// GetBlockByHashParams defines parameters for GetBlockByHash
type GetBlockByHashParams struct{}

// GetMicroblockListParams defines parameters for GetMicroblockList
type GetMicroblockListParams struct {
	// Index of the first microblock to fetch
	Offset int
	// Max number of microblocks to fetch
	Limit int
}

// GetNFTMintsParams defines parameters for GetNFTMints
type GetNFTMintsParams struct {
	// whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times.
	TXMetadata bool
	// index of first event to fetch
	Offset int
	// max number of events to fetch
	Limit int
	// whether or not to include events from unconfirmed transactions
	Unanchored bool
	// token asset class identifier
	AssetIdentifier string
}

// GetAddressMempoolTransactionsParams defines parameters for GetAddressMempoolTransactions
type GetAddressMempoolTransactionsParams struct {
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	Unanchored bool
	// index of first transaction to fetch
	Offset int
	// Transactions for the address
	Address string
	// max number of transactions to fetch
	Limit int
}

// FetchSubdomainsListForNameParams defines parameters for FetchSubdomainsListForName
type FetchSubdomainsListForNameParams struct {
	// fully-qualified name
	Name string
}

// SearchByIDParams defines parameters for SearchByID
type SearchByIDParams struct {
	// This includes the detailed data for purticular hash in the response
	IncludeMetadata bool
	// The hex hash string for a block or transaction, account address, or contract address
	ID string
}

// GetAccountTransactionsParams defines parameters for GetAccountTransactions
type GetAccountTransactionsParams struct {
	// returned data representing the state up until that point in time, rather than the current block.
	UntilBlock string
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
	Principal string
	// max number of account transactions to fetch
	Limit int
	// index of first account transaction to fetch
	Offset int
	// Filter for transactions only at this given block height
	Height int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	Unanchored bool
}

// GetContractEventsByIDParams defines parameters for GetContractEventsByID
type GetContractEventsByIDParams struct {
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	Unanchored bool
	// Contract identifier formatted as `<contract_address>.<contract_name>`
	ContractID string
	// max number of contract events to fetch
	Limit int
	// index of first contract event to fetch
	Offset int
}

// GetStatusParams defines parameters for GetStatus
type GetStatusParams struct{}

// GetMempoolTransactionListParams defines parameters for GetMempoolTransactionList
type GetMempoolTransactionListParams struct {
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	Unanchored bool
	// Filter to only return transactions with this sender address.
	SenderAddress string
	// max number of mempool transactions to fetch
	Limit int
	// Filter to only return transactions with this address as the sender or recipient (recipient only applicable for STX transfer tx types).
	Address string
	// index of first mempool transaction to fetch
	Offset int
	// Filter to only return transactions with this recipient address (only applicable for STX transfer tx types).
	RecipientAddress string
}

// GetTXListDetailsParams defines parameters for GetTXListDetails
type GetTXListDetailsParams struct{}

// GetAccountAssetsParams defines parameters for GetAccountAssets
type GetAccountAssetsParams struct {
	// max number of account assets to fetch
	Limit int
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
	Principal string
	// index of first account assets to fetch
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	Unanchored bool
	// returned data representing the state at that point in time, rather than the current block.
	UntilBlock string
}

// GetMicroblockByHashParams defines parameters for GetMicroblockByHash
type GetMicroblockByHashParams struct{}

// GetSTXSupplyCirculatingPlainParams defines parameters for GetSTXSupplyCirculatingPlain
type GetSTXSupplyCirculatingPlainParams struct{}

// GetContractNFTMetadataParams defines parameters for GetContractNFTMetadata
type GetContractNFTMetadataParams struct {
	// token's contract id
	ContractId string
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
	// Stacks address or a Contract identifier
	Address string
	// Hash of transaction
	TXID string
	// number of items to skip
	Offset int
	// number of items to return
	Limit int
	// Filter the events on event type
	// [SmartContractLog STXLock STXAsset FungibleTokenAsset NonFungibleTokenAsset]
	Type GetFilteredEventsType
}

// GetRawTransactionByIDParams defines parameters for GetRawTransactionByID
type GetRawTransactionByIDParams struct{}

// GetAccountTransactionsWithTransfersParams defines parameters for GetAccountTransactionsWithTransfers
type GetAccountTransactionsWithTransfersParams struct {
	// returned data representing the state up until that point in time, rather than the current block.
	UntilBlock string
	// index of first account transaction to fetch
	Offset int
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
	Principal string
	// max number of account transactions to fetch
	Limit int
	// Filter for transactions only at this given block height
	Height int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	Unanchored bool
}

// GetBlockByHeightParams defines parameters for GetBlockByHeight
type GetBlockByHeightParams struct{}

// GetContractsByTraitParams defines parameters for GetContractsByTrait
type GetContractsByTraitParams struct {
	// index of first contract event to fetch
	Offset int
	// max number of contracts fetch
	Limit int
	// JSON abi of the trait.
	TraitABI string
}

// GetNetworkBlockTimesParams defines parameters for GetNetworkBlockTimes
type GetNetworkBlockTimesParams struct{}

// GetNamesOwnedByAddressParams defines parameters for GetNamesOwnedByAddress
type GetNamesOwnedByAddressParams struct {
	// the address to lookup
	Address string
	// the layer-1 blockchain for the address
	Blockchain string
}

// GetContractInterfaceParams defines parameters for GetContractInterface
type GetContractInterfaceParams struct{}

// GetAccountNoncesParams defines parameters for GetAccountNonces
type GetAccountNoncesParams struct {
	// Optionally get the nonce at a given block hash
	BlockHash string
	// Stacks address (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0`)
	Principal string
	// Optionally get the nonce at a given block height
	BlockHeight int
}

// GetBurnchainRewardsTotalByAddressParams defines parameters for GetBurnchainRewardsTotalByAddress
type GetBurnchainRewardsTotalByAddressParams struct {
	// Reward recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format
	Address string
}

// GetContractFTMetadataParams defines parameters for GetContractFTMetadata
type GetContractFTMetadataParams struct {
	// token's contract id
	ContractId string
}

// GetNameInfoParams defines parameters for GetNameInfo
type GetNameInfoParams struct {
	// fully-qualified name
	Name string
}

// GetAllNamespacesParams defines parameters for GetAllNamespaces
type GetAllNamespacesParams struct{}

// GetNamePriceParams defines parameters for GetNamePrice
type GetNamePriceParams struct {
	// the name to query price information for
	Name string
}

// GetBurnchainRewardListByAddressParams defines parameters for GetBurnchainRewardListByAddress
type GetBurnchainRewardListByAddressParams struct {
	// index of first rewards to fetch
	Offset int
	// Reward recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format
	Address string
	// max number of rewards to fetch
	Limit int
}

// GetSingleTransactionWithTransfersParams defines parameters for GetSingleTransactionWithTransfers
type GetSingleTransactionWithTransfersParams struct {
	// Transaction id
	TXID string
	// Stacks address or a contract identifier
	Principal string
}

// GetBlockListParams defines parameters for GetBlockList
type GetBlockListParams struct {
	// index of first block to fetch
	Offset int
	// max number of blocks to fetch
	Limit int
}

// GetHistoricalZoneFileParams defines parameters for GetHistoricalZoneFile
type GetHistoricalZoneFileParams struct {
	// zone file hash
	ZoneFileHash string
	// fully-qualified name
	Name string
}

// GetCoreAPIInfoParams defines parameters for GetCoreAPIInfo
type GetCoreAPIInfoParams struct{}

// GetPoxInfoParams defines parameters for GetPoxInfo
type GetPoxInfoParams struct{}

// GetAccountInboundParams defines parameters for GetAccountInbound
type GetAccountInboundParams struct {
	// returned data representing the state up until that point in time, rather than the current block.
	UntilBlock string
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
	Principal string
	// number of items to return
	Limit int
	// number of items to skip
	Offset int
	// Filter for transfers only at this given block height
	Height int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	Unanchored bool
}

// GetBlockByBurnBlockHeightParams defines parameters for GetBlockByBurnBlockHeight
type GetBlockByBurnBlockHeightParams struct{}

// GetBurnchainRewardSlotHoldersByAddressParams defines parameters for GetBurnchainRewardSlotHoldersByAddress
type GetBurnchainRewardSlotHoldersByAddressParams struct {
	// index of the first items to fetch
	Offset int
	// Reward slot holder recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format
	Address string
	// max number of items to fetch
	Limit int
}

// PostFeeTransactionParams defines parameters for PostFeeTransaction
type PostFeeTransactionParams struct {
	EstimatedLen       int
	TransactionPayload string
}

// GetAccountSTXBalanceParams defines parameters for GetAccountSTXBalance
type GetAccountSTXBalanceParams struct {
	// returned data representing the state up until that point in time, rather than the current block.
	UntilBlock string
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	Unanchored bool
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
	Principal string
}

// GetNamespaceNamesParams defines parameters for GetNamespaceNames
type GetNamespaceNamesParams struct {
	// names are returned in pages of size 100, so specify the page number.
	Page int
	// the namespace to fetch names from
	TLD string
}

// GetAccountInfoParams defines parameters for GetAccountInfo
type GetAccountInfoParams struct {
	// The Stacks chain tip to query from
	Tip string
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
	Principal string
	// Returns object without the proof field if set to 0
	Proof int
}

// GetFeeTransferParams defines parameters for GetFeeTransfer
type GetFeeTransferParams struct{}

// GetTransactionsByBlockHeightParams defines parameters for GetTransactionsByBlockHeight
type GetTransactionsByBlockHeightParams struct {
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	Unanchored bool
	// Height of block
	Height int
	// max number of transactions to fetch
	Limit int
	// index of first transaction to fetch
	Offset int
}

// GetNFTHistoryParams defines parameters for GetNFTHistory
type GetNFTHistoryParams struct {
	// whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times.
	TXMetadata bool
	// token asset class identifier
	AssetIdentifier string
	// hex representation of the token's unique value
	Value string
	// max number of events to fetch
	Limit int
	// index of first event to fetch
	Offset int
	// whether or not to include events from unconfirmed transactions
	Unanchored bool
}

// GetNFTMetadataListParams defines parameters for GetNFTMetadataList
type GetNFTMetadataListParams struct {
	// index of first tokens to fetch
	Offset int
	// max number of tokens to fetch
	Limit int
}

// GetTransactionByIDParams defines parameters for GetTransactionByID
type GetTransactionByIDParams struct{}

// GetSTXSupplyTotalSupplyPlainParams defines parameters for GetSTXSupplyTotalSupplyPlain
type GetSTXSupplyTotalSupplyPlainParams struct{}

// GetContractByIDParams defines parameters for GetContractByID
type GetContractByIDParams struct{}

// GetNFTHoldingsParams defines parameters for GetNFTHoldings
type GetNFTHoldingsParams struct {
	// whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times.
	TXMetadata bool
	// index of first tokens to fetch
	Offset int
	// whether or not to include tokens from unconfirmed transactions
	Unanchored bool
	// max number of tokens to fetch
	Limit int
	// token owner's STX address or Smart Contract ID
	Principal string
}

// GetAllNamesParams defines parameters for GetAllNames
type GetAllNamesParams struct {
	// names are returned in pages of size 100, so specify the page number.
	Page int
}

// GetContractDataMapEntryParams defines parameters for GetContractDataMapEntry
type GetContractDataMapEntryParams struct {
	// The Stacks chain tip to query from
	Tip string
	// Contract name
	ContractName string
	// Map name
	MapName string
	// Returns object without the proof field when set to 0
	Proof int
	// Stacks address
	ContractAddress string
}

// GetBurnchainRewardListParams defines parameters for GetBurnchainRewardList
type GetBurnchainRewardListParams struct {
	// index of first rewards to fetch
	Offset int
	// max number of rewards to fetch
	Limit int
}

// RunFaucetBTCParams defines parameters for RunFaucetBTC
type RunFaucetBTCParams struct {
	// BTC testnet address
	Address string
}

// RunFaucetSTXParams defines parameters for RunFaucetSTX
type RunFaucetSTXParams struct {
	// Use required number of tokens for stacking
	Stacking bool
	// STX testnet address
	Address string
}

// GetUnanchoredTxsParams defines parameters for GetUnanchoredTxs
type GetUnanchoredTxsParams struct{}

// GetTotalSTXSupplyLegacyFormatParams defines parameters for GetTotalSTXSupplyLegacyFormat
type GetTotalSTXSupplyLegacyFormatParams struct {
	// The block height at which to query supply details from, if not provided then the latest block height is used
	Height int
}

// PostCoreNodeTransactionsParams defines parameters for PostCoreNodeTransactions
type PostCoreNodeTransactionsParams struct{}

// GetAccountBalanceParams defines parameters for GetAccountBalance
type GetAccountBalanceParams struct {
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
	Principal string
	// returned data representing the state up until that point in time, rather than the current block.
	UntilBlock string
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	Unanchored bool
}

// GetFTMetadataListParams defines parameters for GetFTMetadataList
type GetFTMetadataListParams struct {
	// index of first tokens to fetch
	Offset int
	// max number of tokens to fetch
	Limit int
}

// GetSTXSupplyParams defines parameters for GetSTXSupply
type GetSTXSupplyParams struct {
	// The block height at which to query supply details from, if not provided then the latest block height is used
	Height int
}

// FetchFeeRateParams defines parameters for FetchFeeRate
type FetchFeeRateParams struct {
	// A serialized transaction
	Transaction string
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
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	Unanchored bool
	// index of first transaction to fetch
	Offset int
	// max number of transactions to fetch
	Limit int
	// Filter by transaction type
	// [Coinbase TokenTransfer SmartContract ContractCall PoisonMicroblock]
	Type GetTransactionListType
}

// GetTransactionsByBlockHashParams defines parameters for GetTransactionsByBlockHash
type GetTransactionsByBlockHashParams struct {
	// index of first transaction to fetch
	Offset int
	// max number of transactions to fetch
	Limit int
	// Hash of block
	BlockHash string
}

// GetDroppedMempoolTransactionListParams defines parameters for GetDroppedMempoolTransactionList
type GetDroppedMempoolTransactionListParams struct {
	// index of first mempool transaction to fetch
	Offset int
	// max number of mempool transactions to fetch
	Limit int
}

// GetBurnchainRewardSlotHoldersParams defines parameters for GetBurnchainRewardSlotHolders
type GetBurnchainRewardSlotHoldersParams struct {
	// index of the first items to fetch
	Offset int
	// max number of items to fetch
	Limit int
}

// FetchZoneFileParams defines parameters for FetchZoneFile
type FetchZoneFileParams struct {
	// fully-qualified name
	Name string
}

// CallReadOnlyFunctionParams defines parameters for CallReadOnlyFunction
type CallReadOnlyFunctionParams struct {
	// The simulated tx-sender
	Sender string
	// Function name
	FunctionName string
	// The Stacks chain tip to query from
	Tip string
	// An array of hex serialized Clarity values
	// Stacks address
	ContractAddress string
	Arguments       any
	// Contract name
	ContractName string
}

// GetNamespacePriceParams defines parameters for GetNamespacePrice
type GetNamespacePriceParams struct {
	// the namespace to fetch price for
	TLD string
}

// GetAccountNFTParams defines parameters for GetAccountNFT
type GetAccountNFTParams struct {
	// returned data representing the state up until that point in time, rather than the current block.
	UntilBlock string
	// number of items to return
	Limit int
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
	Principal string
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	Unanchored bool
	// number of items to skip
	Offset int
}
