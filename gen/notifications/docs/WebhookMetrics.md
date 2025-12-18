# WebhookMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookId** | **string** | The ID of the webhook that the metrics are for | 
**HttpAttempts** | [**TimePeriodMetric**](TimePeriodMetric.md) | The total number of HTTP requests made (including retries) | 
**HttpSuccessAttempts** | [**TimePeriodMetric**](TimePeriodMetric.md) | The number of HTTP requests that received a 2xx response | 
**HttpFailureAttempts** | [**TimePeriodMetric**](TimePeriodMetric.md) | The number of HTTP requests that failed or received a non-2xx response | 
**MessageCount** | [**TimePeriodMetric**](TimePeriodMetric.md) | The total number of unique messages/events processed (regardless of retries) | 
**MessageSuccessCount** | [**TimePeriodMetric**](TimePeriodMetric.md) | The number of messages that were successfully delivered (after all retries) | 
**SuccessRate** | [**TimePeriodMetric**](TimePeriodMetric.md) | The success rate of the webhook. Only available if there are at least 200 messages | 
**AvgResponseTimeMs** | [**TimePeriodMetric**](TimePeriodMetric.md) | The average response time of the webhook | 
**EwmaSuccessRate** | **NullableFloat32** |  | 
**DlqCount** | **int32** | The number of messages that failed to be delivered, and were put in the DLQ. Messages are stored for 14 days and must be replayed manually. | 

## Methods

### NewWebhookMetrics

`func NewWebhookMetrics(webhookId string, httpAttempts TimePeriodMetric, httpSuccessAttempts TimePeriodMetric, httpFailureAttempts TimePeriodMetric, messageCount TimePeriodMetric, messageSuccessCount TimePeriodMetric, successRate TimePeriodMetric, avgResponseTimeMs TimePeriodMetric, ewmaSuccessRate NullableFloat32, dlqCount int32, ) *WebhookMetrics`

NewWebhookMetrics instantiates a new WebhookMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookMetricsWithDefaults

`func NewWebhookMetricsWithDefaults() *WebhookMetrics`

NewWebhookMetricsWithDefaults instantiates a new WebhookMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookId

