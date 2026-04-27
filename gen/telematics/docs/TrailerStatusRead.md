# TrailerStatusRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FleetId** | **NullableString** |  | 
**FleetRef** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** | Unique identifier of the record at Catena Telematics. | 
**CreatedAt** | **time.Time** | Immutable: The datetime the record was ingested into Catena Telematics. | 
**UpdatedAt** | **time.Time** | The dateime the record was last modified in Catena Telematics. | 
**DeletedAt** | Pointer to **NullableTime** |  | [optional] 
**ConnectionId** | **string** | Unique identifier of the connection at Catena Telematics through which this record was ingested. A connection represents a Fleet/TSP pairing. | 
**TspId** | Pointer to **NullableString** |  | [optional] 
**TspSlug** | Pointer to **NullableString** |  | [optional] 
**SourceName** | [**TspEnum**](TspEnum.md) | The underlying telematics platform that provided this data (e.g., &#x60;samsara&#x60;, &#x60;motive&#x60;, &#x60;hos247&#x60;). Note: Some platforms like &#x60;hos247&#x60; offer white-labeling, so multiple TSPs may share the same source_name — use &#x60;tsp_id&#x60; or &#x60;tsp_slug&#x60; to identify the specific ELD provider. | 
**SourceData** | Pointer to **map[string]interface{}** | Raw source payload as ingested from the TSP. **Note: use it for audit/debugging.** | [optional] 
**SourceId** | **string** | Unique identifier of the record in the TSP. **Note: we generate a unique composite key based on available fields if the TSP does not provide an unique ID.** | 
**SourceDataHash** | **string** | SHA-256 hash of the source data payload. **Note: we use it internally for idempotence and deduplication.** | 
**OccurredAt** | Pointer to **NullableTime** |  | [optional] 
**ExecutionId** | Pointer to **NullableString** |  | [optional] 
**ScheduleId** | Pointer to **NullableString** |  | [optional] 
**Extras** | Pointer to **map[string]interface{}** |  | [optional] 
**TrailerId** | Pointer to **NullableString** |  | [optional] 
**SourceTrailerId** | Pointer to **NullableString** |  | [optional] 
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
**Location** | Pointer to [**NullableLocation5**](Location5.md) |  | [optional] 
**H3Index11** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewTrailerStatusRead

`func NewTrailerStatusRead(fleetId NullableString, id string, createdAt time.Time, updatedAt time.Time, connectionId string, sourceName TspEnum, sourceId string, sourceDataHash string, ) *TrailerStatusRead`

NewTrailerStatusRead instantiates a new TrailerStatusRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrailerStatusReadWithDefaults

`func NewTrailerStatusReadWithDefaults() *TrailerStatusRead`

NewTrailerStatusReadWithDefaults instantiates a new TrailerStatusRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFleetId

`func (o *TrailerStatusRead) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *TrailerStatusRead) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *TrailerStatusRead) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### SetFleetIdNil

`func (o *TrailerStatusRead) SetFleetIdNil(b bool)`

 SetFleetIdNil sets the value for FleetId to be an explicit nil

### UnsetFleetId
`func (o *TrailerStatusRead) UnsetFleetId()`

UnsetFleetId ensures that no value is present for FleetId, not even an explicit nil
### GetFleetRef

`func (o *TrailerStatusRead) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *TrailerStatusRead) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *TrailerStatusRead) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.

### HasFleetRef

`func (o *TrailerStatusRead) HasFleetRef() bool`

HasFleetRef returns a boolean if a field has been set.

### SetFleetRefNil

`func (o *TrailerStatusRead) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *TrailerStatusRead) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetId

`func (o *TrailerStatusRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrailerStatusRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrailerStatusRead) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *TrailerStatusRead) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TrailerStatusRead) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TrailerStatusRead) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *TrailerStatusRead) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TrailerStatusRead) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TrailerStatusRead) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *TrailerStatusRead) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *TrailerStatusRead) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *TrailerStatusRead) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *TrailerStatusRead) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *TrailerStatusRead) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *TrailerStatusRead) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetConnectionId

