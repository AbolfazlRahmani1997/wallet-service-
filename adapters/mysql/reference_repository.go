package mysql

import (
	"gorm.io/gorm"
	"payment/core/domain"
)

type ReferenceRepository struct {
	db *gorm.DB
}

func NewReferenceRepository(db *gorm.DB) *ReferenceRepository {
	return &ReferenceRepository{db: db}
}

func (receiver ReferenceRepository) Create(reference *domain.Reference) error {
	return receiver.db.Create(reference).Error
}

func (receiver ReferenceRepository) FindByToken(token string) domain.Reference {
	var reference domain.Reference
	receiver.db.First(&reference, "token = ?", token)
	return reference
}
