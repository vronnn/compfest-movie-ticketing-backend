package entities

import "github.com/google/uuid"

type Ticket struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Nomor         uint64    `gorm:"type:int" json:"nomor"`
	KodeTransaksi uint64    `gorm:"type:int" json:"kode_transaksi"`
	Jam           string    `gorm:"type:varchar(255)" json:"jam"`
	Studio        string    `gorm:"type:varchar(255)" json:"studio"`

	UserID uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User   User      `gorm:"foreignKey:UserID" json:"-"`

	MovieID uuid.UUID `gorm:"type:uuid" json:"movie_id"`
	Movie   Movies    `gorm:"foreignKey:MovieID" json:"-"`

	Timestamp
}
