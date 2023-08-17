# SearchBibs200ResponseBibRecordsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to [**SearchBibs200ResponseBibRecordsInnerIdentifier**](SearchBibs200ResponseBibRecordsInnerIdentifier.md) |  | [optional] 
**Title** | Pointer to [**SearchBibs200ResponseBibRecordsInnerTitle**](SearchBibs200ResponseBibRecordsInnerTitle.md) |  | [optional] 
**Contributor** | Pointer to [**SearchBibs200ResponseBibRecordsInnerContributor**](SearchBibs200ResponseBibRecordsInnerContributor.md) |  | [optional] 
**Subjects** | Pointer to [**[]SearchBibs200ResponseBibRecordsInnerSubjectsInner**](SearchBibs200ResponseBibRecordsInnerSubjectsInner.md) |  | [optional] 
**Classification** | Pointer to [**SearchBibs200ResponseBibRecordsInnerClassification**](SearchBibs200ResponseBibRecordsInnerClassification.md) |  | [optional] 
**Publishers** | Pointer to [**[]SearchBibs200ResponseBibRecordsInnerPublishersInner**](SearchBibs200ResponseBibRecordsInnerPublishersInner.md) |  | [optional] 
**Date** | Pointer to [**SearchBibs200ResponseBibRecordsInnerDate**](SearchBibs200ResponseBibRecordsInnerDate.md) |  | [optional] 
**Language** | Pointer to [**SearchBibs200ResponseBibRecordsInnerLanguage**](SearchBibs200ResponseBibRecordsInnerLanguage.md) |  | [optional] 
**Edition** | Pointer to [**SearchBibs200ResponseBibRecordsInnerEdition**](SearchBibs200ResponseBibRecordsInnerEdition.md) |  | [optional] 
**Source** | Pointer to [**[]SearchBibs200ResponseBibRecordsInnerSourceInner**](SearchBibs200ResponseBibRecordsInnerSourceInner.md) |  | [optional] 
**Note** | Pointer to [**SearchBibs200ResponseBibRecordsInnerNote**](SearchBibs200ResponseBibRecordsInnerNote.md) |  | [optional] 
**Format** | Pointer to [**SearchBibs200ResponseBibRecordsInnerFormat**](SearchBibs200ResponseBibRecordsInnerFormat.md) |  | [optional] 
**MusicInfo** | Pointer to [**SearchBibs200ResponseBibRecordsInnerMusicInfo**](SearchBibs200ResponseBibRecordsInnerMusicInfo.md) |  | [optional] 
**DigitalAccessAndLocations** | Pointer to [**[]SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner**](SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner.md) |  | [optional] 
**Description** | Pointer to [**SearchBibs200ResponseBibRecordsInnerDescription**](SearchBibs200ResponseBibRecordsInnerDescription.md) |  | [optional] 
**Related** | Pointer to [**SearchBibs200ResponseBibRecordsInnerRelated**](SearchBibs200ResponseBibRecordsInnerRelated.md) |  | [optional] 
**Work** | Pointer to [**SearchBibs200ResponseBibRecordsInnerWork**](SearchBibs200ResponseBibRecordsInnerWork.md) |  | [optional] 
**EditionCluster** | Pointer to [**SearchBibs200ResponseBibRecordsInnerEditionCluster**](SearchBibs200ResponseBibRecordsInnerEditionCluster.md) |  | [optional] 
**TotalEditions** | Pointer to **int32** | Total number of editions in the workgroup | [optional] 
**Database** | Pointer to [**SearchBibs200ResponseBibRecordsInnerDatabase**](SearchBibs200ResponseBibRecordsInnerDatabase.md) |  | [optional] 
**LocalTitle** | Pointer to [**SearchBibs200ResponseBibRecordsInnerTitle**](SearchBibs200ResponseBibRecordsInnerTitle.md) |  | [optional] 
**LocalContributor** | Pointer to [**SearchBibs200ResponseBibRecordsInnerContributor**](SearchBibs200ResponseBibRecordsInnerContributor.md) |  | [optional] 
**LocalSubjects** | Pointer to [**[]SearchBibs200ResponseBibRecordsInnerSubjectsInner**](SearchBibs200ResponseBibRecordsInnerSubjectsInner.md) |  | [optional] 
**LocalNote** | Pointer to [**SearchBibs200ResponseBibRecordsInnerNote**](SearchBibs200ResponseBibRecordsInnerNote.md) |  | [optional] 
**LocalGenres** | Pointer to **[]string** | Genre of work [v655sa,b,v,x,y,z v380sa, (v600,v610,v611,v630,v648, v650,v651,v654 sv) v653sa] | [optional] 

