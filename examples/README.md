# Catena Go SDK Examples

This directory contains runnable examples demonstrating how to use the Catena Go SDK for common tasks.

## Examples

### Onboarding & Management
-   [**Invitations**](invitations/): Create and list fleet invitations.
-   [**Schedules & Executions**](schedules_and_executions/): Navigate the hierarchy of Connections, Schedules, and Executions.

### Data Ingestion
-   [**Ingest Single Vehicle**](ingest_single_vehicle/): Fetch details for a specific vehicle.
-   [**Ingest Vehicle Locations (Paginated)**](ingest_vehicle_locations_paginated/): Efficiently ingest large volumes of vehicle location data using pagination helpers.

### Webhooks & Notifications
-   [**Create Webhook Subscription**](create_webhook_subscription/): Subscribe to real-time events.
-   [**Webhook Metrics**](webhook_metrics/): Monitor the delivery performance of your webhooks.
-   [**Webhook Receiver**](webhook_receiver/): A simple HTTP server that receives and processes webhook events using the SDK's router.

## Running the Examples

Most examples require authentication. You should set the following environment variables before running them:

```bash
export CATENA_CLIENT_ID="your_client_id"
export CATENA_CLIENT_SECRET="your_client_secret"
```

Navigate to an example directory and run it:

```bash
cd invitations
go run main.go
```
