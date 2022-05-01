package sql_transaction

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type SqlTrxKey string

func GetTxFromContext(ctx context.Context) (*sql.Tx, error) {
	var tx *sql.Tx
	k := SqlTrxKey("trxConnection")
	tx = ctx.Value(k).(*sql.Tx)

	fmt.Println("GetTxFromContext")
	fmt.Println(tx)

	if tx == nil {
		return nil, errors.New("transaction connection not found")
	}
	return tx, nil

}
