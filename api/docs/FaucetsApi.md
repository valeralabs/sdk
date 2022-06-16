# \FaucetsApi

All URIs are relative to *https://stacks-node-api.mainnet.stacks.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RunFaucetBTC**](FaucetsApi.md#RunFaucetBTC) | **Post** /extended/v1/faucets/BTC | Add testnet BTC tokens to address
[**RunFaucetSTX**](FaucetsApi.md#RunFaucetSTX) | **Post** /extended/v1/faucets/STX | Get STX testnet tokens



## RunFaucetBTC

> RunFaucetResponse RunFaucetBTC(ctx).Address(address).RunFaucetBTCRequest(runFaucetBTCRequest).Execute()

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
    runFaucetBTCRequest := *openapiclient.NewRunFaucetBTCRequest() // RunFaucetBTCRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FaucetsApi.RunFaucetBTC(context.Background()).Address(address).RunFaucetBTCRequest(runFaucetBTCRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FaucetsApi.RunFaucetBTC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunFaucetBTC`: RunFaucetResponse
    fmt.Fprintf(os.Stdout, "Response from `FaucetsApi.RunFaucetBTC`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunFaucetBTCRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **address** | **string** | A valid testnet BTC address | 
 **runFaucetBTCRequest** | [**RunFaucetBTCRequest**](RunFaucetBTCRequest.md) |  | 

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


## RunFaucetSTX

> RunFaucetResponse RunFaucetSTX(ctx).Address(address).Stacking(stacking).RunFaucetSTXRequest(runFaucetSTXRequest).Execute()

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
    runFaucetSTXRequest := *openapiclient.NewRunFaucetSTXRequest() // RunFaucetSTXRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FaucetsApi.RunFaucetSTX(context.Background()).Address(address).Stacking(stacking).RunFaucetSTXRequest(runFaucetSTXRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FaucetsApi.RunFaucetSTX``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunFaucetSTX`: RunFaucetResponse
    fmt.Fprintf(os.Stdout, "Response from `FaucetsApi.RunFaucetSTX`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunFaucetSTXRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **address** | **string** | A valid testnet STX address | 
 **stacking** | **bool** | Request the amount of STX tokens needed for individual address stacking | [default to false]
 **runFaucetSTXRequest** | [**RunFaucetSTXRequest**](RunFaucetSTXRequest.md) |  | 

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

