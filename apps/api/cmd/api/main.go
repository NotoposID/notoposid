package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	_ "github.com/notopos/api/docs"
	"github.com/notopos/api/internal/config"
	"github.com/notopos/api/internal/database"
	"github.com/notopos/api/internal/middleware"
	"github.com/notopos/api/internal/modules/auth"
	"github.com/notopos/api/internal/modules/users"
)

// @title NOTOPOS AI API
// @version 1.0
// @description Enterprise-grade AI POS API
// @contact.name API Support
// @host localhost:8000
// @BasePath /api/v1
func main() {
	// Load config
	cfg := config.LoadConfig()

	// Initialize Database
	database.ConnectDB(cfg)

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		AppName: "NOTOPOS AI API",
	})

	// Global Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	// Swagger Route
	app.Get("/swagger/*", swagger.HandlerDefault)

	// API Groups
	api := app.Group("/api/v1")

	// Auth Routes
	authHandler := auth.NewHandler(cfg)
	authGroup := api.Group("/auth")
	authGroup.Post("/login", authHandler.Login)

	// Protected Routes (Example)
	protected := api.Group("/", middleware.TenantMiddleware())
	protected.Get("/me", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"tenant_id": c.Locals("tenantID")})
	})

	// Users Routes
	userRepo := users.NewRepository(database.DB)
	userService := users.NewService()
	userHandler := users.NewHandler(userRepo, userService)
	
	usersGroup := protected.Group("/users")
	usersGroup.Post("/", userHandler.Create)
	usersGroup.Get("/", userHandler.GetAll)
	usersGroup.Get("/:id", userHandler.GetByID)
	usersGroup.Put("/:id", userHandler.Update)
	usersGroup.Delete("/:id", userHandler.Delete)

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "healthy",
			"app":    "NOTOPOS AI",
			"env":    cfg.AppEnv,
		})
	})

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
