# \SearchApi

All URIs are relative to *https://stacks-node-api.mainnet.stacks.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchById**](SearchApi.md#SearchById) | **Get** /extended/v1/search/{id} | Search



## SearchById

> SearchResult SearchById(ctx, id).IncludeMetadata(includeMetadata).Execute()

Search



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
    id := "id_example" // string | The hex hash string for a block or transaction, account address, or contract address
    includeMetadata := true // bool | This includes the detailed data for purticular hash in the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.SearchById(context.Background(), id).IncludeMetadata(includeMetadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchById`: SearchResult
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The hex hash string for a block or transaction, account address, or contract address | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeMetadata** | **bool** | This includes the detailed data for purticular hash in the response | 

### Return type

[**SearchResult**](SearchResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

