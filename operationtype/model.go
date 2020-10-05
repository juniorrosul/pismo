package operationtype

// Model struct
type Model struct {
	ID          int    `json:"id" gorm:"primarykey"`
	Description string `json:"description" gorm:"description"`
}
