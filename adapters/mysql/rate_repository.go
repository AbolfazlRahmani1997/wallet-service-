package mysql

import (
	"gorm.io/gorm"
	"payment/core/domain"
)

type RateRepository struct {
	db *gorm.DB
}

func NewRateRepository(db *gorm.DB) RateRepository {

	return RateRepository{db: db}
}

func (r RateRepository) GetRate(Currency domain.Currency) (*domain.Rate, error) {
	var rate domain.Rate
	r.db.Where(&domain.Rate{Currency: Currency}).First(&rate)
	return &rate, nil
}

func (r RateRepository) Update(amount float64, Currency domain.Currency) (*domain.Rate, error) {
	var rate domain.Rate
	r.db.Where(&domain.Rate{Currency: Currency}).First(&rate)
	r.db.Model(&rate).Update("amount", amount)
	return &rate, nil
}
