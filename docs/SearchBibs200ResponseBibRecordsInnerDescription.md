# SearchBibs200ResponseBibRecordsInnerDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhysicalDescription** | Pointer to **string** | Physical Description [v300sa,b,c,d,e,f,g,3] | [optional] 
**DigitalGraphicRepresentation** | Pointer to **string** | Digital Graphic Representation [v352sa,b,c,i,q] | [optional] 
**Genres** | Pointer to **[]string** | Genre of work [v655sa,b,v,x,y,z v380sa, (v600,v610,v611,v630,v648, v650,v651,v654 sv) v653sa] | [optional] 
**CartographicData** | Pointer to **[]string** | Cartographic Mathematical Data [v255sa,b,c] | [optional] 
**ComputerFilesCharacteristics** | Pointer to **[]string** | Computer File Characteristics [v256sa] | [optional] 
**FormOfWorks** | Pointer to **[]string** | Form of Work [v380sa] | [optional] 
**Abstracts** | Pointer to [**[]SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner**](SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner.md) | Abstract [v520sa,b,c] | [optional] 
**Summaries** | Pointer to [**[]SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner**](SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner.md) | Summary [v520sa,b,c] | [optional] 
**Contents** | Pointer to [**[]SearchBibs200ResponseBibRecordsInnerDescriptionContentsInner**](SearchBibs200ResponseBibRecordsInnerDescriptionContentsInner.md) |  | [optional] 
**Bibliographies** | Pointer to [**[]SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner**](SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner.md) | Bibliography [v504sa] | [optional] 
**PeerReviewed** | Pointer to **string** | Is this a peer reviewed work [Admin/OCLCDef/PR] | [optional] 

## Methods

### NewSearchBibs200ResponseBibRecordsInnerDescription

`func NewSearchBibs200ResponseBibRecordsInnerDescription() *SearchBibs200ResponseBibRecordsInnerDescription`

NewSearchBibs200ResponseBibRecordsInnerDescription instantiates a new SearchBibs200ResponseBibRecordsInnerDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchBibs200ResponseBibRecordsInnerDescriptionWithDefaults

`func NewSearchBibs200ResponseBibRecordsInnerDescriptionWithDefaults() *SearchBibs200ResponseBibRecordsInnerDescription`

NewSearchBibs200ResponseBibRecordsInnerDescriptionWithDefaults instantiates a new SearchBibs200ResponseBibRecordsInnerDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhysicalDescription

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetPhysicalDescription() string`

GetPhysicalDescription returns the PhysicalDescription field if non-nil, zero value otherwise.

### GetPhysicalDescriptionOk

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetPhysicalDescriptionOk() (*string, bool)`

GetPhysicalDescriptionOk returns a tuple with the PhysicalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDescription

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) SetPhysicalDescription(v string)`

SetPhysicalDescription sets PhysicalDescription field to given value.

### HasPhysicalDescription

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) HasPhysicalDescription() bool`

HasPhysicalDescription returns a boolean if a field has been set.

### GetDigitalGraphicRepresentation

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetDigitalGraphicRepresentation() string`

GetDigitalGraphicRepresentation returns the DigitalGraphicRepresentation field if non-nil, zero value otherwise.

### GetDigitalGraphicRepresentationOk

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetDigitalGraphicRepresentationOk() (*string, bool)`

GetDigitalGraphicRepresentationOk returns a tuple with the DigitalGraphicRepresentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalGraphicRepresentation

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) SetDigitalGraphicRepresentation(v string)`

SetDigitalGraphicRepresentation sets DigitalGraphicRepresentation field to given value.

### HasDigitalGraphicRepresentation

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) HasDigitalGraphicRepresentation() bool`

HasDigitalGraphicRepresentation returns a boolean if a field has been set.

### GetGenres

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetGenres() []string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetGenresOk() (*[]string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) SetGenres(v []string)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### GetCartographicData

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetCartographicData() []string`

GetCartographicData returns the CartographicData field if non-nil, zero value otherwise.

### GetCartographicDataOk

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetCartographicDataOk() (*[]string, bool)`

