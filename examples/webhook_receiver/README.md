# Webhook Receiver Example

This example demonstrates how to receive and process webhook events from Catena using the SDK's `webhooks` package.

## Prerequisites

- Go 1.18 or later
- A way to expose your local server to the internet (e.g., [ngrok](https://ngrok.com/)) so Catena can send events to it.

## Running the Example

1.  Run the example:

    ```bash
    go run main.go
    ```

    The server will start on port 8080 by default.

2.  Expose your local server (optional, if testing with real Catena events):

    ```bash
    ngrok http 8080
    ```

3.  Configure a Webhook in Catena (using the `create_webhook_subscription` example or the dashboard) to point to your exposed URL.

## Code Highlights

-   **Type-Safe Handlers**: The `webhooks.Router` allows you to define handlers with specific data types (e.g., `[]notificationsapi.BaseVehicleLocation`). The router automatically unmarshals the JSON payload into these structs.
-   **Routing**: Automatically routes events based on the `event_name` field in the payload.
-   **Standard HTTP Handler**: The router implements `http.Handler`, making it easy to integrate with standard Go web servers or frameworks like Gin or Echo.
