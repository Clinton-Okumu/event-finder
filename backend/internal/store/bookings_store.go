package store

import (
	"backend/internal/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

var ErrBookingNotFound = errors.New("booking not found")

type BookingStore interface {
	CreateBooking(ctx context.Context, booking *models.Booking) error
	UpdateBooking(ctx context.Context, booking *models.Booking) error
	GetBooking(ctx context.Context, id uint) (*models.Booking, error)
	GetBookings(ctx context.Context, limit, offset int, orderBy string) ([]models.Booking, error)
	GetBookingsByUserID(ctx context.Context, userID uint) ([]models.Booking, error)
	DeleteBooking(ctx context.Context, id uint) error
}

type bookingStore struct {
	db *gorm.DB
}

func NewBookingStore(db *gorm.DB) BookingStore {
	return &bookingStore{db}
}

func (bs *bookingStore) CreateBooking(ctx context.Context, booking *models.Booking) error {
	return bs.db.WithContext(ctx).Create(booking).Error
}

func (bs *bookingStore) UpdateBooking(ctx context.Context, booking *models.Booking) error {
	return bs.db.WithContext(ctx).Save(booking).Error
}

func (bs *bookingStore) GetBooking(ctx context.Context, id uint) (*models.Booking, error) {
	var booking models.Booking
	err := bs.db.WithContext(ctx).
		Preload("Event").
		Preload("BookingItems").
		Preload("BookingItems.TicketType").
		First(&booking, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrBookingNotFound
	}
	return &booking, err
}

func (bs *bookingStore) GetBookings(ctx context.Context, limit, offset int, orderBy string) ([]models.Booking, error) {
	var bookings []models.Booking
	q := bs.db.WithContext(ctx)
	if orderBy != "" {
		q = q.Order(orderBy)
	}
	if limit > 0 {
		q = q.Limit(limit).Offset(offset)
	}
	err := q.Find(&bookings).Error
	return bookings, err
}

func (bs *bookingStore) DeleteBooking(ctx context.Context, id uint) error {
	result := bs.db.WithContext(ctx).Delete(&models.Booking{}, id)
	if result.RowsAffected == 0 {
		return ErrBookingNotFound
	}
	return result.Error
}

func (bs *bookingStore) GetBookingsByUserID(ctx context.Context, userID uint) ([]models.Booking, error) {
	var bookings []models.Booking
	err := bs.db.WithContext(ctx).
		Preload("Event").
		Preload("BookingItems").
		Preload("BookingItems.TicketType").
		Where("user_id = ?", userID).
		Find(&bookings).Error
	return bookings, err
}
