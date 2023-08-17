# \MemberGeneralHoldingsApi

All URIs are relative to *https://americas.discovery.api.oclc.org/worldcat/search/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindBibDetailedHoldings**](MemberGeneralHoldingsApi.md#FindBibDetailedHoldings) | **Get** /bibs-detailed-holdings | Find detailed info of member holdings for a known item
[**FindBibHoldings**](MemberGeneralHoldingsApi.md#FindBibHoldings) | **Get** /bibs-holdings | Find member holdings for a known item
[**FindBibSummaryHoldings**](MemberGeneralHoldingsApi.md#FindBibSummaryHoldings) | **Get** /bibs-summary-holdings | Get summary of holdings for a known item



## FindBibDetailedHoldings

> FindBibRetainedHoldings200Response FindBibDetailedHoldings(ctx).Accept(accept).OclcNumber(oclcNumber).Isbn(isbn).Issn(issn).HeldByGroup(heldByGroup).HeldBy(heldBy).HeldBySymbol(heldBySymbol).HeldByInstitutionID(heldByInstitutionID).Offset(offset).Limit(limit).Execute()

Find detailed info of member holdings for a known item



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
    isbn := "0439136350" // string | isbn of the bibliographic item (bn index) (optional)
    issn := "0099-1333" // string | issn of the bibliographic item (in index) (optional)
    heldByGroup := "ASRL" // string | Restrict to holdings held by Group Symbol (optional)
    heldBy := "TXH" // string | Institution Symbol (optional)
    heldBySymbol := []string{"Inner_example"} // []string | Institution Symbol (optional)
    heldByInstitutionID := []int32{int32(123)} // []int32 | Institution registryId (optional)
    offset := int32(10) // int32 | start position of the bib records to return (0 based), default 0 (optional) (default to 0)
    limit := int32(50) // int32 | maximum number of records to return, maximum 50, default 10 (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MemberGeneralHoldingsApi.FindBibDetailedHoldings(context.Background()).Accept(accept).OclcNumber(oclcNumber).Isbn(isbn).Issn(issn).HeldByGroup(heldByGroup).HeldBy(heldBy).HeldBySymbol(heldBySymbol).HeldByInstitutionID(heldByInstitutionID).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemberGeneralHoldingsApi.FindBibDetailedHoldings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindBibDetailedHoldings`: FindBibRetainedHoldings200Response
    fmt.Fprintf(os.Stdout, "Response from `MemberGeneralHoldingsApi.FindBibDetailedHoldings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindBibDetailedHoldingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Content Type the client supports and is requesting (application/json currently supported) | [default to &quot;application/json&quot;]
 **oclcNumber** | **int64** | oclc number of the bibliographic item | 
 **isbn** | **string** | isbn of the bibliographic item (bn index) | 
 **issn** | **string** | issn of the bibliographic item (in index) | 
 **heldByGroup** | **string** | Restrict to holdings held by Group Symbol | 
 **heldBy** | **string** | Institution Symbol | 
 **heldBySymbol** | **[]string** | Institution Symbol | 
 **heldByInstitutionID** | **[]int32** | Institution registryId | 
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


## FindBibHoldings

> FindBibRetainedHoldings200Response FindBibHoldings(ctx).Accept(accept).OclcNumber(oclcNumber).Isbn(isbn).Issn(issn).HoldingsAllEditions(holdingsAllEditions).HoldingsAllVariantRecords(holdingsAllVariantRecords).PreferredLanguage(preferredLanguage).HeldInCountry(heldInCountry).HeldBySymbol(heldBySymbol).HeldByInstitutionID(heldByInstitutionID).Lat(lat).Lon(lon).Distance(distance).Unit(unit).NumberNearestHoldings(numberNearestHoldings).Offset(offset).Limit(limit).Execute()

Find member holdings for a known item



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
    isbn := "0439136350" // string | isbn of the bibliographic item (bn index) (optional)
    issn := "0099-1333" // string | issn of the bibliographic item (in index) (optional)
    holdingsAllEditions := true // bool | Get holdings for all editions (optional)
    holdingsAllVariantRecords := true // bool | Get holdings for specific edition across variant records (optional)
    preferredLanguage := "fre" // string | language user would prefer metadata description in (optional)
    heldInCountry := "US" // string | Restrict to holdings held by institutions in requested country (optional)
    heldBySymbol := []string{"Inner_example"} // []string | Institution Symbol (optional)
    heldByInstitutionID := []int32{int32(123)} // []int32 | Institution registryId (optional)
    lat := float64(37.502508) // float64 | latitude (optional)
    lon := float64(-122.22702) // float64 | longitude (optional)
    distance := int32(10) // int32 | distance (optional)
    unit := "unit_example" // string | distance unit (optional) (default to "M")
    numberNearestHoldings := int32(20) // int32 | When a lat/lon is passed, the number of nearest holdings to show even if outside the search radius. (optional)
    offset := int32(10) // int32 | start position of the bib records to return (0 based), default 0 (optional) (default to 0)
    limit := int32(50) // int32 | maximum number of records to return, maximum 50, default 10 (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MemberGeneralHoldingsApi.FindBibHoldings(context.Background()).Accept(accept).OclcNumber(oclcNumber).Isbn(isbn).Issn(issn).HoldingsAllEditions(holdingsAllEditions).HoldingsAllVariantRecords(holdingsAllVariantRecords).PreferredLanguage(preferredLanguage).HeldInCountry(heldInCountry).HeldBySymbol(heldBySymbol).HeldByInstitutionID(heldByInstitutionID).Lat(lat).Lon(lon).Distance(distance).Unit(unit).NumberNearestHoldings(numberNearestHoldings).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemberGeneralHoldingsApi.FindBibHoldings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindBibHoldings`: FindBibRetainedHoldings200Response
    fmt.Fprintf(os.Stdout, "Response from `MemberGeneralHoldingsApi.FindBibHoldings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindBibHoldingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Content Type the client supports and is requesting (application/json currently supported) | [default to &quot;application/json&quot;]
 **oclcNumber** | **int64** | oclc number of the bibliographic item | 
 **isbn** | **string** | isbn of the bibliographic item (bn index) | 
 **issn** | **string** | issn of the bibliographic item (in index) | 
 **holdingsAllEditions** | **bool** | Get holdings for all editions | 
 **holdingsAllVariantRecords** | **bool** | Get holdings for specific edition across variant records | 
 **preferredLanguage** | **string** | language user would prefer metadata description in | 
 **heldInCountry** | **string** | Restrict to holdings held by institutions in requested country | 
 **heldBySymbol** | **[]string** | Institution Symbol | 
 **heldByInstitutionID** | **[]int32** | Institution registryId | 
 **lat** | **float64** | latitude | 
 **lon** | **float64** | longitude | 
 **distance** | **int32** | distance | 
 **unit** | **string** | distance unit | [default to &quot;M&quot;]
 **numberNearestHoldings** | **int32** | When a lat/lon is passed, the number of nearest holdings to show even if outside the search radius. | 
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


## FindBibSummaryHoldings

> FindBibRetainedHoldings200Response FindBibSummaryHoldings(ctx).Accept(accept).OclcNumber(oclcNumber).Isbn(isbn).Issn(issn).HoldingsAllEditions(holdingsAllEditions).HoldingsAllVariantRecords(holdingsAllVariantRecords).PreferredLanguage(preferredLanguage).HeldInCountry(heldInCountry).HeldByGroup(heldByGroup).HeldBy(heldBy).HeldBySymbol(heldBySymbol).HeldByInstitutionID(heldByInstitutionID).Lat(lat).Lon(lon).Distance(distance).Unit(unit).Execute()

Get summary of holdings for a known item



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
    isbn := "0439136350" // string | isbn of the bibliographic item (bn index) (optional)
    issn := "0099-1333" // string | issn of the bibliographic item (in index) (optional)
    holdingsAllEditions := true // bool | Get holdings for all editions (optional)
    holdingsAllVariantRecords := true // bool | Get holdings for specific edition across variant records (optional)
    preferredLanguage := "fre" // string | language user would prefer metadata description in (optional)
    heldInCountry := "US" // string | Restrict to holdings held by institutions in requested country (optional)
    heldByGroup := "ASRL" // string | Restrict to holdings held by Group Symbol (optional)
    heldBy := "TXH" // string | Institution Symbol (optional)
    heldBySymbol := []string{"Inner_example"} // []string | Institution Symbol (optional)
    heldByInstitutionID := []int32{int32(123)} // []int32 | Institution registryId (optional)
    lat := float64(37.502508) // float64 | latitude (optional)
    lon := float64(-122.22702) // float64 | longitude (optional)
    distance := int32(10) // int32 | distance (optional)
    unit := "unit_example" // string | distance unit (optional) (default to "M")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MemberGeneralHoldingsApi.FindBibSummaryHoldings(context.Background()).Accept(accept).OclcNumber(oclcNumber).Isbn(isbn).Issn(issn).HoldingsAllEditions(holdingsAllEditions).HoldingsAllVariantRecords(holdingsAllVariantRecords).PreferredLanguage(preferredLanguage).HeldInCountry(heldInCountry).HeldByGroup(heldByGroup).HeldBy(heldBy).HeldBySymbol(heldBySymbol).HeldByInstitutionID(heldByInstitutionID).Lat(lat).Lon(lon).Distance(distance).Unit(unit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemberGeneralHoldingsApi.FindBibSummaryHoldings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindBibSummaryHoldings`: FindBibRetainedHoldings200Response
    fmt.Fprintf(os.Stdout, "Response from `MemberGeneralHoldingsApi.FindBibSummaryHoldings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindBibSummaryHoldingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Content Type the client supports and is requesting (application/json currently supported) | [default to &quot;application/json&quot;]
 **oclcNumber** | **int64** | oclc number of the bibliographic item | 
 **isbn** | **string** | isbn of the bibliographic item (bn index) | 
 **issn** | **string** | issn of the bibliographic item (in index) | 
 **holdingsAllEditions** | **bool** | Get holdings for all editions | 
 **holdingsAllVariantRecords** | **bool** | Get holdings for specific edition across variant records | 
 **preferredLanguage** | **string** | language user would prefer metadata description in | 
 **heldInCountry** | **string** | Restrict to holdings held by institutions in requested country | 
 **heldByGroup** | **string** | Restrict to holdings held by Group Symbol | 
 **heldBy** | **string** | Institution Symbol | 
 **heldBySymbol** | **[]string** | Institution Symbol | 
 **heldByInstitutionID** | **[]int32** | Institution registryId | 
 **lat** | **float64** | latitude | 
 **lon** | **float64** | longitude | 
 **distance** | **int32** | distance | 
 **unit** | **string** | distance unit | [default to &quot;M&quot;]

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

