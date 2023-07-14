package repository

import (
	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

type SeederRepository interface {
	GetAllBank() ([]entities.ListBank, error)
	GetBankByID(id uint) (entities.ListBank, error)
}

type seederRepository struct {
	connection *gorm.DB
}

func NewSeederRepository(db *gorm.DB) SeederRepository {
	return &seederRepository{
		connection: db,
	}
}

func (sr *seederRepository) GetAllBank() ([]entities.ListBank, error) {
	var listBank []entities.ListBank
	err := sr.connection.Find(&listBank).Error
	if err != nil {
		return listBank, err
	}
	return listBank, nil
}	

func (sr *seederRepository) GetBankByID(id uint) (entities.ListBank, error) {
	var bank entities.ListBank
	err := sr.connection.Where("id = ?", id).First(&bank).Error
	if err != nil {
		return bank, err
	}
	return bank, nil
}