package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerator(t *testing.T) {
	specContent := `{
  "openapi": "3.1.0",
  "webhooks": {
    "webhook.deleted": {
      "post": {
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/WebhookDeleted"
              }
            }
          }
        }
      }
    }
  },
  "components": {
    "schemas": {
      "WebhookDeleted": {
        "properties": {
          "data": {
            "type": "array",
            "items": {
              "$ref": "#/components/schemas/WebhookDeletedEvent"
            }
          }
        }
      }
    }
  }
}`

	tmpDir := t.TempDir()
	specPath := filepath.Join(tmpDir, "openapi.json")
	if err := os.WriteFile(specPath, []byte(specContent), 0644); err != nil {
		t.Fatal(err)
	}

	events, err := discoverWebhookEvents(specPath)
	if err != nil {
		t.Fatalf("discoverWebhookEvents failed: %v", err)
	}

	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}

	if events[0].EventName != "webhook.deleted" {
		t.Errorf("expected event name 'webhook.deleted', got '%s'", events[0].EventName)
	}
	if events[0].GoTypeName != "WebhookDeletedEvent" {
		t.Errorf("expected go type 'WebhookDeletedEvent', got '%s'", events[0].GoTypeName)
	}

	code, err := renderEvents(events)
	if err != nil {
		t.Fatalf("renderEvents failed: %v", err)
	}

	codeStr := string(code)
	if !strings.Contains(codeStr, `"webhook.deleted": decodeWebhookDeletedEvent,`) {
		t.Error("generated code missing map entry")
	}
	if !strings.Contains(codeStr, `func decodeWebhookDeletedEvent(data json.RawMessage) (any, error)`) {
		t.Error("generated code missing decoder function")
	}
	if !strings.Contains(codeStr, `var out []notificationsapi.WebhookDeletedEvent`) {
		t.Error("generated code missing typed slice")
	}
}
