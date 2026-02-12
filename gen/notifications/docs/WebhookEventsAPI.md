# \WebhookEventsAPI

All URIs are relative to *https://api.catenatelematics.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectionCreatedconnectionCreatedPost**](WebhookEventsAPI.md#ConnectionCreatedconnectionCreatedPost) | **Post** /connection.created | Connection Created
[**ConnectionStaledconnectionStaledPost**](WebhookEventsAPI.md#ConnectionStaledconnectionStaledPost) | **Post** /connection.staled | Connection Staled
[**ExecutionFailedexecutionFailedPost**](WebhookEventsAPI.md#ExecutionFailedexecutionFailedPost) | **Post** /execution.failed | Execution Failed
[**ExecutionStaledexecutionStaledPost**](WebhookEventsAPI.md#ExecutionStaledexecutionStaledPost) | **Post** /execution.staled | Execution Staled
[**FleetConnectionCreatedfleetConnectionCreatedPost**](WebhookEventsAPI.md#FleetConnectionCreatedfleetConnectionCreatedPost) | **Post** /fleet_connection.created | Fleet Connection Created
[**HosAvailabilityAddedhosAvailabilityAddedPost**](WebhookEventsAPI.md#HosAvailabilityAddedhosAvailabilityAddedPost) | **Post** /hos_availability.added | Hos Availability Added
[**HosAvailabilityModifiedhosAvailabilityModifiedPost**](WebhookEventsAPI.md#HosAvailabilityModifiedhosAvailabilityModifiedPost) | **Post** /hos_availability.modified | Hos Availability Modified
[**HosAvailabilityRemovedhosAvailabilityRemovedPost**](WebhookEventsAPI.md#HosAvailabilityRemovedhosAvailabilityRemovedPost) | **Post** /hos_availability.removed | Hos Availability Removed
[**HosDailySnapshotAddedhosDailySnapshotAddedPost**](WebhookEventsAPI.md#HosDailySnapshotAddedhosDailySnapshotAddedPost) | **Post** /hos_daily_snapshot.added | Hos Daily Snapshot Added
[**HosDailySnapshotModifiedhosDailySnapshotModifiedPost**](WebhookEventsAPI.md#HosDailySnapshotModifiedhosDailySnapshotModifiedPost) | **Post** /hos_daily_snapshot.modified | Hos Daily Snapshot Modified
[**HosDailySnapshotRemovedhosDailySnapshotRemovedPost**](WebhookEventsAPI.md#HosDailySnapshotRemovedhosDailySnapshotRemovedPost) | **Post** /hos_daily_snapshot.removed | Hos Daily Snapshot Removed
[**HosEventAddedhosEventAddedPost**](WebhookEventsAPI.md#HosEventAddedhosEventAddedPost) | **Post** /hos_event.added | Hos Event Added
[**HosEventModifiedhosEventModifiedPost**](WebhookEventsAPI.md#HosEventModifiedhosEventModifiedPost) | **Post** /hos_event.modified | Hos Event Modified
[**HosEventRemovedhosEventRemovedPost**](WebhookEventsAPI.md#HosEventRemovedhosEventRemovedPost) | **Post** /hos_event.removed | Hos Event Removed
[**HosViolationAddedhosViolationAddedPost**](WebhookEventsAPI.md#HosViolationAddedhosViolationAddedPost) | **Post** /hos_violation.added | Hos Violation Added
[**HosViolationModifiedhosViolationModifiedPost**](WebhookEventsAPI.md#HosViolationModifiedhosViolationModifiedPost) | **Post** /hos_violation.modified | Hos Violation Modified
[**HosViolationRemovedhosViolationRemovedPost**](WebhookEventsAPI.md#HosViolationRemovedhosViolationRemovedPost) | **Post** /hos_violation.removed | Hos Violation Removed
[**IftaSummaryAddediftaSummaryAddedPost**](WebhookEventsAPI.md#IftaSummaryAddediftaSummaryAddedPost) | **Post** /ifta_summary.added | Ifta Summary Added
[**IftaSummaryModifiediftaSummaryModifiedPost**](WebhookEventsAPI.md#IftaSummaryModifiediftaSummaryModifiedPost) | **Post** /ifta_summary.modified | Ifta Summary Modified
[**IftaSummaryRemovediftaSummaryRemovedPost**](WebhookEventsAPI.md#IftaSummaryRemovediftaSummaryRemovedPost) | **Post** /ifta_summary.removed | Ifta Summary Removed
[**InvitationAcceptedinvitationAcceptedPost**](WebhookEventsAPI.md#InvitationAcceptedinvitationAcceptedPost) | **Post** /invitation.accepted | Invitation Accepted
[**InvitationCreatedinvitationCreatedPost**](WebhookEventsAPI.md#InvitationCreatedinvitationCreatedPost) | **Post** /invitation.created | Invitation Created
[**InvitationDeclinedinvitationDeclinedPost**](WebhookEventsAPI.md#InvitationDeclinedinvitationDeclinedPost) | **Post** /invitation.declined | Invitation Declined
[**InvitationDeletedinvitationDeletedPost**](WebhookEventsAPI.md#InvitationDeletedinvitationDeletedPost) | **Post** /invitation.deleted | Invitation Deleted
[**InvitationExpiredinvitationExpiredPost**](WebhookEventsAPI.md#InvitationExpiredinvitationExpiredPost) | **Post** /invitation.expired | Invitation Expired
[**InvitationRejectedinvitationRejectedPost**](WebhookEventsAPI.md#InvitationRejectedinvitationRejectedPost) | **Post** /invitation.rejected | Invitation Rejected
[**InvitationSentinvitationSentPost**](WebhookEventsAPI.md#InvitationSentinvitationSentPost) | **Post** /invitation.sent | Invitation Sent
[**InvitationViewedinvitationViewedPost**](WebhookEventsAPI.md#InvitationViewedinvitationViewedPost) | **Post** /invitation.viewed | Invitation Viewed
[**ScheduleDeactivatedscheduleDeactivatedPost**](WebhookEventsAPI.md#ScheduleDeactivatedscheduleDeactivatedPost) | **Post** /schedule.deactivated | Schedule Deactivated
[**ShareAgreementCreatedshareAgreementCreatedPost**](WebhookEventsAPI.md#ShareAgreementCreatedshareAgreementCreatedPost) | **Post** /share_agreement.created | Share Agreement Created
[**ShareAgreementDeletedshareAgreementDeletedPost**](WebhookEventsAPI.md#ShareAgreementDeletedshareAgreementDeletedPost) | **Post** /share_agreement.deleted | Share Agreement Deleted
[**ShareAgreementUpdatedshareAgreementUpdatedPost**](WebhookEventsAPI.md#ShareAgreementUpdatedshareAgreementUpdatedPost) | **Post** /share_agreement.updated | Share Agreement Updated
[**TrailerAddedtrailerAddedPost**](WebhookEventsAPI.md#TrailerAddedtrailerAddedPost) | **Post** /trailer.added | Trailer Added
[**TrailerLocationAddedtrailerLocationAddedPost**](WebhookEventsAPI.md#TrailerLocationAddedtrailerLocationAddedPost) | **Post** /trailer_location.added | Trailer Location Added
[**TrailerLocationModifiedtrailerLocationModifiedPost**](WebhookEventsAPI.md#TrailerLocationModifiedtrailerLocationModifiedPost) | **Post** /trailer_location.modified | Trailer Location Modified
[**TrailerModifiedtrailerModifiedPost**](WebhookEventsAPI.md#TrailerModifiedtrailerModifiedPost) | **Post** /trailer.modified | Trailer Modified
[**TrailerRemovedtrailerRemovedPost**](WebhookEventsAPI.md#TrailerRemovedtrailerRemovedPost) | **Post** /trailer.removed | Trailer Removed
[**TspCreatedtspCreatedPost**](WebhookEventsAPI.md#TspCreatedtspCreatedPost) | **Post** /tsp.created | Tsp Created
[**UserAddeduserAddedPost**](WebhookEventsAPI.md#UserAddeduserAddedPost) | **Post** /user.added | User Added
[**UserModifieduserModifiedPost**](WebhookEventsAPI.md#UserModifieduserModifiedPost) | **Post** /user.modified | User Modified
[**UserRemoveduserRemovedPost**](WebhookEventsAPI.md#UserRemoveduserRemovedPost) | **Post** /user.removed | User Removed
[**VehicleAddedvehicleAddedPost**](WebhookEventsAPI.md#VehicleAddedvehicleAddedPost) | **Post** /vehicle.added | Vehicle Added
[**VehicleLocationAddedvehicleLocationAddedPost**](WebhookEventsAPI.md#VehicleLocationAddedvehicleLocationAddedPost) | **Post** /vehicle_location.added | Vehicle Location Added
[**VehicleModifiedvehicleModifiedPost**](WebhookEventsAPI.md#VehicleModifiedvehicleModifiedPost) | **Post** /vehicle.modified | Vehicle Modified
[**VehicleRemovedvehicleRemovedPost**](WebhookEventsAPI.md#VehicleRemovedvehicleRemovedPost) | **Post** /vehicle.removed | Vehicle Removed
[**WebhookCreatedwebhookCreatedPost**](WebhookEventsAPI.md#WebhookCreatedwebhookCreatedPost) | **Post** /webhook.created | Webhook Created
[**WebhookDeletedwebhookDeletedPost**](WebhookEventsAPI.md#WebhookDeletedwebhookDeletedPost) | **Post** /webhook.deleted | Webhook Deleted
[**WebhookStaledwebhookStaledPost**](WebhookEventsAPI.md#WebhookStaledwebhookStaledPost) | **Post** /webhook.staled | Webhook Staled
[**WebhookUpdatedwebhookUpdatedPost**](WebhookEventsAPI.md#WebhookUpdatedwebhookUpdatedPost) | **Post** /webhook.updated | Webhook Updated



## ConnectionCreatedconnectionCreatedPost

> interface{} ConnectionCreatedconnectionCreatedPost(ctx).ConnectionCreated(connectionCreated).Execute()

Connection Created



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
	connectionCreated :=  // ConnectionCreated | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.ConnectionCreatedconnectionCreatedPost(context.Background()).ConnectionCreated(connectionCreated).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.ConnectionCreatedconnectionCreatedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionCreatedconnectionCreatedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.ConnectionCreatedconnectionCreatedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConnectionCreatedconnectionCreatedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectionCreated** | [**ConnectionCreated**](ConnectionCreated.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionStaledconnectionStaledPost

> interface{} ConnectionStaledconnectionStaledPost(ctx).ConnectionStaled(connectionStaled).Execute()

Connection Staled



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
	connectionStaled :=  // ConnectionStaled | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.ConnectionStaledconnectionStaledPost(context.Background()).ConnectionStaled(connectionStaled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.ConnectionStaledconnectionStaledPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionStaledconnectionStaledPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.ConnectionStaledconnectionStaledPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConnectionStaledconnectionStaledPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectionStaled** | [**ConnectionStaled**](ConnectionStaled.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecutionFailedexecutionFailedPost

> interface{} ExecutionFailedexecutionFailedPost(ctx).ExecutionFailed(executionFailed).Execute()

Execution Failed



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
	executionFailed :=  // ExecutionFailed | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.ExecutionFailedexecutionFailedPost(context.Background()).ExecutionFailed(executionFailed).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.ExecutionFailedexecutionFailedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecutionFailedexecutionFailedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.ExecutionFailedexecutionFailedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecutionFailedexecutionFailedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **executionFailed** | [**ExecutionFailed**](ExecutionFailed.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecutionStaledexecutionStaledPost

> interface{} ExecutionStaledexecutionStaledPost(ctx).ExecutionStaled(executionStaled).Execute()

Execution Staled



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
	executionStaled :=  // ExecutionStaled | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.ExecutionStaledexecutionStaledPost(context.Background()).ExecutionStaled(executionStaled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.ExecutionStaledexecutionStaledPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecutionStaledexecutionStaledPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.ExecutionStaledexecutionStaledPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecutionStaledexecutionStaledPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **executionStaled** | [**ExecutionStaled**](ExecutionStaled.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FleetConnectionCreatedfleetConnectionCreatedPost

> interface{} FleetConnectionCreatedfleetConnectionCreatedPost(ctx).FleetConnectionCreated(fleetConnectionCreated).Execute()

Fleet Connection Created



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
	fleetConnectionCreated :=  // FleetConnectionCreated | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.FleetConnectionCreatedfleetConnectionCreatedPost(context.Background()).FleetConnectionCreated(fleetConnectionCreated).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.FleetConnectionCreatedfleetConnectionCreatedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FleetConnectionCreatedfleetConnectionCreatedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.FleetConnectionCreatedfleetConnectionCreatedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFleetConnectionCreatedfleetConnectionCreatedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetConnectionCreated** | [**FleetConnectionCreated**](FleetConnectionCreated.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HosAvailabilityAddedhosAvailabilityAddedPost

> interface{} HosAvailabilityAddedhosAvailabilityAddedPost(ctx).HosAvailabilityAdded(hosAvailabilityAdded).Execute()

Hos Availability Added



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
	hosAvailabilityAdded :=  // HosAvailabilityAdded | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.HosAvailabilityAddedhosAvailabilityAddedPost(context.Background()).HosAvailabilityAdded(hosAvailabilityAdded).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.HosAvailabilityAddedhosAvailabilityAddedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HosAvailabilityAddedhosAvailabilityAddedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.HosAvailabilityAddedhosAvailabilityAddedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHosAvailabilityAddedhosAvailabilityAddedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hosAvailabilityAdded** | [**HosAvailabilityAdded**](HosAvailabilityAdded.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HosAvailabilityModifiedhosAvailabilityModifiedPost

> interface{} HosAvailabilityModifiedhosAvailabilityModifiedPost(ctx).HosAvailabilityModified(hosAvailabilityModified).Execute()

Hos Availability Modified



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
	hosAvailabilityModified :=  // HosAvailabilityModified | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.HosAvailabilityModifiedhosAvailabilityModifiedPost(context.Background()).HosAvailabilityModified(hosAvailabilityModified).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.HosAvailabilityModifiedhosAvailabilityModifiedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HosAvailabilityModifiedhosAvailabilityModifiedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.HosAvailabilityModifiedhosAvailabilityModifiedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHosAvailabilityModifiedhosAvailabilityModifiedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hosAvailabilityModified** | [**HosAvailabilityModified**](HosAvailabilityModified.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HosAvailabilityRemovedhosAvailabilityRemovedPost

> interface{} HosAvailabilityRemovedhosAvailabilityRemovedPost(ctx).HosAvailabilityRemoved(hosAvailabilityRemoved).Execute()

Hos Availability Removed



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
	hosAvailabilityRemoved :=  // HosAvailabilityRemoved | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.HosAvailabilityRemovedhosAvailabilityRemovedPost(context.Background()).HosAvailabilityRemoved(hosAvailabilityRemoved).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.HosAvailabilityRemovedhosAvailabilityRemovedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HosAvailabilityRemovedhosAvailabilityRemovedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.HosAvailabilityRemovedhosAvailabilityRemovedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHosAvailabilityRemovedhosAvailabilityRemovedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hosAvailabilityRemoved** | [**HosAvailabilityRemoved**](HosAvailabilityRemoved.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HosDailySnapshotAddedhosDailySnapshotAddedPost

> interface{} HosDailySnapshotAddedhosDailySnapshotAddedPost(ctx).HosDailySnapshotAdded(hosDailySnapshotAdded).Execute()

Hos Daily Snapshot Added



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
	hosDailySnapshotAdded :=  // HosDailySnapshotAdded | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.HosDailySnapshotAddedhosDailySnapshotAddedPost(context.Background()).HosDailySnapshotAdded(hosDailySnapshotAdded).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.HosDailySnapshotAddedhosDailySnapshotAddedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HosDailySnapshotAddedhosDailySnapshotAddedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.HosDailySnapshotAddedhosDailySnapshotAddedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHosDailySnapshotAddedhosDailySnapshotAddedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hosDailySnapshotAdded** | [**HosDailySnapshotAdded**](HosDailySnapshotAdded.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HosDailySnapshotModifiedhosDailySnapshotModifiedPost

> interface{} HosDailySnapshotModifiedhosDailySnapshotModifiedPost(ctx).HosDailySnapshotModified(hosDailySnapshotModified).Execute()

Hos Daily Snapshot Modified



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
	hosDailySnapshotModified :=  // HosDailySnapshotModified | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.HosDailySnapshotModifiedhosDailySnapshotModifiedPost(context.Background()).HosDailySnapshotModified(hosDailySnapshotModified).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.HosDailySnapshotModifiedhosDailySnapshotModifiedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HosDailySnapshotModifiedhosDailySnapshotModifiedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.HosDailySnapshotModifiedhosDailySnapshotModifiedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHosDailySnapshotModifiedhosDailySnapshotModifiedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hosDailySnapshotModified** | [**HosDailySnapshotModified**](HosDailySnapshotModified.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HosDailySnapshotRemovedhosDailySnapshotRemovedPost

> interface{} HosDailySnapshotRemovedhosDailySnapshotRemovedPost(ctx).HosDailySnapshotRemoved(hosDailySnapshotRemoved).Execute()

Hos Daily Snapshot Removed



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
	hosDailySnapshotRemoved :=  // HosDailySnapshotRemoved | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.HosDailySnapshotRemovedhosDailySnapshotRemovedPost(context.Background()).HosDailySnapshotRemoved(hosDailySnapshotRemoved).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.HosDailySnapshotRemovedhosDailySnapshotRemovedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HosDailySnapshotRemovedhosDailySnapshotRemovedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.HosDailySnapshotRemovedhosDailySnapshotRemovedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHosDailySnapshotRemovedhosDailySnapshotRemovedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hosDailySnapshotRemoved** | [**HosDailySnapshotRemoved**](HosDailySnapshotRemoved.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HosEventAddedhosEventAddedPost

> interface{} HosEventAddedhosEventAddedPost(ctx).HosEventAdded(hosEventAdded).Execute()

Hos Event Added



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
	hosEventAdded :=  // HosEventAdded | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.HosEventAddedhosEventAddedPost(context.Background()).HosEventAdded(hosEventAdded).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.HosEventAddedhosEventAddedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HosEventAddedhosEventAddedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.HosEventAddedhosEventAddedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHosEventAddedhosEventAddedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hosEventAdded** | [**HosEventAdded**](HosEventAdded.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HosEventModifiedhosEventModifiedPost

> interface{} HosEventModifiedhosEventModifiedPost(ctx).HosEventModified(hosEventModified).Execute()

Hos Event Modified



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
	hosEventModified :=  // HosEventModified | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.HosEventModifiedhosEventModifiedPost(context.Background()).HosEventModified(hosEventModified).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.HosEventModifiedhosEventModifiedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HosEventModifiedhosEventModifiedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.HosEventModifiedhosEventModifiedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHosEventModifiedhosEventModifiedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hosEventModified** | [**HosEventModified**](HosEventModified.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HosEventRemovedhosEventRemovedPost

> interface{} HosEventRemovedhosEventRemovedPost(ctx).HosEventRemoved(hosEventRemoved).Execute()

Hos Event Removed



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
	hosEventRemoved :=  // HosEventRemoved | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.HosEventRemovedhosEventRemovedPost(context.Background()).HosEventRemoved(hosEventRemoved).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.HosEventRemovedhosEventRemovedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HosEventRemovedhosEventRemovedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.HosEventRemovedhosEventRemovedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHosEventRemovedhosEventRemovedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hosEventRemoved** | [**HosEventRemoved**](HosEventRemoved.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HosViolationAddedhosViolationAddedPost

> interface{} HosViolationAddedhosViolationAddedPost(ctx).HosViolationAdded(hosViolationAdded).Execute()

Hos Violation Added



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
	hosViolationAdded :=  // HosViolationAdded | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.HosViolationAddedhosViolationAddedPost(context.Background()).HosViolationAdded(hosViolationAdded).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.HosViolationAddedhosViolationAddedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HosViolationAddedhosViolationAddedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.HosViolationAddedhosViolationAddedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHosViolationAddedhosViolationAddedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hosViolationAdded** | [**HosViolationAdded**](HosViolationAdded.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HosViolationModifiedhosViolationModifiedPost

> interface{} HosViolationModifiedhosViolationModifiedPost(ctx).HosViolationModified(hosViolationModified).Execute()

Hos Violation Modified



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
	hosViolationModified :=  // HosViolationModified | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.HosViolationModifiedhosViolationModifiedPost(context.Background()).HosViolationModified(hosViolationModified).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.HosViolationModifiedhosViolationModifiedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HosViolationModifiedhosViolationModifiedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.HosViolationModifiedhosViolationModifiedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHosViolationModifiedhosViolationModifiedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hosViolationModified** | [**HosViolationModified**](HosViolationModified.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HosViolationRemovedhosViolationRemovedPost

> interface{} HosViolationRemovedhosViolationRemovedPost(ctx).HosViolationRemoved(hosViolationRemoved).Execute()

Hos Violation Removed



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
	hosViolationRemoved :=  // HosViolationRemoved | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.HosViolationRemovedhosViolationRemovedPost(context.Background()).HosViolationRemoved(hosViolationRemoved).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.HosViolationRemovedhosViolationRemovedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HosViolationRemovedhosViolationRemovedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.HosViolationRemovedhosViolationRemovedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHosViolationRemovedhosViolationRemovedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hosViolationRemoved** | [**HosViolationRemoved**](HosViolationRemoved.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IftaSummaryAddediftaSummaryAddedPost

> interface{} IftaSummaryAddediftaSummaryAddedPost(ctx).IftaSummaryAdded(iftaSummaryAdded).Execute()

Ifta Summary Added



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
	iftaSummaryAdded :=  // IftaSummaryAdded | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.IftaSummaryAddediftaSummaryAddedPost(context.Background()).IftaSummaryAdded(iftaSummaryAdded).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.IftaSummaryAddediftaSummaryAddedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IftaSummaryAddediftaSummaryAddedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.IftaSummaryAddediftaSummaryAddedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIftaSummaryAddediftaSummaryAddedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iftaSummaryAdded** | [**IftaSummaryAdded**](IftaSummaryAdded.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IftaSummaryModifiediftaSummaryModifiedPost

> interface{} IftaSummaryModifiediftaSummaryModifiedPost(ctx).IftaSummaryModified(iftaSummaryModified).Execute()

Ifta Summary Modified



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
	iftaSummaryModified :=  // IftaSummaryModified | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.IftaSummaryModifiediftaSummaryModifiedPost(context.Background()).IftaSummaryModified(iftaSummaryModified).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.IftaSummaryModifiediftaSummaryModifiedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IftaSummaryModifiediftaSummaryModifiedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.IftaSummaryModifiediftaSummaryModifiedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIftaSummaryModifiediftaSummaryModifiedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iftaSummaryModified** | [**IftaSummaryModified**](IftaSummaryModified.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IftaSummaryRemovediftaSummaryRemovedPost

> interface{} IftaSummaryRemovediftaSummaryRemovedPost(ctx).IftaSummaryRemoved(iftaSummaryRemoved).Execute()

Ifta Summary Removed



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
	iftaSummaryRemoved :=  // IftaSummaryRemoved | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.IftaSummaryRemovediftaSummaryRemovedPost(context.Background()).IftaSummaryRemoved(iftaSummaryRemoved).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.IftaSummaryRemovediftaSummaryRemovedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IftaSummaryRemovediftaSummaryRemovedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.IftaSummaryRemovediftaSummaryRemovedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIftaSummaryRemovediftaSummaryRemovedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iftaSummaryRemoved** | [**IftaSummaryRemoved**](IftaSummaryRemoved.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvitationAcceptedinvitationAcceptedPost

> interface{} InvitationAcceptedinvitationAcceptedPost(ctx).InvitationAccepted(invitationAccepted).Execute()

Invitation Accepted



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
	invitationAccepted :=  // InvitationAccepted | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.InvitationAcceptedinvitationAcceptedPost(context.Background()).InvitationAccepted(invitationAccepted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.InvitationAcceptedinvitationAcceptedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvitationAcceptedinvitationAcceptedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.InvitationAcceptedinvitationAcceptedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvitationAcceptedinvitationAcceptedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invitationAccepted** | [**InvitationAccepted**](InvitationAccepted.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvitationCreatedinvitationCreatedPost

> interface{} InvitationCreatedinvitationCreatedPost(ctx).InvitationCreated(invitationCreated).Execute()

Invitation Created



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
	invitationCreated :=  // InvitationCreated | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.InvitationCreatedinvitationCreatedPost(context.Background()).InvitationCreated(invitationCreated).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.InvitationCreatedinvitationCreatedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvitationCreatedinvitationCreatedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.InvitationCreatedinvitationCreatedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvitationCreatedinvitationCreatedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invitationCreated** | [**InvitationCreated**](InvitationCreated.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvitationDeclinedinvitationDeclinedPost

> interface{} InvitationDeclinedinvitationDeclinedPost(ctx).InvitationDeclined(invitationDeclined).Execute()

Invitation Declined



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
	invitationDeclined :=  // InvitationDeclined | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.InvitationDeclinedinvitationDeclinedPost(context.Background()).InvitationDeclined(invitationDeclined).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.InvitationDeclinedinvitationDeclinedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvitationDeclinedinvitationDeclinedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.InvitationDeclinedinvitationDeclinedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvitationDeclinedinvitationDeclinedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invitationDeclined** | [**InvitationDeclined**](InvitationDeclined.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvitationDeletedinvitationDeletedPost

> interface{} InvitationDeletedinvitationDeletedPost(ctx).InvitationDeleted(invitationDeleted).Execute()

Invitation Deleted



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
	invitationDeleted :=  // InvitationDeleted | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.InvitationDeletedinvitationDeletedPost(context.Background()).InvitationDeleted(invitationDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.InvitationDeletedinvitationDeletedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvitationDeletedinvitationDeletedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.InvitationDeletedinvitationDeletedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvitationDeletedinvitationDeletedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invitationDeleted** | [**InvitationDeleted**](InvitationDeleted.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvitationExpiredinvitationExpiredPost

> interface{} InvitationExpiredinvitationExpiredPost(ctx).InvitationExpired(invitationExpired).Execute()

Invitation Expired



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
	invitationExpired :=  // InvitationExpired | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.InvitationExpiredinvitationExpiredPost(context.Background()).InvitationExpired(invitationExpired).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.InvitationExpiredinvitationExpiredPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvitationExpiredinvitationExpiredPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.InvitationExpiredinvitationExpiredPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvitationExpiredinvitationExpiredPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invitationExpired** | [**InvitationExpired**](InvitationExpired.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvitationRejectedinvitationRejectedPost

> interface{} InvitationRejectedinvitationRejectedPost(ctx).InvitationRejected(invitationRejected).Execute()

Invitation Rejected



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
	invitationRejected :=  // InvitationRejected | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.InvitationRejectedinvitationRejectedPost(context.Background()).InvitationRejected(invitationRejected).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.InvitationRejectedinvitationRejectedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvitationRejectedinvitationRejectedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.InvitationRejectedinvitationRejectedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvitationRejectedinvitationRejectedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invitationRejected** | [**InvitationRejected**](InvitationRejected.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvitationSentinvitationSentPost

> interface{} InvitationSentinvitationSentPost(ctx).InvitationSent(invitationSent).Execute()

Invitation Sent



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
	invitationSent :=  // InvitationSent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.InvitationSentinvitationSentPost(context.Background()).InvitationSent(invitationSent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.InvitationSentinvitationSentPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvitationSentinvitationSentPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.InvitationSentinvitationSentPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvitationSentinvitationSentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invitationSent** | [**InvitationSent**](InvitationSent.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvitationViewedinvitationViewedPost

> interface{} InvitationViewedinvitationViewedPost(ctx).InvitationViewed(invitationViewed).Execute()

Invitation Viewed



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
	invitationViewed :=  // InvitationViewed | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.InvitationViewedinvitationViewedPost(context.Background()).InvitationViewed(invitationViewed).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.InvitationViewedinvitationViewedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvitationViewedinvitationViewedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.InvitationViewedinvitationViewedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvitationViewedinvitationViewedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invitationViewed** | [**InvitationViewed**](InvitationViewed.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleDeactivatedscheduleDeactivatedPost

> interface{} ScheduleDeactivatedscheduleDeactivatedPost(ctx).ScheduleDeactivated(scheduleDeactivated).Execute()

Schedule Deactivated



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
	scheduleDeactivated :=  // ScheduleDeactivated | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.ScheduleDeactivatedscheduleDeactivatedPost(context.Background()).ScheduleDeactivated(scheduleDeactivated).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.ScheduleDeactivatedscheduleDeactivatedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScheduleDeactivatedscheduleDeactivatedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.ScheduleDeactivatedscheduleDeactivatedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScheduleDeactivatedscheduleDeactivatedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scheduleDeactivated** | [**ScheduleDeactivated**](ScheduleDeactivated.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShareAgreementCreatedshareAgreementCreatedPost

> interface{} ShareAgreementCreatedshareAgreementCreatedPost(ctx).ShareAgreementCreated(shareAgreementCreated).Execute()

Share Agreement Created



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
	shareAgreementCreated :=  // ShareAgreementCreated | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.ShareAgreementCreatedshareAgreementCreatedPost(context.Background()).ShareAgreementCreated(shareAgreementCreated).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.ShareAgreementCreatedshareAgreementCreatedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShareAgreementCreatedshareAgreementCreatedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.ShareAgreementCreatedshareAgreementCreatedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShareAgreementCreatedshareAgreementCreatedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shareAgreementCreated** | [**ShareAgreementCreated**](ShareAgreementCreated.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShareAgreementDeletedshareAgreementDeletedPost

> interface{} ShareAgreementDeletedshareAgreementDeletedPost(ctx).ShareAgreementDeleted(shareAgreementDeleted).Execute()

Share Agreement Deleted



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
	shareAgreementDeleted :=  // ShareAgreementDeleted | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.ShareAgreementDeletedshareAgreementDeletedPost(context.Background()).ShareAgreementDeleted(shareAgreementDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.ShareAgreementDeletedshareAgreementDeletedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShareAgreementDeletedshareAgreementDeletedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.ShareAgreementDeletedshareAgreementDeletedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShareAgreementDeletedshareAgreementDeletedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shareAgreementDeleted** | [**ShareAgreementDeleted**](ShareAgreementDeleted.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShareAgreementUpdatedshareAgreementUpdatedPost

> interface{} ShareAgreementUpdatedshareAgreementUpdatedPost(ctx).ShareAgreementUpdated(shareAgreementUpdated).Execute()

Share Agreement Updated



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
	shareAgreementUpdated :=  // ShareAgreementUpdated | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.ShareAgreementUpdatedshareAgreementUpdatedPost(context.Background()).ShareAgreementUpdated(shareAgreementUpdated).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.ShareAgreementUpdatedshareAgreementUpdatedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShareAgreementUpdatedshareAgreementUpdatedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.ShareAgreementUpdatedshareAgreementUpdatedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShareAgreementUpdatedshareAgreementUpdatedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shareAgreementUpdated** | [**ShareAgreementUpdated**](ShareAgreementUpdated.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrailerAddedtrailerAddedPost

> interface{} TrailerAddedtrailerAddedPost(ctx).TrailerAdded(trailerAdded).Execute()

Trailer Added



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
	trailerAdded :=  // TrailerAdded | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.TrailerAddedtrailerAddedPost(context.Background()).TrailerAdded(trailerAdded).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.TrailerAddedtrailerAddedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrailerAddedtrailerAddedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.TrailerAddedtrailerAddedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrailerAddedtrailerAddedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trailerAdded** | [**TrailerAdded**](TrailerAdded.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrailerLocationAddedtrailerLocationAddedPost

> interface{} TrailerLocationAddedtrailerLocationAddedPost(ctx).TrailerLocationAdded(trailerLocationAdded).Execute()

Trailer Location Added



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
	trailerLocationAdded :=  // TrailerLocationAdded | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.TrailerLocationAddedtrailerLocationAddedPost(context.Background()).TrailerLocationAdded(trailerLocationAdded).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.TrailerLocationAddedtrailerLocationAddedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrailerLocationAddedtrailerLocationAddedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.TrailerLocationAddedtrailerLocationAddedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrailerLocationAddedtrailerLocationAddedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trailerLocationAdded** | [**TrailerLocationAdded**](TrailerLocationAdded.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrailerLocationModifiedtrailerLocationModifiedPost

> interface{} TrailerLocationModifiedtrailerLocationModifiedPost(ctx).TrailerLocationModified(trailerLocationModified).Execute()

Trailer Location Modified



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
	trailerLocationModified :=  // TrailerLocationModified | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.TrailerLocationModifiedtrailerLocationModifiedPost(context.Background()).TrailerLocationModified(trailerLocationModified).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.TrailerLocationModifiedtrailerLocationModifiedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrailerLocationModifiedtrailerLocationModifiedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.TrailerLocationModifiedtrailerLocationModifiedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrailerLocationModifiedtrailerLocationModifiedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trailerLocationModified** | [**TrailerLocationModified**](TrailerLocationModified.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrailerModifiedtrailerModifiedPost

> interface{} TrailerModifiedtrailerModifiedPost(ctx).TrailerModified(trailerModified).Execute()

Trailer Modified



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
	trailerModified :=  // TrailerModified | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.TrailerModifiedtrailerModifiedPost(context.Background()).TrailerModified(trailerModified).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.TrailerModifiedtrailerModifiedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrailerModifiedtrailerModifiedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.TrailerModifiedtrailerModifiedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrailerModifiedtrailerModifiedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trailerModified** | [**TrailerModified**](TrailerModified.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrailerRemovedtrailerRemovedPost

> interface{} TrailerRemovedtrailerRemovedPost(ctx).TrailerRemoved(trailerRemoved).Execute()

Trailer Removed



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
	trailerRemoved :=  // TrailerRemoved | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.TrailerRemovedtrailerRemovedPost(context.Background()).TrailerRemoved(trailerRemoved).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.TrailerRemovedtrailerRemovedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrailerRemovedtrailerRemovedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.TrailerRemovedtrailerRemovedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrailerRemovedtrailerRemovedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trailerRemoved** | [**TrailerRemoved**](TrailerRemoved.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TspCreatedtspCreatedPost

> interface{} TspCreatedtspCreatedPost(ctx).TspCreated(tspCreated).Execute()

Tsp Created



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
	tspCreated :=  // TspCreated | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.TspCreatedtspCreatedPost(context.Background()).TspCreated(tspCreated).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.TspCreatedtspCreatedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TspCreatedtspCreatedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.TspCreatedtspCreatedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTspCreatedtspCreatedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tspCreated** | [**TspCreated**](TspCreated.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserAddeduserAddedPost

> interface{} UserAddeduserAddedPost(ctx).UserAdded(userAdded).Execute()

User Added



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
	userAdded :=  // UserAdded | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.UserAddeduserAddedPost(context.Background()).UserAdded(userAdded).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.UserAddeduserAddedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserAddeduserAddedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.UserAddeduserAddedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserAddeduserAddedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAdded** | [**UserAdded**](UserAdded.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserModifieduserModifiedPost

> interface{} UserModifieduserModifiedPost(ctx).UserModified(userModified).Execute()

User Modified



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
	userModified :=  // UserModified | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.UserModifieduserModifiedPost(context.Background()).UserModified(userModified).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.UserModifieduserModifiedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserModifieduserModifiedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.UserModifieduserModifiedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserModifieduserModifiedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userModified** | [**UserModified**](UserModified.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserRemoveduserRemovedPost

> interface{} UserRemoveduserRemovedPost(ctx).UserRemoved(userRemoved).Execute()

User Removed



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
	userRemoved :=  // UserRemoved | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.UserRemoveduserRemovedPost(context.Background()).UserRemoved(userRemoved).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.UserRemoveduserRemovedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserRemoveduserRemovedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.UserRemoveduserRemovedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserRemoveduserRemovedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userRemoved** | [**UserRemoved**](UserRemoved.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VehicleAddedvehicleAddedPost

> interface{} VehicleAddedvehicleAddedPost(ctx).VehicleAdded(vehicleAdded).Execute()

Vehicle Added



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
	vehicleAdded :=  // VehicleAdded | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.VehicleAddedvehicleAddedPost(context.Background()).VehicleAdded(vehicleAdded).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.VehicleAddedvehicleAddedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VehicleAddedvehicleAddedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.VehicleAddedvehicleAddedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVehicleAddedvehicleAddedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vehicleAdded** | [**VehicleAdded**](VehicleAdded.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VehicleLocationAddedvehicleLocationAddedPost

> interface{} VehicleLocationAddedvehicleLocationAddedPost(ctx).VehicleLocationAdded(vehicleLocationAdded).Execute()

Vehicle Location Added



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
	vehicleLocationAdded :=  // VehicleLocationAdded | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.VehicleLocationAddedvehicleLocationAddedPost(context.Background()).VehicleLocationAdded(vehicleLocationAdded).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.VehicleLocationAddedvehicleLocationAddedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VehicleLocationAddedvehicleLocationAddedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.VehicleLocationAddedvehicleLocationAddedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVehicleLocationAddedvehicleLocationAddedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vehicleLocationAdded** | [**VehicleLocationAdded**](VehicleLocationAdded.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VehicleModifiedvehicleModifiedPost

> interface{} VehicleModifiedvehicleModifiedPost(ctx).VehicleModified(vehicleModified).Execute()

Vehicle Modified



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
	vehicleModified :=  // VehicleModified | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.VehicleModifiedvehicleModifiedPost(context.Background()).VehicleModified(vehicleModified).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.VehicleModifiedvehicleModifiedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VehicleModifiedvehicleModifiedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.VehicleModifiedvehicleModifiedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVehicleModifiedvehicleModifiedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vehicleModified** | [**VehicleModified**](VehicleModified.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VehicleRemovedvehicleRemovedPost

> interface{} VehicleRemovedvehicleRemovedPost(ctx).VehicleRemoved(vehicleRemoved).Execute()

Vehicle Removed



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
	vehicleRemoved :=  // VehicleRemoved | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.VehicleRemovedvehicleRemovedPost(context.Background()).VehicleRemoved(vehicleRemoved).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.VehicleRemovedvehicleRemovedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VehicleRemovedvehicleRemovedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.VehicleRemovedvehicleRemovedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVehicleRemovedvehicleRemovedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vehicleRemoved** | [**VehicleRemoved**](VehicleRemoved.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookCreatedwebhookCreatedPost

> interface{} WebhookCreatedwebhookCreatedPost(ctx).WebhookCreated(webhookCreated).Execute()

Webhook Created



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
	webhookCreated :=  // WebhookCreated | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.WebhookCreatedwebhookCreatedPost(context.Background()).WebhookCreated(webhookCreated).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.WebhookCreatedwebhookCreatedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookCreatedwebhookCreatedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.WebhookCreatedwebhookCreatedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookCreatedwebhookCreatedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookCreated** | [**WebhookCreated**](WebhookCreated.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDeletedwebhookDeletedPost

> interface{} WebhookDeletedwebhookDeletedPost(ctx).WebhookDeleted(webhookDeleted).Execute()

Webhook Deleted



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
	webhookDeleted :=  // WebhookDeleted | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.WebhookDeletedwebhookDeletedPost(context.Background()).WebhookDeleted(webhookDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.WebhookDeletedwebhookDeletedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookDeletedwebhookDeletedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.WebhookDeletedwebhookDeletedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDeletedwebhookDeletedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDeleted** | [**WebhookDeleted**](WebhookDeleted.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookStaledwebhookStaledPost

> interface{} WebhookStaledwebhookStaledPost(ctx).WebhookStaled(webhookStaled).Execute()

Webhook Staled



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
	webhookStaled :=  // WebhookStaled | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.WebhookStaledwebhookStaledPost(context.Background()).WebhookStaled(webhookStaled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.WebhookStaledwebhookStaledPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookStaledwebhookStaledPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.WebhookStaledwebhookStaledPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookStaledwebhookStaledPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookStaled** | [**WebhookStaled**](WebhookStaled.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookUpdatedwebhookUpdatedPost

> interface{} WebhookUpdatedwebhookUpdatedPost(ctx).WebhookUpdated(webhookUpdated).Execute()

Webhook Updated



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
	webhookUpdated :=  // WebhookUpdated | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEventsAPI.WebhookUpdatedwebhookUpdatedPost(context.Background()).WebhookUpdated(webhookUpdated).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.WebhookUpdatedwebhookUpdatedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookUpdatedwebhookUpdatedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhookEventsAPI.WebhookUpdatedwebhookUpdatedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookUpdatedwebhookUpdatedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookUpdated** | [**WebhookUpdated**](WebhookUpdated.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

