package apiHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "user registered"})
}

func LoginUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "user logged in"})
}
