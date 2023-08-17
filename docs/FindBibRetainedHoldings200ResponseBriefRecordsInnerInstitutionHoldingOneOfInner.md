# FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** | Country of holding institution [Holdings/Library/Country] | [optional] 
**State** | Pointer to **string** | State of holding institution [Holdings/Library/State] | [optional] 
**OclcSymbol** | Pointer to **string** | Oclc Symbol of holding institution [Holdings/Library/InstSym] | [optional] 
**RegistryId** | Pointer to **int32** | The identifier in the WorldCat Registry for the institution | [optional] 
**InstitutionName** | Pointer to **string** | Name of holding institution [Holdings/Library/InstName] | [optional] 
**AlsoCalled** | Pointer to **string** | Some name the library might also be called | [optional] 
**Address** | Pointer to [**FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInnerAddress**](FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInnerAddress.md) |  | [optional] 
**InstitutionType** | Pointer to **string** | The type of institution | [optional] 
**HasOPACLink** | Pointer to **bool** | Whether or not the library has an OPAC deep links | [optional] 
**Distance** | Pointer to **string** | Distance from location if lat/log was specified | [optional] 
**Self** | Pointer to **string** | URI to find more info about the institution | [optional] 
**IllStatus** | Pointer to **string** | ILL status [Holdings/Library/ILLStatus] | [optional] 
**IllGroup** | Pointer to **string** | ILL group [Holdings/Library/ILLGroup] | [optional] 
**HasSharedPrintCommitment** | Pointer to **string** | institution has shared print commitment [Holdings/LibHasSharedPrint] | [optional] 

## Methods

### NewFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner

`func NewFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner() *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner`

NewFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner instantiates a new FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInnerWithDefaults

`func NewFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInnerWithDefaults() *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner`

NewFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInnerWithDefaults instantiates a new FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetState

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetOclcSymbol

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) GetOclcSymbol() string`

GetOclcSymbol returns the OclcSymbol field if non-nil, zero value otherwise.

### GetOclcSymbolOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) GetOclcSymbolOk() (*string, bool)`

GetOclcSymbolOk returns a tuple with the OclcSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOclcSymbol

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) SetOclcSymbol(v string)`

SetOclcSymbol sets OclcSymbol field to given value.

### HasOclcSymbol

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) HasOclcSymbol() bool`

HasOclcSymbol returns a boolean if a field has been set.

### GetRegistryId

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) GetRegistryId() int32`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) GetRegistryIdOk() (*int32, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) SetRegistryId(v int32)`

SetRegistryId sets RegistryId field to given value.

### HasRegistryId

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) HasRegistryId() bool`

HasRegistryId returns a boolean if a field has been set.

### GetInstitutionName

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) GetInstitutionName() string`

GetInstitutionName returns the InstitutionName field if non-nil, zero value otherwise.

### GetInstitutionNameOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) GetInstitutionNameOk() (*string, bool)`

GetInstitutionNameOk returns a tuple with the InstitutionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionName

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) SetInstitutionName(v string)`

SetInstitutionName sets InstitutionName field to given value.

### HasInstitutionName

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) HasInstitutionName() bool`

HasInstitutionName returns a boolean if a field has been set.

### GetAlsoCalled

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) GetAlsoCalled() string`

GetAlsoCalled returns the AlsoCalled field if non-nil, zero value otherwise.

### GetAlsoCalledOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) GetAlsoCalledOk() (*string, bool)`

GetAlsoCalledOk returns a tuple with the AlsoCalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlsoCalled

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) SetAlsoCalled(v string)`

SetAlsoCalled sets AlsoCalled field to given value.

### HasAlsoCalled

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) HasAlsoCalled() bool`

HasAlsoCalled returns a boolean if a field has been set.

### GetAddress

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) GetAddress() FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInnerAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) GetAddressOk() (*FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInnerAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) SetAddress(v FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInnerAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetInstitutionType

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) GetInstitutionType() string`

GetInstitutionType returns the InstitutionType field if non-nil, zero value otherwise.

### GetInstitutionTypeOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) GetInstitutionTypeOk() (*string, bool)`

GetInstitutionTypeOk returns a tuple with the InstitutionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionType

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) SetInstitutionType(v string)`

SetInstitutionType sets InstitutionType field to given value.

### HasInstitutionType

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) HasInstitutionType() bool`

HasInstitutionType returns a boolean if a field has been set.

### GetHasOPACLink

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) GetHasOPACLink() bool`

GetHasOPACLink returns the HasOPACLink field if non-nil, zero value otherwise.

### GetHasOPACLinkOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) GetHasOPACLinkOk() (*bool, bool)`

GetHasOPACLinkOk returns a tuple with the HasOPACLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOPACLink

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) SetHasOPACLink(v bool)`

SetHasOPACLink sets HasOPACLink field to given value.

### HasHasOPACLink

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) HasHasOPACLink() bool`

HasHasOPACLink returns a boolean if a field has been set.

### GetDistance

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) GetDistance() string`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) GetDistanceOk() (*string, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) SetDistance(v string)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetSelf

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetIllStatus

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) GetIllStatus() string`

GetIllStatus returns the IllStatus field if non-nil, zero value otherwise.

### GetIllStatusOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) GetIllStatusOk() (*string, bool)`

GetIllStatusOk returns a tuple with the IllStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIllStatus

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) SetIllStatus(v string)`

SetIllStatus sets IllStatus field to given value.

### HasIllStatus

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) HasIllStatus() bool`

HasIllStatus returns a boolean if a field has been set.

### GetIllGroup

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) GetIllGroup() string`

GetIllGroup returns the IllGroup field if non-nil, zero value otherwise.

### GetIllGroupOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) GetIllGroupOk() (*string, bool)`

GetIllGroupOk returns a tuple with the IllGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIllGroup

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) SetIllGroup(v string)`

SetIllGroup sets IllGroup field to given value.

### HasIllGroup

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) HasIllGroup() bool`

HasIllGroup returns a boolean if a field has been set.

### GetHasSharedPrintCommitment

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) GetHasSharedPrintCommitment() string`

GetHasSharedPrintCommitment returns the HasSharedPrintCommitment field if non-nil, zero value otherwise.

### GetHasSharedPrintCommitmentOk

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) GetHasSharedPrintCommitmentOk() (*string, bool)`

GetHasSharedPrintCommitmentOk returns a tuple with the HasSharedPrintCommitment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSharedPrintCommitment

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) SetHasSharedPrintCommitment(v string)`

SetHasSharedPrintCommitment sets HasSharedPrintCommitment field to given value.

### HasHasSharedPrintCommitment

`func (o *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) HasHasSharedPrintCommitment() bool`

HasHasSharedPrintCommitment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


