package api

import (
	"backend/internal/models"
	"backend/internal/store"
	"backend/internal/tokens"
	"backend/internal/utils"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

type UserHandler struct {
	userStore store.UserStore
	logger    *log.Logger
}

func NewUserHandler(store store.UserStore, logger *log.Logger) *UserHandler {
	return &UserHandler{
		userStore: store,
		logger:    logger,
	}
}

func (h *UserHandler) respondError(w http.ResponseWriter, status int, msg string) {
	utils.WriteJSON(w, status, utils.Envelope{"error": msg})
}

func (h *UserHandler) parseJSON(r *http.Request, dst any) error {
	return json.NewDecoder(r.Body).Decode(dst)
}

// --- Register ---
type registerRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input registerRequest
	if err := h.parseJSON(r, &input); err != nil {
		h.respondError(w, http.StatusBadRequest, "invalid request format")
		return
	}

	input.Username = strings.TrimSpace(input.Username)
	input.Email = strings.TrimSpace(input.Email)
	input.Password = strings.TrimSpace(input.Password)

	if input.Username == "" || input.Email == "" || len(input.Password) < 6 {
		h.respondError(w, http.StatusBadRequest, "name, email, and password (min 6 chars) are required")
		return
	}

	user := &models.User{
		Username: input.Username,
		Email:    input.Email,
		Role:     "user",
	}

	if err := user.SetPassword(input.Password); err != nil {
		h.logger.Printf("ERROR: hashing password: %v", err)
		h.respondError(w, http.StatusInternalServerError, "internal error")
		return
	}

	if err := h.userStore.CreateUser(r.Context(), user); err != nil {
		h.logger.Printf("ERROR: creating user: %v", err)
		h.respondError(w, http.StatusInternalServerError, "could not create user")
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{
		"user": map[string]any{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"role":     user.Role,
		},
	})
}

type TokenHandler struct {
	tokenStore store.TokenStore
	userStore  store.UserStore
	logger     *log.Logger
}

type createTokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewTokenHandler(tokenStore store.TokenStore, userStore store.UserStore, logger *log.Logger) *TokenHandler {
	return &TokenHandler{
		tokenStore: tokenStore,
		userStore:  userStore,
		logger:     logger,
	}
}

func (h *TokenHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req createTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Printf("ERROR decoding createTokenRequest: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "Invalid request body"})
		return
	}

	user, err := h.userStore.GetUserByEmail(r.Context(), req.Email)
	if err != nil {
		h.logger.Printf("ERROR finding user: %v", err)
		utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"error": "Invalid credentials"})
		return
	}

	tokenModel, plaintext, err := h.tokenStore.CreateNewToken(r.Context(), user.ID, 24*time.Hour, tokens.ScopeAuth)
	if err != nil {
		h.logger.Printf("ERROR creating token: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "Failed to generate token"})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{
		"token":  plaintext,
		"expiry": tokenModel.Expiry,
		"user": map[string]any{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"role":     user.Role,
		},
	})
}
