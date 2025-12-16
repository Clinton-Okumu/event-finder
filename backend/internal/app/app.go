package app

import (
	"backend/internal/api"
	"backend/internal/config"
	"backend/internal/middleware"
	"backend/internal/store"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Application struct {
	DB              *gorm.DB
	Logger          *log.Logger
	UserHandler     *api.UserHandler
	CategoryHandler *api.CategoryHandler
	TokenHandler    *api.TokenHandler
	EventHandler    *api.EventsHandler
	Middleware      middleware.UserMiddleware
}

func NewApplication() (*Application, error) {

	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	db, err := config.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := config.RunMigrations(db); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	logger := log.New(os.Stdout, "[Event Finder] ", log.Ldate|log.Ltime)

	//stores
	userStore := store.NewUserStore(db)
	categoryStore := store.NewCategoryStore(db)
	tokenStore := store.NewTokenStore(db)
	eventStore := store.NewEventStore(db)

	//handlers
	userHandler := api.NewUserHandler(userStore, logger)
	categoryHandler := api.NewCategoryHandler(categoryStore, logger)
	tokenHandler := api.NewTokenHandler(tokenStore, userStore, logger)
	eventHandler := api.NewEventsHandler(eventStore, logger)
	middlewareHandler := middleware.UserMiddleware{UserStore: userStore, TokenStore: tokenStore}

	app := &Application{
		DB:              db,
		Logger:          logger,
		UserHandler:     userHandler,
		CategoryHandler: categoryHandler,
		TokenHandler:    tokenHandler,
		EventHandler:    eventHandler,
		Middleware:      middlewareHandler,
	}
	return app, nil
}

func (a *Application) HealthChecker(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available\n")
}

func (a *Application) Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the event-finder API")
}
