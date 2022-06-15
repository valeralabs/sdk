# \NamesApi

All URIs are relative to *https://stacks-node-api.mainnet.stacks.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchSubdomainsListForName**](NamesApi.md#FetchSubdomainsListForName) | **Get** /v1/names/{name}/subdomains | Get Name Subdomains
[**FetchZoneFile**](NamesApi.md#FetchZoneFile) | **Get** /v1/names/{name}/zonefile | Get Zone File
[**GetAllNames**](NamesApi.md#GetAllNames) | **Get** /v1/names | Get All Names
[**GetAllNamespaces**](NamesApi.md#GetAllNamespaces) | **Get** /v1/namespaces | Get All Namespaces
[**GetHistoricalZoneFile**](NamesApi.md#GetHistoricalZoneFile) | **Get** /v1/names/{name}/zonefile/{zoneFileHash} | Get Historical Zone File
[**GetNameInfo**](NamesApi.md#GetNameInfo) | **Get** /v1/names/{name} | Get Name Details
[**GetNamePrice**](NamesApi.md#GetNamePrice) | **Get** /v2/prices/names/{name} | Get Name Price
[**GetNamesOwnedByAddress**](NamesApi.md#GetNamesOwnedByAddress) | **Get** /v1/addresses/{blockchain}/{address} | Get Names Owned by Address
[**GetNamespaceNames**](NamesApi.md#GetNamespaceNames) | **Get** /v1/namespaces/{tld}/names | Get Namespace Names
[**GetNamespacePrice**](NamesApi.md#GetNamespacePrice) | **Get** /v2/prices/namespaces/{tld} | Get Namespace Price



## FetchSubdomainsListForName

> []string FetchSubdomainsListForName(ctx, name).Execute()

Get Name Subdomains



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
    name := "id.blockstack" // string | fully-qualified name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamesApi.FetchSubdomainsListForName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamesApi.FetchSubdomainsListForName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSubdomainsListForName`: []string
    fmt.Fprintf(os.Stdout, "Response from `NamesApi.FetchSubdomainsListForName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | fully-qualified name | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchSubdomainsListForNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchZoneFile

> BnsFetchFileZoneResponse FetchZoneFile(ctx, name).Execute()

Get Zone File



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
    name := "bar.test" // string | fully-qualified name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamesApi.FetchZoneFile(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamesApi.FetchZoneFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchZoneFile`: BnsFetchFileZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `NamesApi.FetchZoneFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | fully-qualified name | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchZoneFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BnsFetchFileZoneResponse**](BnsFetchFileZoneResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllNames

> []string GetAllNames(ctx).Page(page).Execute()

Get All Names



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
    page := int32(23) // int32 | names are returned in pages of size 100, so specify the page number.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamesApi.GetAllNames(context.Background()).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamesApi.GetAllNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllNames`: []string
    fmt.Fprintf(os.Stdout, "Response from `NamesApi.GetAllNames`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | names are returned in pages of size 100, so specify the page number. | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllNamespaces

> BnsGetAllNamespacesResponse GetAllNamespaces(ctx).Execute()

Get All Namespaces



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
    resp, r, err := apiClient.NamesApi.GetAllNamespaces(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamesApi.GetAllNamespaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllNamespaces`: BnsGetAllNamespacesResponse
    fmt.Fprintf(os.Stdout, "Response from `NamesApi.GetAllNamespaces`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllNamespacesRequest struct via the builder pattern


### Return type

[**BnsGetAllNamespacesResponse**](BnsGetAllNamespacesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoricalZoneFile

> BnsFetchHistoricalZoneFileResponse GetHistoricalZoneFile(ctx, name, zoneFileHash).Execute()

Get Historical Zone File



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
    name := "muneeb.id" // string | fully-qualified name
    zoneFileHash := "b100a68235244b012854a95f9114695679002af9" // string | zone file hash

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamesApi.GetHistoricalZoneFile(context.Background(), name, zoneFileHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamesApi.GetHistoricalZoneFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHistoricalZoneFile`: BnsFetchHistoricalZoneFileResponse
    fmt.Fprintf(os.Stdout, "Response from `NamesApi.GetHistoricalZoneFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | fully-qualified name | 
**zoneFileHash** | **string** | zone file hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricalZoneFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BnsFetchHistoricalZoneFileResponse**](BnsFetchHistoricalZoneFileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNameInfo

> BnsGetNameInfoResponse GetNameInfo(ctx, name).Execute()

Get Name Details



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
    name := "muneeb.id" // string | fully-qualified name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamesApi.GetNameInfo(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamesApi.GetNameInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNameInfo`: BnsGetNameInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `NamesApi.GetNameInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | fully-qualified name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNameInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BnsGetNameInfoResponse**](BnsGetNameInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamePrice

> BnsGetNamePriceResponse GetNamePrice(ctx, name).Execute()

Get Name Price



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
    name := "muneeb.id" // string | the name to query price information for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamesApi.GetNamePrice(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamesApi.GetNamePrice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNamePrice`: BnsGetNamePriceResponse
    fmt.Fprintf(os.Stdout, "Response from `NamesApi.GetNamePrice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | the name to query price information for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamePriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BnsGetNamePriceResponse**](BnsGetNamePriceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamesOwnedByAddress

> BnsNamesOwnByAddressResponse GetNamesOwnedByAddress(ctx, blockchain, address).Execute()

Get Names Owned by Address



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
    blockchain := "bitcoin" // string | the layer-1 blockchain for the address
    address := "1QJQxDas5JhdiXhEbNS14iNjr8auFT96GP" // string | the address to lookup

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamesApi.GetNamesOwnedByAddress(context.Background(), blockchain, address).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamesApi.GetNamesOwnedByAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNamesOwnedByAddress`: BnsNamesOwnByAddressResponse
    fmt.Fprintf(os.Stdout, "Response from `NamesApi.GetNamesOwnedByAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | the layer-1 blockchain for the address | 
**address** | **string** | the address to lookup | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamesOwnedByAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BnsNamesOwnByAddressResponse**](BnsNamesOwnByAddressResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespaceNames

> []string GetNamespaceNames(ctx, tld).Page(page).Execute()

Get Namespace Names



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
    tld := "id" // string | the namespace to fetch names from
    page := int32(23) // int32 | names are returned in pages of size 100, so specify the page number.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamesApi.GetNamespaceNames(context.Background(), tld).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamesApi.GetNamespaceNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNamespaceNames`: []string
    fmt.Fprintf(os.Stdout, "Response from `NamesApi.GetNamespaceNames`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tld** | **string** | the namespace to fetch names from | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespaceNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | names are returned in pages of size 100, so specify the page number. | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespacePrice

> BnsGetNamespacePriceResponse GetNamespacePrice(ctx, tld).Execute()

Get Namespace Price



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
    tld := "id" // string | the namespace to fetch price for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamesApi.GetNamespacePrice(context.Background(), tld).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamesApi.GetNamespacePrice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNamespacePrice`: BnsGetNamespacePriceResponse
    fmt.Fprintf(os.Stdout, "Response from `NamesApi.GetNamespacePrice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tld** | **string** | the namespace to fetch price for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespacePriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BnsGetNamespacePriceResponse**](BnsGetNamespacePriceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

