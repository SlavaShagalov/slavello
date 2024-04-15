package usecase

import (
	"github.com/SlavaShagalov/slavello/internal/models"
	"github.com/SlavaShagalov/slavello/internal/users"
)

type usecase struct {
	usersRepo users.Repository
}

func New(rep users.Repository) users.Usecase {
	return &usecase{
		usersRepo: rep,
	}
}

func (uc *usecase) Get(id int) (models.User, error) {
	return uc.usersRepo.Get(id)
}
