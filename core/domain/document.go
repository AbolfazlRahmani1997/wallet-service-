package domain

import (
	"time"
)

type DocumentType string

const (
	CashIn   DocumentType = "cash_in"
	Transfer DocumentType = "transfer"
	CashOut  DocumentType = "cash_out"
)

type DocumentStatus string

const (
	Created   DocumentStatus = "created"
	Processed DocumentStatus = "processed"
	Failed    DocumentStatus = "failed"
	Success   DocumentStatus = "success"
)

type Document struct {
	ID                uint `json:"id" gorm:"primaryKey"`
	WalletOrigin      uint `json:"wallet_origin"`
	WalletDestination uint `json:"wallet_destination"`
	GasFee            float64
	TrackingCode      string
	Type              DocumentType
	Currency          Currency
	Amount            float64
	Status            DocumentStatus
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
