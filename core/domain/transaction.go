package domain

import (
	"time"
)

type TransactionType string

const (
	Deposit  TransactionType = "deposit"
	Withdraw TransactionType = "withdraw"
	GasFee   TransactionType = "gas_fee"
)

type Transaction struct {
	ID          uint `gorm:"primaryKey"`
	DocumentID  uint `gorm:""`
	WalletId    uint
	Amount      float64
	Type        TransactionType
	Description string
	CreatedAt   time.Time
}
