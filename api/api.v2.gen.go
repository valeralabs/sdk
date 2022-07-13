package api

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Retrieves metadata for non fungible tokens for a given contract id. More information on Non-Fungible Tokens on the Stacks blockchain can be found [here](https://docs.stacks.co/write-smart-contracts/tokens#non-fungible-tokens-nfts).
func GetContractNFTMetadata(server Server, params GetContractNFTMetadataParams) (GetContractNFTMetadataResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/tokens/{contractId}/nft/metadata", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetContractNFTMetadataResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetContractNFTMetadataResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetContractNFTMetadataResponse{}, err
		}
		return GetContractNFTMetadataResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Non fungible tokens metadata for contract id) response for GetContractNFTMetadata.
type GetContractNFTMetadataResponse NonFungibleTokenMetadata

// Defines parameters for GetContractNFTMetadata
type GetContractNFTMetadataParams struct {
	// Token's contract id.
	ContractId string
}

/*
Retrieves all recently mined transactions

If using TypeScript, import typings for this response from our types package:

`import type { TransactionResults } from '@stacks/stacks-blockchain-api-types';`
*/
func GetTransactionList(server Server, params GetTransactionListParams) (GetTransactionListResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/tx", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetTransactionListResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetTransactionListResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetTransactionListResponse{}, err
		}
		return GetTransactionListResponse{}, errors.New(respObj.Error)
	}
}

type GetTransactionListType int64

const (
	Coinbase GetTransactionListType = iota
	TokenTransfer
	SmartContract
	ContractCall
	PoisonMicroblock
)

// Defines a 200 (List of transactions) response for GetTransactionList.
type GetTransactionListResponse TransactionResults

// Defines parameters for GetTransactionList
type GetTransactionListParams struct {
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. The default value is false.
	Unanchored bool
	// Max number of transactions to fetch.
	// Optional. The default value is 96.
	// Max value is 200.
	Limit int
	// Index of first transaction to fetch.
	// Optional.
	Offset int
	// Filter by transaction type.
	// Optional.
	Type GetTransactionListType
}

/*
Retrieves a list of all Transactions for a given Address or Contract Identifier. More information on Transaction types can be found [here](https://docs.stacks.co/understand-stacks/transactions#types).

If you need to actively monitor new transactions for an address or contract id, we highly recommend subscribing to [WebSockets or Socket.io](https://github.com/hirosystems/stacks-blockchain-api/tree/master/client) for real-time updates.
*/
func GetAccountTransactions(server Server, params GetAccountTransactionsParams) (GetAccountTransactionsResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/address/{principal}/transactions", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetAccountTransactionsResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetAccountTransactionsResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetAccountTransactionsResponse{}, err
		}
		return GetAccountTransactionsResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for GetAccountTransactions.
type GetAccountTransactionsResponse AddressTransactionsListResponse

// Defines parameters for GetAccountTransactions
type GetAccountTransactionsParams struct {
	// Index of first account transaction to fetch.
	// Optional.
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. The default value is false.
	Unanchored bool
	// Returned data representing the state up until that point in time, rather than the current block.
	// Optional.
	UntilBlock string
	// Filter for transactions only at this given block height.
	// Optional.
	Height int
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
	// Max number of account transactions to fetch.
	// Optional.
	Limit int
}

// Retrieves block details of a specific block for a given chain height
func GetBlockByHash(server Server) (GetBlockByHashResponse, error) {
	url := fmt.Sprintf("%s%s", server, "/extended/v1/block/{hash}")
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetBlockByHashResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetBlockByHashResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetBlockByHashResponse{}, err
		}
		return GetBlockByHashResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Block) response for GetBlockByHash.
type GetBlockByHashResponse Block

// Retrieves the target block times for mainnet and testnet. The block time is hardcoded and will change throughout the implementation phases of the testnet.
func GetNetworkBlockTimes(server Server) (GetNetworkBlockTimesResponse, error) {
	url := fmt.Sprintf("%s%s", server, "/extended/v1/info/network_block_times")
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetNetworkBlockTimesResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetNetworkBlockTimesResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetNetworkBlockTimesResponse{}, err
		}
		return GetNetworkBlockTimesResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for GetNetworkBlockTimes.
type GetNetworkBlockTimesResponse NetworkBlockTimesResponse

// Retrieves the metadata for fungible tokens for a given contract id
func GetContractFTMetadata(server Server, params GetContractFTMetadataParams) (GetContractFTMetadataResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/tokens/{contractId}/ft/metadata", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetContractFTMetadataResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetContractFTMetadataResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetContractFTMetadataResponse{}, err
		}
		return GetContractFTMetadataResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Fungible tokens metadata for contract id) response for GetContractFTMetadata.
type GetContractFTMetadataResponse FungibleTokenMetadata

// Defines parameters for GetContractFTMetadata
type GetContractFTMetadataParams struct {
	// Token's contract id.
	ContractId string
}

// Retrieves a contract interface with a given `contract_address` and `contract name`
func GetContractInterface(server Server) (GetContractInterfaceResponse, error) {
	url := fmt.Sprintf("%s%s", server, "/v2/contracts/interface/{contract_address}/{contract_name}")
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetContractInterfaceResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetContractInterfaceResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetContractInterfaceResponse{}, err
		}
		return GetContractInterfaceResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Contract interface) response for GetContractInterface.
type GetContractInterfaceResponse ContractInterfaceResponse

// Retrieves the Clarity source code of a given contract, along with the block height it was published in, and the MARF proof for the data
func GetContractSource(server Server) (GetContractSourceResponse, error) {
	url := fmt.Sprintf("%s%s", server, "/v2/contracts/source/{contract_address}/{contract_name}")
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetContractSourceResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetContractSourceResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetContractSourceResponse{}, err
		}
		return GetContractSourceResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for GetContractSource.
type GetContractSourceResponse ContractSourceResponse

// Retrieves a list of names within a given namespace.
func GetNamespaceNames(server Server, params GetNamespaceNamesParams) (GetNamespaceNamesResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/v1/namespaces/{TLD}/names", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetNamespaceNamesResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetNamespaceNamesResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetNamespaceNamesResponse{}, err
		}
		return GetNamespaceNamesResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 400 (Error) response for GetNamespaceNames.
type GetNamespaceNamesError400 BnsError

// Defines a 200 (Success) response for GetNamespaceNames.
type GetNamespaceNamesResponse []BnsGetAllNamespacesNamesResponse

// Defines a 404 (Error) response for GetNamespaceNames.
type GetNamespaceNamesError404 BnsError

// Defines parameters for GetNamespaceNames
type GetNamespaceNamesParams struct {
	// The namespace to fetch names from.
	Tld string
	// Names are returned in pages of size 100, so specify the page number.
	Page int
}
type GetContractDataMapEntryBody struct{}

/*
Attempt to fetch data from a contract data map. The contract is identified with Stacks Address `contract_address` and Contract Name `contract_address` in the URL path. The map is identified with [Map Name].

The key to lookup in the map is supplied via the POST body. This should be supplied as the hex string serialization of the key (which should be a Clarity value). Note, this is a JSON string atom.

In the response, `data` is the hex serialization of the map response. Note that map responses are Clarity option types, for non-existent values, this is a serialized none, and for all other responses, it is a serialized (some ...) object.
*/
func GetContractDataMapEntry(server Server, params GetContractDataMapEntryParams, body GetContractDataMapEntryBody) (GetContractDataMapEntryResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/v2/map_entry/{contract_address}/{contract_name}/{map_name}", params))
	var returnedErr error
	sendBody, _ := json.Marshal(body)
	resp, status := makePostReq(url, string(sendBody), &returnedErr)

	if returnedErr != nil {
		return GetContractDataMapEntryResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetContractDataMapEntryResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetContractDataMapEntryResponse{}, err
		}
		return GetContractDataMapEntryResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for GetContractDataMapEntry.
type GetContractDataMapEntryResponse MapEntryResponse

// Defines parameters for GetContractDataMapEntry
type GetContractDataMapEntryParams struct {
	// Returns object without the proof field when set to 0.
	// Optional.
	Proof int
	// Contract name.
	ContractName string
	// Stacks address.
	ContractAddress string
	// The Stacks chain tip to query from.
	// Optional.
	Tip string
}

// Retrieves the total burnchain (e.g. Bitcoin) rewards for a given recipient `address`
func GetBurnchainRewardsTotalByAddress(server Server, params GetBurnchainRewardsTotalByAddressParams) (GetBurnchainRewardsTotalByAddressResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/burnchain/rewards/{address}/total", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetBurnchainRewardsTotalByAddressResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetBurnchainRewardsTotalByAddressResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetBurnchainRewardsTotalByAddressResponse{}, err
		}
		return GetBurnchainRewardsTotalByAddressResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (List of burnchain reward recipients and amounts) response for GetBurnchainRewardsTotalByAddress.
type GetBurnchainRewardsTotalByAddressResponse BurnchainRewardsTotal

// Defines parameters for GetBurnchainRewardsTotalByAddress
type GetBurnchainRewardsTotalByAddressParams struct {
	// Reward recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format.
	Address string
}

// Defines a 200 (success) response for GetSTXSupplyCirculatingPlain.
type GetSTXSupplyCirculatingPlainResponse string

// Retrieves the STX tokens currently in circulation that have been unlocked as plain text.
func GetSTXSupplyCirculatingPlain(server Server) (GetSTXSupplyCirculatingPlainResponse, error) {
	url := fmt.Sprintf("%s%s", server, "/extended/v1/stx_supply/circulating/plain")
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetSTXSupplyCirculatingPlainResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetSTXSupplyCirculatingPlainResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetSTXSupplyCirculatingPlainResponse{}, err
		}
		return GetSTXSupplyCirculatingPlainResponse{}, errors.New(respObj.Error)
	}
}

/*
Retrieves all events relevant to a Non-Fungible Token. Useful to determine the ownership history of a particular asset.

More information on Non-Fungible Tokens on the Stacks blockchain can be found [here](https://docs.stacks.co/write-smart-contracts/tokens#non-fungible-tokens-nfts).
*/
func GetNFTHistory(server Server, params GetNFTHistoryParams) (GetNFTHistoryResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/tokens/nft/history", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetNFTHistoryResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetNFTHistoryResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetNFTHistoryResponse{}, err
		}
		return GetNFTHistoryResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Non-Fungible Token event history) response for GetNFTHistory.
type GetNFTHistoryResponse NonFungibleTokenHistoryEventList

// Defines parameters for GetNFTHistory
type GetNFTHistoryParams struct {
	// Index of first event to fetch.
	// Optional. The default value is 0.
	Offset int
	// Token asset class identifier.
	AssetIdentifier string
	// Whether or not to include events from unconfirmed transactions.
	// Optional. The default value is false.
	Unanchored bool
	// Whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times.
	// Optional. The default value is false.
	TxMetadata bool
	// Max number of events to fetch.
	// Optional. The default value is 50.
	Limit int
}

