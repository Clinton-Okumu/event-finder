package admin

import (
	"backend/internal/config"
	"backend/internal/models"
	"backend/internal/store"
	"context"
	"fmt"
	"log"
	"time"
)

func SeedData() {
	db, err := config.Open()
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	categoryStore := store.NewCategoryStore(db)
	eventStore := store.NewEventStore(db)
	userStore := store.NewUserStore(db)

	ctx := context.Background()

	adminUser, err := userStore.GetUserByEmail(ctx, "admin@eventfinder.com")
	if err != nil || adminUser == nil {
		adminUser = &models.User{
			Username: "admin",
			Email:    "admin@eventfinder.com",
			Role:     "admin",
		}
		if err := adminUser.SetPassword("admin123"); err != nil {
			log.Fatalf("❌ failed to set admin password: %v", err)
		}
		if err := userStore.CreateUser(ctx, adminUser); err != nil {
			log.Fatalf("❌ Failed to create admin user: %v", err)
		}
		fmt.Println("✅ Created admin user: admin@eventfinder.com / admin123")
	}

	categories := []models.Category{
		{
			Name:        "Music",
			Description: "Live music events, concerts, and festivals",
		},
		{
			Name:        "Sports",
			Description: "Sports games, tournaments, and fitness events",
		},
		{
			Name:        "Technology",
			Description: "Tech conferences, hackathons, and workshops",
		},
		{
			Name:        "Arts & Culture",
			Description: "Art exhibitions, theater performances, and cultural events",
		},
		{
			Name:        "Food & Drink",
			Description: "Food festivals, wine tastings, and culinary experiences",
		},
	}

	var categoryIDs []uint
	for _, cat := range categories {
		if err := categoryStore.CreateCategory(ctx, &cat); err != nil {
			log.Printf("⚠️  Warning: failed to create category '%s': %v", cat.Name, err)
		} else {
			categoryIDs = append(categoryIDs, cat.ID)
			fmt.Printf("✅ Created category: %s\n", cat.Name)
		}
	}

	if len(categoryIDs) == 0 {
		log.Fatal("❌ Failed to create any categories")
	}

	now := time.Now()
	events := []models.Event{
		{
			Title:            "Summer Music Festival 2025",
			Description:      "Three days of live music featuring top artists from around the world. Food vendors, art installations, and camping available.",
			Location:         "Central Park, New York",
			StartTime:        now.AddDate(0, 2, 0),
			EndTime:          now.AddDate(0, 2, 3),
			CreatedByID:      adminUser.ID,
			Status:           models.EventConfirmed,
			ImageURL:         "https://images.unsplash.com/photo-1459749411177-287ce3278023?w=800",
			TotalTickets:     5000,
			TicketsRemaining: 5000,
			CategoryID:       categoryIDs[0],
			Price:            99.00,
		},
		{
			Title:            "Tech Innovators Conference",
			Description:      "Annual technology conference bringing together innovators, entrepreneurs, and investors. Keynotes, workshops, and networking events.",
			Location:         "Moscone Center, San Francisco",
			StartTime:        now.AddDate(0, 1, 15),
			EndTime:          now.AddDate(0, 1, 17),
			CreatedByID:      adminUser.ID,
			Status:           models.EventConfirmed,
			ImageURL:         "https://images.unsplash.com/photo-1540575467063-178a50c2df87?w=800",
			TotalTickets:     2000,
			TicketsRemaining: 1800,
			CategoryID:       categoryIDs[2],
			Price:            299.00,
		},
		{
			Title:            "NBA Finals Game 1",
			Description:      "The championship series begins! Watch the two best teams compete for the title. Premium seating and VIP packages available.",
			Location:         "Madison Square Garden, New York",
			StartTime:        now.AddDate(0, 0, 10),
			EndTime:          now.AddDate(0, 0, 10).Add(3 * time.Hour),
			CreatedByID:      adminUser.ID,
			Status:           models.EventConfirmed,
			ImageURL:         "https://images.unsplash.com/photo-1546519638-68e109498ffc?w=800",
			TotalTickets:     20000,
			TicketsRemaining: 8500,
			CategoryID:       categoryIDs[1],
			Price:            150.00,
		},
		{
			Title:            "Modern Art Exhibition",
			Description:      "A stunning collection of contemporary art featuring works from emerging and established artists. Guided tours available daily.",
			Location:         "Museum of Modern Art, New York",
			StartTime:        now.AddDate(0, 0, 5),
			EndTime:          now.AddDate(0, 0, 5).Add(5 * time.Hour),
			CreatedByID:      adminUser.ID,
			Status:           models.EventConfirmed,
			ImageURL:         "https://images.unsplash.com/photo-1531243269054-5ebf6f34081e?w=800",
			TotalTickets:     500,
			TicketsRemaining: 320,
			CategoryID:       categoryIDs[3],
			Price:            25.00,
		},
		{
			Title:            "International Food Festival",
			Description:      "Taste cuisines from over 50 countries! Live cooking demonstrations, wine tastings, and family-friendly activities all day.",
			Location:         "Grant Park, Chicago",
			StartTime:        now.AddDate(0, 1, 0),
			EndTime:          now.AddDate(0, 1, 0).Add(8 * time.Hour),
			CreatedByID:      adminUser.ID,
			Status:           models.EventPending,
			ImageURL:         "https://images.unsplash.com/photo-1555939594-58d7cb561ad1?w=800",
			TotalTickets:     10000,
			TicketsRemaining: 10000,
			CategoryID:       categoryIDs[4],
			Price:            45.00,
		},
		{
			Title:            "Jazz Night Live",
			Description:      "An intimate evening of smooth jazz featuring local and international musicians. Full bar and light dinner menu available.",
			Location:         "Blue Note, New York",
			StartTime:        now.AddDate(0, 0, 3).Add(20 * time.Hour),
			EndTime:          now.AddDate(0, 0, 3).Add(23*time.Hour + 30*time.Minute),
			CreatedByID:      adminUser.ID,
			Status:           models.EventConfirmed,
			ImageURL:         "https://images.unsplash.com/photo-1514320291840-2e0a9bf2a9ae?w=800",
			TotalTickets:     300,
			TicketsRemaining: 250,
			CategoryID:       categoryIDs[0],
			Price:            75.00,
		},
		{
			Title:            "Marathon 2025",
			Description:      "Join thousands of runners for the annual city marathon. Multiple distance options available. Registration includes medal and t-shirt.",
			Location:         "Starting Line, Boston",
			StartTime:        now.AddDate(0, 3, 0).Add(6 * time.Hour),
			EndTime:          now.AddDate(0, 3, 0).Add(12 * time.Hour),
			CreatedByID:      adminUser.ID,
			Status:           models.EventConfirmed,
			ImageURL:         "https://images.unsplash.com/photo-1452626038306-9aae5e071dd3?w=800",
			TotalTickets:     30000,
			TicketsRemaining: 22000,
			CategoryID:       categoryIDs[1],
			Price:            50.00,
		},
		{
			Title:            "AI & Machine Learning Summit",
			Description:      "Explore the latest advancements in artificial intelligence. Expert speakers, hands-on workshops, and networking with industry leaders.",
			Location:         "Convention Center, Seattle",
			StartTime:        now.AddDate(0, 2, 15),
			EndTime:          now.AddDate(0, 2, 17),
			CreatedByID:      adminUser.ID,
			Status:           models.EventConfirmed,
			ImageURL:         "https://images.unsplash.com/photo-1485827404703-89b55fcc595e?w=800",
			TotalTickets:     1500,
			TicketsRemaining: 900,
			CategoryID:       categoryIDs[2],
			Price:            399.00,
		},
	}

	for _, event := range events {
		if err := eventStore.CreateEvent(ctx, &event); err != nil {
			log.Printf("⚠️  Warning: failed to create event '%s': %v", event.Title, err)
		} else {
			fmt.Printf("✅ Created event: %s\n", event.Title)
		}
	}

	fmt.Println("\n🎉 Sample data seeded successfully!")
}
