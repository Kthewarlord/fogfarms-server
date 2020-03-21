package user

import (
	"context"

	"../models"
)

type UseCase interface {
	Fetch(ctx context.Context, cursor string, num int) ([]*models.User, string, error)
	GetByID(ctx context.Context, id int) (*models.User, error)
	Update(ctx context.Context, ar *models.User) error
	GetByTitle(ctx context.Context, title string) (*models.User, error)
	Store(context.Context, *models.User) error
	Delete(ctx context.Context, id int) error
}