## Methods

### NewSearchBibs200ResponseBibRecordsInner

`func NewSearchBibs200ResponseBibRecordsInner() *SearchBibs200ResponseBibRecordsInner`

NewSearchBibs200ResponseBibRecordsInner instantiates a new SearchBibs200ResponseBibRecordsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchBibs200ResponseBibRecordsInnerWithDefaults

`func NewSearchBibs200ResponseBibRecordsInnerWithDefaults() *SearchBibs200ResponseBibRecordsInner`

NewSearchBibs200ResponseBibRecordsInnerWithDefaults instantiates a new SearchBibs200ResponseBibRecordsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *SearchBibs200ResponseBibRecordsInner) GetIdentifier() SearchBibs200ResponseBibRecordsInnerIdentifier`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *SearchBibs200ResponseBibRecordsInner) GetIdentifierOk() (*SearchBibs200ResponseBibRecordsInnerIdentifier, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *SearchBibs200ResponseBibRecordsInner) SetIdentifier(v SearchBibs200ResponseBibRecordsInnerIdentifier)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *SearchBibs200ResponseBibRecordsInner) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetTitle

`func (o *SearchBibs200ResponseBibRecordsInner) GetTitle() SearchBibs200ResponseBibRecordsInnerTitle`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SearchBibs200ResponseBibRecordsInner) GetTitleOk() (*SearchBibs200ResponseBibRecordsInnerTitle, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SearchBibs200ResponseBibRecordsInner) SetTitle(v SearchBibs200ResponseBibRecordsInnerTitle)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SearchBibs200ResponseBibRecordsInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetContributor

`func (o *SearchBibs200ResponseBibRecordsInner) GetContributor() SearchBibs200ResponseBibRecordsInnerContributor`

GetContributor returns the Contributor field if non-nil, zero value otherwise.

### GetContributorOk

`func (o *SearchBibs200ResponseBibRecordsInner) GetContributorOk() (*SearchBibs200ResponseBibRecordsInnerContributor, bool)`

GetContributorOk returns a tuple with the Contributor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributor

`func (o *SearchBibs200ResponseBibRecordsInner) SetContributor(v SearchBibs200ResponseBibRecordsInnerContributor)`

SetContributor sets Contributor field to given value.

### HasContributor

`func (o *SearchBibs200ResponseBibRecordsInner) HasContributor() bool`

HasContributor returns a boolean if a field has been set.

### GetSubjects

`func (o *SearchBibs200ResponseBibRecordsInner) GetSubjects() []SearchBibs200ResponseBibRecordsInnerSubjectsInner`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *SearchBibs200ResponseBibRecordsInner) GetSubjectsOk() (*[]SearchBibs200ResponseBibRecordsInnerSubjectsInner, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *SearchBibs200ResponseBibRecordsInner) SetSubjects(v []SearchBibs200ResponseBibRecordsInnerSubjectsInner)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *SearchBibs200ResponseBibRecordsInner) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.

### GetClassification

`func (o *SearchBibs200ResponseBibRecordsInner) GetClassification() SearchBibs200ResponseBibRecordsInnerClassification`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *SearchBibs200ResponseBibRecordsInner) GetClassificationOk() (*SearchBibs200ResponseBibRecordsInnerClassification, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *SearchBibs200ResponseBibRecordsInner) SetClassification(v SearchBibs200ResponseBibRecordsInnerClassification)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *SearchBibs200ResponseBibRecordsInner) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### GetPublishers

`func (o *SearchBibs200ResponseBibRecordsInner) GetPublishers() []SearchBibs200ResponseBibRecordsInnerPublishersInner`

GetPublishers returns the Publishers field if non-nil, zero value otherwise.

### GetPublishersOk

`func (o *SearchBibs200ResponseBibRecordsInner) GetPublishersOk() (*[]SearchBibs200ResponseBibRecordsInnerPublishersInner, bool)`

GetPublishersOk returns a tuple with the Publishers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishers

`func (o *SearchBibs200ResponseBibRecordsInner) SetPublishers(v []SearchBibs200ResponseBibRecordsInnerPublishersInner)`

SetPublishers sets Publishers field to given value.

### HasPublishers

`func (o *SearchBibs200ResponseBibRecordsInner) HasPublishers() bool`

HasPublishers returns a boolean if a field has been set.

### GetDate

`func (o *SearchBibs200ResponseBibRecordsInner) GetDate() SearchBibs200ResponseBibRecordsInnerDate`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *SearchBibs200ResponseBibRecordsInner) GetDateOk() (*SearchBibs200ResponseBibRecordsInnerDate, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *SearchBibs200ResponseBibRecordsInner) SetDate(v SearchBibs200ResponseBibRecordsInnerDate)`

