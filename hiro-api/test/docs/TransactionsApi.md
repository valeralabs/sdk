# \TransactionsApi

All URIs are relative to *https://stacks-node-api.mainnet.stacks.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAddressMempoolTransactions**](TransactionsApi.md#GetAddressMempoolTransactions) | **Get** /extended/v1/address/{address}/mempool | Transactions for address
[**GetDroppedMempoolTransactionList**](TransactionsApi.md#GetDroppedMempoolTransactionList) | **Get** /extended/v1/tx/mempool/dropped | Get dropped mempool transactions
[**GetFilteredEvents**](TransactionsApi.md#GetFilteredEvents) | **Get** /extended/v1/tx/events | Transaction Events
[**GetMempoolTransactionList**](TransactionsApi.md#GetMempoolTransactionList) | **Get** /extended/v1/tx/mempool | Get mempool transactions
[**GetRawTransactionById**](TransactionsApi.md#GetRawTransactionById) | **Get** /extended/v1/tx/{tx_id}/raw | Get Raw Transaction
[**GetTransactionById**](TransactionsApi.md#GetTransactionById) | **Get** /extended/v1/tx/{tx_id} | Get transaction
[**GetTransactionList**](TransactionsApi.md#GetTransactionList) | **Get** /extended/v1/tx | Get recent transactions
[**GetTransactionsByBlockHash**](TransactionsApi.md#GetTransactionsByBlockHash) | **Get** /extended/v1/tx/block/{block_hash} | Transactions by block hash
[**GetTransactionsByBlockHeight**](TransactionsApi.md#GetTransactionsByBlockHeight) | **Get** /extended/v1/tx/block_height/{height} | Transactions by block height
[**GetTxListDetails**](TransactionsApi.md#GetTxListDetails) | **Get** /extended/v1/tx/multiple | Get list of details for transactions
[**PostCoreNodeTransactions**](TransactionsApi.md#PostCoreNodeTransactions) | **Post** /v2/transactions | Broadcast raw transaction



## GetAddressMempoolTransactions

> Object GetAddressMempoolTransactions(ctx, address).Limit(limit).Offset(offset).Unanchored(unanchored).Execute()

Transactions for address



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
    address := "address_example" // string | Transactions for the address
    limit := int32(56) // int32 | max number of transactions to fetch (optional)
    offset := int32(56) // int32 | index of first transaction to fetch (optional)
    unanchored := true // bool | Include transaction data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.GetAddressMempoolTransactions(context.Background(), address).Limit(limit).Offset(offset).Unanchored(unanchored).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.GetAddressMempoolTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAddressMempoolTransactions`: Object
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.GetAddressMempoolTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | Transactions for the address | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressMempoolTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | max number of transactions to fetch | 
 **offset** | **int32** | index of first transaction to fetch | 
 **unanchored** | **bool** | Include transaction data from unanchored (i.e. unconfirmed) microblocks | [default to false]

### Return type

[**Object**](Object.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDroppedMempoolTransactionList

> Object GetDroppedMempoolTransactionList(ctx).Limit(limit).Offset(offset).Execute()

Get dropped mempool transactions



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
    limit := int32(56) // int32 | max number of mempool transactions to fetch (optional) (default to 96)
    offset := int32(56) // int32 | index of first mempool transaction to fetch (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.GetDroppedMempoolTransactionList(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.GetDroppedMempoolTransactionList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDroppedMempoolTransactionList`: Object
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.GetDroppedMempoolTransactionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDroppedMempoolTransactionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | max number of mempool transactions to fetch | [default to 96]
 **offset** | **int32** | index of first mempool transaction to fetch | 

### Return type

[**Object**](Object.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFilteredEvents

> TransactionEventsResponse GetFilteredEvents(ctx).TxId(txId).Address(address).Limit(limit).Offset(offset).Type_(type_).Execute()

Transaction Events



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
    txId := "0x29e25515652dad41ef675bd0670964e3d537b80ec19cf6ca6f1dd65d5bc642c5" // string | Hash of transaction (optional)
    address := "ST1HB64MAJ1MBV4CQ80GF01DZS4T1DSMX20ADCRA4" // string | Stacks address or a Contract identifier (optional)
    limit := int32(100) // int32 | number of items to return (optional)
    offset := int32(50) // int32 | number of items to skip (optional)
    type_ := []string{"stx_lock"} // []string | Filter the events on event type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.GetFilteredEvents(context.Background()).TxId(txId).Address(address).Limit(limit).Offset(offset).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.GetFilteredEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFilteredEvents`: TransactionEventsResponse
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.GetFilteredEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFilteredEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **txId** | **string** | Hash of transaction | 
 **address** | **string** | Stacks address or a Contract identifier | 
 **limit** | **int32** | number of items to return | 
 **offset** | **int32** | number of items to skip | 
 **type_** | **[]string** | Filter the events on event type | 

### Return type

[**TransactionEventsResponse**](TransactionEventsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMempoolTransactionList

> MempoolTransactionListResponse GetMempoolTransactionList(ctx).SenderAddress(senderAddress).RecipientAddress(recipientAddress).Address(address).Limit(limit).Offset(offset).Unanchored(unanchored).Execute()

Get mempool transactions



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
    senderAddress := "senderAddress_example" // string | Filter to only return transactions with this sender address. (optional)
    recipientAddress := "recipientAddress_example" // string | Filter to only return transactions with this recipient address (only applicable for STX transfer tx types). (optional)
    address := "address_example" // string | Filter to only return transactions with this address as the sender or recipient (recipient only applicable for STX transfer tx types). (optional)
    limit := int32(56) // int32 | max number of mempool transactions to fetch (optional) (default to 96)
    offset := int32(56) // int32 | index of first mempool transaction to fetch (optional)
    unanchored := true // bool | Include transaction data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.GetMempoolTransactionList(context.Background()).SenderAddress(senderAddress).RecipientAddress(recipientAddress).Address(address).Limit(limit).Offset(offset).Unanchored(unanchored).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.GetMempoolTransactionList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMempoolTransactionList`: MempoolTransactionListResponse
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.GetMempoolTransactionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMempoolTransactionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **senderAddress** | **string** | Filter to only return transactions with this sender address. | 
 **recipientAddress** | **string** | Filter to only return transactions with this recipient address (only applicable for STX transfer tx types). | 
 **address** | **string** | Filter to only return transactions with this address as the sender or recipient (recipient only applicable for STX transfer tx types). | 
 **limit** | **int32** | max number of mempool transactions to fetch | [default to 96]
 **offset** | **int32** | index of first mempool transaction to fetch | 
 **unanchored** | **bool** | Include transaction data from unanchored (i.e. unconfirmed) microblocks | [default to false]

### Return type

[**MempoolTransactionListResponse**](MempoolTransactionListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRawTransactionById

> GetRawTransactionResult GetRawTransactionById(ctx, txId).Execute()

Get Raw Transaction



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
    txId := "txId_example" // string | Hash of transaction

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.GetRawTransactionById(context.Background(), txId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.GetRawTransactionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRawTransactionById`: GetRawTransactionResult
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.GetRawTransactionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**txId** | **string** | Hash of transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRawTransactionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRawTransactionResult**](GetRawTransactionResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionById

> Transaction GetTransactionById(ctx, txId).EventOffset(eventOffset).EventLimit(eventLimit).Unanchored(unanchored).Execute()

Get transaction



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
    txId := "txId_example" // string | Hash of transaction
    eventOffset := int32(56) // int32 | The number of events to skip (optional) (default to 0)
    eventLimit := int32(56) // int32 | The numbers of events to return (optional) (default to 96)
    unanchored := true // bool | Include transaction data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.GetTransactionById(context.Background(), txId).EventOffset(eventOffset).EventLimit(eventLimit).Unanchored(unanchored).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.GetTransactionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactionById`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.GetTransactionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**txId** | **string** | Hash of transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventOffset** | **int32** | The number of events to skip | [default to 0]
 **eventLimit** | **int32** | The numbers of events to return | [default to 96]
 **unanchored** | **bool** | Include transaction data from unanchored (i.e. unconfirmed) microblocks | [default to false]

### Return type

[**Transaction**](Transaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionList

> TransactionResults GetTransactionList(ctx).Limit(limit).Offset(offset).Type_(type_).Unanchored(unanchored).Execute()

Get recent transactions



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
    limit := int32(56) // int32 | max number of transactions to fetch (optional) (default to 96)
    offset := int32(56) // int32 | index of first transaction to fetch (optional)
    type_ := []string{"Type_example"} // []string | Filter by transaction type (optional)
    unanchored := true // bool | Include transaction data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.GetTransactionList(context.Background()).Limit(limit).Offset(offset).Type_(type_).Unanchored(unanchored).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.GetTransactionList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactionList`: TransactionResults
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.GetTransactionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | max number of transactions to fetch | [default to 96]
 **offset** | **int32** | index of first transaction to fetch | 
 **type_** | **[]string** | Filter by transaction type | 
 **unanchored** | **bool** | Include transaction data from unanchored (i.e. unconfirmed) microblocks | [default to false]

### Return type

[**TransactionResults**](TransactionResults.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionsByBlockHash

> Object GetTransactionsByBlockHash(ctx, blockHash).Limit(limit).Offset(offset).Execute()

Transactions by block hash



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
    blockHash := "blockHash_example" // string | Hash of block
    limit := int32(56) // int32 | max number of transactions to fetch (optional)
    offset := int32(56) // int32 | index of first transaction to fetch (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.GetTransactionsByBlockHash(context.Background(), blockHash).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.GetTransactionsByBlockHash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactionsByBlockHash`: Object
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.GetTransactionsByBlockHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockHash** | **string** | Hash of block | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsByBlockHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | max number of transactions to fetch | 
 **offset** | **int32** | index of first transaction to fetch | 

### Return type

[**Object**](Object.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionsByBlockHeight

> Object GetTransactionsByBlockHeight(ctx, height).Limit(limit).Offset(offset).Unanchored(unanchored).Execute()

Transactions by block height



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
    height := int32(56) // int32 | Height of block
    limit := int32(56) // int32 | max number of transactions to fetch (optional)
    offset := int32(56) // int32 | index of first transaction to fetch (optional)
    unanchored := true // bool | Include transaction data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.GetTransactionsByBlockHeight(context.Background(), height).Limit(limit).Offset(offset).Unanchored(unanchored).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.GetTransactionsByBlockHeight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactionsByBlockHeight`: Object
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.GetTransactionsByBlockHeight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**height** | **int32** | Height of block | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsByBlockHeightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | max number of transactions to fetch | 
 **offset** | **int32** | index of first transaction to fetch | 
 **unanchored** | **bool** | Include transaction data from unanchored (i.e. unconfirmed) microblocks | [default to false]

### Return type

[**Object**](Object.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTxListDetails

> map[string]TransactionList GetTxListDetails(ctx).TxId(txId).EventOffset(eventOffset).EventLimit(eventLimit).Unanchored(unanchored).Execute()

Get list of details for transactions



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
    txId := []string{"Inner_example"} // []string | Array of transaction ids
    eventOffset := int32(56) // int32 | The number of events to skip (optional) (default to 0)
    eventLimit := int32(56) // int32 | The numbers of events to return (optional) (default to 96)
    unanchored := true // bool | Include transaction data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.GetTxListDetails(context.Background()).TxId(txId).EventOffset(eventOffset).EventLimit(eventLimit).Unanchored(unanchored).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.GetTxListDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTxListDetails`: map[string]TransactionList
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.GetTxListDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTxListDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **txId** | **[]string** | Array of transaction ids | 
 **eventOffset** | **int32** | The number of events to skip | [default to 0]
 **eventLimit** | **int32** | The numbers of events to return | [default to 96]
 **unanchored** | **bool** | Include transaction data from unanchored (i.e. unconfirmed) microblocks | [default to false]

### Return type

[**map[string]TransactionList**](TransactionList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCoreNodeTransactions

> string PostCoreNodeTransactions(ctx).Body(body).Execute()

Broadcast raw transaction



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
    body := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.PostCoreNodeTransactions(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.PostCoreNodeTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCoreNodeTransactions`: string
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.PostCoreNodeTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCoreNodeTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

