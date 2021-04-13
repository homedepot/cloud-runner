package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/homedepot/cloud-runner/internal/api/v1"
	"github.com/homedepot/cloud-runner/internal/middleware"
)

var (
	apiKey string
)

func initialize(r *gin.Engine) {
	{
		api := r.Group("")
		api.GET("/healthz", ok)
	}

	{
		api := r.Group("v1")
		api.POST("/credentials", middleware.RequireAPIKey(apiKey), v1.CreateCredentials)
		api.DELETE("/credentials/:account", middleware.RequireAPIKey(apiKey), v1.DeleteCredentials)
		api.GET("/credentials/:account", v1.GetCredentials)
		api.GET("/credentials", v1.ListCredentials)

		api.POST("/deployments", middleware.RequireAPIKey(apiKey), v1.CreateDeployment)
		api.GET("/deployments/:deploymentID", v1.GetDeployment)
	}
}

func ok(*gin.Context) {}

// WithKey sets the API key required for create and delete operations.
func WithKey(a string) {
	apiKey = a
}
