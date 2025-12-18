package webhooks

import "errors"

var (
	ErrInvalidJSON      = errors.New("invalid json")
	ErrInvalidTimestamp = errors.New("invalid timestamp")
	ErrUnknownEvent     = errors.New("unknown event")
)
