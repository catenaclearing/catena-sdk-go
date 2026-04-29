# \WebhookSubscriptionsAPI

All URIs are relative to *https://api.catenatelematics.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateWebhookSubscription**](WebhookSubscriptionsAPI.md#ActivateWebhookSubscription) | **Post** /v2/notifications/webhooks/{webhook_id}/activate | Activate Webhook Subscription
[**CreateWebhookSubscription**](WebhookSubscriptionsAPI.md#CreateWebhookSubscription) | **Post** /v2/notifications/webhooks | Create Webhook Subscription
[**DeleteWebhookSubscription**](WebhookSubscriptionsAPI.md#DeleteWebhookSubscription) | **Delete** /v2/notifications/webhooks/{webhook_id} | Delete Webhook Subscription
[**GetWebhookSubscription**](WebhookSubscriptionsAPI.md#GetWebhookSubscription) | **Get** /v2/notifications/webhooks/{webhook_id} | Get Webhook Subscription
[**GetWebhookSubscriptionLogs**](WebhookSubscriptionsAPI.md#GetWebhookSubscriptionLogs) | **Get** /v2/notifications/webhooks/{webhook_id}/logs | Get Webhook Subscription Logs
[**GetWebhookSubscriptionMetrics**](WebhookSubscriptionsAPI.md#GetWebhookSubscriptionMetrics) | **Get** /v2/notifications/webhooks/{webhook_id}/metrics | Get Webhook Subscription Metrics
[**ListWebhookSubscriptions**](WebhookSubscriptionsAPI.md#ListWebhookSubscriptions) | **Get** /v2/notifications/webhooks | List Webhook Subscriptions
[**PauseWebhookSubscription**](WebhookSubscriptionsAPI.md#PauseWebhookSubscription) | **Post** /v2/notifications/webhooks/{webhook_id}/pause | Pause Webhook Subscription
[**ReplayUndeliveredMessagesFromDlq**](WebhookSubscriptionsAPI.md#ReplayUndeliveredMessagesFromDlq) | **Post** /v2/notifications/webhooks/{webhook_id}/replay | Replay Undelivered Messages From Dlq
[**UpdateWebhookSubscription**](WebhookSubscriptionsAPI.md#UpdateWebhookSubscription) | **Patch** /v2/notifications/webhooks/{webhook_id} | Update Webhook Subscription



## ActivateWebhookSubscription

> WebhookRead ActivateWebhookSubscription(ctx, webhookId).Execute()

Activate Webhook Subscription



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/notifications"
)

func main() {
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookSubscriptionsAPI.ActivateWebhookSubscription(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookSubscriptionsAPI.ActivateWebhookSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivateWebhookSubscription`: WebhookRead
	fmt.Fprintf(os.Stdout, "Response from `WebhookSubscriptionsAPI.ActivateWebhookSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateWebhookSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookRead**](WebhookRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWebhookSubscription

> ResponseCreateWebhookSubscription CreateWebhookSubscription(ctx).WebhookCreate(webhookCreate).Execute()

Create Webhook Subscription



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/notifications"
)

func main() {
	webhookCreate := *openapiclient.NewWebhookCreate("Url_example", *openapiclient.NewWebhookEventNameUnion()) // WebhookCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookSubscriptionsAPI.CreateWebhookSubscription(context.Background()).WebhookCreate(webhookCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookSubscriptionsAPI.CreateWebhookSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWebhookSubscription`: ResponseCreateWebhookSubscription
	fmt.Fprintf(os.Stdout, "Response from `WebhookSubscriptionsAPI.CreateWebhookSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookCreate** | [**WebhookCreate**](WebhookCreate.md) |  | 

### Return type

[**ResponseCreateWebhookSubscription**](ResponseCreateWebhookSubscription.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhookSubscription

> DeleteWebhookSubscription(ctx, webhookId).Execute()

Delete Webhook Subscription



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/notifications"
)

func main() {
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookSubscriptionsAPI.DeleteWebhookSubscription(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookSubscriptionsAPI.DeleteWebhookSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookSubscription

> WebhookRead GetWebhookSubscription(ctx, webhookId).Execute()

Get Webhook Subscription



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/notifications"
)

func main() {
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookSubscriptionsAPI.GetWebhookSubscription(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookSubscriptionsAPI.GetWebhookSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookSubscription`: WebhookRead
	fmt.Fprintf(os.Stdout, "Response from `WebhookSubscriptionsAPI.GetWebhookSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookRead**](WebhookRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookSubscriptionLogs

> CursorPageTypeVarCustomizedWebhookLogRead GetWebhookSubscriptionLogs(ctx, webhookId).StartDate(startDate).EndDate(endDate).Status(status).Cursor(cursor).Size(size).Execute()

Get Webhook Subscription Logs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/notifications"
)

func main() {
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	startDate := time.Now() // time.Time |  (optional)
	endDate := time.Now() // time.Time |  (optional)
	status := "status_example" // string |  (optional)
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookSubscriptionsAPI.GetWebhookSubscriptionLogs(context.Background(), webhookId).StartDate(startDate).EndDate(endDate).Status(status).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookSubscriptionsAPI.GetWebhookSubscriptionLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookSubscriptionLogs`: CursorPageTypeVarCustomizedWebhookLogRead
	fmt.Fprintf(os.Stdout, "Response from `WebhookSubscriptionsAPI.GetWebhookSubscriptionLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookSubscriptionLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **time.Time** |  | 
 **endDate** | **time.Time** |  | 
 **status** | **string** |  | 
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageTypeVarCustomizedWebhookLogRead**](CursorPageTypeVarCustomizedWebhookLogRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookSubscriptionMetrics

> WebhookMetrics GetWebhookSubscriptionMetrics(ctx, webhookId).Execute()

Get Webhook Subscription Metrics



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/notifications"
)

func main() {
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookSubscriptionsAPI.GetWebhookSubscriptionMetrics(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookSubscriptionsAPI.GetWebhookSubscriptionMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookSubscriptionMetrics`: WebhookMetrics
	fmt.Fprintf(os.Stdout, "Response from `WebhookSubscriptionsAPI.GetWebhookSubscriptionMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookSubscriptionMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookMetrics**](WebhookMetrics.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhookSubscriptions

> CursorPageTypeVarCustomizedWebhookRead ListWebhookSubscriptions(ctx).EventName(eventName).Status(status).Cursor(cursor).Size(size).Execute()

List Webhook Subscriptions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/notifications"
)

func main() {
	eventName := openapiclient.WebhookEventNameEnum("driver_vehicle_association.added") // WebhookEventNameEnum |  (optional)
	status := openapiclient.WebhookStatusEnum("active") // WebhookStatusEnum |  (optional)
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookSubscriptionsAPI.ListWebhookSubscriptions(context.Background()).EventName(eventName).Status(status).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookSubscriptionsAPI.ListWebhookSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWebhookSubscriptions`: CursorPageTypeVarCustomizedWebhookRead
	fmt.Fprintf(os.Stdout, "Response from `WebhookSubscriptionsAPI.ListWebhookSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWebhookSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventName** | [**WebhookEventNameEnum**](WebhookEventNameEnum.md) |  | 
 **status** | [**WebhookStatusEnum**](WebhookStatusEnum.md) |  | 
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageTypeVarCustomizedWebhookRead**](CursorPageTypeVarCustomizedWebhookRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseWebhookSubscription

> WebhookRead PauseWebhookSubscription(ctx, webhookId).Execute()

Pause Webhook Subscription



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/notifications"
)

func main() {
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookSubscriptionsAPI.PauseWebhookSubscription(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookSubscriptionsAPI.PauseWebhookSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PauseWebhookSubscription`: WebhookRead
	fmt.Fprintf(os.Stdout, "Response from `WebhookSubscriptionsAPI.PauseWebhookSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseWebhookSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookRead**](WebhookRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplayUndeliveredMessagesFromDlq

> interface{} ReplayUndeliveredMessagesFromDlq(ctx, webhookId).Execute()

Replay Undelivered Messages From Dlq



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/notifications"
)

func main() {
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookSubscriptionsAPI.ReplayUndeliveredMessagesFromDlq(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookSubscriptionsAPI.ReplayUndeliveredMessagesFromDlq``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplayUndeliveredMessagesFromDlq`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookSubscriptionsAPI.ReplayUndeliveredMessagesFromDlq`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplayUndeliveredMessagesFromDlqRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhookSubscription

> WebhookRead UpdateWebhookSubscription(ctx, webhookId).WebhookUpdate(webhookUpdate).Execute()

Update Webhook Subscription



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/notifications"
)

func main() {
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	webhookUpdate := *openapiclient.NewWebhookUpdate() // WebhookUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookSubscriptionsAPI.UpdateWebhookSubscription(context.Background(), webhookId).WebhookUpdate(webhookUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookSubscriptionsAPI.UpdateWebhookSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebhookSubscription`: WebhookRead
	fmt.Fprintf(os.Stdout, "Response from `WebhookSubscriptionsAPI.UpdateWebhookSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webhookUpdate** | [**WebhookUpdate**](WebhookUpdate.md) |  | 

### Return type

[**WebhookRead**](WebhookRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

