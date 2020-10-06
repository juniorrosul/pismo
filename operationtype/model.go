package operationtype

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by Model to `operation_types`
func (Model) TableName() string {
	return "operation_types"
}

// Model struct
type Model struct {
	ID          uint   `json:"id" gorm:"primarykey"`
	Description string `json:"description" gorm:"description"`
}
