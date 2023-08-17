# SearchBibs200ResponseBibRecordsInnerNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneralNotes** | Pointer to [**[]SearchBibs200ResponseBibRecordsInnerNoteGeneralNotesInner**](SearchBibs200ResponseBibRecordsInnerNoteGeneralNotesInner.md) | General Notes [v500sa, v501sa, v590sa,s3, (v59x sa-z)] | [optional] 
**AudienceNote** | Pointer to **string** | Target Audience Note [v521sa] | [optional] 
**DissertationNotes** | Pointer to [**[]SearchBibs200ResponseBibRecordsInnerNoteDissertationNotesInner**](SearchBibs200ResponseBibRecordsInnerNoteDissertationNotesInner.md) |  | [optional] 
**CastNotes** | Pointer to **[]string** | Cast Notes [v511sa] | [optional] 
**PerformerNotes** | Pointer to **[]string** | Performers Notes [v511sa] | [optional] 
**ParticipantNote** | Pointer to **string** | Participants Notes [v511sa] | [optional] 
**EventNotes** | Pointer to [**[]SearchBibs200ResponseBibRecordsInnerNoteEventNotesInner**](SearchBibs200ResponseBibRecordsInnerNoteEventNotesInner.md) |  | [optional] 
**CreditNotes** | Pointer to **string** | Production Credits Notes [v508sa] | [optional] 
**ScaleNotes** | Pointer to **[]string** | Scale Note for Graphic Material [v507sa,b] | [optional] 
**ReproductionNotes** | Pointer to [**[]SearchBibs200ResponseBibRecordsInnerNoteReproductionNotesInner**](SearchBibs200ResponseBibRecordsInnerNoteReproductionNotesInner.md) | Reproduction Notes [v533] | [optional] 
**UseAndReproductionNotes** | Pointer to **[]string** | Terms Governing Use and Reproduction [v540sa] | [optional] 
**UseownershipAndCustodialHistories** | Pointer to **[]string** | Ownership and Custodial History [v561sa,3,5,u] | [optional] 
**SystemDetailNote** | Pointer to **string** | System Details Note [v538sa,u,i] | [optional] 
**AwardNote** | Pointer to **string** | Awards Note [v586sa] | [optional] 
**LanguageNotes** | Pointer to **[]string** | language note [v546sa] | [optional] 

## Methods

### NewSearchBibs200ResponseBibRecordsInnerNote

`func NewSearchBibs200ResponseBibRecordsInnerNote() *SearchBibs200ResponseBibRecordsInnerNote`

NewSearchBibs200ResponseBibRecordsInnerNote instantiates a new SearchBibs200ResponseBibRecordsInnerNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchBibs200ResponseBibRecordsInnerNoteWithDefaults

`func NewSearchBibs200ResponseBibRecordsInnerNoteWithDefaults() *SearchBibs200ResponseBibRecordsInnerNote`

NewSearchBibs200ResponseBibRecordsInnerNoteWithDefaults instantiates a new SearchBibs200ResponseBibRecordsInnerNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneralNotes

`func (o *SearchBibs200ResponseBibRecordsInnerNote) GetGeneralNotes() []SearchBibs200ResponseBibRecordsInnerNoteGeneralNotesInner`

GetGeneralNotes returns the GeneralNotes field if non-nil, zero value otherwise.

### GetGeneralNotesOk

`func (o *SearchBibs200ResponseBibRecordsInnerNote) GetGeneralNotesOk() (*[]SearchBibs200ResponseBibRecordsInnerNoteGeneralNotesInner, bool)`

GetGeneralNotesOk returns a tuple with the GeneralNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralNotes

`func (o *SearchBibs200ResponseBibRecordsInnerNote) SetGeneralNotes(v []SearchBibs200ResponseBibRecordsInnerNoteGeneralNotesInner)`

SetGeneralNotes sets GeneralNotes field to given value.

### HasGeneralNotes

`func (o *SearchBibs200ResponseBibRecordsInnerNote) HasGeneralNotes() bool`

HasGeneralNotes returns a boolean if a field has been set.

### GetAudienceNote

