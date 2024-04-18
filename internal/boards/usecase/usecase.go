package usecase

import (
	"context"
	"github.com/SlavaShagalov/slavello/internal/boards"
	"github.com/SlavaShagalov/slavello/internal/models"
)

type usecase struct {
	repo boards.Repository
}

func New(repo boards.Repository) boards.Usecase {
	return &usecase{
		repo: repo,
	}
}

func (uc *usecase) ListByWorkspace(ctx context.Context, userID int) ([]models.Board, error) {
	return uc.repo.List(ctx, userID)
}
