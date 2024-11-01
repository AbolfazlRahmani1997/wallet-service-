package main

import (
	"payment/adapters/handler/http"
	"payment/adapters/mysql"
	"payment/api"
	"payment/core/service"
	"payment/initialize"
)

func main() {
	db := initialize.InitializeDatabase()
	initialize.Migration(db)

	walletRepository := mysql.NewWalletRepository(db)
	transactionRepository := mysql.NewTransactionRepository(db)
	documentRepository := mysql.NewDocumentRepository(db)
	dbTransaction := mysql.NewDBTransaction(db)
	WalletService := service.NewWalletService(dbTransaction, walletRepository, documentRepository, "1", transactionRepository)
	WalletHandler := http.NewWalletHandler(WalletService)

	api.InitRouter(WalletHandler)
	err := api.Start("127.0.0.1:9000")
	if err != nil {
		return
	}
}
