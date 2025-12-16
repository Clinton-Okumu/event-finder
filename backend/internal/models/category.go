package models

import (
	"gorm.io/gorm"
)

type Category struct {
	Name        string `gorm:"not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
	gorm.Model
}
