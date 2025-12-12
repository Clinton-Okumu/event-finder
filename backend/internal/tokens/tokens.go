package tokens

import (
	"backend/internal/models"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"
)

const (
	ScopeAuth          = "auth"
	ScopeResetPassword = "reset_password"
	TokenLength        = 32
)

func GenerateToken(userID uint, ttl time.Duration, scope string) (*models.Token, string, error) {
	b := make([]byte, TokenLength)
	_, err := rand.Read(b)
	if err != nil {
		return nil, "", fmt.Errorf("could not generate random bytes: %w", err)
	}
	plaintext := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(b)

	hash := sha256.Sum256([]byte(plaintext))

	token := &models.Token{
		UserID: userID,
		Hash:   hash[:],
		Scope:  scope,
		Expiry: time.Now().Add(ttl),
	}

	return token, plaintext, nil
}

func HashToken(token string) []byte {
	hash := sha256.Sum256([]byte(token))
	return hash[:]
}
