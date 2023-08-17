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

// checks if the SearchBibs200ResponseBibRecordsInnerSourceInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchBibs200ResponseBibRecordsInnerSourceInner{}

// SearchBibs200ResponseBibRecordsInnerSourceInner struct for SearchBibs200ResponseBibRecordsInnerSourceInner
type SearchBibs200ResponseBibRecordsInnerSourceInner struct {
	// Supplements [v949sd]
	Supplements []string `json:"supplements,omitempty"`
	// First Page [v949sf]
	FirstPages []string `json:"firstPages,omitempty"`
	// Issue Dates [v949se]
	IssueDates []string `json:"issueDates,omitempty"`
	// Main Entry [v773sa]
	MainEntry *string `json:"mainEntry,omitempty"`
	// Other Publication Data [v799sa,d,g,o]
	OtherPublicationData []string `json:"otherPublicationData,omitempty"`
	// Related Parts [v773sg]
	RelatedParts []string `json:"relatedParts,omitempty"`
	// Source Title [v799st v773st]
	SourceTitle *string `json:"sourceTitle,omitempty"`
	// Source ISSN [v773sx]
	SourceIssn *string `json:"sourceIssn,omitempty"`
	// Source ISBNs [v799sz]
	SourceIsbns []string `json:"sourceIsbns,omitempty"`
	// Source Author [v799sa]
	SourceAuthor *string `json:"sourceAuthor,omitempty"`
	// Publication Information [v773sd]
	PublicationInformation *string `json:"publicationInformation,omitempty"`
	// Volumes [v799sm v949sa]
	Volumes []string `json:"volumes,omitempty"`
	// Issues [v799sp v949sb]
	Issues []string `json:"issues,omitempty"`
	// Page Range [v949sc]
	PageRanges []string `json:"pageRanges,omitempty"`
}

// NewSearchBibs200ResponseBibRecordsInnerSourceInner instantiates a new SearchBibs200ResponseBibRecordsInnerSourceInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchBibs200ResponseBibRecordsInnerSourceInner() *SearchBibs200ResponseBibRecordsInnerSourceInner {
	this := SearchBibs200ResponseBibRecordsInnerSourceInner{}
	return &this
}

// NewSearchBibs200ResponseBibRecordsInnerSourceInnerWithDefaults instantiates a new SearchBibs200ResponseBibRecordsInnerSourceInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchBibs200ResponseBibRecordsInnerSourceInnerWithDefaults() *SearchBibs200ResponseBibRecordsInnerSourceInner {
	this := SearchBibs200ResponseBibRecordsInnerSourceInner{}
	return &this
}

// GetSupplements returns the Supplements field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) GetSupplements() []string {
	if o == nil || IsNil(o.Supplements) {
		var ret []string
		return ret
	}
	return o.Supplements
}

// GetSupplementsOk returns a tuple with the Supplements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) GetSupplementsOk() ([]string, bool) {
	if o == nil || IsNil(o.Supplements) {
		return nil, false
	}
	return o.Supplements, true
}

// HasSupplements returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) HasSupplements() bool {
	if o != nil && !IsNil(o.Supplements) {
		return true
	}

	return false
}

// SetSupplements gets a reference to the given []string and assigns it to the Supplements field.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) SetSupplements(v []string) {
	o.Supplements = v
}

// GetFirstPages returns the FirstPages field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) GetFirstPages() []string {
	if o == nil || IsNil(o.FirstPages) {
		var ret []string
		return ret
	}
	return o.FirstPages
}

// GetFirstPagesOk returns a tuple with the FirstPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) GetFirstPagesOk() ([]string, bool) {
	if o == nil || IsNil(o.FirstPages) {
		return nil, false
	}
	return o.FirstPages, true
}

// HasFirstPages returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) HasFirstPages() bool {
	if o != nil && !IsNil(o.FirstPages) {
		return true
	}

	return false
}

