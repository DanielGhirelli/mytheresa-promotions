package main

import (
	"log"
	"mytheresa-promotions/config"
	"mytheresa-promotions/http"
	"mytheresa-promotions/pkg/application"
	"mytheresa-promotions/pkg/repo"
	"mytheresa-promotions/pkg/service"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or failed to load, using default values")
	}

	// Load configuration
	cfg := config.LoadConfig()
	if err := cfg.Validate(); err != nil {
		log.Fatalf("Invalid configuration: %v", err)
	}

	// Initialize dependencies
	productRepo, err := repo.NewFileRepository(cfg.DataFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}

	productService := service.NewService(productRepo)
	productHandler := application.NewHandler(productService)

	// Setup server
	server := http.NewServer(cfg.ApiKey)
	server.RegisterRoutes(productHandler)

	// Start server
	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := server.Start(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
