package account

import (
	"context"

	"github.com/saveroz/go-bank/models"
)

type IAccountRepository interface {
	Create(context.Context, *models.Account) error
	GetByAccountNo(context.Context, int) (models.Account, error)
	Update(context.Context, *models.Account) error
}

type service struct {
	repo IAccountRepository
}

func NewService(repo IAccountRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s service) Create(ctx context.Context, acc *models.Account) error {
	err := s.repo.Create(ctx, acc)
	return err
}

func (s service) Get(ctx context.Context, accountNo int) (models.Account, error) {
	account,err := s.repo.GetByAccountNo(ctx, accountNo)
	return account, err
}
