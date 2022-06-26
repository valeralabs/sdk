package api

// GetBlockByHeightParams defines parameters for GetBlockByHeight
type GetBlockByHeightParameters struct{}

// GetBurnchainRewardSlotHoldersByAddressParams defines parameters for GetBurnchainRewardSlotHoldersByAddress
type GetBurnchainRewardSlotHoldersByAddressParameters struct {
	// Reward slot holder recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format
	Address string
	// max number of items to fetch
	Limit int
	// index of the first items to fetch
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

// RosettaMempoolTransactionParams defines parameters for RosettaMempoolTransaction
type RosettaMempoolTransactionParameters struct {
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier any
	// The transaction_identifier uniquely identifies a transaction in a particular network and block or in the mempool.
	TransactionIdentifier any
}

// PostFeeTransactionParams defines parameters for PostFeeTransaction
type PostFeeTransactionParameters struct {
	EstimatedLen       int
	TransactionPayload string
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

// GetContractNFTMetadataParams defines parameters for GetContractNFTMetadata
type GetContractNFTMetadataParameters struct {
	// token's contract id
	ContractId string
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

// RosettaConstructionPayloadsParams defines parameters for RosettaConstructionPayloads
type RosettaConstructionPayloadsParameters struct {
	Metadata any
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier any
	Operations        any
	PublicKeys        any
}

// GetCoreAPIInfoParams defines parameters for GetCoreAPIInfo
type GetCoreAPIInfoParameters struct{}

// FetchSubdomainsListForNameParams defines parameters for FetchSubdomainsListForName
type FetchSubdomainsListForNameParameters struct {
	// fully-qualified name
	Name string
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

// GetBurnchainRewardsTotalByAddressParams defines parameters for GetBurnchainRewardsTotalByAddress
type GetBurnchainRewardsTotalByAddressParameters struct {
	// Reward recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format
	Address string
}

// GetSTXSupplyParams defines parameters for GetSTXSupply
type GetSTXSupplyParameters struct {
	// The block height at which to query supply details from, if not provided then the latest block height is used
	Height int
}

// GetSTXSupplyTotalSupplyPlainParams defines parameters for GetSTXSupplyTotalSupplyPlain
type GetSTXSupplyTotalSupplyPlainParameters struct{}

// GetRawTransactionByIDParams defines parameters for GetRawTransactionByID
type GetRawTransactionByIDParameters struct{}

// RosettaConstructionDeriveParams defines parameters for RosettaConstructionDerive
type RosettaConstructionDeriveParameters struct {
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier any
	// PublicKey contains a public key byte array for a particular CurveType encoded in hex. Note that there is no PrivateKey struct as this is NEVER the concern of an implementation.
	PublicKey any
	Metadata  any
}

// RosettaNetworkStatusParams defines parameters for RosettaNetworkStatus
type RosettaNetworkStatusParameters struct {
	Metadata any
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier any
}

// GetNetworkBlockTimesParams defines parameters for GetNetworkBlockTimes
type GetNetworkBlockTimesParameters struct{}

// RosettaConstructionPreprocessParams defines parameters for RosettaConstructionPreprocess
type RosettaConstructionPreprocessParameters struct {
	Metadata any
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier any
	Operations        any
	//  The caller can also provide a suggested fee multiplier to indicate that the suggested fee should be scaled. This may be used to set higher fees for urgent transactions or to pay lower fees when there is less urgency. It is assumed that providing a very low multiplier (like 0.0001) will never lead to a transaction being created with a fee less than the minimum network fee (if applicable). In the case that the caller provides both a max fee and a suggested fee multiplier, the max fee will set an upper bound on the suggested fee (regardless of the multiplier provided).
	SuggestedFeeMultiplier int
	MaxFee                 any
}

// RosettaNetworkOptionsParams defines parameters for RosettaNetworkOptions
type RosettaNetworkOptionsParameters struct {
	Metadata any
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier any
}

// GetAllNamesParams defines parameters for GetAllNames
type GetAllNamesParameters struct {
	// names are returned in pages of size 100, so specify the page number.
	Page int
}

// GetNamesOwnedByAddressParams defines parameters for GetNamesOwnedByAddress
type GetNamesOwnedByAddressParameters struct {
	// the layer-1 blockchain for the address
	Blockchain string
	// the address to lookup
	Address string
}

// GetBurnchainRewardListParams defines parameters for GetBurnchainRewardList
type GetBurnchainRewardListParameters struct {
	// max number of rewards to fetch
	Limit int
	// index of first rewards to fetch
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

// GetContractFTMetadataParams defines parameters for GetContractFTMetadata
type GetContractFTMetadataParameters struct {
	// token's contract id
	ContractId string
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

// RosettaConstructionCombineParams defines parameters for RosettaConstructionCombine
type RosettaConstructionCombineParameters struct {
	Signatures          any
	UnsignedTransaction string
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier any
}

// RosettaConstructionParseParams defines parameters for RosettaConstructionParse
type RosettaConstructionParseParameters struct {
	// Signed is a boolean indicating whether the transaction is signed.
	Signed bool
	// This must be either the unsigned transaction blob returned by /construction/payloads or the signed transaction blob returned by /construction/combine.
	Transaction string
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier any
}

// RosettaNetworkListParams defines parameters for RosettaNetworkList
type RosettaNetworkListParameters struct{}

// GetAllNamespacesParams defines parameters for GetAllNamespaces
type GetAllNamespacesParameters struct{}

// GetNameInfoParams defines parameters for GetNameInfo
type GetNameInfoParameters struct {
	// fully-qualified name
	Name string
}

// GetBlockByBurnBlockHashParams defines parameters for GetBlockByBurnBlockHash
type GetBlockByBurnBlockHashParameters struct{}

// GetMicroblockByHashParams defines parameters for GetMicroblockByHash
type GetMicroblockByHashParameters struct{}
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

// GetDroppedMempoolTransactionListParams defines parameters for GetDroppedMempoolTransactionList
type GetDroppedMempoolTransactionListParameters struct {
	// max number of mempool transactions to fetch
	Limit int
	// index of first mempool transaction to fetch
	Offset int
}

// GetTransactionByIDParams defines parameters for GetTransactionByID
type GetTransactionByIDParameters struct{}

// RosettaBlockParams defines parameters for RosettaBlock
type RosettaBlockParameters struct {
	// When fetching data by BlockIdentifier, it may be possible to only specify the index or hash. If neither property is specified, it is assumed that the client is making a request at the current block.
	BlockIdentifier any
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier any
}

// RosettaConstructionSubmitParams defines parameters for RosettaConstructionSubmit
type RosettaConstructionSubmitParameters struct {
	// Signed transaction
	SignedTransaction string
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier any
}

// GetHistoricalZoneFileParams defines parameters for GetHistoricalZoneFile
type GetHistoricalZoneFileParameters struct {
	// fully-qualified name
	Name string
	// zone file hash
	ZoneFileHash string
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

// SearchByIDParams defines parameters for SearchByID
type SearchByIDParameters struct {
	// The hex hash string for a block or transaction, account address, or contract address
	ID string
	// This includes the detailed data for purticular hash in the response
	IncludeMetadata bool
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

// GetTotalSTXSupplyLegacyFormatParams defines parameters for GetTotalSTXSupplyLegacyFormat
type GetTotalSTXSupplyLegacyFormatParameters struct {
	// The block height at which to query supply details from, if not provided then the latest block height is used
	Height int
}

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

// RosettaConstructionHashParams defines parameters for RosettaConstructionHash
type RosettaConstructionHashParameters struct {
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier any
	// Signed transaction
	SignedTransaction string
}

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

// GetPoxInfoParams defines parameters for GetPoxInfo
type GetPoxInfoParameters struct{}

// GetBlockListParams defines parameters for GetBlockList
type GetBlockListParameters struct {
	// max number of blocks to fetch
	Limit int
	// index of first block to fetch
	Offset int
}

// RunFaucetBTCParams defines parameters for RunFaucetBTC
type RunFaucetBTCParameters struct {
	// BTC testnet address
	Address string
}

// GetSTXSupplyCirculatingPlainParams defines parameters for GetSTXSupplyCirculatingPlain
type GetSTXSupplyCirculatingPlainParameters struct{}

// GetNFTMetadataListParams defines parameters for GetNFTMetadataList
type GetNFTMetadataListParameters struct {
	// max number of tokens to fetch
	Limit int
	// index of first tokens to fetch
	Offset int
}

// GetAccountSTXBalanceParams defines parameters for GetAccountSTXBalance
type GetAccountSTXBalanceParameters struct {
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
	Principal string
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks
	Unanchored bool
	// returned data representing the state up until that point in time, rather than the current block.
	UntilBlock string
}

// RunFaucetSTXParams defines parameters for RunFaucetSTX
type RunFaucetSTXParameters struct {
	// STX testnet address
	Address string
	// Use required number of tokens for stacking
	Stacking bool
}

// GetMicroblockListParams defines parameters for GetMicroblockList
type GetMicroblockListParameters struct {
	// Max number of microblocks to fetch
	Limit int
	// Index of the first microblock to fetch
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

// FetchZoneFileParams defines parameters for FetchZoneFile
type FetchZoneFileParameters struct {
	// fully-qualified name
	Name string
}

// GetBlockByHashParams defines parameters for GetBlockByHash
type GetBlockByHashParameters struct{}

// GetFTMetadataListParams defines parameters for GetFTMetadataList
type GetFTMetadataListParameters struct {
	// max number of tokens to fetch
	Limit int
	// index of first tokens to fetch
	Offset int
}

// GetNamespaceNamesParams defines parameters for GetNamespaceNames
type GetNamespaceNamesParameters struct {
	// the namespace to fetch names from
	TLD string
	// names are returned in pages of size 100, so specify the page number.
	Page int
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

// GetContractSourceParams defines parameters for GetContractSource
type GetContractSourceParameters struct{}

// GetAccountNoncesParams defines parameters for GetAccountNonces
type GetAccountNoncesParameters struct {
	// Stacks address (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0`)
	Principal string
	// Optionally get the nonce at a given block height
	BlockHeight int
	// Optionally get the nonce at a given block hash
	BlockHash string
}

// GetBlockByBurnBlockHeightParams defines parameters for GetBlockByBurnBlockHeight
type GetBlockByBurnBlockHeightParameters struct{}

// GetNetworkBlockTimeByNetworkParams defines parameters for GetNetworkBlockTimeByNetwork
type GetNetworkBlockTimeByNetworkParameters struct {
	// Which network to retrieve the target block time of
	Network string
}

// RosettaMempoolParams defines parameters for RosettaMempool
type RosettaMempoolParameters struct {
	Metadata any
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier any
}

// GetNamePriceParams defines parameters for GetNamePrice
type GetNamePriceParameters struct {
	// the name to query price information for
	Name string
}

// GetNamespacePriceParams defines parameters for GetNamespacePrice
type GetNamespacePriceParameters struct {
	// the namespace to fetch price for
	TLD string
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

// FetchFeeRateParams defines parameters for FetchFeeRate
type FetchFeeRateParameters struct {
	// A serialized transaction
	Transaction string
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
	Tip string
	// An array of hex serialized Clarity values
	Arguments any
	// The simulated tx-sender
	Sender string
}

// GetContractInterfaceParams defines parameters for GetContractInterface
type GetContractInterfaceParameters struct{}

// GetFeeTransferParams defines parameters for GetFeeTransfer
type GetFeeTransferParameters struct{}

// GetTXListDetailsParams defines parameters for GetTXListDetails
type GetTXListDetailsParameters struct{}

// GetBurnchainRewardSlotHoldersParams defines parameters for GetBurnchainRewardSlotHolders
type GetBurnchainRewardSlotHoldersParameters struct {
	// max number of items to fetch
	Limit int
	// index of the first items to fetch
	Offset int
}

// GetStatusParams defines parameters for GetStatus
type GetStatusParameters struct{}

// RosettaAccountBalanceParams defines parameters for RosettaAccountBalance
type RosettaAccountBalanceParameters struct {
	// The account_identifier uniquely identifies an account within a network. All fields in the account_identifier are utilized to determine this uniqueness (including the metadata field, if populated).
	AccountIdentifier any
	// When fetching data by BlockIdentifier, it may be possible to only specify the index or hash. If neither property is specified, it is assumed that the client is making a request at the current block.
	BlockIdentifier any
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier any
}

// RosettaBlockTransactionParams defines parameters for RosettaBlockTransaction
type RosettaBlockTransactionParameters struct {
	// When fetching data by BlockIdentifier, it may be possible to only specify the index or hash. If neither property is specified, it is assumed that the client is making a request at the current block.
	BlockIdentifier any
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier any
	// The transaction_identifier uniquely identifies a transaction in a particular network and block or in the mempool.
	TransactionIdentifier any
}

// RosettaConstructionMetadataParams defines parameters for RosettaConstructionMetadata
type RosettaConstructionMetadataParameters struct {
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier any
	// The options that will be sent directly to /construction/metadata by the caller.
	Options    any
	PublicKeys any
}

// PostCoreNodeTransactionsParams defines parameters for PostCoreNodeTransactions
type PostCoreNodeTransactionsParameters struct{}
