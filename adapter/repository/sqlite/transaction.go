package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"madsonjr.com/pismo/transaction"
)

type Transaction struct{}

var transactionTable = "transactions"

func checkTransactionTable() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	if db.Migrator().HasTable(transactionTable) == false {
		db.Table(transactionTable).AutoMigrate(&transaction.Model{})
	}
}

// Store implementation
func (t *Transaction) Store(transaction *transaction.Model) error {
	checkTransactionTable()
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.Table(transactionTable).Create(transaction)
	return nil
}
