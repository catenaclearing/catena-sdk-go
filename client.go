package catena

import (
	"net/http"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"

	authenticationapi "github.com/catenaclearing/catena-sdk-go/gen/authentication"
	integrationsapi "github.com/catenaclearing/catena-sdk-go/gen/integrations"
	notificationsapi "github.com/catenaclearing/catena-sdk-go/gen/notifications"
	orgsapi "github.com/catenaclearing/catena-sdk-go/gen/orgs"
	telematicsapi "github.com/catenaclearing/catena-sdk-go/gen/telematics"
)

// Client is the main entry point for the Catena SDK.
type Client struct {
	baseURL    string
	httpClient *http.Client

	// Generated clients
	integrations  *integrationsapi.APIClient
	orgs          *orgsapi.APIClient
	telematics    *telematicsapi.APIClient
	notifications *notificationsapi.APIClient
	auth          *authenticationapi.APIClient

	// Authentication state
	mu           sync.RWMutex
	clientID     string
	clientSecret string
	accessToken  string
	expiresAt    time.Time
	tokenClaims  *TokenClaims
	sf           singleflight.Group
}

// NewClient creates a new Client with the given options.
func NewClient(opts ...Option) *Client {
	c := &Client{
		baseURL:    "https://api.catenatelematics.com",
		httpClient: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(c)
	}

	// Initialize the authentication client immediately as it doesn't need the auth transport
	authConfig := authenticationapi.NewConfiguration()
	authConfig.Servers = authenticationapi.ServerConfigurations{
		{
			URL: "https://auth.catenatelematics.com/realms/catena",
		},
	}
	authConfig.HTTPClient = c.httpClient
	c.auth = authenticationapi.NewAPIClient(authConfig)

	// Initialize other clients with a placeholder configuration
	// They will be fully configured in Authenticate or lazily
	// But we need the transport to be set up.

	// We'll use a custom transport that references the client to get the token
	transport := &authTransport{
		client:    c,
		transport: c.httpClient.Transport,
	}
	if transport.transport == nil {
		transport.transport = http.DefaultTransport
	}

	secureClient := &http.Client{
		Transport: transport,
		Timeout:   c.httpClient.Timeout,
	}

	// Configure generated clients
	integrationsConfig := integrationsapi.NewConfiguration()
	integrationsConfig.Servers = integrationsapi.ServerConfigurations{{URL: c.baseURL}}
	integrationsConfig.HTTPClient = secureClient
	c.integrations = integrationsapi.NewAPIClient(integrationsConfig)

	orgsConfig := orgsapi.NewConfiguration()
	orgsConfig.Servers = orgsapi.ServerConfigurations{{URL: c.baseURL}}
	orgsConfig.HTTPClient = secureClient
	c.orgs = orgsapi.NewAPIClient(orgsConfig)

	telematicsConfig := telematicsapi.NewConfiguration()
	telematicsConfig.Servers = telematicsapi.ServerConfigurations{{URL: c.baseURL}}
	telematicsConfig.HTTPClient = secureClient
	c.telematics = telematicsapi.NewAPIClient(telematicsConfig)

	notificationsConfig := notificationsapi.NewConfiguration()
	notificationsConfig.Servers = notificationsapi.ServerConfigurations{{URL: c.baseURL}}
	notificationsConfig.HTTPClient = secureClient
	c.notifications = notificationsapi.NewAPIClient(notificationsConfig)

	return c
}

// Integrations returns the Integrations API client.
func (c *Client) Integrations() *integrationsapi.APIClient {
	return c.integrations
}

// Orgs returns the Orgs API client.
func (c *Client) Orgs() *orgsapi.APIClient {
	return c.orgs
}

// Telematics returns the Telematics API client.
func (c *Client) Telematics() *telematicsapi.APIClient {
	return c.telematics
}

// Notifications returns the Notifications API client.
func (c *Client) Notifications() *notificationsapi.APIClient {
	return c.notifications
}
