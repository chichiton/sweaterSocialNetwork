package controllers

import (
	"errors"
	controllers "github.com/chichiton/sweaterSocialNetwork/controllers/internal"
	"github.com/chichiton/sweaterSocialNetwork/domain"
	"github.com/chichiton/sweaterSocialNetwork/infrastructure/repositories"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"net/http"
)

type FriendController struct {
	friendRepository *repositories.FriendRepositoryImp
}

func NewFriendController(friendRepository *repositories.FriendRepositoryImp) *FriendController {
	return &FriendController{friendRepository: friendRepository}
}

type Friend struct {
	FriendId domain.UserId `json:"friendId"`
}

func (r *FriendController) GetFriends(c *gin.Context) {
	userId, done := controllers.GetUserId(c)
	if done {
		return
	}
	friends, err := r.friendRepository.GetFriendsByUserId(*userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "request error"})
		c.Abort()
		return
	}
	c.JSON(200, friends)
}

func (r *FriendController) AddFriend(c *gin.Context) {
	userId, done := controllers.GetUserId(c)
	if done {
		return
	}

	var friend Friend
	err := c.BindJSON(&friend)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "request error"})
		c.Abort()
		return
	}

	if *userId == friend.FriendId {
		c.JSON(http.StatusBadRequest, gin.H{"error": "add yourself, jerk"})
		c.Abort()
		return

	}

	err = r.friendRepository.AddFriend(*userId, friend.FriendId)
	if err != nil {
		var dbError *mysql.MySQLError

		if errors.As(err, &dbError) {
			if dbError.Number == 1452 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect friend"})
				c.Abort()
				return
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "add friend error"})
		c.Abort()
		return
	}
}
