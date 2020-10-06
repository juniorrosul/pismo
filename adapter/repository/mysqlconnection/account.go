package mysqlconnection

import (
	"github.com/juniorrosul/pismo/account"
)

// Account struct
type Account struct{}

var accountTable = "accounts"

func checkAccountTable() {
	db := connect()

	if db.Migrator().HasTable(accountTable) == false {
		db.AutoMigrate(&account.Model{})
	}
}

// Initialize implementation
func (a *Account) Initialize() {
	checkAccountTable()
}

// Find implementation
func (a *Account) Find(id int) (*account.Model, error) {
	checkAccountTable()
	db := connect()

	var account account.Model
	result := db.Table(accountTable).First(&account, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &account, nil
}

// Store implementation
func (a *Account) Store(account *account.Model) error {
	checkAccountTable()
	db := connect()

	db.Table(accountTable).Create(account)
	return nil
}
