package middleware

import (
	"backend/internal/models"
	"backend/internal/store"
	"backend/internal/tokens"
	"backend/internal/utils"
	"context"
	"net/http"
	"strings"
)

type UserMiddleware struct {
	UserStore  store.UserStore
	TokenStore store.TokenStore
}

type contextKey string

const userContextKey = contextKey("user")

func SetUser(r *http.Request, user *models.User) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, user)
	return r.WithContext(ctx)
}

func GetUser(r *http.Request) *models.User {
	user, ok := r.Context().Value(userContextKey).(*models.User)
	if !ok {
		panic("missing user in request context")
	}
	return user
}

func (um *UserMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Authorization")

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			r = SetUser(r, models.AnonymousUser)
			next.ServeHTTP(w, r)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"error": "Invalid authorization header"})
			return
		}

		token := parts[1]
		user, err := um.TokenStore.GetUserByToken(r.Context(), token, tokens.ScopeAuth)
		if err != nil || user == nil {
			utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"error": "Invalid or expired token"})
			return
		}

		r = SetUser(r, user)
		next.ServeHTTP(w, r)
	})
}

func (um *UserMiddleware) RequireUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := GetUser(r)
		if user.IsAnonymous() {
			utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"error": "You must be logged in"})
			return
		}
		next.ServeHTTP(w, r)
	}
}
