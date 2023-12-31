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

// checks if the SearchBibs200ResponseBibRecordsInnerRelated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchBibs200ResponseBibRecordsInnerRelated{}

// SearchBibs200ResponseBibRecordsInnerRelated Related [v787] Preceeding [v780] Succeeding [v785] (sa,b,c,d,g,h,i,k,m,n,o,r,s,t,u,w,x,y,z)
type SearchBibs200ResponseBibRecordsInnerRelated struct {
	RelatedItems      []SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem `json:"relatedItems,omitempty"`
	PrecedingEntries  []SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem `json:"precedingEntries,omitempty"`
	SucceedingEntries []SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem `json:"succeedingEntries,omitempty"`
	// [v777]
	IssuedWithEntries             []SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner `json:"issuedWithEntries,omitempty"`
	AdditionalPhysicalFormEntries []SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner `json:"additionalPhysicalFormEntries,omitempty"`
}

// NewSearchBibs200ResponseBibRecordsInnerRelated instantiates a new SearchBibs200ResponseBibRecordsInnerRelated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchBibs200ResponseBibRecordsInnerRelated() *SearchBibs200ResponseBibRecordsInnerRelated {
	this := SearchBibs200ResponseBibRecordsInnerRelated{}
	return &this
}

// NewSearchBibs200ResponseBibRecordsInnerRelatedWithDefaults instantiates a new SearchBibs200ResponseBibRecordsInnerRelated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchBibs200ResponseBibRecordsInnerRelatedWithDefaults() *SearchBibs200ResponseBibRecordsInnerRelated {
	this := SearchBibs200ResponseBibRecordsInnerRelated{}
	return &this
}

// GetRelatedItems returns the RelatedItems field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerRelated) GetRelatedItems() []SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem {
	if o == nil || IsNil(o.RelatedItems) {
		var ret []SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem
		return ret
	}
	return o.RelatedItems
}

// GetRelatedItemsOk returns a tuple with the RelatedItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerRelated) GetRelatedItemsOk() ([]SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem, bool) {
	if o == nil || IsNil(o.RelatedItems) {
		return nil, false
	}
	return o.RelatedItems, true
}

// HasRelatedItems returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerRelated) HasRelatedItems() bool {
	if o != nil && !IsNil(o.RelatedItems) {
		return true
	}

	return false
}

// SetRelatedItems gets a reference to the given []SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem and assigns it to the RelatedItems field.
func (o *SearchBibs200ResponseBibRecordsInnerRelated) SetRelatedItems(v []SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) {
	o.RelatedItems = v
}

// GetPrecedingEntries returns the PrecedingEntries field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerRelated) GetPrecedingEntries() []SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem {
	if o == nil || IsNil(o.PrecedingEntries) {
		var ret []SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem
		return ret
	}
	return o.PrecedingEntries
}

// GetPrecedingEntriesOk returns a tuple with the PrecedingEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerRelated) GetPrecedingEntriesOk() ([]SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem, bool) {
	if o == nil || IsNil(o.PrecedingEntries) {
		return nil, false
	}
	return o.PrecedingEntries, true
}

// HasPrecedingEntries returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerRelated) HasPrecedingEntries() bool {
	if o != nil && !IsNil(o.PrecedingEntries) {
		return true
	}

	return false
}

// SetPrecedingEntries gets a reference to the given []SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem and assigns it to the PrecedingEntries field.
func (o *SearchBibs200ResponseBibRecordsInnerRelated) SetPrecedingEntries(v []SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) {
	o.PrecedingEntries = v
}

// GetSucceedingEntries returns the SucceedingEntries field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerRelated) GetSucceedingEntries() []SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem {
	if o == nil || IsNil(o.SucceedingEntries) {
		var ret []SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem
		return ret
	}
	return o.SucceedingEntries
}

// GetSucceedingEntriesOk returns a tuple with the SucceedingEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerRelated) GetSucceedingEntriesOk() ([]SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem, bool) {
	if o == nil || IsNil(o.SucceedingEntries) {
		return nil, false
	}
	return o.SucceedingEntries, true
}

// HasSucceedingEntries returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerRelated) HasSucceedingEntries() bool {
	if o != nil && !IsNil(o.SucceedingEntries) {
		return true
	}

	return false
}

// SetSucceedingEntries gets a reference to the given []SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem and assigns it to the SucceedingEntries field.
func (o *SearchBibs200ResponseBibRecordsInnerRelated) SetSucceedingEntries(v []SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem) {
	o.SucceedingEntries = v
}

