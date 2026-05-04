# \DriversUsersAPI

All URIs are relative to *https://api.catenatelematics.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroupMessage**](DriversUsersAPI.md#CreateGroupMessage) | **Post** /v2/telematics/group-messages | Create Group Message
[**CreateMessage**](DriversUsersAPI.md#CreateMessage) | **Post** /v2/telematics/messages | Create Message
[**CreateUser**](DriversUsersAPI.md#CreateUser) | **Post** /v2/telematics/users | Create User
[**GetUser**](DriversUsersAPI.md#GetUser) | **Get** /v2/telematics/users/{user_id} | Get User
[**ListGroupMessages**](DriversUsersAPI.md#ListGroupMessages) | **Get** /v2/telematics/group-messages | List Group Messages
[**ListMessages**](DriversUsersAPI.md#ListMessages) | **Get** /v2/telematics/messages | List Messages
[**ListUsers**](DriversUsersAPI.md#ListUsers) | **Get** /v2/telematics/users | List Users
[**UpdateUser**](DriversUsersAPI.md#UpdateUser) | **Patch** /v2/telematics/users/{id} | Update User



## CreateGroupMessage

> ResourceOperation CreateGroupMessage(ctx).GroupMessageCreate(groupMessageCreate).IsSync(isSync).Execute()

Create Group Message



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/telematics"
)

func main() {
	groupMessageCreate := *openapiclient.NewGroupMessageCreate("ConnectionId_example", "SenderId_example", []string{"RecipientIds_example"}) // GroupMessageCreate | 
	isSync := true // bool | Whether to process the request synchronously. If `true`, Catena will attempt to create or update the resource immediately with the TSP and return the result in the response. When performing synchronous operations, you are responsible for handling any necessary retries in case of transient failures.If `false` (default), Catena will create an asynchronous operation, and you can check the status of the operation using the returned operation ID. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriversUsersAPI.CreateGroupMessage(context.Background()).GroupMessageCreate(groupMessageCreate).IsSync(isSync).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriversUsersAPI.CreateGroupMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGroupMessage`: ResourceOperation
	fmt.Fprintf(os.Stdout, "Response from `DriversUsersAPI.CreateGroupMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupMessageCreate** | [**GroupMessageCreate**](GroupMessageCreate.md) |  | 
 **isSync** | **bool** | Whether to process the request synchronously. If &#x60;true&#x60;, Catena will attempt to create or update the resource immediately with the TSP and return the result in the response. When performing synchronous operations, you are responsible for handling any necessary retries in case of transient failures.If &#x60;false&#x60; (default), Catena will create an asynchronous operation, and you can check the status of the operation using the returned operation ID. | [default to false]

### Return type

[**ResourceOperation**](ResourceOperation.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMessage

> ResourceOperation CreateMessage(ctx).MessageCreate(messageCreate).IsSync(isSync).Execute()

Create Message



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/telematics"
)

func main() {
	messageCreate := *openapiclient.NewMessageCreate("ConnectionId_example", "SenderId_example", "RecipientId_example") // MessageCreate | 
	isSync := true // bool | Whether to process the request synchronously. If `true`, Catena will attempt to create or update the resource immediately with the TSP and return the result in the response. When performing synchronous operations, you are responsible for handling any necessary retries in case of transient failures.If `false` (default), Catena will create an asynchronous operation, and you can check the status of the operation using the returned operation ID. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriversUsersAPI.CreateMessage(context.Background()).MessageCreate(messageCreate).IsSync(isSync).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriversUsersAPI.CreateMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMessage`: ResourceOperation
	fmt.Fprintf(os.Stdout, "Response from `DriversUsersAPI.CreateMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **messageCreate** | [**MessageCreate**](MessageCreate.md) |  | 
 **isSync** | **bool** | Whether to process the request synchronously. If &#x60;true&#x60;, Catena will attempt to create or update the resource immediately with the TSP and return the result in the response. When performing synchronous operations, you are responsible for handling any necessary retries in case of transient failures.If &#x60;false&#x60; (default), Catena will create an asynchronous operation, and you can check the status of the operation using the returned operation ID. | [default to false]

### Return type

[**ResourceOperation**](ResourceOperation.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> ResourceOperation CreateUser(ctx).UserCreate(userCreate).IsSync(isSync).Execute()

Create User



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/telematics"
)

func main() {
	userCreate := *openapiclient.NewUserCreate("ConnectionId_example") // UserCreate | 
	isSync := true // bool | Whether to process the request synchronously. If `true`, Catena will attempt to create or update the resource immediately with the TSP and return the result in the response. When performing synchronous operations, you are responsible for handling any necessary retries in case of transient failures.If `false` (default), Catena will create an asynchronous operation, and you can check the status of the operation using the returned operation ID. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriversUsersAPI.CreateUser(context.Background()).UserCreate(userCreate).IsSync(isSync).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriversUsersAPI.CreateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUser`: ResourceOperation
	fmt.Fprintf(os.Stdout, "Response from `DriversUsersAPI.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userCreate** | [**UserCreate**](UserCreate.md) |  | 
 **isSync** | **bool** | Whether to process the request synchronously. If &#x60;true&#x60;, Catena will attempt to create or update the resource immediately with the TSP and return the result in the response. When performing synchronous operations, you are responsible for handling any necessary retries in case of transient failures.If &#x60;false&#x60; (default), Catena will create an asynchronous operation, and you can check the status of the operation using the returned operation ID. | [default to false]

### Return type

[**ResourceOperation**](ResourceOperation.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> UserRead GetUser(ctx, userId).IncludeSourceData(includeSourceData).Execute()

Get User



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/telematics"
)

func main() {
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique identifier of the user.
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriversUsersAPI.GetUser(context.Background(), userId).IncludeSourceData(includeSourceData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriversUsersAPI.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: UserRead
	fmt.Fprintf(os.Stdout, "Response from `DriversUsersAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]

### Return type

[**UserRead**](UserRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroupMessages

> CursorPageGroupMessageRead ListGroupMessages(ctx).FleetIds(fleetIds).FleetRefs(fleetRefs).ConnectionId(connectionId).FromDatetime(fromDatetime).ToDatetime(toDatetime).IncludeSourceData(includeSourceData).SenderIds(senderIds).GroupIds(groupIds).SourceSenderIds(sourceSenderIds).SourceGroupIds(sourceGroupIds).SortBy(sortBy).SortOrder(sortOrder).Cursor(cursor).Size(size).Execute()

List Group Messages



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/telematics"
)

func main() {
	fleetIds := []string{"Inner_example"} // []string | Limit results to specific fleets using Catena's fleet IDs. *For your own fleet identifiers, use `fleet_refs` instead* To specify multiple values, repeat the parameter for each value (e.g., `?fleet_ids=id1&fleet_ids=id2`). (optional)
	fleetRefs := []string{"Inner_example"} // []string | Limit results to specific fleets using your organization's fleet reference identifiers. To specify multiple values, repeat the parameter for each value (e.g., `?fleet_refs=ref1&fleet_refs=ref2`). (optional)
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Limit results to a specific provider connection. This is the UUID assigned by Catena when your fleet connects to a TSP. (optional)
	fromDatetime := time.Now() // time.Time | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at >= from_datetime` **Default value:** `now() - 1 day` **Restriction:** `to_datetime - from_datetime` cannot exceed 45 days (optional)
	toDatetime := time.Now() // time.Time | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at < to_datetime` **Default value:** `now()` **Restriction:** `to_datetime - from_datetime` cannot exceed 45 days (optional)
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)
	senderIds := []string{"Inner_example"} // []string | Limit results to specific senders. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?sender_ids=id1&sender_ids=id2`). (optional)
	groupIds := []string{"Inner_example"} // []string | Limit results to specific group IDs. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?group_ids=id1&group_ids=id2`). (optional)
	sourceSenderIds := []string{"Inner_example"} // []string | Limit results to specific source sender IDs. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?source_sender_ids=id1&source_sender_ids=id2`). (optional)
	sourceGroupIds := []string{"Inner_example"} // []string | Limit results to specific source group IDs. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?source_group_ids=id1&source_group_ids=id2`). (optional)
	sortBy := "sortBy_example" // string | The name of the field to sort results by. If not provided, results will be ordered by `occurred_at`.  (optional)
	sortOrder := "sortOrder_example" // string | The order of sorting, either `asc` for ascending or `desc` for descending. Defaults to `asc` if `sort_by` is provided without `sort_order`. (optional) (default to "asc")
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriversUsersAPI.ListGroupMessages(context.Background()).FleetIds(fleetIds).FleetRefs(fleetRefs).ConnectionId(connectionId).FromDatetime(fromDatetime).ToDatetime(toDatetime).IncludeSourceData(includeSourceData).SenderIds(senderIds).GroupIds(groupIds).SourceSenderIds(sourceSenderIds).SourceGroupIds(sourceGroupIds).SortBy(sortBy).SortOrder(sortOrder).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriversUsersAPI.ListGroupMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGroupMessages`: CursorPageGroupMessageRead
	fmt.Fprintf(os.Stdout, "Response from `DriversUsersAPI.ListGroupMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGroupMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetIds** | **[]string** | Limit results to specific fleets using Catena&#39;s fleet IDs. *For your own fleet identifiers, use &#x60;fleet_refs&#x60; instead* To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_ids&#x3D;id1&amp;fleet_ids&#x3D;id2&#x60;). | 
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers. To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_refs&#x3D;ref1&amp;fleet_refs&#x3D;ref2&#x60;). | 
 **connectionId** | **string** | Limit results to a specific provider connection. This is the UUID assigned by Catena when your fleet connects to a TSP. | 
 **fromDatetime** | **time.Time** | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &gt;&#x3D; from_datetime&#x60; **Default value:** &#x60;now() - 1 day&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 45 days | 
 **toDatetime** | **time.Time** | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &lt; to_datetime&#x60; **Default value:** &#x60;now()&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 45 days | 
 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]
 **senderIds** | **[]string** | Limit results to specific senders. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?sender_ids&#x3D;id1&amp;sender_ids&#x3D;id2&#x60;). | 
 **groupIds** | **[]string** | Limit results to specific group IDs. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?group_ids&#x3D;id1&amp;group_ids&#x3D;id2&#x60;). | 
 **sourceSenderIds** | **[]string** | Limit results to specific source sender IDs. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?source_sender_ids&#x3D;id1&amp;source_sender_ids&#x3D;id2&#x60;). | 
 **sourceGroupIds** | **[]string** | Limit results to specific source group IDs. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?source_group_ids&#x3D;id1&amp;source_group_ids&#x3D;id2&#x60;). | 
 **sortBy** | **string** | The name of the field to sort results by. If not provided, results will be ordered by &#x60;occurred_at&#x60;.  | 
 **sortOrder** | **string** | The order of sorting, either &#x60;asc&#x60; for ascending or &#x60;desc&#x60; for descending. Defaults to &#x60;asc&#x60; if &#x60;sort_by&#x60; is provided without &#x60;sort_order&#x60;. | [default to &quot;asc&quot;]
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageGroupMessageRead**](CursorPageGroupMessageRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMessages

> CursorPageMessageRead ListMessages(ctx).FleetIds(fleetIds).FleetRefs(fleetRefs).ConnectionId(connectionId).FromDatetime(fromDatetime).ToDatetime(toDatetime).IncludeSourceData(includeSourceData).SenderIds(senderIds).RecipientIds(recipientIds).SourceSenderIds(sourceSenderIds).SourceRecipientIds(sourceRecipientIds).SortBy(sortBy).SortOrder(sortOrder).Cursor(cursor).Size(size).Execute()

List Messages



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/telematics"
)

func main() {
	fleetIds := []string{"Inner_example"} // []string | Limit results to specific fleets using Catena's fleet IDs. *For your own fleet identifiers, use `fleet_refs` instead* To specify multiple values, repeat the parameter for each value (e.g., `?fleet_ids=id1&fleet_ids=id2`). (optional)
	fleetRefs := []string{"Inner_example"} // []string | Limit results to specific fleets using your organization's fleet reference identifiers. To specify multiple values, repeat the parameter for each value (e.g., `?fleet_refs=ref1&fleet_refs=ref2`). (optional)
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Limit results to a specific provider connection. This is the UUID assigned by Catena when your fleet connects to a TSP. (optional)
	fromDatetime := time.Now() // time.Time | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at >= from_datetime` **Default value:** `now() - 1 day` **Restriction:** `to_datetime - from_datetime` cannot exceed 45 days (optional)
	toDatetime := time.Now() // time.Time | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at < to_datetime` **Default value:** `now()` **Restriction:** `to_datetime - from_datetime` cannot exceed 45 days (optional)
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)
	senderIds := []string{"Inner_example"} // []string | Limit results to specific senders. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?sender_ids=id1&sender_ids=id2`). (optional)
	recipientIds := []string{"Inner_example"} // []string | Limit results to specific recipients. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?recipient_ids=id1&recipient_ids=id2`). (optional)
	sourceSenderIds := []string{"Inner_example"} // []string | Limit results to specific source sender IDs. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?source_sender_ids=id1&source_sender_ids=id2`). (optional)
	sourceRecipientIds := []string{"Inner_example"} // []string | Limit results to specific source recipient IDs. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?source_recipient_ids=id1&source_recipient_ids=id2`). (optional)
	sortBy := "sortBy_example" // string | The name of the field to sort results by. If not provided, results will be ordered by `occurred_at`.  (optional)
	sortOrder := "sortOrder_example" // string | The order of sorting, either `asc` for ascending or `desc` for descending. Defaults to `asc` if `sort_by` is provided without `sort_order`. (optional) (default to "asc")
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriversUsersAPI.ListMessages(context.Background()).FleetIds(fleetIds).FleetRefs(fleetRefs).ConnectionId(connectionId).FromDatetime(fromDatetime).ToDatetime(toDatetime).IncludeSourceData(includeSourceData).SenderIds(senderIds).RecipientIds(recipientIds).SourceSenderIds(sourceSenderIds).SourceRecipientIds(sourceRecipientIds).SortBy(sortBy).SortOrder(sortOrder).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriversUsersAPI.ListMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMessages`: CursorPageMessageRead
	fmt.Fprintf(os.Stdout, "Response from `DriversUsersAPI.ListMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetIds** | **[]string** | Limit results to specific fleets using Catena&#39;s fleet IDs. *For your own fleet identifiers, use &#x60;fleet_refs&#x60; instead* To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_ids&#x3D;id1&amp;fleet_ids&#x3D;id2&#x60;). | 
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers. To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_refs&#x3D;ref1&amp;fleet_refs&#x3D;ref2&#x60;). | 
 **connectionId** | **string** | Limit results to a specific provider connection. This is the UUID assigned by Catena when your fleet connects to a TSP. | 
 **fromDatetime** | **time.Time** | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &gt;&#x3D; from_datetime&#x60; **Default value:** &#x60;now() - 1 day&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 45 days | 
 **toDatetime** | **time.Time** | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &lt; to_datetime&#x60; **Default value:** &#x60;now()&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 45 days | 
 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]
 **senderIds** | **[]string** | Limit results to specific senders. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?sender_ids&#x3D;id1&amp;sender_ids&#x3D;id2&#x60;). | 
 **recipientIds** | **[]string** | Limit results to specific recipients. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?recipient_ids&#x3D;id1&amp;recipient_ids&#x3D;id2&#x60;). | 
 **sourceSenderIds** | **[]string** | Limit results to specific source sender IDs. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?source_sender_ids&#x3D;id1&amp;source_sender_ids&#x3D;id2&#x60;). | 
 **sourceRecipientIds** | **[]string** | Limit results to specific source recipient IDs. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?source_recipient_ids&#x3D;id1&amp;source_recipient_ids&#x3D;id2&#x60;). | 
 **sortBy** | **string** | The name of the field to sort results by. If not provided, results will be ordered by &#x60;occurred_at&#x60;.  | 
 **sortOrder** | **string** | The order of sorting, either &#x60;asc&#x60; for ascending or &#x60;desc&#x60; for descending. Defaults to &#x60;asc&#x60; if &#x60;sort_by&#x60; is provided without &#x60;sort_order&#x60;. | [default to &quot;asc&quot;]
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageMessageRead**](CursorPageMessageRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers

> CursorPageUserRead ListUsers(ctx).IsDriver(isDriver).FleetIds(fleetIds).FleetRefs(fleetRefs).ConnectionId(connectionId).IncludeSourceData(includeSourceData).UserIds(userIds).SourceIds(sourceIds).SortBy(sortBy).SortOrder(sortOrder).Cursor(cursor).Size(size).Execute()

List Users



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/telematics"
)

func main() {
	isDriver := true // bool | Filter users by whether they are drivers or not. (optional)
	fleetIds := []string{"Inner_example"} // []string | Limit results to specific fleets using Catena's fleet IDs. *For your own fleet identifiers, use `fleet_refs` instead* To specify multiple values, repeat the parameter for each value (e.g., `?fleet_ids=id1&fleet_ids=id2`). (optional)
	fleetRefs := []string{"Inner_example"} // []string | Limit results to specific fleets using your organization's fleet reference identifiers. To specify multiple values, repeat the parameter for each value (e.g., `?fleet_refs=ref1&fleet_refs=ref2`). (optional)
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Limit results to a specific provider connection. This is the UUID assigned by Catena when your fleet connects to a TSP. (optional)
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)
	userIds := []string{"Inner_example"} // []string | Limit results to specific users. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?user_ids=id1&user_ids=id2`). (optional)
	sourceIds := []string{"Inner_example"} // []string | Limit results to specific resources using the identifier provided by the TSP. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?source_ids=id1&source_ids=id2`). (optional)
	sortBy := "sortBy_example" // string | The name of the field to sort results by. If not provided, results will be ordered by `occurred_at`.  (optional)
	sortOrder := "sortOrder_example" // string | The order of sorting, either `asc` for ascending or `desc` for descending. Defaults to `asc` if `sort_by` is provided without `sort_order`. (optional) (default to "asc")
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriversUsersAPI.ListUsers(context.Background()).IsDriver(isDriver).FleetIds(fleetIds).FleetRefs(fleetRefs).ConnectionId(connectionId).IncludeSourceData(includeSourceData).UserIds(userIds).SourceIds(sourceIds).SortBy(sortBy).SortOrder(sortOrder).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriversUsersAPI.ListUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUsers`: CursorPageUserRead
	fmt.Fprintf(os.Stdout, "Response from `DriversUsersAPI.ListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isDriver** | **bool** | Filter users by whether they are drivers or not. | 
 **fleetIds** | **[]string** | Limit results to specific fleets using Catena&#39;s fleet IDs. *For your own fleet identifiers, use &#x60;fleet_refs&#x60; instead* To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_ids&#x3D;id1&amp;fleet_ids&#x3D;id2&#x60;). | 
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers. To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_refs&#x3D;ref1&amp;fleet_refs&#x3D;ref2&#x60;). | 
 **connectionId** | **string** | Limit results to a specific provider connection. This is the UUID assigned by Catena when your fleet connects to a TSP. | 
 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]
 **userIds** | **[]string** | Limit results to specific users. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?user_ids&#x3D;id1&amp;user_ids&#x3D;id2&#x60;). | 
 **sourceIds** | **[]string** | Limit results to specific resources using the identifier provided by the TSP. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?source_ids&#x3D;id1&amp;source_ids&#x3D;id2&#x60;). | 
 **sortBy** | **string** | The name of the field to sort results by. If not provided, results will be ordered by &#x60;occurred_at&#x60;.  | 
 **sortOrder** | **string** | The order of sorting, either &#x60;asc&#x60; for ascending or &#x60;desc&#x60; for descending. Defaults to &#x60;asc&#x60; if &#x60;sort_by&#x60; is provided without &#x60;sort_order&#x60;. | [default to &quot;asc&quot;]
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageUserRead**](CursorPageUserRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> ResourceOperation UpdateUser(ctx, id).UserUpdate(userUpdate).IsSync(isSync).Execute()

Update User



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/telematics"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique identifier of the user
	userUpdate := *openapiclient.NewUserUpdate("ConnectionId_example") // UserUpdate | 
	isSync := true // bool | Whether to process the request synchronously. If `true`, Catena will attempt to create or update the resource immediately with the TSP and return the result in the response. When performing synchronous operations, you are responsible for handling any necessary retries in case of transient failures.If `false` (default), Catena will create an asynchronous operation, and you can check the status of the operation using the returned operation ID. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriversUsersAPI.UpdateUser(context.Background(), id).UserUpdate(userUpdate).IsSync(isSync).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriversUsersAPI.UpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUser`: ResourceOperation
	fmt.Fprintf(os.Stdout, "Response from `DriversUsersAPI.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userUpdate** | [**UserUpdate**](UserUpdate.md) |  | 
 **isSync** | **bool** | Whether to process the request synchronously. If &#x60;true&#x60;, Catena will attempt to create or update the resource immediately with the TSP and return the result in the response. When performing synchronous operations, you are responsible for handling any necessary retries in case of transient failures.If &#x60;false&#x60; (default), Catena will create an asynchronous operation, and you can check the status of the operation using the returned operation ID. | [default to false]

### Return type

[**ResourceOperation**](ResourceOperation.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

