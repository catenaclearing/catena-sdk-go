# CursorPageTypeVarCustomizedWebhookLogRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]WebhookLogRead**](WebhookLogRead.md) |  | 
**CurrentPage** | Pointer to **NullableString** |  | [optional] 
**CurrentPageBackwards** | Pointer to **NullableString** |  | [optional] 
**PreviousPage** | Pointer to **NullableString** |  | [optional] 
**NextPage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCursorPageTypeVarCustomizedWebhookLogRead

`func NewCursorPageTypeVarCustomizedWebhookLogRead(items []WebhookLogRead, ) *CursorPageTypeVarCustomizedWebhookLogRead`

NewCursorPageTypeVarCustomizedWebhookLogRead instantiates a new CursorPageTypeVarCustomizedWebhookLogRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCursorPageTypeVarCustomizedWebhookLogReadWithDefaults

`func NewCursorPageTypeVarCustomizedWebhookLogReadWithDefaults() *CursorPageTypeVarCustomizedWebhookLogRead`

NewCursorPageTypeVarCustomizedWebhookLogReadWithDefaults instantiates a new CursorPageTypeVarCustomizedWebhookLogRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CursorPageTypeVarCustomizedWebhookLogRead) GetItems() []WebhookLogRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CursorPageTypeVarCustomizedWebhookLogRead) GetItemsOk() (*[]WebhookLogRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CursorPageTypeVarCustomizedWebhookLogRead) SetItems(v []WebhookLogRead)`

SetItems sets Items field to given value.


### GetCurrentPage

`func (o *CursorPageTypeVarCustomizedWebhookLogRead) GetCurrentPage() string`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *CursorPageTypeVarCustomizedWebhookLogRead) GetCurrentPageOk() (*string, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *CursorPageTypeVarCustomizedWebhookLogRead) SetCurrentPage(v string)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *CursorPageTypeVarCustomizedWebhookLogRead) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### SetCurrentPageNil

`func (o *CursorPageTypeVarCustomizedWebhookLogRead) SetCurrentPageNil(b bool)`

 SetCurrentPageNil sets the value for CurrentPage to be an explicit nil

### UnsetCurrentPage
`func (o *CursorPageTypeVarCustomizedWebhookLogRead) UnsetCurrentPage()`

UnsetCurrentPage ensures that no value is present for CurrentPage, not even an explicit nil
### GetCurrentPageBackwards

`func (o *CursorPageTypeVarCustomizedWebhookLogRead) GetCurrentPageBackwards() string`

GetCurrentPageBackwards returns the CurrentPageBackwards field if non-nil, zero value otherwise.

### GetCurrentPageBackwardsOk

`func (o *CursorPageTypeVarCustomizedWebhookLogRead) GetCurrentPageBackwardsOk() (*string, bool)`

GetCurrentPageBackwardsOk returns a tuple with the CurrentPageBackwards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPageBackwards

`func (o *CursorPageTypeVarCustomizedWebhookLogRead) SetCurrentPageBackwards(v string)`

SetCurrentPageBackwards sets CurrentPageBackwards field to given value.

### HasCurrentPageBackwards

`func (o *CursorPageTypeVarCustomizedWebhookLogRead) HasCurrentPageBackwards() bool`

HasCurrentPageBackwards returns a boolean if a field has been set.

### SetCurrentPageBackwardsNil

`func (o *CursorPageTypeVarCustomizedWebhookLogRead) SetCurrentPageBackwardsNil(b bool)`

 SetCurrentPageBackwardsNil sets the value for CurrentPageBackwards to be an explicit nil

### UnsetCurrentPageBackwards
`func (o *CursorPageTypeVarCustomizedWebhookLogRead) UnsetCurrentPageBackwards()`

UnsetCurrentPageBackwards ensures that no value is present for CurrentPageBackwards, not even an explicit nil
### GetPreviousPage

`func (o *CursorPageTypeVarCustomizedWebhookLogRead) GetPreviousPage() string`

GetPreviousPage returns the PreviousPage field if non-nil, zero value otherwise.

### GetPreviousPageOk

`func (o *CursorPageTypeVarCustomizedWebhookLogRead) GetPreviousPageOk() (*string, bool)`

GetPreviousPageOk returns a tuple with the PreviousPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPage

`func (o *CursorPageTypeVarCustomizedWebhookLogRead) SetPreviousPage(v string)`

SetPreviousPage sets PreviousPage field to given value.

### HasPreviousPage

`func (o *CursorPageTypeVarCustomizedWebhookLogRead) HasPreviousPage() bool`

HasPreviousPage returns a boolean if a field has been set.

### SetPreviousPageNil

`func (o *CursorPageTypeVarCustomizedWebhookLogRead) SetPreviousPageNil(b bool)`

 SetPreviousPageNil sets the value for PreviousPage to be an explicit nil

### UnsetPreviousPage
`func (o *CursorPageTypeVarCustomizedWebhookLogRead) UnsetPreviousPage()`

UnsetPreviousPage ensures that no value is present for PreviousPage, not even an explicit nil
### GetNextPage

`func (o *CursorPageTypeVarCustomizedWebhookLogRead) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *CursorPageTypeVarCustomizedWebhookLogRead) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *CursorPageTypeVarCustomizedWebhookLogRead) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *CursorPageTypeVarCustomizedWebhookLogRead) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### SetNextPageNil

`func (o *CursorPageTypeVarCustomizedWebhookLogRead) SetNextPageNil(b bool)`

 SetNextPageNil sets the value for NextPage to be an explicit nil

### UnsetNextPage
`func (o *CursorPageTypeVarCustomizedWebhookLogRead) UnsetNextPage()`

UnsetNextPage ensures that no value is present for NextPage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


