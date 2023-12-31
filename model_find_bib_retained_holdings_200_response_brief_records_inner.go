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

// checks if the FindBibRetainedHoldings200ResponseBriefRecordsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FindBibRetainedHoldings200ResponseBriefRecordsInner{}

// FindBibRetainedHoldings200ResponseBriefRecordsInner struct for FindBibRetainedHoldings200ResponseBriefRecordsInner
type FindBibRetainedHoldings200ResponseBriefRecordsInner struct {
	// the oclc number of a given bibliographic record
	OclcNumber int64 `json:"oclcNumber"`
	// Linked [v880sa,b] if present or [v245sa,b]
	Title string `json:"title"`
	// List of creators as single string
	Creator string `json:"creator"`
	// machine readable date 008
	Date string `json:"date"`
	// Language of the item [v041sa,j]
	Language string `json:"language"`
	// General Format Type [Admin/OCLCDef/StdRT]
	GeneralFormat *string `json:"generalFormat,omitempty"`
	// Specific Format Type [Admin/OCLCDef/StdRT2]
	SpecificFormat *string `json:"specificFormat,omitempty"`
	// Edition Statement [v250sa]
	Edition   string `json:"edition"`
	Publisher string `json:"publisher"`
	// International Standard Book Number [v020sa]
	Isbns []string `json:"isbns,omitempty"`
	// International Standard Serial Numbers
	Issns []string `json:"issns,omitempty"`
	// Merged OCLC numbers [v019sa]
	MergedOclcNumbers  []string                                                               `json:"mergedOclcNumbers,omitempty"`
	CatalogingInfo     *FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo     `json:"catalogingInfo,omitempty"`
	InstitutionHolding *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding `json:"institutionHolding,omitempty"`
}

// NewFindBibRetainedHoldings200ResponseBriefRecordsInner instantiates a new FindBibRetainedHoldings200ResponseBriefRecordsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindBibRetainedHoldings200ResponseBriefRecordsInner(oclcNumber int64, title string, creator string, date string, language string, edition string, publisher string) *FindBibRetainedHoldings200ResponseBriefRecordsInner {
	this := FindBibRetainedHoldings200ResponseBriefRecordsInner{}
	this.OclcNumber = oclcNumber
	this.Title = title
	this.Creator = creator
	this.Date = date
	this.Language = language
	this.Edition = edition
	this.Publisher = publisher
	return &this
}

// NewFindBibRetainedHoldings200ResponseBriefRecordsInnerWithDefaults instantiates a new FindBibRetainedHoldings200ResponseBriefRecordsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindBibRetainedHoldings200ResponseBriefRecordsInnerWithDefaults() *FindBibRetainedHoldings200ResponseBriefRecordsInner {
	this := FindBibRetainedHoldings200ResponseBriefRecordsInner{}
	return &this
}

// GetOclcNumber returns the OclcNumber field value
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetOclcNumber() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OclcNumber
}

// GetOclcNumberOk returns a tuple with the OclcNumber field value
// and a boolean to check if the value has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetOclcNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OclcNumber, true
}

// SetOclcNumber sets field value
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) SetOclcNumber(v int64) {
	o.OclcNumber = v
}

// GetTitle returns the Title field value
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) SetTitle(v string) {
	o.Title = v
}

// GetCreator returns the Creator field value
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetCreator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value
// and a boolean to check if the value has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetCreatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Creator, true
}

// SetCreator sets field value
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) SetCreator(v string) {
	o.Creator = v
}

// GetDate returns the Date field value
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) SetDate(v string) {
	o.Date = v
}

// GetLanguage returns the Language field value
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) SetLanguage(v string) {
	o.Language = v
}

// GetGeneralFormat returns the GeneralFormat field value if set, zero value otherwise.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetGeneralFormat() string {
	if o == nil || IsNil(o.GeneralFormat) {
		var ret string
		return ret
	}
	return *o.GeneralFormat
}

// GetGeneralFormatOk returns a tuple with the GeneralFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetGeneralFormatOk() (*string, bool) {
	if o == nil || IsNil(o.GeneralFormat) {
		return nil, false
	}
	return o.GeneralFormat, true
}

