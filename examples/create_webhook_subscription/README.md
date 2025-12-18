# Create Webhook Subscription Example

This example demonstrates how to subscribe to real-time events using webhooks.

## Prerequisites

- Go 1.18 or later
- A Catena Client ID and Client Secret
- A publicly accessible URL to receive webhook events (e.g., using [ngrok](https://ngrok.com/) or a deployed server).

## Running the Example

1.  Set your credentials and webhook URL as environment variables:

    ```bash
    export CATENA_CLIENT_ID="your_client_id"
    export CATENA_CLIENT_SECRET="your_client_secret"
    export WEBHOOK_URL="https://your-server.com/webhook"
    ```

2.  Run the example:

    ```bash
    go run main.go
    ```

## Code Highlights

-   **Event Selection**: Shows how to use the `WebhookEventName` and `WebhookCatchAllEvents` enums to specify which events to subscribe to (e.g., `vehicle_location.*`).
-   **Security**: Demonstrates setting a `Secret` which Catena will use to sign the webhook payloads, allowing you to verify their authenticity.
