package controllers

import (
	"errors"
	"github.com/chichiton/sweaterSocialNetwork/domain"
	"github.com/chichiton/sweaterSocialNetwork/infrastructure/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterController struct {
	userRepository *repositories.UserRepositoryImp
}

func NewRegisterController(userRepository *repositories.UserRepositoryImp) *RegisterController {
	return &RegisterController{userRepository: userRepository}
}

func (r *RegisterController) RegisterUser(c *gin.Context) {
	var registerUser domain.RegisterUser

	if err := c.BindJSON(&registerUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	userId, err := r.userRepository.RegisterUser(registerUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.Unwrap(err)})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"userId": userId})
}
