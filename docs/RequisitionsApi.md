# \RequisitionsApi

All URIs are relative to *https://ob.nordigen.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRequisitionByIdV2**](RequisitionsApi.md#DeleteRequisitionByIdV2) | **Delete** /api/v2/requisitions/{id}/ | 
[**RequisitionById**](RequisitionsApi.md#RequisitionById) | **Get** /api/v2/requisitions/{id}/ | 
[**RequisitionCreated**](RequisitionsApi.md#RequisitionCreated) | **Post** /api/v2/requisitions/ | 
[**RetrieveAllRequisitions**](RequisitionsApi.md#RetrieveAllRequisitions) | **Get** /api/v2/requisitions/ | 



## DeleteRequisitionByIdV2

> map[string]interface{} DeleteRequisitionByIdV2(ctx, id).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this requisition.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequisitionsApi.DeleteRequisitionByIdV2(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequisitionsApi.DeleteRequisitionByIdV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRequisitionByIdV2`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RequisitionsApi.DeleteRequisitionByIdV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this requisition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequisitionByIdV2Request struct via the builder pattern


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


## RequisitionById

> RequisitionV2 RequisitionById(ctx, id).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this requisition.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequisitionsApi.RequisitionById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequisitionsApi.RequisitionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequisitionById`: RequisitionV2
    fmt.Fprintf(os.Stdout, "Response from `RequisitionsApi.RequisitionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this requisition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequisitionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequisitionV2**](RequisitionV2.md)

### Authorization

[jwtAuth](../README.md#jwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequisitionCreated

> SpectacularRequisitionV2 RequisitionCreated(ctx).RequisitionV2Request(requisitionV2Request).Execute()





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
    requisitionV2Request := *openapiclient.NewRequisitionV2Request("Redirect_example", "InstitutionId_example") // RequisitionV2Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequisitionsApi.RequisitionCreated(context.Background()).RequisitionV2Request(requisitionV2Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequisitionsApi.RequisitionCreated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequisitionCreated`: SpectacularRequisitionV2
    fmt.Fprintf(os.Stdout, "Response from `RequisitionsApi.RequisitionCreated`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequisitionCreatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requisitionV2Request** | [**RequisitionV2Request**](RequisitionV2Request.md) |  | 

### Return type

[**SpectacularRequisitionV2**](SpectacularRequisitionV2.md)

### Authorization

[jwtAuth](../README.md#jwtAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAllRequisitions

> PaginatedRequisitionV2List RetrieveAllRequisitions(ctx).Limit(limit).Offset(offset).Execute()





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
    resp, r, err := apiClient.RequisitionsApi.RetrieveAllRequisitions(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequisitionsApi.RetrieveAllRequisitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveAllRequisitions`: PaginatedRequisitionV2List
    fmt.Fprintf(os.Stdout, "Response from `RequisitionsApi.RetrieveAllRequisitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAllRequisitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedRequisitionV2List**](PaginatedRequisitionV2List.md)

### Authorization

[jwtAuth](../README.md#jwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

