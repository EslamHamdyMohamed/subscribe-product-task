package test

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"transfer_money/app/exchange/requests"
	"transfer_money/database"
	"transfer_money/database/models/transfer"
)

var err error

func connectToSqlite() {
	database.DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		fmt.Println("Error ", err.Error())
	}
}

func init() {
	os.Truncate("gorm.db", 0)
	connectToSqlite()
	database.DB.AutoMigrate(&transfer.Account{})
	mockAccounts()
}

func setRequest(from, to string, amount float64) requests.TransferRequest {
	return requests.TransferRequest{
		From:   from,
		To:     to,
		Amount: amount,
	}
}

func mockAccounts() {
	accounts := []transfer.Account{
		{
			ID:      "1",
			Name:    "A",
			Balance: 100,
		},
		{
			ID:      "2",
			Name:    "B",
			Balance: 50,
		},
	}
	database.DB.Create(&accounts)
}
