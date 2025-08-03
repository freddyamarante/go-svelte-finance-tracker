package database

import (
	"fmt"
	"log"

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
	err = DB.AutoMigrate(&models.Transaction{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("Database migration completed")
	return nil
}

// InitSampleData initializes sample data if the database is empty
func InitSampleData() error {
	var count int64
	DB.Model(&models.Transaction{}).Count(&count)

	if count == 0 {
		sampleTransactions := []models.Transaction{
			{ID: "1", Description: "Initial Savings", Amount: 1000.0, Type: "income"},
			{ID: "2", Description: "Groceries", Amount: 50.0, Type: "expense"},
			{ID: "3", Description: "Salary", Amount: 2500.0, Type: "income"},
			{ID: "4", Description: "Gas", Amount: 45.0, Type: "expense"},
			{ID: "5", Description: "Freelance Work", Amount: 300.0, Type: "income"},
		}

		for _, t := range sampleTransactions {
			if err := DB.Create(&t).Error; err != nil {
				return fmt.Errorf("failed to create sample transaction: %w", err)
			}
		}

		log.Println("Sample data initialized")
	}

	return nil
} 