# SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | **string** | URI of resource [v856su] | 
**MaterialSpecified** | Pointer to **string** | Materials specified |3 | [optional] 
**Instructions** | Pointer to **[]string** | Instruction [v856si] | [optional] 
**LinkText** | Pointer to **[]string** | Text associated with a link |y | [optional] 
**NonPublicNote** | Pointer to **string** | Nonpublic note [v856sx] | [optional] 
**PublicNote** | Pointer to **[]string** | A note pertaining to the electronic location of the source identified in the field |z | [optional] 
**AccessMethod** | Pointer to **string** | Access method [v856s2] | [optional] 
**AccessStatus** | Pointer to **string** | Access method [v856s7] | [optional] 
**Relationship** | Pointer to **string** | Relationship to the item described in the record as a whole [v856 second indicator] | [optional] 

## Methods

### NewSearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner

`func NewSearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner(uri string, ) *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner`

NewSearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner instantiates a new SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInnerWithDefaults

`func NewSearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInnerWithDefaults() *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner`

NewSearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInnerWithDefaults instantiates a new SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUri

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) SetUri(v string)`

SetUri sets Uri field to given value.


### GetMaterialSpecified

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) GetMaterialSpecified() string`

GetMaterialSpecified returns the MaterialSpecified field if non-nil, zero value otherwise.

### GetMaterialSpecifiedOk

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) GetMaterialSpecifiedOk() (*string, bool)`

GetMaterialSpecifiedOk returns a tuple with the MaterialSpecified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterialSpecified

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) SetMaterialSpecified(v string)`

SetMaterialSpecified sets MaterialSpecified field to given value.

### HasMaterialSpecified

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) HasMaterialSpecified() bool`

HasMaterialSpecified returns a boolean if a field has been set.

### GetInstructions

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) GetInstructions() []string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) GetInstructionsOk() (*[]string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) SetInstructions(v []string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetLinkText

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) GetLinkText() []string`

GetLinkText returns the LinkText field if non-nil, zero value otherwise.

### GetLinkTextOk

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) GetLinkTextOk() (*[]string, bool)`

GetLinkTextOk returns a tuple with the LinkText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkText

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) SetLinkText(v []string)`

SetLinkText sets LinkText field to given value.

### HasLinkText

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) HasLinkText() bool`

HasLinkText returns a boolean if a field has been set.

### GetNonPublicNote

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) GetNonPublicNote() string`

GetNonPublicNote returns the NonPublicNote field if non-nil, zero value otherwise.

### GetNonPublicNoteOk

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) GetNonPublicNoteOk() (*string, bool)`

GetNonPublicNoteOk returns a tuple with the NonPublicNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonPublicNote

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) SetNonPublicNote(v string)`

SetNonPublicNote sets NonPublicNote field to given value.

### HasNonPublicNote

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) HasNonPublicNote() bool`

HasNonPublicNote returns a boolean if a field has been set.

### GetPublicNote

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) GetPublicNote() []string`

GetPublicNote returns the PublicNote field if non-nil, zero value otherwise.

### GetPublicNoteOk

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) GetPublicNoteOk() (*[]string, bool)`

GetPublicNoteOk returns a tuple with the PublicNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNote

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) SetPublicNote(v []string)`

SetPublicNote sets PublicNote field to given value.

### HasPublicNote

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) HasPublicNote() bool`

HasPublicNote returns a boolean if a field has been set.

### GetAccessMethod

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) GetAccessMethod() string`

GetAccessMethod returns the AccessMethod field if non-nil, zero value otherwise.

### GetAccessMethodOk

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) GetAccessMethodOk() (*string, bool)`

GetAccessMethodOk returns a tuple with the AccessMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMethod

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) SetAccessMethod(v string)`

SetAccessMethod sets AccessMethod field to given value.

### HasAccessMethod

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) HasAccessMethod() bool`

HasAccessMethod returns a boolean if a field has been set.

### GetAccessStatus

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) GetAccessStatus() string`

GetAccessStatus returns the AccessStatus field if non-nil, zero value otherwise.

### GetAccessStatusOk

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) GetAccessStatusOk() (*string, bool)`

GetAccessStatusOk returns a tuple with the AccessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessStatus

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) SetAccessStatus(v string)`

SetAccessStatus sets AccessStatus field to given value.

### HasAccessStatus

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) HasAccessStatus() bool`

HasAccessStatus returns a boolean if a field has been set.

### GetRelationship

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) GetRelationship() string`

GetRelationship returns the Relationship field if non-nil, zero value otherwise.

### GetRelationshipOk

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) GetRelationshipOk() (*string, bool)`

GetRelationshipOk returns a tuple with the Relationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationship

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) SetRelationship(v string)`

SetRelationship sets Relationship field to given value.

### HasRelationship

`func (o *SearchBibs200ResponseBibRecordsInnerDigitalAccessAndLocationsInner) HasRelationship() bool`

HasRelationship returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