// Retrieves a list of non fungible tokens with their metadata. More information on Non-Fungible Tokens on the Stacks blockchain can be found [here](https://docs.stacks.co/write-smart-contracts/tokens#non-fungible-tokens-nfts).
func GetNFTMetadataList(server Server, params GetNFTMetadataListParams) (GetNFTMetadataListResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/tokens/nft/metadata", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetNFTMetadataListResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetNFTMetadataListResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetNFTMetadataListResponse{}, err
		}
		return GetNFTMetadataListResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (List of non fungible tokens metadata) response for GetNFTMetadataList.
type GetNFTMetadataListResponse NonFungibleTokensMetadataList

// Defines parameters for GetNFTMetadataList
type GetNFTMetadataListParams struct {
	// Max number of tokens to fetch.
	// Optional.
	Limit int
	// Index of first tokens to fetch.
	// Optional.
	Offset int
}

// Retrieves the price of a namespace. The `amount` given will be in the smallest possible units of the currency.
func GetNamespacePrice(server Server, params GetNamespacePriceParams) (GetNamespacePriceResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/v2/prices/namespaces/{TLD}", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetNamespacePriceResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetNamespacePriceResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetNamespacePriceResponse{}, err
		}
		return GetNamespacePriceResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for GetNamespacePrice.
type GetNamespacePriceResponse BnsGetNamespacePriceResponse

// Defines parameters for GetNamespacePrice
type GetNamespacePriceParams struct {
	// The namespace to fetch price for.
	Tld string
}

// Retrieves the price of a name. The `amount` given will be in the smallest possible units of the currency.
func GetNamePrice(server Server, params GetNamePriceParams) (GetNamePriceResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/v2/prices/names/{name}", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetNamePriceResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetNamePriceResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetNamePriceResponse{}, err
		}
		return GetNamePriceResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for GetNamePrice.
type GetNamePriceResponse BnsGetNamePriceResponse

// Defines parameters for GetNamePrice
type GetNamePriceParams struct {
	// The name to query price information for.
	Name string
}

// Retrieves the target block time for a given network. The network can be mainnet or testnet. The block time is hardcoded and will change throughout the implementation phases of the testnet.
func GetNetworkBlockTimeByNetwork(server Server, params GetNetworkBlockTimeByNetworkParams) (GetNetworkBlockTimeByNetworkResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/info/network_block_time/{network}", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetNetworkBlockTimeByNetworkResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetNetworkBlockTimeByNetworkResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetNetworkBlockTimeByNetworkResponse{}, err
		}
		return GetNetworkBlockTimeByNetworkResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for GetNetworkBlockTimeByNetwork.
type GetNetworkBlockTimeByNetworkResponse any

// Defines parameters for GetNetworkBlockTimeByNetwork
type GetNetworkBlockTimeByNetworkParams struct {
	// Which network to retrieve the target block time of.
	Network string
}

// Retrieves total supply of STX tokens including those currently in circulation that have been unlocked.
func GetTotalSTXSupplyLegacyFormat(server Server, params GetTotalSTXSupplyLegacyFormatParams) (GetTotalSTXSupplyLegacyFormatResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/stx_supply/legacy_format", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetTotalSTXSupplyLegacyFormatResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetTotalSTXSupplyLegacyFormatResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetTotalSTXSupplyLegacyFormatResponse{}, err
		}
		return GetTotalSTXSupplyLegacyFormatResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for GetTotalSTXSupplyLegacyFormat.
type GetTotalSTXSupplyLegacyFormatResponse GetStxSupplyLegacyFormatResponse

// Defines parameters for GetTotalSTXSupplyLegacyFormat
type GetTotalSTXSupplyLegacyFormatParams struct {
	// The block height at which to query supply details from, if not provided then the latest block height is used.
	// Optional.
	Height int
}

// Retrieves list of fungible tokens with their metadata. More information on Fungible Tokens on the Stacks blockchain can be found [here](https://docs.stacks.co/write-smart-contracts/tokens#fungible-tokens).
func GetFTMetadataList(server Server, params GetFTMetadataListParams) (GetFTMetadataListResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/tokens/ft/metadata", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetFTMetadataListResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetFTMetadataListResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetFTMetadataListResponse{}, err
		}
		return GetFTMetadataListResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (List of fungible tokens metadata) response for GetFTMetadataList.
type GetFTMetadataListResponse FungibleTokensMetadataList

// Defines parameters for GetFTMetadataList
type GetFTMetadataListParams struct {
	// Index of first tokens to fetch.
	// Optional.
	Offset int
	// Max number of tokens to fetch.
	// Optional.
	Limit int
}

// Retrieves the list of events filtered by principal (STX address or Smart Contract ID), transaction id or event types. The list of event types is ('smart_contract_log', 'stx_lock', 'stx_asset', 'fungible_token_asset', 'non_fungible_token_asset').
func GetFilteredEvents(server Server, params GetFilteredEventsParams) (GetFilteredEventsResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/tx/events", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetFilteredEventsResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetFilteredEventsResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetFilteredEventsResponse{}, err
		}
		return GetFilteredEventsResponse{}, errors.New(respObj.Error)
	}
}

type GetFilteredEventsType int64

const (
	SmartContractLog GetFilteredEventsType = iota
	StxLock
	StxAsset
	FungibleTokenAsset
	NonFungibleTokenAsset
)

// Defines a 200 (Success) response for GetFilteredEvents.
type GetFilteredEventsResponse TransactionEventsResponse

// Defines parameters for GetFilteredEvents
type GetFilteredEventsParams struct {
	// Stacks address or a Contract identifier.
	// Optional.
	Address string
	// Hash of transaction.
	// Optional.
	TxID string
	// Number of items to return.
	// Optional.
	Limit int
	// Filter the events on event type.
	// Optional.
	Type GetFilteredEventsType
}

// Retrieves the list of subdomains for a specific name
func FetchSubdomainsListForName(server Server, params FetchSubdomainsListForNameParams) (FetchSubdomainsListForNameResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/v1/names/{name}/subdomains", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return FetchSubdomainsListForNameResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj FetchSubdomainsListForNameResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return FetchSubdomainsListForNameResponse{}, err
		}
		return FetchSubdomainsListForNameResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for FetchSubdomainsListForName.
type FetchSubdomainsListForNameResponse []GetAllSubdomainsInName

// Defines parameters for FetchSubdomainsListForName
type FetchSubdomainsListForNameParams struct {
	// Fully-qualified name.
	Name string
}

// Retrieves block details of a specific block for a given burnchain block hash
func GetBlockByBurnBlockHash(server Server) (GetBlockByBurnBlockHashResponse, error) {
	url := fmt.Sprintf("%s%s", server, "/extended/v1/block/by_burn_block_hash/{burn_block_hash}")
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetBlockByBurnBlockHashResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetBlockByBurnBlockHashResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetBlockByBurnBlockHashResponse{}, err
		}
		return GetBlockByBurnBlockHashResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Block) response for GetBlockByBurnBlockHash.
type GetBlockByBurnBlockHashResponse Block

// Retrieves details of a contract with a given `contract_id`
func GetContractByID(server Server) (GetContractByIDResponse, error) {
	url := fmt.Sprintf("%s%s", server, "/extended/v1/contract/{contract_ID}")
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetContractByIDResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetContractByIDResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetContractByIDResponse{}, err
		}
		return GetContractByIDResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Contract found) response for GetContractByID.
type GetContractByIDResponse SmartContractTransaction

/*
Retrieves a hex encoded serialized transaction for a given ID
*/
func GetRawTransactionByID(server Server) (GetRawTransactionByIDResponse, error) {
	url := fmt.Sprintf("%s%s", server, "/extended/v1/tx/{TX_ID}/raw")
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetRawTransactionByIDResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetRawTransactionByIDResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetRawTransactionByIDResponse{}, err
		}
		return GetRawTransactionByIDResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Hex encoded serialized transaction) response for GetRawTransactionByID.
type GetRawTransactionByIDResponse GetRawTransactionResult

// Retrieves a list of contracts based on the following traits listed in JSON format -  functions, variables, maps, fungible tokens and non-fungible tokens
func GetContractsByTrait(server Server, params GetContractsByTraitParams) (GetContractsByTraitResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/contract/by_trait", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetContractsByTraitResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetContractsByTraitResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetContractsByTraitResponse{}, err
		}
		return GetContractsByTraitResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (List of contracts implement given trait) response for GetContractsByTrait.
type GetContractsByTraitResponse ContractListResponse

// Defines parameters for GetContractsByTrait
type GetContractsByTraitParams struct {
	// Max number of contracts fetch.
	// Optional.
	Limit int
	// Index of first contract event to fetch.
	// Optional.
	Offset int
	// JSON abi of the trait.
	TraitABI string
}
type RunFaucetBTCBody struct{}

/*
Add 1 BTC token to the specified testnet BTC address.

The endpoint returns the transaction ID, which you can use to view the transaction in a testnet Bitcoin block
explorer. The tokens are delivered once the transaction has been included in a block.

**Note:** This is a testnet only endpoint. This endpoint will not work on the mainnet.
*/
func RunFaucetBTC(server Server, body RunFaucetBTCBody) (RunFaucetBTCResponse, error) {
	url := fmt.Sprintf("%s%s", server, "/extended/v1/faucets/btc")
	var returnedErr error
	sendBody, _ := json.Marshal(body)
	resp, status := makePostReq(url, string(sendBody), &returnedErr)

	if returnedErr != nil {
		return RunFaucetBTCResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj RunFaucetBTCResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return RunFaucetBTCResponse{}, err
		}
		return RunFaucetBTCResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for RunFaucetBTC.
type RunFaucetBTCResponse RunFaucetResponse

// Search blocks, transactions, contracts, or accounts by hash/ID
func SearchByID(server Server, params SearchByIDParams) (SearchByIDResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/search/{ID}", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return SearchByIDResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj SearchByIDResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return SearchByIDResponse{}, err
		}
		return SearchByIDResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for SearchByID.
type SearchByIDResponse SearchResult

// Defines parameters for SearchByID
type SearchByIDParams struct {
	// This includes the detailed data for purticular hash in the response.
	// Optional.
	IncludeMetadata bool
	// The hex hash string for a block or transaction, account address, or contract address.
	Id string
}

// Defines a 200 (Success) response for GetFeeTransfer.
type GetFeeTransferResponse string

// Retrieves an estimated fee rate for STX transfer transactions. This a a fee rate / byte, and is returned as a JSON integer
func GetFeeTransfer(server Server) (GetFeeTransferResponse, error) {
	url := fmt.Sprintf("%s%s", server, "/v2/fees/transfer")
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetFeeTransferResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetFeeTransferResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetFeeTransferResponse{}, err
		}
		return GetFeeTransferResponse{}, errors.New(respObj.Error)
	}
}

// Retrieves a list of all assets events associated with an account or a Contract Identifier. This includes Transfers, Mints.
func GetAccountAssets(server Server, params GetAccountAssetsParams) (GetAccountAssetsResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/address/{principal}/assets", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetAccountAssetsResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetAccountAssetsResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetAccountAssetsResponse{}, err
		}
		return GetAccountAssetsResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for GetAccountAssets.
