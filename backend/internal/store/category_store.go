package store

import (
	"backend/internal/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

var ErrCategoryNotFound = errors.New("category not found")

type CategoryStore interface {
	CreateCategory(ctx context.Context, category *models.Category) error
	UpdateCategory(ctx context.Context, category *models.Category) error
	GetCategory(ctx context.Context, id uint) (*models.Category, error)
	GetCategories(ctx context.Context, limit, offset int, orderBy string) ([]models.Category, error)
	DeleteCategory(ctx context.Context, id uint) error
}

type categoryStore struct {
	db *gorm.DB
}

func NewCategoryStore(db *gorm.DB) CategoryStore {
	return &categoryStore{db}
}

func (cs *categoryStore) CreateCategory(ctx context.Context, category *models.Category) error {
	return cs.db.WithContext(ctx).Create(category).Error
}

func (cs *categoryStore) UpdateCategory(ctx context.Context, category *models.Category) error {
	return cs.db.WithContext(ctx).Save(category).Error
}

func (cs *categoryStore) GetCategory(ctx context.Context, id uint) (*models.Category, error) {
	var category models.Category
	err := cs.db.WithContext(ctx).First(&category, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrCategoryNotFound
	}
	return &category, err
}

func (cs *categoryStore) GetCategories(ctx context.Context, limit, offset int, orderBy string) ([]models.Category, error) {
	var categories []models.Category

	q := cs.db.WithContext(ctx)
	if orderBy != "" {
		q = q.Order(orderBy)
	}
	if limit > 0 {
		q = q.Limit(limit).Offset(offset)
	}

	err := q.Find(&categories).Error
	return categories, err
}

func (cs *categoryStore) DeleteCategory(ctx context.Context, id uint) error {
	result := cs.db.WithContext(ctx).Delete(&models.Category{}, id)
	if result.RowsAffected == 0 {
		return ErrCategoryNotFound
	}
	return result.Error
}
