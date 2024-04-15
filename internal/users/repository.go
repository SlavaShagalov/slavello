package users

import (
	"context"
	"github.com/SlavaShagalov/slavello/internal/models"
)

type CreateParams struct {
	Name           string
	Username       string
	Email          string
	HashedPassword string
}

type Repository interface {
	Create(ctx context.Context, params *CreateParams) (models.User, error)
	GetByUsername(ctx context.Context, username string) (models.User, error)
	Get(id int) (models.User, error)
}
