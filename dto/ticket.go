package dto

import "github.com/google/uuid"

type TicketCreateDTO struct {
	Amount     int      `json:"amount" binding:"required"`
	Nomor      []uint64 `json:"nomor" binding:"required"`
	Jam 			string   `json:"jam" binding:"required"`
	Studio 		string   `json:"studio" binding:"required"`

	MovieID uuid.UUID `gorm:"foreignKey" json:"movie_id"`
}

type GetAllTicketMovie struct {
	MovieID uuid.UUID `gorm:"foreignKey" json:"movie_id"`
}

type SendScheduleMovie struct {
	Jam 			string   `json:"jam" binding:"required"`
	Studio 		string   `json:"studio" binding:"required"`
}