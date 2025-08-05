package middleware

import (
	"go-events-api/cmd/api/dto"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates JWT token in Authorization header
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		parts := strings.Fields(authHeader)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Bearer token is required"})
			return
		}

		/*
			More things to check here:
			- Token signature is valid
			- Token is not expired, etc
			- Check if user exists in database, check if user account is ACTIVE
			- Set user in context
		*/

		user := dto.User{
			ID:    1,
			Name:  "Paras Agrawal",
			Email: "paras.agrawal@gmail.com",
		}
		c.Set("user", user)

		c.Next()
	}
}