SetDate sets Date field to given value.

### HasDate

`func (o *SearchBibs200ResponseBibRecordsInner) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetLanguage

`func (o *SearchBibs200ResponseBibRecordsInner) GetLanguage() SearchBibs200ResponseBibRecordsInnerLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *SearchBibs200ResponseBibRecordsInner) GetLanguageOk() (*SearchBibs200ResponseBibRecordsInnerLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *SearchBibs200ResponseBibRecordsInner) SetLanguage(v SearchBibs200ResponseBibRecordsInnerLanguage)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *SearchBibs200ResponseBibRecordsInner) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetEdition

`func (o *SearchBibs200ResponseBibRecordsInner) GetEdition() SearchBibs200ResponseBibRecordsInnerEdition`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *SearchBibs200ResponseBibRecordsInner) GetEditionOk() (*SearchBibs200ResponseBibRecordsInnerEdition, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *SearchBibs200ResponseBibRecordsInner) SetEdition(v SearchBibs200ResponseBibRecordsInnerEdition)`

SetEdition sets Edition field to given value.

### HasEdition

`func (o *SearchBibs200ResponseBibRecordsInner) HasEdition() bool`

HasEdition returns a boolean if a field has been set.

### GetSource

`func (o *SearchBibs200ResponseBibRecordsInner) GetSource() []SearchBibs200ResponseBibRecordsInnerSourceInner`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SearchBibs200ResponseBibRecordsInner) GetSourceOk() (*[]SearchBibs200ResponseBibRecordsInnerSourceInner, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SearchBibs200ResponseBibRecordsInner) SetSource(v []SearchBibs200ResponseBibRecordsInnerSourceInner)`

SetSource sets Source field to given value.

### HasSource

`func (o *SearchBibs200ResponseBibRecordsInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetNote

`func (o *SearchBibs200ResponseBibRecordsInner) GetNote() SearchBibs200ResponseBibRecordsInnerNote`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *SearchBibs200ResponseBibRecordsInner) GetNoteOk() (*SearchBibs200ResponseBibRecordsInnerNote, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *SearchBibs200ResponseBibRecordsInner) SetNote(v SearchBibs200ResponseBibRecordsInnerNote)`

SetNote sets Note field to given value.

### HasNote

`func (o *SearchBibs200ResponseBibRecordsInner) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetFormat

`func (o *SearchBibs200ResponseBibRecordsInner) GetFormat() SearchBibs200ResponseBibRecordsInnerFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *SearchBibs200ResponseBibRecordsInner) GetFormatOk() (*SearchBibs200ResponseBibRecordsInnerFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *SearchBibs200ResponseBibRecordsInner) SetFormat(v SearchBibs200ResponseBibRecordsInnerFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *SearchBibs200ResponseBibRecordsInner) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetMusicInfo

`func (o *SearchBibs200ResponseBibRecordsInner) GetMusicInfo() SearchBibs200ResponseBibRecordsInnerMusicInfo`

GetMusicInfo returns the MusicInfo field if non-nil, zero value otherwise.

### GetMusicInfoOk

`func (o *SearchBibs200ResponseBibRecordsInner) GetMusicInfoOk() (*SearchBibs200ResponseBibRecordsInnerMusicInfo, bool)`

GetMusicInfoOk returns a tuple with the MusicInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicInfo

`func (o *SearchBibs200ResponseBibRecordsInner) SetMusicInfo(v SearchBibs200ResponseBibRecordsInnerMusicInfo)`

SetMusicInfo sets MusicInfo field to given value.

### HasMusicInfo

`func (o *SearchBibs200ResponseBibRecordsInner) HasMusicInfo() bool`

HasMusicInfo returns a boolean if a field has been set.

### GetDigitalAccessAndLocations

`func (o *SearchBibs200ResponseBibRecordsInner) GetDigitalAccessAndLocations() []SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner`

GetDigitalAccessAndLocations returns the DigitalAccessAndLocations field if non-nil, zero value otherwise.

### GetDigitalAccessAndLocationsOk

`func (o *SearchBibs200ResponseBibRecordsInner) GetDigitalAccessAndLocationsOk() (*[]SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner, bool)`