// SetFirstPages gets a reference to the given []string and assigns it to the FirstPages field.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) SetFirstPages(v []string) {
	o.FirstPages = v
}

// GetIssueDates returns the IssueDates field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) GetIssueDates() []string {
	if o == nil || IsNil(o.IssueDates) {
		var ret []string
		return ret
	}
	return o.IssueDates
}

// GetIssueDatesOk returns a tuple with the IssueDates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) GetIssueDatesOk() ([]string, bool) {
	if o == nil || IsNil(o.IssueDates) {
		return nil, false
	}
	return o.IssueDates, true
}

// HasIssueDates returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) HasIssueDates() bool {
	if o != nil && !IsNil(o.IssueDates) {
		return true
	}

	return false
}

// SetIssueDates gets a reference to the given []string and assigns it to the IssueDates field.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) SetIssueDates(v []string) {
	o.IssueDates = v
}

// GetMainEntry returns the MainEntry field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) GetMainEntry() string {
	if o == nil || IsNil(o.MainEntry) {
		var ret string
		return ret
	}
	return *o.MainEntry
}

// GetMainEntryOk returns a tuple with the MainEntry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) GetMainEntryOk() (*string, bool) {
	if o == nil || IsNil(o.MainEntry) {
		return nil, false
	}
	return o.MainEntry, true
}

// HasMainEntry returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) HasMainEntry() bool {
	if o != nil && !IsNil(o.MainEntry) {
		return true
	}

	return false
}

// SetMainEntry gets a reference to the given string and assigns it to the MainEntry field.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) SetMainEntry(v string) {
	o.MainEntry = &v
}

// GetOtherPublicationData returns the OtherPublicationData field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) GetOtherPublicationData() []string {
	if o == nil || IsNil(o.OtherPublicationData) {
		var ret []string
		return ret
	}
	return o.OtherPublicationData
}

// GetOtherPublicationDataOk returns a tuple with the OtherPublicationData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) GetOtherPublicationDataOk() ([]string, bool) {
	if o == nil || IsNil(o.OtherPublicationData) {
		return nil, false
	}
	return o.OtherPublicationData, true
}

// HasOtherPublicationData returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) HasOtherPublicationData() bool {
	if o != nil && !IsNil(o.OtherPublicationData) {
		return true
	}

	return false
}

// SetOtherPublicationData gets a reference to the given []string and assigns it to the OtherPublicationData field.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) SetOtherPublicationData(v []string) {
	o.OtherPublicationData = v
}

// GetRelatedParts returns the RelatedParts field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) GetRelatedParts() []string {
	if o == nil || IsNil(o.RelatedParts) {
		var ret []string
		return ret
	}
	return o.RelatedParts
}

// GetRelatedPartsOk returns a tuple with the RelatedParts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) GetRelatedPartsOk() ([]string, bool) {
	if o == nil || IsNil(o.RelatedParts) {
		return nil, false
	}
	return o.RelatedParts, true
}

// HasRelatedParts returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) HasRelatedParts() bool {
	if o != nil && !IsNil(o.RelatedParts) {
		return true
	}

	return false
}

// SetRelatedParts gets a reference to the given []string and assigns it to the RelatedParts field.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) SetRelatedParts(v []string) {
	o.RelatedParts = v
}

// GetSourceTitle returns the SourceTitle field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) GetSourceTitle() string {
	if o == nil || IsNil(o.SourceTitle) {
		var ret string
		return ret
	}
	return *o.SourceTitle
}

// GetSourceTitleOk returns a tuple with the SourceTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) GetSourceTitleOk() (*string, bool) {
	if o == nil || IsNil(o.SourceTitle) {
		return nil, false
	}
	return o.SourceTitle, true
}

// HasSourceTitle returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) HasSourceTitle() bool {
	if o != nil && !IsNil(o.SourceTitle) {
		return true
	}

	return false
}

// SetSourceTitle gets a reference to the given string and assigns it to the SourceTitle field.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) SetSourceTitle(v string) {
	o.SourceTitle = &v
}