type GetAccountAssetsResponse AddressAssetsListResponse

// Defines parameters for GetAccountAssets
type GetAccountAssetsParams struct {
	// Index of first account assets to fetch.
	// Optional.
	Offset int
	// Returned data representing the state at that point in time, rather than the current block.
	// Optional.
	UntilBlock string
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. The default value is false.
	Unanchored bool
	// Max number of account assets to fetch.
	// Optional.
	Limit int
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
}

/*
Retrieves a list of STX transfers with memos to the given principal. This includes regular transfers from a stx-transfer transaction type,
and transfers from contract-call transactions a the `send-many-memo` bulk sending contract.
*/
func GetAccountInbound(server Server, params GetAccountInboundParams) (GetAccountInboundResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/address/{principal}/stx_inbound", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetAccountInboundResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetAccountInboundResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetAccountInboundResponse{}, err
		}
		return GetAccountInboundResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for GetAccountInbound.
type GetAccountInboundResponse AddressStxInboundListResponse

// Defines parameters for GetAccountInbound
type GetAccountInboundParams struct {
	// Filter for transfers only at this given block height.
	// Optional.
	Height int
	// Number of items to skip.
	// Optional.
	Offset int
	// Number of items to return.
	// Optional.
	Limit int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. The default value is false.
	Unanchored bool
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
	// Returned data representing the state up until that point in time, rather than the current block.
	// Optional.
	UntilBlock string
}

// Retrieve all transactions for an account or contract identifier including STX transfers for each transaction.
func TransactionsWithTransfers(server Server, params TransactionsWithTransfersParams) (TransactionsWithTransfersResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/address/{principal}/transactions_with_transfers", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return TransactionsWithTransfersResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj TransactionsWithTransfersResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return TransactionsWithTransfersResponse{}, err
		}
		return TransactionsWithTransfersResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for TransactionsWithTransfers.
type TransactionsWithTransfersResponse AddressTransactionsWithTransfersListResponse

// Defines parameters for TransactionsWithTransfers
type TransactionsWithTransfersParams struct {
	// Filter for transactions only at this given block height.
	// Optional.
	Height int
	// Index of first account transaction to fetch.
	// Optional.
	Offset int
	// Returned data representing the state up until that point in time, rather than the current block.
	// Optional.
	UntilBlock string
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. The default value is false.
	Unanchored bool
	// Max number of account transactions to fetch.
	// Optional.
	Limit int
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
}

// Retrieves a list of recent burnchain (e.g. Bitcoin) rewards for the given recipient with the associated amounts and block info
func GetBurnchainRewardListByAddress(server Server, params GetBurnchainRewardListByAddressParams) (GetBurnchainRewardListByAddressResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/burnchain/rewards/{address}", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetBurnchainRewardListByAddressResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetBurnchainRewardListByAddressResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetBurnchainRewardListByAddressResponse{}, err
		}
		return GetBurnchainRewardListByAddressResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (List of burnchain reward recipients and amounts) response for GetBurnchainRewardListByAddress.
type GetBurnchainRewardListByAddressResponse BurnchainRewardListResponse

// Defines parameters for GetBurnchainRewardListByAddress
type GetBurnchainRewardListByAddressParams struct {
	// Max number of rewards to fetch.
	// Optional.
	Limit int
	// Index of first rewards to fetch.
	// Optional.
	Offset int
	// Reward recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format.
	Address string
}

// Retrieves a list of all names known to the node.
func GetAllNames(server Server, params GetAllNamesParams) (GetAllNamesResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/v1/names", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetAllNamesResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetAllNamesResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetAllNamesResponse{}, err
		}
		return GetAllNamesResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for GetAllNames.
type GetAllNamesResponse []BnsGetAllNamesResponse

// Defines a 400 (Error) response for GetAllNames.
type GetAllNamesError400 BnsError

// Defines parameters for GetAllNames
type GetAllNamesParams struct {
	// Names are returned in pages of size 100, so specify the page number.
	Page int
}
type CallReadOnlyFunctionBody struct{}

/*
Call a read-only public function on a given smart contract.

The smart contract and function are specified using the URL path. The arguments and the simulated tx-sender are supplied via the POST body in the following JSON format:
*/
func CallReadOnlyFunction(server Server, params CallReadOnlyFunctionParams, body CallReadOnlyFunctionBody) (CallReadOnlyFunctionResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/v2/contracts/call-read/{contract_address}/{contract_name}/{function_name}", params))
	var returnedErr error
	sendBody, _ := json.Marshal(body)
	resp, status := makePostReq(url, string(sendBody), &returnedErr)

	if returnedErr != nil {
		return CallReadOnlyFunctionResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj CallReadOnlyFunctionResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return CallReadOnlyFunctionResponse{}, err
		}
		return CallReadOnlyFunctionResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for CallReadOnlyFunction.
type CallReadOnlyFunctionResponse ReadOnlyFunctionSuccessResponse

// Defines parameters for CallReadOnlyFunction
type CallReadOnlyFunctionParams struct {
	// The Stacks chain tip to query from.
	// Optional.
	Tip string
	// Contract name.
	ContractName string
	// Stacks address.
	ContractAddress string
	// Function name.
	FunctionName string
}
type PostCoreNodeTransactionsBody struct{}

// Broadcasts raw transactions on the network. You can use the [@stacks/transactions](https://github.com/blockstack/stacks.js) project to generate a raw transaction payload.
func PostCoreNodeTransactions(server Server, body PostCoreNodeTransactionsBody) (PostCoreNodeTransactionsResponse, error) {
	url := fmt.Sprintf("%s%s", server, "/v2/transactions")
	var returnedErr error
	sendBody, _ := json.Marshal(body)
	resp, status := makePostReq(url, string(sendBody), &returnedErr)

	if returnedErr != nil {
		return PostCoreNodeTransactionsResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj PostCoreNodeTransactionsResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return PostCoreNodeTransactionsResponse{}, err
		}
		return PostCoreNodeTransactionsResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Transaction id of successful post of a raw tx to the node's mempool) response for PostCoreNodeTransactions.
type PostCoreNodeTransactionsResponse string

// Defines a 400 (Rejections result in a 400 error) response for PostCoreNodeTransactions.
type PostCoreNodeTransactionsError400 PostCoreNodeTransactionsError

// Retrieves STX token balance for a given Address or Contract Identifier.
func GetAccountSTXBalance(server Server, params GetAccountSTXBalanceParams) (GetAccountSTXBalanceResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/address/{principal}/stx", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetAccountSTXBalanceResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetAccountSTXBalanceResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetAccountSTXBalanceResponse{}, err
		}
		return GetAccountSTXBalanceResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for GetAccountSTXBalance.
type GetAccountSTXBalanceResponse AddressStxBalanceResponse

// Defines parameters for GetAccountSTXBalance
type GetAccountSTXBalanceParams struct {
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. The default value is false.
	Unanchored bool
	// Returned data representing the state up until that point in time, rather than the current block.
	// Optional.
	UntilBlock string
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
}

// Retrieves a list of recent burnchain (e.g. Bitcoin) reward recipients with the associated amounts and block info
func GetBurnchainRewardList(server Server, params GetBurnchainRewardListParams) (GetBurnchainRewardListResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/burnchain/rewards", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetBurnchainRewardListResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetBurnchainRewardListResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetBurnchainRewardListResponse{}, err
		}
		return GetBurnchainRewardListResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (List of burnchain reward recipients and amounts) response for GetBurnchainRewardList.
type GetBurnchainRewardListResponse BurnchainRewardListResponse

// Defines parameters for GetBurnchainRewardList
type GetBurnchainRewardListParams struct {
	// Max number of rewards to fetch.
	// Optional. The default value is 96.
	// Max value is 250.
	Limit int
	// Index of first rewards to fetch.
	// Optional.
	Offset int
}

// Retrieves the total supply for STX tokens as plain text.
func GetSTXSupplyTotalSupplyPlain(server Server) (GetSTXSupplyTotalSupplyPlainResponse, error) {
	url := fmt.Sprintf("%s%s", server, "/extended/v1/stx_supply/total/plain")
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetSTXSupplyTotalSupplyPlainResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetSTXSupplyTotalSupplyPlainResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetSTXSupplyTotalSupplyPlainResponse{}, err
		}
		return GetSTXSupplyTotalSupplyPlainResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (success) response for GetSTXSupplyTotalSupplyPlain.
type GetSTXSupplyTotalSupplyPlainResponse string

// Retrieves a list of names owned by the address provided.
func GetNamesOwnedByAddress(server Server, params GetNamesOwnedByAddressParams) (GetNamesOwnedByAddressResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/v1/addresses/{blockchain}/{address}", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetNamesOwnedByAddressResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetNamesOwnedByAddressResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetNamesOwnedByAddressResponse{}, err
		}
		return GetNamesOwnedByAddressResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 404 (Error) response for GetNamesOwnedByAddress.
type GetNamesOwnedByAddressError404 BnsError

// Defines a 200 (Success) response for GetNamesOwnedByAddress.
type GetNamesOwnedByAddressResponse BnsNamesOwnByAddressResponse

// Defines parameters for GetNamesOwnedByAddress
type GetNamesOwnedByAddressParams struct {
	// The address to lookup.
	Address string
	// The layer-1 blockchain for the address.
	Blockchain string
}
type FetchFeeRateBody struct{}

/*
**NOTE:** This endpoint is deprecated in favor of [Get approximate fees for the given transaction](#operation/post_fee_transaction).

Retrieves estimated fee rate.
*/
func FetchFeeRate(server Server, body FetchFeeRateBody) (FetchFeeRateResponse, error) {
	url := fmt.Sprintf("%s%s", server, "/extended/v1/fee_rate")
	var returnedErr error
	sendBody, _ := json.Marshal(body)
	resp, status := makePostReq(url, string(sendBody), &returnedErr)

	if returnedErr != nil {
		return FetchFeeRateResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj FetchFeeRateResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return FetchFeeRateResponse{}, err
		}
		return FetchFeeRateResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Transaction fee rate) response for FetchFeeRate.
type FetchFeeRateResponse FeeRate

/*
Retrieves a list of microblocks.

If you need to actively monitor new microblocks, we highly recommend subscribing to [WebSockets or Socket.io](https://github.com/hirosystems/stacks-blockchain-api/tree/master/client) for real-time updates.
*/
func GetMicroblockList(server Server, params GetMicroblockListParams) (GetMicroblockListResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/microblock", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetMicroblockListResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetMicroblockListResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetMicroblockListResponse{}, err
		}
		return GetMicroblockListResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (List of microblocks) response for GetMicroblockList.
type GetMicroblockListResponse MicroblockListResponse

// Defines parameters for GetMicroblockList
type GetMicroblockListParams struct {
	// Index of the first microblock to fetch.
	// Optional.
	Offset int
	// Max number of microblocks to fetch.
	// Optional. The default value is 20.
	// Max value is 200.
	Limit int
}

/*
Retrieves the list of Non-Fungible Tokens owned by the given principal (STX address or Smart Contract ID).
Results can be filtered by one or more asset identifiers and can include metadata about the transaction that made the principal own each token.

More information on Non-Fungible Tokens on the Stacks blockchain can be found [here](https://docs.stacks.co/write-smart-contracts/tokens#non-fungible-tokens-nfts).
*/
func GetNFTHoldings(server Server, params GetNFTHoldingsParams) (GetNFTHoldingsResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/tokens/nft/holdings", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetNFTHoldingsResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetNFTHoldingsResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetNFTHoldingsResponse{}, err
		}
		return GetNFTHoldingsResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (List of Non-Fungible Token holdings) response for GetNFTHoldings.
type GetNFTHoldingsResponse NonFungibleTokenHoldingsList

// Defines parameters for GetNFTHoldings
type GetNFTHoldingsParams struct {
	// Token owner's STX address or Smart Contract ID.
	Principal string
	// Index of first tokens to fetch.
	// Optional. The default value is 0.
	Offset int
	// Whether or not to include tokens from unconfirmed transactions.
	// Optional. The default value is false.
	Unanchored bool
	// Whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times.
	// Optional. The default value is false.
	TxMetadata bool
	// Max number of tokens to fetch.
	// Optional. The default value is 50.
	Limit int
	// Identifiers of the token asset classes to filter for.
	// Optional.
	AssetIdentifiers []GetNFTHoldingsAssetIdentifiers
}

/*
Retrieves a list of transactions for a given list of transaction IDs

If using TypeScript, import typings for this response from our types package:

`import type { Transaction } from '@stacks/stacks-blockchain-api-types';`
*/
func GetTXListDetails(server Server) (GetTXListDetailsResponse, error) {
	url := fmt.Sprintf("%s%s", server, "/extended/v1/tx/multiple")
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetTXListDetailsResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetTXListDetailsResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetTXListDetailsResponse{}, err
		}
		return GetTXListDetailsResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Returns list of transactions with their details for corresponding requested tx_ids.) response for GetTXListDetails.
type GetTXListDetailsResponse TransactionList

// Retrieves block details of a specific block for a given burn chain height
func GetBlockByBurnBlockHeight(server Server) (GetBlockByBurnBlockHeightResponse, error) {
	url := fmt.Sprintf("%s%s", server, "/extended/v1/block/by_burn_block_height/{burn_block_height}")
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetBlockByBurnBlockHeightResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetBlockByBurnBlockHeightResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetBlockByBurnBlockHeightResponse{}, err
		}
		return GetBlockByBurnBlockHeightResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Block) response for GetBlockByBurnBlockHeight.
