# CursorPageShareAgreementRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ShareAgreementRead**](ShareAgreementRead.md) |  | 
**Total** | **int32** |  | 
**CurrentPage** | Pointer to **NullableString** |  | [optional] 
**CurrentPageBackwards** | Pointer to **NullableString** |  | [optional] 
**PreviousPage** | Pointer to **NullableString** |  | [optional] 
**NextPage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCursorPageShareAgreementRead

`func NewCursorPageShareAgreementRead(items []ShareAgreementRead, total int32, ) *CursorPageShareAgreementRead`

NewCursorPageShareAgreementRead instantiates a new CursorPageShareAgreementRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCursorPageShareAgreementReadWithDefaults

`func NewCursorPageShareAgreementReadWithDefaults() *CursorPageShareAgreementRead`

NewCursorPageShareAgreementReadWithDefaults instantiates a new CursorPageShareAgreementRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CursorPageShareAgreementRead) GetItems() []ShareAgreementRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CursorPageShareAgreementRead) GetItemsOk() (*[]ShareAgreementRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CursorPageShareAgreementRead) SetItems(v []ShareAgreementRead)`

SetItems sets Items field to given value.


### GetTotal

`func (o *CursorPageShareAgreementRead) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CursorPageShareAgreementRead) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CursorPageShareAgreementRead) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetCurrentPage

`func (o *CursorPageShareAgreementRead) GetCurrentPage() string`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *CursorPageShareAgreementRead) GetCurrentPageOk() (*string, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *CursorPageShareAgreementRead) SetCurrentPage(v string)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *CursorPageShareAgreementRead) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### SetCurrentPageNil

`func (o *CursorPageShareAgreementRead) SetCurrentPageNil(b bool)`

 SetCurrentPageNil sets the value for CurrentPage to be an explicit nil

### UnsetCurrentPage
`func (o *CursorPageShareAgreementRead) UnsetCurrentPage()`

UnsetCurrentPage ensures that no value is present for CurrentPage, not even an explicit nil
### GetCurrentPageBackwards

`func (o *CursorPageShareAgreementRead) GetCurrentPageBackwards() string`

GetCurrentPageBackwards returns the CurrentPageBackwards field if non-nil, zero value otherwise.

### GetCurrentPageBackwardsOk

`func (o *CursorPageShareAgreementRead) GetCurrentPageBackwardsOk() (*string, bool)`

GetCurrentPageBackwardsOk returns a tuple with the CurrentPageBackwards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPageBackwards

`func (o *CursorPageShareAgreementRead) SetCurrentPageBackwards(v string)`

SetCurrentPageBackwards sets CurrentPageBackwards field to given value.

### HasCurrentPageBackwards

`func (o *CursorPageShareAgreementRead) HasCurrentPageBackwards() bool`

HasCurrentPageBackwards returns a boolean if a field has been set.

### SetCurrentPageBackwardsNil

`func (o *CursorPageShareAgreementRead) SetCurrentPageBackwardsNil(b bool)`

 SetCurrentPageBackwardsNil sets the value for CurrentPageBackwards to be an explicit nil

### UnsetCurrentPageBackwards
`func (o *CursorPageShareAgreementRead) UnsetCurrentPageBackwards()`

UnsetCurrentPageBackwards ensures that no value is present for CurrentPageBackwards, not even an explicit nil
### GetPreviousPage

`func (o *CursorPageShareAgreementRead) GetPreviousPage() string`

GetPreviousPage returns the PreviousPage field if non-nil, zero value otherwise.

### GetPreviousPageOk

`func (o *CursorPageShareAgreementRead) GetPreviousPageOk() (*string, bool)`

GetPreviousPageOk returns a tuple with the PreviousPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPage

`func (o *CursorPageShareAgreementRead) SetPreviousPage(v string)`

SetPreviousPage sets PreviousPage field to given value.

### HasPreviousPage

`func (o *CursorPageShareAgreementRead) HasPreviousPage() bool`

HasPreviousPage returns a boolean if a field has been set.

### SetPreviousPageNil

`func (o *CursorPageShareAgreementRead) SetPreviousPageNil(b bool)`

 SetPreviousPageNil sets the value for PreviousPage to be an explicit nil

### UnsetPreviousPage
`func (o *CursorPageShareAgreementRead) UnsetPreviousPage()`

UnsetPreviousPage ensures that no value is present for PreviousPage, not even an explicit nil
### GetNextPage

`func (o *CursorPageShareAgreementRead) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *CursorPageShareAgreementRead) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *CursorPageShareAgreementRead) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *CursorPageShareAgreementRead) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### SetNextPageNil

`func (o *CursorPageShareAgreementRead) SetNextPageNil(b bool)`

 SetNextPageNil sets the value for NextPage to be an explicit nil

### UnsetNextPage
`func (o *CursorPageShareAgreementRead) UnsetNextPage()`

UnsetNextPage ensures that no value is present for NextPage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


