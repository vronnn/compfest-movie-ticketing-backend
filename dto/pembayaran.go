package dto

type CreatePembayaranDTO struct {
	Harga float64 `gorm:"type:float" json:"harga"`

	BankID uint `gorm:"type:uint" json:"bank_id"`
}