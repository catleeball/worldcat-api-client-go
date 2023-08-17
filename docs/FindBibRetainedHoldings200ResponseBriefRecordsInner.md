# FindBibRetainedHoldings200ResponseBriefRecordsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OclcNumber** | **int64** | the oclc number of a given bibliographic record | 
**Title** | **string** | Linked [v880sa,b] if present or [v245sa,b] | 
**Creator** | **string** | List of creators as single string | 
**Date** | **string** | machine readable date 008 | 
**Language** | **string** | Language of the item [v041sa,j] | 
**GeneralFormat** | Pointer to **string** | General Format Type [Admin/OCLCDef/StdRT] | [optional] 
**SpecificFormat** | Pointer to **string** | Specific Format Type [Admin/OCLCDef/StdRT2] | [optional] 
**Edition** | **string** | Edition Statement [v250sa] | 
**Publisher** | **string** |  | 
**Isbns** | Pointer to **[]string** | International Standard Book Number [v020sa] | [optional] 
**Issns** | Pointer to **[]string** | International Standard Serial Numbers | [optional] 
**MergedOclcNumbers** | Pointer to **[]string** | Merged OCLC numbers [v019sa] | [optional] 
**CatalogingInfo** | Pointer to [**FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo**](FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo.md) |  | [optional] 
**InstitutionHolding** | Pointer to [**FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding**](FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding.md) |  | [optional] 

## Methods

### NewFindBibRetainedHoldings200ResponseBriefRecordsInner

`func NewFindBibRetainedHoldings200ResponseBriefRecordsInner(oclcNumber int64, title string, creator string, date string, language string, edition string, publisher string, ) *FindBibRetainedHoldings200ResponseBriefRecordsInner`

NewFindBibRetainedHoldings200ResponseBriefRecordsInner instantiates a new FindBibRetainedHoldings200ResponseBriefRecordsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindBibRetainedHoldings200ResponseBriefRecordsInnerWithDefaults

`func NewFindBibRetainedHoldings200ResponseBriefRecordsInnerWithDefaults() *FindBibRetainedHoldings200ResponseBriefRecordsInner`

NewFindBibRetainedHoldings200ResponseBriefRecordsInnerWithDefaults instantiates a new FindBibRetainedHoldings200ResponseBriefRecordsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOclcNumber

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetOclcNumber() int64`

GetOclcNumber returns the OclcNumber field if non-nil, zero value otherwise.

### GetOclcNumberOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetOclcNumberOk() (*int64, bool)`

GetOclcNumberOk returns a tuple with the OclcNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOclcNumber

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) SetOclcNumber(v int64)`

SetOclcNumber sets OclcNumber field to given value.


### GetTitle

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetCreator

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) SetCreator(v string)`

SetCreator sets Creator field to given value.


### GetDate

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) SetDate(v string)`

SetDate sets Date field to given value.


### GetLanguage

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetGeneralFormat

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetGeneralFormat() string`

GetGeneralFormat returns the GeneralFormat field if non-nil, zero value otherwise.

### GetGeneralFormatOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetGeneralFormatOk() (*string, bool)`

GetGeneralFormatOk returns a tuple with the GeneralFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralFormat

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) SetGeneralFormat(v string)`

SetGeneralFormat sets GeneralFormat field to given value.

### HasGeneralFormat

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) HasGeneralFormat() bool`

HasGeneralFormat returns a boolean if a field has been set.

### GetSpecificFormat

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetSpecificFormat() string`

GetSpecificFormat returns the SpecificFormat field if non-nil, zero value otherwise.

### GetSpecificFormatOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetSpecificFormatOk() (*string, bool)`

GetSpecificFormatOk returns a tuple with the SpecificFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificFormat

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) SetSpecificFormat(v string)`

SetSpecificFormat sets SpecificFormat field to given value.

### HasSpecificFormat

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) HasSpecificFormat() bool`

HasSpecificFormat returns a boolean if a field has been set.

### GetEdition

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetEdition() string`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetEditionOk() (*string, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) SetEdition(v string)`

SetEdition sets Edition field to given value.


### GetPublisher

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.


### GetIsbns

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetIsbns() []string`

GetIsbns returns the Isbns field if non-nil, zero value otherwise.

### GetIsbnsOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetIsbnsOk() (*[]string, bool)`

GetIsbnsOk returns a tuple with the Isbns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsbns

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) SetIsbns(v []string)`

SetIsbns sets Isbns field to given value.

### HasIsbns

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) HasIsbns() bool`

HasIsbns returns a boolean if a field has been set.

### GetIssns

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetIssns() []string`

GetIssns returns the Issns field if non-nil, zero value otherwise.

### GetIssnsOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetIssnsOk() (*[]string, bool)`

GetIssnsOk returns a tuple with the Issns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssns

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) SetIssns(v []string)`

SetIssns sets Issns field to given value.

### HasIssns

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) HasIssns() bool`

HasIssns returns a boolean if a field has been set.

### GetMergedOclcNumbers

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetMergedOclcNumbers() []string`

GetMergedOclcNumbers returns the MergedOclcNumbers field if non-nil, zero value otherwise.

### GetMergedOclcNumbersOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetMergedOclcNumbersOk() (*[]string, bool)`

GetMergedOclcNumbersOk returns a tuple with the MergedOclcNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedOclcNumbers

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) SetMergedOclcNumbers(v []string)`

SetMergedOclcNumbers sets MergedOclcNumbers field to given value.

### HasMergedOclcNumbers

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) HasMergedOclcNumbers() bool`

HasMergedOclcNumbers returns a boolean if a field has been set.

### GetCatalogingInfo

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetCatalogingInfo() FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo`

GetCatalogingInfo returns the CatalogingInfo field if non-nil, zero value otherwise.

### GetCatalogingInfoOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetCatalogingInfoOk() (*FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo, bool)`

GetCatalogingInfoOk returns a tuple with the CatalogingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogingInfo

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) SetCatalogingInfo(v FindBibRetainedHoldings200ResponseBriefRecordsInnerCatalogingInfo)`

SetCatalogingInfo sets CatalogingInfo field to given value.

### HasCatalogingInfo

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) HasCatalogingInfo() bool`

HasCatalogingInfo returns a boolean if a field has been set.

### GetInstitutionHolding

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetInstitutionHolding() FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding`

GetInstitutionHolding returns the InstitutionHolding field if non-nil, zero value otherwise.

### GetInstitutionHoldingOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) GetInstitutionHoldingOk() (*FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding, bool)`

GetInstitutionHoldingOk returns a tuple with the InstitutionHolding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionHolding

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) SetInstitutionHolding(v FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding)`

SetInstitutionHolding sets InstitutionHolding field to given value.

### HasInstitutionHolding

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInner) HasInstitutionHolding() bool`

HasInstitutionHolding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