type GetBlockByBurnBlockHeightResponse Block

// Retrieves block details of a specific block at a given block height
func GetBlockByHeight(server Server) (GetBlockByHeightResponse, error) {
	url := fmt.Sprintf("%s%s", server, "/extended/v1/block/by_height/{height}")
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetBlockByHeightResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetBlockByHeightResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetBlockByHeightResponse{}, err
		}
		return GetBlockByHeightResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Block) response for GetBlockByHeight.
type GetBlockByHeightResponse Block

// Retrieves a list of the Bitcoin addresses that would validly receive Proof-of-Transfer commitments.
func GetBurnchainRewardSlotHolders(server Server, params GetBurnchainRewardSlotHoldersParams) (GetBurnchainRewardSlotHoldersResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/burnchain/reward_slot_holders", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetBurnchainRewardSlotHoldersResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetBurnchainRewardSlotHoldersResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetBurnchainRewardSlotHoldersResponse{}, err
		}
		return GetBurnchainRewardSlotHoldersResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (List of burnchain reward recipients and amounts) response for GetBurnchainRewardSlotHolders.
type GetBurnchainRewardSlotHoldersResponse BurnchainRewardSlotHolderListResponse

// Defines parameters for GetBurnchainRewardSlotHolders
type GetBurnchainRewardSlotHoldersParams struct {
	// Index of the first items to fetch.
	// Optional.
	Offset int
	// Max number of items to fetch.
	// Optional. The default value is 96.
	// Max value is 250.
	Limit int
}
type RunFaucetSTXBody struct{}

/*
Add 500 STX tokens to the specified testnet address. Testnet STX addresses begin with `ST`. If the `stacking`
parameter is set to `true`, the faucet will add the required number of tokens for individual stacking to the
specified testnet address.

The endpoint returns the transaction ID, which you can use to view the transaction in the
[Stacks Explorer](https://explorer.stacks.co/?chain=testnet). The tokens are delivered once the transaction has
been included in an anchor block.

A common reason for failed faucet transactions is that the faucet has run out of tokens. If you are experiencing
failed faucet transactions to a testnet address, you can get help in [Discord](https://stacks.chat).

**Note:** This is a testnet only endpoint. This endpoint will not work on the mainnet.
*/
func RunFaucetSTX(server Server, body RunFaucetSTXBody) (RunFaucetSTXResponse, error) {
	url := fmt.Sprintf("%s%s", server, "/extended/v1/faucets/stx")
	var returnedErr error
	sendBody, _ := json.Marshal(body)
	resp, status := makePostReq(url, string(sendBody), &returnedErr)

	if returnedErr != nil {
		return RunFaucetSTXResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj RunFaucetSTXResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return RunFaucetSTXResponse{}, err
		}
		return RunFaucetSTXResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for RunFaucetSTX.
type RunFaucetSTXResponse RunFaucetResponse

/*
Retrieves transaction details for a given transaction ID

`import type { Transaction } from '@stacks/stacks-blockchain-api-types';`
*/
func GetTransactionByID(server Server) (GetTransactionByIDResponse, error) {
	url := fmt.Sprintf("%s%s", server, "/extended/v1/tx/{TX_ID}")
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetTransactionByIDResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetTransactionByIDResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetTransactionByIDResponse{}, err
		}
		return GetTransactionByIDResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Transaction) response for GetTransactionByID.
type GetTransactionByIDResponse Transaction

// Defines a 200 (Success) response for FetchZoneFile.
type FetchZoneFileResponse any

// Retrieves a user’s raw zone file. This only works for RFC-compliant zone files. This method returns an error for names that have non-standard zone files.
func FetchZoneFile(server Server, params FetchZoneFileParams) (FetchZoneFileResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/v1/names/{name}/zonefile", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return FetchZoneFileResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj FetchZoneFileResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return FetchZoneFileResponse{}, err
		}
		return FetchZoneFileResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 400 (Error) response for FetchZoneFile.
type FetchZoneFileError400 BnsError

// Defines a 404 (Error) response for FetchZoneFile.
type FetchZoneFileError404 BnsError

// Defines parameters for FetchZoneFile
type FetchZoneFileParams struct {
	// Fully-qualified name.
	Name string
}

/*
Retrieves the account data for a given Account or a Contract Identifier

Where balance is the hex encoding of a unsigned 128-bit integer (big-endian), nonce is an unsigned 64-bit integer, and the proofs are provided as hex strings.

For non-existent accounts, this does not return a 404 error, rather it returns an object with balance and nonce of 0.
*/
func GetAccountInfo(server Server, params GetAccountInfoParams) (GetAccountInfoResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/v2/accounts/{principal}", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetAccountInfoResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetAccountInfoResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetAccountInfoResponse{}, err
		}
		return GetAccountInfoResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for GetAccountInfo.
type GetAccountInfoResponse AccountDataResponse

// Defines parameters for GetAccountInfo
type GetAccountInfoParams struct {
	// Returns object without the proof field if set to 0.
	// Optional.
	Proof int
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
	// The Stacks chain tip to query from.
	// Optional.
	Tip string
}

/*
**NOTE:** This endpoint is deprecated in favor of [Non-Fungible Token holdings](#operation/get_nft_holdings).

Retrieves a list of all nfts owned by an address, contains the clarity value of the identifier of the nft.
*/
func GetAccountNFT(server Server, params GetAccountNFTParams) (GetAccountNFTResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/address/{principal}/nft_events", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetAccountNFTResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetAccountNFTResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetAccountNFTResponse{}, err
		}
		return GetAccountNFTResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for GetAccountNFT.
type GetAccountNFTResponse AddressNftListResponse

// Defines parameters for GetAccountNFT
type GetAccountNFTParams struct {
	// Number of items to skip.
	// Optional.
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. The default value is false.
	Unanchored bool
	// Returned data representing the state up until that point in time, rather than the current block.
	// Optional.
	UntilBlock string
	// Number of items to return.
	// Optional.
	Limit int
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
}

// Retrieves transactions that have been streamed in microblocks but not yet accepted or rejected in an anchor block
func GetUnanchoredTxs(server Server) (GetUnanchoredTxsResponse, error) {
	url := fmt.Sprintf("%s%s", server, "/extended/v1/microblock/unanchored/txs")
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetUnanchoredTxsResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetUnanchoredTxsResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetUnanchoredTxsResponse{}, err
		}
		return GetUnanchoredTxsResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Transactions) response for GetUnanchoredTxs.
type GetUnanchoredTxsResponse UnanchoredTransactionListResponse

// Retrieves details of a given name including the `address`, `status` and last transaction id - `last_txid`.
func GetNameInfo(server Server, params GetNameInfoParams) (GetNameInfoResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/v1/names/{name}", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetNameInfoResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetNameInfoResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetNameInfoResponse{}, err
		}
		return GetNameInfoResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 400 (Error) response for GetNameInfo.
type GetNameInfoError400 BnsError

// Defines a 404 (Error) response for GetNameInfo.
type GetNameInfoError404 BnsError

// Defines a 200 (Success) response for GetNameInfo.
type GetNameInfoResponse BnsGetNameInfoResponse

// Defines parameters for GetNameInfo
type GetNameInfoParams struct {
	// Fully-qualified name.
	Name string
}

// Defines a 200 (Success) response for GetHistoricalZoneFile.
type GetHistoricalZoneFileResponse any

// Retrieves the historical zonefile specified by the username and zone hash.
func GetHistoricalZoneFile(server Server, params GetHistoricalZoneFileParams) (GetHistoricalZoneFileResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/v1/names/{name}/zonefile/{zoneFileHash}", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetHistoricalZoneFileResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetHistoricalZoneFileResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetHistoricalZoneFileResponse{}, err
		}
		return GetHistoricalZoneFileResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 400 (Error) response for GetHistoricalZoneFile.
type GetHistoricalZoneFileError400 BnsError

// Defines a 404 (Error) response for GetHistoricalZoneFile.
type GetHistoricalZoneFileError404 BnsError

// Defines parameters for GetHistoricalZoneFile
type GetHistoricalZoneFileParams struct {
	// Fully-qualified name.
	Name string
	// Zone file hash.
	ZoneFileHash string
}

