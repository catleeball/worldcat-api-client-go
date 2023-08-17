# \MemberSharedPrintHoldingsApi

All URIs are relative to *https://americas.discovery.api.oclc.org/worldcat/search/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindBibRetainedHoldings**](MemberSharedPrintHoldingsApi.md#FindBibRetainedHoldings) | **Get** /bibs-retained-holdings | Find member shared print holdings for a known item



## FindBibRetainedHoldings

> FindBibRetainedHoldings200Response FindBibRetainedHoldings(ctx).Accept(accept).OclcNumber(oclcNumber).Isbn(isbn).Issn(issn).HeldByGroup(heldByGroup).HeldInState(heldInState).Offset(offset).Limit(limit).Execute()

Find member shared print holdings for a known item



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/catleeball/worldcat-api-client-go"
)

func main() {
    accept := "accept_example" // string | Content Type the client supports and is requesting (application/json currently supported) (optional) (default to "application/json")
    oclcNumber := int64(41266045) // int64 | oclc number of the bibliographic item (optional)
    isbn := "0439136350" // string | isbn of the bibliographic item (bn index) (optional)
    issn := "0099-1333" // string | issn of the bibliographic item (in index) (optional)
    heldByGroup := "ASRL" // string | Restrict to holdings held by Group Symbol (optional)
    heldInState := "US-TX" // string | Restrict to holdings held by Institutions in requested State/Province, per ISO 3166-2. (optional)
    offset := int32(10) // int32 | start position of the bib records to return (0 based), default 0 (optional) (default to 0)
    limit := int32(50) // int32 | maximum number of records to return, maximum 50, default 10 (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MemberSharedPrintHoldingsApi.FindBibRetainedHoldings(context.Background()).Accept(accept).OclcNumber(oclcNumber).Isbn(isbn).Issn(issn).HeldByGroup(heldByGroup).HeldInState(heldInState).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemberSharedPrintHoldingsApi.FindBibRetainedHoldings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindBibRetainedHoldings`: FindBibRetainedHoldings200Response
    fmt.Fprintf(os.Stdout, "Response from `MemberSharedPrintHoldingsApi.FindBibRetainedHoldings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindBibRetainedHoldingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Content Type the client supports and is requesting (application/json currently supported) | [default to &quot;application/json&quot;]
 **oclcNumber** | **int64** | oclc number of the bibliographic item | 
 **isbn** | **string** | isbn of the bibliographic item (bn index) | 
 **issn** | **string** | issn of the bibliographic item (in index) | 
 **heldByGroup** | **string** | Restrict to holdings held by Group Symbol | 
 **heldInState** | **string** | Restrict to holdings held by Institutions in requested State/Province, per ISO 3166-2. | 
 **offset** | **int32** | start position of the bib records to return (0 based), default 0 | [default to 0]
 **limit** | **int32** | maximum number of records to return, maximum 50, default 10 | [default to 10]

### Return type

[**FindBibRetainedHoldings200Response**](FindBibRetainedHoldings200Response.md)

### Authorization

[worldcat_search_auth](../README.md#worldcat_search_auth), [worldcat_search_auth](../README.md#worldcat_search_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

