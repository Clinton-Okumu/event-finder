package tokens

import (
	"backend/internal/models"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"time"

	"gorm.io/gorm"
)

const (
	ScopeAuth          = "auth"
	ScopeResetPassword = "reset_password"
	TokenLength        = 32
)

func GenerateToken(userID uint, ttl time.Duration, scope string) (*models.Token, string, error) {
	b := make([]byte, TokenLength)
	if _, err := rand.Read(b); err != nil {
		return nil, "", fmt.Errorf("could not generate random bytes: %w", err)
	}

	plaintext := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(b)
	hash := sha256.Sum256([]byte(plaintext))

	token := &models.Token{
		UserID: userID,
		Hash:   hash[:],
		Scope:  scope,
		Expiry: time.Now().UTC().Add(ttl),
	}

	return token, plaintext, nil
}

func VerifyToken(provided string, dbToken *models.Token) bool {
	hash := sha256.Sum256([]byte(provided))
	return subtle.ConstantTimeCompare(hash[:], dbToken.Hash) == 1 && dbToken.Expiry.After(time.Now().UTC())
}

func CleanupExpiredTokens(db *gorm.DB) error {
	return db.Where("expiry < ?", time.Now().UTC()).Delete(&models.Token{}).Error
}
