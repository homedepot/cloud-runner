package middleware

import (
	"github.com/gin-gonic/gin"
	"github.homedepot.com/cd/cloud-runner/pkg/snowql"
)

// SetSnowQLClient attaches a snowql.Client to the gin context.
func SetSnowQLClient(sc snowql.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(snowql.ClientInstanceKey, sc)
		c.Next()
	}
}
