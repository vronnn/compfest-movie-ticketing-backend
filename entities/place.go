package entities

type Place struct {
	ID   uint   `gorm:"type:int;primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(100)" json:"name"`
}