package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TopupRepository interface {
	CreateTopup(ctx context.Context, topup entities.Topup, bankID int) (entities.Topup, error)
	GetAllTopupUser(ctx context.Context, userID uuid.UUID) ([]entities.Topup, error)
	GetTopupByID(ctx context.Context, topupID uuid.UUID) (entities.Topup, error)
}

type topupRepository struct {
	connection *gorm.DB
}

func NewTopupRepository(db *gorm.DB) TopupRepository {
	return &topupRepository{
		connection: db,
	}
}

func (tr *topupRepository) CreateTopup(ctx context.Context, topup entities.Topup, bankID int) (entities.Topup, error) {
	var user entities.User
	var bank entities.ListBank

	if err := tr.connection.Where("id", bankID).First(&bank).Error; err != nil {
		return entities.Topup{}, err
	}

	if err := tr.connection.Where("id", topup.UserID).First(&user).Error; err != nil {
		return entities.Topup{}, err
	}

	user.Saldo += topup.Jumlah
	tr.connection.Save(&user)
	topup.BankName = bank.Name

	if err := tr.connection.Create(&topup).Error; err != nil {
		return entities.Topup{}, err
	}
	return topup, nil
}

func (tr *topupRepository) GetAllTopupUser(ctx context.Context, userID uuid.UUID) ([]entities.Topup, error) {
	var topup []entities.Topup
	if err := tr.connection.Where("user_id", userID).Find(&topup).Error; err != nil {
		return nil, err
	}
	return topup, nil
}

func (tr *topupRepository) GetTopupByID(ctx context.Context, topupID uuid.UUID) (entities.Topup, error) {
	var topup entities.Topup
	if err := tr.connection.Where("id = ?", topupID).Take(&topup).Error; err != nil {
		return entities.Topup{}, err
	}
	return topup, nil
}