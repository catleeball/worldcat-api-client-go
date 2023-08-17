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

// checks if the SearchBibs200ResponseBibRecordsInnerContributor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchBibs200ResponseBibRecordsInnerContributor{}

// SearchBibs200ResponseBibRecordsInnerContributor struct for SearchBibs200ResponseBibRecordsInnerContributor
type SearchBibs200ResponseBibRecordsInnerContributor struct {
	// Creators of work
	Creators []SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner `json:"creators,omitempty"`
	// Additional creators of work [(v700,v710,v711,v790,v791,v792,v796,v797,v798) - sa,b,c,d,q,n,k,e,j,4]
	AdditionalCreators        []SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInner          `json:"additionalCreators,omitempty"`
	StatementOfResponsibility *SearchBibs200ResponseBibRecordsInnerContributorStatementOfResponsibility `json:"statementOfResponsibility,omitempty"`
}

// NewSearchBibs200ResponseBibRecordsInnerContributor instantiates a new SearchBibs200ResponseBibRecordsInnerContributor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchBibs200ResponseBibRecordsInnerContributor() *SearchBibs200ResponseBibRecordsInnerContributor {
	this := SearchBibs200ResponseBibRecordsInnerContributor{}
	return &this
}

// NewSearchBibs200ResponseBibRecordsInnerContributorWithDefaults instantiates a new SearchBibs200ResponseBibRecordsInnerContributor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchBibs200ResponseBibRecordsInnerContributorWithDefaults() *SearchBibs200ResponseBibRecordsInnerContributor {
	this := SearchBibs200ResponseBibRecordsInnerContributor{}
	return &this
}

// GetCreators returns the Creators field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerContributor) GetCreators() []SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner {
	if o == nil || IsNil(o.Creators) {
		var ret []SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner
		return ret
	}
	return o.Creators
}

// GetCreatorsOk returns a tuple with the Creators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributor) GetCreatorsOk() ([]SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner, bool) {
	if o == nil || IsNil(o.Creators) {
		return nil, false
	}
	return o.Creators, true
}

// HasCreators returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributor) HasCreators() bool {
	if o != nil && !IsNil(o.Creators) {
		return true
	}

	return false
}

// SetCreators gets a reference to the given []SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner and assigns it to the Creators field.
func (o *SearchBibs200ResponseBibRecordsInnerContributor) SetCreators(v []SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) {
	o.Creators = v
}

// GetAdditionalCreators returns the AdditionalCreators field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerContributor) GetAdditionalCreators() []SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInner {
	if o == nil || IsNil(o.AdditionalCreators) {
		var ret []SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInner
		return ret
	}
	return o.AdditionalCreators
}

// GetAdditionalCreatorsOk returns a tuple with the AdditionalCreators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributor) GetAdditionalCreatorsOk() ([]SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInner, bool) {
	if o == nil || IsNil(o.AdditionalCreators) {
		return nil, false
	}
	return o.AdditionalCreators, true
}

// HasAdditionalCreators returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributor) HasAdditionalCreators() bool {
	if o != nil && !IsNil(o.AdditionalCreators) {
		return true
	}

	return false
}

// SetAdditionalCreators gets a reference to the given []SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInner and assigns it to the AdditionalCreators field.
func (o *SearchBibs200ResponseBibRecordsInnerContributor) SetAdditionalCreators(v []SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInner) {
	o.AdditionalCreators = v
}

// GetStatementOfResponsibility returns the StatementOfResponsibility field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerContributor) GetStatementOfResponsibility() SearchBibs200ResponseBibRecordsInnerContributorStatementOfResponsibility {
	if o == nil || IsNil(o.StatementOfResponsibility) {
		var ret SearchBibs200ResponseBibRecordsInnerContributorStatementOfResponsibility
		return ret
	}
	return *o.StatementOfResponsibility
}

// GetStatementOfResponsibilityOk returns a tuple with the StatementOfResponsibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributor) GetStatementOfResponsibilityOk() (*SearchBibs200ResponseBibRecordsInnerContributorStatementOfResponsibility, bool) {
	if o == nil || IsNil(o.StatementOfResponsibility) {
		return nil, false
	}
	return o.StatementOfResponsibility, true
}

// HasStatementOfResponsibility returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerContributor) HasStatementOfResponsibility() bool {
	if o != nil && !IsNil(o.StatementOfResponsibility) {
		return true
	}

	return false
}

// SetStatementOfResponsibility gets a reference to the given SearchBibs200ResponseBibRecordsInnerContributorStatementOfResponsibility and assigns it to the StatementOfResponsibility field.
func (o *SearchBibs200ResponseBibRecordsInnerContributor) SetStatementOfResponsibility(v SearchBibs200ResponseBibRecordsInnerContributorStatementOfResponsibility) {
	o.StatementOfResponsibility = &v
}

func (o SearchBibs200ResponseBibRecordsInnerContributor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchBibs200ResponseBibRecordsInnerContributor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Creators) {
		toSerialize["creators"] = o.Creators
	}
	if !IsNil(o.AdditionalCreators) {
		toSerialize["additionalCreators"] = o.AdditionalCreators
	}
	if !IsNil(o.StatementOfResponsibility) {
		toSerialize["statementOfResponsibility"] = o.StatementOfResponsibility
	}
	return toSerialize, nil
}

type NullableSearchBibs200ResponseBibRecordsInnerContributor struct {
	value *SearchBibs200ResponseBibRecordsInnerContributor
	isSet bool
}

func (v NullableSearchBibs200ResponseBibRecordsInnerContributor) Get() *SearchBibs200ResponseBibRecordsInnerContributor {
	return v.value
}

func (v *NullableSearchBibs200ResponseBibRecordsInnerContributor) Set(val *SearchBibs200ResponseBibRecordsInnerContributor) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchBibs200ResponseBibRecordsInnerContributor) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchBibs200ResponseBibRecordsInnerContributor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchBibs200ResponseBibRecordsInnerContributor(val *SearchBibs200ResponseBibRecordsInnerContributor) *NullableSearchBibs200ResponseBibRecordsInnerContributor {
	return &NullableSearchBibs200ResponseBibRecordsInnerContributor{value: val, isSet: true}
}

func (v NullableSearchBibs200ResponseBibRecordsInnerContributor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchBibs200ResponseBibRecordsInnerContributor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
