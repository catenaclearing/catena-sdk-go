# ExtTspIntegrationRequestRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique integration request identifier | 
**ExtTspId** | **string** | The external TSP that was requested for integration | 
**ExtTspName** | **string** | Name of the external TSP at the time of request | 
**ExtTspComplianceStatus** | [**TspComplianceStatusEnum**](TspComplianceStatusEnum.md) | Compliance status at the time of request. &#39;compliant&#39; &#x3D; at least one device registered, &#39;non_compliant&#39; &#x3D; ALL devices revoked, &#39;unknown&#39; &#x3D; not linked to a registry | 
**FleetId** | **string** | The fleet requesting integration | 
**FleetRef** | **string** | External fleet identifier from the partner&#39;s system | 
**PartnerId** | **string** | The partner requesting the integration | 
**InvitationId** | **string** | The related invitation for the fleet requesting the integration | 
**FleetContactName** | Pointer to **NullableString** |  | [optional] 
**FleetContactNotes** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** | When the integration request was created | 
**UpdatedAt** | **time.Time** | Last modification timestamp | 
**IntegratedTspIds** | Pointer to **[]string** | List of TSP IDs that are integrated for this external TSP. Empty if not yet integrated. | [optional] 

## Methods

### NewExtTspIntegrationRequestRead

`func NewExtTspIntegrationRequestRead(id string, extTspId string, extTspName string, extTspComplianceStatus TspComplianceStatusEnum, fleetId string, fleetRef string, partnerId string, invitationId string, createdAt time.Time, updatedAt time.Time, ) *ExtTspIntegrationRequestRead`

NewExtTspIntegrationRequestRead instantiates a new ExtTspIntegrationRequestRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtTspIntegrationRequestReadWithDefaults

`func NewExtTspIntegrationRequestReadWithDefaults() *ExtTspIntegrationRequestRead`

NewExtTspIntegrationRequestReadWithDefaults instantiates a new ExtTspIntegrationRequestRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExtTspIntegrationRequestRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExtTspIntegrationRequestRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExtTspIntegrationRequestRead) SetId(v string)`

SetId sets Id field to given value.


### GetExtTspId

`func (o *ExtTspIntegrationRequestRead) GetExtTspId() string`

GetExtTspId returns the ExtTspId field if non-nil, zero value otherwise.

### GetExtTspIdOk

`func (o *ExtTspIntegrationRequestRead) GetExtTspIdOk() (*string, bool)`

GetExtTspIdOk returns a tuple with the ExtTspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtTspId

`func (o *ExtTspIntegrationRequestRead) SetExtTspId(v string)`

SetExtTspId sets ExtTspId field to given value.


### GetExtTspName

`func (o *ExtTspIntegrationRequestRead) GetExtTspName() string`

GetExtTspName returns the ExtTspName field if non-nil, zero value otherwise.

### GetExtTspNameOk

`func (o *ExtTspIntegrationRequestRead) GetExtTspNameOk() (*string, bool)`

GetExtTspNameOk returns a tuple with the ExtTspName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtTspName

`func (o *ExtTspIntegrationRequestRead) SetExtTspName(v string)`

SetExtTspName sets ExtTspName field to given value.


### GetExtTspComplianceStatus

`func (o *ExtTspIntegrationRequestRead) GetExtTspComplianceStatus() TspComplianceStatusEnum`

GetExtTspComplianceStatus returns the ExtTspComplianceStatus field if non-nil, zero value otherwise.

### GetExtTspComplianceStatusOk

`func (o *ExtTspIntegrationRequestRead) GetExtTspComplianceStatusOk() (*TspComplianceStatusEnum, bool)`

GetExtTspComplianceStatusOk returns a tuple with the ExtTspComplianceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtTspComplianceStatus

`func (o *ExtTspIntegrationRequestRead) SetExtTspComplianceStatus(v TspComplianceStatusEnum)`

SetExtTspComplianceStatus sets ExtTspComplianceStatus field to given value.


### GetFleetId

`func (o *ExtTspIntegrationRequestRead) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *ExtTspIntegrationRequestRead) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *ExtTspIntegrationRequestRead) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetFleetRef

