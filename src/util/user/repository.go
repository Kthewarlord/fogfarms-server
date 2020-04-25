package user

import "github.com/KitaPDev/fogfarms-server/models"

type Repository interface {
	GetAllUsers() []models.User
	GetUserByUsername(username string) (*models.User, error)
	GetUserByID(UserID int) (*models.User, error)
	CreateUser(username string, password string, isAdministrator bool)
	ValidateUser(username string, inputPassword string) bool
}
