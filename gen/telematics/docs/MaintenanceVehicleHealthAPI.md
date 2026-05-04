# \MaintenanceVehicleHealthAPI

All URIs are relative to *https://api.catenatelematics.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFuelTransaction**](MaintenanceVehicleHealthAPI.md#CreateFuelTransaction) | **Post** /v2/telematics/fuel-transactions | Create Fuel Transaction
[**ListEngineLogs**](MaintenanceVehicleHealthAPI.md#ListEngineLogs) | **Get** /v2/telematics/engine-logs | List Engine Logs
[**ListFuelTransactions**](MaintenanceVehicleHealthAPI.md#ListFuelTransactions) | **Get** /v2/telematics/fuel-transactions | List Fuel Transactions
[**ListVehicleSensorEvents**](MaintenanceVehicleHealthAPI.md#ListVehicleSensorEvents) | **Get** /v2/telematics/vehicle-sensor-events | List Vehicle Sensor Events



## CreateFuelTransaction

> ResourceOperation CreateFuelTransaction(ctx).FuelTransactionCreate(fuelTransactionCreate).IsSync(isSync).Execute()

Create Fuel Transaction



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
	fuelTransactionCreate := *openapiclient.NewFuelTransactionCreate("ConnectionId_example", "DriverId_example", "VehicleId_example", float32(123), openapiclient.FuelVolumeUnitEnum("CUBIC_METER")) // FuelTransactionCreate | 
	isSync := true // bool | Whether to process the request synchronously. If `true`, Catena will attempt to create or update the resource immediately with the TSP and return the result in the response. When performing synchronous operations, you are responsible for handling any necessary retries in case of transient failures.If `false` (default), Catena will create an asynchronous operation, and you can check the status of the operation using the returned operation ID. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceVehicleHealthAPI.CreateFuelTransaction(context.Background()).FuelTransactionCreate(fuelTransactionCreate).IsSync(isSync).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceVehicleHealthAPI.CreateFuelTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFuelTransaction`: ResourceOperation
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceVehicleHealthAPI.CreateFuelTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFuelTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fuelTransactionCreate** | [**FuelTransactionCreate**](FuelTransactionCreate.md) |  | 
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


## ListEngineLogs

> CursorPageEngineLogRead ListEngineLogs(ctx).FleetIds(fleetIds).FleetRefs(fleetRefs).ConnectionId(connectionId).FromDatetime(fromDatetime).ToDatetime(toDatetime).IncludeSourceData(includeSourceData).DriverIds(driverIds).VehicleIds(vehicleIds).SortBy(sortBy).SortOrder(sortOrder).Cursor(cursor).Size(size).Execute()

List Engine Logs



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
	driverIds := []string{"Inner_example"} // []string | Limit results to specific drivers. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?driver_ids=id1&driver_ids=id2`). (optional)
	vehicleIds := []string{"Inner_example"} // []string | Limit results to specific vehicles. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?vehicle_ids=id1&vehicle_ids=id2`). (optional)
	sortBy := "sortBy_example" // string | The name of the field to sort results by. If not provided, results will be ordered by `occurred_at`.  (optional)
	sortOrder := "sortOrder_example" // string | The order of sorting, either `asc` for ascending or `desc` for descending. Defaults to `asc` if `sort_by` is provided without `sort_order`. (optional) (default to "asc")
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceVehicleHealthAPI.ListEngineLogs(context.Background()).FleetIds(fleetIds).FleetRefs(fleetRefs).ConnectionId(connectionId).FromDatetime(fromDatetime).ToDatetime(toDatetime).IncludeSourceData(includeSourceData).DriverIds(driverIds).VehicleIds(vehicleIds).SortBy(sortBy).SortOrder(sortOrder).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceVehicleHealthAPI.ListEngineLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEngineLogs`: CursorPageEngineLogRead
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceVehicleHealthAPI.ListEngineLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEngineLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetIds** | **[]string** | Limit results to specific fleets using Catena&#39;s fleet IDs. *For your own fleet identifiers, use &#x60;fleet_refs&#x60; instead* To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_ids&#x3D;id1&amp;fleet_ids&#x3D;id2&#x60;). | 
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers. To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_refs&#x3D;ref1&amp;fleet_refs&#x3D;ref2&#x60;). | 
 **connectionId** | **string** | Limit results to a specific provider connection. This is the UUID assigned by Catena when your fleet connects to a TSP. | 
 **fromDatetime** | **time.Time** | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &gt;&#x3D; from_datetime&#x60; **Default value:** &#x60;now() - 1 day&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 45 days | 
 **toDatetime** | **time.Time** | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &lt; to_datetime&#x60; **Default value:** &#x60;now()&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 45 days | 
 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]
 **driverIds** | **[]string** | Limit results to specific drivers. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?driver_ids&#x3D;id1&amp;driver_ids&#x3D;id2&#x60;). | 
 **vehicleIds** | **[]string** | Limit results to specific vehicles. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?vehicle_ids&#x3D;id1&amp;vehicle_ids&#x3D;id2&#x60;). | 
 **sortBy** | **string** | The name of the field to sort results by. If not provided, results will be ordered by &#x60;occurred_at&#x60;.  | 
 **sortOrder** | **string** | The order of sorting, either &#x60;asc&#x60; for ascending or &#x60;desc&#x60; for descending. Defaults to &#x60;asc&#x60; if &#x60;sort_by&#x60; is provided without &#x60;sort_order&#x60;. | [default to &quot;asc&quot;]
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageEngineLogRead**](CursorPageEngineLogRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFuelTransactions

> CursorPageFuelTransactionRead ListFuelTransactions(ctx).FleetIds(fleetIds).FleetRefs(fleetRefs).ConnectionId(connectionId).FromDatetime(fromDatetime).ToDatetime(toDatetime).IncludeSourceData(includeSourceData).DriverIds(driverIds).SourceDriverIds(sourceDriverIds).VehicleIds(vehicleIds).SortBy(sortBy).SortOrder(sortOrder).Cursor(cursor).Size(size).Execute()

List Fuel Transactions



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
	driverIds := []string{"Inner_example"} // []string | Limit results to specific drivers. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?driver_ids=id1&driver_ids=id2`). (optional)
	sourceDriverIds := []string{"Inner_example"} // []string | Limit results to specific source driver IDs. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?source_driver_ids=id1&source_driver_ids=id2`). (optional)
	vehicleIds := []string{"Inner_example"} // []string | Limit results to specific vehicles. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?vehicle_ids=id1&vehicle_ids=id2`). (optional)
	sortBy := "sortBy_example" // string | The name of the field to sort results by. If not provided, results will be ordered by `occurred_at`.  (optional)
	sortOrder := "sortOrder_example" // string | The order of sorting, either `asc` for ascending or `desc` for descending. Defaults to `asc` if `sort_by` is provided without `sort_order`. (optional) (default to "asc")
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceVehicleHealthAPI.ListFuelTransactions(context.Background()).FleetIds(fleetIds).FleetRefs(fleetRefs).ConnectionId(connectionId).FromDatetime(fromDatetime).ToDatetime(toDatetime).IncludeSourceData(includeSourceData).DriverIds(driverIds).SourceDriverIds(sourceDriverIds).VehicleIds(vehicleIds).SortBy(sortBy).SortOrder(sortOrder).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceVehicleHealthAPI.ListFuelTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFuelTransactions`: CursorPageFuelTransactionRead
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceVehicleHealthAPI.ListFuelTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFuelTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetIds** | **[]string** | Limit results to specific fleets using Catena&#39;s fleet IDs. *For your own fleet identifiers, use &#x60;fleet_refs&#x60; instead* To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_ids&#x3D;id1&amp;fleet_ids&#x3D;id2&#x60;). | 
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers. To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_refs&#x3D;ref1&amp;fleet_refs&#x3D;ref2&#x60;). | 
 **connectionId** | **string** | Limit results to a specific provider connection. This is the UUID assigned by Catena when your fleet connects to a TSP. | 
 **fromDatetime** | **time.Time** | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &gt;&#x3D; from_datetime&#x60; **Default value:** &#x60;now() - 1 day&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 45 days | 
 **toDatetime** | **time.Time** | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &lt; to_datetime&#x60; **Default value:** &#x60;now()&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 45 days | 
 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]
 **driverIds** | **[]string** | Limit results to specific drivers. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?driver_ids&#x3D;id1&amp;driver_ids&#x3D;id2&#x60;). | 
 **sourceDriverIds** | **[]string** | Limit results to specific source driver IDs. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?source_driver_ids&#x3D;id1&amp;source_driver_ids&#x3D;id2&#x60;). | 
 **vehicleIds** | **[]string** | Limit results to specific vehicles. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?vehicle_ids&#x3D;id1&amp;vehicle_ids&#x3D;id2&#x60;). | 
 **sortBy** | **string** | The name of the field to sort results by. If not provided, results will be ordered by &#x60;occurred_at&#x60;.  | 
 **sortOrder** | **string** | The order of sorting, either &#x60;asc&#x60; for ascending or &#x60;desc&#x60; for descending. Defaults to &#x60;asc&#x60; if &#x60;sort_by&#x60; is provided without &#x60;sort_order&#x60;. | [default to &quot;asc&quot;]
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageFuelTransactionRead**](CursorPageFuelTransactionRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVehicleSensorEvents

> CursorPageVehicleSensorRead ListVehicleSensorEvents(ctx).FleetIds(fleetIds).FleetRefs(fleetRefs).ConnectionId(connectionId).FromDatetime(fromDatetime).ToDatetime(toDatetime).IncludeSourceData(includeSourceData).VehicleIds(vehicleIds).SortBy(sortBy).SortOrder(sortOrder).Cursor(cursor).Size(size).Execute()

List Vehicle Sensor Events



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
	vehicleIds := []string{"Inner_example"} // []string | Limit results to specific vehicles. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?vehicle_ids=id1&vehicle_ids=id2`). (optional)
	sortBy := "sortBy_example" // string | The name of the field to sort results by. If not provided, results will be ordered by `occurred_at`.  (optional)
	sortOrder := "sortOrder_example" // string | The order of sorting, either `asc` for ascending or `desc` for descending. Defaults to `asc` if `sort_by` is provided without `sort_order`. (optional) (default to "asc")
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceVehicleHealthAPI.ListVehicleSensorEvents(context.Background()).FleetIds(fleetIds).FleetRefs(fleetRefs).ConnectionId(connectionId).FromDatetime(fromDatetime).ToDatetime(toDatetime).IncludeSourceData(includeSourceData).VehicleIds(vehicleIds).SortBy(sortBy).SortOrder(sortOrder).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceVehicleHealthAPI.ListVehicleSensorEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVehicleSensorEvents`: CursorPageVehicleSensorRead
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceVehicleHealthAPI.ListVehicleSensorEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVehicleSensorEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetIds** | **[]string** | Limit results to specific fleets using Catena&#39;s fleet IDs. *For your own fleet identifiers, use &#x60;fleet_refs&#x60; instead* To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_ids&#x3D;id1&amp;fleet_ids&#x3D;id2&#x60;). | 
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers. To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_refs&#x3D;ref1&amp;fleet_refs&#x3D;ref2&#x60;). | 
 **connectionId** | **string** | Limit results to a specific provider connection. This is the UUID assigned by Catena when your fleet connects to a TSP. | 
 **fromDatetime** | **time.Time** | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &gt;&#x3D; from_datetime&#x60; **Default value:** &#x60;now() - 1 day&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 45 days | 
 **toDatetime** | **time.Time** | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &lt; to_datetime&#x60; **Default value:** &#x60;now()&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 45 days | 
 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]
 **vehicleIds** | **[]string** | Limit results to specific vehicles. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?vehicle_ids&#x3D;id1&amp;vehicle_ids&#x3D;id2&#x60;). | 
 **sortBy** | **string** | The name of the field to sort results by. If not provided, results will be ordered by &#x60;occurred_at&#x60;.  | 
 **sortOrder** | **string** | The order of sorting, either &#x60;asc&#x60; for ascending or &#x60;desc&#x60; for descending. Defaults to &#x60;asc&#x60; if &#x60;sort_by&#x60; is provided without &#x60;sort_order&#x60;. | [default to &quot;asc&quot;]
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageVehicleSensorRead**](CursorPageVehicleSensorRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

