# \TokenApi

All URIs are relative to *https://ob.nordigen.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JWTObtain**](TokenApi.md#JWTObtain) | **Post** /api/v2/token/new/ | 
[**JWTRefresh**](TokenApi.md#JWTRefresh) | **Post** /api/v2/token/refresh/ | 



## JWTObtain

> SpectacularJWTObtain JWTObtain(ctx).JWTObtainPairRequest(jWTObtainPairRequest).Execute()





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
    jWTObtainPairRequest := *openapiclient.NewJWTObtainPairRequest("SecretId_example", "SecretKey_example") // JWTObtainPairRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenApi.JWTObtain(context.Background()).JWTObtainPairRequest(jWTObtainPairRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenApi.JWTObtain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `JWTObtain`: SpectacularJWTObtain
    fmt.Fprintf(os.Stdout, "Response from `TokenApi.JWTObtain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiJWTObtainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jWTObtainPairRequest** | [**JWTObtainPairRequest**](JWTObtainPairRequest.md) |  | 

### Return type

[**SpectacularJWTObtain**](SpectacularJWTObtain.md)

### Authorization

[jwtAuth](../README.md#jwtAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JWTRefresh

> SpectacularJWTRefresh JWTRefresh(ctx).JWTRefreshRequest(jWTRefreshRequest).Execute()





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
    jWTRefreshRequest := *openapiclient.NewJWTRefreshRequest("Refresh_example") // JWTRefreshRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenApi.JWTRefresh(context.Background()).JWTRefreshRequest(jWTRefreshRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenApi.JWTRefresh``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `JWTRefresh`: SpectacularJWTRefresh
    fmt.Fprintf(os.Stdout, "Response from `TokenApi.JWTRefresh`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiJWTRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jWTRefreshRequest** | [**JWTRefreshRequest**](JWTRefreshRequest.md) |  | 

### Return type

[**SpectacularJWTRefresh**](SpectacularJWTRefresh.md)

### Authorization

[jwtAuth](../README.md#jwtAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

