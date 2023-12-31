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

// checks if the SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem{}

// SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem struct for SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem
type SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem struct {
	TypeOfRelationship *string  `json:"typeOfRelationship,omitempty"`
	RelatedItemTitle   *string  `json:"relatedItemTitle,omitempty"`
	RelationshipInfo   []string `json:"relationshipInfo,omitempty"`
	RelatedParts       []string `json:"relatedParts,omitempty"`
	MainEntryHeading   *string  `json:"mainEntryHeading,omitempty"`
	// Edition Statement [v250sa]
	Edition                  *string  `json:"edition,omitempty"`
	QualifyingInfo           *string  `json:"qualifyingInfo,omitempty"`
	PublicationInfo          *string  `json:"publicationInfo,omitempty"`
	PhysicalDescription      *string  `json:"physicalDescription,omitempty"`
	SeriesData               []string `json:"seriesData,omitempty"`
	MaterialSpecificDetail   *string  `json:"materialSpecificDetail,omitempty"`
	RelatedItemNotes         []string `json:"relatedItemNotes,omitempty"`
	OtherItemIdentifiers     []string `json:"otherItemIdentifiers,omitempty"`
	ReportNumbers            []string `json:"reportNumbers,omitempty"`
	UniformTitle             *string  `json:"uniformTitle,omitempty"`
	StandardTechReportNumber *string  `json:"standardTechReportNumber,omitempty"`
	RecordControlNumbers     []string `json:"recordControlNumbers,omitempty"`
	// International Standard Serial Numbers
	Issns []string `json:"issns,omitempty"`
	Coden *string  `json:"coden,omitempty"`
	// International Standard Book Number [v020sa]
	Isbns []string `json:"isbns,omitempty"`
}

// NewSearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem instantiates a new SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem() *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem {
	this := SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem{}
	return &this
}

// NewSearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItemWithDefaults instantiates a new SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItemWithDefaults() *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem {
	this := SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem{}
	return &this
}

// GetTypeOfRelationship returns the TypeOfRelationship field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetTypeOfRelationship() string {
	if o == nil || IsNil(o.TypeOfRelationship) {
		var ret string
		return ret
	}
	return *o.TypeOfRelationship
}

// GetTypeOfRelationshipOk returns a tuple with the TypeOfRelationship field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetTypeOfRelationshipOk() (*string, bool) {
	if o == nil || IsNil(o.TypeOfRelationship) {
		return nil, false
	}
	return o.TypeOfRelationship, true
}

// HasTypeOfRelationship returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) HasTypeOfRelationship() bool {
	if o != nil && !IsNil(o.TypeOfRelationship) {
		return true
	}

	return false
}

// SetTypeOfRelationship gets a reference to the given string and assigns it to the TypeOfRelationship field.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) SetTypeOfRelationship(v string) {
	o.TypeOfRelationship = &v
}

// GetRelatedItemTitle returns the RelatedItemTitle field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetRelatedItemTitle() string {
	if o == nil || IsNil(o.RelatedItemTitle) {
		var ret string
		return ret
	}
	return *o.RelatedItemTitle
}

// GetRelatedItemTitleOk returns a tuple with the RelatedItemTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetRelatedItemTitleOk() (*string, bool) {
	if o == nil || IsNil(o.RelatedItemTitle) {
		return nil, false
	}
	return o.RelatedItemTitle, true
}

// HasRelatedItemTitle returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) HasRelatedItemTitle() bool {
	if o != nil && !IsNil(o.RelatedItemTitle) {
		return true
	}

	return false
}

// SetRelatedItemTitle gets a reference to the given string and assigns it to the RelatedItemTitle field.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) SetRelatedItemTitle(v string) {
	o.RelatedItemTitle = &v
}

// GetRelationshipInfo returns the RelationshipInfo field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetRelationshipInfo() []string {
	if o == nil || IsNil(o.RelationshipInfo) {
		var ret []string
		return ret
	}
	return o.RelationshipInfo
}

// GetRelationshipInfoOk returns a tuple with the RelationshipInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetRelationshipInfoOk() ([]string, bool) {
	if o == nil || IsNil(o.RelationshipInfo) {
		return nil, false
	}
	return o.RelationshipInfo, true
}

