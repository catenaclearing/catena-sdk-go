package catena_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/catenaclearing/catena-sdk-go"
	integrationsapi "github.com/catenaclearing/catena-sdk-go/gen/integrations"
	"github.com/catenaclearing/catena-sdk-go/pagination"
)

type MockTransport struct {
}

func (t *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.URL.Host == "auth.catenatelematics.com" {
		// Return dummy token
		// JWT: header.payload.signature
		// payload: {"exp": 9999999999, "sub": "test"}
		dummyJWT := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjk5OTk5OTk5OTksInN1YiI6InRlc3QifQ.signature"
		respBody := `{"access_token":"` + dummyJWT + `","expires_in":3600,"token_type":"Bearer"}`
		resp := &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(respBody)),
			Header:     make(http.Header),
		}
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	}
	return http.DefaultTransport.RoundTrip(req)
}

func TestListConnectionsPagination(t *testing.T) {
	// Mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/integrations/connections" {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}

		cursor := r.URL.Query().Get("cursor")

		var resp map[string]interface{}

		switch cursor {
		case "":
			// First page
			resp = map[string]interface{}{
				"items": []map[string]interface{}{
					{"id": "1", "source_name": "geotab", "created_at": "2023-01-01T00:00:00Z", "updated_at": "2023-01-01T00:00:00Z", "fleet_id": "f1", "tsp_id": "t1", "credentials": map[string]interface{}{"api_key": "k1"}, "status": "active", "description": "d1"},
					{"id": "2", "source_name": "samsara", "created_at": "2023-01-02T00:00:00Z", "updated_at": "2023-01-02T00:00:00Z", "fleet_id": "f2", "tsp_id": "t2", "credentials": map[string]interface{}{"api_key": "k2"}, "status": "active", "description": "d2"},
				},
				"next_page": "page2",
			}
		case "page2":
			// Second page
			resp = map[string]interface{}{
				"items": []map[string]interface{}{
					{"id": "3", "source_name": "motive", "created_at": "2023-01-03T00:00:00Z", "updated_at": "2023-01-03T00:00:00Z", "fleet_id": "f3", "tsp_id": "t3", "credentials": map[string]interface{}{"api_key": "k3"}, "status": "active", "description": "d3"},
				},
				"next_page": "", // End
			}
		default:
			http.Error(w, "Invalid cursor", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Client
	httpClient := &http.Client{Transport: &MockTransport{}}
	client := catena.NewClient(
		catena.WithBaseURL(server.URL),
		catena.WithClientID("test"),
		catena.WithClientSecret("test"),
		catena.WithHTTPClient(httpClient),
	)

	// Authenticate
	err := client.Authenticate(context.Background(), "test", "test")
	if err != nil {
		t.Fatalf("Authenticate failed: %v", err)
	}

	// Call wrapper
	var items []integrationsapi.ConnectionRead
	err = pagination.ListConnectionsEach(client, context.Background(), pagination.ListConnectionsPaginationOptions{}, func(item integrationsapi.ConnectionRead) error {
		items = append(items, item)
		return nil
	})

	if err != nil {
		t.Fatalf("ListConnectionsEach failed: %v", err)
	}

	if len(items) != 3 {
		t.Errorf("expected 3 items, got %d", len(items))
	}

	if len(items) > 0 && items[0].Id != "1" {
		t.Errorf("expected item 1 id 1, got %s", items[0].Id)
	}
	if len(items) > 2 && items[2].Id != "3" {
		t.Errorf("expected item 3 id 3, got %s", items[2].Id)
	}
}
