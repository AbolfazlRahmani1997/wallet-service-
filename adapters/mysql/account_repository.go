package mysql

import (
	"gorm.io/gorm"
	"payment/core/domain"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (a AccountRepository) Create(account *domain.Account) error {

	return a.db.Create(&account).Error
}

func (a AccountRepository) FindById(id string) (*domain.Account, error) {
	var account domain.Account
	return &account, a.db.First(&account, "id = ?", id).Error
}
func (a AccountRepository) Update(account domain.Account) error {
	return a.db.Save(account).Error
}
