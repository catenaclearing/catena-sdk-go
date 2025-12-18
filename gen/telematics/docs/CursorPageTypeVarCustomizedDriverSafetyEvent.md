# CursorPageTypeVarCustomizedDriverSafetyEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]DriverSafetyEvent**](DriverSafetyEvent.md) |  | 
**CurrentPage** | Pointer to **NullableString** |  | [optional] 
**CurrentPageBackwards** | Pointer to **NullableString** |  | [optional] 
**PreviousPage** | Pointer to **NullableString** |  | [optional] 
**NextPage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCursorPageTypeVarCustomizedDriverSafetyEvent

`func NewCursorPageTypeVarCustomizedDriverSafetyEvent(items []DriverSafetyEvent, ) *CursorPageTypeVarCustomizedDriverSafetyEvent`

NewCursorPageTypeVarCustomizedDriverSafetyEvent instantiates a new CursorPageTypeVarCustomizedDriverSafetyEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCursorPageTypeVarCustomizedDriverSafetyEventWithDefaults

`func NewCursorPageTypeVarCustomizedDriverSafetyEventWithDefaults() *CursorPageTypeVarCustomizedDriverSafetyEvent`

NewCursorPageTypeVarCustomizedDriverSafetyEventWithDefaults instantiates a new CursorPageTypeVarCustomizedDriverSafetyEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CursorPageTypeVarCustomizedDriverSafetyEvent) GetItems() []DriverSafetyEvent`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CursorPageTypeVarCustomizedDriverSafetyEvent) GetItemsOk() (*[]DriverSafetyEvent, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CursorPageTypeVarCustomizedDriverSafetyEvent) SetItems(v []DriverSafetyEvent)`

SetItems sets Items field to given value.


### GetCurrentPage

`func (o *CursorPageTypeVarCustomizedDriverSafetyEvent) GetCurrentPage() string`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *CursorPageTypeVarCustomizedDriverSafetyEvent) GetCurrentPageOk() (*string, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *CursorPageTypeVarCustomizedDriverSafetyEvent) SetCurrentPage(v string)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *CursorPageTypeVarCustomizedDriverSafetyEvent) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### SetCurrentPageNil

`func (o *CursorPageTypeVarCustomizedDriverSafetyEvent) SetCurrentPageNil(b bool)`

 SetCurrentPageNil sets the value for CurrentPage to be an explicit nil

### UnsetCurrentPage
`func (o *CursorPageTypeVarCustomizedDriverSafetyEvent) UnsetCurrentPage()`

UnsetCurrentPage ensures that no value is present for CurrentPage, not even an explicit nil
### GetCurrentPageBackwards

`func (o *CursorPageTypeVarCustomizedDriverSafetyEvent) GetCurrentPageBackwards() string`

GetCurrentPageBackwards returns the CurrentPageBackwards field if non-nil, zero value otherwise.

### GetCurrentPageBackwardsOk

`func (o *CursorPageTypeVarCustomizedDriverSafetyEvent) GetCurrentPageBackwardsOk() (*string, bool)`

GetCurrentPageBackwardsOk returns a tuple with the CurrentPageBackwards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPageBackwards

`func (o *CursorPageTypeVarCustomizedDriverSafetyEvent) SetCurrentPageBackwards(v string)`

SetCurrentPageBackwards sets CurrentPageBackwards field to given value.

### HasCurrentPageBackwards

`func (o *CursorPageTypeVarCustomizedDriverSafetyEvent) HasCurrentPageBackwards() bool`

HasCurrentPageBackwards returns a boolean if a field has been set.

### SetCurrentPageBackwardsNil

`func (o *CursorPageTypeVarCustomizedDriverSafetyEvent) SetCurrentPageBackwardsNil(b bool)`

 SetCurrentPageBackwardsNil sets the value for CurrentPageBackwards to be an explicit nil

### UnsetCurrentPageBackwards
`func (o *CursorPageTypeVarCustomizedDriverSafetyEvent) UnsetCurrentPageBackwards()`

UnsetCurrentPageBackwards ensures that no value is present for CurrentPageBackwards, not even an explicit nil
### GetPreviousPage

`func (o *CursorPageTypeVarCustomizedDriverSafetyEvent) GetPreviousPage() string`

GetPreviousPage returns the PreviousPage field if non-nil, zero value otherwise.

### GetPreviousPageOk

`func (o *CursorPageTypeVarCustomizedDriverSafetyEvent) GetPreviousPageOk() (*string, bool)`

GetPreviousPageOk returns a tuple with the PreviousPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPage

`func (o *CursorPageTypeVarCustomizedDriverSafetyEvent) SetPreviousPage(v string)`

SetPreviousPage sets PreviousPage field to given value.

### HasPreviousPage

`func (o *CursorPageTypeVarCustomizedDriverSafetyEvent) HasPreviousPage() bool`

HasPreviousPage returns a boolean if a field has been set.

### SetPreviousPageNil

`func (o *CursorPageTypeVarCustomizedDriverSafetyEvent) SetPreviousPageNil(b bool)`

 SetPreviousPageNil sets the value for PreviousPage to be an explicit nil

### UnsetPreviousPage
`func (o *CursorPageTypeVarCustomizedDriverSafetyEvent) UnsetPreviousPage()`

UnsetPreviousPage ensures that no value is present for PreviousPage, not even an explicit nil
### GetNextPage

`func (o *CursorPageTypeVarCustomizedDriverSafetyEvent) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *CursorPageTypeVarCustomizedDriverSafetyEvent) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *CursorPageTypeVarCustomizedDriverSafetyEvent) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *CursorPageTypeVarCustomizedDriverSafetyEvent) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### SetNextPageNil

`func (o *CursorPageTypeVarCustomizedDriverSafetyEvent) SetNextPageNil(b bool)`

 SetNextPageNil sets the value for NextPage to be an explicit nil

### UnsetNextPage
`func (o *CursorPageTypeVarCustomizedDriverSafetyEvent) UnsetNextPage()`

UnsetNextPage ensures that no value is present for NextPage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


