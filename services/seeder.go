package services

import (
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/repository"
)

type SeederService interface {
	GetAllBank() ([]entities.ListBank, error)
	GetBankByID(id uint) (entities.ListBank, error)
}

type seederService struct {
	seederRepository repository.SeederRepository
}

func NewSeederService(sr repository.SeederRepository) SeederService {
	return &seederService{
		seederRepository: sr,
	}
}

func (ss *seederService) GetAllBank() ([]entities.ListBank, error) {
	return ss.seederRepository.GetAllBank()
}

func (ss *seederService) GetBankByID(id uint) (entities.ListBank, error) {
	return ss.seederRepository.GetBankByID(id)
}