// HasRelationshipInfo returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) HasRelationshipInfo() bool {
	if o != nil && !IsNil(o.RelationshipInfo) {
		return true
	}

	return false
}

// SetRelationshipInfo gets a reference to the given []string and assigns it to the RelationshipInfo field.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) SetRelationshipInfo(v []string) {
	o.RelationshipInfo = v
}

// GetRelatedParts returns the RelatedParts field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetRelatedParts() []string {
	if o == nil || IsNil(o.RelatedParts) {
		var ret []string
		return ret
	}
	return o.RelatedParts
}

// GetRelatedPartsOk returns a tuple with the RelatedParts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetRelatedPartsOk() ([]string, bool) {
	if o == nil || IsNil(o.RelatedParts) {
		return nil, false
	}
	return o.RelatedParts, true
}

// HasRelatedParts returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) HasRelatedParts() bool {
	if o != nil && !IsNil(o.RelatedParts) {
		return true
	}

	return false
}

// SetRelatedParts gets a reference to the given []string and assigns it to the RelatedParts field.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) SetRelatedParts(v []string) {
	o.RelatedParts = v
}

// GetMainEntryHeading returns the MainEntryHeading field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetMainEntryHeading() string {
	if o == nil || IsNil(o.MainEntryHeading) {
		var ret string
		return ret
	}
	return *o.MainEntryHeading
}

// GetMainEntryHeadingOk returns a tuple with the MainEntryHeading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetMainEntryHeadingOk() (*string, bool) {
	if o == nil || IsNil(o.MainEntryHeading) {
		return nil, false
	}
	return o.MainEntryHeading, true
}

// HasMainEntryHeading returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) HasMainEntryHeading() bool {
	if o != nil && !IsNil(o.MainEntryHeading) {
		return true
	}

	return false
}

// SetMainEntryHeading gets a reference to the given string and assigns it to the MainEntryHeading field.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) SetMainEntryHeading(v string) {
	o.MainEntryHeading = &v
}

// GetEdition returns the Edition field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetEdition() string {
	if o == nil || IsNil(o.Edition) {
		var ret string
		return ret
	}
	return *o.Edition
}

// GetEditionOk returns a tuple with the Edition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetEditionOk() (*string, bool) {
	if o == nil || IsNil(o.Edition) {
		return nil, false
	}
	return o.Edition, true
}

// HasEdition returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) HasEdition() bool {
	if o != nil && !IsNil(o.Edition) {
		return true
	}

	return false
}

// SetEdition gets a reference to the given string and assigns it to the Edition field.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) SetEdition(v string) {
	o.Edition = &v
}

// GetQualifyingInfo returns the QualifyingInfo field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetQualifyingInfo() string {
	if o == nil || IsNil(o.QualifyingInfo) {
		var ret string
		return ret
	}
	return *o.QualifyingInfo
}

// GetQualifyingInfoOk returns a tuple with the QualifyingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetQualifyingInfoOk() (*string, bool) {
	if o == nil || IsNil(o.QualifyingInfo) {
		return nil, false
	}
	return o.QualifyingInfo, true
}

// HasQualifyingInfo returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) HasQualifyingInfo() bool {
	if o != nil && !IsNil(o.QualifyingInfo) {
		return true
	}

	return false
}

// SetQualifyingInfo gets a reference to the given string and assigns it to the QualifyingInfo field.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) SetQualifyingInfo(v string) {
	o.QualifyingInfo = &v
}

// GetPublicationInfo returns the PublicationInfo field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetPublicationInfo() string {
	if o == nil || IsNil(o.PublicationInfo) {
		var ret string
		return ret
	}
	return *o.PublicationInfo
}

// GetPublicationInfoOk returns a tuple with the PublicationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetPublicationInfoOk() (*string, bool) {
	if o == nil || IsNil(o.PublicationInfo) {
		return nil, false
	}
	return o.PublicationInfo, true
}

// HasPublicationInfo returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) HasPublicationInfo() bool {
	if o != nil && !IsNil(o.PublicationInfo) {
		return true
	}

	return false
}

