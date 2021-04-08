package api

import (
	"github.com/gin-gonic/gin"
	"github.com/homedepot/cloud-runner/internal/pkg/fiat"
	"github.com/homedepot/cloud-runner/internal/pkg/gcloud"
	"github.com/homedepot/cloud-runner/internal/pkg/middleware"
	"github.com/homedepot/cloud-runner/internal/pkg/sql"
)

// Server hold the gin engine and any clients we need for the API.
type Server struct {
	e          *gin.Engine
	builder    gcloud.CloudRunCommandBuilder
	fiatClient fiat.Client
	sqlClient  sql.Client
}

// NewServer returns a new instance of Server.
func NewServer() Server {
	return Server{}
}

// WithBuilder sets the gcloud command builder.
func (s *Server) WithBuilder(b gcloud.CloudRunCommandBuilder) {
	s.builder = b
}

// WithEngine sets the gin engine instance for the server.
func (s *Server) WithEngine(e *gin.Engine) {
	s.e = e
}

// WithFiatClient sets the Fiat Client.
func (s *Server) WithFiatClient(fc fiat.Client) {
	s.fiatClient = fc
}

// WithSQLClient sets the sql client instance for the server.
func (s *Server) WithSQLClient(sc sql.Client) {
	s.sqlClient = sc
}

// Setup sets any global middlewares then initializes the API.
func (s *Server) Setup() {
	s.e.Use(middleware.SetBuilder(s.builder))
	s.e.Use(middleware.SetFiatClient(s.fiatClient))
	s.e.Use(middleware.SetSQLClient(s.sqlClient))

	initialize(s.e)
}
