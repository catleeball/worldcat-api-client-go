# SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to [**SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner**](SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner.md) |  | [optional] 
**SecondName** | Pointer to [**SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner**](SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner.md) |  | [optional] 
**NonPersonName** | Pointer to [**SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner**](SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**CreatorNotes** | Pointer to **[]string** | subfield d | [optional] 
**Relators** | Pointer to [**[]SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner**](SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner.md) | subfield e | [optional] 
**Affiliation** | Pointer to **string** | subfield u | [optional] 
**IsPrimary** | Pointer to **bool** | whether or not this contributor is the primary one (the 1xx) | [optional] 
**RelatedItem** | Pointer to [**SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem**](SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem.md) |  | [optional] 

## Methods

### NewSearchBibs200ResponseBibRecordsInnerContributorCreatorsInner

`func NewSearchBibs200ResponseBibRecordsInnerContributorCreatorsInner() *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner`

NewSearchBibs200ResponseBibRecordsInnerContributorCreatorsInner instantiates a new SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerWithDefaults

`func NewSearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerWithDefaults() *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner`

NewSearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerWithDefaults instantiates a new SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) GetFirstName() SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) GetFirstNameOk() (*SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) SetFirstName(v SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetSecondName

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) GetSecondName() SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner`

GetSecondName returns the SecondName field if non-nil, zero value otherwise.

### GetSecondNameOk

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) GetSecondNameOk() (*SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner, bool)`

GetSecondNameOk returns a tuple with the SecondName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondName

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) SetSecondName(v SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner)`

SetSecondName sets SecondName field to given value.

### HasSecondName

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) HasSecondName() bool`

HasSecondName returns a boolean if a field has been set.

### GetNonPersonName

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) GetNonPersonName() SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner`

GetNonPersonName returns the NonPersonName field if non-nil, zero value otherwise.

### GetNonPersonNameOk

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) GetNonPersonNameOk() (*SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner, bool)`

GetNonPersonNameOk returns a tuple with the NonPersonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonPersonName

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) SetNonPersonName(v SearchBibs200ResponseBibRecordsInnerTitleMainTitlesInner)`

SetNonPersonName sets NonPersonName field to given value.

### HasNonPersonName

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) HasNonPersonName() bool`

HasNonPersonName returns a boolean if a field has been set.

### GetType

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatorNotes

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) GetCreatorNotes() []string`

GetCreatorNotes returns the CreatorNotes field if non-nil, zero value otherwise.

### GetCreatorNotesOk

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) GetCreatorNotesOk() (*[]string, bool)`

GetCreatorNotesOk returns a tuple with the CreatorNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorNotes

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) SetCreatorNotes(v []string)`

SetCreatorNotes sets CreatorNotes field to given value.

### HasCreatorNotes

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) HasCreatorNotes() bool`

HasCreatorNotes returns a boolean if a field has been set.

### GetRelators

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) GetRelators() []SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner`

GetRelators returns the Relators field if non-nil, zero value otherwise.

### GetRelatorsOk

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) GetRelatorsOk() (*[]SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner, bool)`

GetRelatorsOk returns a tuple with the Relators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelators

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) SetRelators(v []SearchBibs200ResponseBibRecordsInnerTitleAdditionalTitlesInnerRelatorsInner)`

SetRelators sets Relators field to given value.

### HasRelators

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) HasRelators() bool`

HasRelators returns a boolean if a field has been set.

### GetAffiliation

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) GetAffiliation() string`

GetAffiliation returns the Affiliation field if non-nil, zero value otherwise.

### GetAffiliationOk

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) GetAffiliationOk() (*string, bool)`

GetAffiliationOk returns a tuple with the Affiliation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliation

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) SetAffiliation(v string)`

SetAffiliation sets Affiliation field to given value.

### HasAffiliation

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) HasAffiliation() bool`

HasAffiliation returns a boolean if a field has been set.

### GetIsPrimary

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetRelatedItem

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) GetRelatedItem() SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem`

GetRelatedItem returns the RelatedItem field if non-nil, zero value otherwise.

### GetRelatedItemOk

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) GetRelatedItemOk() (*SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem, bool)`

GetRelatedItemOk returns a tuple with the RelatedItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedItem

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) SetRelatedItem(v SearchBibs200ResponseBibRecordsInnerContributorCreatorsInnerRelatedItem)`

SetRelatedItem sets RelatedItem field to given value.

### HasRelatedItem

`func (o *SearchBibs200ResponseBibRecordsInnerContributorCreatorsInner) HasRelatedItem() bool`

HasRelatedItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