// SetPublicationInfo gets a reference to the given string and assigns it to the PublicationInfo field.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) SetPublicationInfo(v string) {
	o.PublicationInfo = &v
}

// GetPhysicalDescription returns the PhysicalDescription field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetPhysicalDescription() string {
	if o == nil || IsNil(o.PhysicalDescription) {
		var ret string
		return ret
	}
	return *o.PhysicalDescription
}

// GetPhysicalDescriptionOk returns a tuple with the PhysicalDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetPhysicalDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.PhysicalDescription) {
		return nil, false
	}
	return o.PhysicalDescription, true
}

// HasPhysicalDescription returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) HasPhysicalDescription() bool {
	if o != nil && !IsNil(o.PhysicalDescription) {
		return true
	}

	return false
}

// SetPhysicalDescription gets a reference to the given string and assigns it to the PhysicalDescription field.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) SetPhysicalDescription(v string) {
	o.PhysicalDescription = &v
}

// GetSeriesData returns the SeriesData field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetSeriesData() []string {
	if o == nil || IsNil(o.SeriesData) {
		var ret []string
		return ret
	}
	return o.SeriesData
}

// GetSeriesDataOk returns a tuple with the SeriesData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetSeriesDataOk() ([]string, bool) {
	if o == nil || IsNil(o.SeriesData) {
		return nil, false
	}
	return o.SeriesData, true
}

// HasSeriesData returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) HasSeriesData() bool {
	if o != nil && !IsNil(o.SeriesData) {
		return true
	}

	return false
}

// SetSeriesData gets a reference to the given []string and assigns it to the SeriesData field.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) SetSeriesData(v []string) {
	o.SeriesData = v
}

// GetMaterialSpecificDetail returns the MaterialSpecificDetail field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetMaterialSpecificDetail() string {
	if o == nil || IsNil(o.MaterialSpecificDetail) {
		var ret string
		return ret
	}
	return *o.MaterialSpecificDetail
}

// GetMaterialSpecificDetailOk returns a tuple with the MaterialSpecificDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetMaterialSpecificDetailOk() (*string, bool) {
	if o == nil || IsNil(o.MaterialSpecificDetail) {
		return nil, false
	}
	return o.MaterialSpecificDetail, true
}

// HasMaterialSpecificDetail returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) HasMaterialSpecificDetail() bool {
	if o != nil && !IsNil(o.MaterialSpecificDetail) {
		return true
	}

	return false
}

// SetMaterialSpecificDetail gets a reference to the given string and assigns it to the MaterialSpecificDetail field.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) SetMaterialSpecificDetail(v string) {
	o.MaterialSpecificDetail = &v
}

// GetRelatedItemNotes returns the RelatedItemNotes field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetRelatedItemNotes() []string {
	if o == nil || IsNil(o.RelatedItemNotes) {
		var ret []string
		return ret
	}
	return o.RelatedItemNotes
}

// GetRelatedItemNotesOk returns a tuple with the RelatedItemNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetRelatedItemNotesOk() ([]string, bool) {
	if o == nil || IsNil(o.RelatedItemNotes) {
		return nil, false
	}
	return o.RelatedItemNotes, true
}

// HasRelatedItemNotes returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) HasRelatedItemNotes() bool {
	if o != nil && !IsNil(o.RelatedItemNotes) {
		return true
	}

	return false
}

// SetRelatedItemNotes gets a reference to the given []string and assigns it to the RelatedItemNotes field.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) SetRelatedItemNotes(v []string) {
	o.RelatedItemNotes = v
}

// GetOtherItemIdentifiers returns the OtherItemIdentifiers field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetOtherItemIdentifiers() []string {
	if o == nil || IsNil(o.OtherItemIdentifiers) {
		var ret []string
		return ret
	}
	return o.OtherItemIdentifiers
}

// GetOtherItemIdentifiersOk returns a tuple with the OtherItemIdentifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetOtherItemIdentifiersOk() ([]string, bool) {
	if o == nil || IsNil(o.OtherItemIdentifiers) {
		return nil, false
	}
	return o.OtherItemIdentifiers, true
}

// HasOtherItemIdentifiers returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) HasOtherItemIdentifiers() bool {
	if o != nil && !IsNil(o.OtherItemIdentifiers) {
		return true
	}

	return false
}

