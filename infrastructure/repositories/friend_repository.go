package repositories

import (
	"fmt"
	"github.com/chichiton/sweaterSocialNetwork/domain"
	"github.com/chichiton/sweaterSocialNetwork/infrastructure/db_connector"
	repositories "github.com/chichiton/sweaterSocialNetwork/infrastructure/repositories/internal"
)

type FriendRepositoryImp struct {
	DbConnector *db_connector.MySqlConnector
}

func NewFriendRepository(dbConnector *db_connector.MySqlConnector) *FriendRepositoryImp {
	return &FriendRepositoryImp{DbConnector: dbConnector}
}

func (r FriendRepositoryImp) AddFriend(userId domain.UserId, friendId domain.UserId) error {

	db := r.DbConnector.GetConnection()
	defer db.Close()

	addFriendSqlQuery := "insert into sweater_db.user_friend (user_id, friend_id) VALUES (?,?)"

	_, err := db.Exec(addFriendSqlQuery, userId, friendId)
	if err != nil {
		return fmt.Errorf("insert friend error: %w", err)
	}

	_, err = db.Exec(addFriendSqlQuery, friendId, userId)
	if err != nil {
		return fmt.Errorf("insert friend error: %w", err)
	}

	return nil
}

func (r FriendRepositoryImp) GetFriendsByUserId(userId domain.UserId) ([]domain.UserProfile, error) {
	db := r.DbConnector.GetConnection()
	defer db.Close()

	getFriendsSqlQuery :=
		`select p.user_id, p.first_name, p.last_name, p.age, p.gender,p.city from sweater_db.user_friend f
				join sweater_db.user_profile p on f.friend_id = p.user_id
				where f.user_id = ?`

	rows, err := db.Query(getFriendsSqlQuery, userId)
	if err != nil {
		return nil, fmt.Errorf("select friend error: %w", err)
	}

	var userFriends []domain.UserProfile
	for rows.Next() {
		var friend domain.UserProfile

		err = rows.Scan(&friend.UserId, &friend.FirstName, &friend.LastName, &friend.Age, &friend.Gender, &friend.City)
		if err != nil {
			return nil, fmt.Errorf("user_friends select error: %w", err)
		}

		interests, err := repositories.GetInterestsByUserId(db, friend.UserId)
		if err != nil {
			return nil, fmt.Errorf("user_friends select error: %w", err)
		}

		friend.Interests = interests

		userFriends = append(userFriends, friend)
	}

	return userFriends, nil
}
