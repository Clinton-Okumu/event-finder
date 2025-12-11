package models

import (
	"time"

	"gorm.io/gorm"
)

type BookingItem struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	BookingID       uint           `gorm:"not null" json:"booking_id"`
	Booking         *Booking       `gorm:"foreignKey:BookingID" json:"booking,omitempty"`
	TicketTypeID    uint           `gorm:"not null" json:"ticket_type_id"`
	TicketType      *EventTicket   `gorm:"foreignKey:TicketTypeID" json:"ticket_type,omitempty"`
	Quantity        int            `gorm:"not null" json:"quantity"`
	PriceAtPurchase float64        `gorm:"type:numeric(10,2);not null" json:"price_at_purchase"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}

func (bi *BookingItem) TotalPrice() float64 {
	return float64(bi.Quantity) * bi.PriceAtPurchase
}