`func (o *TrailerStatusRead) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *TrailerStatusRead) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *TrailerStatusRead) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetTspId

`func (o *TrailerStatusRead) GetTspId() string`

GetTspId returns the TspId field if non-nil, zero value otherwise.

### GetTspIdOk

`func (o *TrailerStatusRead) GetTspIdOk() (*string, bool)`

GetTspIdOk returns a tuple with the TspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspId

`func (o *TrailerStatusRead) SetTspId(v string)`

SetTspId sets TspId field to given value.

### HasTspId

`func (o *TrailerStatusRead) HasTspId() bool`

HasTspId returns a boolean if a field has been set.

### SetTspIdNil

`func (o *TrailerStatusRead) SetTspIdNil(b bool)`

 SetTspIdNil sets the value for TspId to be an explicit nil

### UnsetTspId
`func (o *TrailerStatusRead) UnsetTspId()`

UnsetTspId ensures that no value is present for TspId, not even an explicit nil
### GetTspSlug

`func (o *TrailerStatusRead) GetTspSlug() string`

GetTspSlug returns the TspSlug field if non-nil, zero value otherwise.

### GetTspSlugOk

`func (o *TrailerStatusRead) GetTspSlugOk() (*string, bool)`

GetTspSlugOk returns a tuple with the TspSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspSlug

`func (o *TrailerStatusRead) SetTspSlug(v string)`

SetTspSlug sets TspSlug field to given value.

### HasTspSlug

`func (o *TrailerStatusRead) HasTspSlug() bool`

HasTspSlug returns a boolean if a field has been set.

### SetTspSlugNil

`func (o *TrailerStatusRead) SetTspSlugNil(b bool)`

 SetTspSlugNil sets the value for TspSlug to be an explicit nil

### UnsetTspSlug
`func (o *TrailerStatusRead) UnsetTspSlug()`

UnsetTspSlug ensures that no value is present for TspSlug, not even an explicit nil
### GetSourceName

`func (o *TrailerStatusRead) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *TrailerStatusRead) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *TrailerStatusRead) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetSourceData

`func (o *TrailerStatusRead) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *TrailerStatusRead) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *TrailerStatusRead) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *TrailerStatusRead) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetSourceId

`func (o *TrailerStatusRead) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *TrailerStatusRead) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *TrailerStatusRead) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSourceDataHash

`func (o *TrailerStatusRead) GetSourceDataHash() string`

GetSourceDataHash returns the SourceDataHash field if non-nil, zero value otherwise.

### GetSourceDataHashOk

`func (o *TrailerStatusRead) GetSourceDataHashOk() (*string, bool)`

GetSourceDataHashOk returns a tuple with the SourceDataHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataHash

`func (o *TrailerStatusRead) SetSourceDataHash(v string)`

SetSourceDataHash sets SourceDataHash field to given value.


### GetOccurredAt

`func (o *TrailerStatusRead) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *TrailerStatusRead) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *TrailerStatusRead) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *TrailerStatusRead) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### SetOccurredAtNil

`func (o *TrailerStatusRead) SetOccurredAtNil(b bool)`

 SetOccurredAtNil sets the value for OccurredAt to be an explicit nil

### UnsetOccurredAt
`func (o *TrailerStatusRead) UnsetOccurredAt()`

UnsetOccurredAt ensures that no value is present for OccurredAt, not even an explicit nil
### GetExecutionId

`func (o *TrailerStatusRead) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *TrailerStatusRead) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *TrailerStatusRead) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *TrailerStatusRead) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *TrailerStatusRead) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *TrailerStatusRead) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *TrailerStatusRead) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *TrailerStatusRead) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *TrailerStatusRead) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *TrailerStatusRead) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *TrailerStatusRead) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *TrailerStatusRead) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetExtras

