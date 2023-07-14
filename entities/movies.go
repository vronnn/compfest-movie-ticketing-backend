package entities

import (
	"time"

	"github.com/google/uuid"
)

type Movies struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Title       string     `gorm:"type:varchar(255)" json:"title"`
	Description string     `gorm:"type:text" json:"description"`
	ReleaseDate time.Time  `gorm:"timestamp with time zone" json:"release_date"`
	PosterUrl   string     `gorm:"type:varchar(255)" json:"poster_url"`
	AgeRating   int        `gorm:"type:int" json:"age_rating"`
	TicketPrice float64    `gorm:"type:float" json:"ticket_price"`

	Ticket      []Ticket   `gorm:"foreignKey:MovieID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Timestamp
}