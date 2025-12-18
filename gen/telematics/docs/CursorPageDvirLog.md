# CursorPageDvirLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]DvirLog**](DvirLog.md) |  | 
**Total** | **int32** |  | 
**CurrentPage** | Pointer to **NullableString** |  | [optional] 
**CurrentPageBackwards** | Pointer to **NullableString** |  | [optional] 
**PreviousPage** | Pointer to **NullableString** |  | [optional] 
**NextPage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCursorPageDvirLog

`func NewCursorPageDvirLog(items []DvirLog, total int32, ) *CursorPageDvirLog`

NewCursorPageDvirLog instantiates a new CursorPageDvirLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCursorPageDvirLogWithDefaults

`func NewCursorPageDvirLogWithDefaults() *CursorPageDvirLog`

NewCursorPageDvirLogWithDefaults instantiates a new CursorPageDvirLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CursorPageDvirLog) GetItems() []DvirLog`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CursorPageDvirLog) GetItemsOk() (*[]DvirLog, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CursorPageDvirLog) SetItems(v []DvirLog)`

SetItems sets Items field to given value.


### GetTotal

`func (o *CursorPageDvirLog) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CursorPageDvirLog) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CursorPageDvirLog) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetCurrentPage

`func (o *CursorPageDvirLog) GetCurrentPage() string`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *CursorPageDvirLog) GetCurrentPageOk() (*string, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *CursorPageDvirLog) SetCurrentPage(v string)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *CursorPageDvirLog) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### SetCurrentPageNil

`func (o *CursorPageDvirLog) SetCurrentPageNil(b bool)`

 SetCurrentPageNil sets the value for CurrentPage to be an explicit nil

### UnsetCurrentPage
`func (o *CursorPageDvirLog) UnsetCurrentPage()`

UnsetCurrentPage ensures that no value is present for CurrentPage, not even an explicit nil
### GetCurrentPageBackwards

`func (o *CursorPageDvirLog) GetCurrentPageBackwards() string`

GetCurrentPageBackwards returns the CurrentPageBackwards field if non-nil, zero value otherwise.

### GetCurrentPageBackwardsOk

`func (o *CursorPageDvirLog) GetCurrentPageBackwardsOk() (*string, bool)`

GetCurrentPageBackwardsOk returns a tuple with the CurrentPageBackwards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPageBackwards

`func (o *CursorPageDvirLog) SetCurrentPageBackwards(v string)`

SetCurrentPageBackwards sets CurrentPageBackwards field to given value.

### HasCurrentPageBackwards

`func (o *CursorPageDvirLog) HasCurrentPageBackwards() bool`

HasCurrentPageBackwards returns a boolean if a field has been set.

### SetCurrentPageBackwardsNil

`func (o *CursorPageDvirLog) SetCurrentPageBackwardsNil(b bool)`

 SetCurrentPageBackwardsNil sets the value for CurrentPageBackwards to be an explicit nil

### UnsetCurrentPageBackwards
`func (o *CursorPageDvirLog) UnsetCurrentPageBackwards()`

UnsetCurrentPageBackwards ensures that no value is present for CurrentPageBackwards, not even an explicit nil
### GetPreviousPage

`func (o *CursorPageDvirLog) GetPreviousPage() string`

GetPreviousPage returns the PreviousPage field if non-nil, zero value otherwise.

### GetPreviousPageOk

`func (o *CursorPageDvirLog) GetPreviousPageOk() (*string, bool)`

GetPreviousPageOk returns a tuple with the PreviousPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPage

`func (o *CursorPageDvirLog) SetPreviousPage(v string)`

SetPreviousPage sets PreviousPage field to given value.

### HasPreviousPage

`func (o *CursorPageDvirLog) HasPreviousPage() bool`

HasPreviousPage returns a boolean if a field has been set.

### SetPreviousPageNil

`func (o *CursorPageDvirLog) SetPreviousPageNil(b bool)`

 SetPreviousPageNil sets the value for PreviousPage to be an explicit nil

### UnsetPreviousPage
`func (o *CursorPageDvirLog) UnsetPreviousPage()`

UnsetPreviousPage ensures that no value is present for PreviousPage, not even an explicit nil
### GetNextPage

`func (o *CursorPageDvirLog) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *CursorPageDvirLog) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *CursorPageDvirLog) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *CursorPageDvirLog) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### SetNextPageNil

`func (o *CursorPageDvirLog) SetNextPageNil(b bool)`

 SetNextPageNil sets the value for NextPage to be an explicit nil

### UnsetNextPage
`func (o *CursorPageDvirLog) UnsetNextPage()`

UnsetNextPage ensures that no value is present for NextPage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


