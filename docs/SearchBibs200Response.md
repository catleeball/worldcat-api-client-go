# SearchBibs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfRecords** | Pointer to **int32** |  | [optional] 
**BibRecords** | Pointer to [**[]SearchBibs200ResponseBibRecordsInner**](SearchBibs200ResponseBibRecordsInner.md) |  | [optional] 
**SearchFacets** | Pointer to [**[]FindBibRetainedHoldings200ResponseSearchFacetsInner**](FindBibRetainedHoldings200ResponseSearchFacetsInner.md) |  | [optional] 

## Methods

### NewSearchBibs200Response

`func NewSearchBibs200Response() *SearchBibs200Response`

NewSearchBibs200Response instantiates a new SearchBibs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchBibs200ResponseWithDefaults

`func NewSearchBibs200ResponseWithDefaults() *SearchBibs200Response`

NewSearchBibs200ResponseWithDefaults instantiates a new SearchBibs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfRecords

`func (o *SearchBibs200Response) GetNumberOfRecords() int32`

GetNumberOfRecords returns the NumberOfRecords field if non-nil, zero value otherwise.

### GetNumberOfRecordsOk

`func (o *SearchBibs200Response) GetNumberOfRecordsOk() (*int32, bool)`

GetNumberOfRecordsOk returns a tuple with the NumberOfRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfRecords

`func (o *SearchBibs200Response) SetNumberOfRecords(v int32)`

SetNumberOfRecords sets NumberOfRecords field to given value.

### HasNumberOfRecords

`func (o *SearchBibs200Response) HasNumberOfRecords() bool`

HasNumberOfRecords returns a boolean if a field has been set.

### GetBibRecords

`func (o *SearchBibs200Response) GetBibRecords() []SearchBibs200ResponseBibRecordsInner`

GetBibRecords returns the BibRecords field if non-nil, zero value otherwise.

### GetBibRecordsOk

`func (o *SearchBibs200Response) GetBibRecordsOk() (*[]SearchBibs200ResponseBibRecordsInner, bool)`

GetBibRecordsOk returns a tuple with the BibRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBibRecords

`func (o *SearchBibs200Response) SetBibRecords(v []SearchBibs200ResponseBibRecordsInner)`

SetBibRecords sets BibRecords field to given value.

### HasBibRecords

`func (o *SearchBibs200Response) HasBibRecords() bool`

HasBibRecords returns a boolean if a field has been set.

### GetSearchFacets

`func (o *SearchBibs200Response) GetSearchFacets() []FindBibRetainedHoldings200ResponseSearchFacetsInner`

GetSearchFacets returns the SearchFacets field if non-nil, zero value otherwise.

### GetSearchFacetsOk

`func (o *SearchBibs200Response) GetSearchFacetsOk() (*[]FindBibRetainedHoldings200ResponseSearchFacetsInner, bool)`

GetSearchFacetsOk returns a tuple with the SearchFacets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFacets

`func (o *SearchBibs200Response) SetSearchFacets(v []FindBibRetainedHoldings200ResponseSearchFacetsInner)`

SetSearchFacets sets SearchFacets field to given value.

### HasSearchFacets

`func (o *SearchBibs200Response) HasSearchFacets() bool`

HasSearchFacets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


