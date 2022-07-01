package api

// GetAccountTransactionsWithTransfersParams defines parameters for GetAccountTransactionsWithTransfers
type GetAccountTransactionsWithTransfersParams struct {
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
	// returned data representing the state up until that point in time, rather than the current block.
	UntilBlock string
}

// GetSingleTransactionWithTransfersParams defines parameters for GetSingleTransactionWithTransfers
type GetSingleTransactionWithTransfersParams struct {
	// Stacks address or a contract identifier
	Principal string
	// Transaction id
	TXID string
}

// GetBlockListParams defines parameters for GetBlockList
type GetBlockListParams struct {
	// max number of blocks to fetch
	Limit int
	// index of first block to fetch
	Offset int
}

// GetTransactionListParams defines parameters for GetTransactionList
type GetTransactionListParams struct {
	// max number of transactions to fetch
	Limit int
	// index of first transaction to fetch
	Offset int
	// Filter by transaction type
	// [Coinbase TokenTransfer SmartContract ContractCall PoisonMicroblock]
	Type Type
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	Unanchored bool
}

// GetContractDataMapEntryParams defines parameters for GetContractDataMapEntry
type GetContractDataMapEntryParams struct {
	// Stacks address
	ContractAddress string
	// Contract name
	ContractName string
	// Map name
	MapName string
	// Returns object without the proof field when set to 0
	Proof int
	// The Stacks chain tip to query from
	Tip string
}

// GetBlockByHeightParams defines parameters for GetBlockByHeight
type GetBlockByHeightParams struct{}

// FetchSubdomainsListForNameParams defines parameters for FetchSubdomainsListForName
type FetchSubdomainsListForNameParams struct {
	// fully-qualified name
	Name string
}

// FetchZoneFileParams defines parameters for FetchZoneFile
type FetchZoneFileParams struct {
	// fully-qualified name
	Name string
}

// GetPoxInfoParams defines parameters for GetPoxInfo
type GetPoxInfoParams struct{}

// GetAccountBalanceParams defines parameters for GetAccountBalance
type GetAccountBalanceParams struct {
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
	Principal string
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	Unanchored bool
	// returned data representing the state up until that point in time, rather than the current block.
	UntilBlock string
}

// GetAccountSTXBalanceParams defines parameters for GetAccountSTXBalance
type GetAccountSTXBalanceParams struct {
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
	Principal string
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	Unanchored bool
	// returned data representing the state up until that point in time, rather than the current block.
	UntilBlock string
}

// GetNFTHoldingsParams defines parameters for GetNFTHoldings
type GetNFTHoldingsParams struct {
	// token owner's STX address or Smart Contract ID
	Principal string
	// max number of tokens to fetch
	Limit int
	// index of first tokens to fetch
	Offset int
	// whether or not to include tokens from unconfirmed transactions
	Unanchored bool
	// whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times.
	TXMetadata bool
}

// GetContractNFTMetadataParams defines parameters for GetContractNFTMetadata
type GetContractNFTMetadataParams struct {
	// token's contract id
	ContractId string
}

// PostCoreNodeTransactionsParams defines parameters for PostCoreNodeTransactions
type PostCoreNodeTransactionsParams struct{}

// GetTransactionsByBlockHashParams defines parameters for GetTransactionsByBlockHash
type GetTransactionsByBlockHashParams struct {
	// Hash of block
	BlockHash string
	// max number of transactions to fetch
	Limit int
	// index of first transaction to fetch
	Offset int
}

// GetNameInfoParams defines parameters for GetNameInfo
type GetNameInfoParams struct {
	// fully-qualified name
	Name string
}

// GetAccountInfoParams defines parameters for GetAccountInfo
type GetAccountInfoParams struct {
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
	Principal string
	// Returns object without the proof field if set to 0
	Proof int
	// The Stacks chain tip to query from
	Tip string
}

