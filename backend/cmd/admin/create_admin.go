package admin

import (
	"backend/internal/config"
	"backend/internal/models"
	"backend/internal/store"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/term"
)

func CreateAdmin() {
	db, err := config.Open()
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	userStore := store.NewUserStore(db)
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Enter email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Enter password: ")
	bytePassword, _ := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	password := strings.TrimSpace(string(bytePassword))

	if username == "" || email == "" || password == "" {
		log.Fatal("❌ username, email, and password are required")
	}

	admin := &models.User{
		Username: username,
		Email:    email,
		Role:     "admin",
	}

	if err := admin.SetPassword(password); err != nil {
		log.Fatalf("❌ failed to set password: %v", err)
	}

	if err := userStore.CreateUser(context.Background(), admin); err != nil {
		log.Fatalf("❌ Failed to create admin user: %v", err)
	}

	fmt.Printf("✅ Admin user '%s' created successfully\n", email)

}
