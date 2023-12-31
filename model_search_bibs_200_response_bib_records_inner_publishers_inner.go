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

// checks if the SearchBibs200ResponseBibRecordsInnerPublishersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchBibs200ResponseBibRecordsInnerPublishersInner{}

// SearchBibs200ResponseBibRecordsInnerPublishersInner Publisher Information [v260sa,b v264sa,b v880sa,b]
type SearchBibs200ResponseBibRecordsInnerPublishersInner struct {
	PublisherName *SearchBibs200ResponseBibRecordsInnerPublishersInnerPublisherName `json:"publisherName,omitempty"`
	// Publication Place [v260sa v264sa]
	PublicationPlace *string `json:"publicationPlace,omitempty"`
}

// NewSearchBibs200ResponseBibRecordsInnerPublishersInner instantiates a new SearchBibs200ResponseBibRecordsInnerPublishersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchBibs200ResponseBibRecordsInnerPublishersInner() *SearchBibs200ResponseBibRecordsInnerPublishersInner {
	this := SearchBibs200ResponseBibRecordsInnerPublishersInner{}
	return &this
}

// NewSearchBibs200ResponseBibRecordsInnerPublishersInnerWithDefaults instantiates a new SearchBibs200ResponseBibRecordsInnerPublishersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchBibs200ResponseBibRecordsInnerPublishersInnerWithDefaults() *SearchBibs200ResponseBibRecordsInnerPublishersInner {
	this := SearchBibs200ResponseBibRecordsInnerPublishersInner{}
	return &this
}

// GetPublisherName returns the PublisherName field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerPublishersInner) GetPublisherName() SearchBibs200ResponseBibRecordsInnerPublishersInnerPublisherName {
	if o == nil || IsNil(o.PublisherName) {
		var ret SearchBibs200ResponseBibRecordsInnerPublishersInnerPublisherName
		return ret
	}
	return *o.PublisherName
}

// GetPublisherNameOk returns a tuple with the PublisherName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerPublishersInner) GetPublisherNameOk() (*SearchBibs200ResponseBibRecordsInnerPublishersInnerPublisherName, bool) {
	if o == nil || IsNil(o.PublisherName) {
		return nil, false
	}
	return o.PublisherName, true
}

// HasPublisherName returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerPublishersInner) HasPublisherName() bool {
	if o != nil && !IsNil(o.PublisherName) {
		return true
	}

	return false
}

// SetPublisherName gets a reference to the given SearchBibs200ResponseBibRecordsInnerPublishersInnerPublisherName and assigns it to the PublisherName field.
func (o *SearchBibs200ResponseBibRecordsInnerPublishersInner) SetPublisherName(v SearchBibs200ResponseBibRecordsInnerPublishersInnerPublisherName) {
	o.PublisherName = &v
}

// GetPublicationPlace returns the PublicationPlace field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerPublishersInner) GetPublicationPlace() string {
	if o == nil || IsNil(o.PublicationPlace) {
		var ret string
		return ret
	}
	return *o.PublicationPlace
}

// GetPublicationPlaceOk returns a tuple with the PublicationPlace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerPublishersInner) GetPublicationPlaceOk() (*string, bool) {
	if o == nil || IsNil(o.PublicationPlace) {
		return nil, false
	}
	return o.PublicationPlace, true
}

// HasPublicationPlace returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerPublishersInner) HasPublicationPlace() bool {
	if o != nil && !IsNil(o.PublicationPlace) {
		return true
	}

	return false
}

// SetPublicationPlace gets a reference to the given string and assigns it to the PublicationPlace field.
func (o *SearchBibs200ResponseBibRecordsInnerPublishersInner) SetPublicationPlace(v string) {
	o.PublicationPlace = &v
}

func (o SearchBibs200ResponseBibRecordsInnerPublishersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchBibs200ResponseBibRecordsInnerPublishersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PublisherName) {
		toSerialize["publisherName"] = o.PublisherName
	}
	if !IsNil(o.PublicationPlace) {
		toSerialize["publicationPlace"] = o.PublicationPlace
	}
	return toSerialize, nil
}

type NullableSearchBibs200ResponseBibRecordsInnerPublishersInner struct {
	value *SearchBibs200ResponseBibRecordsInnerPublishersInner
	isSet bool
}

func (v NullableSearchBibs200ResponseBibRecordsInnerPublishersInner) Get() *SearchBibs200ResponseBibRecordsInnerPublishersInner {
	return v.value
}

func (v *NullableSearchBibs200ResponseBibRecordsInnerPublishersInner) Set(val *SearchBibs200ResponseBibRecordsInnerPublishersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchBibs200ResponseBibRecordsInnerPublishersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchBibs200ResponseBibRecordsInnerPublishersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchBibs200ResponseBibRecordsInnerPublishersInner(val *SearchBibs200ResponseBibRecordsInnerPublishersInner) *NullableSearchBibs200ResponseBibRecordsInnerPublishersInner {
	return &NullableSearchBibs200ResponseBibRecordsInnerPublishersInner{value: val, isSet: true}
}

func (v NullableSearchBibs200ResponseBibRecordsInnerPublishersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchBibs200ResponseBibRecordsInnerPublishersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
