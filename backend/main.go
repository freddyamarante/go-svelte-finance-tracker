package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/freddyamarante/go-svelte-finance-tracker/internal/config"
	"github.com/freddyamarante/go-svelte-finance-tracker/internal/database"
	"github.com/freddyamarante/go-svelte-finance-tracker/internal/server"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Print configuration for debugging
	log.Printf("Starting backend with configuration:")
	log.Printf("  Environment: %s", cfg.Env)
	log.Printf("  Port: %s", cfg.Server.Port)
	log.Printf("  Host: %s", cfg.Server.Host)
	log.Printf("  Database Path: %s", cfg.Database.Path)
	log.Printf("  Log Requests: %t", cfg.Logging.Requests)
	log.Printf("  Log SQL: %t", cfg.Logging.SQL)
	log.Printf("  Log CORS: %t", cfg.Logging.CORS)

	// Ensure database directory exists
	dbDir := filepath.Dir(cfg.Database.Path)
	if dbDir != "." {
		log.Printf("Creating database directory: %s", dbDir)
		if err := os.MkdirAll(dbDir, 0755); err != nil {
			log.Printf("Failed to create database directory: %v", err)
		}
	}

	// Initialize database
	log.Printf("Connecting to database: %s", cfg.Database.Path)
	if err := database.InitDatabase(cfg.Database.Path, cfg.Logging.SQL); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize sample data
	if err := database.InitSampleData(); err != nil {
		log.Fatalf("Failed to initialize sample data: %v", err)
	}

	// Create and start server
	srv := server.New(cfg)
	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
