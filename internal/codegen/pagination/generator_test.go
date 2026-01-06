package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGenerator(t *testing.T) {
	// Setup temp dir
	tmpDir, err := os.MkdirTemp("", "catena-gen-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Create specs/integrations/openapi.json
	specsDir := filepath.Join(tmpDir, "specs", "integrations")
	if err := os.MkdirAll(specsDir, 0755); err != nil {
		t.Fatal(err)
	}

	specContent := `{
  "paths": {
    "/v2/integrations/connections": {
      "get": {
        "tags": ["Connections"],
        "operationId": "list_connections",
        "parameters": [
          {"name": "cursor", "in": "query", "schema": {"type": "string"}},
          {"name": "size", "in": "query", "schema": {"type": "integer"}}
        ],
        "responses": {
          "200": {
            "content": {
              "application/json": {
                "schema": {"$ref": "#/components/schemas/ConnectionList"}
              }
            }
          }
        }
      }
    }
  },
  "components": {
    "schemas": {
      "ConnectionList": {
        "properties": {
          "items": {"type": "array", "items": {"type": "string"}},
          "next_page": {"type": "string"}
        }
      }
    }
  }
}`
	if err := os.WriteFile(filepath.Join(specsDir, "openapi.json"), []byte(specContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Create gen/integrations/api_connections.go
	genDir := filepath.Join(tmpDir, "gen", "integrations")
	if err := os.MkdirAll(genDir, 0755); err != nil {
		t.Fatal(err)
	}

	goContent := `package integrationsapi

import (
"context"
"net/http"
)

type ConnectionsAPI interface {
	ListConnections(ctx context.Context) ApiListConnectionsRequest
}

type ConnectionsAPIService struct {}

func (a *ConnectionsAPIService) ListConnections(ctx context.Context) ApiListConnectionsRequest {
	return ApiListConnectionsRequest{}
}

type ApiListConnectionsRequest struct {}

func (r ApiListConnectionsRequest) Cursor(cursor string) ApiListConnectionsRequest { return r }
func (r ApiListConnectionsRequest) Size(size int32) ApiListConnectionsRequest { return r }
func (r ApiListConnectionsRequest) Execute() (*ConnectionList, *http.Response, error) { return nil, nil, nil }

type ConnectionList struct {
	Items []string
	NextPage *string
}
`
	if err := os.WriteFile(filepath.Join(genDir, "api_connections.go"), []byte(goContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Run discovery
	ops, err := discoverPaginatedOperations(filepath.Join(tmpDir, "specs"))
	if err != nil {
		t.Fatalf("discoverPaginatedOperations failed: %v", err)
	}

	if len(ops) != 1 {
		t.Fatalf("expected 1 operation, got %d", len(ops))
	}

	op := ops[0]
	if op.OperationID != "list_connections" {
		t.Errorf("expected operationId list_connections, got %s", op.OperationID)
	}

	// Run inspection
	ops, err = inspectGeneratedCode(tmpDir, ops)
	if err != nil {
		t.Fatalf("inspectGeneratedCode failed: %v", err)
	}

	op = ops[0]
	if op.MethodName != "ListConnections" {
		t.Errorf("expected MethodName ListConnections, got %s", op.MethodName)
	}
	if op.CursorGoName != "Cursor" {
		t.Errorf("expected CursorGoName Cursor, got %s", op.CursorGoName)
	}
	if op.LimitGoName != "Size" {
		t.Errorf("expected LimitGoName Size, got %s", op.LimitGoName)
	}
	if op.ItemType != "string" { // In dummy go code, items is []string
		t.Errorf("expected ItemType string, got %s", op.ItemType)
	}
}