`func (o *SearchBibs200ResponseBibRecordsInnerNote) GetAudienceNote() string`

GetAudienceNote returns the AudienceNote field if non-nil, zero value otherwise.

### GetAudienceNoteOk

`func (o *SearchBibs200ResponseBibRecordsInnerNote) GetAudienceNoteOk() (*string, bool)`

GetAudienceNoteOk returns a tuple with the AudienceNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceNote

`func (o *SearchBibs200ResponseBibRecordsInnerNote) SetAudienceNote(v string)`

SetAudienceNote sets AudienceNote field to given value.

### HasAudienceNote

`func (o *SearchBibs200ResponseBibRecordsInnerNote) HasAudienceNote() bool`

HasAudienceNote returns a boolean if a field has been set.

### GetDissertationNotes

`func (o *SearchBibs200ResponseBibRecordsInnerNote) GetDissertationNotes() []SearchBibs200ResponseBibRecordsInnerNoteDissertationNotesInner`

GetDissertationNotes returns the DissertationNotes field if non-nil, zero value otherwise.

### GetDissertationNotesOk

`func (o *SearchBibs200ResponseBibRecordsInnerNote) GetDissertationNotesOk() (*[]SearchBibs200ResponseBibRecordsInnerNoteDissertationNotesInner, bool)`

GetDissertationNotesOk returns a tuple with the DissertationNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDissertationNotes

`func (o *SearchBibs200ResponseBibRecordsInnerNote) SetDissertationNotes(v []SearchBibs200ResponseBibRecordsInnerNoteDissertationNotesInner)`

SetDissertationNotes sets DissertationNotes field to given value.

### HasDissertationNotes

`func (o *SearchBibs200ResponseBibRecordsInnerNote) HasDissertationNotes() bool`

HasDissertationNotes returns a boolean if a field has been set.

### GetCastNotes

`func (o *SearchBibs200ResponseBibRecordsInnerNote) GetCastNotes() []string`

GetCastNotes returns the CastNotes field if non-nil, zero value otherwise.

### GetCastNotesOk

`func (o *SearchBibs200ResponseBibRecordsInnerNote) GetCastNotesOk() (*[]string, bool)`

GetCastNotesOk returns a tuple with the CastNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastNotes

`func (o *SearchBibs200ResponseBibRecordsInnerNote) SetCastNotes(v []string)`

SetCastNotes sets CastNotes field to given value.

### HasCastNotes

`func (o *SearchBibs200ResponseBibRecordsInnerNote) HasCastNotes() bool`

HasCastNotes returns a boolean if a field has been set.

### GetPerformerNotes

`func (o *SearchBibs200ResponseBibRecordsInnerNote) GetPerformerNotes() []string`

GetPerformerNotes returns the PerformerNotes field if non-nil, zero value otherwise.

### GetPerformerNotesOk

`func (o *SearchBibs200ResponseBibRecordsInnerNote) GetPerformerNotesOk() (*[]string, bool)`

GetPerformerNotesOk returns a tuple with the PerformerNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformerNotes

`func (o *SearchBibs200ResponseBibRecordsInnerNote) SetPerformerNotes(v []string)`

SetPerformerNotes sets PerformerNotes field to given value.

### HasPerformerNotes

`func (o *SearchBibs200ResponseBibRecordsInnerNote) HasPerformerNotes() bool`

HasPerformerNotes returns a boolean if a field has been set.

### GetParticipantNote

`func (o *SearchBibs200ResponseBibRecordsInnerNote) GetParticipantNote() string`

GetParticipantNote returns the ParticipantNote field if non-nil, zero value otherwise.

### GetParticipantNoteOk

`func (o *SearchBibs200ResponseBibRecordsInnerNote) GetParticipantNoteOk() (*string, bool)`

GetParticipantNoteOk returns a tuple with the ParticipantNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantNote

`func (o *SearchBibs200ResponseBibRecordsInnerNote) SetParticipantNote(v string)`

SetParticipantNote sets ParticipantNote field to given value.

### HasParticipantNote

`func (o *SearchBibs200ResponseBibRecordsInnerNote) HasParticipantNote() bool`

