package mysql

import (
	"gorm.io/gorm"
)

type DBTransaction struct {
	db *gorm.DB
}

func NewDBTransaction(db *gorm.DB) *DBTransaction {
	return &DBTransaction{db: db}
}

func (t *DBTransaction) Begin() *gorm.DB {
	return t.db.Begin()
}

func (t *DBTransaction) Commit(tx *gorm.DB) *gorm.DB {
	return tx.Commit()
}

func (t *DBTransaction) Rollback(tx *gorm.DB) error {
	return tx.Rollback().Error
}
