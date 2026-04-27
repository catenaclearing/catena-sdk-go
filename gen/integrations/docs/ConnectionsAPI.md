# \ConnectionsAPI

All URIs are relative to *https://api.catenatelematics.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateSchedule**](ConnectionsAPI.md#ActivateSchedule) | **Post** /v2/integrations/connections/{connection_id}/schedules/{schedule_id}/activate | Activate Schedule
[**ActivateSchedules**](ConnectionsAPI.md#ActivateSchedules) | **Post** /v2/integrations/connections/{connection_id}/schedules/activate-schedules | Activate Schedules
[**BackfillSchedules**](ConnectionsAPI.md#BackfillSchedules) | **Post** /v2/integrations/connections/schedules/backfill | Backfill Schedules
[**CreateConnection**](ConnectionsAPI.md#CreateConnection) | **Post** /v2/integrations/connections | Create Connection
[**CreateConnectionPreAuth**](ConnectionsAPI.md#CreateConnectionPreAuth) | **Post** /v2/integrations/connection-pre-auths | Create Connection Pre Auth
[**CreateSchedule**](ConnectionsAPI.md#CreateSchedule) | **Post** /v2/integrations/connections/{connection_id}/schedules | Create Schedule
[**DeleteConnection**](ConnectionsAPI.md#DeleteConnection) | **Delete** /v2/integrations/connections/{connection_id} | Delete Connection
[**DeleteSchedule**](ConnectionsAPI.md#DeleteSchedule) | **Delete** /v2/integrations/connections/{connection_id}/schedules/{schedule_id} | Delete Schedule
[**GetConnection**](ConnectionsAPI.md#GetConnection) | **Get** /v2/integrations/connections/{connection_id} | Get Connection
[**GetDataFreshness**](ConnectionsAPI.md#GetDataFreshness) | **Get** /v2/integrations/connections/data-freshness | Get Data Freshness
[**GetSchedule**](ConnectionsAPI.md#GetSchedule) | **Get** /v2/integrations/connections/{connection_id}/schedules/{schedule_id} | Get Schedule
[**ListConnections**](ConnectionsAPI.md#ListConnections) | **Get** /v2/integrations/connections | List Connections
[**ListExecutions**](ConnectionsAPI.md#ListExecutions) | **Get** /v2/integrations/connections/{connection_id}/schedules/{schedule_id}/executions | List Executions
[**ListSchedules**](ConnectionsAPI.md#ListSchedules) | **Get** /v2/integrations/connections/{connection_id}/schedules | List Schedules
[**UpdateConnection**](ConnectionsAPI.md#UpdateConnection) | **Patch** /v2/integrations/connections/{connection_id} | Update Connection
[**UpdateSchedule**](ConnectionsAPI.md#UpdateSchedule) | **Patch** /v2/integrations/connections/{connection_id}/schedules/{schedule_id} | Update Schedule



## ActivateSchedule

> ScheduleRead ActivateSchedule(ctx, scheduleId).Execute()

Activate Schedule



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
	scheduleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.ActivateSchedule(context.Background(), scheduleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ActivateSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivateSchedule`: ScheduleRead
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.ActivateSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScheduleRead**](ScheduleRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivateSchedules

> []ScheduleRead ActivateSchedules(ctx, connectionId).Execute()

Activate Schedules



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
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.ActivateSchedules(context.Background(), connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ActivateSchedules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivateSchedules`: []ScheduleRead
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.ActivateSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ScheduleRead**](ScheduleRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackfillSchedules

> ScheduleBackfillResponse BackfillSchedules(ctx).ScheduleBackfillRequest(scheduleBackfillRequest).Execute()

Backfill Schedules



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
	scheduleBackfillRequest := *openapiclient.NewScheduleBackfillRequest(openapiclient.ResourceEnum("asset")) // ScheduleBackfillRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.BackfillSchedules(context.Background()).ScheduleBackfillRequest(scheduleBackfillRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.BackfillSchedules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackfillSchedules`: ScheduleBackfillResponse
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.BackfillSchedules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackfillSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scheduleBackfillRequest** | [**ScheduleBackfillRequest**](ScheduleBackfillRequest.md) |  | 

### Return type

[**ScheduleBackfillResponse**](ScheduleBackfillResponse.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConnection

> ConnectionRead CreateConnection(ctx).ConnectionCreate(connectionCreate).Execute()

Create Connection



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
	connectionCreate := *openapiclient.NewConnectionCreate("TspId_example", *openapiclient.NewCredentials("AccessToken_example", "AccountId_example", "Host_example", "Username_example", "Password_example", "ApiKey_example", "Url_example", "AuthCode_example", "Token_example", "RedirectUri_example", "ClientId_example", "ClientSecret_example", "CompanyId_example", openapiclient.DatabaseDriverEnum("postgresql+psycopg2"), int32(123), "Database_example", "CarrierIdentifier_example", "CarrierIdentifierType_example", "AppId_example", "AppKey_example", "ClientKey_example", "SecretKey_example", "ResourceOwnerId_example", "ResourceOwnerSecret_example", "SignatureMethod_example", "Realm_example", "TokenUrl_example", "Code_example", "PrivateKey_example", "ProviderToken_example", "DotNumber_example", "ConsumerKey_example", "AccessKey_example", "BucketName_example", "Region_example", "ApiId_example", "CarrierId_example", "RestUsername_example", "RestPassword_example")) // ConnectionCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.CreateConnection(context.Background()).ConnectionCreate(connectionCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.CreateConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConnection`: ConnectionRead
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.CreateConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectionCreate** | [**ConnectionCreate**](ConnectionCreate.md) |  | 

### Return type

[**ConnectionRead**](ConnectionRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConnectionPreAuth

> string CreateConnectionPreAuth(ctx).ConnectionPreAuthCreate(connectionPreAuthCreate).Execute()

Create Connection Pre Auth



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
	connectionPreAuthCreate := *openapiclient.NewConnectionPreAuthCreate("TspId_example", *openapiclient.NewApiBasicCredsInput("Username_example", "Password_example")) // ConnectionPreAuthCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.CreateConnectionPreAuth(context.Background()).ConnectionPreAuthCreate(connectionPreAuthCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.CreateConnectionPreAuth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConnectionPreAuth`: string
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.CreateConnectionPreAuth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectionPreAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectionPreAuthCreate** | [**ConnectionPreAuthCreate**](ConnectionPreAuthCreate.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSchedule

> ScheduleRead CreateSchedule(ctx, connectionId).ScheduleCreate(scheduleCreate).Execute()

Create Schedule



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
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	scheduleCreate := *openapiclient.NewScheduleCreate(openapiclient.ResourceEnum("asset")) // ScheduleCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.CreateSchedule(context.Background(), connectionId).ScheduleCreate(scheduleCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.CreateSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSchedule`: ScheduleRead
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.CreateSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scheduleCreate** | [**ScheduleCreate**](ScheduleCreate.md) |  | 

### Return type

[**ScheduleRead**](ScheduleRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnection

> DeleteConnection(ctx, connectionId).Execute()

Delete Connection



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
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectionsAPI.DeleteConnection(context.Background(), connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.DeleteConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectionRequest struct via the builder pattern


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


## DeleteSchedule

> DeleteSchedule(ctx, connectionId, scheduleId).Execute()

Delete Schedule



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
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	scheduleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectionsAPI.DeleteSchedule(context.Background(), connectionId, scheduleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.DeleteSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** |  | 
**scheduleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScheduleRequest struct via the builder pattern


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


## GetConnection

> ConnectionRead GetConnection(ctx, connectionId).Execute()

Get Connection



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
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.GetConnection(context.Background(), connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.GetConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnection`: ConnectionRead
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.GetConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectionRead**](ConnectionRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataFreshness

> []DataFreshness GetDataFreshness(ctx).FleetIds(fleetIds).FleetRefs(fleetRefs).Execute()

Get Data Freshness



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
	fleetIds := []string{"Inner_example"} // []string | Limit results to specific fleets. **Maximum:** 100 IDs. To specify multiple values, repeat the parameter for each value (e.g., `?fleet_ids=id1&fleet_ids=id2`). (optional)
	fleetRefs := []string{"Inner_example"} // []string | Limit results to specific fleet references. **Maximum:** 100 references. To specify multiple values, repeat the parameter for each value (e.g., `?fleet_refs=ref1&fleet_refs=ref2`). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.GetDataFreshness(context.Background()).FleetIds(fleetIds).FleetRefs(fleetRefs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.GetDataFreshness``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataFreshness`: []DataFreshness
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.GetDataFreshness`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDataFreshnessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetIds** | **[]string** | Limit results to specific fleets. **Maximum:** 100 IDs. To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_ids&#x3D;id1&amp;fleet_ids&#x3D;id2&#x60;). | 
 **fleetRefs** | **[]string** | Limit results to specific fleet references. **Maximum:** 100 references. To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_refs&#x3D;ref1&amp;fleet_refs&#x3D;ref2&#x60;). | 

### Return type

[**[]DataFreshness**](DataFreshness.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchedule

> ScheduleRead GetSchedule(ctx, connectionId, scheduleId).Execute()

Get Schedule



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
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	scheduleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.GetSchedule(context.Background(), connectionId, scheduleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.GetSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSchedule`: ScheduleRead
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.GetSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** |  | 
**scheduleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ScheduleRead**](ScheduleRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnections

> CursorPageTypeVarCustomizedConnectionRead ListConnections(ctx).TspId(tspId).FleetIds(fleetIds).FleetRefs(fleetRefs).Cursor(cursor).Size(size).Execute()

List Connections



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
	tspId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	fleetIds := []*string{"Inner_example"} // []*string | Limit results to specific fleets. **Maximum:** 100 IDs. To specify multiple values, repeat the parameter for each value (e.g., `?fleet_ids=id1&fleet_ids=id2`). (optional)
	fleetRefs := []string{"Inner_example"} // []string | Limit results to specific fleet references. **Maximum:** 100 references. To specify multiple values, repeat the parameter for each value (e.g., `?fleet_refs=ref1&fleet_refs=ref2`). (optional)
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.ListConnections(context.Background()).TspId(tspId).FleetIds(fleetIds).FleetRefs(fleetRefs).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ListConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConnections`: CursorPageTypeVarCustomizedConnectionRead
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.ListConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tspId** | **string** |  | 
 **fleetIds** | **[]string** | Limit results to specific fleets. **Maximum:** 100 IDs. To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_ids&#x3D;id1&amp;fleet_ids&#x3D;id2&#x60;). | 
 **fleetRefs** | **[]string** | Limit results to specific fleet references. **Maximum:** 100 references. To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_refs&#x3D;ref1&amp;fleet_refs&#x3D;ref2&#x60;). | 
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageTypeVarCustomizedConnectionRead**](CursorPageTypeVarCustomizedConnectionRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExecutions

> CursorPageTypeVarCustomizedExecutionRead ListExecutions(ctx, connectionId, scheduleId).Cursor(cursor).Size(size).Execute()

List Executions



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
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	scheduleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.ListExecutions(context.Background(), connectionId, scheduleId).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ListExecutions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListExecutions`: CursorPageTypeVarCustomizedExecutionRead
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.ListExecutions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** |  | 
**scheduleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListExecutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageTypeVarCustomizedExecutionRead**](CursorPageTypeVarCustomizedExecutionRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSchedules

> CursorPageTypeVarCustomizedScheduleRead ListSchedules(ctx, connectionId).Resource(resource).Cursor(cursor).Size(size).Execute()

List Schedules



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
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	resource := openapiclient.ResourceEnum("asset") // ResourceEnum |  (optional)
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.ListSchedules(context.Background(), connectionId).Resource(resource).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ListSchedules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSchedules`: CursorPageTypeVarCustomizedScheduleRead
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.ListSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resource** | [**ResourceEnum**](ResourceEnum.md) |  | 
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageTypeVarCustomizedScheduleRead**](CursorPageTypeVarCustomizedScheduleRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnection

> ConnectionRead UpdateConnection(ctx, connectionId).ConnectionUpdate(connectionUpdate).Execute()

Update Connection



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
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	connectionUpdate := *openapiclient.NewConnectionUpdate() // ConnectionUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.UpdateConnection(context.Background(), connectionId).ConnectionUpdate(connectionUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.UpdateConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateConnection`: ConnectionRead
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.UpdateConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectionUpdate** | [**ConnectionUpdate**](ConnectionUpdate.md) |  | 

### Return type

[**ConnectionRead**](ConnectionRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSchedule

> ScheduleRead UpdateSchedule(ctx, connectionId, scheduleId).ScheduleUpdate(scheduleUpdate).Execute()

Update Schedule



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
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	scheduleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	scheduleUpdate := *openapiclient.NewScheduleUpdate() // ScheduleUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.UpdateSchedule(context.Background(), connectionId, scheduleId).ScheduleUpdate(scheduleUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.UpdateSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSchedule`: ScheduleRead
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.UpdateSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** |  | 
**scheduleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scheduleUpdate** | [**ScheduleUpdate**](ScheduleUpdate.md) |  | 

### Return type

[**ScheduleRead**](ScheduleRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

