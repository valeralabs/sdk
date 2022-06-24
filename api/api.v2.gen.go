package api

// GetSTXSupplyCirculatingPlainParams defines parameters for GetSTXSupplyCirculatingPlain
type GetSTXSupplyCirculatingPlainParams struct{}

// GetTransactionsByBlockHeightParams defines parameters for GetTransactionsByBlockHeight
type GetTransactionsByBlockHeightParams struct {
	// Height of block
	height int
	// max number of transactions to fetch
	limit int
	// index of first transaction to fetch
	offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	unanchored bool
}

// FetchZoneFileParams defines parameters for FetchZoneFile
type FetchZoneFileParams struct {
	// fully-qualified name
	name string
}

// GetHistoricalZoneFileParams defines parameters for GetHistoricalZoneFile
type GetHistoricalZoneFileParams struct {
	// fully-qualified name
	name string
	// zone file hash
	zoneFileHash string
}

// GetAccountNoncesParams defines parameters for GetAccountNonces
type GetAccountNoncesParams struct {
	// Stacks address (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0`)
	principal string
	// Optionally get the nonce at a given block height
	block_height int
	// Optionally get the nonce at a given block hash
	block_hash string
}

// GetAccountTransactionsParams defines parameters for GetAccountTransactions
type GetAccountTransactionsParams struct {
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
	principal string
	// max number of account transactions to fetch
	limit int
	// index of first account transaction to fetch
	offset int
	// Filter for transactions only at this given block height
	height int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	unanchored bool
	// returned data representing the state up until that point in time, rather than the current block.
	until_block string
}

// GetSTXSupplyParams defines parameters for GetSTXSupply
type GetSTXSupplyParams struct {
	// The block height at which to query supply details from, if not provided then the latest block height is used
	height int
}

// GetNamePriceParams defines parameters for GetNamePrice
type GetNamePriceParams struct {
	// the name to query price information for
	name string
}

// GetAllNamespacesParams defines parameters for GetAllNamespaces
type GetAllNamespacesParams struct{}

// CallReadOnlyFunctionParams defines parameters for CallReadOnlyFunction
type CallReadOnlyFunctionParams struct {
	// Stacks address
	contract_address string
	// Contract name
	contract_name string
	// Function name
	function_name string
	// The Stacks chain tip to query from
	tip string
}

// PostCoreNodeTransactionsParams defines parameters for PostCoreNodeTransactions
type PostCoreNodeTransactionsParams struct{}

// GetRawTransactionByIDParams defines parameters for GetRawTransactionByID
type GetRawTransactionByIDParams struct{}

// GetAccountSTXBalanceParams defines parameters for GetAccountSTXBalance
type GetAccountSTXBalanceParams struct {
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
	principal string
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	unanchored bool
	// returned data representing the state up until that point in time, rather than the current block.
	until_block string
}

// RosettaConstructionCombineParams defines parameters for RosettaConstructionCombine
type RosettaConstructionCombineParams struct{}

// RosettaConstructionPayloadsParams defines parameters for RosettaConstructionPayloads
type RosettaConstructionPayloadsParams struct{}

// RosettaMempoolParams defines parameters for RosettaMempool
type RosettaMempoolParams struct{}

// GetFilteredEventsParams defines parameters for GetFilteredEvents
type GetFilteredEventsParams struct {
	// Hash of transaction
	TX_ID string
	// Stacks address or a Contract identifier
	address string
	// number of items to return
	limit int
	// number of items to skip
	offset int
	// Filter the events on event type
	type_ interface{}
}

// GetCoreAPIInfoParams defines parameters for GetCoreAPIInfo
type GetCoreAPIInfoParams struct{}

// GetPoxInfoParams defines parameters for GetPoxInfo
type GetPoxInfoParams struct{}

// GetContractEventsByIDParams defines parameters for GetContractEventsByID
type GetContractEventsByIDParams struct {
	// Contract identifier formatted as `<contract_address>.<contract_name>`
	contract_ID string
	// max number of contract events to fetch
	limit int
	// index of first contract event to fetch
	offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	unanchored bool
}

// GetBlockByBurnBlockHeightParams defines parameters for GetBlockByBurnBlockHeight
type GetBlockByBurnBlockHeightParams struct{}

// GetUnanchoredTXsParams defines parameters for GetUnanchoredTXs
type GetUnanchoredTXsParams struct{}

