package store

import (
	"backend/internal/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

var ErrEventTicketNotFound = errors.New("event ticket not found")

type EventTicketStore interface {
	CreateEventTicket(ctx context.Context, eventTicket *models.EventTicket) error
	UpdateEventTicket(ctx context.Context, eventTicket *models.EventTicket) error
	GetEventTicket(ctx context.Context, id uint) (*models.EventTicket, error)
	GetEventTickets(ctx context.Context, limit, offset int, orderBy string) ([]models.EventTicket, error)
	DeleteEventTicket(ctx context.Context, id uint) error
}

type eventTicketStore struct {
	db *gorm.DB
}

func NewEventTicketStore(db *gorm.DB) EventTicketStore {
	return &eventTicketStore{db}
}

func (es *eventTicketStore) CreateEventTicket(ctx context.Context, eventTicket *models.EventTicket) error {
	return es.db.WithContext(ctx).Create(eventTicket).Error
}

func (es *eventTicketStore) UpdateEventTicket(ctx context.Context, eventTicket *models.EventTicket) error {
	return es.db.WithContext(ctx).Save(eventTicket).Error
}

func (es *eventTicketStore) GetEventTicket(ctx context.Context, id uint) (*models.EventTicket, error) {
	var eventTicket models.EventTicket
	err := es.db.WithContext(ctx).First(&eventTicket, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrEventTicketNotFound
	}
	return &eventTicket, err
}

func (es *eventTicketStore) GetEventTickets(ctx context.Context, limit, offset int, orderBy string) ([]models.EventTicket, error) {
	var eventTickets []models.EventTicket

	q := es.db.WithContext(ctx)
	if orderBy != "" {
		q = q.Order(orderBy)
	}
	if limit > 0 {
		q = q.Limit(limit).Offset(offset)
	}

	err := q.Find(&eventTickets).Error
	return eventTickets, err
}

func (es *eventTicketStore) DeleteEventTicket(ctx context.Context, id uint) error {
	result := es.db.WithContext(ctx).Delete(&models.EventTicket{}, id)
	if result.RowsAffected == 0 {
		return ErrEventTicketNotFound
	}
	return result.Error
}
