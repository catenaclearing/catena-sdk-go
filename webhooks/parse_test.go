package webhooks

import (
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

	env, payload, err = ParseBytes([]byte(unknownJSON))
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
