package app

import (
	"backend/internal/api"
	"backend/internal/config"
	"backend/internal/store"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Application struct {
	DB          *gorm.DB
	Logger      *log.Logger
	UserHandler *api.UserHandler
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

	//handlers
	userHandler := api.NewUserHandler(userStore, logger)

	app := &Application{
		DB:          db,
		Logger:      logger,
		UserHandler: userHandler,
	}
	return app, nil
}

func (a *Application) HealthChecker(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available\n")
}

func (a *Application) Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the event-finder API")
}
