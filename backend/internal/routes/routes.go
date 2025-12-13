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
		// All API routes require authentication
		r.Use(app.Middleware.Authenticate)

		// Users
		r.Route("/users", func(r chi.Router) {
			// Add user-specific routes here if needed
		})

		// Categories
		r.Route("/categories", func(r chi.Router) {
			// Public routes
			r.Get("/", app.CategoryHandler.GetCategories)
			r.Get("/{id}", app.CategoryHandler.GetCategoryByID)

			// Admin-only routes for categories
			r.Group(func(r chi.Router) {
				r.Use(app.Middleware.RequireAdmin)
				r.Post("/", app.CategoryHandler.CreateCategory)
				r.Put("/{id}", app.CategoryHandler.UpdateCategory)
				r.Delete("/{id}", app.CategoryHandler.DeleteCategory)
			})
		})

		// Events
		r.Route("/events", func(r chi.Router) {
			// All event routes require authentication
			r.Get("/", app.EventHandler.GetEvents)
			r.Get("/{id}", app.EventHandler.GetEventByID)

			// Admin-only routes for events
			r.Group(func(r chi.Router) {
				r.Use(app.Middleware.RequireAdmin)
				r.Post("/", app.EventHandler.CreateEvent)
				r.Put("/{id}", app.EventHandler.UpdateEvent)
				r.Delete("/{id}", app.EventHandler.DeleteEvent)
			})
		})
	})

	return r
}
