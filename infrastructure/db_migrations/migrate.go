package db_migrations

import (
	"database/sql"
	"fmt"
	"github.com/chichiton/sweaterSocialNetwork/infrastructure/db_connector"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"time"
)

func CreateDatabase(config db_connector.DbConfig) {
	dbName := config.DbName
	config.DbName = ""

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%d)/%v", config.UserName, config.Password, config.Host, config.Port, config.DbName)

	var db *sql.DB
	var connectionError error

	for i := 0; i <= 10; i++ {
		db, connectionError = sql.Open(config.DriverName, connectionString)
		if connectionError == nil {
			connectionError = db.Ping()

			if connectionError != nil {
				if i < 10 {
					log.Println(connectionError)
					time.Sleep(time.Duration(i) * time.Second)
					continue
				} else {
					panic(connectionError)
				}
			}
		}
	}

	//db, err := sql.Open(config.DriverName, connectionString)
	//if err != nil {
	//	panic(err)
	//}

	//db := db_connector.NewMySqlConnector(config).GetConnection()

	defer db.Close()

	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		panic(err)
	}
}

func Migrate(connector *db_connector.MySqlConnector, config db_connector.DbConfig) {
	db := connector.GetConnection()
	defer db.Close()

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		`file://./infrastructure/db_migrations`,
		config.DbName, driver)
	if err != nil {
		panic(err)
	}
	err = m.Up()
	if err != nil && err.Error() != "no change" {
		panic(err)
	}
}

func pingDb(db *sql.DB) {
	for i := 0; i < 5; i++ {
		err := db.Ping()
		time.Sleep(time.Duration(i) * time.Second)
		if err != nil {
			log.Println(err)
			if i == 4 {
				panic(err)
			}
		} else {
			break
		}
	}
}
