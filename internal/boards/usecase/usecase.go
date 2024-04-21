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

func (uc *usecase) Create(ctx context.Context, params *boards.CreateParams) (models.Board, error) {
	return uc.repo.Create(ctx, params)
}

func (uc *usecase) ListByWorkspace(ctx context.Context, userID int) ([]models.Board, error) {
	return uc.repo.List(ctx, userID)
}

func (uc *usecase) ListByTitle(ctx context.Context, title string, userID int) ([]models.Board, error) {
	return uc.repo.ListByTitle(ctx, title, userID)
}

func (uc *usecase) Get(ctx context.Context, id int) (models.Board, error) {
	return uc.repo.Get(ctx, id)
}

func (uc *usecase) FullUpdate(ctx context.Context, params *boards.FullUpdateParams) (models.Board, error) {
	return uc.repo.FullUpdate(ctx, params)
}

func (uc *usecase) PartialUpdate(ctx context.Context, params *boards.PartialUpdateParams) (models.Board, error) {
	return uc.repo.PartialUpdate(ctx, params)
}

func (uc *usecase) Delete(ctx context.Context, id int) error {
	return uc.repo.Delete(ctx, id)
}
