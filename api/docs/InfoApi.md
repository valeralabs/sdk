# \InfoApi

All URIs are relative to *https://stacks-node-api.mainnet.stacks.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCoreApiInfo**](InfoApi.md#GetCoreApiInfo) | **Get** /v2/info | Get Core API info
[**GetNetworkBlockTimeByNetwork**](InfoApi.md#GetNetworkBlockTimeByNetwork) | **Get** /extended/v1/info/network_block_time/{network} | Get a given network&#39;s target block time
[**GetNetworkBlockTimes**](InfoApi.md#GetNetworkBlockTimes) | **Get** /extended/v1/info/network_block_times | Get the network target block time
[**GetPoxInfo**](InfoApi.md#GetPoxInfo) | **Get** /v2/pox | Get Proof-of-Transfer details
[**GetSTXSupply**](InfoApi.md#GetSTXSupply) | **Get** /extended/v1/STX_supply | Get total and unlocked STX supply
[**GetSTXSupplyCirculatingPlain**](InfoApi.md#GetSTXSupplyCirculatingPlain) | **Get** /extended/v1/STX_supply/circulating/plain | Get circulating STX supply in plain text format
[**GetSTXSupplyTotalSupplyPlain**](InfoApi.md#GetSTXSupplyTotalSupplyPlain) | **Get** /extended/v1/STX_supply/total/plain | Get total STX supply in plain text format
[**GetStatus**](InfoApi.md#GetStatus) | **Get** /extended/v1/status | API status
[**GetTotalSTXSupplyLegacyFormat**](InfoApi.md#GetTotalSTXSupplyLegacyFormat) | **Get** /extended/v1/STX_supply/legacy_format | Get total and unlocked STX supply (results formatted the same as the legacy 1.0 API)



## GetCoreApiInfo

> CoreNodeInfoResponse GetCoreApiInfo(ctx).Execute()

Get Core API info



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
    resp, r, err := apiClient.InfoApi.GetCoreApiInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfoApi.GetCoreApiInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCoreApiInfo`: CoreNodeInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `InfoApi.GetCoreApiInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCoreApiInfoRequest struct via the builder pattern


### Return type

[**CoreNodeInfoResponse**](CoreNodeInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkBlockTimeByNetwork

> NetworkBlockTimeResponse GetNetworkBlockTimeByNetwork(ctx, network).Execute()

Get a given network's target block time



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
    network := "network_example" // string | Which network to retrieve the target block time of

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfoApi.GetNetworkBlockTimeByNetwork(context.Background(), network).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfoApi.GetNetworkBlockTimeByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkBlockTimeByNetwork`: NetworkBlockTimeResponse
    fmt.Fprintf(os.Stdout, "Response from `InfoApi.GetNetworkBlockTimeByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | Which network to retrieve the target block time of | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkBlockTimeByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkBlockTimeResponse**](NetworkBlockTimeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkBlockTimes

> NetworkBlockTimesResponse GetNetworkBlockTimes(ctx).Execute()

Get the network target block time



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
    resp, r, err := apiClient.InfoApi.GetNetworkBlockTimes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfoApi.GetNetworkBlockTimes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkBlockTimes`: NetworkBlockTimesResponse
    fmt.Fprintf(os.Stdout, "Response from `InfoApi.GetNetworkBlockTimes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkBlockTimesRequest struct via the builder pattern


### Return type

[**NetworkBlockTimesResponse**](NetworkBlockTimesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoxInfo

> CoreNodePoxResponse GetPoxInfo(ctx).Execute()

Get Proof-of-Transfer details



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
    resp, r, err := apiClient.InfoApi.GetPoxInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfoApi.GetPoxInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPoxInfo`: CoreNodePoxResponse
    fmt.Fprintf(os.Stdout, "Response from `InfoApi.GetPoxInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoxInfoRequest struct via the builder pattern


### Return type

[**CoreNodePoxResponse**](CoreNodePoxResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSTXSupply

> GetStxSupplyResponse GetSTXSupply(ctx).Height(height).Execute()

Get total and unlocked STX supply



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
    height := float32(8.14) // float32 | The block height at which to query supply details from, if not provided then the latest block height is used (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfoApi.GetSTXSupply(context.Background()).Height(height).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfoApi.GetSTXSupply``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSTXSupply`: GetStxSupplyResponse
    fmt.Fprintf(os.Stdout, "Response from `InfoApi.GetSTXSupply`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSTXSupplyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **height** | **float32** | The block height at which to query supply details from, if not provided then the latest block height is used | 

### Return type

[**GetStxSupplyResponse**](GetStxSupplyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSTXSupplyCirculatingPlain

> string GetSTXSupplyCirculatingPlain(ctx).Execute()

Get circulating STX supply in plain text format



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
    resp, r, err := apiClient.InfoApi.GetSTXSupplyCirculatingPlain(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfoApi.GetSTXSupplyCirculatingPlain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSTXSupplyCirculatingPlain`: string
    fmt.Fprintf(os.Stdout, "Response from `InfoApi.GetSTXSupplyCirculatingPlain`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSTXSupplyCirculatingPlainRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSTXSupplyTotalSupplyPlain

> string GetSTXSupplyTotalSupplyPlain(ctx).Execute()

Get total STX supply in plain text format



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
    resp, r, err := apiClient.InfoApi.GetSTXSupplyTotalSupplyPlain(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfoApi.GetSTXSupplyTotalSupplyPlain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSTXSupplyTotalSupplyPlain`: string
    fmt.Fprintf(os.Stdout, "Response from `InfoApi.GetSTXSupplyTotalSupplyPlain`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSTXSupplyTotalSupplyPlainRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatus

> ServerStatusResponse GetStatus(ctx).Execute()

API status



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
    resp, r, err := apiClient.InfoApi.GetStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfoApi.GetStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatus`: ServerStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `InfoApi.GetStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusRequest struct via the builder pattern


### Return type

[**ServerStatusResponse**](ServerStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTotalSTXSupplyLegacyFormat

> GetStxSupplyLegacyFormatResponse GetTotalSTXSupplyLegacyFormat(ctx).Height(height).Execute()

Get total and unlocked STX supply (results formatted the same as the legacy 1.0 API)



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
    height := float32(8.14) // float32 | The block height at which to query supply details from, if not provided then the latest block height is used (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfoApi.GetTotalSTXSupplyLegacyFormat(context.Background()).Height(height).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfoApi.GetTotalSTXSupplyLegacyFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTotalSTXSupplyLegacyFormat`: GetStxSupplyLegacyFormatResponse
    fmt.Fprintf(os.Stdout, "Response from `InfoApi.GetTotalSTXSupplyLegacyFormat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTotalSTXSupplyLegacyFormatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **height** | **float32** | The block height at which to query supply details from, if not provided then the latest block height is used | 

### Return type

[**GetStxSupplyLegacyFormatResponse**](GetStxSupplyLegacyFormatResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

