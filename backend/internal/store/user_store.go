package store

import (
	"backend/internal/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

var ErrUserNotFound = errors.New("user not found")

type UserStore interface {
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	DeleteUserByID(ctx context.Context, id uint) error
}

type userStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) UserStore {
	return &userStore{db}
}

func (us *userStore) CreateUser(ctx context.Context, user *models.User) error {
	return us.db.WithContext(ctx).Create(user).Error
}

func (us *userStore) UpdateUser(ctx context.Context, user *models.User) error {
	return us.db.WithContext(ctx).Save(user).Error
}

func (us *userStore) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := us.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}
	return &user, err
}

func (us *userStore) DeleteUserByID(ctx context.Context, id uint) error {
	result := us.db.WithContext(ctx).Delete(&models.User{}, id)
	if result.RowsAffected == 0 {
		return ErrUserNotFound
	}
	return result.Error
}
