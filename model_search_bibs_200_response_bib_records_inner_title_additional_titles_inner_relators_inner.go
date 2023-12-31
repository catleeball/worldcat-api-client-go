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

// checks if the SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner{}

// SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner struct for SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner
type SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner struct {
	Term          *string `json:"term,omitempty"`
	AlternateTerm *string `json:"alternateTerm,omitempty"`
}

// NewSearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner instantiates a new SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner() *SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner {
	this := SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner{}
	return &this
}

// NewSearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInnerWithDefaults instantiates a new SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInnerWithDefaults() *SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner {
	this := SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner{}
	return &this
}

// GetTerm returns the Term field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner) GetTerm() string {
	if o == nil || IsNil(o.Term) {
		var ret string
		return ret
	}
	return *o.Term
}

// GetTermOk returns a tuple with the Term field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner) GetTermOk() (*string, bool) {
	if o == nil || IsNil(o.Term) {
		return nil, false
	}
	return o.Term, true
}

// HasTerm returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner) HasTerm() bool {
	if o != nil && !IsNil(o.Term) {
		return true
	}

	return false
}

// SetTerm gets a reference to the given string and assigns it to the Term field.
func (o *SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner) SetTerm(v string) {
	o.Term = &v
}

// GetAlternateTerm returns the AlternateTerm field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner) GetAlternateTerm() string {
	if o == nil || IsNil(o.AlternateTerm) {
		var ret string
		return ret
	}
	return *o.AlternateTerm
}

// GetAlternateTermOk returns a tuple with the AlternateTerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner) GetAlternateTermOk() (*string, bool) {
	if o == nil || IsNil(o.AlternateTerm) {
		return nil, false
	}
	return o.AlternateTerm, true
}

// HasAlternateTerm returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner) HasAlternateTerm() bool {
	if o != nil && !IsNil(o.AlternateTerm) {
		return true
	}

	return false
}

// SetAlternateTerm gets a reference to the given string and assigns it to the AlternateTerm field.
func (o *SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner) SetAlternateTerm(v string) {
	o.AlternateTerm = &v
}

func (o SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Term) {
		toSerialize["term"] = o.Term
	}
	if !IsNil(o.AlternateTerm) {
		toSerialize["alternateTerm"] = o.AlternateTerm
	}
	return toSerialize, nil
}

type NullableSearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner struct {
	value *SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner
	isSet bool
}

func (v NullableSearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner) Get() *SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner {
	return v.value
}

func (v *NullableSearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner) Set(val *SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner(val *SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner) *NullableSearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner {
	return &NullableSearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner{value: val, isSet: true}
}

func (v NullableSearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
