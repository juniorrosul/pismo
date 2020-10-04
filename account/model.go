package account

import "gorm.io/gorm"

// Model struct
type Model struct {
	gorm.Model
	ID             int    `json:"id"`
	DocumentNumber string `json:"document_number"`
}
