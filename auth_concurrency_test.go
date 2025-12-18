package catena_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/catenaclearing/catena-sdk-go"
)

type RedirectTransport struct {
	AuthURL       string
	BaseTransport http.RoundTripper
}

func (t *RedirectTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.URL.Host == "auth.catenatelematics.com" {
		// Redirect to mock auth server
		parsedAuthURL, _ := url.Parse(t.AuthURL)
		req.URL.Scheme = parsedAuthURL.Scheme
		req.URL.Host = parsedAuthURL.Host
		return t.BaseTransport.RoundTrip(req)
	}
	return t.BaseTransport.RoundTrip(req)
}

func TestAuthConcurrency(t *testing.T) {
	var authCalls int32

	// Mock Auth Server
	authServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&authCalls, 1)
		time.Sleep(100 * time.Millisecond) // Simulate slow auth
		w.Header().Set("Content-Type", "application/json")
		// Return a valid token
		w.Write([]byte(`{"access_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjk5OTk5OTk5OTksInN1YiI6InRlc3QifQ.sig","expires_in":3600,"token_type":"Bearer"}`))
	}))
	defer authServer.Close()

	// Mock API Server (to trigger 401)
	apiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Always return 401 to force refresh
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer apiServer.Close()

	transport := &RedirectTransport{
		AuthURL:       authServer.URL,
		BaseTransport: http.DefaultTransport,
	}

	client := catena.NewClient(
		catena.WithBaseURL(apiServer.URL),
		catena.WithClientID("test"),
		catena.WithClientSecret("test"),
		catena.WithHTTPClient(&http.Client{Transport: transport}),
	)

	// Initial Auth
	err := client.Authenticate(context.Background(), "test", "test")
	if err != nil {
		t.Fatalf("Authenticate failed: %v", err)
	}

	// Reset counter (Authenticate calls it once)
	atomic.StoreInt32(&authCalls, 0)

	// Spawn goroutines
	var wg sync.WaitGroup
	concurrency := 50

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Call an API that returns 401, triggering refresh
			client.Integrations().ConnectionsAPI.ListConnections(context.Background()).Execute()
		}()
	}

	wg.Wait()

	calls := atomic.LoadInt32(&authCalls)
	// We expect exactly 1 call because singleflight should coalesce them.
	if calls != 1 {
		t.Errorf("Expected 1 auth call, got %d", calls)
	}
}
