package boards

import (
	"context"
	"github.com/SlavaShagalov/slavello/internal/models"
)

type Repository interface {
	List(ctx context.Context, workspaceID int) ([]models.Board, error)
}
