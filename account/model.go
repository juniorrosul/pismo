package account

// Model struct
type Model struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	DocumentNumber string `json:"document_number" gorm:"document_number"`
}
