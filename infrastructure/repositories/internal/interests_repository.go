package repositories

import (
	"database/sql"
	"fmt"
	"github.com/chichiton/sweaterSocialNetwork/domain"
)

func GetInterestsByUserId(db *sql.DB, userId domain.UserId) ([]domain.Interest, error) {
	getInterestsQuery := "select title from sweater_db.user_interest where user_id =?"

	rows, err := db.Query(getInterestsQuery, userId)

	var interests []domain.Interest

	for rows.Next() {
		var interest domain.Interest

		err = rows.Scan(&interest)
		if err != nil {
			return nil, fmt.Errorf("user_interst select error: %w", err)
		}
		interests = append(interests, interest)
	}
	return interests, nil
}
