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

// checks if the FindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner{}

// FindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner struct for FindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner
type FindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner struct {
	Value *string `json:"value,omitempty"`
	Count *int32  `json:"count,omitempty"`
}

// NewFindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner instantiates a new FindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner() *FindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner {
	this := FindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner{}
	return &this
}

// NewFindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInnerWithDefaults instantiates a new FindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInnerWithDefaults() *FindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner {
	this := FindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *FindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner) SetValue(v string) {
	o.Value = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *FindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *FindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *FindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner) SetCount(v int32) {
	o.Count = &v
}

func (o FindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableFindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner struct {
	value *FindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner
	isSet bool
}

func (v NullableFindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner) Get() *FindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner {
	return v.value
}

func (v *NullableFindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner) Set(val *FindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner(val *FindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner) *NullableFindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner {
	return &NullableFindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner{value: val, isSet: true}
}

func (v NullableFindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindBibRetainedHoldings200ResponseSearchFacetsInnerValuesInnerSubFacetValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
