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

// checks if the SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner{}

// SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner struct for SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner
type SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner struct {
	Text          *string `json:"text,omitempty"`
	RomanizedText *string `json:"romanizedText,omitempty"`
	LanguageCode  *string `json:"languageCode,omitempty"`
	TextDirection *string `json:"textDirection,omitempty"`
}

// NewSearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner instantiates a new SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner() *SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner {
	this := SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner{}
	return &this
}

// NewSearchBibs200ResponseBibRecordsInnerTitleMainTitlesInnerWithDefaults instantiates a new SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchBibs200ResponseBibRecordsInnerTitleMainTitlesInnerWithDefaults() *SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner {
	this := SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner) SetText(v string) {
	o.Text = &v
}

// GetRomanizedText returns the RomanizedText field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner) GetRomanizedText() string {
	if o == nil || IsNil(o.RomanizedText) {
		var ret string
		return ret
	}
	return *o.RomanizedText
}

// GetRomanizedTextOk returns a tuple with the RomanizedText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner) GetRomanizedTextOk() (*string, bool) {
	if o == nil || IsNil(o.RomanizedText) {
		return nil, false
	}
	return o.RomanizedText, true
}

// HasRomanizedText returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner) HasRomanizedText() bool {
	if o != nil && !IsNil(o.RomanizedText) {
		return true
	}

	return false
}

// SetRomanizedText gets a reference to the given string and assigns it to the RomanizedText field.
func (o *SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner) SetRomanizedText(v string) {
	o.RomanizedText = &v
}

// GetLanguageCode returns the LanguageCode field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner) GetLanguageCode() string {
	if o == nil || IsNil(o.LanguageCode) {
		var ret string
		return ret
	}
	return *o.LanguageCode
}

// GetLanguageCodeOk returns a tuple with the LanguageCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner) GetLanguageCodeOk() (*string, bool) {
	if o == nil || IsNil(o.LanguageCode) {
		return nil, false
	}
	return o.LanguageCode, true
}

// HasLanguageCode returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner) HasLanguageCode() bool {
	if o != nil && !IsNil(o.LanguageCode) {
		return true
	}

	return false
}

// SetLanguageCode gets a reference to the given string and assigns it to the LanguageCode field.
func (o *SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner) SetLanguageCode(v string) {
	o.LanguageCode = &v
}

// GetTextDirection returns the TextDirection field value if set, zero value otherwise.
func (o *SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner) GetTextDirection() string {
	if o == nil || IsNil(o.TextDirection) {
		var ret string
		return ret
	}
	return *o.TextDirection
}

// GetTextDirectionOk returns a tuple with the TextDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner) GetTextDirectionOk() (*string, bool) {
	if o == nil || IsNil(o.TextDirection) {
		return nil, false
	}
	return o.TextDirection, true
}

// HasTextDirection returns a boolean if a field has been set.
func (o *SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner) HasTextDirection() bool {
	if o != nil && !IsNil(o.TextDirection) {
		return true
	}

	return false
}

// SetTextDirection gets a reference to the given string and assigns it to the TextDirection field.
func (o *SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner) SetTextDirection(v string) {
	o.TextDirection = &v
}

func (o SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.RomanizedText) {
		toSerialize["romanizedText"] = o.RomanizedText
	}
	if !IsNil(o.LanguageCode) {
		toSerialize["languageCode"] = o.LanguageCode
	}
	if !IsNil(o.TextDirection) {
		toSerialize["textDirection"] = o.TextDirection
	}
	return toSerialize, nil
}

type NullableSearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner struct {
	value *SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner
	isSet bool
}

func (v NullableSearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner) Get() *SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner {
	return v.value
}

func (v *NullableSearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner) Set(val *SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner(val *SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner) *NullableSearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner {
	return &NullableSearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner{value: val, isSet: true}
}

func (v NullableSearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
