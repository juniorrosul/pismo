package mysqlconnection

import (
	"fmt"
	"time"

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
	checkOperationTypeTable()

	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if db.Migrator().HasTable(transactionTable) == false {
		db.AutoMigrate(&transaction.Model{})
	}
}

// Store implementation
func (t *Transaction) Store(transaction *transaction.Model) error {
	checkTransactionTable()
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	transaction.EventDate = time.Now()

	db.Create(transaction)
	return nil
}
