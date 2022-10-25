package repositories

import (
	"fmt"
	"github.com/chichiton/sweaterSocialNetwork/domain"
	"github.com/chichiton/sweaterSocialNetwork/infrastructure/db_connector"
	"github.com/chichiton/sweaterSocialNetwork/infrastructure/repositories/internal"
	"strings"
)

type UserRepositoryImp struct {
	dbConnector *db_connector.MySqlConnector
}

func NewUserRepositoryImp(dbConnector *db_connector.MySqlConnector) *UserRepositoryImp {
	return &UserRepositoryImp{dbConnector: dbConnector}
}

func (r UserRepositoryImp) RegisterUser(user domain.RegisterUser) (userId domain.UserId, err error) {
	db := r.dbConnector.GetConnection()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return userId, fmt.Errorf("db transaction error: %w", err)
	}

	defer tx.Rollback()

	sqlInsertUser := "insert into sweater_db.user(login, password_hash) values (?,?)"

	hash, err := user.Password.Hash()
	if err != nil {
		return userId, err
	}

	result, err := tx.Exec(sqlInsertUser, user.Login, hash)
	if err != nil {
		tx.Rollback()
		return userId, fmt.Errorf("insert error: %w", err)
	}

	id, err := result.LastInsertId()
	userId = domain.UserId(id)

	if err != nil {
		tx.Rollback()
		return userId, fmt.Errorf("inserted id error: %w", err)
	}

	sqlInsertProfile := "insert into user_profile(user_id, first_name, last_name, age, gender, city) values (?,?,?,?,?,?)"

	_, err = tx.Exec(sqlInsertProfile, userId, user.FirstName, user.LastName, user.Age, user.Gender, user.City)
	if err != nil {
		tx.Rollback()
		return userId, fmt.Errorf("insert error: %w", err)
	}

	var interestsStrings []string
	var interestsArgs []interface{}
	for _, in := range user.Interests {
		interestsStrings = append(interestsStrings, "(?, ?)")

		interestsArgs = append(interestsArgs, in)
		interestsArgs = append(interestsArgs, userId)
	}

	sqlInsertInterests := "insert ignore into sweater_db.user_interest(title, user_id) values %s"

	sqlInsertInterests = fmt.Sprintf(sqlInsertInterests, strings.Join(interestsStrings, ","))

	_, err = tx.Exec(sqlInsertInterests, interestsArgs...)
	if err != nil {
		tx.Rollback()
		return userId, fmt.Errorf("insert error: %w", err)
	}

	tx.Commit()

	return userId, err
}

func (r UserRepositoryImp) GetUserAuthByLogin(login domain.Login) (auth domain.Auth, err error) {
	db := r.dbConnector.GetConnection()
	defer db.Close()

	getAuthQuery := "select id, login, password_hash from sweater_db.user where login = ?"

	err = db.QueryRow(getAuthQuery, login).Scan(&auth.UserId, &auth.Login, &auth.PasswordHash)
	if err != nil {
		return auth, fmt.Errorf("user select error: %w", err)
	}
	return auth, nil
}

func (r UserRepositoryImp) GetUserProfile(userId domain.UserId) (domain.UserProfile, error) {
	db := r.dbConnector.GetConnection()
	defer db.Close()

	getProfileQuery := "select user_id, first_name, last_name, age, gender, city  from sweater_db.user_profile where user_id = ?"

	user := domain.UserProfile{}
	err := db.QueryRow(getProfileQuery, userId).Scan(&user.UserId, &user.FirstName, &user.LastName, &user.Age, &user.Gender, &user.City)
	if err != nil {
		return user, fmt.Errorf("user select error: %w", err)
	}

	user.Interests, err = repositories.GetInterestsByUserId(db, userId)
	if err != nil {
		return user, err
	}

	return user, nil
}
