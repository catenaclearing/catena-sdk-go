package main

import (
	"bytes"
	_ "embed"
	"go/format"
	"sort"
	"text/template"
)

//go:embed templates/notifications_events.go.tmpl
var eventsTemplate string

func renderEvents(events []WebhookEvent) ([]byte, error) {
	uniqueTypes := make(map[string]bool)
	var types []string
	for _, e := range events {
		if !uniqueTypes[e.GoTypeName] {
			uniqueTypes[e.GoTypeName] = true
			types = append(types, e.GoTypeName)
		}
	}
	sort.Strings(types)

	data := RenderData{
		Events:      events,
		UniqueTypes: types,
	}

	tmpl, err := template.New("events").Parse(eventsTemplate)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, err
	}

	return format.Source(buf.Bytes())
}
