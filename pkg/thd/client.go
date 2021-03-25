package thd

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	IdentityClientInstanceKey = `THDIdentityClient`
)

var (
	mux         sync.Mutex
	expiration  time.Time
	cachedToken string
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . IdentityClient

// IdentityClient makes a request for a client token.
//
// See https://om-curriculum.apps-np.homedepot.com/application-security/api-security/01_service_to_service_oauth2/50_client-cred-flow/
type IdentityClient interface {
	Token() (string, error)
	WithClientID(string)
	WithClientSecret(string)
	WithResource(string)
	WithTokenEndpoint(string)
}

// NewIdentityClient returns an implementation of IdentityClient using a default http client.
func NewIdentityClient() IdentityClient {
	// Reset expiration for a new client instance.
	expiration = time.Time{}

	return &client{
		c: http.DefaultClient,
	}
}

type client struct {
	c             *http.Client
	clientID      string
	clientSecret  string
	resource      string
	tokenEndpoint string
}

type token struct {
	AccessToken string `json:"access_token"`
	// TokenType   string `json:"token_type"`
	ExpiresIn int `json:"expires_in"`
}

type errorResponse struct {
	ErrorDescription string `json:"error_description"`
	Error            string `json:"error"`
}

// Token returns a cached token if it has not expired, otherwise it
// retrieves a new access token and sets the cached token.
func (c *client) Token() (string, error) {
	mux.Lock()
	defer mux.Unlock()

	// If the cached token has not expired just return it.
	if time.Now().In(time.UTC).Before(expiration) {
		return cachedToken, nil
	}

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", c.clientID)
	data.Set("client_secret", c.clientSecret)
	data.Set("resource", c.resource)

	// Create request and URL encode the form.
	r, err := http.NewRequestWithContext(context.Background(),
		http.MethodPost,
		c.tokenEndpoint,
		strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("error making request for THD IDP: %w", err)
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.c.Do(r)
	if err != nil {
		return "", fmt.Errorf("error doing request for new THD IDP token: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 399 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", fmt.Errorf("error getting token from THD IDP: %s", res.Status)
		}

		var e errorResponse

		err = json.Unmarshal(body, &e)
		if err != nil {
			return "", fmt.Errorf("error getting token from THD IDP: %s", res.Status)
		}

		return "", fmt.Errorf("error getting token from THD IDP: %s", e.ErrorDescription)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading body from THD IDP: %w", err)
	}

	var t token

	err = json.Unmarshal(body, &t)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling body from THD IDP: %w", err)
	}

	cachedToken = t.AccessToken
	expiration = time.Now().In(time.UTC).Add(time.Second * time.Duration((t.ExpiresIn/10)*9))

	return cachedToken, nil
}

// WithClientID sets the client ID, for example spiffe://homedepot.dev/om-api-security-client.
func (c *client) WithClientID(clientID string) {
	c.clientID = clientID
}

// WithClientSecret sets the client secret.
func (c *client) WithClientSecret(clientSecret string) {
	c.clientSecret = clientSecret
}

// WithResource sets the resource, for example spiffe://homedepot.dev/om-api-security-api.
func (c *client) WithResource(resource string) {
	c.resource = resource
}

// WithTokenEndpoint sets the token endpoint, for example https://identity-qa.homedepot.com/as/token.oauth2.
func (c *client) WithTokenEndpoint(tokenEndpoint string) {
	c.tokenEndpoint = tokenEndpoint
}

// Instance returns the client instance attached to the gin context.
func Instance(c *gin.Context) IdentityClient {
	return c.MustGet(IdentityClientInstanceKey).(IdentityClient)
}
