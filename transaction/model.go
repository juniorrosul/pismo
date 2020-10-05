package transaction

import "time"

// Model struct
type Model struct {
	ID              int       `json:"id" gorm:"primaryKey"`
	AccountID       int       `json:"account_id,string" gorm:"account_id"`
	OperationTypeID int       `json:"operation_type_id,string" gorm:"operation_type_id"`
	Amount          float64   `json:"amount" gorm:"amount"`
	EventDate       time.Time `gorm:"event_date"`
}
