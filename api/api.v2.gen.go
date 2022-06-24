package api

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

// GetContractFTMetadataParams defines parameters for GetContractFTMetadata
type GetContractFTMetadataParams struct {
	// token's contract id
	ContractId string
}

// RosettaNetworkOptionsParams defines parameters for RosettaNetworkOptions
type RosettaNetworkOptionsParams struct {
	Metadata interface{}
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier interface{}
}

// FetchSubdomainsListForNameParams defines parameters for FetchSubdomainsListForName
type FetchSubdomainsListForNameParams struct {
	// fully-qualified name
	Name string
}

// RosettaConstructionDeriveParams defines parameters for RosettaConstructionDerive
type RosettaConstructionDeriveParams struct {
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier interface{}
	// PublicKey contains a public key byte array for a particular CurveType encoded in hex. Note that there is no PrivateKey struct as this is NEVER the concern of an implementation.
	PublicKey interface{}
	Metadata  interface{}
}

// RosettaConstructionParseParams defines parameters for RosettaConstructionParse
type RosettaConstructionParseParams struct {
	// This must be either the unsigned transaction blob returned by /construction/payloads or the signed transaction blob returned by /construction/combine.
	Transaction string
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier interface{}
	// Signed is a boolean indicating whether the transaction is signed.
	Signed bool
}

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

// GetNFTMetadataListParams defines parameters for GetNFTMetadataList
type GetNFTMetadataListParams struct {
	// max number of tokens to fetch
	Limit int
	// index of first tokens to fetch
	Offset int
}

// GetContractNFTMetadataParams defines parameters for GetContractNFTMetadata
type GetContractNFTMetadataParams struct {
	// token's contract id
	ContractId string
}

// RosettaAccountBalanceParams defines parameters for RosettaAccountBalance
type RosettaAccountBalanceParams struct {
	// When fetching data by BlockIdentifier, it may be possible to only specify the index or hash. If neither property is specified, it is assumed that the client is making a request at the current block.
	BlockIdentifier interface{}
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier interface{}
	// The account_identifier uniquely identifies an account within a network. All fields in the account_identifier are utilized to determine this uniqueness (including the metadata field, if populated).
	AccountIdentifier interface{}
}

