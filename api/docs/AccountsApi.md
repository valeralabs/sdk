# \AccountsApi

All URIs are relative to *https://stacks-node-api.mainnet.stacks.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountAssets**](AccountsApi.md#GetAccountAssets) | **Get** /extended/v1/address/{principal}/assets | Get account assets
[**GetAccountBalance**](AccountsApi.md#GetAccountBalance) | **Get** /extended/v1/address/{principal}/balances | Get account balances
[**GetAccountInbound**](AccountsApi.md#GetAccountInbound) | **Get** /extended/v1/address/{principal}/STX_inbound | Get inbound STX transfers
[**GetAccountInfo**](AccountsApi.md#GetAccountInfo) | **Get** /v2/accounts/{principal} | Get account info
[**GetAccountNFT**](AccountsApi.md#GetAccountNFT) | **Get** /extended/v1/address/{principal}/NFT_events | Get NFT events
[**GetAccountNonces**](AccountsApi.md#GetAccountNonces) | **Get** /extended/v1/address/{principal}/nonces | Get the latest nonce used by an account
[**GetAccountSTXBalance**](AccountsApi.md#GetAccountSTXBalance) | **Get** /extended/v1/address/{principal}/STX | Get account STX balance
[**GetAccountTransactions**](AccountsApi.md#GetAccountTransactions) | **Get** /extended/v1/address/{principal}/transactions | Get account transactions
[**GetAccountTransactionsWithTransfers**](AccountsApi.md#GetAccountTransactionsWithTransfers) | **Get** /extended/v1/address/{principal}/transactions_with_transfers | Get account transactions including STX transfers for each transaction.
[**GetSingleTransactionWithTransfers**](AccountsApi.md#GetSingleTransactionWithTransfers) | **Get** /extended/v1/address/{principal}/{tx_id}/with_transfers | Get account transaction information for specific transaction



## GetAccountAssets

> AddressAssetsListResponse GetAccountAssets(ctx, principal).Limit(limit).Offset(offset).Unanchored(unanchored).UntilBlock(untilBlock).Execute()

Get account assets



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    principal := "principal_example" // string | Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
    limit := int32(56) // int32 | max number of account assets to fetch (optional)
    offset := int32(56) // int32 | index of first account assets to fetch (optional)
    unanchored := true // bool | Include transaction data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)
    untilBlock := "untilBlock_example" // string | returned data representing the state at that point in time, rather than the current block. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.GetAccountAssets(context.Background(), principal).Limit(limit).Offset(offset).Unanchored(unanchored).UntilBlock(untilBlock).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAccountAssets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountAssets`: AddressAssetsListResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAccountAssets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string** | Stacks address or a Contract identifier (e.g. &#x60;SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info&#x60;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | max number of account assets to fetch | 
 **offset** | **int32** | index of first account assets to fetch | 
 **unanchored** | **bool** | Include transaction data from unanchored (i.e. unconfirmed) microblocks | [default to false]
 **untilBlock** | **string** | returned data representing the state at that point in time, rather than the current block. | 

### Return type

[**AddressAssetsListResponse**](AddressAssetsListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountBalance

> AddressBalanceResponse GetAccountBalance(ctx, principal).Unanchored(unanchored).UntilBlock(untilBlock).Execute()

Get account balances



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    principal := "principal_example" // string | Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
    unanchored := true // bool | Include transaction data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)
    untilBlock := "untilBlock_example" // string | returned data representing the state up until that point in time, rather than the current block. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.GetAccountBalance(context.Background(), principal).Unanchored(unanchored).UntilBlock(untilBlock).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAccountBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountBalance`: AddressBalanceResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAccountBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string** | Stacks address or a Contract identifier (e.g. &#x60;SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info&#x60;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unanchored** | **bool** | Include transaction data from unanchored (i.e. unconfirmed) microblocks | [default to false]
 **untilBlock** | **string** | returned data representing the state up until that point in time, rather than the current block. | 

### Return type

[**AddressBalanceResponse**](AddressBalanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountInbound

> AddressStxInboundListResponse GetAccountInbound(ctx, principal).Limit(limit).Offset(offset).Height(height).Unanchored(unanchored).UntilBlock(untilBlock).Execute()

Get inbound STX transfers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    principal := "principal_example" // string | Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
    limit := int32(56) // int32 | number of items to return (optional)
    offset := int32(56) // int32 | number of items to skip (optional)
    height := float32(8.14) // float32 | Filter for transfers only at this given block height (optional)
    unanchored := true // bool | Include transaction data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)
    untilBlock := "untilBlock_example" // string | returned data representing the state up until that point in time, rather than the current block. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.GetAccountInbound(context.Background(), principal).Limit(limit).Offset(offset).Height(height).Unanchored(unanchored).UntilBlock(untilBlock).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAccountInbound``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountInbound`: AddressStxInboundListResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAccountInbound`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string** | Stacks address or a Contract identifier (e.g. &#x60;SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info&#x60;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountInboundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | number of items to return | 
 **offset** | **int32** | number of items to skip | 
 **height** | **float32** | Filter for transfers only at this given block height | 
 **unanchored** | **bool** | Include transaction data from unanchored (i.e. unconfirmed) microblocks | [default to false]
 **untilBlock** | **string** | returned data representing the state up until that point in time, rather than the current block. | 

### Return type

[**AddressStxInboundListResponse**](AddressStxInboundListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountInfo

> AccountDataResponse GetAccountInfo(ctx, principal).Proof(proof).Tip(tip).Execute()

Get account info



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    principal := "principal_example" // string | Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
    proof := int32(56) // int32 | Returns object without the proof field if set to 0 (optional)
    tip := "tip_example" // string | The Stacks chain tip to query from (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.GetAccountInfo(context.Background(), principal).Proof(proof).Tip(tip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAccountInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountInfo`: AccountDataResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAccountInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string** | Stacks address or a Contract identifier (e.g. &#x60;SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info&#x60;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **proof** | **int32** | Returns object without the proof field if set to 0 | 
 **tip** | **string** | The Stacks chain tip to query from | 

### Return type

[**AccountDataResponse**](AccountDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountNFT

> AddressNftListResponse GetAccountNFT(ctx, principal).Limit(limit).Offset(offset).Unanchored(unanchored).UntilBlock(untilBlock).Execute()

Get NFT events



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    principal := "principal_example" // string | Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
    limit := int32(56) // int32 | number of items to return (optional)
    offset := int32(56) // int32 | number of items to skip (optional)
    unanchored := true // bool | Include transaction data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)
    untilBlock := "untilBlock_example" // string | returned data representing the state up until that point in time, rather than the current block. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.GetAccountNFT(context.Background(), principal).Limit(limit).Offset(offset).Unanchored(unanchored).UntilBlock(untilBlock).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAccountNFT``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountNFT`: AddressNftListResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAccountNFT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string** | Stacks address or a Contract identifier (e.g. &#x60;SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info&#x60;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountNFTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | number of items to return | 
 **offset** | **int32** | number of items to skip | 
 **unanchored** | **bool** | Include transaction data from unanchored (i.e. unconfirmed) microblocks | [default to false]
 **untilBlock** | **string** | returned data representing the state up until that point in time, rather than the current block. | 

### Return type

[**AddressNftListResponse**](AddressNftListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountNonces

> AddressNonces GetAccountNonces(ctx, principal).BlockHeight(blockHeight).BlockHash(blockHash).Execute()

Get the latest nonce used by an account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    principal := "principal_example" // string | Stacks address (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0`)
    blockHeight := float32(8.14) // float32 | Optionally get the nonce at a given block height (optional)
    blockHash := "blockHash_example" // string | Optionally get the nonce at a given block hash (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.GetAccountNonces(context.Background(), principal).BlockHeight(blockHeight).BlockHash(blockHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAccountNonces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountNonces`: AddressNonces
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAccountNonces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string** | Stacks address (e.g. &#x60;SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0&#x60;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountNoncesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blockHeight** | **float32** | Optionally get the nonce at a given block height | 
 **blockHash** | **string** | Optionally get the nonce at a given block hash | 

### Return type

[**AddressNonces**](AddressNonces.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountSTXBalance

> AddressStxBalanceResponse GetAccountSTXBalance(ctx, principal).Unanchored(unanchored).UntilBlock(untilBlock).Execute()

Get account STX balance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    principal := "principal_example" // string | Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
    unanchored := true // bool | Include transaction data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)
    untilBlock := "untilBlock_example" // string | returned data representing the state up until that point in time, rather than the current block. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.GetAccountSTXBalance(context.Background(), principal).Unanchored(unanchored).UntilBlock(untilBlock).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAccountSTXBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountSTXBalance`: AddressStxBalanceResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAccountSTXBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string** | Stacks address or a Contract identifier (e.g. &#x60;SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info&#x60;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountSTXBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unanchored** | **bool** | Include transaction data from unanchored (i.e. unconfirmed) microblocks | [default to false]
 **untilBlock** | **string** | returned data representing the state up until that point in time, rather than the current block. | 

### Return type

[**AddressStxBalanceResponse**](AddressStxBalanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountTransactions

> AddressTransactionsListResponse GetAccountTransactions(ctx, principal).Limit(limit).Offset(offset).Height(height).Unanchored(unanchored).UntilBlock(untilBlock).Execute()

Get account transactions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    principal := "principal_example" // string | Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
    limit := int32(56) // int32 | max number of account transactions to fetch (optional)
    offset := int32(56) // int32 | index of first account transaction to fetch (optional)
    height := float32(8.14) // float32 | Filter for transactions only at this given block height (optional)
    unanchored := true // bool | Include transaction data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)
    untilBlock := "untilBlock_example" // string | returned data representing the state up until that point in time, rather than the current block. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.GetAccountTransactions(context.Background(), principal).Limit(limit).Offset(offset).Height(height).Unanchored(unanchored).UntilBlock(untilBlock).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAccountTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountTransactions`: AddressTransactionsListResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAccountTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string** | Stacks address or a Contract identifier (e.g. &#x60;SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info&#x60;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | max number of account transactions to fetch | 
 **offset** | **int32** | index of first account transaction to fetch | 
 **height** | **float32** | Filter for transactions only at this given block height | 
 **unanchored** | **bool** | Include transaction data from unanchored (i.e. unconfirmed) microblocks | [default to false]
 **untilBlock** | **string** | returned data representing the state up until that point in time, rather than the current block. | 

### Return type

[**AddressTransactionsListResponse**](AddressTransactionsListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountTransactionsWithTransfers

> AddressTransactionsWithTransfersListResponse GetAccountTransactionsWithTransfers(ctx, principal).Limit(limit).Offset(offset).Height(height).Unanchored(unanchored).UntilBlock(untilBlock).Execute()

Get account transactions including STX transfers for each transaction.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    principal := "principal_example" // string | Stacks address or a Contract identifier (e.g. `SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info`)
    limit := int32(56) // int32 | max number of account transactions to fetch (optional)
    offset := int32(56) // int32 | index of first account transaction to fetch (optional)
    height := float32(8.14) // float32 | Filter for transactions only at this given block height (optional)
    unanchored := true // bool | Include transaction data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)
    untilBlock := "untilBlock_example" // string | returned data representing the state up until that point in time, rather than the current block. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.GetAccountTransactionsWithTransfers(context.Background(), principal).Limit(limit).Offset(offset).Height(height).Unanchored(unanchored).UntilBlock(untilBlock).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAccountTransactionsWithTransfers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountTransactionsWithTransfers`: AddressTransactionsWithTransfersListResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAccountTransactionsWithTransfers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string** | Stacks address or a Contract identifier (e.g. &#x60;SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0.get-info&#x60;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountTransactionsWithTransfersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | max number of account transactions to fetch | 
 **offset** | **int32** | index of first account transaction to fetch | 
 **height** | **float32** | Filter for transactions only at this given block height | 
 **unanchored** | **bool** | Include transaction data from unanchored (i.e. unconfirmed) microblocks | [default to false]
 **untilBlock** | **string** | returned data representing the state up until that point in time, rather than the current block. | 

### Return type

[**AddressTransactionsWithTransfersListResponse**](AddressTransactionsWithTransfersListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingleTransactionWithTransfers

> AddressTransactionWithTransfers GetSingleTransactionWithTransfers(ctx, principal, txId).Execute()

Get account transaction information for specific transaction



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    principal := "principal_example" // string | Stacks address or a contract identifier
    txId := "txId_example" // string | Transaction id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.GetSingleTransactionWithTransfers(context.Background(), principal, txId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetSingleTransactionWithTransfers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSingleTransactionWithTransfers`: AddressTransactionWithTransfers
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetSingleTransactionWithTransfers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string** | Stacks address or a contract identifier | 
**txId** | **string** | Transaction id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleTransactionWithTransfersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AddressTransactionWithTransfers**](AddressTransactionWithTransfers.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