// GetIssuedWithEntries returns the IssuedWithEntries field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerRelated) GetIssuedWithEntries() []SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner {
	if o == nil || IsNil(o.IssuedWithEntries) {
		var ret []SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner
		return ret
	}
	return o.IssuedWithEntries
}

// GetIssuedWithEntriesOk returns a tuple with the IssuedWithEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerRelated) GetIssuedWithEntriesOk() ([]SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner, bool) {
	if o == nil || IsNil(o.IssuedWithEntries) {
		return nil, false
	}
	return o.IssuedWithEntries, true
}

// HasIssuedWithEntries returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerRelated) HasIssuedWithEntries() bool {
	if o != nil && !IsNil(o.IssuedWithEntries) {
		return true
	}

	return false
}

// SetIssuedWithEntries gets a reference to the given []SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner and assigns it to the IssuedWithEntries field.
func (o *SearchBibs200ResponseBibRecordsInnerRelated) SetIssuedWithEntries(v []SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) {
	o.IssuedWithEntries = v
}

// GetAdditionalPhysicalFormEntries returns the AdditionalPhysicalFormEntries field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerRelated) GetAdditionalPhysicalFormEntries() []SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner {
	if o == nil || IsNil(o.AdditionalPhysicalFormEntries) {
		var ret []SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner
		return ret
	}
	return o.AdditionalPhysicalFormEntries
}

// GetAdditionalPhysicalFormEntriesOk returns a tuple with the AdditionalPhysicalFormEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerRelated) GetAdditionalPhysicalFormEntriesOk() ([]SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner, bool) {
	if o == nil || IsNil(o.AdditionalPhysicalFormEntries) {
		return nil, false
	}
	return o.AdditionalPhysicalFormEntries, true
}

// HasAdditionalPhysicalFormEntries returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerRelated) HasAdditionalPhysicalFormEntries() bool {
	if o != nil && !IsNil(o.AdditionalPhysicalFormEntries) {
		return true
	}

	return false
}

// SetAdditionalPhysicalFormEntries gets a reference to the given []SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner and assigns it to the AdditionalPhysicalFormEntries field.
func (o *SearchBibs200ResponseBibRecordsInnerRelated) SetAdditionalPhysicalFormEntries(v []SearchBibs200ResponseBibRecordsInnerRelatedIssuedWithEntriesInner) {
	o.AdditionalPhysicalFormEntries = v
}

func (o SearchBibs200ResponseBibRecordsInnerRelated) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchBibs200ResponseBibRecordsInnerRelated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RelatedItems) {
		toSerialize["relatedItems"] = o.RelatedItems
	}
	if !IsNil(o.PrecedingEntries) {
		toSerialize["precedingEntries"] = o.PrecedingEntries
	}
	if !IsNil(o.SucceedingEntries) {
		toSerialize["succeedingEntries"] = o.SucceedingEntries
	}
	if !IsNil(o.IssuedWithEntries) {
		toSerialize["issuedWithEntries"] = o.IssuedWithEntries
	}
	if !IsNil(o.AdditionalPhysicalFormEntries) {
		toSerialize["additionalPhysicalFormEntries"] = o.AdditionalPhysicalFormEntries
	}
	return toSerialize, nil
}

type NullableSearchBibs200ResponseBibRecordsInnerRelated struct {
	value *SearchBibs200ResponseBibRecordsInnerRelated
	isSet bool
}

func (v NullableSearchBibs200ResponseBibRecordsInnerRelated) Get() *SearchBibs200ResponseBibRecordsInnerRelated {
	return v.value
}

func (v *NullableSearchBibs200ResponseBibRecordsInnerRelated) Set(val *SearchBibs200ResponseBibRecordsInnerRelated) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchBibs200ResponseBibRecordsInnerRelated) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchBibs200ResponseBibRecordsInnerRelated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchBibs200ResponseBibRecordsInnerRelated(val *SearchBibs200ResponseBibRecordsInnerRelated) *NullableSearchBibs200ResponseBibRecordsInnerRelated {
	return &NullableSearchBibs200ResponseBibRecordsInnerRelated{value: val, isSet: true}
}

func (v NullableSearchBibs200ResponseBibRecordsInnerRelated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchBibs200ResponseBibRecordsInnerRelated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
