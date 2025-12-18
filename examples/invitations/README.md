# Invitations Example

This example demonstrates how to use the Catena Go SDK to:
1.  Authenticate with the API.
2.  Create a new fleet invitation.
3.  List all existing invitations using the pagination helper.

## Prerequisites

- Go 1.18 or later
- A Catena Client ID and Client Secret

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

-   **Authentication**: Uses `c.Authenticate(ctx, clientID, clientSecret)` to obtain a token.
-   **Creating Resources**: Uses the generated `orgsapi` models (e.g., `InvitationCreate`) and the fluent builder pattern (`CreateInvitation(...).InvitationCreate(...).Execute()`).
-   **Pagination**: Uses `pagination.ListInvitationsEach` to automatically handle cursor-based pagination, simplifying the retrieval of large lists.
