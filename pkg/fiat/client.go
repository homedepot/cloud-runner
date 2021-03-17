package fiat

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ClientInstanceKey = `FiatClient`
	defaultFiatURL    = "http://spin-fiat.spinnaker:7003"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Client

// Client holds the implementation to call fiat.
type Client interface {
	Roles(string) (Roles, error)
	WithURL(string)
}

// NewClient returns an implementation of Client.
func NewClient() Client {
	return &client{}
}

// NewDefaultClient returns an implementation of Client with the URL
// set to http://spin-fiat.spinnaker:7003.
func NewDefaultClient() Client {
	c := NewClient()
	c.WithURL(defaultFiatURL)

	return c
}

type client struct {
	url string
}

type Roles []Role

type Role struct {
	Name   string `json:"name"`
	Source string `json:"source"`
}

// Roles lists roles for a given account.
func (c *client) Roles(account string) (Roles, error) {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet,
		fmt.Sprintf("%s/authorize/%s/roles", c.url, account), nil)
	if err != nil {
		return Roles{}, fmt.Errorf("fiat: error creating request: %w", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Roles{}, fmt.Errorf("fiat: error making request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 399 {
		return Roles{}, fmt.Errorf("fiat: error getting account roles: %s", res.Status)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Roles{}, fmt.Errorf("fiat: error reading response body: %w", err)
	}

	roles := Roles{}

	err = json.Unmarshal(b, &roles)
	if err != nil {
		return Roles{}, fmt.Errorf("fiat: error unmarshaling roles: %w", err)
	}

	return roles, nil
}

func (c *client) WithURL(url string) {
	c.url = url
}

// Instance returns the client instance attached to the gin context.
func Instance(c *gin.Context) Client {
	return c.MustGet(ClientInstanceKey).(Client)
}
