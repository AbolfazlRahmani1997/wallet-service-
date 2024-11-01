package initialize

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"payment/core/domain"
)

func InitializeDatabase() *gorm.DB {
	var dns string
	if os.Getenv("ENV") == "production" {
		dns = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASS"))
	} else {
		fmt.Print("test")
		dns = "host=localhost user=niclub password=niclub dbname=payments port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	}

	db, _ := gorm.Open(postgres.Open(dns), &gorm.Config{})
	return db
}

func Migration(db *gorm.DB) {
	err := db.AutoMigrate(&domain.Rate{}, &domain.Account{}, &domain.Transaction{}, &domain.Wallet{}, &domain.Document{})
	if err != nil {
		fmt.Print(err)
		return
	}
	var Account domain.Account
	var WalletIr domain.Wallet
	var WalletGold domain.Wallet
	db.Where(&domain.Account{Id: 1, PhoneNumber: "09107879978", NationalCode: "0020747489"}).FirstOrCreate(&Account)
	db.Where(&domain.Wallet{AccountId: Account.Id, Currency: domain.IRR, WalletType: domain.Master}).FirstOrCreate(&WalletIr)
	db.Where(&domain.Wallet{AccountId: Account.Id, Currency: domain.Gold, WalletType: domain.Master}).FirstOrCreate(&WalletGold)
}
