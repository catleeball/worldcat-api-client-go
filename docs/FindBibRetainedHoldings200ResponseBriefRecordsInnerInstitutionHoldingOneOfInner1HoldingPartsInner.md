# FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkingAndSequenceNumber** | Pointer to **string** | Field link and sequence number [852/863-865/866-868/876-878 | s8] | [optional] 
**SummaryOfHoldings** | Pointer to **string** | Summary of Holdings [https://www.loc.gov/marc/holdings/echdcntf.html] | [optional] 
**Enumerations** | Pointer to [**[]FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInnerEnumerationsInner**](FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInnerEnumerationsInner.md) | Enumerations [853-855/863-865 |a-h] | [optional] 
**Chronologies** | Pointer to [**[]FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInnerEnumerationsInner**](FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInnerEnumerationsInner.md) | Chronologies [853-855/863-865 |i-v] | [optional] 
**TextualHoldings** | Pointer to **string** | [866-8 |a] | [optional] 
**ItemMaterialSpecified** | Pointer to **string** | [876-8 |3] | [optional] 
**PieceDesignation** | Pointer to **string** | Piece Designation [852/863-865/866-868/876-878 |p] | [optional] 
**CancelledPieceDesignations** | Pointer to **[]string** | Cancelled Piece Designation [876-878 |r] | [optional] 
**TemporaryLocation** | Pointer to **string** | Temporary location [876-878 |l] | [optional] 
**PublicNotes** | Pointer to **[]string** | Public notes [852/863-865/866-868/876-878 |z] | [optional] 
**YearRanges** | Pointer to [**[]FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInnerYearRangesInner**](FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInnerYearRangesInner.md) | Year as Range [si] | [optional] 
**VolumeRanges** | Pointer to [**[]FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInnerYearRangesInner**](FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInnerYearRangesInner.md) | Volume Number As Range [sb] | [optional] 

## Methods

### NewFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner

`func NewFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner() *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner`

NewFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner instantiates a new FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInnerWithDefaults

`func NewFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInnerWithDefaults() *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner`

NewFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInnerWithDefaults instantiates a new FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkingAndSequenceNumber

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) GetLinkingAndSequenceNumber() string`

GetLinkingAndSequenceNumber returns the LinkingAndSequenceNumber field if non-nil, zero value otherwise.

### GetLinkingAndSequenceNumberOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) GetLinkingAndSequenceNumberOk() (*string, bool)`

GetLinkingAndSequenceNumberOk returns a tuple with the LinkingAndSequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkingAndSequenceNumber

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) SetLinkingAndSequenceNumber(v string)`

SetLinkingAndSequenceNumber sets LinkingAndSequenceNumber field to given value.

### HasLinkingAndSequenceNumber

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) HasLinkingAndSequenceNumber() bool`

HasLinkingAndSequenceNumber returns a boolean if a field has been set.

### GetSummaryOfHoldings

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) GetSummaryOfHoldings() string`

GetSummaryOfHoldings returns the SummaryOfHoldings field if non-nil, zero value otherwise.

### GetSummaryOfHoldingsOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) GetSummaryOfHoldingsOk() (*string, bool)`

GetSummaryOfHoldingsOk returns a tuple with the SummaryOfHoldings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryOfHoldings

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) SetSummaryOfHoldings(v string)`

SetSummaryOfHoldings sets SummaryOfHoldings field to given value.

### HasSummaryOfHoldings

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) HasSummaryOfHoldings() bool`

HasSummaryOfHoldings returns a boolean if a field has been set.

### GetEnumerations

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) GetEnumerations() []FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInnerEnumerationsInner`

GetEnumerations returns the Enumerations field if non-nil, zero value otherwise.

### GetEnumerationsOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) GetEnumerationsOk() (*[]FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInnerEnumerationsInner, bool)`

GetEnumerationsOk returns a tuple with the Enumerations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumerations

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) SetEnumerations(v []FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInnerEnumerationsInner)`

SetEnumerations sets Enumerations field to given value.

### HasEnumerations

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) HasEnumerations() bool`

HasEnumerations returns a boolean if a field has been set.

### GetChronologies

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) GetChronologies() []FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInnerEnumerationsInner`

GetChronologies returns the Chronologies field if non-nil, zero value otherwise.

### GetChronologiesOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) GetChronologiesOk() (*[]FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInnerEnumerationsInner, bool)`

GetChronologiesOk returns a tuple with the Chronologies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChronologies

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) SetChronologies(v []FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInnerEnumerationsInner)`

SetChronologies sets Chronologies field to given value.

### HasChronologies

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) HasChronologies() bool`

HasChronologies returns a boolean if a field has been set.

