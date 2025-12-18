# WebhookCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The target URL of the webhook | 
**Filters** | Pointer to [**NullableWebhookFilters**](WebhookFilters.md) |  | [optional] 
**Secret** | Pointer to **NullableString** |  | [optional] 
**EventName** | [**WebhookEventNameUnion**](WebhookEventNameUnion.md) |  | 

## Methods

### NewWebhookCreate

`func NewWebhookCreate(url string, eventName WebhookEventNameUnion, ) *WebhookCreate`

NewWebhookCreate instantiates a new WebhookCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookCreateWithDefaults

`func NewWebhookCreateWithDefaults() *WebhookCreate`

NewWebhookCreateWithDefaults instantiates a new WebhookCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *WebhookCreate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookCreate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookCreate) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetFilters

`func (o *WebhookCreate) GetFilters() WebhookFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *WebhookCreate) GetFiltersOk() (*WebhookFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *WebhookCreate) SetFilters(v WebhookFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *WebhookCreate) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *WebhookCreate) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *WebhookCreate) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetSecret

`func (o *WebhookCreate) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *WebhookCreate) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *WebhookCreate) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *WebhookCreate) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### SetSecretNil

`func (o *WebhookCreate) SetSecretNil(b bool)`

 SetSecretNil sets the value for Secret to be an explicit nil

### UnsetSecret
`func (o *WebhookCreate) UnsetSecret()`

UnsetSecret ensures that no value is present for Secret, not even an explicit nil
### GetEventName

`func (o *WebhookCreate) GetEventName() WebhookEventNameUnion`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *WebhookCreate) GetEventNameOk() (*WebhookEventNameUnion, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *WebhookCreate) SetEventName(v WebhookEventNameUnion)`

SetEventName sets EventName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


