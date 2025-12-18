# Webhook Metrics Example

This example demonstrates how to monitor the health and performance of your webhook subscriptions.

## Prerequisites

- Go 1.18 or later
- A Catena Client ID and Client Secret
- At least one active webhook subscription.

## Running the Example

1.  Set your credentials as environment variables:

    ```bash
    export CATENA_CLIENT_ID="your_client_id"
    export CATENA_CLIENT_SECRET="your_client_secret"
    ```

2.  Run the example:

    ```bash
    go run main.go
    ```

## Code Highlights

-   **Monitoring**: Shows how to retrieve metrics like `TotalDelivered`, `TotalFailed`, and `SuccessRate` to ensure your webhook integration is functioning correctly.
-   **Resource Identification**: Uses `ListWebhookSubscriptionsEach` to find a webhook ID to query.
