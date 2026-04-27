# TrailerPretripStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to **NullableString** |  | [optional] 
**InitiatedAt** | Pointer to **NullableTime** |  | [optional] 
**CompletedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewTrailerPretripStatus

`func NewTrailerPretripStatus() *TrailerPretripStatus`

NewTrailerPretripStatus instantiates a new TrailerPretripStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrailerPretripStatusWithDefaults

`func NewTrailerPretripStatusWithDefaults() *TrailerPretripStatus`

NewTrailerPretripStatusWithDefaults instantiates a new TrailerPretripStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *TrailerPretripStatus) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TrailerPretripStatus) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TrailerPretripStatus) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *TrailerPretripStatus) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *TrailerPretripStatus) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *TrailerPretripStatus) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetInitiatedAt

`func (o *TrailerPretripStatus) GetInitiatedAt() time.Time`

GetInitiatedAt returns the InitiatedAt field if non-nil, zero value otherwise.

### GetInitiatedAtOk

`func (o *TrailerPretripStatus) GetInitiatedAtOk() (*time.Time, bool)`

GetInitiatedAtOk returns a tuple with the InitiatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatedAt

`func (o *TrailerPretripStatus) SetInitiatedAt(v time.Time)`

SetInitiatedAt sets InitiatedAt field to given value.

### HasInitiatedAt

`func (o *TrailerPretripStatus) HasInitiatedAt() bool`

HasInitiatedAt returns a boolean if a field has been set.

### SetInitiatedAtNil

`func (o *TrailerPretripStatus) SetInitiatedAtNil(b bool)`

 SetInitiatedAtNil sets the value for InitiatedAt to be an explicit nil

### UnsetInitiatedAt
`func (o *TrailerPretripStatus) UnsetInitiatedAt()`

UnsetInitiatedAt ensures that no value is present for InitiatedAt, not even an explicit nil
### GetCompletedAt

`func (o *TrailerPretripStatus) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *TrailerPretripStatus) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *TrailerPretripStatus) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *TrailerPretripStatus) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *TrailerPretripStatus) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *TrailerPretripStatus) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


