package models

import (
	"time"

	"gorm.io/gorm"
)

type TicketType string

const (
	TicketEarlyBird TicketType = "Early Bird"
	TicketVIP       TicketType = "VIP"
	TicketRegular   TicketType = "Regular"
	TicketVVIP      TicketType = "VVIP"
)

type EventTicket struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	EventID     uint           `gorm:"not null" json:"event_id"`
	Event       *Event         `gorm:"foreignKey:EventID" json:"event,omitempty"`
	Name        TicketType     `gorm:"type:enum('Early Bird','VIP','Regular','VVIP');not null" json:"name"`
	Price       float64        `gorm:"type:numeric(10,2);not null" json:"price"`
	Quantity    int            `gorm:"not null" json:"quantity"`
	Description string         `gorm:"type:text" json:"description,omitempty"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