`func (o *TrailerStatusRead) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *TrailerStatusRead) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *TrailerStatusRead) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *TrailerStatusRead) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### SetExtrasNil

`func (o *TrailerStatusRead) SetExtrasNil(b bool)`

 SetExtrasNil sets the value for Extras to be an explicit nil

### UnsetExtras
`func (o *TrailerStatusRead) UnsetExtras()`

UnsetExtras ensures that no value is present for Extras, not even an explicit nil
### GetTrailerId

`func (o *TrailerStatusRead) GetTrailerId() string`

GetTrailerId returns the TrailerId field if non-nil, zero value otherwise.

### GetTrailerIdOk

`func (o *TrailerStatusRead) GetTrailerIdOk() (*string, bool)`

GetTrailerIdOk returns a tuple with the TrailerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerId

`func (o *TrailerStatusRead) SetTrailerId(v string)`

SetTrailerId sets TrailerId field to given value.

### HasTrailerId

`func (o *TrailerStatusRead) HasTrailerId() bool`

HasTrailerId returns a boolean if a field has been set.

### SetTrailerIdNil

`func (o *TrailerStatusRead) SetTrailerIdNil(b bool)`

 SetTrailerIdNil sets the value for TrailerId to be an explicit nil

### UnsetTrailerId
`func (o *TrailerStatusRead) UnsetTrailerId()`

UnsetTrailerId ensures that no value is present for TrailerId, not even an explicit nil
### GetSourceTrailerId

`func (o *TrailerStatusRead) GetSourceTrailerId() string`

GetSourceTrailerId returns the SourceTrailerId field if non-nil, zero value otherwise.

### GetSourceTrailerIdOk

`func (o *TrailerStatusRead) GetSourceTrailerIdOk() (*string, bool)`

GetSourceTrailerIdOk returns a tuple with the SourceTrailerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTrailerId

`func (o *TrailerStatusRead) SetSourceTrailerId(v string)`

SetSourceTrailerId sets SourceTrailerId field to given value.

### HasSourceTrailerId

`func (o *TrailerStatusRead) HasSourceTrailerId() bool`

HasSourceTrailerId returns a boolean if a field has been set.

### SetSourceTrailerIdNil

`func (o *TrailerStatusRead) SetSourceTrailerIdNil(b bool)`

 SetSourceTrailerIdNil sets the value for SourceTrailerId to be an explicit nil

### UnsetSourceTrailerId
`func (o *TrailerStatusRead) UnsetSourceTrailerId()`

UnsetSourceTrailerId ensures that no value is present for SourceTrailerId, not even an explicit nil
### GetPowerOn

`func (o *TrailerStatusRead) GetPowerOn() bool`

GetPowerOn returns the PowerOn field if non-nil, zero value otherwise.

### GetPowerOnOk

`func (o *TrailerStatusRead) GetPowerOnOk() (*bool, bool)`

GetPowerOnOk returns a tuple with the PowerOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOn

`func (o *TrailerStatusRead) SetPowerOn(v bool)`

SetPowerOn sets PowerOn field to given value.

### HasPowerOn

`func (o *TrailerStatusRead) HasPowerOn() bool`

HasPowerOn returns a boolean if a field has been set.

### SetPowerOnNil

`func (o *TrailerStatusRead) SetPowerOnNil(b bool)`

 SetPowerOnNil sets the value for PowerOn to be an explicit nil

### UnsetPowerOn
`func (o *TrailerStatusRead) UnsetPowerOn()`

UnsetPowerOn ensures that no value is present for PowerOn, not even an explicit nil
### GetPowerSource

`func (o *TrailerStatusRead) GetPowerSource() ReeferPowerSourceEnum`

GetPowerSource returns the PowerSource field if non-nil, zero value otherwise.

### GetPowerSourceOk

`func (o *TrailerStatusRead) GetPowerSourceOk() (*ReeferPowerSourceEnum, bool)`

GetPowerSourceOk returns a tuple with the PowerSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSource

`func (o *TrailerStatusRead) SetPowerSource(v ReeferPowerSourceEnum)`

SetPowerSource sets PowerSource field to given value.

### HasPowerSource

`func (o *TrailerStatusRead) HasPowerSource() bool`

HasPowerSource returns a boolean if a field has been set.

### SetPowerSourceNil

`func (o *TrailerStatusRead) SetPowerSourceNil(b bool)`

 SetPowerSourceNil sets the value for PowerSource to be an explicit nil

### UnsetPowerSource
`func (o *TrailerStatusRead) UnsetPowerSource()`

UnsetPowerSource ensures that no value is present for PowerSource, not even an explicit nil
### GetBatteryVoltage

`func (o *TrailerStatusRead) GetBatteryVoltage() float32`

GetBatteryVoltage returns the BatteryVoltage field if non-nil, zero value otherwise.

### GetBatteryVoltageOk

`func (o *TrailerStatusRead) GetBatteryVoltageOk() (*float32, bool)`

GetBatteryVoltageOk returns a tuple with the BatteryVoltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryVoltage

`func (o *TrailerStatusRead) SetBatteryVoltage(v float32)`

SetBatteryVoltage sets BatteryVoltage field to given value.

### HasBatteryVoltage

`func (o *TrailerStatusRead) HasBatteryVoltage() bool`

HasBatteryVoltage returns a boolean if a field has been set.

### SetBatteryVoltageNil

`func (o *TrailerStatusRead) SetBatteryVoltageNil(b bool)`

 SetBatteryVoltageNil sets the value for BatteryVoltage to be an explicit nil

### UnsetBatteryVoltage
`func (o *TrailerStatusRead) UnsetBatteryVoltage()`

UnsetBatteryVoltage ensures that no value is present for BatteryVoltage, not even an explicit nil
### GetBatteryPercent

`func (o *TrailerStatusRead) GetBatteryPercent() float32`

GetBatteryPercent returns the BatteryPercent field if non-nil, zero value otherwise.

### GetBatteryPercentOk

`func (o *TrailerStatusRead) GetBatteryPercentOk() (*float32, bool)`

GetBatteryPercentOk returns a tuple with the BatteryPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryPercent

`func (o *TrailerStatusRead) SetBatteryPercent(v float32)`

SetBatteryPercent sets BatteryPercent field to given value.

### HasBatteryPercent

`func (o *TrailerStatusRead) HasBatteryPercent() bool`

HasBatteryPercent returns a boolean if a field has been set.

### SetBatteryPercentNil

`func (o *TrailerStatusRead) SetBatteryPercentNil(b bool)`

 SetBatteryPercentNil sets the value for BatteryPercent to be an explicit nil

### UnsetBatteryPercent
`func (o *TrailerStatusRead) UnsetBatteryPercent()`

UnsetBatteryPercent ensures that no value is present for BatteryPercent, not even an explicit nil
### GetBatteryStatus

`func (o *TrailerStatusRead) GetBatteryStatus() ReeferBatteryStatusEnum`

GetBatteryStatus returns the BatteryStatus field if non-nil, zero value otherwise.

### GetBatteryStatusOk

`func (o *TrailerStatusRead) GetBatteryStatusOk() (*ReeferBatteryStatusEnum, bool)`

GetBatteryStatusOk returns a tuple with the BatteryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryStatus

`func (o *TrailerStatusRead) SetBatteryStatus(v ReeferBatteryStatusEnum)`

SetBatteryStatus sets BatteryStatus field to given value.

### HasBatteryStatus

`func (o *TrailerStatusRead) HasBatteryStatus() bool`

HasBatteryStatus returns a boolean if a field has been set.

### SetBatteryStatusNil

`func (o *TrailerStatusRead) SetBatteryStatusNil(b bool)`

 SetBatteryStatusNil sets the value for BatteryStatus to be an explicit nil

### UnsetBatteryStatus
`func (o *TrailerStatusRead) UnsetBatteryStatus()`

UnsetBatteryStatus ensures that no value is present for BatteryStatus, not even an explicit nil
### GetReeferOperationalStatus

`func (o *TrailerStatusRead) GetReeferOperationalStatus() ReeferOperationalStatusEnum`

GetReeferOperationalStatus returns the ReeferOperationalStatus field if non-nil, zero value otherwise.

### GetReeferOperationalStatusOk

`func (o *TrailerStatusRead) GetReeferOperationalStatusOk() (*ReeferOperationalStatusEnum, bool)`

GetReeferOperationalStatusOk returns a tuple with the ReeferOperationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReeferOperationalStatus

`func (o *TrailerStatusRead) SetReeferOperationalStatus(v ReeferOperationalStatusEnum)`

SetReeferOperationalStatus sets ReeferOperationalStatus field to given value.

### HasReeferOperationalStatus

`func (o *TrailerStatusRead) HasReeferOperationalStatus() bool`

HasReeferOperationalStatus returns a boolean if a field has been set.

### SetReeferOperationalStatusNil

`func (o *TrailerStatusRead) SetReeferOperationalStatusNil(b bool)`

 SetReeferOperationalStatusNil sets the value for ReeferOperationalStatus to be an explicit nil

### UnsetReeferOperationalStatus
`func (o *TrailerStatusRead) UnsetReeferOperationalStatus()`

UnsetReeferOperationalStatus ensures that no value is present for ReeferOperationalStatus, not even an explicit nil
### GetOperationMode

`func (o *TrailerStatusRead) GetOperationMode() ReeferOperationModeEnum`

GetOperationMode returns the OperationMode field if non-nil, zero value otherwise.

### GetOperationModeOk

`func (o *TrailerStatusRead) GetOperationModeOk() (*ReeferOperationModeEnum, bool)`

GetOperationModeOk returns a tuple with the OperationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationMode

`func (o *TrailerStatusRead) SetOperationMode(v ReeferOperationModeEnum)`

SetOperationMode sets OperationMode field to given value.

### HasOperationMode

`func (o *TrailerStatusRead) HasOperationMode() bool`

HasOperationMode returns a boolean if a field has been set.

### SetOperationModeNil

`func (o *TrailerStatusRead) SetOperationModeNil(b bool)`

 SetOperationModeNil sets the value for OperationMode to be an explicit nil

### UnsetOperationMode
`func (o *TrailerStatusRead) UnsetOperationMode()`

UnsetOperationMode ensures that no value is present for OperationMode, not even an explicit nil
### GetControlMode

`func (o *TrailerStatusRead) GetControlMode() ReeferControlModeEnum`

GetControlMode returns the ControlMode field if non-nil, zero value otherwise.

### GetControlModeOk

`func (o *TrailerStatusRead) GetControlModeOk() (*ReeferControlModeEnum, bool)`

GetControlModeOk returns a tuple with the ControlMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlMode

`func (o *TrailerStatusRead) SetControlMode(v ReeferControlModeEnum)`

SetControlMode sets ControlMode field to given value.

### HasControlMode

`func (o *TrailerStatusRead) HasControlMode() bool`

HasControlMode returns a boolean if a field has been set.

### SetControlModeNil

`func (o *TrailerStatusRead) SetControlModeNil(b bool)`

 SetControlModeNil sets the value for ControlMode to be an explicit nil

### UnsetControlMode
`func (o *TrailerStatusRead) UnsetControlMode()`

UnsetControlMode ensures that no value is present for ControlMode, not even an explicit nil
### GetDoorOpen

`func (o *TrailerStatusRead) GetDoorOpen() bool`

GetDoorOpen returns the DoorOpen field if non-nil, zero value otherwise.

### GetDoorOpenOk

`func (o *TrailerStatusRead) GetDoorOpenOk() (*bool, bool)`

GetDoorOpenOk returns a tuple with the DoorOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoorOpen

`func (o *TrailerStatusRead) SetDoorOpen(v bool)`

SetDoorOpen sets DoorOpen field to given value.

### HasDoorOpen

`func (o *TrailerStatusRead) HasDoorOpen() bool`

HasDoorOpen returns a boolean if a field has been set.

### SetDoorOpenNil

`func (o *TrailerStatusRead) SetDoorOpenNil(b bool)`

 SetDoorOpenNil sets the value for DoorOpen to be an explicit nil

### UnsetDoorOpen
`func (o *TrailerStatusRead) UnsetDoorOpen()`

UnsetDoorOpen ensures that no value is present for DoorOpen, not even an explicit nil
### GetCargoStatus

`func (o *TrailerStatusRead) GetCargoStatus() ReeferCargoStatusEnum`

GetCargoStatus returns the CargoStatus field if non-nil, zero value otherwise.

### GetCargoStatusOk

`func (o *TrailerStatusRead) GetCargoStatusOk() (*ReeferCargoStatusEnum, bool)`

GetCargoStatusOk returns a tuple with the CargoStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCargoStatus

`func (o *TrailerStatusRead) SetCargoStatus(v ReeferCargoStatusEnum)`

SetCargoStatus sets CargoStatus field to given value.

### HasCargoStatus

`func (o *TrailerStatusRead) HasCargoStatus() bool`

HasCargoStatus returns a boolean if a field has been set.

### SetCargoStatusNil

`func (o *TrailerStatusRead) SetCargoStatusNil(b bool)`

 SetCargoStatusNil sets the value for CargoStatus to be an explicit nil

### UnsetCargoStatus
`func (o *TrailerStatusRead) UnsetCargoStatus()`

UnsetCargoStatus ensures that no value is present for CargoStatus, not even an explicit nil
### GetFuelLevel

`func (o *TrailerStatusRead) GetFuelLevel() float32`

GetFuelLevel returns the FuelLevel field if non-nil, zero value otherwise.

### GetFuelLevelOk

`func (o *TrailerStatusRead) GetFuelLevelOk() (*float32, bool)`

GetFuelLevelOk returns a tuple with the FuelLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelLevel

`func (o *TrailerStatusRead) SetFuelLevel(v float32)`

SetFuelLevel sets FuelLevel field to given value.

### HasFuelLevel

`func (o *TrailerStatusRead) HasFuelLevel() bool`

HasFuelLevel returns a boolean if a field has been set.

### SetFuelLevelNil

`func (o *TrailerStatusRead) SetFuelLevelNil(b bool)`

 SetFuelLevelNil sets the value for FuelLevel to be an explicit nil

### UnsetFuelLevel
`func (o *TrailerStatusRead) UnsetFuelLevel()`

UnsetFuelLevel ensures that no value is present for FuelLevel, not even an explicit nil
### GetReeferEngineHours

`func (o *TrailerStatusRead) GetReeferEngineHours() float32`

GetReeferEngineHours returns the ReeferEngineHours field if non-nil, zero value otherwise.

### GetReeferEngineHoursOk

`func (o *TrailerStatusRead) GetReeferEngineHoursOk() (*float32, bool)`

GetReeferEngineHoursOk returns a tuple with the ReeferEngineHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReeferEngineHours

`func (o *TrailerStatusRead) SetReeferEngineHours(v float32)`

SetReeferEngineHours sets ReeferEngineHours field to given value.

### HasReeferEngineHours

`func (o *TrailerStatusRead) HasReeferEngineHours() bool`

HasReeferEngineHours returns a boolean if a field has been set.

### SetReeferEngineHoursNil

`func (o *TrailerStatusRead) SetReeferEngineHoursNil(b bool)`

 SetReeferEngineHoursNil sets the value for ReeferEngineHours to be an explicit nil

### UnsetReeferEngineHours
`func (o *TrailerStatusRead) UnsetReeferEngineHours()`

UnsetReeferEngineHours ensures that no value is present for ReeferEngineHours, not even an explicit nil
### GetAmbientTemperature

`func (o *TrailerStatusRead) GetAmbientTemperature() float32`

GetAmbientTemperature returns the AmbientTemperature field if non-nil, zero value otherwise.

### GetAmbientTemperatureOk

`func (o *TrailerStatusRead) GetAmbientTemperatureOk() (*float32, bool)`

GetAmbientTemperatureOk returns a tuple with the AmbientTemperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmbientTemperature

`func (o *TrailerStatusRead) SetAmbientTemperature(v float32)`

SetAmbientTemperature sets AmbientTemperature field to given value.

### HasAmbientTemperature

`func (o *TrailerStatusRead) HasAmbientTemperature() bool`

HasAmbientTemperature returns a boolean if a field has been set.

### SetAmbientTemperatureNil

`func (o *TrailerStatusRead) SetAmbientTemperatureNil(b bool)`

 SetAmbientTemperatureNil sets the value for AmbientTemperature to be an explicit nil

### UnsetAmbientTemperature
`func (o *TrailerStatusRead) UnsetAmbientTemperature()`

UnsetAmbientTemperature ensures that no value is present for AmbientTemperature, not even an explicit nil
### GetZones

`func (o *TrailerStatusRead) GetZones() []TrailerZoneStatus`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *TrailerStatusRead) GetZonesOk() (*[]TrailerZoneStatus, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *TrailerStatusRead) SetZones(v []TrailerZoneStatus)`

