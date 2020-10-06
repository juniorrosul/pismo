package repository

import (
	"github.com/juniorrosul/pismo/transaction"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

// Initialize implementation
func (t *Transaction) Initialize() error {
	checkTransactionTable()
	return nil
}
