# \NonFungibleTokensApi

All URIs are relative to *https://stacks-node-api.mainnet.stacks.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContractNftMetadata**](NonFungibleTokensApi.md#GetContractNftMetadata) | **Get** /extended/v1/tokens/{contractId}/nft/metadata | Non fungible tokens metadata for contract id
[**GetNftHistory**](NonFungibleTokensApi.md#GetNftHistory) | **Get** /extended/v1/tokens/nft/history | Non-Fungible Token history
[**GetNftHoldings**](NonFungibleTokensApi.md#GetNftHoldings) | **Get** /extended/v1/tokens/nft/holdings | Non-Fungible Token holdings
[**GetNftMetadataList**](NonFungibleTokensApi.md#GetNftMetadataList) | **Get** /extended/v1/tokens/nft/metadata | Non fungible tokens metadata list
[**GetNftMints**](NonFungibleTokensApi.md#GetNftMints) | **Get** /extended/v1/tokens/nft/mints | Non-Fungible Token mints



## GetContractNftMetadata

> NonFungibleTokenMetadata GetContractNftMetadata(ctx, contractId).Execute()

Non fungible tokens metadata for contract id



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
    contractId := "contractId_example" // string | token's contract id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonFungibleTokensApi.GetContractNftMetadata(context.Background(), contractId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonFungibleTokensApi.GetContractNftMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContractNftMetadata`: NonFungibleTokenMetadata
    fmt.Fprintf(os.Stdout, "Response from `NonFungibleTokensApi.GetContractNftMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | token&#39;s contract id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractNftMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NonFungibleTokenMetadata**](NonFungibleTokenMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNftHistory

> NonFungibleTokenHistoryEventList GetNftHistory(ctx).AssetIdentifier(assetIdentifier).Value(value).Limit(limit).Offset(offset).Unanchored(unanchored).TxMetadata(txMetadata).Execute()

Non-Fungible Token history



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
    assetIdentifier := "SP2X0TZ59D5SZ8ACQ6YMCHHNR2ZN51Z32E2CJ173.the-explorer-guild::The-Explorer-Guild" // string | token asset class identifier
    value := "0x0100000000000000000000000000000803" // string | hex representation of the token's unique value
    limit := int32(56) // int32 | max number of events to fetch (optional) (default to 50)
    offset := int32(56) // int32 | index of first event to fetch (optional) (default to 0)
    unanchored := true // bool | whether or not to include events from unconfirmed transactions (optional) (default to false)
    txMetadata := true // bool | whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonFungibleTokensApi.GetNftHistory(context.Background()).AssetIdentifier(assetIdentifier).Value(value).Limit(limit).Offset(offset).Unanchored(unanchored).TxMetadata(txMetadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonFungibleTokensApi.GetNftHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNftHistory`: NonFungibleTokenHistoryEventList
    fmt.Fprintf(os.Stdout, "Response from `NonFungibleTokensApi.GetNftHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNftHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetIdentifier** | **string** | token asset class identifier | 
 **value** | **string** | hex representation of the token&#39;s unique value | 
 **limit** | **int32** | max number of events to fetch | [default to 50]
 **offset** | **int32** | index of first event to fetch | [default to 0]
 **unanchored** | **bool** | whether or not to include events from unconfirmed transactions | [default to false]
 **txMetadata** | **bool** | whether or not to include the complete transaction metadata instead of just &#x60;tx_id&#x60;. Enabling this option can affect performance and response times. | [default to false]

### Return type

[**NonFungibleTokenHistoryEventList**](NonFungibleTokenHistoryEventList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNftHoldings

> NonFungibleTokenHoldingsList GetNftHoldings(ctx).Principal(principal).AssetIdentifiers(assetIdentifiers).Limit(limit).Offset(offset).Unanchored(unanchored).TxMetadata(txMetadata).Execute()

Non-Fungible Token holdings



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
    principal := "SPNWZ5V2TPWGQGVDR6T7B6RQ4XMGZ4PXTEE0VQ0S.marketplace-v3" // string | token owner's STX address or Smart Contract ID
    assetIdentifiers := []string{"Inner_example"} // []string | identifiers of the token asset classes to filter for (optional)
    limit := int32(56) // int32 | max number of tokens to fetch (optional) (default to 50)
    offset := int32(56) // int32 | index of first tokens to fetch (optional) (default to 0)
    unanchored := true // bool | whether or not to include tokens from unconfirmed transactions (optional) (default to false)
    txMetadata := true // bool | whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonFungibleTokensApi.GetNftHoldings(context.Background()).Principal(principal).AssetIdentifiers(assetIdentifiers).Limit(limit).Offset(offset).Unanchored(unanchored).TxMetadata(txMetadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonFungibleTokensApi.GetNftHoldings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNftHoldings`: NonFungibleTokenHoldingsList
    fmt.Fprintf(os.Stdout, "Response from `NonFungibleTokensApi.GetNftHoldings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNftHoldingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **principal** | **string** | token owner&#39;s STX address or Smart Contract ID | 
 **assetIdentifiers** | **[]string** | identifiers of the token asset classes to filter for | 
 **limit** | **int32** | max number of tokens to fetch | [default to 50]
 **offset** | **int32** | index of first tokens to fetch | [default to 0]
 **unanchored** | **bool** | whether or not to include tokens from unconfirmed transactions | [default to false]
 **txMetadata** | **bool** | whether or not to include the complete transaction metadata instead of just &#x60;tx_id&#x60;. Enabling this option can affect performance and response times. | [default to false]

### Return type

[**NonFungibleTokenHoldingsList**](NonFungibleTokenHoldingsList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNftMetadataList

> NonFungibleTokensMetadataList GetNftMetadataList(ctx).Limit(limit).Offset(offset).Execute()

Non fungible tokens metadata list



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
    limit := int32(56) // int32 | max number of tokens to fetch (optional)
    offset := int32(56) // int32 | index of first tokens to fetch (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonFungibleTokensApi.GetNftMetadataList(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonFungibleTokensApi.GetNftMetadataList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNftMetadataList`: NonFungibleTokensMetadataList
    fmt.Fprintf(os.Stdout, "Response from `NonFungibleTokensApi.GetNftMetadataList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNftMetadataListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | max number of tokens to fetch | 
 **offset** | **int32** | index of first tokens to fetch | 

### Return type

[**NonFungibleTokensMetadataList**](NonFungibleTokensMetadataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNftMints

> NonFungibleTokenMintList GetNftMints(ctx).AssetIdentifier(assetIdentifier).Limit(limit).Offset(offset).Unanchored(unanchored).TxMetadata(txMetadata).Execute()

Non-Fungible Token mints



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
    assetIdentifier := "SP2X0TZ59D5SZ8ACQ6YMCHHNR2ZN51Z32E2CJ173.the-explorer-guild::The-Explorer-Guild" // string | token asset class identifier
    limit := int32(56) // int32 | max number of events to fetch (optional) (default to 50)
    offset := int32(56) // int32 | index of first event to fetch (optional) (default to 0)
    unanchored := true // bool | whether or not to include events from unconfirmed transactions (optional) (default to false)
    txMetadata := true // bool | whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonFungibleTokensApi.GetNftMints(context.Background()).AssetIdentifier(assetIdentifier).Limit(limit).Offset(offset).Unanchored(unanchored).TxMetadata(txMetadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonFungibleTokensApi.GetNftMints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNftMints`: NonFungibleTokenMintList
    fmt.Fprintf(os.Stdout, "Response from `NonFungibleTokensApi.GetNftMints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNftMintsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetIdentifier** | **string** | token asset class identifier | 
 **limit** | **int32** | max number of events to fetch | [default to 50]
 **offset** | **int32** | index of first event to fetch | [default to 0]
 **unanchored** | **bool** | whether or not to include events from unconfirmed transactions | [default to false]
 **txMetadata** | **bool** | whether or not to include the complete transaction metadata instead of just &#x60;tx_id&#x60;. Enabling this option can affect performance and response times. | [default to false]

### Return type

[**NonFungibleTokenMintList**](NonFungibleTokenMintList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

