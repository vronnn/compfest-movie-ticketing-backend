package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketRepository interface {
	CreateTicket(ctx context.Context, ticket entities.Ticket, amountTicket []uint64) (entities.Ticket, error)
	GetAllTicketMovie(ctx context.Context, ticket entities.Ticket, jam string, studio string) ([]int, error)
	GetTicketUser(ctx context.Context, userID uuid.UUID) ([]map[string]interface{}, error)
}

type ticketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &ticketRepository{
		db: db,
	}
}

func (tr *ticketRepository) CreateTicket(ctx context.Context, ticket entities.Ticket, amountTicket []uint64) (entities.Ticket, error) {
	var ticketEntity entities.Ticket
	var movieEntity entities.Movies
	var pembayaranEntity entities.Pembayaran
	var user entities.User

	if err := tr.db.Where("id = ?", ticket.UserID).First(&user).Error; err != nil {
		return entities.Ticket{}, err
	}

	if err := tr.db.Where("id = ?", ticket.MovieID).First(&movieEntity).Error; err != nil {
		return entities.Ticket{}, err
	}

	// Check Umur User < Age Movie
	// if user.Age < movieEntity.AgeRating {
	// 	return entities.Ticket{}, errors.New("umur tidak memenuhi syarat")
	// }

	amount := len(amountTicket)
	totalHarga := movieEntity.TicketPrice * float64(amount)
	pembayaranEntity.ListBankID = 1
	pembayaranEntity.Harga = totalHarga
	pembayaranEntity.UserID = ticket.UserID

	if user.Saldo < totalHarga {
		return entities.Ticket{}, errors.New("saldo tidak cukup")
	}
	
	user.Saldo -= totalHarga
	tr.db.Save(&user)

	if err := tr.db.Create(&pembayaranEntity).Error; err != nil {
		fmt.Print(err)
		return entities.Ticket{}, err
	}
	
	for _, nomor := range amountTicket {
		ticketID := uuid.New()
		ticketEntity.ID = ticketID

		var existingTicket entities.Ticket
		check := tr.db.Where("nomor = ? AND movie_id = ?", nomor, ticket.MovieID).First(&existingTicket)

		if check.RowsAffected > 0 {
			return entities.Ticket{}, errors.New("nomor ticket sudah di booking")
		}

		ticketEntity.Nomor = uint64(nomor)
		ticketEntity.UserID = ticket.UserID
		ticketEntity.MovieID = ticket.MovieID
		ticketEntity.KodeTransaksi = ticket.KodeTransaksi
		ticketEntity.Jam = ticket.Jam
		ticketEntity.Studio = ticket.Studio

		booked := tr.db.Create(&ticketEntity)
		if booked.Error != nil {
			return entities.Ticket{}, booked.Error
		}
	}
	return ticketEntity, nil
}

func (tr *ticketRepository) GetAllTicketMovie(ctx context.Context, ticket entities.Ticket, jam string, studio string) ([]int, error) {
	var ticketEntity []entities.Ticket
	var availableTicket [65]int

	for i := 1; i <= 64; i++ {
		availableTicket[i] = 1
	}

	if err := tr.db.Where("movie_id = ? AND jam = ? AND studio = ?", ticket.MovieID, jam, studio).Find(&ticketEntity).Error; err != nil {
		return []int{}, err
	}
	
	for _, ticket := range ticketEntity {
		availableTicket[ticket.Nomor] = 0
	}

	return availableTicket[1:], nil
}

func (tr *ticketRepository) GetTicketUser(ctx context.Context, userID uuid.UUID) ([]map[string]interface{}, error) {
	ticketMap := make(map[int][]int)
	var tickets []entities.Ticket

	if err := tr.db.Where("user_id = ?", userID).Find(&tickets).Error; err != nil {
		return nil, err
	}

	for _, ticket := range tickets {
		ticketMap[int(ticket.KodeTransaksi)] = append(ticketMap[int(ticket.KodeTransaksi)], int(ticket.Nomor))
	}

	var result []map[string]interface{}
	for kodeTransaksi, nomors := range ticketMap {
		ticket := make(map[string]interface{})
		ticket["nomor"] = nomors
		ticket["kode_transaksi"] = kodeTransaksi
		ticket["jam"] = tickets[0].Jam
		ticket["studio"] = tickets[0].Studio
		ticket["user_id"] = tickets[0].UserID
		ticket["movie_id"] = tickets[0].MovieID
		ticket["created_at"] = tickets[0].CreatedAt
		ticket["updated_at"] = tickets[0].UpdatedAt

		result = append(result, ticket)
	}

	return result, nil
}