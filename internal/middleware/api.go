package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireAPIKey(apiKey string) gin.HandlerFunc {
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
