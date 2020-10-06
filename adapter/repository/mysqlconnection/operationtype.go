package mysqlconnection

import (
	"fmt"

	"github.com/juniorrosul/pismo/operationtype"
)

// OperationType struct
type OperationType struct{}

var operationTypeTable = "operation_types"

func checkOperationTypeTable() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if db.Migrator().HasTable(operationTypeTable) == false {
		db.AutoMigrate(&operationtype.Model{})
		operationTypes := []operationtype.Model{
			{
				ID:          1,
				Description: "COMPRA A VISTA",
			},
			{
				ID:          2,
				Description: "COMPRA PARCELADA",
			},
			{
				ID:          3,
				Description: "SAQUE",
			},
			{
				ID:          4,
				Description: "PAGAMENTO",
			},
		}
		db.Create(operationTypes)
	}
}

// Initialize database
func (ot *OperationType) Initialize() {
	checkOperationTypeTable()
}

// Find implementation
func (ot *OperationType) Find(id int) (*operationtype.Model, error) {
	checkOperationTypeTable()
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var operationType operationtype.Model
	result := db.First(&operationType, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &operationType, nil
}

// Store implementation
func (ot *OperationType) Store(operationType *operationtype.Model) error {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	db.Table(operationTypeTable).Create(operationType)
	return nil
}
