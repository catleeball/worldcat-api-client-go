# FindBibRetainedHoldings400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | A URI reference [RFC3986] that identifies the problem type. | [optional] 
**Title** | Pointer to **string** | A short, human-readable summary of the problem type. | [optional] 
**Instance** | Pointer to **string** | A URI reference that identifies the specific occurrence of the problem. | [optional] 
**Detail** | Pointer to **string** | A human-readable explanation specific to this occurrence of the problem., | [optional] 
**InvalidParams** | Pointer to [**[]FindBibRetainedHoldings400ResponseInvalidParamsInner**](FindBibRetainedHoldings400ResponseInvalidParamsInner.md) | An array of validation errors. | [optional] 

## Methods

### NewFindBibRetainedHoldings400Response

`func NewFindBibRetainedHoldings400Response() *FindBibRetainedHoldings400Response`

NewFindBibRetainedHoldings400Response instantiates a new FindBibRetainedHoldings400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindBibRetainedHoldings400ResponseWithDefaults

`func NewFindBibRetainedHoldings400ResponseWithDefaults() *FindBibRetainedHoldings400Response`

NewFindBibRetainedHoldings400ResponseWithDefaults instantiates a new FindBibRetainedHoldings400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FindBibRetainedHoldings400Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FindBibRetainedHoldings400Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FindBibRetainedHoldings400Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FindBibRetainedHoldings400Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *FindBibRetainedHoldings400Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FindBibRetainedHoldings400Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FindBibRetainedHoldings400Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FindBibRetainedHoldings400Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetInstance

`func (o *FindBibRetainedHoldings400Response) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *FindBibRetainedHoldings400Response) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *FindBibRetainedHoldings400Response) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *FindBibRetainedHoldings400Response) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetDetail

`func (o *FindBibRetainedHoldings400Response) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *FindBibRetainedHoldings400Response) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *FindBibRetainedHoldings400Response) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *FindBibRetainedHoldings400Response) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetInvalidParams

`func (o *FindBibRetainedHoldings400Response) GetInvalidParams() []FindBibRetainedHoldings400ResponseInvalidParamsInner`

GetInvalidParams returns the InvalidParams field if non-nil, zero value otherwise.

### GetInvalidParamsOk

`func (o *FindBibRetainedHoldings400Response) GetInvalidParamsOk() (*[]FindBibRetainedHoldings400ResponseInvalidParamsInner, bool)`

GetInvalidParamsOk returns a tuple with the InvalidParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidParams

`func (o *FindBibRetainedHoldings400Response) SetInvalidParams(v []FindBibRetainedHoldings400ResponseInvalidParamsInner)`

SetInvalidParams sets InvalidParams field to given value.

### HasInvalidParams

`func (o *FindBibRetainedHoldings400Response) HasInvalidParams() bool`

HasInvalidParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