// RosettaConstructionPayloadsParams defines parameters for RosettaConstructionPayloads
type RosettaConstructionPayloadsParams struct {
	PublicKeys interface{}
	Metadata   interface{}
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier interface{}
	Operations        interface{}
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

// GetSingleTransactionWithTransfersParams defines parameters for GetSingleTransactionWithTransfers
type GetSingleTransactionWithTransfersParams struct {
	// Stacks address or a contract identifier
	Principal string
	// Transaction id
	TXID string
}

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

// GetAllNamesParams defines parameters for GetAllNames
type GetAllNamesParams struct {
	// names are returned in pages of size 100, so specify the page number.
	Page int
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

// RosettaConstructionCombineParams defines parameters for RosettaConstructionCombine
type RosettaConstructionCombineParams struct {
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier   interface{}
	Signatures          interface{}
	UnsignedTransaction string
}

// RosettaConstructionPreprocessParams defines parameters for RosettaConstructionPreprocess
type RosettaConstructionPreprocessParams struct {
	MaxFee   interface{}
	Metadata interface{}
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier interface{}
	Operations        interface{}
	//  The caller can also provide a suggested fee multiplier to indicate that the suggested fee should be scaled. This may be used to set higher fees for urgent transactions or to pay lower fees when there is less urgency. It is assumed that providing a very low multiplier (like 0.0001) will never lead to a transaction being created with a fee less than the minimum network fee (if applicable). In the case that the caller provides both a max fee and a suggested fee multiplier, the max fee will set an upper bound on the suggested fee (regardless of the multiplier provided).
	SuggestedFeeMultiplier int
}

// GetBlockByBurnBlockHashParams defines parameters for GetBlockByBurnBlockHash
type GetBlockByBurnBlockHashParams struct{}

// GetBlockByBurnBlockHeightParams defines parameters for GetBlockByBurnBlockHeight
type GetBlockByBurnBlockHeightParams struct{}

// RunFaucetSTXParams defines parameters for RunFaucetSTX
type RunFaucetSTXParams struct {
	// STX testnet address
	Address string
	// Use required number of tokens for stacking
	Stacking bool
}

// GetMicroblockListParams defines parameters for GetMicroblockList
type GetMicroblockListParams struct {
	// Max number of microblocks to fetch
	Limit int
	// Index of the first microblock to fetch
	Offset int
}

// GetTXListDetailsParams defines parameters for GetTXListDetails
type GetTXListDetailsParams struct{}

// RosettaBlockTransactionParams defines parameters for RosettaBlockTransaction
type RosettaBlockTransactionParams struct {
	// When fetching data by BlockIdentifier, it may be possible to only specify the index or hash. If neither property is specified, it is assumed that the client is making a request at the current block.
	BlockIdentifier interface{}
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier interface{}
	// The transaction_identifier uniquely identifies a transaction in a particular network and block or in the mempool.
	TransactionIdentifier interface{}
}

// FetchZoneFileParams defines parameters for FetchZoneFile
type FetchZoneFileParams struct {
	// fully-qualified name
	Name string
}

// GetFeeTransferParams defines parameters for GetFeeTransfer
type GetFeeTransferParams struct{}

// GetCoreAPIInfoParams defines parameters for GetCoreAPIInfo
type GetCoreAPIInfoParams struct{}

// GetBurnchainRewardSlotHoldersByAddressParams defines parameters for GetBurnchainRewardSlotHoldersByAddress
type GetBurnchainRewardSlotHoldersByAddressParams struct {
	// Reward slot holder recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format
	Address string
	// max number of items to fetch
	Limit int
	// index of the first items to fetch
	Offset int
}

// GetNFTMintsParams defines parameters for GetNFTMints
type GetNFTMintsParams struct {
	// token asset class identifier
	AssetIDentifier string
	// max number of events to fetch
	Limit int
	// index of first event to fetch
	Offset int
	// whether or not to include events from unconfirmed transactions
	Unanchored bool
	// whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times.
	TXMetadata bool
}

// GetTransactionsByBlockHashParams defines parameters for GetTransactionsByBlockHash
type GetTransactionsByBlockHashParams struct {
	// Hash of block
	BlockHash string
	// max number of transactions to fetch
	Limit int
	// index of first transaction to fetch
	Offset int
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
type GetFilteredEventsType int64

const (
	SmartContractLog GetFilteredEventsType = iota
	StxLock
	StxAsset
	FungibleTokenAsset
	NonFungibleTokenAsset
)

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
	// [SmartContractLog StxLock StxAsset FungibleTokenAsset NonFungibleTokenAsset]
	Type GetFilteredEventsType
}

// RosettaMempoolParams defines parameters for RosettaMempool
type RosettaMempoolParams struct {
	Metadata interface{}
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier interface{}
}

// PostFeeTransactionParams defines parameters for PostFeeTransaction
type PostFeeTransactionParams struct {
	EstimatedLen       int
	TransactionPayload string
}

// GetTransactionByIDParams defines parameters for GetTransactionByID
type GetTransactionByIDParams struct{}

// GetAccountNoncesParams defines parameters for GetAccountNonces
type GetAccountNoncesParams struct {
	// Stacks address (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0`)
	Principal string
	// Optionally get the nonce at a given block height
	BlockHeight int
	// Optionally get the nonce at a given block hash
	BlockHash string
}

// GetSTXSupplyParams defines parameters for GetSTXSupply
type GetSTXSupplyParams struct {
	// The block height at which to query supply details from, if not provided then the latest block height is used
	Height int
}

// GetNamesOwnedByAddressParams defines parameters for GetNamesOwnedByAddress
type GetNamesOwnedByAddressParams struct {
	// the layer-1 blockchain for the address
	Blockchain string
	// the address to lookup
	Address string
}

// GetNameInfoParams defines parameters for GetNameInfo
type GetNameInfoParams struct {
	// fully-qualified name
	Name string
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
	Arguments interface{}
	// The simulated tx-sender
	Sender string
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

// GetBurnchainRewardListByAddressParams defines parameters for GetBurnchainRewardListByAddress
type GetBurnchainRewardListByAddressParams struct {
	// Reward recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format
	Address string
	// max number of rewards to fetch
	Limit int
	// index of first rewards to fetch
	Offset int
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

// RosettaConstructionMetadataParams defines parameters for RosettaConstructionMetadata
type RosettaConstructionMetadataParams struct {
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier interface{}
	// The options that will be sent directly to /construction/metadata by the caller.
	Options    interface{}
	PublicKeys interface{}
}

// PostCoreNodeTransactionsParams defines parameters for PostCoreNodeTransactions
type PostCoreNodeTransactionsParams struct{}

// GetContractsByTraitParams defines parameters for GetContractsByTrait
type GetContractsByTraitParams struct {
	// JSON abi of the trait.
	TraitAbi string
	// max number of contracts fetch
	Limit int
	// index of first contract event to fetch
	Offset int
}

// GetNetworkBlockTimesParams defines parameters for GetNetworkBlockTimes
type GetNetworkBlockTimesParams struct{}

// GetSTXSupplyTotalSupplyPlainParams defines parameters for GetSTXSupplyTotalSupplyPlain
type GetSTXSupplyTotalSupplyPlainParams struct{}

// GetDroppedMempoolTransactionListParams defines parameters for GetDroppedMempoolTransactionList
type GetDroppedMempoolTransactionListParams struct {
	// max number of mempool transactions to fetch
	Limit int
	// index of first mempool transaction to fetch
	Offset int
}

// RosettaConstructionHashParams defines parameters for RosettaConstructionHash
type RosettaConstructionHashParams struct {
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier interface{}
	// Signed transaction
	SignedTransaction string
}

// GetBurnchainRewardsTotalByAddressParams defines parameters for GetBurnchainRewardsTotalByAddress
type GetBurnchainRewardsTotalByAddressParams struct {
	// Reward recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format
	Address string
}

// RunFaucetBTCParams defines parameters for RunFaucetBTC
type RunFaucetBTCParams struct {
	// BTC testnet address
	Address string
}

// RosettaConstructionSubmitParams defines parameters for RosettaConstructionSubmit
type RosettaConstructionSubmitParams struct {
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier interface{}
	// Signed transaction
	SignedTransaction string
}

// GetNamePriceParams defines parameters for GetNamePrice
type GetNamePriceParams struct {
	// the name to query price information for
	Name string
}

// SearchByIDParams defines parameters for SearchByID
type SearchByIDParams struct {
	// The hex hash string for a block or transaction, account address, or contract address
	ID string
	// This includes the detailed data for purticular hash in the response
	IncludeMetadata bool
}

// RosettaMempoolTransactionParams defines parameters for RosettaMempoolTransaction
type RosettaMempoolTransactionParams struct {
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier interface{}
	// The transaction_identifier uniquely identifies a transaction in a particular network and block or in the mempool.
	TransactionIdentifier interface{}
}

// RosettaNetworkStatusParams defines parameters for RosettaNetworkStatus
type RosettaNetworkStatusParams struct {
	Metadata interface{}
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier interface{}
}

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

// GetBlockListParams defines parameters for GetBlockList
type GetBlockListParams struct {
	// max number of blocks to fetch
	Limit int
	// index of first block to fetch
	Offset int
}

// GetBurnchainRewardListParams defines parameters for GetBurnchainRewardList
type GetBurnchainRewardListParams struct {
	// max number of rewards to fetch
	Limit int
	// index of first rewards to fetch
	Offset int
}

// FetchFeeRateParams defines parameters for FetchFeeRate
type FetchFeeRateParams struct {
	// A serialized transaction
	Transaction string
}

// GetStatusParams defines parameters for GetStatus
type GetStatusParams struct{}

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

// GetHistoricalZoneFileParams defines parameters for GetHistoricalZoneFile
type GetHistoricalZoneFileParams struct {
	// fully-qualified name
	Name string
	// zone file hash
	ZoneFileHash string
}

// GetPoxInfoParams defines parameters for GetPoxInfo
type GetPoxInfoParams struct{}

// GetRawTransactionByIDParams defines parameters for GetRawTransactionByID
type GetRawTransactionByIDParams struct{}

// GetNetworkBlockTimeByNetworkParams defines parameters for GetNetworkBlockTimeByNetwork
type GetNetworkBlockTimeByNetworkParams struct {
	// Which network to retrieve the target block time of
	Network string
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
	// max number of transactions to fetch
	Limit int
	// index of first transaction to fetch
	Offset int
	// Filter by transaction type
	// [Coinbase TokenTransfer SmartContract ContractCall PoisonMicroblock]
	Type GetTransactionListType
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	Unanchored bool
}

// RosettaNetworkListParams defines parameters for RosettaNetworkList
type RosettaNetworkListParams struct{}

// GetNamespacePriceParams defines parameters for GetNamespacePrice
type GetNamespacePriceParams struct {
	// the namespace to fetch price for
	TLD string
}

// GetBlockByHeightParams defines parameters for GetBlockByHeight
type GetBlockByHeightParams struct{}

// GetUnanchoredTXsParams defines parameters for GetUnanchoredTXs
type GetUnanchoredTXsParams struct{}

// GetFTMetadataListParams defines parameters for GetFTMetadataList
type GetFTMetadataListParams struct {
	// max number of tokens to fetch
	Limit int
	// index of first tokens to fetch
	Offset int
}

// GetContractInterfaceParams defines parameters for GetContractInterface
type GetContractInterfaceParams struct{}

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

// GetBlockByHashParams defines parameters for GetBlockByHash
type GetBlockByHashParams struct{}

// GetSTXSupplyCirculatingPlainParams defines parameters for GetSTXSupplyCirculatingPlain
type GetSTXSupplyCirculatingPlainParams struct{}

// GetNFTHistoryParams defines parameters for GetNFTHistory
type GetNFTHistoryParams struct {
	// token asset class identifier
	AssetIDentifier string
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

// GetAllNamespacesParams defines parameters for GetAllNamespaces
type GetAllNamespacesParams struct{}

// GetContractByIDParams defines parameters for GetContractByID
type GetContractByIDParams struct{}

// RosettaBlockParams defines parameters for RosettaBlock
type RosettaBlockParams struct {
	// When fetching data by BlockIdentifier, it may be possible to only specify the index or hash. If neither property is specified, it is assumed that the client is making a request at the current block.
	BlockIdentifier interface{}
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier interface{}
}

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

// GetContractSourceParams defines parameters for GetContractSource
type GetContractSourceParams struct{}

// GetNamespaceNamesParams defines parameters for GetNamespaceNames
type GetNamespaceNamesParams struct {
	// the namespace to fetch names from
	TLD string
	// names are returned in pages of size 100, so specify the page number.
	Page int
}