// HasGeneralFormat returns a boolean if a field has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) HasGeneralFormat() bool {
	if o != nil && !IsNil(o.GeneralFormat) {
		return true
	}

	return false
}

// SetGeneralFormat gets a reference to the given string and assigns it to the GeneralFormat field.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) SetGeneralFormat(v string) {
	o.GeneralFormat = &v
}

// GetSpecificFormat returns the SpecificFormat field value if set, zero value otherwise.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetSpecificFormat() string {
	if o == nil || IsNil(o.SpecificFormat) {
		var ret string
		return ret
	}
	return *o.SpecificFormat
}

// GetSpecificFormatOk returns a tuple with the SpecificFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetSpecificFormatOk() (*string, bool) {
	if o == nil || IsNil(o.SpecificFormat) {
		return nil, false
	}
	return o.SpecificFormat, true
}

// HasSpecificFormat returns a boolean if a field has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) HasSpecificFormat() bool {
	if o != nil && !IsNil(o.SpecificFormat) {
		return true
	}

	return false
}

// SetSpecificFormat gets a reference to the given string and assigns it to the SpecificFormat field.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) SetSpecificFormat(v string) {
	o.SpecificFormat = &v
}

// GetEdition returns the Edition field value
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetEdition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Edition
}

// GetEditionOk returns a tuple with the Edition field value
// and a boolean to check if the value has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetEditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Edition, true
}

// SetEdition sets field value
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) SetEdition(v string) {
	o.Edition = v
}

// GetPublisher returns the Publisher field value
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetPublisher() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Publisher
}

// GetPublisherOk returns a tuple with the Publisher field value
// and a boolean to check if the value has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetPublisherOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Publisher, true
}

// SetPublisher sets field value
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) SetPublisher(v string) {
	o.Publisher = v
}

// GetIsbns returns the Isbns field value if set, zero value otherwise.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetIsbns() []string {
	if o == nil || IsNil(o.Isbns) {
		var ret []string
		return ret
	}
	return o.Isbns
}

// GetIsbnsOk returns a tuple with the Isbns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetIsbnsOk() ([]string, bool) {
	if o == nil || IsNil(o.Isbns) {
		return nil, false
	}
	return o.Isbns, true
}

// HasIsbns returns a boolean if a field has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) HasIsbns() bool {
	if o != nil && !IsNil(o.Isbns) {
		return true
	}

	return false
}

// SetIsbns gets a reference to the given []string and assigns it to the Isbns field.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) SetIsbns(v []string) {
	o.Isbns = v
}

// GetIssns returns the Issns field value if set, zero value otherwise.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetIssns() []string {
	if o == nil || IsNil(o.Issns) {
		var ret []string
		return ret
	}
	return o.Issns
}

// GetIssnsOk returns a tuple with the Issns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetIssnsOk() ([]string, bool) {
	if o == nil || IsNil(o.Issns) {
		return nil, false
	}
	return o.Issns, true
}

// HasIssns returns a boolean if a field has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) HasIssns() bool {
	if o != nil && !IsNil(o.Issns) {
		return true
	}

	return false
}

// SetIssns gets a reference to the given []string and assigns it to the Issns field.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) SetIssns(v []string) {
	o.Issns = v
}

// GetMergedOclcNumbers returns the MergedOclcNumbers field value if set, zero value otherwise.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetMergedOclcNumbers() []string {
	if o == nil || IsNil(o.MergedOclcNumbers) {
		var ret []string
		return ret
	}
	return o.MergedOclcNumbers
}

// GetMergedOclcNumbersOk returns a tuple with the MergedOclcNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetMergedOclcNumbersOk() ([]string, bool) {
	if o == nil || IsNil(o.MergedOclcNumbers) {
		return nil, false
	}
	return o.MergedOclcNumbers, true
}

// HasMergedOclcNumbers returns a boolean if a field has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) HasMergedOclcNumbers() bool {
	if o != nil && !IsNil(o.MergedOclcNumbers) {
		return true
	}

	return false
}

// SetMergedOclcNumbers gets a reference to the given []string and assigns it to the MergedOclcNumbers field.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) SetMergedOclcNumbers(v []string) {
	o.MergedOclcNumbers = v
}

