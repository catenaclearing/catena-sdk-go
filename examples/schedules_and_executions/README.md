# Schedules and Executions Example

This example demonstrates how to traverse the hierarchy of Integrations resources:
1.  **Connections**: Represents a link to a Telematics Service Provider (TSP).
2.  **Schedules**: Defines how often data is synced for a connection.
3.  **Executions**: Individual runs of a schedule.

## Prerequisites

- Go 1.18 or later
- A Catena Client ID and Client Secret
- At least one existing Connection in your Catena account.

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

-   **Hierarchical Data Access**: Shows how to use IDs from parent resources (Connection ID -> Schedule ID) to access child resources (Executions).
-   **Pagination Helpers**: Uses `pagination.ListConnectionsEach`, `pagination.ListSchedulesEach`, and `pagination.ListExecutionsEach` to simplify data retrieval.
-   **Context Management**: Uses `context.WithTimeout` for the initial authentication to ensure the application doesn't hang indefinitely if the network is down.
