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

// checks if the FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo{}

// FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo struct for FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo
type FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo struct {
	// Cataloging Agency [v040/sa]
	CatalogingAgency *string `json:"catalogingAgency,omitempty"`
	// Transcribing Agency [v040/sc]
	TranscribingAgency *string `json:"transcribingAgency,omitempty"`
	// Language record was cataloged in [v040sb]
	CatalogingLanguage *string `json:"catalogingLanguage,omitempty"`
	// Level of cataloging [LDR position 5]
	LevelOfCataloging *string `json:"levelOfCataloging,omitempty"`
}

// NewFindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo instantiates a new FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo() *FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo {
	this := FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo{}
	return &this
}

// NewFindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfoWithDefaults instantiates a new FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfoWithDefaults() *FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo {
	this := FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo{}
	return &this
}

// GetCatalogingAgency returns the CatalogingAgency field value if set, zero value otherwise.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo) GetCatalogingAgency() string {
	if o == nil || IsNil(o.CatalogingAgency) {
		var ret string
		return ret
	}
	return *o.CatalogingAgency
}

// GetCatalogingAgencyOk returns a tuple with the CatalogingAgency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo) GetCatalogingAgencyOk() (*string, bool) {
	if o == nil || IsNil(o.CatalogingAgency) {
		return nil, false
	}
	return o.CatalogingAgency, true
}

// HasCatalogingAgency returns a boolean if a field has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo) HasCatalogingAgency() bool {
	if o != nil && !IsNil(o.CatalogingAgency) {
		return true
	}

	return false
}

// SetCatalogingAgency gets a reference to the given string and assigns it to the CatalogingAgency field.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo) SetCatalogingAgency(v string) {
	o.CatalogingAgency = &v
}

// GetTranscribingAgency returns the TranscribingAgency field value if set, zero value otherwise.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo) GetTranscribingAgency() string {
	if o == nil || IsNil(o.TranscribingAgency) {
		var ret string
		return ret
	}
	return *o.TranscribingAgency
}

// GetTranscribingAgencyOk returns a tuple with the TranscribingAgency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo) GetTranscribingAgencyOk() (*string, bool) {
	if o == nil || IsNil(o.TranscribingAgency) {
		return nil, false
	}
	return o.TranscribingAgency, true
}

// HasTranscribingAgency returns a boolean if a field has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo) HasTranscribingAgency() bool {
	if o != nil && !IsNil(o.TranscribingAgency) {
		return true
	}

	return false
}

// SetTranscribingAgency gets a reference to the given string and assigns it to the TranscribingAgency field.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo) SetTranscribingAgency(v string) {
	o.TranscribingAgency = &v
}

// GetCatalogingLanguage returns the CatalogingLanguage field value if set, zero value otherwise.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo) GetCatalogingLanguage() string {
	if o == nil || IsNil(o.CatalogingLanguage) {
		var ret string
		return ret
	}
	return *o.CatalogingLanguage
}

// GetCatalogingLanguageOk returns a tuple with the CatalogingLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo) GetCatalogingLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.CatalogingLanguage) {
		return nil, false
	}
	return o.CatalogingLanguage, true
}

// HasCatalogingLanguage returns a boolean if a field has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo) HasCatalogingLanguage() bool {
	if o != nil && !IsNil(o.CatalogingLanguage) {
		return true
	}

	return false
}

// SetCatalogingLanguage gets a reference to the given string and assigns it to the CatalogingLanguage field.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo) SetCatalogingLanguage(v string) {
	o.CatalogingLanguage = &v
}

// GetLevelOfCataloging returns the LevelOfCataloging field value if set, zero value otherwise.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo) GetLevelOfCataloging() string {
	if o == nil || IsNil(o.LevelOfCataloging) {
		var ret string
		return ret
	}
	return *o.LevelOfCataloging
}

// GetLevelOfCatalogingOk returns a tuple with the LevelOfCataloging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo) GetLevelOfCatalogingOk() (*string, bool) {
	if o == nil || IsNil(o.LevelOfCataloging) {
		return nil, false
	}
	return o.LevelOfCataloging, true
}

// HasLevelOfCataloging returns a boolean if a field has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo) HasLevelOfCataloging() bool {
	if o != nil && !IsNil(o.LevelOfCataloging) {
		return true
	}

	return false
}

// SetLevelOfCataloging gets a reference to the given string and assigns it to the LevelOfCataloging field.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo) SetLevelOfCataloging(v string) {
	o.LevelOfCataloging = &v
}

func (o FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CatalogingAgency) {
		toSerialize["catalogingAgency"] = o.CatalogingAgency
	}
	if !IsNil(o.TranscribingAgency) {
		toSerialize["transcribingAgency"] = o.TranscribingAgency
	}
	if !IsNil(o.CatalogingLanguage) {
		toSerialize["catalogingLanguage"] = o.CatalogingLanguage
	}
	if !IsNil(o.LevelOfCataloging) {
		toSerialize["levelOfCataloging"] = o.LevelOfCataloging
	}
	return toSerialize, nil
}

type NullableFindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo struct {
	value *FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo
	isSet bool
}

func (v NullableFindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo) Get() *FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo {
	return v.value
}

func (v *NullableFindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo) Set(val *FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo(val *FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo) *NullableFindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo {
	return &NullableFindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo{value: val, isSet: true}
}

func (v NullableFindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
