package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/catenaclearing/catena-sdk-go"
	integrationsapi "github.com/catenaclearing/catena-sdk-go/gen/integrations"
	"github.com/catenaclearing/catena-sdk-go/pagination"
)

var ErrStop = errors.New("stop iteration")

func main() {
	// 1. Initialize the client
	clientID := os.Getenv("CATENA_CLIENT_ID")
	clientSecret := os.Getenv("CATENA_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		log.Fatal("Please set CATENA_CLIENT_ID and CATENA_CLIENT_SECRET environment variables")
	}

	c := catena.NewClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := c.Authenticate(ctx, clientID, clientSecret); err != nil {
		log.Printf("Failed to authenticate: %v", err)
		return
	}

	fmt.Println("Successfully authenticated!")

	// 2. List Connections to find a valid Connection ID
	// We need a connection ID to list schedules.
	fmt.Println("\n--- Finding a Connection ---")
	var connectionID string

	// We'll just take the first connection we find.
	err := pagination.ListConnectionsEach(
		c,
		context.Background(),
		pagination.ListConnectionsPaginationOptions{},
		func(conn integrationsapi.ConnectionRead) error {
			connectionID = conn.GetId()
			fmt.Printf("Found Connection: %s (TSP: %s)\n", connectionID, conn.GetTspId())
			return ErrStop // Stop after finding the first one
		},
	)

	if err != nil && err != ErrStop {
		log.Printf("Error listing connections: %v", err)
		return
	}

	if connectionID == "" {
		log.Println("No connections found. Please create a connection first.")
		return
	}

	// 3. List Schedules for the Connection
	fmt.Printf("\n--- Listing Schedules for Connection %s ---\n", connectionID)
	var scheduleID string

	err = pagination.ListSchedulesEach(
		c,
		context.Background(),
		pagination.ListSchedulesPaginationOptions{
			ConnectionId: connectionID,
		},
		func(schedule integrationsapi.ScheduleRead) error {
			scheduleID = schedule.GetId()
			fmt.Printf("Found Schedule: %s (Type: %s)\n", scheduleID, schedule.GetResource())
			return ErrStop // Stop after finding the first one
		},
	)

	if err != nil && err != ErrStop {
		log.Printf("Error listing schedules: %v", err)
		return
	}

	if scheduleID == "" {
		log.Println("No schedules found for this connection.")
		return
	}

	// 4. List Executions for the Schedule
	fmt.Printf("\n--- Listing Executions for Schedule %s ---\n", scheduleID)

	count := 0
	err = pagination.ListExecutionsEach(
		c,
		context.Background(),
		pagination.ListExecutionsPaginationOptions{
			ConnectionId: connectionID,
			ScheduleId:   scheduleID,
		},
		func(execution integrationsapi.ExecutionRead) error {
			count++
			fmt.Printf("%d. Execution ID: %s, Status: %s, Created: %v\n",
				count,
				execution.GetId(),
				execution.GetStatus(),
				execution.GetCreatedAt(),
			)
			return nil // Continue listing all executions
		},
	)

	if err != nil {
		log.Printf("Error listing executions: %v", err)
		return
	}

	fmt.Printf("\nTotal executions listed: %d\n", count)
}
