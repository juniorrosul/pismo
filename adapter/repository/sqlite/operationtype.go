package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"madsonjr.com/pismo/operationtype"
)

// OperationType struct
type OperationType struct{}

var operationTypeTable = "operation_types"

func checkOperationTypeTable() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	if db.Migrator().HasTable(operationTypeTable) == false {
		db.Table(operationTypeTable).AutoMigrate(&operationtype.Model{})
	}
}

// Find implementation
func (ot *OperationType) Find(id int) (*operationtype.Model, error) {
	checkOperationTypeTable()
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	var operationType operationtype.Model
	result := db.Table(operationTypeTable).First(&operationType, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &operationType, nil
}