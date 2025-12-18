package catena

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"
)

// TokenClaims represents the claims extracted from the JWT.
//
// Note: These claims are decoded from the token without verification.
// Do not trust these claims for authorization decisions.
type TokenClaims struct {
	Exp int64               `json:"exp"`
	Sub string              `json:"sub"`
	Org map[string]OrgClaim `json:"org"`
}

// OrgClaim represents the organization details in the JWT.
type OrgClaim struct {
	ID      string   `json:"id"`
	OrgType []string `json:"org_type"`
}

// Authenticate performs the initial authentication using client credentials.
func (c *Client) Authenticate(ctx context.Context, clientID, clientSecret string) error {
	c.mu.Lock()
	c.clientID = clientID
	c.clientSecret = clientSecret
	c.mu.Unlock()

	return c.refreshToken(ctx)
}

// refreshToken refreshes the access token using singleflight to prevent thundering herd.
func (c *Client) refreshToken(ctx context.Context) error {
	// Fast path: check if token is valid
	c.mu.RLock()
	if time.Now().Before(c.expiresAt) {
		c.mu.RUnlock()
		return nil
	}
	c.mu.RUnlock()

	_, err, _ := c.sf.Do("refresh", func() (interface{}, error) {
		// Double check in case another goroutine refreshed it
		c.mu.RLock()
		if time.Now().Before(c.expiresAt) {
			c.mu.RUnlock()
			return nil, nil
		}
		clientID := c.clientID
		clientSecret := c.clientSecret
		c.mu.RUnlock()

		if clientID == "" || clientSecret == "" {
			return nil, &AuthError{Op: "refresh_token", Err: errors.New("client credentials not set")}
		}

		// Perform network request without holding the lock
		req := c.auth.OAuth20API.Token(ctx).
			GrantType("client_credentials").
			ClientId(clientID).
			ClientSecret(clientSecret).
			Scope("organization")

		resp, _, err := req.Execute()
		if err != nil {
			return nil, &AuthError{Op: "token_exchange", Err: err}
		}

		if resp == nil || resp.AccessToken == "" {
			return nil, &AuthError{Op: "token_exchange", Err: errors.New("empty access token received")}
		}

		claims, err := decodeJWT(resp.AccessToken)
		if err != nil {
			return nil, &AuthError{Op: "decode_jwt", Err: err}
		}

		c.mu.Lock()
		defer c.mu.Unlock()
		c.accessToken = resp.AccessToken
		// Set expiration with a safety buffer of 30 seconds
		c.expiresAt = time.Unix(int64(claims.Exp), 0).Add(-30 * time.Second)
		c.tokenClaims = claims

		return nil, nil
	})

	return err
}

// invalidateTokenIfMatching marks the current token as expired only if it matches the provided token.
func (c *Client) invalidateTokenIfMatching(token string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.accessToken == token {
		c.expiresAt = time.Time{}
	}
}

// decodeJWT decodes the JWT payload without verification.
// WARNING: These claims are NOT verified. Do not use them for security decisions.
func decodeJWT(token string) (*TokenClaims, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, &TokenError{Reason: "invalid token format", Err: errors.New("expected 3 parts")}
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, &TokenError{Reason: "base64 decode failed", Err: err}
	}

	var claims TokenClaims
	if err := json.Unmarshal(payload, &claims); err != nil {
		return nil, &TokenError{Reason: "json unmarshal failed", Err: err}
	}

	return &claims, nil
}

// AccessToken returns the current access token.
func (c *Client) AccessToken() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.accessToken
}

// TokenExpiresAt returns the expiration time of the current token.
func (c *Client) TokenExpiresAt() time.Time {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.expiresAt
}

// UserID returns the user ID (sub) from the token.
func (c *Client) UserID() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.tokenClaims == nil {
		return ""
	}
	return c.tokenClaims.Sub
}

// OrgID returns the organization ID from the token.
func (c *Client) OrgID() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.tokenClaims == nil || len(c.tokenClaims.Org) == 0 {
		return ""
	}
	for _, org := range c.tokenClaims.Org {
		return org.ID
	}
	return ""
}

// OrgName returns the organization name from the token.
func (c *Client) OrgName() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.tokenClaims == nil || len(c.tokenClaims.Org) == 0 {
		return ""
	}
	for name := range c.tokenClaims.Org {
		return name
	}
	return ""
}

// OrgType returns the organization type from the token.
func (c *Client) OrgType() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.tokenClaims == nil || len(c.tokenClaims.Org) == 0 {
		return ""
	}
	for _, org := range c.tokenClaims.Org {
		if len(org.OrgType) > 0 {
			return org.OrgType[0]
		}
	}
	return ""
}

// authTransport is an http.RoundTripper that handles authentication.
type authTransport struct {
	client    *Client
	transport http.RoundTripper
}

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Get the token, refreshing if necessary
	token, err := t.getToken(req.Context())
	if err != nil {
		return nil, err
	}

	// Clone the request to avoid modifying the original
	clonedReq := req.Clone(req.Context())
	clonedReq.Header.Set("Authorization", "Bearer "+token)

	// Execute the request
	resp, err := t.transport.RoundTrip(clonedReq)
	if err != nil {
		return nil, err
	}

	// Handle 401/403
	if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
		// Close the previous response body
		resp.Body.Close()

		// Invalidate token to force refresh
		t.client.invalidateTokenIfMatching(token)

		// Force refresh
		token, err = t.refreshToken(req.Context())
		if err != nil {
			return nil, err
		}

		// Retry with new token
		clonedReq.Header.Set("Authorization", "Bearer "+token)
		return t.transport.RoundTrip(clonedReq)
	}

	return resp, nil
}

func (t *authTransport) getToken(ctx context.Context) (string, error) {
	t.client.mu.RLock()
	token := t.client.accessToken
	expiresAt := t.client.expiresAt
	t.client.mu.RUnlock()

	if token == "" {
		return "", &AuthError{Op: "get_token", Err: errors.New("client not authenticated")}
	}

	if time.Now().After(expiresAt) {
		return t.refreshToken(ctx)
	}

	return token, nil
}

func (t *authTransport) refreshToken(ctx context.Context) (string, error) {
	if err := t.client.refreshToken(ctx); err != nil {
		return "", err
	}
	return t.client.AccessToken(), nil
}
