package api

// Defines a 500 (Faucet call failed) response for RunFaucetBTC. type RunFaucetBTC500Response struct{}
// Defines a 200 (Success) response for RunFaucetBTC. type RunFaucetBTC200Response struct{}
// Defines parameters for RunFaucetBTC
type RunFaucetBTCParams struct {
	// BTC testnet address.
	Address string
}

// Defines a 404 (Cannot find transaction for given ID) response for GetRawTransactionByID. type GetRawTransactionByID404Response struct{}
// Defines a 200 (Hex encoded serialized transaction) response for GetRawTransactionByID. type GetRawTransactionByID200Response struct{}
// Defines parameters for GetRawTransactionByID
type GetRawTransactionByIDParams struct{}

// Defines a 200 (Success) response for GetAccountNonces. type GetAccountNonces200Response struct{}
// Defines parameters for GetAccountNonces
type GetAccountNoncesParams struct {
	// Stacks address (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0`).
	Principal string
	// Optionally get the nonce at a given block height.
	BlockHeight int
	// Optionally get the nonce at a given block hash.
	// Optional. Pass an empty string to use the default value.
	BlockHash string
}

// Defines a 200 (success) response for GetSTXSupplyCirculatingPlain. type GetSTXSupplyCirculatingPlain200Response struct{}
// Defines parameters for GetSTXSupplyCirculatingPlain
type GetSTXSupplyCirculatingPlainParams struct{}

// Defines a 200 (List of events) response for GetContractEventsByID. type GetContractEventsByID200Response struct{}
// Defines parameters for GetContractEventsByID
type GetContractEventsByIDParams struct {
	// Contract identifier formatted as `<contract_address>.<contract_name>`.
	ContractID string
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Max number of contract events to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Index of first contract event to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
}

// Defines a 200 (Transactions) response for GetUnanchoredTxs. type GetUnanchoredTxs200Response struct{}
// Defines parameters for GetUnanchoredTxs
type GetUnanchoredTxsParams struct{}

// Defines a 404 (Cannot find microblock with given hash) response for GetMicroblockByHash. type GetMicroblockByHash404Response struct{}
// Defines a 200 (Microblock) response for GetMicroblockByHash. type GetMicroblockByHash200Response struct{}
// Defines parameters for GetMicroblockByHash
type GetMicroblockByHashParams struct{}

// Defines a 200 (List of fungible tokens metadata) response for GetFTMetadataList. type GetFTMetadataList200Response struct{}
// Defines parameters for GetFTMetadataList
type GetFTMetadataListParams struct {
	// Index of first tokens to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Max number of tokens to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
}

// Defines a 200 (List of transactions) response for GetTransactionList. type GetTransactionList200Response struct{}
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

// Defines a 404 (Could not find any transaction by ID) response for GetTXListDetails. type GetTXListDetails404Response struct{}
// Defines a 200 (Returns list of transactions with their details for corresponding requested tx_ids.) response for GetTXListDetails. type GetTXListDetails200Response struct{}
// Defines parameters for GetTXListDetails
type GetTXListDetailsParams struct{}

// Defines a 404 (Error) response for GetHistoricalZoneFile. type GetHistoricalZoneFile404Response struct{}
// Defines a 200 (Success) response for GetHistoricalZoneFile. type GetHistoricalZoneFile200Response struct{}
// Defines a 400 (Error) response for GetHistoricalZoneFile. type GetHistoricalZoneFile400Response struct{}
// Defines parameters for GetHistoricalZoneFile
type GetHistoricalZoneFileParams struct {
	// Zone file hash.
	ZoneFileHash string
	// Fully-qualified name.
	Name string
}

// Defines a 200 (Success) response for GetAccountBalance. type GetAccountBalance200Response struct{}
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

// Defines a 404 (Cannot find block with given height) response for GetBlockByBurnBlockHash. type GetBlockByBurnBlockHash404Response struct{}
// Defines a 200 (Block) response for GetBlockByBurnBlockHash. type GetBlockByBurnBlockHash200Response struct{}
// Defines parameters for GetBlockByBurnBlockHash
type GetBlockByBurnBlockHashParams struct{}

// Defines a 200 (Block) response for GetBlockByBurnBlockHeight. type GetBlockByBurnBlockHeight200Response struct{}
// Defines a 404 (Cannot find block with given height) response for GetBlockByBurnBlockHeight. type GetBlockByBurnBlockHeight404Response struct{}
// Defines parameters for GetBlockByBurnBlockHeight
type GetBlockByBurnBlockHeightParams struct{}

