# \LocalHoldingsResourcesApi

All URIs are relative to *https://americas.discovery.api.oclc.org/worldcat/search/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveLhr**](LocalHoldingsResourcesApi.md#RetrieveLhr) | **Get** /my-holdings/{controlNumber} | Retrieve LHR Resource
[**SearchLhr**](LocalHoldingsResourcesApi.md#SearchLhr) | **Get** /my-holdings | Search LHR Resources
[**SearchRetrainedLhr**](LocalHoldingsResourcesApi.md#SearchRetrainedLhr) | **Get** /retained-holdings | Search for shared print LHR Resources



## RetrieveLhr

> FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1 RetrieveLhr(ctx, controlNumber).Accept(accept).Execute()

Retrieve LHR Resource



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    controlNumber := "238374600" // string | LHR control number
    accept := "accept_example" // string | Content Type the client supports and is requesting (application/json currently supported) (optional) (default to "application/json")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocalHoldingsResourcesApi.RetrieveLhr(context.Background(), controlNumber).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalHoldingsResourcesApi.RetrieveLhr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveLhr`: FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1
    fmt.Fprintf(os.Stdout, "Response from `LocalHoldingsResourcesApi.RetrieveLhr`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**controlNumber** | **string** | LHR control number | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveLhrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** | Content Type the client supports and is requesting (application/json currently supported) | [default to &quot;application/json&quot;]

### Return type

[**FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1**](FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1.md)

### Authorization

[worldcat_search_auth](../README.md#worldcat_search_auth), [worldcat_search_auth](../README.md#worldcat_search_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchLhr

> SearchLhr200Response SearchLhr(ctx).Accept(accept).OclcNumber(oclcNumber).Barcode(barcode).OrderBy(orderBy).Offset(offset).Limit(limit).Execute()

Search LHR Resources



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    accept := "accept_example" // string | Content Type the client supports and is requesting (application/json currently supported) (optional) (default to "application/json")
    oclcNumber := int64(41266045) // int64 | oclc number of the bibliographic item (optional)
    barcode := "K123456789" // string | barcode (optional)
    orderBy := "orderBy_example" // string | result sort key (optional) (default to "oclcSymbol")
    offset := int32(10) // int32 | start position of the bib records to return (0 based), default 0 (optional) (default to 0)
    limit := int32(50) // int32 | maximum number of records to return, maximum 50, default 10 (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocalHoldingsResourcesApi.SearchLhr(context.Background()).Accept(accept).OclcNumber(oclcNumber).Barcode(barcode).OrderBy(orderBy).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalHoldingsResourcesApi.SearchLhr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchLhr`: SearchLhr200Response
    fmt.Fprintf(os.Stdout, "Response from `LocalHoldingsResourcesApi.SearchLhr`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchLhrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Content Type the client supports and is requesting (application/json currently supported) | [default to &quot;application/json&quot;]
 **oclcNumber** | **int64** | oclc number of the bibliographic item | 
 **barcode** | **string** | barcode | 
 **orderBy** | **string** | result sort key | [default to &quot;oclcSymbol&quot;]
 **offset** | **int32** | start position of the bib records to return (0 based), default 0 | [default to 0]
 **limit** | **int32** | maximum number of records to return, maximum 50, default 10 | [default to 10]

### Return type

[**SearchLhr200Response**](SearchLhr200Response.md)

### Authorization

[worldcat_search_auth](../README.md#worldcat_search_auth), [worldcat_search_auth](../README.md#worldcat_search_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchRetrainedLhr

> SearchLhr200Response SearchRetrainedLhr(ctx).Accept(accept).OclcNumber(oclcNumber).Barcode(barcode).HeldBy(heldBy).HeldBySymbol(heldBySymbol).HeldByInstitutionID(heldByInstitutionID).SpProgram(spProgram).OrderBy(orderBy).Offset(offset).Limit(limit).Execute()

Search for shared print LHR Resources



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    accept := "accept_example" // string | Content Type the client supports and is requesting (application/json currently supported) (optional) (default to "application/json")
    oclcNumber := int64(41266045) // int64 | oclc number of the bibliographic item (optional)
    barcode := "K123456789" // string | barcode (optional)
    heldBy := "TXH" // string | Institution Symbol (optional)
    heldBySymbol := []string{"Inner_example"} // []string | Institution Symbol (optional)
    heldByInstitutionID := []int32{int32(123)} // []int32 | Institution registryId (optional)
    spProgram := "WEST" // string | Limiter to restrict the response to bibliographic records associated with a particular shared print program (l8 index). (optional)
    orderBy := "orderBy_example" // string | result sort key (optional) (default to "oclcSymbol")
    offset := int32(10) // int32 | start position of the bib records to return (0 based), default 0 (optional) (default to 0)
    limit := int32(50) // int32 | maximum number of records to return, maximum 50, default 10 (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocalHoldingsResourcesApi.SearchRetrainedLhr(context.Background()).Accept(accept).OclcNumber(oclcNumber).Barcode(barcode).HeldBy(heldBy).HeldBySymbol(heldBySymbol).HeldByInstitutionID(heldByInstitutionID).SpProgram(spProgram).OrderBy(orderBy).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalHoldingsResourcesApi.SearchRetrainedLhr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchRetrainedLhr`: SearchLhr200Response
    fmt.Fprintf(os.Stdout, "Response from `LocalHoldingsResourcesApi.SearchRetrainedLhr`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRetrainedLhrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Content Type the client supports and is requesting (application/json currently supported) | [default to &quot;application/json&quot;]
 **oclcNumber** | **int64** | oclc number of the bibliographic item | 
 **barcode** | **string** | barcode | 
 **heldBy** | **string** | Institution Symbol | 
 **heldBySymbol** | **[]string** | Institution Symbol | 
 **heldByInstitutionID** | **[]int32** | Institution registryId | 
 **spProgram** | **string** | Limiter to restrict the response to bibliographic records associated with a particular shared print program (l8 index). | 
 **orderBy** | **string** | result sort key | [default to &quot;oclcSymbol&quot;]
 **offset** | **int32** | start position of the bib records to return (0 based), default 0 | [default to 0]
 **limit** | **int32** | maximum number of records to return, maximum 50, default 10 | [default to 10]

### Return type

[**SearchLhr200Response**](SearchLhr200Response.md)

### Authorization

[worldcat_search_auth](../README.md#worldcat_search_auth), [worldcat_search_auth](../README.md#worldcat_search_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