// GetBlockByBurnBlockHashParams defines parameters for GetBlockByBurnBlockHash
type GetBlockByBurnBlockHashParams struct{}

// GetContractByIDParams defines parameters for GetContractByID
type GetContractByIDParams struct{}

// FetchFeeRateParams defines parameters for FetchFeeRate
type FetchFeeRateParams struct {
	// A serialized transaction
	Transaction string
}

// GetContractSourceParams defines parameters for GetContractSource
type GetContractSourceParams struct{}

// GetAccountAssetsParams defines parameters for GetAccountAssets
type GetAccountAssetsParams struct {
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
	Principal string
	// max number of account assets to fetch
	Limit int
	// index of first account assets to fetch
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	Unanchored bool
	// returned data representing the state at that point in time, rather than the current block.
	UntilBlock string
}

// GetAccountNoncesParams defines parameters for GetAccountNonces
type GetAccountNoncesParams struct {
	// Stacks address (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0`)
	Principal string
	// Optionally get the nonce at a given block height
	BlockHeight int
	// Optionally get the nonce at a given block hash
	BlockHash string
}

// GetBlockByBurnBlockHeightParams defines parameters for GetBlockByBurnBlockHeight
type GetBlockByBurnBlockHeightParams struct{}

// GetContractsByTraitParams defines parameters for GetContractsByTrait
type GetContractsByTraitParams struct {
	// JSON abi of the trait.
	TraitABI string
	// max number of contracts fetch
	Limit int
	// index of first contract event to fetch
	Offset int
}

// GetSTXSupplyTotalSupplyPlainParams defines parameters for GetSTXSupplyTotalSupplyPlain
type GetSTXSupplyTotalSupplyPlainParams struct{}

// GetRawTransactionByIDParams defines parameters for GetRawTransactionByID
type GetRawTransactionByIDParams struct{}

// GetAccountNFTParams defines parameters for GetAccountNFT
type GetAccountNFTParams struct {
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
	Principal string
	// number of items to return
	Limit int
	// number of items to skip
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	Unanchored bool
	// returned data representing the state up until that point in time, rather than the current block.
	UntilBlock string
}

// GetAccountTransactionsParams defines parameters for GetAccountTransactions
type GetAccountTransactionsParams struct {
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
	// returned data representing the state up until that point in time, rather than the current block.
	UntilBlock string
}

// GetSTXSupplyParams defines parameters for GetSTXSupply
type GetSTXSupplyParams struct {
	// The block height at which to query supply details from, if not provided then the latest block height is used
	Height int
}

// GetTransactionByIDParams defines parameters for GetTransactionByID
type GetTransactionByIDParams struct{}

// GetNamePriceParams defines parameters for GetNamePrice
type GetNamePriceParams struct {
	// the name to query price information for
	Name string
}

// GetNamespacePriceParams defines parameters for GetNamespacePrice
type GetNamespacePriceParams struct {
	// the namespace to fetch price for
	TLD string
}

// GetAddressMempoolTransactionsParams defines parameters for GetAddressMempoolTransactions
type GetAddressMempoolTransactionsParams struct {
	// Transactions for the address
	Address string
	// max number of transactions to fetch
	Limit int
	// index of first transaction to fetch
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	Unanchored bool
}

// GetBurnchainRewardListParams defines parameters for GetBurnchainRewardList
type GetBurnchainRewardListParams struct {
	// max number of rewards to fetch
	Limit int
	// index of first rewards to fetch
	Offset int
}

// GetSTXSupplyCirculatingPlainParams defines parameters for GetSTXSupplyCirculatingPlain
type GetSTXSupplyCirculatingPlainParams struct{}

// GetAccountInboundParams defines parameters for GetAccountInbound
type GetAccountInboundParams struct {
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
	// returned data representing the state up until that point in time, rather than the current block.
	UntilBlock string
}

// GetMicroblockListParams defines parameters for GetMicroblockList
type GetMicroblockListParams struct {
	// Max number of microblocks to fetch
	Limit int
	// Index of the first microblock to fetch
	Offset int
}