// Retrieves a list of the Bitcoin addresses that would validly receive Proof-of-Transfer commitments for a given reward slot holder recipient address.
func GetBurnchainRewardSlotHoldersByAddress(server Server, params GetBurnchainRewardSlotHoldersByAddressParams) (GetBurnchainRewardSlotHoldersByAddressResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/burnchain/reward_slot_holders/{address}", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetBurnchainRewardSlotHoldersByAddressResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetBurnchainRewardSlotHoldersByAddressResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetBurnchainRewardSlotHoldersByAddressResponse{}, err
		}
		return GetBurnchainRewardSlotHoldersByAddressResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (List of burnchain reward recipients and amounts) response for GetBurnchainRewardSlotHoldersByAddress.
type GetBurnchainRewardSlotHoldersByAddressResponse BurnchainRewardSlotHolderListResponse

// Defines parameters for GetBurnchainRewardSlotHoldersByAddress
type GetBurnchainRewardSlotHoldersByAddressParams struct {
	// Max number of items to fetch.
	// Optional.
	Limit int
	// Index of the first items to fetch.
	// Optional.
	Offset int
	// Reward slot holder recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format.
	Address string
}

/*
Retrieves all mint events for a Non-Fungible Token asset class. Useful to determine which NFTs of a particular collection have been claimed.

More information on Non-Fungible Tokens on the Stacks blockchain can be found [here](https://docs.stacks.co/write-smart-contracts/tokens#non-fungible-tokens-nfts).
*/
func GetNFTMints(server Server, params GetNFTMintsParams) (GetNFTMintsResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/tokens/nft/mints", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetNFTMintsResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetNFTMintsResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetNFTMintsResponse{}, err
		}
		return GetNFTMintsResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Non-Fungible Token mints) response for GetNFTMints.
type GetNFTMintsResponse NonFungibleTokenMintList

// Defines parameters for GetNFTMints
type GetNFTMintsParams struct {
	// Index of first event to fetch.
	// Optional. The default value is 0.
	Offset int
	// Token asset class identifier.
	AssetIdentifier string
	// Max number of events to fetch.
	// Optional. The default value is 50.
	Limit int
	// Whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times.
	// Optional. The default value is false.
	TxMetadata bool
	// Whether or not to include events from unconfirmed transactions.
	// Optional. The default value is false.
	Unanchored bool
}

// Retrieves a list of all transactions within a block for a given block hash.
func GetTransactionsByBlockHash(server Server, params GetTransactionsByBlockHashParams) (GetTransactionsByBlockHashResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/tx/block/{block_hash}", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetTransactionsByBlockHashResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetTransactionsByBlockHashResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetTransactionsByBlockHashResponse{}, err
		}
		return GetTransactionsByBlockHashResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (List of Transactions) response for GetTransactionsByBlockHash.
type GetTransactionsByBlockHashResponse TransactionResults

// Defines parameters for GetTransactionsByBlockHash
type GetTransactionsByBlockHashParams struct {
	// Index of first transaction to fetch.
	// Optional.
	Offset int
	// Hash of block.
	BlockHash string
	// Max number of transactions to fetch.
	// Optional.
	Limit int
}

// Retrieves information about the Core API including the server version
func GetCoreAPIInfo(server Server) (GetCoreAPIInfoResponse, error) {
	url := fmt.Sprintf("%s%s", server, "/v2/info")
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetCoreAPIInfoResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetCoreAPIInfoResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetCoreAPIInfoResponse{}, err
		}
		return GetCoreAPIInfoResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for GetCoreAPIInfo.
type GetCoreAPIInfoResponse CoreNodeInfoResponse

// Retrieves Proof-of-Transfer (PoX) information. Can be used for Stacking.
func GetPoxInfo(server Server) (GetPoxInfoResponse, error) {
	url := fmt.Sprintf("%s%s", server, "/v2/pox")
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetPoxInfoResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetPoxInfoResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetPoxInfoResponse{}, err
		}
		return GetPoxInfoResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for GetPoxInfo.
type GetPoxInfoResponse CoreNodePoxResponse

/*
Retrieves a list of recently mined blocks

If you need to actively monitor new blocks, we highly recommend subscribing to [WebSockets or Socket.io](https://github.com/hirosystems/stacks-blockchain-api/tree/master/client) for real-time updates.
*/
func GetBlockList(server Server, params GetBlockListParams) (GetBlockListResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/block", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetBlockListResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetBlockListResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetBlockListResponse{}, err
		}
		return GetBlockListResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (List of blocks) response for GetBlockList.
type GetBlockListResponse BlockListResponse

// Defines parameters for GetBlockList
type GetBlockListParams struct {
	// Index of first block to fetch.
	// Optional.
	Offset int
	// Max number of blocks to fetch.
	// Optional. The default value is 20.
	// Max value is 200.
	Limit int
}

// Retrieves the total and unlocked STX supply. More information on Stacking can be found [here] (https://docs.stacks.co/understand-stacks/stacking).
func GetSTXSupply(server Server, params GetSTXSupplyParams) (GetSTXSupplyResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/stx_supply", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetSTXSupplyResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetSTXSupplyResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetSTXSupplyResponse{}, err
		}
		return GetSTXSupplyResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for GetSTXSupply.
type GetSTXSupplyResponse GetStxSupplyResponse

// Defines parameters for GetSTXSupply
type GetSTXSupplyParams struct {
	// The block height at which to query supply details from, if not provided then the latest block height is used.
	// Optional.
	Height int
}

// Retrieves a list of all namespaces known to the node.
func GetAllNamespaces(server Server) (GetAllNamespacesResponse, error) {
	url := fmt.Sprintf("%s%s", server, "/v1/namespaces")
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetAllNamespacesResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetAllNamespacesResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetAllNamespacesResponse{}, err
		}
		return GetAllNamespacesResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for GetAllNamespaces.
type GetAllNamespacesResponse BnsGetAllNamespacesResponse

// Retrieves a specific microblock by `hash`
func GetMicroblockByHash(server Server) (GetMicroblockByHashResponse, error) {
	url := fmt.Sprintf("%s%s", server, "/extended/v1/microblock/{hash}")
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetMicroblockByHashResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetMicroblockByHashResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetMicroblockByHashResponse{}, err
		}
		return GetMicroblockByHashResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Microblock) response for GetMicroblockByHash.
type GetMicroblockByHashResponse Microblock

// Retrieves all transactions within a block at a given height
func GetTransactionsByBlockHeight(server Server, params GetTransactionsByBlockHeightParams) (GetTransactionsByBlockHeightResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/tx/block_height/{height}", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetTransactionsByBlockHeightResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetTransactionsByBlockHeightResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetTransactionsByBlockHeightResponse{}, err
		}
		return GetTransactionsByBlockHeightResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (List of Transactions) response for GetTransactionsByBlockHeight.
type GetTransactionsByBlockHeightResponse TransactionResults

// Defines parameters for GetTransactionsByBlockHeight
type GetTransactionsByBlockHeightParams struct {
	// Max number of transactions to fetch.
	// Optional.
	Limit int
	// Index of first transaction to fetch.
	// Optional.
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. The default value is false.
	Unanchored bool
	// Height of block.
	Height int
}

// Retrieves all transactions for a given address that are currently in mempool
func GetAddressMempoolTransactions(server Server, params GetAddressMempoolTransactionsParams) (GetAddressMempoolTransactionsResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/address/{address}/mempool", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetAddressMempoolTransactionsResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetAddressMempoolTransactionsResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetAddressMempoolTransactionsResponse{}, err
		}
		return GetAddressMempoolTransactionsResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (List of Transactions) response for GetAddressMempoolTransactions.
type GetAddressMempoolTransactionsResponse MempoolTransactionListResponse

// Defines parameters for GetAddressMempoolTransactions
type GetAddressMempoolTransactionsParams struct {
	// Index of first transaction to fetch.
	// Optional.
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. The default value is false.
	Unanchored bool
	// Transactions for the address.
	Address string
	// Max number of transactions to fetch.
	// Optional.
	Limit int
}

// Retrieves transaction details for a given Transcation Id `tx_id`, for a given account or contract Identifier.
func GetSingleTransactionWithTransfers(server Server, params GetSingleTransactionWithTransfersParams) (GetSingleTransactionWithTransfersResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/address/{principal}/{TX_ID}/with_transfers", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetSingleTransactionWithTransfersResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetSingleTransactionWithTransfersResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetSingleTransactionWithTransfersResponse{}, err
		}
		return GetSingleTransactionWithTransfersResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for GetSingleTransactionWithTransfers.
type GetSingleTransactionWithTransfersResponse AddressTransactionWithTransfers

// Defines parameters for GetSingleTransactionWithTransfers
type GetSingleTransactionWithTransfersParams struct {
	// Stacks address or a contract identifier.
	Principal string
	// Transaction id.
	TxID string
}
type PostFeeTransactionBody struct{}

/*
Get an estimated fee for the supplied transaction.  This
estimates the execution cost of the transaction, the current
fee rate of the network, and returns estimates for fee
amounts.
* `transaction_payload` is a hex-encoded serialization of
  the TransactionPayload for the transaction.
* `estimated_len` is an optional argument that provides the
  endpoint with an estimation of the final length (in bytes)
  of the transaction, including any post-conditions and
  signatures
If the node cannot provide an estimate for the transaction
(e.g., if the node has never seen a contract-call for the
given contract and function) or if estimation is not
configured on this node, a 400 response is returned.
The 400 response will be a JSON error containing a `reason`
field which can be one of the following:
* `DatabaseError` - this Stacks node has had an internal
  database error while trying to estimate the costs of the
  supplied transaction.
* `NoEstimateAvailable` - this Stacks node has not seen this
  kind of contract-call before, and it cannot provide an
  estimate yet.
* `CostEstimationDisabled` - this Stacks node does not perform
  fee or cost estimation, and it cannot respond on this
  endpoint.
The 200 response contains the following data:
* `estimated_cost` - the estimated multi-dimensional cost of
  executing the Clarity VM on the provided transaction.
* `estimated_cost_scalar` - a unitless integer that the Stacks
  node uses to compare how much of the block limit is consumed
  by different transactions. This value incorporates the
  estimated length of the transaction and the estimated
  execution cost of the transaction. The range of this integer
  may vary between different Stacks nodes. In order to compute
  an estimate of total fee amount for the transaction, this
  value is multiplied by the same Stacks node's estimated fee
  rate.
* `cost_scalar_change_by_byte` - a float value that indicates how
  much the `estimated_cost_scalar` value would increase for every
  additional byte in the final transaction.
* `estimations` - an array of estimated fee rates and total fees to
  pay in microSTX for the transaction. This array provides a range of
  estimates (default: 3) that may be used. Each element of the array
  contains the following fields:
    * `fee_rate` - the estimated value for the current fee
      rates in the network
    * `fee` - the estimated value for the total fee in
      microSTX that the given transaction should pay. These
      values are the result of computing:
      `fee_rate` x `estimated_cost_scalar`.
      If the estimated fees are less than the minimum relay
      fee `(1 ustx x estimated_len)`, then that minimum relay
      fee will be returned here instead.
Note: If the final transaction's byte size is larger than
supplied to `estimated_len`, then applications should increase
this fee amount by:
  `fee_rate` x `cost_scalar_change_by_byte` x (`final_size` - `estimated_size`)
*/
func PostFeeTransaction(server Server, body PostFeeTransactionBody) (PostFeeTransactionResponse, error) {
	url := fmt.Sprintf("%s%s", server, "/v2/fees/transaction")
	var returnedErr error
	sendBody, _ := json.Marshal(body)
	resp, status := makePostReq(url, string(sendBody), &returnedErr)

	if returnedErr != nil {
		return PostFeeTransactionResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj PostFeeTransactionResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return PostFeeTransactionResponse{}, err
		}
		return PostFeeTransactionResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Estimated fees for the transaction) response for PostFeeTransaction.
