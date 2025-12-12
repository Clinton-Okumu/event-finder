package store

import (
	"backend/internal/models"
	"backend/internal/tokens"
	"context"
	"time"

	"gorm.io/gorm"
)

type TokenStore interface {
	CreateNewToken(ctx context.Context, userID uint, ttl time.Duration, scope string) (*models.Token, string, error)
	DeleteAllTokensForUser(ctx context.Context, userID uint, scope string) error
	GetUserByToken(ctx context.Context, tokenPlaintext string, scope string) (*models.User, error)
}

type tokenStore struct {
	db *gorm.DB
}

func NewTokenStore(db *gorm.DB) TokenStore {
	return &tokenStore{db}
}

func (ts *tokenStore) CreateNewToken(ctx context.Context, userID uint, ttl time.Duration, scope string) (*models.Token, string, error) {
	token, plaintext, err := tokens.GenerateToken(userID, ttl, scope)
	if err != nil {
		return nil, "", err
	}

	if err := ts.db.WithContext(ctx).Create(token).Error; err != nil {
		return nil, "", err
	}

	return token, plaintext, nil
}

func (ts *tokenStore) DeleteAllTokensForUser(ctx context.Context, userID uint, scope string) error {
	return ts.db.WithContext(ctx).
		Where("user_id = ? AND scope = ?", userID, scope).
		Delete(&models.Token{}).Error
}

func (ts *tokenStore) GetUserByToken(ctx context.Context, tokenPlaintext string, scope string) (*models.User, error) {
	hash := tokens.HashToken(tokenPlaintext)

	var token models.Token
	if err := ts.db.WithContext(ctx).
		Where("hash = ? AND scope = ? AND expiry > ?", hash, scope, time.Now()).
		First(&token).Error; err != nil {
		return nil, err
	}

	var user models.User
	if err := ts.db.WithContext(ctx).
		First(&user, token.UserID).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
