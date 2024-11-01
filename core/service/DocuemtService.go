package service

import (
	"payment/core/domain"
	"payment/ports"
)

type DocumentService struct {
	DocumentRepository ports.DocumentRepository
}

func NewDocumentService(documentRepository ports.DocumentRepository) *DocumentService {
	return &DocumentService{documentRepository}
}

func (s DocumentService) GetAll() ([]domain.Document, error) {
	return s.DocumentRepository.GetAll()

}
