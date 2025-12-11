package routes

import (
	"backend/internal/app"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func SetUpRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	// CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// System routes
	r.Get("/", app.Welcome)
	r.Get("/health", app.HealthChecker)

	// API routes
	r.Route("/users", func(r chi.Router) {
		r.Post("/register", app.UserHandler.Register)
		r.Post("/login", app.UserHandler.Login)
	})

	return r
}