// GetNFTHoldingsParams defines parameters for GetNFTHoldings
type GetNFTHoldingsParams struct {
	// token owner's STX address or Smart Contract ID
	principal string
	// identifiers of the token asset classes to filter for
	asset_IDentifiers interface{}
	// max number of tokens to fetch
	limit int
	// index of first tokens to fetch
	offset int
	// whether or not to include tokens from unconfirmed transactions
	unanchored bool
	// whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times.
	TX_metadata bool
}

// GetNamespaceNamesParams defines parameters for GetNamespaceNames
type GetNamespaceNamesParams struct {
	// the namespace to fetch names from
	TLD string
	// names are returned in pages of size 100, so specify the page number.
	page int
}

// GetContractByIDParams defines parameters for GetContractByID
type GetContractByIDParams struct{}

// GetBlockListParams defines parameters for GetBlockList
type GetBlockListParams struct {
	// max number of blocks to fetch
	limit int
	// index of first block to fetch
	offset int
}

// GetBlockByHeightParams defines parameters for GetBlockByHeight
type GetBlockByHeightParams struct{}

// GetMempoolTransactionListParams defines parameters for GetMempoolTransactionList
type GetMempoolTransactionListParams struct {
	// Filter to only return transactions with this sender address.
	sender_address string
	// Filter to only return transactions with this recipient address (only applicable for STX transfer tx types).
	recipient_address string
	// Filter to only return transactions with this address as the sender or recipient (recipient only applicable for STX transfer tx types).
	address string
	// max number of mempool transactions to fetch
	limit int
	// index of first mempool transaction to fetch
	offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	unanchored bool
}

// RosettaConstructionSubmitParams defines parameters for RosettaConstructionSubmit
type RosettaConstructionSubmitParams struct{}

// RosettaNetworkStatusParams defines parameters for RosettaNetworkStatus
type RosettaNetworkStatusParams struct{}

// GetNamesOwnedByAddressParams defines parameters for GetNamesOwnedByAddress
type GetNamesOwnedByAddressParams struct {
	// the layer-1 blockchain for the address
	blockchain string
	// the address to lookup
	address string
}

// GetContractDataMapEntryParams defines parameters for GetContractDataMapEntry
type GetContractDataMapEntryParams struct {
	// Stacks address
	contract_address string
	// Contract name
	contract_name string
	// Map name
	map_name string
	// Returns object without the proof field when set to 0
	proof int
	// The Stacks chain tip to query from
	tip string
}

// SearchByIDParams defines parameters for SearchByID
type SearchByIDParams struct {
	// The hex hash string for a block or transaction, account address, or contract address
	ID string
	// This includes the detailed data for purticular hash in the response
	include_metadata bool
}

// GetBurnchainRewardSlotHoldersByAddressParams defines parameters for GetBurnchainRewardSlotHoldersByAddress
type GetBurnchainRewardSlotHoldersByAddressParams struct {
	// Reward slot holder recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format
	address string
	// max number of items to fetch
	limit int
	// index of the first items to fetch
	offset int
}

// RunFaucetSTXParams defines parameters for RunFaucetSTX
type RunFaucetSTXParams struct{}

// RosettaAccountBalanceParams defines parameters for RosettaAccountBalance
type RosettaAccountBalanceParams struct{}

// RosettaConstructionPreprocessParams defines parameters for RosettaConstructionPreprocess
type RosettaConstructionPreprocessParams struct{}

// RosettaBlockParams defines parameters for RosettaBlock
type RosettaBlockParams struct{}

// RosettaBlockTransactionParams defines parameters for RosettaBlockTransaction
type RosettaBlockTransactionParams struct{}

// RosettaConstructionHashParams defines parameters for RosettaConstructionHash
type RosettaConstructionHashParams struct{}

// GetAllNamesParams defines parameters for GetAllNames
type GetAllNamesParams struct {
	// names are returned in pages of size 100, so specify the page number.
	page int
}

// GetAddressMempoolTransactionsParams defines parameters for GetAddressMempoolTransactions
type GetAddressMempoolTransactionsParams struct {
	// Transactions for the address
	address string
	// max number of transactions to fetch
	limit int
	// index of first transaction to fetch
	offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	unanchored bool
}

// GetAccountNFTParams defines parameters for GetAccountNFT
type GetAccountNFTParams struct {
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
	principal string
	// number of items to return
	limit int
	// number of items to skip
	offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	unanchored bool
	// returned data representing the state up until that point in time, rather than the current block.
	until_block string
}

