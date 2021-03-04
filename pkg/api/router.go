package api

import (
	"github.com/gin-gonic/gin"
)

// Init defines the Cloud Runner API.
func Init(r *gin.Engine) {
	{
		api := r.Group("")
		api.GET("/healthz", OK)
	}
}

func OK(*gin.Context) {}
