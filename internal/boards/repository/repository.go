package repository

import (
	"context"
	"database/sql"
	pkgBoards "github.com/SlavaShagalov/slavello/internal/boards"
	"github.com/SlavaShagalov/slavello/internal/models"
	"github.com/SlavaShagalov/slavello/internal/pkg/constants"
	pkgErrors "github.com/SlavaShagalov/slavello/internal/pkg/errors"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type repository struct {
	db  *sql.DB
	log *zap.Logger
}

func New(db *sql.DB, log *zap.Logger) pkgBoards.Repository {
	return &repository{db: db, log: log}
}

const listCmd = `
	SELECT id, workspace_id, title, description, background, created_at, updated_at
	FROM boards
	WHERE workspace_id = $1;`

func (repo *repository) List(ctx context.Context, workspaceID int) ([]models.Board, error) {
	rows, err := repo.db.Query(listCmd, workspaceID)
	if err != nil {
		repo.log.Error(constants.DBError, zap.Error(err), zap.String("sql_query", listCmd),
			zap.Int("workspace_id", workspaceID))
		return nil, errors.Wrap(pkgErrors.ErrDb, err.Error())
	}
	defer func() {
		_ = rows.Close()
	}()

	boards := []models.Board{}
	var board models.Board
	var description sql.NullString
	background := new(sql.NullString)
	for rows.Next() {
		err = rows.Scan(
			&board.ID,
			&board.WorkspaceID,
			&board.Title,
			&description,
			background,
			&board.CreatedAt,
			&board.UpdatedAt,
		)
		if err != nil {
			repo.log.Error(constants.DBScanError, zap.Error(err), zap.String("sql_query", listCmd),
				zap.Int("workspace_id", workspaceID))
			return nil, errors.Wrap(pkgErrors.ErrDb, err.Error())
		}

		if background.Valid {
			board.Background = &background.String
		} else {
			board.Background = nil
		}
		board.Description = description.String

		boards = append(boards, board)
	}

	return boards, nil
}
