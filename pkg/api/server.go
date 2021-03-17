package api

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.homedepot.com/cd/cloud-runner/pkg/fiat"
	"github.homedepot.com/cd/cloud-runner/pkg/middleware"
	"github.homedepot.com/cd/cloud-runner/pkg/snowql"
	"github.homedepot.com/cd/cloud-runner/pkg/sql"
	"github.homedepot.com/cd/cloud-runner/pkg/thd"
)

var (
	errEngineNotDefined = errors.New("engine not defined")
)

// Server hold the gin engine and any clients we need for the API.
type Server struct {
	e                 *gin.Engine
	fiatClient        fiat.Client
	snowQLClient      snowql.Client
	sqlClient         sql.Client
	thdIdentityClient thd.IdentityClient
}

// NewServer returns a new instance of Server.
func NewServer() Server {
	return Server{}
}

// WithEngine sets the gin engine instance for the server.
func (s *Server) WithEngine(e *gin.Engine) {
	s.e = e
}

// WithFiatClient sets the Fiat Client.
func (s *Server) WithFiatClient(fc fiat.Client) {
	s.fiatClient = fc
}

// WithSnowQLClient sets the SnowQL Client.
func (s *Server) WithSnowQLClient(sc snowql.Client) {
	s.snowQLClient = sc
}

// WithSQLClient sets the sql client instance for the server.
func (s *Server) WithSQLClient(sc sql.Client) {
	s.sqlClient = sc
}

// WithTHDIdentityClient sets the THD Identity Client.
func (s *Server) WithTHDIdentityClient(tc thd.IdentityClient) {
	s.thdIdentityClient = tc
}

// Setup sets any global middlewares then initalizes the API.
func (s *Server) Setup() error {
	if s.e == nil {
		return errEngineNotDefined
	}

	s.e.Use(middleware.SetFiatClient(s.fiatClient))
	s.e.Use(middleware.SetSnowQLClient(s.snowQLClient))
	s.e.Use(middleware.SetSQLClient(s.sqlClient))
	s.e.Use(middleware.SetTHDIdentityClient(s.thdIdentityClient))

	initialize(s.e)

	return nil
}
