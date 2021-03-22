package middleware

import (
	"github.com/gin-gonic/gin"
	"github.homedepot.com/cd/cloud-runner/pkg/gcloud"
)

// SetBuilder attaches a gcloud.CloudRunCommandBuilder to the gin context.
func SetBuilder(b gcloud.CloudRunCommandBuilder) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(gcloud.BuilderInstanceKey, b)
		c.Next()
	}
}
