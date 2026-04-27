# \FleetOperationsTrackingAPI

All URIs are relative to *https://api.catenatelematics.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVehicle**](FleetOperationsTrackingAPI.md#CreateVehicle) | **Post** /v2/telematics/vehicles | Create Vehicle
[**GetTrailer**](FleetOperationsTrackingAPI.md#GetTrailer) | **Get** /v2/telematics/trailers/{trailer_id} | Get Trailer
[**GetVehicle**](FleetOperationsTrackingAPI.md#GetVehicle) | **Get** /v2/telematics/vehicles/{vehicle_id} | Get Vehicle
[**ListDriverVehicleAssociations**](FleetOperationsTrackingAPI.md#ListDriverVehicleAssociations) | **Get** /v2/telematics/driver-vehicle-associations | List Driver Vehicle Associations
[**ListEngineStatuses**](FleetOperationsTrackingAPI.md#ListEngineStatuses) | **Get** /v2/telematics/engine-statuses | List Engine Statuses
[**ListTrailerLocations**](FleetOperationsTrackingAPI.md#ListTrailerLocations) | **Get** /v2/telematics/trailer-locations | List Trailer Locations
[**ListTrailerStatuses**](FleetOperationsTrackingAPI.md#ListTrailerStatuses) | **Get** /v2/telematics/trailer-statuses | List Trailer Statuses
[**ListTrailerVehicleAssociations**](FleetOperationsTrackingAPI.md#ListTrailerVehicleAssociations) | **Get** /v2/telematics/trailer-vehicle-associations | List Trailer Vehicle Associations
[**ListTrailers**](FleetOperationsTrackingAPI.md#ListTrailers) | **Get** /v2/telematics/trailers | List Trailers
[**ListVehicleLocations**](FleetOperationsTrackingAPI.md#ListVehicleLocations) | **Get** /v2/telematics/vehicle-locations | List Vehicle Locations
[**ListVehicleRegionSegments**](FleetOperationsTrackingAPI.md#ListVehicleRegionSegments) | **Get** /v2/telematics/vehicle-region-segments | List Vehicle Region Segments
[**ListVehicles**](FleetOperationsTrackingAPI.md#ListVehicles) | **Get** /v2/telematics/vehicles | List Vehicles
[**UpdateVehicle**](FleetOperationsTrackingAPI.md#UpdateVehicle) | **Patch** /v2/telematics/vehicles/{source_id} | Update Vehicle



## CreateVehicle

> ResourceOperationAccept CreateVehicle(ctx).VehicleCreate(vehicleCreate).Execute()

Create Vehicle



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
	vehicleCreate := *openapiclient.NewVehicleCreate("ConnectionId_example") // VehicleCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetOperationsTrackingAPI.CreateVehicle(context.Background()).VehicleCreate(vehicleCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetOperationsTrackingAPI.CreateVehicle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVehicle`: ResourceOperationAccept
	fmt.Fprintf(os.Stdout, "Response from `FleetOperationsTrackingAPI.CreateVehicle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVehicleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vehicleCreate** | [**VehicleCreate**](VehicleCreate.md) |  | 

### Return type

[**ResourceOperationAccept**](ResourceOperationAccept.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## ListDriverVehicleAssociations

> CursorPageDriverVehicleAssociationRead ListDriverVehicleAssociations(ctx).OnlyLatest(onlyLatest).FleetIds(fleetIds).FleetRefs(fleetRefs).ConnectionId(connectionId).FromDatetime(fromDatetime).ToDatetime(toDatetime).DriverIds(driverIds).VehicleIds(vehicleIds).IncludeSourceData(includeSourceData).Cursor(cursor).Size(size).Execute()

List Driver Vehicle Associations



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
	onlyLatest := true // bool | If true, only the latest known association per driver is returned. (optional) (default to true)
	fleetIds := []string{"Inner_example"} // []string | Limit results to specific fleets using Catena's fleet IDs. *For your own fleet identifiers, use `fleet_refs` instead* To specify multiple values, repeat the parameter for each value (e.g., `?fleet_ids=id1&fleet_ids=id2`). (optional)
	fleetRefs := []string{"Inner_example"} // []string | Limit results to specific fleets using your organization's fleet reference identifiers. To specify multiple values, repeat the parameter for each value (e.g., `?fleet_refs=ref1&fleet_refs=ref2`). (optional)
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Limit results to a specific provider connection. This is the UUID assigned by Catena when your fleet connects to a TSP. (optional)
	fromDatetime := time.Now() // time.Time | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at >= from_datetime` **Default value:** `now() - 1 day` **Restriction:** `to_datetime - from_datetime` cannot exceed 45 days (optional)
	toDatetime := time.Now() // time.Time | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at < to_datetime` **Default value:** `now()` **Restriction:** `to_datetime - from_datetime` cannot exceed 45 days (optional)
	driverIds := []string{"Inner_example"} // []string | Limit results to specific drivers. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?driver_ids=id1&driver_ids=id2`). (optional)
	vehicleIds := []string{"Inner_example"} // []string | Limit results to specific vehicles. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?vehicle_ids=id1&vehicle_ids=id2`). (optional)
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetOperationsTrackingAPI.ListDriverVehicleAssociations(context.Background()).OnlyLatest(onlyLatest).FleetIds(fleetIds).FleetRefs(fleetRefs).ConnectionId(connectionId).FromDatetime(fromDatetime).ToDatetime(toDatetime).DriverIds(driverIds).VehicleIds(vehicleIds).IncludeSourceData(includeSourceData).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetOperationsTrackingAPI.ListDriverVehicleAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDriverVehicleAssociations`: CursorPageDriverVehicleAssociationRead
	fmt.Fprintf(os.Stdout, "Response from `FleetOperationsTrackingAPI.ListDriverVehicleAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDriverVehicleAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onlyLatest** | **bool** | If true, only the latest known association per driver is returned. | [default to true]
 **fleetIds** | **[]string** | Limit results to specific fleets using Catena&#39;s fleet IDs. *For your own fleet identifiers, use &#x60;fleet_refs&#x60; instead* To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_ids&#x3D;id1&amp;fleet_ids&#x3D;id2&#x60;). | 
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers. To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_refs&#x3D;ref1&amp;fleet_refs&#x3D;ref2&#x60;). | 
 **connectionId** | **string** | Limit results to a specific provider connection. This is the UUID assigned by Catena when your fleet connects to a TSP. | 
 **fromDatetime** | **time.Time** | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &gt;&#x3D; from_datetime&#x60; **Default value:** &#x60;now() - 1 day&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 45 days | 
 **toDatetime** | **time.Time** | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &lt; to_datetime&#x60; **Default value:** &#x60;now()&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 45 days | 
 **driverIds** | **[]string** | Limit results to specific drivers. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?driver_ids&#x3D;id1&amp;driver_ids&#x3D;id2&#x60;). | 
 **vehicleIds** | **[]string** | Limit results to specific vehicles. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?vehicle_ids&#x3D;id1&amp;vehicle_ids&#x3D;id2&#x60;). | 
 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageDriverVehicleAssociationRead**](CursorPageDriverVehicleAssociationRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEngineStatuses

> CursorPageEngineStatusRead ListEngineStatuses(ctx).OnlyLatest(onlyLatest).FleetIds(fleetIds).FleetRefs(fleetRefs).ConnectionId(connectionId).FromDatetime(fromDatetime).ToDatetime(toDatetime).DriverIds(driverIds).VehicleIds(vehicleIds).IncludeSourceData(includeSourceData).Cursor(cursor).Size(size).Execute()

List Engine Statuses



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
	onlyLatest := true // bool | If true, only the latest known engine status per vehicle is returned. (optional) (default to true)
	fleetIds := []string{"Inner_example"} // []string | Limit results to specific fleets using Catena's fleet IDs. *For your own fleet identifiers, use `fleet_refs` instead* To specify multiple values, repeat the parameter for each value (e.g., `?fleet_ids=id1&fleet_ids=id2`). (optional)
	fleetRefs := []string{"Inner_example"} // []string | Limit results to specific fleets using your organization's fleet reference identifiers. To specify multiple values, repeat the parameter for each value (e.g., `?fleet_refs=ref1&fleet_refs=ref2`). (optional)
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Limit results to a specific provider connection. This is the UUID assigned by Catena when your fleet connects to a TSP. (optional)
	fromDatetime := time.Now() // time.Time | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at >= from_datetime` **Default value:** `now() - 1 day` **Restriction:** `to_datetime - from_datetime` cannot exceed 45 days (optional)
	toDatetime := time.Now() // time.Time | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at < to_datetime` **Default value:** `now()` **Restriction:** `to_datetime - from_datetime` cannot exceed 45 days (optional)
	driverIds := []string{"Inner_example"} // []string | Limit results to specific drivers. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?driver_ids=id1&driver_ids=id2`). (optional)
	vehicleIds := []string{"Inner_example"} // []string | Limit results to specific vehicles. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?vehicle_ids=id1&vehicle_ids=id2`). (optional)
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetOperationsTrackingAPI.ListEngineStatuses(context.Background()).OnlyLatest(onlyLatest).FleetIds(fleetIds).FleetRefs(fleetRefs).ConnectionId(connectionId).FromDatetime(fromDatetime).ToDatetime(toDatetime).DriverIds(driverIds).VehicleIds(vehicleIds).IncludeSourceData(includeSourceData).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetOperationsTrackingAPI.ListEngineStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEngineStatuses`: CursorPageEngineStatusRead
	fmt.Fprintf(os.Stdout, "Response from `FleetOperationsTrackingAPI.ListEngineStatuses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEngineStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onlyLatest** | **bool** | If true, only the latest known engine status per vehicle is returned. | [default to true]
 **fleetIds** | **[]string** | Limit results to specific fleets using Catena&#39;s fleet IDs. *For your own fleet identifiers, use &#x60;fleet_refs&#x60; instead* To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_ids&#x3D;id1&amp;fleet_ids&#x3D;id2&#x60;). | 
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers. To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_refs&#x3D;ref1&amp;fleet_refs&#x3D;ref2&#x60;). | 
 **connectionId** | **string** | Limit results to a specific provider connection. This is the UUID assigned by Catena when your fleet connects to a TSP. | 
 **fromDatetime** | **time.Time** | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &gt;&#x3D; from_datetime&#x60; **Default value:** &#x60;now() - 1 day&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 45 days | 
 **toDatetime** | **time.Time** | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &lt; to_datetime&#x60; **Default value:** &#x60;now()&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 45 days | 
 **driverIds** | **[]string** | Limit results to specific drivers. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?driver_ids&#x3D;id1&amp;driver_ids&#x3D;id2&#x60;). | 
 **vehicleIds** | **[]string** | Limit results to specific vehicles. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?vehicle_ids&#x3D;id1&amp;vehicle_ids&#x3D;id2&#x60;). | 
 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageEngineStatusRead**](CursorPageEngineStatusRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTrailerLocations

> CursorPageTrailerLocationRead ListTrailerLocations(ctx).OnlyLatest(onlyLatest).FleetIds(fleetIds).FleetRefs(fleetRefs).ConnectionId(connectionId).FromDatetime(fromDatetime).ToDatetime(toDatetime).TrailerIds(trailerIds).VehicleIds(vehicleIds).IncludeSourceData(includeSourceData).Cursor(cursor).Size(size).Execute()

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
	fleetIds := []string{"Inner_example"} // []string | Limit results to specific fleets using Catena's fleet IDs. *For your own fleet identifiers, use `fleet_refs` instead* To specify multiple values, repeat the parameter for each value (e.g., `?fleet_ids=id1&fleet_ids=id2`). (optional)
	fleetRefs := []string{"Inner_example"} // []string | Limit results to specific fleets using your organization's fleet reference identifiers. To specify multiple values, repeat the parameter for each value (e.g., `?fleet_refs=ref1&fleet_refs=ref2`). (optional)
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Limit results to a specific provider connection. This is the UUID assigned by Catena when your fleet connects to a TSP. (optional)
	fromDatetime := time.Now() // time.Time | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at >= from_datetime` **Default value:** `now() - 1 day` **Restriction:** `to_datetime - from_datetime` cannot exceed 45 days (optional)
	toDatetime := time.Now() // time.Time | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at < to_datetime` **Default value:** `now()` **Restriction:** `to_datetime - from_datetime` cannot exceed 45 days (optional)
	trailerIds := []string{"Inner_example"} // []string | Limit results to specific trailers. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?trailer_ids=id1&trailer_ids=id2`). (optional)
	vehicleIds := []string{"Inner_example"} // []string | Limit results to specific vehicles. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?vehicle_ids=id1&vehicle_ids=id2`). (optional)
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetOperationsTrackingAPI.ListTrailerLocations(context.Background()).OnlyLatest(onlyLatest).FleetIds(fleetIds).FleetRefs(fleetRefs).ConnectionId(connectionId).FromDatetime(fromDatetime).ToDatetime(toDatetime).TrailerIds(trailerIds).VehicleIds(vehicleIds).IncludeSourceData(includeSourceData).Cursor(cursor).Size(size).Execute()
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
 **fleetIds** | **[]string** | Limit results to specific fleets using Catena&#39;s fleet IDs. *For your own fleet identifiers, use &#x60;fleet_refs&#x60; instead* To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_ids&#x3D;id1&amp;fleet_ids&#x3D;id2&#x60;). | 
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers. To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_refs&#x3D;ref1&amp;fleet_refs&#x3D;ref2&#x60;). | 
 **connectionId** | **string** | Limit results to a specific provider connection. This is the UUID assigned by Catena when your fleet connects to a TSP. | 
 **fromDatetime** | **time.Time** | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &gt;&#x3D; from_datetime&#x60; **Default value:** &#x60;now() - 1 day&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 45 days | 
 **toDatetime** | **time.Time** | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &lt; to_datetime&#x60; **Default value:** &#x60;now()&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 45 days | 
 **trailerIds** | **[]string** | Limit results to specific trailers. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?trailer_ids&#x3D;id1&amp;trailer_ids&#x3D;id2&#x60;). | 
 **vehicleIds** | **[]string** | Limit results to specific vehicles. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?vehicle_ids&#x3D;id1&amp;vehicle_ids&#x3D;id2&#x60;). | 
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


## ListTrailerStatuses

> CursorPageTrailerStatusRead ListTrailerStatuses(ctx).OnlyLatest(onlyLatest).FleetIds(fleetIds).FleetRefs(fleetRefs).ConnectionId(connectionId).FromDatetime(fromDatetime).ToDatetime(toDatetime).TrailerIds(trailerIds).IncludeSourceData(includeSourceData).Cursor(cursor).Size(size).Execute()

List Trailer Statuses



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
	onlyLatest := true // bool | If true, only the latest known status per trailer is returned. (optional) (default to true)
	fleetIds := []string{"Inner_example"} // []string | Limit results to specific fleets using Catena's fleet IDs. *For your own fleet identifiers, use `fleet_refs` instead* To specify multiple values, repeat the parameter for each value (e.g., `?fleet_ids=id1&fleet_ids=id2`). (optional)
	fleetRefs := []string{"Inner_example"} // []string | Limit results to specific fleets using your organization's fleet reference identifiers. To specify multiple values, repeat the parameter for each value (e.g., `?fleet_refs=ref1&fleet_refs=ref2`). (optional)
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Limit results to a specific provider connection. This is the UUID assigned by Catena when your fleet connects to a TSP. (optional)
	fromDatetime := time.Now() // time.Time | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at >= from_datetime` **Default value:** `now() - 1 day` **Restriction:** `to_datetime - from_datetime` cannot exceed 45 days (optional)
	toDatetime := time.Now() // time.Time | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at < to_datetime` **Default value:** `now()` **Restriction:** `to_datetime - from_datetime` cannot exceed 45 days (optional)
	trailerIds := []string{"Inner_example"} // []string | Limit results to specific trailers. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?trailer_ids=id1&trailer_ids=id2`). (optional)
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetOperationsTrackingAPI.ListTrailerStatuses(context.Background()).OnlyLatest(onlyLatest).FleetIds(fleetIds).FleetRefs(fleetRefs).ConnectionId(connectionId).FromDatetime(fromDatetime).ToDatetime(toDatetime).TrailerIds(trailerIds).IncludeSourceData(includeSourceData).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetOperationsTrackingAPI.ListTrailerStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTrailerStatuses`: CursorPageTrailerStatusRead
	fmt.Fprintf(os.Stdout, "Response from `FleetOperationsTrackingAPI.ListTrailerStatuses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTrailerStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onlyLatest** | **bool** | If true, only the latest known status per trailer is returned. | [default to true]
 **fleetIds** | **[]string** | Limit results to specific fleets using Catena&#39;s fleet IDs. *For your own fleet identifiers, use &#x60;fleet_refs&#x60; instead* To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_ids&#x3D;id1&amp;fleet_ids&#x3D;id2&#x60;). | 
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers. To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_refs&#x3D;ref1&amp;fleet_refs&#x3D;ref2&#x60;). | 
 **connectionId** | **string** | Limit results to a specific provider connection. This is the UUID assigned by Catena when your fleet connects to a TSP. | 
 **fromDatetime** | **time.Time** | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &gt;&#x3D; from_datetime&#x60; **Default value:** &#x60;now() - 1 day&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 45 days | 
 **toDatetime** | **time.Time** | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &lt; to_datetime&#x60; **Default value:** &#x60;now()&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 45 days | 
 **trailerIds** | **[]string** | Limit results to specific trailers. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?trailer_ids&#x3D;id1&amp;trailer_ids&#x3D;id2&#x60;). | 
 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageTrailerStatusRead**](CursorPageTrailerStatusRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTrailerVehicleAssociations

> CursorPageTrailerVehicleAssociationRead ListTrailerVehicleAssociations(ctx).OnlyLatest(onlyLatest).FleetIds(fleetIds).FleetRefs(fleetRefs).ConnectionId(connectionId).FromDatetime(fromDatetime).ToDatetime(toDatetime).TrailerIds(trailerIds).VehicleIds(vehicleIds).IncludeSourceData(includeSourceData).Cursor(cursor).Size(size).Execute()

List Trailer Vehicle Associations



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
	onlyLatest := true // bool | If true, only the latest known association per trailer is returned. (optional) (default to true)
	fleetIds := []string{"Inner_example"} // []string | Limit results to specific fleets using Catena's fleet IDs. *For your own fleet identifiers, use `fleet_refs` instead* To specify multiple values, repeat the parameter for each value (e.g., `?fleet_ids=id1&fleet_ids=id2`). (optional)
	fleetRefs := []string{"Inner_example"} // []string | Limit results to specific fleets using your organization's fleet reference identifiers. To specify multiple values, repeat the parameter for each value (e.g., `?fleet_refs=ref1&fleet_refs=ref2`). (optional)
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Limit results to a specific provider connection. This is the UUID assigned by Catena when your fleet connects to a TSP. (optional)
	fromDatetime := time.Now() // time.Time | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at >= from_datetime` **Default value:** `now() - 1 day` **Restriction:** `to_datetime - from_datetime` cannot exceed 45 days (optional)
	toDatetime := time.Now() // time.Time | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at < to_datetime` **Default value:** `now()` **Restriction:** `to_datetime - from_datetime` cannot exceed 45 days (optional)
	trailerIds := []string{"Inner_example"} // []string | Limit results to specific trailers. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?trailer_ids=id1&trailer_ids=id2`). (optional)
	vehicleIds := []string{"Inner_example"} // []string | Limit results to specific vehicles. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?vehicle_ids=id1&vehicle_ids=id2`). (optional)
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetOperationsTrackingAPI.ListTrailerVehicleAssociations(context.Background()).OnlyLatest(onlyLatest).FleetIds(fleetIds).FleetRefs(fleetRefs).ConnectionId(connectionId).FromDatetime(fromDatetime).ToDatetime(toDatetime).TrailerIds(trailerIds).VehicleIds(vehicleIds).IncludeSourceData(includeSourceData).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetOperationsTrackingAPI.ListTrailerVehicleAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTrailerVehicleAssociations`: CursorPageTrailerVehicleAssociationRead
	fmt.Fprintf(os.Stdout, "Response from `FleetOperationsTrackingAPI.ListTrailerVehicleAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTrailerVehicleAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onlyLatest** | **bool** | If true, only the latest known association per trailer is returned. | [default to true]
 **fleetIds** | **[]string** | Limit results to specific fleets using Catena&#39;s fleet IDs. *For your own fleet identifiers, use &#x60;fleet_refs&#x60; instead* To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_ids&#x3D;id1&amp;fleet_ids&#x3D;id2&#x60;). | 
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers. To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_refs&#x3D;ref1&amp;fleet_refs&#x3D;ref2&#x60;). | 
 **connectionId** | **string** | Limit results to a specific provider connection. This is the UUID assigned by Catena when your fleet connects to a TSP. | 
 **fromDatetime** | **time.Time** | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &gt;&#x3D; from_datetime&#x60; **Default value:** &#x60;now() - 1 day&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 45 days | 
 **toDatetime** | **time.Time** | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &lt; to_datetime&#x60; **Default value:** &#x60;now()&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 45 days | 
 **trailerIds** | **[]string** | Limit results to specific trailers. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?trailer_ids&#x3D;id1&amp;trailer_ids&#x3D;id2&#x60;). | 
 **vehicleIds** | **[]string** | Limit results to specific vehicles. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?vehicle_ids&#x3D;id1&amp;vehicle_ids&#x3D;id2&#x60;). | 
 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageTrailerVehicleAssociationRead**](CursorPageTrailerVehicleAssociationRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTrailers

> CursorPageTrailerRead ListTrailers(ctx).FleetIds(fleetIds).FleetRefs(fleetRefs).ConnectionId(connectionId).IncludeSourceData(includeSourceData).TrailerIds(trailerIds).SourceIds(sourceIds).SortBy(sortBy).SortOrder(sortOrder).Cursor(cursor).Size(size).Execute()

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
	fleetIds := []string{"Inner_example"} // []string | Limit results to specific fleets using Catena's fleet IDs. *For your own fleet identifiers, use `fleet_refs` instead* To specify multiple values, repeat the parameter for each value (e.g., `?fleet_ids=id1&fleet_ids=id2`). (optional)
	fleetRefs := []string{"Inner_example"} // []string | Limit results to specific fleets using your organization's fleet reference identifiers. To specify multiple values, repeat the parameter for each value (e.g., `?fleet_refs=ref1&fleet_refs=ref2`). (optional)
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Limit results to a specific provider connection. This is the UUID assigned by Catena when your fleet connects to a TSP. (optional)
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)
	trailerIds := []string{"Inner_example"} // []string | Limit results to specific trailers. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?trailer_ids=id1&trailer_ids=id2`). (optional)
	sourceIds := []string{"Inner_example"} // []string | Limit results to specific resources using the identifier provided by the TSP. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?source_ids=id1&source_ids=id2`). (optional)
	sortBy := "sortBy_example" // string | The name of the field to sort results by. If not provided, results will be ordered by `occurred_at`.  (optional)
	sortOrder := "sortOrder_example" // string | The order of sorting, either `asc` for ascending or `desc` for descending. Defaults to `asc` if `sort_by` is provided without `sort_order`. (optional) (default to "asc")
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetOperationsTrackingAPI.ListTrailers(context.Background()).FleetIds(fleetIds).FleetRefs(fleetRefs).ConnectionId(connectionId).IncludeSourceData(includeSourceData).TrailerIds(trailerIds).SourceIds(sourceIds).SortBy(sortBy).SortOrder(sortOrder).Cursor(cursor).Size(size).Execute()
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
 **fleetIds** | **[]string** | Limit results to specific fleets using Catena&#39;s fleet IDs. *For your own fleet identifiers, use &#x60;fleet_refs&#x60; instead* To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_ids&#x3D;id1&amp;fleet_ids&#x3D;id2&#x60;). | 
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers. To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_refs&#x3D;ref1&amp;fleet_refs&#x3D;ref2&#x60;). | 
 **connectionId** | **string** | Limit results to a specific provider connection. This is the UUID assigned by Catena when your fleet connects to a TSP. | 
 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]
 **trailerIds** | **[]string** | Limit results to specific trailers. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?trailer_ids&#x3D;id1&amp;trailer_ids&#x3D;id2&#x60;). | 
 **sourceIds** | **[]string** | Limit results to specific resources using the identifier provided by the TSP. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?source_ids&#x3D;id1&amp;source_ids&#x3D;id2&#x60;). | 
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

> CursorPageVehicleLocationRead ListVehicleLocations(ctx).OnlyLatest(onlyLatest).FleetIds(fleetIds).FleetRefs(fleetRefs).ConnectionId(connectionId).FromDatetime(fromDatetime).ToDatetime(toDatetime).DriverIds(driverIds).VehicleIds(vehicleIds).IncludeSourceData(includeSourceData).Cursor(cursor).Size(size).Execute()

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
	fleetIds := []string{"Inner_example"} // []string | Limit results to specific fleets using Catena's fleet IDs. *For your own fleet identifiers, use `fleet_refs` instead* To specify multiple values, repeat the parameter for each value (e.g., `?fleet_ids=id1&fleet_ids=id2`). (optional)
	fleetRefs := []string{"Inner_example"} // []string | Limit results to specific fleets using your organization's fleet reference identifiers. To specify multiple values, repeat the parameter for each value (e.g., `?fleet_refs=ref1&fleet_refs=ref2`). (optional)
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Limit results to a specific provider connection. This is the UUID assigned by Catena when your fleet connects to a TSP. (optional)
	fromDatetime := time.Now() // time.Time | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at >= from_datetime` **Default value:** `now() - 1 day` **Restriction:** `to_datetime - from_datetime` cannot exceed 45 days (optional)
	toDatetime := time.Now() // time.Time | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at < to_datetime` **Default value:** `now()` **Restriction:** `to_datetime - from_datetime` cannot exceed 45 days (optional)
	driverIds := []string{"Inner_example"} // []string | Limit results to specific drivers. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?driver_ids=id1&driver_ids=id2`). (optional)
	vehicleIds := []string{"Inner_example"} // []string | Limit results to specific vehicles. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?vehicle_ids=id1&vehicle_ids=id2`). (optional)
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetOperationsTrackingAPI.ListVehicleLocations(context.Background()).OnlyLatest(onlyLatest).FleetIds(fleetIds).FleetRefs(fleetRefs).ConnectionId(connectionId).FromDatetime(fromDatetime).ToDatetime(toDatetime).DriverIds(driverIds).VehicleIds(vehicleIds).IncludeSourceData(includeSourceData).Cursor(cursor).Size(size).Execute()
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
 **fleetIds** | **[]string** | Limit results to specific fleets using Catena&#39;s fleet IDs. *For your own fleet identifiers, use &#x60;fleet_refs&#x60; instead* To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_ids&#x3D;id1&amp;fleet_ids&#x3D;id2&#x60;). | 
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers. To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_refs&#x3D;ref1&amp;fleet_refs&#x3D;ref2&#x60;). | 
 **connectionId** | **string** | Limit results to a specific provider connection. This is the UUID assigned by Catena when your fleet connects to a TSP. | 
 **fromDatetime** | **time.Time** | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &gt;&#x3D; from_datetime&#x60; **Default value:** &#x60;now() - 1 day&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 45 days | 
 **toDatetime** | **time.Time** | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &lt; to_datetime&#x60; **Default value:** &#x60;now()&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 45 days | 
 **driverIds** | **[]string** | Limit results to specific drivers. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?driver_ids&#x3D;id1&amp;driver_ids&#x3D;id2&#x60;). | 
 **vehicleIds** | **[]string** | Limit results to specific vehicles. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?vehicle_ids&#x3D;id1&amp;vehicle_ids&#x3D;id2&#x60;). | 
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


## ListVehicleRegionSegments

> CursorPageVehicleRegionSegmentRead ListVehicleRegionSegments(ctx).ProcessingDate(processingDate).GroupBy(groupBy).FleetIds(fleetIds).FleetRefs(fleetRefs).ConnectionId(connectionId).VehicleIds(vehicleIds).Cursor(cursor).Size(size).Execute()

List Vehicle Region Segments



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
	processingDate := time.Now() // string | Processing date in YYYY-MM-DD format
	groupBy := "groupBy_example" // string | Optional grouping mode. Use `segment` to collapse repeated visits to the same region for a vehicle within the filtered window, using the min start time and max end time. (optional)
	fleetIds := []string{"Inner_example"} // []string | Limit results to specific fleets using Catena's fleet IDs. *For your own fleet identifiers, use `fleet_refs` instead* To specify multiple values, repeat the parameter for each value (e.g., `?fleet_ids=id1&fleet_ids=id2`). (optional)
	fleetRefs := []string{"Inner_example"} // []string | Limit results to specific fleets using your organization's fleet reference identifiers. To specify multiple values, repeat the parameter for each value (e.g., `?fleet_refs=ref1&fleet_refs=ref2`). (optional)
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Limit results to a specific provider connection. This is the UUID assigned by Catena when your fleet connects to a TSP. (optional)
	vehicleIds := []string{"Inner_example"} // []string | Limit results to specific vehicles. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?vehicle_ids=id1&vehicle_ids=id2`). (optional)
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetOperationsTrackingAPI.ListVehicleRegionSegments(context.Background()).ProcessingDate(processingDate).GroupBy(groupBy).FleetIds(fleetIds).FleetRefs(fleetRefs).ConnectionId(connectionId).VehicleIds(vehicleIds).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetOperationsTrackingAPI.ListVehicleRegionSegments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVehicleRegionSegments`: CursorPageVehicleRegionSegmentRead
	fmt.Fprintf(os.Stdout, "Response from `FleetOperationsTrackingAPI.ListVehicleRegionSegments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVehicleRegionSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processingDate** | **string** | Processing date in YYYY-MM-DD format | 
 **groupBy** | **string** | Optional grouping mode. Use &#x60;segment&#x60; to collapse repeated visits to the same region for a vehicle within the filtered window, using the min start time and max end time. | 
 **fleetIds** | **[]string** | Limit results to specific fleets using Catena&#39;s fleet IDs. *For your own fleet identifiers, use &#x60;fleet_refs&#x60; instead* To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_ids&#x3D;id1&amp;fleet_ids&#x3D;id2&#x60;). | 
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers. To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_refs&#x3D;ref1&amp;fleet_refs&#x3D;ref2&#x60;). | 
 **connectionId** | **string** | Limit results to a specific provider connection. This is the UUID assigned by Catena when your fleet connects to a TSP. | 
 **vehicleIds** | **[]string** | Limit results to specific vehicles. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?vehicle_ids&#x3D;id1&amp;vehicle_ids&#x3D;id2&#x60;). | 
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageVehicleRegionSegmentRead**](CursorPageVehicleRegionSegmentRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVehicles

> CursorPageVehicleRead ListVehicles(ctx).FleetIds(fleetIds).FleetRefs(fleetRefs).ConnectionId(connectionId).IncludeSourceData(includeSourceData).VehicleIds(vehicleIds).SourceIds(sourceIds).SortBy(sortBy).SortOrder(sortOrder).Cursor(cursor).Size(size).Execute()

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
	fleetIds := []string{"Inner_example"} // []string | Limit results to specific fleets using Catena's fleet IDs. *For your own fleet identifiers, use `fleet_refs` instead* To specify multiple values, repeat the parameter for each value (e.g., `?fleet_ids=id1&fleet_ids=id2`). (optional)
	fleetRefs := []string{"Inner_example"} // []string | Limit results to specific fleets using your organization's fleet reference identifiers. To specify multiple values, repeat the parameter for each value (e.g., `?fleet_refs=ref1&fleet_refs=ref2`). (optional)
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Limit results to a specific provider connection. This is the UUID assigned by Catena when your fleet connects to a TSP. (optional)
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)
	vehicleIds := []string{"Inner_example"} // []string | Limit results to specific vehicles. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?vehicle_ids=id1&vehicle_ids=id2`). (optional)
	sourceIds := []string{"Inner_example"} // []string | Limit results to specific resources using the identifier provided by the TSP. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., `?source_ids=id1&source_ids=id2`). (optional)
	sortBy := "sortBy_example" // string | The name of the field to sort results by. If not provided, results will be ordered by `occurred_at`.  (optional)
	sortOrder := "sortOrder_example" // string | The order of sorting, either `asc` for ascending or `desc` for descending. Defaults to `asc` if `sort_by` is provided without `sort_order`. (optional) (default to "asc")
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetOperationsTrackingAPI.ListVehicles(context.Background()).FleetIds(fleetIds).FleetRefs(fleetRefs).ConnectionId(connectionId).IncludeSourceData(includeSourceData).VehicleIds(vehicleIds).SourceIds(sourceIds).SortBy(sortBy).SortOrder(sortOrder).Cursor(cursor).Size(size).Execute()
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
 **fleetIds** | **[]string** | Limit results to specific fleets using Catena&#39;s fleet IDs. *For your own fleet identifiers, use &#x60;fleet_refs&#x60; instead* To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_ids&#x3D;id1&amp;fleet_ids&#x3D;id2&#x60;). | 
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers. To specify multiple values, repeat the parameter for each value (e.g., &#x60;?fleet_refs&#x3D;ref1&amp;fleet_refs&#x3D;ref2&#x60;). | 
 **connectionId** | **string** | Limit results to a specific provider connection. This is the UUID assigned by Catena when your fleet connects to a TSP. | 
 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]
 **vehicleIds** | **[]string** | Limit results to specific vehicles. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?vehicle_ids&#x3D;id1&amp;vehicle_ids&#x3D;id2&#x60;). | 
 **sourceIds** | **[]string** | Limit results to specific resources using the identifier provided by the TSP. **Maximum:** 100 IDs To specify multiple values, repeat the parameter for each value (e.g., &#x60;?source_ids&#x3D;id1&amp;source_ids&#x3D;id2&#x60;). | 
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


## UpdateVehicle

> ResourceOperationAccept UpdateVehicle(ctx, sourceId).VehicleUpdate(vehicleUpdate).Execute()

Update Vehicle



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
	sourceId := "sourceId_example" // string | The unique identifier of the vehicle in the TSP.
	vehicleUpdate := *openapiclient.NewVehicleUpdate("ConnectionId_example") // VehicleUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetOperationsTrackingAPI.UpdateVehicle(context.Background(), sourceId).VehicleUpdate(vehicleUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetOperationsTrackingAPI.UpdateVehicle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVehicle`: ResourceOperationAccept
	fmt.Fprintf(os.Stdout, "Response from `FleetOperationsTrackingAPI.UpdateVehicle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The unique identifier of the vehicle in the TSP. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVehicleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vehicleUpdate** | [**VehicleUpdate**](VehicleUpdate.md) |  | 

### Return type

[**ResourceOperationAccept**](ResourceOperationAccept.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

