package entities

import (
	"time"

	"github.com/google/uuid"
)

type WithDrawal struct {
	ID               uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	JumlahPenarikan  float64   `gorm:"type:float" json:"jumlah_penarikan"`
	TanggalTransaksi time.Time `gorm:"timestamp with time zone" json:"tanggal_transaksi"`
	BankName         string    `gorm:"type:varchar(100)" json:"bank_name"`
	NoRek		  string    `gorm:"type:varchar(100)" json:"no_rek"`

	UserID uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User   User      `gorm:"foreignKey:UserID" json:"-"`

	Timestamp
}
