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

// checks if the SearchLhr200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchLhr200Response{}

// SearchLhr200Response struct for SearchLhr200Response
type SearchLhr200Response struct {
	NumberOfHoldings *int32                                                                             `json:"numberOfHoldings,omitempty"`
	DetailedHoldings []FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1 `json:"detailedHoldings,omitempty"`
}

// NewSearchLhr200Response instantiates a new SearchLhr200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchLhr200Response() *SearchLhr200Response {
	this := SearchLhr200Response{}
	return &this
}

// NewSearchLhr200ResponseWithDefaults instantiates a new SearchLhr200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchLhr200ResponseWithDefaults() *SearchLhr200Response {
	this := SearchLhr200Response{}
	return &this
}

// GetNumberOfHoldings returns the NumberOfHoldings field value if set, zero value otherwise.
func (o *SearchLhr200Response) GetNumberOfHoldings() int32 {
	if o == nil || IsNil(o.NumberOfHoldings) {
		var ret int32
		return ret
	}
	return *o.NumberOfHoldings
}

// GetNumberOfHoldingsOk returns a tuple with the NumberOfHoldings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchLhr200Response) GetNumberOfHoldingsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfHoldings) {
		return nil, false
	}
	return o.NumberOfHoldings, true
}

// HasNumberOfHoldings returns a boolean if a field has been set.
func (o *SearchLhr200Response) HasNumberOfHoldings() bool {
	if o != nil && !IsNil(o.NumberOfHoldings) {
		return true
	}

	return false
}

// SetNumberOfHoldings gets a reference to the given int32 and assigns it to the NumberOfHoldings field.
func (o *SearchLhr200Response) SetNumberOfHoldings(v int32) {
	o.NumberOfHoldings = &v
}

// GetDetailedHoldings returns the DetailedHoldings field value if set, zero value otherwise.
func (o *SearchLhr200Response) GetDetailedHoldings() []FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1 {
	if o == nil || IsNil(o.DetailedHoldings) {
		var ret []FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1
		return ret
	}
	return o.DetailedHoldings
}

// GetDetailedHoldingsOk returns a tuple with the DetailedHoldings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchLhr200Response) GetDetailedHoldingsOk() ([]FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1, bool) {
	if o == nil || IsNil(o.DetailedHoldings) {
		return nil, false
	}
	return o.DetailedHoldings, true
}

// HasDetailedHoldings returns a boolean if a field has been set.
func (o *SearchLhr200Response) HasDetailedHoldings() bool {
	if o != nil && !IsNil(o.DetailedHoldings) {
		return true
	}

	return false
}

// SetDetailedHoldings gets a reference to the given []FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1 and assigns it to the DetailedHoldings field.
func (o *SearchLhr200Response) SetDetailedHoldings(v []FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) {
	o.DetailedHoldings = v
}

func (o SearchLhr200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchLhr200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NumberOfHoldings) {
		toSerialize["numberOfHoldings"] = o.NumberOfHoldings
	}
	if !IsNil(o.DetailedHoldings) {
		toSerialize["detailedHoldings"] = o.DetailedHoldings
	}
	return toSerialize, nil
}

type NullableSearchLhr200Response struct {
	value *SearchLhr200Response
	isSet bool
}

func (v NullableSearchLhr200Response) Get() *SearchLhr200Response {
	return v.value
}

func (v *NullableSearchLhr200Response) Set(val *SearchLhr200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchLhr200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchLhr200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchLhr200Response(val *SearchLhr200Response) *NullableSearchLhr200Response {
	return &NullableSearchLhr200Response{value: val, isSet: true}
}

func (v NullableSearchLhr200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchLhr200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
