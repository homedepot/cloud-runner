package middleware

import (
	"github.com/gin-gonic/gin"
	"github.homedepot.com/cd/cloud-runner/pkg/thd"
)

// SetTHDIdentityClient attaches a thd.IdentityClient to the gin context.
func SetTHDIdentityClient(tc thd.IdentityClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(thd.IdentityClientInstanceKey, tc)
		c.Next()
	}
}
