# \FleetsAPI

All URIs are relative to *https://api.catenatelematics.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFleet**](FleetsAPI.md#CreateFleet) | **Post** /v2/orgs/fleets | Create Fleet
[**CreateFleetProperties**](FleetsAPI.md#CreateFleetProperties) | **Post** /v2/orgs/fleets/{fleet_id}/properties | Create Fleet Properties
[**DeleteFleet**](FleetsAPI.md#DeleteFleet) | **Delete** /v2/orgs/fleets/{fleet_id} | Delete Fleet
[**DeleteFleetProperty**](FleetsAPI.md#DeleteFleetProperty) | **Delete** /v2/orgs/fleets/{fleet_id}/properties/{property_id} | Delete Fleet Property
[**GetFleet**](FleetsAPI.md#GetFleet) | **Get** /v2/orgs/fleets/{fleet_id} | Get Fleet
[**ListFleetProperties**](FleetsAPI.md#ListFleetProperties) | **Get** /v2/orgs/fleets/{fleet_id}/properties | List Fleet Properties
[**ListFleets**](FleetsAPI.md#ListFleets) | **Get** /v2/orgs/fleets | List Fleets
[**UpdateFleet**](FleetsAPI.md#UpdateFleet) | **Patch** /v2/orgs/fleets/{fleet_id} | Update Fleet



## CreateFleet

> FleetRead CreateFleet(ctx).FleetCreate(fleetCreate).Execute()

Create Fleet



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/orgs"
)

func main() {
	fleetCreate := *openapiclient.NewFleetCreate("Name_example") // FleetCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetsAPI.CreateFleet(context.Background()).FleetCreate(fleetCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetsAPI.CreateFleet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFleet`: FleetRead
	fmt.Fprintf(os.Stdout, "Response from `FleetsAPI.CreateFleet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFleetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetCreate** | [**FleetCreate**](FleetCreate.md) |  | 

### Return type

[**FleetRead**](FleetRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFleetProperties

> []FleetPropertyRead CreateFleetProperties(ctx, fleetId).FleetPropertyCreate(fleetPropertyCreate).Execute()

Create Fleet Properties



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/orgs"
)

func main() {
	fleetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	fleetPropertyCreate := []openapiclient.FleetPropertyCreate{*openapiclient.NewFleetPropertyCreate(openapiclient.FleetPropertyKeyEnum("dot_number"), "Value_example")} // []FleetPropertyCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetsAPI.CreateFleetProperties(context.Background(), fleetId).FleetPropertyCreate(fleetPropertyCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetsAPI.CreateFleetProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFleetProperties`: []FleetPropertyRead
	fmt.Fprintf(os.Stdout, "Response from `FleetsAPI.CreateFleetProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFleetPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fleetPropertyCreate** | [**[]FleetPropertyCreate**](FleetPropertyCreate.md) |  | 

### Return type

[**[]FleetPropertyRead**](FleetPropertyRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFleet

> DeleteFleet(ctx, fleetId).Execute()

Delete Fleet



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/orgs"
)

func main() {
	fleetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FleetsAPI.DeleteFleet(context.Background(), fleetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetsAPI.DeleteFleet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFleetRequest struct via the builder pattern


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


## DeleteFleetProperty

> DeleteFleetProperty(ctx, fleetId, propertyId).Execute()

Delete Fleet Property



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/orgs"
)

func main() {
	fleetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	propertyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FleetsAPI.DeleteFleetProperty(context.Background(), fleetId, propertyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetsAPI.DeleteFleetProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetId** | **string** |  | 
**propertyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFleetPropertyRequest struct via the builder pattern


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


## GetFleet

> FleetRead GetFleet(ctx, fleetId).Execute()

Get Fleet



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/orgs"
)

func main() {
	fleetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetsAPI.GetFleet(context.Background(), fleetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetsAPI.GetFleet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFleet`: FleetRead
	fmt.Fprintf(os.Stdout, "Response from `FleetsAPI.GetFleet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFleetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FleetRead**](FleetRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFleetProperties

> []FleetPropertyRead ListFleetProperties(ctx, fleetId).Execute()

List Fleet Properties



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/orgs"
)

func main() {
	fleetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetsAPI.ListFleetProperties(context.Background(), fleetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetsAPI.ListFleetProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFleetProperties`: []FleetPropertyRead
	fmt.Fprintf(os.Stdout, "Response from `FleetsAPI.ListFleetProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFleetPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FleetPropertyRead**](FleetPropertyRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFleets

> CursorPageCustomizedFleetRead ListFleets(ctx).FleetRefs(fleetRefs).Cursor(cursor).Size(size).Execute()

List Fleets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/orgs"
)

func main() {
	fleetRefs := []string{"Inner_example"} // []string | Limit results to specific fleets using your organization's fleet reference identifiers (optional)
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetsAPI.ListFleets(context.Background()).FleetRefs(fleetRefs).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetsAPI.ListFleets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFleets`: CursorPageCustomizedFleetRead
	fmt.Fprintf(os.Stdout, "Response from `FleetsAPI.ListFleets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFleetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers | 
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 500]

### Return type

[**CursorPageCustomizedFleetRead**](CursorPageCustomizedFleetRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFleet

> FleetRead UpdateFleet(ctx, fleetId).FleetUpdate(fleetUpdate).Execute()

Update Fleet



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/orgs"
)

func main() {
	fleetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	fleetUpdate := *openapiclient.NewFleetUpdate() // FleetUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetsAPI.UpdateFleet(context.Background(), fleetId).FleetUpdate(fleetUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetsAPI.UpdateFleet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFleet`: FleetRead
	fmt.Fprintf(os.Stdout, "Response from `FleetsAPI.UpdateFleet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFleetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fleetUpdate** | [**FleetUpdate**](FleetUpdate.md) |  | 

### Return type

[**FleetRead**](FleetRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

