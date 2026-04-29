# BaseTrailerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal unique identifier for the telematics event record (Catena PK). | 
**FleetId** | **string** | The Catena fleet this record belongs to (multi-tenant scope). | 
**FleetRef** | **NullableString** |  | 
**TspId** | Pointer to **NullableString** |  | [optional] 
**TspSlug** | Pointer to **NullableString** |  | [optional] 
**SourceName** | [**TspEnum**](TspEnum.md) | The underlying telematics platform that provided this data (e.g., &#x60;samsara&#x60;, &#x60;motive&#x60;, &#x60;hos247&#x60;). Note: Some platforms like &#x60;hos247&#x60; offer white-labeling, so multiple TSPs may share the same source_name — use &#x60;tsp_id&#x60; or &#x60;tsp_slug&#x60; to identify the specific ELD provider. | 
**ConnectionId** | **string** | The specific fleet↔TSP connection through which this record was sourced. | 
**SourceId** | **string** | The ID of the record in the TSP or a deterministic ID/Hash generated from a composite unique key | 
**CreatedAt** | **time.Time** | Immutable: first time this record was ingested into our system. | 
**UpdatedAt** | **time.Time** | Last time we modified this record in our system. | 
**DeletedAt** | Pointer to **NullableTime** |  | [optional] 
**OccurredAt** | **time.Time** | When the underlying event/observation occurred, as reported by the TSP, or the moment it was ingested by us if not available. | 
**ExecutionId** | Pointer to **NullableString** |  | [optional] 
**ScheduleId** | Pointer to **NullableString** |  | [optional] 
**Extras** | Pointer to **map[string]interface{}** |  | [optional] 
**TrailerId** | Pointer to **NullableString** |  | [optional] 
**SourceTrailerId** | **string** | External source system trailer identifier (from the TSP). | 
**PowerOn** | Pointer to **NullableBool** |  | [optional] 
**PowerSource** | Pointer to [**NullableReeferPowerSourceEnum**](ReeferPowerSourceEnum.md) |  | [optional] 
**BatteryVoltage** | Pointer to **NullableFloat32** |  | [optional] 
**BatteryPercent** | Pointer to **NullableFloat32** |  | [optional] 
**BatteryStatus** | Pointer to [**NullableReeferBatteryStatusEnum**](ReeferBatteryStatusEnum.md) |  | [optional] 
**ReeferOperationalStatus** | Pointer to [**NullableReeferOperationalStatusEnum**](ReeferOperationalStatusEnum.md) |  | [optional] 
**OperationMode** | Pointer to [**NullableReeferOperationModeEnum**](ReeferOperationModeEnum.md) |  | [optional] 
**ControlMode** | Pointer to [**NullableReeferControlModeEnum**](ReeferControlModeEnum.md) |  | [optional] 
**DoorOpen** | Pointer to **NullableBool** |  | [optional] 
**CargoStatus** | Pointer to [**NullableReeferCargoStatusEnum**](ReeferCargoStatusEnum.md) |  | [optional] 
**FuelLevel** | Pointer to **NullableFloat32** |  | [optional] 
**ReeferEngineHours** | Pointer to **NullableFloat32** |  | [optional] 
**AmbientTemperature** | Pointer to **NullableFloat32** |  | [optional] 
**Zones** | Pointer to [**[]TrailerZoneStatus**](TrailerZoneStatus.md) |  | [optional] 
**RemoteProbeStatuses** | Pointer to [**[]RemoteProbeStatus**](RemoteProbeStatus.md) |  | [optional] 
**PreTripStatus** | Pointer to [**NullableTrailerPretripStatus**](TrailerPretripStatus.md) |  | [optional] 
**SuctionPressureKpa** | Pointer to **NullableFloat32** |  | [optional] 
**DischargePressureKpa** | Pointer to **NullableFloat32** |  | [optional] 
**HumidityPercent** | Pointer to **NullableFloat32** |  | [optional] 
**ActiveAlarms** | Pointer to [**[]TrailerAlarm**](TrailerAlarm.md) |  | [optional] 
**Location** | Pointer to [**NullablePoint**](Point.md) |  | [optional] 
**H3Index11** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewBaseTrailerStatus

`func NewBaseTrailerStatus(id string, fleetId string, fleetRef NullableString, sourceName TspEnum, connectionId string, sourceId string, createdAt time.Time, updatedAt time.Time, occurredAt time.Time, sourceTrailerId string, ) *BaseTrailerStatus`

NewBaseTrailerStatus instantiates a new BaseTrailerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseTrailerStatusWithDefaults

`func NewBaseTrailerStatusWithDefaults() *BaseTrailerStatus`

NewBaseTrailerStatusWithDefaults instantiates a new BaseTrailerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseTrailerStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseTrailerStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseTrailerStatus) SetId(v string)`

SetId sets Id field to given value.


### GetFleetId

`func (o *BaseTrailerStatus) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *BaseTrailerStatus) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *BaseTrailerStatus) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetFleetRef

`func (o *BaseTrailerStatus) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *BaseTrailerStatus) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *BaseTrailerStatus) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.


### SetFleetRefNil

`func (o *BaseTrailerStatus) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *BaseTrailerStatus) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetTspId

`func (o *BaseTrailerStatus) GetTspId() string`

GetTspId returns the TspId field if non-nil, zero value otherwise.

### GetTspIdOk

`func (o *BaseTrailerStatus) GetTspIdOk() (*string, bool)`

GetTspIdOk returns a tuple with the TspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspId

`func (o *BaseTrailerStatus) SetTspId(v string)`

SetTspId sets TspId field to given value.

### HasTspId

`func (o *BaseTrailerStatus) HasTspId() bool`

HasTspId returns a boolean if a field has been set.

### SetTspIdNil

`func (o *BaseTrailerStatus) SetTspIdNil(b bool)`

 SetTspIdNil sets the value for TspId to be an explicit nil

### UnsetTspId
`func (o *BaseTrailerStatus) UnsetTspId()`

UnsetTspId ensures that no value is present for TspId, not even an explicit nil
### GetTspSlug

`func (o *BaseTrailerStatus) GetTspSlug() string`

GetTspSlug returns the TspSlug field if non-nil, zero value otherwise.

### GetTspSlugOk

`func (o *BaseTrailerStatus) GetTspSlugOk() (*string, bool)`

GetTspSlugOk returns a tuple with the TspSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspSlug

`func (o *BaseTrailerStatus) SetTspSlug(v string)`

SetTspSlug sets TspSlug field to given value.

### HasTspSlug

`func (o *BaseTrailerStatus) HasTspSlug() bool`

HasTspSlug returns a boolean if a field has been set.

### SetTspSlugNil

`func (o *BaseTrailerStatus) SetTspSlugNil(b bool)`

 SetTspSlugNil sets the value for TspSlug to be an explicit nil

### UnsetTspSlug
`func (o *BaseTrailerStatus) UnsetTspSlug()`

UnsetTspSlug ensures that no value is present for TspSlug, not even an explicit nil
### GetSourceName

`func (o *BaseTrailerStatus) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *BaseTrailerStatus) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *BaseTrailerStatus) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetConnectionId

`func (o *BaseTrailerStatus) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *BaseTrailerStatus) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *BaseTrailerStatus) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetSourceId

`func (o *BaseTrailerStatus) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *BaseTrailerStatus) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *BaseTrailerStatus) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetCreatedAt

`func (o *BaseTrailerStatus) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BaseTrailerStatus) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BaseTrailerStatus) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BaseTrailerStatus) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BaseTrailerStatus) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BaseTrailerStatus) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *BaseTrailerStatus) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *BaseTrailerStatus) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *BaseTrailerStatus) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *BaseTrailerStatus) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *BaseTrailerStatus) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *BaseTrailerStatus) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetOccurredAt

`func (o *BaseTrailerStatus) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *BaseTrailerStatus) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *BaseTrailerStatus) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.


### GetExecutionId

`func (o *BaseTrailerStatus) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *BaseTrailerStatus) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *BaseTrailerStatus) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *BaseTrailerStatus) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *BaseTrailerStatus) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *BaseTrailerStatus) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *BaseTrailerStatus) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *BaseTrailerStatus) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *BaseTrailerStatus) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *BaseTrailerStatus) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *BaseTrailerStatus) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *BaseTrailerStatus) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetExtras

`func (o *BaseTrailerStatus) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *BaseTrailerStatus) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *BaseTrailerStatus) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *BaseTrailerStatus) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### SetExtrasNil

`func (o *BaseTrailerStatus) SetExtrasNil(b bool)`

 SetExtrasNil sets the value for Extras to be an explicit nil

### UnsetExtras
`func (o *BaseTrailerStatus) UnsetExtras()`

UnsetExtras ensures that no value is present for Extras, not even an explicit nil
### GetTrailerId

`func (o *BaseTrailerStatus) GetTrailerId() string`

GetTrailerId returns the TrailerId field if non-nil, zero value otherwise.

### GetTrailerIdOk

`func (o *BaseTrailerStatus) GetTrailerIdOk() (*string, bool)`

GetTrailerIdOk returns a tuple with the TrailerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerId

`func (o *BaseTrailerStatus) SetTrailerId(v string)`

SetTrailerId sets TrailerId field to given value.

### HasTrailerId

`func (o *BaseTrailerStatus) HasTrailerId() bool`

HasTrailerId returns a boolean if a field has been set.

### SetTrailerIdNil

`func (o *BaseTrailerStatus) SetTrailerIdNil(b bool)`

 SetTrailerIdNil sets the value for TrailerId to be an explicit nil

### UnsetTrailerId
`func (o *BaseTrailerStatus) UnsetTrailerId()`

UnsetTrailerId ensures that no value is present for TrailerId, not even an explicit nil
### GetSourceTrailerId

`func (o *BaseTrailerStatus) GetSourceTrailerId() string`

GetSourceTrailerId returns the SourceTrailerId field if non-nil, zero value otherwise.

### GetSourceTrailerIdOk

`func (o *BaseTrailerStatus) GetSourceTrailerIdOk() (*string, bool)`

GetSourceTrailerIdOk returns a tuple with the SourceTrailerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTrailerId

`func (o *BaseTrailerStatus) SetSourceTrailerId(v string)`

SetSourceTrailerId sets SourceTrailerId field to given value.


### GetPowerOn

`func (o *BaseTrailerStatus) GetPowerOn() bool`

GetPowerOn returns the PowerOn field if non-nil, zero value otherwise.

### GetPowerOnOk

`func (o *BaseTrailerStatus) GetPowerOnOk() (*bool, bool)`

GetPowerOnOk returns a tuple with the PowerOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOn

`func (o *BaseTrailerStatus) SetPowerOn(v bool)`

SetPowerOn sets PowerOn field to given value.

### HasPowerOn

`func (o *BaseTrailerStatus) HasPowerOn() bool`

HasPowerOn returns a boolean if a field has been set.

### SetPowerOnNil

`func (o *BaseTrailerStatus) SetPowerOnNil(b bool)`

 SetPowerOnNil sets the value for PowerOn to be an explicit nil

### UnsetPowerOn
`func (o *BaseTrailerStatus) UnsetPowerOn()`

UnsetPowerOn ensures that no value is present for PowerOn, not even an explicit nil
### GetPowerSource

`func (o *BaseTrailerStatus) GetPowerSource() ReeferPowerSourceEnum`

GetPowerSource returns the PowerSource field if non-nil, zero value otherwise.

### GetPowerSourceOk

`func (o *BaseTrailerStatus) GetPowerSourceOk() (*ReeferPowerSourceEnum, bool)`

GetPowerSourceOk returns a tuple with the PowerSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSource

`func (o *BaseTrailerStatus) SetPowerSource(v ReeferPowerSourceEnum)`

SetPowerSource sets PowerSource field to given value.

### HasPowerSource

`func (o *BaseTrailerStatus) HasPowerSource() bool`

HasPowerSource returns a boolean if a field has been set.

### SetPowerSourceNil

`func (o *BaseTrailerStatus) SetPowerSourceNil(b bool)`

 SetPowerSourceNil sets the value for PowerSource to be an explicit nil

### UnsetPowerSource
`func (o *BaseTrailerStatus) UnsetPowerSource()`

UnsetPowerSource ensures that no value is present for PowerSource, not even an explicit nil
### GetBatteryVoltage

`func (o *BaseTrailerStatus) GetBatteryVoltage() float32`

GetBatteryVoltage returns the BatteryVoltage field if non-nil, zero value otherwise.

### GetBatteryVoltageOk

`func (o *BaseTrailerStatus) GetBatteryVoltageOk() (*float32, bool)`

GetBatteryVoltageOk returns a tuple with the BatteryVoltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryVoltage

`func (o *BaseTrailerStatus) SetBatteryVoltage(v float32)`

SetBatteryVoltage sets BatteryVoltage field to given value.

### HasBatteryVoltage

`func (o *BaseTrailerStatus) HasBatteryVoltage() bool`

HasBatteryVoltage returns a boolean if a field has been set.

### SetBatteryVoltageNil

`func (o *BaseTrailerStatus) SetBatteryVoltageNil(b bool)`

 SetBatteryVoltageNil sets the value for BatteryVoltage to be an explicit nil

### UnsetBatteryVoltage
`func (o *BaseTrailerStatus) UnsetBatteryVoltage()`

UnsetBatteryVoltage ensures that no value is present for BatteryVoltage, not even an explicit nil
### GetBatteryPercent

`func (o *BaseTrailerStatus) GetBatteryPercent() float32`

GetBatteryPercent returns the BatteryPercent field if non-nil, zero value otherwise.

### GetBatteryPercentOk

`func (o *BaseTrailerStatus) GetBatteryPercentOk() (*float32, bool)`

GetBatteryPercentOk returns a tuple with the BatteryPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryPercent

`func (o *BaseTrailerStatus) SetBatteryPercent(v float32)`

SetBatteryPercent sets BatteryPercent field to given value.

### HasBatteryPercent

`func (o *BaseTrailerStatus) HasBatteryPercent() bool`

HasBatteryPercent returns a boolean if a field has been set.

### SetBatteryPercentNil

`func (o *BaseTrailerStatus) SetBatteryPercentNil(b bool)`

 SetBatteryPercentNil sets the value for BatteryPercent to be an explicit nil

### UnsetBatteryPercent
`func (o *BaseTrailerStatus) UnsetBatteryPercent()`

UnsetBatteryPercent ensures that no value is present for BatteryPercent, not even an explicit nil
### GetBatteryStatus

`func (o *BaseTrailerStatus) GetBatteryStatus() ReeferBatteryStatusEnum`

GetBatteryStatus returns the BatteryStatus field if non-nil, zero value otherwise.

### GetBatteryStatusOk

`func (o *BaseTrailerStatus) GetBatteryStatusOk() (*ReeferBatteryStatusEnum, bool)`

GetBatteryStatusOk returns a tuple with the BatteryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryStatus

`func (o *BaseTrailerStatus) SetBatteryStatus(v ReeferBatteryStatusEnum)`

SetBatteryStatus sets BatteryStatus field to given value.

### HasBatteryStatus

`func (o *BaseTrailerStatus) HasBatteryStatus() bool`

HasBatteryStatus returns a boolean if a field has been set.

### SetBatteryStatusNil

`func (o *BaseTrailerStatus) SetBatteryStatusNil(b bool)`

 SetBatteryStatusNil sets the value for BatteryStatus to be an explicit nil

### UnsetBatteryStatus
`func (o *BaseTrailerStatus) UnsetBatteryStatus()`

UnsetBatteryStatus ensures that no value is present for BatteryStatus, not even an explicit nil
### GetReeferOperationalStatus

`func (o *BaseTrailerStatus) GetReeferOperationalStatus() ReeferOperationalStatusEnum`

GetReeferOperationalStatus returns the ReeferOperationalStatus field if non-nil, zero value otherwise.

### GetReeferOperationalStatusOk

`func (o *BaseTrailerStatus) GetReeferOperationalStatusOk() (*ReeferOperationalStatusEnum, bool)`

GetReeferOperationalStatusOk returns a tuple with the ReeferOperationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReeferOperationalStatus

`func (o *BaseTrailerStatus) SetReeferOperationalStatus(v ReeferOperationalStatusEnum)`

SetReeferOperationalStatus sets ReeferOperationalStatus field to given value.

### HasReeferOperationalStatus

`func (o *BaseTrailerStatus) HasReeferOperationalStatus() bool`

HasReeferOperationalStatus returns a boolean if a field has been set.

### SetReeferOperationalStatusNil

`func (o *BaseTrailerStatus) SetReeferOperationalStatusNil(b bool)`

 SetReeferOperationalStatusNil sets the value for ReeferOperationalStatus to be an explicit nil

### UnsetReeferOperationalStatus
`func (o *BaseTrailerStatus) UnsetReeferOperationalStatus()`

UnsetReeferOperationalStatus ensures that no value is present for ReeferOperationalStatus, not even an explicit nil
### GetOperationMode

`func (o *BaseTrailerStatus) GetOperationMode() ReeferOperationModeEnum`

GetOperationMode returns the OperationMode field if non-nil, zero value otherwise.

### GetOperationModeOk

`func (o *BaseTrailerStatus) GetOperationModeOk() (*ReeferOperationModeEnum, bool)`

GetOperationModeOk returns a tuple with the OperationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationMode

`func (o *BaseTrailerStatus) SetOperationMode(v ReeferOperationModeEnum)`

SetOperationMode sets OperationMode field to given value.

### HasOperationMode

`func (o *BaseTrailerStatus) HasOperationMode() bool`

HasOperationMode returns a boolean if a field has been set.

### SetOperationModeNil

`func (o *BaseTrailerStatus) SetOperationModeNil(b bool)`

 SetOperationModeNil sets the value for OperationMode to be an explicit nil

### UnsetOperationMode
`func (o *BaseTrailerStatus) UnsetOperationMode()`

UnsetOperationMode ensures that no value is present for OperationMode, not even an explicit nil
### GetControlMode

`func (o *BaseTrailerStatus) GetControlMode() ReeferControlModeEnum`

GetControlMode returns the ControlMode field if non-nil, zero value otherwise.

### GetControlModeOk

`func (o *BaseTrailerStatus) GetControlModeOk() (*ReeferControlModeEnum, bool)`

GetControlModeOk returns a tuple with the ControlMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlMode

`func (o *BaseTrailerStatus) SetControlMode(v ReeferControlModeEnum)`

SetControlMode sets ControlMode field to given value.

### HasControlMode

`func (o *BaseTrailerStatus) HasControlMode() bool`

HasControlMode returns a boolean if a field has been set.

### SetControlModeNil

`func (o *BaseTrailerStatus) SetControlModeNil(b bool)`

 SetControlModeNil sets the value for ControlMode to be an explicit nil

### UnsetControlMode
`func (o *BaseTrailerStatus) UnsetControlMode()`

UnsetControlMode ensures that no value is present for ControlMode, not even an explicit nil
### GetDoorOpen

`func (o *BaseTrailerStatus) GetDoorOpen() bool`

GetDoorOpen returns the DoorOpen field if non-nil, zero value otherwise.

### GetDoorOpenOk

`func (o *BaseTrailerStatus) GetDoorOpenOk() (*bool, bool)`

GetDoorOpenOk returns a tuple with the DoorOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoorOpen

`func (o *BaseTrailerStatus) SetDoorOpen(v bool)`

SetDoorOpen sets DoorOpen field to given value.

### HasDoorOpen

`func (o *BaseTrailerStatus) HasDoorOpen() bool`

HasDoorOpen returns a boolean if a field has been set.

### SetDoorOpenNil

`func (o *BaseTrailerStatus) SetDoorOpenNil(b bool)`

 SetDoorOpenNil sets the value for DoorOpen to be an explicit nil

### UnsetDoorOpen
`func (o *BaseTrailerStatus) UnsetDoorOpen()`

UnsetDoorOpen ensures that no value is present for DoorOpen, not even an explicit nil
### GetCargoStatus

`func (o *BaseTrailerStatus) GetCargoStatus() ReeferCargoStatusEnum`

GetCargoStatus returns the CargoStatus field if non-nil, zero value otherwise.

### GetCargoStatusOk

`func (o *BaseTrailerStatus) GetCargoStatusOk() (*ReeferCargoStatusEnum, bool)`

GetCargoStatusOk returns a tuple with the CargoStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCargoStatus

`func (o *BaseTrailerStatus) SetCargoStatus(v ReeferCargoStatusEnum)`

SetCargoStatus sets CargoStatus field to given value.

### HasCargoStatus

`func (o *BaseTrailerStatus) HasCargoStatus() bool`

HasCargoStatus returns a boolean if a field has been set.

### SetCargoStatusNil

`func (o *BaseTrailerStatus) SetCargoStatusNil(b bool)`

 SetCargoStatusNil sets the value for CargoStatus to be an explicit nil

### UnsetCargoStatus
`func (o *BaseTrailerStatus) UnsetCargoStatus()`

UnsetCargoStatus ensures that no value is present for CargoStatus, not even an explicit nil
### GetFuelLevel

`func (o *BaseTrailerStatus) GetFuelLevel() float32`

GetFuelLevel returns the FuelLevel field if non-nil, zero value otherwise.

### GetFuelLevelOk

`func (o *BaseTrailerStatus) GetFuelLevelOk() (*float32, bool)`

GetFuelLevelOk returns a tuple with the FuelLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelLevel

`func (o *BaseTrailerStatus) SetFuelLevel(v float32)`

SetFuelLevel sets FuelLevel field to given value.

### HasFuelLevel

`func (o *BaseTrailerStatus) HasFuelLevel() bool`

HasFuelLevel returns a boolean if a field has been set.

### SetFuelLevelNil

`func (o *BaseTrailerStatus) SetFuelLevelNil(b bool)`

 SetFuelLevelNil sets the value for FuelLevel to be an explicit nil

### UnsetFuelLevel
`func (o *BaseTrailerStatus) UnsetFuelLevel()`

UnsetFuelLevel ensures that no value is present for FuelLevel, not even an explicit nil
### GetReeferEngineHours

`func (o *BaseTrailerStatus) GetReeferEngineHours() float32`

GetReeferEngineHours returns the ReeferEngineHours field if non-nil, zero value otherwise.

### GetReeferEngineHoursOk

`func (o *BaseTrailerStatus) GetReeferEngineHoursOk() (*float32, bool)`

GetReeferEngineHoursOk returns a tuple with the ReeferEngineHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReeferEngineHours

`func (o *BaseTrailerStatus) SetReeferEngineHours(v float32)`

SetReeferEngineHours sets ReeferEngineHours field to given value.

### HasReeferEngineHours

`func (o *BaseTrailerStatus) HasReeferEngineHours() bool`

HasReeferEngineHours returns a boolean if a field has been set.

### SetReeferEngineHoursNil

`func (o *BaseTrailerStatus) SetReeferEngineHoursNil(b bool)`

 SetReeferEngineHoursNil sets the value for ReeferEngineHours to be an explicit nil

### UnsetReeferEngineHours
`func (o *BaseTrailerStatus) UnsetReeferEngineHours()`

UnsetReeferEngineHours ensures that no value is present for ReeferEngineHours, not even an explicit nil
### GetAmbientTemperature

`func (o *BaseTrailerStatus) GetAmbientTemperature() float32`

GetAmbientTemperature returns the AmbientTemperature field if non-nil, zero value otherwise.

### GetAmbientTemperatureOk

`func (o *BaseTrailerStatus) GetAmbientTemperatureOk() (*float32, bool)`

GetAmbientTemperatureOk returns a tuple with the AmbientTemperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmbientTemperature

`func (o *BaseTrailerStatus) SetAmbientTemperature(v float32)`

SetAmbientTemperature sets AmbientTemperature field to given value.

### HasAmbientTemperature

`func (o *BaseTrailerStatus) HasAmbientTemperature() bool`

HasAmbientTemperature returns a boolean if a field has been set.

### SetAmbientTemperatureNil

`func (o *BaseTrailerStatus) SetAmbientTemperatureNil(b bool)`

 SetAmbientTemperatureNil sets the value for AmbientTemperature to be an explicit nil

### UnsetAmbientTemperature
`func (o *BaseTrailerStatus) UnsetAmbientTemperature()`

UnsetAmbientTemperature ensures that no value is present for AmbientTemperature, not even an explicit nil
### GetZones

`func (o *BaseTrailerStatus) GetZones() []TrailerZoneStatus`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *BaseTrailerStatus) GetZonesOk() (*[]TrailerZoneStatus, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *BaseTrailerStatus) SetZones(v []TrailerZoneStatus)`