type PostFeeTransactionResponse TransactionFeeEstimateResponse

/*
Retrieves all transactions that have been recently broadcast to the mempool. These are pending transactions awaiting confirmation.

If you need to monitor new transactions, we highly recommend subscribing to [WebSockets or Socket.io](https://github.com/hirosystems/stacks-blockchain-api/tree/master/client) for real-time updates.
*/
func GetMempoolTransactionList(server Server, params GetMempoolTransactionListParams) (GetMempoolTransactionListResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/tx/mempool", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetMempoolTransactionListResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetMempoolTransactionListResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetMempoolTransactionListResponse{}, err
		}
		return GetMempoolTransactionListResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (List of mempool transactions) response for GetMempoolTransactionList.
type GetMempoolTransactionListResponse MempoolTransactionListResponse

// Defines parameters for GetMempoolTransactionList
type GetMempoolTransactionListParams struct {
	// Max number of mempool transactions to fetch.
	// Optional. The default value is 96.
	// Max value is 200.
	Limit int
	// Filter to only return transactions with this sender address.
	// Optional.
	SenderAddress string
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. The default value is false.
	Unanchored bool
	// Index of first mempool transaction to fetch.
	// Optional.
	Offset int
	// Filter to only return transactions with this recipient address (only applicable for STX transfer tx types).
	// Optional.
	RecipientAddress string
	// Filter to only return transactions with this address as the sender or recipient (recipient only applicable for STX transfer tx types).
	// Optional.
	Address string
}

/*
Retrieves all recently-broadcast transactions that have been dropped from the mempool.

Transactions are dropped from the mempool if:
 * they were stale and awaiting garbage collection or,
 * were expensive,  or
 * were replaced with a new fee
*/
func GetDroppedMempoolTransactionList(server Server, params GetDroppedMempoolTransactionListParams) (GetDroppedMempoolTransactionListResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/tx/mempool/dropped", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetDroppedMempoolTransactionListResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetDroppedMempoolTransactionListResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetDroppedMempoolTransactionListResponse{}, err
		}
		return GetDroppedMempoolTransactionListResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (List of dropped mempool transactions) response for GetDroppedMempoolTransactionList.
type GetDroppedMempoolTransactionListResponse MempoolTransactionListResponse

// Defines parameters for GetDroppedMempoolTransactionList
type GetDroppedMempoolTransactionListParams struct {
	// Max number of mempool transactions to fetch.
	// Optional. The default value is 96.
	// Max value is 200.
	Limit int
	// Index of first mempool transaction to fetch.
	// Optional.
	Offset int
}

// Retrieves total account balance information for a given Address or Contract Identifier. This includes the balances of  STX Tokens, Fungible Tokens and Non-Fungible Tokens for the account.
func GetAccountBalance(server Server, params GetAccountBalanceParams) (GetAccountBalanceResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/address/{principal}/balances", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetAccountBalanceResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetAccountBalanceResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetAccountBalanceResponse{}, err
		}
		return GetAccountBalanceResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for GetAccountBalance.
type GetAccountBalanceResponse AddressBalanceResponse

// Defines parameters for GetAccountBalance
type GetAccountBalanceParams struct {
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. The default value is false.
	Unanchored bool
	// Returned data representing the state up until that point in time, rather than the current block.
	// Optional.
	UntilBlock string
	// Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`).
	Principal string
}

// Retrieves the latest nonce values used by an account by inspecting the mempool, microblock transactions, and anchored transactions.
func GetAccountNonces(server Server, params GetAccountNoncesParams) (GetAccountNoncesResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/address/{principal}/nonces", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetAccountNoncesResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetAccountNoncesResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetAccountNoncesResponse{}, err
		}
		return GetAccountNoncesResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for GetAccountNonces.
type GetAccountNoncesResponse AddressNonces

// Defines parameters for GetAccountNonces
type GetAccountNoncesParams struct {
	// Stacks address (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0`).
	Principal string
	// Optionally get the nonce at a given block height.
	// Optional.
	BlockHeight int
	// Optionally get the nonce at a given block hash.
	// Optional.
	BlockHash string
}

// Retrieves a list of events that have been triggered by a given `contract_id`
func GetContractEventsByID(server Server, params GetContractEventsByIDParams) (GetContractEventsByIDResponse, error) {
	url := fmt.Sprintf("%s%s", server, fillPath("/extended/v1/contract/{contract_ID}/events", params))
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetContractEventsByIDResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetContractEventsByIDResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetContractEventsByIDResponse{}, err
		}
		return GetContractEventsByIDResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (List of events) response for GetContractEventsByID.
type GetContractEventsByIDResponse TransactionEvent

// Defines parameters for GetContractEventsByID
type GetContractEventsByIDParams struct {
	// Index of first contract event to fetch.
	// Optional.
	Offset int
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. The default value is false.
	Unanchored bool
	// Max number of contract events to fetch.
	// Optional.
	Limit int
	// Contract identifier formatted as `<contract_address>.<contract_name>`.
	ContractID string
}

// Retrieves the running status of the Stacks Blockchain API, including the server version and current chain tip information.
func GetStatus(server Server) (GetStatusResponse, error) {
	url := fmt.Sprintf("%s%s", server, "/extended/v1/status")
	var returnedErr error
	resp, status := makeGetReq(url, &returnedErr)

	if returnedErr != nil {
		return GetStatusResponse{}, returnedErr
	}

	switch status {
	case 200:
		var respObj GetStatusResponse
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return respObj, err
		}
		return respObj, nil
	default:
		var respObj struct {
			Error string `json:"error"`
		}
		err := json.Unmarshal(resp, &respObj)
		if err != nil {
			return GetStatusResponse{}, err
		}
		return GetStatusResponse{}, errors.New(respObj.Error)
	}
}

