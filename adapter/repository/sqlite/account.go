package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"madsonjr.com/pismo/account"
)

// Account struct
type Account struct{}

var accountTable = "accounts"

func checkAccountTable() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	if db.Migrator().HasTable(accountTable) == false {
		db.Table(accountTable).AutoMigrate(&account.Model{})
	}
}

// Find implementation
func (a *Account) Find(id int) (*account.Model, error) {
	checkAccountTable()
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

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
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.Table(accountTable).Create(account)
	return nil
}