// Defines a 200 (Block) response for GetBlockByHash. type GetBlockByHash200Response struct{}
// Defines a 404 (Cannot find block with given ID) response for GetBlockByHash. type GetBlockByHash404Response struct{}
// Defines parameters for GetBlockByHash
type GetBlockByHashParams struct{}

// Defines a 200 (List of burnchain reward recipients and amounts) response for GetBurnchainRewardSlotHolders. type GetBurnchainRewardSlotHolders200Response struct{}
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

// Defines a 200 (List of burnchain reward recipients and amounts) response for GetBurnchainRewardList. type GetBurnchainRewardList200Response struct{}
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

// Defines a 200 (Success) response for GetTotalSTXSupplyLegacyFormat. type GetTotalSTXSupplyLegacyFormat200Response struct{}
// Defines parameters for GetTotalSTXSupplyLegacyFormat
type GetTotalSTXSupplyLegacyFormatParams struct {
	// The block height at which to query supply details from, if not provided then the latest block height is used.
	Height int
}

// Defines a 200 (Non fungible tokens metadata for contract id) response for GetContractNFTMetadata. type GetContractNFTMetadata200Response struct{}
// Defines parameters for GetContractNFTMetadata
type GetContractNFTMetadataParams struct {
	// Token's contract id.
	ContractId string
}

// Defines a 200 (Success) response for GetAccountInfo. type GetAccountInfo200Response struct{}
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

// Defines a 200 (Success) response for GetContractSource. type GetContractSource200Response struct{}
// Defines parameters for GetContractSource
type GetContractSourceParams struct{}

// Defines a 200 (List of burnchain reward recipients and amounts) response for GetBurnchainRewardsTotalByAddress. type GetBurnchainRewardsTotalByAddress200Response struct{}
// Defines parameters for GetBurnchainRewardsTotalByAddress
type GetBurnchainRewardsTotalByAddressParams struct {
	// Reward recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format.
	Address string
}

// Defines a 200 (Success) response for GetNetworkBlockTimeByNetwork. type GetNetworkBlockTimeByNetwork200Response struct{}
// Defines parameters for GetNetworkBlockTimeByNetwork
type GetNetworkBlockTimeByNetworkParams struct {
	// Which network to retrieve the target block time of.
	Network string
}

// Defines a 200 (List of Non-Fungible Token holdings) response for GetNFTHoldings. type GetNFTHoldings200Response struct{}
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

// Defines a 400 (Error) response for GetAllNames. type GetAllNames400Response struct{}
// Defines a 200 (Success) response for GetAllNames. type GetAllNames200Response struct{}
// Defines parameters for GetAllNames
type GetAllNamesParams struct {
	// Names are returned in pages of size 100, so specify the page number.
	Page int
}

// Defines a 200 (Success) response for GetAccountAssets. type GetAccountAssets200Response struct{}
// Defines parameters for GetAccountAssets
type GetAccountAssetsParams struct {
	// Index of first account assets to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
	// Max number of account assets to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Returned data representing the state at that point in time, rather than the current block.
	// Optional. Pass an empty string to use the default value.
	UntilBlock string
}

// Defines a 200 (Success) response for GetAccountInbound. type GetAccountInbound200Response struct{}
// Defines parameters for GetAccountInbound
type GetAccountInboundParams struct {
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
	// Returned data representing the state up until that point in time, rather than the current block.
	// Optional. Pass an empty string to use the default value.
	UntilBlock string
}

// Defines a 200 (Block) response for GetBlockByHeight. type GetBlockByHeight200Response struct{}
// Defines a 404 (Cannot find block with given height) response for GetBlockByHeight. type GetBlockByHeight404Response struct{}
// Defines parameters for GetBlockByHeight
type GetBlockByHeightParams struct{}

// Defines a 200 (List of contracts implement given trait) response for GetContractsByTrait. type GetContractsByTrait200Response struct{}
// Defines parameters for GetContractsByTrait
type GetContractsByTraitParams struct {
	// Max number of contracts fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// JSON abi of the trait.
	TraitABI string
	// Index of first contract event to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
}

// Defines a 200 (Success) response for GetNetworkBlockTimes. type GetNetworkBlockTimes200Response struct{}
// Defines parameters for GetNetworkBlockTimes
type GetNetworkBlockTimesParams struct{}

// Defines a 200 (List of microblocks) response for GetMicroblockList. type GetMicroblockList200Response struct{}
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

// Defines a 200 (Success) response for GetStatus. type GetStatus200Response struct{}
// Defines parameters for GetStatus
type GetStatusParams struct{}