// Defines a 200 (Success) response for GetStatus.
type GetStatusResponse ServerStatusResponse
type PostFeeTransactionResponseEstimatedCost struct {
	ReadCount   int `json:"read_count"`
	ReadLength  int `json:"read_length"`
	Runtime     int `json:"runtime"`
	WriteCount  int `json:"write_count"`
	WriteLength int `json:"write_length"`
}
type GetAccountBalanceResponseFungibleTokens struct{}
type StxBalance struct {
	TotalMinerRewardsReceived string `json:"total_miner_rewards_received"`
	TotalReceived             string `json:"total_received"`
	Balance                   string `json:"balance"`
	Locked                    string `json:"locked"`
	LockHeight                int    `json:"lock_height"`
	LockTXID                  string `json:"lock_tx_id"`
	TotalFeesSent             string `json:"total_fees_sent"`
	TotalSent                 string `json:"total_sent"`
	BurnchainLockHeight       int    `json:"burnchain_lock_height"`
	BurnchainUnlockHeight     int    `json:"burnchain_unlock_height"`
}
type GetAllSubdomainsInName struct {
	GetAllSubdomainsInName string
}
type TransactionsWithTransfersResponseResultsAddressTransactionWithTransfersFtTransfersTransactionsWithTransfersResponseResultsAddressTransactionWithTransfersFtTransfers struct {
	AssetIdentifier string `json:"asset_identifier"`
	Recipient       string `json:"recipient"`
	Sender          string `json:"sender"`
	Amount          string `json:"amount"`
}
type BlockListResponse struct {
	Limit   int     `json:"limit"`
	Offset  int     `json:"offset"`
	Results []Block `json:"results"`
	Total   int     `json:"total"`
}
type GetSingleTransactionWithTransfersResponseFtTransfersGetSingleTransactionWithTransfersResponseFtTransfers struct {
	Amount          string `json:"amount"`
	AssetIdentifier string `json:"asset_identifier"`
	Recipient       string `json:"recipient"`
	Sender          string `json:"sender"`
}
type TransactionsWithTransfersResponseResultsAddressTransactionWithTransfersNftTransfersTransactionsWithTransfersResponseResultsAddressTransactionWithTransfersNftTransfers struct {
	Recipient       string                                                                                                                                                                      `json:"recipient"`
	Sender          string                                                                                                                                                                      `json:"sender"`
	Value           TransactionsWithTransfersResponseResultsAddressTransactionWithTransfersNftTransfersTransactionsWithTransfersResponseResultsAddressTransactionWithTransfersNftTransfersValue `json:"value"`
	AssetIdentifier string                                                                                                                                                                      `json:"asset_identifier"`
}
type AddressTransactionWithTransfers struct {
	FtTransfers  []GetSingleTransactionWithTransfersResponseFtTransfers  `json:"ft_transfers"`
	NftTransfers []GetSingleTransactionWithTransfersResponseNftTransfers `json:"nft_transfers"`
	StxReceived  string                                                  `json:"stx_received"`
	StxSent      string                                                  `json:"stx_sent"`
	StxTransfers []GetSingleTransactionWithTransfersResponseStxTransfers `json:"stx_transfers"`
	Tx           Transaction                                             `json:"tx"`
}
type TransactionList struct{}
type GetBlockByBurnBlockHeightResponseTxs struct {
	GetBlockByBurnBlockHeightResponseTxs string
}
type NftEvent struct {
	NftEvent NftEvent
}
type GetTransactionsByBlockHeightResponseResults struct {
	// This is an anyOf type, and so any of the following fields may be set
}
type GetAddressMempoolTransactionsResponseResults struct {
	// This is an anyOf type, and so any of the following fields may be set
}
type NetworkBlockTimesResponse struct {
	Mainnet TargetBlockTime `json:"mainnet"`
	Testnet TargetBlockTime `json:"testnet"`
}
type GetContractInterfaceResponseNonFungibleTokensGetContractInterfaceResponseNonFungibleTokens struct{}
type SmartContract struct {
	SmartContract SmartContract
}
type PostCoreNodeTransactionsError struct {
	Txid       string                                     `json:"txid"`
	Error      string                                     `json:"error"`
	Reason     string                                     `json:"reason"`
	ReasonData PostCoreNodeTransactionsError400ReasonData `json:"reason_data"`
}
type MicroblockListResponse struct {
	Limit   int          `json:"limit"`
	Offset  int          `json:"offset"`
	Results []Microblock `json:"results"`
	Total   int          `json:"total"`
}
type BnsGetNameInfoResponse struct {
	Blockchain   string `json:"blockchain"`
	ExpireBlock  int    `json:"expire_block"`
	GracePeriod  int    `json:"grace_period"`
	Resolver     string `json:"resolver"`
	ZonefileHash string `json:"zonefile_hash"`
	Address      string `json:"address"`
	LastTXID     string `json:"last_txid"`
	Status       string `json:"status"`
	Zonefile     string `json:"zonefile"`
}
type GetBlockListResponseResultsBlockTxs struct {
	GetBlockListResponseResultsBlockTxs string
}
type GetSingleTransactionWithTransfersResponseFtTransfers struct {
	GetSingleTransactionWithTransfersResponseFtTransfers GetSingleTransactionWithTransfersResponseFtTransfersGetSingleTransactionWithTransfersResponseFtTransfers
}
type GetContractInterfaceResponseFunctionsGetContractInterfaceResponseFunctions struct{}
type GetBlockByBurnBlockHashResponseTxs struct {
	GetBlockByBurnBlockHashResponseTxs string
}
type AddressTokenOfferingLocked struct {
	TotalUnlocked  string                  `json:"total_unlocked"`
	UnlockSchedule []AddressUnlockSchedule `json:"unlock_schedule"`
	TotalLocked    string                  `json:"total_locked"`
}
type GetAccountNoncesResponseDetectedMissingNonces struct {
	GetAccountNoncesResponseDetectedMissingNonces int
}
type PostFeeTransactionResponseEstimations struct {
	PostFeeTransactionResponseEstimations PostFeeTransactionResponseEstimationsPostFeeTransactionResponseEstimations
}
type AddressUnlockSchedule struct {
	AddressUnlockSchedule AddressUnlockSchedule
}
type ReadOnlyFunctionSuccessResponse struct {
	Cause  string `json:"cause"`
	Okay   bool   `json:"okay"`
	Result string `json:"result"`
}
type GetSingleTransactionWithTransfersResponseNftTransfersGetSingleTransactionWithTransfersResponseNftTransfers struct {
	AssetIdentifier string                                                                                                          `json:"asset_identifier"`
	Recipient       string                                                                                                          `json:"recipient"`
	Sender          string                                                                                                          `json:"sender"`
	Value           GetSingleTransactionWithTransfersResponseNftTransfersGetSingleTransactionWithTransfersResponseNftTransfersValue `json:"value"`
}
type BnsError struct {
	Error string `json:"error"`
}
type TransactionEventsResponse struct {
	Limit   int                              `json:"limit"`
	Offset  int                              `json:"offset"`
	Results GetFilteredEventsResponseResults `json:"results"`
}
type GetBlockByBurnBlockHeightResponseMicroblocksStreamed struct {
	GetBlockByBurnBlockHeightResponseMicroblocksStreamed string
}
type GetAllNamespacesResponseNamespaces struct {
	GetAllNamespacesResponseNamespaces string
}
type RunFaucetResponse struct {
	Success bool   `json:"success"`
	TxId    string `json:"txId"`
	TxRaw   string `json:"txRaw"`
}
type TransactionsWithTransfersResponseResultsAddressTransactionWithTransfersNftTransfersTransactionsWithTransfersResponseResultsAddressTransactionWithTransfersNftTransfersValue struct {
	Hex  string `json:"hex"`
	Repr string `json:"repr"`
}
type GetNFTMintsResponseResults struct {
	// This is an anyOf type, and so any of the following fields may be set
}
type MempoolTransactionListResponse struct {
	Offset  int                                             `json:"offset"`
	Results GetDroppedMempoolTransactionListResponseResults `json:"results"`
	Total   int                                             `json:"total"`
	Limit   int                                             `json:"limit"`
}
type GetTransactionListResponseResults struct {
	// This is an anyOf type, and so any of the following fields may be set
}
type GetFilteredEventsResponseResults struct {
	// This is an anyOf type, and so any of the following fields may be set
}
type AddressAssetsListResponse struct {
	Limit   int                             `json:"limit"`
	Offset  int                             `json:"offset"`
	Results GetAccountAssetsResponseResults `json:"results"`
	Total   int                             `json:"total"`
}
type GetBlockByBurnBlockHeightResponseMicroblocksAccepted struct {
	GetBlockByBurnBlockHeightResponseMicroblocksAccepted string
}
type GetAccountNFTResponseNftEventsNftEventValue struct {
	Hex  string `json:"hex"`
	Repr string `json:"repr"`
}
type BnsGetAllNamespacesResponse struct {
	Namespaces []GetAllNamespacesResponseNamespaces `json:"namespaces"`
}
type TransactionFeeEstimateResponse struct {
	CostScalarChangeByByte int                                     `json:"cost_scalar_change_by_byte"`
	EstimatedCost          PostFeeTransactionResponseEstimatedCost `json:"estimated_cost"`
	EstimatedCostScalar    int                                     `json:"estimated_cost_scalar"`
	Estimations            []PostFeeTransactionResponseEstimations `json:"estimations"`
}
type AddressBalanceResponse struct {
	FungibleTokens      GetAccountBalanceResponseFungibleTokens    `json:"fungible_tokens"`
	NonFungibleTokens   GetAccountBalanceResponseNonFungibleTokens `json:"non_fungible_tokens"`
	Stx                 StxBalance                                 `json:"stx"`
	TokenOfferingLocked AddressTokenOfferingLocked                 `json:"token_offering_locked"`
}
type GetContractInterfaceResponseFungibleTokensGetContractInterfaceResponseFungibleTokens struct{}
type SmartContractTransaction struct{}
type Microblock struct {
	ParentBlockHash        string                           `json:"parent_block_hash"`
	ParentBurnBlockHeight  int                              `json:"parent_burn_block_height"`
	Txs                    []GetMicroblockByHashResponseTxs `json:"txs"`
	MicroblockCanonical    bool                             `json:"microblock_canonical"`
	MicroblockParentHash   string                           `json:"microblock_parent_hash"`
	ParentBlockHeight      int                              `json:"parent_block_height"`
	ParentBurnBlockHash    string                           `json:"parent_burn_block_hash"`
	ParentBurnBlockTime    int                              `json:"parent_burn_block_time"`
	BlockHeight            int                              `json:"block_height"`
	MicroblockSequence     int                              `json:"microblock_sequence"`
	MicroblockHash         string                           `json:"microblock_hash"`
	ParentBurnBlockTimeISO string                           `json:"parent_burn_block_time_iso"`
	BlockHash              string                           `json:"block_hash"`
	Canonical              bool                             `json:"canonical"`
}
type GetNFTHoldingsResponseResults struct {
	// This is an anyOf type, and so any of the following fields may be set
}
type GetNFTHoldingsAssetIdentifiers struct {
	GetNFTHoldingsAssetIdentifiers string
}
type GetBlockListResponseResultsBlockMicroblocksStreamed struct {
	GetBlockListResponseResultsBlockMicroblocksStreamed string
}
type GetBlockListResponseResultsBlockMicroblocksAccepted struct {
	GetBlockListResponseResultsBlockMicroblocksAccepted string
}
type GetBlockByHashResponseMicroblocksAccepted struct {
	GetBlockByHashResponseMicroblocksAccepted string
}
type GetAccountAssetsResponseResults struct {
	// This is an anyOf type, and so any of the following fields may be set
}
type CoreNodePoxResponse struct {
	RewardCycleLength          int    `json:"reward_cycle_length"`
	ContractID                 string `json:"contract_id"`
	FirstBurnchainBlockHeight  int    `json:"first_burnchain_block_height"`
	RejectionVotesLeftRequired int    `json:"rejection_votes_left_required"`
	RewardCycleID              int    `json:"reward_cycle_id"`
	TotalLiquidSupplyUstx      int    `json:"total_liquid_supply_ustx"`
	MinAmountUstx              int    `json:"min_amount_ustx"`
	RegistrationWindowLength   int    `json:"registration_window_length"`
	RejectionFraction          int    `json:"rejection_fraction"`
}
type PostFeeTransactionResponseEstimationsPostFeeTransactionResponseEstimations struct {
	Fee     int `json:"fee"`
	FeeRate int `json:"fee_rate"`
}
type GetContractInterfaceResponseVariablesGetContractInterfaceResponseVariables struct{}
type UnanchoredTransactionListResponse struct {
	Results GetUnanchoredTxsResponseResults `json:"results"`
	Total   int                             `json:"total"`
}
type GetTransactionsByBlockHashResponseResults struct {
	// This is an anyOf type, and so any of the following fields may be set
}
type TransactionEvent struct{}
type FeeRate struct {
	FeeRate int `json:"fee_rate"`
}
type BurnchainRewardSlotHolder struct {
	BurnchainRewardSlotHolder BurnchainRewardSlotHolder
}
type ContractInterfaceResponse struct {
	Functions         []GetContractInterfaceResponseFunctions         `json:"functions"`
	FungibleTokens    []GetContractInterfaceResponseFungibleTokens    `json:"fungible_tokens"`
	Maps              []GetContractInterfaceResponseMaps              `json:"maps"`
	NonFungibleTokens []GetContractInterfaceResponseNonFungibleTokens `json:"non_fungible_tokens"`
	Variables         []GetContractInterfaceResponseVariables         `json:"variables"`
}
type AddressStxBalanceResponse struct{}
type GetBlockByHeightResponseTxs struct {
	GetBlockByHeightResponseTxs string
}
type BurnchainRewardSlotHolderListResponse struct {
	Limit   int                         `json:"limit"`
	Offset  int                         `json:"offset"`
	Results []BurnchainRewardSlotHolder `json:"results"`
	Total   int                         `json:"total"`
}
type GetAccountBalanceResponseNonFungibleTokens struct{}
type GetBlockByHashResponseMicroblocksStreamed struct {
	GetBlockByHashResponseMicroblocksStreamed string
}
type GetContractInterfaceResponseNonFungibleTokens struct {
	GetContractInterfaceResponseNonFungibleTokens GetContractInterfaceResponseNonFungibleTokensGetContractInterfaceResponseNonFungibleTokens
}
type AddressStxInboundListResponse struct {
	Limit   int                  `json:"limit"`
	Offset  int                  `json:"offset"`
	Results []InboundStxTransfer `json:"results"`
	Total   int                  `json:"total"`
}
type GetSingleTransactionWithTransfersResponseNftTransfers struct {
	GetSingleTransactionWithTransfersResponseNftTransfers GetSingleTransactionWithTransfersResponseNftTransfersGetSingleTransactionWithTransfersResponseNftTransfers
}
type Transaction struct{}
type GetBlockByHeightResponseMicroblocksStreamed struct {
	GetBlockByHeightResponseMicroblocksStreamed string
}
type GetSingleTransactionWithTransfersResponseNftTransfersGetSingleTransactionWithTransfersResponseNftTransfersValue struct {
	Repr string `json:"repr"`
	Hex  string `json:"hex"`
}
type AddressNonces struct {
	DetectedMissingNonces []GetAccountNoncesResponseDetectedMissingNonces `json:"detected_missing_nonces"`
	LastExecutedTXNonce   int                                             `json:"last_executed_tx_nonce"`
	LastMempoolTXNonce    int                                             `json:"last_mempool_tx_nonce"`
	PossibleNextNonce     int                                             `json:"possible_next_nonce"`
}
type FungibleTokenMetadata struct {
	FungibleTokenMetadata FungibleTokenMetadata
}
type GetContractInterfaceResponseMapsGetContractInterfaceResponseMaps struct{}
type GetBlockByBurnBlockHashResponseMicroblocksStreamed struct {
	GetBlockByBurnBlockHashResponseMicroblocksStreamed string
}
type AddressTransactionsWithTransfersListResponse struct {
	Limit   int                               `json:"limit"`
	Offset  int                               `json:"offset"`
	Results []AddressTransactionWithTransfers `json:"results"`
	Total   int                               `json:"total"`
}
type BurnchainRewardListResponse struct {
	Limit   int               `json:"limit"`
	Offset  int               `json:"offset"`
	Results []BurnchainReward `json:"results"`
}
type GetMicroblockByHashResponseTxs struct {
	GetMicroblockByHashResponseTxs string
}
type GetBlockByHashResponseTxs struct {
	GetBlockByHashResponseTxs string
}
type GetStxSupplyLegacyFormatResponse struct {
	UnlockedPercent         string `json:"unlockedPercent"`
	UnlockedSupply          string `json:"unlockedSupply"`
	UnlockedSupplyFormatted string `json:"unlockedSupplyFormatted"`
	BlockHeight             string `json:"blockHeight"`
	TotalStacks             string `json:"totalStacks"`
	TotalStacksFormatted    string `json:"totalStacksFormatted"`
}
type PostCoreNodeTransactionsError400ReasonData struct{}
type CoreNodeInfoResponse struct {
	PeerVersion            int    `json:"peer_version"`
	PoxConsensus           string `json:"pox_consensus"`
	ServerVersion          string `json:"server_version"`
	StableBurnBlockHeight  int    `json:"stable_burn_block_height"`
	StablePoxConsensus     string `json:"stable_pox_consensus"`
	BurnBlockHeight        int    `json:"burn_block_height"`
	NetworkID              int    `json:"network_id"`
	ParentNetworkID        int    `json:"parent_network_id"`
	StacksTip              string `json:"stacks_tip"`
	StacksTipConsensusHash string `json:"stacks_tip_consensus_hash"`
	ExitAtBlockHeight      int    `json:"exit_at_block_height"`
	StacksTipHeight        int    `json:"stacks_tip_height"`
	UnanchoredTip          string `json:"unanchored_tip"`
}
type ContractListResponse struct {
	Limit   int             `json:"limit"`
	Offset  int             `json:"offset"`
	Results []SmartContract `json:"results"`
}
type TransactionsWithTransfersResponseResultsAddressTransactionWithTransfersFtTransfers struct {
	TransactionsWithTransfersResponseResultsAddressTransactionWithTransfersFtTransfers TransactionsWithTransfersResponseResultsAddressTransactionWithTransfersFtTransfersTransactionsWithTransfersResponseResultsAddressTransactionWithTransfersFtTransfers
}
type GetContractInterfaceResponseMaps struct {
	GetContractInterfaceResponseMaps GetContractInterfaceResponseMapsGetContractInterfaceResponseMaps
}
type BnsGetAllNamespacesNamesResponse struct {
	BnsGetAllNamespacesNamesResponse string
}
type GetBlockByHeightResponseMicroblocksAccepted struct {
	GetBlockByHeightResponseMicroblocksAccepted string
}
type GetSingleTransactionWithTransfersResponseStxTransfers struct {
	GetSingleTransactionWithTransfersResponseStxTransfers GetSingleTransactionWithTransfersResponseStxTransfersGetSingleTransactionWithTransfersResponseStxTransfers
}
type ServerStatusResponse struct {
	ChainTip      ChainTip `json:"chain_tip"`
	ServerVersion string   `json:"server_version"`
	Status        string   `json:"status"`
}
type AddressTransactionsListResponse struct {
	Results GetAccountTransactionsResponseResults `json:"results"`
	Total   int                                   `json:"total"`
	Limit   int                                   `json:"limit"`
	Offset  int                                   `json:"offset"`
}
type Block struct {
	Block Block
}
type ChainTip struct {
	BlockHash          string `json:"block_hash"`
	BlockHeight        int    `json:"block_height"`
	IndexBlockHash     string `json:"index_block_hash"`
	MicroblockHash     string `json:"microblock_hash"`
	MicroblockSequence int    `json:"microblock_sequence"`
}
type GetFilteredEventsType struct{}
type TransactionsWithTransfersResponseResultsAddressTransactionWithTransfersStxTransfers struct {
	TransactionsWithTransfersResponseResultsAddressTransactionWithTransfersStxTransfers TransactionsWithTransfersResponseResultsAddressTransactionWithTransfersStxTransfersTransactionsWithTransfersResponseResultsAddressTransactionWithTransfersStxTransfers
}
type MapEntryResponse struct {
	Data  string `json:"data"`
	Proof string `json:"proof"`
}
type BurnchainRewardsTotal struct {
	RewardAmount    string `json:"reward_amount"`
	RewardRecipient string `json:"reward_recipient"`
}
type GetTransactionListType struct{}
type TransactionResults struct {
	Limit   int                                         `json:"limit"`
	Offset  int                                         `json:"offset"`
	Results GetTransactionsByBlockHeightResponseResults `json:"results"`
	Total   int                                         `json:"total"`
}
type GetNamesOwnedByAddressResponseNames struct {
	GetNamesOwnedByAddressResponseNames string
}
type BnsNamesOwnByAddressResponse struct {
	Names []GetNamesOwnedByAddressResponseNames `json:"names"`
}
type AddressNftListResponse struct {
	Offset    int        `json:"offset"`
	Total     int        `json:"total"`
	Limit     int        `json:"limit"`
	NftEvents []NftEvent `json:"nft_events"`
}
type GetNFTHistoryResponseResults struct {
	// This is an anyOf type, and so any of the following fields may be set
}
type BnsGetNamespacePriceResponse struct {
	Amount string `json:"amount"`
	Units  string `json:"units"`
}
type GetMicroblockListResponseResultsMicroblockTxs struct {
	GetMicroblockListResponseResultsMicroblockTxs string
}
type NonFungibleTokenHistoryEventList struct {
	Limit   int                          `json:"limit"`
	Offset  int                          `json:"offset"`
	Results GetNFTHistoryResponseResults `json:"results"`
	Total   int                          `json:"total"`
}
type NonFungibleTokensMetadataList struct {
	Limit   int                        `json:"limit"`
	Offset  int                        `json:"offset"`
	Results []NonFungibleTokenMetadata `json:"results"`
	Total   int                        `json:"total"`
}
type TransactionsWithTransfersResponseResultsAddressTransactionWithTransfersStxTransfersTransactionsWithTransfersResponseResultsAddressTransactionWithTransfersStxTransfers struct {
	Amount    string `json:"amount"`
	Recipient string `json:"recipient"`
	Sender    string `json:"sender"`
}
type NonFungibleTokenMintList struct {
	Results GetNFTMintsResponseResults `json:"results"`
	Total   int                        `json:"total"`
	Limit   int                        `json:"limit"`
	Offset  int                        `json:"offset"`
}
type GetStxSupplyResponse struct {
	BlockHeight     int    `json:"block_height"`
	TotalSTX        string `json:"total_stx"`
	UnlockedPercent string `json:"unlocked_percent"`
	UnlockedSTX     string `json:"unlocked_stx"`
}
type GetAccountTransactionsResponseResults struct {
	// This is an anyOf type, and so any of the following fields may be set
}
type FungibleTokensMetadataList struct {
	Offset  int                     `json:"offset"`
	Results []FungibleTokenMetadata `json:"results"`
	Total   int                     `json:"total"`
	Limit   int                     `json:"limit"`
}
type AccountDataResponse struct {
	Balance      string `json:"balance"`
	BalanceProof string `json:"balance_proof"`
	Locked       string `json:"locked"`
	Nonce        int    `json:"nonce"`
	NonceProof   string `json:"nonce_proof"`
	UnlockHeight int    `json:"unlock_height"`
}
type GetSingleTransactionWithTransfersResponseStxTransfersGetSingleTransactionWithTransfersResponseStxTransfers struct {
	Sender    string `json:"sender"`
	Amount    string `json:"amount"`
	Recipient string `json:"recipient"`
}
type GetDroppedMempoolTransactionListResponseResults struct {
	// This is an anyOf type, and so any of the following fields may be set
}
type ContractSourceResponse struct {
	Source        string `json:"source"`
	Proof         string `json:"proof"`
	PublishHeight int    `json:"publish_height"`
}
type BurnchainReward struct {
	BurnchainReward BurnchainReward
}
type BnsGetNamePriceResponse struct {
	Amount string `json:"amount"`
	Units  string `json:"units"`
}
type GetUnanchoredTxsResponseResults struct {
	// This is an anyOf type, and so any of the following fields may be set
}
type GetContractInterfaceResponseFunctions struct {
	GetContractInterfaceResponseFunctions GetContractInterfaceResponseFunctionsGetContractInterfaceResponseFunctions
}
type GetContractInterfaceResponseFungibleTokens struct {
	GetContractInterfaceResponseFungibleTokens GetContractInterfaceResponseFungibleTokensGetContractInterfaceResponseFungibleTokens
}
type InboundStxTransfer struct {
	InboundStxTransfer InboundStxTransfer
}
type BnsGetAllNamesResponse struct {
	BnsGetAllNamesResponse string
}
type GetContractInterfaceResponseVariables struct {
	GetContractInterfaceResponseVariables GetContractInterfaceResponseVariablesGetContractInterfaceResponseVariables
}
type GetBlockByBurnBlockHashResponseMicroblocksAccepted struct {
	GetBlockByBurnBlockHashResponseMicroblocksAccepted string
}
type GetRawTransactionResult struct {
	RawTX string `json:"raw_tx"`
}
type TransactionsWithTransfersResponseResultsAddressTransactionWithTransfersNftTransfers struct {
	TransactionsWithTransfersResponseResultsAddressTransactionWithTransfersNftTransfers TransactionsWithTransfersResponseResultsAddressTransactionWithTransfersNftTransfersTransactionsWithTransfersResponseResultsAddressTransactionWithTransfersNftTransfers
}
type NonFungibleTokenHoldingsList struct {
	Offset  int                           `json:"offset"`
	Results GetNFTHoldingsResponseResults `json:"results"`
	Total   int                           `json:"total"`
	Limit   int                           `json:"limit"`
}
type NonFungibleTokenMetadata struct {
	NonFungibleTokenMetadata NonFungibleTokenMetadata
}
type SearchResult struct{}
type TargetBlockTime struct {
	TargetBlockTime int `json:"target_block_time"`
}
type GetMempoolTransactionListResponseResults struct {
	// This is an anyOf type, and so any of the following fields may be set
}
type CallReadOnlyFunctionArguments struct {
	// The simulated tx-sender
	Sender string
	// An array of hex serialized Clarity values
	Arguments []string
}