`func (o *WebhookMetrics) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebhookMetrics) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *WebhookMetrics) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.


### GetHttpAttempts

`func (o *WebhookMetrics) GetHttpAttempts() TimePeriodMetric`

GetHttpAttempts returns the HttpAttempts field if non-nil, zero value otherwise.

### GetHttpAttemptsOk

`func (o *WebhookMetrics) GetHttpAttemptsOk() (*TimePeriodMetric, bool)`

GetHttpAttemptsOk returns a tuple with the HttpAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpAttempts

`func (o *WebhookMetrics) SetHttpAttempts(v TimePeriodMetric)`

SetHttpAttempts sets HttpAttempts field to given value.


### GetHttpSuccessAttempts

`func (o *WebhookMetrics) GetHttpSuccessAttempts() TimePeriodMetric`

GetHttpSuccessAttempts returns the HttpSuccessAttempts field if non-nil, zero value otherwise.

### GetHttpSuccessAttemptsOk

`func (o *WebhookMetrics) GetHttpSuccessAttemptsOk() (*TimePeriodMetric, bool)`

GetHttpSuccessAttemptsOk returns a tuple with the HttpSuccessAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpSuccessAttempts

`func (o *WebhookMetrics) SetHttpSuccessAttempts(v TimePeriodMetric)`

SetHttpSuccessAttempts sets HttpSuccessAttempts field to given value.


### GetHttpFailureAttempts

`func (o *WebhookMetrics) GetHttpFailureAttempts() TimePeriodMetric`

GetHttpFailureAttempts returns the HttpFailureAttempts field if non-nil, zero value otherwise.

### GetHttpFailureAttemptsOk

`func (o *WebhookMetrics) GetHttpFailureAttemptsOk() (*TimePeriodMetric, bool)`

GetHttpFailureAttemptsOk returns a tuple with the HttpFailureAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpFailureAttempts

`func (o *WebhookMetrics) SetHttpFailureAttempts(v TimePeriodMetric)`

SetHttpFailureAttempts sets HttpFailureAttempts field to given value.


### GetMessageCount

`func (o *WebhookMetrics) GetMessageCount() TimePeriodMetric`

GetMessageCount returns the MessageCount field if non-nil, zero value otherwise.

### GetMessageCountOk

`func (o *WebhookMetrics) GetMessageCountOk() (*TimePeriodMetric, bool)`

GetMessageCountOk returns a tuple with the MessageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCount

`func (o *WebhookMetrics) SetMessageCount(v TimePeriodMetric)`

SetMessageCount sets MessageCount field to given value.


### GetMessageSuccessCount

`func (o *WebhookMetrics) GetMessageSuccessCount() TimePeriodMetric`

GetMessageSuccessCount returns the MessageSuccessCount field if non-nil, zero value otherwise.

### GetMessageSuccessCountOk

`func (o *WebhookMetrics) GetMessageSuccessCountOk() (*TimePeriodMetric, bool)`

GetMessageSuccessCountOk returns a tuple with the MessageSuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSuccessCount

`func (o *WebhookMetrics) SetMessageSuccessCount(v TimePeriodMetric)`

SetMessageSuccessCount sets MessageSuccessCount field to given value.


### GetSuccessRate

`func (o *WebhookMetrics) GetSuccessRate() TimePeriodMetric`

GetSuccessRate returns the SuccessRate field if non-nil, zero value otherwise.

### GetSuccessRateOk

`func (o *WebhookMetrics) GetSuccessRateOk() (*TimePeriodMetric, bool)`

GetSuccessRateOk returns a tuple with the SuccessRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRate

`func (o *WebhookMetrics) SetSuccessRate(v TimePeriodMetric)`

SetSuccessRate sets SuccessRate field to given value.


### GetAvgResponseTimeMs

`func (o *WebhookMetrics) GetAvgResponseTimeMs() TimePeriodMetric`

GetAvgResponseTimeMs returns the AvgResponseTimeMs field if non-nil, zero value otherwise.

### GetAvgResponseTimeMsOk

`func (o *WebhookMetrics) GetAvgResponseTimeMsOk() (*TimePeriodMetric, bool)`

GetAvgResponseTimeMsOk returns a tuple with the AvgResponseTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgResponseTimeMs

`func (o *WebhookMetrics) SetAvgResponseTimeMs(v TimePeriodMetric)`

SetAvgResponseTimeMs sets AvgResponseTimeMs field to given value.


### GetEwmaSuccessRate

`func (o *WebhookMetrics) GetEwmaSuccessRate() float32`

GetEwmaSuccessRate returns the EwmaSuccessRate field if non-nil, zero value otherwise.

### GetEwmaSuccessRateOk

`func (o *WebhookMetrics) GetEwmaSuccessRateOk() (*float32, bool)`

GetEwmaSuccessRateOk returns a tuple with the EwmaSuccessRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEwmaSuccessRate

`func (o *WebhookMetrics) SetEwmaSuccessRate(v float32)`

SetEwmaSuccessRate sets EwmaSuccessRate field to given value.


### SetEwmaSuccessRateNil

`func (o *WebhookMetrics) SetEwmaSuccessRateNil(b bool)`

 SetEwmaSuccessRateNil sets the value for EwmaSuccessRate to be an explicit nil

### UnsetEwmaSuccessRate
`func (o *WebhookMetrics) UnsetEwmaSuccessRate()`

UnsetEwmaSuccessRate ensures that no value is present for EwmaSuccessRate, not even an explicit nil
### GetDlqCount

`func (o *WebhookMetrics) GetDlqCount() int32`

GetDlqCount returns the DlqCount field if non-nil, zero value otherwise.

### GetDlqCountOk

`func (o *WebhookMetrics) GetDlqCountOk() (*int32, bool)`

GetDlqCountOk returns a tuple with the DlqCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlqCount

`func (o *WebhookMetrics) SetDlqCount(v int32)`

SetDlqCount sets DlqCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


