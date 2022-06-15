# \BlocksApi

All URIs are relative to *https://stacks-node-api.mainnet.stacks.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBlockByBurnBlockHash**](BlocksApi.md#GetBlockByBurnBlockHash) | **Get** /extended/v1/block/by_burn_block_hash/{burn_block_hash} | Get block by burnchain block hash
[**GetBlockByBurnBlockHeight**](BlocksApi.md#GetBlockByBurnBlockHeight) | **Get** /extended/v1/block/by_burn_block_height/{burn_block_height} | Get block by burnchain height
[**GetBlockByHash**](BlocksApi.md#GetBlockByHash) | **Get** /extended/v1/block/{hash} | Get block by hash
[**GetBlockByHeight**](BlocksApi.md#GetBlockByHeight) | **Get** /extended/v1/block/by_height/{height} | Get block by height
[**GetBlockList**](BlocksApi.md#GetBlockList) | **Get** /extended/v1/block | Get recent blocks



## GetBlockByBurnBlockHash

> Object GetBlockByBurnBlockHash(ctx, burnBlockHash).Execute()

Get block by burnchain block hash



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
    burnBlockHash := "burnBlockHash_example" // string | Hash of the burnchain block

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlocksApi.GetBlockByBurnBlockHash(context.Background(), burnBlockHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlocksApi.GetBlockByBurnBlockHash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlockByBurnBlockHash`: Object
    fmt.Fprintf(os.Stdout, "Response from `BlocksApi.GetBlockByBurnBlockHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**burnBlockHash** | **string** | Hash of the burnchain block | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockByBurnBlockHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetBlockByBurnBlockHeight

> Object GetBlockByBurnBlockHeight(ctx, burnBlockHeight).Execute()

Get block by burnchain height



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
    burnBlockHeight := float32(8.14) // float32 | Height of the burn chain block

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlocksApi.GetBlockByBurnBlockHeight(context.Background(), burnBlockHeight).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlocksApi.GetBlockByBurnBlockHeight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlockByBurnBlockHeight`: Object
    fmt.Fprintf(os.Stdout, "Response from `BlocksApi.GetBlockByBurnBlockHeight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**burnBlockHeight** | **float32** | Height of the burn chain block | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockByBurnBlockHeightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetBlockByHash

> Block GetBlockByHash(ctx, hash).Execute()

Get block by hash



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
    hash := "hash_example" // string | Hash of the block

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlocksApi.GetBlockByHash(context.Background(), hash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlocksApi.GetBlockByHash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlockByHash`: Block
    fmt.Fprintf(os.Stdout, "Response from `BlocksApi.GetBlockByHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | Hash of the block | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockByHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Block**](Block.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockByHeight

> Object GetBlockByHeight(ctx, height).Execute()

Get block by height



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
    height := float32(8.14) // float32 | Height of the block

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlocksApi.GetBlockByHeight(context.Background(), height).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlocksApi.GetBlockByHeight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlockByHeight`: Object
    fmt.Fprintf(os.Stdout, "Response from `BlocksApi.GetBlockByHeight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**height** | **float32** | Height of the block | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockByHeightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetBlockList

> BlockListResponse GetBlockList(ctx).Limit(limit).Offset(offset).Execute()

Get recent blocks



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
    limit := int32(56) // int32 | max number of blocks to fetch (optional) (default to 20)
    offset := int32(56) // int32 | index of first block to fetch (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlocksApi.GetBlockList(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlocksApi.GetBlockList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlockList`: BlockListResponse
    fmt.Fprintf(os.Stdout, "Response from `BlocksApi.GetBlockList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | max number of blocks to fetch | [default to 20]
 **offset** | **int32** | index of first block to fetch | 

### Return type

[**BlockListResponse**](BlockListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