// Defines a 200 (Success) response for CallReadOnlyFunction. type CallReadOnlyFunction200Response struct{}
// Defines parameters for CallReadOnlyFunction
type CallReadOnlyFunctionParams struct {
	// The Stacks chain tip to query from.
	// Optional. Pass an empty string to use the default value.
	Tip string
	// An array of hex serialized Clarity values.
	Arguments CallReadOnlyFunctionArguments
	// The simulated tx-sender.
	Sender string
	// Contract name.
	ContractName string
	// Function name.
	FunctionName string
}

// Defines a 200 (Success) response for GetAccountTransactionsWithTransfers. type GetAccountTransactionsWithTransfers200Response struct{}
// Defines parameters for GetAccountTransactionsWithTransfers
type GetAccountTransactionsWithTransfersParams struct {
	// Filter for transactions only at this given block height.
	Height int
	// Index of first account transaction to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
	// Max number of account transactions to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Returned data representing the state up until that point in time, rather than the current block.
	// Optional. Pass an empty string to use the default value.
	UntilBlock string
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
}

// Defines a 404 (Not found) response for SearchByID. type SearchByID404Response struct{}
// Defines a 200 (Success) response for SearchByID. type SearchByID200Response struct{}
// Defines parameters for SearchByID
type SearchByIDParams struct {
	// This includes the detailed data for purticular hash in the response.
	// Optional. Use `<nil>` as default.
	IncludeMetadata bool
	// The hex hash string for a block or transaction, account address, or contract address.
	ID string
}

// Defines a 200 (Success) response for GetFeeTransfer. type GetFeeTransfer200Response struct{}
// Defines parameters for GetFeeTransfer
type GetFeeTransferParams struct{}

// Defines a 200 (Success) response for GetAccountSTXBalance. type GetAccountSTXBalance200Response struct{}
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

// Defines a 200 (Success) response for GetAccountTransactions. type GetAccountTransactions200Response struct{}
// Defines parameters for GetAccountTransactions
type GetAccountTransactionsParams struct {
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
	// Filter for transactions only at this given block height.
	Height int
	// Index of first account transaction to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Max number of account transactions to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Returned data representing the state up until that point in time, rather than the current block.
	// Optional. Pass an empty string to use the default value.
	UntilBlock string
}

// Defines a 200 (List of Transactions) response for GetTransactionsByBlockHeight. type GetTransactionsByBlockHeight200Response struct{}
// Defines parameters for GetTransactionsByBlockHeight
type GetTransactionsByBlockHeightParams struct {
	// Index of first transaction to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Max number of transactions to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Height of block.
	Height int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
}

// Defines a 200 (Success) response for GetAllNamespaces. type GetAllNamespaces200Response struct{}
// Defines parameters for GetAllNamespaces
type GetAllNamespacesParams struct{}

// Defines a 404 (Error) response for GetNamespaceNames. type GetNamespaceNames404Response struct{}
// Defines a 200 (Success) response for GetNamespaceNames. type GetNamespaceNames200Response struct{}
// Defines a 400 (Error) response for GetNamespaceNames. type GetNamespaceNames400Response struct{}
// Defines parameters for GetNamespaceNames
type GetNamespaceNamesParams struct {
	// The namespace to fetch names from.
	TLD string
	// Names are returned in pages of size 100, so specify the page number.
	Page int
}

// Defines a 200 (Success) response for GetSingleTransactionWithTransfers. type GetSingleTransactionWithTransfers200Response struct{}
// Defines a 404 (Not found) response for GetSingleTransactionWithTransfers. type GetSingleTransactionWithTransfers404Response struct{}
// Defines parameters for GetSingleTransactionWithTransfers
type GetSingleTransactionWithTransfersParams struct {
	// Stacks address or a contract identifier.
	Principal string
	// Transaction id.
	TXID string
}

// Defines a 200 (Transaction fee rate) response for FetchFeeRate. type FetchFeeRate200Response struct{}
// Defines parameters for FetchFeeRate
type FetchFeeRateParams struct {
	// A serialized transaction.
	Transaction string
}

// Defines a 200 (Non-Fungible Token mints) response for GetNFTMints. type GetNFTMints200Response struct{}
// Defines parameters for GetNFTMints
type GetNFTMintsParams struct {
	// Whether or not to include events from unconfirmed transactions.
	// Optional. Use `false` as default.
	Unanchored bool
	// Whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times.
	// Optional. Use `false` as default.
	TXMetadata bool
	// Token asset class identifier.
	AssetIdentifier string
	// Max number of events to fetch.
	// Optional. Use `-1` to use the default value (50).
	Limit int
	// Index of first event to fetch.
	// Optional. Use `-1` to use the default value (0).
	Offset int
}

