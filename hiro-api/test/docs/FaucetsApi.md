# \FaucetsApi

All URIs are relative to *https://stacks-node-api.mainnet.stacks.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RunFaucetBtc**](FaucetsApi.md#RunFaucetBtc) | **Post** /extended/v1/faucets/btc | Add testnet BTC tokens to address
[**RunFaucetStx**](FaucetsApi.md#RunFaucetStx) | **Post** /extended/v1/faucets/stx | Get STX testnet tokens



## RunFaucetBtc

> Object RunFaucetBtc(ctx).Address(address).RunFaucetBtcRequest(runFaucetBtcRequest).Execute()

Add testnet BTC tokens to address



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
    address := "2N4M94S1ZPt8HfxydXzL2P7qyzgVq7MHWts" // string | A valid testnet BTC address
    runFaucetBtcRequest := *openapiclient.NewRunFaucetBtcRequest() // RunFaucetBtcRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FaucetsApi.RunFaucetBtc(context.Background()).Address(address).RunFaucetBtcRequest(runFaucetBtcRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FaucetsApi.RunFaucetBtc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunFaucetBtc`: Object
    fmt.Fprintf(os.Stdout, "Response from `FaucetsApi.RunFaucetBtc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunFaucetBtcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **address** | **string** | A valid testnet BTC address | 
 **runFaucetBtcRequest** | [**RunFaucetBtcRequest**](RunFaucetBtcRequest.md) |  | 

### Return type

[**Object**](Object.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunFaucetStx

> RunFaucetResponse RunFaucetStx(ctx).Address(address).Stacking(stacking).RunFaucetStxRequest(runFaucetStxRequest).Execute()

Get STX testnet tokens



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
    address := "ST3M7N9Q9HDRM7RVP1Q26P0EE69358PZZAZD7KMXQ" // string | A valid testnet STX address
    stacking := true // bool | Request the amount of STX tokens needed for individual address stacking (optional) (default to false)
    runFaucetStxRequest := *openapiclient.NewRunFaucetStxRequest() // RunFaucetStxRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FaucetsApi.RunFaucetStx(context.Background()).Address(address).Stacking(stacking).RunFaucetStxRequest(runFaucetStxRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FaucetsApi.RunFaucetStx``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunFaucetStx`: RunFaucetResponse
    fmt.Fprintf(os.Stdout, "Response from `FaucetsApi.RunFaucetStx`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunFaucetStxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **address** | **string** | A valid testnet STX address | 
 **stacking** | **bool** | Request the amount of STX tokens needed for individual address stacking | [default to false]
 **runFaucetStxRequest** | [**RunFaucetStxRequest**](RunFaucetStxRequest.md) |  | 

### Return type

[**RunFaucetResponse**](RunFaucetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

