package core

import (
	"context"

	"github.com/pkg/errors"
	"github.com/valiant1012/transaction-service/src/models/postgres"
)

type TransactionRequest struct {
	ID       int64   `json:"-"`
	Amount   float64 `json:"amount"`
	Type     string  `json:"type"`
	ParentId int64   `json:"parent_id"`
}

func (t *TransactionRequest) Validate() error {
	if t.Type == "" {
		return errors.New("missing transaction type")
	}
	return nil
}

func CreateTransaction(ctx context.Context, transactionRequest TransactionRequest) (postgres.Transaction, error) {
	transaction := postgres.Transaction{
		Amount:   transactionRequest.Amount,
		Type:     transactionRequest.Type,
		ParentID: transactionRequest.ParentId,
	}

	err := postgres.InsertTransaction(ctx, &transaction)
	if err != nil {
		return transaction, errors.Wrap(err, "insert transaction")
	}

	return transaction, nil
}

func GetTransactionByID(ctx context.Context, id int64) (postgres.Transaction, bool, error) {
	return postgres.GetTransactionByID(ctx, id)
}

func GetTransactionByType(ctx context.Context, transactionType string) ([]int64, error) {
	transactionIDs, err := postgres.GetTransactionsByType(ctx, transactionType)
	if err != nil {
		return nil, errors.Wrap(err, "get transactions")
	}

	return transactionIDs, nil
}
