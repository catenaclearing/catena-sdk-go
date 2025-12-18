# \WebhookEventSchemasAPI

All URIs are relative to *https://api.catenatelematics.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEventSchema**](WebhookEventSchemasAPI.md#GetEventSchema) | **Get** /v2/notifications/schemas/{event_name} | Get Event Schema
[**GetEventSchemaVersions**](WebhookEventSchemasAPI.md#GetEventSchemaVersions) | **Get** /v2/notifications/schemas/{event_name}/versions/{version} | Get Event Schema Version



## GetEventSchema

> EventSchema GetEventSchema(ctx, eventName).Execute()

Get Event Schema



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/notificationsapi"
)

func main() {
	eventName := "eventName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventSchemasAPI.GetEventSchema(context.Background(), eventName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventSchemasAPI.GetEventSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventSchema`: EventSchema
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventSchemasAPI.GetEventSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventSchema**](EventSchema.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventSchemaVersions

> EventSchema GetEventSchemaVersions(ctx, eventName, version).Execute()

Get Event Schema Version



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/notificationsapi"
)

func main() {
	eventName := "eventName_example" // string | 
	version := "version_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventSchemasAPI.GetEventSchemaVersions(context.Background(), eventName, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventSchemasAPI.GetEventSchemaVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventSchemaVersions`: EventSchema
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventSchemasAPI.GetEventSchemaVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventName** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventSchemaVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EventSchema**](EventSchema.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

