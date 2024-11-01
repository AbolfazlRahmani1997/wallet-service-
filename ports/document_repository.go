package ports

import (
	"github.com/oklog/ulid/v2"
	"payment/core/domain"
)

type DocumentRepository interface {
	Create(document *domain.Document) error
	GetByID(id ulid.ULID) (*domain.Document, error)
	Update(document *domain.Document) error
	GetAll() ([]domain.Document, error)
}
