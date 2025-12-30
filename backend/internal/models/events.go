package models

import (
	"time"

	"gorm.io/gorm"
)

type EventStatus string

const (
	EventPending   EventStatus = "pending"
	EventConfirmed EventStatus = "confirmed"
	EventCanceled  EventStatus = "canceled"
)

type Event struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	Title            string         `gorm:"not null" json:"title"`
	Description      string         `gorm:"type:text;not null" json:"description"`
	Location         string         `gorm:"not null" json:"location"`
	StartTime        time.Time      `gorm:"not null" json:"start_time"`
	EndTime          time.Time      `gorm:"not null" json:"end_time"`
	CreatedByID      uint           `gorm:"not null" json:"created_by"`
	Creator          *User          `gorm:"foreignKey:CreatedByID" json:"creator,omitempty"`
	Status           EventStatus    `gorm:"type:varchar(20);default:'pending';not null" json:"status"`
	ImageURL         string         `gorm:"size:500;not null" json:"image_url"`
	TotalTickets     int            `gorm:"not null" json:"total_tickets"`
	TicketsRemaining int            `gorm:"not null" json:"tickets_remaining"`
	CategoryID       uint           `gorm:"not null" json:"category_id"`
	Category         *Category      `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Price            float64        `gorm:"type:numeric(10,2);default:0" json:"price"`
	CreatedAt        time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}
