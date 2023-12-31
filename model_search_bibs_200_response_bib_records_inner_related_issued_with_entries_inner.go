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

// checks if the SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner{}

// SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner struct for SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner
type SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner struct {
	DisplayConstant          *string  `json:"displayConstant,omitempty"`
	Titles                   []string `json:"titles,omitempty"`
	RecordControlOclcNumbers []string `json:"recordControlOclcNumbers,omitempty"`
	// International Standard Serial Numbers
	Issns []string `json:"issns,omitempty"`
	// International Standard Book Number [v020sa]
	Isbns             []string `json:"isbns,omitempty"`
	MainEntryHeadings []string `json:"mainEntryHeadings,omitempty"`
	RelatedParts      []string `json:"relatedParts,omitempty"`
	UniformTitle      *string  `json:"uniformTitle,omitempty"`
}

// NewSearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner instantiates a new SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner() *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner {
	this := SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner{}
	return &this
}

// NewSearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInnerWithDefaults instantiates a new SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInnerWithDefaults() *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner {
	this := SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner{}
	return &this
}

// GetDisplayConstant returns the DisplayConstant field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) GetDisplayConstant() string {
	if o == nil || IsNil(o.DisplayConstant) {
		var ret string
		return ret
	}
	return *o.DisplayConstant
}

// GetDisplayConstantOk returns a tuple with the DisplayConstant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) GetDisplayConstantOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayConstant) {
		return nil, false
	}
	return o.DisplayConstant, true
}

// HasDisplayConstant returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) HasDisplayConstant() bool {
	if o != nil && !IsNil(o.DisplayConstant) {
		return true
	}

	return false
}

// SetDisplayConstant gets a reference to the given string and assigns it to the DisplayConstant field.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) SetDisplayConstant(v string) {
	o.DisplayConstant = &v
}

// GetTitles returns the Titles field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) GetTitles() []string {
	if o == nil || IsNil(o.Titles) {
		var ret []string
		return ret
	}
	return o.Titles
}

// GetTitlesOk returns a tuple with the Titles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) GetTitlesOk() ([]string, bool) {
	if o == nil || IsNil(o.Titles) {
		return nil, false
	}
	return o.Titles, true
}

// HasTitles returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) HasTitles() bool {
	if o != nil && !IsNil(o.Titles) {
		return true
	}

	return false
}

// SetTitles gets a reference to the given []string and assigns it to the Titles field.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) SetTitles(v []string) {
	o.Titles = v
}

// GetRecordControlOclcNumbers returns the RecordControlOclcNumbers field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) GetRecordControlOclcNumbers() []string {
	if o == nil || IsNil(o.RecordControlOclcNumbers) {
		var ret []string
		return ret
	}
	return o.RecordControlOclcNumbers
}

// GetRecordControlOclcNumbersOk returns a tuple with the RecordControlOclcNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) GetRecordControlOclcNumbersOk() ([]string, bool) {
	if o == nil || IsNil(o.RecordControlOclcNumbers) {
		return nil, false
	}
	return o.RecordControlOclcNumbers, true
}

// HasRecordControlOclcNumbers returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) HasRecordControlOclcNumbers() bool {
	if o != nil && !IsNil(o.RecordControlOclcNumbers) {
		return true
	}

	return false
}

// SetRecordControlOclcNumbers gets a reference to the given []string and assigns it to the RecordControlOclcNumbers field.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) SetRecordControlOclcNumbers(v []string) {
	o.RecordControlOclcNumbers = v
}

// GetIssns returns the Issns field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) GetIssns() []string {
	if o == nil || IsNil(o.Issns) {
		var ret []string
		return ret
	}
	return o.Issns
}

// GetIssnsOk returns a tuple with the Issns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) GetIssnsOk() ([]string, bool) {
	if o == nil || IsNil(o.Issns) {
		return nil, false
	}
	return o.Issns, true
}

// HasIssns returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) HasIssns() bool {
	if o != nil && !IsNil(o.Issns) {
		return true
	}

	return false
}

// SetIssns gets a reference to the given []string and assigns it to the Issns field.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) SetIssns(v []string) {
	o.Issns = v
}

// GetIsbns returns the Isbns field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) GetIsbns() []string {
	if o == nil || IsNil(o.Isbns) {
		var ret []string
		return ret
	}
	return o.Isbns
}

