package account

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by Model to `accounts`
func (Model) TableName() string {
	return "accounts"
}

// Model struct
type Model struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	DocumentNumber string `json:"document_number" gorm:"document_number"`
}
