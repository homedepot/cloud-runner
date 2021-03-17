package snowql

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	queryGetLCP = `query getLCP($name:String!) {
			application(name: $name) {
				lifecyclePhase
			}
		}`
	ClientInstanceKey = `SnowQLClient`
)

var (
	variablesGetLCP = `{
			"name": "%s"
		}`
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Client

// Client implements calls to the SnowQL API.
type Client interface {
	GetLCP(string, string) (string, error)
	WithURL(string)
}

type client struct {
	c   *http.Client
	url string
}

// NewClient returns an instance of Client with the *http.Client
// set to the default client.
func NewClient() Client {
	return &client{
		c: http.DefaultClient,
	}
}

type getLCPRequest struct {
	Query     string `json:"query"`
	Variables string `json:"variables"`
}

type getLCPResponse struct {
	Data struct {
		Application *struct {
			LifecyclePhase string `json:"lifecyclePhase"`
		} `json:"application"`
	} `json:"data"`
}

// GetLCP returns the lifecycle phase for a given project ID. It takes in
// the bearer token since this is generated using THD IDP and expires every hour.
//
// We could call THD IDP from this client to get the bearer token, but I don't like attaching clients to clients.
func (c *client) GetLCP(projectID, token string) (string, error) {
	requestBody := getLCPRequest{
		Query:     queryGetLCP,
		Variables: fmt.Sprintf(variablesGetLCP, projectID),
	}

	b, err := json.Marshal(&requestBody)
	if err != nil {
		return "", fmt.Errorf("snowql: error marshaling request body: %w", err)
	}

	r, err := http.NewRequestWithContext(context.Background(), http.MethodPost, fmt.Sprintf("%s/graphql", c.url), bytes.NewBuffer(b))
	if err != nil {
		return "", fmt.Errorf("snowql: error creating request: %w", err)
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := c.c.Do(r)
	if err != nil {
		return "", fmt.Errorf("snowql: error making request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 399 {
		return "", fmt.Errorf("error getting LCP: %s", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("snowql: error reading response body: %w", err)
	}

	var lcpResp getLCPResponse

	err = json.Unmarshal(body, &lcpResp)
	if err != nil {
		return "", fmt.Errorf("snowql: error unmarshaling response body: %w", err)
	}

	if lcpResp.Data.Application == nil {
		return "", fmt.Errorf("project ID '%s' not found", projectID)
	}

	return lcpResp.Data.Application.LifecyclePhase, nil
}

// WithURL sets the SnowQL URL, for example https://snowql-dot-io1-datalake-dev.appspot.com.
func (c *client) WithURL(url string) {
	c.url = url
}

// Instance returns the client instance attached to the gin context.
func Instance(c *gin.Context) Client {
	return c.MustGet(ClientInstanceKey).(Client)
}