SetZones sets Zones field to given value.

### HasZones

`func (o *TrailerStatusRead) HasZones() bool`

HasZones returns a boolean if a field has been set.

### SetZonesNil

`func (o *TrailerStatusRead) SetZonesNil(b bool)`

 SetZonesNil sets the value for Zones to be an explicit nil

### UnsetZones
`func (o *TrailerStatusRead) UnsetZones()`

UnsetZones ensures that no value is present for Zones, not even an explicit nil
### GetRemoteProbeStatuses

`func (o *TrailerStatusRead) GetRemoteProbeStatuses() []RemoteProbeStatus`

GetRemoteProbeStatuses returns the RemoteProbeStatuses field if non-nil, zero value otherwise.

### GetRemoteProbeStatusesOk

`func (o *TrailerStatusRead) GetRemoteProbeStatusesOk() (*[]RemoteProbeStatus, bool)`

GetRemoteProbeStatusesOk returns a tuple with the RemoteProbeStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteProbeStatuses

`func (o *TrailerStatusRead) SetRemoteProbeStatuses(v []RemoteProbeStatus)`

SetRemoteProbeStatuses sets RemoteProbeStatuses field to given value.

### HasRemoteProbeStatuses

`func (o *TrailerStatusRead) HasRemoteProbeStatuses() bool`

