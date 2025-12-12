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
	r.Post("/register", app.UserHandler.Register)
	r.Post("/login", app.TokenHandler.Login)

	// API routes
	r.Group(func(r chi.Router) {
		r.Use(app.Middleware.Authenticate)
		r.Route("/users", func(r chi.Router) {
		})
		r.Route("/categories", func(r chi.Router) {
			r.Get("/", app.CategoryHandler.GetCategories)
			r.Get("/{id}", app.CategoryHandler.GetCategoryByID)
			r.Post("/", app.CategoryHandler.CreateCategory)
			r.Put("/{id}", app.CategoryHandler.UpdateCategory)
			r.Delete("/{id}", app.CategoryHandler.DeleteCategory)
		})
	})

	return r
}