HasParticipantNote returns a boolean if a field has been set.

### GetEventNotes

`func (o *SearchBibs200ResponseBibRecordsInnerNote) GetEventNotes() []SearchBibs200ResponseBibRecordsInnerNoteEventNotesInner`

GetEventNotes returns the EventNotes field if non-nil, zero value otherwise.

### GetEventNotesOk

`func (o *SearchBibs200ResponseBibRecordsInnerNote) GetEventNotesOk() (*[]SearchBibs200ResponseBibRecordsInnerNoteEventNotesInner, bool)`

GetEventNotesOk returns a tuple with the EventNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotes

`func (o *SearchBibs200ResponseBibRecordsInnerNote) SetEventNotes(v []SearchBibs200ResponseBibRecordsInnerNoteEventNotesInner)`

SetEventNotes sets EventNotes field to given value.

### HasEventNotes

`func (o *SearchBibs200ResponseBibRecordsInnerNote) HasEventNotes() bool`

HasEventNotes returns a boolean if a field has been set.

### GetCreditNotes

`func (o *SearchBibs200ResponseBibRecordsInnerNote) GetCreditNotes() string`

GetCreditNotes returns the CreditNotes field if non-nil, zero value otherwise.

### GetCreditNotesOk

`func (o *SearchBibs200ResponseBibRecordsInnerNote) GetCreditNotesOk() (*string, bool)`

GetCreditNotesOk returns a tuple with the CreditNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNotes

`func (o *SearchBibs200ResponseBibRecordsInnerNote) SetCreditNotes(v string)`

SetCreditNotes sets CreditNotes field to given value.

### HasCreditNotes

`func (o *SearchBibs200ResponseBibRecordsInnerNote) HasCreditNotes() bool`

HasCreditNotes returns a boolean if a field has been set.

### GetScaleNotes

`func (o *SearchBibs200ResponseBibRecordsInnerNote) GetScaleNotes() []string`

GetScaleNotes returns the ScaleNotes field if non-nil, zero value otherwise.

### GetScaleNotesOk

`func (o *SearchBibs200ResponseBibRecordsInnerNote) GetScaleNotesOk() (*[]string, bool)`

GetScaleNotesOk returns a tuple with the ScaleNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleNotes

`func (o *SearchBibs200ResponseBibRecordsInnerNote) SetScaleNotes(v []string)`

SetScaleNotes sets ScaleNotes field to given value.

### HasScaleNotes

`func (o *SearchBibs200ResponseBibRecordsInnerNote) HasScaleNotes() bool`

HasScaleNotes returns a boolean if a field has been set.

### GetReproductionNotes

`func (o *SearchBibs200ResponseBibRecordsInnerNote) GetReproductionNotes() []SearchBibs200ResponseBibRecordsInnerNoteReproductionNotesInner`

GetReproductionNotes returns the ReproductionNotes field if non-nil, zero value otherwise.

### GetReproductionNotesOk

`func (o *SearchBibs200ResponseBibRecordsInnerNote) GetReproductionNotesOk() (*[]SearchBibs200ResponseBibRecordsInnerNoteReproductionNotesInner, bool)`

GetReproductionNotesOk returns a tuple with the ReproductionNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReproductionNotes

`func (o *SearchBibs200ResponseBibRecordsInnerNote) SetReproductionNotes(v []SearchBibs200ResponseBibRecordsInnerNoteReproductionNotesInner)`

SetReproductionNotes sets ReproductionNotes field to given value.

### HasReproductionNotes

`func (o *SearchBibs200ResponseBibRecordsInnerNote) HasReproductionNotes() bool`

HasReproductionNotes returns a boolean if a field has been set.

### GetUseAndReproductionNotes

`func (o *SearchBibs200ResponseBibRecordsInnerNote) GetUseAndReproductionNotes() []string`

GetUseAndReproductionNotes returns the UseAndReproductionNotes field if non-nil, zero value otherwise.

### GetUseAndReproductionNotesOk

`func (o *SearchBibs200ResponseBibRecordsInnerNote) GetUseAndReproductionNotesOk() (*[]string, bool)`

