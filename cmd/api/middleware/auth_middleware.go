package middleware

import (
	"go-events-api/cmd/api/config"
	"go-events-api/cmd/api/helpers"
	"go-events-api/cmd/api/models"
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
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" || parts[1] == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Bearer token is required"})
			return
		}

		// Verify JWT token
		token := parts[1]
		claims, err := helpers.VerifyJWT(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// check if user exists
		var user models.User
		config.DB.Raw("SELECT * FROM users WHERE id = ?;", claims.User.ID).Scan(&user)
		if user.ID == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not registered"})
			return
		}

		// Set user in context
		c.Set("user", claims.User)

		c.Next()
	}
}
