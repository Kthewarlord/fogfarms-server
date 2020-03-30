package repository

import (
	"database/sql"
	"fmt"
	"github.com/ddfsdd/fogfarms-server/models"
	"time"
)

const (
	DbHost = "localhost"
	DbPort = 5432
	DbUser = "fogfarms"
	DbPass = "fogfarms"
	DbName = "fogfarms-01"
)

var connectionString string = fmt.Sprintf("port=%d host=%s user=%s "+
	"password=%s dbname=%s sslmode=disable",
	DbPort, DbHost, DbUser, DbPass, DbName)

func GetAllUsers() []models.User {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	rows, err := db.Query("SELECT * FROM Users;")
	if err != nil {
		panic(err)
	}

	var users []models.User
	for rows.Next() {
		var id string
		var username string
		var hash string
		var salt string
		var r string
		err := rows.Scan(&id, &username, &hash, &salt, &r)
		if err != nil {
			panic(err)
		}

		role := models.AuthorizedUser
		if r == "Administrator" {
			role = models.Administrator
		}

		user := models.User{
			UserID:    id,
			Username:  username,
			Salt:      salt,
			Hash:      hash,
			Role:      role,
			CreatedAt: time.Time{},
		}

		users = append(users, user)
	}

	return users
}