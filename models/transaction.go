package models

import (
	"context"
	"time"
)

// transaction type
const (
	topUp    = "TOPUP"    // money you move to bank account
	withdraw = "WITHDRAW" // withdraw your money from bank account
	transfer = "TRANSFER" // transfer your fund to another account or receive money from your account
)

// type to refer money in and out
const (
	debit  = "DEBIT"  // money out
	credit = "CREDIT" // money in
)

type Transaction struct {
	ID              string    `json:"id"`
	AccountNo       int       `json:"account_no" validate:"required"` //receiver of the money
	Amount          int       `json:"amount" validate:"required"`
	Type            string    `json:"type"`
	TransactionType string    `json:"transaction_type" validate:"required"`
	Description     string    `json:"description"`
	Created         time.Time `json:"created"`
	Updated         time.Time `json:"updated"`
}

type ITransactionService interface {
	// Transfer(ctx context.Context, t *Transaction) error
	Withdraw(ctx context.Context, t *Transaction) error
	TopUp(ctx context.Context, t *Transaction) error
}