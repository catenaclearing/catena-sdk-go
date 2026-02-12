# \FleetOperationsTrackingAPI

All URIs are relative to *https://api.catenatelematics.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTrailer**](FleetOperationsTrackingAPI.md#GetTrailer) | **Get** /v2/telematics/trailers/{trailer_id} | Get Trailer
[**GetVehicle**](FleetOperationsTrackingAPI.md#GetVehicle) | **Get** /v2/telematics/vehicles/{vehicle_id} | Get Vehicle
[**GetVehicleSensorEvents**](FleetOperationsTrackingAPI.md#GetVehicleSensorEvents) | **Get** /v2/telematics/vehicles/{vehicle_id}/sensor-events | Get Vehicle Sensor Events
[**ListTrailerLocations**](FleetOperationsTrackingAPI.md#ListTrailerLocations) | **Get** /v2/telematics/trailer-locations | List Trailer Locations
[**ListTrailers**](FleetOperationsTrackingAPI.md#ListTrailers) | **Get** /v2/telematics/trailers | List Trailers
[**ListVehicleLocations**](FleetOperationsTrackingAPI.md#ListVehicleLocations) | **Get** /v2/telematics/vehicle-locations | List Vehicle Locations
[**ListVehicleSensorEvents**](FleetOperationsTrackingAPI.md#ListVehicleSensorEvents) | **Get** /v2/telematics/vehicle-sensor-events | List Vehicle Sensor Events
[**ListVehicles**](FleetOperationsTrackingAPI.md#ListVehicles) | **Get** /v2/telematics/vehicles | List Vehicles



## GetTrailer

> TrailerRead GetTrailer(ctx, trailerId).IncludeSourceData(includeSourceData).Execute()

Get Trailer



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
	trailerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique identifier of the trailer.
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetOperationsTrackingAPI.GetTrailer(context.Background(), trailerId).IncludeSourceData(includeSourceData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetOperationsTrackingAPI.GetTrailer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrailer`: TrailerRead
	fmt.Fprintf(os.Stdout, "Response from `FleetOperationsTrackingAPI.GetTrailer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trailerId** | **string** | The unique identifier of the trailer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrailerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]

### Return type

[**TrailerRead**](TrailerRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVehicle

> VehicleRead GetVehicle(ctx, vehicleId).IncludeSourceData(includeSourceData).Execute()

Get Vehicle



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
	vehicleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique identifier of the vehicle.
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetOperationsTrackingAPI.GetVehicle(context.Background(), vehicleId).IncludeSourceData(includeSourceData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetOperationsTrackingAPI.GetVehicle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVehicle`: VehicleRead
	fmt.Fprintf(os.Stdout, "Response from `FleetOperationsTrackingAPI.GetVehicle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vehicleId** | **string** | The unique identifier of the vehicle. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVehicleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]

### Return type

[**VehicleRead**](VehicleRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVehicleSensorEvents

> CursorPageVehicleSensorRead GetVehicleSensorEvents(ctx, vehicleId).FleetIds(fleetIds).FleetRefs(fleetRefs).IncludeSourceData(includeSourceData).SortBy(sortBy).SortOrder(sortOrder).Cursor(cursor).Size(size).Execute()

Get Vehicle Sensor Events



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
	vehicleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique identifier of the vehicle
	fleetIds := []string{"Inner_example"} // []string | Limit results to specific fleets using Catena's fleet IDs. *For your own fleet identifiers, use `fleet_refs` instead* (optional)
	fleetRefs := []string{"Inner_example"} // []string | Limit results to specific fleets using your organization's fleet reference identifiers (optional)
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)
	sortBy := "sortBy_example" // string | The name of the field to sort results by. If not provided, results will be ordered by `occurred_at`.  (optional)
	sortOrder := "sortOrder_example" // string | The order of sorting, either `asc` for ascending or `desc` for descending. Defaults to `asc` if `sort_by` is provided without `sort_order`. (optional) (default to "asc")
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetOperationsTrackingAPI.GetVehicleSensorEvents(context.Background(), vehicleId).FleetIds(fleetIds).FleetRefs(fleetRefs).IncludeSourceData(includeSourceData).SortBy(sortBy).SortOrder(sortOrder).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetOperationsTrackingAPI.GetVehicleSensorEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVehicleSensorEvents`: CursorPageVehicleSensorRead
	fmt.Fprintf(os.Stdout, "Response from `FleetOperationsTrackingAPI.GetVehicleSensorEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vehicleId** | **string** | The unique identifier of the vehicle | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVehicleSensorEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fleetIds** | **[]string** | Limit results to specific fleets using Catena&#39;s fleet IDs. *For your own fleet identifiers, use &#x60;fleet_refs&#x60; instead* | 
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers | 
 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]
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


## ListTrailerLocations

> CursorPageTrailerLocationRead ListTrailerLocations(ctx).OnlyLatest(onlyLatest).FleetIds(fleetIds).FleetRefs(fleetRefs).FromDatetime(fromDatetime).ToDatetime(toDatetime).TrailerIds(trailerIds).VehicleIds(vehicleIds).IncludeSourceData(includeSourceData).Cursor(cursor).Size(size).Execute()

List Trailer Locations



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
	onlyLatest := true // bool | If true, only the latest known location per trailer is returned. (optional) (default to true)
	fleetIds := []string{"Inner_example"} // []string | Limit results to specific fleets using Catena's fleet IDs. *For your own fleet identifiers, use `fleet_refs` instead* (optional)
	fleetRefs := []string{"Inner_example"} // []string | Limit results to specific fleets using your organization's fleet reference identifiers (optional)
	fromDatetime := time.Now() // time.Time | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at >= from_datetime` **Default value:** `now() - 1 day` **Restriction:** `to_datetime - from_datetime` cannot exceed 45 days (optional)
	toDatetime := time.Now() // time.Time | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at < to_datetime` **Default value:** `now()` **Restriction:** `to_datetime - from_datetime` cannot exceed 45 days (optional)
	trailerIds := []string{"Inner_example"} // []string | Limit results to specific trailers. **Maximum:** 100 IDs (optional)
	vehicleIds := []string{"Inner_example"} // []string | Limit results to specific vehicles. **Maximum:** 100 IDs (optional)
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetOperationsTrackingAPI.ListTrailerLocations(context.Background()).OnlyLatest(onlyLatest).FleetIds(fleetIds).FleetRefs(fleetRefs).FromDatetime(fromDatetime).ToDatetime(toDatetime).TrailerIds(trailerIds).VehicleIds(vehicleIds).IncludeSourceData(includeSourceData).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetOperationsTrackingAPI.ListTrailerLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTrailerLocations`: CursorPageTrailerLocationRead
	fmt.Fprintf(os.Stdout, "Response from `FleetOperationsTrackingAPI.ListTrailerLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTrailerLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onlyLatest** | **bool** | If true, only the latest known location per trailer is returned. | [default to true]
 **fleetIds** | **[]string** | Limit results to specific fleets using Catena&#39;s fleet IDs. *For your own fleet identifiers, use &#x60;fleet_refs&#x60; instead* | 
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers | 
 **fromDatetime** | **time.Time** | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &gt;&#x3D; from_datetime&#x60; **Default value:** &#x60;now() - 1 day&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 45 days | 
 **toDatetime** | **time.Time** | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &lt; to_datetime&#x60; **Default value:** &#x60;now()&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 45 days | 
 **trailerIds** | **[]string** | Limit results to specific trailers. **Maximum:** 100 IDs | 
 **vehicleIds** | **[]string** | Limit results to specific vehicles. **Maximum:** 100 IDs | 
 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageTrailerLocationRead**](CursorPageTrailerLocationRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTrailers

> CursorPageTrailerRead ListTrailers(ctx).FleetIds(fleetIds).FleetRefs(fleetRefs).IncludeSourceData(includeSourceData).TrailerIds(trailerIds).SourceIds(sourceIds).SortBy(sortBy).SortOrder(sortOrder).Cursor(cursor).Size(size).Execute()

List Trailers



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
	fleetIds := []string{"Inner_example"} // []string | Limit results to specific fleets using Catena's fleet IDs. *For your own fleet identifiers, use `fleet_refs` instead* (optional)
	fleetRefs := []string{"Inner_example"} // []string | Limit results to specific fleets using your organization's fleet reference identifiers (optional)
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)
	trailerIds := []string{"Inner_example"} // []string | Limit results to specific trailers. **Maximum:** 100 IDs (optional)
	sourceIds := []string{"Inner_example"} // []string | Limit results to specific resources using the identifier provided by the TSP. **Maximum:** 100 IDs (optional)
	sortBy := "sortBy_example" // string | The name of the field to sort results by. If not provided, results will be ordered by `occurred_at`.  (optional)
	sortOrder := "sortOrder_example" // string | The order of sorting, either `asc` for ascending or `desc` for descending. Defaults to `asc` if `sort_by` is provided without `sort_order`. (optional) (default to "asc")
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetOperationsTrackingAPI.ListTrailers(context.Background()).FleetIds(fleetIds).FleetRefs(fleetRefs).IncludeSourceData(includeSourceData).TrailerIds(trailerIds).SourceIds(sourceIds).SortBy(sortBy).SortOrder(sortOrder).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetOperationsTrackingAPI.ListTrailers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTrailers`: CursorPageTrailerRead
	fmt.Fprintf(os.Stdout, "Response from `FleetOperationsTrackingAPI.ListTrailers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTrailersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetIds** | **[]string** | Limit results to specific fleets using Catena&#39;s fleet IDs. *For your own fleet identifiers, use &#x60;fleet_refs&#x60; instead* | 
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers | 
 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]
 **trailerIds** | **[]string** | Limit results to specific trailers. **Maximum:** 100 IDs | 
 **sourceIds** | **[]string** | Limit results to specific resources using the identifier provided by the TSP. **Maximum:** 100 IDs | 
 **sortBy** | **string** | The name of the field to sort results by. If not provided, results will be ordered by &#x60;occurred_at&#x60;.  | 
 **sortOrder** | **string** | The order of sorting, either &#x60;asc&#x60; for ascending or &#x60;desc&#x60; for descending. Defaults to &#x60;asc&#x60; if &#x60;sort_by&#x60; is provided without &#x60;sort_order&#x60;. | [default to &quot;asc&quot;]
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageTrailerRead**](CursorPageTrailerRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVehicleLocations

> CursorPageVehicleLocationRead ListVehicleLocations(ctx).OnlyLatest(onlyLatest).FleetIds(fleetIds).FleetRefs(fleetRefs).FromDatetime(fromDatetime).ToDatetime(toDatetime).DriverIds(driverIds).VehicleIds(vehicleIds).IncludeSourceData(includeSourceData).Cursor(cursor).Size(size).Execute()

List Vehicle Locations



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
	onlyLatest := true // bool | If true, only the latest known location per vehicle is returned. (optional) (default to true)
	fleetIds := []string{"Inner_example"} // []string | Limit results to specific fleets using Catena's fleet IDs. *For your own fleet identifiers, use `fleet_refs` instead* (optional)
	fleetRefs := []string{"Inner_example"} // []string | Limit results to specific fleets using your organization's fleet reference identifiers (optional)
	fromDatetime := time.Now() // time.Time | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at >= from_datetime` **Default value:** `now() - 1 day` **Restriction:** `to_datetime - from_datetime` cannot exceed 45 days (optional)
	toDatetime := time.Now() // time.Time | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at < to_datetime` **Default value:** `now()` **Restriction:** `to_datetime - from_datetime` cannot exceed 45 days (optional)
	driverIds := []string{"Inner_example"} // []string | Limit results to specific drivers. **Maximum:** 100 IDs (optional)
	vehicleIds := []string{"Inner_example"} // []string | Limit results to specific vehicles. **Maximum:** 100 IDs (optional)
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetOperationsTrackingAPI.ListVehicleLocations(context.Background()).OnlyLatest(onlyLatest).FleetIds(fleetIds).FleetRefs(fleetRefs).FromDatetime(fromDatetime).ToDatetime(toDatetime).DriverIds(driverIds).VehicleIds(vehicleIds).IncludeSourceData(includeSourceData).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetOperationsTrackingAPI.ListVehicleLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVehicleLocations`: CursorPageVehicleLocationRead
	fmt.Fprintf(os.Stdout, "Response from `FleetOperationsTrackingAPI.ListVehicleLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVehicleLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onlyLatest** | **bool** | If true, only the latest known location per vehicle is returned. | [default to true]
 **fleetIds** | **[]string** | Limit results to specific fleets using Catena&#39;s fleet IDs. *For your own fleet identifiers, use &#x60;fleet_refs&#x60; instead* | 
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers | 
 **fromDatetime** | **time.Time** | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &gt;&#x3D; from_datetime&#x60; **Default value:** &#x60;now() - 1 day&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 45 days | 
 **toDatetime** | **time.Time** | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &lt; to_datetime&#x60; **Default value:** &#x60;now()&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 45 days | 
 **driverIds** | **[]string** | Limit results to specific drivers. **Maximum:** 100 IDs | 
 **vehicleIds** | **[]string** | Limit results to specific vehicles. **Maximum:** 100 IDs | 
 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageVehicleLocationRead**](CursorPageVehicleLocationRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVehicleSensorEvents

> CursorPageVehicleSensorRead ListVehicleSensorEvents(ctx).FleetIds(fleetIds).FleetRefs(fleetRefs).FromDatetime(fromDatetime).ToDatetime(toDatetime).IncludeSourceData(includeSourceData).VehicleIds(vehicleIds).SortBy(sortBy).SortOrder(sortOrder).Cursor(cursor).Size(size).Execute()

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
	fleetIds := []string{"Inner_example"} // []string | Limit results to specific fleets using Catena's fleet IDs. *For your own fleet identifiers, use `fleet_refs` instead* (optional)
	fleetRefs := []string{"Inner_example"} // []string | Limit results to specific fleets using your organization's fleet reference identifiers (optional)
	fromDatetime := time.Now() // time.Time | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at >= from_datetime` **Default value:** `now() - 1 day` **Restriction:** `to_datetime - from_datetime` cannot exceed 45 days (optional)
	toDatetime := time.Now() // time.Time | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at < to_datetime` **Default value:** `now()` **Restriction:** `to_datetime - from_datetime` cannot exceed 45 days (optional)
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)
	vehicleIds := []string{"Inner_example"} // []string | Limit results to specific vehicles. **Maximum:** 100 IDs (optional)
	sortBy := "sortBy_example" // string | The name of the field to sort results by. If not provided, results will be ordered by `occurred_at`.  (optional)
	sortOrder := "sortOrder_example" // string | The order of sorting, either `asc` for ascending or `desc` for descending. Defaults to `asc` if `sort_by` is provided without `sort_order`. (optional) (default to "asc")
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetOperationsTrackingAPI.ListVehicleSensorEvents(context.Background()).FleetIds(fleetIds).FleetRefs(fleetRefs).FromDatetime(fromDatetime).ToDatetime(toDatetime).IncludeSourceData(includeSourceData).VehicleIds(vehicleIds).SortBy(sortBy).SortOrder(sortOrder).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetOperationsTrackingAPI.ListVehicleSensorEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVehicleSensorEvents`: CursorPageVehicleSensorRead
	fmt.Fprintf(os.Stdout, "Response from `FleetOperationsTrackingAPI.ListVehicleSensorEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVehicleSensorEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetIds** | **[]string** | Limit results to specific fleets using Catena&#39;s fleet IDs. *For your own fleet identifiers, use &#x60;fleet_refs&#x60; instead* | 
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers | 
 **fromDatetime** | **time.Time** | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &gt;&#x3D; from_datetime&#x60; **Default value:** &#x60;now() - 1 day&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 45 days | 
 **toDatetime** | **time.Time** | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &lt; to_datetime&#x60; **Default value:** &#x60;now()&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 45 days | 
 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]
 **vehicleIds** | **[]string** | Limit results to specific vehicles. **Maximum:** 100 IDs | 
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


## ListVehicles

> CursorPageVehicleRead ListVehicles(ctx).FleetIds(fleetIds).FleetRefs(fleetRefs).IncludeSourceData(includeSourceData).VehicleIds(vehicleIds).SourceIds(sourceIds).SortBy(sortBy).SortOrder(sortOrder).Cursor(cursor).Size(size).Execute()

List Vehicles



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
	fleetIds := []string{"Inner_example"} // []string | Limit results to specific fleets using Catena's fleet IDs. *For your own fleet identifiers, use `fleet_refs` instead* (optional)
	fleetRefs := []string{"Inner_example"} // []string | Limit results to specific fleets using your organization's fleet reference identifiers (optional)
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)
	vehicleIds := []string{"Inner_example"} // []string | Limit results to specific vehicles. **Maximum:** 100 IDs (optional)
	sourceIds := []string{"Inner_example"} // []string | Limit results to specific resources using the identifier provided by the TSP. **Maximum:** 100 IDs (optional)
	sortBy := "sortBy_example" // string | The name of the field to sort results by. If not provided, results will be ordered by `occurred_at`.  (optional)
	sortOrder := "sortOrder_example" // string | The order of sorting, either `asc` for ascending or `desc` for descending. Defaults to `asc` if `sort_by` is provided without `sort_order`. (optional) (default to "asc")
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetOperationsTrackingAPI.ListVehicles(context.Background()).FleetIds(fleetIds).FleetRefs(fleetRefs).IncludeSourceData(includeSourceData).VehicleIds(vehicleIds).SourceIds(sourceIds).SortBy(sortBy).SortOrder(sortOrder).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetOperationsTrackingAPI.ListVehicles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVehicles`: CursorPageVehicleRead
	fmt.Fprintf(os.Stdout, "Response from `FleetOperationsTrackingAPI.ListVehicles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVehiclesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetIds** | **[]string** | Limit results to specific fleets using Catena&#39;s fleet IDs. *For your own fleet identifiers, use &#x60;fleet_refs&#x60; instead* | 
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers | 
 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]
 **vehicleIds** | **[]string** | Limit results to specific vehicles. **Maximum:** 100 IDs | 
 **sourceIds** | **[]string** | Limit results to specific resources using the identifier provided by the TSP. **Maximum:** 100 IDs | 
 **sortBy** | **string** | The name of the field to sort results by. If not provided, results will be ordered by &#x60;occurred_at&#x60;.  | 
 **sortOrder** | **string** | The order of sorting, either &#x60;asc&#x60; for ascending or &#x60;desc&#x60; for descending. Defaults to &#x60;asc&#x60; if &#x60;sort_by&#x60; is provided without &#x60;sort_order&#x60;. | [default to &quot;asc&quot;]
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageVehicleRead**](CursorPageVehicleRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

