package repositories

import (
	"github.com/chichiton/sweaterSocialNetwork/domain"
	"github.com/chichiton/sweaterSocialNetwork/infrastructure/db_connector"
)

type UserRepository interface {
	db_connector.DbConnector
	GetUserAuthByLogin(login domain.Login) (auth domain.Auth, err error)
	GetUserProfile(userId domain.UserId) (domain.UserProfile, error)
	RegisterUser(user domain.RegisterUser) (userId domain.UserId, err error)
}

type FriendRepository interface {
	db_connector.DbConnector
	AddFriend(userId domain.UserId, friendId domain.UserId) error
	GetFriendsByUserId(userId domain.UserId) ([]domain.UserProfile, error)
}
