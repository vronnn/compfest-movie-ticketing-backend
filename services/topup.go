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

type TopupService interface {
	CreateTopup(ctx context.Context, topup dto.TopupCreateDTO, userID uuid.UUID) (entities.Topup, error)
	GetAllTopupUser(ctx context.Context, userID uuid.UUID) ([]entities.Topup, error)
	GetTopupByID(ctx context.Context, topupID uuid.UUID) (entities.Topup, error)
}

type topupService struct {
	topupRepository repository.TopupRepository
}

func NewTopupService(tr repository.TopupRepository) TopupService {
	return &topupService{
		topupRepository: tr,
	}
}

func (ts *topupService) CreateTopup(ctx context.Context, topupDTO dto.TopupCreateDTO, userID uuid.UUID) (entities.Topup, error) {
	topup := entities.Topup{}
	err := smapping.FillStruct(&topup, smapping.MapFields(&topupDTO))
	if err != nil {
		return entities.Topup{}, err
	}

	topup.UserID = userID
	topup.TanggalTransaksi = time.Now()
	return ts.topupRepository.CreateTopup(ctx, topup, int(topupDTO.BankID))
}

func (ts *topupService) GetAllTopupUser(ctx context.Context, userID uuid.UUID) ([]entities.Topup, error) {
	return ts.topupRepository.GetAllTopupUser(ctx, userID)
}

func (ts *topupService) GetTopupByID(ctx context.Context, topupID uuid.UUID) (entities.Topup, error) {
	return ts.topupRepository.GetTopupByID(ctx, topupID)
}