// GetFilteredEventsParams defines parameters for GetFilteredEvents
type GetFilteredEventsParams struct {
	// Hash of transaction
	TXID string
	// Stacks address or a Contract identifier
	Address string
	// number of items to return
	Limit int
	// number of items to skip
	Offset int
	// Filter the events on event type
	// [SmartContractLog STXLock STXAsset FungibleTokenAsset NonFungibleTokenAsset]
	Type Type
}

// GetNamespaceNamesParams defines parameters for GetNamespaceNames
type GetNamespaceNamesParams struct {
	// the namespace to fetch names from
	TLD string
	// names are returned in pages of size 100, so specify the page number.
	Page int
}

// PostFeeTransactionParams defines parameters for PostFeeTransaction
type PostFeeTransactionParams struct {
	EstimatedLen       int
	TransactionPayload string
}

// GetNetworkBlockTimeByNetworkParams defines parameters for GetNetworkBlockTimeByNetwork
type GetNetworkBlockTimeByNetworkParams struct {
	// Which network to retrieve the target block time of
	Network string
}

// GetFTMetadataListParams defines parameters for GetFTMetadataList
type GetFTMetadataListParams struct {
	// max number of tokens to fetch
	Limit int
	// index of first tokens to fetch
	Offset int
}

// GetNFTMetadataListParams defines parameters for GetNFTMetadataList
type GetNFTMetadataListParams struct {
	// max number of tokens to fetch
	Limit int
	// index of first tokens to fetch
	Offset int
}

// GetContractFTMetadataParams defines parameters for GetContractFTMetadata
type GetContractFTMetadataParams struct {
	// token's contract id
	ContractId string
}

// GetMempoolTransactionListParams defines parameters for GetMempoolTransactionList
type GetMempoolTransactionListParams struct {
	// Filter to only return transactions with this sender address.
	SenderAddress string
	// Filter to only return transactions with this recipient address (only applicable for STX transfer tx types).
	RecipientAddress string
	// Filter to only return transactions with this address as the sender or recipient (recipient only applicable for STX transfer tx types).
	Address string
	// max number of mempool transactions to fetch
	Limit int
	// index of first mempool transaction to fetch
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	Unanchored bool
}

// GetAllNamespacesParams defines parameters for GetAllNamespaces
type GetAllNamespacesParams struct{}

// GetBlockByHashParams defines parameters for GetBlockByHash
type GetBlockByHashParams struct{}

// GetBurnchainRewardSlotHoldersParams defines parameters for GetBurnchainRewardSlotHolders
type GetBurnchainRewardSlotHoldersParams struct {
	// max number of items to fetch
	Limit int
	// index of the first items to fetch
	Offset int
}

// GetMicroblockByHashParams defines parameters for GetMicroblockByHash
type GetMicroblockByHashParams struct{}

// GetTotalSTXSupplyLegacyFormatParams defines parameters for GetTotalSTXSupplyLegacyFormat
type GetTotalSTXSupplyLegacyFormatParams struct {
	// The block height at which to query supply details from, if not provided then the latest block height is used
	Height int
}

// GetNFTMintsParams defines parameters for GetNFTMints
type GetNFTMintsParams struct {
	// token asset class identifier
	AssetIdentifier string
	// max number of events to fetch
	Limit int
	// index of first event to fetch
	Offset int
	// whether or not to include events from unconfirmed transactions
	Unanchored bool
	// whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times.
	TXMetadata bool
}

// GetBurnchainRewardListByAddressParams defines parameters for GetBurnchainRewardListByAddress
type GetBurnchainRewardListByAddressParams struct {
	// Reward recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format
	Address string
	// max number of rewards to fetch
	Limit int
	// index of first rewards to fetch
	Offset int
}

// SearchByIDParams defines parameters for SearchByID
type SearchByIDParams struct {
	// The hex hash string for a block or transaction, account address, or contract address
	ID string
	// This includes the detailed data for purticular hash in the response
	IncludeMetadata bool
}

