package account

import (
	"context"

	"github.com/saveroz/go-bank/models"
)

type IAccountRepository interface {
	Create(c context.Context, a *models.Account) error
}

type service struct {
	repo IAccountRepository
}

func NewService(repo IAccountRepository) service {
	return service{
		repo: repo,
	}
}

func (s service) Create(ctx context.Context, acc *models.Account) error {
	err := s.repo.Create(ctx, acc)
	return err
}
