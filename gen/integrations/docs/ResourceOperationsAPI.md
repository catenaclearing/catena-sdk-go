# \ResourceOperationsAPI

All URIs are relative to *https://api.catenatelematics.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetResourceOperation**](ResourceOperationsAPI.md#GetResourceOperation) | **Get** /v2/integrations/connections/{connection_id}/resource-operations/{resource_operation_id} | Get Resource Operation
[**ListResourceOperations**](ResourceOperationsAPI.md#ListResourceOperations) | **Get** /v2/integrations/connections/resource-operations | List Resource Operations



## GetResourceOperation

> ResourceOperationRead GetResourceOperation(ctx, connectionId, resourceOperationId).Execute()

Get Resource Operation



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/integrations"
)

func main() {
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique identifier of the connection.
	resourceOperationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique identifier of the resource operation.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceOperationsAPI.GetResourceOperation(context.Background(), connectionId, resourceOperationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceOperationsAPI.GetResourceOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResourceOperation`: ResourceOperationRead
	fmt.Fprintf(os.Stdout, "Response from `ResourceOperationsAPI.GetResourceOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | The unique identifier of the connection. | 
**resourceOperationId** | **string** | The unique identifier of the resource operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResourceOperationRead**](ResourceOperationRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResourceOperations

> CursorPageTypeVarCustomizedResourceOperationRead ListResourceOperations(ctx).Resource(resource).ConnectionId(connectionId).SourceName(sourceName).FleetRefs(fleetRefs).Operation(operation).SourceId(sourceId).FromDatetime(fromDatetime).ToDatetime(toDatetime).Cursor(cursor).Size(size).Execute()

List Resource Operations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/integrations"
)

func main() {
	resource := openapiclient.ResourceEnum("asset") // ResourceEnum | Filter resource operations by resource type (e.g., `vehicle`, `user`, `fuel_charge`). (optional)
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter resource operations by connection ID. (optional)
	sourceName := openapiclient.TspEnum("ada") // TspEnum | Filter resource operations by source name. (optional)
	fleetRefs := []string{"Inner_example"} // []string | Filter resource operations by one or more fleet references. **Maximum:** 100 values. To specify multiple values, repeat the parameter for each value (e.g., `?fleet_refs=ref1&fleet_refs=ref2`). (optional)
	operation := openapiclient.ResourceOperationTypeEnum("create") // ResourceOperationTypeEnum | Filter resource operations by operation type (e.g., `create`, `update`). (optional)
	sourceId := "sourceId_example" // string | Filter resource operations by source ID in the TSP. Requires the `resource` parameter. (optional)
	fromDatetime := time.Now() // time.Time | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at >= from_datetime` **Default value:** `now() - 1 day` **Restriction:** `to_datetime - from_datetime` cannot exceed 45 days (optional)
	toDatetime := time.Now() // time.Time | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at < to_datetime` **Default value:** `now()` **Restriction:** `to_datetime - from_datetime` cannot exceed 45 days (optional)
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceOperationsAPI.ListResourceOperations(context.Background()).Resource(resource).ConnectionId(connectionId).SourceName(sourceName).FleetRefs(fleetRefs).Operation(operation).SourceId(sourceId).FromDatetime(fromDatetime).ToDatetime(toDatetime).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceOperationsAPI.ListResourceOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListResourceOperations`: CursorPageTypeVarCustomizedResourceOperationRead
	fmt.Fprintf(os.Stdout, "Response from `ResourceOperationsAPI.ListResourceOperations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListResourceOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resource** | [**ResourceEnum**](ResourceEnum.md) | Filter resource operations by resource type (e.g., &#x60;vehicle&#x60;, &#x60;user&#x60;, &#x60;fuel_charge&#x60;). | 
 **connectionId** | **string** | Filter resource operations by connection ID. | 
 **sourceName** | [**TspEnum**](TspEnum.md) | Filter resource operations by source name. | 
 **fleetRefs** | **[]string** | Filter resource operations by one or more fleet references. **Maximum:** 100 values. To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_refs&#x3D;ref1&amp;fleet_refs&#x3D;ref2&#x60;). | 
 **operation** | [**ResourceOperationTypeEnum**](ResourceOperationTypeEnum.md) | Filter resource operations by operation type (e.g., &#x60;create&#x60;, &#x60;update&#x60;). | 
 **sourceId** | **string** | Filter resource operations by source ID in the TSP. Requires the &#x60;resource&#x60; parameter. | 
 **fromDatetime** | **time.Time** | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &gt;&#x3D; from_datetime&#x60; **Default value:** &#x60;now() - 1 day&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 45 days | 
 **toDatetime** | **time.Time** | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &lt; to_datetime&#x60; **Default value:** &#x60;now()&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 45 days | 
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageTypeVarCustomizedResourceOperationRead**](CursorPageTypeVarCustomizedResourceOperationRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

