package mysql

import (
	"gorm.io/gorm"
	"payment/core/domain"
)

type WalletRepositoryImpl struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepositoryImpl {
	return &WalletRepositoryImpl{db: db}
}

func (r *WalletRepositoryImpl) Create(wallet *domain.Wallet) error {
	return r.db.Create(wallet).Error
}

func (r *WalletRepositoryImpl) GetByID(id uint) (*domain.Wallet, error) {
	var wallet domain.Wallet
	if err := r.db.First(&wallet, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *WalletRepositoryImpl) GetByUserId(UserId uint, Currency domain.Currency) (*domain.Wallet, error) {
	var wallet domain.Wallet
	if err := r.db.Where("user_id = ? AND currency = ?", UserId, Currency).First(&wallet).Error; err != nil {
	}
	return &wallet, nil
}

func (r *WalletRepositoryImpl) Update(wallet *domain.Wallet) error {
	return r.db.Save(wallet).Error
}

func (r *WalletRepositoryImpl) GetMasterByCurrency(currency domain.Currency) (*domain.Wallet, error) {
	var wallet domain.Wallet
	return &wallet, r.db.Where("currency = ?", currency).Where("wallet_type = ?", "Master").First(&wallet).Error
}
