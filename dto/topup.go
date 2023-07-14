package dto

import (
	"github.com/google/uuid"
)

type TopupCreateDTO struct {
	ID     uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Jumlah float64   `gorm:"type:float" json:"jumlah"`
	BankID int       `gorm:"type:int" json:"bank_id"`
}
