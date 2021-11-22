package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"database/sql"
	"math/rand"

	"github.com/saveroz/go-bank/models"
)

const MIN = 10000
const MAX = 99999

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) AccountRepository {
	return AccountRepository{
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

	sqlStatement := `INSERT INTO account (id, name, balance, account_no, created, updated) VALUES ($1, $2, $3, $4, $5, $6)`
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

func (ar AccountRepository) Fetch(ctx context.Context, accountNo string) (models.Account, error) {
	var err error
	var account models.Account

	sqlStatement := `SELECT id, name, balance, account_no, created, updated WHERE account_no=$1`
	row := ar.db.QueryRowContext(
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
		fmt.Println("error saat fetch account")
		return account,err
	}

	return account,err
}
