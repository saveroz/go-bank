package transaction

import (
	"context"

	"github.com/saveroz/go-bank/models"
)

type ITransactionRepository interface {
	Create(ctx context.Context, t *models.Transaction) error
}

type service struct {
	repo ITransactionRepository
}

func NewTransaction(repo ITransactionRepository) service {
	return service{
		repo: repo,
	}
}

func (s service) Create(ctx context.Context, tr *models.Transaction) error {
	var err error
	err = s.repo.Create(ctx, tr)
	return err
}
