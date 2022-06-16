# \FeesApi

All URIs are relative to *https://stacks-node-api.mainnet.stacks.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchFeeRate**](FeesApi.md#FetchFeeRate) | **Post** /extended/v1/fee_rate | Fetch fee rate
[**GetFeeTransfer**](FeesApi.md#GetFeeTransfer) | **Get** /v2/fees/transfer | Get estimated fee
[**PostFeeTransaction**](FeesApi.md#PostFeeTransaction) | **Post** /v2/fees/transaction | Get approximate fees for the given transaction



## FetchFeeRate

> FeeRate FetchFeeRate(ctx).FeeRateRequest(feeRateRequest).Execute()

Fetch fee rate



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
    feeRateRequest := *openapiclient.NewFeeRateRequest("Transaction_example") // FeeRateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeesApi.FetchFeeRate(context.Background()).FeeRateRequest(feeRateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeesApi.FetchFeeRate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchFeeRate`: FeeRate
    fmt.Fprintf(os.Stdout, "Response from `FeesApi.FetchFeeRate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchFeeRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feeRateRequest** | [**FeeRateRequest**](FeeRateRequest.md) |  | 

### Return type

[**FeeRate**](FeeRate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeeTransfer

> map[string]interface{} GetFeeTransfer(ctx).Execute()

Get estimated fee



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
    resp, r, err := apiClient.FeesApi.GetFeeTransfer(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeesApi.GetFeeTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeeTransfer`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `FeesApi.GetFeeTransfer`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeeTransferRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFeeTransaction

> TransactionFeeEstimateResponse PostFeeTransaction(ctx).TransactionFeeEstimateRequest(transactionFeeEstimateRequest).Execute()

Get approximate fees for the given transaction



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
    transactionFeeEstimateRequest := *openapiclient.NewTransactionFeeEstimateRequest("TransactionPayload_example") // TransactionFeeEstimateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeesApi.PostFeeTransaction(context.Background()).TransactionFeeEstimateRequest(transactionFeeEstimateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeesApi.PostFeeTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostFeeTransaction`: TransactionFeeEstimateResponse
    fmt.Fprintf(os.Stdout, "Response from `FeesApi.PostFeeTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFeeTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionFeeEstimateRequest** | [**TransactionFeeEstimateRequest**](TransactionFeeEstimateRequest.md) |  | 

### Return type

[**TransactionFeeEstimateResponse**](TransactionFeeEstimateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

