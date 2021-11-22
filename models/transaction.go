package models

import (
	"context"
	"time"
)

type Transaction struct {
	ID              string    `json:"id"`
	DebitAccountNo  string    `json:"debit_account_no" validate:"required"`  //receiver of the money
	CreditAccountNo string    `json:"credit_account_no" validate:"required"` // sender of the money
	Money           int       `json:"money" validate:"required"`
	Description     string    `json:"description"`
	Created       time.Time   `json:"created"`
	Updated       time.Time   `json:"updated"`
}

type ITransactionService interface {
	Create(ctx context.Context, t *Transaction) error
}
