package webhooks

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	notificationsapi "github.com/catenaclearing/catena-sdk-go/gen/notifications"
)

func TestRouter(t *testing.T) {
	router := NewRouter()

	called := false
	router.Handle("connection.created", func(ctx context.Context, env *Envelope, data []notificationsapi.BaseConnectionEvent) error {
		called = true
		if len(data) != 1 {
			return errors.New("expected 1 item")
		}
		return nil
	})

	// Test valid request
	body := `{
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

	req := httptest.NewRequest("POST", "/webhooks", bytes.NewBufferString(body))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("expected status 204, got %d", w.Code)
	}
	if !called {
		t.Error("handler not called")
	}

	// Test unregistered event
	bodyUnregistered := `{
		"version": "1.0",
		"event_name": "other.event",
		"webhook_id": "wh_123",
		"timestamp": "2023-01-01T00:00:00Z",
		"data": []
	}`
	req = httptest.NewRequest("POST", "/webhooks", bytes.NewBufferString(bodyUnregistered))
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusNotFound {
		t.Errorf("expected status 404 for unregistered event, got %d", w.Code)
	}

	// Test handler error
	router.Handle("connection.created", func(ctx context.Context, env *Envelope, data []notificationsapi.BaseConnectionEvent) error {
		return errors.New("handler error")
	})

	req = httptest.NewRequest("POST", "/webhooks", bytes.NewBufferString(body))
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected status 500 for handler error, got %d", w.Code)
	}
}
