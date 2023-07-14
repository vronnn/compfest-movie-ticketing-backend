package entities

import "github.com/google/uuid"

type Pembayaran struct {
	ID     uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Harga float64   `gorm:"type:float" json:"harga"`

	ListBankID uint     `gorm:"type:uint" json:"list_bank_id"`
	ListBank   ListBank `gorm:"foreignKey:ListBankID" json:"-"`

	UserID uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User   User      `gorm:"foreignKey:UserID" json:"user,omitempty"`

	Timestamp
}