// GetSourceIssn returns the SourceIssn field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) GetSourceIssn() string {
	if o == nil || IsNil(o.SourceIssn) {
		var ret string
		return ret
	}
	return *o.SourceIssn
}

// GetSourceIssnOk returns a tuple with the SourceIssn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) GetSourceIssnOk() (*string, bool) {
	if o == nil || IsNil(o.SourceIssn) {
		return nil, false
	}
	return o.SourceIssn, true
}

// HasSourceIssn returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) HasSourceIssn() bool {
	if o != nil && !IsNil(o.SourceIssn) {
		return true
	}

	return false
}

// SetSourceIssn gets a reference to the given string and assigns it to the SourceIssn field.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) SetSourceIssn(v string) {
	o.SourceIssn = &v
}

// GetSourceIsbns returns the SourceIsbns field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) GetSourceIsbns() []string {
	if o == nil || IsNil(o.SourceIsbns) {
		var ret []string
		return ret
	}
	return o.SourceIsbns
}

// GetSourceIsbnsOk returns a tuple with the SourceIsbns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) GetSourceIsbnsOk() ([]string, bool) {
	if o == nil || IsNil(o.SourceIsbns) {
		return nil, false
	}
	return o.SourceIsbns, true
}

// HasSourceIsbns returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) HasSourceIsbns() bool {
	if o != nil && !IsNil(o.SourceIsbns) {
		return true
	}

	return false
}

// SetSourceIsbns gets a reference to the given []string and assigns it to the SourceIsbns field.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) SetSourceIsbns(v []string) {
	o.SourceIsbns = v
}

// GetSourceAuthor returns the SourceAuthor field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) GetSourceAuthor() string {
	if o == nil || IsNil(o.SourceAuthor) {
		var ret string
		return ret
	}
	return *o.SourceAuthor
}

// GetSourceAuthorOk returns a tuple with the SourceAuthor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) GetSourceAuthorOk() (*string, bool) {
	if o == nil || IsNil(o.SourceAuthor) {
		return nil, false
	}
	return o.SourceAuthor, true
}

// HasSourceAuthor returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) HasSourceAuthor() bool {
	if o != nil && !IsNil(o.SourceAuthor) {
		return true
	}

	return false
}

// SetSourceAuthor gets a reference to the given string and assigns it to the SourceAuthor field.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) SetSourceAuthor(v string) {
	o.SourceAuthor = &v
}

// GetPublicationInformation returns the PublicationInformation field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) GetPublicationInformation() string {
	if o == nil || IsNil(o.PublicationInformation) {
		var ret string
		return ret
	}
	return *o.PublicationInformation
}

// GetPublicationInformationOk returns a tuple with the PublicationInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) GetPublicationInformationOk() (*string, bool) {
	if o == nil || IsNil(o.PublicationInformation) {
		return nil, false
	}
	return o.PublicationInformation, true
}

// HasPublicationInformation returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) HasPublicationInformation() bool {
	if o != nil && !IsNil(o.PublicationInformation) {
		return true
	}

	return false
}

// SetPublicationInformation gets a reference to the given string and assigns it to the PublicationInformation field.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) SetPublicationInformation(v string) {
	o.PublicationInformation = &v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) GetVolumes() []string {
	if o == nil || IsNil(o.Volumes) {
		var ret []string
		return ret
	}
	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) GetVolumesOk() ([]string, bool) {
	if o == nil || IsNil(o.Volumes) {
		return nil, false
	}
	return o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) HasVolumes() bool {
	if o != nil && !IsNil(o.Volumes) {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []string and assigns it to the Volumes field.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) SetVolumes(v []string) {
	o.Volumes = v
}

// GetIssues returns the Issues field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) GetIssues() []string {
	if o == nil || IsNil(o.Issues) {
		var ret []string
		return ret
	}
	return o.Issues
}

// GetIssuesOk returns a tuple with the Issues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) GetIssuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Issues) {
		return nil, false
	}
	return o.Issues, true
}