// GetNFTHistoryParams defines parameters for GetNFTHistory
type GetNFTHistoryParams struct {
	// token asset class identifier
	asset_IDentifier string
	// hex representation of the token's unique value
	value string
	// max number of events to fetch
	limit int
	// index of first event to fetch
	offset int
	// whether or not to include events from unconfirmed transactions
	unanchored bool
	// whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times.
	TX_metadata bool
}

// GetTransactionsByBlockHashParams defines parameters for GetTransactionsByBlockHash
type GetTransactionsByBlockHashParams struct {
	// Hash of block
	block_hash string
	// max number of transactions to fetch
	limit int
	// index of first transaction to fetch
	offset int
}

// GetContractSourceParams defines parameters for GetContractSource
type GetContractSourceParams struct{}

// PostFeeTransactionParams defines parameters for PostFeeTransaction
type PostFeeTransactionParams struct{}

// GetFeeTransferParams defines parameters for GetFeeTransfer
type GetFeeTransferParams struct{}

// FetchFeeRateParams defines parameters for FetchFeeRate
type FetchFeeRateParams struct{}

// RosettaNetworkOptionsParams defines parameters for RosettaNetworkOptions
type RosettaNetworkOptionsParams struct{}

// GetAccountInfoParams defines parameters for GetAccountInfo
type GetAccountInfoParams struct {
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
	principal string
	// Returns object without the proof field if set to 0
	proof int
	// The Stacks chain tip to query from
	tip string
}

// GetTransactionByIDParams defines parameters for GetTransactionByID
type GetTransactionByIDParams struct{}

// RosettaConstructionDeriveParams defines parameters for RosettaConstructionDerive
type RosettaConstructionDeriveParams struct{}

// GetContractInterfaceParams defines parameters for GetContractInterface
type GetContractInterfaceParams struct{}

// GetAccountInboundParams defines parameters for GetAccountInbound
type GetAccountInboundParams struct {
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
	principal string
	// number of items to return
	limit int
	// number of items to skip
	offset int
	// Filter for transfers only at this given block height
	height int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	unanchored bool
	// returned data representing the state up until that point in time, rather than the current block.
	until_block string
}

// GetBurnchainRewardsTotalByAddressParams defines parameters for GetBurnchainRewardsTotalByAddress
type GetBurnchainRewardsTotalByAddressParams struct {
	// Reward recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format
	address string
}

// GetFTMetadataListParams defines parameters for GetFTMetadataList
type GetFTMetadataListParams struct {
	// max number of tokens to fetch
	limit int
	// index of first tokens to fetch
	offset int
}

// GetNFTMetadataListParams defines parameters for GetNFTMetadataList
type GetNFTMetadataListParams struct {
	// max number of tokens to fetch
	limit int
	// index of first tokens to fetch
	offset int
}

// GetStatusParams defines parameters for GetStatus
type GetStatusParams struct{}

// RosettaNetworkListParams defines parameters for RosettaNetworkList
type RosettaNetworkListParams struct{}

// FetchSubdomainsListForNameParams defines parameters for FetchSubdomainsListForName
type FetchSubdomainsListForNameParams struct {
	// fully-qualified name
	name string
}

// GetAccountTransactionsWithTransfersParams defines parameters for GetAccountTransactionsWithTransfers
type GetAccountTransactionsWithTransfersParams struct {
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
	principal string
	// max number of account transactions to fetch
	limit int
	// index of first account transaction to fetch
	offset int
	// Filter for transactions only at this given block height
	height int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	unanchored bool
	// returned data representing the state up until that point in time, rather than the current block.
	until_block string
}

// RunFaucetBTCParams defines parameters for RunFaucetBTC
type RunFaucetBTCParams struct{}

// GetNamespacePriceParams defines parameters for GetNamespacePrice
type GetNamespacePriceParams struct {
	// the namespace to fetch price for
	TLD string
}

// GetBurnchainRewardListByAddressParams defines parameters for GetBurnchainRewardListByAddress
type GetBurnchainRewardListByAddressParams struct {
	// Reward recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format
	address string
	// max number of rewards to fetch
	limit int
	// index of first rewards to fetch
	offset int
}

// GetMicroblockByHashParams defines parameters for GetMicroblockByHash
type GetMicroblockByHashParams struct{}

// RosettaConstructionParseParams defines parameters for RosettaConstructionParse
type RosettaConstructionParseParams struct{}

// RosettaMempoolTransactionParams defines parameters for RosettaMempoolTransaction
type RosettaMempoolTransactionParams struct{}

