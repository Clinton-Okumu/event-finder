package store

import (
	"backend/internal/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

var ErrEventNotFound = errors.New("event not found")

type EventStore interface {
	CreateEvent(ctx context.Context, event *models.Event) error
	UpdateEvent(ctx context.Context, event *models.Event) error
	GetEvent(ctx context.Context, id uint) (*models.Event, error)
	GetEvents(ctx context.Context, limit, offset int, orderBy string) ([]models.Event, error)
	DeleteEvent(ctx context.Context, id uint) error
}

type eventStore struct {
	db *gorm.DB
}

func NewEventStore(db *gorm.DB) EventStore {
	return &eventStore{db}
}

func (es *eventStore) CreateEvent(ctx context.Context, event *models.Event) error {
	return es.db.WithContext(ctx).Create(event).Error
}

func (es *eventStore) UpdateEvent(ctx context.Context, event *models.Event) error {
	return es.db.WithContext(ctx).Save(event).Error
}

func (es *eventStore) GetEvent(ctx context.Context, id uint) (*models.Event, error) {
	var event models.Event
	err := es.db.WithContext(ctx).First(&event, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrEventNotFound
	}
	return &event, err
}

func (es *eventStore) GetEvents(ctx context.Context, limit, offset int, orderBy string) ([]models.Event, error) {
	var events []models.Event

	q := es.db.WithContext(ctx)
	if orderBy != "" {
		q = q.Order(orderBy)
	}
	if limit > 0 {
		q = q.Limit(limit).Offset(offset)
	}

	err := q.Find(&events).Error
	return events, err
}

func (es *eventStore) DeleteEvent(ctx context.Context, id uint) error {
	result := es.db.WithContext(ctx).Delete(&models.Event{}, id)
	if result.RowsAffected == 0 {
		return ErrEventNotFound
	}
	return result.Error
}
