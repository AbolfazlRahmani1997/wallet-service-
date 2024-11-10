package service

import (
	"payment/adapters/mysql"
	"payment/core/domain"
)

type AccountService struct {
	accountRepository *mysql.AccountRepository
}

func NewAccountService(accountRepository *mysql.AccountRepository) *AccountService {
	return &AccountService{accountRepository: accountRepository}
}

func (receiver AccountService) GetAccount(id string) (*domain.Account, error) {
	return receiver.accountRepository.FindById(id)
}

func (receiver AccountService) SaveAccount(account *domain.Account) error {
	err := receiver.accountRepository.Create(account)
	if err != nil {
		return err
	}
	return nil

}
