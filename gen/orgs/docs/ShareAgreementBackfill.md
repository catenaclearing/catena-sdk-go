# ShareAgreementBackfill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | [**ResourceEnum**](ResourceEnum.md) | The resource to add to share agreements that don&#39;t have it | 
**ShareLevel** | Pointer to [**ShareLevelEnum**](ShareLevelEnum.md) | The permission level to assign to the resource | [optional] 
**FleetIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewShareAgreementBackfill

`func NewShareAgreementBackfill(resource ResourceEnum, ) *ShareAgreementBackfill`

NewShareAgreementBackfill instantiates a new ShareAgreementBackfill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareAgreementBackfillWithDefaults

`func NewShareAgreementBackfillWithDefaults() *ShareAgreementBackfill`

NewShareAgreementBackfillWithDefaults instantiates a new ShareAgreementBackfill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *ShareAgreementBackfill) GetResource() ResourceEnum`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ShareAgreementBackfill) GetResourceOk() (*ResourceEnum, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ShareAgreementBackfill) SetResource(v ResourceEnum)`

SetResource sets Resource field to given value.


### GetShareLevel

`func (o *ShareAgreementBackfill) GetShareLevel() ShareLevelEnum`

GetShareLevel returns the ShareLevel field if non-nil, zero value otherwise.

### GetShareLevelOk

`func (o *ShareAgreementBackfill) GetShareLevelOk() (*ShareLevelEnum, bool)`

GetShareLevelOk returns a tuple with the ShareLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareLevel

`func (o *ShareAgreementBackfill) SetShareLevel(v ShareLevelEnum)`

SetShareLevel sets ShareLevel field to given value.

### HasShareLevel

`func (o *ShareAgreementBackfill) HasShareLevel() bool`

HasShareLevel returns a boolean if a field has been set.

### GetFleetIds

`func (o *ShareAgreementBackfill) GetFleetIds() []string`

GetFleetIds returns the FleetIds field if non-nil, zero value otherwise.

### GetFleetIdsOk

`func (o *ShareAgreementBackfill) GetFleetIdsOk() (*[]string, bool)`

GetFleetIdsOk returns a tuple with the FleetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetIds

`func (o *ShareAgreementBackfill) SetFleetIds(v []string)`

SetFleetIds sets FleetIds field to given value.

### HasFleetIds

`func (o *ShareAgreementBackfill) HasFleetIds() bool`

HasFleetIds returns a boolean if a field has been set.

### SetFleetIdsNil

`func (o *ShareAgreementBackfill) SetFleetIdsNil(b bool)`

 SetFleetIdsNil sets the value for FleetIds to be an explicit nil

### UnsetFleetIds
`func (o *ShareAgreementBackfill) UnsetFleetIds()`

UnsetFleetIds ensures that no value is present for FleetIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


