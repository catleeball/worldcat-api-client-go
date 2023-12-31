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

// checks if the FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location{}

// FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location Local holding location [LHR -> v852sa,b,c]
type FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location struct {
	// Institution Symbol [LHR -> v852sa]
	HoldingLocation *string `json:"holdingLocation,omitempty"`
	// Sublocation or collection [LHR -> v852sb]
	SublocationCollection *string `json:"sublocationCollection,omitempty"`
	// Shelving location [LHR -> v852sc]
	ShelvingLocation *string `json:"shelvingLocation,omitempty"`
}

// NewFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location instantiates a new FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location() *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location {
	this := FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location{}
	return &this
}

// NewFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1LocationWithDefaults instantiates a new FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1LocationWithDefaults() *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location {
	this := FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location{}
	return &this
}

// GetHoldingLocation returns the HoldingLocation field value if set, zero value otherwise.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location) GetHoldingLocation() string {
	if o == nil || IsNil(o.HoldingLocation) {
		var ret string
		return ret
	}
	return *o.HoldingLocation
}

// GetHoldingLocationOk returns a tuple with the HoldingLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location) GetHoldingLocationOk() (*string, bool) {
	if o == nil || IsNil(o.HoldingLocation) {
		return nil, false
	}
	return o.HoldingLocation, true
}

// HasHoldingLocation returns a boolean if a field has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location) HasHoldingLocation() bool {
	if o != nil && !IsNil(o.HoldingLocation) {
		return true
	}

	return false
}

// SetHoldingLocation gets a reference to the given string and assigns it to the HoldingLocation field.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location) SetHoldingLocation(v string) {
	o.HoldingLocation = &v
}

// GetSublocationCollection returns the SublocationCollection field value if set, zero value otherwise.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location) GetSublocationCollection() string {
	if o == nil || IsNil(o.SublocationCollection) {
		var ret string
		return ret
	}
	return *o.SublocationCollection
}

// GetSublocationCollectionOk returns a tuple with the SublocationCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location) GetSublocationCollectionOk() (*string, bool) {
	if o == nil || IsNil(o.SublocationCollection) {
		return nil, false
	}
	return o.SublocationCollection, true
}

// HasSublocationCollection returns a boolean if a field has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location) HasSublocationCollection() bool {
	if o != nil && !IsNil(o.SublocationCollection) {
		return true
	}

	return false
}

// SetSublocationCollection gets a reference to the given string and assigns it to the SublocationCollection field.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location) SetSublocationCollection(v string) {
	o.SublocationCollection = &v
}

// GetShelvingLocation returns the ShelvingLocation field value if set, zero value otherwise.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location) GetShelvingLocation() string {
	if o == nil || IsNil(o.ShelvingLocation) {
		var ret string
		return ret
	}
	return *o.ShelvingLocation
}

// GetShelvingLocationOk returns a tuple with the ShelvingLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location) GetShelvingLocationOk() (*string, bool) {
	if o == nil || IsNil(o.ShelvingLocation) {
		return nil, false
	}
	return o.ShelvingLocation, true
}

// HasShelvingLocation returns a boolean if a field has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location) HasShelvingLocation() bool {
	if o != nil && !IsNil(o.ShelvingLocation) {
		return true
	}

	return false
}

// SetShelvingLocation gets a reference to the given string and assigns it to the ShelvingLocation field.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location) SetShelvingLocation(v string) {
	o.ShelvingLocation = &v
}

func (o FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HoldingLocation) {
		toSerialize["holdingLocation"] = o.HoldingLocation
	}
	if !IsNil(o.SublocationCollection) {
		toSerialize["sublocationCollection"] = o.SublocationCollection
	}
	if !IsNil(o.ShelvingLocation) {
		toSerialize["shelvingLocation"] = o.ShelvingLocation
	}
	return toSerialize, nil
}

type NullableFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location struct {
	value *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location
	isSet bool
}

func (v NullableFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location) Get() *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location {
	return v.value
}

func (v *NullableFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location) Set(val *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location) {
	v.value = val
	v.isSet = true
}

func (v NullableFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location) IsSet() bool {
	return v.isSet
}

func (v *NullableFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location(val *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location) *NullableFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location {
	return &NullableFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location{value: val, isSet: true}
}

func (v NullableFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
