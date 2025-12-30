package main

import (
	"backend/cmd/admin"
	"backend/internal/app"
	"backend/internal/routes"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {

	// cli commands(admin)
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "create-admin":
			admin.CreateAdmin()
			return
		case "seed-data":
			admin.SeedData()
			return
		}
	}

	var port int
	serverPortStr := os.Getenv("PORT")
	port, err := strconv.Atoi(serverPortStr)
	if err != nil || port <= 0 {
		port = 5000
	}

	// initialize the Application struct
	application, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	// Close the underlying *sql.DB used by GORM
	sqlDB, err := application.DB.DB()
	if err != nil {
		application.Logger.Fatalf("failed to get raw DB: %v", err)
	}
	defer sqlDB.Close()

	// setup routes
	r := routes.SetUpRoutes(application)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	application.Logger.Printf("🚀 Starting server on http://localhost:%d", port)
	if err := server.ListenAndServe(); err != nil {
		application.Logger.Fatal(err)
	}
}