// Defines a 404 (Error) response for FetchZoneFile. type FetchZoneFile404Response struct{}
// Defines a 400 (Error) response for FetchZoneFile. type FetchZoneFile400Response struct{}
// Defines a 200 (Success) response for FetchZoneFile. type FetchZoneFile200Response struct{}
// Defines parameters for FetchZoneFile
type FetchZoneFileParams struct {
	// Fully-qualified name.
	Name string
}

// Defines a 200 (Non-Fungible Token event history) response for GetNFTHistory. type GetNFTHistory200Response struct{}
// Defines parameters for GetNFTHistory
type GetNFTHistoryParams struct {
	// Hex representation of the token's unique value.
	Value string
	// Token asset class identifier.
	AssetIdentifier string
	// Whether or not to include events from unconfirmed transactions.
	// Optional. Use `false` as default.
	Unanchored bool
	// Index of first event to fetch.
	// Optional. Use `-1` to use the default value (0).
	Offset int
	// Max number of events to fetch.
	// Optional. Use `-1` to use the default value (50).
	Limit int
}

// Defines a 200 (List of burnchain reward recipients and amounts) response for GetBurnchainRewardListByAddress. type GetBurnchainRewardListByAddress200Response struct{}
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

// Defines a 200 (List of dropped mempool transactions) response for GetDroppedMempoolTransactionList. type GetDroppedMempoolTransactionList200Response struct{}
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

// Defines a 200 (List of burnchain reward recipients and amounts) response for GetBurnchainRewardSlotHoldersByAddress. type GetBurnchainRewardSlotHoldersByAddress200Response struct{}
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

// Defines a 200 (List of Transactions) response for GetTransactionsByBlockHash. type GetTransactionsByBlockHash200Response struct{}
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

// Defines a 200 (Success) response for GetCoreAPIInfo. type GetCoreAPIInfo200Response struct{}
// Defines parameters for GetCoreAPIInfo
type GetCoreAPIInfoParams struct{}

// Defines a 200 (List of mempool transactions) response for GetMempoolTransactionList. type GetMempoolTransactionList200Response struct{}
// Defines parameters for GetMempoolTransactionList
type GetMempoolTransactionListParams struct {
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
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
}

// Defines a 200 (Success) response for GetNameInfo. type GetNameInfo200Response struct{}
// Defines a 400 (Error) response for GetNameInfo. type GetNameInfo400Response struct{}
// Defines a 404 (Error) response for GetNameInfo. type GetNameInfo404Response struct{}
// Defines parameters for GetNameInfo
type GetNameInfoParams struct {
	// Fully-qualified name.
	Name string
}

// Defines a 200 (Contract interface) response for GetContractInterface. type GetContractInterface200Response struct{}
// Defines parameters for GetContractInterface
type GetContractInterfaceParams struct{}

// Defines a 200 (List of blocks) response for GetBlockList. type GetBlockList200Response struct{}
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

// Defines a 404 (Cannot find contract of given ID) response for GetContractByID. type GetContractByID404Response struct{}
// Defines a 200 (Contract found) response for GetContractByID. type GetContractByID200Response struct{}
// Defines parameters for GetContractByID
type GetContractByIDParams struct{}

// Defines a 200 (Fungible tokens metadata for contract id) response for GetContractFTMetadata. type GetContractFTMetadata200Response struct{}
// Defines parameters for GetContractFTMetadata
type GetContractFTMetadataParams struct {
	// Token's contract id.
	ContractId string
}

// Defines a 400 (Failed loading data map) response for GetContractDataMapEntry. type GetContractDataMapEntry400Response struct{}
// Defines a 200 (Success) response for GetContractDataMapEntry. type GetContractDataMapEntry200Response struct{}
// Defines parameters for GetContractDataMapEntry
type GetContractDataMapEntryParams struct {
	// Stacks address.
	ContractAddress string
	// The Stacks chain tip to query from.
	// Optional. Pass an empty string to use the default value.
	Tip string
	// Returns object without the proof field when set to 0.
	// Optional. Use `-1` to use the default value.
	Proof int
	// Contract name.
	ContractName string
	// Map name.
	MapName string
}

// Defines a 200 (Success) response for GetPoxInfo. type GetPoxInfo200Response struct{}
// Defines parameters for GetPoxInfo
type GetPoxInfoParams struct{}

// Defines a 200 (Success) response for GetNamePrice. type GetNamePrice200Response struct{}
// Defines parameters for GetNamePrice
type GetNamePriceParams struct {
	// The name to query price information for.
	Name string
}

