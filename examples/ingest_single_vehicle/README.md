# Ingest Single Vehicle Example

This example demonstrates how to retrieve detailed information about a specific vehicle.

## Prerequisites

- Go 1.18 or later
- A Catena Client ID and Client Secret
- At least one vehicle in your connected fleets.

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

-   **Direct Resource Access**: Shows how to use `GetVehicle(ctx, vehicleID)` to fetch a single resource by its ID.
-   **Model Accessors**: Demonstrates using the generated model methods (e.g., `GetVin()`, `GetMake()`) to safely access field values, handling potential nil pointers automatically.
