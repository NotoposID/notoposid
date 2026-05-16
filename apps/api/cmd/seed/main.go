package main

import (
	"log"

	"github.com/notopos/api/internal/config"
	"github.com/notopos/api/internal/database"
	"github.com/notopos/api/internal/modules/tenants"
	"github.com/notopos/api/internal/modules/users"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	cfg := config.LoadConfig()
	database.ConnectDB(cfg)
	db := database.DB

	log.Println("Running Migrations...")
	err := db.AutoMigrate(&tenants.Tenant{}, &users.User{})
	if err != nil {
		log.Fatal("Migration failed: ", err)
	}

	log.Println("Seeding Data...")

	// Create Default Tenant
	tenant := tenants.Tenant{
		Name:             "Default Tenant",
		Domain:           "notopos.com",
		SubscriptionPlan: "pro",
	}
	db.Where(tenants.Tenant{Domain: "notopos.com"}).FirstOrCreate(&tenant)

	// Create Admin User
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), 14)
	user := users.User{
		TenantID:     tenant.ID,
		Name:         "Admin User",
		Email:        "admin@notopos.com",
		PasswordHash: string(hashedPassword),
		Role:         "admin",
	}
	db.Where(users.User{Email: "admin@notopos.com"}).FirstOrCreate(&user)

	log.Println("Seeding completed successfully!")
	log.Printf("Default Tenant ID: %s", tenant.ID)
	log.Printf("Admin Email: %s", user.Email)
}
