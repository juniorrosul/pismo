package mysqlconnection

import (
	"github.com/juniorrosul/pismo/transaction"
)

// Transaction struct
type Transaction struct{}

var transactionTable = "transactions"

// Initialize implementation
func (t *Transaction) Initialize() error {
	checkTransactionTable()
	return nil
}

func checkTransactionTable() {
	db = connect()
	defer db.Close()

	if db.Migrator().HasTable(transactionTable) == false {
		db.Table(transactionTable).AutoMigrate(&transaction.Model{})
	}
}

// Store implementation
func (t *Transaction) Store(transaction *transaction.Model) error {
	db = connect()
	defer db.Close()

	db.Table(transactionTable).Create(transaction)
	return nil
}
