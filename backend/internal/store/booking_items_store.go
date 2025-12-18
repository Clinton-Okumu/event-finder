package store

import (
	"backend/internal/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

var ErrBookingItemNotFound = errors.New("booking item not found")

type BookingItemStore interface {
	CreateBookingItem(ctx context.Context, bookingItem *models.BookingItem) error
	UpdateBookingItem(ctx context.Context, bookingItem *models.BookingItem) error
	GetBookingItem(ctx context.Context, id uint) (*models.BookingItem, error)
	GetBookingItems(ctx context.Context, limit, offset int, orderBy string) ([]models.BookingItem, error)
	DeleteBookingItem(ctx context.Context, id uint) error
}

type bookingItemStore struct {
	db *gorm.DB
}

func NewBookingItemStore(db *gorm.DB) BookingItemStore {
	return &bookingItemStore{db}
}

func (bis *bookingItemStore) CreateBookingItem(ctx context.Context, bookingItem *models.BookingItem) error {
	return bis.db.WithContext(ctx).Create(bookingItem).Error
}

func (bis *bookingItemStore) UpdateBookingItem(ctx context.Context, bookingItem *models.BookingItem) error {
	return bis.db.WithContext(ctx).Save(bookingItem).Error
}

func (bis *bookingItemStore) GetBookingItem(ctx context.Context, id uint) (*models.BookingItem, error) {
	var bookingItem models.BookingItem
	err := bis.db.WithContext(ctx).First(&bookingItem, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrBookingItemNotFound
	}
	return &bookingItem, err
}

func (bis *bookingItemStore) GetBookingItems(ctx context.Context, limit, offset int, orderBy string) ([]models.BookingItem, error) {
	var bookingItems []models.BookingItem
	q := bis.db.WithContext(ctx)
	if orderBy != "" {
		q = q.Order(orderBy)
	}
	if limit > 0 {
		q = q.Limit(limit).Offset(offset)
	}
	err := q.Find(&bookingItems).Error
	return bookingItems, err
}

func (bis *bookingItemStore) DeleteBookingItem(ctx context.Context, id uint) error {
	result := bis.db.WithContext(ctx).Delete(&models.BookingItem{}, id)
	if result.RowsAffected == 0 {
		return ErrBookingItemNotFound
	}
	return result.Error
}
