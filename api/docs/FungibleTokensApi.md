# \FungibleTokensApi

All URIs are relative to *https://stacks-node-api.mainnet.stacks.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContractFtMetadata**](FungibleTokensApi.md#GetContractFtMetadata) | **Get** /extended/v1/tokens/{contractId}/ft/metadata | Fungible tokens metadata for contract id
[**GetFtMetadataList**](FungibleTokensApi.md#GetFtMetadataList) | **Get** /extended/v1/tokens/ft/metadata | Fungible tokens metadata list



## GetContractFtMetadata

> FungibleTokenMetadata GetContractFtMetadata(ctx, contractId).Execute()

Fungible tokens metadata for contract id



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
    resp, r, err := apiClient.FungibleTokensApi.GetContractFtMetadata(context.Background(), contractId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FungibleTokensApi.GetContractFtMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContractFtMetadata`: FungibleTokenMetadata
    fmt.Fprintf(os.Stdout, "Response from `FungibleTokensApi.GetContractFtMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | token&#39;s contract id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractFtMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FungibleTokenMetadata**](FungibleTokenMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFtMetadataList

> FungibleTokensMetadataList GetFtMetadataList(ctx).Limit(limit).Offset(offset).Execute()

Fungible tokens metadata list



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
    resp, r, err := apiClient.FungibleTokensApi.GetFtMetadataList(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FungibleTokensApi.GetFtMetadataList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFtMetadataList`: FungibleTokensMetadataList
    fmt.Fprintf(os.Stdout, "Response from `FungibleTokensApi.GetFtMetadataList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFtMetadataListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | max number of tokens to fetch | 
 **offset** | **int32** | index of first tokens to fetch | 

### Return type

[**FungibleTokensMetadataList**](FungibleTokensMetadataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

