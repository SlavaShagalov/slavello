package users

import (
	"github.com/SlavaShagalov/slavello/internal/models"
)

type Usecase interface {
	Get(id int) (models.User, error)
}
