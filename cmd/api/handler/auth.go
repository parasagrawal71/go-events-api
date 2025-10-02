package apiHandler

import (
	"go-events-api/cmd/api/config"
	"go-events-api/cmd/api/dto"
	"go-events-api/cmd/api/helpers"
	"go-events-api/cmd/api/models"
	"go-events-api/cmd/api/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterUser registers a new user
//
// @Summary Registers a new user
// @Description Registers a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body dto.RegisterUser true "User details"
// @Success 201 {object} dto.User
// @Router /api/v1/auth/register [post]
func RegisterUser(c *gin.Context) {
	var user dto.RegisterUser
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userToCreate := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	createdUser, err := repository.UserRepo.Create(&userToCreate)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, createdUser)
}

// LoginUser logs in a user
//
// @Summary Logs in a user
// @Description Logs in a user
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body dto.LoginUser true "User details"
// @Success 200 {object} gin.H
// @Router /api/v1/auth/login [post]
func LoginUser(c *gin.Context) {
	var user dto.LoginUser
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFromDB models.User
	config.DB.Where("email = ? AND password = ?", user.Email, user.Password).Find(&models.User{}).Scan(&userFromDB)
	if userFromDB.ID == 0 {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	token, err := helpers.GenerateJWT(userFromDB)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"token": token})
}
