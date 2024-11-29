package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "No authorization header",
			})
			return
		}

		// TODO: Implement your token validation logic here
		// This could involve JWT validation, database checks, etc.

		c.Next()
	}
}
