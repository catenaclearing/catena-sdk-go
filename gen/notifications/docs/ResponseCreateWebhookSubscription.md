# ResponseCreateWebhookSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The target URL of the webhook | 
**Filters** | Pointer to [**WebhookFilters**](WebhookFilters.md) |  | [optional] 
**Secret** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**PartnerId** | **string** | The ID of the partner that owns the webhook | 
**Status** | Pointer to [**WebhookStatusEnum**](WebhookStatusEnum.md) | The current status of the webhook | [optional] 
**EventName** | [**WebhookEventNameEnum**](WebhookEventNameEnum.md) | The event name that triggers the webhook | 

## Methods

### NewResponseCreateWebhookSubscription

`func NewResponseCreateWebhookSubscription(url string, partnerId string, eventName WebhookEventNameEnum, ) *ResponseCreateWebhookSubscription`

NewResponseCreateWebhookSubscription instantiates a new ResponseCreateWebhookSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseCreateWebhookSubscriptionWithDefaults

`func NewResponseCreateWebhookSubscriptionWithDefaults() *ResponseCreateWebhookSubscription`

NewResponseCreateWebhookSubscriptionWithDefaults instantiates a new ResponseCreateWebhookSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ResponseCreateWebhookSubscription) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ResponseCreateWebhookSubscription) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ResponseCreateWebhookSubscription) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetFilters

`func (o *ResponseCreateWebhookSubscription) GetFilters() WebhookFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ResponseCreateWebhookSubscription) GetFiltersOk() (*WebhookFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ResponseCreateWebhookSubscription) SetFilters(v WebhookFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ResponseCreateWebhookSubscription) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSecret

`func (o *ResponseCreateWebhookSubscription) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ResponseCreateWebhookSubscription) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ResponseCreateWebhookSubscription) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ResponseCreateWebhookSubscription) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ResponseCreateWebhookSubscription) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResponseCreateWebhookSubscription) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResponseCreateWebhookSubscription) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ResponseCreateWebhookSubscription) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ResponseCreateWebhookSubscription) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResponseCreateWebhookSubscription) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResponseCreateWebhookSubscription) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ResponseCreateWebhookSubscription) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ResponseCreateWebhookSubscription) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ResponseCreateWebhookSubscription) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ResponseCreateWebhookSubscription) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ResponseCreateWebhookSubscription) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetId

`func (o *ResponseCreateWebhookSubscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseCreateWebhookSubscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseCreateWebhookSubscription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResponseCreateWebhookSubscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPartnerId

`func (o *ResponseCreateWebhookSubscription) GetPartnerId() string`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *ResponseCreateWebhookSubscription) GetPartnerIdOk() (*string, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *ResponseCreateWebhookSubscription) SetPartnerId(v string)`

SetPartnerId sets PartnerId field to given value.


### GetStatus

`func (o *ResponseCreateWebhookSubscription) GetStatus() WebhookStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseCreateWebhookSubscription) GetStatusOk() (*WebhookStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseCreateWebhookSubscription) SetStatus(v WebhookStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponseCreateWebhookSubscription) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEventName

`func (o *ResponseCreateWebhookSubscription) GetEventName() WebhookEventNameEnum`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *ResponseCreateWebhookSubscription) GetEventNameOk() (*WebhookEventNameEnum, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *ResponseCreateWebhookSubscription) SetEventName(v WebhookEventNameEnum)`

SetEventName sets EventName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


