package models

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	Hash   []byte    `json:"-"`
	UserID uint      `gorm:"index" json:"-"`
	Expiry time.Time `gorm:"index" json:"expiry"`
	Scope  string    `gorm:"index" json:"-"`
}