GetCartographicDataOk returns a tuple with the CartographicData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartographicData

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) SetCartographicData(v []string)`

SetCartographicData sets CartographicData field to given value.

### HasCartographicData

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) HasCartographicData() bool`

HasCartographicData returns a boolean if a field has been set.

### GetComputerFilesCharacteristics

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetComputerFilesCharacteristics() []string`

GetComputerFilesCharacteristics returns the ComputerFilesCharacteristics field if non-nil, zero value otherwise.

### GetComputerFilesCharacteristicsOk

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetComputerFilesCharacteristicsOk() (*[]string, bool)`

GetComputerFilesCharacteristicsOk returns a tuple with the ComputerFilesCharacteristics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerFilesCharacteristics

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) SetComputerFilesCharacteristics(v []string)`

SetComputerFilesCharacteristics sets ComputerFilesCharacteristics field to given value.

### HasComputerFilesCharacteristics

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) HasComputerFilesCharacteristics() bool`

HasComputerFilesCharacteristics returns a boolean if a field has been set.

### GetFormOfWorks

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetFormOfWorks() []string`

GetFormOfWorks returns the FormOfWorks field if non-nil, zero value otherwise.

### GetFormOfWorksOk

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetFormOfWorksOk() (*[]string, bool)`

GetFormOfWorksOk returns a tuple with the FormOfWorks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormOfWorks

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) SetFormOfWorks(v []string)`

SetFormOfWorks sets FormOfWorks field to given value.

### HasFormOfWorks

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) HasFormOfWorks() bool`

HasFormOfWorks returns a boolean if a field has been set.

### GetAbstracts

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetAbstracts() []SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner`

GetAbstracts returns the Abstracts field if non-nil, zero value otherwise.

### GetAbstractsOk

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetAbstractsOk() (*[]SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner, bool)`

GetAbstractsOk returns a tuple with the Abstracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstracts

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) SetAbstracts(v []SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner)`

SetAbstracts sets Abstracts field to given value.

### HasAbstracts

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) HasAbstracts() bool`

HasAbstracts returns a boolean if a field has been set.

### GetSummaries

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetSummaries() []SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner`

GetSummaries returns the Summaries field if non-nil, zero value otherwise.

### GetSummariesOk

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetSummariesOk() (*[]SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner, bool)`

GetSummariesOk returns a tuple with the Summaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaries

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) SetSummaries(v []SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner)`

SetSummaries sets Summaries field to given value.

### HasSummaries

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) HasSummaries() bool`

HasSummaries returns a boolean if a field has been set.

### GetContents

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetContents() []SearchBibs200ResponseBibRecordsInnerDescriptionContentsInner`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetContentsOk() (*[]SearchBibs200ResponseBibRecordsInnerDescriptionContentsInner, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) SetContents(v []SearchBibs200ResponseBibRecordsInnerDescriptionContentsInner)`

SetContents sets Contents field to given value.

### HasContents

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetBibliographies

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetBibliographies() []SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner`

GetBibliographies returns the Bibliographies field if non-nil, zero value otherwise.

### GetBibliographiesOk

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetBibliographiesOk() (*[]SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner, bool)`

GetBibliographiesOk returns a tuple with the Bibliographies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBibliographies

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) SetBibliographies(v []SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner)`

SetBibliographies sets Bibliographies field to given value.

### HasBibliographies

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) HasBibliographies() bool`

HasBibliographies returns a boolean if a field has been set.

### GetPeerReviewed

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetPeerReviewed() string`

GetPeerReviewed returns the PeerReviewed field if non-nil, zero value otherwise.

### GetPeerReviewedOk

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetPeerReviewedOk() (*string, bool)`

GetPeerReviewedOk returns a tuple with the PeerReviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerReviewed

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) SetPeerReviewed(v string)`

SetPeerReviewed sets PeerReviewed field to given value.

### HasPeerReviewed

`func (o *SearchBibs200ResponseBibRecordsInnerDescription) HasPeerReviewed() bool`

HasPeerReviewed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


