package models

import (
	"time"

	"gorm.io/gorm"
)

type BookingStatus string

const (
	BookingPending   BookingStatus = "pending"
	BookingConfirmed BookingStatus = "confirmed"
	BookingCanceled  BookingStatus = "canceled"
)

type Booking struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	User      *User          `gorm:"foreignKey:UserID" json:"user,omitempty"`
	EventID   uint           `gorm:"not null" json:"event_id"`
	Event     *Event         `gorm:"foreignKey:EventID" json:"event,omitempty"`
	Status    BookingStatus  `gorm:"type:text;check:status IN ('pending', 'confirmed', 'canceled');default:'pending';not null" json:"status"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
