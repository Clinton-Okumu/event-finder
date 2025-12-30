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
	r.Get("/validate-token", app.TokenHandler.ValidateToken)
	r.Post("/logout", app.TokenHandler.Logout)

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

		//Event tickets
		r.Route("/event_tickets", func(r chi.Router) {
			// All event ticket routes require authentication
			r.Get("/", app.EventTicketHandler.GetEventTickets)
			r.Get("/{id}", app.EventTicketHandler.GetEventTicketByID)

			// Admin-only routes for event tickets
			r.Group(func(r chi.Router) {
				r.Use(app.Middleware.RequireAdmin)
				r.Post("/", app.EventTicketHandler.CreateEventTicket)
				r.Put("/{id}", app.EventTicketHandler.UpdateEventTicket)
				r.Delete("/{id}", app.EventTicketHandler.DeleteEventTicket)
			})
		})

		//Bookings
		r.Route("/bookings", func(r chi.Router) {
			// All booking routes require authentication
			r.Get("/", app.BookingHandler.GetBookings)
			r.Get("/{id}", app.BookingHandler.GetBookingByID)

			// Admin-only routes for bookings
			r.Group(func(r chi.Router) {
				r.Use(app.Middleware.RequireAdmin)
				r.Post("/", app.BookingHandler.CreateBooking)
				r.Put("/{id}", app.BookingHandler.UpdateBooking)
				r.Delete("/{id}", app.BookingHandler.DeleteBooking)
			})
		})

		// Booking items
		r.Route("/booking_items", func(r chi.Router) {
			// All booking item routes require authentication
			r.Get("/", app.BookingItemHandler.GetBookingItems)
			r.Get("/{id}", app.BookingItemHandler.GetBookingItemByID)

			// Admin-only routes for booking items
			r.Group(func(r chi.Router) {
				r.Use(app.Middleware.RequireAdmin)
				r.Post("/", app.BookingItemHandler.CreateBookingItem)
				r.Put("/{id}", app.BookingItemHandler.UpdateBookingItem)
				r.Delete("/{id}", app.BookingItemHandler.DeleteBookingItem)
			})
		})
	})

	return r
}