SetZones sets Zones field to given value.

### HasZones

`func (o *BaseTrailerStatus) HasZones() bool`

HasZones returns a boolean if a field has been set.

### SetZonesNil

`func (o *BaseTrailerStatus) SetZonesNil(b bool)`

 SetZonesNil sets the value for Zones to be an explicit nil

### UnsetZones
`func (o *BaseTrailerStatus) UnsetZones()`

UnsetZones ensures that no value is present for Zones, not even an explicit nil
### GetRemoteProbeStatuses

`func (o *BaseTrailerStatus) GetRemoteProbeStatuses() []RemoteProbeStatus`

GetRemoteProbeStatuses returns the RemoteProbeStatuses field if non-nil, zero value otherwise.

### GetRemoteProbeStatusesOk

`func (o *BaseTrailerStatus) GetRemoteProbeStatusesOk() (*[]RemoteProbeStatus, bool)`

GetRemoteProbeStatusesOk returns a tuple with the RemoteProbeStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteProbeStatuses

`func (o *BaseTrailerStatus) SetRemoteProbeStatuses(v []RemoteProbeStatus)`

SetRemoteProbeStatuses sets RemoteProbeStatuses field to given value.

### HasRemoteProbeStatuses

`func (o *BaseTrailerStatus) HasRemoteProbeStatuses() bool`

