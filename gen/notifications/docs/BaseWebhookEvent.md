# BaseWebhookEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the webhook | 
**Url** | **string** | The URL we will send the event to | 
**EventName** | [**WebhookEventNameEnum**](WebhookEventNameEnum.md) | The name of the event | 
**Filters** | Pointer to [**NullableWebhookFilters**](WebhookFilters.md) |  | [optional] 
**Status** | [**WebhookStatusEnum**](WebhookStatusEnum.md) | The status of the webhook | 

## Methods

### NewBaseWebhookEvent

`func NewBaseWebhookEvent(id string, url string, eventName WebhookEventNameEnum, status WebhookStatusEnum, ) *BaseWebhookEvent`

NewBaseWebhookEvent instantiates a new BaseWebhookEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseWebhookEventWithDefaults

`func NewBaseWebhookEventWithDefaults() *BaseWebhookEvent`

NewBaseWebhookEventWithDefaults instantiates a new BaseWebhookEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseWebhookEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseWebhookEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseWebhookEvent) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *BaseWebhookEvent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BaseWebhookEvent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BaseWebhookEvent) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetEventName

`func (o *BaseWebhookEvent) GetEventName() WebhookEventNameEnum`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *BaseWebhookEvent) GetEventNameOk() (*WebhookEventNameEnum, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *BaseWebhookEvent) SetEventName(v WebhookEventNameEnum)`

SetEventName sets EventName field to given value.


### GetFilters

`func (o *BaseWebhookEvent) GetFilters() WebhookFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *BaseWebhookEvent) GetFiltersOk() (*WebhookFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *BaseWebhookEvent) SetFilters(v WebhookFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *BaseWebhookEvent) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *BaseWebhookEvent) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *BaseWebhookEvent) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetStatus

`func (o *BaseWebhookEvent) GetStatus() WebhookStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaseWebhookEvent) GetStatusOk() (*WebhookStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaseWebhookEvent) SetStatus(v WebhookStatusEnum)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


