package webhooks

import (
	"encoding/json"
	"time"
)

type Envelope struct {
	Version         string          `json:"version"`
	EventName       string          `json:"event_name"`
	Data            json.RawMessage `json:"data"` // IMPORTANT: array payload
	WebhookID       string          `json:"webhook_id"`
	Timestamp       time.Time       `json:"timestamp"`
	ID              string          `json:"id"`
	DeliveryAttempt int             `json:"delivery_attempt"`
}