HasRemoteProbeStatuses returns a boolean if a field has been set.

### SetRemoteProbeStatusesNil

`func (o *BaseTrailerStatus) SetRemoteProbeStatusesNil(b bool)`

 SetRemoteProbeStatusesNil sets the value for RemoteProbeStatuses to be an explicit nil

### UnsetRemoteProbeStatuses
`func (o *BaseTrailerStatus) UnsetRemoteProbeStatuses()`

UnsetRemoteProbeStatuses ensures that no value is present for RemoteProbeStatuses, not even an explicit nil
### GetPreTripStatus

`func (o *BaseTrailerStatus) GetPreTripStatus() TrailerPretripStatus`

GetPreTripStatus returns the PreTripStatus field if non-nil, zero value otherwise.

### GetPreTripStatusOk

`func (o *BaseTrailerStatus) GetPreTripStatusOk() (*TrailerPretripStatus, bool)`

GetPreTripStatusOk returns a tuple with the PreTripStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreTripStatus

`func (o *BaseTrailerStatus) SetPreTripStatus(v TrailerPretripStatus)`

SetPreTripStatus sets PreTripStatus field to given value.

### HasPreTripStatus

`func (o *BaseTrailerStatus) HasPreTripStatus() bool`

HasPreTripStatus returns a boolean if a field has been set.

