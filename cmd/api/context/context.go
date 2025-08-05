package appContext

import (
	"go-events-api/cmd/api/dto"

	"github.com/gin-gonic/gin"
)

func GetUserFromContext(c *gin.Context) dto.User {
	user, exist := c.Get("user")
	if !exist {
		return dto.User{}
	}

	return user.(dto.User)
}
