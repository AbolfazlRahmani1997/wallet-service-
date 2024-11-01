package service

import (
	"errors"
	"fmt"
	"github.com/sqids/sqids-go"
	_ "github.com/sqids/sqids-go"
	"payment/adapters/mysql"
	"payment/core/domain"
	"time"
)

type WalletService struct {
	walletIdMaster        uint
	WalletRepository      *mysql.WalletRepositoryImpl
	RateRepository        *mysql.RateRepository
	DocumentRepository    *mysql.DocumentRepositoryImpl
	TransactionRepository *mysql.TransactionRepositoryImpl
	dbTransaction         *mysql.DBTransaction
}

func NewWalletService(dbTransaction *mysql.DBTransaction, WalletRepository *mysql.WalletRepositoryImpl, DocumentRepository *mysql.DocumentRepositoryImpl, walletIdMaster string, TransactionRepository *mysql.TransactionRepositoryImpl) WalletService {
	return WalletService{dbTransaction: dbTransaction, WalletRepository: WalletRepository, DocumentRepository: DocumentRepository, walletIdMaster: 1, TransactionRepository: TransactionRepository}
}

func (ws *WalletService) CreateWallet() (domain.Wallet, error) {

	wallet := domain.Wallet{Balance: 0, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	err := ws.WalletRepository.Create(&wallet)
	if err != nil {
		return wallet, err
	}
	return wallet, nil
}

//func (ws *WalletService) GetWallet(id ulid.ULID) (*Wallet, error) {
//
//}
//
//func (ws *WalletService) UpdateWallet(wallet *Wallet) error {
//
//}

func (ws *WalletService) CashIn(WalletId uint, Amount float64) (domain.Document, error) {

	tx := ws.dbTransaction.Begin()
	defer func() {
		if p := recover(); p != nil {
			ws.dbTransaction.Rollback(tx)

		} else {
			ws.dbTransaction.Commit(tx)
		}
	}()
	wallet, err := ws.WalletRepository.GetByID(WalletId)
	if err != nil {
		return domain.Document{}, err
	}
	document := domain.Document{

		WalletOrigin:      ws.walletIdMaster,
		WalletDestination: wallet.ID,
		TrackingCode:      "",
		Type:              domain.CashIn,
		Amount:            Amount,
		Status:            domain.Created,
		CreatedAt:         time.Time{},
		UpdatedAt:         time.Time{},
	}
	err = ws.DocumentRepository.Create(&document)
	if err != nil {
		return domain.Document{}, err
	}
	document.Status = domain.Processed
	err = ws.DocumentRepository.Update(&document)
	if err != nil {
		return domain.Document{}, err
	}
	transaction := domain.Transaction{
		WalletId:   ws.walletIdMaster,
		Amount:     -Amount,
		Type:       domain.Deposit,
		DocumentID: document.ID,
	}
	err = ws.TransactionRepository.Create(&transaction)
	if err != nil {
		document.Status = domain.Failed
		err := ws.DocumentRepository.Update(&document)
		if err != nil {
			return domain.Document{}, err
		}
		return domain.Document{}, err
	}
	transactionWithDraw := domain.Transaction{
		WalletId:   wallet.ID,
		Amount:     +Amount,
		Type:       domain.Withdraw,
		DocumentID: document.ID,
	}
	err = ws.TransactionRepository.Create(&transactionWithDraw)
	if err != nil {
		document.Status = domain.Failed
		err := ws.DocumentRepository.Update(&document)
		if err != nil {
			return domain.Document{}, err
		}
		return domain.Document{}, err
	}
	document.Status = domain.Success
	err = ws.DocumentRepository.Update(&document)
	if err != nil {
		return domain.Document{}, err
	}
	return document, nil
}

func (ws *WalletService) Transfer(amount float64, WalletOriginId uint, WalletDestinationId uint) (domain.Document, error) {

	tx := ws.dbTransaction.Begin()
	defer func() {

		if p := recover(); p != nil {
			ws.dbTransaction.Rollback(tx)

		} else {
			ws.dbTransaction.Commit(tx)
		}
	}()
	WalletOrigin, err := ws.WalletRepository.GetByID(WalletOriginId)
	if err != nil {
		fmt.Print(err)
		return domain.Document{}, err
	}

	if WalletOrigin.Balance < amount+100 {
		return domain.Document{}, errors.New("not enough money")
	}

	WalletDestination, err := ws.WalletRepository.GetByID(WalletDestinationId)
	document := domain.Document{
		WalletOrigin:      WalletOrigin.ID,
		WalletDestination: WalletDestination.ID,
		TrackingCode:      GenerateTrackId(),
		GasFee:            100,
		Type:              domain.Transfer,
		Amount:            amount,
		Status:            domain.Created,
		CreatedAt:         time.Time{},
		UpdatedAt:         time.Time{},
	}
	err = ws.DocumentRepository.Create(&document)

	transaction := &domain.Transaction{
		WalletId:    WalletOrigin.ID,
		DocumentID:  document.ID,
		Amount:      -amount,
		Type:        domain.Withdraw,
		CreatedAt:   time.Time{},
		Description: "Withdraw",
	}
	err = ws.TransactionRepository.Create(transaction)
	if err != nil {
		return domain.Document{}, err
	}
	transaction = &domain.Transaction{
		WalletId:    WalletOrigin.ID,
		DocumentID:  document.ID,
		Amount:      -document.GasFee,
		Type:        domain.GasFee,
		CreatedAt:   time.Time{},
		Description: "GasFee",
	}
	err = ws.TransactionRepository.Create(transaction)
	if err != nil {
		return domain.Document{}, err
	}
	transaction = &domain.Transaction{
		WalletId:    ws.walletIdMaster,
		DocumentID:  document.ID,
		Amount:      document.GasFee,
		Type:        domain.GasFee,
		CreatedAt:   time.Time{},
		Description: "GasFee",
	}
	err = ws.TransactionRepository.Create(transaction)
	if err != nil {
		return domain.Document{}, err
	}
	transaction = &domain.Transaction{
		WalletId:    WalletDestination.ID,
		DocumentID:  document.ID,
		Amount:      amount,
		Type:        domain.Deposit,
		CreatedAt:   time.Time{},
		Description: "Deposit",
	}
	err = ws.TransactionRepository.Create(transaction)
	document.Status = domain.Success
	err = ws.DocumentRepository.Update(&document)
	if err != nil {
		return domain.Document{}, err
	}

	return document, nil
}

func (s *WalletService) Change(WalletOriginId uint, amount float64, to domain.Currency) (*domain.Document, error) {
	tx := s.dbTransaction.Begin()
	walletOrigin, _ := s.WalletRepository.GetByID(WalletOriginId)
	walletDestinationMaster, _ := s.WalletRepository.GetMasterByCurrency(to)
	if walletOrigin.Balance < amount+100 {
		return nil, errors.New("not enough money")
	}
	WalletDestination, err := s.WalletRepository.GetByUserId(walletOrigin.AccountId, to)
	document := &domain.Document{WalletOrigin: WalletOriginId, WalletDestination: WalletDestination.ID, TrackingCode: GenerateTrackId(), Amount: amount, Currency: to}
	err = s.DocumentRepository.Create(document)
	if err != nil {
		return nil, err
	}

	DestinationRate, _ := s.RateRepository.GetRate(to)
	OriginRate, _ := s.RateRepository.GetRate(walletOrigin.Currency)
	ConvertedAmount := amount * OriginRate.Amount / DestinationRate.Amount
	err = s.DocumentRepository.Create(document)
	if err != nil {
		return &domain.Document{}, err
	}
	transaction := &domain.Transaction{WalletId: WalletOriginId, DocumentID: document.ID, Amount: -amount, Description: "براشت از حساب"}
	s.TransactionRepository.Create(transaction)
	transactionWage := &domain.Transaction{WalletId: WalletOriginId, DocumentID: document.ID, Amount: -100, Description: "براشت از حساب"}
	s.TransactionRepository.Create(transactionWage)
	transactionIntro := &domain.Transaction{WalletId: s.walletIdMaster, DocumentID: document.ID, Amount: amount, Description: "واریز"}
	s.TransactionRepository.Create(transactionIntro)
	transactionInWage := &domain.Transaction{WalletId: s.walletIdMaster, DocumentID: document.ID, Amount: 100, Description: "واریز"}
	s.TransactionRepository.Create(transactionInWage)
	transactionTo := &domain.Transaction{WalletId: walletDestinationMaster.ID, DocumentID: document.ID, Amount: -ConvertedAmount, Description: "واریز"}
	s.TransactionRepository.Create(transactionTo)
	transactionChangeTo := &domain.Transaction{WalletId: WalletDestination.ID, DocumentID: document.ID, Amount: ConvertedAmount, Description: "واریز"}
	s.TransactionRepository.Create(transactionChangeTo)

	s.dbTransaction.Commit(tx)
	return document, nil
}

func GenerateTrackId() string {
	s, _ := sqids.New()
	id, _ := s.Encode([]uint64{1, 2, 3}) // "86Rf07"
	return id
}
