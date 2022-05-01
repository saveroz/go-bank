package tidb

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"database/sql"
	"math/rand"

	"github.com/saveroz/go-bank/models"
	"github.com/saveroz/go-bank/pkg/sql_transaction"
)

const MIN = 10000
const MAX = 99999

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (ar AccountRepository) Create(ctx context.Context, ac *models.Account) error {
	var err error

	randomNum := rand.Intn((MAX - MIN) + MIN)
	ac.ID = uuid.New().String()
	ac.Balance = 0
	ac.AccountNo = randomNum
	ac.Created = time.Now().UTC()
	ac.Updated = time.Now().UTC()

	fmt.Println(ac.ID)

	sqlStatement := `INSERT INTO account (id, name, balance, account_no, created, updated) VALUES (?, ?, ?, ?, ?, ?)`
	result, err := ar.db.ExecContext(
		ctx,
		sqlStatement,
		ac.ID,
		ac.Name,
		ac.Balance,
		ac.AccountNo,
		ac.Created,
		ac.Updated,
	)

	if err != nil {
		fmt.Println("error saat create account")
		return err
	}

	insertedRow, err := result.RowsAffected()

	if insertedRow != 1 || err != nil {
		fmt.Println("pembuatan account tidak sukses")
		return errors.New("failed to create account")
	}

	return err
}

func (ar AccountRepository) GetByAccountNo(ctx context.Context, accountNo int) (models.Account, error) {
	var err error
	var account models.Account

	// add for update to lock the account row
	sqlStatement := `SELECT id, name, balance, account_no, created, updated FROM account WHERE account_no=$1`

	tx, err := sql_transaction.GetTxFromContext(ctx)
	fmt.Println("found value:", tx)
	if err != nil {
		return account, err
	}

	row := tx.QueryRowContext(
		ctx,
		sqlStatement,
		accountNo,
	)

	err = row.Scan(
		&account.ID,
		&account.Name,
		&account.Balance,
		&account.AccountNo,
		&account.Created,
		&account.Updated,
	)

	if err != nil {
		fmt.Println("mamamia")
		fmt.Println(accountNo)
		fmt.Println("error saat fetch account")
		return account, err
	}

	return account, err
}

func (ar AccountRepository) Update(ctx context.Context, ac *models.Account) error {
	var err error

	ac.Updated = time.Now().UTC()
	sqlStatement := `UPDATE account SET name=$1, balance=$2, created=$3, updated=$4
	WHERE id=$5
	`

	tx, err := sql_transaction.GetTxFromContext(ctx)
	fmt.Println("found value in update account:", tx)
	if err != nil {
		return err
	}
	result, err := tx.ExecContext(
		ctx,
		sqlStatement,
		ac.Name,
		ac.Balance,
		ac.Created,
		ac.Updated,
		ac.ID,
	)

	if err != nil {
		fmt.Println("error saat update account")
		return err
	}

	insertedRow, err := result.RowsAffected()

	if insertedRow != 1 || err != nil {
		fmt.Println("pembuatan account tidak sukses")
		return errors.New("failed to create account")
	}

	return err
}
