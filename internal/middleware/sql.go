package middleware

import (
	"github.com/gin-gonic/gin"
	"github.homedepot.com/cd/cloud-runner/internal/sql"
)

// SetSQLClient attaches a sql.Client to the gin context.
func SetSQLClient(sc sql.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(sql.ClientInstanceKey, sc)
		c.Next()
	}
}