// Defines a 200 (Success) response for GetNamespacePrice. type GetNamespacePrice200Response struct{}
// Defines parameters for GetNamespacePrice
type GetNamespacePriceParams struct {
	// The namespace to fetch price for.
	TLD string
}

// Defines a 200 (List of Transactions) response for GetAddressMempoolTransactions. type GetAddressMempoolTransactions200Response struct{}
// Defines parameters for GetAddressMempoolTransactions
type GetAddressMempoolTransactionsParams struct {
	// Max number of transactions to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Transactions for the address.
	Address string
	// Index of first transaction to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
}

// Defines a 200 (Success) response for GetAccountNFT. type GetAccountNFT200Response struct{}
// Defines parameters for GetAccountNFT
type GetAccountNFTParams struct {
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Returned data representing the state up until that point in time, rather than the current block.
	// Optional. Pass an empty string to use the default value.
	UntilBlock string
	// Number of items to skip.
	// Optional. Use `-1` to use the default value.
	Offset int
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
	// Number of items to return.
	// Optional. Use `-1` to use the default value.
	Limit int
}

// Defines a 500 (Faucet call failed) response for RunFaucetSTX. type RunFaucetSTX500Response struct{}
// Defines a 200 (Success) response for RunFaucetSTX. type RunFaucetSTX200Response struct{}
// Defines parameters for RunFaucetSTX
type RunFaucetSTXParams struct {
	// STX testnet address.
	Address string
	// Use required number of tokens for stacking.
	Stacking bool
}

// Defines a 200 (Success) response for GetSTXSupply. type GetSTXSupply200Response struct{}
// Defines parameters for GetSTXSupply
type GetSTXSupplyParams struct {
	// The block height at which to query supply details from, if not provided then the latest block height is used.
	Height int
}

// Defines a 200 (List of non fungible tokens metadata) response for GetNFTMetadataList. type GetNFTMetadataList200Response struct{}
// Defines parameters for GetNFTMetadataList
type GetNFTMetadataListParams struct {
	// Max number of tokens to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Index of first tokens to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
}

// Defines a 404 (Cannot find transaction for given ID) response for GetTransactionByID. type GetTransactionByID404Response struct{}
// Defines a 200 (Transaction) response for GetTransactionByID. type GetTransactionByID200Response struct{}
// Defines parameters for GetTransactionByID
type GetTransactionByIDParams struct{}

// Defines a 200 (Estimated fees for the transaction) response for PostFeeTransaction. type PostFeeTransaction200Response struct{}
// Defines parameters for PostFeeTransaction
type PostFeeTransactionParams struct {
	//
	EstimatedLen int
	//
	TransactionPayload string
}

// Defines a 400 (Rejections result in a 400 error) response for PostCoreNodeTransactions. type PostCoreNodeTransactions400Response struct{}
// Defines a 200 (Transaction id of successful post of a raw tx to the node's mempool) response for PostCoreNodeTransactions. type PostCoreNodeTransactions200Response struct{}
// Defines parameters for PostCoreNodeTransactions
type PostCoreNodeTransactionsParams struct{}

// Defines a 200 (success) response for GetSTXSupplyTotalSupplyPlain. type GetSTXSupplyTotalSupplyPlain200Response struct{}
// Defines parameters for GetSTXSupplyTotalSupplyPlain
type GetSTXSupplyTotalSupplyPlainParams struct{}

// Defines a 200 (Success) response for GetFilteredEvents. type GetFilteredEvents200Response struct{}
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
	// Stacks address or a Contract identifier.
	// Optional. Pass an empty string to use the default value.
	Address string
	// Number of items to return.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Filter the events on event type.
	// [SmartContractLog STXLock STXAsset FungibleTokenAsset NonFungibleTokenAsset]
	Type GetFilteredEventsType
	// Number of items to skip.
	// Optional. Use `-1` to use the default value.
	Offset int
}

// Defines a 404 (Error) response for GetNamesOwnedByAddress. type GetNamesOwnedByAddress404Response struct{}
// Defines a 200 (Success) response for GetNamesOwnedByAddress. type GetNamesOwnedByAddress200Response struct{}
// Defines parameters for GetNamesOwnedByAddress
type GetNamesOwnedByAddressParams struct {
	// The address to lookup.
	Address string
	// The layer-1 blockchain for the address.
	Blockchain string
}

// Defines a 200 (Success) response for FetchSubdomainsListForName. type FetchSubdomainsListForName200Response struct{}
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
