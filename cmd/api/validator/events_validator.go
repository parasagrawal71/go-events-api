package validator

import (
	"bytes"
	"encoding/json"
	"go-events-api/cmd/api/dto"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ValidateCreateEvent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var event dto.Event

		/*
			ISSUE:  empty request body {"error": "EOF"}
				The problem is how Gin handles request bodies:
				- In Go’s http.Request, the body is a stream (io.ReadCloser).
				- Once you read it (like with c.BindJSON(&event) in your middleware), it is consumed.
				- When your handler later tries c.BindJSON(&newEvent), the body is already empty → "EOF".
		*/
		// if err := c.BindJSON(&event); err != nil {
		// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }

		// Read the raw request body
		body, err := c.GetRawData()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Unmarshal for validation
		if err := json.Unmarshal(body, &event); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate the event
		if strings.Contains(event.Name, "@") {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "event name cannot contain special characters"})
			return
		}

		// Restore the body so the handler can read it
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		c.Next()
	}
}
