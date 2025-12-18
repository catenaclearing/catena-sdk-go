package webhooks

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/catenaclearing/catena-sdk-go/webhooks_gen"
)

type UnknownEvent struct {
	EventName string
	RawData   json.RawMessage
}

func ParseBytes(b []byte) (*Envelope, any, error) {
	var env Envelope
	if err := json.Unmarshal(b, &env); err != nil {
		return nil, nil, fmt.Errorf("%w: %v", ErrInvalidJSON, err)
	}

	decoder, ok := webhooks_gen.NotificationsDecoders[env.EventName]
	if !ok {
		return &env, UnknownEvent{EventName: env.EventName, RawData: env.Data}, nil
	}

	data, err := decoder(env.Data)
	if err != nil {
		return &env, nil, fmt.Errorf("%w: %v", ErrInvalidJSON, err)
	}

	return &env, data, nil
}

func ParseRequest(r *http.Request) (*Envelope, any, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read body: %w", err)
	}
	defer r.Body.Close()

	return ParseBytes(body)
}