GetDigitalAccessAndLocationsOk returns a tuple with the DigitalAccessAndLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalAccessAndLocations

`func (o *SearchBibs200ResponseBibRecordsInner) SetDigitalAccessAndLocations(v []SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner)`

SetDigitalAccessAndLocations sets DigitalAccessAndLocations field to given value.

### HasDigitalAccessAndLocations

`func (o *SearchBibs200ResponseBibRecordsInner) HasDigitalAccessAndLocations() bool`

HasDigitalAccessAndLocations returns a boolean if a field has been set.

### GetDescription

`func (o *SearchBibs200ResponseBibRecordsInner) GetDescription() SearchBibs200ResponseBibRecordsInnerDescription`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SearchBibs200ResponseBibRecordsInner) GetDescriptionOk() (*SearchBibs200ResponseBibRecordsInnerDescription, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SearchBibs200ResponseBibRecordsInner) SetDescription(v SearchBibs200ResponseBibRecordsInnerDescription)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SearchBibs200ResponseBibRecordsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRelated

`func (o *SearchBibs200ResponseBibRecordsInner) GetRelated() SearchBibs200ResponseBibRecordsInnerRelated`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *SearchBibs200ResponseBibRecordsInner) GetRelatedOk() (*SearchBibs200ResponseBibRecordsInnerRelated, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *SearchBibs200ResponseBibRecordsInner) SetRelated(v SearchBibs200ResponseBibRecordsInnerRelated)`

SetRelated sets Related field to given value.

### HasRelated

`func (o *SearchBibs200ResponseBibRecordsInner) HasRelated() bool`

HasRelated returns a boolean if a field has been set.

### GetWork

`func (o *SearchBibs200ResponseBibRecordsInner) GetWork() SearchBibs200ResponseBibRecordsInnerWork`

GetWork returns the Work field if non-nil, zero value otherwise.

### GetWorkOk

`func (o *SearchBibs200ResponseBibRecordsInner) GetWorkOk() (*SearchBibs200ResponseBibRecordsInnerWork, bool)`

GetWorkOk returns a tuple with the Work field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWork

`func (o *SearchBibs200ResponseBibRecordsInner) SetWork(v SearchBibs200ResponseBibRecordsInnerWork)`

SetWork sets Work field to given value.

### HasWork

`func (o *SearchBibs200ResponseBibRecordsInner) HasWork() bool`

HasWork returns a boolean if a field has been set.

### GetEditionCluster

`func (o *SearchBibs200ResponseBibRecordsInner) GetEditionCluster() SearchBibs200ResponseBibRecordsInnerEditionCluster`

GetEditionCluster returns the EditionCluster field if non-nil, zero value otherwise.

### GetEditionClusterOk

`func (o *SearchBibs200ResponseBibRecordsInner) GetEditionClusterOk() (*SearchBibs200ResponseBibRecordsInnerEditionCluster, bool)`

GetEditionClusterOk returns a tuple with the EditionCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditionCluster

`func (o *SearchBibs200ResponseBibRecordsInner) SetEditionCluster(v SearchBibs200ResponseBibRecordsInnerEditionCluster)`

SetEditionCluster sets EditionCluster field to given value.

### HasEditionCluster

`func (o *SearchBibs200ResponseBibRecordsInner) HasEditionCluster() bool`

HasEditionCluster returns a boolean if a field has been set.

### GetTotalEditions

`func (o *SearchBibs200ResponseBibRecordsInner) GetTotalEditions() int32`

GetTotalEditions returns the TotalEditions field if non-nil, zero value otherwise.

### GetTotalEditionsOk

`func (o *SearchBibs200ResponseBibRecordsInner) GetTotalEditionsOk() (*int32, bool)`

GetTotalEditionsOk returns a tuple with the TotalEditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEditions

`func (o *SearchBibs200ResponseBibRecordsInner) SetTotalEditions(v int32)`

SetTotalEditions sets TotalEditions field to given value.

### HasTotalEditions

`func (o *SearchBibs200ResponseBibRecordsInner) HasTotalEditions() bool`

HasTotalEditions returns a boolean if a field has been set.

### GetDatabase

`func (o *SearchBibs200ResponseBibRecordsInner) GetDatabase() SearchBibs200ResponseBibRecordsInnerDatabase`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *SearchBibs200ResponseBibRecordsInner) GetDatabaseOk() (*SearchBibs200ResponseBibRecordsInnerDatabase, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *SearchBibs200ResponseBibRecordsInner) SetDatabase(v SearchBibs200ResponseBibRecordsInnerDatabase)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *SearchBibs200ResponseBibRecordsInner) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetLocalTitle

`func (o *SearchBibs200ResponseBibRecordsInner) GetLocalTitle() SearchBibs200ResponseBibRecordsInnerTitle`

GetLocalTitle returns the LocalTitle field if non-nil, zero value otherwise.

### GetLocalTitleOk

`func (o *SearchBibs200ResponseBibRecordsInner) GetLocalTitleOk() (*SearchBibs200ResponseBibRecordsInnerTitle, bool)`

GetLocalTitleOk returns a tuple with the LocalTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalTitle

`func (o *SearchBibs200ResponseBibRecordsInner) SetLocalTitle(v SearchBibs200ResponseBibRecordsInnerTitle)`

SetLocalTitle sets LocalTitle field to given value.

### HasLocalTitle

`func (o *SearchBibs200ResponseBibRecordsInner) HasLocalTitle() bool`

HasLocalTitle returns a boolean if a field has been set.

### GetLocalContributor

`func (o *SearchBibs200ResponseBibRecordsInner) GetLocalContributor() SearchBibs200ResponseBibRecordsInnerContributor`

GetLocalContributor returns the LocalContributor field if non-nil, zero value otherwise.

### GetLocalContributorOk

`func (o *SearchBibs200ResponseBibRecordsInner) GetLocalContributorOk() (*SearchBibs200ResponseBibRecordsInnerContributor, bool)`

GetLocalContributorOk returns a tuple with the LocalContributor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalContributor

`func (o *SearchBibs200ResponseBibRecordsInner) SetLocalContributor(v SearchBibs200ResponseBibRecordsInnerContributor)`

SetLocalContributor sets LocalContributor field to given value.

### HasLocalContributor

`func (o *SearchBibs200ResponseBibRecordsInner) HasLocalContributor() bool`

HasLocalContributor returns a boolean if a field has been set.

### GetLocalSubjects

`func (o *SearchBibs200ResponseBibRecordsInner) GetLocalSubjects() []SearchBibs200ResponseBibRecordsInnerSubjectsInner`

GetLocalSubjects returns the LocalSubjects field if non-nil, zero value otherwise.

### GetLocalSubjectsOk

`func (o *SearchBibs200ResponseBibRecordsInner) GetLocalSubjectsOk() (*[]SearchBibs200ResponseBibRecordsInnerSubjectsInner, bool)`

GetLocalSubjectsOk returns a tuple with the LocalSubjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSubjects

`func (o *SearchBibs200ResponseBibRecordsInner) SetLocalSubjects(v []SearchBibs200ResponseBibRecordsInnerSubjectsInner)`

SetLocalSubjects sets LocalSubjects field to given value.

### HasLocalSubjects

`func (o *SearchBibs200ResponseBibRecordsInner) HasLocalSubjects() bool`

HasLocalSubjects returns a boolean if a field has been set.

### GetLocalNote

`func (o *SearchBibs200ResponseBibRecordsInner) GetLocalNote() SearchBibs200ResponseBibRecordsInnerNote`

GetLocalNote returns the LocalNote field if non-nil, zero value otherwise.

### GetLocalNoteOk

`func (o *SearchBibs200ResponseBibRecordsInner) GetLocalNoteOk() (*SearchBibs200ResponseBibRecordsInnerNote, bool)`

GetLocalNoteOk returns a tuple with the LocalNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalNote

`func (o *SearchBibs200ResponseBibRecordsInner) SetLocalNote(v SearchBibs200ResponseBibRecordsInnerNote)`

SetLocalNote sets LocalNote field to given value.

### HasLocalNote

`func (o *SearchBibs200ResponseBibRecordsInner) HasLocalNote() bool`

HasLocalNote returns a boolean if a field has been set.

### GetLocalGenres

`func (o *SearchBibs200ResponseBibRecordsInner) GetLocalGenres() []string`

GetLocalGenres returns the LocalGenres field if non-nil, zero value otherwise.

### GetLocalGenresOk

`func (o *SearchBibs200ResponseBibRecordsInner) GetLocalGenresOk() (*[]string, bool)`

GetLocalGenresOk returns a tuple with the LocalGenres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalGenres

`func (o *SearchBibs200ResponseBibRecordsInner) SetLocalGenres(v []string)`

SetLocalGenres sets LocalGenres field to given value.

### HasLocalGenres

`func (o *SearchBibs200ResponseBibRecordsInner) HasLocalGenres() bool`

HasLocalGenres returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


