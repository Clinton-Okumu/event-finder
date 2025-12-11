package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique;not null" json:"username"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`
	Role     string `gorm:"type:varchar(20);default:user;not null" json:"role"`
	gorm.Model
}

func (u *User) SetPassword(raw string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

func (u *User) CheckPassword(raw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(raw))
	return err == nil
}

func (u *User) IsAdmin() bool {
	return u.Role == "admin"
}
