# SearchBibs200ResponseBibRecordsInnerIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OclcNumber** | Pointer to **int64** | OCLC Number [Admin/BaseOCLCNo || Admin/OCLCNo || C001] | [optional] 
**Lccn** | Pointer to **string** | Library of Congress Control Number [v010sa] | [optional] 
**Isbns** | Pointer to **[]string** | International Standard Book Number [v020sa] | [optional] 
**Issns** | Pointer to **[]string** | International Standard Serial Numbers | [optional] 
**ExternalIdentifiers** | Pointer to [**[]SearchBibs200ResponseBibRecordsInnerIdentifierExternalIdentifiersInner**](SearchBibs200ResponseBibRecordsInnerIdentifierExternalIdentifiersInner.md) | External System Control Number [v029sa,b,c,t] | [optional] 
**OtherStandardIdentifiers** | Pointer to [**[]SearchBibs200ResponseBibRecordsInnerIdentifierOtherStandardIdentifiersInner**](SearchBibs200ResponseBibRecordsInnerIdentifierOtherStandardIdentifiersInner.md) | Other Standard Identifier [v024sa] | [optional] 
**Dois** | Pointer to **[]string** | Digital Object Identifier [v901sb v024sa] | [optional] 
**MergedOclcNumbers** | Pointer to **[]string** | Merged OCLC numbers [v019sa] | [optional] 
**GpoNumber** | Pointer to **[]string** | GPO Item number [v074sa] | [optional] 

## Methods

### NewSearchBibs200ResponseBibRecordsInnerIdentifier

`func NewSearchBibs200ResponseBibRecordsInnerIdentifier() *SearchBibs200ResponseBibRecordsInnerIdentifier`

NewSearchBibs200ResponseBibRecordsInnerIdentifier instantiates a new SearchBibs200ResponseBibRecordsInnerIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchBibs200ResponseBibRecordsInnerIdentifierWithDefaults

`func NewSearchBibs200ResponseBibRecordsInnerIdentifierWithDefaults() *SearchBibs200ResponseBibRecordsInnerIdentifier`

NewSearchBibs200ResponseBibRecordsInnerIdentifierWithDefaults instantiates a new SearchBibs200ResponseBibRecordsInnerIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOclcNumber

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) GetOclcNumber() int64`

GetOclcNumber returns the OclcNumber field if non-nil, zero value otherwise.

### GetOclcNumberOk

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) GetOclcNumberOk() (*int64, bool)`

GetOclcNumberOk returns a tuple with the OclcNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOclcNumber

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) SetOclcNumber(v int64)`

SetOclcNumber sets OclcNumber field to given value.

### HasOclcNumber

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) HasOclcNumber() bool`

HasOclcNumber returns a boolean if a field has been set.

### GetLccn

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) GetLccn() string`

GetLccn returns the Lccn field if non-nil, zero value otherwise.

### GetLccnOk

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) GetLccnOk() (*string, bool)`

GetLccnOk returns a tuple with the Lccn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLccn

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) SetLccn(v string)`

SetLccn sets Lccn field to given value.

### HasLccn

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) HasLccn() bool`

HasLccn returns a boolean if a field has been set.

### GetIsbns

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) GetIsbns() []string`

GetIsbns returns the Isbns field if non-nil, zero value otherwise.

### GetIsbnsOk

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) GetIsbnsOk() (*[]string, bool)`

GetIsbnsOk returns a tuple with the Isbns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsbns

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) SetIsbns(v []string)`

SetIsbns sets Isbns field to given value.

### HasIsbns

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) HasIsbns() bool`

HasIsbns returns a boolean if a field has been set.

### GetIssns

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) GetIssns() []string`

GetIssns returns the Issns field if non-nil, zero value otherwise.

### GetIssnsOk

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) GetIssnsOk() (*[]string, bool)`

GetIssnsOk returns a tuple with the Issns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssns

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) SetIssns(v []string)`

SetIssns sets Issns field to given value.

### HasIssns

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) HasIssns() bool`

HasIssns returns a boolean if a field has been set.

### GetExternalIdentifiers

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) GetExternalIdentifiers() []SearchBibs200ResponseBibRecordsInnerIdentifierExternalIdentifiersInner`

GetExternalIdentifiers returns the ExternalIdentifiers field if non-nil, zero value otherwise.

### GetExternalIdentifiersOk

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) GetExternalIdentifiersOk() (*[]SearchBibs200ResponseBibRecordsInnerIdentifierExternalIdentifiersInner, bool)`

GetExternalIdentifiersOk returns a tuple with the ExternalIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifiers

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) SetExternalIdentifiers(v []SearchBibs200ResponseBibRecordsInnerIdentifierExternalIdentifiersInner)`

SetExternalIdentifiers sets ExternalIdentifiers field to given value.

### HasExternalIdentifiers

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) HasExternalIdentifiers() bool`

HasExternalIdentifiers returns a boolean if a field has been set.

### GetOtherStandardIdentifiers

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) GetOtherStandardIdentifiers() []SearchBibs200ResponseBibRecordsInnerIdentifierOtherStandardIdentifiersInner`

GetOtherStandardIdentifiers returns the OtherStandardIdentifiers field if non-nil, zero value otherwise.

### GetOtherStandardIdentifiersOk

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) GetOtherStandardIdentifiersOk() (*[]SearchBibs200ResponseBibRecordsInnerIdentifierOtherStandardIdentifiersInner, bool)`

GetOtherStandardIdentifiersOk returns a tuple with the OtherStandardIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherStandardIdentifiers

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) SetOtherStandardIdentifiers(v []SearchBibs200ResponseBibRecordsInnerIdentifierOtherStandardIdentifiersInner)`

SetOtherStandardIdentifiers sets OtherStandardIdentifiers field to given value.

### HasOtherStandardIdentifiers

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) HasOtherStandardIdentifiers() bool`

HasOtherStandardIdentifiers returns a boolean if a field has been set.

### GetDois

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) GetDois() []string`

GetDois returns the Dois field if non-nil, zero value otherwise.

### GetDoisOk

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) GetDoisOk() (*[]string, bool)`

GetDoisOk returns a tuple with the Dois field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDois

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) SetDois(v []string)`

SetDois sets Dois field to given value.

### HasDois

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) HasDois() bool`

HasDois returns a boolean if a field has been set.

### GetMergedOclcNumbers

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) GetMergedOclcNumbers() []string`

GetMergedOclcNumbers returns the MergedOclcNumbers field if non-nil, zero value otherwise.

### GetMergedOclcNumbersOk

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) GetMergedOclcNumbersOk() (*[]string, bool)`

GetMergedOclcNumbersOk returns a tuple with the MergedOclcNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedOclcNumbers

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) SetMergedOclcNumbers(v []string)`

SetMergedOclcNumbers sets MergedOclcNumbers field to given value.

### HasMergedOclcNumbers

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) HasMergedOclcNumbers() bool`

HasMergedOclcNumbers returns a boolean if a field has been set.

### GetGpoNumber

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) GetGpoNumber() []string`

GetGpoNumber returns the GpoNumber field if non-nil, zero value otherwise.

### GetGpoNumberOk

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) GetGpoNumberOk() (*[]string, bool)`

GetGpoNumberOk returns a tuple with the GpoNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpoNumber

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) SetGpoNumber(v []string)`

SetGpoNumber sets GpoNumber field to given value.

### HasGpoNumber

`func (o *SearchBibs200ResponseBibRecordsInnerIdentifier) HasGpoNumber() bool`

HasGpoNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


