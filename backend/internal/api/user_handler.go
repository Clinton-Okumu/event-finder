package api

import (
	"backend/internal/models"
	"backend/internal/store"
	"backend/internal/utils"
	"encoding/json"
	"log"
	"net/http"
	"strings"
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

// --- Login ---
type loginRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input loginRequest
	if err := h.parseJSON(r, &input); err != nil {
		h.respondError(w, http.StatusBadRequest, "invalid request format")
		return
	}

	input.Email = strings.TrimSpace(input.Email)
	input.Password = strings.TrimSpace(input.Password)

	if input.Email == "" || input.Password == "" {
		h.respondError(w, http.StatusBadRequest, "email and password are required")
		return
	}

	user, err := h.userStore.GetUserByEmail(r.Context(), input.Email)
	if err != nil {
		h.logger.Printf("ERROR: retrieving user: %v", err)
		h.respondError(w, http.StatusUnauthorized, "invalid email or password")
		return
	}

	if !user.CheckPassword(input.Password) {
		h.respondError(w, http.StatusUnauthorized, "invalid email or password")
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{
		"user": map[string]any{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"role":     user.Role,
		},
	})
}
