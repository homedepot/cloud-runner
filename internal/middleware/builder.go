package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/homedepot/cloud-runner/internal/gcloud"
)

// SetBuilder attaches a gcloud.CloudRunCommandBuilder to the gin context.
func SetBuilder(b gcloud.CloudRunCommandBuilder) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(gcloud.BuilderInstanceKey, b)
		c.Next()
	}
}