`func (o *ExtTspIntegrationRequestRead) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *ExtTspIntegrationRequestRead) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *ExtTspIntegrationRequestRead) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.


### GetPartnerId

`func (o *ExtTspIntegrationRequestRead) GetPartnerId() string`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *ExtTspIntegrationRequestRead) GetPartnerIdOk() (*string, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *ExtTspIntegrationRequestRead) SetPartnerId(v string)`

SetPartnerId sets PartnerId field to given value.


### GetInvitationId

`func (o *ExtTspIntegrationRequestRead) GetInvitationId() string`

GetInvitationId returns the InvitationId field if non-nil, zero value otherwise.

### GetInvitationIdOk

`func (o *ExtTspIntegrationRequestRead) GetInvitationIdOk() (*string, bool)`

GetInvitationIdOk returns a tuple with the InvitationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationId

`func (o *ExtTspIntegrationRequestRead) SetInvitationId(v string)`

SetInvitationId sets InvitationId field to given value.


### GetFleetContactName

`func (o *ExtTspIntegrationRequestRead) GetFleetContactName() string`

GetFleetContactName returns the FleetContactName field if non-nil, zero value otherwise.

### GetFleetContactNameOk

`func (o *ExtTspIntegrationRequestRead) GetFleetContactNameOk() (*string, bool)`

GetFleetContactNameOk returns a tuple with the FleetContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetContactName

`func (o *ExtTspIntegrationRequestRead) SetFleetContactName(v string)`

SetFleetContactName sets FleetContactName field to given value.

### HasFleetContactName

`func (o *ExtTspIntegrationRequestRead) HasFleetContactName() bool`

HasFleetContactName returns a boolean if a field has been set.

### SetFleetContactNameNil

`func (o *ExtTspIntegrationRequestRead) SetFleetContactNameNil(b bool)`

 SetFleetContactNameNil sets the value for FleetContactName to be an explicit nil

### UnsetFleetContactName
`func (o *ExtTspIntegrationRequestRead) UnsetFleetContactName()`

UnsetFleetContactName ensures that no value is present for FleetContactName, not even an explicit nil
### GetFleetContactNotes

`func (o *ExtTspIntegrationRequestRead) GetFleetContactNotes() string`

GetFleetContactNotes returns the FleetContactNotes field if non-nil, zero value otherwise.

### GetFleetContactNotesOk

`func (o *ExtTspIntegrationRequestRead) GetFleetContactNotesOk() (*string, bool)`

GetFleetContactNotesOk returns a tuple with the FleetContactNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetContactNotes

`func (o *ExtTspIntegrationRequestRead) SetFleetContactNotes(v string)`

SetFleetContactNotes sets FleetContactNotes field to given value.

### HasFleetContactNotes

`func (o *ExtTspIntegrationRequestRead) HasFleetContactNotes() bool`

HasFleetContactNotes returns a boolean if a field has been set.

### SetFleetContactNotesNil

`func (o *ExtTspIntegrationRequestRead) SetFleetContactNotesNil(b bool)`

 SetFleetContactNotesNil sets the value for FleetContactNotes to be an explicit nil

### UnsetFleetContactNotes
`func (o *ExtTspIntegrationRequestRead) UnsetFleetContactNotes()`

UnsetFleetContactNotes ensures that no value is present for FleetContactNotes, not even an explicit nil
### GetCreatedAt

`func (o *ExtTspIntegrationRequestRead) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExtTspIntegrationRequestRead) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExtTspIntegrationRequestRead) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ExtTspIntegrationRequestRead) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ExtTspIntegrationRequestRead) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ExtTspIntegrationRequestRead) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIntegratedTspIds

`func (o *ExtTspIntegrationRequestRead) GetIntegratedTspIds() []string`

GetIntegratedTspIds returns the IntegratedTspIds field if non-nil, zero value otherwise.

### GetIntegratedTspIdsOk

`func (o *ExtTspIntegrationRequestRead) GetIntegratedTspIdsOk() (*[]string, bool)`

GetIntegratedTspIdsOk returns a tuple with the IntegratedTspIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegratedTspIds

`func (o *ExtTspIntegrationRequestRead) SetIntegratedTspIds(v []string)`

SetIntegratedTspIds sets IntegratedTspIds field to given value.

### HasIntegratedTspIds

`func (o *ExtTspIntegrationRequestRead) HasIntegratedTspIds() bool`

HasIntegratedTspIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


