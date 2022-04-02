# \PremiumApi

All URIs are relative to *https://ob.nordigen.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveAccountBalancesV2**](PremiumApi.md#RetrieveAccountBalancesV2) | **Get** /api/v2/accounts/premium/{id}/balances/ | 
[**RetrieveAccountDetailsV2**](PremiumApi.md#RetrieveAccountDetailsV2) | **Get** /api/v2/accounts/premium/{id}/details/ | 
[**RetrieveAccountTransactionsV2**](PremiumApi.md#RetrieveAccountTransactionsV2) | **Get** /api/v2/accounts/premium/{id}/transactions/ | 



## RetrieveAccountBalancesV2

> RetrieveAccountBalancesV2(ctx, id).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PremiumApi.RetrieveAccountBalancesV2(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PremiumApi.RetrieveAccountBalancesV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAccountBalancesV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[jwtAuth](../README.md#jwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAccountDetailsV2

> RetrieveAccountDetailsV2(ctx, id).Country(country).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    country := "AT" // string | ISO 3166 two-character country code (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PremiumApi.RetrieveAccountDetailsV2(context.Background(), id).Country(country).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PremiumApi.RetrieveAccountDetailsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAccountDetailsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **country** | **string** | ISO 3166 two-character country code | 

### Return type

 (empty response body)

### Authorization

[jwtAuth](../README.md#jwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAccountTransactionsV2

> RetrieveAccountTransactionsV2(ctx, id).Country(country).DateFrom(dateFrom).DateTo(dateTo).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    country := "AT" // string | ISO 3166 two-character country code (optional)
    dateFrom := time.Now() // string |  (optional)
    dateTo := time.Now() // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PremiumApi.RetrieveAccountTransactionsV2(context.Background(), id).Country(country).DateFrom(dateFrom).DateTo(dateTo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PremiumApi.RetrieveAccountTransactionsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAccountTransactionsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **country** | **string** | ISO 3166 two-character country code | 
 **dateFrom** | **string** |  | 
 **dateTo** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[jwtAuth](../README.md#jwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

