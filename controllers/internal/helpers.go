package controllers

import (
	"github.com/chichiton/sweaterSocialNetwork/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserId(c *gin.Context) (*domain.UserId, bool) {
	key, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userId not found in token"})
		c.Abort()
		return nil, true
	}

	userId, ok := key.(*domain.UserId)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userId not found in token"})
		c.Abort()
		return nil, true
	}
	return userId, false
}