HasRemoteProbeStatuses returns a boolean if a field has been set.

### SetRemoteProbeStatusesNil

`func (o *TrailerStatusRead) SetRemoteProbeStatusesNil(b bool)`

 SetRemoteProbeStatusesNil sets the value for RemoteProbeStatuses to be an explicit nil

### UnsetRemoteProbeStatuses
`func (o *TrailerStatusRead) UnsetRemoteProbeStatuses()`

UnsetRemoteProbeStatuses ensures that no value is present for RemoteProbeStatuses, not even an explicit nil
### GetPreTripStatus

`func (o *TrailerStatusRead) GetPreTripStatus() TrailerPretripStatus`

GetPreTripStatus returns the PreTripStatus field if non-nil, zero value otherwise.

### GetPreTripStatusOk

`func (o *TrailerStatusRead) GetPreTripStatusOk() (*TrailerPretripStatus, bool)`

GetPreTripStatusOk returns a tuple with the PreTripStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreTripStatus

`func (o *TrailerStatusRead) SetPreTripStatus(v TrailerPretripStatus)`

SetPreTripStatus sets PreTripStatus field to given value.

### HasPreTripStatus

`func (o *TrailerStatusRead) HasPreTripStatus() bool`

HasPreTripStatus returns a boolean if a field has been set.

