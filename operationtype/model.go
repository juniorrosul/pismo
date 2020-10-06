package operationtype

// Model struct
type Model struct {
	ID          uint   `json:"id" gorm:"primarykey"`
	Description string `json:"description" gorm:"description"`
}
