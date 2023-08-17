# \BibliographicResourcesApi

All URIs are relative to *https://americas.discovery.api.oclc.org/worldcat/search/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveBib**](BibliographicResourcesApi.md#RetrieveBib) | **Get** /bibs/{oclcNumber} | Retrieve specific Bibliographic Resource
[**RetrieveBriefBib**](BibliographicResourcesApi.md#RetrieveBriefBib) | **Get** /brief-bibs/{oclcNumber} | Retrieve specific Bibliographic Resource
[**RetrieveOtherEditions**](BibliographicResourcesApi.md#RetrieveOtherEditions) | **Get** /brief-bibs/{oclcNumber}/other-editions | Retrieve Other Editions related to a particular Bibliographic Resource
[**SearchBibs**](BibliographicResourcesApi.md#SearchBibs) | **Get** /bibs | Search Bibliographic Resources
[**SearchBriefBibs**](BibliographicResourcesApi.md#SearchBriefBibs) | **Get** /brief-bibs | Search Brief Bibliographic Resources



## RetrieveBib

> SearchBibs200ResponseBibRecordsInner RetrieveBib(ctx, oclcNumber).Accept(accept).Execute()

Retrieve specific Bibliographic Resource



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
    oclcNumber := int64(41266045) // int64 | record OCLC number
    accept := "accept_example" // string | Content Type the client supports and is requesting (application/json currently supported) (optional) (default to "application/json")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BibliographicResourcesApi.RetrieveBib(context.Background(), oclcNumber).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BibliographicResourcesApi.RetrieveBib``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveBib`: SearchBibs200ResponseBibRecordsInner
    fmt.Fprintf(os.Stdout, "Response from `BibliographicResourcesApi.RetrieveBib`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oclcNumber** | **int64** | record OCLC number | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveBibRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** | Content Type the client supports and is requesting (application/json currently supported) | [default to &quot;application/json&quot;]

### Return type

[**SearchBibs200ResponseBibRecordsInner**](SearchBibs200ResponseBibRecordsInner.md)

### Authorization

[worldcat_search_auth](../README.md#worldcat_search_auth), [worldcat_search_auth](../README.md#worldcat_search_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveBriefBib

> FindBibRetainedHoldings200ResponseBriefRecordsInner RetrieveBriefBib(ctx, oclcNumber).Accept(accept).ShowHoldingsIndicators(showHoldingsIndicators).Execute()

Retrieve specific Bibliographic Resource



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
    oclcNumber := int64(41266045) // int64 | record OCLC number
    accept := "accept_example" // string | Content Type the client supports and is requesting (application/json currently supported) (optional) (default to "application/json")
    showHoldingsIndicators := true // bool | whether or not to show holdings indicators in response (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BibliographicResourcesApi.RetrieveBriefBib(context.Background(), oclcNumber).Accept(accept).ShowHoldingsIndicators(showHoldingsIndicators).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BibliographicResourcesApi.RetrieveBriefBib``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveBriefBib`: FindBibRetainedHoldings200ResponseBriefRecordsInner
    fmt.Fprintf(os.Stdout, "Response from `BibliographicResourcesApi.RetrieveBriefBib`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oclcNumber** | **int64** | record OCLC number | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveBriefBibRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** | Content Type the client supports and is requesting (application/json currently supported) | [default to &quot;application/json&quot;]
 **showHoldingsIndicators** | **bool** | whether or not to show holdings indicators in response | [default to false]

### Return type

[**FindBibRetainedHoldings200ResponseBriefRecordsInner**](FindBibRetainedHoldings200ResponseBriefRecordsInner.md)

### Authorization

[worldcat_search_auth](../README.md#worldcat_search_auth), [worldcat_search_auth](../README.md#worldcat_search_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveOtherEditions

> FindBibRetainedHoldings200Response RetrieveOtherEditions(ctx, oclcNumber).Accept(accept).DeweyNumber(deweyNumber).DatePublished(datePublished).HeldByGroup(heldByGroup).HeldBy(heldBy).HeldBySymbol(heldBySymbol).HeldByInstitutionID(heldByInstitutionID).InLanguage(inLanguage).InCatalogLanguage(inCatalogLanguage).MaterialType(materialType).CatalogSource(catalogSource).ItemType(itemType).ItemSubType(itemSubType).RetentionCommitments(retentionCommitments).SpProgram(spProgram).Genre(genre).Topic(topic).Subtopic(subtopic).Audience(audience).Content(content).OpenAccess(openAccess).PeerReviewed(peerReviewed).Facets(facets).GroupVariantRecords(groupVariantRecords).PreferredLanguage(preferredLanguage).ShowHoldingsIndicators(showHoldingsIndicators).Offset(offset).Limit(limit).OrderBy(orderBy).Execute()

Retrieve Other Editions related to a particular Bibliographic Resource



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
    oclcNumber := int64(41266045) // int64 | record OCLC number
    accept := "accept_example" // string | Content Type the client supports and is requesting (application/json currently supported) (optional) (default to "application/json")
    deweyNumber := []string{"Inner_example"} // []string | Limiter to restrict the response to the specified dewey classification number(s). For multiple values repeat the URL parameter. (dd index) (optional)
    datePublished := []string{"Inner_example"} // []string | Limiter to restrict the response to one or more dates or to a range (in index). (2000, 2000-2010, 1990,1991) (optional)
    heldByGroup := "ASRL" // string | Restrict to holdings held by Group Symbol (optional)
    heldBy := "TXH" // string | Institution Symbol (optional)
    heldBySymbol := []string{"Inner_example"} // []string | Institution Symbol (optional)
    heldByInstitutionID := []int32{int32(123)} // []int32 | Institution registryId (optional)
    inLanguage := []string{"Inner_example"} // []string | Limiter to restrict the response to the single specified language (ln index). (eng, fre, ger) (optional)
    inCatalogLanguage := "eng" // string | Limiter to restrict the response to the single specified cataloging language (ll index). (dut, eng) (optional)
    materialType := "URL" // string | Limiter to restrict the response to the specified material type (mt index). (URL, LPS) (optional)
    catalogSource := "OCL" // string | Limiter to restrict the response to the single OCLC symbol as the cataloging source (cs index). (OSL) (optional)
    itemType := []string{"ItemType_example"} // []string | Limiter to restrict the response to the single specified OCLC top-level facet type (x0 index). (book, image) (optional)
    itemSubType := []string{"ItemSubType_example"} // []string | Limiter to restrict the response to the single specified OCLC sub facet type (x1 index). (dvd, digital) (optional)
    retentionCommitments := true // bool | Limiter to restrict the response to bibliographic records with a retention commitment (l8 index). (optional) (default to false)
    spProgram := "WEST" // string | Limiter to restrict the response to bibliographic records associated with a particular shared print program (l8 index). (optional)
    genre := "fiction" // string | Genre to limit results to (ge index) (optional)
    topic := "34000000" // string | topic to limit results to (s0 index). Based on OCLC Conspectus Division (optional)
    subtopic := "34932000" // string | subtopic to limit results to (s1 index). Based on OCLC Conspectus Category (optional)
    audience := "juv" // string | Audience to limit results to (optional)
    content := []string{"Content_example"} // []string | Content to limit results to (optional)
    openAccess := true // bool | filter to just open access content (optional)
    peerReviewed := true // bool | filter to just peer reviewed content (optional)
    facets := []string{"Facets_example"} // []string | The requested list of facets requested in the response (optional)
    groupVariantRecords := true // bool | Whether or not to group variant records. Default is 'false' (optional) (default to false)
    preferredLanguage := "fre" // string | language user would prefer metadata description in (optional)
    showHoldingsIndicators := true // bool | whether or not to show holdings indicators in response (optional) (default to false)
    offset := int32(10) // int32 | start position of the bib records to return (0 based), default 0 (optional) (default to 0)
    limit := int32(50) // int32 | maximum number of records to return, maximum 50, default 10 (optional) (default to 10)
    orderBy := "language" // string | how to sort the edition resultsresults (optional) (default to "publicationDateDesc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BibliographicResourcesApi.RetrieveOtherEditions(context.Background(), oclcNumber).Accept(accept).DeweyNumber(deweyNumber).DatePublished(datePublished).HeldByGroup(heldByGroup).HeldBy(heldBy).HeldBySymbol(heldBySymbol).HeldByInstitutionID(heldByInstitutionID).InLanguage(inLanguage).InCatalogLanguage(inCatalogLanguage).MaterialType(materialType).CatalogSource(catalogSource).ItemType(itemType).ItemSubType(itemSubType).RetentionCommitments(retentionCommitments).SpProgram(spProgram).Genre(genre).Topic(topic).Subtopic(subtopic).Audience(audience).Content(content).OpenAccess(openAccess).PeerReviewed(peerReviewed).Facets(facets).GroupVariantRecords(groupVariantRecords).PreferredLanguage(preferredLanguage).ShowHoldingsIndicators(showHoldingsIndicators).Offset(offset).Limit(limit).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BibliographicResourcesApi.RetrieveOtherEditions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveOtherEditions`: FindBibRetainedHoldings200Response
    fmt.Fprintf(os.Stdout, "Response from `BibliographicResourcesApi.RetrieveOtherEditions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oclcNumber** | **int64** | record OCLC number | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveOtherEditionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** | Content Type the client supports and is requesting (application/json currently supported) | [default to &quot;application/json&quot;]
 **deweyNumber** | **[]string** | Limiter to restrict the response to the specified dewey classification number(s). For multiple values repeat the URL parameter. (dd index) | 
 **datePublished** | **[]string** | Limiter to restrict the response to one or more dates or to a range (in index). (2000, 2000-2010, 1990,1991) | 
 **heldByGroup** | **string** | Restrict to holdings held by Group Symbol | 
 **heldBy** | **string** | Institution Symbol | 
 **heldBySymbol** | **[]string** | Institution Symbol | 
 **heldByInstitutionID** | **[]int32** | Institution registryId | 
 **inLanguage** | **[]string** | Limiter to restrict the response to the single specified language (ln index). (eng, fre, ger) | 
 **inCatalogLanguage** | **string** | Limiter to restrict the response to the single specified cataloging language (ll index). (dut, eng) | 
 **materialType** | **string** | Limiter to restrict the response to the specified material type (mt index). (URL, LPS) | 
 **catalogSource** | **string** | Limiter to restrict the response to the single OCLC symbol as the cataloging source (cs index). (OSL) | 
 **itemType** | **[]string** | Limiter to restrict the response to the single specified OCLC top-level facet type (x0 index). (book, image) | 
 **itemSubType** | **[]string** | Limiter to restrict the response to the single specified OCLC sub facet type (x1 index). (dvd, digital) | 
 **retentionCommitments** | **bool** | Limiter to restrict the response to bibliographic records with a retention commitment (l8 index). | [default to false]
 **spProgram** | **string** | Limiter to restrict the response to bibliographic records associated with a particular shared print program (l8 index). | 
 **genre** | **string** | Genre to limit results to (ge index) | 
 **topic** | **string** | topic to limit results to (s0 index). Based on OCLC Conspectus Division | 
 **subtopic** | **string** | subtopic to limit results to (s1 index). Based on OCLC Conspectus Category | 
 **audience** | **string** | Audience to limit results to | 
 **content** | **[]string** | Content to limit results to | 
 **openAccess** | **bool** | filter to just open access content | 
 **peerReviewed** | **bool** | filter to just peer reviewed content | 
 **facets** | **[]string** | The requested list of facets requested in the response | 
 **groupVariantRecords** | **bool** | Whether or not to group variant records. Default is &#39;false&#39; | [default to false]
 **preferredLanguage** | **string** | language user would prefer metadata description in | 
 **showHoldingsIndicators** | **bool** | whether or not to show holdings indicators in response | [default to false]
 **offset** | **int32** | start position of the bib records to return (0 based), default 0 | [default to 0]
 **limit** | **int32** | maximum number of records to return, maximum 50, default 10 | [default to 10]
 **orderBy** | **string** | how to sort the edition resultsresults | [default to &quot;publicationDateDesc&quot;]

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


## SearchBibs

> SearchBibs200Response SearchBibs(ctx).Q(q).Accept(accept).DeweyNumber(deweyNumber).DatePublished(datePublished).HeldByGroup(heldByGroup).HeldBy(heldBy).HeldBySymbol(heldBySymbol).HeldByInstitutionID(heldByInstitutionID).InLanguage(inLanguage).InCatalogLanguage(inCatalogLanguage).MaterialType(materialType).CatalogSource(catalogSource).ItemType(itemType).ItemSubType(itemSubType).RetentionCommitments(retentionCommitments).SpProgram(spProgram).Facets(facets).GroupRelatedEditions(groupRelatedEditions).GroupVariantRecords(groupVariantRecords).PreferredLanguage(preferredLanguage).OrderBy(orderBy).Offset(offset).Limit(limit).Execute()

Search Bibliographic Resources



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
    q := "ti: Simon's Cat" // string | query in the form of a keyword search or fielded search, for example  (au:gazihan OR au:womacka) AND kw:java See - https://help.oclc.org/Librarian_Toolbox/Searching_WorldCat_Indexes/Bibliographic_records/Bibliographic_record_indexes/Bibliographic_record_index_lists
    accept := "accept_example" // string | Content Type the client supports and is requesting (application/json currently supported) (optional) (default to "application/json")
    deweyNumber := []string{"Inner_example"} // []string | Limiter to restrict the response to the specified dewey classification number(s). For multiple values repeat the URL parameter. (dd index) (optional)
    datePublished := []string{"Inner_example"} // []string | Limiter to restrict the response to one or more dates or to a range (in index). (2000, 2000-2010, 1990,1991) (optional)
    heldByGroup := "ASRL" // string | Restrict to holdings held by Group Symbol (optional)
    heldBy := "TXH" // string | Institution Symbol (optional)
    heldBySymbol := []string{"Inner_example"} // []string | Institution Symbol (optional)
    heldByInstitutionID := []int32{int32(123)} // []int32 | Institution registryId (optional)
    inLanguage := []string{"Inner_example"} // []string | Limiter to restrict the response to the single specified language (ln index). (eng, fre, ger) (optional)
    inCatalogLanguage := "eng" // string | Limiter to restrict the response to the single specified cataloging language (ll index). (dut, eng) (optional)
    materialType := "URL" // string | Limiter to restrict the response to the specified material type (mt index). (URL, LPS) (optional)
    catalogSource := "OCL" // string | Limiter to restrict the response to the single OCLC symbol as the cataloging source (cs index). (OSL) (optional)
    itemType := []string{"ItemType_example"} // []string | Limiter to restrict the response to the single specified OCLC top-level facet type (x0 index). (book, image) (optional)
    itemSubType := []string{"ItemSubType_example"} // []string | Limiter to restrict the response to the single specified OCLC sub facet type (x1 index). (dvd, digital) (optional)
    retentionCommitments := true // bool | Limiter to restrict the response to bibliographic records with a retention commitment (l8 index). (optional) (default to false)
    spProgram := "WEST" // string | Limiter to restrict the response to bibliographic records associated with a particular shared print program (l8 index). (optional)
    facets := []string{"Facets_example"} // []string | The requested list of facets requested in the response (optional)
    groupRelatedEditions := true // bool | Whether or not to use FRBR grouping. Default is 'false' (optional) (default to false)
    groupVariantRecords := true // bool | Whether or not to group variant records. Default is 'false' (optional) (default to false)
    preferredLanguage := "fre" // string | language user would prefer metadata description in (optional)
    orderBy := "orderBy_example" // string | result sort key (optional) (default to "bestMatch")
    offset := int32(10) // int32 | start position of the bib records to return (0 based), default 0 (optional) (default to 0)
    limit := int32(50) // int32 | maximum number of records to return, maximum 50, default 10 (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BibliographicResourcesApi.SearchBibs(context.Background()).Q(q).Accept(accept).DeweyNumber(deweyNumber).DatePublished(datePublished).HeldByGroup(heldByGroup).HeldBy(heldBy).HeldBySymbol(heldBySymbol).HeldByInstitutionID(heldByInstitutionID).InLanguage(inLanguage).InCatalogLanguage(inCatalogLanguage).MaterialType(materialType).CatalogSource(catalogSource).ItemType(itemType).ItemSubType(itemSubType).RetentionCommitments(retentionCommitments).SpProgram(spProgram).Facets(facets).GroupRelatedEditions(groupRelatedEditions).GroupVariantRecords(groupVariantRecords).PreferredLanguage(preferredLanguage).OrderBy(orderBy).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BibliographicResourcesApi.SearchBibs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchBibs`: SearchBibs200Response
    fmt.Fprintf(os.Stdout, "Response from `BibliographicResourcesApi.SearchBibs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchBibsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | query in the form of a keyword search or fielded search, for example  (au:gazihan OR au:womacka) AND kw:java See - https://help.oclc.org/Librarian_Toolbox/Searching_WorldCat_Indexes/Bibliographic_records/Bibliographic_record_indexes/Bibliographic_record_index_lists | 
 **accept** | **string** | Content Type the client supports and is requesting (application/json currently supported) | [default to &quot;application/json&quot;]
 **deweyNumber** | **[]string** | Limiter to restrict the response to the specified dewey classification number(s). For multiple values repeat the URL parameter. (dd index) | 
 **datePublished** | **[]string** | Limiter to restrict the response to one or more dates or to a range (in index). (2000, 2000-2010, 1990,1991) | 
 **heldByGroup** | **string** | Restrict to holdings held by Group Symbol | 
 **heldBy** | **string** | Institution Symbol | 
 **heldBySymbol** | **[]string** | Institution Symbol | 
 **heldByInstitutionID** | **[]int32** | Institution registryId | 
 **inLanguage** | **[]string** | Limiter to restrict the response to the single specified language (ln index). (eng, fre, ger) | 
 **inCatalogLanguage** | **string** | Limiter to restrict the response to the single specified cataloging language (ll index). (dut, eng) | 
 **materialType** | **string** | Limiter to restrict the response to the specified material type (mt index). (URL, LPS) | 
 **catalogSource** | **string** | Limiter to restrict the response to the single OCLC symbol as the cataloging source (cs index). (OSL) | 
 **itemType** | **[]string** | Limiter to restrict the response to the single specified OCLC top-level facet type (x0 index). (book, image) | 
 **itemSubType** | **[]string** | Limiter to restrict the response to the single specified OCLC sub facet type (x1 index). (dvd, digital) | 
 **retentionCommitments** | **bool** | Limiter to restrict the response to bibliographic records with a retention commitment (l8 index). | [default to false]
 **spProgram** | **string** | Limiter to restrict the response to bibliographic records associated with a particular shared print program (l8 index). | 
 **facets** | **[]string** | The requested list of facets requested in the response | 
 **groupRelatedEditions** | **bool** | Whether or not to use FRBR grouping. Default is &#39;false&#39; | [default to false]
 **groupVariantRecords** | **bool** | Whether or not to group variant records. Default is &#39;false&#39; | [default to false]
 **preferredLanguage** | **string** | language user would prefer metadata description in | 
 **orderBy** | **string** | result sort key | [default to &quot;bestMatch&quot;]
 **offset** | **int32** | start position of the bib records to return (0 based), default 0 | [default to 0]
 **limit** | **int32** | maximum number of records to return, maximum 50, default 10 | [default to 10]

### Return type

[**SearchBibs200Response**](SearchBibs200Response.md)

### Authorization

[worldcat_search_auth](../README.md#worldcat_search_auth), [worldcat_search_auth](../README.md#worldcat_search_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchBriefBibs

> FindBibRetainedHoldings200Response SearchBriefBibs(ctx).Q(q).Accept(accept).DeweyNumber(deweyNumber).DatePublished(datePublished).HeldByGroup(heldByGroup).HeldBy(heldBy).HeldBySymbol(heldBySymbol).HeldByInstitutionID(heldByInstitutionID).InLanguage(inLanguage).InCatalogLanguage(inCatalogLanguage).MaterialType(materialType).CatalogSource(catalogSource).ItemType(itemType).ItemSubType(itemSubType).RetentionCommitments(retentionCommitments).SpProgram(spProgram).Facets(facets).GroupRelatedEditions(groupRelatedEditions).GroupVariantRecords(groupVariantRecords).PreferredLanguage(preferredLanguage).ShowHoldingsIndicators(showHoldingsIndicators).OrderBy(orderBy).Offset(offset).Limit(limit).Execute()

Search Brief Bibliographic Resources



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
    q := "ti: Simon's Cat" // string | query in the form of a keyword search or fielded search, for example  (au:gazihan OR au:womacka) AND kw:java See - https://help.oclc.org/Librarian_Toolbox/Searching_WorldCat_Indexes/Bibliographic_records/Bibliographic_record_indexes/Bibliographic_record_index_lists
    accept := "accept_example" // string | Content Type the client supports and is requesting (application/json currently supported) (optional) (default to "application/json")
    deweyNumber := []string{"Inner_example"} // []string | Limiter to restrict the response to the specified dewey classification number(s). For multiple values repeat the URL parameter. (dd index) (optional)
    datePublished := []string{"Inner_example"} // []string | Limiter to restrict the response to one or more dates or to a range (in index). (2000, 2000-2010, 1990,1991) (optional)
    heldByGroup := "ASRL" // string | Restrict to holdings held by Group Symbol (optional)
    heldBy := "TXH" // string | Institution Symbol (optional)
    heldBySymbol := []string{"Inner_example"} // []string | Institution Symbol (optional)
    heldByInstitutionID := []int32{int32(123)} // []int32 | Institution registryId (optional)
    inLanguage := []string{"Inner_example"} // []string | Limiter to restrict the response to the single specified language (ln index). (eng, fre, ger) (optional)
    inCatalogLanguage := "eng" // string | Limiter to restrict the response to the single specified cataloging language (ll index). (dut, eng) (optional)
    materialType := "URL" // string | Limiter to restrict the response to the specified material type (mt index). (URL, LPS) (optional)
    catalogSource := "OCL" // string | Limiter to restrict the response to the single OCLC symbol as the cataloging source (cs index). (OSL) (optional)
    itemType := []string{"ItemType_example"} // []string | Limiter to restrict the response to the single specified OCLC top-level facet type (x0 index). (book, image) (optional)
    itemSubType := []string{"ItemSubType_example"} // []string | Limiter to restrict the response to the single specified OCLC sub facet type (x1 index). (dvd, digital) (optional)
    retentionCommitments := true // bool | Limiter to restrict the response to bibliographic records with a retention commitment (l8 index). (optional) (default to false)
    spProgram := "WEST" // string | Limiter to restrict the response to bibliographic records associated with a particular shared print program (l8 index). (optional)
    facets := []string{"Facets_example"} // []string | The requested list of facets requested in the response (optional)
    groupRelatedEditions := true // bool | Whether or not to use FRBR grouping. Default is 'false' (optional) (default to false)
    groupVariantRecords := true // bool | Whether or not to group variant records. Default is 'false' (optional) (default to false)
    preferredLanguage := "fre" // string | language user would prefer metadata description in (optional)
    showHoldingsIndicators := true // bool | whether or not to show holdings indicators in response (optional) (default to false)
    orderBy := "orderBy_example" // string | result sort key (optional) (default to "bestMatch")
    offset := int32(10) // int32 | start position of the bib records to return (0 based), default 0 (optional) (default to 0)
    limit := int32(50) // int32 | maximum number of records to return, maximum 50, default 10 (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BibliographicResourcesApi.SearchBriefBibs(context.Background()).Q(q).Accept(accept).DeweyNumber(deweyNumber).DatePublished(datePublished).HeldByGroup(heldByGroup).HeldBy(heldBy).HeldBySymbol(heldBySymbol).HeldByInstitutionID(heldByInstitutionID).InLanguage(inLanguage).InCatalogLanguage(inCatalogLanguage).MaterialType(materialType).CatalogSource(catalogSource).ItemType(itemType).ItemSubType(itemSubType).RetentionCommitments(retentionCommitments).SpProgram(spProgram).Facets(facets).GroupRelatedEditions(groupRelatedEditions).GroupVariantRecords(groupVariantRecords).PreferredLanguage(preferredLanguage).ShowHoldingsIndicators(showHoldingsIndicators).OrderBy(orderBy).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BibliographicResourcesApi.SearchBriefBibs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchBriefBibs`: FindBibRetainedHoldings200Response
    fmt.Fprintf(os.Stdout, "Response from `BibliographicResourcesApi.SearchBriefBibs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchBriefBibsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | query in the form of a keyword search or fielded search, for example  (au:gazihan OR au:womacka) AND kw:java See - https://help.oclc.org/Librarian_Toolbox/Searching_WorldCat_Indexes/Bibliographic_records/Bibliographic_record_indexes/Bibliographic_record_index_lists | 
 **accept** | **string** | Content Type the client supports and is requesting (application/json currently supported) | [default to &quot;application/json&quot;]
 **deweyNumber** | **[]string** | Limiter to restrict the response to the specified dewey classification number(s). For multiple values repeat the URL parameter. (dd index) | 
 **datePublished** | **[]string** | Limiter to restrict the response to one or more dates or to a range (in index). (2000, 2000-2010, 1990,1991) | 
 **heldByGroup** | **string** | Restrict to holdings held by Group Symbol | 
 **heldBy** | **string** | Institution Symbol | 
 **heldBySymbol** | **[]string** | Institution Symbol | 
 **heldByInstitutionID** | **[]int32** | Institution registryId | 
 **inLanguage** | **[]string** | Limiter to restrict the response to the single specified language (ln index). (eng, fre, ger) | 
 **inCatalogLanguage** | **string** | Limiter to restrict the response to the single specified cataloging language (ll index). (dut, eng) | 
 **materialType** | **string** | Limiter to restrict the response to the specified material type (mt index). (URL, LPS) | 
 **catalogSource** | **string** | Limiter to restrict the response to the single OCLC symbol as the cataloging source (cs index). (OSL) | 
 **itemType** | **[]string** | Limiter to restrict the response to the single specified OCLC top-level facet type (x0 index). (book, image) | 
 **itemSubType** | **[]string** | Limiter to restrict the response to the single specified OCLC sub facet type (x1 index). (dvd, digital) | 
 **retentionCommitments** | **bool** | Limiter to restrict the response to bibliographic records with a retention commitment (l8 index). | [default to false]
 **spProgram** | **string** | Limiter to restrict the response to bibliographic records associated with a particular shared print program (l8 index). | 
 **facets** | **[]string** | The requested list of facets requested in the response | 
 **groupRelatedEditions** | **bool** | Whether or not to use FRBR grouping. Default is &#39;false&#39; | [default to false]
 **groupVariantRecords** | **bool** | Whether or not to group variant records. Default is &#39;false&#39; | [default to false]
 **preferredLanguage** | **string** | language user would prefer metadata description in | 
 **showHoldingsIndicators** | **bool** | whether or not to show holdings indicators in response | [default to false]
 **orderBy** | **string** | result sort key | [default to &quot;bestMatch&quot;]
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