### SetPreTripStatusNil

`func (o *BaseTrailerStatus) SetPreTripStatusNil(b bool)`

 SetPreTripStatusNil sets the value for PreTripStatus to be an explicit nil

### UnsetPreTripStatus
`func (o *BaseTrailerStatus) UnsetPreTripStatus()`

UnsetPreTripStatus ensures that no value is present for PreTripStatus, not even an explicit nil
### GetSuctionPressureKpa

`func (o *BaseTrailerStatus) GetSuctionPressureKpa() float32`

GetSuctionPressureKpa returns the SuctionPressureKpa field if non-nil, zero value otherwise.

### GetSuctionPressureKpaOk

`func (o *BaseTrailerStatus) GetSuctionPressureKpaOk() (*float32, bool)`

GetSuctionPressureKpaOk returns a tuple with the SuctionPressureKpa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuctionPressureKpa

`func (o *BaseTrailerStatus) SetSuctionPressureKpa(v float32)`

SetSuctionPressureKpa sets SuctionPressureKpa field to given value.

### HasSuctionPressureKpa

`func (o *BaseTrailerStatus) HasSuctionPressureKpa() bool`

HasSuctionPressureKpa returns a boolean if a field has been set.

### SetSuctionPressureKpaNil

`func (o *BaseTrailerStatus) SetSuctionPressureKpaNil(b bool)`

 SetSuctionPressureKpaNil sets the value for SuctionPressureKpa to be an explicit nil

