package transaction

import (
	"time"

	"github.com/juniorrosul/pismo/account"
	"github.com/juniorrosul/pismo/operationtype"
)

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by Model to `transactions`
func (Model) TableName() string {
	return "transactions"
}

// Model struct
type Model struct {
	ID              uint                `json:"id" gorm:"primaryKey"`
	AccountID       uint                `json:"account_id,string" gorm:"account_id"`
	Account         account.Model       `gorm:"references:id"`
	OperationTypeID uint                `json:"operation_type_id,string" gorm:"operation_type_id"`
	OperationType   operationtype.Model `gorm:"references:id"`
	Amount          float64             `json:"amount" gorm:"amount"`
	EventDate       time.Time           `gorm:"event_date,default:CURRENT_TIMESTAMP"`
}
