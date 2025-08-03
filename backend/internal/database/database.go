package database

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/freddyamarante/go-svelte-finance-tracker/internal/models"
)

var DB *gorm.DB

// InitDatabase initializes the database connection
func InitDatabase(dbPath string, logSQL bool) error {
	var gormConfig gorm.Config
	if !logSQL {
		gormConfig.Logger = logger.Default.LogMode(logger.Silent)
	}

	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gormConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connected successfully")

	// Auto migrate the schema
	err = DB.AutoMigrate(&models.User{}, &models.Transaction{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("Database migration completed")
	return nil
}

// generateID generates a random ID
func generateID() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

// InitSampleData initializes sample data if the database is empty
func InitSampleData() error {
	var count int64
	DB.Model(&models.Transaction{}).Count(&count)

	if count == 0 {
		// Hash the sample password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("demo123"), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("failed to hash sample password: %w", err)
		}

		// Create a sample user with proper password hash
		sampleUser := models.User{
			ID:       generateID(),
			Email:    "demo@example.com",
			Password: string(hashedPassword),
		}

		if err := DB.Create(&sampleUser).Error; err != nil {
			return fmt.Errorf("failed to create sample user: %w", err)
		}

		sampleTransactions := []models.Transaction{
			{ID: generateID(), UserID: sampleUser.ID, Description: "Initial Savings", Amount: 1000.0, Type: "income"},
			{ID: generateID(), UserID: sampleUser.ID, Description: "Groceries", Amount: 50.0, Type: "expense"},
			{ID: generateID(), UserID: sampleUser.ID, Description: "Salary", Amount: 2500.0, Type: "income"},
			{ID: generateID(), UserID: sampleUser.ID, Description: "Gas", Amount: 45.0, Type: "expense"},
			{ID: generateID(), UserID: sampleUser.ID, Description: "Freelance Work", Amount: 300.0, Type: "income"},
		}

		for _, t := range sampleTransactions {
			if err := DB.Create(&t).Error; err != nil {
				return fmt.Errorf("failed to create sample transaction: %w", err)
			}
		}

		log.Println("Sample data initialized with user: demo@example.com / demo123")
	}

	return nil
}
