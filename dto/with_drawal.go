package dto

type CreateWithDrawalDTO struct {
	JumlahPenarikan float64 `gorm:"type:float" json:"jumlah_penarikan"`
	NoRek 				 string  `gorm:"type:varchar(100)" json:"no_rek"`
	BankID          uint    `gorm:"type:uint" json:"bank_id"`
}