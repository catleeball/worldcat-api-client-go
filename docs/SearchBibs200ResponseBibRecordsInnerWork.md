# SearchBibs200ResponseBibRecordsInnerWork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of the workset [Admin/Work/ID] | [optional] 
**Count** | Pointer to **int32** | Number of items in Workset [Admin/Work/Count] | [optional] 
**ArticleCluster** | Pointer to [**[]SearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner**](SearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner.md) | Items in Article Workset [WorkSet/NO] | [optional] 

## Methods

### NewSearchBibs200ResponseBibRecordsInnerWork

`func NewSearchBibs200ResponseBibRecordsInnerWork() *SearchBibs200ResponseBibRecordsInnerWork`

NewSearchBibs200ResponseBibRecordsInnerWork instantiates a new SearchBibs200ResponseBibRecordsInnerWork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchBibs200ResponseBibRecordsInnerWorkWithDefaults

`func NewSearchBibs200ResponseBibRecordsInnerWorkWithDefaults() *SearchBibs200ResponseBibRecordsInnerWork`

NewSearchBibs200ResponseBibRecordsInnerWorkWithDefaults instantiates a new SearchBibs200ResponseBibRecordsInnerWork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SearchBibs200ResponseBibRecordsInnerWork) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchBibs200ResponseBibRecordsInnerWork) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchBibs200ResponseBibRecordsInnerWork) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SearchBibs200ResponseBibRecordsInnerWork) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCount

`func (o *SearchBibs200ResponseBibRecordsInnerWork) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SearchBibs200ResponseBibRecordsInnerWork) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SearchBibs200ResponseBibRecordsInnerWork) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *SearchBibs200ResponseBibRecordsInnerWork) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetArticleCluster

`func (o *SearchBibs200ResponseBibRecordsInnerWork) GetArticleCluster() []SearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner`

GetArticleCluster returns the ArticleCluster field if non-nil, zero value otherwise.

### GetArticleClusterOk

`func (o *SearchBibs200ResponseBibRecordsInnerWork) GetArticleClusterOk() (*[]SearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner, bool)`

GetArticleClusterOk returns a tuple with the ArticleCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleCluster

`func (o *SearchBibs200ResponseBibRecordsInnerWork) SetArticleCluster(v []SearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner)`

SetArticleCluster sets ArticleCluster field to given value.

### HasArticleCluster

`func (o *SearchBibs200ResponseBibRecordsInnerWork) HasArticleCluster() bool`

HasArticleCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


