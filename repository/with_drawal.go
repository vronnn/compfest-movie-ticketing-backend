package repository

import (
	"context"
	"errors"

	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WithDrawalRepository interface {
	CreateWithDrawal(ctx context.Context, withDrawal entities.WithDrawal, bankID int) (entities.WithDrawal, error)
	GetAllWithDrawalUser(ctx context.Context, userID uuid.UUID) ([]entities.WithDrawal, error)
}

type withDrawalRepository struct {
	db *gorm.DB
}

func NewWithDrawalRepository(db *gorm.DB) WithDrawalRepository {
	return &withDrawalRepository{
		db: db,
	}
}

func (wr *withDrawalRepository) CreateWithDrawal(ctx context.Context, withDrawal entities.WithDrawal, bankID int) (entities.WithDrawal, error) {
	var user entities.User
	var bank entities.ListBank

	if err := wr.db.Where("id = ?", bankID).First(&bank).Error; err != nil {
		return entities.WithDrawal{}, err
	}

	if err := wr.db.Where("id = ?", withDrawal.UserID).First(&user).Error; err != nil {
		return entities.WithDrawal{}, err
	}

	if user.Saldo < withDrawal.JumlahPenarikan {
		return entities.WithDrawal{}, errors.New("saldo tidak cukup")
	}

	user.Saldo -= withDrawal.JumlahPenarikan
	wr.db.Save(&user)
	withDrawal.BankName = bank.Name

	if err := wr.db.Create(&withDrawal).Error; err != nil {
		return entities.WithDrawal{}, err
	}

	return withDrawal, nil
}

func (wr *withDrawalRepository) GetAllWithDrawalUser(ctx context.Context, userID uuid.UUID) ([]entities.WithDrawal, error) {
	var withDrawal []entities.WithDrawal
	if err := wr.db.Where("user_id = ?", userID).Find(&withDrawal).Error; err != nil {
		return nil, err
	}
	return withDrawal, nil
}