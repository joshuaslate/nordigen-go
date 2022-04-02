# \AgreementsApi

All URIs are relative to *https://ob.nordigen.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptEUA**](AgreementsApi.md#AcceptEUA) | **Put** /api/v2/agreements/enduser/{id}/accept/ | 
[**CreateEUAV2**](AgreementsApi.md#CreateEUAV2) | **Post** /api/v2/agreements/enduser/ | 
[**DeleteEUAByIdV2**](AgreementsApi.md#DeleteEUAByIdV2) | **Delete** /api/v2/agreements/enduser/{id}/ | 
[**RetrieveAllEUAsForAnEndUserV2**](AgreementsApi.md#RetrieveAllEUAsForAnEndUserV2) | **Get** /api/v2/agreements/enduser/ | 
[**RetrieveEUAByIdV2**](AgreementsApi.md#RetrieveEUAByIdV2) | **Get** /api/v2/agreements/enduser/{id}/ | 



## AcceptEUA

> EndUserAgreement AcceptEUA(ctx, id).EnduserAcceptanceDetailsRequest(enduserAcceptanceDetailsRequest).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this end user agreement.
    enduserAcceptanceDetailsRequest := *openapiclient.NewEnduserAcceptanceDetailsRequest("UserAgent_example", "IpAddress_example") // EnduserAcceptanceDetailsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgreementsApi.AcceptEUA(context.Background(), id).EnduserAcceptanceDetailsRequest(enduserAcceptanceDetailsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsApi.AcceptEUA``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcceptEUA`: EndUserAgreement
    fmt.Fprintf(os.Stdout, "Response from `AgreementsApi.AcceptEUA`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this end user agreement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptEUARequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enduserAcceptanceDetailsRequest** | [**EnduserAcceptanceDetailsRequest**](EnduserAcceptanceDetailsRequest.md) |  | 

### Return type

[**EndUserAgreement**](EndUserAgreement.md)

### Authorization

[jwtAuth](../README.md#jwtAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEUAV2

> EndUserAgreement CreateEUAV2(ctx).EndUserAgreementRequest(endUserAgreementRequest).Execute()





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
    endUserAgreementRequest := *openapiclient.NewEndUserAgreementRequest("InstitutionId_example") // EndUserAgreementRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgreementsApi.CreateEUAV2(context.Background()).EndUserAgreementRequest(endUserAgreementRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsApi.CreateEUAV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEUAV2`: EndUserAgreement
    fmt.Fprintf(os.Stdout, "Response from `AgreementsApi.CreateEUAV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEUAV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endUserAgreementRequest** | [**EndUserAgreementRequest**](EndUserAgreementRequest.md) |  | 

### Return type

[**EndUserAgreement**](EndUserAgreement.md)

### Authorization

[jwtAuth](../README.md#jwtAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEUAByIdV2

> map[string]interface{} DeleteEUAByIdV2(ctx, id).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this end user agreement.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgreementsApi.DeleteEUAByIdV2(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsApi.DeleteEUAByIdV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteEUAByIdV2`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AgreementsApi.DeleteEUAByIdV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this end user agreement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEUAByIdV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[jwtAuth](../README.md#jwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAllEUAsForAnEndUserV2

> PaginatedEndUserAgreementList RetrieveAllEUAsForAnEndUserV2(ctx).Limit(limit).Offset(offset).Execute()





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
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgreementsApi.RetrieveAllEUAsForAnEndUserV2(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsApi.RetrieveAllEUAsForAnEndUserV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveAllEUAsForAnEndUserV2`: PaginatedEndUserAgreementList
    fmt.Fprintf(os.Stdout, "Response from `AgreementsApi.RetrieveAllEUAsForAnEndUserV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAllEUAsForAnEndUserV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedEndUserAgreementList**](PaginatedEndUserAgreementList.md)

### Authorization

[jwtAuth](../README.md#jwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveEUAByIdV2

> EndUserAgreement RetrieveEUAByIdV2(ctx, id).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this end user agreement.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgreementsApi.RetrieveEUAByIdV2(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsApi.RetrieveEUAByIdV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveEUAByIdV2`: EndUserAgreement
    fmt.Fprintf(os.Stdout, "Response from `AgreementsApi.RetrieveEUAByIdV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this end user agreement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveEUAByIdV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EndUserAgreement**](EndUserAgreement.md)

### Authorization

[jwtAuth](../README.md#jwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

