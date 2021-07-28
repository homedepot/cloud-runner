package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.homedepot.com/cd/cloud-runner/internal/api/v1"
	"github.homedepot.com/cd/cloud-runner/internal/fiat"
	"github.homedepot.com/cd/cloud-runner/internal/gcloud"
	"github.homedepot.com/cd/cloud-runner/internal/sql"
)

// Server hold the gin engine and any clients we need for the API.
type Server struct {
	apiKey     string
	e          *gin.Engine
	builder    gcloud.CloudRunCommandBuilder
	fiatClient fiat.Client
	sqlClient  sql.Client
}

// NewServer returns a new instance of Server.
func NewServer() Server {
	return Server{}
}

// WithAPIKey sets the API key required for create and delete operations.
func (s *Server) WithAPIKey(apiKey string) {
	s.apiKey = apiKey
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
	{
		api := s.e.Group("")
		api.GET("/healthz", ok)
	}

	{
		// Declare a controller to hold non request-scoped services.
		cc := &v1.Controller{
			FiatClient: s.fiatClient,
			SqlClient:  s.sqlClient,
			Builder:    s.builder,
		}
		api := s.e.Group("v1")
		api.POST("/credentials", requireAPIKey(s.apiKey), cc.CreateCredentials)
		api.DELETE("/credentials/:account", requireAPIKey(s.apiKey), cc.DeleteCredentials)
		api.GET("/credentials/:account", cc.GetCredentials)
		api.GET("/credentials", cc.ListCredentials)

		api.POST("/deployments", requireAPIKey(s.apiKey), cc.CreateDeployment)
		api.GET("/deployments/:deploymentID", cc.GetDeployment)
	}
}

func ok(*gin.Context) {}

// requireAPIKey is a middleware that returns 401 if the API-Key header
// is not set and 403 if the the API-Key is not correct.
func requireAPIKey(apiKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("API-Key") == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "API-Key header not set"})

			return
		}

		if c.GetHeader("API-Key") != apiKey {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "bad api key"})

			return
		}

		c.Next()
	}
}
