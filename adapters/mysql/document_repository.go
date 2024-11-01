package mysql

import (
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
	"payment/core/domain"
)

type DocumentRepositoryImpl struct {
	db *gorm.DB
}

func NewDocumentRepository(db *gorm.DB) *DocumentRepositoryImpl {
	return &DocumentRepositoryImpl{db: db}
}

func (r *DocumentRepositoryImpl) GetAll() ([]domain.Document, error) {
	documents := make([]domain.Document, 0)
	r.db.Find(&documents)
	return documents, nil

}

func (r *DocumentRepositoryImpl) Create(document *domain.Document) error {
	return r.db.Create(document).Error
}

func (r *DocumentRepositoryImpl) GetByID(id ulid.ULID) (*domain.Document, error) {
	var document domain.Document
	if err := r.db.First(&document, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &document, nil
}

func (r *DocumentRepositoryImpl) Update(document *domain.Document) error {
	return r.db.Save(document).Error
}
