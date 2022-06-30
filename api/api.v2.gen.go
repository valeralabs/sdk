package api

type GetFilteredEventsType int64

const (
	SmartContractLog GetFilteredEventsType = iota
	STXLock
	STXAsset
	FungibleTokenAsset
	NonFungibleTokenAsset
)

// GetFilteredEventsParams defines parameters for GetFilteredEvents
type GetFilteredEventsParameters struct {
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
	Type GetFilteredEventsType
}

// RosettaConstructionParseParams defines parameters for RosettaConstructionParse
type RosettaConstructionParseParameters struct {
	NetworkIdentifier any
	// The network_identifier specifies which network a particular object is associated with.
	Signed bool
	// Signed is a boolean indicating whether the transaction is signed.
	Transaction string
	// This must be either the unsigned transaction blob returned by /construction/payloads or the signed transaction blob returned by /construction/combine.
}

// GetBurnchainRewardSlotHoldersParams defines parameters for GetBurnchainRewardSlotHolders
type GetBurnchainRewardSlotHoldersParameters struct {
	// max number of items to fetch
	Limit int
	// index of the first items to fetch
	Offset int
}

// GetTransactionsByBlockHashParams defines parameters for GetTransactionsByBlockHash
type GetTransactionsByBlockHashParameters struct {
	// Hash of block
	BlockHash string
	// max number of transactions to fetch
	Limit int
	// index of first transaction to fetch
	Offset int
}

// GetTXListDetailsParams defines parameters for GetTXListDetails
type GetTXListDetailsParameters struct{}

// RosettaConstructionDeriveParams defines parameters for RosettaConstructionDerive
type RosettaConstructionDeriveParameters struct {
	PublicKey any
	// PublicKey contains a public key byte array for a particular CurveType encoded in hex. Note that there is no PrivateKey struct as this is NEVER the concern of an implementation.
	Metadata          any
	NetworkIdentifier any
	// The network_identifier specifies which network a particular object is associated with.
}

// GetMicroblockByHashParams defines parameters for GetMicroblockByHash
type GetMicroblockByHashParameters struct{}

