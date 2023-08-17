# SearchLhr200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfHoldings** | Pointer to **int32** |  | [optional] 
**DetailedHoldings** | Pointer to [**[]FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1**](FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1.md) |  | [optional] 

## Methods

### NewSearchLhr200Response

`func NewSearchLhr200Response() *SearchLhr200Response`

NewSearchLhr200Response instantiates a new SearchLhr200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchLhr200ResponseWithDefaults

`func NewSearchLhr200ResponseWithDefaults() *SearchLhr200Response`

NewSearchLhr200ResponseWithDefaults instantiates a new SearchLhr200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfHoldings

`func (o *SearchLhr200Response) GetNumberOfHoldings() int32`

GetNumberOfHoldings returns the NumberOfHoldings field if non-nil, zero value otherwise.

### GetNumberOfHoldingsOk

`func (o *SearchLhr200Response) GetNumberOfHoldingsOk() (*int32, bool)`

GetNumberOfHoldingsOk returns a tuple with the NumberOfHoldings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfHoldings

`func (o *SearchLhr200Response) SetNumberOfHoldings(v int32)`

SetNumberOfHoldings sets NumberOfHoldings field to given value.

### HasNumberOfHoldings

`func (o *SearchLhr200Response) HasNumberOfHoldings() bool`

HasNumberOfHoldings returns a boolean if a field has been set.

### GetDetailedHoldings

`func (o *SearchLhr200Response) GetDetailedHoldings() []FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1`

GetDetailedHoldings returns the DetailedHoldings field if non-nil, zero value otherwise.

### GetDetailedHoldingsOk

`func (o *SearchLhr200Response) GetDetailedHoldingsOk() (*[]FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1, bool)`

GetDetailedHoldingsOk returns a tuple with the DetailedHoldings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedHoldings

`func (o *SearchLhr200Response) SetDetailedHoldings(v []FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1)`

SetDetailedHoldings sets DetailedHoldings field to given value.

### HasDetailedHoldings

`func (o *SearchLhr200Response) HasDetailedHoldings() bool`

HasDetailedHoldings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


