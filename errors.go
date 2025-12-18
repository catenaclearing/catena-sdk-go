package catena

import "fmt"

// AuthError represents an error during authentication.
type AuthError struct {
	Op  string
	Err error
}

func (e *AuthError) Error() string {
	if e.Err == nil {
		return fmt.Sprintf("authentication error: %s", e.Op)
	}
	return fmt.Sprintf("authentication error during %s: %v", e.Op, e.Err)
}

func (e *AuthError) Unwrap() error {
	return e.Err
}

// TokenError represents an error with the token (invalid, expired, etc).
type TokenError struct {
	Reason string
	Err    error
}

func (e *TokenError) Error() string {
	if e.Err == nil {
		return fmt.Sprintf("token error: %s", e.Reason)
	}
	return fmt.Sprintf("token error: %s: %v", e.Reason, e.Err)
}

func (e *TokenError) Unwrap() error {
	return e.Err
}
