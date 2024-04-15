package bcrypt

import (
	"context"
	pkgHasher "github.com/SlavaShagalov/slavello/internal/pkg/hasher"
	"golang.org/x/crypto/bcrypt"
)

type hasher struct{}

func New() pkgHasher.Hasher {
	return &hasher{}
}

func (h *hasher) GetHashedPassword(ctx context.Context, password string) (string, error) {
	pswd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(pswd), err
}

func (h *hasher) CompareHashAndPassword(ctx context.Context, hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