// GetContractsByTraitParams defines parameters for GetContractsByTrait
type GetContractsByTraitParams struct {
	// JSON abi of the trait.
	trait_abi string
	// max number of contracts fetch
	limit int
	// index of first contract event to fetch
	offset int
}

// GetTotalSTXSupplyLegacyFormatParams defines parameters for GetTotalSTXSupplyLegacyFormat
type GetTotalSTXSupplyLegacyFormatParams struct {
	// The block height at which to query supply details from, if not provided then the latest block height is used
	height int
}

// GetSTXSupplyTotalSupplyPlainParams defines parameters for GetSTXSupplyTotalSupplyPlain
type GetSTXSupplyTotalSupplyPlainParams struct{}

// GetContractNFTMetadataParams defines parameters for GetContractNFTMetadata
type GetContractNFTMetadataParams struct {
	// token's contract id
	contractId string
}

// GetAccountBalanceParams defines parameters for GetAccountBalance
type GetAccountBalanceParams struct {
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
	principal string
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	unanchored bool
	// returned data representing the state up until that point in time, rather than the current block.
	until_block string
}

// GetBlockByBurnBlockHashParams defines parameters for GetBlockByBurnBlockHash
type GetBlockByBurnBlockHashParams struct{}

// GetBurnchainRewardSlotHoldersParams defines parameters for GetBurnchainRewardSlotHolders
type GetBurnchainRewardSlotHoldersParams struct {
	// max number of items to fetch
	limit int
	// index of the first items to fetch
	offset int
}

// GetBurnchainRewardListParams defines parameters for GetBurnchainRewardList
type GetBurnchainRewardListParams struct {
	// max number of rewards to fetch
	limit int
	// index of first rewards to fetch
	offset int
}

// GetDroppedMempoolTransactionListParams defines parameters for GetDroppedMempoolTransactionList
type GetDroppedMempoolTransactionListParams struct {
	// max number of mempool transactions to fetch
	limit int
	// index of first mempool transaction to fetch
	offset int
}

// GetTXListDetailsParams defines parameters for GetTXListDetails
type GetTXListDetailsParams struct{}

// GetSingleTransactionWithTransfersParams defines parameters for GetSingleTransactionWithTransfers
type GetSingleTransactionWithTransfersParams struct {
	// Stacks address or a contract identifier
	principal string
	// Transaction id
	TX_ID string
}

// GetNFTMintsParams defines parameters for GetNFTMints
type GetNFTMintsParams struct {
	// token asset class identifier
	asset_IDentifier string
	// max number of events to fetch
	limit int
	// index of first event to fetch
	offset int
	// whether or not to include events from unconfirmed transactions
	unanchored bool
	// whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times.
	TX_metadata bool
}

// GetContractFTMetadataParams defines parameters for GetContractFTMetadata
type GetContractFTMetadataParams struct {
	// token's contract id
	contractId string
}

// GetTransactionListParams defines parameters for GetTransactionList
type GetTransactionListParams struct {
	// max number of transactions to fetch
	limit int
	// index of first transaction to fetch
	offset int
	// Filter by transaction type
	type_ interface{}
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	unanchored bool
}

// GetAccountAssetsParams defines parameters for GetAccountAssets
type GetAccountAssetsParams struct {
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
	principal string
	// max number of account assets to fetch
	limit int
	// index of first account assets to fetch
	offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	unanchored bool
	// returned data representing the state at that point in time, rather than the current block.
	until_block string
}

// GetNetworkBlockTimeByNetworkParams defines parameters for GetNetworkBlockTimeByNetwork
type GetNetworkBlockTimeByNetworkParams struct {
	// Which network to retrieve the target block time of
	network string
}

// GetMicroblockListParams defines parameters for GetMicroblockList
type GetMicroblockListParams struct {
	// Max number of microblocks to fetch
	limit int
	// Index of the first microblock to fetch
	offset int
}

// RosettaConstructionMetadataParams defines parameters for RosettaConstructionMetadata
type RosettaConstructionMetadataParams struct{}

// GetBlockByHashParams defines parameters for GetBlockByHash
type GetBlockByHashParams struct{}

// GetNetworkBlockTimesParams defines parameters for GetNetworkBlockTimes
type GetNetworkBlockTimesParams struct{}

// GetNameInfoParams defines parameters for GetNameInfo
type GetNameInfoParams struct {
	// fully-qualified name
	name string
}