// SetOtherItemIdentifiers gets a reference to the given []string and assigns it to the OtherItemIdentifiers field.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) SetOtherItemIdentifiers(v []string) {
	o.OtherItemIdentifiers = v
}

// GetReportNumbers returns the ReportNumbers field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetReportNumbers() []string {
	if o == nil || IsNil(o.ReportNumbers) {
		var ret []string
		return ret
	}
	return o.ReportNumbers
}

// GetReportNumbersOk returns a tuple with the ReportNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetReportNumbersOk() ([]string, bool) {
	if o == nil || IsNil(o.ReportNumbers) {
		return nil, false
	}
	return o.ReportNumbers, true
}

// HasReportNumbers returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) HasReportNumbers() bool {
	if o != nil && !IsNil(o.ReportNumbers) {
		return true
	}

	return false
}

// SetReportNumbers gets a reference to the given []string and assigns it to the ReportNumbers field.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) SetReportNumbers(v []string) {
	o.ReportNumbers = v
}

// GetUniformTitle returns the UniformTitle field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetUniformTitle() string {
	if o == nil || IsNil(o.UniformTitle) {
		var ret string
		return ret
	}
	return *o.UniformTitle
}

// GetUniformTitleOk returns a tuple with the UniformTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetUniformTitleOk() (*string, bool) {
	if o == nil || IsNil(o.UniformTitle) {
		return nil, false
	}
	return o.UniformTitle, true
}

// HasUniformTitle returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) HasUniformTitle() bool {
	if o != nil && !IsNil(o.UniformTitle) {
		return true
	}

	return false
}

// SetUniformTitle gets a reference to the given string and assigns it to the UniformTitle field.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) SetUniformTitle(v string) {
	o.UniformTitle = &v
}

// GetStandardTechReportNumber returns the StandardTechReportNumber field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetStandardTechReportNumber() string {
	if o == nil || IsNil(o.StandardTechReportNumber) {
		var ret string
		return ret
	}
	return *o.StandardTechReportNumber
}

// GetStandardTechReportNumberOk returns a tuple with the StandardTechReportNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetStandardTechReportNumberOk() (*string, bool) {
	if o == nil || IsNil(o.StandardTechReportNumber) {
		return nil, false
	}
	return o.StandardTechReportNumber, true
}

// HasStandardTechReportNumber returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) HasStandardTechReportNumber() bool {
	if o != nil && !IsNil(o.StandardTechReportNumber) {
		return true
	}

	return false
}

// SetStandardTechReportNumber gets a reference to the given string and assigns it to the StandardTechReportNumber field.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) SetStandardTechReportNumber(v string) {
	o.StandardTechReportNumber = &v
}

// GetRecordControlNumbers returns the RecordControlNumbers field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetRecordControlNumbers() []string {
	if o == nil || IsNil(o.RecordControlNumbers) {
		var ret []string
		return ret
	}
	return o.RecordControlNumbers
}

// GetRecordControlNumbersOk returns a tuple with the RecordControlNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetRecordControlNumbersOk() ([]string, bool) {
	if o == nil || IsNil(o.RecordControlNumbers) {
		return nil, false
	}
	return o.RecordControlNumbers, true
}

// HasRecordControlNumbers returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) HasRecordControlNumbers() bool {
	if o != nil && !IsNil(o.RecordControlNumbers) {
		return true
	}

	return false
}

// SetRecordControlNumbers gets a reference to the given []string and assigns it to the RecordControlNumbers field.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) SetRecordControlNumbers(v []string) {
	o.RecordControlNumbers = v
}

// GetIssns returns the Issns field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetIssns() []string {
	if o == nil || IsNil(o.Issns) {
		var ret []string
		return ret
	}
	return o.Issns
}

// GetIssnsOk returns a tuple with the Issns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetIssnsOk() ([]string, bool) {
	if o == nil || IsNil(o.Issns) {
		return nil, false
	}
	return o.Issns, true
}

// HasIssns returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) HasIssns() bool {
	if o != nil && !IsNil(o.Issns) {
		return true
	}

	return false
}

// SetIssns gets a reference to the given []string and assigns it to the Issns field.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) SetIssns(v []string) {
	o.Issns = v
}

