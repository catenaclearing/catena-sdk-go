# Ingest Vehicle Locations (Paginated) Example

This example demonstrates how to efficiently ingest large volumes of vehicle location data using the Catena SDK's pagination helpers.

## Prerequisites

- Go 1.18 or later
- A Catena Client ID and Client Secret
- Access to telematics data (fleets connected to your account).

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

-   **Efficient Pagination**: The `pagination.ListVehicleLocationsEach` function handles the complexity of cursor-based pagination, allowing you to process an indefinite number of records as a continuous stream.
-   **Filtering**: Shows how to use `ListVehicleLocationsPaginationOptions` to filter data by time range (`FromDatetime`).
-   **Processing Loop**: The callback function allows you to process each record individually (e.g., save to a database) without loading the entire dataset into memory.
