package transaction

import (
	"context"
	"fmt"

	"github.com/saveroz/go-bank/models"
	"github.com/saveroz/go-bank/pkg/sql_transaction"
)

type ITransactionRepository interface {
	Create(ctx context.Context, t *models.Transaction) error
	WithSqlTransaction(context.Context, func(context.Context) error) error
}

type IAccountRepository interface {
	GetByAccountNo(context.Context, int) (models.Account, error)
	Update(context.Context, *models.Account) error
}

type service struct {
	transactionRepo ITransactionRepository
	accountRepo     IAccountRepository
}

func NewService(transactionRepo ITransactionRepository, accountRepo IAccountRepository) *service {
	return &service{
		transactionRepo: transactionRepo,
		accountRepo:     accountRepo,
	}
}

func (s service) TopUp(ctx context.Context, tr *models.Transaction) error {
	var err error
	k := sql_transaction.SqlTrxKey("trxConnection")

	err = s.transactionRepo.WithSqlTransaction(ctx, func(c context.Context) error {
		fmt.Println("DI SERVICE")
		fmt.Println(c.Value(k))
		fmt.Println("DI SERVICE")
		account, e := s.accountRepo.GetByAccountNo(c, tr.AccountNo)
		if e != nil {
			fmt.Println(e)
			return e
		}
		e = s.transactionRepo.Create(c, tr)
		if err != nil {
			//TODO: add rollback
			return e
		}
		account.Balance += tr.Amount
		e = s.accountRepo.Update(c, &account)

		if e != nil {
			//TODO: add rollback
			return e
		}
		return e
	})
	return err
}
