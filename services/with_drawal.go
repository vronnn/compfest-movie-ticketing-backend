package services

import (
	"context"
	"time"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/repository"
	"github.com/google/uuid"
	"github.com/mashingan/smapping"
)

type WithDrawalService interface {
	CreateWithDrawal(ctx context.Context, withDrawalDTO dto.CreateWithDrawalDTO, userID uuid.UUID) (entities.WithDrawal, error)
	GetAllWithDrawalUser(ctx context.Context, userID uuid.UUID) ([]entities.WithDrawal, error)
}

type withDrawalService struct {
	withDrawalRepository repository.WithDrawalRepository
}

func NewWithDrawalService(wr repository.WithDrawalRepository) WithDrawalService {
	return &withDrawalService{
		withDrawalRepository: wr,
	}
}

func (ws *withDrawalService) CreateWithDrawal(ctx context.Context, withDrawalDTO dto.CreateWithDrawalDTO, userID uuid.UUID) (entities.WithDrawal, error) {
	withDrawal := entities.WithDrawal{}
	err := smapping.FillStruct(&withDrawal, smapping.MapFields(&withDrawalDTO))
	if err != nil {
		return entities.WithDrawal{}, err
	}

	withDrawal.UserID = userID
	withDrawal.TanggalTransaksi = time.Now()
	return ws.withDrawalRepository.CreateWithDrawal(ctx, withDrawal, int(withDrawalDTO.BankID))
}

func (ws *withDrawalService) GetAllWithDrawalUser(ctx context.Context, userID uuid.UUID) ([]entities.WithDrawal, error) {
	return ws.withDrawalRepository.GetAllWithDrawalUser(ctx, userID)
}