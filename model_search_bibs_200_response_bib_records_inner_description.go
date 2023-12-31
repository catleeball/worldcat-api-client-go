/*
WorldCat Search API v. 2

Searching of WorldCat

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SearchBibs200ResponseBibRecordsInnerDescription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchBibs200ResponseBibRecordsInnerDescription{}

// SearchBibs200ResponseBibRecordsInnerDescription struct for SearchBibs200ResponseBibRecordsInnerDescription
type SearchBibs200ResponseBibRecordsInnerDescription struct {
	// Physical Description [v300sa,b,c,d,e,f,g,3]
	PhysicalDescription *string `json:"physicalDescription,omitempty"`
	// Digital Graphic Representation [v352sa,b,c,i,q]
	DigitalGraphicRepresentation *string `json:"digitalGraphicRepresentation,omitempty"`
	// Genre of work [v655sa,b,v,x,y,z v380sa, (v600,v610,v611,v630,v648, v650,v651,v654 sv) v653sa]
	Genres []string `json:"genres,omitempty"`
	// Cartographic Mathematical Data [v255sa,b,c]
	CartographicData []string `json:"cartographicData,omitempty"`
	// Computer File Characteristics [v256sa]
	ComputerFilesCharacteristics []string `json:"computerFilesCharacteristics,omitempty"`
	// Form of Work [v380sa]
	FormOfWorks []string `json:"formOfWorks,omitempty"`
	// Abstract [v520sa,b,c]
	Abstracts []SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner `json:"abstracts,omitempty"`
	// Summary [v520sa,b,c]
	Summaries []SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner     `json:"summaries,omitempty"`
	Contents  []SearchBibs200ResponseBibRecordsInnerDescriptionContentsInner `json:"contents,omitempty"`
	// Bibliography [v504sa]
	Bibliographies []SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner `json:"bibliographies,omitempty"`
	// Is this a peer reviewed work [Admin/OCLCDef/PR]
	PeerReviewed *string `json:"peerReviewed,omitempty"`
}

// NewSearchBibs200ResponseBibRecordsInnerDescription instantiates a new SearchBibs200ResponseBibRecordsInnerDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchBibs200ResponseBibRecordsInnerDescription() *SearchBibs200ResponseBibRecordsInnerDescription {
	this := SearchBibs200ResponseBibRecordsInnerDescription{}
	return &this
}

// NewSearchBibs200ResponseBibRecordsInnerDescriptionWithDefaults instantiates a new SearchBibs200ResponseBibRecordsInnerDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchBibs200ResponseBibRecordsInnerDescriptionWithDefaults() *SearchBibs200ResponseBibRecordsInnerDescription {
	this := SearchBibs200ResponseBibRecordsInnerDescription{}
	return &this
}

// GetPhysicalDescription returns the PhysicalDescription field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetPhysicalDescription() string {
	if o == nil || IsNil(o.PhysicalDescription) {
		var ret string
		return ret
	}
	return *o.PhysicalDescription
}

// GetPhysicalDescriptionOk returns a tuple with the PhysicalDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetPhysicalDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.PhysicalDescription) {
		return nil, false
	}
	return o.PhysicalDescription, true
}

// HasPhysicalDescription returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) HasPhysicalDescription() bool {
	if o != nil && !IsNil(o.PhysicalDescription) {
		return true
	}

	return false
}

// SetPhysicalDescription gets a reference to the given string and assigns it to the PhysicalDescription field.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) SetPhysicalDescription(v string) {
	o.PhysicalDescription = &v
}

// GetDigitalGraphicRepresentation returns the DigitalGraphicRepresentation field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetDigitalGraphicRepresentation() string {
	if o == nil || IsNil(o.DigitalGraphicRepresentation) {
		var ret string
		return ret
	}
	return *o.DigitalGraphicRepresentation
}

// GetDigitalGraphicRepresentationOk returns a tuple with the DigitalGraphicRepresentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetDigitalGraphicRepresentationOk() (*string, bool) {
	if o == nil || IsNil(o.DigitalGraphicRepresentation) {
		return nil, false
	}
	return o.DigitalGraphicRepresentation, true
}

// HasDigitalGraphicRepresentation returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) HasDigitalGraphicRepresentation() bool {
	if o != nil && !IsNil(o.DigitalGraphicRepresentation) {
		return true
	}

	return false
}

// SetDigitalGraphicRepresentation gets a reference to the given string and assigns it to the DigitalGraphicRepresentation field.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) SetDigitalGraphicRepresentation(v string) {
	o.DigitalGraphicRepresentation = &v
}

// GetGenres returns the Genres field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetGenres() []string {
	if o == nil || IsNil(o.Genres) {
		var ret []string
		return ret
	}
	return o.Genres
}

// GetGenresOk returns a tuple with the Genres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetGenresOk() ([]string, bool) {
	if o == nil || IsNil(o.Genres) {
		return nil, false
	}
	return o.Genres, true
}

// HasGenres returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) HasGenres() bool {
	if o != nil && !IsNil(o.Genres) {
		return true
	}

	return false
}

// SetGenres gets a reference to the given []string and assigns it to the Genres field.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) SetGenres(v []string) {
	o.Genres = v
}

// GetCartographicData returns the CartographicData field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetCartographicData() []string {
	if o == nil || IsNil(o.CartographicData) {
		var ret []string
		return ret
	}
	return o.CartographicData
}

// GetCartographicDataOk returns a tuple with the CartographicData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetCartographicDataOk() ([]string, bool) {
	if o == nil || IsNil(o.CartographicData) {
		return nil, false
	}
	return o.CartographicData, true
}

// HasCartographicData returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) HasCartographicData() bool {
	if o != nil && !IsNil(o.CartographicData) {
		return true
	}

	return false
}

// SetCartographicData gets a reference to the given []string and assigns it to the CartographicData field.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) SetCartographicData(v []string) {
	o.CartographicData = v
}

// GetComputerFilesCharacteristics returns the ComputerFilesCharacteristics field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetComputerFilesCharacteristics() []string {
	if o == nil || IsNil(o.ComputerFilesCharacteristics) {
		var ret []string
		return ret
	}
	return o.ComputerFilesCharacteristics
}

// GetComputerFilesCharacteristicsOk returns a tuple with the ComputerFilesCharacteristics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetComputerFilesCharacteristicsOk() ([]string, bool) {
	if o == nil || IsNil(o.ComputerFilesCharacteristics) {
		return nil, false
	}
	return o.ComputerFilesCharacteristics, true
}

// HasComputerFilesCharacteristics returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) HasComputerFilesCharacteristics() bool {
	if o != nil && !IsNil(o.ComputerFilesCharacteristics) {
		return true
	}

	return false
}

// SetComputerFilesCharacteristics gets a reference to the given []string and assigns it to the ComputerFilesCharacteristics field.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) SetComputerFilesCharacteristics(v []string) {
	o.ComputerFilesCharacteristics = v
}

// GetFormOfWorks returns the FormOfWorks field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetFormOfWorks() []string {
	if o == nil || IsNil(o.FormOfWorks) {
		var ret []string
		return ret
	}
	return o.FormOfWorks
}

// GetFormOfWorksOk returns a tuple with the FormOfWorks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetFormOfWorksOk() ([]string, bool) {
	if o == nil || IsNil(o.FormOfWorks) {
		return nil, false
	}
	return o.FormOfWorks, true
}

// HasFormOfWorks returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) HasFormOfWorks() bool {
	if o != nil && !IsNil(o.FormOfWorks) {
		return true
	}

	return false
}

// SetFormOfWorks gets a reference to the given []string and assigns it to the FormOfWorks field.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) SetFormOfWorks(v []string) {
	o.FormOfWorks = v
}

// GetAbstracts returns the Abstracts field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetAbstracts() []SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner {
	if o == nil || IsNil(o.Abstracts) {
		var ret []SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner
		return ret
	}
	return o.Abstracts
}

// GetAbstractsOk returns a tuple with the Abstracts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetAbstractsOk() ([]SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner, bool) {
	if o == nil || IsNil(o.Abstracts) {
		return nil, false
	}
	return o.Abstracts, true
}

// HasAbstracts returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) HasAbstracts() bool {
	if o != nil && !IsNil(o.Abstracts) {
		return true
	}

	return false
}

// SetAbstracts gets a reference to the given []SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner and assigns it to the Abstracts field.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) SetAbstracts(v []SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner) {
	o.Abstracts = v
}

// GetSummaries returns the Summaries field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetSummaries() []SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner {
	if o == nil || IsNil(o.Summaries) {
		var ret []SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner
		return ret
	}
	return o.Summaries
}

// GetSummariesOk returns a tuple with the Summaries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetSummariesOk() ([]SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner, bool) {
	if o == nil || IsNil(o.Summaries) {
		return nil, false
	}
	return o.Summaries, true
}

// HasSummaries returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) HasSummaries() bool {
	if o != nil && !IsNil(o.Summaries) {
		return true
	}

	return false
}

// SetSummaries gets a reference to the given []SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner and assigns it to the Summaries field.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) SetSummaries(v []SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner) {
	o.Summaries = v
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetContents() []SearchBibs200ResponseBibRecordsInnerDescriptionContentsInner {
	if o == nil || IsNil(o.Contents) {
		var ret []SearchBibs200ResponseBibRecordsInnerDescriptionContentsInner
		return ret
	}
	return o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetContentsOk() ([]SearchBibs200ResponseBibRecordsInnerDescriptionContentsInner, bool) {
	if o == nil || IsNil(o.Contents) {
		return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) HasContents() bool {
	if o != nil && !IsNil(o.Contents) {
		return true
	}

	return false
}

// SetContents gets a reference to the given []SearchBibs200ResponseBibRecordsInnerDescriptionContentsInner and assigns it to the Contents field.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) SetContents(v []SearchBibs200ResponseBibRecordsInnerDescriptionContentsInner) {
	o.Contents = v
}

// GetBibliographies returns the Bibliographies field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetBibliographies() []SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner {
	if o == nil || IsNil(o.Bibliographies) {
		var ret []SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner
		return ret
	}
	return o.Bibliographies
}

// GetBibliographiesOk returns a tuple with the Bibliographies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetBibliographiesOk() ([]SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner, bool) {
	if o == nil || IsNil(o.Bibliographies) {
		return nil, false
	}
	return o.Bibliographies, true
}

// HasBibliographies returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) HasBibliographies() bool {
	if o != nil && !IsNil(o.Bibliographies) {
		return true
	}

	return false
}

// SetBibliographies gets a reference to the given []SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner and assigns it to the Bibliographies field.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) SetBibliographies(v []SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner) {
	o.Bibliographies = v
}

// GetPeerReviewed returns the PeerReviewed field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetPeerReviewed() string {
	if o == nil || IsNil(o.PeerReviewed) {
		var ret string
		return ret
	}
	return *o.PeerReviewed
}

// GetPeerReviewedOk returns a tuple with the PeerReviewed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) GetPeerReviewedOk() (*string, bool) {
	if o == nil || IsNil(o.PeerReviewed) {
		return nil, false
	}
	return o.PeerReviewed, true
}

// HasPeerReviewed returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) HasPeerReviewed() bool {
	if o != nil && !IsNil(o.PeerReviewed) {
		return true
	}

	return false
}

// SetPeerReviewed gets a reference to the given string and assigns it to the PeerReviewed field.
func (o *SearchBibs200ResponseBibRecordsInnerDescription) SetPeerReviewed(v string) {
	o.PeerReviewed = &v
}

func (o SearchBibs200ResponseBibRecordsInnerDescription) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchBibs200ResponseBibRecordsInnerDescription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PhysicalDescription) {
		toSerialize["physicalDescription"] = o.PhysicalDescription
	}
	if !IsNil(o.DigitalGraphicRepresentation) {
		toSerialize["digitalGraphicRepresentation"] = o.DigitalGraphicRepresentation
	}
	if !IsNil(o.Genres) {
		toSerialize["genres"] = o.Genres
	}
	if !IsNil(o.CartographicData) {
		toSerialize["cartographicData"] = o.CartographicData
	}
	if !IsNil(o.ComputerFilesCharacteristics) {
		toSerialize["computerFilesCharacteristics"] = o.ComputerFilesCharacteristics
	}
	if !IsNil(o.FormOfWorks) {
		toSerialize["formOfWorks"] = o.FormOfWorks
	}
	if !IsNil(o.Abstracts) {
		toSerialize["abstracts"] = o.Abstracts
	}
	if !IsNil(o.Summaries) {
		toSerialize["summaries"] = o.Summaries
	}
	if !IsNil(o.Contents) {
		toSerialize["contents"] = o.Contents
	}
	if !IsNil(o.Bibliographies) {
		toSerialize["bibliographies"] = o.Bibliographies
	}
	if !IsNil(o.PeerReviewed) {
		toSerialize["peerReviewed"] = o.PeerReviewed
	}
	return toSerialize, nil
}

type NullableSearchBibs200ResponseBibRecordsInnerDescription struct {
	value *SearchBibs200ResponseBibRecordsInnerDescription
	isSet bool
}

func (v NullableSearchBibs200ResponseBibRecordsInnerDescription) Get() *SearchBibs200ResponseBibRecordsInnerDescription {
	return v.value
}

func (v *NullableSearchBibs200ResponseBibRecordsInnerDescription) Set(val *SearchBibs200ResponseBibRecordsInnerDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchBibs200ResponseBibRecordsInnerDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchBibs200ResponseBibRecordsInnerDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchBibs200ResponseBibRecordsInnerDescription(val *SearchBibs200ResponseBibRecordsInnerDescription) *NullableSearchBibs200ResponseBibRecordsInnerDescription {
	return &NullableSearchBibs200ResponseBibRecordsInnerDescription{value: val, isSet: true}
}

func (v NullableSearchBibs200ResponseBibRecordsInnerDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchBibs200ResponseBibRecordsInnerDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
