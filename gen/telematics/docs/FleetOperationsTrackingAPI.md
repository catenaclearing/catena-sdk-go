# \FleetOperationsTrackingAPI

All URIs are relative to *https://api.catenatelematics.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTrailer**](FleetOperationsTrackingAPI.md#GetTrailer) | **Get** /v2/telematics/trailers/{trailer_id} | Get Trailer
[**GetVehicle**](FleetOperationsTrackingAPI.md#GetVehicle) | **Get** /v2/telematics/vehicles/{vehicle_id} | Get Vehicle
[**GetVehicleSensorEvents**](FleetOperationsTrackingAPI.md#GetVehicleSensorEvents) | **Get** /v2/telematics/vehicles/{vehicle_id}/sensor-events | Get Vehicle Sensor Events
[**ListTrailers**](FleetOperationsTrackingAPI.md#ListTrailers) | **Get** /v2/telematics/trailers | List Trailers
[**ListVehicleLocations**](FleetOperationsTrackingAPI.md#ListVehicleLocations) | **Get** /v2/telematics/vehicle-locations | List Vehicle Locations
[**ListVehicleSensorEvents**](FleetOperationsTrackingAPI.md#ListVehicleSensorEvents) | **Get** /v2/telematics/vehicle-sensor-events | List Vehicle Sensor Events
[**ListVehicles**](FleetOperationsTrackingAPI.md#ListVehicles) | **Get** /v2/telematics/vehicles | List Vehicles



## GetTrailer

> Trailer GetTrailer(ctx, trailerId).IncludeSourceData(includeSourceData).Execute()

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
	// response from `GetTrailer`: Trailer
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

[**Trailer**](Trailer.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVehicle

> Vehicle GetVehicle(ctx, vehicleId).IncludeSourceData(includeSourceData).Execute()

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
	// response from `GetVehicle`: Vehicle
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

[**Vehicle**](Vehicle.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVehicleSensorEvents

> CursorPageTypeVarCustomizedVehicleSensor GetVehicleSensorEvents(ctx, vehicleId).FleetIds(fleetIds).FleetRefs(fleetRefs).IncludeSourceData(includeSourceData).Cursor(cursor).Size(size).Execute()

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
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetOperationsTrackingAPI.GetVehicleSensorEvents(context.Background(), vehicleId).FleetIds(fleetIds).FleetRefs(fleetRefs).IncludeSourceData(includeSourceData).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetOperationsTrackingAPI.GetVehicleSensorEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVehicleSensorEvents`: CursorPageTypeVarCustomizedVehicleSensor
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
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageTypeVarCustomizedVehicleSensor**](CursorPageTypeVarCustomizedVehicleSensor.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTrailers

> CursorPageTypeVarCustomizedTrailer ListTrailers(ctx).FleetIds(fleetIds).FleetRefs(fleetRefs).IncludeSourceData(includeSourceData).TrailerIds(trailerIds).Cursor(cursor).Size(size).Execute()

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
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetOperationsTrackingAPI.ListTrailers(context.Background()).FleetIds(fleetIds).FleetRefs(fleetRefs).IncludeSourceData(includeSourceData).TrailerIds(trailerIds).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetOperationsTrackingAPI.ListTrailers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTrailers`: CursorPageTypeVarCustomizedTrailer
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
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageTypeVarCustomizedTrailer**](CursorPageTypeVarCustomizedTrailer.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVehicleLocations

> CursorPageTypeVarCustomizedVehicleLocation ListVehicleLocations(ctx).OnlyLatest(onlyLatest).FleetIds(fleetIds).FleetRefs(fleetRefs).FromDatetime(fromDatetime).ToDatetime(toDatetime).DriverIds(driverIds).VehicleIds(vehicleIds).IncludeSourceData(includeSourceData).Cursor(cursor).Size(size).Execute()

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
	fromDatetime := time.Now() // time.Time | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at >= from_datetime` **Default value:** `now() - 1 day` **Restriction:** `to_datetime - from_datetime` cannot exceed 15 days (optional)
	toDatetime := time.Now() // time.Time | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at < to_datetime` **Default value:** `now()` **Restriction:** `to_datetime - from_datetime` cannot exceed 15 days (optional)
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
	// response from `ListVehicleLocations`: CursorPageTypeVarCustomizedVehicleLocation
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
 **fromDatetime** | **time.Time** | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &gt;&#x3D; from_datetime&#x60; **Default value:** &#x60;now() - 1 day&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 15 days | 
 **toDatetime** | **time.Time** | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &lt; to_datetime&#x60; **Default value:** &#x60;now()&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 15 days | 
 **driverIds** | **[]string** | Limit results to specific drivers. **Maximum:** 100 IDs | 
 **vehicleIds** | **[]string** | Limit results to specific vehicles. **Maximum:** 100 IDs | 
 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageTypeVarCustomizedVehicleLocation**](CursorPageTypeVarCustomizedVehicleLocation.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVehicleSensorEvents

> CursorPageTypeVarCustomizedVehicleSensor ListVehicleSensorEvents(ctx).FleetIds(fleetIds).FleetRefs(fleetRefs).FromDatetime(fromDatetime).ToDatetime(toDatetime).IncludeSourceData(includeSourceData).VehicleIds(vehicleIds).Cursor(cursor).Size(size).Execute()

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
	fromDatetime := time.Now() // time.Time | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at >= from_datetime` **Default value:** `now() - 1 day` **Restriction:** `to_datetime - from_datetime` cannot exceed 15 days (optional)
	toDatetime := time.Now() // time.Time | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at < to_datetime` **Default value:** `now()` **Restriction:** `to_datetime - from_datetime` cannot exceed 15 days (optional)
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)
	vehicleIds := []string{"Inner_example"} // []string | Limit results to specific vehicles. **Maximum:** 100 IDs (optional)
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetOperationsTrackingAPI.ListVehicleSensorEvents(context.Background()).FleetIds(fleetIds).FleetRefs(fleetRefs).FromDatetime(fromDatetime).ToDatetime(toDatetime).IncludeSourceData(includeSourceData).VehicleIds(vehicleIds).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetOperationsTrackingAPI.ListVehicleSensorEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVehicleSensorEvents`: CursorPageTypeVarCustomizedVehicleSensor
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
 **fromDatetime** | **time.Time** | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &gt;&#x3D; from_datetime&#x60; **Default value:** &#x60;now() - 1 day&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 15 days | 
 **toDatetime** | **time.Time** | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &lt; to_datetime&#x60; **Default value:** &#x60;now()&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 15 days | 
 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]
 **vehicleIds** | **[]string** | Limit results to specific vehicles. **Maximum:** 100 IDs | 
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageTypeVarCustomizedVehicleSensor**](CursorPageTypeVarCustomizedVehicleSensor.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVehicles

> CursorPageTypeVarCustomizedVehicle ListVehicles(ctx).FleetIds(fleetIds).FleetRefs(fleetRefs).IncludeSourceData(includeSourceData).VehicleIds(vehicleIds).Cursor(cursor).Size(size).Execute()

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
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetOperationsTrackingAPI.ListVehicles(context.Background()).FleetIds(fleetIds).FleetRefs(fleetRefs).IncludeSourceData(includeSourceData).VehicleIds(vehicleIds).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetOperationsTrackingAPI.ListVehicles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVehicles`: CursorPageTypeVarCustomizedVehicle
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
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageTypeVarCustomizedVehicle**](CursorPageTypeVarCustomizedVehicle.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