// GetCatalogingInfo returns the CatalogingInfo field value if set, zero value otherwise.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetCatalogingInfo() FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo {
	if o == nil || IsNil(o.CatalogingInfo) {
		var ret FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo
		return ret
	}
	return *o.CatalogingInfo
}

// GetCatalogingInfoOk returns a tuple with the CatalogingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetCatalogingInfoOk() (*FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo, bool) {
	if o == nil || IsNil(o.CatalogingInfo) {
		return nil, false
	}
	return o.CatalogingInfo, true
}

// HasCatalogingInfo returns a boolean if a field has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) HasCatalogingInfo() bool {
	if o != nil && !IsNil(o.CatalogingInfo) {
		return true
	}

	return false
}

// SetCatalogingInfo gets a reference to the given FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo and assigns it to the CatalogingInfo field.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) SetCatalogingInfo(v FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo) {
	o.CatalogingInfo = &v
}

// GetInstitutionHolding returns the InstitutionHolding field value if set, zero value otherwise.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetInstitutionHolding() FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding {
	if o == nil || IsNil(o.InstitutionHolding) {
		var ret FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding
		return ret
	}
	return *o.InstitutionHolding
}

// GetInstitutionHoldingOk returns a tuple with the InstitutionHolding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetInstitutionHoldingOk() (*FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding, bool) {
	if o == nil || IsNil(o.InstitutionHolding) {
		return nil, false
	}
	return o.InstitutionHolding, true
}

// HasInstitutionHolding returns a boolean if a field has been set.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) HasInstitutionHolding() bool {
	if o != nil && !IsNil(o.InstitutionHolding) {
		return true
	}

	return false
}

// SetInstitutionHolding gets a reference to the given FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding and assigns it to the InstitutionHolding field.
func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) SetInstitutionHolding(v FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding) {
	o.InstitutionHolding = &v
}

func (o FindBibRetainedHoldings200ResponseBriefRecordsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FindBibRetainedHoldings200ResponseBriefRecordsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["oclcNumber"] = o.OclcNumber
	toSerialize["title"] = o.Title
	toSerialize["creator"] = o.Creator
	toSerialize["date"] = o.Date
	toSerialize["language"] = o.Language
	if !IsNil(o.GeneralFormat) {
		toSerialize["generalFormat"] = o.GeneralFormat
	}
	if !IsNil(o.SpecificFormat) {
		toSerialize["specificFormat"] = o.SpecificFormat
	}
	toSerialize["edition"] = o.Edition
	toSerialize["publisher"] = o.Publisher
	if !IsNil(o.Isbns) {
		toSerialize["isbns"] = o.Isbns
	}
	if !IsNil(o.Issns) {
		toSerialize["issns"] = o.Issns
	}
	if !IsNil(o.MergedOclcNumbers) {
		toSerialize["mergedOclcNumbers"] = o.MergedOclcNumbers
	}
	if !IsNil(o.CatalogingInfo) {
		toSerialize["catalogingInfo"] = o.CatalogingInfo
	}
	if !IsNil(o.InstitutionHolding) {
		toSerialize["institutionHolding"] = o.InstitutionHolding
	}
	return toSerialize, nil
}

type NullableFindBibRetainedHoldings200ResponseBriefRecordsInner struct {
	value *FindBibRetainedHoldings200ResponseBriefRecordsInner
	isSet bool
}

func (v NullableFindBibRetainedHoldings200ResponseBriefRecordsInner) Get() *FindBibRetainedHoldings200ResponseBriefRecordsInner {
	return v.value
}

func (v *NullableFindBibRetainedHoldings200ResponseBriefRecordsInner) Set(val *FindBibRetainedHoldings200ResponseBriefRecordsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFindBibRetainedHoldings200ResponseBriefRecordsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFindBibRetainedHoldings200ResponseBriefRecordsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindBibRetainedHoldings200ResponseBriefRecordsInner(val *FindBibRetainedHoldings200ResponseBriefRecordsInner) *NullableFindBibRetainedHoldings200ResponseBriefRecordsInner {
	return &NullableFindBibRetainedHoldings200ResponseBriefRecordsInner{value: val, isSet: true}
}

func (v NullableFindBibRetainedHoldings200ResponseBriefRecordsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindBibRetainedHoldings200ResponseBriefRecordsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
