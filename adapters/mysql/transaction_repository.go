package mysql

import (
	_ "github.com/oklog/ulid/v2"
	"gorm.io/gorm"
	"payment/core/domain"
)

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{db: db}
}

func (r *TransactionRepositoryImpl) Create(transaction *domain.Transaction) error {
	return r.db.Create(transaction).Error
}
