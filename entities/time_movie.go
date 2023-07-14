package entities

type TimeMovie struct {
	ID   uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Time string `gorm:"type:varchar(255)" json:"time"`
	Type int 	`gorm:"type:int" json:"type"`
}