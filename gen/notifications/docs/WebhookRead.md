# WebhookRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the webhook subscription | 
**Url** | **interface{}** |  | 
**EventName** | [**WebhookEventNameEnum**](WebhookEventNameEnum.md) | The event name that triggers this webhook | 
**Filters** | Pointer to [**NullableWebhookFilters**](WebhookFilters.md) |  | [optional] 
**Secret** | **interface{}** |  | 
**Status** | [**WebhookStatusEnum**](WebhookStatusEnum.md) | The current status of the webhook subscription (active, inactive, stale, or deleted) | 
**CreatedAt** | **time.Time** | The timestamp when the webhook subscription was created | 
**UpdatedAt** | **time.Time** | The timestamp when the webhook subscription was last updated | 

## Methods

### NewWebhookRead

`func NewWebhookRead(id string, url interface{}, eventName WebhookEventNameEnum, secret interface{}, status WebhookStatusEnum, createdAt time.Time, updatedAt time.Time, ) *WebhookRead`

NewWebhookRead instantiates a new WebhookRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookReadWithDefaults

`func NewWebhookReadWithDefaults() *WebhookRead`

NewWebhookReadWithDefaults instantiates a new WebhookRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookRead) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *WebhookRead) GetUrl() interface{}`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookRead) GetUrlOk() (*interface{}, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookRead) SetUrl(v interface{})`

SetUrl sets Url field to given value.


### SetUrlNil

`func (o *WebhookRead) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *WebhookRead) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetEventName

`func (o *WebhookRead) GetEventName() WebhookEventNameEnum`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *WebhookRead) GetEventNameOk() (*WebhookEventNameEnum, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *WebhookRead) SetEventName(v WebhookEventNameEnum)`

SetEventName sets EventName field to given value.


### GetFilters

`func (o *WebhookRead) GetFilters() WebhookFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *WebhookRead) GetFiltersOk() (*WebhookFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *WebhookRead) SetFilters(v WebhookFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *WebhookRead) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *WebhookRead) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *WebhookRead) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetSecret

`func (o *WebhookRead) GetSecret() interface{}`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *WebhookRead) GetSecretOk() (*interface{}, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *WebhookRead) SetSecret(v interface{})`

SetSecret sets Secret field to given value.


### SetSecretNil

`func (o *WebhookRead) SetSecretNil(b bool)`

 SetSecretNil sets the value for Secret to be an explicit nil

### UnsetSecret
`func (o *WebhookRead) UnsetSecret()`

UnsetSecret ensures that no value is present for Secret, not even an explicit nil
### GetStatus

`func (o *WebhookRead) GetStatus() WebhookStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookRead) GetStatusOk() (*WebhookStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookRead) SetStatus(v WebhookStatusEnum)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *WebhookRead) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookRead) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookRead) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *WebhookRead) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WebhookRead) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WebhookRead) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


