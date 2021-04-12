package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/homedepot/cloud-runner/internal/fiat"
)

// SetFiatClient attaches a fiat.Client to the gin context.
func SetFiatClient(fc fiat.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(fiat.ClientInstanceKey, fc)
		c.Next()
	}
}