// GetNFTHistoryParams defines parameters for GetNFTHistory
type GetNFTHistoryParameters struct {
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

// GetSTXSupplyCirculatingPlainParams defines parameters for GetSTXSupplyCirculatingPlain
type GetSTXSupplyCirculatingPlainParameters struct{}

// RosettaConstructionPayloadsParams defines parameters for RosettaConstructionPayloads
type RosettaConstructionPayloadsParameters struct {
	PublicKeys        any
	Metadata          any
	NetworkIdentifier any
	// The network_identifier specifies which network a particular object is associated with.
	Operations any
}

// RosettaNetworkOptionsParams defines parameters for RosettaNetworkOptions
type RosettaNetworkOptionsParameters struct {
	Metadata          any
	NetworkIdentifier any
	// The network_identifier specifies which network a particular object is associated with.
}

// GetContractSourceParams defines parameters for GetContractSource
type GetContractSourceParameters struct{}

// GetContractDataMapEntryParams defines parameters for GetContractDataMapEntry
type GetContractDataMapEntryParameters struct {
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

// GetBurnchainRewardListByAddressParams defines parameters for GetBurnchainRewardListByAddress
type GetBurnchainRewardListByAddressParameters struct {
	// Reward recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format
	Address string
	// max number of rewards to fetch
	Limit int
	// index of first rewards to fetch
	Offset int
}

// FetchFeeRateParams defines parameters for FetchFeeRate
type FetchFeeRateParameters struct {
	Transaction string
	// A serialized transaction
}

// GetNetworkBlockTimesParams defines parameters for GetNetworkBlockTimes
type GetNetworkBlockTimesParameters struct{}

// GetStatusParams defines parameters for GetStatus
type GetStatusParameters struct{}

// GetContractNFTMetadataParams defines parameters for GetContractNFTMetadata
type GetContractNFTMetadataParameters struct {
	// token's contract id
	ContractId string
}

// RosettaConstructionMetadataParams defines parameters for RosettaConstructionMetadata
type RosettaConstructionMetadataParameters struct {
	NetworkIdentifier any
	// The network_identifier specifies which network a particular object is associated with.
	Options any
	// The options that will be sent directly to /construction/metadata by the caller.
	PublicKeys any
}

// RosettaMempoolTransactionParams defines parameters for RosettaMempoolTransaction
type RosettaMempoolTransactionParameters struct {
	NetworkIdentifier any
	// The network_identifier specifies which network a particular object is associated with.
	TransactionIdentifier any
	// The transaction_identifier uniquely identifies a transaction in a particular network and block or in the mempool.
}

// GetFeeTransferParams defines parameters for GetFeeTransfer
type GetFeeTransferParameters struct{}

// GetAddressMempoolTransactionsParams defines parameters for GetAddressMempoolTransactions
type GetAddressMempoolTransactionsParameters struct {
	// Transactions for the address
	Address string
	// max number of transactions to fetch
	Limit int
	// index of first transaction to fetch
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	Unanchored bool
}

// GetBlockByBurnBlockHeightParams defines parameters for GetBlockByBurnBlockHeight
type GetBlockByBurnBlockHeightParameters struct{}

// GetNamePriceParams defines parameters for GetNamePrice
type GetNamePriceParameters struct {
	// the name to query price information for
	Name string
}

// GetNamespaceNamesParams defines parameters for GetNamespaceNames
type GetNamespaceNamesParameters struct {
	// the namespace to fetch names from
	TLD string
	// names are returned in pages of size 100, so specify the page number.
	Page int
}

// GetTransactionByIDParams defines parameters for GetTransactionByID
type GetTransactionByIDParameters struct{}

// RosettaMempoolParams defines parameters for RosettaMempool
type RosettaMempoolParameters struct {
	Metadata          any
	NetworkIdentifier any
	// The network_identifier specifies which network a particular object is associated with.
}

// GetMicroblockListParams defines parameters for GetMicroblockList
type GetMicroblockListParameters struct {
	// Max number of microblocks to fetch
	Limit int
	// Index of the first microblock to fetch
	Offset int
}

// GetNFTMintsParams defines parameters for GetNFTMints
type GetNFTMintsParameters struct {
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

// GetNameInfoParams defines parameters for GetNameInfo
type GetNameInfoParameters struct {
	// fully-qualified name
	Name string
}

// FetchSubdomainsListForNameParams defines parameters for FetchSubdomainsListForName
type FetchSubdomainsListForNameParameters struct {
	// fully-qualified name
	Name string
}

// GetCoreAPIInfoParams defines parameters for GetCoreAPIInfo
type GetCoreAPIInfoParameters struct{}

// PostCoreNodeTransactionsParams defines parameters for PostCoreNodeTransactions
type PostCoreNodeTransactionsParameters struct{}

// GetAccountBalanceParams defines parameters for GetAccountBalance
type GetAccountBalanceParameters struct {
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
	Principal string
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	Unanchored bool
	// returned data representing the state up until that point in time, rather than the current block.
	UntilBlock string
}

// GetBurnchainRewardsTotalByAddressParams defines parameters for GetBurnchainRewardsTotalByAddress
type GetBurnchainRewardsTotalByAddressParameters struct {
	// Reward recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format
	Address string
}

// GetBlockByHeightParams defines parameters for GetBlockByHeight
type GetBlockByHeightParameters struct{}

// GetTotalSTXSupplyLegacyFormatParams defines parameters for GetTotalSTXSupplyLegacyFormat
type GetTotalSTXSupplyLegacyFormatParameters struct {
	// The block height at which to query supply details from, if not provided then the latest block height is used
	Height int
}

// GetDroppedMempoolTransactionListParams defines parameters for GetDroppedMempoolTransactionList
type GetDroppedMempoolTransactionListParameters struct {
	// max number of mempool transactions to fetch
	Limit int
	// index of first mempool transaction to fetch
	Offset int
}

// RosettaConstructionSubmitParams defines parameters for RosettaConstructionSubmit
type RosettaConstructionSubmitParameters struct {
	NetworkIdentifier any
	// The network_identifier specifies which network a particular object is associated with.
	SignedTransaction string
	// Signed transaction
}

// GetAccountNoncesParams defines parameters for GetAccountNonces
type GetAccountNoncesParameters struct {
	// Stacks address (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0`)
	Principal string
	// Optionally get the nonce at a given block height
	BlockHeight int
	// Optionally get the nonce at a given block hash
	BlockHash string
}

// GetSingleTransactionWithTransfersParams defines parameters for GetSingleTransactionWithTransfers
type GetSingleTransactionWithTransfersParameters struct {
	// Stacks address or a contract identifier
	Principal string
	// Transaction id
	TXID string
}

// PostFeeTransactionParams defines parameters for PostFeeTransaction
type PostFeeTransactionParameters struct {
	EstimatedLen       int
	TransactionPayload string
}

// GetBurnchainRewardListParams defines parameters for GetBurnchainRewardList
type GetBurnchainRewardListParameters struct {
	// max number of rewards to fetch
	Limit int
	// index of first rewards to fetch
	Offset int
}

// GetNFTHoldingsParams defines parameters for GetNFTHoldings
type GetNFTHoldingsParameters struct {
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

// GetNFTMetadataListParams defines parameters for GetNFTMetadataList
type GetNFTMetadataListParameters struct {
	// max number of tokens to fetch
	Limit int
	// index of first tokens to fetch
	Offset int
}

// GetRawTransactionByIDParams defines parameters for GetRawTransactionByID
type GetRawTransactionByIDParameters struct{}

// GetBlockByBurnBlockHashParams defines parameters for GetBlockByBurnBlockHash
type GetBlockByBurnBlockHashParameters struct{}

// RunFaucetSTXParams defines parameters for RunFaucetSTX
type RunFaucetSTXParameters struct {
	Stacking bool
	// Use required number of tokens for stacking
	Address string
	// STX testnet address
}

// RosettaConstructionHashParams defines parameters for RosettaConstructionHash
type RosettaConstructionHashParameters struct {
	SignedTransaction string
	// Signed transaction
	NetworkIdentifier any
	// The network_identifier specifies which network a particular object is associated with.
}

// RosettaNetworkListParams defines parameters for RosettaNetworkList
type RosettaNetworkListParameters struct{}

// GetAllNamespacesParams defines parameters for GetAllNamespaces
type GetAllNamespacesParameters struct{}

// GetBlockByHashParams defines parameters for GetBlockByHash
type GetBlockByHashParameters struct{}

// RunFaucetBTCParams defines parameters for RunFaucetBTC
type RunFaucetBTCParameters struct {
	Address string
	// BTC testnet address
}

// GetNetworkBlockTimeByNetworkParams defines parameters for GetNetworkBlockTimeByNetwork
type GetNetworkBlockTimeByNetworkParameters struct {
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
type GetTransactionListParameters struct {
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

// RosettaConstructionPreprocessParams defines parameters for RosettaConstructionPreprocess
type RosettaConstructionPreprocessParameters struct {
	MaxFee            any
	Metadata          any
	NetworkIdentifier any
	// The network_identifier specifies which network a particular object is associated with.
	Operations             any
	SuggestedFeeMultiplier int
	//  The caller can also provide a suggested fee multiplier to indicate that the suggested fee should be scaled. This may be used to set higher fees for urgent transactions or to pay lower fees when there is less urgency. It is assumed that providing a very low multiplier (like 0.0001) will never lead to a transaction being created with a fee less than the minimum network fee (if applicable). In the case that the caller provides both a max fee and a suggested fee multiplier, the max fee will set an upper bound on the suggested fee (regardless of the multiplier provided).
}

// GetAccountNFTParams defines parameters for GetAccountNFT
type GetAccountNFTParameters struct {
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
type GetAccountTransactionsParameters struct {
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

// GetContractByIDParams defines parameters for GetContractByID
type GetContractByIDParameters struct{}

// GetAccountSTXBalanceParams defines parameters for GetAccountSTXBalance
type GetAccountSTXBalanceParameters struct {
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
	Principal string
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	Unanchored bool
	// returned data representing the state up until that point in time, rather than the current block.
	UntilBlock string
}

// GetAccountInboundParams defines parameters for GetAccountInbound
type GetAccountInboundParameters struct {
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
type GetContractFTMetadataParameters struct {
	// token's contract id
	ContractId string
}

// GetMempoolTransactionListParams defines parameters for GetMempoolTransactionList
type GetMempoolTransactionListParameters struct {
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

// RosettaConstructionCombineParams defines parameters for RosettaConstructionCombine
type RosettaConstructionCombineParameters struct {
	NetworkIdentifier any
	// The network_identifier specifies which network a particular object is associated with.
	Signatures          any
	UnsignedTransaction string
}

// RosettaNetworkStatusParams defines parameters for RosettaNetworkStatus
type RosettaNetworkStatusParameters struct {
	Metadata          any
	NetworkIdentifier any
	// The network_identifier specifies which network a particular object is associated with.
}

// GetAccountInfoParams defines parameters for GetAccountInfo
type GetAccountInfoParameters struct {
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
	Principal string
	// Returns object without the proof field if set to 0
	Proof int
	// The Stacks chain tip to query from
	Tip string
}

// CallReadOnlyFunctionParams defines parameters for CallReadOnlyFunction
type CallReadOnlyFunctionParameters struct {
	// Stacks address
	ContractAddress string
	// Contract name
	ContractName string
	// Function name
	FunctionName string
	// The Stacks chain tip to query from
	Tip       string
	Arguments any
	// An array of hex serialized Clarity values
	Sender string
	// The simulated tx-sender
}

// GetSTXSupplyParams defines parameters for GetSTXSupply
type GetSTXSupplyParameters struct {
	// The block height at which to query supply details from, if not provided then the latest block height is used
	Height int
}

// GetFTMetadataListParams defines parameters for GetFTMetadataList
type GetFTMetadataListParameters struct {
	// max number of tokens to fetch
	Limit int
	// index of first tokens to fetch
	Offset int
}

// GetContractsByTraitParams defines parameters for GetContractsByTrait
type GetContractsByTraitParameters struct {
	// JSON abi of the trait.
	TraitABI string
	// max number of contracts fetch
	Limit int
	// index of first contract event to fetch
	Offset int
}

// GetSTXSupplyTotalSupplyPlainParams defines parameters for GetSTXSupplyTotalSupplyPlain
type GetSTXSupplyTotalSupplyPlainParameters struct{}

// GetHistoricalZoneFileParams defines parameters for GetHistoricalZoneFile
type GetHistoricalZoneFileParameters struct {
	// fully-qualified name
	Name string
	// zone file hash
	ZoneFileHash string
}

// GetAccountAssetsParams defines parameters for GetAccountAssets
type GetAccountAssetsParameters struct {
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

// GetBurnchainRewardSlotHoldersByAddressParams defines parameters for GetBurnchainRewardSlotHoldersByAddress
type GetBurnchainRewardSlotHoldersByAddressParameters struct {
	// Reward slot holder recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format
	Address string
	// max number of items to fetch
	Limit int
	// index of the first items to fetch
	Offset int
}

// GetNamesOwnedByAddressParams defines parameters for GetNamesOwnedByAddress
type GetNamesOwnedByAddressParameters struct {
	// the layer-1 blockchain for the address
	Blockchain string
	// the address to lookup
	Address string
}

// GetAllNamesParams defines parameters for GetAllNames
type GetAllNamesParameters struct {
	// names are returned in pages of size 100, so specify the page number.
	Page int
}

// FetchZoneFileParams defines parameters for FetchZoneFile
type FetchZoneFileParameters struct {
	// fully-qualified name
	Name string
}

// GetContractInterfaceParams defines parameters for GetContractInterface
type GetContractInterfaceParameters struct{}

// GetNamespacePriceParams defines parameters for GetNamespacePrice
type GetNamespacePriceParameters struct {
	// the namespace to fetch price for
	TLD string
}

// GetContractEventsByIDParams defines parameters for GetContractEventsByID
type GetContractEventsByIDParameters struct {
	// Contract identifier formatted as `<contract_address>.<contract_name>`
	ContractID string
	// max number of contract events to fetch
	Limit int
	// index of first contract event to fetch
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	Unanchored bool
}

// RosettaAccountBalanceParams defines parameters for RosettaAccountBalance
type RosettaAccountBalanceParameters struct {
	AccountIdentifier any
	// The account_identifier uniquely identifies an account within a network. All fields in the account_identifier are utilized to determine this uniqueness (including the metadata field, if populated).
	BlockIdentifier any
	// When fetching data by BlockIdentifier, it may be possible to only specify the index or hash. If neither property is specified, it is assumed that the client is making a request at the current block.
	NetworkIdentifier any
	// The network_identifier specifies which network a particular object is associated with.
}

// GetUnanchoredTxsParams defines parameters for GetUnanchoredTxs
type GetUnanchoredTxsParameters struct{}

// SearchByIDParams defines parameters for SearchByID
type SearchByIDParameters struct {
	// The hex hash string for a block or transaction, account address, or contract address
	ID string
	// This includes the detailed data for purticular hash in the response
	IncludeMetadata bool
}

// GetTransactionsByBlockHeightParams defines parameters for GetTransactionsByBlockHeight
type GetTransactionsByBlockHeightParameters struct {
	// Height of block
	Height int
	// max number of transactions to fetch
	Limit int
	// index of first transaction to fetch
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	Unanchored bool
}

// RosettaBlockParams defines parameters for RosettaBlock
type RosettaBlockParameters struct {
	BlockIdentifier any
	// When fetching data by BlockIdentifier, it may be possible to only specify the index or hash. If neither property is specified, it is assumed that the client is making a request at the current block.
	NetworkIdentifier any
	// The network_identifier specifies which network a particular object is associated with.
}

// RosettaBlockTransactionParams defines parameters for RosettaBlockTransaction
type RosettaBlockTransactionParameters struct {
	BlockIdentifier any
	// When fetching data by BlockIdentifier, it may be possible to only specify the index or hash. If neither property is specified, it is assumed that the client is making a request at the current block.
	NetworkIdentifier any
	// The network_identifier specifies which network a particular object is associated with.
	TransactionIdentifier any
	// The transaction_identifier uniquely identifies a transaction in a particular network and block or in the mempool.
}

// GetPoxInfoParams defines parameters for GetPoxInfo
type GetPoxInfoParameters struct{}

// GetAccountTransactionsWithTransfersParams defines parameters for GetAccountTransactionsWithTransfers
type GetAccountTransactionsWithTransfersParameters struct {
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

// GetBlockListParams defines parameters for GetBlockList
type GetBlockListParameters struct {
	// max number of blocks to fetch
	Limit int
	// index of first block to fetch
	Offset int
}