// GetCoden returns the Coden field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetCoden() string {
	if o == nil || IsNil(o.Coden) {
		var ret string
		return ret
	}
	return *o.Coden
}

// GetCodenOk returns a tuple with the Coden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetCodenOk() (*string, bool) {
	if o == nil || IsNil(o.Coden) {
		return nil, false
	}
	return o.Coden, true
}

// HasCoden returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) HasCoden() bool {
	if o != nil && !IsNil(o.Coden) {
		return true
	}

	return false
}

// SetCoden gets a reference to the given string and assigns it to the Coden field.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) SetCoden(v string) {
	o.Coden = &v
}

// GetIsbns returns the Isbns field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetIsbns() []string {
	if o == nil || IsNil(o.Isbns) {
		var ret []string
		return ret
	}
	return o.Isbns
}

// GetIsbnsOk returns a tuple with the Isbns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) GetIsbnsOk() ([]string, bool) {
	if o == nil || IsNil(o.Isbns) {
		return nil, false
	}
	return o.Isbns, true
}

// HasIsbns returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) HasIsbns() bool {
	if o != nil && !IsNil(o.Isbns) {
		return true
	}

	return false
}

// SetIsbns gets a reference to the given []string and assigns it to the Isbns field.
func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) SetIsbns(v []string) {
	o.Isbns = v
}

func (o SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TypeOfRelationship) {
		toSerialize["typeOfRelationship"] = o.TypeOfRelationship
	}
	if !IsNil(o.RelatedItemTitle) {
		toSerialize["relatedItemTitle"] = o.RelatedItemTitle
	}
	if !IsNil(o.RelationshipInfo) {
		toSerialize["relationshipInfo"] = o.RelationshipInfo
	}
	if !IsNil(o.RelatedParts) {
		toSerialize["relatedParts"] = o.RelatedParts
	}
	if !IsNil(o.MainEntryHeading) {
		toSerialize["mainEntryHeading"] = o.MainEntryHeading
	}
	if !IsNil(o.Edition) {
		toSerialize["edition"] = o.Edition
	}
	if !IsNil(o.QualifyingInfo) {
		toSerialize["qualifyingInfo"] = o.QualifyingInfo
	}
	if !IsNil(o.PublicationInfo) {
		toSerialize["publicationInfo"] = o.PublicationInfo
	}
	if !IsNil(o.PhysicalDescription) {
		toSerialize["physicalDescription"] = o.PhysicalDescription
	}
	if !IsNil(o.SeriesData) {
		toSerialize["seriesData"] = o.SeriesData
	}
	if !IsNil(o.MaterialSpecificDetail) {
		toSerialize["materialSpecificDetail"] = o.MaterialSpecificDetail
	}
	if !IsNil(o.RelatedItemNotes) {
		toSerialize["relatedItemNotes"] = o.RelatedItemNotes
	}
	if !IsNil(o.OtherItemIdentifiers) {
		toSerialize["otherItemIdentifiers"] = o.OtherItemIdentifiers
	}
	if !IsNil(o.ReportNumbers) {
		toSerialize["reportNumbers"] = o.ReportNumbers
	}
	if !IsNil(o.UniformTitle) {
		toSerialize["uniformTitle"] = o.UniformTitle
	}
	if !IsNil(o.StandardTechReportNumber) {
		toSerialize["standardTechReportNumber"] = o.StandardTechReportNumber
	}
	if !IsNil(o.RecordControlNumbers) {
		toSerialize["recordControlNumbers"] = o.RecordControlNumbers
	}
	if !IsNil(o.Issns) {
		toSerialize["issns"] = o.Issns
	}
	if !IsNil(o.Coden) {
		toSerialize["coden"] = o.Coden
	}
	if !IsNil(o.Isbns) {
		toSerialize["isbns"] = o.Isbns
	}
	return toSerialize, nil
}

type NullableSearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem struct {
	value *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem
	isSet bool
}

func (v NullableSearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) Get() *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem {
	return v.value
}

func (v *NullableSearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) Set(val *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem(val *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) *NullableSearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem {
	return &NullableSearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem{value: val, isSet: true}
}

func (v NullableSearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
