package domain

import (
	"github.com/oklog/ulid/v2"
	"time"
)

type Currency string

const (
	IRR  Currency = "rial"
	Gold Currency = "gold"
)

type WalletType string

const (
	Master WalletType = "Master"
	User   WalletType = "User"
)

type Wallet struct {
	ID         uint     `gorm:"primaryKey"`
	Currency   Currency `gorm:"index:,unique,composite:Currency_Account"`
	AccountId  uint     `json:"account_id" gorm:"index:,unique,composite:Currency_Account"`
	Balance    float64
	WalletType WalletType
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type WalletServiceInterface interface {
	CreateWallet(currency string) (*Wallet, error)
	GetWallet(id ulid.ULID) (*Wallet, error)
	UpdateWallet(wallet *Wallet) error
}
