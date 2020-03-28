package user

import (
	"../../models"
	"./repository"
)

var registeredUsers []models.User

func GetAllUsers() []models.User {
	if len(registeredUsers) == 0 {
		registeredUsers = repository.GetAllUsers()
	}

	return registeredUsers
}

func GetUser(username string) models.User {
	
}

func Exists(username string) bool {
	for _, user := range registeredUsers {
		if user.Username == username {
			return true
		}
	}
	return false
}

func ValidateUser(username string, password string) bool {
	if Exists(username) {
		for _, user := range registeredUsers {
			if user.Username == username {

			}
		}
	}
	return false
}

func hash(username string, password string) string {

}