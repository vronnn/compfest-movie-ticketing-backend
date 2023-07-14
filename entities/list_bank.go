package entities

type ListBank struct {
	ID   uint   `gorm:"type:int;primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(100)" json:"name"`

	// Pembayaran Pembayaran `gorm:"foreignKey:ListBankID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
