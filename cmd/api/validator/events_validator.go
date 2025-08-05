package validator

import (
	"go-events-api/cmd/api/dto"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ValidateCreateEvent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var event dto.Event
		if err := c.BindJSON(&event); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if strings.Contains(event.Name, "@") {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "event name cannot contain special characters"})
			return
		}

		c.Next()
	}
}
