package mysqlconnection

import (
	"github.com/juniorrosul/pismo/account"
)

const (
	dbName     = "pismo"
	dbHost     = "localhost"
	dbUser     = "root"
	dbPassword = "root"
	dbPort     = "3306"
)

// Account struct
type Account struct{}

var accountTable = "accounts"

func checkAccountTable() {
	db = connect()
	defer db.Close()

	if db.Migrator().HasTable(accountTable) == false {
		db.Table(accountTable).AutoMigrate(&account.Model{})
	}
}

// Initialize implementation
func (a *Account) Initialize() {
	checkAccountTable()
}

// Find implementation
func (a *Account) Find(id int) (*account.Model, error) {
	checkAccountTable()
	db = connect()
	defer db.Close()

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
	db = connect()
	defer db.Close()

	db.Table(accountTable).Create(account)
	return nil
}
