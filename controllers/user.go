package controllers

import (
	"github.com/chichiton/sweaterSocialNetwork/domain"
	"github.com/chichiton/sweaterSocialNetwork/infrastructure/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	userRepository *repositories.UserRepositoryImp
}

func NewUserController(userRepository *repositories.UserRepositoryImp) *UserController {
	return &UserController{userRepository: userRepository}
}

func (r *UserController) GetUserProfile(c *gin.Context) {
	userId := c.Keys["id"].(*domain.UserId)
	userProfile, err := r.userRepository.GetUserProfile(*userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, userProfile)
}
