package models

import (
	"time"

	"gorm.io/gorm"
)

type PaymentStatus string

const (
	PaymentPending   PaymentStatus = "pending"
	PaymentCompleted PaymentStatus = "completed"
	PaymentFailed    PaymentStatus = "failed"
	PaymentRefunded  PaymentStatus = "refunded"
)

type Payment struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	BookingID     uint           `gorm:"not null" json:"booking_id"`
	Booking       *Booking       `gorm:"foreignKey:BookingID" json:"booking,omitempty"`
	Method        string         `gorm:"size:50;not null;default:'card'" json:"method"`
	Amount        float64        `gorm:"type:numeric(10,2);not null" json:"amount"`
	Status        PaymentStatus  `gorm:"type:enum('pending','completed','failed','refunded');default:'pending';not null" json:"status"`
	TransactionID *string        `gorm:"size:255;unique" json:"transaction_id,omitempty"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}
