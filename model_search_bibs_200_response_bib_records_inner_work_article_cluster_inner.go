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

// checks if the SearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner{}

// SearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner struct for SearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner
type SearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner struct {
	DatabaseCollection *string `json:"databaseCollection,omitempty"`
	// the oclc number of a given bibliographic record
	OclcNumber *int64 `json:"oclcNumber,omitempty"`
}

// NewSearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner instantiates a new SearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner() *SearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner {
	this := SearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner{}
	return &this
}

// NewSearchBibs200ResponseBibRecordsInnerWorkArticleClusterInnerWithDefaults instantiates a new SearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchBibs200ResponseBibRecordsInnerWorkArticleClusterInnerWithDefaults() *SearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner {
	this := SearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner{}
	return &this
}

// GetDatabaseCollection returns the DatabaseCollection field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner) GetDatabaseCollection() string {
	if o == nil || IsNil(o.DatabaseCollection) {
		var ret string
		return ret
	}
	return *o.DatabaseCollection
}

// GetDatabaseCollectionOk returns a tuple with the DatabaseCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner) GetDatabaseCollectionOk() (*string, bool) {
	if o == nil || IsNil(o.DatabaseCollection) {
		return nil, false
	}
	return o.DatabaseCollection, true
}

// HasDatabaseCollection returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner) HasDatabaseCollection() bool {
	if o != nil && !IsNil(o.DatabaseCollection) {
		return true
	}

	return false
}

// SetDatabaseCollection gets a reference to the given string and assigns it to the DatabaseCollection field.
func (o *SearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner) SetDatabaseCollection(v string) {
	o.DatabaseCollection = &v
}

// GetOclcNumber returns the OclcNumber field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner) GetOclcNumber() int64 {
	if o == nil || IsNil(o.OclcNumber) {
		var ret int64
		return ret
	}
	return *o.OclcNumber
}

// GetOclcNumberOk returns a tuple with the OclcNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner) GetOclcNumberOk() (*int64, bool) {
	if o == nil || IsNil(o.OclcNumber) {
		return nil, false
	}
	return o.OclcNumber, true
}

// HasOclcNumber returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner) HasOclcNumber() bool {
	if o != nil && !IsNil(o.OclcNumber) {
		return true
	}

	return false
}

// SetOclcNumber gets a reference to the given int64 and assigns it to the OclcNumber field.
func (o *SearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner) SetOclcNumber(v int64) {
	o.OclcNumber = &v
}

func (o SearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DatabaseCollection) {
		toSerialize["databaseCollection"] = o.DatabaseCollection
	}
	if !IsNil(o.OclcNumber) {
		toSerialize["oclcNumber"] = o.OclcNumber
	}
	return toSerialize, nil
}

type NullableSearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner struct {
	value *SearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner
	isSet bool
}

func (v NullableSearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner) Get() *SearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner {
	return v.value
}

func (v *NullableSearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner) Set(val *SearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner(val *SearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner) *NullableSearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner {
	return &NullableSearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner{value: val, isSet: true}
}

func (v NullableSearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchBibs200ResponseBibRecordsInnerWorkArticleClusterInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
