package boards

import (
	"context"
	"github.com/SlavaShagalov/slavello/internal/models"
)

type Usecase interface {
	ListByWorkspace(ctx context.Context, workspaceID int) ([]models.Board, error)
}