### GetTextualHoldings

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) GetTextualHoldings() string`

GetTextualHoldings returns the TextualHoldings field if non-nil, zero value otherwise.

### GetTextualHoldingsOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) GetTextualHoldingsOk() (*string, bool)`

GetTextualHoldingsOk returns a tuple with the TextualHoldings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextualHoldings

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) SetTextualHoldings(v string)`

SetTextualHoldings sets TextualHoldings field to given value.

### HasTextualHoldings

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) HasTextualHoldings() bool`

HasTextualHoldings returns a boolean if a field has been set.

### GetItemMaterialSpecified

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) GetItemMaterialSpecified() string`

GetItemMaterialSpecified returns the ItemMaterialSpecified field if non-nil, zero value otherwise.

### GetItemMaterialSpecifiedOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) GetItemMaterialSpecifiedOk() (*string, bool)`

GetItemMaterialSpecifiedOk returns a tuple with the ItemMaterialSpecified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemMaterialSpecified

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) SetItemMaterialSpecified(v string)`

SetItemMaterialSpecified sets ItemMaterialSpecified field to given value.

### HasItemMaterialSpecified

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) HasItemMaterialSpecified() bool`

HasItemMaterialSpecified returns a boolean if a field has been set.

### GetPieceDesignation

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) GetPieceDesignation() string`

GetPieceDesignation returns the PieceDesignation field if non-nil, zero value otherwise.

### GetPieceDesignationOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) GetPieceDesignationOk() (*string, bool)`

GetPieceDesignationOk returns a tuple with the PieceDesignation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPieceDesignation

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) SetPieceDesignation(v string)`

SetPieceDesignation sets PieceDesignation field to given value.

### HasPieceDesignation

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) HasPieceDesignation() bool`

HasPieceDesignation returns a boolean if a field has been set.

### GetCancelledPieceDesignations

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) GetCancelledPieceDesignations() []string`

GetCancelledPieceDesignations returns the CancelledPieceDesignations field if non-nil, zero value otherwise.

### GetCancelledPieceDesignationsOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) GetCancelledPieceDesignationsOk() (*[]string, bool)`

GetCancelledPieceDesignationsOk returns a tuple with the CancelledPieceDesignations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledPieceDesignations

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) SetCancelledPieceDesignations(v []string)`

SetCancelledPieceDesignations sets CancelledPieceDesignations field to given value.

### HasCancelledPieceDesignations

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) HasCancelledPieceDesignations() bool`

HasCancelledPieceDesignations returns a boolean if a field has been set.

### GetTemporaryLocation

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) GetTemporaryLocation() string`

GetTemporaryLocation returns the TemporaryLocation field if non-nil, zero value otherwise.

### GetTemporaryLocationOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) GetTemporaryLocationOk() (*string, bool)`

GetTemporaryLocationOk returns a tuple with the TemporaryLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryLocation

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) SetTemporaryLocation(v string)`

SetTemporaryLocation sets TemporaryLocation field to given value.

### HasTemporaryLocation

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) HasTemporaryLocation() bool`

HasTemporaryLocation returns a boolean if a field has been set.

### GetPublicNotes

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) GetPublicNotes() []string`

GetPublicNotes returns the PublicNotes field if non-nil, zero value otherwise.

### GetPublicNotesOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) GetPublicNotesOk() (*[]string, bool)`

GetPublicNotesOk returns a tuple with the PublicNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNotes

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) SetPublicNotes(v []string)`

SetPublicNotes sets PublicNotes field to given value.

### HasPublicNotes

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) HasPublicNotes() bool`

HasPublicNotes returns a boolean if a field has been set.

### GetYearRanges

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) GetYearRanges() []FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInnerYearRangesInner`

GetYearRanges returns the YearRanges field if non-nil, zero value otherwise.

### GetYearRangesOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) GetYearRangesOk() (*[]FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInnerYearRangesInner, bool)`

GetYearRangesOk returns a tuple with the YearRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearRanges

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) SetYearRanges(v []FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInnerYearRangesInner)`

SetYearRanges sets YearRanges field to given value.

### HasYearRanges

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) HasYearRanges() bool`

HasYearRanges returns a boolean if a field has been set.

### GetVolumeRanges

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) GetVolumeRanges() []FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInnerYearRangesInner`

GetVolumeRanges returns the VolumeRanges field if non-nil, zero value otherwise.

### GetVolumeRangesOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) GetVolumeRangesOk() (*[]FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInnerYearRangesInner, bool)`

GetVolumeRangesOk returns a tuple with the VolumeRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeRanges

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) SetVolumeRanges(v []FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInnerYearRangesInner)`

SetVolumeRanges sets VolumeRanges field to given value.

### HasVolumeRanges

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner) HasVolumeRanges() bool`

HasVolumeRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


