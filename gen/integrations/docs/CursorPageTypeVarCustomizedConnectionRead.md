# CursorPageTypeVarCustomizedConnectionRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ConnectionRead**](ConnectionRead.md) |  | 
**CurrentPage** | Pointer to **NullableString** |  | [optional] 
**CurrentPageBackwards** | Pointer to **NullableString** |  | [optional] 
**PreviousPage** | Pointer to **NullableString** |  | [optional] 
**NextPage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCursorPageTypeVarCustomizedConnectionRead

`func NewCursorPageTypeVarCustomizedConnectionRead(items []ConnectionRead, ) *CursorPageTypeVarCustomizedConnectionRead`

NewCursorPageTypeVarCustomizedConnectionRead instantiates a new CursorPageTypeVarCustomizedConnectionRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCursorPageTypeVarCustomizedConnectionReadWithDefaults

`func NewCursorPageTypeVarCustomizedConnectionReadWithDefaults() *CursorPageTypeVarCustomizedConnectionRead`

NewCursorPageTypeVarCustomizedConnectionReadWithDefaults instantiates a new CursorPageTypeVarCustomizedConnectionRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CursorPageTypeVarCustomizedConnectionRead) GetItems() []ConnectionRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CursorPageTypeVarCustomizedConnectionRead) GetItemsOk() (*[]ConnectionRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CursorPageTypeVarCustomizedConnectionRead) SetItems(v []ConnectionRead)`

SetItems sets Items field to given value.


### GetCurrentPage

`func (o *CursorPageTypeVarCustomizedConnectionRead) GetCurrentPage() string`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *CursorPageTypeVarCustomizedConnectionRead) GetCurrentPageOk() (*string, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *CursorPageTypeVarCustomizedConnectionRead) SetCurrentPage(v string)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *CursorPageTypeVarCustomizedConnectionRead) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### SetCurrentPageNil

`func (o *CursorPageTypeVarCustomizedConnectionRead) SetCurrentPageNil(b bool)`

 SetCurrentPageNil sets the value for CurrentPage to be an explicit nil

### UnsetCurrentPage
`func (o *CursorPageTypeVarCustomizedConnectionRead) UnsetCurrentPage()`

UnsetCurrentPage ensures that no value is present for CurrentPage, not even an explicit nil
### GetCurrentPageBackwards

`func (o *CursorPageTypeVarCustomizedConnectionRead) GetCurrentPageBackwards() string`

GetCurrentPageBackwards returns the CurrentPageBackwards field if non-nil, zero value otherwise.

### GetCurrentPageBackwardsOk

`func (o *CursorPageTypeVarCustomizedConnectionRead) GetCurrentPageBackwardsOk() (*string, bool)`

GetCurrentPageBackwardsOk returns a tuple with the CurrentPageBackwards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPageBackwards

`func (o *CursorPageTypeVarCustomizedConnectionRead) SetCurrentPageBackwards(v string)`

SetCurrentPageBackwards sets CurrentPageBackwards field to given value.

### HasCurrentPageBackwards

`func (o *CursorPageTypeVarCustomizedConnectionRead) HasCurrentPageBackwards() bool`

HasCurrentPageBackwards returns a boolean if a field has been set.

### SetCurrentPageBackwardsNil

`func (o *CursorPageTypeVarCustomizedConnectionRead) SetCurrentPageBackwardsNil(b bool)`

 SetCurrentPageBackwardsNil sets the value for CurrentPageBackwards to be an explicit nil

### UnsetCurrentPageBackwards
`func (o *CursorPageTypeVarCustomizedConnectionRead) UnsetCurrentPageBackwards()`

UnsetCurrentPageBackwards ensures that no value is present for CurrentPageBackwards, not even an explicit nil
### GetPreviousPage

`func (o *CursorPageTypeVarCustomizedConnectionRead) GetPreviousPage() string`

GetPreviousPage returns the PreviousPage field if non-nil, zero value otherwise.

### GetPreviousPageOk

`func (o *CursorPageTypeVarCustomizedConnectionRead) GetPreviousPageOk() (*string, bool)`

GetPreviousPageOk returns a tuple with the PreviousPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPage

`func (o *CursorPageTypeVarCustomizedConnectionRead) SetPreviousPage(v string)`

SetPreviousPage sets PreviousPage field to given value.

### HasPreviousPage

`func (o *CursorPageTypeVarCustomizedConnectionRead) HasPreviousPage() bool`

HasPreviousPage returns a boolean if a field has been set.

### SetPreviousPageNil

`func (o *CursorPageTypeVarCustomizedConnectionRead) SetPreviousPageNil(b bool)`

 SetPreviousPageNil sets the value for PreviousPage to be an explicit nil

### UnsetPreviousPage
`func (o *CursorPageTypeVarCustomizedConnectionRead) UnsetPreviousPage()`

UnsetPreviousPage ensures that no value is present for PreviousPage, not even an explicit nil
### GetNextPage

`func (o *CursorPageTypeVarCustomizedConnectionRead) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *CursorPageTypeVarCustomizedConnectionRead) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *CursorPageTypeVarCustomizedConnectionRead) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *CursorPageTypeVarCustomizedConnectionRead) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### SetNextPageNil

`func (o *CursorPageTypeVarCustomizedConnectionRead) SetNextPageNil(b bool)`

 SetNextPageNil sets the value for NextPage to be an explicit nil

### UnsetNextPage
`func (o *CursorPageTypeVarCustomizedConnectionRead) UnsetNextPage()`

UnsetNextPage ensures that no value is present for NextPage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