// GetIsbnsOk returns a tuple with the Isbns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) GetIsbnsOk() ([]string, bool) {
	if o == nil || IsNil(o.Isbns) {
		return nil, false
	}
	return o.Isbns, true
}

// HasIsbns returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) HasIsbns() bool {
	if o != nil && !IsNil(o.Isbns) {
		return true
	}

	return false
}

// SetIsbns gets a reference to the given []string and assigns it to the Isbns field.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) SetIsbns(v []string) {
	o.Isbns = v
}

// GetMainEntryHeadings returns the MainEntryHeadings field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) GetMainEntryHeadings() []string {
	if o == nil || IsNil(o.MainEntryHeadings) {
		var ret []string
		return ret
	}
	return o.MainEntryHeadings
}

// GetMainEntryHeadingsOk returns a tuple with the MainEntryHeadings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) GetMainEntryHeadingsOk() ([]string, bool) {
	if o == nil || IsNil(o.MainEntryHeadings) {
		return nil, false
	}
	return o.MainEntryHeadings, true
}

// HasMainEntryHeadings returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) HasMainEntryHeadings() bool {
	if o != nil && !IsNil(o.MainEntryHeadings) {
		return true
	}

	return false
}

// SetMainEntryHeadings gets a reference to the given []string and assigns it to the MainEntryHeadings field.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) SetMainEntryHeadings(v []string) {
	o.MainEntryHeadings = v
}

// GetRelatedParts returns the RelatedParts field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) GetRelatedParts() []string {
	if o == nil || IsNil(o.RelatedParts) {
		var ret []string
		return ret
	}
	return o.RelatedParts
}

// GetRelatedPartsOk returns a tuple with the RelatedParts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) GetRelatedPartsOk() ([]string, bool) {
	if o == nil || IsNil(o.RelatedParts) {
		return nil, false
	}
	return o.RelatedParts, true
}

// HasRelatedParts returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) HasRelatedParts() bool {
	if o != nil && !IsNil(o.RelatedParts) {
		return true
	}

	return false
}

// SetRelatedParts gets a reference to the given []string and assigns it to the RelatedParts field.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) SetRelatedParts(v []string) {
	o.RelatedParts = v
}

// GetUniformTitle returns the UniformTitle field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) GetUniformTitle() string {
	if o == nil || IsNil(o.UniformTitle) {
		var ret string
		return ret
	}
	return *o.UniformTitle
}

// GetUniformTitleOk returns a tuple with the UniformTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) GetUniformTitleOk() (*string, bool) {
	if o == nil || IsNil(o.UniformTitle) {
		return nil, false
	}
	return o.UniformTitle, true
}

// HasUniformTitle returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) HasUniformTitle() bool {
	if o != nil && !IsNil(o.UniformTitle) {
		return true
	}

	return false
}

// SetUniformTitle gets a reference to the given string and assigns it to the UniformTitle field.
func (o *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) SetUniformTitle(v string) {
	o.UniformTitle = &v
}

func (o SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayConstant) {
		toSerialize["displayConstant"] = o.DisplayConstant
	}
	if !IsNil(o.Titles) {
		toSerialize["titles"] = o.Titles
	}
	if !IsNil(o.RecordControlOclcNumbers) {
		toSerialize["recordControlOclcNumbers"] = o.RecordControlOclcNumbers
	}
	if !IsNil(o.Issns) {
		toSerialize["issns"] = o.Issns
	}
	if !IsNil(o.Isbns) {
		toSerialize["isbns"] = o.Isbns
	}
	if !IsNil(o.MainEntryHeadings) {
		toSerialize["mainEntryHeadings"] = o.MainEntryHeadings
	}
	if !IsNil(o.RelatedParts) {
		toSerialize["relatedParts"] = o.RelatedParts
	}
	if !IsNil(o.UniformTitle) {
		toSerialize["uniformTitle"] = o.UniformTitle
	}
	return toSerialize, nil
}

type NullableSearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner struct {
	value *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner
	isSet bool
}

func (v NullableSearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) Get() *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner {
	return v.value
}

func (v *NullableSearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) Set(val *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner(val *SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) *NullableSearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner {
	return &NullableSearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner{value: val, isSet: true}
}

func (v NullableSearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
