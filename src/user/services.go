package user

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/user/repository"
	"golang.org/x/crypto/bcrypt"
)

var registeredUsers []models.User

func GetAllUsers() []models.User {
	if len(registeredUsers) == 0 {
		registeredUsers = repository.GetAllUsers()
	}

	return registeredUsers
}

func GetUser(username string) *models.User {
	if exists, user := Exists(username); exists {
		return user
	}
}

func Exists(username string) (bool, *models.User) {
	for _, user := range registeredUsers {
		if user.Username == username {
			return true, &user
		}
	}
	return false, nil
}

func ValidateUser(username string, password string) bool {
	if exists, user := Exists(username); exists {
		if user.Username == username {
			if user.Hash == hash(password, user.Salt) {
				return true
			}
		}
	}
	return false
}

func hash(password string, salt string) string {
	s := password + salt
	h, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(h)
}