### SetPreTripStatusNil

`func (o *TrailerStatusRead) SetPreTripStatusNil(b bool)`

 SetPreTripStatusNil sets the value for PreTripStatus to be an explicit nil

### UnsetPreTripStatus
`func (o *TrailerStatusRead) UnsetPreTripStatus()`

UnsetPreTripStatus ensures that no value is present for PreTripStatus, not even an explicit nil
### GetSuctionPressureKpa

`func (o *TrailerStatusRead) GetSuctionPressureKpa() float32`

GetSuctionPressureKpa returns the SuctionPressureKpa field if non-nil, zero value otherwise.

### GetSuctionPressureKpaOk

`func (o *TrailerStatusRead) GetSuctionPressureKpaOk() (*float32, bool)`

GetSuctionPressureKpaOk returns a tuple with the SuctionPressureKpa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuctionPressureKpa

`func (o *TrailerStatusRead) SetSuctionPressureKpa(v float32)`

SetSuctionPressureKpa sets SuctionPressureKpa field to given value.

### HasSuctionPressureKpa

`func (o *TrailerStatusRead) HasSuctionPressureKpa() bool`

HasSuctionPressureKpa returns a boolean if a field has been set.

### SetSuctionPressureKpaNil

`func (o *TrailerStatusRead) SetSuctionPressureKpaNil(b bool)`

 SetSuctionPressureKpaNil sets the value for SuctionPressureKpa to be an explicit nil