// GetNFTHistoryParams defines parameters for GetNFTHistory
type GetNFTHistoryParams struct {
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
	// whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times.
	TXMetadata bool
}

// GetTransactionsByBlockHeightParams defines parameters for GetTransactionsByBlockHeight
type GetTransactionsByBlockHeightParams struct {
	// Height of block
	Height int
	// max number of transactions to fetch
	Limit int
	// index of first transaction to fetch
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	Unanchored bool
}

// GetDroppedMempoolTransactionListParams defines parameters for GetDroppedMempoolTransactionList
type GetDroppedMempoolTransactionListParams struct {
	// max number of mempool transactions to fetch
	Limit int
	// index of first mempool transaction to fetch
	Offset int
}

// GetTXListDetailsParams defines parameters for GetTXListDetails
type GetTXListDetailsParams struct{}

// GetCoreAPIInfoParams defines parameters for GetCoreAPIInfo
type GetCoreAPIInfoParams struct{}

// RunFaucetBTCParams defines parameters for RunFaucetBTC
type RunFaucetBTCParams struct {
	// BTC testnet address
	Address string
}

// GetUnanchoredTxsParams defines parameters for GetUnanchoredTxs
type GetUnanchoredTxsParams struct{}

// GetNamesOwnedByAddressParams defines parameters for GetNamesOwnedByAddress
type GetNamesOwnedByAddressParams struct {
	// the layer-1 blockchain for the address
	Blockchain string
	// the address to lookup
	Address string
}

// GetHistoricalZoneFileParams defines parameters for GetHistoricalZoneFile
type GetHistoricalZoneFileParams struct {
	// fully-qualified name
	Name string
	// zone file hash
	ZoneFileHash string
}

// GetContractInterfaceParams defines parameters for GetContractInterface
type GetContractInterfaceParams struct{}

// GetFeeTransferParams defines parameters for GetFeeTransfer
type GetFeeTransferParams struct{}

// RunFaucetSTXParams defines parameters for RunFaucetSTX
type RunFaucetSTXParams struct {
	// STX testnet address
	Address string
	// Use required number of tokens for stacking
	Stacking bool
}

// GetStatusParams defines parameters for GetStatus
type GetStatusParams struct{}

// CallReadOnlyFunctionParams defines parameters for CallReadOnlyFunction
type CallReadOnlyFunctionParams struct {
	// Stacks address
	ContractAddress string
	// Contract name
	ContractName string
	// Function name
	FunctionName string
	// The Stacks chain tip to query from
	Tip string
	// An array of hex serialized Clarity values
	Arguments any
	// The simulated tx-sender
	Sender string
}

// GetBurnchainRewardSlotHoldersByAddressParams defines parameters for GetBurnchainRewardSlotHoldersByAddress
type GetBurnchainRewardSlotHoldersByAddressParams struct {
	// Reward slot holder recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format
	Address string
	// max number of items to fetch
	Limit int
	// index of the first items to fetch
	Offset int
}

// GetBurnchainRewardsTotalByAddressParams defines parameters for GetBurnchainRewardsTotalByAddress
type GetBurnchainRewardsTotalByAddressParams struct {
	// Reward recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format
	Address string
}

// GetContractEventsByIDParams defines parameters for GetContractEventsByID
type GetContractEventsByIDParams struct {
	// Contract identifier formatted as `<contract_address>.<contract_name>`
	ContractID string
	// max number of contract events to fetch
	Limit int
	// index of first contract event to fetch
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	Unanchored bool
}

// GetNetworkBlockTimesParams defines parameters for GetNetworkBlockTimes
type GetNetworkBlockTimesParams struct{}

// GetAllNamesParams defines parameters for GetAllNames
type GetAllNamesParams struct {
	// names are returned in pages of size 100, so specify the page number.
	Page int
}