// HasIssues returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) HasIssues() bool {
	if o != nil && !IsNil(o.Issues) {
		return true
	}

	return false
}

// SetIssues gets a reference to the given []string and assigns it to the Issues field.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) SetIssues(v []string) {
	o.Issues = v
}

// GetPageRanges returns the PageRanges field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) GetPageRanges() []string {
	if o == nil || IsNil(o.PageRanges) {
		var ret []string
		return ret
	}
	return o.PageRanges
}

// GetPageRangesOk returns a tuple with the PageRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) GetPageRangesOk() ([]string, bool) {
	if o == nil || IsNil(o.PageRanges) {
		return nil, false
	}
	return o.PageRanges, true
}

// HasPageRanges returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) HasPageRanges() bool {
	if o != nil && !IsNil(o.PageRanges) {
		return true
	}

	return false
}

// SetPageRanges gets a reference to the given []string and assigns it to the PageRanges field.
func (o *SearchBibs200ResponseBibRecordsInnerSourceInner) SetPageRanges(v []string) {
	o.PageRanges = v
}

func (o SearchBibs200ResponseBibRecordsInnerSourceInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchBibs200ResponseBibRecordsInnerSourceInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Supplements) {
		toSerialize["supplements"] = o.Supplements
	}
	if !IsNil(o.FirstPages) {
		toSerialize["firstPages"] = o.FirstPages
	}
	if !IsNil(o.IssueDates) {
		toSerialize["issueDates"] = o.IssueDates
	}
	if !IsNil(o.MainEntry) {
		toSerialize["mainEntry"] = o.MainEntry
	}
	if !IsNil(o.OtherPublicationData) {
		toSerialize["otherPublicationData"] = o.OtherPublicationData
	}
	if !IsNil(o.RelatedParts) {
		toSerialize["relatedParts"] = o.RelatedParts
	}
	if !IsNil(o.SourceTitle) {
		toSerialize["sourceTitle"] = o.SourceTitle
	}
	if !IsNil(o.SourceIssn) {
		toSerialize["sourceIssn"] = o.SourceIssn
	}
	if !IsNil(o.SourceIsbns) {
		toSerialize["sourceIsbns"] = o.SourceIsbns
	}
	if !IsNil(o.SourceAuthor) {
		toSerialize["sourceAuthor"] = o.SourceAuthor
	}
	if !IsNil(o.PublicationInformation) {
		toSerialize["publicationInformation"] = o.PublicationInformation
	}
	if !IsNil(o.Volumes) {
		toSerialize["volumes"] = o.Volumes
	}
	if !IsNil(o.Issues) {
		toSerialize["issues"] = o.Issues
	}
	if !IsNil(o.PageRanges) {
		toSerialize["pageRanges"] = o.PageRanges
	}
	return toSerialize, nil
}

type NullableSearchBibs200ResponseBibRecordsInnerSourceInner struct {
	value *SearchBibs200ResponseBibRecordsInnerSourceInner
	isSet bool
}

func (v NullableSearchBibs200ResponseBibRecordsInnerSourceInner) Get() *SearchBibs200ResponseBibRecordsInnerSourceInner {
	return v.value
}

func (v *NullableSearchBibs200ResponseBibRecordsInnerSourceInner) Set(val *SearchBibs200ResponseBibRecordsInnerSourceInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchBibs200ResponseBibRecordsInnerSourceInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchBibs200ResponseBibRecordsInnerSourceInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchBibs200ResponseBibRecordsInnerSourceInner(val *SearchBibs200ResponseBibRecordsInnerSourceInner) *NullableSearchBibs200ResponseBibRecordsInnerSourceInner {
	return &NullableSearchBibs200ResponseBibRecordsInnerSourceInner{value: val, isSet: true}
}

func (v NullableSearchBibs200ResponseBibRecordsInnerSourceInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchBibs200ResponseBibRecordsInnerSourceInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}