### UnsetSuctionPressureKpa
`func (o *TrailerStatusRead) UnsetSuctionPressureKpa()`

UnsetSuctionPressureKpa ensures that no value is present for SuctionPressureKpa, not even an explicit nil
### GetDischargePressureKpa

`func (o *TrailerStatusRead) GetDischargePressureKpa() float32`

GetDischargePressureKpa returns the DischargePressureKpa field if non-nil, zero value otherwise.

### GetDischargePressureKpaOk

`func (o *TrailerStatusRead) GetDischargePressureKpaOk() (*float32, bool)`

GetDischargePressureKpaOk returns a tuple with the DischargePressureKpa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDischargePressureKpa

`func (o *TrailerStatusRead) SetDischargePressureKpa(v float32)`

SetDischargePressureKpa sets DischargePressureKpa field to given value.

### HasDischargePressureKpa

`func (o *TrailerStatusRead) HasDischargePressureKpa() bool`

HasDischargePressureKpa returns a boolean if a field has been set.

### SetDischargePressureKpaNil

`func (o *TrailerStatusRead) SetDischargePressureKpaNil(b bool)`

 SetDischargePressureKpaNil sets the value for DischargePressureKpa to be an explicit nil

### UnsetDischargePressureKpa
`func (o *TrailerStatusRead) UnsetDischargePressureKpa()`

