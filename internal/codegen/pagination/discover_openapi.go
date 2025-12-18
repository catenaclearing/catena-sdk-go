package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type OpenAPISpec struct {
	Paths      map[string]map[string]Operation `json:"paths"`
	Components Components                      `json:"components"`
}

type Operation struct {
	Tags        []string            `json:"tags"`
	OperationID string              `json:"operationId"`
	Parameters  []Parameter         `json:"parameters"`
	Responses   map[string]Response `json:"responses"`
}

type Parameter struct {
	Name     string `json:"name"`
	In       string `json:"in"`
	Required bool   `json:"required"`
	Schema   Schema `json:"schema"`
}

type Response struct {
	Content map[string]MediaType `json:"content"`
	Ref     string               `json:"$ref"`
}

type MediaType struct {
	Schema Schema `json:"schema"`
}

type Schema struct {
	Ref        string            `json:"$ref"`
	Type       string            `json:"type"`
	Format     string            `json:"format"`
	Properties map[string]Schema `json:"properties"`
	Items      *Schema           `json:"items"`
	AnyOf      []Schema          `json:"anyOf"`
}

type Components struct {
	Schemas map[string]Schema `json:"schemas"`
}

func discoverPaginatedOperations(specsDir string) ([]PaginatedOperation, error) {
	var operations []PaginatedOperation

	services := []string{"integrations", "orgs", "telematics", "notifications"}

	for _, service := range services {
		specPath := filepath.Join(specsDir, service, "openapi.json")
		if _, err := os.Stat(specPath); os.IsNotExist(err) {
			fmt.Printf("Skipping %s: %v\n", service, err)
			continue
		}

		specBytes, err := os.ReadFile(specPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read spec %s: %w", specPath, err)
		}

		var spec OpenAPISpec
		if err := json.Unmarshal(specBytes, &spec); err != nil {
			return nil, fmt.Errorf("failed to parse spec %s: %w", specPath, err)
		}

		for _, methods := range spec.Paths {
			for method, op := range methods {
				if method != "get" {
					continue
				}

				// Check for pagination params
				var cursorParam, limitParam string
				var hasCursor, hasLimit bool
				var requiredParams []Param
				var optionalParams []Param

				for _, param := range op.Parameters {
					p := Param{
						Name: param.Name,
						In:   param.In,
					}

					schema := param.Schema
					// Handle AnyOf (usually nullable or union types)
					if len(schema.AnyOf) > 0 {
						for _, s := range schema.AnyOf {
							if s.Type != "null" && s.Type != "" {
								schema = s
								break
							}
							// Also check if it has items (array)
							if s.Type == "array" {
								schema = s
								break
							}
						}
					}

					// Simple type mapping
					if schema.Type == "integer" {
						p.Type = "int"
						p.IsPointer = !param.Required
					} else if schema.Type == "boolean" {
						p.Type = "bool"
						p.IsPointer = !param.Required
					} else if schema.Type == "array" {
						p.Type = "[]string"
						p.IsPointer = !param.Required
					} else if schema.Type == "string" && schema.Format == "date-time" {
						p.Type = "time.Time"
						p.IsPointer = !param.Required
					} else {
						p.Type = "string"
						p.IsPointer = !param.Required
					}

					if param.Name == "cursor" || param.Name == "page" || param.Name == "next_page" || param.Name == "page_token" || param.Name == "after" {
						cursorParam = param.Name
						hasCursor = true
					} else if param.Name == "size" || param.Name == "limit" || param.Name == "page_size" || param.Name == "pageSize" || param.Name == "per_page" || param.Name == "perPage" {
						limitParam = param.Name
						hasLimit = true
					} else {
						if param.Required {
							requiredParams = append(requiredParams, p)
						} else {
							optionalParams = append(optionalParams, p)
						}
					}
				}

				if !hasCursor {
					continue
				}

				// Check response schema
				resp200, ok := op.Responses["200"]
				if !ok {
					continue
				}

				schema := resp200.Content["application/json"].Schema
				schemaName := ""
				if schema.Ref != "" {
					parts := strings.Split(schema.Ref, "/")
					schemaName = parts[len(parts)-1]
				}

				// Resolve schema to check properties
				if schemaName != "" {
					compSchema, ok := spec.Components.Schemas[schemaName]
					if ok {
						_, hasItems := compSchema.Properties["items"]
						_, hasNextPage := compSchema.Properties["next_page"]

						if hasItems && hasNextPage {
							// Found paginated operation
							tag := "Default"
							if len(op.Tags) > 0 {
								tag = op.Tags[0]
							}

							operations = append(operations, PaginatedOperation{
								Service:         service,
								ServiceAccessor: strings.ToUpper(service[:1]) + service[1:],
								Tag:             tag,
								OperationID:     op.OperationID,
								CursorParam:     cursorParam,
								LimitParam:      limitParam,
								ResponseSchema:  schemaName,
								RequiredParams:  requiredParams,
								OptionalParams:  optionalParams,
								HasCursor:       hasCursor,
								HasLimit:        hasLimit,
							})
						}
					}
				}
			}
		}
	}

	return operations, nil
}
