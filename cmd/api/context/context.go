package appContext

import (
	"go-events-api/cmd/api/helpers"

	"github.com/gin-gonic/gin"
)

func GetUserFromContext(c *gin.Context) helpers.User {
	user, exist := c.Get("user")
	if !exist {
		return helpers.User{}
	}

	return user.(helpers.User)
}
