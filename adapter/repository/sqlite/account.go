package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"madsonjr.com/pismo/account"
)

type Account struct{}

var tableName = "accounts"

func checkTable() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	if db.Migrator().HasTable(tableName) == false {
		db.Migrator().CreateTable(&account.Model{})
	}
}

func (a *Account) Find(id int) (*account.Model, error) {
	checkTable()
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	var account account.Model
	db.Table(tableName).First(&account, id)

	return &account, nil
}

func (a *Account) Store(account *account.Model) error {
	checkTable()
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.Table("accounts").Create(account)
	return nil
}
