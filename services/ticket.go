package services

import (
	"context"
	"crypto/rand"
	"math/big"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/repository"
	"github.com/google/uuid"
)

type TicketService interface {
	CreateTicket(ctx context.Context, ticketDTO dto.TicketCreateDTO, userID uuid.UUID) (entities.Ticket, error)
	GetAllTicketMovie(ctx context.Context, ticket dto.GetAllTicketMovie, jam string, studio string) ([]int, error)
	GetTicketUser(ctx context.Context, userID uuid.UUID) ([]map[string]interface{}, error)
}

type ticketService struct {
	ticketRepository repository.TicketRepository
}

func NewTicketService(ticketRepository repository.TicketRepository) TicketService {
	return &ticketService{
		ticketRepository: ticketRepository,
	}
}

func (ts *ticketService) CreateTicket(ctx context.Context, ticketDTO dto.TicketCreateDTO, userID uuid.UUID) (entities.Ticket, error) {
	var ticket entities.Ticket

	amountTicket := ticketDTO.Nomor
	ticket.MovieID = ticketDTO.MovieID
	ticket.UserID = userID
	ticket.Jam = ticketDTO.Jam
	ticket.Studio = ticketDTO.Studio
	
	kodeTransaksi, err := generateRandom(4)
	if err != nil {
		return entities.Ticket{}, err
	}
	
	ticket.KodeTransaksi = kodeTransaksi.Uint64()

	return ts.ticketRepository.CreateTicket(ctx, ticket, amountTicket)
}

func (ts *ticketService) GetAllTicketMovie(ctx context.Context, ticket dto.GetAllTicketMovie, jam string, studio string) ([]int, error) {
	var ticketEntity entities.Ticket
	ticketEntity.MovieID = ticket.MovieID

	return ts.ticketRepository.GetAllTicketMovie(ctx, ticketEntity, jam, studio)
}

func (ts *ticketService) GetTicketUser(ctx context.Context, userID uuid.UUID) ([]map[string]interface{}, error) {
	return ts.ticketRepository.GetTicketUser(ctx, userID)
}

func generateRandom(numBytes int) (*big.Int, error) {
	value := make([]byte, numBytes)
	_, err := rand.Reader.Read(value)
	if err != nil {
		return nil, err
	}

	for {
		if value[0] != 0 {
			break
		}

		firstByte := value[:1]
		_, err := rand.Reader.Read(firstByte)
		if err != nil {
			return nil, err
		}
	}

	return (&big.Int{}).SetBytes(value), nil
}