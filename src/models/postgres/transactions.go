package postgres

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const (
	TransactionColumnID = "id"
)

type Transaction struct {
	ID           int64         `gorm:"column:id;primaryKey;" json:"id"`
	Amount       float64       `gorm:"column:amount;" json:"amount"`
	Type         string        `gorm:"column:type;" json:"type"`
	ParentID     *int64        `gorm:"column:parent_id;index;" json:"parent_id,omitempty"`
	Transactions []Transaction `gorm:"foreignKey:parent_id;references:id" json:"-"`

	// Timestamps
	CreatedAt time.Time `gorm:"column:created_at;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at;" json:"deleted_at"`
}

func MigrateTransactions(ctx context.Context) error {
	err := DB().WithContext(ctx).AutoMigrate(&Transaction{})
	if err != nil {
		return errors.Wrap(err, "migrate transactions")
	}
	return nil
}

func InsertTransaction(ctx context.Context, transaction *Transaction) error {
	result := DB().WithContext(ctx).Model(&Transaction{}).Create(transaction)
	if result.Error != nil {
		return errors.Wrap(result.Error, "insert transaction")
	}
	return nil
}

func GetTransactionByID(ctx context.Context, id int64) (Transaction, bool, error) {
	var transaction Transaction
	result := DB().WithContext(ctx).Model(&Transaction{}).First(&transaction, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return transaction, false, nil
		}
		return transaction, false, errors.Wrap(result.Error, "find transaction")
	}
	return transaction, true, nil
}

func GetTransactionsByType(ctx context.Context, transactionType string) ([]int64, error) {
	transactionIDs := []int64{}
	result := DB().WithContext(ctx).Model(&Transaction{}).Where(&Transaction{Type: transactionType}).Pluck(TransactionColumnID, &transactionIDs)
	if result.Error != nil {
		return transactionIDs, errors.Wrap(result.Error, "find transaction")
	}
	return transactionIDs, nil
}

// Did not work, have to manually define level of nesting required
func GetTransactionsAmountSumByParentId(ctx context.Context, rootId int64) (Transaction, error) {
	var transactions Transaction
	result := DB().WithContext(ctx).Model(&Transaction{}).Preload("Transactions").First(&transactions, rootId)
	if result.Error != nil {
		return transactions, errors.Wrap(result.Error, "associate transactions")
	}

	return transactions, nil
}

func GetTransactionSum(ctx context.Context, parentID int64) (float64, error) {
	var transactions []Transaction
	err := DB().WithContext(ctx).Raw(
		`WITH RECURSIVE transaction_tree AS (
			SELECT id, amount, parent_id FROM transactions WHERE id = ?  
			UNION ALL
			SELECT t.id, t.amount, t.parent_id FROM transactions t INNER JOIN transaction_tree tt ON t.parent_id = tt.id
		)
		SELECT * FROM transaction_tree
	`, parentID).Scan(&transactions).Error
	if err != nil {
		return 0, err
	}

	var sum float64
	for _, txn := range transactions {
		sum += txn.Amount
	}
	return sum, nil
}
