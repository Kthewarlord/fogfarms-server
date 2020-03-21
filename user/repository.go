package user

import (
	"../models"
)

//Repository repository interface
type Repository interface {
	Find(userID string) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	FindByChangePasswordHash(hash string) (*models.User, error)
	FindAll() ([]*models.User, error)
	Update(user *models.User) error
	Store(user *models.User) (string, error)
}
