package main

type WebhookEvent struct {
	EventName  string
	GoTypeName string // e.g. WebhookCreated
}

type RenderData struct {
	Events      []WebhookEvent
	UniqueTypes []string
}
