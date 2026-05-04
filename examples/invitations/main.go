package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/catenaclearing/catena-sdk-go"
	orgsapi "github.com/catenaclearing/catena-sdk-go/gen/orgs"
	"github.com/catenaclearing/catena-sdk-go/pagination"
)

func stringPtr(s string) *string {
	return &s
}

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

	// 2. Create an Invitation
	fmt.Println("\n--- Creating Invitation ---")
	newInvitation := orgsapi.InvitationCreate{
		PartnerProvidedFleetName:  *orgsapi.NewNullableString(stringPtr("Acme Fleet")),
		PartnerProvidedFleetEmail: *orgsapi.NewNullableString(stringPtr("fleet@acme.com")),
	}

	invitation, resp, err := c.Orgs().InvitationsAPI.
		CreateInvitation(context.Background()).
		InvitationCreate(newInvitation).
		Execute()

	if err != nil {
		log.Printf("Error creating invitation: %v\nResponse: %v", err, resp)
		return
	}

	if invitation != nil {
		fmt.Printf("Created invitation %s; status: %s\n", invitation.GetId(), invitation.GetStatus())
	} else if resp != nil {
		fmt.Printf("Created invitation; HTTP status: %s\n", resp.Status)
	} else {
		fmt.Println("Created invitation")
	}

	// 3. List Invitations
	fmt.Println("\n--- Listing Invitations ---")

	count := 0
	err = pagination.ListInvitationsEach(
		c,
		context.Background(),
		pagination.ListInvitationsPaginationOptions{},
		func(inv orgsapi.InvitationRead) error {
			count++
			fmt.Printf("%d. Invitation ID: %s, Status: %s\n", count, inv.GetId(), inv.GetStatus())
			return nil
		},
	)

	if err != nil {
		log.Printf("Error listing invitations: %v", err)
		return
	}
}