### UnsetSuctionPressureKpa
`func (o *BaseTrailerStatus) UnsetSuctionPressureKpa()`

UnsetSuctionPressureKpa ensures that no value is present for SuctionPressureKpa, not even an explicit nil
### GetDischargePressureKpa

`func (o *BaseTrailerStatus) GetDischargePressureKpa() float32`

GetDischargePressureKpa returns the DischargePressureKpa field if non-nil, zero value otherwise.

### GetDischargePressureKpaOk

`func (o *BaseTrailerStatus) GetDischargePressureKpaOk() (*float32, bool)`

GetDischargePressureKpaOk returns a tuple with the DischargePressureKpa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDischargePressureKpa

`func (o *BaseTrailerStatus) SetDischargePressureKpa(v float32)`

SetDischargePressureKpa sets DischargePressureKpa field to given value.

### HasDischargePressureKpa

`func (o *BaseTrailerStatus) HasDischargePressureKpa() bool`

HasDischargePressureKpa returns a boolean if a field has been set.

### SetDischargePressureKpaNil

`func (o *BaseTrailerStatus) SetDischargePressureKpaNil(b bool)`

 SetDischargePressureKpaNil sets the value for DischargePressureKpa to be an explicit nil

### UnsetDischargePressureKpa
`func (o *BaseTrailerStatus) UnsetDischargePressureKpa()`

