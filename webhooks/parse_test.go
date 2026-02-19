package webhooks

import (
	"net/http/httptest"
	"strings"
	"testing"

	notificationsapi "github.com/catenaclearing/catena-sdk-go/gen/notifications"
	"github.com/catenaclearing/catena-sdk-go/webhooks_gen"
)

func TestParseBytes(t *testing.T) {
	// Valid event
	validJSON := `{
		"version": "1.0",
		"event_name": "connection.created",
		"webhook_id": "wh_123",
		"timestamp": "2023-01-01T00:00:00Z",
		"data": [{
			"id": "conn_123",
			"source_name": "azuga",
			"fleet_id": "fleet_123",
			"tsp_id": "tsp_123",
			"status": "active"
		}]
	}`

	env, payload, err := ParseBytes([]byte(validJSON))
	if err != nil {
		t.Fatalf("ParseBytes failed: %v", err)
	}

	if env.EventName != "connection.created" {
		t.Errorf("expected event name 'connection.created', got '%s'", env.EventName)
	}

	if _, ok := payload.([]notificationsapi.BaseConnectionEvent); !ok {
		t.Errorf("expected payload type []notificationsapi.BaseConnectionEvent, got %T", payload)
	}

	// Unknown event
	unknownJSON := `{
		"version": "1.0",
		"event_name": "unknown.event",
		"webhook_id": "wh_123",
		"timestamp": "2023-01-01T00:00:00Z",
		"data": [{"foo": "bar"}]
	}`

	_, payload, err = ParseBytes([]byte(unknownJSON))
	if err != nil {
		t.Fatalf("ParseBytes failed for unknown event: %v", err)
	}

	if _, ok := payload.(UnknownEvent); !ok {
		t.Errorf("expected payload type UnknownEvent, got %T", payload)
	}

	// Invalid JSON
	invalidJSON := `{ "invalid": `
	_, _, err = ParseBytes([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error for invalid JSON, got nil")
	}
}

func TestParseBytesUserAddedWithFleetRef(t *testing.T) {
	userAdded := `{
		"version": "1.0",
		"event_name": "user.added",
		"webhook_id": "wh_123",
		"data": [{
			"id": "user_123",
			"fleet_id": "fleet_123",
			"fleet_ref": "partner-fleet-abc",
			"source_name": "samsara",
			"connection_id": "conn_123",
			"source_id": "source_123",
			"created_at": "2023-01-01T00:00:00Z",
			"updated_at": "2023-01-01T00:00:00Z",
			"occurred_at": "2023-01-01T00:00:00Z"
		}]
	}`

	_, payload, err := ParseBytes([]byte(userAdded))
	if err != nil {
		t.Fatalf("ParseBytes failed: %v", err)
	}

	users, ok := payload.([]notificationsapi.BaseUser)
	if !ok {
		t.Fatalf("expected payload type []notificationsapi.BaseUser, got %T", payload)
	}
	if len(users) != 1 {
		t.Fatalf("expected 1 user, got %d", len(users))
	}
	if got := users[0].GetFleetRef(); got != "partner-fleet-abc" {
		t.Fatalf("expected fleet_ref 'partner-fleet-abc', got %q", got)
	}
}

func TestParseBytesIgnoresUnknownFields(t *testing.T) {
	withUnknownFields := `{
		"version": "1.0",
		"event_name": "connection.created",
		"webhook_id": "wh_123",
		"data": [{
			"id": "conn_123",
			"source_name": "azuga",
			"fleet_id": "fleet_123",
			"tsp_id": "tsp_123",
			"status": "active",
			"future_field": "new-value"
		}]
	}`

	_, payload, err := ParseBytes([]byte(withUnknownFields))
	if err != nil {
		t.Fatalf("ParseBytes failed with unknown fields: %v", err)
	}

	if _, ok := payload.([]notificationsapi.BaseConnectionEvent); !ok {
		t.Fatalf("expected payload type []notificationsapi.BaseConnectionEvent, got %T", payload)
	}
}

func TestParseRequest(t *testing.T) {
	body := `{
		"version": "1.0",
		"event_name": "connection.created",
		"webhook_id": "wh_123",
		"data": [{
			"id": "conn_123",
			"source_name": "azuga",
			"fleet_id": "fleet_123",
			"tsp_id": "tsp_123",
			"status": "active"
		}]
	}`

	req := httptest.NewRequest("POST", "/webhooks", strings.NewReader(body))
	env, payload, err := ParseRequest(req)
	if err != nil {
		t.Fatalf("ParseRequest failed: %v", err)
	}
	if env.EventName != "connection.created" {
		t.Fatalf("expected event_name connection.created, got %s", env.EventName)
	}
	if _, ok := payload.([]notificationsapi.BaseConnectionEvent); !ok {
		t.Fatalf("expected []notificationsapi.BaseConnectionEvent payload, got %T", payload)
	}
}

func TestSupportedEvents(t *testing.T) {
	expectedEvents := []string{
		"connection.created",
		"connection.staled",
		"webhook.deleted",
	}

	for _, event := range expectedEvents {
		if _, ok := webhooks_gen.NotificationsDecoders[event]; !ok {
			t.Errorf("expected event '%s' to be supported", event)
		}
	}
}
