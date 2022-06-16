# \MicroblocksApi

All URIs are relative to *https://stacks-node-api.mainnet.stacks.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMicroblockByHash**](MicroblocksApi.md#GetMicroblockByHash) | **Get** /extended/v1/microblock/{hash} | Get microblock
[**GetMicroblockList**](MicroblocksApi.md#GetMicroblockList) | **Get** /extended/v1/microblock | Get recent microblocks
[**GetUnanchoredTxs**](MicroblocksApi.md#GetUnanchoredTxs) | **Get** /extended/v1/microblock/unanchored/txs | Get the list of current transactions that belong to unanchored microblocks



## GetMicroblockByHash

> Microblock GetMicroblockByHash(ctx, hash).Execute()

Get microblock



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
    hash := "hash_example" // string | Hash of the microblock

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MicroblocksApi.GetMicroblockByHash(context.Background(), hash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MicroblocksApi.GetMicroblockByHash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMicroblockByHash`: Microblock
    fmt.Fprintf(os.Stdout, "Response from `MicroblocksApi.GetMicroblockByHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | Hash of the microblock | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMicroblockByHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Microblock**](Microblock.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMicroblockList

> MicroblockListResponse GetMicroblockList(ctx).Limit(limit).Offset(offset).Execute()

Get recent microblocks



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
    limit := int32(56) // int32 | Max number of microblocks to fetch (optional) (default to 20)
    offset := int32(56) // int32 | Index of the first microblock to fetch (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MicroblocksApi.GetMicroblockList(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MicroblocksApi.GetMicroblockList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMicroblockList`: MicroblockListResponse
    fmt.Fprintf(os.Stdout, "Response from `MicroblocksApi.GetMicroblockList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMicroblockListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of microblocks to fetch | [default to 20]
 **offset** | **int32** | Index of the first microblock to fetch | 

### Return type

[**MicroblockListResponse**](MicroblockListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUnanchoredTxs

> UnanchoredTransactionListResponse GetUnanchoredTxs(ctx).Execute()

Get the list of current transactions that belong to unanchored microblocks



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MicroblocksApi.GetUnanchoredTxs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MicroblocksApi.GetUnanchoredTxs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUnanchoredTxs`: UnanchoredTransactionListResponse
    fmt.Fprintf(os.Stdout, "Response from `MicroblocksApi.GetUnanchoredTxs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUnanchoredTxsRequest struct via the builder pattern


### Return type

[**UnanchoredTransactionListResponse**](UnanchoredTransactionListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

