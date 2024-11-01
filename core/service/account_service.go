package service

import (
	"payment/adapters/mysql"
	"payment/core/domain"
)

type AccountService struct {
	accountRepository mysql.AccountRepository
}

func NewAccountService(accountRepository mysql.AccountRepository) *AccountService {
	return &AccountService{accountRepository: accountRepository}
}

func (receiver AccountService) getAccount(id string) (*domain.Account, error) {
	return receiver.accountRepository.FindById(id)
}