UnsetDischargePressureKpa ensures that no value is present for DischargePressureKpa, not even an explicit nil
### GetHumidityPercent

`func (o *BaseTrailerStatus) GetHumidityPercent() float32`

GetHumidityPercent returns the HumidityPercent field if non-nil, zero value otherwise.

### GetHumidityPercentOk

`func (o *BaseTrailerStatus) GetHumidityPercentOk() (*float32, bool)`

GetHumidityPercentOk returns a tuple with the HumidityPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumidityPercent

`func (o *BaseTrailerStatus) SetHumidityPercent(v float32)`

SetHumidityPercent sets HumidityPercent field to given value.

### HasHumidityPercent

`func (o *BaseTrailerStatus) HasHumidityPercent() bool`

HasHumidityPercent returns a boolean if a field has been set.

### SetHumidityPercentNil

`func (o *BaseTrailerStatus) SetHumidityPercentNil(b bool)`

 SetHumidityPercentNil sets the value for HumidityPercent to be an explicit nil

### UnsetHumidityPercent
`func (o *BaseTrailerStatus) UnsetHumidityPercent()`

UnsetHumidityPercent ensures that no value is present for HumidityPercent, not even an explicit nil
### GetActiveAlarms

`func (o *BaseTrailerStatus) GetActiveAlarms() []TrailerAlarm`

GetActiveAlarms returns the ActiveAlarms field if non-nil, zero value otherwise.

### GetActiveAlarmsOk

`func (o *BaseTrailerStatus) GetActiveAlarmsOk() (*[]TrailerAlarm, bool)`

GetActiveAlarmsOk returns a tuple with the ActiveAlarms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAlarms

`func (o *BaseTrailerStatus) SetActiveAlarms(v []TrailerAlarm)`

SetActiveAlarms sets ActiveAlarms field to given value.

### HasActiveAlarms

`func (o *BaseTrailerStatus) HasActiveAlarms() bool`

HasActiveAlarms returns a boolean if a field has been set.

### SetActiveAlarmsNil

`func (o *BaseTrailerStatus) SetActiveAlarmsNil(b bool)`

 SetActiveAlarmsNil sets the value for ActiveAlarms to be an explicit nil

### UnsetActiveAlarms
`func (o *BaseTrailerStatus) UnsetActiveAlarms()`

UnsetActiveAlarms ensures that no value is present for ActiveAlarms, not even an explicit nil
### GetLocation

`func (o *BaseTrailerStatus) GetLocation() Point`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BaseTrailerStatus) GetLocationOk() (*Point, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BaseTrailerStatus) SetLocation(v Point)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *BaseTrailerStatus) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *BaseTrailerStatus) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *BaseTrailerStatus) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetH3Index11

`func (o *BaseTrailerStatus) GetH3Index11() int32`

GetH3Index11 returns the H3Index11 field if non-nil, zero value otherwise.

### GetH3Index11Ok

`func (o *BaseTrailerStatus) GetH3Index11Ok() (*int32, bool)`

GetH3Index11Ok returns a tuple with the H3Index11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH3Index11

`func (o *BaseTrailerStatus) SetH3Index11(v int32)`

SetH3Index11 sets H3Index11 field to given value.

### HasH3Index11

`func (o *BaseTrailerStatus) HasH3Index11() bool`

HasH3Index11 returns a boolean if a field has been set.

### SetH3Index11Nil

`func (o *BaseTrailerStatus) SetH3Index11Nil(b bool)`

 SetH3Index11Nil sets the value for H3Index11 to be an explicit nil

### UnsetH3Index11
`func (o *BaseTrailerStatus) UnsetH3Index11()`

UnsetH3Index11 ensures that no value is present for H3Index11, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


