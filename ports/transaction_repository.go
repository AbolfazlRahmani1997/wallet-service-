package ports

import (
	"payment/core/domain"
)

type TransactionRepository interface {
	Create(transaction *domain.Transaction) error
}