UnsetDischargePressureKpa ensures that no value is present for DischargePressureKpa, not even an explicit nil
### GetHumidityPercent

`func (o *TrailerStatusRead) GetHumidityPercent() float32`

GetHumidityPercent returns the HumidityPercent field if non-nil, zero value otherwise.

### GetHumidityPercentOk

`func (o *TrailerStatusRead) GetHumidityPercentOk() (*float32, bool)`

GetHumidityPercentOk returns a tuple with the HumidityPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumidityPercent

`func (o *TrailerStatusRead) SetHumidityPercent(v float32)`

SetHumidityPercent sets HumidityPercent field to given value.

### HasHumidityPercent

`func (o *TrailerStatusRead) HasHumidityPercent() bool`

HasHumidityPercent returns a boolean if a field has been set.

### SetHumidityPercentNil

`func (o *TrailerStatusRead) SetHumidityPercentNil(b bool)`

 SetHumidityPercentNil sets the value for HumidityPercent to be an explicit nil

### UnsetHumidityPercent
`func (o *TrailerStatusRead) UnsetHumidityPercent()`

UnsetHumidityPercent ensures that no value is present for HumidityPercent, not even an explicit nil
### GetActiveAlarms

`func (o *TrailerStatusRead) GetActiveAlarms() []TrailerAlarm`

GetActiveAlarms returns the ActiveAlarms field if non-nil, zero value otherwise.

### GetActiveAlarmsOk

`func (o *TrailerStatusRead) GetActiveAlarmsOk() (*[]TrailerAlarm, bool)`

GetActiveAlarmsOk returns a tuple with the ActiveAlarms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAlarms

`func (o *TrailerStatusRead) SetActiveAlarms(v []TrailerAlarm)`

SetActiveAlarms sets ActiveAlarms field to given value.

### HasActiveAlarms

`func (o *TrailerStatusRead) HasActiveAlarms() bool`

HasActiveAlarms returns a boolean if a field has been set.

### SetActiveAlarmsNil

`func (o *TrailerStatusRead) SetActiveAlarmsNil(b bool)`

 SetActiveAlarmsNil sets the value for ActiveAlarms to be an explicit nil

### UnsetActiveAlarms
`func (o *TrailerStatusRead) UnsetActiveAlarms()`

UnsetActiveAlarms ensures that no value is present for ActiveAlarms, not even an explicit nil
### GetLocation

`func (o *TrailerStatusRead) GetLocation() Location5`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *TrailerStatusRead) GetLocationOk() (*Location5, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *TrailerStatusRead) SetLocation(v Location5)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *TrailerStatusRead) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *TrailerStatusRead) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *TrailerStatusRead) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetH3Index11

`func (o *TrailerStatusRead) GetH3Index11() int32`

GetH3Index11 returns the H3Index11 field if non-nil, zero value otherwise.

### GetH3Index11Ok

`func (o *TrailerStatusRead) GetH3Index11Ok() (*int32, bool)`

GetH3Index11Ok returns a tuple with the H3Index11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH3Index11

`func (o *TrailerStatusRead) SetH3Index11(v int32)`

SetH3Index11 sets H3Index11 field to given value.

### HasH3Index11

`func (o *TrailerStatusRead) HasH3Index11() bool`

HasH3Index11 returns a boolean if a field has been set.

### SetH3Index11Nil

`func (o *TrailerStatusRead) SetH3Index11Nil(b bool)`

 SetH3Index11Nil sets the value for H3Index11 to be an explicit nil

### UnsetH3Index11
`func (o *TrailerStatusRead) UnsetH3Index11()`

UnsetH3Index11 ensures that no value is present for H3Index11, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