GetUseAndReproductionNotesOk returns a tuple with the UseAndReproductionNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAndReproductionNotes

`func (o *SearchBibs200ResponseBibRecordsInnerNote) SetUseAndReproductionNotes(v []string)`

SetUseAndReproductionNotes sets UseAndReproductionNotes field to given value.

### HasUseAndReproductionNotes

`func (o *SearchBibs200ResponseBibRecordsInnerNote) HasUseAndReproductionNotes() bool`

HasUseAndReproductionNotes returns a boolean if a field has been set.

### GetUseownershipAndCustodialHistories

`func (o *SearchBibs200ResponseBibRecordsInnerNote) GetUseownershipAndCustodialHistories() []string`

GetUseownershipAndCustodialHistories returns the UseownershipAndCustodialHistories field if non-nil, zero value otherwise.

### GetUseownershipAndCustodialHistoriesOk

`func (o *SearchBibs200ResponseBibRecordsInnerNote) GetUseownershipAndCustodialHistoriesOk() (*[]string, bool)`

GetUseownershipAndCustodialHistoriesOk returns a tuple with the UseownershipAndCustodialHistories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseownershipAndCustodialHistories

`func (o *SearchBibs200ResponseBibRecordsInnerNote) SetUseownershipAndCustodialHistories(v []string)`

SetUseownershipAndCustodialHistories sets UseownershipAndCustodialHistories field to given value.

### HasUseownershipAndCustodialHistories

`func (o *SearchBibs200ResponseBibRecordsInnerNote) HasUseownershipAndCustodialHistories() bool`

HasUseownershipAndCustodialHistories returns a boolean if a field has been set.

### GetSystemDetailNote

`func (o *SearchBibs200ResponseBibRecordsInnerNote) GetSystemDetailNote() string`

GetSystemDetailNote returns the SystemDetailNote field if non-nil, zero value otherwise.

### GetSystemDetailNoteOk

`func (o *SearchBibs200ResponseBibRecordsInnerNote) GetSystemDetailNoteOk() (*string, bool)`

GetSystemDetailNoteOk returns a tuple with the SystemDetailNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDetailNote

`func (o *SearchBibs200ResponseBibRecordsInnerNote) SetSystemDetailNote(v string)`

SetSystemDetailNote sets SystemDetailNote field to given value.

### HasSystemDetailNote

`func (o *SearchBibs200ResponseBibRecordsInnerNote) HasSystemDetailNote() bool`

HasSystemDetailNote returns a boolean if a field has been set.

### GetAwardNote

`func (o *SearchBibs200ResponseBibRecordsInnerNote) GetAwardNote() string`

GetAwardNote returns the AwardNote field if non-nil, zero value otherwise.

### GetAwardNoteOk

`func (o *SearchBibs200ResponseBibRecordsInnerNote) GetAwardNoteOk() (*string, bool)`

GetAwardNoteOk returns a tuple with the AwardNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwardNote

`func (o *SearchBibs200ResponseBibRecordsInnerNote) SetAwardNote(v string)`

SetAwardNote sets AwardNote field to given value.

### HasAwardNote

`func (o *SearchBibs200ResponseBibRecordsInnerNote) HasAwardNote() bool`

HasAwardNote returns a boolean if a field has been set.

### GetLanguageNotes

`func (o *SearchBibs200ResponseBibRecordsInnerNote) GetLanguageNotes() []string`

GetLanguageNotes returns the LanguageNotes field if non-nil, zero value otherwise.

### GetLanguageNotesOk

`func (o *SearchBibs200ResponseBibRecordsInnerNote) GetLanguageNotesOk() (*[]string, bool)`

GetLanguageNotesOk returns a tuple with the LanguageNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageNotes

`func (o *SearchBibs200ResponseBibRecordsInnerNote) SetLanguageNotes(v []string)`

SetLanguageNotes sets LanguageNotes field to given value.

### HasLanguageNotes

`func (o *SearchBibs200ResponseBibRecordsInnerNote) HasLanguageNotes() bool`

HasLanguageNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


