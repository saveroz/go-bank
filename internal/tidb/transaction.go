package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"database/sql"

	"github.com/saveroz/go-bank/models"
	"github.com/saveroz/go-bank/pkg/sql_transaction"
)

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *transactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (trxRepo transactionRepository) Create(ctx context.Context, t *models.Transaction) error {
	var err error

	t.ID = uuid.New().String()
	t.Created = time.Now().UTC()
	t.Updated = time.Now().UTC()

	tx, err := sql_transaction.GetTxFromContext(ctx)
	fmt.Println("found value in create transaction account:", tx)
	if err != nil {
		return err
	}

	sqlStatement := `INSERT INTO transaction
	(id, account_no, amount, type, transaction_type, description, created, updated)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	result, err := tx.ExecContext(
		ctx,
		sqlStatement,
		t.ID,
		t.AccountNo,
		t.Amount,
		t.Type,
		t.TransactionType,
		t.Description,
		t.Created,
		t.Updated,
	)

	if err != nil {
		fmt.Println(err)
		fmt.Println("error saat create transaction")
		return err
	}

	insertedRow, err := result.RowsAffected()

	if insertedRow != 1 || err != nil {
		fmt.Println("pembuatan account tidak sukses")
		return errors.New("failed to create account")
	}

	return err
}

func (trxRepo *transactionRepository) WithSqlTransaction(ctx context.Context, fn func(c context.Context) error) error {
	trxOption := sql.TxOptions{
		Isolation: 2,
		ReadOnly: false,
	}

	tx, err := trxRepo.db.BeginTx(ctx, &trxOption)

	if err != nil {
		fmt.Println("Failed to start transaction")
		return err
	}

	k := sql_transaction.SqlTrxKey("trxConnection")
	c := context.WithValue(ctx, k, tx)
	fmt.Println("DIDALEM TRANSCTION PROCESS")
	fmt.Println(c.Value(k))

	err = fn(c)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
