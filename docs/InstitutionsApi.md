# \InstitutionsApi

All URIs are relative to *https://ob.nordigen.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveAllSupportedInstitutionsInAGivenCountry**](InstitutionsApi.md#RetrieveAllSupportedInstitutionsInAGivenCountry) | **Get** /api/v2/institutions/ | 
[**RetrieveInstitution**](InstitutionsApi.md#RetrieveInstitution) | **Get** /api/v2/institutions/{id}/ | 



## RetrieveAllSupportedInstitutionsInAGivenCountry

> []Aspsp RetrieveAllSupportedInstitutionsInAGivenCountry(ctx).Country(country).Execute()





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
    country := "AT" // string | ISO 3166 two-character country code

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstitutionsApi.RetrieveAllSupportedInstitutionsInAGivenCountry(context.Background()).Country(country).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstitutionsApi.RetrieveAllSupportedInstitutionsInAGivenCountry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveAllSupportedInstitutionsInAGivenCountry`: []Aspsp
    fmt.Fprintf(os.Stdout, "Response from `InstitutionsApi.RetrieveAllSupportedInstitutionsInAGivenCountry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAllSupportedInstitutionsInAGivenCountryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **string** | ISO 3166 two-character country code | 

### Return type

[**[]Aspsp**](Aspsp.md)

### Authorization

[jwtAuth](../README.md#jwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveInstitution

> Aspsp RetrieveInstitution(ctx, id).Execute()





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
    id := "N26_NTSBDEB1" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstitutionsApi.RetrieveInstitution(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstitutionsApi.RetrieveInstitution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveInstitution`: Aspsp
    fmt.Fprintf(os.Stdout, "Response from `InstitutionsApi.RetrieveInstitution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveInstitutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Aspsp**](Aspsp.md)

### Authorization

[jwtAuth](../README.md#jwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

