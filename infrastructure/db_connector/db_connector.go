package db_connector

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strconv"
)

type DbConnector interface {
	GetConnection() *sql.DB
}

type MySqlConnector struct {
	config DbConfig
}

func NewMySqlConnector(config DbConfig) *MySqlConnector {
	return &MySqlConnector{config: config}
}

func (c MySqlConnector) GetConnection() *sql.DB {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%d)/%v", c.config.UserName, c.config.Password, c.config.Host, c.config.Port, c.config.DbName)

	db, err := sql.Open(c.config.DriverName, connectionString)
	if err != nil {
		panic(err)
	}

	return db
}

type DbConfig struct {
	DriverName string
	UserName   string
	Password   string
	Host       string
	DbName     string
	Port       int
}

func (c *DbConfig) SetDbConfig(mode string) {
	if mode == "debug" {

		c.DriverName = "mysql"
		c.DbName = "sweater_db"
		c.Host = "localhost"
		c.Port = 3306
		c.UserName = "root"
		c.Password = "qwerty12345"

	} else {
		port, err := strconv.Atoi(os.Getenv("DB_PORT"))
		if err != nil {
			panic(err)
		}

		c.DriverName = "mysql"
		c.DbName = os.Getenv("DB_NAME")
		c.Host = os.Getenv("DB_HOST")
		c.Port = port
		c.UserName = os.Getenv("DB_USER")
		c.Password = os.Getenv("DB_PASSWORD")
	}
}
