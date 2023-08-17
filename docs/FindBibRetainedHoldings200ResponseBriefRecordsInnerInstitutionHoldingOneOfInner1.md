# FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LhrControlNumber** | **string** | LHR Control Number [LHR -&gt; HldDetRec/c001] | 
**LhrDateEntered** | **string** | LHR Date Entered [LHR -&gt; HldDetRec/Admin/CreateDate] | 
**LhrLastUpdated** | **string** | LHR Date Last Used [LHR -&gt; HldDetRec/Admin/ReplacedDate] | 
**OclcNumber** | **string** | OCLC Number [LHR -&gt; HldDetRec/c004] | 
**Format** | **string** | material format [LHR -&gt; HldDetRec/c007] | 
**Location** | Pointer to [**FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location**](FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location.md) |  | [optional] 
**CopyNumber** | Pointer to **string** | copy number [LHR -&gt; HldDetRec/v852st] | [optional] 
**CallNumber** | Pointer to [**FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1CallNumber**](FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1CallNumber.md) |  | [optional] 
**LendingPolicy** | **string** | local lending policy [LHR -&gt; HldDetRec/Admin/Lend] | 
**HasSharedPrintCommitment** | **string** | institution has shared print commitment [Holdings/LibHasSharedPrint || LHR -&gt; HldDetRec/Admin/SPReg] | 
**SharedPrintCommitments** | Pointer to [**[]FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1SharedPrintCommitmentsInner**](FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1SharedPrintCommitmentsInner.md) | shared print details [LHR -&gt; HldDetRec/v583] | [optional] 
**Summary** | Pointer to **string** | summary statement [LHR -&gt; HldSummRec/v966sa] | [optional] 
**HoldingParts** | Pointer to [**[]FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner**](FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner.md) | Textual Holding [LHR -&gt; HldDetRec/v866,v867,v868] Item Information Holding [LHR -&gt; HldDetRec/v876,v877,v878] Enumeration Chronology Holding [LHR -&gt; HldDetRec/v863,v864,v865] | [optional] 

## Methods

### NewFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1

`func NewFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1(lhrControlNumber string, lhrDateEntered string, lhrLastUpdated string, oclcNumber string, format string, lendingPolicy string, hasSharedPrintCommitment string, ) *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1`

NewFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1 instantiates a new FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1WithDefaults

`func NewFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1WithDefaults() *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1`

NewFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1WithDefaults instantiates a new FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLhrControlNumber

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) GetLhrControlNumber() string`

GetLhrControlNumber returns the LhrControlNumber field if non-nil, zero value otherwise.

### GetLhrControlNumberOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) GetLhrControlNumberOk() (*string, bool)`

GetLhrControlNumberOk returns a tuple with the LhrControlNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLhrControlNumber

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) SetLhrControlNumber(v string)`

SetLhrControlNumber sets LhrControlNumber field to given value.


### GetLhrDateEntered

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) GetLhrDateEntered() string`

GetLhrDateEntered returns the LhrDateEntered field if non-nil, zero value otherwise.

### GetLhrDateEnteredOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) GetLhrDateEnteredOk() (*string, bool)`

GetLhrDateEnteredOk returns a tuple with the LhrDateEntered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLhrDateEntered

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) SetLhrDateEntered(v string)`

SetLhrDateEntered sets LhrDateEntered field to given value.


### GetLhrLastUpdated

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) GetLhrLastUpdated() string`

GetLhrLastUpdated returns the LhrLastUpdated field if non-nil, zero value otherwise.

### GetLhrLastUpdatedOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) GetLhrLastUpdatedOk() (*string, bool)`

GetLhrLastUpdatedOk returns a tuple with the LhrLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLhrLastUpdated

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) SetLhrLastUpdated(v string)`

SetLhrLastUpdated sets LhrLastUpdated field to given value.


### GetOclcNumber

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) GetOclcNumber() string`

GetOclcNumber returns the OclcNumber field if non-nil, zero value otherwise.

### GetOclcNumberOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) GetOclcNumberOk() (*string, bool)`

GetOclcNumberOk returns a tuple with the OclcNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOclcNumber

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) SetOclcNumber(v string)`

SetOclcNumber sets OclcNumber field to given value.


### GetFormat

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetLocation

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) GetLocation() FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) GetLocationOk() (*FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) SetLocation(v FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetCopyNumber

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) GetCopyNumber() string`

GetCopyNumber returns the CopyNumber field if non-nil, zero value otherwise.

### GetCopyNumberOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) GetCopyNumberOk() (*string, bool)`

GetCopyNumberOk returns a tuple with the CopyNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyNumber

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) SetCopyNumber(v string)`

SetCopyNumber sets CopyNumber field to given value.

### HasCopyNumber

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) HasCopyNumber() bool`

HasCopyNumber returns a boolean if a field has been set.

### GetCallNumber

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) GetCallNumber() FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1CallNumber`

GetCallNumber returns the CallNumber field if non-nil, zero value otherwise.

### GetCallNumberOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) GetCallNumberOk() (*FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1CallNumber, bool)`

GetCallNumberOk returns a tuple with the CallNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallNumber

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) SetCallNumber(v FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1CallNumber)`

SetCallNumber sets CallNumber field to given value.

### HasCallNumber

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) HasCallNumber() bool`

HasCallNumber returns a boolean if a field has been set.

### GetLendingPolicy

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) GetLendingPolicy() string`

GetLendingPolicy returns the LendingPolicy field if non-nil, zero value otherwise.

### GetLendingPolicyOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) GetLendingPolicyOk() (*string, bool)`

GetLendingPolicyOk returns a tuple with the LendingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLendingPolicy

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) SetLendingPolicy(v string)`

SetLendingPolicy sets LendingPolicy field to given value.


### GetHasSharedPrintCommitment

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) GetHasSharedPrintCommitment() string`

GetHasSharedPrintCommitment returns the HasSharedPrintCommitment field if non-nil, zero value otherwise.

### GetHasSharedPrintCommitmentOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) GetHasSharedPrintCommitmentOk() (*string, bool)`

GetHasSharedPrintCommitmentOk returns a tuple with the HasSharedPrintCommitment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSharedPrintCommitment

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) SetHasSharedPrintCommitment(v string)`

SetHasSharedPrintCommitment sets HasSharedPrintCommitment field to given value.


### GetSharedPrintCommitments

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) GetSharedPrintCommitments() []FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1SharedPrintCommitmentsInner`

GetSharedPrintCommitments returns the SharedPrintCommitments field if non-nil, zero value otherwise.

### GetSharedPrintCommitmentsOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) GetSharedPrintCommitmentsOk() (*[]FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1SharedPrintCommitmentsInner, bool)`

GetSharedPrintCommitmentsOk returns a tuple with the SharedPrintCommitments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedPrintCommitments

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) SetSharedPrintCommitments(v []FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1SharedPrintCommitmentsInner)`

SetSharedPrintCommitments sets SharedPrintCommitments field to given value.

### HasSharedPrintCommitments

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) HasSharedPrintCommitments() bool`

HasSharedPrintCommitments returns a boolean if a field has been set.

### GetSummary

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetHoldingParts

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) GetHoldingParts() []FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner`

GetHoldingParts returns the HoldingParts field if non-nil, zero value otherwise.

### GetHoldingPartsOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) GetHoldingPartsOk() (*[]FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner, bool)`

GetHoldingPartsOk returns a tuple with the HoldingParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldingParts

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) SetHoldingParts(v []FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1HoldingPartsInner)`

SetHoldingParts sets HoldingParts field to given value.

### HasHoldingParts

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) HasHoldingParts() bool`

HasHoldingParts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


