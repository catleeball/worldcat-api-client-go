# SearchBibs200ResponseBibRecordsInnerSubjectsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectName** | Pointer to [**SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner**](SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner.md) |  | [optional] 
**Vocabulary** | Pointer to **string** | subject authority vocabularies (MESH, FAST, LCSH, RVM) | [optional] 
**Relators** | Pointer to [**[]SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner**](SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner.md) |  | [optional] 
**SubjectType** | Pointer to **string** |  | [optional] 
**Uri** | Pointer to **string** | URI [6xx s1] | [optional] 

## Methods

### NewSearchBibs200ResponseBibRecordsInnerSubjectsInner

`func NewSearchBibs200ResponseBibRecordsInnerSubjectsInner() *SearchBibs200ResponseBibRecordsInnerSubjectsInner`

NewSearchBibs200ResponseBibRecordsInnerSubjectsInner instantiates a new SearchBibs200ResponseBibRecordsInnerSubjectsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchBibs200ResponseBibRecordsInnerSubjectsInnerWithDefaults

`func NewSearchBibs200ResponseBibRecordsInnerSubjectsInnerWithDefaults() *SearchBibs200ResponseBibRecordsInnerSubjectsInner`

NewSearchBibs200ResponseBibRecordsInnerSubjectsInnerWithDefaults instantiates a new SearchBibs200ResponseBibRecordsInnerSubjectsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjectName

`func (o *SearchBibs200ResponseBibRecordsInnerSubjectsInner) GetSubjectName() SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner`

GetSubjectName returns the SubjectName field if non-nil, zero value otherwise.

### GetSubjectNameOk

`func (o *SearchBibs200ResponseBibRecordsInnerSubjectsInner) GetSubjectNameOk() (*SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner, bool)`

GetSubjectNameOk returns a tuple with the SubjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectName

`func (o *SearchBibs200ResponseBibRecordsInnerSubjectsInner) SetSubjectName(v SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner)`

SetSubjectName sets SubjectName field to given value.

### HasSubjectName

`func (o *SearchBibs200ResponseBibRecordsInnerSubjectsInner) HasSubjectName() bool`

HasSubjectName returns a boolean if a field has been set.

### GetVocabulary

`func (o *SearchBibs200ResponseBibRecordsInnerSubjectsInner) GetVocabulary() string`

GetVocabulary returns the Vocabulary field if non-nil, zero value otherwise.

### GetVocabularyOk

`func (o *SearchBibs200ResponseBibRecordsInnerSubjectsInner) GetVocabularyOk() (*string, bool)`

GetVocabularyOk returns a tuple with the Vocabulary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVocabulary

`func (o *SearchBibs200ResponseBibRecordsInnerSubjectsInner) SetVocabulary(v string)`

SetVocabulary sets Vocabulary field to given value.

### HasVocabulary

`func (o *SearchBibs200ResponseBibRecordsInnerSubjectsInner) HasVocabulary() bool`

HasVocabulary returns a boolean if a field has been set.

### GetRelators

`func (o *SearchBibs200ResponseBibRecordsInnerSubjectsInner) GetRelators() []SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner`

GetRelators returns the Relators field if non-nil, zero value otherwise.

### GetRelatorsOk

`func (o *SearchBibs200ResponseBibRecordsInnerSubjectsInner) GetRelatorsOk() (*[]SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner, bool)`

GetRelatorsOk returns a tuple with the Relators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelators

`func (o *SearchBibs200ResponseBibRecordsInnerSubjectsInner) SetRelators(v []SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner)`

SetRelators sets Relators field to given value.

### HasRelators

`func (o *SearchBibs200ResponseBibRecordsInnerSubjectsInner) HasRelators() bool`

HasRelators returns a boolean if a field has been set.

### GetSubjectType

`func (o *SearchBibs200ResponseBibRecordsInnerSubjectsInner) GetSubjectType() string`

GetSubjectType returns the SubjectType field if non-nil, zero value otherwise.

### GetSubjectTypeOk

`func (o *SearchBibs200ResponseBibRecordsInnerSubjectsInner) GetSubjectTypeOk() (*string, bool)`

GetSubjectTypeOk returns a tuple with the SubjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectType

`func (o *SearchBibs200ResponseBibRecordsInnerSubjectsInner) SetSubjectType(v string)`

SetSubjectType sets SubjectType field to given value.

### HasSubjectType

`func (o *SearchBibs200ResponseBibRecordsInnerSubjectsInner) HasSubjectType() bool`

HasSubjectType returns a boolean if a field has been set.

### GetUri

`func (o *SearchBibs200ResponseBibRecordsInnerSubjectsInner) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *SearchBibs200ResponseBibRecordsInnerSubjectsInner) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *SearchBibs200ResponseBibRecordsInnerSubjectsInner) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *SearchBibs200ResponseBibRecordsInnerSubjectsInner) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


