package ports

import (
	"github.com/oklog/ulid/v2"
	"payment/core/domain"
)

type WalletRepository interface {
	Create(wallet *domain.Wallet) error
	GetByID(id ulid.ULID) (*domain.Wallet, error)
	Update(wallet *domain.Wallet) error
}
