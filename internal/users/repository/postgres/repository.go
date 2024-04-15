package postgres

import (
	"context"
	"database/sql"
	"github.com/SlavaShagalov/slavello/internal/models"
	"github.com/SlavaShagalov/slavello/internal/pkg/constants"
	pkgErrors "github.com/SlavaShagalov/slavello/internal/pkg/errors"
	pkgUsers "github.com/SlavaShagalov/slavello/internal/users"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type repository struct {
	db  *sql.DB
	log *zap.Logger
}

func New(db *sql.DB, log *zap.Logger) pkgUsers.Repository {
	return &repository{
		db:  db,
		log: log,
	}
}

const createCmd = `
	INSERT INTO users (name, username, email, hashed_password)
	VALUES ($1, $2, $3, $4)
	RETURNING id, username, hashed_password, email, name, avatar, created_at, updated_at;`

func (repo *repository) Create(ctx context.Context, params *pkgUsers.CreateParams) (models.User, error) {
	row := repo.db.QueryRow(createCmd, params.Name, params.Username, params.Email, params.HashedPassword)

	var user models.User
	err := scanUser(row, &user)
	if err != nil {
		repo.log.Error(constants.DBScanError, zap.Error(err), zap.String("sql_query", createCmd),
			zap.Any("create_params", params))
		return models.User{}, errors.Wrap(pkgErrors.ErrDb, err.Error())
	}

	repo.log.Debug("User created", zap.Int("id", user.ID), zap.String("username", user.Username))
	return user, nil
}

const getCmd = `
	SELECT id, username, hashed_password, email, name, avatar, created_at, updated_at
	FROM users
	WHERE id = $1;`

func (repo *repository) Get(id int) (models.User, error) {
	row := repo.db.QueryRow(getCmd, id)

	var user models.User
	err := scanUser(row, &user)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, errors.Wrap(pkgErrors.ErrUserNotFound, err.Error())
		}

		repo.log.Error(constants.DBScanError, zap.Error(err), zap.String("sql_query", getCmd),
			zap.Int("id", id))
		return models.User{}, errors.Wrap(pkgErrors.ErrDb, err.Error())
	}

	return user, nil
}

const getByUsernameCmd = `
	SELECT id, username, hashed_password, email, name, avatar, created_at, updated_at
	FROM users
	WHERE username = $1;`

func (repo *repository) GetByUsername(ctx context.Context, username string) (models.User, error) {
	row := repo.db.QueryRow(getByUsernameCmd, username)

	var user models.User
	err := scanUser(row, &user)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, pkgErrors.ErrUserNotFound
		}

		repo.log.Error(constants.DBScanError, zap.Error(err), zap.String("sql_query", getByUsernameCmd),
			zap.String("username", username))
		return models.User{}, pkgErrors.ErrDb
	}

	return user, nil
}

func scanUser(row *sql.Row, user *models.User) error {
	avatar := new(sql.NullString)
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.Name,
		avatar,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return err
	}

	if avatar.Valid {
		user.Avatar = &avatar.String
	} else {
		user.Avatar = nil
	}

	return nil
}
