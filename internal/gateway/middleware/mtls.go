package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MTLSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.TLS == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "TLS required"})
			return
		}

		if len(c.Request.TLS.PeerCertificates) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Client certificate required"})
			return
		}

		c.Next()
	}
}
