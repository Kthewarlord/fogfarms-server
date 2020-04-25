package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var db *sql.DB

// const (
// 	DbHost = "localhost"
// 	DbPort = 5432
// 	DbUser = "fogfarms"
// 	DbPass = "fogfarms"
// 	DbName = "fogfarms-01"
// 	SSLMODE = "disable"
// )
// const (
// 	DbHost  = "localhost"
// 	DbPort  = 5432
// 	DbUser  = "postgres"
// 	DbPass  = "postgres"
// 	DbName  = "fogfarms-01"
// 	SSLMODE = "disable"
// )

func GetDB() *sql.DB {
	var err error
	var DbPass string
	var DbPort int
	var DbHost string
	var DbUser string
	var DbName string
	var port = os.Getenv("PORT")
	var SSLMODE string
	if port != "" {
		DbPass = os.Getenv("DB_PASS")
		DbPort, err = strconv.Atoi(os.Getenv("DB_PORT"))
		DbHost = os.Getenv("DB_Host")
		DbUser = os.Getenv("DB_USER")
		DbName = os.Getenv("DB_NAME")
		SSLMODE = "require"

	} else {
		DbHost = "localhost"
		DbPort = 5432
		DbUser = "postgres"
		DbPass = "postgres"
		DbName = "fogfarms-01"
		SSLMODE = "disable"
	}

	if db == nil {
		connectionString := fmt.Sprintf("port=%d host=%s user=%s "+
			"password=%s dbname=%s sslmode=%s",
			DbPort, DbHost, DbUser, DbPass, DbName, SSLMODE)

		db, err = sql.Open("postgres", connectionString)
		if err != nil {
			panic(err)
		}
	}

	return db
}
