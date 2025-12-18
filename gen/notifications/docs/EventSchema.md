# EventSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventName** | **string** | The name of the event | 
**Version** | **string** | The version of the schema | 
**CreatedAt** | Pointer to **time.Time** | The date and time the schema was created | [optional] 
**OpenapiSchema** | **string** | The complete OpenAPI 3.0 document containing the event schema as a JSON string | 
**Changelog** | Pointer to **string** | The changes made to this version of the schema, when compared to the previous version | [optional] [default to "Initial schema version"]

## Methods

### NewEventSchema

`func NewEventSchema(eventName string, version string, openapiSchema string, ) *EventSchema`

NewEventSchema instantiates a new EventSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventSchemaWithDefaults

`func NewEventSchemaWithDefaults() *EventSchema`

NewEventSchemaWithDefaults instantiates a new EventSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventName

`func (o *EventSchema) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *EventSchema) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *EventSchema) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetVersion

`func (o *EventSchema) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EventSchema) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EventSchema) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetCreatedAt

`func (o *EventSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EventSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EventSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EventSchema) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetOpenapiSchema

`func (o *EventSchema) GetOpenapiSchema() string`

GetOpenapiSchema returns the OpenapiSchema field if non-nil, zero value otherwise.

### GetOpenapiSchemaOk

`func (o *EventSchema) GetOpenapiSchemaOk() (*string, bool)`

GetOpenapiSchemaOk returns a tuple with the OpenapiSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenapiSchema

`func (o *EventSchema) SetOpenapiSchema(v string)`

SetOpenapiSchema sets OpenapiSchema field to given value.


### GetChangelog

`func (o *EventSchema) GetChangelog() string`

GetChangelog returns the Changelog field if non-nil, zero value otherwise.

### GetChangelogOk

`func (o *EventSchema) GetChangelogOk() (*string, bool)`

GetChangelogOk returns a tuple with the Changelog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelog

`func (o *EventSchema) SetChangelog(v string)`

SetChangelog sets Changelog field to given value.

### HasChangelog

`func (o *EventSchema) HasChangelog() bool`

HasChangelog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


