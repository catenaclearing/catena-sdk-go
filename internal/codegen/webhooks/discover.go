package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

type OpenAPISpec struct {
	Webhooks   map[string]map[string]Operation `json:"webhooks"`
	Components Components                      `json:"components"`
}

type Operation struct {
	RequestBody RequestBody `json:"requestBody"`
}

type RequestBody struct {
	Content map[string]MediaType `json:"content"`
}

type MediaType struct {
	Schema Schema `json:"schema"`
}

type Schema struct {
	Ref        string            `json:"$ref"`
	Type       string            `json:"type"`
	Properties map[string]Schema `json:"properties"`
	Items      *Schema           `json:"items"`
}

type Components struct {
	Schemas map[string]Schema `json:"schemas"`
}

func discoverWebhookEvents(specPath string) ([]WebhookEvent, error) {
	data, err := os.ReadFile(specPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read spec: %w", err)
	}

	var spec OpenAPISpec
	if err := json.Unmarshal(data, &spec); err != nil {
		return nil, fmt.Errorf("failed to parse spec: %w", err)
	}

	var events []WebhookEvent

	for eventName, methods := range spec.Webhooks {
		// Usually "post"
		op, ok := methods["post"]
		if !ok {
			continue
		}

		schema := op.RequestBody.Content["application/json"].Schema
		if schema.Ref == "" {
			continue
		}

		// Resolve Ref to get the actual schema
		schemaName := path.Base(schema.Ref)
		def, ok := spec.Components.Schemas[schemaName]
		if !ok {
			continue
		}

		// Look for "data" property which should be an array
		dataProp, ok := def.Properties["data"]
		if !ok || dataProp.Type != "array" || dataProp.Items == nil {
			continue
		}

		// Get the items ref
		itemsRef := dataProp.Items.Ref
		if itemsRef == "" {
			continue
		}

		goTypeName := path.Base(itemsRef)

		events = append(events, WebhookEvent{
			EventName:  eventName,
			GoTypeName: goTypeName,
		})
	}

	return events, nil
}
