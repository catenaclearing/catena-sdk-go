package catena

import (
	"testing"
)

func TestClientStructure(t *testing.T) {
	client := NewClient(WithBaseURL("https://example.com"))

	if client.Integrations() == nil {
		t.Error("Integrations client is nil")
	}
	if client.Orgs() == nil {
		t.Error("Orgs client is nil")
	}
	if client.Telematics() == nil {
		t.Error("Telematics client is nil")
	}
	if client.Notifications() == nil {
		t.Error("Notifications client is nil")
	}
}

func TestAuthMethods(t *testing.T) {
	client := NewClient()

	// Just checking method signatures exist
	_ = client.AccessToken()
	_ = client.TokenExpiresAt()
	_ = client.UserID()
	_ = client.OrgID()
	_ = client.OrgName()
	_ = client.OrgType()

	// Authenticate signature check
	// err := client.Authenticate(context.Background(), "id", "secret")
}
